package rpcHost

import (
	"errors"
	"fmt"
	"jachunPM_commom/db"
	"protocol"
	"strconv"
	"time"

	"github.com/luyu6056/cache"
)

type HostServer struct {
}

var Host HostServer

func (HostServer) SendMsg(msg *protocol.Msg, remote uint16, out protocol.MSG_DATA) {
	protocol.SendMsg(msg, protocol.HostServerNo, remote, out, rpcHostMsgInChan, rpcServerOutChan[protocol.HostServerNo])
}
func (HostServer) SendMsgWaitResult(msg *protocol.Msg, remote uint16, out protocol.MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	return protocol.SendMsgWaitResult(msg, protocol.HostServerNo, remote, out, result, rpcHostMsgInChan, rpcServerOutChan[protocol.HostServerNo], timeout...)
}

func GetMsg() *protocol.Msg {
	msg := &protocol.Msg{DB: &protocol.MsgDB{DB: db.DB}}
	msg.SetServer(Host)
	return msg
}
func (HostServer) GetUserCacheById(id int32) (user *protocol.MSG_USER_INFO_cache, err error) {
	r := cache.Hget(strconv.Itoa(int(id)), strconv.Itoa(protocol.UserServerNo)+"_"+protocol.PATH_USER_INFO_CACHE)
	var b []byte
	if ok := r.Get("value", &b); ok {
		if len(b) > 4 {
			buf := protocol.BufPoolGet()
			defer func() {
				if r := recover(); r != nil {
					err = errors.New("cache反序列化错误" + fmt.Sprint(r))
				}
				buf.Reset()
				protocol.BufPoolPut(buf)
			}()
			buf.Write(b[4:])
			user = protocol.READ_MSG_USER_INFO_cache(buf)
		} else {
			return nil, errors.New("获取userCache失败，消息不够长，不足以读取一条缓存")
		}
	}
	return
}
