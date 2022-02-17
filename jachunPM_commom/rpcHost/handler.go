package rpcHost

import (
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"jachunPM_commom/db"
	"libraries"
	"os"
	"os/exec"
	_filepath "path/filepath"
	"protocol"
	"reflect"
	"runtime"
	"runtime/debug"
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
	fileTmpPath                    string //清除过期文件代码在rpcserver.tick()
	errTransactionNo               = errors.New("TransactionNotFoundNo")
	errTransactionNotAllowCommit   = errors.New("TransactionNotAllowCommit")
	errTransactionNotFoundSvr      = errors.New("TransactionNotFoundServer")
	errTransactionNotFoundWaitChan = errors.New("TransactionNotFoundWaitChan")
	errTransactionNotFoundErrChan  = errors.New("TransactionNotFoundErrChan")
	hostGoPool, _                  = ants.NewPool(10000)
	rpcHostMsgInChan               = make(chan *protocol.Msg, protocol.Rpcmsgnum)
	tmpFileId                      int64
	fileLock                       sync.Mutex
)

func init() {
	//删除所有缓存上的事务，如果common重启，事务将变得不可信
	cache.Hdel_all(TransactionCacheKey)
	var err error
	filepath, err = libraries.GetBaseRootPath()
	if err != nil {
		panic("获取运行根目录失败" + err.Error())
	}

	if runtime.GOOS == "windows" {
		fileTmpPath = filepath + `\tmp\`
		filepath += `\upload\`
	} else {
		fileTmpPath = filepath + "/tmp/"
		filepath += "/upload/"
	}
	if ok, _ := libraries.PathExists(filepath); !ok {
		os.Mkdir(filepath, 0777)
	}
	if ok, _ := libraries.PathExists(fileTmpPath); !ok {
		os.Mkdir(fileTmpPath, 0777)
	}
	go deleteTempFile()
	go deleteMsgno()
}
func deleteTempFile() {
	defer func() {
		if err := recover(); err != nil {
			libraries.ReleaseLog("deleteTempFile 发生错误 %v", err)
		}
		time.Sleep(time.Second)
		go deleteTempFile()
	}()
	for {
		list, err := libraries.ListDirAll(fileTmpPath, "")
		if err != nil {
			libraries.ReleaseLog("deleteTempFile 获取目录内容错误 %v", err)
		}
		now := time.Now()
		for _, file := range list {
			s, err := os.Stat(file)
			if err != nil {
				libraries.ReleaseLog("deleteTempFile 获取文件%s错误 %v", file, err)
			}
			//临时文件保存7天
			if now.Unix()-s.ModTime().Unix() > 86400*7 {
				os.Remove(file)
			}
		}
		time.Sleep(time.Minute)
	}
}
func deleteMsgno() {
	defer func() {
		if err := recover(); err != nil {
			libraries.ReleaseLog("deleteMsgno 发生错误 %v", err)
		}
		time.Sleep(time.Second)
		go deleteMsgno()
	}()
	for {
		now := time.Now()
		msgnoTtl.Range(func(k, v interface{}) bool {
			logs, ok := v.([]*db.Log_msg)
			if !ok || len(logs) == 0 {
				msgnoTtl.Delete(k)
			}
			last := logs[0]
			if now.Unix()-last.Timestamp.Unix() > int64(last.TimeOut) {
				msgnoTtl.Delete(k)
			}
			return true
		})
		time.Sleep(time.Minute)
	}
}
func HandlerMsg(data []byte, c gnet.Conn) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	index := 0
	for index < len(data) {
		in, l, err := protocol.ReadOneMsgFromBytes(data[index:])
		index += l
		if err != nil {
			return errors.New("读消息出错" + err.Error())
		} else {
			db.ToZstd(data[index-l:index], in.Cmd)
			in.Addr = c.RemoteAddr().String()
			context := c.Context().(*RpcServer)
			in.SetServer(context)
			if context.Id != -1 {
				if in.Local != uint16(context.ServerNo)|uint16(context.Id)<<8 { //检查local
					return errors.New(fmt.Sprintf("消息的local来源不对,in.Local%d,local%d", in.Local, uint16(context.ServerNo)|uint16(context.Id)<<8))
				}
				if byte(in.GetRemoteID()) == protocol.HostServerNo || (in.GetRemoteID() == 0 && byte(in.Cmd) == protocol.HostServerNo) {
					if !checkTTL(in) {
						continue
					}
					in.ReadData()
					in.DB.DB = db.DB
					hostAsyncHand.Invoke(in)
					//hostAsyncHand是异步执行，如果有需要同步的需要在这里进行处理
					protocol.BufPoolPut(in.Buf())
				} else {

					rpcHostMsgInChan <- in
				}
			} else {
				if in.Cmd == protocol.CMD_MSG_HOST_regServer {
					in.ReadData()
					data, ok := in.Data.(*protocol.MSG_HOST_regServer)
					if ok {
						go func() {
							context.Start(data.No, data.IpPort, data.Window, in)
							//拉取注册前的消息，推入队列
							for i := len(context.inChan); i > 0; i-- {
								in := <-context.inChan
								//in.Local=uint16(context.ServerNo)|uint16(context.Id)<<8
								if byte(in.GetRemoteID()) == protocol.HostServerNo || (in.GetRemoteID() == 0 && byte(in.Cmd) == protocol.HostServerNo) {
									in.ReadData()
									in.DB.DB = db.DB
									hostAsyncHand.Invoke(in)
									protocol.BufPoolPut(in.Buf())
								} else {
									rpcHostMsgInChan <- in
								}

							}
						}()

					} else {
						libraries.ReleaseLog("协议错误，MSG_HOST_regServer读取错误")
					}

				} else {
					//消息推入chan，等注册后再执行
					select {
					case context.inChan <- in:
					default:
						c.Close()
						libraries.ReleaseLog("%s服务未注册，缓存已满", c.RemoteAddr().String())
						return nil
					}
				}

			}
		}
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
		in.DB.DB = db.DB
		cmdSvrNo := byte(in.Cmd)
		//检查msgno

		if !checkTTL(in) {
			continue
		}

		if in.GetRemoteID() == 0 {
			//libraries.DebugLog("from %d cmd %s ,to %d", in.Local, protocol.CmdToName[in.Cmd], cmdSvrNo)
			rpcServerOutChan[cmdSvrNo] <- buf
		} else {
			svrNo := byte(in.GetRemoteID())
			id := byte(in.GetRemoteID() >> 8)
			rpcLock.RLock()
			if v, ok := rpcServerIdList[svrNo]; ok {
				if remoteSvr := v[id]; remoteSvr != nil {
					remoteSvr.outChan <- buf
				} else {
					rpcServerOutChan[svrNo] <- buf
				}
			}
			rpcLock.RUnlock()
		}

	}

}

var hostAsyncHand, _ = ants.NewPoolWithFunc(10000, func(args interface{}) {
	in, ok := args.(*protocol.Msg)
	if !ok {
		return
	}
	var svr *RpcServer
	//if in.Local == protocol.HostServerNo {
	//	in.SetServer(Host)
	//} else {
	svr, _ = in.Svr.(*RpcServer)
	//}

	i := in.Data
	//defer i.Put()  SetMsgQuery里面不能回收，所以不能defer回收

	switch data := i.(type) {
	case *protocol.MSG_HOST_WINDOW_UPDATE:
		if svr != nil {
			atomic.AddInt32(&svr.window, data.Add)
			//尝试把半开的服务恢复正常
			//libraries.DebugLog("服务%v，ID%v，更新窗口%d", svr.ServerNo, svr.Id, svr.window)
			select {
			case svr.setStatusOpenChan <- "更新窗口值":
			default:
			}
		}

	case *protocol.MSG_HOST_PONG:
		if svr != nil {
			//libraries.DebugLog("服务%d,ID%d,收到pong,%d", svr.ServerNo, svr.Id, data.Rand)
			svr.pongTime = time.Now().Unix()
			if svr.window > 0 {
				//尝试把pong响应超时导致半开的服务恢复正常
				select {
				case svr.setStatusOpenChan <- "收到pong":
				default:
				}
			}
		}
	case *protocol.MSG_HOST_CACHE_DEL:
		cache.Hdel(data.Name, data.Path)
	case *protocol.MSG_HOST_CACHE_DelPath:
		cache.Hdel_all(data.Path)
	case *protocol.MSG_HOST_CACHE_SET:
		//libraries.DebugLog("Path:%s Name:%s Value:%s", data.Path, data.Name, libraries.MD5_B(data.Value))
		cache.Hset(data.Name, map[string][]byte{"value": data.Value}, data.Path, data.Expire)
	case *protocol.MSG_HOST_CACHE_GET:

		r := cache.Hget(data.Name, data.Path)
		out := protocol.GET_MSG_HOST_CACHE_GET_result()
		r.Get("value", &out.Value)
		//libraries.DebugLog("Path:%s Name:%s Value:%s",data.Path,data.Name,libraries.MD5_B(out.Value))
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_HOST_CACHE_GETPATH:
		out := protocol.GET_MSG_HOST_CACHE_GETPATH_result()
		if len(data.Names) == 0 {
			cache.RangePath(data.Path, func(key string, v *cache.Hashvalue) bool {
				var value []byte
				if v.Get("value", &value) {
					out.Value = append(out.Value, value)
				}
				return true
			})
		} else {
			cache.RangePath(data.Path, func(key string, v *cache.Hashvalue) bool {

				for _, name := range data.Names {
					if key == name {
						var value []byte
						v.Get("value", &value)
						out.Value = append(out.Value, value)
					}
				}
				return true
			})
		}

		in.SendResult(out)
		out.Put()
	case *protocol.MSG_HOST_GET_Msgno:
		libraries.DebugLog("还有GET的")
	case *protocol.MSG_HOST_regServer:
		//common掉线可能会导致其他服务反复发送reg
		if data.No != svr.ServerNo {
			libraries.DebugLog("注册的serverNo不对，注册%d,实际%d", data.No, svr.ServerNo)
		}
	case *protocol.MSG_HOST_ResetWindow:
		libraries.DebugLog("%d 重置 %d", svr.ServerNo, data.Window)
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

		var file = &db.File{
			Pathname:   dir + data.Name,
			Title:      data.Name,
			Extension:  data.Name[strings.LastIndex(data.Name, ".")+1:],
			Size:       int64(len(data.Data)),
			ObjectType: data.ObjectType,
			ObjectID:   data.ObjectID,
			AddedBy:    data.AddBy,
			AddedDate:  now,
			Type:       data.Type,
		}
		var id int64
		id, err = in.DB.Table(db.TABLE_FILE).Insert(file)
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
		err := in.DB.Table(db.TABLE_FILE).Prepare().Where("Id=?", data.FileID).Find(&file)
		if err != nil {
			in.WriteErr(err)
			return
		}
		os.Remove(filepath + file.Pathname)
		_, err = in.DB.Table(db.TABLE_FILE).Where("Id=?", data.FileID).Delete()
		in.WriteErr(err)
	case *protocol.MSG_FILE_getByID:

		out := protocol.GET_MSG_FILE_getByID_result()
		if data.FileID > 0 {
			var file *db.File
			err := in.DB.Table(db.TABLE_FILE).Prepare().Where("Id=?", data.FileID).Find(&file)
			if file == nil && err == nil {
				err = errors.New(protocol.Err_FileNotFound.String())
			}
			if err != nil {
				in.WriteErr(err)
				return
			}
			out.Ext = file.Extension
			out.Name = file.Title
			out.Type = file.Type
			out.Size = file.Size
			out.FileID = file.Id
			out.AddedDate = file.AddedDate
			out.ObjectType = file.ObjectType
			out.ObjectID = file.ObjectID
		} else {
			name := strconv.FormatInt(data.FileID, 10) + ".tar"
			f, err := os.Stat(fileTmpPath + name)
			if err != nil {
				in.WriteErr(errors.New(protocol.Err_FileNotFound.String()))
				return
			}
			out.FileID = data.FileID
			out.Size = f.Size()
			out.Name = name
		}
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_FILE_getByObject:
		var files []*db.File
		err := in.DB.Table(db.TABLE_FILE).Prepare().Where("ObjectType=? and ObjectID=?", data.ObjectType, data.ObjectID).Limit(0).Select(&files)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_FILE_getByObject_result()
		for _, file := range files {
			/*b, err := ioutil.ReadFile(filepath + file.Pathname)
			if err != nil {
				err = errors.New(protocol.Err_FileNotFound.String() + " err:" + err.Error())
				in.WriteErr(err)
				return
			}*/
			tmp := protocol.GET_MSG_FILE_getByID_result()
			//tmp.Data = b
			tmp.Ext = file.Extension
			tmp.Name = file.Title
			tmp.Type = file.Type
			tmp.FileID = file.Id
			tmp.Size = file.Size
			tmp.AddedDate = file.AddedDate
			tmp.ObjectType = file.ObjectType
			tmp.ObjectID = file.ObjectID
			out.List = append(out.List, tmp)
		}
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_FILE_RangeDown:
		if data.End <= data.Start {
			in.WriteErr(errors.New(protocol.Err_FileDownLoadStartEnd.String()))
			return
		}
		var f *os.File
		var err error
		if data.FileID > 0 {
			var file *db.File
			err = in.DB.Table(db.TABLE_FILE).Prepare().Where("Id=?", data.FileID).Find(&file)
			if file == nil && err == nil {
				err = errors.New(protocol.Err_FileNotFound.String())
			}
			if err != nil {
				in.WriteErr(err)
				return
			}
			f, err = getFile(file)
			if err != nil {
				in.WriteErr(err)
				return
			}
		} else {
			name := strconv.FormatInt(data.FileID, 10) + ".tar"
			f, err = os.Open(fileTmpPath + name)
			if err != nil {
				in.WriteErr(err)
				return
			}
		}

		defer f.Close()
		f.Seek(data.Start, 0)
		out := protocol.GET_MSG_FILE_RangeDown_result()
		out.Byte = make([]byte, data.End-data.Start)
		i, err := f.Read(out.Byte)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out.Byte = out.Byte[:i]
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_FILE_updateMapByWhere:
		_, err := in.DB.Table(db.TABLE_FILE).Where(data.Where).Update(data.Update)
		in.WriteErr(err)
	case *protocol.MSG_HOST_BeginTransaction:
		//发起事务的服务器
		startSvr := data.TransactionNo == 0
		for data.TransactionNo == 0 {
			data.TransactionNo = uint32(transactionNoCache.INCRBY("NO", 1))
		}
		out := protocol.GET_MSG_HOST_BeginTransaction_result()
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
				cache.Hdel(strconv.Itoa(int(transactionno)), TransactionCacheKey) //删掉以避免后续func带上事务
				out := protocol.GET_MSG_HOST_Transaction_Commit()
				out.No = transactionno
				for _, svr := range svrList {
					if svr.local == protocol.HostServerNo {
						protocol.MsgTransactionCommit(out)
					} else {
						svr.SendMsg(in, svr.local, out)
					}

				}
				return nil
			}
			rollbackfunc := func() error {
				svrList, err := getTransactionsvrList(transactionCache)
				if err != nil {
					return err
				}
				cache.Hdel(strconv.Itoa(int(transactionno)), TransactionCacheKey) //删掉以避免后续func带上事务
				out := protocol.GET_MSG_HOST_Transaction_RollBack()
				out.No = transactionno
				for _, svr := range svrList {
					if svr.local == protocol.HostServerNo {
						protocol.MsgTransactionRollBack(out)
					} else {
						svr.SendMsg(in, svr.local, out)
					}

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
	case *protocol.MSG_HOST_Transaction_Commit:
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
		out := protocol.GET_MSG_HOST_Transaction_Check()
		out.No = data.No
		for _, svr := range svrList {
			if svr.local == in.Local {
				continue
			}
			var err error

			if svr.local == protocol.HostServerNo {
				err = protocol.MsgTransactionCheck(out)
			} else {
				err = svr.SendMsgWaitResult(nil, svr.local, out, nil)
			}

			if err != nil {
				libraries.DebugLog("%d,%+v,%+v", svr.local, svr, err)
				waitChan <- TransactionOptionRollback
				in.WriteErr(err)
				return
			}
		}
		out.Put()

		waitChan <- TransactionOptionCommit
		err = <-errChan
		in.WriteErr(err)
	case *protocol.MSG_HOST_Transaction_RollBack:

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
	case *protocol.MSG_FILE_uploadTmp:
		if data.Index == 0 {
			os.Remove(fileTmpPath + data.Name)
		}
		f, err := os.OpenFile(fileTmpPath+data.Name, os.O_CREATE|os.O_RDWR, 0777)
		if err == nil {
			offset := int64(data.Index * data.BlockSize)
			if stat, err := f.Stat(); err != nil {
				in.WriteErr(err)
				f.Close()
				return
			} else if stat.Size() < offset {
				in.WriteErr(errors.New("Error File Create"))
				f.Close()
				return
			}
			f.Seek(offset, 0)
			f.Write(data.Data)
			f.Sync()
			f.Close()
		}
		in.WriteErr(err)
	case *protocol.MSG_FILE_updateTmp:
		var insertList []*db.File
		now := time.Now()
		for _, file := range data.Files {
			insert := &db.File{
				Title:      file.Title,
				ObjectType: file.ObjectType,
				ObjectID:   file.ObjectID,
				AddedBy:    file.AddBy,
				AddedDate:  now,
				Type:       file.Type,
			}
			if stat, err := os.Stat(fileTmpPath + file.Name); err != nil {
				in.WriteErr(err)
				return
			} else {
				insert.Size = stat.Size()
			}
			insert.Pathname = file.Code + "/" + file.Type + now.Format("/2006/01/02/") + file.Name

			newpath := filepath + insert.Pathname
			if runtime.GOOS == "windows" {
				newpath = strings.ReplaceAll(newpath, "/", "\\")
			}
			dir := _filepath.Dir(newpath)
			_, err := os.Stat(dir)
			if err != nil {
				os.MkdirAll(dir, 0777)
			}
			if err := os.Rename(fileTmpPath+file.Name, newpath); err != nil {
				in.WriteErr(err)
				for _, i := range insertList {
					if runtime.GOOS == "windows" {
						os.Remove(filepath + strings.ReplaceAll(i.Pathname, "/", "\\"))
					} else {
						os.Remove(filepath + i.Pathname)
					}

				}
				return
			}
			insertList = append(insertList, insert)
		}
		var err error
		if len(insertList) > 0 {
			_, err = in.DB.Table(db.TABLE_FILE).InsertAll(insertList)
			if err != nil {
				for _, i := range insertList {
					os.Remove(filepath + i.Pathname)
				}
			}
		}
		in.WriteErr(err)
	case *protocol.MSG_FILE_edit:
		var file *db.File
		if err := in.DB.Table(db.TABLE_FILE).Prepare().Where("Id=?", data.FileID).Find(&file); err != nil {
			in.WriteErr(err)
			return
		}
		if file.Title != data.Name {
			if _, err := in.DB.Table(db.TABLE_FILE).Prepare().Where("Id=?", data.FileID).Update("Title=?", data.Name); err != nil {
				in.WriteErr(err)
				return
			}
			actionID, _ := in.ActionCreate(file.ObjectType, file.ObjectID, "editfile", "", data.Name, nil, 0)
			if actionID > 0 {
				var change protocol.ChangeHistory = []*protocol.MSG_LOG_History{&protocol.MSG_LOG_History{
					Field: "fileName",
					Old:   file.Title,
					New:   data.Name,
				}}
				change.Add(actionID, in)
			}
		}
		in.WriteErr(nil)
	case *protocol.MSG_FILE_getByWhere:
		out := protocol.GET_MSG_FILE_getByWhere_result()
		var files []*db.File
		var err error
		if err = in.DB.Table(db.TABLE_FILE).Where(data.Where).Limit((data.Page-1)*data.PerPage, data.PerPage).Select(&files); err != nil {
			in.WriteErr(err)
			return
		}
		if data.Total == 0 {
			if out.Total, err = in.DB.Table(db.TABLE_FILE).Where(data.Where).Count(); err != nil {
				in.WriteErr(err)
				return
			}
		}
		for _, file := range files {
			tmp := protocol.GET_MSG_FILE_getByID_result()
			tmp.FileID = file.Id
			tmp.Size = file.Size
			tmp.Type = file.Type
			tmp.Name = file.Title
			tmp.ObjectID = file.ObjectID
			tmp.ObjectType = file.ObjectType
			tmp.AddedDate = file.AddedDate
			tmp.Ext = file.Extension
			out.List = append(out.List, tmp)
		}
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_FILE_download_byIds:
		var files []*db.File
		var err error
		if err = in.DB.Table(db.TABLE_FILE).Where(map[string]interface{}{"Id": data.Ids}).Limit(0).Select(&files); err != nil {
			in.WriteErr(err)
			return
		}
		//生成一个临时id
		var id int64
		for {
			id = atomic.AddInt64(&tmpFileId, -1)
			if id > 0 {
				id = -1
				tmpFileId = -1
			}
			break
		}
		d, _ := os.Create(fileTmpPath + strconv.FormatInt(id, 10) + ".tar")
		defer d.Close()
		tw := tar.NewWriter(d)
		defer tw.Close()

		for _, file := range files {
			f, err := getFile(file)
			if err != nil {
				in.WriteErr(err)
				return
			}
			err = compress(f, tw)
			if err != nil {
				in.WriteErr(err)
				return
			}
		}

		out := protocol.GET_MSG_FILE_download_byIds_result()
		out.FileID = id
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_HOST_getCenterSvrId:
		out := protocol.GET_MSG_HOST_getCenterSvrId_result()
		rpcLock.RLock()
		if v, ok := rpcServerIdList[data.No]; ok {
			for _, s := range v {
				if s.isCenter {
					out.Id = uint16(svr.ServerNo) + uint16(s.Id)<<8
				}
			}
		}
		defer rpcLock.RUnlock()
		in.SendResult(out)
		out.Put()
	default:
		if protocol.SetMsgQuery(in) {
			return
		}
		libraries.ReleaseLog("host未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
	}
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
		if int(uint8(id)) >= len(rpcServerIdList) || int(uint8(id>>8)) >= len(rpcServerIdList[uint8(id)]) {
			libraries.DebugLog("commit找不到id为" + key + "的服务器")
			err = errTransactionNotFoundSvr
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
func getFile(file *db.File) (*os.File, error) {
	path := filepath + file.Pathname[:strings.LastIndex(file.Pathname, "/")]
	f, err := os.Open(filepath + file.Pathname)
	if err != nil {
		e, ok := func() (error, bool) {
			if runtime.GOOS == "windows" {
				//尝试从svn远程拉文件
				path = strings.Replace(path, "/", `\`, -1)

				isExist, e := libraries.PathExists(path)
				if e != nil {
					return e, false
				}
				if !isExist {
					c := "d: && cd " + strings.Replace(filepath, "/", `\`, -1) + " && svn co --depth=empty http://192.168.6.99/svn/project/upload/1/" + file.Pathname[:strings.LastIndex(file.Pathname, "/")] + " " + file.Pathname[:strings.LastIndex(file.Pathname, "/")]
					cmd := exec.Command("cmd.exe", "/c", c)
					cmd.Start()
					cmd.Wait()
				}
				c := "d: && cd " + path + "\\ && svn up " + file.Title
				cmd := exec.Command("cmd.exe", "/c", c)
				cmd.Start()
				cmd.Wait()
				return nil, true
			}
			return nil, false
		}()
		if !ok {
			if e != nil {
				err = errors.New(err.Error() + " & " + e.Error())
			}
			err = errors.New(protocol.Err_FileNotFound.String() + " err:" + err.Error())
			return nil, err
		}
		f, err = os.Open(path + string(os.PathSeparator) + file.Title)
		if err != nil {
			return nil, err
		}
	}
	return f, nil

}

func compress(file *os.File, tw *tar.Writer) error {

	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {

		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = header.Name
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func checkTTL(in *protocol.Msg) bool {
	b := in.Buf().Bytes()
	//检查事务
	if transactionNo := int(b[10]) | int(b[11])<<8 | int(b[12])<<16 | int(b[13])<<24; transactionNo > 0 {
		if _, has := cache.Has(strconv.Itoa(transactionNo), TransactionCacheKey); !has {
			b[10], b[11], b[12], b[13] = 0, 0, 0, 0
		}
	}
	if in.Msgno == 0 {
		msgno := atomic.AddUint32(&globalMsgno, 1)
		in.Msgno = msgno
		b[0] = byte(in.Msgno)
		b[1] = byte(in.Msgno >> 8)
		b[2] = byte(in.Msgno >> 16)
		b[3] = byte(in.Msgno >> 24)
		msgnoTtl.Store(in.Msgno, db.MsgtoLog(in, nil))
	} else {
		if v, ok := msgnoTtl.Load(in.Msgno); ok {
			if _, ttlNoCheck := protocol.CMD_NO_CHECK_TTL[in.Cmd]; !ttlNoCheck {
				msgnoTtl.Store(in.Msgno, db.MsgtoLog(in, v.([]*db.Log_msg)))
				var lastTTl int
				if len(v.([]*db.Log_msg)) > 0 {
					lastTTl = v.([]*db.Log_msg)[len(v.([]*db.Log_msg))-1].Ttl
				}
				if lastTTl >= protocol.MaxMsgTtl {
					libraries.ReleaseLog("ttl过大,local %d remoted %d cmd %s msgno %d", in.Local, in.GetRemoteID(), protocol.CmdToName[in.Cmd], in.Msgno)
					//抛弃消息
					protocol.BufPoolPut(in.Buf())
					return false
				}
			}

		} else {
			libraries.DebugLog("无效的msgno %d,%s %+v", in.Msgno, protocol.CmdToName[in.Cmd], in)
			protocol.BufPoolPut(in.Buf())
			return false //抛弃消息
		}
	}
	return true
}
