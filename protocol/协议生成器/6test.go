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
	Id int32
}

type MSG_TEST_testsuite_getById_result struct {
	Info *MSG_TEST_testsuite_info
}
type MSG_TEST_bug_getCount struct {
	Where map[string]interface{}
}
type MSG_TEST_buf_getCount_result struct {
	Count int
}
type MSG_TEST_product_deleteBranch_check struct {
	BranchID int32
}
type MSG_TEST_product_deleteBranch_result struct {
	Result ErrCode
}
type MSG_TEST_bug_updateMapById struct {
	Id     int32
	Update map[string]interface{}
}
type MSG_TEST_testtask_getById struct {
	Id int32
}
type MSG_TEST_testtask_getById_result struct {
	Info *MSG_TEST_Testtask_info
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
	Where map[string]interface{}
}
type MSG_TEST_bug_getPairs_result struct {
	List []HtmlKeyValueStr
}
type MSG_TEST_bug_getCountByWhere struct {
	Where map[string]interface{}
}
type MSG_TEST_bug_getCountByWhere_result struct {
	Count int
}
type MSG_TEST_bug_getById struct {
	Id int32
}
type MSG_TEST_bug_getById_result struct {
	Info *MSG_TEST_bug
}
type MSG_TEST_bug struct {
	Id             int32
	Product        int32
	Branch         int32
	Module         int32
	Project        int32
	Plan           int32
	Story          int32
	StoryVersion   int16
	Task           int32
	ToTask         int32
	ToStory        int32
	Title          string
	Keywords       string
	Severity       int8
	Pri            int8
	Type           string
	Os             string
	Browser        string
	Hardware       string
	Found          string
	Steps          string
	Status         string
	Color          string
	Confirmed      int8
	ActivatedCount int16
	ActivatedDate  time.Time
	Mailto         []int32
	OpenedBy       int32
	OpenedDate     time.Time
	OpenedBuild    string
	AssignedTo     int32
	AssignedDate   time.Time
	Deadline       time.Time
	ResolvedBy     int32
	Resolution     string
	ResolvedBuild  string
	ResolvedDate   time.Time
	ClosedBy       int32
	ClosedDate     time.Time
	DuplicateBug   int32
	LinkBug        []int32
	Case           int32
	CaseVersion    int16
	Result         int32
	Testtask       int32
	LastEditedBy   int32
	LastEditedDate time.Time
	Deleted        bool
}
type MSG_TEST_CASE_getTaskCasePairs struct {
	Story        int32
	StoryVersion int16
}
type MSG_TEST_CASE_getTaskCasePairs_result struct {
	List []HtmlKeyValueStr
}
