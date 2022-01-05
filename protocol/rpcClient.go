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

type RpcQueryResult interface {
}

//封装一个用于发消息的，外部包通过msg发消息
type RpclientSend struct {
	*RpcClient
}

var (
	RpcClientQueryResultErrType = errors.New("result结构体必须为& *MSG_,并且包含QueryResultID")
	RpcClientQueryTimeOutErr    = errors.New("请求超时")
	rpcHanleMsgNum              = runtime.NumCPU()
	rpcClientQueryLock          sync.RWMutex
	rpcClientQueryMap           = make(map[uint32]chan RpcQueryResult)
	rpcClientQueryId            = cache.Hget("rpcClientQueryId", "rpcClientQuery")
)

type RpcClient struct {
	No                     uint8 //服务no
	Id                     uint8 //服务自身id
	ErrNum                 uint32
	CloseChan              chan bool //只允许在Close()里面发起chan
	ErrChan                chan string
	HandleMsg              func(*Msg)
	Addr                   string
	Status                 int
	IsMaster               bool //主服务器，维护host的cache
	DB                     *mysql.MysqlDB
	inchan                 chan *Msg
	outchan                chan *libraries.MsgBuffer
	conn                   net.Conn
	reconnect              chan []byte
	waitshutdown           sync.WaitGroup
	window                 int32 //接收窗口
	tokenKey               string
	cache                  *RpcCache
	tick                   *time.Ticker
	handleTick             func(time.Time)
	sendStruct             *RpclientSend
	decodebuf1, decodebuf2 *libraries.MsgBuffer
	decoder                *zstd.Decoder
}

func NewClient(no uint8, hostAddr string, tokenKey string) (*RpcClient, error) {

	client := &RpcClient{
		inchan:    make(chan *Msg, rpcHanleMsgNum*4),
		outchan:   make(chan *libraries.MsgBuffer, Rpcmsgnum),
		CloseChan: make(chan bool, 1),
		ErrChan:   make(chan string),
		reconnect: make(chan []byte, Rpcmsgnum),

		No:         no,
		Addr:       hostAddr,
		window:     DefaultWindowSize,
		Status:     RpcClientStatuShutdown,
		tokenKey:   tokenKey,
		tick:       time.NewTicker(RpcTickDefaultTime * time.Second),
		decodebuf1: new(libraries.MsgBuffer),
		decodebuf2: new(libraries.MsgBuffer),
	}
	cache := &RpcClient{
		inchan:    make(chan *Msg, rpcHanleMsgNum),
		outchan:   make(chan *libraries.MsgBuffer, Rpcmsgnum),
		CloseChan: make(chan bool, 1),
		ErrChan:   make(chan string),
		reconnect: make(chan []byte, Rpcmsgnum),

		No:         no + 128,
		Addr:       hostAddr,
		window:     DefaultWindowSize,
		Status:     RpcClientStatuShutdown,
		tokenKey:   tokenKey,
		tick:       time.NewTicker(RpcTickDefaultTime * time.Second),
		HandleMsg:  HandleCache,
		decodebuf1: new(libraries.MsgBuffer),
		decodebuf2: new(libraries.MsgBuffer),
	}
	err := client.Dial()
	if err != nil {
		return nil, err
	}
	err = cache.Dial()
	if err != nil {
		return nil, err
	}

	client.sendStruct = &RpclientSend{client}
	cache.sendStruct = &RpclientSend{cache}
	client.cache = &RpcCache{Svr: cache.sendStruct}
	client.decoder, _ = zstd.NewReader(client.decodebuf1)
	cache.decoder, _ = zstd.NewReader(cache.decodebuf1)
	return client, nil
}
func (client *RpcClient) Dial() error {
	conn, err := dail("tcp4", client.Addr, client)
	if err != nil {
		return err
	}
	client.conn = conn
	return nil
}
func (client *RpcClient) Start() {
	go client.handleWrite()
	go client.runTick()
	go client.handleRead()
	for i := 0; i < runtime.NumCPU(); i++ {
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
	data := GET_MSG_HOST_regServer()
	data.No = client.No
	data.Time = time.Now().Unix()
	data.Token = libraries.SHA256_S(client.tokenKey + strconv.Itoa(int(data.Time)))
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
			if msgNum > CompressMinNum {
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
			time.Sleep(time.Second * 2)
			n := 1
			libraries.ReleaseLog("重连")
			for ; n < MaxReconnectNum || MaxReconnectNum <= 0; n++ {
				err := client.Dial()
				if err != nil {
					libraries.ReleaseLog("尝试重新连接服务器失败%v", err)
					time.Sleep(time.Second)
					continue
				}

				break
			}
			if MaxReconnectNum > 0 && n >= MaxReconnectNum {
				client.Close("重连host失败次数过多，关闭本服务")
				return
			}
			libraries.ReleaseLog("重连ok")

			client.reg()
			rebuildReconnectData := func(data []byte) {
				//保证reconnect顺序不错乱
				var tmpchan = make(chan []byte, len(client.reconnect)+1)
				tmpchan <- data
				for i := 0; i < len(client.reconnect); i++ {
					tmpchan <- <-client.reconnect
				}
			}
			if data != nil {
				_, err := client.conn.Write(data)
				if err != nil {
					libraries.ReleaseLog("reconnectSusess消息，发消息错误%v", err)
					rebuildReconnectData(data)
					return
				}
			}
			//重新把之前发送失败的消息，发一遍
			for i := 0; i < len(client.reconnect); i++ {
				data := <-client.reconnect
				if data != nil {
					_, err := client.conn.Write(data)
					if err != nil {
						libraries.ReleaseLog("reconnectSusess消息，发消息错误%v", err)
						rebuildReconnectData(data)
						return
					}
				}

			}

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

func (client *RpcClient) handleMsg() {

	defer func() {
		if err := recover(); err != nil {
			go func() {
				client.ErrChan <- fmt.Sprint(err) + string("\r\n")
			}()
			atomic.AddUint32(&client.ErrNum, 1)
			go client.handleMsg()
		}
	}()

	for {
		select {
		case msg := <-client.inchan:
			msg.DB.DB = client.DB

			i := msg.Data

			msg.SetServer(client.sendStruct)

			switch data := i.(type) {
			case *MSG_HOST_regServer_result:
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
			case *MSG_HOST_StartTicker:
				client.IsMaster = true
				client.Status |= RpcTickStatusFirst
				client.tick.Reset(RpcTickDefaultTime * time.Second)
			case *MSG_HOST_PING:
				out := GET_MSG_HOST_PONG()
				msg.SendResult(out)
				out.Put()
			case *MSG_HOST_Transaction_Check:
				msg.WriteErr(MsgTransactionCheck(data))
			case *MSG_HOST_Transaction_Commit:
				MsgTransactionCommit(data)
			case *MSG_HOST_Transaction_RollBack:
				MsgTransactionRollBack(data)
			default:
				if SetMsgQuery(msg) {
					//这里不能回收
					continue
				} else {
					client.HandleMsg(msg)
				}
			}
			i.Put()
			BufPoolPut(msg.buf)
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

func (client *RpclientSend) SendMsg(remote uint16, msgno uint32, ttl uint16, transactionNo uint32, queryID uint32, out MSG_DATA) {
	SendMsg(uint16(client.No)|uint16(client.Id)<<8, remote, msgno, ttl, transactionNo, queryID, out, client.outchan)
}

func (client *RpclientSend) SendMsgWaitResult(remote uint16, msgno uint32, ttl uint16, transactionNo uint32, out MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	return SendMsgWaitResult(uint16(client.No)|uint16(client.Id)<<8, remote, msgno, ttl, transactionNo, out, result, client.outchan, timeout...)
}

//没有remote,msgno,ttl,transactionNo发送
func (client *RpclientSend) SendMsgToDefault(out MSG_DATA) {
	SendMsg(uint16(client.No)|uint16(client.Id)<<8, 0, 0, 0, 0, 0, out, client.outchan)
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
func (client *RpcClient) GetMsg() (*Msg, error) {
	data := GET_MSG_HOST_GET_Msgno()
	defer data.Put()
	var resdata *MSG_HOST_GET_Msgno_result
	err := client.sendStruct.SendMsgWaitResultToDefault(data, &resdata)
	if err != nil {
		return nil, err
	}
	msg := &Msg{Msgno: resdata.Msgno, DB: &MsgDB{DB: client.DB}}
	msg.SetServer(client.sendStruct)
	return msg, nil
}

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
func (client *RpcClient) Decompress(in []byte) (out []byte) {
	client.decodebuf2.Reset()
	client.decodebuf2.WriteByte(in[0] - 128)
	client.decodebuf1.Reset()
	client.decodebuf1.Write(in[1:])
	client.decoder.Reset(client.decodebuf1)
	io.Copy(client.decodebuf2, client.decoder)
	return client.decodebuf2.Bytes()
}
