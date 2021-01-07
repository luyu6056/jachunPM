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
