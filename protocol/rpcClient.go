package protocol

import (
	"errors"
	"fmt"
	"io"
	"libraries"
	"mysql"
	"net"
	"reflect"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/klauspost/compress/zstd"
	"github.com/luyu6056/cache"
)

const (
	RpcClientStatuShutdown = 1 << iota
	RpcClientStatuNormal
	RpcTickStatusFirst //tick first时要初始化缓存
)

type RpcQuery interface {
	getQueryID() uint32
	setQueryID(id uint32)
}
type RpcQueryResult interface {
	getQueryResultID() uint32
	setQueryResultID(id uint32)
}

//封装一个用于发消息的，外部包通过msg发消息
type RpclientSend struct {
	*RpcClient
}

var (
	RpcClientQueryResultErrType = errors.New("result结构体必须为& *MSG_,并且包含QueryResultID")
	RpcClientQueryTimeOutErr    = errors.New("请求超时")
	rpcHanleMsgNum              = runtime.NumCPU()
	rpcClientQueryLock          sync.Mutex
	rpcClientQueryMap           = make(map[uint32]chan RpcQueryResult)
	rpcClientQueryId            = cache.Hget("rpcClientQueryId", "rpcClientQuery")
)

type RpcClient struct {
	No                       uint8 //服务no
	Id                       uint8 //服务自身id
	ErrNum                   uint32
	CloseChan                chan bool //只允许在Close()里面发起chan
	ErrChan                  chan string
	HandleMsg                func(*Msg)
	Addr                     string
	Status                   int
	IsMaster                 bool //主服务器，维护host的cache
	DB                       *mysql.MysqlDB
	inchan                   chan *Msg
	outchan                  chan *libraries.MsgBuffer
	conn                     net.Conn
	reconnect, reconnectData chan []byte
	wait                     chan bool
	waitshutdown             sync.WaitGroup
	window                   int32 //接收窗口
	tokenKey                 string
	cache                    *RpcCache
	tick                     *time.Ticker
	handleTick               func(time.Time)
	sendStruct               *RpclientSend
}

func NewClient(no uint8, hostAddr string, tokenKey string) (*RpcClient, error) {
	conn, err := net.Dial("tcp4", hostAddr)
	if err != nil {
		return nil, err
	}
	cacheConn, err := net.Dial("tcp4", hostAddr)
	if err != nil {
		return nil, err
	}
	client := &RpcClient{
		inchan:        make(chan *Msg, rpcHanleMsgNum),
		outchan:       make(chan *libraries.MsgBuffer, Rpcmsgnum),
		CloseChan:     make(chan bool, 1),
		ErrChan:       make(chan string),
		reconnect:     make(chan []byte, Rpcmsgnum),
		reconnectData: make(chan []byte, Rpcmsgnum),
		wait:          make(chan bool),

		No:       no,
		Addr:     hostAddr,
		conn:     conn,
		window:   DefaultWindowSize,
		Status:   RpcClientStatuShutdown,
		tokenKey: tokenKey,
		tick:     time.NewTicker(RpcTickDefaultTime * time.Second),
	}
	cache := &RpcClient{
		inchan:        make(chan *Msg, rpcHanleMsgNum),
		outchan:       make(chan *libraries.MsgBuffer, Rpcmsgnum),
		CloseChan:     make(chan bool, 1),
		ErrChan:       make(chan string),
		reconnect:     make(chan []byte, Rpcmsgnum),
		reconnectData: make(chan []byte, Rpcmsgnum),
		wait:          make(chan bool),

		No:        no + 128,
		Addr:      hostAddr,
		conn:      cacheConn,
		window:    DefaultWindowSize,
		Status:    RpcClientStatuShutdown,
		tokenKey:  tokenKey,
		tick:      time.NewTicker(RpcTickDefaultTime * time.Second),
		HandleMsg: HandleCache,
	}
	client.sendStruct = &RpclientSend{client}
	cache.sendStruct = &RpclientSend{cache}
	client.cache = &RpcCache{Svr: cache.sendStruct}
	return client, nil
}
func (client *RpcClient) Start() {
	go client.handleWrite()
	go client.runTick()
	go client.handleRead()
	for i := 0; i < rpcHanleMsgNum; i++ {
		go client.handleMsg()
	}
	if client.cache != nil {
		client.cache.Svr.(*RpclientSend).Start()
		client.waitshutdown.Add(6)
	}

	client.reg()
	client.waitshutdown.Wait()
}
func (client *RpcClient) Close(reason string) {
	libraries.ReleaseLog(reason)

	client.Status = RpcClientStatuShutdown
	client.CloseChan <- true
}
func (client *RpcClient) reg() {
	//注册rpc服务
	data := GET_MSG_COMMON_regServer()
	data.No = client.No
	data.Time = time.Now().Unix()
	data.Token = libraries.SHA256_S(client.tokenKey + strconv.Itoa(int(data.Time)))
	data.Window = client.window
	client.sendStruct.SendMsgToDefault(data)
	data.Put()
}
func (client *RpcClient) resetWindow() {
	//重置窗口，请不要在其他地方调用，下面在conn.wait卡住的情况下调用了，保证窗口的唯一性
	data := GET_MSG_COMMON_ResetWindow()
	client.window = DefaultWindowSize
	data.Window = client.window
	client.sendStruct.SendMsgToDefault(data)
	data.Put()
}

func (client *RpcClient) handleWrite() {
	defer func() {
		if err := recover(); err != nil {
			client.ErrChan <- fmt.Sprintf("%v\r\n", err) + string(debug.Stack())
			atomic.AddUint32(&client.ErrNum, 1)
		}
		if client.outchan != nil {
			go client.handleWrite()
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
			if msgNum > CompressMinNum || out.Len()+o.Len() > CompressMinLen {
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
	writeAllMsg := func() {

		for i := 0; i < len(client.outchan) && i < MaxMsgNum-1; i++ {
			o1 := <-client.outchan
			if msglen+o1.Len() > MaxOutLen {
				client.outchan <- o1
				break
			}
			msgNum++
			writeToBuf(o1)
			o1.Reset()
			//BufPoolPut(o1)

		}
		if compress {
			//有压缩
			zstdWriter.Close()
			msglen = zstdbuf.Len()
			if msglen > MaxOutLen {
				panic("消息大于最大允许长度,压缩后更长？？？，请检查代码")
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
			if msglen > MaxOutLen {
				panic("消息大于最大允许长度,前面的限制有错，请重写代码")
				return
			}
			if msglen != out.Len() {
				panic("msglen计算有误，请检查代码")
				return
			}
			b := msgbuf.Make(5 + msglen)
			b[0] = byte(msglen)
			b[1] = byte(msglen >> 8)
			b[2] = byte(msglen >> 16)
			b[3] = byte(msglen >> 24)
			b[4] = byte(msgNum)
			copy(b[5:], out.Bytes())
		}
		_, err := client.conn.Write(msgbuf.Bytes())
		if err != nil {
			libraries.ReleaseLog("正常消息，发消息错误%v", err)
			data := make([]byte, msgbuf.Len())
			copy(data, msgbuf.Bytes())
			client.reconnect <- data
		}
	}
	for {
		select {
		case o := <-client.outchan:
			compress = false
			out.Reset()
			msgbuf.Reset()
			zstdbuf.Reset()
			msglen = 0
			writeToBuf(o)
			msgNum = 1
			o.Reset()
			//BufPoolPut(o)

			writeAllMsg()
		case data := <-client.reconnect: //重连
			client.tick.Stop()
			client.IsMaster = false
			client.conn.Close()
			n := 1
			libraries.ReleaseLog("重连")
			for ; n < MaxReconnectNum || MaxReconnectNum <= 0; n++ {
				newconn, err := net.Dial("tcp4", client.Addr)
				if err != nil {
					libraries.ReleaseLog("尝试重新连接服务器失败%v", err)
					time.Sleep(time.Second)
					continue
				}
				client.conn = newconn
				break
			}
			if MaxReconnectNum > 0 && n >= MaxReconnectNum {
				client.Close("重连host失败次数过多，关闭本服务")
				return
			}
			libraries.ReleaseLog("重连ok")
			//连接成功后，把发送失败的消息重新发一遍
			client.reconnectData <- data
			for i := 0; i < len(client.reconnect); i++ {
				client.reconnectData <- <-client.reconnect
			}
			client.reg()

			//重新把之前发送失败的消息，发一遍
			for i := 0; i < len(client.reconnectData); i++ {
				data := <-client.reconnectData
				if data != nil {
					_, err := client.conn.Write(data)
					if err != nil {
						libraries.ReleaseLog("reconnectSusess消息，发消息错误%v", err)
						client.reconnect <- data
					}
				}

			}
			//重置窗口，避免发送失败的消息增加到不对的窗口值
			client.resetWindow()
			//让handleRead继续工作
			client.wait <- true
		case <-client.CloseChan:
			writeAllMsg()
			close(client.outchan)
			client.outchan = nil
			client.CloseChan <- true
			client.waitshutdown.Done()
			return
		}
	}
}

func (client *RpcClient) handleRead() {
	defer func() {
		if err := recover(); err != nil {
			client.ErrChan <- fmt.Sprintf("%v\r\n", err) + string(debug.Stack())
			atomic.AddUint32(&client.ErrNum, 1)
			go client.handleRead()
		}
	}()
	var in *libraries.MsgBuffer
	buf := &libraries.MsgBuffer{}
	decodebuf1 := &libraries.MsgBuffer{}
	decodebuf2 := &libraries.MsgBuffer{}
	decoder, _ := zstd.NewReader(decodebuf1)
	lenbuf := make([]byte, 5)
	for {
		n, err := io.ReadFull(client.conn, lenbuf)
		if err != nil || n != 5 {
			libraries.ReleaseLog("读消息失败，err %v，读取数量 %d", err, n)
			client.reconnect <- nil
			select {
			case <-client.wait:

			case <-client.CloseChan:
				client.waitshutdown.Done()
				client.CloseChan <- true
				return
			}
			continue
		}
		msglen := int(lenbuf[0]) | int(lenbuf[1])<<8 | int(lenbuf[2])<<16 | int(lenbuf[3])<<24
		msgnum := lenbuf[4] & 127
		compress := lenbuf[4]>>7 == 1
		buf.Reset()
		b := buf.Make(msglen)
		n, err = io.ReadFull(client.conn, b)
		if err != nil || n != msglen {
			libraries.ReleaseLog("读消息失败，err %v，读取数量 %d", err, n)
			client.reconnect <- nil
			select {
			case <-client.wait:

			case <-client.CloseChan:
				client.waitshutdown.Done()
				client.CloseChan <- true
				return
			}
			continue
		}
		//解压缩
		if compress {
			decodebuf1.Reset()
			decodebuf1.Write(b)
			decoder.Reset(decodebuf1)
			io.Copy(decodebuf2, decoder)
			in = decodebuf2
		} else {
			in = buf
		}
		n = 0
		for ; in.Len() > 0; n++ {
			msg, err := ReadOneMsg(in)
			if err != nil {
				libraries.ReleaseLog("读消息错误%v", err)
				in.Reset()
			} else {
				msg.ReadData()
				client.inchan <- msg
			}

		}
		if n != int(msgnum) {
			libraries.DebugLog("读消息数量错误，请检查协议，消息总量%d,已读%d", msgnum, n)
		}
		client.window -= int32(msgnum)
		if client.Status&RpcClientStatuNormal == RpcClientStatuNormal && client.window < DefaultWindowSize/2 {
			data := GET_MSG_COMMON_WINDOW_UPDATE()
			data.Add = DefaultWindowSize - client.window
			client.window = DefaultWindowSize
			//libraries.DebugLog("增加窗口%d，实际窗口%d", data.Add, client.window)
			client.sendStruct.SendMsgToDefault(data)
			data.Put()
		}

	}

}
func (client *RpcClient) handleMsg() {

	defer func() {
		if err := recover(); err != nil {
			client.ErrChan <- fmt.Sprint(err) + string("\r\n") + string(debug.Stack())
			atomic.AddUint32(&client.ErrNum, 1)
			go client.handleMsg()
		}
	}()

	for {
		select {
		case msg := <-client.inchan:
			msg.DB.DB = client.DB

			i := msg.Data
			if client.cache != nil {
				msg.SetServer(client.cache.Svr) //使用cache的conn处理msg.SendResult等消息内发送，减少对主连接的协程阻塞
			} else {
				msg.SetServer(client.sendStruct)
			}
			switch data := i.(type) {
			case *MSG_COMMON_regServer_result:
				if client.Status&RpcClientStatuShutdown == RpcClientStatuShutdown {
					client.Status -= RpcClientStatuShutdown
				}
				client.Id = data.Id
				client.Status |= RpcClientStatuNormal
				if client.cache != nil {
					libraries.ReleaseLog("连接host成功，本服务ID%d", client.Id)
				} else {
					libraries.ReleaseLog("连接host_cache成功,cacheId%d", client.Id)
				}
			case *MSG_COMMON_StartTicker:
				client.IsMaster = true
				client.Status |= RpcTickStatusFirst
				client.tick.Reset(RpcTickDefaultTime * time.Second)
			case *MSG_COMMON_PING:
				out := GET_MSG_COMMON_PONG()
				client.sendStruct.SendMsgToDefault(out)
				out.Put()
			case *MSG_COMMON_Transaction_Check:

				msgTransactionCheck(data, msg)
			case *MSG_COMMON_Transaction_Commit:
				msgTransactionCommit(data, msg)
			case *MSG_COMMON_Transaction_RollBack:
				msgTransactionRollBack(data, msg)
			default:
				if SetMsgQuery(i) {
					//这里不能回收
					continue
				} else {
					client.HandleMsg(msg)
				}
			}
			i.Put()
		case errstr := <-client.ErrChan:
			libraries.ReleaseLog(errstr)
		case <-client.CloseChan:
			client.CloseChan <- true
			return
		}
	}
}

func (client *RpcClient) runTick() {
	for {
		select {
		case now := <-client.tick.C:
			if client.handleTick != nil {
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println("定时任务发生错误", err)
							debug.PrintStack()
						}
					}()
					client.handleTick(now)
				}()
			}
		case <-client.CloseChan:
			client.waitshutdown.Done()
			client.CloseChan <- true
			return
		}
	}
}

func (client *RpclientSend) SendMsg(remote uint16, msgno uint32, ttl uint8, transactionNo uint32, out MSG_DATA) {
	SendMsg(uint16(client.No)|uint16(client.Id)<<8, remote, msgno, ttl, transactionNo, out, client.outchan)
}

func (client *RpclientSend) SendMsgWaitResult(remote uint16, msgno uint32, ttl uint8, transactionNo uint32, out MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	return SendMsgWaitResult(uint16(client.No)|uint16(client.Id)<<8, remote, msgno, ttl, transactionNo, out, result, client.outchan, timeout...)
}

//没有remote,msgno,ttl,transactionNo发送
func (client *RpclientSend) SendMsgToDefault(out MSG_DATA) {
	SendMsg(uint16(client.No)|uint16(client.Id)<<8, 0, 0, 0, 0, out, client.outchan)
}

func (client *RpclientSend) SendMsgWaitResultToDefault(out MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	return SendMsgWaitResult(uint16(client.No)|uint16(client.Id)<<8, 0, 0, 0, 0, out, result, client.outchan, timeout...)
}
func (client *RpcClient) CacheGet(serverNo uint8, path, key string, value interface{}) (err error) {
	b, err := client.cache.Get(key, strconv.Itoa(int(serverNo))+"_"+path)
	if len(b) > 4 {
		cmd := int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
		buf := BufPoolGet()
		defer func() {
			if r := recover(); r != nil {
				err = errors.New("cache反序列化错误" + fmt.Sprint(r))
			}
			buf.Reset()
			BufPoolPut(buf)
		}()
		buf.Write(b[4:])
		if v, ok := value.(MSG_DATA); ok {
			if cmd == v.cmd() {
				v.read(buf)
				return
			}
		} else {
			r := reflect.ValueOf(value)
			if r.Elem().IsNil() {
				if f, ok := cmdMapFunc[cmd]; ok {
					v := reflect.ValueOf(f(buf))
					if v.Kind() == r.Elem().Kind() {
						r.Elem().Set(v)
						return
					}
				}
			}

		}
		return errors.New("找不到反序列化方法或者cmd值不对")
	} else {
		return errors.New("消息不够长，不足以读取一条缓存")
	}
	return err
}
func (client *RpcClient) CacheGetPath(serverNo uint8, path string) ([][]byte, error) {
	return client.cache.GetPath(strconv.Itoa(int(serverNo)) + "_" + path)
}

//只允许修改本服务的缓存
func (client *RpcClient) CacheSet(path, key string, value MSG_DATA, expire int64) error {
	buf := BufPoolGet()
	value.write(buf)
	err := client.cache.Set(key, strconv.Itoa(int(client.No))+"_"+path, buf.Bytes(), expire)
	buf.Reset()
	BufPoolPut(buf)
	return err
}
func (client *RpcClient) CacheDel(path, key string) error {
	return client.cache.Del(key, strconv.Itoa(int(client.No))+"_"+path)
}
func (client *RpcClient) CacheDelPath(path string) error {
	return client.cache.DelPath(strconv.Itoa(int(client.No)) + "_" + path)
}

//暂时不开放getMsg接口
/*func (client *RpcClient) getMsg() (*Msg, error) {
	data := GET_MSG_COMMON_GET_Msgno()
	defer data.Put()
	var resdata *MSG_COMMON_GET_Msgno_result
	err := client.sendStruct.SendMsgWaitResultToDefault(data, &resdata)
	if err != nil {
		return nil, err
	}
	msg := &Msg{Msgno: resdata.Msgno, DB: &MsgDB{DB: client.DB}}
	msg.SetServer(client.sendStruct)
	return msg, nil
}*/

func (client *RpcClient) SetConfig(lang CountryNo, key string, config map[string]map[string]interface{}) (err error) {
	return client.cache.Set(key, PATH_CONFIG_CACHE+lang.String(), libraries.JsonMarshal(config), 0)
}
func (client *RpcClient) GetUserCacheById(id int32) (user *MSG_USER_INFO_cache) {
	if id <= 0 {
		return nil
	}
	err := client.CacheGet(UserServerNo, PATH_USER_INFO_CACHE, strconv.Itoa(int(id)), &user)
	if err != nil {
		libraries.DebugLog("获取user缓存失败%+v", err)
	}
	return
}
func (client *RpcClient) GetTreeById(moduleID int32) (res *MSG_PROJECT_tree_cache) {
	err := client.CacheGet(ProjectServerNo, PATH_PROJECT_TREE_CACHE, strconv.Itoa(int(moduleID)), &res)
	if err != nil {
		libraries.DebugLog("获取tree缓存失败%+v", err)
	}
	return
}
func (client *RpcClient) GetProductById(productID int32) (res *MSG_PROJECT_product_cache) {
	err := client.CacheGet(ProjectServerNo, PATH_PROJECT_PRODUCT_CACHE, strconv.Itoa(int(productID)), &res)
	if err != nil {
		libraries.DebugLog("获取Product缓存失败%+v", err)
	}
	return
}
func (client *RpcClient) SetTickHand(f func(time.Time)) {
	client.handleTick = f
}
func (client *RpcClient) GetProjectById(id int32) (res *MSG_PROJECT_project_cache) {
	err := client.CacheGet(ProjectServerNo, PATH_PROJECT_PROJECT_CACHE, strconv.Itoa(int(id)), &res)
	if err != nil {
		libraries.DebugLog("获取project缓存失败%+v", err)
	}
	return
}
