package rpcHost

import (
	"errors"
	"fmt"
	"jachunPM_commom/db"
	"libraries"
	"os"
	"protocol"
	"reflect"
	"runtime/debug"
	"server"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/luyu6056/cache"
	"github.com/luyu6056/gnet"
)

var filepath string

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
			buf.Write(in.Next())
			rpcServerOutChan[protocol.HostServerNo] <- buf
		}
	}
	if int(msgnum) != n {
		libraries.DebugLog("读消息数量错误，请检查协议，消息总量%d,已读%d", msgnum, n)
	}
	return nil
}

var msgnoTtl sync.Map

func HostServerHandler() {
	if libraries.IsRelease {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				debug.PrintStack()
				go HostServerHandler()
			}
		}()

	}
	for buf := range rpcServerOutChan[protocol.HostServerNo] {

		in, err := protocol.ReadOneMsg(buf)
		if err != nil {
			//上面读过一遍了，这里不应该出错
			libraries.DebugLog("host读消息出错,错误 %v", err)
			buf.Reset()
			protocol.BufPoolPut(buf)
			continue
		}
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
		func() {
			rpcLock.RLock()
			defer rpcLock.RUnlock()
			svr := rpcServerIdList[uint8(in.Local)][uint8(in.Local>>8)]
			if svr == nil {
				//可能服务掉线了,暂不处理
				libraries.DebugLog("host收到不存在的svr,No%d,Id%d", uint8(in.Local), uint8(in.Local>>8))
				return
			}
			in.ReadData()
			in.SetServer(svr)
			i := in.Data
			defer func() {
				i.Put()
				buf.Reset()
				protocol.BufPoolPut(buf)
			}()
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
				if svr != nil {
					r := cache.Hget(data.Name, data.Path)
					out := protocol.GET_MSG_COMMON_CACHE_GET_result()
					r.Get("value", &out.Value)
					out.QueryResultID = data.QueryID
					svr.SendMsg(svr.local, 0, 0, out)
					out.Put()

				}
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
				svr.SendMsg(svr.local, 0, 0, out)
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
					svr.SendMsg(svr.local, 0, 0, out)
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
			default:
				libraries.ReleaseLog("host未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
			}
		}()

	}

}
