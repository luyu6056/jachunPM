package rpcHost

import (
	"errors"
	"fmt"
	"io/ioutil"
	"jachunPM_commom/db"
	"libraries"
	"os"
	"protocol"
	"reflect"
	"runtime/debug"
	"server"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/luyu6056/cache"
	"github.com/luyu6056/gnet"
	"github.com/panjf2000/ants/v2"
)

type transactionOption bool

const (
	TransactionCacheKey       = "TransactionCache"
	TransactionOptionCommit   = transactionOption(true)
	TransactionOptionRollback = transactionOption(false)
)

var (
	filepath                       string
	errTransactionNo               = errors.New("TransactionNotFoundNo")
	errTransactionNotAllowCommit   = errors.New("TransactionNotAllowCommit")
	errTransactionNotFoundSvr      = errors.New("TransactionNotFoundServer")
	errTransactionNotFoundWaitChan = errors.New("TransactionNotFoundWaitChan")
	errTransactionNotFoundErrChan  = errors.New("TransactionNotFoundErrChan")
	hostGoPool, _                  = ants.NewPool(10000)
	rpcHostMsgInChan               = make(chan *protocol.Msg, protocol.Rpcmsgnum)
)

func init() {
	var err error
	filepath, err = libraries.GetBaseRootPath()
	if err != nil {
		panic("获取运行根目录失败" + err.Error())
	}
	filepath += "/upload/"
}
func SendMsgToRemote(ctx *server.Context, c gnet.Conn) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
			if svr, ok := c.Context().(*RpcServer); ok {
				atomic.AddUint32(&svr.ErrNum, 1)
			}

		}
	}()
	b := ctx.In.Next(1)
	msgnum := b[0]
	var n int
	for ctx.In.Len() > 0 {
		n++
		in, err := protocol.ReadOneMsg(ctx.In)
		if err != nil {
			return errors.New("读消息出错" + err.Error())
		} else {
			if c.Context() == nil {
				in.ReadData()
				i := in.Data
				switch data := i.(type) {
				case *protocol.MSG_COMMON_regServer:
					svr := NewRpcServer(c)
					svr.Start(data.No, data.IpPort, data.Window)
					c.SetContext(svr)
				case nil:
					libraries.ReleaseLog("未注册服务，消息读取失败nil,请检查协议")
				default:
					libraries.ReleaseLog("未注册服务，host未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
				}
				continue
			}

			svr, ok := c.Context().(*RpcServer)
			if !ok {
				return errors.New("收到非rpcserver消息")
			} else if in.Local != uint16(svr.ServerNo)|uint16(svr.Id)<<8 { //检查local
				fmt.Println(in.Cmd)
				return errors.New(fmt.Sprintf("消息的local来源不对,in.Local%d,local%d", in.Local, uint16(svr.ServerNo)|uint16(svr.Id)<<8))
			}

			buf := protocol.BufPoolGet()
			in.NextWithSetInBuf(buf)
			rpcHostMsgInChan <- in
		}
	}
	if int(msgnum) != n {
		libraries.DebugLog("读消息数量错误，请检查协议，消息总量%d,已读%d", msgnum, n)
	}
	return nil
}

var msgnoTtl sync.Map
var transactionNoCache = cache.Hget("transactionNo", "Common")

func HostServerHandlerOutChan() {
	if libraries.IsRelease {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				debug.PrintStack()
				go HostServerHandlerOutChan()
			}
		}()
	}
	//把buf转化为msg
	for buf := range rpcServerOutChan[protocol.HostServerNo] {
		in, err := protocol.ReadOneMsg(buf)
		if err != nil {
			libraries.ReleaseLog("hostOutChan读消息出错" + err.Error())
			continue
		}
		rpcHostMsgInChan <- in
	}

}
func HostServerHandlerMsgIn() {
	if libraries.IsRelease {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				debug.PrintStack()
				go HostServerHandlerMsgIn()
			}
		}()

	}
	for in := range rpcHostMsgInChan {
		buf := in.Buf()
		cmdSvrNo := byte(in.Cmd)
		if cmdSvrNo != protocol.HostServerNo || cmdSvrNo == protocol.HostServerNo && in.Remote != 0 {
			//检查msgno
			b := buf.Bytes()
			if in.Msgno == 0 {
				msgno := atomic.AddUint32(&globalMsgno, 1)
				in.Msgno = msgno
				b[0] = byte(in.Msgno)
				b[1] = byte(in.Msgno >> 8)
				b[2] = byte(in.Msgno >> 16)
				b[3] = byte(in.Msgno >> 24)
				ttl := int32(0)
				msgnoTtl.Store(in.Msgno, &ttl)
				time.AfterFunc(protocol.MsgTimeOut*time.Second, func() { msgnoTtl.Delete(msgno) })
			} else {
				if v, ok := msgnoTtl.Load(in.Msgno); ok {
					ttl := v.(*int32)
					newTtl := atomic.AddInt32(ttl, 1)
					if newTtl >= protocol.MaxMsgTtl {
						//抛弃消息，记录log
						db.WriteMsgLog(in)
						protocol.BufPoolPut(buf)
						continue
					}
					in.Ttl = uint8(newTtl)
				} else {
					libraries.DebugLog("无效的msgno %d", in.Msgno)
					protocol.BufPoolPut(buf)
					continue //抛弃消息
				}
			}
			db.WriteMsgLog(in)

			b[4] = in.Ttl
			if in.Remote == 0 {
				libraries.DebugLog("cmd%s ,No%d", protocol.CmdToName[in.Cmd], cmdSvrNo)
				rpcServerOutChan[cmdSvrNo] <- buf
			} else {
				svrNo := byte(in.Remote)
				id := byte(in.Remote >> 8)
				rpcLock.RLock()
				if remoteSvr := rpcServerIdList[svrNo][id]; remoteSvr != nil {
					remoteSvr.outChan <- buf
				} else {
					//服务挂壁了，转发到公共消息
					rpcServerOutChan[svrNo] <- buf
				}
				rpcLock.RUnlock()
			}
			continue
		}

		in.ReadData()
		//hostAsyncHand是异步执行，如果有需要同步的需要在这里进行处理
		hostAsyncHand.Invoke(in)
		buf.Reset()
		protocol.BufPoolPut(buf)
	}

}

var hostAsyncHand, _ = ants.NewPoolWithFunc(10000, func(args interface{}) {
	in, ok := args.(*protocol.Msg)
	if !ok {
		return
	}
	rpcLock.RLock()
	defer rpcLock.RUnlock()
	svr := rpcServerIdList[uint8(in.Local)][uint8(in.Local>>8)]
	if svr == nil {
		//可能服务掉线了,暂不处理
		libraries.DebugLog("host收到不存在的svr,No%d,Id%d", uint8(in.Local), uint8(in.Local>>8))
		return
	}
	in.SetServer(svr)
	i := in.Data
	//defer i.Put()  SetMsgQuery里面不能回收，所以不能defer回收

	switch data := i.(type) {
	case *protocol.MSG_COMMON_WINDOW_UPDATE:
		if svr != nil {
			atomic.AddInt32(&svr.window, data.Add)
			//尝试把半开的服务恢复正常
			//libraries.DebugLog("服务%v，ID%v，更新窗口%d", svr.ServerNo, svr.Id, svr.window)
			select {
			case svr.setStatusOpenChan <- "更新窗口值":
			default:
			}
		}

	case *protocol.MSG_COMMON_PONG:
		if svr != nil {
			//libraries.DebugLog("服务%d,ID%d,收到pong", svr.ServerNo, svr.Id)
			svr.pongTime = time.Now().Unix()
			if svr.window > 0 {
				//尝试把pong响应超时导致半开的服务恢复正常
				select {
				case svr.setStatusOpenChan <- "收到pong":
				default:
				}
			}
		}
	case *protocol.MSG_COMMON_CACHE_DEL:
		cache.Hdel(data.Name, data.Path)
	case *protocol.MSG_COMMON_CACHE_DelPath:
		cache.Hdel_all(data.Path)
	case *protocol.MSG_COMMON_CACHE_SET:
		cache.Hset(data.Name, map[string][]byte{"value": data.Value}, data.Path, data.Expire)
	case *protocol.MSG_COMMON_CACHE_GET:
		r := cache.Hget(data.Name, data.Path)
		out := protocol.GET_MSG_COMMON_CACHE_GET_result()
		r.Get("value", &out.Value)
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_COMMON_CACHE_GETPATH:
		out := protocol.GET_MSG_COMMON_CACHE_GETPATH_result()
		cache.RangePath(data.Path, func(key string, v *cache.Hashvalue) bool {
			var value []byte
			if v.Get("value", &value) {
				out.Value = append(out.Value, value)
			}
			return true
		})
		out.QueryResultID = data.QueryID
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_COMMON_CACHE_GET_result:
		protocol.HandleCache(in)
	case *protocol.MSG_COMMON_GET_Msgno:
		if svr != nil {
			out := protocol.GET_MSG_COMMON_GET_Msgno_result()
			msgno := atomic.AddUint32(&globalMsgno, 1)
			out.Msgno = msgno
			ttl := int32(0)
			msgnoTtl.Store(msgno, &ttl)
			time.AfterFunc(protocol.MsgTimeOut*time.Second, func() { msgnoTtl.Delete(msgno) })
			out.QueryResultID = data.QueryID
			in.SendResult(out)
			out.Put()
		}
	case *protocol.MSG_COMMON_regServer:
		//common掉线可能会导致其他服务反复发送reg
		if data.No != svr.ServerNo {
			libraries.DebugLog("注册的serverNo不对，注册%d,实际%d", data.No, svr.ServerNo)
		}
	case *protocol.MSG_COMMON_ResetWindow:
		svr.window = data.Window
	case *protocol.MSG_FILE_upload:
		var dir string
		now := time.Now()
		if data.Code != "" {
			dir = data.Code + "/" + data.Type + "/" + now.Format("2006") + "/" + now.Format("01") + "/" + now.Format("02") + "/"
		} else {
			dir = now.Format("200601") + "/"
		}
		filedir := strings.ReplaceAll(filepath+dir, "/", string(os.PathSeparator))
		exist, err := libraries.PathExists(filedir)
		defer func() {
			if err != nil {
				in.WriteErr(err)
			}
		}()
		if err != nil {
			return
		}
		if !exist {
			err = os.MkdirAll(filedir, os.ModePerm)
			if err != nil {
				return
			}
		}
		var f *os.File
		f, err = os.Create(filedir + data.Name)
		if err != nil {
			return
		}
		defer f.Close()
		_, err = f.Write(data.Data)
		if err != nil {
			return
		}
		adduser, _ := Host.GetUserCacheById(data.AddBy)
		var addedby string
		if adduser != nil {
			addedby = adduser.Account
		}
		var file = &db.File{
			Pathname:   dir + data.Name,
			Title:      data.Name,
			Extension:  data.Name[strings.LastIndex(data.Name, ".")+1:],
			Size:       len(data.Data),
			ObjectType: data.ObjectType,
			ObjectID:   data.ObjectID,
			AddedBy:    addedby,
			AddedDate:  now,
			Type:       data.Type,
		}
		var id int64
		id, err = db.DB.Table(db.TABLE_FILE).Insert(file)
		if err != nil {
			os.Remove(filepath + dir + data.Name)
			return
		}
		out := protocol.GET_MSG_FILE_upload_result()
		out.FileID = id
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_FILE_DeleteByID:
		var file *db.File
		err := db.DB.Table(db.TABLE_FILE).Prepare().Where("Id=?", data.FileID).Find(&file)
		if err != nil {
			in.WriteErr(err)
			return
		}
		os.Remove(filepath + file.Pathname)
		db.DB.Table(db.TABLE_FILE).Where("Id=?", data.FileID).Delete()
	case *protocol.MSG_FILE_getByID:
		var file *db.File
		err := db.DB.Table(db.TABLE_FILE).Prepare().Where("Id=?", data.FileID).Find(&file)
		if err != nil {
			in.WriteErr(err)
			return
		}
		b, err := ioutil.ReadFile(filepath + file.Pathname)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_FILE_getByID_result()
		out.Data = b
		out.Ext = file.Extension
		out.Name = file.Title
		out.Type = file.Type
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_COMMON_BeginTransaction:
		//发起事务的服务器
		startSvr := data.TransactionNo == 0
		for data.TransactionNo == 0 {
			data.TransactionNo = uint32(transactionNoCache.INCRBY("NO", 1))
		}
		out := protocol.GET_MSG_COMMON_BeginTransaction_result()
		out.TransactionNo = data.TransactionNo

		if startSvr {
			transactionCache := cache.Hget(strconv.Itoa(int(data.TransactionNo)), TransactionCacheKey)
			transactionCache.Expire(protocol.MsgTimeOut * 2)
			waitChan := make(chan transactionOption, 0xffff) //理论上服务有多少个，chan长度就至少要有多少
			errChan := make(chan error, 1)
			transactionCache.Set(strconv.Itoa(int(in.Local)), true)
			transactionCache.Set("startSvr", in.Local)
			transactionCache.Set("waitChan", waitChan)
			transactionCache.Set("errChan", errChan)
			transactionno := data.TransactionNo
			commitfunc := func() error {
				svrList, err := getTransactionsvrList(transactionCache)
				if err != nil {
					return err
				}
				out := protocol.GET_MSG_COMMON_Transaction_Commit()
				out.No = transactionno
				for _, svr := range svrList {
					svr.SendMsg(svr.local, 0, 0, 0, out)
				}
				return nil
			}
			rollbackfunc := func() error {
				svrList, err := getTransactionsvrList(transactionCache)
				if err != nil {
					return err
				}
				out := protocol.GET_MSG_COMMON_Transaction_RollBack()
				out.No = transactionno
				for _, svr := range svrList {
					svr.SendMsg(svr.local, 0, 0, 0, out)
				}
				return nil
			}
			go func() {
				select {
				case option := <-waitChan:
					if option == TransactionOptionCommit {
						errChan <- commitfunc()
					} else {
						errChan <- rollbackfunc()
					}
				case <-time.After(protocol.MsgTimeOut * time.Second * 2):
					errChan <- rollbackfunc()
				}
			}()
		} else {
			transactionCache := cache.Hget(strconv.Itoa(int(data.TransactionNo)), TransactionCacheKey)
			transactionCache.Set(strconv.Itoa(int(in.Local)), true)
		}
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_COMMON_Transaction_Commit:
		transactionCache, ok := cache.Has(strconv.Itoa(int(data.No)), TransactionCacheKey)
		if !ok {
			in.WriteErr(errTransactionNo)
			return
		}
		//只有发起事务的服务器才允许进行commit
		if uint16(transactionCache.Load_int16("startSvr")) != in.Local {
			in.WriteErr(errTransactionNotAllowCommit)
			return
		}
		var waitChan chan transactionOption
		var errChan chan error
		if ok := transactionCache.Get("waitChan", &waitChan); !ok {
			in.WriteErr(errTransactionNotFoundWaitChan)
			return
		}
		if ok := transactionCache.Get("errChan", &errChan); !ok {
			in.WriteErr(errTransactionNotFoundErrChan)
			return
		}
		svrList, err := getTransactionsvrList(transactionCache)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_COMMON_Transaction_Check()
		out.No = data.No
		for _, svr := range svrList {
			if svr.local == in.Local {
				continue
			}
			err := svr.SendMsgWaitResult(svr.local, 0, 0, 0, out, nil)
			if err != nil {
				waitChan <- TransactionOptionRollback
				in.WriteErr(err)
				return
			}
		}
		out.Put()
		waitChan <- TransactionOptionCommit
		err = <-errChan
		in.WriteErr(err)
	case *protocol.MSG_COMMON_Transaction_RollBack:
		transactionCache, ok := cache.Has(strconv.Itoa(int(data.No)), TransactionCacheKey)
		if !ok {
			in.WriteErr(errTransactionNo)
			return
		}

		var waitChan chan transactionOption
		var errChan chan error
		if ok := transactionCache.Get("waitChan", &waitChan); !ok {
			in.WriteErr(errTransactionNotFoundWaitChan)
			return
		}
		if ok := transactionCache.Get("errChan", &errChan); !ok {
			in.WriteErr(errTransactionNotFoundErrChan)
			return
		}
		waitChan <- TransactionOptionRollback
		err := <-errChan
		in.WriteErr(err)

	default:
		if protocol.SetMsgQuery(i) {
			return //return掉，避免i.Put被回收
		} else {
			libraries.ReleaseLog("host未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
		}

	}
	i.Put()
})

func getTransactionsvrList(transactionCache *cache.Hashvalue) (svrList []*RpcServer, err error) {
	var id int
	transactionCache.Range(func(key string, _ interface{}) bool {
		if key == "startSvr" || key == "waitChan" || key == "errChan" {
			return true
		}
		id, err = strconv.Atoi(key)
		if err != nil {
			err = errors.New("commit出现错误，缓存id不是数字" + key)
			return false
		}
		if svr := rpcServerIdList[uint8(id)][uint8(id>>8)]; svr == nil {
			libraries.DebugLog("commit找不到id为" + key + "的服务器")
			err = errTransactionNotFoundSvr
			return false
		} else {
			svrList = append(svrList, svr)
		}
		return true
	})
	return
}
