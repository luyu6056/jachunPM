package rpcHost

import (
	"io"
	"jachunPM_commom/db"
	"libraries"
	"os"
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
	rpcServerIdList[protocol.HostServerNo] = make([]*RpcServer, 1)
	rpcServerIdList[protocol.HostServerNo][0] = &RpcServer{ //common服务器
		ServerNo:          protocol.HostServerNo,
		Id:                0,
		ServerConn:        nil,
		zstdDecoder:       nil,
		zstdDecodeBuf1:    nil,
		zstdDecodeBuf2:    nil,
		setStatusOpenChan: nil,
		closeChan:         nil,
		outChan:           rpcServerOutChan[protocol.HostServerNo],
		startTime:         0,
		busyTime:          0,
		Ip:                "",
		pongTime:          0,
		status:            0,
		isCenter:          false,
		window:            0,
		local:             protocol.HostServerNo,
		CacheServer:       nil,
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
	inChan              chan *protocol.Msg
	startTime, busyTime int64 //时间统计
	Ip                  string
	pongTime            int64
	status              int
	isCenter            bool   //是否中心服
	window              int32  //发送窗口
	local               uint16 //ServerNo与Id编码后的数值
	CacheServer         *RpcServer
}

func NewRpcServer(c gnet.Conn) *RpcServer {
	s := &RpcServer{ServerConn: c, Id: -1, inChan: make(chan *protocol.Msg, 65535)}
	return s
}
func (svr *RpcServer) SendMsg(msg *protocol.Msg, remote uint16, out protocol.MSG_DATA) {
	protocol.SendMsg(msg, protocol.HostServerNo, remote, out, rpcServerOutChan[protocol.HostServerNo])
}
func (svr *RpcServer) SendMsgWaitResult(msg *protocol.Msg, remote uint16, out protocol.MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	return protocol.SendMsgWaitResult(msg, protocol.HostServerNo, remote, out, result, rpcServerOutChan[protocol.HostServerNo], timeout...)
}
func (svr *RpcServer) Start(no uint8, ipport string, window int32, in *protocol.Msg) {
	svr.ServerNo = no
	svr.Ip = ipport
	svr.zstdDecodeBuf1 = &libraries.MsgBuffer{}
	svr.zstdDecodeBuf2 = &libraries.MsgBuffer{}
	svr.zstdDecoder, _ = zstd.NewReader(svr.zstdDecodeBuf1,zstd.WithDecoderDicts(protocol.ZstdDict))
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
	go svr.handlerMsgOut(outChan)

	libraries.DebugLog("服务%v，ID%v,启动", svr.ServerNo, svr.Id)

	data := protocol.GET_MSG_HOST_regServer_result()
	data.Id = uint8(svr.Id)
	svr.SendMsg(in, svr.local, data)
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
func (svr *RpcServer) tick(now time.Time) {
	files, _ := libraries.ListDir(fileTmpPath, "")
	for _, f := range files {
		if s, _ := os.Stat(f); s != nil {
			if now.Unix()-s.ModTime().Unix() > 86400 { //清除一天前文件
				os.Remove(f)
			}
		}
	}
}
func (svr *RpcServer) setCenter() {
	rpcServerCenterId[svr.ServerNo] = uint8(svr.Id)
	data := protocol.GET_MSG_HOST_StartTicker()
	svr.isCenter = true
	svr.SendMsg(nil, svr.local, data)
	data.Put()
}

//以下为server的消息in和out，client的解码编码在protocol里
func (svr *RpcServer) handlerMsgOut(outChan chan *libraries.MsgBuffer) {
	defer func() {
		if r := recover(); r != nil {
			libraries.DebugLog("%v", r)
			go svr.handlerMsgOut(outChan)
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
	zstdWriter, _ := zstd.NewWriter(zstdbuf, zstd.WithEncoderLevel(protocol.ZstdLevel),zstd.WithEncoderDict(protocol.ZstdDict))
	writeToBuf := func(o *libraries.MsgBuffer) {
		msglen += o.Len()
		if compress {
			zstdWriter.Write(o.Bytes())
		} else {
			if msgNum > protocol.CompressMinNum || msglen > protocol.CompressMinLen {
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
				if msglen+o1.Len() > protocol.MaxOutLen {
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
			if msglen > protocol.MaxOutLen {
				libraries.ReleaseLog("消息大于最大允许长度,压缩后更长？？？")
				return
			}
			b := msgbuf.Make(4 + msglen)
			b[0] = byte(msglen)
			b[1] = byte(msglen >> 8)
			b[2] = byte(msglen >> 16)
			b[3] = byte(msglen>>24) + 1<<7
			copy(b[4:], zstdbuf.Bytes())
		} else {
			//无压缩
			b := msgbuf.Make(4 + msglen)
			b[0] = byte(msglen)
			b[1] = byte(msglen >> 8)
			b[2] = byte(msglen >> 16)
			b[3] = byte(msglen >> 24)
			copy(b[4:], out.Bytes())
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
		data := protocol.GET_MSG_HOST_PING()
		svr.SendMsg(nil, svr.local, data)
		data.Put()
		svr.tick(now)
	}
	pingTick := time.NewTicker(rpcPingTime)
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
			case now := <-pingTick.C:
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
			case now := <-pingTick.C:
				ping(now)
			}
		case rpcStatusClose:
			return
		}

	}
}
func (svr *RpcServer) Decompress(in []byte) (out []byte) {
	svr.zstdDecodeBuf2.Reset()
	svr.zstdDecodeBuf1.Reset()
	svr.zstdDecodeBuf1.Write(in)
	svr.zstdDecoder.Reset(svr.zstdDecodeBuf1)
	io.Copy(svr.zstdDecodeBuf2, svr.zstdDecoder)
	//libraries.DebugLog("解压前%d,解压后%d,压缩率%d",len(in),svr.zstdDecodeBuf2.Len(),len(in)*100/svr.zstdDecodeBuf2.Len())
	return svr.zstdDecodeBuf2.Bytes()
}
