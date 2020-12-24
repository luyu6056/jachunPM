package protocol

import (
	"sync"
	"libraries"
)

const (
	CMD_MSG_PROJECT_tree_getLinePairs = -1108380155
	CMD_MSG_PROJECT_tree_getLinePairs_result = -1262905851
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

