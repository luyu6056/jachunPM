package protocol

import (
	"errors"
	"fmt"
	"libraries"
	"mysql"
	"reflect"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

type MSG_DATA interface {
	Put()
	write(buf *libraries.MsgBuffer)
	cmd() int32
	read(buf *libraries.MsgBuffer)
}
type Msg struct {
	Msgno              uint32
	Ttl                uint8
	Local              uint16
	Remote             uint16
	Cmd                int32
	Data               MSG_DATA
	buf                *libraries.MsgBuffer
	datalen            int
	svr                MsgServer
	cache              RpcCache
	transactionTimeOut time.Duration
	DB                 *MsgDB
}
type MsgDB struct {
	msg           *Msg
	db            *mysql.MysqlDB
	transaction   *MsgDBTransaction
	transactionNo uint32
}
type MsgDBTransaction struct {
	t                *mysql.Transaction
	msg              *Msg
	newTransactionNo uint32
}

var (
	errMsgLen              = errors.New("消息长度不够解析一条msg")
	errMsgDataLen          = errors.New("消息长度不够解析data")
	errTransactionTimeOut  = errors.New("当前消息开启了事务，消息超时不能大于默认时长")
	transactionMap         sync.Map
	errTransactionNotFound = errors.New("DB NotFoundTransaction")
)

func ReadOneMsg(buf *libraries.MsgBuffer) (msg *Msg, err error) {

	data := buf.Bytes()
	if len(data) < MsgHeadLen {
		return nil, errMsgLen
	}
	msg = &Msg{DB: &MsgDB{}}
	msg.Msgno = uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
	msg.Ttl = data[4] & 127
	msg.Local = uint16(data[5]) | uint16(data[6])<<8
	msg.Remote = uint16(data[7]) | uint16(data[8])<<8
	msg.DB.transactionNo = uint32(data[9]) | uint32(data[10])<<8 | uint32(data[11])<<16 | uint32(data[12])<<24
	msg.DB.msg = msg
	msg.datalen = int(data[13]) | int(data[14])<<8 | int(data[15])<<16
	if len(data) < MsgHeadLen+msg.datalen {
		return nil, errMsgDataLen
	}
	msg.Cmd = int32(data[16]) | int32(data[17])<<8 | int32(data[18])<<16 | int32(data[19])<<24
	//msg.Data = data[MsgHeadLen : MsgHeadLen+datalen]
	msg.buf = buf
	return msg, nil
}
func (m *Msg) ReadData() {
	if f, ok := cmdMapFunc[m.Cmd]; ok {
		/*屏蔽掉是安全拷贝读取
		m.buf.Next(MsgHeadLen)
		buf := BufPoolGet()
		buf.Write(m.buf.Next(m.datalen)[4:]) //跳过cmd部分
		m.Data = f(buf)
		buf.Reset()
		BufPoolPut(buf)
		先用高效率方案*/

		//高效处理方案，缺点一旦读错，将导致整个buf流失效
		m.buf.Next(MsgHeadLen + 4)
		m.Data = f(m.buf)
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
func (m *Msg) NextWithSetInBuf(buf *libraries.MsgBuffer) {
	buf.Write(m.buf.Next(MsgHeadLen + m.datalen))
	m.buf = buf
}
func (m *Msg) Buf() *libraries.MsgBuffer {
	return m.buf
}

type MsgServer interface {
	SendMsg(remote uint16, msgno uint32, ttl uint8, transactionNo uint32, out MSG_DATA)
	SendMsgWaitResult(remote uint16, msgno uint32, ttl uint8, transactionNo uint32, out MSG_DATA, result interface{}, timeout ...time.Duration) (err error)
}

//实现msg读写,由msg发出的消息，继承msgno和ttl
func (m *Msg) SetServer(svr MsgServer) {
	m.svr = svr
	m.cache.Svr = svr
}

//指定服务器发送
func (m *Msg) SendMsg(remote uint16, out MSG_DATA) {
	transactionNo := m.DB.transactionNo
	if transactionNo == 0 && m.DB.transaction != nil {
		transactionNo = m.DB.transaction.newTransactionNo
	}
	m.svr.SendMsg(remote, m.Msgno, m.Ttl, transactionNo, out)
}
func (m *Msg) SendMsgWaitResult(remote uint16, out MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	if m.DB.transaction != nil && timeout != nil && timeout[0] > MsgTimeOut*time.Second { //事务有默认超时，不允许timeout大于事务超时
		return errTransactionTimeOut
	}
	transactionNo := m.DB.transactionNo
	if transactionNo == 0 && m.DB.transaction != nil {
		transactionNo = m.DB.transaction.newTransactionNo
	}
	return m.svr.SendMsgWaitResult(remote, m.Msgno, m.Ttl, transactionNo, out, result, timeout...)
}

//原路返回
func (m *Msg) SendResult(out MSG_DATA) {
	if in, ok := m.Data.(RpcQuery); ok {
		if outQuery, ok1 := out.(RpcQueryResult); ok1 {
			outQuery.setQueryResultID(in.getQueryID())
		}
	}
	transactionNo := m.DB.transactionNo
	if transactionNo == 0 && m.DB.transaction != nil {
		transactionNo = m.DB.transaction.newTransactionNo
	}
	m.svr.SendMsg(m.Local, m.Msgno, m.Ttl, transactionNo, out)
}

//原路返回err
func (m *Msg) WriteErr(err error) {
	data := GET_MSG_COMMON_QueryErr()
	if in, ok := m.Data.(RpcQuery); ok {
		data.QueryResultID = in.getQueryID()
	}
	if data.QueryResultID == 0 && err != nil {
		libraries.DebugLog("返回queryID=0的err%v", err)
	}
	if err != nil {
		data.Err = err.Error()
		data.Stack = debug.Stack()
	}
	transactionNo := m.DB.transactionNo
	if transactionNo == 0 && m.DB.transaction != nil {
		transactionNo = m.DB.transaction.newTransactionNo
	}
	m.svr.SendMsg(m.Local, m.Msgno, m.Ttl, transactionNo, data)
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
func SendMsgWaitResult(local, remote uint16, msgno uint32, ttl uint8, transactionNo uint32, out MSG_DATA, result interface{}, outchan chan *libraries.MsgBuffer, timeout ...time.Duration) (err error) {
	query, ok := out.(RpcQuery)
	if !ok {
		return errors.New(fmt.Sprintf("out结构体%s不含QueryID无法使用Result模式", reflect.TypeOf(out).Elem().Name()))
	}
	id := uint32(rpcClientQueryId.INCRBY("id", 1))
	if id == 0 { //不允许为0
		id = uint32(rpcClientQueryId.INCRBY("id", 1))
	}
	query.setQueryID(id)
	resultchan := make(chan RpcQueryResult, 1)
	rpcClientQueryMap[id] = resultchan
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
	b[9] = byte(transactionNo)
	b[10] = byte(transactionNo >> 8)
	b[11] = byte(transactionNo >> 16)
	b[12] = byte(transactionNo >> 24)
	out.write(buf)
	if buf.Len() > MaxMsgLen {
		libraries.ReleaseLog("消息发送失败，包体超过限制" + string(debug.Stack()))
		return errors.New("消息发送失败，包体超过限制" + string(debug.Stack()))
	} else {
		b = buf.Bytes()
		msglen := buf.Len() - MsgHeadLen
		b[13] = byte(msglen)
		b[14] = byte(msglen >> 8)
		b[15] = byte(msglen >> 16)
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
		rpcClientQueryLock.Lock()
		defer rpcClientQueryLock.Unlock()
		delete(rpcClientQueryMap, id)
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
func SetMsgQuery(i interface{}) bool {
	if rpcResult, ok := i.(RpcQueryResult); ok {
		rpcClientQueryLock.RLock()
		if v, ok := rpcClientQueryMap[rpcResult.getQueryResultID()]; ok {
			v <- rpcResult
		}
		rpcClientQueryLock.RUnlock()
		return true
	}
	return false
}
func SendMsg(local, remote uint16, msgno uint32, ttl uint8, transactionNo uint32, out MSG_DATA, outchan chan *libraries.MsgBuffer) {
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
	b[9] = byte(transactionNo)
	b[10] = byte(transactionNo >> 8)
	b[11] = byte(transactionNo >> 16)
	b[12] = byte(transactionNo >> 24)

	out.write(buf)
	if buf.Len() > MaxMsgLen {
		libraries.ReleaseLog("消息发送失败，包体超过限制" + string(debug.Stack()))
	} else {
		b = buf.Bytes()
		msglen := buf.Len() - MsgHeadLen
		b[13] = byte(msglen)
		b[14] = byte(msglen >> 8)
		b[15] = byte(msglen >> 16)
		outchan <- buf
	}
}
func (db *MsgDB) BeginTransaction() (*MsgDBTransaction, error) {
	msg := db.msg
	if msg.DB.transaction == nil {
		result := GET_MSG_COMMON_BeginTransaction_result()
		out := GET_MSG_COMMON_BeginTransaction()
		out.TransactionNo = msg.DB.transactionNo
		err := msg.SendMsgWaitResult(0, out, &result)
		if err != nil {
			return nil, err
		}
		transaction, err := msg.DB.db.BeginTransaction()
		if err != nil {
			return nil, err
		}
		no := result.TransactionNo
		transactionMap.Store(no, transaction)
		msg.DB.transaction = &MsgDBTransaction{t: transaction, msg: msg, newTransactionNo: no}
		time.AfterFunc(MsgTimeOut*10*time.Second, func() {
			transactionMap.Delete(no)
			msg.DB.transaction.t.EndTransaction()

		})
		result.Put()
		out.Put()
	}
	return msg.DB.transaction, nil
}

//封装一下，如果事务存在，则自动切换到事务
func (db *MsgDB) Table(tablename string) *mysql.Mysql_Build {
	if db.transaction == nil {
		if db.transactionNo == 0 {
			return db.db.Table(tablename)
		}
		session, err := db.BeginTransaction()
		//session失败的时候使用mysqlbuild去传递err
		if err != nil {
			b := db.db.Table(tablename)
			b.SetErr(err)
			return b
		}
		return session.Table(tablename)

	}
	return db.transaction.Table(tablename)
}

//封装一下事务
func (t *MsgDBTransaction) Table(tablename string) *mysql.Mysql_Build {
	return t.t.Table(tablename)
}
func (t *MsgDBTransaction) Commit() error {
	//拦截住no不为0的commit
	if t.msg.DB.transactionNo > 0 {
		return nil
	}

	out := GET_MSG_COMMON_Transaction_Commit()
	out.No = t.newTransactionNo
	err := t.msg.SendMsgWaitResult(0, out, nil)
	out.Put()
	if err == nil {
		//马上执行commit避免提前EndTransaction
		if v, ok := transactionMap.LoadAndDelete(t.newTransactionNo); ok {
			v.(*mysql.Transaction).Commit()
		}
	}

	return err
}
func (t *MsgDBTransaction) Rollback() error {
	//允许任意节点发出rollback
	out := GET_MSG_COMMON_Transaction_RollBack()
	out.No = t.newTransactionNo
	err := t.msg.SendMsgWaitResult(0, out, nil)
	out.Put()
	return err
}
func (t *MsgDBTransaction) EndTransaction() {
	//拦截住no不为0的end
	if t.msg.DB.transactionNo > 0 {
		return
	}
	t.t.EndTransaction()
	//并尝试告诉网关 关闭其他可能还存在的事务
	if _, ok := transactionMap.Load(t.newTransactionNo); ok {
		out := GET_MSG_COMMON_Transaction_RollBack()
		out.No = t.newTransactionNo
		t.msg.SendMsgWaitResult(0, out, nil)
		out.Put()
	}

}
func (t *MsgDBTransaction) CommitCallback(f func()) {
	t.t.CommitCallback(f)
}
func (t *MsgDBTransaction) RollbackCallback(f func()) {
	t.t.RollbackCallback(f)
}
func msgTransactionCheck(data *MSG_COMMON_Transaction_Check, in *Msg) {
	if _, ok := transactionMap.Load(data.No); ok {
		in.WriteErr(nil)
	} else {
		in.WriteErr(errTransactionNotFound)
	}
}

func msgTransactionCommit(data *MSG_COMMON_Transaction_Commit, in *Msg) {
	if v, ok := transactionMap.LoadAndDelete(data.No); ok {
		v.(*mysql.Transaction).Commit()
	}
}

func msgTransactionRollBack(data *MSG_COMMON_Transaction_RollBack, in *Msg) {
	if v, ok := transactionMap.LoadAndDelete(data.No); ok {
		v.(*mysql.Transaction).Rollback()
	}
}
