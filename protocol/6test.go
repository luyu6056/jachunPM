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
	CMD_MSG_TEST_bug_getById = 671857158
	CMD_MSG_TEST_bug_getById_result = -2057750522
	CMD_MSG_TEST_bug = 1019057158
	CMD_MSG_TEST_CASE_getTaskCasePairs = -225686266
	CMD_MSG_TEST_CASE_getTaskCasePairs_result = -730127098
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
	data.AddedDate = time.UnixMicro(0)
	data.LastEditedBy = 0
	data.LastEditedByAccount = ``
	data.LastEditedDate = time.UnixMicro(0)
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
	WRITE_int64(data.AddedDate.UnixMicro(), buf)
	WRITE_int32(data.LastEditedBy, buf)
	WRITE_string(data.LastEditedByAccount, buf)
	WRITE_int64(data.LastEditedDate.UnixMicro(), buf)
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
	data.AddedDate = time.UnixMicro(READ_int64(buf))
	data.LastEditedBy = READ_int32(buf)
	data.LastEditedByAccount = READ_string(buf)
	data.LastEditedDate = time.UnixMicro(READ_int64(buf))
	data.Deleted = READ_bool(buf)

}

type MSG_TEST_testsuite_getById struct {
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
	data.Id = 0
	pool_MSG_TEST_testsuite_getById.Put(data)
}
func (data *MSG_TEST_testsuite_getById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_testsuite_getById,buf)
	WRITE_MSG_TEST_testsuite_getById(data, buf)
}

func WRITE_MSG_TEST_testsuite_getById(data *MSG_TEST_testsuite_getById, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
}

func READ_MSG_TEST_testsuite_getById(buf *libraries.MsgBuffer) *MSG_TEST_testsuite_getById {
	data := pool_MSG_TEST_testsuite_getById.Get().(*MSG_TEST_testsuite_getById)
	data.read(buf)
	return data
}

func (data *MSG_TEST_testsuite_getById) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)

}

type MSG_TEST_testsuite_getById_result struct {
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
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_TEST_testsuite_info(buf)
	}else{
		data.Info = nil
	}

}

type MSG_TEST_bug_getCount struct {
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
	data.Where = nil
	pool_MSG_TEST_bug_getCount.Put(data)
}
func (data *MSG_TEST_bug_getCount) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getCount,buf)
	WRITE_MSG_TEST_bug_getCount(data, buf)
}

func WRITE_MSG_TEST_bug_getCount(data *MSG_TEST_bug_getCount, buf *libraries.MsgBuffer) {
	WRITE_map(data.Where,buf)
}

func READ_MSG_TEST_bug_getCount(buf *libraries.MsgBuffer) *MSG_TEST_bug_getCount {
	data := pool_MSG_TEST_bug_getCount.Get().(*MSG_TEST_bug_getCount)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_getCount) read(buf *libraries.MsgBuffer) {
	READ_map(&data.Where,buf)

}

type MSG_TEST_buf_getCount_result struct {
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
	data.Count = 0
	pool_MSG_TEST_buf_getCount_result.Put(data)
}
func (data *MSG_TEST_buf_getCount_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_buf_getCount_result,buf)
	WRITE_MSG_TEST_buf_getCount_result(data, buf)
}

func WRITE_MSG_TEST_buf_getCount_result(data *MSG_TEST_buf_getCount_result, buf *libraries.MsgBuffer) {
	WRITE_int(data.Count, buf)
}

func READ_MSG_TEST_buf_getCount_result(buf *libraries.MsgBuffer) *MSG_TEST_buf_getCount_result {
	data := pool_MSG_TEST_buf_getCount_result.Get().(*MSG_TEST_buf_getCount_result)
	data.read(buf)
	return data
}

func (data *MSG_TEST_buf_getCount_result) read(buf *libraries.MsgBuffer) {
	data.Count = READ_int(buf)

}

type MSG_TEST_product_deleteBranch_check struct {
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
	data.BranchID = 0
	pool_MSG_TEST_product_deleteBranch_check.Put(data)
}
func (data *MSG_TEST_product_deleteBranch_check) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_product_deleteBranch_check,buf)
	WRITE_MSG_TEST_product_deleteBranch_check(data, buf)
}

func WRITE_MSG_TEST_product_deleteBranch_check(data *MSG_TEST_product_deleteBranch_check, buf *libraries.MsgBuffer) {
	WRITE_int32(data.BranchID, buf)
}

func READ_MSG_TEST_product_deleteBranch_check(buf *libraries.MsgBuffer) *MSG_TEST_product_deleteBranch_check {
	data := pool_MSG_TEST_product_deleteBranch_check.Get().(*MSG_TEST_product_deleteBranch_check)
	data.read(buf)
	return data
}

func (data *MSG_TEST_product_deleteBranch_check) read(buf *libraries.MsgBuffer) {
	data.BranchID = READ_int32(buf)

}

type MSG_TEST_product_deleteBranch_result struct {
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
	data.Result = 0
	pool_MSG_TEST_product_deleteBranch_result.Put(data)
}
func (data *MSG_TEST_product_deleteBranch_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_product_deleteBranch_result,buf)
	WRITE_MSG_TEST_product_deleteBranch_result(data, buf)
}

func WRITE_MSG_TEST_product_deleteBranch_result(data *MSG_TEST_product_deleteBranch_result, buf *libraries.MsgBuffer) {
	WRITE_ErrCode(data.Result, buf)
}

func READ_MSG_TEST_product_deleteBranch_result(buf *libraries.MsgBuffer) *MSG_TEST_product_deleteBranch_result {
	data := pool_MSG_TEST_product_deleteBranch_result.Get().(*MSG_TEST_product_deleteBranch_result)
	data.read(buf)
	return data
}

func (data *MSG_TEST_product_deleteBranch_result) read(buf *libraries.MsgBuffer) {
	data.Result = READ_ErrCode(buf)

}

type MSG_TEST_bug_updateMapById struct {
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
	data.Id = 0
	data.Update = nil
	pool_MSG_TEST_bug_updateMapById.Put(data)
}
func (data *MSG_TEST_bug_updateMapById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_updateMapById,buf)
	WRITE_MSG_TEST_bug_updateMapById(data, buf)
}

func WRITE_MSG_TEST_bug_updateMapById(data *MSG_TEST_bug_updateMapById, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_map(data.Update,buf)
}

func READ_MSG_TEST_bug_updateMapById(buf *libraries.MsgBuffer) *MSG_TEST_bug_updateMapById {
	data := pool_MSG_TEST_bug_updateMapById.Get().(*MSG_TEST_bug_updateMapById)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_updateMapById) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	READ_map(&data.Update,buf)

}

type MSG_TEST_testtask_getById struct {
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
	data.Id = 0
	pool_MSG_TEST_testtask_getById.Put(data)
}
func (data *MSG_TEST_testtask_getById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_testtask_getById,buf)
	WRITE_MSG_TEST_testtask_getById(data, buf)
}

func WRITE_MSG_TEST_testtask_getById(data *MSG_TEST_testtask_getById, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
}

func READ_MSG_TEST_testtask_getById(buf *libraries.MsgBuffer) *MSG_TEST_testtask_getById {
	data := pool_MSG_TEST_testtask_getById.Get().(*MSG_TEST_testtask_getById)
	data.read(buf)
	return data
}

func (data *MSG_TEST_testtask_getById) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)

}

type MSG_TEST_testtask_getById_result struct {
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
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_TEST_Testtask_info(buf)
	}else{
		data.Info = nil
	}

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
	data.Begin = time.UnixMicro(0)
	data.End = time.UnixMicro(0)
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
	WRITE_int64(data.Begin.UnixMicro(), buf)
	WRITE_int64(data.End.UnixMicro(), buf)
	WRITE_int(len(data.Mailto), buf)
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
	data.Begin = time.UnixMicro(READ_int64(buf))
	data.End = time.UnixMicro(READ_int64(buf))
	Mailto_len := READ_int(buf)
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
	data.Where = nil
	pool_MSG_TEST_bug_getPairs.Put(data)
}
func (data *MSG_TEST_bug_getPairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getPairs,buf)
	WRITE_MSG_TEST_bug_getPairs(data, buf)
}

func WRITE_MSG_TEST_bug_getPairs(data *MSG_TEST_bug_getPairs, buf *libraries.MsgBuffer) {
	WRITE_map(data.Where,buf)
}

func READ_MSG_TEST_bug_getPairs(buf *libraries.MsgBuffer) *MSG_TEST_bug_getPairs {
	data := pool_MSG_TEST_bug_getPairs.Get().(*MSG_TEST_bug_getPairs)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_getPairs) read(buf *libraries.MsgBuffer) {
	READ_map(&data.Where,buf)

}

type MSG_TEST_bug_getPairs_result struct {
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
	data.List = data.List[:0]
	pool_MSG_TEST_bug_getPairs_result.Put(data)
}
func (data *MSG_TEST_bug_getPairs_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getPairs_result,buf)
	WRITE_MSG_TEST_bug_getPairs_result(data, buf)
}

func WRITE_MSG_TEST_bug_getPairs_result(data *MSG_TEST_bug_getPairs_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
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
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]HtmlKeyValueStr, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_HtmlKeyValueStr(buf)
	}

}

type MSG_TEST_bug_getCountByWhere struct {
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
	data.Where = nil
	pool_MSG_TEST_bug_getCountByWhere.Put(data)
}
func (data *MSG_TEST_bug_getCountByWhere) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getCountByWhere,buf)
	WRITE_MSG_TEST_bug_getCountByWhere(data, buf)
}

func WRITE_MSG_TEST_bug_getCountByWhere(data *MSG_TEST_bug_getCountByWhere, buf *libraries.MsgBuffer) {
	WRITE_map(data.Where,buf)
}

func READ_MSG_TEST_bug_getCountByWhere(buf *libraries.MsgBuffer) *MSG_TEST_bug_getCountByWhere {
	data := pool_MSG_TEST_bug_getCountByWhere.Get().(*MSG_TEST_bug_getCountByWhere)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_getCountByWhere) read(buf *libraries.MsgBuffer) {
	READ_map(&data.Where,buf)

}

type MSG_TEST_bug_getCountByWhere_result struct {
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
	data.Count = 0
	pool_MSG_TEST_bug_getCountByWhere_result.Put(data)
}
func (data *MSG_TEST_bug_getCountByWhere_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getCountByWhere_result,buf)
	WRITE_MSG_TEST_bug_getCountByWhere_result(data, buf)
}

func WRITE_MSG_TEST_bug_getCountByWhere_result(data *MSG_TEST_bug_getCountByWhere_result, buf *libraries.MsgBuffer) {
	WRITE_int(data.Count, buf)
}

func READ_MSG_TEST_bug_getCountByWhere_result(buf *libraries.MsgBuffer) *MSG_TEST_bug_getCountByWhere_result {
	data := pool_MSG_TEST_bug_getCountByWhere_result.Get().(*MSG_TEST_bug_getCountByWhere_result)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_getCountByWhere_result) read(buf *libraries.MsgBuffer) {
	data.Count = READ_int(buf)

}

type MSG_TEST_bug_getById struct {
	Id int32
}

var pool_MSG_TEST_bug_getById = sync.Pool{New: func() interface{} { return &MSG_TEST_bug_getById{} }}

func GET_MSG_TEST_bug_getById() *MSG_TEST_bug_getById {
	return pool_MSG_TEST_bug_getById.Get().(*MSG_TEST_bug_getById)
}

func (data *MSG_TEST_bug_getById) cmd() int32 {
	return CMD_MSG_TEST_bug_getById
}

func (data *MSG_TEST_bug_getById) Put() {
	data.Id = 0
	pool_MSG_TEST_bug_getById.Put(data)
}
func (data *MSG_TEST_bug_getById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getById,buf)
	WRITE_MSG_TEST_bug_getById(data, buf)
}

func WRITE_MSG_TEST_bug_getById(data *MSG_TEST_bug_getById, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
}

func READ_MSG_TEST_bug_getById(buf *libraries.MsgBuffer) *MSG_TEST_bug_getById {
	data := pool_MSG_TEST_bug_getById.Get().(*MSG_TEST_bug_getById)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_getById) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)

}

type MSG_TEST_bug_getById_result struct {
	Info *MSG_TEST_bug
}

var pool_MSG_TEST_bug_getById_result = sync.Pool{New: func() interface{} { return &MSG_TEST_bug_getById_result{} }}

func GET_MSG_TEST_bug_getById_result() *MSG_TEST_bug_getById_result {
	return pool_MSG_TEST_bug_getById_result.Get().(*MSG_TEST_bug_getById_result)
}

func (data *MSG_TEST_bug_getById_result) cmd() int32 {
	return CMD_MSG_TEST_bug_getById_result
}

func (data *MSG_TEST_bug_getById_result) Put() {
	if data.Info != nil {
		data.Info.Put()
		data.Info = nil
	}
	pool_MSG_TEST_bug_getById_result.Put(data)
}
func (data *MSG_TEST_bug_getById_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug_getById_result,buf)
	WRITE_MSG_TEST_bug_getById_result(data, buf)
}

func WRITE_MSG_TEST_bug_getById_result(data *MSG_TEST_bug_getById_result, buf *libraries.MsgBuffer) {
	if data.Info == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_TEST_bug(data.Info, buf)
	}
}

func READ_MSG_TEST_bug_getById_result(buf *libraries.MsgBuffer) *MSG_TEST_bug_getById_result {
	data := pool_MSG_TEST_bug_getById_result.Get().(*MSG_TEST_bug_getById_result)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug_getById_result) read(buf *libraries.MsgBuffer) {
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_TEST_bug(buf)
	}else{
		data.Info = nil
	}

}

type MSG_TEST_bug struct {
	Id int32
	Product int32
	Branch int32
	Module int32
	Project int32
	Plan int32
	Story int32
	StoryVersion int16
	Task int32
	ToTask int32
	ToStory int32
	Title string
	Keywords string
	Severity int8
	Pri int8
	Type string
	Os string
	Browser string
	Hardware string
	Found string
	Steps string
	Status string
	Color string
	Confirmed int8
	ActivatedCount int16
	ActivatedDate time.Time
	Mailto []int32
	OpenedBy int32
	OpenedDate time.Time
	OpenedBuild string
	AssignedTo int32
	AssignedDate time.Time
	Deadline time.Time
	ResolvedBy int32
	Resolution string
	ResolvedBuild string
	ResolvedDate time.Time
	ClosedBy int32
	ClosedDate time.Time
	DuplicateBug int32
	LinkBug []int32
	Case int32
	CaseVersion int16
	Result int32
	Testtask int32
	LastEditedBy int32
	LastEditedDate time.Time
	Deleted bool
}

var pool_MSG_TEST_bug = sync.Pool{New: func() interface{} { return &MSG_TEST_bug{} }}

func GET_MSG_TEST_bug() *MSG_TEST_bug {
	return pool_MSG_TEST_bug.Get().(*MSG_TEST_bug)
}

func (data *MSG_TEST_bug) cmd() int32 {
	return CMD_MSG_TEST_bug
}

func (data *MSG_TEST_bug) Put() {
	data.Id = 0
	data.Product = 0
	data.Branch = 0
	data.Module = 0
	data.Project = 0
	data.Plan = 0
	data.Story = 0
	data.StoryVersion = 0
	data.Task = 0
	data.ToTask = 0
	data.ToStory = 0
	data.Title = ``
	data.Keywords = ``
	data.Severity = 0
	data.Pri = 0
	data.Type = ``
	data.Os = ``
	data.Browser = ``
	data.Hardware = ``
	data.Found = ``
	data.Steps = ``
	data.Status = ``
	data.Color = ``
	data.Confirmed = 0
	data.ActivatedCount = 0
	data.ActivatedDate = time.UnixMicro(0)
	data.Mailto = data.Mailto[:0]
	data.OpenedBy = 0
	data.OpenedDate = time.UnixMicro(0)
	data.OpenedBuild = ``
	data.AssignedTo = 0
	data.AssignedDate = time.UnixMicro(0)
	data.Deadline = time.UnixMicro(0)
	data.ResolvedBy = 0
	data.Resolution = ``
	data.ResolvedBuild = ``
	data.ResolvedDate = time.UnixMicro(0)
	data.ClosedBy = 0
	data.ClosedDate = time.UnixMicro(0)
	data.DuplicateBug = 0
	data.LinkBug = data.LinkBug[:0]
	data.Case = 0
	data.CaseVersion = 0
	data.Result = 0
	data.Testtask = 0
	data.LastEditedBy = 0
	data.LastEditedDate = time.UnixMicro(0)
	data.Deleted = false
	pool_MSG_TEST_bug.Put(data)
}
func (data *MSG_TEST_bug) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_bug,buf)
	WRITE_MSG_TEST_bug(data, buf)
}

func WRITE_MSG_TEST_bug(data *MSG_TEST_bug, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_int32(data.Module, buf)
	WRITE_int32(data.Project, buf)
	WRITE_int32(data.Plan, buf)
	WRITE_int32(data.Story, buf)
	WRITE_int16(data.StoryVersion, buf)
	WRITE_int32(data.Task, buf)
	WRITE_int32(data.ToTask, buf)
	WRITE_int32(data.ToStory, buf)
	WRITE_string(data.Title, buf)
	WRITE_string(data.Keywords, buf)
	WRITE_int8(data.Severity, buf)
	WRITE_int8(data.Pri, buf)
	WRITE_string(data.Type, buf)
	WRITE_string(data.Os, buf)
	WRITE_string(data.Browser, buf)
	WRITE_string(data.Hardware, buf)
	WRITE_string(data.Found, buf)
	WRITE_string(data.Steps, buf)
	WRITE_string(data.Status, buf)
	WRITE_string(data.Color, buf)
	WRITE_int8(data.Confirmed, buf)
	WRITE_int16(data.ActivatedCount, buf)
	WRITE_int64(data.ActivatedDate.UnixMicro(), buf)
	WRITE_int(len(data.Mailto), buf)
	for _, v := range data.Mailto{
		WRITE_int32(v, buf)
	}
	WRITE_int32(data.OpenedBy, buf)
	WRITE_int64(data.OpenedDate.UnixMicro(), buf)
	WRITE_string(data.OpenedBuild, buf)
	WRITE_int32(data.AssignedTo, buf)
	WRITE_int64(data.AssignedDate.UnixMicro(), buf)
	WRITE_int64(data.Deadline.UnixMicro(), buf)
	WRITE_int32(data.ResolvedBy, buf)
	WRITE_string(data.Resolution, buf)
	WRITE_string(data.ResolvedBuild, buf)
	WRITE_int64(data.ResolvedDate.UnixMicro(), buf)
	WRITE_int32(data.ClosedBy, buf)
	WRITE_int64(data.ClosedDate.UnixMicro(), buf)
	WRITE_int32(data.DuplicateBug, buf)
	WRITE_int(len(data.LinkBug), buf)
	for _, v := range data.LinkBug{
		WRITE_int32(v, buf)
	}
	WRITE_int32(data.Case, buf)
	WRITE_int16(data.CaseVersion, buf)
	WRITE_int32(data.Result, buf)
	WRITE_int32(data.Testtask, buf)
	WRITE_int32(data.LastEditedBy, buf)
	WRITE_int64(data.LastEditedDate.UnixMicro(), buf)
	WRITE_bool(data.Deleted, buf)
}

func READ_MSG_TEST_bug(buf *libraries.MsgBuffer) *MSG_TEST_bug {
	data := pool_MSG_TEST_bug.Get().(*MSG_TEST_bug)
	data.read(buf)
	return data
}

func (data *MSG_TEST_bug) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.Module = READ_int32(buf)
	data.Project = READ_int32(buf)
	data.Plan = READ_int32(buf)
	data.Story = READ_int32(buf)
	data.StoryVersion = READ_int16(buf)
	data.Task = READ_int32(buf)
	data.ToTask = READ_int32(buf)
	data.ToStory = READ_int32(buf)
	data.Title = READ_string(buf)
	data.Keywords = READ_string(buf)
	data.Severity = READ_int8(buf)
	data.Pri = READ_int8(buf)
	data.Type = READ_string(buf)
	data.Os = READ_string(buf)
	data.Browser = READ_string(buf)
	data.Hardware = READ_string(buf)
	data.Found = READ_string(buf)
	data.Steps = READ_string(buf)
	data.Status = READ_string(buf)
	data.Color = READ_string(buf)
	data.Confirmed = READ_int8(buf)
	data.ActivatedCount = READ_int16(buf)
	data.ActivatedDate = time.UnixMicro(READ_int64(buf))
	Mailto_len := READ_int(buf)
	if Mailto_len>cap(data.Mailto){
		data.Mailto= make([]int32, Mailto_len)
	}else{
		data.Mailto = data.Mailto[:Mailto_len]
	}
	for i := 0; i < Mailto_len; i++ {
		data.Mailto[i] = READ_int32(buf)
	}
	data.OpenedBy = READ_int32(buf)
	data.OpenedDate = time.UnixMicro(READ_int64(buf))
	data.OpenedBuild = READ_string(buf)
	data.AssignedTo = READ_int32(buf)
	data.AssignedDate = time.UnixMicro(READ_int64(buf))
	data.Deadline = time.UnixMicro(READ_int64(buf))
	data.ResolvedBy = READ_int32(buf)
	data.Resolution = READ_string(buf)
	data.ResolvedBuild = READ_string(buf)
	data.ResolvedDate = time.UnixMicro(READ_int64(buf))
	data.ClosedBy = READ_int32(buf)
	data.ClosedDate = time.UnixMicro(READ_int64(buf))
	data.DuplicateBug = READ_int32(buf)
	LinkBug_len := READ_int(buf)
	if LinkBug_len>cap(data.LinkBug){
		data.LinkBug= make([]int32, LinkBug_len)
	}else{
		data.LinkBug = data.LinkBug[:LinkBug_len]
	}
	for i := 0; i < LinkBug_len; i++ {
		data.LinkBug[i] = READ_int32(buf)
	}
	data.Case = READ_int32(buf)
	data.CaseVersion = READ_int16(buf)
	data.Result = READ_int32(buf)
	data.Testtask = READ_int32(buf)
	data.LastEditedBy = READ_int32(buf)
	data.LastEditedDate = time.UnixMicro(READ_int64(buf))
	data.Deleted = READ_bool(buf)

}

type MSG_TEST_CASE_getTaskCasePairs struct {
	Story int32
	StoryVersion int16
}

var pool_MSG_TEST_CASE_getTaskCasePairs = sync.Pool{New: func() interface{} { return &MSG_TEST_CASE_getTaskCasePairs{} }}

func GET_MSG_TEST_CASE_getTaskCasePairs() *MSG_TEST_CASE_getTaskCasePairs {
	return pool_MSG_TEST_CASE_getTaskCasePairs.Get().(*MSG_TEST_CASE_getTaskCasePairs)
}

func (data *MSG_TEST_CASE_getTaskCasePairs) cmd() int32 {
	return CMD_MSG_TEST_CASE_getTaskCasePairs
}

func (data *MSG_TEST_CASE_getTaskCasePairs) Put() {
	data.Story = 0
	data.StoryVersion = 0
	pool_MSG_TEST_CASE_getTaskCasePairs.Put(data)
}
func (data *MSG_TEST_CASE_getTaskCasePairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_CASE_getTaskCasePairs,buf)
	WRITE_MSG_TEST_CASE_getTaskCasePairs(data, buf)
}

func WRITE_MSG_TEST_CASE_getTaskCasePairs(data *MSG_TEST_CASE_getTaskCasePairs, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Story, buf)
	WRITE_int16(data.StoryVersion, buf)
}

func READ_MSG_TEST_CASE_getTaskCasePairs(buf *libraries.MsgBuffer) *MSG_TEST_CASE_getTaskCasePairs {
	data := pool_MSG_TEST_CASE_getTaskCasePairs.Get().(*MSG_TEST_CASE_getTaskCasePairs)
	data.read(buf)
	return data
}

func (data *MSG_TEST_CASE_getTaskCasePairs) read(buf *libraries.MsgBuffer) {
	data.Story = READ_int32(buf)
	data.StoryVersion = READ_int16(buf)

}

type MSG_TEST_CASE_getTaskCasePairs_result struct {
	List []HtmlKeyValueStr
}

var pool_MSG_TEST_CASE_getTaskCasePairs_result = sync.Pool{New: func() interface{} { return &MSG_TEST_CASE_getTaskCasePairs_result{} }}

func GET_MSG_TEST_CASE_getTaskCasePairs_result() *MSG_TEST_CASE_getTaskCasePairs_result {
	return pool_MSG_TEST_CASE_getTaskCasePairs_result.Get().(*MSG_TEST_CASE_getTaskCasePairs_result)
}

func (data *MSG_TEST_CASE_getTaskCasePairs_result) cmd() int32 {
	return CMD_MSG_TEST_CASE_getTaskCasePairs_result
}

func (data *MSG_TEST_CASE_getTaskCasePairs_result) Put() {
	data.List = data.List[:0]
	pool_MSG_TEST_CASE_getTaskCasePairs_result.Put(data)
}
func (data *MSG_TEST_CASE_getTaskCasePairs_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_TEST_CASE_getTaskCasePairs_result,buf)
	WRITE_MSG_TEST_CASE_getTaskCasePairs_result(data, buf)
}

func WRITE_MSG_TEST_CASE_getTaskCasePairs_result(data *MSG_TEST_CASE_getTaskCasePairs_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_TEST_CASE_getTaskCasePairs_result(buf *libraries.MsgBuffer) *MSG_TEST_CASE_getTaskCasePairs_result {
	data := pool_MSG_TEST_CASE_getTaskCasePairs_result.Get().(*MSG_TEST_CASE_getTaskCasePairs_result)
	data.read(buf)
	return data
}

func (data *MSG_TEST_CASE_getTaskCasePairs_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]HtmlKeyValueStr, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_HtmlKeyValueStr(buf)
	}

}

