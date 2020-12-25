package protocol

import (
	"errors"
	"fmt"
	"libraries"
	"reflect"
	"runtime/debug"
	"strconv"
	"sync/atomic"
	"time"
)

type MSG_DATA interface {
	Put()
	write(buf *libraries.MsgBuffer)
	cmd() int32
	read(buf *libraries.MsgBuffer)
}
type Msg struct {
	Msgno   uint32
	Ttl     uint8
	Local   uint16
	Remote  uint16
	Cmd     int32
	Data    MSG_DATA
	buf     *libraries.MsgBuffer
	datalen int
	svr     MsgServer
	cache   RpcCache
}

var (
	errMsgLen     = errors.New("消息长度不够解析一条msg")
	errMsgDataLen = errors.New("消息长度不够解析data")
)

func ReadOneMsg(buf *libraries.MsgBuffer) (msg *Msg, err error) {

	data := buf.Bytes()
	if len(data) < MsgHeadLen {
		return nil, errMsgLen
	}
	msg = &Msg{}
	msg.Msgno = uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
	msg.Ttl = data[4]
	msg.Local = uint16(data[5]) | uint16(data[6])<<8
	msg.Remote = uint16(data[7]) | uint16(data[8])<<8

	msg.datalen = int(data[9]) | int(data[10])<<8 | int(data[11])<<16
	if len(data) < MsgHeadLen+msg.datalen {
		return nil, errMsgDataLen
	}
	msg.Cmd = int32(data[12]) | int32(data[13])<<8 | int32(data[14])<<16 | int32(data[15])<<24
	//msg.Data = data[MsgHeadLen : MsgHeadLen+datalen]
	msg.buf = buf
	return msg, nil
}
func (m *Msg) ReadData() {
	if f, ok := cmdMapFunc[m.Cmd]; ok {
		m.buf.Next(MsgHeadLen)
		buf := BufPoolGet()
		buf.Write(m.buf.Next(m.datalen)[4:]) //跳过cmd部分
		m.Data = f(buf)
		buf.Reset()
		BufPoolPut(buf)
	}
	return
}
func (m *Msg) ReadDataWithCopy() {
	if f, ok := cmdMapFunc[m.Cmd]; ok {
		buf := BufPoolGet()
		buf.Write(m.buf.Bytes()[MsgHeadLen+4 : MsgHeadLen+m.datalen])
		m.Data = f(buf)
		BufPoolPut(buf)
	}
}
func (m *Msg) Next() []byte {
	return m.buf.Next(MsgHeadLen + m.datalen)
}

type MsgServer interface {
	SendMsg(remote uint16, msgno uint32, ttl uint8, out MSG_DATA)
	SendMsgWaitResult(remote uint16, msgno uint32, ttl uint8, out MSG_DATA, result interface{}, timeout ...time.Duration) (err error)
}

//实现msg读写,由msg发出的消息，继承msgno和ttl
func (m *Msg) SetServer(svr MsgServer) {
	m.svr = svr
	m.cache.Svr = svr
}

//指定服务器发送
func (m *Msg) SendMsg(remote uint16, out MSG_DATA) {
	m.svr.SendMsg(remote, m.Msgno, m.Ttl, out)
}
func (m *Msg) SendMsgWaitResult(remote uint16, out MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	return m.svr.SendMsgWaitResult(remote, m.Msgno, m.Ttl, out, result, timeout...)
}

//原路返回
func (m *Msg) SendResult(out MSG_DATA) {
	if in, ok := m.Data.(RpcQuery); ok {
		if outQuery, ok1 := out.(RpcQueryResult); ok1 {
			outQuery.setQueryResultID(in.getQueryID())
		}
	}
	m.svr.SendMsg(m.Local, m.Msgno, m.Ttl, out)
}

//原路返回err
func (m *Msg) WriteErr(err error) {
	data := GET_MSG_COMMON_QueryErr()
	if in, ok := m.Data.(RpcQuery); ok {
		data.QueryResultID = in.getQueryID()
	}
	if data.QueryResultID == 0 && err != nil {
		libraries.DebugLog("返回queryID=0的err", err)
	}
	if err != nil {
		data.Err = err.Error()
		data.Stack = debug.Stack()
	}
	m.svr.SendMsg(m.Local, m.Msgno, m.Ttl, data)
	data.Put()
}

func (m *Msg) Cache_Get(key string, value MSG_DATA) (err error) {
	b, err := m.cache.Get(key, "Msg:"+strconv.Itoa(int(m.Msgno)))
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
		if cmd == value.cmd() {
			if value == nil {
				r := reflect.ValueOf(value)
				r.Set(reflect.New(r.Type().Elem()))
			}
			value.read(buf)
		}
	}
	return err
}

func (m *Msg) Cache_Set(key string, value MSG_DATA) error {
	buf := BufPoolGet()
	value.write(buf)
	err := m.cache.Set(key, "Msg:"+strconv.Itoa(int(m.Msgno)), buf.Bytes(), MsgTimeOut*2)
	buf.Reset()
	BufPoolPut(buf)
	return err
}
func (m *Msg) Cache_Del(key string) error {
	return m.cache.Del(key, "Msg:"+strconv.Itoa(int(m.Msgno)))
}

//解决其他地方无法调用小写方法
func MSG_DATA_Write(data MSG_DATA, buf *libraries.MsgBuffer) {
	data.write(buf)
}

//result可以传入nil，但是返回非MSG_COMMON_QueryErr就会报错
func SendMsgWaitResult(local, remote uint16, msgno uint32, ttl uint8, out MSG_DATA, result interface{}, outchan chan *libraries.MsgBuffer, timeout ...time.Duration) (err error) {
	query, ok := out.(RpcQuery)
	if !ok {
		return errors.New(fmt.Sprintf("out结构体%s不含QueryID无法使用Result模式", reflect.TypeOf(out).Elem().Name()))
	}
	id := atomic.AddUint32(&RpcClientQueryId, 1)
	if id == 0 { //不允许为0
		id = atomic.AddUint32(&RpcClientQueryId, 1)
	}
	query.setQueryID(id)
	resultchan := make(chan RpcQueryResult, 1)
	RpcClientQueryMap[id] = resultchan
	buf := BufPoolGet()
	b := buf.Make(MsgHeadLen)
	b[0] = byte(msgno)
	b[1] = byte(msgno >> 8)
	b[2] = byte(msgno >> 16)
	b[3] = byte(msgno >> 24)
	b[4] = ttl
	b[5] = byte(local)
	b[6] = byte(local >> 8)
	b[7] = byte(remote)
	b[8] = byte(remote >> 8)
	out.write(buf)
	if buf.Len() > MaxMsgLen {
		libraries.ReleaseLog("消息发送失败，包体超过限制" + string(debug.Stack()))
		return errors.New("消息发送失败，包体超过限制" + string(debug.Stack()))
	} else {
		b = buf.Bytes()
		msglen := buf.Len() - MsgHeadLen
		b[9] = byte(msglen)
		b[10] = byte(msglen >> 8)
		b[11] = byte(msglen >> 16)
		outchan <- buf
	}
	_timeout := MsgTimeOut * time.Second
	if len(timeout) == 1 {
		_timeout = timeout[0]
	}
	checkAndSetResult := func(resultmsg RpcQueryResult) error {
		r1 := reflect.ValueOf(result)
		if r1.Kind() != reflect.Ptr {
			return RpcClientQueryResultErrType
		}
		r2 := reflect.ValueOf(resultmsg)
		r1 = r1.Elem()
		if r1.Type().Elem().Name() != r2.Elem().Type().Name() {
			return errors.New(fmt.Sprintf("实际返回的结果为%s,与请求的%s不相符", r2.Elem().Type().Name(), r1.Type().Elem().Name()))
		}
		r1.Set(r2)
		return nil
	}
	select {
	case r := <-resultchan:
		if err, ok := r.(*MSG_COMMON_QueryErr); ok {
			if err.Err != "" {
				return errors.New(err.Err)
			} else if result == nil {
				return nil
			} else {
				r1 := reflect.ValueOf(result)
				return errors.New(fmt.Sprintf("实际返回的结果为MSG_COMMON_QueryErr,与请求的%s不相符", r1.Elem().Elem().Type().Name()))
			}
		}
		return checkAndSetResult(r)
	case <-time.After(_timeout):
		RpcClientQueryLock.Lock()
		defer RpcClientQueryLock.Unlock()
		delete(RpcClientQueryMap, id)
		select {
		case r := <-resultchan:
			defer r.(MSG_DATA).Put()
			if err, ok := r.(*MSG_COMMON_QueryErr); ok {
				if err.Err != "" {
					return errors.New(err.Err)
				} else if result == nil {
					return nil
				} else {
					r1 := reflect.ValueOf(result)
					return errors.New(fmt.Sprintf("实际返回的结果为MSG_COMMON_QueryErr,与请求的%s不相符", r1.Elem().Elem().Type().Name()))
				}
			}
			return checkAndSetResult(r)
		default:
		}
	}
	return RpcClientQueryTimeOutErr

}
func SendMsg(local, remote uint16, msgno uint32, ttl uint8, out MSG_DATA, outchan chan *libraries.MsgBuffer) {
	buf := BufPoolGet()
	b := buf.Make(MsgHeadLen)
	b[0] = byte(msgno)
	b[1] = byte(msgno >> 8)
	b[2] = byte(msgno >> 16)
	b[3] = byte(msgno >> 24)
	b[4] = ttl
	b[5] = byte(local)
	b[6] = byte(local >> 8)
	b[7] = byte(remote)
	b[8] = byte(remote >> 8)
	out.write(buf)
	if buf.Len() > MaxMsgLen {
		libraries.ReleaseLog("消息发送失败，包体超过限制" + string(debug.Stack()))
	} else {
		b = buf.Bytes()
		msglen := buf.Len() - MsgHeadLen
		b[9] = byte(msglen)
		b[10] = byte(msglen >> 8)
		b[11] = byte(msglen >> 16)
		outchan <- buf
	}
}
