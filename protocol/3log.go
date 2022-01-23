package protocol

import (
	"sync"
	"libraries"
	"unsafe"
	"time"
)

const (
	CMD_MSG_LOG_Action = -1093696765
	CMD_MSG_LOG_History = -1004505085
	CMD_MSG_LOG_Action_Create = -1137975037
	CMD_MSG_LOG_Action_Create_result = -1288845053
	CMD_MSG_LOG_Action_GetByWhereMap = -1161866493
	CMD_MSG_LOG_Action_GetByID = 488167427
	CMD_MSG_LOG_Action_GetByID_result = 1253762819
	CMD_MSG_LOG_Action_GetByWhereMap_result = -810685181
	CMD_MSG_LOG_Action_transformActions = 1875485187
	CMD_MSG_LOG_Action_transformActions_result = -1642266109
	CMD_MSG_LOG_transformActions_info = 1937124355
	CMD_MSG_LOG_Action_AddHistory = 882194947
	CMD_MSG_LOG_Action_set_read = -576769789
)

type MSG_LOG_Action struct {
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
	Historys []*MSG_LOG_History
	AppendLink string `db:"-"`
}

var pool_MSG_LOG_Action = sync.Pool{New: func() interface{} { return &MSG_LOG_Action{} }}

func GET_MSG_LOG_Action() *MSG_LOG_Action {
	return pool_MSG_LOG_Action.Get().(*MSG_LOG_Action)
}

func (data *MSG_LOG_Action) cmd() int32 {
	return CMD_MSG_LOG_Action
}

func (data *MSG_LOG_Action) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_Action()))
}

func (data *MSG_LOG_Action) Put() {
	data.Id = 0
	data.ObjectType = ``
	data.ObjectID = 0
	data.Product = 0
	data.Project = 0
	data.ActorId = 0
	data.Actor = ``
	data.Action = ``
	data.Date = time.UnixMicro(0)
	data.Comment = ``
	data.Extra = ``
	data.Read = false
	data.Historys = data.Historys[:0]
	data.AppendLink = ``
	pool_MSG_LOG_Action.Put(data)
}
func (data *MSG_LOG_Action) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Action,buf)
	WRITE_MSG_LOG_Action(data, buf)
}

func WRITE_MSG_LOG_Action(data *MSG_LOG_Action, buf *libraries.MsgBuffer) {
	WRITE_int64(data.Id, buf)
	WRITE_string(data.ObjectType, buf)
	WRITE_int32(data.ObjectID, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Project, buf)
	WRITE_int32(data.ActorId, buf)
	WRITE_string(data.Actor, buf)
	WRITE_string(data.Action, buf)
	WRITE_int64(data.Date.UnixMicro(), buf)
	WRITE_string(data.Comment, buf)
	WRITE_string(data.Extra, buf)
	WRITE_bool(data.Read, buf)
	WRITE_int(len(data.Historys), buf)
	for _, v := range data.Historys{
		WRITE_MSG_LOG_History(v, buf)
	}
	WRITE_string(data.AppendLink, buf)
}

func READ_MSG_LOG_Action(buf *libraries.MsgBuffer) *MSG_LOG_Action {
	data := pool_MSG_LOG_Action.Get().(*MSG_LOG_Action)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Action) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int64(buf)
	data.ObjectType = READ_string(buf)
	data.ObjectID = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Project = READ_int32(buf)
	data.ActorId = READ_int32(buf)
	data.Actor = READ_string(buf)
	data.Action = READ_string(buf)
	data.Date = time.UnixMicro(READ_int64(buf))
	data.Comment = READ_string(buf)
	data.Extra = READ_string(buf)
	data.Read = READ_bool(buf)
	Historys_len := READ_int(buf)
	if Historys_len>cap(data.Historys){
		data.Historys= make([]*MSG_LOG_History, Historys_len)
	}else{
		data.Historys = data.Historys[:Historys_len]
	}
	for i := 0; i < Historys_len; i++ {
		data.Historys[i] = READ_MSG_LOG_History(buf)
	}
	data.AppendLink = READ_string(buf)

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

func (data *MSG_LOG_History) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_History()))
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

type MSG_LOG_Action_Create struct {
	ObjectType string
	ObjectID int32
	ActionType string
	Comment string
	Extra string
	ActorId int32
	Products []int32
	Project int32
}

var pool_MSG_LOG_Action_Create = sync.Pool{New: func() interface{} { return &MSG_LOG_Action_Create{} }}

func GET_MSG_LOG_Action_Create() *MSG_LOG_Action_Create {
	return pool_MSG_LOG_Action_Create.Get().(*MSG_LOG_Action_Create)
}

func (data *MSG_LOG_Action_Create) cmd() int32 {
	return CMD_MSG_LOG_Action_Create
}

func (data *MSG_LOG_Action_Create) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_Action_Create()))
}

func (data *MSG_LOG_Action_Create) Put() {
	data.ObjectType = ``
	data.ObjectID = 0
	data.ActionType = ``
	data.Comment = ``
	data.Extra = ``
	data.ActorId = 0
	data.Products = data.Products[:0]
	data.Project = 0
	pool_MSG_LOG_Action_Create.Put(data)
}
func (data *MSG_LOG_Action_Create) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Action_Create,buf)
	WRITE_MSG_LOG_Action_Create(data, buf)
}

func WRITE_MSG_LOG_Action_Create(data *MSG_LOG_Action_Create, buf *libraries.MsgBuffer) {
	WRITE_string(data.ObjectType, buf)
	WRITE_int32(data.ObjectID, buf)
	WRITE_string(data.ActionType, buf)
	WRITE_string(data.Comment, buf)
	WRITE_string(data.Extra, buf)
	WRITE_int32(data.ActorId, buf)
	WRITE_int(len(data.Products), buf)
	for _, v := range data.Products{
		WRITE_int32(v, buf)
	}
	WRITE_int32(data.Project, buf)
}

func READ_MSG_LOG_Action_Create(buf *libraries.MsgBuffer) *MSG_LOG_Action_Create {
	data := pool_MSG_LOG_Action_Create.Get().(*MSG_LOG_Action_Create)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Action_Create) read(buf *libraries.MsgBuffer) {
	data.ObjectType = READ_string(buf)
	data.ObjectID = READ_int32(buf)
	data.ActionType = READ_string(buf)
	data.Comment = READ_string(buf)
	data.Extra = READ_string(buf)
	data.ActorId = READ_int32(buf)
	Products_len := READ_int(buf)
	if Products_len>cap(data.Products){
		data.Products= make([]int32, Products_len)
	}else{
		data.Products = data.Products[:Products_len]
	}
	for i := 0; i < Products_len; i++ {
		data.Products[i] = READ_int32(buf)
	}
	data.Project = READ_int32(buf)

}

type MSG_LOG_Action_Create_result struct {
	ActionId int64
}

var pool_MSG_LOG_Action_Create_result = sync.Pool{New: func() interface{} { return &MSG_LOG_Action_Create_result{} }}

func GET_MSG_LOG_Action_Create_result() *MSG_LOG_Action_Create_result {
	return pool_MSG_LOG_Action_Create_result.Get().(*MSG_LOG_Action_Create_result)
}

func (data *MSG_LOG_Action_Create_result) cmd() int32 {
	return CMD_MSG_LOG_Action_Create_result
}

func (data *MSG_LOG_Action_Create_result) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_Action_Create_result()))
}

func (data *MSG_LOG_Action_Create_result) Put() {
	data.ActionId = 0
	pool_MSG_LOG_Action_Create_result.Put(data)
}
func (data *MSG_LOG_Action_Create_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Action_Create_result,buf)
	WRITE_MSG_LOG_Action_Create_result(data, buf)
}

func WRITE_MSG_LOG_Action_Create_result(data *MSG_LOG_Action_Create_result, buf *libraries.MsgBuffer) {
	WRITE_int64(data.ActionId, buf)
}

func READ_MSG_LOG_Action_Create_result(buf *libraries.MsgBuffer) *MSG_LOG_Action_Create_result {
	data := pool_MSG_LOG_Action_Create_result.Get().(*MSG_LOG_Action_Create_result)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Action_Create_result) read(buf *libraries.MsgBuffer) {
	data.ActionId = READ_int64(buf)

}

type MSG_LOG_Action_GetByWhereMap struct {
	Star time.Time
	Where map[string]interface{}
	Order string
}

var pool_MSG_LOG_Action_GetByWhereMap = sync.Pool{New: func() interface{} { return &MSG_LOG_Action_GetByWhereMap{} }}

func GET_MSG_LOG_Action_GetByWhereMap() *MSG_LOG_Action_GetByWhereMap {
	return pool_MSG_LOG_Action_GetByWhereMap.Get().(*MSG_LOG_Action_GetByWhereMap)
}

func (data *MSG_LOG_Action_GetByWhereMap) cmd() int32 {
	return CMD_MSG_LOG_Action_GetByWhereMap
}

func (data *MSG_LOG_Action_GetByWhereMap) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_Action_GetByWhereMap()))
}

func (data *MSG_LOG_Action_GetByWhereMap) Put() {
	data.Star = time.UnixMicro(0)
	data.Where = nil
	data.Order = ``
	pool_MSG_LOG_Action_GetByWhereMap.Put(data)
}
func (data *MSG_LOG_Action_GetByWhereMap) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Action_GetByWhereMap,buf)
	WRITE_MSG_LOG_Action_GetByWhereMap(data, buf)
}

func WRITE_MSG_LOG_Action_GetByWhereMap(data *MSG_LOG_Action_GetByWhereMap, buf *libraries.MsgBuffer) {
	WRITE_int64(data.Star.UnixMicro(), buf)
	WRITE_map(data.Where,buf)
	WRITE_string(data.Order, buf)
}

func READ_MSG_LOG_Action_GetByWhereMap(buf *libraries.MsgBuffer) *MSG_LOG_Action_GetByWhereMap {
	data := pool_MSG_LOG_Action_GetByWhereMap.Get().(*MSG_LOG_Action_GetByWhereMap)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Action_GetByWhereMap) read(buf *libraries.MsgBuffer) {
	data.Star = time.UnixMicro(READ_int64(buf))
	READ_map(&data.Where,buf)
	data.Order = READ_string(buf)

}

type MSG_LOG_Action_GetByID struct {
	Id int64
}

var pool_MSG_LOG_Action_GetByID = sync.Pool{New: func() interface{} { return &MSG_LOG_Action_GetByID{} }}

func GET_MSG_LOG_Action_GetByID() *MSG_LOG_Action_GetByID {
	return pool_MSG_LOG_Action_GetByID.Get().(*MSG_LOG_Action_GetByID)
}

func (data *MSG_LOG_Action_GetByID) cmd() int32 {
	return CMD_MSG_LOG_Action_GetByID
}

func (data *MSG_LOG_Action_GetByID) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_Action_GetByID()))
}

func (data *MSG_LOG_Action_GetByID) Put() {
	data.Id = 0
	pool_MSG_LOG_Action_GetByID.Put(data)
}
func (data *MSG_LOG_Action_GetByID) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Action_GetByID,buf)
	WRITE_MSG_LOG_Action_GetByID(data, buf)
}

func WRITE_MSG_LOG_Action_GetByID(data *MSG_LOG_Action_GetByID, buf *libraries.MsgBuffer) {
	WRITE_int64(data.Id, buf)
}

func READ_MSG_LOG_Action_GetByID(buf *libraries.MsgBuffer) *MSG_LOG_Action_GetByID {
	data := pool_MSG_LOG_Action_GetByID.Get().(*MSG_LOG_Action_GetByID)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Action_GetByID) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int64(buf)

}

type MSG_LOG_Action_GetByID_result struct {
	Info *MSG_LOG_Action
}

var pool_MSG_LOG_Action_GetByID_result = sync.Pool{New: func() interface{} { return &MSG_LOG_Action_GetByID_result{} }}

func GET_MSG_LOG_Action_GetByID_result() *MSG_LOG_Action_GetByID_result {
	return pool_MSG_LOG_Action_GetByID_result.Get().(*MSG_LOG_Action_GetByID_result)
}

func (data *MSG_LOG_Action_GetByID_result) cmd() int32 {
	return CMD_MSG_LOG_Action_GetByID_result
}

func (data *MSG_LOG_Action_GetByID_result) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_Action_GetByID_result()))
}

func (data *MSG_LOG_Action_GetByID_result) Put() {
	if data.Info != nil {
		data.Info.Put()
		data.Info = nil
	}
	pool_MSG_LOG_Action_GetByID_result.Put(data)
}
func (data *MSG_LOG_Action_GetByID_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Action_GetByID_result,buf)
	WRITE_MSG_LOG_Action_GetByID_result(data, buf)
}

func WRITE_MSG_LOG_Action_GetByID_result(data *MSG_LOG_Action_GetByID_result, buf *libraries.MsgBuffer) {
	if data.Info == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_LOG_Action(data.Info, buf)
	}
}

func READ_MSG_LOG_Action_GetByID_result(buf *libraries.MsgBuffer) *MSG_LOG_Action_GetByID_result {
	data := pool_MSG_LOG_Action_GetByID_result.Get().(*MSG_LOG_Action_GetByID_result)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Action_GetByID_result) read(buf *libraries.MsgBuffer) {
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_LOG_Action(buf)
	}else{
		data.Info = nil
	}

}

type MSG_LOG_Action_GetByWhereMap_result struct {
	List []*MSG_LOG_Action
}

var pool_MSG_LOG_Action_GetByWhereMap_result = sync.Pool{New: func() interface{} { return &MSG_LOG_Action_GetByWhereMap_result{} }}

func GET_MSG_LOG_Action_GetByWhereMap_result() *MSG_LOG_Action_GetByWhereMap_result {
	return pool_MSG_LOG_Action_GetByWhereMap_result.Get().(*MSG_LOG_Action_GetByWhereMap_result)
}

func (data *MSG_LOG_Action_GetByWhereMap_result) cmd() int32 {
	return CMD_MSG_LOG_Action_GetByWhereMap_result
}

func (data *MSG_LOG_Action_GetByWhereMap_result) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_Action_GetByWhereMap_result()))
}

func (data *MSG_LOG_Action_GetByWhereMap_result) Put() {
	data.List = data.List[:0]
	pool_MSG_LOG_Action_GetByWhereMap_result.Put(data)
}
func (data *MSG_LOG_Action_GetByWhereMap_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Action_GetByWhereMap_result,buf)
	WRITE_MSG_LOG_Action_GetByWhereMap_result(data, buf)
}

func WRITE_MSG_LOG_Action_GetByWhereMap_result(data *MSG_LOG_Action_GetByWhereMap_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_LOG_Action(v, buf)
	}
}

func READ_MSG_LOG_Action_GetByWhereMap_result(buf *libraries.MsgBuffer) *MSG_LOG_Action_GetByWhereMap_result {
	data := pool_MSG_LOG_Action_GetByWhereMap_result.Get().(*MSG_LOG_Action_GetByWhereMap_result)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Action_GetByWhereMap_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_LOG_Action, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_LOG_Action(buf)
	}

}

type MSG_LOG_Action_transformActions struct {
	Where map[string]interface{}
	Order string
	Page int
	PerPage int
	Total int
}

var pool_MSG_LOG_Action_transformActions = sync.Pool{New: func() interface{} { return &MSG_LOG_Action_transformActions{} }}

func GET_MSG_LOG_Action_transformActions() *MSG_LOG_Action_transformActions {
	return pool_MSG_LOG_Action_transformActions.Get().(*MSG_LOG_Action_transformActions)
}

func (data *MSG_LOG_Action_transformActions) cmd() int32 {
	return CMD_MSG_LOG_Action_transformActions
}

func (data *MSG_LOG_Action_transformActions) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_Action_transformActions()))
}

func (data *MSG_LOG_Action_transformActions) Put() {
	data.Where = nil
	data.Order = ``
	data.Page = 0
	data.PerPage = 0
	data.Total = 0
	pool_MSG_LOG_Action_transformActions.Put(data)
}
func (data *MSG_LOG_Action_transformActions) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Action_transformActions,buf)
	WRITE_MSG_LOG_Action_transformActions(data, buf)
}

func WRITE_MSG_LOG_Action_transformActions(data *MSG_LOG_Action_transformActions, buf *libraries.MsgBuffer) {
	WRITE_map(data.Where,buf)
	WRITE_string(data.Order, buf)
	WRITE_int(data.Page, buf)
	WRITE_int(data.PerPage, buf)
	WRITE_int(data.Total, buf)
}

func READ_MSG_LOG_Action_transformActions(buf *libraries.MsgBuffer) *MSG_LOG_Action_transformActions {
	data := pool_MSG_LOG_Action_transformActions.Get().(*MSG_LOG_Action_transformActions)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Action_transformActions) read(buf *libraries.MsgBuffer) {
	READ_map(&data.Where,buf)
	data.Order = READ_string(buf)
	data.Page = READ_int(buf)
	data.PerPage = READ_int(buf)
	data.Total = READ_int(buf)

}

type MSG_LOG_Action_transformActions_result struct {
	List []*MSG_LOG_transformActions_info
	Total int
}

var pool_MSG_LOG_Action_transformActions_result = sync.Pool{New: func() interface{} { return &MSG_LOG_Action_transformActions_result{} }}

func GET_MSG_LOG_Action_transformActions_result() *MSG_LOG_Action_transformActions_result {
	return pool_MSG_LOG_Action_transformActions_result.Get().(*MSG_LOG_Action_transformActions_result)
}

func (data *MSG_LOG_Action_transformActions_result) cmd() int32 {
	return CMD_MSG_LOG_Action_transformActions_result
}

func (data *MSG_LOG_Action_transformActions_result) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_Action_transformActions_result()))
}

func (data *MSG_LOG_Action_transformActions_result) Put() {
	data.List = data.List[:0]
	data.Total = 0
	pool_MSG_LOG_Action_transformActions_result.Put(data)
}
func (data *MSG_LOG_Action_transformActions_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Action_transformActions_result,buf)
	WRITE_MSG_LOG_Action_transformActions_result(data, buf)
}

func WRITE_MSG_LOG_Action_transformActions_result(data *MSG_LOG_Action_transformActions_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_LOG_transformActions_info(v, buf)
	}
	WRITE_int(data.Total, buf)
}

func READ_MSG_LOG_Action_transformActions_result(buf *libraries.MsgBuffer) *MSG_LOG_Action_transformActions_result {
	data := pool_MSG_LOG_Action_transformActions_result.Get().(*MSG_LOG_Action_transformActions_result)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Action_transformActions_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_LOG_transformActions_info, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_LOG_transformActions_info(buf)
	}
	data.Total = READ_int(buf)

}

type MSG_LOG_transformActions_info struct {
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
	ActionLabel string `db:"-"`
	ObjectLabel string `db:"-"`
	ObjectName string `db:"-"`
	OriginalDate string `db:"-"`
	ObjectLink string `db:"-"`
	Major bool `db:"-"`
}

var pool_MSG_LOG_transformActions_info = sync.Pool{New: func() interface{} { return &MSG_LOG_transformActions_info{} }}

func GET_MSG_LOG_transformActions_info() *MSG_LOG_transformActions_info {
	return pool_MSG_LOG_transformActions_info.Get().(*MSG_LOG_transformActions_info)
}

func (data *MSG_LOG_transformActions_info) cmd() int32 {
	return CMD_MSG_LOG_transformActions_info
}

func (data *MSG_LOG_transformActions_info) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_transformActions_info()))
}

func (data *MSG_LOG_transformActions_info) Put() {
	data.Id = 0
	data.ObjectType = ``
	data.ObjectID = 0
	data.Product = 0
	data.Project = 0
	data.ActorId = 0
	data.Actor = ``
	data.Action = ``
	data.Date = time.UnixMicro(0)
	data.Comment = ``
	data.Extra = ``
	data.Read = false
	data.ActionLabel = ``
	data.ObjectLabel = ``
	data.ObjectName = ``
	data.OriginalDate = ``
	data.ObjectLink = ``
	data.Major = false
	pool_MSG_LOG_transformActions_info.Put(data)
}
func (data *MSG_LOG_transformActions_info) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_transformActions_info,buf)
	WRITE_MSG_LOG_transformActions_info(data, buf)
}

func WRITE_MSG_LOG_transformActions_info(data *MSG_LOG_transformActions_info, buf *libraries.MsgBuffer) {
	WRITE_int64(data.Id, buf)
	WRITE_string(data.ObjectType, buf)
	WRITE_int32(data.ObjectID, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Project, buf)
	WRITE_int32(data.ActorId, buf)
	WRITE_string(data.Actor, buf)
	WRITE_string(data.Action, buf)
	WRITE_int64(data.Date.UnixMicro(), buf)
	WRITE_string(data.Comment, buf)
	WRITE_string(data.Extra, buf)
	WRITE_bool(data.Read, buf)
	WRITE_string(data.ActionLabel, buf)
	WRITE_string(data.ObjectLabel, buf)
	WRITE_string(data.ObjectName, buf)
	WRITE_string(data.OriginalDate, buf)
	WRITE_string(data.ObjectLink, buf)
	WRITE_bool(data.Major, buf)
}

func READ_MSG_LOG_transformActions_info(buf *libraries.MsgBuffer) *MSG_LOG_transformActions_info {
	data := pool_MSG_LOG_transformActions_info.Get().(*MSG_LOG_transformActions_info)
	data.read(buf)
	return data
}

func (data *MSG_LOG_transformActions_info) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int64(buf)
	data.ObjectType = READ_string(buf)
	data.ObjectID = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Project = READ_int32(buf)
	data.ActorId = READ_int32(buf)
	data.Actor = READ_string(buf)
	data.Action = READ_string(buf)
	data.Date = time.UnixMicro(READ_int64(buf))
	data.Comment = READ_string(buf)
	data.Extra = READ_string(buf)
	data.Read = READ_bool(buf)
	data.ActionLabel = READ_string(buf)
	data.ObjectLabel = READ_string(buf)
	data.ObjectName = READ_string(buf)
	data.OriginalDate = READ_string(buf)
	data.ObjectLink = READ_string(buf)
	data.Major = READ_bool(buf)

}

type MSG_LOG_Action_AddHistory struct {
	Id int64
	History []*MSG_LOG_History
}

var pool_MSG_LOG_Action_AddHistory = sync.Pool{New: func() interface{} { return &MSG_LOG_Action_AddHistory{} }}

func GET_MSG_LOG_Action_AddHistory() *MSG_LOG_Action_AddHistory {
	return pool_MSG_LOG_Action_AddHistory.Get().(*MSG_LOG_Action_AddHistory)
}

func (data *MSG_LOG_Action_AddHistory) cmd() int32 {
	return CMD_MSG_LOG_Action_AddHistory
}

func (data *MSG_LOG_Action_AddHistory) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_Action_AddHistory()))
}

func (data *MSG_LOG_Action_AddHistory) Put() {
	data.Id = 0
	data.History = data.History[:0]
	pool_MSG_LOG_Action_AddHistory.Put(data)
}
func (data *MSG_LOG_Action_AddHistory) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Action_AddHistory,buf)
	WRITE_MSG_LOG_Action_AddHistory(data, buf)
}

func WRITE_MSG_LOG_Action_AddHistory(data *MSG_LOG_Action_AddHistory, buf *libraries.MsgBuffer) {
	WRITE_int64(data.Id, buf)
	WRITE_int(len(data.History), buf)
	for _, v := range data.History{
		WRITE_MSG_LOG_History(v, buf)
	}
}

func READ_MSG_LOG_Action_AddHistory(buf *libraries.MsgBuffer) *MSG_LOG_Action_AddHistory {
	data := pool_MSG_LOG_Action_AddHistory.Get().(*MSG_LOG_Action_AddHistory)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Action_AddHistory) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int64(buf)
	History_len := READ_int(buf)
	if History_len>cap(data.History){
		data.History= make([]*MSG_LOG_History, History_len)
	}else{
		data.History = data.History[:History_len]
	}
	for i := 0; i < History_len; i++ {
		data.History[i] = READ_MSG_LOG_History(buf)
	}

}

type MSG_LOG_Action_set_read struct {
	ObjectType string
	ObjectID int32
}

var pool_MSG_LOG_Action_set_read = sync.Pool{New: func() interface{} { return &MSG_LOG_Action_set_read{} }}

func GET_MSG_LOG_Action_set_read() *MSG_LOG_Action_set_read {
	return pool_MSG_LOG_Action_set_read.Get().(*MSG_LOG_Action_set_read)
}

func (data *MSG_LOG_Action_set_read) cmd() int32 {
	return CMD_MSG_LOG_Action_set_read
}

func (data *MSG_LOG_Action_set_read) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_LOG_Action_set_read()))
}

func (data *MSG_LOG_Action_set_read) Put() {
	data.ObjectType = ``
	data.ObjectID = 0
	pool_MSG_LOG_Action_set_read.Put(data)
}
func (data *MSG_LOG_Action_set_read) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_LOG_Action_set_read,buf)
	WRITE_MSG_LOG_Action_set_read(data, buf)
}

func WRITE_MSG_LOG_Action_set_read(data *MSG_LOG_Action_set_read, buf *libraries.MsgBuffer) {
	WRITE_string(data.ObjectType, buf)
	WRITE_int32(data.ObjectID, buf)
}

func READ_MSG_LOG_Action_set_read(buf *libraries.MsgBuffer) *MSG_LOG_Action_set_read {
	data := pool_MSG_LOG_Action_set_read.Get().(*MSG_LOG_Action_set_read)
	data.read(buf)
	return data
}

func (data *MSG_LOG_Action_set_read) read(buf *libraries.MsgBuffer) {
	data.ObjectType = READ_string(buf)
	data.ObjectID = READ_int32(buf)

}

