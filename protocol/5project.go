package protocol

import (
	"sync"
	"libraries"
)

const (
	CMD_MSG_PROJECT_tree_getLinePairs = -1108380155
	CMD_MSG_PROJECT_tree_getLinePairs_result = -1262905851
	CMD_MSG_PROJECT_product_cache = -553153019
	CMD_MSG_PROJECT_product_insert = -504988411
	CMD_MSG_PROJECT_product_insert_result = -686336507
)

type MSG_PROJECT_tree_getLinePairs struct {
	QueryID uint32
}

var pool_MSG_PROJECT_tree_getLinePairs = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_getLinePairs{} }}

func GET_MSG_PROJECT_tree_getLinePairs() *MSG_PROJECT_tree_getLinePairs {
	return pool_MSG_PROJECT_tree_getLinePairs.Get().(*MSG_PROJECT_tree_getLinePairs)
}

func (data *MSG_PROJECT_tree_getLinePairs) cmd() int32 {
	return CMD_MSG_PROJECT_tree_getLinePairs
}

func (data *MSG_PROJECT_tree_getLinePairs) Put() {
	data.QueryID = 0
	pool_MSG_PROJECT_tree_getLinePairs.Put(data)
}
func (data *MSG_PROJECT_tree_getLinePairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_getLinePairs,buf)
	WRITE_MSG_PROJECT_tree_getLinePairs(data, buf)
}

func WRITE_MSG_PROJECT_tree_getLinePairs(data *MSG_PROJECT_tree_getLinePairs, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
}

func READ_MSG_PROJECT_tree_getLinePairs(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_getLinePairs {
	data := pool_MSG_PROJECT_tree_getLinePairs.Get().(*MSG_PROJECT_tree_getLinePairs)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_getLinePairs) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)

}
func (data *MSG_PROJECT_tree_getLinePairs) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_tree_getLinePairs) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_tree_getLinePairs_result struct {
	QueryResultID uint32
	List []HtmlKeyValueStr
}

var pool_MSG_PROJECT_tree_getLinePairs_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_getLinePairs_result{} }}

func GET_MSG_PROJECT_tree_getLinePairs_result() *MSG_PROJECT_tree_getLinePairs_result {
	return pool_MSG_PROJECT_tree_getLinePairs_result.Get().(*MSG_PROJECT_tree_getLinePairs_result)
}

func (data *MSG_PROJECT_tree_getLinePairs_result) cmd() int32 {
	return CMD_MSG_PROJECT_tree_getLinePairs_result
}

func (data *MSG_PROJECT_tree_getLinePairs_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	pool_MSG_PROJECT_tree_getLinePairs_result.Put(data)
}
func (data *MSG_PROJECT_tree_getLinePairs_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_getLinePairs_result,buf)
	WRITE_MSG_PROJECT_tree_getLinePairs_result(data, buf)
}

func WRITE_MSG_PROJECT_tree_getLinePairs_result(data *MSG_PROJECT_tree_getLinePairs_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_PROJECT_tree_getLinePairs_result(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_getLinePairs_result {
	data := pool_MSG_PROJECT_tree_getLinePairs_result.Get().(*MSG_PROJECT_tree_getLinePairs_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_getLinePairs_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	for i := 0; i < List_len; i++ {
		data.List = append(data.List, READ_HtmlKeyValueStr(buf))
	}

}
func (data *MSG_PROJECT_tree_getLinePairs_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_tree_getLinePairs_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_product_cache struct {
	Id int32
	Name string
	Code string
	Line int32
	Type string
	Status string
	Desc string
	PO int32
	QD int32
	RD int32
	Acl string
	Whitelist []int32
	CreatedBy int32
	CreatedDate int64
	Order int32
	Deleted bool
	TimeStamp int64
}

var pool_MSG_PROJECT_product_cache = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_cache{} }}

func GET_MSG_PROJECT_product_cache() *MSG_PROJECT_product_cache {
	return pool_MSG_PROJECT_product_cache.Get().(*MSG_PROJECT_product_cache)
}

func (data *MSG_PROJECT_product_cache) cmd() int32 {
	return CMD_MSG_PROJECT_product_cache
}

func (data *MSG_PROJECT_product_cache) Put() {
	data.Id = 0
	data.Name = ``
	data.Code = ``
	data.Line = 0
	data.Type = ``
	data.Status = ``
	data.Desc = ``
	data.PO = 0
	data.QD = 0
	data.RD = 0
	data.Acl = ``
	data.Whitelist = data.Whitelist[:0]
	data.CreatedBy = 0
	data.CreatedDate = 0
	data.Order = 0
	data.Deleted = false
	data.TimeStamp = 0
	pool_MSG_PROJECT_product_cache.Put(data)
}
func (data *MSG_PROJECT_product_cache) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_cache,buf)
	WRITE_MSG_PROJECT_product_cache(data, buf)
}

func WRITE_MSG_PROJECT_product_cache(data *MSG_PROJECT_product_cache, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Code, buf)
	WRITE_int32(data.Line, buf)
	WRITE_string(data.Type, buf)
	WRITE_string(data.Status, buf)
	WRITE_string(data.Desc, buf)
	WRITE_int32(data.PO, buf)
	WRITE_int32(data.QD, buf)
	WRITE_int32(data.RD, buf)
	WRITE_string(data.Acl, buf)
	WRITE_int32(int32(len(data.Whitelist)), buf)
	for _, v := range data.Whitelist{
		WRITE_int32(v, buf)
	}
	WRITE_int32(data.CreatedBy, buf)
	WRITE_int64(data.CreatedDate, buf)
	WRITE_int32(data.Order, buf)
	WRITE_bool(data.Deleted, buf)
	WRITE_int64(data.TimeStamp, buf)
}

func READ_MSG_PROJECT_product_cache(buf *libraries.MsgBuffer) *MSG_PROJECT_product_cache {
	data := pool_MSG_PROJECT_product_cache.Get().(*MSG_PROJECT_product_cache)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_cache) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Code = READ_string(buf)
	data.Line = READ_int32(buf)
	data.Type = READ_string(buf)
	data.Status = READ_string(buf)
	data.Desc = READ_string(buf)
	data.PO = READ_int32(buf)
	data.QD = READ_int32(buf)
	data.RD = READ_int32(buf)
	data.Acl = READ_string(buf)
	Whitelist_len := int(READ_int32(buf))
	for i := 0; i < Whitelist_len; i++ {
		data.Whitelist = append(data.Whitelist, READ_int32(buf))
	}
	data.CreatedBy = READ_int32(buf)
	data.CreatedDate = READ_int64(buf)
	data.Order = READ_int32(buf)
	data.Deleted = READ_bool(buf)
	data.TimeStamp = READ_int64(buf)

}

type MSG_PROJECT_product_insert struct {
	QueryID uint32
	DocName string
	Data *MSG_PROJECT_product_cache
}

var pool_MSG_PROJECT_product_insert = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_insert{} }}

func GET_MSG_PROJECT_product_insert() *MSG_PROJECT_product_insert {
	return pool_MSG_PROJECT_product_insert.Get().(*MSG_PROJECT_product_insert)
}

func (data *MSG_PROJECT_product_insert) cmd() int32 {
	return CMD_MSG_PROJECT_product_insert
}

func (data *MSG_PROJECT_product_insert) Put() {
	data.QueryID = 0
	data.DocName = ``
	if data.Data != nil {
		data.Data.Put()
		data.Data = nil
	}
	pool_MSG_PROJECT_product_insert.Put(data)
}
func (data *MSG_PROJECT_product_insert) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_insert,buf)
	WRITE_MSG_PROJECT_product_insert(data, buf)
}

func WRITE_MSG_PROJECT_product_insert(data *MSG_PROJECT_product_insert, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_string(data.DocName, buf)
	if data.Data == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_PROJECT_product_cache(data.Data, buf)
	}
}

func READ_MSG_PROJECT_product_insert(buf *libraries.MsgBuffer) *MSG_PROJECT_product_insert {
	data := pool_MSG_PROJECT_product_insert.Get().(*MSG_PROJECT_product_insert)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_insert) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.DocName = READ_string(buf)
	Data_len := int(READ_int8(buf))
	if Data_len == 1 {
		data.Data = READ_MSG_PROJECT_product_cache(buf)
	}else{
		data.Data = nil
	}

}
func (data *MSG_PROJECT_product_insert) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_product_insert) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_product_insert_result struct {
	QueryResultID uint32
	ID int32
}

var pool_MSG_PROJECT_product_insert_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_insert_result{} }}

func GET_MSG_PROJECT_product_insert_result() *MSG_PROJECT_product_insert_result {
	return pool_MSG_PROJECT_product_insert_result.Get().(*MSG_PROJECT_product_insert_result)
}

func (data *MSG_PROJECT_product_insert_result) cmd() int32 {
	return CMD_MSG_PROJECT_product_insert_result
}

func (data *MSG_PROJECT_product_insert_result) Put() {
	data.QueryResultID = 0
	data.ID = 0
	pool_MSG_PROJECT_product_insert_result.Put(data)
}
func (data *MSG_PROJECT_product_insert_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_insert_result,buf)
	WRITE_MSG_PROJECT_product_insert_result(data, buf)
}

func WRITE_MSG_PROJECT_product_insert_result(data *MSG_PROJECT_product_insert_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(data.ID, buf)
}

func READ_MSG_PROJECT_product_insert_result(buf *libraries.MsgBuffer) *MSG_PROJECT_product_insert_result {
	data := pool_MSG_PROJECT_product_insert_result.Get().(*MSG_PROJECT_product_insert_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_insert_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.ID = READ_int32(buf)

}
func (data *MSG_PROJECT_product_insert_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_product_insert_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

