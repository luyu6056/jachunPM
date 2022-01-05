package protocol

import (
	"errors"
	"fmt"
	"libraries"
	"mysql"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
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
	Ttl                uint16
	Local, remoteID    uint16
	Cmd                int32
	Data               MSG_DATA
	buf                *libraries.MsgBuffer
	datalen            int
	Svr                MsgServer
	cache              RpcCache
	transactionTimeOut time.Duration
	DB                 *MsgDB
	lang               CountryNo
	QueryID            uint32
	Addr               string
}
type MsgDB struct {
	msg           *Msg
	DB            *mysql.MysqlDB
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

	datalen := int(data[MsgHeadLen-3]) | int(data[MsgHeadLen-2])<<8 | int(data[MsgHeadLen-1])<<16
	if len(data) < MsgHeadLen+datalen {
		return nil, errMsgDataLen
	}
	msg = &Msg{DB: &MsgDB{}}
	msg.datalen = datalen
	msg.Msgno = uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
	msg.Ttl = uint16(data[4]) | uint16(data[5])<<8
	msg.Local = uint16(data[6]) | uint16(data[7])<<8
	msg.remoteID = uint16(data[8]) | uint16(data[9])<<8
	msg.DB.transactionNo = uint32(data[10]) | uint32(data[11])<<8 | uint32(data[12])<<16 | uint32(data[13])<<24
	msg.QueryID = uint32(data[14]) | uint32(data[15])<<8 | uint32(data[16])<<16 | uint32(data[17])<<24
	msg.DB.msg = msg
	msg.Cmd = int32(data[MsgHeadLen]) | int32(data[MsgHeadLen+1])<<8 | int32(data[MsgHeadLen+2])<<16 | int32(data[MsgHeadLen+3])<<24
	msg.lang = DefaultLang //暂时默认语言
	msg.buf = buf
	return msg, nil
}
func ReadOneMsgFromBytes(data []byte) (msg *Msg, length int, err error) {
	if len(data) < MsgHeadLen {
		return nil, 0, errMsgLen
	}

	datalen := int(data[MsgHeadLen-3]) | int(data[MsgHeadLen-2])<<8 | int(data[MsgHeadLen-1])<<16
	if len(data) < MsgHeadLen+datalen {
		return nil, 0, errMsgDataLen
	}
	msg = &Msg{DB: &MsgDB{}}
	msg.datalen = datalen
	msg.Msgno = uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
	msg.Ttl = uint16(data[4]) | uint16(data[5])<<8
	msg.Local = uint16(data[6]) | uint16(data[7])<<8
	msg.remoteID = uint16(data[8]) | uint16(data[9])<<8
	msg.DB.transactionNo = uint32(data[10]) | uint32(data[11])<<8 | uint32(data[12])<<16 | uint32(data[13])<<24
	msg.QueryID = uint32(data[14]) | uint32(data[15])<<8 | uint32(data[16])<<16 | uint32(data[17])<<24
	msg.DB.msg = msg
	msg.Cmd = int32(data[MsgHeadLen]) | int32(data[MsgHeadLen+1])<<8 | int32(data[MsgHeadLen+2])<<16 | int32(data[MsgHeadLen+3])<<24
	msg.lang = DefaultLang //暂时默认语言
	//msg.Data = data[MsgHeadLen : MsgHeadLen+datalen]
	msg.buf = BufPoolGet()
	msg.buf.Write(data[:MsgHeadLen+msg.datalen])
	return msg, MsgHeadLen + msg.datalen, nil
}

func (m *Msg) GetRemoteID() uint16 {
	return m.remoteID
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
func (m *Msg) Buf() *libraries.MsgBuffer {
	return m.buf
}

type MsgServer interface {
	SendMsg(remote uint16, msgno uint32, ttl uint16, transactionNo uint32, queryID uint32, out MSG_DATA)
	SendMsgWaitResult(remote uint16, msgno uint32, ttl uint16, transactionNo uint32, out MSG_DATA, result interface{}, timeout ...time.Duration) (err error)
}

//实现msg读写,由msg发出的消息，继承msgno和ttl
func (m *Msg) SetServer(svr MsgServer) {
	m.Svr = svr
	m.cache.Svr = svr
}

//指定服务器发送
func (m *Msg) SendMsg(remote uint16, out MSG_DATA) {
	transactionNo := m.DB.transactionNo
	if transactionNo == 0 && m.DB.transaction != nil {
		transactionNo = m.DB.transaction.newTransactionNo
	}
	m.Svr.SendMsg(remote, m.Msgno, m.Ttl, transactionNo, 0, out)
}
func (m *Msg) SendMsgWaitResult(remote uint16, out MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	if m.DB.transaction != nil && timeout != nil && timeout[0] > MsgTimeOut*time.Second { //事务有默认超时，不允许timeout大于事务超时
		return errTransactionTimeOut
	}
	transactionNo := m.DB.transactionNo
	if transactionNo == 0 && m.DB.transaction != nil {
		transactionNo = m.DB.transaction.newTransactionNo
	}
	return m.Svr.SendMsgWaitResult(remote, m.Msgno, m.Ttl, transactionNo, out, result, timeout...)
}

//原路返回
func (m *Msg) SendResult(out MSG_DATA) {
	transactionNo := m.DB.transactionNo
	if transactionNo == 0 && m.DB.transaction != nil {
		transactionNo = m.DB.transaction.newTransactionNo
	}
	m.Svr.SendMsg(m.Local, m.Msgno, m.Ttl, transactionNo, m.QueryID, out)
}

//原路返回err
func (m *Msg) WriteErr(err error) {
	data := GET_MSG_HOST_QueryErr()
	if m.QueryID == 0 {
		//libraries.DebugLog("返回queryID=0的err%v", err)
		return
	}
	if err != nil {
		data.Err = err.Error()
		data.Stack = debug.Stack()
	}
	transactionNo := m.DB.transactionNo
	if transactionNo == 0 && m.DB.transaction != nil {
		transactionNo = m.DB.transaction.newTransactionNo
	}
	m.Svr.SendMsg(m.Local, m.Msgno, m.Ttl, transactionNo, m.QueryID, data)
	data.Put()
	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
	}

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
func (m *Msg) LoadConfig(key string) (res map[string]map[string]interface{}, err error) {
	b, err := m.cache.Get(key, PATH_CONFIG_CACHE+m.lang.String())
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return
	}
	err = libraries.JsonUnmarshal(b, &res)
	return res, err
}

//解析具体某个值
func (m *Msg) LoadConfigToValue(key, key1, key2 string, value interface{}) error {
	b, err := m.cache.Get(key, PATH_CONFIG_CACHE+m.lang.String())
	if err != nil {
		return err
	}
	if len(b) == 0 {
		return nil
	}
	var res map[string]map[string]interface{}
	err = libraries.JsonUnmarshal(b, &res)
	if err != nil {
		return err
	}
	if v1, ok := res[key1]; ok {
		if v2, ok := v1[key2]; ok {
			err = libraries.JsonUnmarshal(libraries.JsonMarshal(v2), value)
		}
	}
	return err
}

//解决其他地方无法调用小写方法
func MSG_DATA_Write(data MSG_DATA, buf *libraries.MsgBuffer) {
	data.write(buf)
}

//result可以传入nil，但是返回非MSG_HOST_QueryErr就会报错
func SendMsgWaitResult(local, remote uint16, msgno uint32, ttl uint16, transactionNo uint32, out MSG_DATA, result interface{}, outchan chan *libraries.MsgBuffer, timeout ...time.Duration) (err error) {
	resultchan := make(chan RpcQueryResult, 1)
	rpcClientQueryLock.Lock()
	var queryId uint32
	for {
		queryId = uint32(rpcClientQueryId.INCRBY("id", 1))
		if queryId == 0 { //不允许为0
			queryId = uint32(rpcClientQueryId.INCRBY("id", 1))
		}
		if _, ok := rpcClientQueryMap[queryId]; !ok {
			break
		}
	}
	rpcClientQueryMap[queryId] = resultchan
	rpcClientQueryLock.Unlock()
	buf := BufPoolGet()
	b := buf.Make(MsgHeadLen)
	b[0] = byte(msgno)
	b[1] = byte(msgno >> 8)
	b[2] = byte(msgno >> 16)
	b[3] = byte(msgno >> 24)
	b[4] = byte(ttl)
	b[5] = byte(ttl >> 8)
	b[6] = byte(local)
	b[7] = byte(local >> 8)
	b[8] = byte(remote)
	b[9] = byte(remote >> 8)
	b[10] = byte(transactionNo)
	b[11] = byte(transactionNo >> 8)
	b[12] = byte(transactionNo >> 16)
	b[13] = byte(transactionNo >> 24)

	b[14] = byte(queryId)
	b[15] = byte(queryId >> 8)
	b[16] = byte(queryId >> 16)
	b[17] = byte(queryId >> 24)
	out.write(buf)
	if buf.Len() > MaxMsgLen {
		libraries.ReleaseLog("消息发送失败，包体超过限制" + string(debug.Stack()))
		rpcClientQueryLock.Lock()
		delete(rpcClientQueryMap, queryId)
		rpcClientQueryLock.Unlock()
		return errors.New("消息发送失败，包体超过限制" + string(debug.Stack()))
	} else {
		b = buf.Bytes()
		msglen := buf.Len() - MsgHeadLen
		b[MsgHeadLen-3] = byte(msglen)
		b[MsgHeadLen-2] = byte(msglen >> 8)
		b[MsgHeadLen-1] = byte(msglen >> 16)
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
		if r1.Kind() != reflect.Ptr {
			return RpcClientQueryResultErrType
		}
		if r1.Type().Elem().Name() != r2.Elem().Type().Name() {
			return errors.New(fmt.Sprintf("实际返回的结果为%s,与请求的%s不相符", r2.Elem().Type().Name(), r1.Type().Elem().Name()))
		}
		r1.Set(r2)
		return nil
	}
	select {
	case r := <-resultchan:
		rpcClientQueryLock.Lock()
		defer rpcClientQueryLock.Unlock()
		delete(rpcClientQueryMap, queryId)
		if err, ok := r.(*MSG_HOST_QueryErr); ok {
			if err.Err != "" {
				return errors.New(err.Err)
			} else if result == nil {
				return nil
			} else {
				r1 := reflect.ValueOf(result)
				return errors.New(fmt.Sprintf("实际返回的结果为MSG_HOST_QueryErr,与请求的%s不相符", r1.Type().Elem().Elem().Name()))
			}
		}
		return checkAndSetResult(r)
	case <-time.After(_timeout):
		rpcClientQueryLock.Lock()
		defer rpcClientQueryLock.Unlock()
		delete(rpcClientQueryMap, queryId)
		select {
		case r := <-resultchan:
			defer r.(MSG_DATA).Put()
			if err, ok := r.(*MSG_HOST_QueryErr); ok {
				if err.Err != "" {
					return errors.New(err.Err)
				} else if result == nil {
					return nil
				} else {
					r1 := reflect.ValueOf(result)
					return errors.New(fmt.Sprintf("实际返回的结果为MSG_HOST_QueryErr,与请求的%s不相符", r1.Elem().Elem().Type().Name()))
				}
			}
			return checkAndSetResult(r)
		default:
		}
	}
	return RpcClientQueryTimeOutErr

}
func SetMsgQuery(in *Msg) bool {
	queryID := in.QueryID
	rpcClientQueryLock.RLock()
	defer rpcClientQueryLock.RUnlock()
	if v, ok := rpcClientQueryMap[queryID]; ok {
		if in.Data == nil {
			in.ReadData()
		}
		v <- in.Data.(RpcQueryResult)
	} else {
		return false
	}

	return true

}
func SendMsg(local, remote uint16, msgno uint32, ttl uint16, transactionNo uint32, queryId uint32, out MSG_DATA, outchan chan *libraries.MsgBuffer) {
	buf := BufPoolGet()
	b := buf.Make(MsgHeadLen)
	b[0] = byte(msgno)
	b[1] = byte(msgno >> 8)
	b[2] = byte(msgno >> 16)
	b[3] = byte(msgno >> 24)
	b[4] = byte(ttl)
	b[5] = byte(ttl >> 8)
	b[6] = byte(local)
	b[7] = byte(local >> 8)
	b[8] = byte(remote)
	b[9] = byte(remote >> 8)
	b[10] = byte(transactionNo)
	b[11] = byte(transactionNo >> 8)
	b[12] = byte(transactionNo >> 16)
	b[13] = byte(transactionNo >> 24)

	b[14] = byte(queryId)
	b[15] = byte(queryId >> 8)
	b[16] = byte(queryId >> 16)
	b[17] = byte(queryId >> 24)
	out.write(buf)
	if buf.Len() > MaxMsgLen {
		libraries.ReleaseLog("消息发送失败，包体超过限制" + string(debug.Stack()))
	} else {
		b = buf.Bytes()
		msglen := buf.Len() - MsgHeadLen
		b[MsgHeadLen-3] = byte(msglen)
		b[MsgHeadLen-2] = byte(msglen >> 8)
		b[MsgHeadLen-1] = byte(msglen >> 16)
		outchan <- buf
	}
}

var msgTransactionLock sync.Mutex

func (msg *Msg) BeginTransaction() (*MsgDBTransaction, error) {
	if msg.DB.transaction == nil {
		msgTransactionLock.Lock()
		defer func() {
			msgTransactionLock.Unlock()
		}()

		if v, ok := transactionMap.Load(msg.DB.transactionNo); ok {
			msg.DB.transaction = &MsgDBTransaction{t: v.(*mysql.Transaction), msg: msg, newTransactionNo: msg.DB.transactionNo}
			return msg.DB.transaction, nil
		}
		result := GET_MSG_HOST_BeginTransaction_result()
		out := GET_MSG_HOST_BeginTransaction()
		out.TransactionNo = msg.DB.transactionNo
		err := msg.SendMsgWaitResult(0, out, &result)
		if err != nil {
			return nil, err
		}
		no := result.TransactionNo
		msg.DB.transaction = &MsgDBTransaction{t: nil, msg: msg, newTransactionNo: no}
		if msg.DB.DB != nil {
			transaction, err := msg.DB.DB.BeginTransaction()
			if err != nil {
				return nil, err
			}

			msg.DB.transaction.t = transaction
			transactionMap.Store(no, transaction)
		} else {
			transactionMap.Store(no, nil)
		}

		time.AfterFunc(MsgTimeOut*10*time.Second, func() {
			if _, ok := transactionMap.LoadAndDelete(no); ok {
				out := GET_MSG_HOST_Transaction_RollBack()
				out.No = no
				msg.SendMsgWaitResult(0, out, nil)
				out.Put()
				if msg.DB.transaction.t != nil {
					msg.DB.transaction.t.Rollback()
					msg.DB.transaction.t.EndTransaction()
				}

			}

		})
		result.Put()
		out.Put()
		return msg.DB.transaction, nil
	}
	//沿用旧的事务conn给予新的DB，transactionNo，拦截后续申请的子事务
	newMsg := &Msg{
		Msgno:              msg.Msgno,
		Ttl:                msg.Ttl,
		Local:              msg.Local,
		remoteID:           msg.remoteID,
		Cmd:                msg.Cmd,
		Data:               msg.Data,
		buf:                msg.buf,
		datalen:            msg.datalen,
		Svr:                msg.Svr,
		cache:              msg.cache,
		transactionTimeOut: msg.transactionTimeOut,
		DB:                 &MsgDB{transactionNo: msg.DB.transaction.newTransactionNo},
		lang:               msg.lang,
	}
	newMsg.DB.msg = newMsg
	newMsg.DB.transaction = &MsgDBTransaction{t: msg.DB.transaction.t, msg: newMsg, newTransactionNo: msg.DB.transaction.newTransactionNo}
	return newMsg.DB.transaction, nil
}

//封装一下，如果事务存在，则自动切换到事务
func (db *MsgDB) Table(tablename string) *mysql.Mysql_Table {
	if db.transaction == nil {
		if db.transactionNo == 0 {
			return db.DB.Table(tablename)
		}
		session, err := db.msg.BeginTransaction()
		//session失败的时候使用mysqlbuild去传递err
		if err != nil {
			b := db.DB.Table(tablename)
			b.SetErr(err)
			return b
		}
		return session.Table(tablename)

	}
	return db.transaction.Table(tablename)
}
func (db *MsgDB) Raw(sql string, arg ...interface{}) *mysql.Mysql_RawBuild {
	return db.DB.Raw(sql, arg...)
}

//封装一下事务
func (t *MsgDBTransaction) Table(tablename string) *mysql.Mysql_Table {
	return t.t.Table(tablename)
}
func (t *MsgDBTransaction) Raw(sql string, arg ...interface{}) *mysql.Mysql_RawBuild {
	return t.msg.DB.Raw(sql, arg...)
}
func (t *MsgDBTransaction) Commit() error {

	//拦截住no不为0的commit
	if t.msg.DB.transactionNo > 0 {

		return nil
	}

	if _, ok := transactionMap.Load(t.newTransactionNo); ok {
		out := GET_MSG_HOST_Transaction_Commit()
		out.No = t.newTransactionNo
		err := t.msg.SendMsgWaitResult(0, out, nil)
		out.Put()
		return err
	}
	return nil
}
func (t *MsgDBTransaction) Rollback() error {

	if v, ok := transactionMap.LoadAndDelete(t.newTransactionNo); ok {
		//允许任意节点发出rollback
		out := GET_MSG_HOST_Transaction_RollBack()
		out.No = t.newTransactionNo
		err := t.msg.SendMsgWaitResult(0, out, nil)
		out.Put()
		if t, ok := v.(*mysql.Transaction); ok {
			t.Rollback()
			t.EndTransaction()
		}

		return err
	}
	return nil
}

func (t *MsgDBTransaction) CommitCallback(f func()) {
	t.t.CommitCallback(f)
}
func (t *MsgDBTransaction) RollbackCallback(f func()) {
	t.t.RollbackCallback(f)
}
func MsgTransactionCheck(data *MSG_HOST_Transaction_Check) error {

	if _, ok := transactionMap.Load(data.No); !ok {
		return errTransactionNotFound
	}
	return nil
}

func MsgTransactionCommit(data *MSG_HOST_Transaction_Commit) {
	if v, ok := transactionMap.LoadAndDelete(data.No); ok {
		if t, ok := v.(*mysql.Transaction); ok {
			t.Commit()
			t.EndTransaction()
		}
	}
}

func MsgTransactionRollBack(data *MSG_HOST_Transaction_RollBack) {
	if v, ok := transactionMap.LoadAndDelete(data.No); ok {
		if t, ok := v.(*mysql.Transaction); ok {
			t.Rollback()
			t.EndTransaction()
		}
	}
}
func (m *Msg) ActionCreate(objectType string, objectID int32, actionType, comment, extra string, products, projects []int32) (actionID int64, err error) {
	out := GET_MSG_LOG_Action_Create()
	out.ObjectType = objectType
	out.ObjectID = objectID
	out.ActionType = actionType //操作类型
	out.Comment = comment       //信息
	out.Extra = extra           //额外信息
	out.ActorId = m.GetUserID()
	out.Products = products
	out.Projects = projects
	var result *MSG_LOG_Action_Create_result

	if err = m.SendMsgWaitResult(0, out, &result); err != nil {
		return
	}
	out.Put()
	return result.ActionId, nil

}
func (m *Msg) GetUserID() int32 {
	b, err := m.cache.Get("Uid", "Msg:"+strconv.Itoa(int(m.Msgno)))
	if err == nil {
		return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
	}
	return 0
}
func (m *Msg) HasPriv(moduleName, methodName string, ext ...interface{}) bool {
	return true
}

type diff struct {
	key   string
	value string
}

func (m *Msg) ActionLogHistory(actionID int64, oldObj, newObj interface{}) ([]*MSG_LOG_History, error) {
	change, err := GetDiffChange(oldObj, newObj)
	if len(change) > 0 {
		out := GET_MSG_LOG_Action_AddHistory()
		out.History = change
		out.Id = actionID
		m.SendMsg(0, out)
	}
	return change, err
}

type ChangeHistory []*MSG_LOG_History

func (change ChangeHistory) Add(actionID int64, m *Msg) {
	if len(change) > 0 {
		out := GET_MSG_LOG_Action_AddHistory()
		out.History = change
		out.Id = actionID
		m.SendMsg(0, out)
	}
}
func WRITE_ChangeHistory(change ChangeHistory, buf *libraries.MsgBuffer) {
	WRITE_int(len(change), buf)
	for _, c := range change {
		WRITE_MSG_LOG_History(c, buf)
	}
}
func READ_ChangeHistory(buf *libraries.MsgBuffer) ChangeHistory {
	var change ChangeHistory
	l := READ_int(buf)
	for i := 0; i < l; i++ {
		change = append(change, READ_MSG_LOG_History(buf))
	}
	return change
}
func GetDiffChange(oldObj, newObj interface{}) (change ChangeHistory, err error) {
	r_o := reflect.ValueOf(oldObj)
	r_n := reflect.ValueOf(newObj)
	for r_o.Kind() == reflect.Ptr {
		r_o = r_o.Elem()
	}
	for r_n.Kind() == reflect.Ptr {
		r_n = r_n.Elem()
	}
	if r_o.Kind() != reflect.Struct && r_o.Kind() != reflect.Map && r_o.Kind() != r_n.Kind() {
		return nil, errors.New(fmt.Sprintf("ActionLogHistory传入不接受的格式 old %v,new %v", r_o.Kind(), r_n.Kind()))
	}
	switch r_o.Kind() {
	case reflect.Struct:
		if r_o.Type().String() != r_n.Type().String() {
			return nil, errors.New(fmt.Sprintf("ActionLogHistory传入不一样的结构体 old %v,new %v", r_o.Type().String(), r_n.Type().String()))
		}
		for i := 0; i < r_o.NumField(); i++ {
			field := r_n.Field(i)
			t := r_n.Type().Field(i)
			if field.Kind() == reflect.Ptr || (field.Kind() == reflect.Struct && t.Type.String() != "time.Time") || field.Kind() == reflect.Map {
				continue
			}
			lowerKey := strings.ToLower(t.Name)
			value := libraries.I2S(field.Interface())
			oldValue := libraries.I2S(r_o.Field(i).Interface())

			if t.Type.String() == "time.Time" {
				if time, _ := field.Interface().(time.Time); time.Unix() > 946656000 { //随便定个时间，暂时定做2000年1月1日的时间戳
					value = time.Format("2006-01-02 16:04:05")
				} else {
					value = ""
				}
			}
			if lowerKey == "lastediteddate" || lowerKey == "lasteditedby" || lowerKey == "assigneddate" || lowerKey == "editedby" || lowerKey == "editeddate" || lowerKey == "uid" || (lowerKey == "finisheddate" && value == "") || (lowerKey == "canceleddate" && value == "") || (lowerKey == "closeddate" && value == "") {
				continue
			}
			if oldValue == value {
				continue
			}
			var diffString []string
			if strings.Contains(value, "\n") || strings.Contains(oldValue, "\n") || strings.Contains("name,title,desc,spec,steps,content,digest,verify,report", lowerKey) {
				//text1 = str_replace('&nbsp;', '', trim(text1));
				//text2 = str_replace('&nbsp;', '', trim(text2));
				w := strings.Split(oldValue, "\n")
				o := strings.Split(value, "\n")
				w1 := string_array_diff_assoc(w, o)
				o1 := string_array_diff_assoc(o, w)
				var w2 []diff

				for idx, val := range w1 {
					w2 = append(w2, diff{
						key:   fmt.Sprintf("%03d<", idx),
						value: fmt.Sprintf("%03d- ", idx+1) + "<del>" + val + "</del>",
					})
				}
				for idx, val := range o1 {
					w2 = append(w2, diff{
						key:   fmt.Sprintf("%03d>", idx),
						value: fmt.Sprintf("%03d+ ", idx+1) + "<ins>" + val + "</ins>",
					})
				}
				order_diff(w2)
				for _, v := range w2 {
					diffString = append(diffString, v.value)
				}
			}
			if len(diffString) > 0 {
				change = append(change, &MSG_LOG_History{
					Field: lowerKey,
					Old:   oldValue,
					New:   value,
					Diff:  strings.Join(diffString, "\n"),
				})
			} else {
				change = append(change, &MSG_LOG_History{
					Field: lowerKey,
					Old:   oldValue,
					New:   value,
				})
			}

		}
	}
	return
}

func string_array_diff_assoc(a, b []string) []string {
	for i := len(a) - 1; i >= 0; i-- {
		v1 := a[i]
		for _, v2 := range b {
			if v1 == v2 {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}
	return a
}
func CopyObj(oldobj interface{}, newobj interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
	}()
	r_o := reflect.ValueOf(oldobj)
	r_n := reflect.ValueOf(newobj)
	r_nt := reflect.TypeOf(newobj)
	for r_o.Kind() == reflect.Ptr {
		r_o = r_o.Elem()

	}
	for r_n.Kind() == reflect.Ptr {
		if r_n.Elem().Kind() == reflect.Invalid {
			r_n.Set(reflect.New(r_nt.Elem()))
		}
		r_n = r_n.Elem()
		r_nt = r_nt.Elem()
	}

	if r_o.Kind() != reflect.Struct && r_o.Kind() != reflect.Map && r_o.Kind() != r_n.Kind() {
		return errors.New(fmt.Sprintf("CopyObj传入不接受的格式 old %v,new %v", r_o.Kind(), r_n.Kind()))
	}
	if r_o.Type().String() != r_nt.String() {
		return errors.New(fmt.Sprintf("CopyObj传入不一样的结构体 old %v,new %v", r_o.Type().String(), r_n.Type().String()))
	}
	for i := 0; i < r_o.NumField(); i++ {
		field := r_n.Field(i)
		field.Set(reflect.ValueOf(r_o.Field(i).Interface()))
	}
	return nil
}
