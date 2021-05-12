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
	CMD_MSG_TEST_bug_updateMapById = 1967635718
	CMD_MSG_TEST_testtask_getById = -999177978
	CMD_MSG_TEST_testtask_getById_result = -582706938
	CMD_MSG_TEST_Testtask_info = -2134257914
	CMD_MSG_TEST_bug_getPairs = -1911489786
	CMD_MSG_TEST_bug_getPairs_result = -1036028410
	CMD_MSG_TEST_bug_getCountByWhere = 76296454
	CMD_MSG_TEST_bug_getCountByWhere_result = 1504611846
)

type MSG_TEST_testsuite_info struct {
	Id int32
	Product int32
	Name string
	Desc string
	Type string
	AddedBy int32
	AddedByAccount string
	AddedDate time.Time
	LastEditedBy int32
	LastEditedByAccount string
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
	data.AddedBy = 0
	data.AddedByAccount = ``
	data.AddedDate = time.Unix(0,0)
	data.LastEditedBy = 0
	data.LastEditedByAccount = ``
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
	WRITE_int32(data.AddedBy, buf)
	WRITE_string(data.AddedByAccount, buf)
	WRITE_int64(data.AddedDate.UnixNano(), buf)
	WRITE_int32(data.LastEditedBy, buf)
	WRITE_string(data.LastEditedByAccount, buf)
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
	data.AddedBy = READ_int32(buf)
	data.AddedByAccount = READ_string(buf)
	data.AddedDate = time.Unix(0, READ_int64(buf))
	data.LastEditedBy = READ_int32(buf)
	data.LastEditedByAccount = READ_string(buf)
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

type MSG_TEST_bug_updateMapById struct {
	QueryID uint32
	Id int32
	Update map[string]interface{}
}

var pool_MSG_TEST_bug_updateMapById = sync.Pool{New: func() interface{} { return &MSG_TEST_bug_updateMapById{} }}

func GET_MSG_TEST_bug_updateMapById() *MSG_TEST_bug_updateMapById {
	return pool_MSG_TEST_bug_updateMapById.Get().(*MSG_TEST_bug_updateMapById)
}

func (data *MSG_TEST_bug_updateMapById) cmd() int32 {
	return CMD_MSG_TEST_bug_updateMapById
}

func (data *MSG_TEST_bug_updateMapById) Put() {
	data.QueryID = 0
	data.Id = 0
	data.Update = nil
	pool_MSG_TEST_bug_updateMapById.Put(data)
}
func (data *MSG_TEST_bug_updateMapById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_updateMapById,buf)
	WRITE_MSG_TEST_bug_updateMapById(data, buf)
}

func WRITE_MSG_TEST_bug_updateMapById(data *MSG_TEST_bug_updateMapById, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
	WRITE_map(data.Update,buf)
}

func READ_MSG_TEST_bug_updateMapById(buf *libraries.MsgBuffer) *MSG_TEST_bug_updateMapById {
	data := pool_MSG_TEST_bug_updateMapById.Get().(*MSG_TEST_bug_updateMapById)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_updateMapById) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)
	READ_map(&data.Update,buf)

}
func (data *MSG_TEST_bug_updateMapById) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_TEST_bug_updateMapById) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_TEST_testtask_getById struct {
	QueryID uint32
	Id int32
}

var pool_MSG_TEST_testtask_getById = sync.Pool{New: func() interface{} { return &MSG_TEST_testtask_getById{} }}

func GET_MSG_TEST_testtask_getById() *MSG_TEST_testtask_getById {
	return pool_MSG_TEST_testtask_getById.Get().(*MSG_TEST_testtask_getById)
}

func (data *MSG_TEST_testtask_getById) cmd() int32 {
	return CMD_MSG_TEST_testtask_getById
}

func (data *MSG_TEST_testtask_getById) Put() {
	data.QueryID = 0
	data.Id = 0
	pool_MSG_TEST_testtask_getById.Put(data)
}
func (data *MSG_TEST_testtask_getById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_testtask_getById,buf)
	WRITE_MSG_TEST_testtask_getById(data, buf)
}

func WRITE_MSG_TEST_testtask_getById(data *MSG_TEST_testtask_getById, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
}

func READ_MSG_TEST_testtask_getById(buf *libraries.MsgBuffer) *MSG_TEST_testtask_getById {
	data := pool_MSG_TEST_testtask_getById.Get().(*MSG_TEST_testtask_getById)
	data.read(buf)
	return data
}

func (data *MSG_TEST_testtask_getById) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)

}
func (data *MSG_TEST_testtask_getById) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_TEST_testtask_getById) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_TEST_testtask_getById_result struct {
	QueryResultID uint32
	Info *MSG_TEST_Testtask_info
}

var pool_MSG_TEST_testtask_getById_result = sync.Pool{New: func() interface{} { return &MSG_TEST_testtask_getById_result{} }}

func GET_MSG_TEST_testtask_getById_result() *MSG_TEST_testtask_getById_result {
	return pool_MSG_TEST_testtask_getById_result.Get().(*MSG_TEST_testtask_getById_result)
}

func (data *MSG_TEST_testtask_getById_result) cmd() int32 {
	return CMD_MSG_TEST_testtask_getById_result
}

func (data *MSG_TEST_testtask_getById_result) Put() {
	data.QueryResultID = 0
	if data.Info != nil {
		data.Info.Put()
		data.Info = nil
	}
	pool_MSG_TEST_testtask_getById_result.Put(data)
}
func (data *MSG_TEST_testtask_getById_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_testtask_getById_result,buf)
	WRITE_MSG_TEST_testtask_getById_result(data, buf)
}

func WRITE_MSG_TEST_testtask_getById_result(data *MSG_TEST_testtask_getById_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	if data.Info == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_TEST_Testtask_info(data.Info, buf)
	}
}

func READ_MSG_TEST_testtask_getById_result(buf *libraries.MsgBuffer) *MSG_TEST_testtask_getById_result {
	data := pool_MSG_TEST_testtask_getById_result.Get().(*MSG_TEST_testtask_getById_result)
	data.read(buf)
	return data
}

func (data *MSG_TEST_testtask_getById_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_TEST_Testtask_info(buf)
	}else{
		data.Info = nil
	}

}
func (data *MSG_TEST_testtask_getById_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_TEST_testtask_getById_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_TEST_Testtask_info struct {
	Id int32
	Name string
	Product int32
	Project int32
	Build string
	OwnerId int32
	Owner string
	Pri int8
	Begin time.Time
	End time.Time
	Mailto []int32
	Desc string
	Report string
	Status string
	Deleted bool
}

var pool_MSG_TEST_Testtask_info = sync.Pool{New: func() interface{} { return &MSG_TEST_Testtask_info{} }}

func GET_MSG_TEST_Testtask_info() *MSG_TEST_Testtask_info {
	return pool_MSG_TEST_Testtask_info.Get().(*MSG_TEST_Testtask_info)
}

func (data *MSG_TEST_Testtask_info) cmd() int32 {
	return CMD_MSG_TEST_Testtask_info
}

func (data *MSG_TEST_Testtask_info) Put() {
	data.Id = 0
	data.Name = ``
	data.Product = 0
	data.Project = 0
	data.Build = ``
	data.OwnerId = 0
	data.Owner = ``
	data.Pri = 0
	data.Begin = time.Unix(0,0)
	data.End = time.Unix(0,0)
	data.Mailto = data.Mailto[:0]
	data.Desc = ``
	data.Report = ``
	data.Status = ``
	data.Deleted = false
	pool_MSG_TEST_Testtask_info.Put(data)
}
func (data *MSG_TEST_Testtask_info) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_Testtask_info,buf)
	WRITE_MSG_TEST_Testtask_info(data, buf)
}

func WRITE_MSG_TEST_Testtask_info(data *MSG_TEST_Testtask_info, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_string(data.Name, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Project, buf)
	WRITE_string(data.Build, buf)
	WRITE_int32(data.OwnerId, buf)
	WRITE_string(data.Owner, buf)
	WRITE_int8(data.Pri, buf)
	WRITE_int64(data.Begin.UnixNano(), buf)
	WRITE_int64(data.End.UnixNano(), buf)
	WRITE_int32(int32(len(data.Mailto)), buf)
	for _, v := range data.Mailto{
		WRITE_int32(v, buf)
	}
	WRITE_string(data.Desc, buf)
	WRITE_string(data.Report, buf)
	WRITE_string(data.Status, buf)
	WRITE_bool(data.Deleted, buf)
}

func READ_MSG_TEST_Testtask_info(buf *libraries.MsgBuffer) *MSG_TEST_Testtask_info {
	data := pool_MSG_TEST_Testtask_info.Get().(*MSG_TEST_Testtask_info)
	data.read(buf)
	return data
}

func (data *MSG_TEST_Testtask_info) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Product = READ_int32(buf)
	data.Project = READ_int32(buf)
	data.Build = READ_string(buf)
	data.OwnerId = READ_int32(buf)
	data.Owner = READ_string(buf)
	data.Pri = READ_int8(buf)
	data.Begin = time.Unix(0, READ_int64(buf))
	data.End = time.Unix(0, READ_int64(buf))
	Mailto_len := int(READ_int32(buf))
	if Mailto_len>cap(data.Mailto){
		data.Mailto= make([]int32, Mailto_len)
	}else{
		data.Mailto = data.Mailto[:Mailto_len]
	}
	for i := 0; i < Mailto_len; i++ {
		data.Mailto[i] = READ_int32(buf)
	}
	data.Desc = READ_string(buf)
	data.Report = READ_string(buf)
	data.Status = READ_string(buf)
	data.Deleted = READ_bool(buf)

}

type MSG_TEST_bug_getPairs struct {
	QueryID uint32
	Where map[string]interface{}
}

var pool_MSG_TEST_bug_getPairs = sync.Pool{New: func() interface{} { return &MSG_TEST_bug_getPairs{} }}

func GET_MSG_TEST_bug_getPairs() *MSG_TEST_bug_getPairs {
	return pool_MSG_TEST_bug_getPairs.Get().(*MSG_TEST_bug_getPairs)
}

func (data *MSG_TEST_bug_getPairs) cmd() int32 {
	return CMD_MSG_TEST_bug_getPairs
}

func (data *MSG_TEST_bug_getPairs) Put() {
	data.QueryID = 0
	data.Where = nil
	pool_MSG_TEST_bug_getPairs.Put(data)
}
func (data *MSG_TEST_bug_getPairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getPairs,buf)
	WRITE_MSG_TEST_bug_getPairs(data, buf)
}

func WRITE_MSG_TEST_bug_getPairs(data *MSG_TEST_bug_getPairs, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_map(data.Where,buf)
}

func READ_MSG_TEST_bug_getPairs(buf *libraries.MsgBuffer) *MSG_TEST_bug_getPairs {
	data := pool_MSG_TEST_bug_getPairs.Get().(*MSG_TEST_bug_getPairs)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_getPairs) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	READ_map(&data.Where,buf)

}
func (data *MSG_TEST_bug_getPairs) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_TEST_bug_getPairs) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_TEST_bug_getPairs_result struct {
	QueryResultID uint32
	List []HtmlKeyValueStr
}

var pool_MSG_TEST_bug_getPairs_result = sync.Pool{New: func() interface{} { return &MSG_TEST_bug_getPairs_result{} }}

func GET_MSG_TEST_bug_getPairs_result() *MSG_TEST_bug_getPairs_result {
	return pool_MSG_TEST_bug_getPairs_result.Get().(*MSG_TEST_bug_getPairs_result)
}

func (data *MSG_TEST_bug_getPairs_result) cmd() int32 {
	return CMD_MSG_TEST_bug_getPairs_result
}

func (data *MSG_TEST_bug_getPairs_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	pool_MSG_TEST_bug_getPairs_result.Put(data)
}
func (data *MSG_TEST_bug_getPairs_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getPairs_result,buf)
	WRITE_MSG_TEST_bug_getPairs_result(data, buf)
}

func WRITE_MSG_TEST_bug_getPairs_result(data *MSG_TEST_bug_getPairs_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_TEST_bug_getPairs_result(buf *libraries.MsgBuffer) *MSG_TEST_bug_getPairs_result {
	data := pool_MSG_TEST_bug_getPairs_result.Get().(*MSG_TEST_bug_getPairs_result)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_getPairs_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	if List_len>cap(data.List){
		data.List= make([]HtmlKeyValueStr, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_HtmlKeyValueStr(buf)
	}

}
func (data *MSG_TEST_bug_getPairs_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_TEST_bug_getPairs_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_TEST_bug_getCountByWhere struct {
	QueryID uint32
	Where map[string]interface{}
}

var pool_MSG_TEST_bug_getCountByWhere = sync.Pool{New: func() interface{} { return &MSG_TEST_bug_getCountByWhere{} }}

func GET_MSG_TEST_bug_getCountByWhere() *MSG_TEST_bug_getCountByWhere {
	return pool_MSG_TEST_bug_getCountByWhere.Get().(*MSG_TEST_bug_getCountByWhere)
}

func (data *MSG_TEST_bug_getCountByWhere) cmd() int32 {
	return CMD_MSG_TEST_bug_getCountByWhere
}

func (data *MSG_TEST_bug_getCountByWhere) Put() {
	data.QueryID = 0
	data.Where = nil
	pool_MSG_TEST_bug_getCountByWhere.Put(data)
}
func (data *MSG_TEST_bug_getCountByWhere) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getCountByWhere,buf)
	WRITE_MSG_TEST_bug_getCountByWhere(data, buf)
}

func WRITE_MSG_TEST_bug_getCountByWhere(data *MSG_TEST_bug_getCountByWhere, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_map(data.Where,buf)
}

func READ_MSG_TEST_bug_getCountByWhere(buf *libraries.MsgBuffer) *MSG_TEST_bug_getCountByWhere {
	data := pool_MSG_TEST_bug_getCountByWhere.Get().(*MSG_TEST_bug_getCountByWhere)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_getCountByWhere) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	READ_map(&data.Where,buf)

}
func (data *MSG_TEST_bug_getCountByWhere) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_TEST_bug_getCountByWhere) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_TEST_bug_getCountByWhere_result struct {
	QueryResultID uint32
	Count int
}

var pool_MSG_TEST_bug_getCountByWhere_result = sync.Pool{New: func() interface{} { return &MSG_TEST_bug_getCountByWhere_result{} }}

func GET_MSG_TEST_bug_getCountByWhere_result() *MSG_TEST_bug_getCountByWhere_result {
	return pool_MSG_TEST_bug_getCountByWhere_result.Get().(*MSG_TEST_bug_getCountByWhere_result)
}

func (data *MSG_TEST_bug_getCountByWhere_result) cmd() int32 {
	return CMD_MSG_TEST_bug_getCountByWhere_result
}

func (data *MSG_TEST_bug_getCountByWhere_result) Put() {
	data.QueryResultID = 0
	data.Count = 0
	pool_MSG_TEST_bug_getCountByWhere_result.Put(data)
}
func (data *MSG_TEST_bug_getCountByWhere_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getCountByWhere_result,buf)
	WRITE_MSG_TEST_bug_getCountByWhere_result(data, buf)
}

func WRITE_MSG_TEST_bug_getCountByWhere_result(data *MSG_TEST_bug_getCountByWhere_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int(data.Count, buf)
}

func READ_MSG_TEST_bug_getCountByWhere_result(buf *libraries.MsgBuffer) *MSG_TEST_bug_getCountByWhere_result {
	data := pool_MSG_TEST_bug_getCountByWhere_result.Get().(*MSG_TEST_bug_getCountByWhere_result)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_getCountByWhere_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Count = READ_int(buf)

}
func (data *MSG_TEST_bug_getCountByWhere_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_TEST_bug_getCountByWhere_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

