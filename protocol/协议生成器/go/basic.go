package protocol

import (
	"sync"
	"bbs/libraries"
)

const (
	CMD_MSG_U2WS_Ping = -2064054626
	CMD_MSG_WS2U_Ping = -1181151646
	CMD_MSG_U2WS_Gettoken = -160848332
	CMD_MSG_WS2U_Gettoken = 157251877
	CMD_MSG_WS2U_CommonResult = 436275943
	CMD_MSG_U2WS_logout = -843349899
	CMD_MSG_U2WS_settoken = -135332941
	CMD_MSG_U2WS_Getseccode = -1981134582
	CMD_MSG_WS2U_Getseccode = -866166167
	CMD_MSG_WS2U_Common_Head = 1185077681
	CMD_MSG_WS2U_Server_OK = 1710690038
	CMD_MSG_U2WS_tpl_load_js = -761168918
	CMD_MSG_WS2U_tpl_load_js = 106954267
)

type MSG_U2WS_Ping struct {
}

var pool_MSG_U2WS_Ping = sync.Pool{New: func() interface{} { return &MSG_U2WS_Ping{} }}

func GET_MSG_U2WS_Ping() *MSG_U2WS_Ping {
	return pool_MSG_U2WS_Ping.Get().(*MSG_U2WS_Ping)
}

func (data *MSG_U2WS_Ping) Put() {
	pool_MSG_U2WS_Ping.Put(data)
}
func (data *MSG_U2WS_Ping) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_Ping)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_Ping(data, buf)
}

func WRITE_MSG_U2WS_Ping(data *MSG_U2WS_Ping, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_Ping(buf *libraries.MsgBuffer) (data *MSG_U2WS_Ping) {
	data = pool_MSG_U2WS_Ping.Get().(*MSG_U2WS_Ping)
	return
}

type MSG_WS2U_Ping struct {
	Result int16
}

var pool_MSG_WS2U_Ping = sync.Pool{New: func() interface{} { return &MSG_WS2U_Ping{} }}

func GET_MSG_WS2U_Ping() *MSG_WS2U_Ping {
	return pool_MSG_WS2U_Ping.Get().(*MSG_WS2U_Ping)
}

func (data *MSG_WS2U_Ping) Put() {
	data.Result = 0
	pool_MSG_WS2U_Ping.Put(data)
}
func (data *MSG_WS2U_Ping) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_Ping)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_Ping(data, buf)
}

func WRITE_MSG_WS2U_Ping(data *MSG_WS2U_Ping, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
}

func READ_MSG_WS2U_Ping(buf *libraries.MsgBuffer) (data *MSG_WS2U_Ping) {
	data = pool_MSG_WS2U_Ping.Get().(*MSG_WS2U_Ping)
	data.Result = READ_int16(buf)
	return
}

type MSG_U2WS_Gettoken struct {
}

var pool_MSG_U2WS_Gettoken = sync.Pool{New: func() interface{} { return &MSG_U2WS_Gettoken{} }}

func GET_MSG_U2WS_Gettoken() *MSG_U2WS_Gettoken {
	return pool_MSG_U2WS_Gettoken.Get().(*MSG_U2WS_Gettoken)
}

func (data *MSG_U2WS_Gettoken) Put() {
	pool_MSG_U2WS_Gettoken.Put(data)
}
func (data *MSG_U2WS_Gettoken) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_Gettoken)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_Gettoken(data, buf)
}

func WRITE_MSG_U2WS_Gettoken(data *MSG_U2WS_Gettoken, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_Gettoken(buf *libraries.MsgBuffer) (data *MSG_U2WS_Gettoken) {
	data = pool_MSG_U2WS_Gettoken.Get().(*MSG_U2WS_Gettoken)
	return
}

type MSG_WS2U_Gettoken struct {
	Token []byte
	Head *MSG_WS2U_Common_Head
}

var pool_MSG_WS2U_Gettoken = sync.Pool{New: func() interface{} { return &MSG_WS2U_Gettoken{} }}

func GET_MSG_WS2U_Gettoken() *MSG_WS2U_Gettoken {
	return pool_MSG_WS2U_Gettoken.Get().(*MSG_WS2U_Gettoken)
}

func (data *MSG_WS2U_Gettoken) Put() {
	data.Token = data.Token[:0]
	if data.Head != nil {
		data.Head.Put()
		data.Head = nil
	}
	pool_MSG_WS2U_Gettoken.Put(data)
}
func (data *MSG_WS2U_Gettoken) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_Gettoken)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_Gettoken(data, buf)
}

func WRITE_MSG_WS2U_Gettoken(data *MSG_WS2U_Gettoken, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Token)), buf)
	buf.Write(data.Token)
	if data.Head == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_WS2U_Common_Head(data.Head, buf)
	}
}

func READ_MSG_WS2U_Gettoken(buf *libraries.MsgBuffer) (data *MSG_WS2U_Gettoken) {
	data = pool_MSG_WS2U_Gettoken.Get().(*MSG_WS2U_Gettoken)
	Token_len := int(READ_int32(buf))
	data.Token = make([]byte, Token_len)
	copy(data.Token,buf.Next(Token_len))
	Head_len := int(READ_int8(buf))
	if Head_len == 1 {
		data.Head = READ_MSG_WS2U_Common_Head(buf)
	}else{
		data.Head = nil
	}
	return
}

type MSG_WS2U_CommonResult struct {
	Result int16
	Err_url string
}

var pool_MSG_WS2U_CommonResult = sync.Pool{New: func() interface{} { return &MSG_WS2U_CommonResult{} }}

func GET_MSG_WS2U_CommonResult() *MSG_WS2U_CommonResult {
	return pool_MSG_WS2U_CommonResult.Get().(*MSG_WS2U_CommonResult)
}

func (data *MSG_WS2U_CommonResult) Put() {
	data.Result = 0
	data.Err_url = ``
	pool_MSG_WS2U_CommonResult.Put(data)
}
func (data *MSG_WS2U_CommonResult) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_CommonResult)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_CommonResult(data, buf)
}

func WRITE_MSG_WS2U_CommonResult(data *MSG_WS2U_CommonResult, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
	WRITE_string(data.Err_url, buf)
}

func READ_MSG_WS2U_CommonResult(buf *libraries.MsgBuffer) (data *MSG_WS2U_CommonResult) {
	data = pool_MSG_WS2U_CommonResult.Get().(*MSG_WS2U_CommonResult)
	data.Result = READ_int16(buf)
	data.Err_url = READ_string(buf)
	return
}

type MSG_U2WS_logout struct {
}

var pool_MSG_U2WS_logout = sync.Pool{New: func() interface{} { return &MSG_U2WS_logout{} }}

func GET_MSG_U2WS_logout() *MSG_U2WS_logout {
	return pool_MSG_U2WS_logout.Get().(*MSG_U2WS_logout)
}

func (data *MSG_U2WS_logout) Put() {
	pool_MSG_U2WS_logout.Put(data)
}
func (data *MSG_U2WS_logout) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_logout)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_logout(data, buf)
}

func WRITE_MSG_U2WS_logout(data *MSG_U2WS_logout, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_logout(buf *libraries.MsgBuffer) (data *MSG_U2WS_logout) {
	data = pool_MSG_U2WS_logout.Get().(*MSG_U2WS_logout)
	return
}

type MSG_U2WS_settoken struct {
	Token []byte
}

var pool_MSG_U2WS_settoken = sync.Pool{New: func() interface{} { return &MSG_U2WS_settoken{} }}

func GET_MSG_U2WS_settoken() *MSG_U2WS_settoken {
	return pool_MSG_U2WS_settoken.Get().(*MSG_U2WS_settoken)
}

func (data *MSG_U2WS_settoken) Put() {
	data.Token = data.Token[:0]
	pool_MSG_U2WS_settoken.Put(data)
}
func (data *MSG_U2WS_settoken) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_settoken)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_settoken(data, buf)
}

func WRITE_MSG_U2WS_settoken(data *MSG_U2WS_settoken, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Token)), buf)
	buf.Write(data.Token)
}

func READ_MSG_U2WS_settoken(buf *libraries.MsgBuffer) (data *MSG_U2WS_settoken) {
	data = pool_MSG_U2WS_settoken.Get().(*MSG_U2WS_settoken)
	Token_len := int(READ_int32(buf))
	data.Token = make([]byte, Token_len)
	copy(data.Token,buf.Next(Token_len))
	return
}

type MSG_U2WS_Getseccode struct {
}

var pool_MSG_U2WS_Getseccode = sync.Pool{New: func() interface{} { return &MSG_U2WS_Getseccode{} }}

func GET_MSG_U2WS_Getseccode() *MSG_U2WS_Getseccode {
	return pool_MSG_U2WS_Getseccode.Get().(*MSG_U2WS_Getseccode)
}

func (data *MSG_U2WS_Getseccode) Put() {
	pool_MSG_U2WS_Getseccode.Put(data)
}
func (data *MSG_U2WS_Getseccode) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_Getseccode)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_Getseccode(data, buf)
}

func WRITE_MSG_U2WS_Getseccode(data *MSG_U2WS_Getseccode, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_Getseccode(buf *libraries.MsgBuffer) (data *MSG_U2WS_Getseccode) {
	data = pool_MSG_U2WS_Getseccode.Get().(*MSG_U2WS_Getseccode)
	return
}

type MSG_WS2U_Getseccode struct {
	Img []byte
}

var pool_MSG_WS2U_Getseccode = sync.Pool{New: func() interface{} { return &MSG_WS2U_Getseccode{} }}

func GET_MSG_WS2U_Getseccode() *MSG_WS2U_Getseccode {
	return pool_MSG_WS2U_Getseccode.Get().(*MSG_WS2U_Getseccode)
}

func (data *MSG_WS2U_Getseccode) Put() {
	data.Img = data.Img[:0]
	pool_MSG_WS2U_Getseccode.Put(data)
}
func (data *MSG_WS2U_Getseccode) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_Getseccode)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_Getseccode(data, buf)
}

func WRITE_MSG_WS2U_Getseccode(data *MSG_WS2U_Getseccode, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Img)), buf)
	buf.Write(data.Img)
}

func READ_MSG_WS2U_Getseccode(buf *libraries.MsgBuffer) (data *MSG_WS2U_Getseccode) {
	data = pool_MSG_WS2U_Getseccode.Get().(*MSG_WS2U_Getseccode)
	Img_len := int(READ_int32(buf))
	data.Img = make([]byte, Img_len)
	copy(data.Img,buf.Next(Img_len))
	return
}

type MSG_WS2U_Common_Head struct {
	Bbname string
	Sitename string
	Username string
	Grouptitle string
	Avatar string
	Portalcp int8
	Admincp int8
	Adminid int16
	Groupid int16
	Diy int8
	Uid int32
	Unread_num int8
	Send_botton int8
	Timestamp int32
}

var pool_MSG_WS2U_Common_Head = sync.Pool{New: func() interface{} { return &MSG_WS2U_Common_Head{} }}

func GET_MSG_WS2U_Common_Head() *MSG_WS2U_Common_Head {
	return pool_MSG_WS2U_Common_Head.Get().(*MSG_WS2U_Common_Head)
}

func (data *MSG_WS2U_Common_Head) Put() {
	data.Bbname = ``
	data.Sitename = ``
	data.Username = ``
	data.Grouptitle = ``
	data.Avatar = ``
	data.Portalcp = 0
	data.Admincp = 0
	data.Adminid = 0
	data.Groupid = 0
	data.Diy = 0
	data.Uid = 0
	data.Unread_num = 0
	data.Send_botton = 0
	data.Timestamp = 0
	pool_MSG_WS2U_Common_Head.Put(data)
}
func (data *MSG_WS2U_Common_Head) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_Common_Head)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_Common_Head(data, buf)
}

func WRITE_MSG_WS2U_Common_Head(data *MSG_WS2U_Common_Head, buf *libraries.MsgBuffer) {
	WRITE_string(data.Bbname, buf)
	WRITE_string(data.Sitename, buf)
	WRITE_string(data.Username, buf)
	WRITE_string(data.Grouptitle, buf)
	WRITE_string(data.Avatar, buf)
	WRITE_int8(data.Portalcp, buf)
	WRITE_int8(data.Admincp, buf)
	WRITE_int16(data.Adminid, buf)
	WRITE_int16(data.Groupid, buf)
	WRITE_int8(data.Diy, buf)
	WRITE_int32(data.Uid, buf)
	WRITE_int8(data.Unread_num, buf)
	WRITE_int8(data.Send_botton, buf)
	WRITE_int32(data.Timestamp, buf)
}

func READ_MSG_WS2U_Common_Head(buf *libraries.MsgBuffer) (data *MSG_WS2U_Common_Head) {
	data = pool_MSG_WS2U_Common_Head.Get().(*MSG_WS2U_Common_Head)
	data.Bbname = READ_string(buf)
	data.Sitename = READ_string(buf)
	data.Username = READ_string(buf)
	data.Grouptitle = READ_string(buf)
	data.Avatar = READ_string(buf)
	data.Portalcp = READ_int8(buf)
	data.Admincp = READ_int8(buf)
	data.Adminid = READ_int16(buf)
	data.Groupid = READ_int16(buf)
	data.Diy = READ_int8(buf)
	data.Uid = READ_int32(buf)
	data.Unread_num = READ_int8(buf)
	data.Send_botton = READ_int8(buf)
	data.Timestamp = READ_int32(buf)
	return
}

type MSG_WS2U_Server_OK struct {
}

var pool_MSG_WS2U_Server_OK = sync.Pool{New: func() interface{} { return &MSG_WS2U_Server_OK{} }}

func GET_MSG_WS2U_Server_OK() *MSG_WS2U_Server_OK {
	return pool_MSG_WS2U_Server_OK.Get().(*MSG_WS2U_Server_OK)
}

func (data *MSG_WS2U_Server_OK) Put() {
	pool_MSG_WS2U_Server_OK.Put(data)
}
func (data *MSG_WS2U_Server_OK) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_Server_OK)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_Server_OK(data, buf)
}

func WRITE_MSG_WS2U_Server_OK(data *MSG_WS2U_Server_OK, buf *libraries.MsgBuffer) {
}

func READ_MSG_WS2U_Server_OK(buf *libraries.MsgBuffer) (data *MSG_WS2U_Server_OK) {
	data = pool_MSG_WS2U_Server_OK.Get().(*MSG_WS2U_Server_OK)
	return
}

type MSG_U2WS_tpl_load_js struct {
	Name string
}

var pool_MSG_U2WS_tpl_load_js = sync.Pool{New: func() interface{} { return &MSG_U2WS_tpl_load_js{} }}

func GET_MSG_U2WS_tpl_load_js() *MSG_U2WS_tpl_load_js {
	return pool_MSG_U2WS_tpl_load_js.Get().(*MSG_U2WS_tpl_load_js)
}

func (data *MSG_U2WS_tpl_load_js) Put() {
	data.Name = ``
	pool_MSG_U2WS_tpl_load_js.Put(data)
}
func (data *MSG_U2WS_tpl_load_js) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_tpl_load_js)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_tpl_load_js(data, buf)
}

func WRITE_MSG_U2WS_tpl_load_js(data *MSG_U2WS_tpl_load_js, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
}

func READ_MSG_U2WS_tpl_load_js(buf *libraries.MsgBuffer) (data *MSG_U2WS_tpl_load_js) {
	data = pool_MSG_U2WS_tpl_load_js.Get().(*MSG_U2WS_tpl_load_js)
	data.Name = READ_string(buf)
	return
}

type MSG_WS2U_tpl_load_js struct {
	Name string
	Result string
}

var pool_MSG_WS2U_tpl_load_js = sync.Pool{New: func() interface{} { return &MSG_WS2U_tpl_load_js{} }}

func GET_MSG_WS2U_tpl_load_js() *MSG_WS2U_tpl_load_js {
	return pool_MSG_WS2U_tpl_load_js.Get().(*MSG_WS2U_tpl_load_js)
}

func (data *MSG_WS2U_tpl_load_js) Put() {
	data.Name = ``
	data.Result = ``
	pool_MSG_WS2U_tpl_load_js.Put(data)
}
func (data *MSG_WS2U_tpl_load_js) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_tpl_load_js)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_tpl_load_js(data, buf)
}

func WRITE_MSG_WS2U_tpl_load_js(data *MSG_WS2U_tpl_load_js, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
	WRITE_string(data.Result, buf)
}

func READ_MSG_WS2U_tpl_load_js(buf *libraries.MsgBuffer) (data *MSG_WS2U_tpl_load_js) {
	data = pool_MSG_WS2U_tpl_load_js.Get().(*MSG_WS2U_tpl_load_js)
	data.Name = READ_string(buf)
	data.Result = READ_string(buf)
	return
}

