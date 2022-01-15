package protocol

import (
	"sync"
	"libraries"
	"time"
)

const (
	CMD_MSG_HOST_regServer = -102044159
	CMD_MSG_HOST_regServer_result = -1926345727
	CMD_MSG_HOST_StartTicker = -756565247
	CMD_MSG_HOST_PING = -1848741631
	CMD_MSG_HOST_PONG = -1790767871
	CMD_MSG_HOST_WINDOW_UPDATE = -1489369599
	CMD_MSG_HOST_CACHE_GET = 516345601
	CMD_MSG_HOST_CACHE_GET_result = 939602945
	CMD_MSG_HOST_CACHE_GETPATH = 1095125505
	CMD_MSG_HOST_CACHE_GETPATH_result = -1722911999
	CMD_MSG_HOST_CACHE_SET = 99206145
	CMD_MSG_HOST_CACHE_DEL = 267184385
	CMD_MSG_HOST_CACHE_DelPath = 1895097345
	CMD_MSG_HOST_GET_Msgno = 1583876353
	CMD_MSG_HOST_GET_Msgno_result = -1315762175
	CMD_MSG_HOST_QueryErr = 557953793
	CMD_MSG_HOST_ResetWindow = -630575871
	CMD_MSG_FILE_upload = 1110878977
	CMD_MSG_FILE_upload_result = -2057389055
	CMD_MSG_FILE_getByID = -1871273215
	CMD_MSG_FILE_getByID_result = -1484540159
	CMD_MSG_FILE_updateMapByWhere = -671290623
	CMD_MSG_FILE_DeleteByID = -1568529407
	CMD_MSG_HOST_BeginTransaction = 2097391617
	CMD_MSG_HOST_BeginTransaction_result = -1940757759
	CMD_MSG_HOST_Transaction_Commit = 1920226305
	CMD_MSG_HOST_Transaction_RollBack = 2009659137
	CMD_MSG_HOST_Transaction_Check = -1626340351
	CMD_MSG_FILE_getByObject = 232008193
	CMD_MSG_FILE_getByObject_result = 665022465
	CMD_MSG_HOST_GET_MsgUserId = 1184964097
	CMD_MSG_HOST_GET_MsgUserId_result = 1153682945
	CMD_MSG_FILE_uploadTmp = 1057561345
	CMD_MSG_FILE_RangeDown = 636129793
	CMD_MSG_FILE_RangeDown_result = 226334977
	CMD_MSG_FILE_updateTmp = 642796289
	CMD_MSG_FILE_edit = 1409978369
	CMD_MSG_FILE_getByWhere = 1857790721
	CMD_MSG_FILE_getByWhere_result = -453068031
	CMD_MSG_FILE_download_byIds = 743980545
	CMD_MSG_FILE_download_byIds_result = -1058626559
)

type MSG_HOST_regServer struct {
	No uint8
	IpPort string
	Time int64
	Token string
	Window int32
}

var pool_MSG_HOST_regServer = sync.Pool{New: func() interface{} { return &MSG_HOST_regServer{} }}

func GET_MSG_HOST_regServer() *MSG_HOST_regServer {
	return pool_MSG_HOST_regServer.Get().(*MSG_HOST_regServer)
}

func (data *MSG_HOST_regServer) cmd() int32 {
	return CMD_MSG_HOST_regServer
}

func (data *MSG_HOST_regServer) Put() {
	data.No = 0
	data.IpPort = ``
	data.Time = 0
	data.Token = ``
	data.Window = 0
	pool_MSG_HOST_regServer.Put(data)
}
func (data *MSG_HOST_regServer) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_regServer,buf)
	WRITE_MSG_HOST_regServer(data, buf)
}

func WRITE_MSG_HOST_regServer(data *MSG_HOST_regServer, buf *libraries.MsgBuffer) {
	WRITE_uint8(data.No, buf)
	WRITE_string(data.IpPort, buf)
	WRITE_int64(data.Time, buf)
	WRITE_string(data.Token, buf)
	WRITE_int32(data.Window, buf)
}

func READ_MSG_HOST_regServer(buf *libraries.MsgBuffer) *MSG_HOST_regServer {
	data := pool_MSG_HOST_regServer.Get().(*MSG_HOST_regServer)
	data.read(buf)
	return data
}

func (data *MSG_HOST_regServer) read(buf *libraries.MsgBuffer) {
	data.No = READ_uint8(buf)
	data.IpPort = READ_string(buf)
	data.Time = READ_int64(buf)
	data.Token = READ_string(buf)
	data.Window = READ_int32(buf)

}

type MSG_HOST_regServer_result struct {
	Id uint8
}

var pool_MSG_HOST_regServer_result = sync.Pool{New: func() interface{} { return &MSG_HOST_regServer_result{} }}

func GET_MSG_HOST_regServer_result() *MSG_HOST_regServer_result {
	return pool_MSG_HOST_regServer_result.Get().(*MSG_HOST_regServer_result)
}

func (data *MSG_HOST_regServer_result) cmd() int32 {
	return CMD_MSG_HOST_regServer_result
}

func (data *MSG_HOST_regServer_result) Put() {
	data.Id = 0
	pool_MSG_HOST_regServer_result.Put(data)
}
func (data *MSG_HOST_regServer_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_regServer_result,buf)
	WRITE_MSG_HOST_regServer_result(data, buf)
}

func WRITE_MSG_HOST_regServer_result(data *MSG_HOST_regServer_result, buf *libraries.MsgBuffer) {
	WRITE_uint8(data.Id, buf)
}

func READ_MSG_HOST_regServer_result(buf *libraries.MsgBuffer) *MSG_HOST_regServer_result {
	data := pool_MSG_HOST_regServer_result.Get().(*MSG_HOST_regServer_result)
	data.read(buf)
	return data
}

func (data *MSG_HOST_regServer_result) read(buf *libraries.MsgBuffer) {
	data.Id = READ_uint8(buf)

}

type MSG_HOST_StartTicker struct {
}

var pool_MSG_HOST_StartTicker = sync.Pool{New: func() interface{} { return &MSG_HOST_StartTicker{} }}

func GET_MSG_HOST_StartTicker() *MSG_HOST_StartTicker {
	return pool_MSG_HOST_StartTicker.Get().(*MSG_HOST_StartTicker)
}

func (data *MSG_HOST_StartTicker) cmd() int32 {
	return CMD_MSG_HOST_StartTicker
}

func (data *MSG_HOST_StartTicker) Put() {
	pool_MSG_HOST_StartTicker.Put(data)
}
func (data *MSG_HOST_StartTicker) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_StartTicker,buf)
	WRITE_MSG_HOST_StartTicker(data, buf)
}

func WRITE_MSG_HOST_StartTicker(data *MSG_HOST_StartTicker, buf *libraries.MsgBuffer) {
}

func READ_MSG_HOST_StartTicker(buf *libraries.MsgBuffer) *MSG_HOST_StartTicker {
	data := pool_MSG_HOST_StartTicker.Get().(*MSG_HOST_StartTicker)
	data.read(buf)
	return data
}

func (data *MSG_HOST_StartTicker) read(buf *libraries.MsgBuffer) {

}

type MSG_HOST_PING struct {
	Rand int
}

var pool_MSG_HOST_PING = sync.Pool{New: func() interface{} { return &MSG_HOST_PING{} }}

func GET_MSG_HOST_PING() *MSG_HOST_PING {
	return pool_MSG_HOST_PING.Get().(*MSG_HOST_PING)
}

func (data *MSG_HOST_PING) cmd() int32 {
	return CMD_MSG_HOST_PING
}

func (data *MSG_HOST_PING) Put() {
	data.Rand = 0
	pool_MSG_HOST_PING.Put(data)
}
func (data *MSG_HOST_PING) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_PING,buf)
	WRITE_MSG_HOST_PING(data, buf)
}

func WRITE_MSG_HOST_PING(data *MSG_HOST_PING, buf *libraries.MsgBuffer) {
	WRITE_int(data.Rand, buf)
}

func READ_MSG_HOST_PING(buf *libraries.MsgBuffer) *MSG_HOST_PING {
	data := pool_MSG_HOST_PING.Get().(*MSG_HOST_PING)
	data.read(buf)
	return data
}

func (data *MSG_HOST_PING) read(buf *libraries.MsgBuffer) {
	data.Rand = READ_int(buf)

}

type MSG_HOST_PONG struct {
	Rand int
}

var pool_MSG_HOST_PONG = sync.Pool{New: func() interface{} { return &MSG_HOST_PONG{} }}

func GET_MSG_HOST_PONG() *MSG_HOST_PONG {
	return pool_MSG_HOST_PONG.Get().(*MSG_HOST_PONG)
}

func (data *MSG_HOST_PONG) cmd() int32 {
	return CMD_MSG_HOST_PONG
}

func (data *MSG_HOST_PONG) Put() {
	data.Rand = 0
	pool_MSG_HOST_PONG.Put(data)
}
func (data *MSG_HOST_PONG) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_PONG,buf)
	WRITE_MSG_HOST_PONG(data, buf)
}

func WRITE_MSG_HOST_PONG(data *MSG_HOST_PONG, buf *libraries.MsgBuffer) {
	WRITE_int(data.Rand, buf)
}

func READ_MSG_HOST_PONG(buf *libraries.MsgBuffer) *MSG_HOST_PONG {
	data := pool_MSG_HOST_PONG.Get().(*MSG_HOST_PONG)
	data.read(buf)
	return data
}

func (data *MSG_HOST_PONG) read(buf *libraries.MsgBuffer) {
	data.Rand = READ_int(buf)

}

type MSG_HOST_WINDOW_UPDATE struct {
	Add int32
}

var pool_MSG_HOST_WINDOW_UPDATE = sync.Pool{New: func() interface{} { return &MSG_HOST_WINDOW_UPDATE{} }}

func GET_MSG_HOST_WINDOW_UPDATE() *MSG_HOST_WINDOW_UPDATE {
	return pool_MSG_HOST_WINDOW_UPDATE.Get().(*MSG_HOST_WINDOW_UPDATE)
}

func (data *MSG_HOST_WINDOW_UPDATE) cmd() int32 {
	return CMD_MSG_HOST_WINDOW_UPDATE
}

func (data *MSG_HOST_WINDOW_UPDATE) Put() {
	data.Add = 0
	pool_MSG_HOST_WINDOW_UPDATE.Put(data)
}
func (data *MSG_HOST_WINDOW_UPDATE) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_WINDOW_UPDATE,buf)
	WRITE_MSG_HOST_WINDOW_UPDATE(data, buf)
}

func WRITE_MSG_HOST_WINDOW_UPDATE(data *MSG_HOST_WINDOW_UPDATE, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Add, buf)
}

func READ_MSG_HOST_WINDOW_UPDATE(buf *libraries.MsgBuffer) *MSG_HOST_WINDOW_UPDATE {
	data := pool_MSG_HOST_WINDOW_UPDATE.Get().(*MSG_HOST_WINDOW_UPDATE)
	data.read(buf)
	return data
}

func (data *MSG_HOST_WINDOW_UPDATE) read(buf *libraries.MsgBuffer) {
	data.Add = READ_int32(buf)

}

type MSG_HOST_CACHE_GET struct {
	Path string
	Name string
}

var pool_MSG_HOST_CACHE_GET = sync.Pool{New: func() interface{} { return &MSG_HOST_CACHE_GET{} }}

func GET_MSG_HOST_CACHE_GET() *MSG_HOST_CACHE_GET {
	return pool_MSG_HOST_CACHE_GET.Get().(*MSG_HOST_CACHE_GET)
}

func (data *MSG_HOST_CACHE_GET) cmd() int32 {
	return CMD_MSG_HOST_CACHE_GET
}

func (data *MSG_HOST_CACHE_GET) Put() {
	data.Path = ``
	data.Name = ``
	pool_MSG_HOST_CACHE_GET.Put(data)
}
func (data *MSG_HOST_CACHE_GET) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_CACHE_GET,buf)
	WRITE_MSG_HOST_CACHE_GET(data, buf)
}

func WRITE_MSG_HOST_CACHE_GET(data *MSG_HOST_CACHE_GET, buf *libraries.MsgBuffer) {
	WRITE_string(data.Path, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_HOST_CACHE_GET(buf *libraries.MsgBuffer) *MSG_HOST_CACHE_GET {
	data := pool_MSG_HOST_CACHE_GET.Get().(*MSG_HOST_CACHE_GET)
	data.read(buf)
	return data
}

func (data *MSG_HOST_CACHE_GET) read(buf *libraries.MsgBuffer) {
	data.Path = READ_string(buf)
	data.Name = READ_string(buf)

}

type MSG_HOST_CACHE_GET_result struct {
	Value []byte
}

var pool_MSG_HOST_CACHE_GET_result = sync.Pool{New: func() interface{} { return &MSG_HOST_CACHE_GET_result{} }}

func GET_MSG_HOST_CACHE_GET_result() *MSG_HOST_CACHE_GET_result {
	return pool_MSG_HOST_CACHE_GET_result.Get().(*MSG_HOST_CACHE_GET_result)
}

func (data *MSG_HOST_CACHE_GET_result) cmd() int32 {
	return CMD_MSG_HOST_CACHE_GET_result
}

func (data *MSG_HOST_CACHE_GET_result) Put() {
	data.Value = data.Value[:0]
	pool_MSG_HOST_CACHE_GET_result.Put(data)
}
func (data *MSG_HOST_CACHE_GET_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_CACHE_GET_result,buf)
	WRITE_MSG_HOST_CACHE_GET_result(data, buf)
}

func WRITE_MSG_HOST_CACHE_GET_result(data *MSG_HOST_CACHE_GET_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.Value), buf)
	buf.Write(data.Value)
}

func READ_MSG_HOST_CACHE_GET_result(buf *libraries.MsgBuffer) *MSG_HOST_CACHE_GET_result {
	data := pool_MSG_HOST_CACHE_GET_result.Get().(*MSG_HOST_CACHE_GET_result)
	data.read(buf)
	return data
}

func (data *MSG_HOST_CACHE_GET_result) read(buf *libraries.MsgBuffer) {
	Value_len := READ_int(buf)
	data.Value = make([]byte, Value_len)
	copy(data.Value,buf.Next(Value_len))

}

type MSG_HOST_CACHE_GETPATH struct {
	Path string
	Names []string
}

var pool_MSG_HOST_CACHE_GETPATH = sync.Pool{New: func() interface{} { return &MSG_HOST_CACHE_GETPATH{} }}

func GET_MSG_HOST_CACHE_GETPATH() *MSG_HOST_CACHE_GETPATH {
	return pool_MSG_HOST_CACHE_GETPATH.Get().(*MSG_HOST_CACHE_GETPATH)
}

func (data *MSG_HOST_CACHE_GETPATH) cmd() int32 {
	return CMD_MSG_HOST_CACHE_GETPATH
}

func (data *MSG_HOST_CACHE_GETPATH) Put() {
	data.Path = ``
	data.Names = data.Names[:0]
	pool_MSG_HOST_CACHE_GETPATH.Put(data)
}
func (data *MSG_HOST_CACHE_GETPATH) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_CACHE_GETPATH,buf)
	WRITE_MSG_HOST_CACHE_GETPATH(data, buf)
}

func WRITE_MSG_HOST_CACHE_GETPATH(data *MSG_HOST_CACHE_GETPATH, buf *libraries.MsgBuffer) {
	WRITE_string(data.Path, buf)
	WRITE_int(len(data.Names), buf)
	for _, v := range data.Names{
		WRITE_string(v, buf)
	}
}

func READ_MSG_HOST_CACHE_GETPATH(buf *libraries.MsgBuffer) *MSG_HOST_CACHE_GETPATH {
	data := pool_MSG_HOST_CACHE_GETPATH.Get().(*MSG_HOST_CACHE_GETPATH)
	data.read(buf)
	return data
}

func (data *MSG_HOST_CACHE_GETPATH) read(buf *libraries.MsgBuffer) {
	data.Path = READ_string(buf)
	Names_len := READ_int(buf)
	if Names_len>cap(data.Names){
		data.Names= make([]string, Names_len)
	}else{
		data.Names = data.Names[:Names_len]
	}
	for i := 0; i < Names_len; i++ {
		data.Names[i] = READ_string(buf)
	}

}

type MSG_HOST_CACHE_GETPATH_result struct {
	Value [][]byte
}

var pool_MSG_HOST_CACHE_GETPATH_result = sync.Pool{New: func() interface{} { return &MSG_HOST_CACHE_GETPATH_result{} }}

func GET_MSG_HOST_CACHE_GETPATH_result() *MSG_HOST_CACHE_GETPATH_result {
	return pool_MSG_HOST_CACHE_GETPATH_result.Get().(*MSG_HOST_CACHE_GETPATH_result)
}

func (data *MSG_HOST_CACHE_GETPATH_result) cmd() int32 {
	return CMD_MSG_HOST_CACHE_GETPATH_result
}

func (data *MSG_HOST_CACHE_GETPATH_result) Put() {
	data.Value = data.Value[:0]
	pool_MSG_HOST_CACHE_GETPATH_result.Put(data)
}
func (data *MSG_HOST_CACHE_GETPATH_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_CACHE_GETPATH_result,buf)
	WRITE_MSG_HOST_CACHE_GETPATH_result(data, buf)
}

func WRITE_MSG_HOST_CACHE_GETPATH_result(data *MSG_HOST_CACHE_GETPATH_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.Value), buf)
	for _, v := range data.Value{
		WRITE_int(len(v), buf)
		buf.Write(v)
	}
}

func READ_MSG_HOST_CACHE_GETPATH_result(buf *libraries.MsgBuffer) *MSG_HOST_CACHE_GETPATH_result {
	data := pool_MSG_HOST_CACHE_GETPATH_result.Get().(*MSG_HOST_CACHE_GETPATH_result)
	data.read(buf)
	return data
}

func (data *MSG_HOST_CACHE_GETPATH_result) read(buf *libraries.MsgBuffer) {
	Value_len := READ_int(buf)
	for i := 0; i < Value_len; i++ {
		l := READ_int(buf)
		b := make([]byte,l)
		copy(b,buf.Next(int(l)))
		data.Value = append(data.Value, b)
	}

}

type MSG_HOST_CACHE_SET struct {
	Path string
	Name string
	Value []byte
	Expire int64
}

var pool_MSG_HOST_CACHE_SET = sync.Pool{New: func() interface{} { return &MSG_HOST_CACHE_SET{} }}

func GET_MSG_HOST_CACHE_SET() *MSG_HOST_CACHE_SET {
	return pool_MSG_HOST_CACHE_SET.Get().(*MSG_HOST_CACHE_SET)
}

func (data *MSG_HOST_CACHE_SET) cmd() int32 {
	return CMD_MSG_HOST_CACHE_SET
}

func (data *MSG_HOST_CACHE_SET) Put() {
	data.Path = ``
	data.Name = ``
	data.Value = data.Value[:0]
	data.Expire = 0
	pool_MSG_HOST_CACHE_SET.Put(data)
}
func (data *MSG_HOST_CACHE_SET) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_CACHE_SET,buf)
	WRITE_MSG_HOST_CACHE_SET(data, buf)
}

func WRITE_MSG_HOST_CACHE_SET(data *MSG_HOST_CACHE_SET, buf *libraries.MsgBuffer) {
	WRITE_string(data.Path, buf)
	WRITE_string(data.Name, buf)
	WRITE_int(len(data.Value), buf)
	buf.Write(data.Value)
	WRITE_int64(data.Expire, buf)
}

func READ_MSG_HOST_CACHE_SET(buf *libraries.MsgBuffer) *MSG_HOST_CACHE_SET {
	data := pool_MSG_HOST_CACHE_SET.Get().(*MSG_HOST_CACHE_SET)
	data.read(buf)
	return data
}

func (data *MSG_HOST_CACHE_SET) read(buf *libraries.MsgBuffer) {
	data.Path = READ_string(buf)
	data.Name = READ_string(buf)
	Value_len := READ_int(buf)
	data.Value = make([]byte, Value_len)
	copy(data.Value,buf.Next(Value_len))
	data.Expire = READ_int64(buf)

}

type MSG_HOST_CACHE_DEL struct {
	Path string
	Name string
}

var pool_MSG_HOST_CACHE_DEL = sync.Pool{New: func() interface{} { return &MSG_HOST_CACHE_DEL{} }}

func GET_MSG_HOST_CACHE_DEL() *MSG_HOST_CACHE_DEL {
	return pool_MSG_HOST_CACHE_DEL.Get().(*MSG_HOST_CACHE_DEL)
}

func (data *MSG_HOST_CACHE_DEL) cmd() int32 {
	return CMD_MSG_HOST_CACHE_DEL
}

func (data *MSG_HOST_CACHE_DEL) Put() {
	data.Path = ``
	data.Name = ``
	pool_MSG_HOST_CACHE_DEL.Put(data)
}
func (data *MSG_HOST_CACHE_DEL) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_CACHE_DEL,buf)
	WRITE_MSG_HOST_CACHE_DEL(data, buf)
}

func WRITE_MSG_HOST_CACHE_DEL(data *MSG_HOST_CACHE_DEL, buf *libraries.MsgBuffer) {
	WRITE_string(data.Path, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_HOST_CACHE_DEL(buf *libraries.MsgBuffer) *MSG_HOST_CACHE_DEL {
	data := pool_MSG_HOST_CACHE_DEL.Get().(*MSG_HOST_CACHE_DEL)
	data.read(buf)
	return data
}

func (data *MSG_HOST_CACHE_DEL) read(buf *libraries.MsgBuffer) {
	data.Path = READ_string(buf)
	data.Name = READ_string(buf)

}

type MSG_HOST_CACHE_DelPath struct {
	Path string
}

var pool_MSG_HOST_CACHE_DelPath = sync.Pool{New: func() interface{} { return &MSG_HOST_CACHE_DelPath{} }}

func GET_MSG_HOST_CACHE_DelPath() *MSG_HOST_CACHE_DelPath {
	return pool_MSG_HOST_CACHE_DelPath.Get().(*MSG_HOST_CACHE_DelPath)
}

func (data *MSG_HOST_CACHE_DelPath) cmd() int32 {
	return CMD_MSG_HOST_CACHE_DelPath
}

func (data *MSG_HOST_CACHE_DelPath) Put() {
	data.Path = ``
	pool_MSG_HOST_CACHE_DelPath.Put(data)
}
func (data *MSG_HOST_CACHE_DelPath) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_CACHE_DelPath,buf)
	WRITE_MSG_HOST_CACHE_DelPath(data, buf)
}

func WRITE_MSG_HOST_CACHE_DelPath(data *MSG_HOST_CACHE_DelPath, buf *libraries.MsgBuffer) {
	WRITE_string(data.Path, buf)
}

func READ_MSG_HOST_CACHE_DelPath(buf *libraries.MsgBuffer) *MSG_HOST_CACHE_DelPath {
	data := pool_MSG_HOST_CACHE_DelPath.Get().(*MSG_HOST_CACHE_DelPath)
	data.read(buf)
	return data
}

func (data *MSG_HOST_CACHE_DelPath) read(buf *libraries.MsgBuffer) {
	data.Path = READ_string(buf)

}

type MSG_HOST_GET_Msgno struct {
	Uid int32
}

var pool_MSG_HOST_GET_Msgno = sync.Pool{New: func() interface{} { return &MSG_HOST_GET_Msgno{} }}

func GET_MSG_HOST_GET_Msgno() *MSG_HOST_GET_Msgno {
	return pool_MSG_HOST_GET_Msgno.Get().(*MSG_HOST_GET_Msgno)
}

func (data *MSG_HOST_GET_Msgno) cmd() int32 {
	return CMD_MSG_HOST_GET_Msgno
}

func (data *MSG_HOST_GET_Msgno) Put() {
	data.Uid = 0
	pool_MSG_HOST_GET_Msgno.Put(data)
}
func (data *MSG_HOST_GET_Msgno) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_GET_Msgno,buf)
	WRITE_MSG_HOST_GET_Msgno(data, buf)
}

func WRITE_MSG_HOST_GET_Msgno(data *MSG_HOST_GET_Msgno, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
}

func READ_MSG_HOST_GET_Msgno(buf *libraries.MsgBuffer) *MSG_HOST_GET_Msgno {
	data := pool_MSG_HOST_GET_Msgno.Get().(*MSG_HOST_GET_Msgno)
	data.read(buf)
	return data
}

func (data *MSG_HOST_GET_Msgno) read(buf *libraries.MsgBuffer) {
	data.Uid = READ_int32(buf)

}

type MSG_HOST_GET_Msgno_result struct {
	Msgno uint32
}

var pool_MSG_HOST_GET_Msgno_result = sync.Pool{New: func() interface{} { return &MSG_HOST_GET_Msgno_result{} }}

func GET_MSG_HOST_GET_Msgno_result() *MSG_HOST_GET_Msgno_result {
	return pool_MSG_HOST_GET_Msgno_result.Get().(*MSG_HOST_GET_Msgno_result)
}

func (data *MSG_HOST_GET_Msgno_result) cmd() int32 {
	return CMD_MSG_HOST_GET_Msgno_result
}

func (data *MSG_HOST_GET_Msgno_result) Put() {
	data.Msgno = 0
	pool_MSG_HOST_GET_Msgno_result.Put(data)
}
func (data *MSG_HOST_GET_Msgno_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_GET_Msgno_result,buf)
	WRITE_MSG_HOST_GET_Msgno_result(data, buf)
}

func WRITE_MSG_HOST_GET_Msgno_result(data *MSG_HOST_GET_Msgno_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.Msgno, buf)
}

func READ_MSG_HOST_GET_Msgno_result(buf *libraries.MsgBuffer) *MSG_HOST_GET_Msgno_result {
	data := pool_MSG_HOST_GET_Msgno_result.Get().(*MSG_HOST_GET_Msgno_result)
	data.read(buf)
	return data
}

func (data *MSG_HOST_GET_Msgno_result) read(buf *libraries.MsgBuffer) {
	data.Msgno = READ_uint32(buf)

}

type MSG_HOST_QueryErr struct {
	Err string
	Stack []byte
}

var pool_MSG_HOST_QueryErr = sync.Pool{New: func() interface{} { return &MSG_HOST_QueryErr{} }}

func GET_MSG_HOST_QueryErr() *MSG_HOST_QueryErr {
	return pool_MSG_HOST_QueryErr.Get().(*MSG_HOST_QueryErr)
}

func (data *MSG_HOST_QueryErr) cmd() int32 {
	return CMD_MSG_HOST_QueryErr
}

func (data *MSG_HOST_QueryErr) Put() {
	data.Err = ``
	data.Stack = data.Stack[:0]
	pool_MSG_HOST_QueryErr.Put(data)
}
func (data *MSG_HOST_QueryErr) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_QueryErr,buf)
	WRITE_MSG_HOST_QueryErr(data, buf)
}

func WRITE_MSG_HOST_QueryErr(data *MSG_HOST_QueryErr, buf *libraries.MsgBuffer) {
	WRITE_string(data.Err, buf)
	WRITE_int(len(data.Stack), buf)
	buf.Write(data.Stack)
}

func READ_MSG_HOST_QueryErr(buf *libraries.MsgBuffer) *MSG_HOST_QueryErr {
	data := pool_MSG_HOST_QueryErr.Get().(*MSG_HOST_QueryErr)
	data.read(buf)
	return data
}

func (data *MSG_HOST_QueryErr) read(buf *libraries.MsgBuffer) {
	data.Err = READ_string(buf)
	Stack_len := READ_int(buf)
	data.Stack = make([]byte, Stack_len)
	copy(data.Stack,buf.Next(Stack_len))

}

type MSG_HOST_ResetWindow struct {
	Window int32
}

var pool_MSG_HOST_ResetWindow = sync.Pool{New: func() interface{} { return &MSG_HOST_ResetWindow{} }}

func GET_MSG_HOST_ResetWindow() *MSG_HOST_ResetWindow {
	return pool_MSG_HOST_ResetWindow.Get().(*MSG_HOST_ResetWindow)
}

func (data *MSG_HOST_ResetWindow) cmd() int32 {
	return CMD_MSG_HOST_ResetWindow
}

func (data *MSG_HOST_ResetWindow) Put() {
	data.Window = 0
	pool_MSG_HOST_ResetWindow.Put(data)
}
func (data *MSG_HOST_ResetWindow) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_ResetWindow,buf)
	WRITE_MSG_HOST_ResetWindow(data, buf)
}

func WRITE_MSG_HOST_ResetWindow(data *MSG_HOST_ResetWindow, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Window, buf)
}

func READ_MSG_HOST_ResetWindow(buf *libraries.MsgBuffer) *MSG_HOST_ResetWindow {
	data := pool_MSG_HOST_ResetWindow.Get().(*MSG_HOST_ResetWindow)
	data.read(buf)
	return data
}

func (data *MSG_HOST_ResetWindow) read(buf *libraries.MsgBuffer) {
	data.Window = READ_int32(buf)

}

type MSG_FILE_upload struct {
	Name string
	Title string
	Data []byte
	AddBy int32
	ObjectType string
	ObjectID int32
	Code string
	Type string
}

var pool_MSG_FILE_upload = sync.Pool{New: func() interface{} { return &MSG_FILE_upload{} }}

func GET_MSG_FILE_upload() *MSG_FILE_upload {
	return pool_MSG_FILE_upload.Get().(*MSG_FILE_upload)
}

func (data *MSG_FILE_upload) cmd() int32 {
	return CMD_MSG_FILE_upload
}

func (data *MSG_FILE_upload) Put() {
	data.Name = ``
	data.Title = ``
	data.Data = data.Data[:0]
	data.AddBy = 0
	data.ObjectType = ``
	data.ObjectID = 0
	data.Code = ``
	data.Type = ``
	pool_MSG_FILE_upload.Put(data)
}
func (data *MSG_FILE_upload) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_upload,buf)
	WRITE_MSG_FILE_upload(data, buf)
}

func WRITE_MSG_FILE_upload(data *MSG_FILE_upload, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
	WRITE_string(data.Title, buf)
	WRITE_int(len(data.Data), buf)
	buf.Write(data.Data)
	WRITE_int32(data.AddBy, buf)
	WRITE_string(data.ObjectType, buf)
	WRITE_int32(data.ObjectID, buf)
	WRITE_string(data.Code, buf)
	WRITE_string(data.Type, buf)
}

func READ_MSG_FILE_upload(buf *libraries.MsgBuffer) *MSG_FILE_upload {
	data := pool_MSG_FILE_upload.Get().(*MSG_FILE_upload)
	data.read(buf)
	return data
}

func (data *MSG_FILE_upload) read(buf *libraries.MsgBuffer) {
	data.Name = READ_string(buf)
	data.Title = READ_string(buf)
	Data_len := READ_int(buf)
	data.Data = make([]byte, Data_len)
	copy(data.Data,buf.Next(Data_len))
	data.AddBy = READ_int32(buf)
	data.ObjectType = READ_string(buf)
	data.ObjectID = READ_int32(buf)
	data.Code = READ_string(buf)
	data.Type = READ_string(buf)

}

type MSG_FILE_upload_result struct {
	FileID int64
}

var pool_MSG_FILE_upload_result = sync.Pool{New: func() interface{} { return &MSG_FILE_upload_result{} }}

func GET_MSG_FILE_upload_result() *MSG_FILE_upload_result {
	return pool_MSG_FILE_upload_result.Get().(*MSG_FILE_upload_result)
}

func (data *MSG_FILE_upload_result) cmd() int32 {
	return CMD_MSG_FILE_upload_result
}

func (data *MSG_FILE_upload_result) Put() {
	data.FileID = 0
	pool_MSG_FILE_upload_result.Put(data)
}
func (data *MSG_FILE_upload_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_upload_result,buf)
	WRITE_MSG_FILE_upload_result(data, buf)
}

func WRITE_MSG_FILE_upload_result(data *MSG_FILE_upload_result, buf *libraries.MsgBuffer) {
	WRITE_int64(data.FileID, buf)
}

func READ_MSG_FILE_upload_result(buf *libraries.MsgBuffer) *MSG_FILE_upload_result {
	data := pool_MSG_FILE_upload_result.Get().(*MSG_FILE_upload_result)
	data.read(buf)
	return data
}

func (data *MSG_FILE_upload_result) read(buf *libraries.MsgBuffer) {
	data.FileID = READ_int64(buf)

}

type MSG_FILE_getByID struct {
	FileID int64
	NoData bool
}

var pool_MSG_FILE_getByID = sync.Pool{New: func() interface{} { return &MSG_FILE_getByID{} }}

func GET_MSG_FILE_getByID() *MSG_FILE_getByID {
	return pool_MSG_FILE_getByID.Get().(*MSG_FILE_getByID)
}

func (data *MSG_FILE_getByID) cmd() int32 {
	return CMD_MSG_FILE_getByID
}

func (data *MSG_FILE_getByID) Put() {
	data.FileID = 0
	data.NoData = false
	pool_MSG_FILE_getByID.Put(data)
}
func (data *MSG_FILE_getByID) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_getByID,buf)
	WRITE_MSG_FILE_getByID(data, buf)
}

func WRITE_MSG_FILE_getByID(data *MSG_FILE_getByID, buf *libraries.MsgBuffer) {
	WRITE_int64(data.FileID, buf)
	WRITE_bool(data.NoData, buf)
}

func READ_MSG_FILE_getByID(buf *libraries.MsgBuffer) *MSG_FILE_getByID {
	data := pool_MSG_FILE_getByID.Get().(*MSG_FILE_getByID)
	data.read(buf)
	return data
}

func (data *MSG_FILE_getByID) read(buf *libraries.MsgBuffer) {
	data.FileID = READ_int64(buf)
	data.NoData = READ_bool(buf)

}

type MSG_FILE_getByID_result struct {
	FileID int64
	Name string
	Ext string
	Size int64
	Type string
	AddedDate time.Time
	ObjectType string
	ObjectID int32
}

var pool_MSG_FILE_getByID_result = sync.Pool{New: func() interface{} { return &MSG_FILE_getByID_result{} }}

func GET_MSG_FILE_getByID_result() *MSG_FILE_getByID_result {
	return pool_MSG_FILE_getByID_result.Get().(*MSG_FILE_getByID_result)
}

func (data *MSG_FILE_getByID_result) cmd() int32 {
	return CMD_MSG_FILE_getByID_result
}

func (data *MSG_FILE_getByID_result) Put() {
	data.FileID = 0
	data.Name = ``
	data.Ext = ``
	data.Size = 0
	data.Type = ``
	data.AddedDate = time.UnixMicro(0)
	data.ObjectType = ``
	data.ObjectID = 0
	pool_MSG_FILE_getByID_result.Put(data)
}
func (data *MSG_FILE_getByID_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_getByID_result,buf)
	WRITE_MSG_FILE_getByID_result(data, buf)
}

func WRITE_MSG_FILE_getByID_result(data *MSG_FILE_getByID_result, buf *libraries.MsgBuffer) {
	WRITE_int64(data.FileID, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Ext, buf)
	WRITE_int64(data.Size, buf)
	WRITE_string(data.Type, buf)
	WRITE_int64(data.AddedDate.UnixMicro(), buf)
	WRITE_string(data.ObjectType, buf)
	WRITE_int32(data.ObjectID, buf)
}

func READ_MSG_FILE_getByID_result(buf *libraries.MsgBuffer) *MSG_FILE_getByID_result {
	data := pool_MSG_FILE_getByID_result.Get().(*MSG_FILE_getByID_result)
	data.read(buf)
	return data
}

func (data *MSG_FILE_getByID_result) read(buf *libraries.MsgBuffer) {
	data.FileID = READ_int64(buf)
	data.Name = READ_string(buf)
	data.Ext = READ_string(buf)
	data.Size = READ_int64(buf)
	data.Type = READ_string(buf)
	data.AddedDate = time.UnixMicro(READ_int64(buf))
	data.ObjectType = READ_string(buf)
	data.ObjectID = READ_int32(buf)

}

type MSG_FILE_updateMapByWhere struct {
	Where map[string]interface{}
	Update map[string]interface{}
}

var pool_MSG_FILE_updateMapByWhere = sync.Pool{New: func() interface{} { return &MSG_FILE_updateMapByWhere{} }}

func GET_MSG_FILE_updateMapByWhere() *MSG_FILE_updateMapByWhere {
	return pool_MSG_FILE_updateMapByWhere.Get().(*MSG_FILE_updateMapByWhere)
}

func (data *MSG_FILE_updateMapByWhere) cmd() int32 {
	return CMD_MSG_FILE_updateMapByWhere
}

func (data *MSG_FILE_updateMapByWhere) Put() {
	data.Where = nil
	data.Update = nil
	pool_MSG_FILE_updateMapByWhere.Put(data)
}
func (data *MSG_FILE_updateMapByWhere) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_updateMapByWhere,buf)
	WRITE_MSG_FILE_updateMapByWhere(data, buf)
}

func WRITE_MSG_FILE_updateMapByWhere(data *MSG_FILE_updateMapByWhere, buf *libraries.MsgBuffer) {
	WRITE_map(data.Where,buf)
	WRITE_map(data.Update,buf)
}

func READ_MSG_FILE_updateMapByWhere(buf *libraries.MsgBuffer) *MSG_FILE_updateMapByWhere {
	data := pool_MSG_FILE_updateMapByWhere.Get().(*MSG_FILE_updateMapByWhere)
	data.read(buf)
	return data
}

func (data *MSG_FILE_updateMapByWhere) read(buf *libraries.MsgBuffer) {
	READ_map(&data.Where,buf)
	READ_map(&data.Update,buf)

}

type MSG_FILE_DeleteByID struct {
	FileID int64
}

var pool_MSG_FILE_DeleteByID = sync.Pool{New: func() interface{} { return &MSG_FILE_DeleteByID{} }}

func GET_MSG_FILE_DeleteByID() *MSG_FILE_DeleteByID {
	return pool_MSG_FILE_DeleteByID.Get().(*MSG_FILE_DeleteByID)
}

func (data *MSG_FILE_DeleteByID) cmd() int32 {
	return CMD_MSG_FILE_DeleteByID
}

func (data *MSG_FILE_DeleteByID) Put() {
	data.FileID = 0
	pool_MSG_FILE_DeleteByID.Put(data)
}
func (data *MSG_FILE_DeleteByID) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_DeleteByID,buf)
	WRITE_MSG_FILE_DeleteByID(data, buf)
}

func WRITE_MSG_FILE_DeleteByID(data *MSG_FILE_DeleteByID, buf *libraries.MsgBuffer) {
	WRITE_int64(data.FileID, buf)
}

func READ_MSG_FILE_DeleteByID(buf *libraries.MsgBuffer) *MSG_FILE_DeleteByID {
	data := pool_MSG_FILE_DeleteByID.Get().(*MSG_FILE_DeleteByID)
	data.read(buf)
	return data
}

func (data *MSG_FILE_DeleteByID) read(buf *libraries.MsgBuffer) {
	data.FileID = READ_int64(buf)

}

type MSG_HOST_BeginTransaction struct {
	TransactionNo uint32
}

var pool_MSG_HOST_BeginTransaction = sync.Pool{New: func() interface{} { return &MSG_HOST_BeginTransaction{} }}

func GET_MSG_HOST_BeginTransaction() *MSG_HOST_BeginTransaction {
	return pool_MSG_HOST_BeginTransaction.Get().(*MSG_HOST_BeginTransaction)
}

func (data *MSG_HOST_BeginTransaction) cmd() int32 {
	return CMD_MSG_HOST_BeginTransaction
}

func (data *MSG_HOST_BeginTransaction) Put() {
	data.TransactionNo = 0
	pool_MSG_HOST_BeginTransaction.Put(data)
}
func (data *MSG_HOST_BeginTransaction) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_BeginTransaction,buf)
	WRITE_MSG_HOST_BeginTransaction(data, buf)
}

func WRITE_MSG_HOST_BeginTransaction(data *MSG_HOST_BeginTransaction, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.TransactionNo, buf)
}

func READ_MSG_HOST_BeginTransaction(buf *libraries.MsgBuffer) *MSG_HOST_BeginTransaction {
	data := pool_MSG_HOST_BeginTransaction.Get().(*MSG_HOST_BeginTransaction)
	data.read(buf)
	return data
}

func (data *MSG_HOST_BeginTransaction) read(buf *libraries.MsgBuffer) {
	data.TransactionNo = READ_uint32(buf)

}

type MSG_HOST_BeginTransaction_result struct {
	TransactionNo uint32
}

var pool_MSG_HOST_BeginTransaction_result = sync.Pool{New: func() interface{} { return &MSG_HOST_BeginTransaction_result{} }}

func GET_MSG_HOST_BeginTransaction_result() *MSG_HOST_BeginTransaction_result {
	return pool_MSG_HOST_BeginTransaction_result.Get().(*MSG_HOST_BeginTransaction_result)
}

func (data *MSG_HOST_BeginTransaction_result) cmd() int32 {
	return CMD_MSG_HOST_BeginTransaction_result
}

func (data *MSG_HOST_BeginTransaction_result) Put() {
	data.TransactionNo = 0
	pool_MSG_HOST_BeginTransaction_result.Put(data)
}
func (data *MSG_HOST_BeginTransaction_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_BeginTransaction_result,buf)
	WRITE_MSG_HOST_BeginTransaction_result(data, buf)
}

func WRITE_MSG_HOST_BeginTransaction_result(data *MSG_HOST_BeginTransaction_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.TransactionNo, buf)
}

func READ_MSG_HOST_BeginTransaction_result(buf *libraries.MsgBuffer) *MSG_HOST_BeginTransaction_result {
	data := pool_MSG_HOST_BeginTransaction_result.Get().(*MSG_HOST_BeginTransaction_result)
	data.read(buf)
	return data
}

func (data *MSG_HOST_BeginTransaction_result) read(buf *libraries.MsgBuffer) {
	data.TransactionNo = READ_uint32(buf)

}

type MSG_HOST_Transaction_Commit struct {
	No uint32
}

var pool_MSG_HOST_Transaction_Commit = sync.Pool{New: func() interface{} { return &MSG_HOST_Transaction_Commit{} }}

func GET_MSG_HOST_Transaction_Commit() *MSG_HOST_Transaction_Commit {
	return pool_MSG_HOST_Transaction_Commit.Get().(*MSG_HOST_Transaction_Commit)
}

func (data *MSG_HOST_Transaction_Commit) cmd() int32 {
	return CMD_MSG_HOST_Transaction_Commit
}

func (data *MSG_HOST_Transaction_Commit) Put() {
	data.No = 0
	pool_MSG_HOST_Transaction_Commit.Put(data)
}
func (data *MSG_HOST_Transaction_Commit) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_Transaction_Commit,buf)
	WRITE_MSG_HOST_Transaction_Commit(data, buf)
}

func WRITE_MSG_HOST_Transaction_Commit(data *MSG_HOST_Transaction_Commit, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.No, buf)
}

func READ_MSG_HOST_Transaction_Commit(buf *libraries.MsgBuffer) *MSG_HOST_Transaction_Commit {
	data := pool_MSG_HOST_Transaction_Commit.Get().(*MSG_HOST_Transaction_Commit)
	data.read(buf)
	return data
}

func (data *MSG_HOST_Transaction_Commit) read(buf *libraries.MsgBuffer) {
	data.No = READ_uint32(buf)

}

type MSG_HOST_Transaction_RollBack struct {
	No uint32
}

var pool_MSG_HOST_Transaction_RollBack = sync.Pool{New: func() interface{} { return &MSG_HOST_Transaction_RollBack{} }}

func GET_MSG_HOST_Transaction_RollBack() *MSG_HOST_Transaction_RollBack {
	return pool_MSG_HOST_Transaction_RollBack.Get().(*MSG_HOST_Transaction_RollBack)
}

func (data *MSG_HOST_Transaction_RollBack) cmd() int32 {
	return CMD_MSG_HOST_Transaction_RollBack
}

func (data *MSG_HOST_Transaction_RollBack) Put() {
	data.No = 0
	pool_MSG_HOST_Transaction_RollBack.Put(data)
}
func (data *MSG_HOST_Transaction_RollBack) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_Transaction_RollBack,buf)
	WRITE_MSG_HOST_Transaction_RollBack(data, buf)
}

func WRITE_MSG_HOST_Transaction_RollBack(data *MSG_HOST_Transaction_RollBack, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.No, buf)
}

func READ_MSG_HOST_Transaction_RollBack(buf *libraries.MsgBuffer) *MSG_HOST_Transaction_RollBack {
	data := pool_MSG_HOST_Transaction_RollBack.Get().(*MSG_HOST_Transaction_RollBack)
	data.read(buf)
	return data
}

func (data *MSG_HOST_Transaction_RollBack) read(buf *libraries.MsgBuffer) {
	data.No = READ_uint32(buf)

}

type MSG_HOST_Transaction_Check struct {
	No uint32
}

var pool_MSG_HOST_Transaction_Check = sync.Pool{New: func() interface{} { return &MSG_HOST_Transaction_Check{} }}

func GET_MSG_HOST_Transaction_Check() *MSG_HOST_Transaction_Check {
	return pool_MSG_HOST_Transaction_Check.Get().(*MSG_HOST_Transaction_Check)
}

func (data *MSG_HOST_Transaction_Check) cmd() int32 {
	return CMD_MSG_HOST_Transaction_Check
}

func (data *MSG_HOST_Transaction_Check) Put() {
	data.No = 0
	pool_MSG_HOST_Transaction_Check.Put(data)
}
func (data *MSG_HOST_Transaction_Check) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_Transaction_Check,buf)
	WRITE_MSG_HOST_Transaction_Check(data, buf)
}

func WRITE_MSG_HOST_Transaction_Check(data *MSG_HOST_Transaction_Check, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.No, buf)
}

func READ_MSG_HOST_Transaction_Check(buf *libraries.MsgBuffer) *MSG_HOST_Transaction_Check {
	data := pool_MSG_HOST_Transaction_Check.Get().(*MSG_HOST_Transaction_Check)
	data.read(buf)
	return data
}

func (data *MSG_HOST_Transaction_Check) read(buf *libraries.MsgBuffer) {
	data.No = READ_uint32(buf)

}

type MSG_FILE_getByObject struct {
	ObjectType string
	ObjectID int32
}

var pool_MSG_FILE_getByObject = sync.Pool{New: func() interface{} { return &MSG_FILE_getByObject{} }}

func GET_MSG_FILE_getByObject() *MSG_FILE_getByObject {
	return pool_MSG_FILE_getByObject.Get().(*MSG_FILE_getByObject)
}

func (data *MSG_FILE_getByObject) cmd() int32 {
	return CMD_MSG_FILE_getByObject
}

func (data *MSG_FILE_getByObject) Put() {
	data.ObjectType = ``
	data.ObjectID = 0
	pool_MSG_FILE_getByObject.Put(data)
}
func (data *MSG_FILE_getByObject) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_getByObject,buf)
	WRITE_MSG_FILE_getByObject(data, buf)
}

func WRITE_MSG_FILE_getByObject(data *MSG_FILE_getByObject, buf *libraries.MsgBuffer) {
	WRITE_string(data.ObjectType, buf)
	WRITE_int32(data.ObjectID, buf)
}

func READ_MSG_FILE_getByObject(buf *libraries.MsgBuffer) *MSG_FILE_getByObject {
	data := pool_MSG_FILE_getByObject.Get().(*MSG_FILE_getByObject)
	data.read(buf)
	return data
}

func (data *MSG_FILE_getByObject) read(buf *libraries.MsgBuffer) {
	data.ObjectType = READ_string(buf)
	data.ObjectID = READ_int32(buf)

}

type MSG_FILE_getByObject_result struct {
	List []*MSG_FILE_getByID_result
}

var pool_MSG_FILE_getByObject_result = sync.Pool{New: func() interface{} { return &MSG_FILE_getByObject_result{} }}

func GET_MSG_FILE_getByObject_result() *MSG_FILE_getByObject_result {
	return pool_MSG_FILE_getByObject_result.Get().(*MSG_FILE_getByObject_result)
}

func (data *MSG_FILE_getByObject_result) cmd() int32 {
	return CMD_MSG_FILE_getByObject_result
}

func (data *MSG_FILE_getByObject_result) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_FILE_getByObject_result.Put(data)
}
func (data *MSG_FILE_getByObject_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_getByObject_result,buf)
	WRITE_MSG_FILE_getByObject_result(data, buf)
}

func WRITE_MSG_FILE_getByObject_result(data *MSG_FILE_getByObject_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_FILE_getByID_result(v, buf)
	}
}

func READ_MSG_FILE_getByObject_result(buf *libraries.MsgBuffer) *MSG_FILE_getByObject_result {
	data := pool_MSG_FILE_getByObject_result.Get().(*MSG_FILE_getByObject_result)
	data.read(buf)
	return data
}

func (data *MSG_FILE_getByObject_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_FILE_getByID_result, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_FILE_getByID_result(buf)
	}

}

type MSG_HOST_GET_MsgUserId struct {
}

var pool_MSG_HOST_GET_MsgUserId = sync.Pool{New: func() interface{} { return &MSG_HOST_GET_MsgUserId{} }}

func GET_MSG_HOST_GET_MsgUserId() *MSG_HOST_GET_MsgUserId {
	return pool_MSG_HOST_GET_MsgUserId.Get().(*MSG_HOST_GET_MsgUserId)
}

func (data *MSG_HOST_GET_MsgUserId) cmd() int32 {
	return CMD_MSG_HOST_GET_MsgUserId
}

func (data *MSG_HOST_GET_MsgUserId) Put() {
	pool_MSG_HOST_GET_MsgUserId.Put(data)
}
func (data *MSG_HOST_GET_MsgUserId) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_GET_MsgUserId,buf)
	WRITE_MSG_HOST_GET_MsgUserId(data, buf)
}

func WRITE_MSG_HOST_GET_MsgUserId(data *MSG_HOST_GET_MsgUserId, buf *libraries.MsgBuffer) {
}

func READ_MSG_HOST_GET_MsgUserId(buf *libraries.MsgBuffer) *MSG_HOST_GET_MsgUserId {
	data := pool_MSG_HOST_GET_MsgUserId.Get().(*MSG_HOST_GET_MsgUserId)
	data.read(buf)
	return data
}

func (data *MSG_HOST_GET_MsgUserId) read(buf *libraries.MsgBuffer) {

}

type MSG_HOST_GET_MsgUserId_result struct {
	Uid int32
}

var pool_MSG_HOST_GET_MsgUserId_result = sync.Pool{New: func() interface{} { return &MSG_HOST_GET_MsgUserId_result{} }}

func GET_MSG_HOST_GET_MsgUserId_result() *MSG_HOST_GET_MsgUserId_result {
	return pool_MSG_HOST_GET_MsgUserId_result.Get().(*MSG_HOST_GET_MsgUserId_result)
}

func (data *MSG_HOST_GET_MsgUserId_result) cmd() int32 {
	return CMD_MSG_HOST_GET_MsgUserId_result
}

func (data *MSG_HOST_GET_MsgUserId_result) Put() {
	data.Uid = 0
	pool_MSG_HOST_GET_MsgUserId_result.Put(data)
}
func (data *MSG_HOST_GET_MsgUserId_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_HOST_GET_MsgUserId_result,buf)
	WRITE_MSG_HOST_GET_MsgUserId_result(data, buf)
}

func WRITE_MSG_HOST_GET_MsgUserId_result(data *MSG_HOST_GET_MsgUserId_result, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
}

func READ_MSG_HOST_GET_MsgUserId_result(buf *libraries.MsgBuffer) *MSG_HOST_GET_MsgUserId_result {
	data := pool_MSG_HOST_GET_MsgUserId_result.Get().(*MSG_HOST_GET_MsgUserId_result)
	data.read(buf)
	return data
}

func (data *MSG_HOST_GET_MsgUserId_result) read(buf *libraries.MsgBuffer) {
	data.Uid = READ_int32(buf)

}

type MSG_FILE_uploadTmp struct {
	Name string
	Data []byte
	Index int
	BlockSize int
}

var pool_MSG_FILE_uploadTmp = sync.Pool{New: func() interface{} { return &MSG_FILE_uploadTmp{} }}

func GET_MSG_FILE_uploadTmp() *MSG_FILE_uploadTmp {
	return pool_MSG_FILE_uploadTmp.Get().(*MSG_FILE_uploadTmp)
}

func (data *MSG_FILE_uploadTmp) cmd() int32 {
	return CMD_MSG_FILE_uploadTmp
}

func (data *MSG_FILE_uploadTmp) Put() {
	data.Name = ``
	data.Data = data.Data[:0]
	data.Index = 0
	data.BlockSize = 0
	pool_MSG_FILE_uploadTmp.Put(data)
}
func (data *MSG_FILE_uploadTmp) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_uploadTmp,buf)
	WRITE_MSG_FILE_uploadTmp(data, buf)
}

func WRITE_MSG_FILE_uploadTmp(data *MSG_FILE_uploadTmp, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
	WRITE_int(len(data.Data), buf)
	buf.Write(data.Data)
	WRITE_int(data.Index, buf)
	WRITE_int(data.BlockSize, buf)
}

func READ_MSG_FILE_uploadTmp(buf *libraries.MsgBuffer) *MSG_FILE_uploadTmp {
	data := pool_MSG_FILE_uploadTmp.Get().(*MSG_FILE_uploadTmp)
	data.read(buf)
	return data
}

func (data *MSG_FILE_uploadTmp) read(buf *libraries.MsgBuffer) {
	data.Name = READ_string(buf)
	Data_len := READ_int(buf)
	data.Data = make([]byte, Data_len)
	copy(data.Data,buf.Next(Data_len))
	data.Index = READ_int(buf)
	data.BlockSize = READ_int(buf)

}

type MSG_FILE_RangeDown struct {
	FileID int64
	Start int64
	End int64
}

var pool_MSG_FILE_RangeDown = sync.Pool{New: func() interface{} { return &MSG_FILE_RangeDown{} }}

func GET_MSG_FILE_RangeDown() *MSG_FILE_RangeDown {
	return pool_MSG_FILE_RangeDown.Get().(*MSG_FILE_RangeDown)
}

func (data *MSG_FILE_RangeDown) cmd() int32 {
	return CMD_MSG_FILE_RangeDown
}

func (data *MSG_FILE_RangeDown) Put() {
	data.FileID = 0
	data.Start = 0
	data.End = 0
	pool_MSG_FILE_RangeDown.Put(data)
}
func (data *MSG_FILE_RangeDown) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_RangeDown,buf)
	WRITE_MSG_FILE_RangeDown(data, buf)
}

func WRITE_MSG_FILE_RangeDown(data *MSG_FILE_RangeDown, buf *libraries.MsgBuffer) {
	WRITE_int64(data.FileID, buf)
	WRITE_int64(data.Start, buf)
	WRITE_int64(data.End, buf)
}

func READ_MSG_FILE_RangeDown(buf *libraries.MsgBuffer) *MSG_FILE_RangeDown {
	data := pool_MSG_FILE_RangeDown.Get().(*MSG_FILE_RangeDown)
	data.read(buf)
	return data
}

func (data *MSG_FILE_RangeDown) read(buf *libraries.MsgBuffer) {
	data.FileID = READ_int64(buf)
	data.Start = READ_int64(buf)
	data.End = READ_int64(buf)

}

type MSG_FILE_RangeDown_result struct {
	Byte []byte
}

var pool_MSG_FILE_RangeDown_result = sync.Pool{New: func() interface{} { return &MSG_FILE_RangeDown_result{} }}

func GET_MSG_FILE_RangeDown_result() *MSG_FILE_RangeDown_result {
	return pool_MSG_FILE_RangeDown_result.Get().(*MSG_FILE_RangeDown_result)
}

func (data *MSG_FILE_RangeDown_result) cmd() int32 {
	return CMD_MSG_FILE_RangeDown_result
}

func (data *MSG_FILE_RangeDown_result) Put() {
	data.Byte = data.Byte[:0]
	pool_MSG_FILE_RangeDown_result.Put(data)
}
func (data *MSG_FILE_RangeDown_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_RangeDown_result,buf)
	WRITE_MSG_FILE_RangeDown_result(data, buf)
}

func WRITE_MSG_FILE_RangeDown_result(data *MSG_FILE_RangeDown_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.Byte), buf)
	buf.Write(data.Byte)
}

func READ_MSG_FILE_RangeDown_result(buf *libraries.MsgBuffer) *MSG_FILE_RangeDown_result {
	data := pool_MSG_FILE_RangeDown_result.Get().(*MSG_FILE_RangeDown_result)
	data.read(buf)
	return data
}

func (data *MSG_FILE_RangeDown_result) read(buf *libraries.MsgBuffer) {
	Byte_len := READ_int(buf)
	data.Byte = make([]byte, Byte_len)
	copy(data.Byte,buf.Next(Byte_len))

}

type MSG_FILE_updateTmp struct {
	Files []*MSG_FILE_upload
}

var pool_MSG_FILE_updateTmp = sync.Pool{New: func() interface{} { return &MSG_FILE_updateTmp{} }}

func GET_MSG_FILE_updateTmp() *MSG_FILE_updateTmp {
	return pool_MSG_FILE_updateTmp.Get().(*MSG_FILE_updateTmp)
}

func (data *MSG_FILE_updateTmp) cmd() int32 {
	return CMD_MSG_FILE_updateTmp
}

func (data *MSG_FILE_updateTmp) Put() {
	for _,v := range data.Files {
		v.Put()
	}
	data.Files = data.Files[:0]
	pool_MSG_FILE_updateTmp.Put(data)
}
func (data *MSG_FILE_updateTmp) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_updateTmp,buf)
	WRITE_MSG_FILE_updateTmp(data, buf)
}

func WRITE_MSG_FILE_updateTmp(data *MSG_FILE_updateTmp, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.Files), buf)
	for _, v := range data.Files{
		WRITE_MSG_FILE_upload(v, buf)
	}
}

func READ_MSG_FILE_updateTmp(buf *libraries.MsgBuffer) *MSG_FILE_updateTmp {
	data := pool_MSG_FILE_updateTmp.Get().(*MSG_FILE_updateTmp)
	data.read(buf)
	return data
}

func (data *MSG_FILE_updateTmp) read(buf *libraries.MsgBuffer) {
	Files_len := READ_int(buf)
	if Files_len>cap(data.Files){
		data.Files= make([]*MSG_FILE_upload, Files_len)
	}else{
		data.Files = data.Files[:Files_len]
	}
	for i := 0; i < Files_len; i++ {
		data.Files[i] = READ_MSG_FILE_upload(buf)
	}

}

type MSG_FILE_edit struct {
	FileID int64
	Name string
}

var pool_MSG_FILE_edit = sync.Pool{New: func() interface{} { return &MSG_FILE_edit{} }}

func GET_MSG_FILE_edit() *MSG_FILE_edit {
	return pool_MSG_FILE_edit.Get().(*MSG_FILE_edit)
}

func (data *MSG_FILE_edit) cmd() int32 {
	return CMD_MSG_FILE_edit
}

func (data *MSG_FILE_edit) Put() {
	data.FileID = 0
	data.Name = ``
	pool_MSG_FILE_edit.Put(data)
}
func (data *MSG_FILE_edit) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_edit,buf)
	WRITE_MSG_FILE_edit(data, buf)
}

func WRITE_MSG_FILE_edit(data *MSG_FILE_edit, buf *libraries.MsgBuffer) {
	WRITE_int64(data.FileID, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_FILE_edit(buf *libraries.MsgBuffer) *MSG_FILE_edit {
	data := pool_MSG_FILE_edit.Get().(*MSG_FILE_edit)
	data.read(buf)
	return data
}

func (data *MSG_FILE_edit) read(buf *libraries.MsgBuffer) {
	data.FileID = READ_int64(buf)
	data.Name = READ_string(buf)

}

type MSG_FILE_getByWhere struct {
	Where map[string]interface{}
	Page int
	PerPage int
	Total int
}

var pool_MSG_FILE_getByWhere = sync.Pool{New: func() interface{} { return &MSG_FILE_getByWhere{} }}

func GET_MSG_FILE_getByWhere() *MSG_FILE_getByWhere {
	return pool_MSG_FILE_getByWhere.Get().(*MSG_FILE_getByWhere)
}

func (data *MSG_FILE_getByWhere) cmd() int32 {
	return CMD_MSG_FILE_getByWhere
}

func (data *MSG_FILE_getByWhere) Put() {
	data.Where = nil
	data.Page = 0
	data.PerPage = 0
	data.Total = 0
	pool_MSG_FILE_getByWhere.Put(data)
}
func (data *MSG_FILE_getByWhere) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_getByWhere,buf)
	WRITE_MSG_FILE_getByWhere(data, buf)
}

func WRITE_MSG_FILE_getByWhere(data *MSG_FILE_getByWhere, buf *libraries.MsgBuffer) {
	WRITE_map(data.Where,buf)
	WRITE_int(data.Page, buf)
	WRITE_int(data.PerPage, buf)
	WRITE_int(data.Total, buf)
}

func READ_MSG_FILE_getByWhere(buf *libraries.MsgBuffer) *MSG_FILE_getByWhere {
	data := pool_MSG_FILE_getByWhere.Get().(*MSG_FILE_getByWhere)
	data.read(buf)
	return data
}

func (data *MSG_FILE_getByWhere) read(buf *libraries.MsgBuffer) {
	READ_map(&data.Where,buf)
	data.Page = READ_int(buf)
	data.PerPage = READ_int(buf)
	data.Total = READ_int(buf)

}

type MSG_FILE_getByWhere_result struct {
	List []*MSG_FILE_getByID_result
	Total int
}

var pool_MSG_FILE_getByWhere_result = sync.Pool{New: func() interface{} { return &MSG_FILE_getByWhere_result{} }}

func GET_MSG_FILE_getByWhere_result() *MSG_FILE_getByWhere_result {
	return pool_MSG_FILE_getByWhere_result.Get().(*MSG_FILE_getByWhere_result)
}

func (data *MSG_FILE_getByWhere_result) cmd() int32 {
	return CMD_MSG_FILE_getByWhere_result
}

func (data *MSG_FILE_getByWhere_result) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	data.Total = 0
	pool_MSG_FILE_getByWhere_result.Put(data)
}
func (data *MSG_FILE_getByWhere_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_getByWhere_result,buf)
	WRITE_MSG_FILE_getByWhere_result(data, buf)
}

func WRITE_MSG_FILE_getByWhere_result(data *MSG_FILE_getByWhere_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_FILE_getByID_result(v, buf)
	}
	WRITE_int(data.Total, buf)
}

func READ_MSG_FILE_getByWhere_result(buf *libraries.MsgBuffer) *MSG_FILE_getByWhere_result {
	data := pool_MSG_FILE_getByWhere_result.Get().(*MSG_FILE_getByWhere_result)
	data.read(buf)
	return data
}

func (data *MSG_FILE_getByWhere_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_FILE_getByID_result, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_FILE_getByID_result(buf)
	}
	data.Total = READ_int(buf)

}

type MSG_FILE_download_byIds struct {
	Ids []int64
}

var pool_MSG_FILE_download_byIds = sync.Pool{New: func() interface{} { return &MSG_FILE_download_byIds{} }}

func GET_MSG_FILE_download_byIds() *MSG_FILE_download_byIds {
	return pool_MSG_FILE_download_byIds.Get().(*MSG_FILE_download_byIds)
}

func (data *MSG_FILE_download_byIds) cmd() int32 {
	return CMD_MSG_FILE_download_byIds
}

func (data *MSG_FILE_download_byIds) Put() {
	data.Ids = data.Ids[:0]
	pool_MSG_FILE_download_byIds.Put(data)
}
func (data *MSG_FILE_download_byIds) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_download_byIds,buf)
	WRITE_MSG_FILE_download_byIds(data, buf)
}

func WRITE_MSG_FILE_download_byIds(data *MSG_FILE_download_byIds, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.Ids), buf)
	for _, v := range data.Ids{
		WRITE_int64(v, buf)
	}
}

func READ_MSG_FILE_download_byIds(buf *libraries.MsgBuffer) *MSG_FILE_download_byIds {
	data := pool_MSG_FILE_download_byIds.Get().(*MSG_FILE_download_byIds)
	data.read(buf)
	return data
}

func (data *MSG_FILE_download_byIds) read(buf *libraries.MsgBuffer) {
	Ids_len := READ_int(buf)
	if Ids_len>cap(data.Ids){
		data.Ids= make([]int64, Ids_len)
	}else{
		data.Ids = data.Ids[:Ids_len]
	}
	for i := 0; i < Ids_len; i++ {
		data.Ids[i] = READ_int64(buf)
	}

}

type MSG_FILE_download_byIds_result struct {
	FileID int64
}

var pool_MSG_FILE_download_byIds_result = sync.Pool{New: func() interface{} { return &MSG_FILE_download_byIds_result{} }}

func GET_MSG_FILE_download_byIds_result() *MSG_FILE_download_byIds_result {
	return pool_MSG_FILE_download_byIds_result.Get().(*MSG_FILE_download_byIds_result)
}

func (data *MSG_FILE_download_byIds_result) cmd() int32 {
	return CMD_MSG_FILE_download_byIds_result
}

func (data *MSG_FILE_download_byIds_result) Put() {
	data.FileID = 0
	pool_MSG_FILE_download_byIds_result.Put(data)
}
func (data *MSG_FILE_download_byIds_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_download_byIds_result,buf)
	WRITE_MSG_FILE_download_byIds_result(data, buf)
}

func WRITE_MSG_FILE_download_byIds_result(data *MSG_FILE_download_byIds_result, buf *libraries.MsgBuffer) {
	WRITE_int64(data.FileID, buf)
}

func READ_MSG_FILE_download_byIds_result(buf *libraries.MsgBuffer) *MSG_FILE_download_byIds_result {
	data := pool_MSG_FILE_download_byIds_result.Get().(*MSG_FILE_download_byIds_result)
	data.read(buf)
	return data
}

func (data *MSG_FILE_download_byIds_result) read(buf *libraries.MsgBuffer) {
	data.FileID = READ_int64(buf)

}

