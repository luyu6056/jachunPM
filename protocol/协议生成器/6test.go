package main

import "time"

type MSG_TEST_testsuite_info struct {
	Id             int32
	Product        int32
	Name           string
	Desc           string
	Type           string
	AddedBy        string
	AddedDate      time.Time
	LastEditedBy   string
	LastEditedDate time.Time
	Deleted        bool
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
