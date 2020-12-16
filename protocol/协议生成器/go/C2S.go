package protocol

import (
	"sync"
	"bbs/libraries"
)

const (
	CMD_MSG_C2S_Conn_Client = 340003950
	CMD_MSG_S2C_Conn_Client = -467585541
	CMD_MSG_Conn_Down = 691372614
	CMD_MSG_C2S_Regedit = -1221014720
)

type MSG_C2S_Conn_Client struct {
	Fd int32
	Ip string
	UserAgent string
}

var pool_MSG_C2S_Conn_Client = sync.Pool{New: func() interface{} { return &MSG_C2S_Conn_Client{} }}

func GET_MSG_C2S_Conn_Client() *MSG_C2S_Conn_Client {
	return pool_MSG_C2S_Conn_Client.Get().(*MSG_C2S_Conn_Client)
}

func (data *MSG_C2S_Conn_Client) Put() {
	data.Fd = 0
	data.Ip = ``
	data.UserAgent = ``
	pool_MSG_C2S_Conn_Client.Put(data)
}
func (data *MSG_C2S_Conn_Client) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_C2S_Conn_Client)
	WRITE_int32(cmd, buf)
	WRITE_MSG_C2S_Conn_Client(data, buf)
}

func WRITE_MSG_C2S_Conn_Client(data *MSG_C2S_Conn_Client, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fd, buf)
	WRITE_string(data.Ip, buf)
	WRITE_string(data.UserAgent, buf)
}

func READ_MSG_C2S_Conn_Client(buf *libraries.MsgBuffer) (data *MSG_C2S_Conn_Client) {
	data = pool_MSG_C2S_Conn_Client.Get().(*MSG_C2S_Conn_Client)
	data.Fd = READ_int32(buf)
	data.Ip = READ_string(buf)
	data.UserAgent = READ_string(buf)
	return
}

type MSG_S2C_Conn_Client struct {
	Fd int32
}

var pool_MSG_S2C_Conn_Client = sync.Pool{New: func() interface{} { return &MSG_S2C_Conn_Client{} }}

func GET_MSG_S2C_Conn_Client() *MSG_S2C_Conn_Client {
	return pool_MSG_S2C_Conn_Client.Get().(*MSG_S2C_Conn_Client)
}

func (data *MSG_S2C_Conn_Client) Put() {
	data.Fd = 0
	pool_MSG_S2C_Conn_Client.Put(data)
}
func (data *MSG_S2C_Conn_Client) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_S2C_Conn_Client)
	WRITE_int32(cmd, buf)
	WRITE_MSG_S2C_Conn_Client(data, buf)
}

func WRITE_MSG_S2C_Conn_Client(data *MSG_S2C_Conn_Client, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fd, buf)
}

func READ_MSG_S2C_Conn_Client(buf *libraries.MsgBuffer) (data *MSG_S2C_Conn_Client) {
	data = pool_MSG_S2C_Conn_Client.Get().(*MSG_S2C_Conn_Client)
	data.Fd = READ_int32(buf)
	return
}

type MSG_Conn_Down struct {
	Fd int32
	GroupId int8
}

var pool_MSG_Conn_Down = sync.Pool{New: func() interface{} { return &MSG_Conn_Down{} }}

func GET_MSG_Conn_Down() *MSG_Conn_Down {
	return pool_MSG_Conn_Down.Get().(*MSG_Conn_Down)
}

func (data *MSG_Conn_Down) Put() {
	data.Fd = 0
	data.GroupId = 0
	pool_MSG_Conn_Down.Put(data)
}
func (data *MSG_Conn_Down) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_Conn_Down)
	WRITE_int32(cmd, buf)
	WRITE_MSG_Conn_Down(data, buf)
}

func WRITE_MSG_Conn_Down(data *MSG_Conn_Down, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fd, buf)
	WRITE_int8(data.GroupId, buf)
}

func READ_MSG_Conn_Down(buf *libraries.MsgBuffer) (data *MSG_Conn_Down) {
	data = pool_MSG_Conn_Down.Get().(*MSG_Conn_Down)
	data.Fd = READ_int32(buf)
	data.GroupId = READ_int8(buf)
	return
}

type MSG_C2S_Regedit struct {
	GroupId int8
	Serverid int8
	Key string
	Time int64
}

var pool_MSG_C2S_Regedit = sync.Pool{New: func() interface{} { return &MSG_C2S_Regedit{} }}

func GET_MSG_C2S_Regedit() *MSG_C2S_Regedit {
	return pool_MSG_C2S_Regedit.Get().(*MSG_C2S_Regedit)
}

func (data *MSG_C2S_Regedit) Put() {
	data.GroupId = 0
	data.Serverid = 0
	data.Key = ``
	data.Time = 0
	pool_MSG_C2S_Regedit.Put(data)
}
func (data *MSG_C2S_Regedit) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_C2S_Regedit)
	WRITE_int32(cmd, buf)
	WRITE_MSG_C2S_Regedit(data, buf)
}

func WRITE_MSG_C2S_Regedit(data *MSG_C2S_Regedit, buf *libraries.MsgBuffer) {
	WRITE_int8(data.GroupId, buf)
	WRITE_int8(data.Serverid, buf)
	WRITE_string(data.Key, buf)
	WRITE_int64(data.Time, buf)
}

func READ_MSG_C2S_Regedit(buf *libraries.MsgBuffer) (data *MSG_C2S_Regedit) {
	data = pool_MSG_C2S_Regedit.Get().(*MSG_C2S_Regedit)
	data.GroupId = READ_int8(buf)
	data.Serverid = READ_int8(buf)
	data.Key = READ_string(buf)
	data.Time = READ_int64(buf)
	return
}

