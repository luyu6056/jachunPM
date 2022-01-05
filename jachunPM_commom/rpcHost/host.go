package rpcHost

import (
	"errors"
	"fmt"
	"protocol"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/luyu6056/cache"
)

type HostServer struct {
}

var Host HostServer

func (HostServer) SendMsg(remote uint16, msgno uint32, ttl uint16, transactionNo, queryID uint32, out protocol.MSG_DATA) {
	protocol.SendMsg(protocol.HostServerNo, remote, msgno, ttl, transactionNo, queryID, out, rpcServerOutChan[protocol.HostServerNo])
}
func (HostServer) SendMsgWaitResult(remote uint16, msgno uint32, ttl uint16, transactionNo uint32, out protocol.MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	return protocol.SendMsgWaitResult(protocol.HostServerNo, remote, msgno, ttl, transactionNo, out, result, rpcServerOutChan[protocol.HostServerNo], timeout...)
}

func GetOneMsg() *protocol.Msg {
	m := &protocol.Msg{}
	msgno := atomic.AddUint32(&globalMsgno, 1)
	m.Msgno = msgno
	ttl := int32(0)
	msgnoTtl.Store(m.Msgno, &ttl)
	time.AfterFunc(protocol.MsgTimeOut*time.Second, func() { msgnoTtl.Delete(msgno) })
	m.SetServer(Host)
	return m
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
			return nil, errors.New("消息不够长，不足以读取一条缓存")
		}
	}
	return
}
