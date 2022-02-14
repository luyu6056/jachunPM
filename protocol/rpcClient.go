package protocol

import (
	"errors"
	"fmt"
	"libraries"
	"mysql"
	"net"
	"reflect"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

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
	rpcClientQueryMap           = make(map[uint32]chan *Msg)
	rpcClientQueryId            = cache.Hget("rpcClientQueryId", "rpcClientQuery")
)

type RpcClient struct {
	No                   uint8           //服务no
	Id                   uint8           //服务自身id
	CloseChan            chan bool       //只允许在Close()里面发起chan
	HandleMsg            func(*Msg) bool //找到路由返回true，没找到false，以便于进行下一步操作
	Addr                 string
	Status               int
	IsMaster             bool //主服务器，维护host的cache
	DB                   *mysql.MysqlDB
	inchan               chan *Msg
	outchan              chan *libraries.MsgBuffer
	conn                 net.Conn
	reconnect            chan []byte
	waitshutdown         sync.WaitGroup
	window               int32 //接收窗口
	tokenKey             string
	cache                *RpcCache
	tick                 *time.Ticker
	handleTick           func(time.Time)
	sendStruct           *RpclientSend
	encodebuf, decodebuf *[]byte
}

func NewClient(no uint8, hostAddr string, tokenKey string) (*RpcClient, error) {
	buf1 := make([]byte, 65536)
	buf2 := make([]byte, 65536)
	buf3 := make([]byte, 65536)
	buf4 := make([]byte, 65536)
	client := &RpcClient{
		inchan:    make(chan *Msg, rpcHanleMsgNum*4),
		outchan:   make(chan *libraries.MsgBuffer, Rpcmsgnum),
		CloseChan: make(chan bool, 1),
		reconnect: make(chan []byte, Rpcmsgnum),
		No:        no,
		Addr:      hostAddr,
		window:    DefaultWindowSize,
		Status:    RpcClientStatuShutdown,
		tokenKey:  tokenKey,
		tick:      time.NewTicker(RpcTickDefaultTime * time.Second),
		encodebuf: &buf1,
		decodebuf: &buf2,
	}
	cache := &RpcClient{
		inchan:    make(chan *Msg, rpcHanleMsgNum),
		outchan:   make(chan *libraries.MsgBuffer, Rpcmsgnum),
		CloseChan: make(chan bool, 1),
		reconnect: make(chan []byte, Rpcmsgnum),

		No:        no + 128,
		Addr:      hostAddr,
		window:    DefaultWindowSize,
		Status:    RpcClientStatuShutdown,
		tokenKey:  tokenKey,
		tick:      time.NewTicker(RpcTickDefaultTime * time.Second),
		HandleMsg: HandleCache,
		encodebuf: &buf3,
		decodebuf: &buf4,
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
	client.sendStruct.SendMsgToDefault(nil, data)
	data.Put()
}
func (client *RpcClient) Local() uint16 {
	return uint16(client.No) | uint16(client.Id)<<8
}

func (client *RpcClient) EncodeBuf() (out *[]byte) {
	return client.encodebuf
}
func (client *RpcClient) DecodeBuf() (out *[]byte) {
	return client.decodebuf
}
func (client *RpcClient) handleWrite() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
		if client.outchan != nil {
			go client.handleWrite()
		}
	}()
	var (
		out    = &libraries.MsgBuffer{}
		msglen int
		//busyTime int64 = time.Now().UnixNano()

	)

	writeToBuf := func(o *libraries.MsgBuffer) {
		msglen += o.Len()
		out.Write(o.Bytes())
	}
	writeAllMsg := func() {
		for {
			select {
			case o1 := <-client.outchan:
				if msglen+o1.Len() > MaxOutLen {
					client.outchan <- o1
					break
				}
				writeToBuf(o1)
				o1.Reset()
				BufPoolPut(o1)
			default:
				if _, err := client.conn.Write(out.Bytes()); err != nil {
					libraries.ReleaseLog("正常消息，发消息错误%v", err)
					data := make([]byte, out.Len())
					copy(data, out.Bytes())
					client.reconnect <- data
				}
				return
			}
		}
	}
	for {
		select {
		case o := <-client.outchan:
			out.Reset()
			msglen = 0
			writeToBuf(o)
			o.Reset()
			BufPoolPut(o)
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

func (client *RpclientSend) SendMsg(msg *Msg, remote uint16, out MSG_DATA) {
	SendMsg(msg, client.Local(), remote, out, client.inchan, client.outchan)
}

func (client *RpclientSend) SendMsgWaitResult(msg *Msg, remote uint16, out MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	return SendMsgWaitResult(msg, client.Local(), remote, out, result, client.inchan, client.outchan, timeout...)
}

//没有remote,msgno,ttl,transactionNo发送
func (client *RpclientSend) SendMsgToDefault(msg *Msg, out MSG_DATA) {
	SendMsg(msg, client.Local(), 0, out, client.inchan, client.outchan)
}

func (client *RpclientSend) SendMsgWaitResultToDefault(msg *Msg, out MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	return SendMsgWaitResult(msg, client.Local(), 0, out, result, client.inchan, client.outchan, timeout...)
}
func (client *RpcClient) CacheGet(serverNo uint8, path string, key string, value interface{}) (err error) {
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
		return errors.New("CacheGet消息不够长，不足以读取一条缓存")
	}
	return err
}
func (client *RpcClient) CacheGetPath(serverNo uint8, path string, names ...string) ([][]byte, error) {

	return client.cache.GetPath(strconv.Itoa(int(serverNo))+"_"+path, names...)
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

func (client *RpcClient) GetMsg() (*Msg, error) {
	msg := &Msg{DB: &MsgDB{DB: client.DB}}
	msg.SetServer(client.sendStruct)
	return msg, nil
}

func (client *RpcClient) SetConfig(lang CountryNo, key string, config map[string]map[string]interface{}) (err error) {
	return client.cache.Set(key, PATH_CONFIG_CACHE+lang.String(), libraries.JsonMarshal(config), 0)
}
func (client *RpcClient) GetUserCacheById(id int32) (user *MSG_USER_INFO_cache) {
	err := client.CacheGet(UserServerNo, PATH_USER_INFO_CACHE, strconv.Itoa(int(id)), &user)
	if err != nil {
		libraries.DebugLog("获取user缓存失败%+v", err)
	}
	return
}
func (client *RpcClient) GetUserCacheByIds(ids []int32) (user []*MSG_USER_INFO_cache) {
	if len(ids) == 0 {
		return
	}
	var idstr = make([]string, len(ids))
	for k, v := range ids {
		idstr[k] = strconv.Itoa(int(v))
	}
	res, err := client.cache.GetPath(strconv.Itoa(int(UserServerNo))+"_"+PATH_USER_INFO_CACHE, idstr...)
	if err != nil {
		return nil
	}

	buf := BufPoolGet()
	for _, b := range res {
		buf.Reset()
		buf.Write(b)
		if v, ok := READ_MSG_DATA(buf).(*MSG_USER_INFO_cache); ok {
			user = append(user, v)
		} else {
			user = append(user, nil)
		}
	}
	buf.Reset()
	BufPoolPut(buf)
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
func (client *RpcClient) GetdeptCacheById(deptId int32) (deptinfo *MSG_USER_Dept_cache, err error) {
	if err = client.CacheGet(UserServerNo, PATH_USER_DEPT_CACHE, strconv.Itoa(int(deptId)), &deptinfo); err != nil {
		err = errors.New(fmt.Sprintf("无法获取dept id%d的缓存,%+v", deptId, err))
	}
	return
}
func (client *RpcClient) SetTickHand(f func(time.Time)) {
	client.handleTick = f
}
func (client *RpcClient) GetProjectById(id int32) (res *MSG_PROJECT_project_cache) {
	err := client.CacheGet(ProjectServerNo, PATH_PROJECT_PROJECT_CACHE, strconv.Itoa(int(id)), &res)
	if err != nil {
		libraries.DebugLog("获取project  id %d 缓存失败%+v", id, err)
	}
	return
}

func (client *RpcClient) handleMsg() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
		go client.handleMsg()
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
				if !client.HandleMsg(msg) {
					if SetMsgQuery(msg) {
						//这里不能回收
						continue
					} else {
						if v, ok := CmdToName[msg.Cmd]; ok {
							libraries.ReleaseLog("未设置消息CMD%s处理", v)
						} else {
							libraries.ReleaseLog("未设置消息CMD%d处理", msg.Cmd)
						}
					}
				}
			}
			if i != nil {
				i.Put()
			}
			BufPoolPut(msg.buf)
		case <-client.CloseChan:
			client.CloseChan <- true
			return
		}
	}
}
