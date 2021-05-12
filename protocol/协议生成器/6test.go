package main

import "time"

type MSG_TEST_testsuite_info struct {
	Id                  int32
	Product             int32
	Name                string
	Desc                string
	Type                string
	AddedBy             int32
	AddedByAccount      string
	AddedDate           time.Time
	LastEditedBy        int32
	LastEditedByAccount string
	LastEditedDate      time.Time
	Deleted             bool
}
type MSG_TEST_testsuite_getById struct {
	QueryID uint32
	Id      int32
}

type MSG_TEST_testsuite_getById_result struct {
	QueryResultID uint32
	Info          *MSG_TEST_testsuite_info
}
type MSG_TEST_bug_getCount struct {
	QueryID uint32
	Where   map[string]interface{}
}
type MSG_TEST_buf_getCount_result struct {
	QueryResultID uint32
	Count         int
}
type MSG_TEST_product_deleteBranch_check struct {
	QueryID  uint32
	BranchID int32
}
type MSG_TEST_product_deleteBranch_result struct {
	QueryResultID uint32
	Result        ErrCode
}
type MSG_TEST_bug_updateMapById struct {
	QueryID uint32
	Id      int32
	Update  map[string]interface{}
}
type MSG_TEST_testtask_getById struct {
	QueryID uint32
	Id      int32
}
type MSG_TEST_testtask_getById_result struct {
	QueryResultID uint32
	Info          *MSG_TEST_Testtask_info
}
type MSG_TEST_Testtask_info struct {
	Id      int32
	Name    string
	Product int32
	Project int32
	Build   string
	OwnerId int32
	Owner   string
	Pri     int8
	Begin   time.Time
	End     time.Time
	Mailto  []int32
	Desc    string
	Report  string
	Status  string
	Deleted bool
}

type MSG_TEST_bug_getPairs struct {
	QueryID uint32
	Where   map[string]interface{}
}
type MSG_TEST_bug_getPairs_result struct {
	QueryResultID uint32
	List          []HtmlKeyValueStr
}
type MSG_TEST_bug_getCountByWhere struct {
	QueryID uint32
	Where   map[string]interface{}
}
type MSG_TEST_bug_getCountByWhere_result struct {
	QueryResultID uint32
	Count         int
}
