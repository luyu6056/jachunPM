package protocol

import (
	"sync"
	"libraries"
	"time"
)

const (
	CMD_MSG_LOG_Ation = 865382915
	CMD_MSG_LOG_History = -1004505085
)

type MSG_LOG_Ation struct {
	Id int64
	ObjectType string
	ObjectID int32
	Product int32
	Project int32
	ActorId int32
	Actor string
	Action string
	Date time.Time
	Comment string
	Extra string
	Read bool
	historys []*MSG_LOG_History
}

var pool_MSG_LOG_Ation = sync.Pool{New: func() interface{} { return &MSG_LOG_Ation{} }}

func GET_MSG_LOG_Ation() *MSG_LOG_Ation {
	return pool_MSG_LOG_Ation.Get().(*MSG_LOG_Ation)
}

func (data *MSG_LOG_Ation) cmd() int32 {
	return CMD_MSG_LOG_Ation
}

func (data *MSG_LOG_Ation) Put() {
	data.Id = 0
	data.ObjectType = ``
	data.ObjectID = 0
	data.Product = 0
	data.Project = 0
	data.ActorId = 0
	data.Actor = ``
	data.Action = ``
	data.Date = time.Unix(0,0)
	data.Comment = ``
	data.Extra = ``
	data.Read = false
	for _,v := range data.historys {
		v.Put()
	}
	data.historys = data.historys[:0]
	pool_MSG_LOG_Ation.Put(data)
}
func (data *MSG_LOG_Ation) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Ation,buf)
	WRITE_MSG_LOG_Ation(data, buf)
}

func WRITE_MSG_LOG_Ation(data *MSG_LOG_Ation, buf *libraries.MsgBuffer) {
	WRITE_int64(data.Id, buf)
	WRITE_string(data.ObjectType, buf)
	WRITE_int32(data.ObjectID, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Project, buf)
	WRITE_int32(data.ActorId, buf)
	WRITE_string(data.Actor, buf)
	WRITE_string(data.Action, buf)
	WRITE_int64(data.Date.UnixNano(), buf)
	WRITE_string(data.Comment, buf)
	WRITE_string(data.Extra, buf)
	WRITE_bool(data.Read, buf)
	WRITE_int32(int32(len(data.historys)), buf)
	for _, v := range data.historys{
		WRITE_MSG_LOG_History(v, buf)
	}
}

func READ_MSG_LOG_Ation(buf *libraries.MsgBuffer) *MSG_LOG_Ation {
	data := pool_MSG_LOG_Ation.Get().(*MSG_LOG_Ation)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Ation) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int64(buf)
	data.ObjectType = READ_string(buf)
	data.ObjectID = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Project = READ_int32(buf)
	data.ActorId = READ_int32(buf)
	data.Actor = READ_string(buf)
	data.Action = READ_string(buf)
	data.Date = time.Unix(0, READ_int64(buf))
	data.Comment = READ_string(buf)
	data.Extra = READ_string(buf)
	data.Read = READ_bool(buf)
	historys_len := int(READ_int32(buf))
	if historys_len>cap(data.historys){
		data.historys= make([]*MSG_LOG_History, historys_len)
	}else{
		data.historys = data.historys[:historys_len]
	}
	for i := 0; i < historys_len; i++ {
		data.historys[i] = READ_MSG_LOG_History(buf)
	}

}

type MSG_LOG_History struct {
	Field string
	Old string
	New string
	Diff string
	FieldLabel string
}

var pool_MSG_LOG_History = sync.Pool{New: func() interface{} { return &MSG_LOG_History{} }}

func GET_MSG_LOG_History() *MSG_LOG_History {
	return pool_MSG_LOG_History.Get().(*MSG_LOG_History)
}

func (data *MSG_LOG_History) cmd() int32 {
	return CMD_MSG_LOG_History
}

func (data *MSG_LOG_History) Put() {
	data.Field = ``
	data.Old = ``
	data.New = ``
	data.Diff = ``
	data.FieldLabel = ``
	pool_MSG_LOG_History.Put(data)
}
func (data *MSG_LOG_History) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_History,buf)
	WRITE_MSG_LOG_History(data, buf)
}

func WRITE_MSG_LOG_History(data *MSG_LOG_History, buf *libraries.MsgBuffer) {
	WRITE_string(data.Field, buf)
	WRITE_string(data.Old, buf)
	WRITE_string(data.New, buf)
	WRITE_string(data.Diff, buf)
	WRITE_string(data.FieldLabel, buf)
}

func READ_MSG_LOG_History(buf *libraries.MsgBuffer) *MSG_LOG_History {
	data := pool_MSG_LOG_History.Get().(*MSG_LOG_History)
	data.read(buf)
	return data
}

func (data *MSG_LOG_History) read(buf *libraries.MsgBuffer) {
	data.Field = READ_string(buf)
	data.Old = READ_string(buf)
	data.New = READ_string(buf)
	data.Diff = READ_string(buf)
	data.FieldLabel = READ_string(buf)

}

