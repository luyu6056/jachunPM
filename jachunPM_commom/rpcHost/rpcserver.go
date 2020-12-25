package rpcHost

import (
	"errors"
	"io"
	"jachunPM_commom/db"
	"libraries"
	"protocol"
	"sync"
	"sync/atomic"
	"time"

	"github.com/klauspost/compress/zstd"
	"github.com/luyu6056/gnet"
)

const (
	rpcStatusOpen     = iota //熔断状态正常
	rpcStatusHalfOpen        //半开
	rpcStatusClose
)

/**
 *只有网关host才允许作为rpc server服务
 *网关中心————发消息流程
 *GET一个消息体并赋值
 *从rpcMsgChan池子里拿出一个buffer
 *消息体.Write(buffer)
 *消息体.Put()
 *buffer发送到指定的chan
 **/
var (
	rpcServerIdList     = make(map[uint8][]*RpcServer)
	rpcServerOutChan    = make(map[uint8]chan *libraries.MsgBuffer) //公共消息chan，实现均衡负载
	rpcLock             sync.RWMutex
	rpcPingTime                = time.Second * 10
	rpcPingHalfOpenTime        = time.Second * 30 //ping响应超时
	rpcServerCenterId          = make(map[uint8]uint8)
	globalMsgno         uint32 = 1
)

func init() {

	for i := uint8(0); i < protocol.MaxServerNoNum; i++ {
		rpcServerOutChan[i] = make(chan *libraries.MsgBuffer, protocol.Rpcmsgnum)
	}
}

func MsgnoInit() {
	var msg db.Log_msg
	err := db.DB.Table(db.TABLE_LOG_MSG).Field("max(msgno) as Msgno").Find(&msg)
	if err != nil {
		libraries.ReleaseLog("初始化msgno号失败%v", err)
	} else {
		globalMsgno = msg.Msgno + 1
	}
}

type RpcServer struct {
	ServerNo                       uint8 //服务序号
	Id                             int16 //服务Id，有效值0-255
	ServerConn                     gnet.Conn
	zstdDecoder                    *zstd.Decoder        //codec解压相关
	zstdDecodeBuf1, zstdDecodeBuf2 *libraries.MsgBuffer //codec解压相关

	setStatusOpenChan   chan string
	closeChan           chan int
	outChan             chan *libraries.MsgBuffer //指定本服务接收的消息
	startTime, busyTime int64                     //时间统计
	ErrNum              uint32
	Ip                  string
	pongTime            int64
	status              int
	isCenter            bool   //是否中心服
	window              int32  //发送窗口
	local               uint16 //ServerNo与Id编码后的数值
	CacheServer         *RpcServer
}

func NewRpcServer(c gnet.Conn) *RpcServer {
	s := &RpcServer{ServerConn: c, Id: -1}
	return s
}
func (svr *RpcServer) SendMsg(remote uint16, msgno uint32, ttl uint8, out protocol.MSG_DATA) {
	protocol.SendMsg(svr.local, remote, msgno, ttl, out, svr.outChan)
}
func (svr *RpcServer) SendMsgWaitResult(remote uint16, msgno uint32, ttl uint8, out protocol.MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	return protocol.SendMsgWaitResult(0, remote, msgno, ttl, out, result, rpcServerOutChan[protocol.HostServerNo], timeout...)
}
func (svr *RpcServer) Start(no uint8, ipport string, window int32) {
	svr.ServerNo = no
	svr.Ip = ipport
	svr.zstdDecodeBuf1 = &libraries.MsgBuffer{}
	svr.zstdDecodeBuf2 = &libraries.MsgBuffer{}
	svr.zstdDecoder, _ = zstd.NewReader(svr.zstdDecodeBuf1)
	svr.outChan = make(chan *libraries.MsgBuffer, protocol.Rpcmsgnum)
	svr.closeChan = make(chan int, 1)
	svr.setStatusOpenChan = make(chan string)
	svr.pongTime = time.Now().Unix()
	svr.window = window
	rpcLock.Lock()
	defer rpcLock.Unlock()
	outChan := rpcServerOutChan[svr.ServerNo]
	//分配id
	if rpcServerIdList[svr.ServerNo] == nil {
		rpcServerIdList[svr.ServerNo] = make([]*RpcServer, 256)
	}
	for id, v := range rpcServerIdList[svr.ServerNo] {
		if v == nil && id < 256 {
			svr.Id = int16(id)
			rpcServerIdList[svr.ServerNo][id] = svr
			break
		}
	}
	if svr.Id < 0 || svr.Id > 255 {
		svr.ServerConn.Close()
		libraries.ReleaseLog("服务注册失败,服务No%d,获得的ID%d", no, svr.Id)
		return
	}
	svr.local = uint16(svr.ServerNo) | uint16(svr.Id<<8)
	//分配center
	if _, ok := rpcServerCenterId[svr.ServerNo]; !ok {
		svr.setCenter()
	}
	go svr.handleMsgOut(outChan)

	libraries.DebugLog("服务%v，ID%v,启动", svr.ServerNo, svr.Id)

	data := protocol.GET_MSG_COMMON_regServer_result()
	data.Id = uint8(svr.Id)
	svr.SendMsg(svr.local, 0, 0, data)
	data.Put()
}

func (svr *RpcServer) Close() {
	if svr.closeChan != nil {
		select {
		case svr.closeChan <- 1:
		default:
		}
	}
	rpcLock.Lock()
	defer rpcLock.Unlock()
	if svr.Id > -1 && rpcServerIdList[svr.ServerNo] != nil {
		libraries.DebugLog("服务%v，ID%v，关闭", svr.ServerNo, svr.Id)
		rpcServerIdList[svr.ServerNo][svr.Id] = nil
	}
	if svr.isCenter {
		delete(rpcServerCenterId, svr.ServerNo)
		//寻找一个正常的服务并将它设置为中心服务
		for _, v := range rpcServerIdList[svr.ServerNo] {
			if v != nil {
				v.setCenter()
				break
			}
		}
	}

}

func (svr *RpcServer) setCenter() {
	rpcServerCenterId[svr.ServerNo] = uint8(svr.Id)
	data := protocol.GET_MSG_COMMON_StartTicker()
	svr.isCenter = true
	svr.SendMsg(svr.local, 0, 0, data)
	data.Put()
}

//以下为server的消息in和out，client的解码编码在protocol里
func (svr *RpcServer) handleMsgOut(outChan chan *libraries.MsgBuffer) {
	defer func() {
		if r := recover(); r != nil {
			libraries.DebugLog("%v", r)
			go svr.handleMsgOut(outChan)
		}
	}()
	var (
		out      = &libraries.MsgBuffer{}
		msgbuf   = &libraries.MsgBuffer{}
		zstdbuf  = &libraries.MsgBuffer{}
		compress bool
		msglen   int
		//busyTime int64 = time.Now().UnixNano()
		msgNum int
	)
	zstdWriter, _ := zstd.NewWriter(zstdbuf)
	writeToBuf := func(o *libraries.MsgBuffer) {
		msglen += o.Len()
		if compress {
			zstdWriter.Write(o.Bytes())
		} else {
			if msgNum > protocol.CompressMinNum || out.Len()+o.Len() > protocol.CompressMinLen {
				compress = true
				zstdWriter.Reset(zstdbuf)
				zstdWriter.Write(out.Bytes())
				out.Reset()
				zstdWriter.Write(o.Bytes())
			} else {
				out.Write(o.Bytes())
			}
		}
	}
	write := func(o *libraries.MsgBuffer, c chan *libraries.MsgBuffer) {
		compress = false
		out.Reset()
		msgbuf.Reset()
		zstdbuf.Reset()
		msglen = 0
		writeToBuf(o)
		o.Reset()
		protocol.BufPoolPut(o)
		msgNum = 1
	out:
		for i := 0; i < len(c) && i < int(svr.window)-1 && i < protocol.MaxMsgNum; i++ {
			select {
			case o1 := <-c:
				if msglen+o1.Len() > protocol.MaxMsgLen {
					c <- o1
					break out
				}
				msgNum++
				writeToBuf(o1)
				o1.Reset()
				protocol.BufPoolPut(o1)
			default:
			}

		}

		atomic.AddInt32(&svr.window, -1*int32(msgNum))
		//窗口控制熔断
		if svr.window <= 0 {
			libraries.DebugLog("服务%v，ID%v，因窗口不够，进入半开状态", svr.ServerNo, svr.Id)
			svr.status = rpcStatusHalfOpen
		}
		if compress {
			//有压缩
			zstdWriter.Close()
			msglen = zstdbuf.Len()
			if msglen > protocol.MaxMsgLen {
				libraries.ReleaseLog("消息大于最大允许长度,压缩后更长？？？，请检查代码")
				return
			}
			b := msgbuf.Make(5 + msglen)
			b[0] = byte(msglen)
			b[1] = byte(msglen >> 8)
			b[2] = byte(msglen >> 16)
			b[3] = byte(msglen >> 24)
			b[4] = byte(msgNum) + 1<<7
			copy(b[5:], zstdbuf.Bytes())
		} else {
			//无压缩
			if msglen > protocol.MaxMsgLen {
				libraries.ReleaseLog("消息大于最大允许长度,前面的限制有错，请重写代码")
				return
			}
			if msglen != out.Len() {
				libraries.ReleaseLog("msglen计算有误，请检查代码")
			}
			b := msgbuf.Make(5 + msglen)
			b[0] = byte(msglen)
			b[1] = byte(msglen >> 8)
			b[2] = byte(msglen >> 16)
			b[3] = byte(msglen >> 24)
			b[4] = byte(msgNum)
			copy(b[5:], out.Bytes())
		}
		svr.ServerConn.AsyncWrite(msgbuf.Bytes())
	}
	ping := func(now time.Time) {
		checktime := now.Add(rpcPingHalfOpenTime * -1)
		//超时熔断
		if checktime.Unix() > svr.pongTime {
			libraries.DebugLog("服务%v，ID%v，cache%v,ping响应超时，进入半开状态", svr.ServerNo, svr.Id, svr.CacheServer)
			svr.status = rpcStatusHalfOpen
		}
		data := protocol.GET_MSG_COMMON_PING()
		svr.SendMsg(svr.local, 0, 0, data)
		data.Put()
	}
	for {

		switch svr.status {
		case rpcStatusOpen:
			select {
			case o := <-svr.outChan:
				write(o, svr.outChan)
			case o := <-outChan:
				write(o, outChan)
			case <-svr.closeChan:
				return
			case now := <-time.After(rpcPingTime):
				ping(now)
			}
		case rpcStatusHalfOpen:
			//半开状态，只处理指定了本服务的消息
			select {
			case o := <-svr.outChan:
				write(o, svr.outChan)
			case reason := <-svr.setStatusOpenChan:
				//由半开恢复到正常状态
				libraries.DebugLog("服务%v，ID%v，%s，进入正常状态", svr.ServerNo, svr.Id, reason)
				svr.status = rpcStatusOpen
			case <-svr.closeChan:
				return
			case now := <-time.After(rpcPingTime):
				ping(now)
			}
		case rpcStatusClose:
			return
		}

	}
}

type RpcCodec struct {
}

var errRpcContext = errors.New("错误的rpcContext")

func (code RpcCodec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	return buf, nil
}

func (code RpcCodec) Decode(c gnet.Conn) (data []byte, err error) {
	if c.BufferLength() > 5 {
		data = c.Read()
		msglen := int(data[0]) | int(data[1])<<8 | int(data[2])<<16 | int(data[3])<<24
		if len(data) < msglen+5 { //消息长度不够
			return nil, nil
		}
		c.ShiftN(msglen + 5)
		//解压缩
		if data[4]>>7 == 1 {
			if ctx, ok := c.Context().(*RpcServer); ok {
				if decoder := ctx.zstdDecoder; decoder != nil {
					ctx.zstdDecodeBuf1.Reset()
					ctx.zstdDecodeBuf1.Write(data[5 : msglen+5])
					ctx.zstdDecodeBuf2.Reset()
					ctx.zstdDecodeBuf2.WriteByte(data[4] - 128)
					decoder.Reset(ctx.zstdDecodeBuf1)
					io.Copy(ctx.zstdDecodeBuf2, decoder)
					return ctx.zstdDecodeBuf2.Bytes(), nil
				}
			}
			return nil, errRpcContext
		}
		return data[4 : msglen+5], nil
	}
	return nil, nil
}
