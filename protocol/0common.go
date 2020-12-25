package protocol

import (
	"sync"
	"libraries"
)

const (
	CMD_MSG_COMMON_regServer = 1508627456
	CMD_MSG_COMMON_regServer_result = -1423609088
	CMD_MSG_COMMON_StartTicker = 845234432
	CMD_MSG_COMMON_PING = -2041209088
	CMD_MSG_COMMON_PONG = -2099715328
	CMD_MSG_COMMON_WINDOW_UPDATE = 1835280896
	CMD_MSG_COMMON_CACHE_GET = -1094191872
	CMD_MSG_COMMON_CACHE_GET_result = 503974656
	CMD_MSG_COMMON_CACHE_GETPATH = -1961335296
	CMD_MSG_COMMON_CACHE_GETPATH_result = -1979298816
	CMD_MSG_COMMON_CACHE_SET = -1511457280
	CMD_MSG_COMMON_CACHE_DEL = -1343368448
	CMD_MSG_COMMON_CACHE_DelPath = -1163212800
	CMD_MSG_COMMON_GET_Msgno = -26660096
	CMD_MSG_COMMON_GET_Msgno_result = -1751634176
	CMD_MSG_COMMON_QueryErr = 714684672
	CMD_MSG_COMMON_ResetWindow = 988600064
	CMD_MSG_FILE_upload = 1110878976
	CMD_MSG_FILE_upload_result = -2057389056
	CMD_MSG_FILE_getByID = -1871273216
	CMD_MSG_FILE_getByID_result = -1484540160
	CMD_MSG_FILE_updateByIDMap = -1777243392
	CMD_MSG_FILE_DeleteByID = -1568529408
)

type MSG_COMMON_regServer struct {
	No uint8
	IpPort string
	Time int64
	Token string
	Window int32
}

var pool_MSG_COMMON_regServer = sync.Pool{New: func() interface{} { return &MSG_COMMON_regServer{} }}

func GET_MSG_COMMON_regServer() *MSG_COMMON_regServer {
	return pool_MSG_COMMON_regServer.Get().(*MSG_COMMON_regServer)
}

func (data *MSG_COMMON_regServer) cmd() int32 {
	return CMD_MSG_COMMON_regServer
}

func (data *MSG_COMMON_regServer) Put() {
	data.No = 0
	data.IpPort = ``
	data.Time = 0
	data.Token = ``
	data.Window = 0
	pool_MSG_COMMON_regServer.Put(data)
}
func (data *MSG_COMMON_regServer) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_regServer,buf)
	WRITE_MSG_COMMON_regServer(data, buf)
}

func WRITE_MSG_COMMON_regServer(data *MSG_COMMON_regServer, buf *libraries.MsgBuffer) {
	WRITE_uint8(data.No, buf)
	WRITE_string(data.IpPort, buf)
	WRITE_int64(data.Time, buf)
	WRITE_string(data.Token, buf)
	WRITE_int32(data.Window, buf)
}

func READ_MSG_COMMON_regServer(buf *libraries.MsgBuffer) *MSG_COMMON_regServer {
	data := pool_MSG_COMMON_regServer.Get().(*MSG_COMMON_regServer)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_regServer) read(buf *libraries.MsgBuffer) {
	data.No = READ_uint8(buf)
	data.IpPort = READ_string(buf)
	data.Time = READ_int64(buf)
	data.Token = READ_string(buf)
	data.Window = READ_int32(buf)

}

type MSG_COMMON_regServer_result struct {
	Id uint8
}

var pool_MSG_COMMON_regServer_result = sync.Pool{New: func() interface{} { return &MSG_COMMON_regServer_result{} }}

func GET_MSG_COMMON_regServer_result() *MSG_COMMON_regServer_result {
	return pool_MSG_COMMON_regServer_result.Get().(*MSG_COMMON_regServer_result)
}

func (data *MSG_COMMON_regServer_result) cmd() int32 {
	return CMD_MSG_COMMON_regServer_result
}

func (data *MSG_COMMON_regServer_result) Put() {
	data.Id = 0
	pool_MSG_COMMON_regServer_result.Put(data)
}
func (data *MSG_COMMON_regServer_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_regServer_result,buf)
	WRITE_MSG_COMMON_regServer_result(data, buf)
}

func WRITE_MSG_COMMON_regServer_result(data *MSG_COMMON_regServer_result, buf *libraries.MsgBuffer) {
	WRITE_uint8(data.Id, buf)
}

func READ_MSG_COMMON_regServer_result(buf *libraries.MsgBuffer) *MSG_COMMON_regServer_result {
	data := pool_MSG_COMMON_regServer_result.Get().(*MSG_COMMON_regServer_result)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_regServer_result) read(buf *libraries.MsgBuffer) {
	data.Id = READ_uint8(buf)

}

type MSG_COMMON_StartTicker struct {
}

var pool_MSG_COMMON_StartTicker = sync.Pool{New: func() interface{} { return &MSG_COMMON_StartTicker{} }}

func GET_MSG_COMMON_StartTicker() *MSG_COMMON_StartTicker {
	return pool_MSG_COMMON_StartTicker.Get().(*MSG_COMMON_StartTicker)
}

func (data *MSG_COMMON_StartTicker) cmd() int32 {
	return CMD_MSG_COMMON_StartTicker
}

func (data *MSG_COMMON_StartTicker) Put() {
	pool_MSG_COMMON_StartTicker.Put(data)
}
func (data *MSG_COMMON_StartTicker) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_StartTicker,buf)
	WRITE_MSG_COMMON_StartTicker(data, buf)
}

func WRITE_MSG_COMMON_StartTicker(data *MSG_COMMON_StartTicker, buf *libraries.MsgBuffer) {
}

func READ_MSG_COMMON_StartTicker(buf *libraries.MsgBuffer) *MSG_COMMON_StartTicker {
	data := pool_MSG_COMMON_StartTicker.Get().(*MSG_COMMON_StartTicker)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_StartTicker) read(buf *libraries.MsgBuffer) {

}

type MSG_COMMON_PING struct {
}

var pool_MSG_COMMON_PING = sync.Pool{New: func() interface{} { return &MSG_COMMON_PING{} }}

func GET_MSG_COMMON_PING() *MSG_COMMON_PING {
	return pool_MSG_COMMON_PING.Get().(*MSG_COMMON_PING)
}

func (data *MSG_COMMON_PING) cmd() int32 {
	return CMD_MSG_COMMON_PING
}

func (data *MSG_COMMON_PING) Put() {
	pool_MSG_COMMON_PING.Put(data)
}
func (data *MSG_COMMON_PING) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_PING,buf)
	WRITE_MSG_COMMON_PING(data, buf)
}

func WRITE_MSG_COMMON_PING(data *MSG_COMMON_PING, buf *libraries.MsgBuffer) {
}

func READ_MSG_COMMON_PING(buf *libraries.MsgBuffer) *MSG_COMMON_PING {
	data := pool_MSG_COMMON_PING.Get().(*MSG_COMMON_PING)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_PING) read(buf *libraries.MsgBuffer) {

}

type MSG_COMMON_PONG struct {
}

var pool_MSG_COMMON_PONG = sync.Pool{New: func() interface{} { return &MSG_COMMON_PONG{} }}

func GET_MSG_COMMON_PONG() *MSG_COMMON_PONG {
	return pool_MSG_COMMON_PONG.Get().(*MSG_COMMON_PONG)
}

func (data *MSG_COMMON_PONG) cmd() int32 {
	return CMD_MSG_COMMON_PONG
}

func (data *MSG_COMMON_PONG) Put() {
	pool_MSG_COMMON_PONG.Put(data)
}
func (data *MSG_COMMON_PONG) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_PONG,buf)
	WRITE_MSG_COMMON_PONG(data, buf)
}

func WRITE_MSG_COMMON_PONG(data *MSG_COMMON_PONG, buf *libraries.MsgBuffer) {
}

func READ_MSG_COMMON_PONG(buf *libraries.MsgBuffer) *MSG_COMMON_PONG {
	data := pool_MSG_COMMON_PONG.Get().(*MSG_COMMON_PONG)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_PONG) read(buf *libraries.MsgBuffer) {

}

type MSG_COMMON_WINDOW_UPDATE struct {
	Add int32
}

var pool_MSG_COMMON_WINDOW_UPDATE = sync.Pool{New: func() interface{} { return &MSG_COMMON_WINDOW_UPDATE{} }}

func GET_MSG_COMMON_WINDOW_UPDATE() *MSG_COMMON_WINDOW_UPDATE {
	return pool_MSG_COMMON_WINDOW_UPDATE.Get().(*MSG_COMMON_WINDOW_UPDATE)
}

func (data *MSG_COMMON_WINDOW_UPDATE) cmd() int32 {
	return CMD_MSG_COMMON_WINDOW_UPDATE
}

func (data *MSG_COMMON_WINDOW_UPDATE) Put() {
	data.Add = 0
	pool_MSG_COMMON_WINDOW_UPDATE.Put(data)
}
func (data *MSG_COMMON_WINDOW_UPDATE) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_WINDOW_UPDATE,buf)
	WRITE_MSG_COMMON_WINDOW_UPDATE(data, buf)
}

func WRITE_MSG_COMMON_WINDOW_UPDATE(data *MSG_COMMON_WINDOW_UPDATE, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Add, buf)
}

func READ_MSG_COMMON_WINDOW_UPDATE(buf *libraries.MsgBuffer) *MSG_COMMON_WINDOW_UPDATE {
	data := pool_MSG_COMMON_WINDOW_UPDATE.Get().(*MSG_COMMON_WINDOW_UPDATE)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_WINDOW_UPDATE) read(buf *libraries.MsgBuffer) {
	data.Add = READ_int32(buf)

}

type MSG_COMMON_CACHE_GET struct {
	QueryID uint32
	Path string
	Name string
}

var pool_MSG_COMMON_CACHE_GET = sync.Pool{New: func() interface{} { return &MSG_COMMON_CACHE_GET{} }}

func GET_MSG_COMMON_CACHE_GET() *MSG_COMMON_CACHE_GET {
	return pool_MSG_COMMON_CACHE_GET.Get().(*MSG_COMMON_CACHE_GET)
}

func (data *MSG_COMMON_CACHE_GET) cmd() int32 {
	return CMD_MSG_COMMON_CACHE_GET
}

func (data *MSG_COMMON_CACHE_GET) Put() {
	data.QueryID = 0
	data.Path = ``
	data.Name = ``
	pool_MSG_COMMON_CACHE_GET.Put(data)
}
func (data *MSG_COMMON_CACHE_GET) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_CACHE_GET,buf)
	WRITE_MSG_COMMON_CACHE_GET(data, buf)
}

func WRITE_MSG_COMMON_CACHE_GET(data *MSG_COMMON_CACHE_GET, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_string(data.Path, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_COMMON_CACHE_GET(buf *libraries.MsgBuffer) *MSG_COMMON_CACHE_GET {
	data := pool_MSG_COMMON_CACHE_GET.Get().(*MSG_COMMON_CACHE_GET)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_CACHE_GET) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Path = READ_string(buf)
	data.Name = READ_string(buf)

}
func (data *MSG_COMMON_CACHE_GET) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_COMMON_CACHE_GET) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_COMMON_CACHE_GET_result struct {
	QueryResultID uint32
	Value []byte
}

var pool_MSG_COMMON_CACHE_GET_result = sync.Pool{New: func() interface{} { return &MSG_COMMON_CACHE_GET_result{} }}

func GET_MSG_COMMON_CACHE_GET_result() *MSG_COMMON_CACHE_GET_result {
	return pool_MSG_COMMON_CACHE_GET_result.Get().(*MSG_COMMON_CACHE_GET_result)
}

func (data *MSG_COMMON_CACHE_GET_result) cmd() int32 {
	return CMD_MSG_COMMON_CACHE_GET_result
}

func (data *MSG_COMMON_CACHE_GET_result) Put() {
	data.QueryResultID = 0
	data.Value = data.Value[:0]
	pool_MSG_COMMON_CACHE_GET_result.Put(data)
}
func (data *MSG_COMMON_CACHE_GET_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_CACHE_GET_result,buf)
	WRITE_MSG_COMMON_CACHE_GET_result(data, buf)
}

func WRITE_MSG_COMMON_CACHE_GET_result(data *MSG_COMMON_CACHE_GET_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.Value)), buf)
	buf.Write(data.Value)
}

func READ_MSG_COMMON_CACHE_GET_result(buf *libraries.MsgBuffer) *MSG_COMMON_CACHE_GET_result {
	data := pool_MSG_COMMON_CACHE_GET_result.Get().(*MSG_COMMON_CACHE_GET_result)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_CACHE_GET_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	Value_len := int(READ_int32(buf))
	data.Value = make([]byte, Value_len)
	copy(data.Value,buf.Next(Value_len))

}
func (data *MSG_COMMON_CACHE_GET_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_COMMON_CACHE_GET_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_COMMON_CACHE_GETPATH struct {
	QueryID uint32
	Path string
}

var pool_MSG_COMMON_CACHE_GETPATH = sync.Pool{New: func() interface{} { return &MSG_COMMON_CACHE_GETPATH{} }}

func GET_MSG_COMMON_CACHE_GETPATH() *MSG_COMMON_CACHE_GETPATH {
	return pool_MSG_COMMON_CACHE_GETPATH.Get().(*MSG_COMMON_CACHE_GETPATH)
}

func (data *MSG_COMMON_CACHE_GETPATH) cmd() int32 {
	return CMD_MSG_COMMON_CACHE_GETPATH
}

func (data *MSG_COMMON_CACHE_GETPATH) Put() {
	data.QueryID = 0
	data.Path = ``
	pool_MSG_COMMON_CACHE_GETPATH.Put(data)
}
func (data *MSG_COMMON_CACHE_GETPATH) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_CACHE_GETPATH,buf)
	WRITE_MSG_COMMON_CACHE_GETPATH(data, buf)
}

func WRITE_MSG_COMMON_CACHE_GETPATH(data *MSG_COMMON_CACHE_GETPATH, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_string(data.Path, buf)
}

func READ_MSG_COMMON_CACHE_GETPATH(buf *libraries.MsgBuffer) *MSG_COMMON_CACHE_GETPATH {
	data := pool_MSG_COMMON_CACHE_GETPATH.Get().(*MSG_COMMON_CACHE_GETPATH)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_CACHE_GETPATH) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Path = READ_string(buf)

}
func (data *MSG_COMMON_CACHE_GETPATH) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_COMMON_CACHE_GETPATH) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_COMMON_CACHE_GETPATH_result struct {
	QueryResultID uint32
	Value [][]byte
}

var pool_MSG_COMMON_CACHE_GETPATH_result = sync.Pool{New: func() interface{} { return &MSG_COMMON_CACHE_GETPATH_result{} }}

func GET_MSG_COMMON_CACHE_GETPATH_result() *MSG_COMMON_CACHE_GETPATH_result {
	return pool_MSG_COMMON_CACHE_GETPATH_result.Get().(*MSG_COMMON_CACHE_GETPATH_result)
}

func (data *MSG_COMMON_CACHE_GETPATH_result) cmd() int32 {
	return CMD_MSG_COMMON_CACHE_GETPATH_result
}

func (data *MSG_COMMON_CACHE_GETPATH_result) Put() {
	data.QueryResultID = 0
	data.Value = data.Value[:0]
	pool_MSG_COMMON_CACHE_GETPATH_result.Put(data)
}
func (data *MSG_COMMON_CACHE_GETPATH_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_CACHE_GETPATH_result,buf)
	WRITE_MSG_COMMON_CACHE_GETPATH_result(data, buf)
}

func WRITE_MSG_COMMON_CACHE_GETPATH_result(data *MSG_COMMON_CACHE_GETPATH_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.Value)), buf)
	for _, v := range data.Value{
		WRITE_int32(int32(len(v)), buf)
		buf.Write(v)
	}
}

func READ_MSG_COMMON_CACHE_GETPATH_result(buf *libraries.MsgBuffer) *MSG_COMMON_CACHE_GETPATH_result {
	data := pool_MSG_COMMON_CACHE_GETPATH_result.Get().(*MSG_COMMON_CACHE_GETPATH_result)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_CACHE_GETPATH_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	Value_len := int(READ_int32(buf))
	for i := 0; i < Value_len; i++ {
		l := READ_int32(buf)
		b := make([]byte,l)
		copy(b,buf.Next(int(l)))
		data.Value = append(data.Value, b)
	}

}
func (data *MSG_COMMON_CACHE_GETPATH_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_COMMON_CACHE_GETPATH_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_COMMON_CACHE_SET struct {
	Path string
	Name string
	Value []byte
	Expire int64
}

var pool_MSG_COMMON_CACHE_SET = sync.Pool{New: func() interface{} { return &MSG_COMMON_CACHE_SET{} }}

func GET_MSG_COMMON_CACHE_SET() *MSG_COMMON_CACHE_SET {
	return pool_MSG_COMMON_CACHE_SET.Get().(*MSG_COMMON_CACHE_SET)
}

func (data *MSG_COMMON_CACHE_SET) cmd() int32 {
	return CMD_MSG_COMMON_CACHE_SET
}

func (data *MSG_COMMON_CACHE_SET) Put() {
	data.Path = ``
	data.Name = ``
	data.Value = data.Value[:0]
	data.Expire = 0
	pool_MSG_COMMON_CACHE_SET.Put(data)
}
func (data *MSG_COMMON_CACHE_SET) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_CACHE_SET,buf)
	WRITE_MSG_COMMON_CACHE_SET(data, buf)
}

func WRITE_MSG_COMMON_CACHE_SET(data *MSG_COMMON_CACHE_SET, buf *libraries.MsgBuffer) {
	WRITE_string(data.Path, buf)
	WRITE_string(data.Name, buf)
	WRITE_int32(int32(len(data.Value)), buf)
	buf.Write(data.Value)
	WRITE_int64(data.Expire, buf)
}

func READ_MSG_COMMON_CACHE_SET(buf *libraries.MsgBuffer) *MSG_COMMON_CACHE_SET {
	data := pool_MSG_COMMON_CACHE_SET.Get().(*MSG_COMMON_CACHE_SET)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_CACHE_SET) read(buf *libraries.MsgBuffer) {
	data.Path = READ_string(buf)
	data.Name = READ_string(buf)
	Value_len := int(READ_int32(buf))
	data.Value = make([]byte, Value_len)
	copy(data.Value,buf.Next(Value_len))
	data.Expire = READ_int64(buf)

}

type MSG_COMMON_CACHE_DEL struct {
	Path string
	Name string
}

var pool_MSG_COMMON_CACHE_DEL = sync.Pool{New: func() interface{} { return &MSG_COMMON_CACHE_DEL{} }}

func GET_MSG_COMMON_CACHE_DEL() *MSG_COMMON_CACHE_DEL {
	return pool_MSG_COMMON_CACHE_DEL.Get().(*MSG_COMMON_CACHE_DEL)
}

func (data *MSG_COMMON_CACHE_DEL) cmd() int32 {
	return CMD_MSG_COMMON_CACHE_DEL
}

func (data *MSG_COMMON_CACHE_DEL) Put() {
	data.Path = ``
	data.Name = ``
	pool_MSG_COMMON_CACHE_DEL.Put(data)
}
func (data *MSG_COMMON_CACHE_DEL) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_CACHE_DEL,buf)
	WRITE_MSG_COMMON_CACHE_DEL(data, buf)
}

func WRITE_MSG_COMMON_CACHE_DEL(data *MSG_COMMON_CACHE_DEL, buf *libraries.MsgBuffer) {
	WRITE_string(data.Path, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_COMMON_CACHE_DEL(buf *libraries.MsgBuffer) *MSG_COMMON_CACHE_DEL {
	data := pool_MSG_COMMON_CACHE_DEL.Get().(*MSG_COMMON_CACHE_DEL)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_CACHE_DEL) read(buf *libraries.MsgBuffer) {
	data.Path = READ_string(buf)
	data.Name = READ_string(buf)

}

type MSG_COMMON_CACHE_DelPath struct {
	Path string
}

var pool_MSG_COMMON_CACHE_DelPath = sync.Pool{New: func() interface{} { return &MSG_COMMON_CACHE_DelPath{} }}

func GET_MSG_COMMON_CACHE_DelPath() *MSG_COMMON_CACHE_DelPath {
	return pool_MSG_COMMON_CACHE_DelPath.Get().(*MSG_COMMON_CACHE_DelPath)
}

func (data *MSG_COMMON_CACHE_DelPath) cmd() int32 {
	return CMD_MSG_COMMON_CACHE_DelPath
}

func (data *MSG_COMMON_CACHE_DelPath) Put() {
	data.Path = ``
	pool_MSG_COMMON_CACHE_DelPath.Put(data)
}
func (data *MSG_COMMON_CACHE_DelPath) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_CACHE_DelPath,buf)
	WRITE_MSG_COMMON_CACHE_DelPath(data, buf)
}

func WRITE_MSG_COMMON_CACHE_DelPath(data *MSG_COMMON_CACHE_DelPath, buf *libraries.MsgBuffer) {
	WRITE_string(data.Path, buf)
}

func READ_MSG_COMMON_CACHE_DelPath(buf *libraries.MsgBuffer) *MSG_COMMON_CACHE_DelPath {
	data := pool_MSG_COMMON_CACHE_DelPath.Get().(*MSG_COMMON_CACHE_DelPath)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_CACHE_DelPath) read(buf *libraries.MsgBuffer) {
	data.Path = READ_string(buf)

}

type MSG_COMMON_GET_Msgno struct {
	QueryID uint32
}

var pool_MSG_COMMON_GET_Msgno = sync.Pool{New: func() interface{} { return &MSG_COMMON_GET_Msgno{} }}

func GET_MSG_COMMON_GET_Msgno() *MSG_COMMON_GET_Msgno {
	return pool_MSG_COMMON_GET_Msgno.Get().(*MSG_COMMON_GET_Msgno)
}

func (data *MSG_COMMON_GET_Msgno) cmd() int32 {
	return CMD_MSG_COMMON_GET_Msgno
}

func (data *MSG_COMMON_GET_Msgno) Put() {
	data.QueryID = 0
	pool_MSG_COMMON_GET_Msgno.Put(data)
}
func (data *MSG_COMMON_GET_Msgno) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_GET_Msgno,buf)
	WRITE_MSG_COMMON_GET_Msgno(data, buf)
}

func WRITE_MSG_COMMON_GET_Msgno(data *MSG_COMMON_GET_Msgno, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
}

func READ_MSG_COMMON_GET_Msgno(buf *libraries.MsgBuffer) *MSG_COMMON_GET_Msgno {
	data := pool_MSG_COMMON_GET_Msgno.Get().(*MSG_COMMON_GET_Msgno)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_GET_Msgno) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)

}
func (data *MSG_COMMON_GET_Msgno) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_COMMON_GET_Msgno) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_COMMON_GET_Msgno_result struct {
	QueryResultID uint32
	Msgno uint32
}

var pool_MSG_COMMON_GET_Msgno_result = sync.Pool{New: func() interface{} { return &MSG_COMMON_GET_Msgno_result{} }}

func GET_MSG_COMMON_GET_Msgno_result() *MSG_COMMON_GET_Msgno_result {
	return pool_MSG_COMMON_GET_Msgno_result.Get().(*MSG_COMMON_GET_Msgno_result)
}

func (data *MSG_COMMON_GET_Msgno_result) cmd() int32 {
	return CMD_MSG_COMMON_GET_Msgno_result
}

func (data *MSG_COMMON_GET_Msgno_result) Put() {
	data.QueryResultID = 0
	data.Msgno = 0
	pool_MSG_COMMON_GET_Msgno_result.Put(data)
}
func (data *MSG_COMMON_GET_Msgno_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_GET_Msgno_result,buf)
	WRITE_MSG_COMMON_GET_Msgno_result(data, buf)
}

func WRITE_MSG_COMMON_GET_Msgno_result(data *MSG_COMMON_GET_Msgno_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_uint32(data.Msgno, buf)
}

func READ_MSG_COMMON_GET_Msgno_result(buf *libraries.MsgBuffer) *MSG_COMMON_GET_Msgno_result {
	data := pool_MSG_COMMON_GET_Msgno_result.Get().(*MSG_COMMON_GET_Msgno_result)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_GET_Msgno_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Msgno = READ_uint32(buf)

}
func (data *MSG_COMMON_GET_Msgno_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_COMMON_GET_Msgno_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_COMMON_QueryErr struct {
	QueryResultID uint32
	Err string
	Stack []byte
}

var pool_MSG_COMMON_QueryErr = sync.Pool{New: func() interface{} { return &MSG_COMMON_QueryErr{} }}

func GET_MSG_COMMON_QueryErr() *MSG_COMMON_QueryErr {
	return pool_MSG_COMMON_QueryErr.Get().(*MSG_COMMON_QueryErr)
}

func (data *MSG_COMMON_QueryErr) cmd() int32 {
	return CMD_MSG_COMMON_QueryErr
}

func (data *MSG_COMMON_QueryErr) Put() {
	data.QueryResultID = 0
	data.Err = ``
	data.Stack = data.Stack[:0]
	pool_MSG_COMMON_QueryErr.Put(data)
}
func (data *MSG_COMMON_QueryErr) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_QueryErr,buf)
	WRITE_MSG_COMMON_QueryErr(data, buf)
}

func WRITE_MSG_COMMON_QueryErr(data *MSG_COMMON_QueryErr, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_string(data.Err, buf)
	WRITE_int32(int32(len(data.Stack)), buf)
	buf.Write(data.Stack)
}

func READ_MSG_COMMON_QueryErr(buf *libraries.MsgBuffer) *MSG_COMMON_QueryErr {
	data := pool_MSG_COMMON_QueryErr.Get().(*MSG_COMMON_QueryErr)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_QueryErr) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Err = READ_string(buf)
	Stack_len := int(READ_int32(buf))
	data.Stack = make([]byte, Stack_len)
	copy(data.Stack,buf.Next(Stack_len))

}
func (data *MSG_COMMON_QueryErr) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_COMMON_QueryErr) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_COMMON_ResetWindow struct {
	Window int32
}

var pool_MSG_COMMON_ResetWindow = sync.Pool{New: func() interface{} { return &MSG_COMMON_ResetWindow{} }}

func GET_MSG_COMMON_ResetWindow() *MSG_COMMON_ResetWindow {
	return pool_MSG_COMMON_ResetWindow.Get().(*MSG_COMMON_ResetWindow)
}

func (data *MSG_COMMON_ResetWindow) cmd() int32 {
	return CMD_MSG_COMMON_ResetWindow
}

func (data *MSG_COMMON_ResetWindow) Put() {
	data.Window = 0
	pool_MSG_COMMON_ResetWindow.Put(data)
}
func (data *MSG_COMMON_ResetWindow) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_COMMON_ResetWindow,buf)
	WRITE_MSG_COMMON_ResetWindow(data, buf)
}

func WRITE_MSG_COMMON_ResetWindow(data *MSG_COMMON_ResetWindow, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Window, buf)
}

func READ_MSG_COMMON_ResetWindow(buf *libraries.MsgBuffer) *MSG_COMMON_ResetWindow {
	data := pool_MSG_COMMON_ResetWindow.Get().(*MSG_COMMON_ResetWindow)
	data.read(buf)
	return data
}

func (data *MSG_COMMON_ResetWindow) read(buf *libraries.MsgBuffer) {
	data.Window = READ_int32(buf)

}

type MSG_FILE_upload struct {
	QueryID uint32
	Name string
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
	data.QueryID = 0
	data.Name = ``
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
	WRITE_uint32(data.QueryID, buf)
	WRITE_string(data.Name, buf)
	WRITE_int32(int32(len(data.Data)), buf)
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
	data.QueryID = READ_uint32(buf)
	data.Name = READ_string(buf)
	Data_len := int(READ_int32(buf))
	data.Data = make([]byte, Data_len)
	copy(data.Data,buf.Next(Data_len))
	data.AddBy = READ_int32(buf)
	data.ObjectType = READ_string(buf)
	data.ObjectID = READ_int32(buf)
	data.Code = READ_string(buf)
	data.Type = READ_string(buf)

}
func (data *MSG_FILE_upload) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_FILE_upload) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_FILE_upload_result struct {
	QueryResultID uint32
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
	data.QueryResultID = 0
	data.FileID = 0
	pool_MSG_FILE_upload_result.Put(data)
}
func (data *MSG_FILE_upload_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_upload_result,buf)
	WRITE_MSG_FILE_upload_result(data, buf)
}

func WRITE_MSG_FILE_upload_result(data *MSG_FILE_upload_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int64(data.FileID, buf)
}

func READ_MSG_FILE_upload_result(buf *libraries.MsgBuffer) *MSG_FILE_upload_result {
	data := pool_MSG_FILE_upload_result.Get().(*MSG_FILE_upload_result)
	data.read(buf)
	return data
}

func (data *MSG_FILE_upload_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.FileID = READ_int64(buf)

}
func (data *MSG_FILE_upload_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_FILE_upload_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_FILE_getByID struct {
	QueryID uint32
	FileID int64
}

var pool_MSG_FILE_getByID = sync.Pool{New: func() interface{} { return &MSG_FILE_getByID{} }}

func GET_MSG_FILE_getByID() *MSG_FILE_getByID {
	return pool_MSG_FILE_getByID.Get().(*MSG_FILE_getByID)
}

func (data *MSG_FILE_getByID) cmd() int32 {
	return CMD_MSG_FILE_getByID
}

func (data *MSG_FILE_getByID) Put() {
	data.QueryID = 0
	data.FileID = 0
	pool_MSG_FILE_getByID.Put(data)
}
func (data *MSG_FILE_getByID) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_getByID,buf)
	WRITE_MSG_FILE_getByID(data, buf)
}

func WRITE_MSG_FILE_getByID(data *MSG_FILE_getByID, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int64(data.FileID, buf)
}

func READ_MSG_FILE_getByID(buf *libraries.MsgBuffer) *MSG_FILE_getByID {
	data := pool_MSG_FILE_getByID.Get().(*MSG_FILE_getByID)
	data.read(buf)
	return data
}

func (data *MSG_FILE_getByID) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.FileID = READ_int64(buf)

}
func (data *MSG_FILE_getByID) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_FILE_getByID) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_FILE_getByID_result struct {
	QueryResultID uint32
	Name string
	Ext string
	Data []byte
	Type string
}

var pool_MSG_FILE_getByID_result = sync.Pool{New: func() interface{} { return &MSG_FILE_getByID_result{} }}

func GET_MSG_FILE_getByID_result() *MSG_FILE_getByID_result {
	return pool_MSG_FILE_getByID_result.Get().(*MSG_FILE_getByID_result)
}

func (data *MSG_FILE_getByID_result) cmd() int32 {
	return CMD_MSG_FILE_getByID_result
}

func (data *MSG_FILE_getByID_result) Put() {
	data.QueryResultID = 0
	data.Name = ``
	data.Ext = ``
	data.Data = data.Data[:0]
	data.Type = ``
	pool_MSG_FILE_getByID_result.Put(data)
}
func (data *MSG_FILE_getByID_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_getByID_result,buf)
	WRITE_MSG_FILE_getByID_result(data, buf)
}

func WRITE_MSG_FILE_getByID_result(data *MSG_FILE_getByID_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Ext, buf)
	WRITE_int32(int32(len(data.Data)), buf)
	buf.Write(data.Data)
	WRITE_string(data.Type, buf)
}

func READ_MSG_FILE_getByID_result(buf *libraries.MsgBuffer) *MSG_FILE_getByID_result {
	data := pool_MSG_FILE_getByID_result.Get().(*MSG_FILE_getByID_result)
	data.read(buf)
	return data
}

func (data *MSG_FILE_getByID_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Name = READ_string(buf)
	data.Ext = READ_string(buf)
	Data_len := int(READ_int32(buf))
	data.Data = make([]byte, Data_len)
	copy(data.Data,buf.Next(Data_len))
	data.Type = READ_string(buf)

}
func (data *MSG_FILE_getByID_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_FILE_getByID_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_FILE_updateByIDMap struct {
	QueryID uint32
	FileID int64
	Update map[string]string
}

var pool_MSG_FILE_updateByIDMap = sync.Pool{New: func() interface{} { return &MSG_FILE_updateByIDMap{} }}

func GET_MSG_FILE_updateByIDMap() *MSG_FILE_updateByIDMap {
	return pool_MSG_FILE_updateByIDMap.Get().(*MSG_FILE_updateByIDMap)
}

func (data *MSG_FILE_updateByIDMap) cmd() int32 {
	return CMD_MSG_FILE_updateByIDMap
}

func (data *MSG_FILE_updateByIDMap) Put() {
	data.QueryID = 0
	data.FileID = 0
	data.Update = nil
	pool_MSG_FILE_updateByIDMap.Put(data)
}
func (data *MSG_FILE_updateByIDMap) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_updateByIDMap,buf)
	WRITE_MSG_FILE_updateByIDMap(data, buf)
}

func WRITE_MSG_FILE_updateByIDMap(data *MSG_FILE_updateByIDMap, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int64(data.FileID, buf)
	WRITE_map(data.Update,buf)
}

func READ_MSG_FILE_updateByIDMap(buf *libraries.MsgBuffer) *MSG_FILE_updateByIDMap {
	data := pool_MSG_FILE_updateByIDMap.Get().(*MSG_FILE_updateByIDMap)
	data.read(buf)
	return data
}

func (data *MSG_FILE_updateByIDMap) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.FileID = READ_int64(buf)
	READ_map(&data.Update,buf)

}
func (data *MSG_FILE_updateByIDMap) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_FILE_updateByIDMap) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_FILE_DeleteByID struct {
	QueryID uint32
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
	data.QueryID = 0
	data.FileID = 0
	pool_MSG_FILE_DeleteByID.Put(data)
}
func (data *MSG_FILE_DeleteByID) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_FILE_DeleteByID,buf)
	WRITE_MSG_FILE_DeleteByID(data, buf)
}

func WRITE_MSG_FILE_DeleteByID(data *MSG_FILE_DeleteByID, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int64(data.FileID, buf)
}

func READ_MSG_FILE_DeleteByID(buf *libraries.MsgBuffer) *MSG_FILE_DeleteByID {
	data := pool_MSG_FILE_DeleteByID.Get().(*MSG_FILE_DeleteByID)
	data.read(buf)
	return data
}

func (data *MSG_FILE_DeleteByID) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.FileID = READ_int64(buf)

}
func (data *MSG_FILE_DeleteByID) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_FILE_DeleteByID) setQueryID(id uint32) {
	data.QueryID = id
}

