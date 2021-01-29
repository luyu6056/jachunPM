package protocol

import (
	"sync"
	"libraries"
	"time"
)

const (
	CMD_MSG_TEST_testsuite_info = 829781510
	CMD_MSG_TEST_testsuite_getById = -143122426
	CMD_MSG_TEST_testsuite_getById_result = -1951638010
	CMD_MSG_TEST_bug_getCount = 614527750
	CMD_MSG_TEST_buf_getCount_result = 298479622
	CMD_MSG_TEST_product_deleteBranch_check = 1974311942
	CMD_MSG_TEST_product_deleteBranch_result = -1650651898
)

type MSG_TEST_testsuite_info struct {
	Id int32
	Product int32
	Name string
	Desc string
	Type string
	AddedBy string
	AddedDate time.Time
	LastEditedBy string
	LastEditedDate time.Time
	Deleted bool
}

var pool_MSG_TEST_testsuite_info = sync.Pool{New: func() interface{} { return &MSG_TEST_testsuite_info{} }}

func GET_MSG_TEST_testsuite_info() *MSG_TEST_testsuite_info {
	return pool_MSG_TEST_testsuite_info.Get().(*MSG_TEST_testsuite_info)
}

func (data *MSG_TEST_testsuite_info) cmd() int32 {
	return CMD_MSG_TEST_testsuite_info
}

func (data *MSG_TEST_testsuite_info) Put() {
	data.Id = 0
	data.Product = 0
	data.Name = ``
	data.Desc = ``
	data.Type = ``
	data.AddedBy = ``
	data.AddedDate = time.Unix(0,0)
	data.LastEditedBy = ``
	data.LastEditedDate = time.Unix(0,0)
	data.Deleted = false
	pool_MSG_TEST_testsuite_info.Put(data)
}
func (data *MSG_TEST_testsuite_info) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_testsuite_info,buf)
	WRITE_MSG_TEST_testsuite_info(data, buf)
}

func WRITE_MSG_TEST_testsuite_info(data *MSG_TEST_testsuite_info, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Product, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Desc, buf)
	WRITE_string(data.Type, buf)
	WRITE_string(data.AddedBy, buf)
	WRITE_int64(data.AddedDate.UnixNano(), buf)
	WRITE_string(data.LastEditedBy, buf)
	WRITE_int64(data.LastEditedDate.UnixNano(), buf)
	WRITE_bool(data.Deleted, buf)
}

func READ_MSG_TEST_testsuite_info(buf *libraries.MsgBuffer) *MSG_TEST_testsuite_info {
	data := pool_MSG_TEST_testsuite_info.Get().(*MSG_TEST_testsuite_info)
	data.read(buf)
	return data
}

func (data *MSG_TEST_testsuite_info) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Desc = READ_string(buf)
	data.Type = READ_string(buf)
	data.AddedBy = READ_string(buf)
	data.AddedDate = time.Unix(0, READ_int64(buf))
	data.LastEditedBy = READ_string(buf)
	data.LastEditedDate = time.Unix(0, READ_int64(buf))
	data.Deleted = READ_bool(buf)

}

type MSG_TEST_testsuite_getById struct {
	QueryID uint32
	Id int32
}

var pool_MSG_TEST_testsuite_getById = sync.Pool{New: func() interface{} { return &MSG_TEST_testsuite_getById{} }}

func GET_MSG_TEST_testsuite_getById() *MSG_TEST_testsuite_getById {
	return pool_MSG_TEST_testsuite_getById.Get().(*MSG_TEST_testsuite_getById)
}

func (data *MSG_TEST_testsuite_getById) cmd() int32 {
	return CMD_MSG_TEST_testsuite_getById
}

func (data *MSG_TEST_testsuite_getById) Put() {
	data.QueryID = 0
	data.Id = 0
	pool_MSG_TEST_testsuite_getById.Put(data)
}
func (data *MSG_TEST_testsuite_getById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_testsuite_getById,buf)
	WRITE_MSG_TEST_testsuite_getById(data, buf)
}

func WRITE_MSG_TEST_testsuite_getById(data *MSG_TEST_testsuite_getById, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
}

func READ_MSG_TEST_testsuite_getById(buf *libraries.MsgBuffer) *MSG_TEST_testsuite_getById {
	data := pool_MSG_TEST_testsuite_getById.Get().(*MSG_TEST_testsuite_getById)
	data.read(buf)
	return data
}

func (data *MSG_TEST_testsuite_getById) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)

}
func (data *MSG_TEST_testsuite_getById) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_TEST_testsuite_getById) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_TEST_testsuite_getById_result struct {
	QueryResultID uint32
	Info *MSG_TEST_testsuite_info
}

var pool_MSG_TEST_testsuite_getById_result = sync.Pool{New: func() interface{} { return &MSG_TEST_testsuite_getById_result{} }}

func GET_MSG_TEST_testsuite_getById_result() *MSG_TEST_testsuite_getById_result {
	return pool_MSG_TEST_testsuite_getById_result.Get().(*MSG_TEST_testsuite_getById_result)
}

func (data *MSG_TEST_testsuite_getById_result) cmd() int32 {
	return CMD_MSG_TEST_testsuite_getById_result
}

func (data *MSG_TEST_testsuite_getById_result) Put() {
	data.QueryResultID = 0
	if data.Info != nil {
		data.Info.Put()
		data.Info = nil
	}
	pool_MSG_TEST_testsuite_getById_result.Put(data)
}
func (data *MSG_TEST_testsuite_getById_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_testsuite_getById_result,buf)
	WRITE_MSG_TEST_testsuite_getById_result(data, buf)
}

func WRITE_MSG_TEST_testsuite_getById_result(data *MSG_TEST_testsuite_getById_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	if data.Info == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_TEST_testsuite_info(data.Info, buf)
	}
}

func READ_MSG_TEST_testsuite_getById_result(buf *libraries.MsgBuffer) *MSG_TEST_testsuite_getById_result {
	data := pool_MSG_TEST_testsuite_getById_result.Get().(*MSG_TEST_testsuite_getById_result)
	data.read(buf)
	return data
}

func (data *MSG_TEST_testsuite_getById_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_TEST_testsuite_info(buf)
	}else{
		data.Info = nil
	}

}
func (data *MSG_TEST_testsuite_getById_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_TEST_testsuite_getById_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_TEST_bug_getCount struct {
	QueryID uint32
	Where map[string]interface{}
}

var pool_MSG_TEST_bug_getCount = sync.Pool{New: func() interface{} { return &MSG_TEST_bug_getCount{} }}

func GET_MSG_TEST_bug_getCount() *MSG_TEST_bug_getCount {
	return pool_MSG_TEST_bug_getCount.Get().(*MSG_TEST_bug_getCount)
}

func (data *MSG_TEST_bug_getCount) cmd() int32 {
	return CMD_MSG_TEST_bug_getCount
}

func (data *MSG_TEST_bug_getCount) Put() {
	data.QueryID = 0
	data.Where = nil
	pool_MSG_TEST_bug_getCount.Put(data)
}
func (data *MSG_TEST_bug_getCount) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getCount,buf)
	WRITE_MSG_TEST_bug_getCount(data, buf)
}

func WRITE_MSG_TEST_bug_getCount(data *MSG_TEST_bug_getCount, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_map(data.Where,buf)
}

func READ_MSG_TEST_bug_getCount(buf *libraries.MsgBuffer) *MSG_TEST_bug_getCount {
	data := pool_MSG_TEST_bug_getCount.Get().(*MSG_TEST_bug_getCount)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_getCount) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	READ_map(&data.Where,buf)

}
func (data *MSG_TEST_bug_getCount) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_TEST_bug_getCount) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_TEST_buf_getCount_result struct {
	QueryResultID uint32
	Count int
}

var pool_MSG_TEST_buf_getCount_result = sync.Pool{New: func() interface{} { return &MSG_TEST_buf_getCount_result{} }}

func GET_MSG_TEST_buf_getCount_result() *MSG_TEST_buf_getCount_result {
	return pool_MSG_TEST_buf_getCount_result.Get().(*MSG_TEST_buf_getCount_result)
}

func (data *MSG_TEST_buf_getCount_result) cmd() int32 {
	return CMD_MSG_TEST_buf_getCount_result
}

func (data *MSG_TEST_buf_getCount_result) Put() {
	data.QueryResultID = 0
	data.Count = 0
	pool_MSG_TEST_buf_getCount_result.Put(data)
}
func (data *MSG_TEST_buf_getCount_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_buf_getCount_result,buf)
	WRITE_MSG_TEST_buf_getCount_result(data, buf)
}

func WRITE_MSG_TEST_buf_getCount_result(data *MSG_TEST_buf_getCount_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int(data.Count, buf)
}

func READ_MSG_TEST_buf_getCount_result(buf *libraries.MsgBuffer) *MSG_TEST_buf_getCount_result {
	data := pool_MSG_TEST_buf_getCount_result.Get().(*MSG_TEST_buf_getCount_result)
	data.read(buf)
	return data
}

func (data *MSG_TEST_buf_getCount_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Count = READ_int(buf)

}
func (data *MSG_TEST_buf_getCount_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_TEST_buf_getCount_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_TEST_product_deleteBranch_check struct {
	QueryID uint32
	BranchID int32
}

var pool_MSG_TEST_product_deleteBranch_check = sync.Pool{New: func() interface{} { return &MSG_TEST_product_deleteBranch_check{} }}

func GET_MSG_TEST_product_deleteBranch_check() *MSG_TEST_product_deleteBranch_check {
	return pool_MSG_TEST_product_deleteBranch_check.Get().(*MSG_TEST_product_deleteBranch_check)
}

func (data *MSG_TEST_product_deleteBranch_check) cmd() int32 {
	return CMD_MSG_TEST_product_deleteBranch_check
}

func (data *MSG_TEST_product_deleteBranch_check) Put() {
	data.QueryID = 0
	data.BranchID = 0
	pool_MSG_TEST_product_deleteBranch_check.Put(data)
}
func (data *MSG_TEST_product_deleteBranch_check) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_product_deleteBranch_check,buf)
	WRITE_MSG_TEST_product_deleteBranch_check(data, buf)
}

func WRITE_MSG_TEST_product_deleteBranch_check(data *MSG_TEST_product_deleteBranch_check, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.BranchID, buf)
}

func READ_MSG_TEST_product_deleteBranch_check(buf *libraries.MsgBuffer) *MSG_TEST_product_deleteBranch_check {
	data := pool_MSG_TEST_product_deleteBranch_check.Get().(*MSG_TEST_product_deleteBranch_check)
	data.read(buf)
	return data
}

func (data *MSG_TEST_product_deleteBranch_check) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.BranchID = READ_int32(buf)

}
func (data *MSG_TEST_product_deleteBranch_check) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_TEST_product_deleteBranch_check) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_TEST_product_deleteBranch_result struct {
	QueryResultID uint32
	Result ErrCode
}

var pool_MSG_TEST_product_deleteBranch_result = sync.Pool{New: func() interface{} { return &MSG_TEST_product_deleteBranch_result{} }}

func GET_MSG_TEST_product_deleteBranch_result() *MSG_TEST_product_deleteBranch_result {
	return pool_MSG_TEST_product_deleteBranch_result.Get().(*MSG_TEST_product_deleteBranch_result)
}

func (data *MSG_TEST_product_deleteBranch_result) cmd() int32 {
	return CMD_MSG_TEST_product_deleteBranch_result
}

func (data *MSG_TEST_product_deleteBranch_result) Put() {
	data.QueryResultID = 0
	data.Result = 0
	pool_MSG_TEST_product_deleteBranch_result.Put(data)
}
func (data *MSG_TEST_product_deleteBranch_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_product_deleteBranch_result,buf)
	WRITE_MSG_TEST_product_deleteBranch_result(data, buf)
}

func WRITE_MSG_TEST_product_deleteBranch_result(data *MSG_TEST_product_deleteBranch_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_ErrCode(data.Result, buf)
}

func READ_MSG_TEST_product_deleteBranch_result(buf *libraries.MsgBuffer) *MSG_TEST_product_deleteBranch_result {
	data := pool_MSG_TEST_product_deleteBranch_result.Get().(*MSG_TEST_product_deleteBranch_result)
	data.read(buf)
	return data
}

func (data *MSG_TEST_product_deleteBranch_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Result = READ_ErrCode(buf)

}
func (data *MSG_TEST_product_deleteBranch_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_TEST_product_deleteBranch_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

