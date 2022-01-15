package main

import "time"

type MSG_LOG_Action struct {
	Id         int64
	ObjectType string
	ObjectID   int32
	Product    int32
	Project    int32
	ActorId    int32
	Actor      string
	Action     string
	Date       time.Time
	Comment    string
	Extra      string
	Read       bool
	Historys   []*MSG_LOG_History
	AppendLink string `db:"-"`
}
type MSG_LOG_History struct {
	Field      string
	Old        string
	New        string
	Diff       string
	FieldLabel string
}
type MSG_LOG_Action_Create struct {
	ObjectType string
	ObjectID   int32
	ActionType string //操作类型
	Comment    string //信息
	Extra      string //额外信息
	ActorId    int32  //操作者
	Products   []int32
	Project    int32
}
type MSG_LOG_Action_Create_result struct {
	ActionId int64
}
type MSG_LOG_Action_GetByWhereMap struct {
	Star  time.Time
	Where map[string]interface{}
	Order string
}
type MSG_LOG_Action_GetByID struct {
	Id int64
}
type MSG_LOG_Action_GetByID_result struct {
	Info *MSG_LOG_Action
}

type MSG_LOG_Action_GetByWhereMap_result struct {
	List []*MSG_LOG_Action
}
type MSG_LOG_Action_transformActions struct {
	Where   map[string]interface{}
	Order   string
	Page    int
	PerPage int
	Total   int
}
type MSG_LOG_Action_transformActions_result struct {
	List  []*MSG_LOG_transformActions_info
	Total int
}
type MSG_LOG_transformActions_info struct {
	Id         int64
	ObjectType string
	ObjectID   int32
	Product    int32
	Project    int32
	ActorId    int32
	Actor      string
	Action     string
	Date       time.Time
	Comment    string
	Extra      string
	Read       bool
	//Historys    []*MSG_LOG_History
	ActionLabel  string `db:"-"`
	ObjectLabel  string `db:"-"`
	ObjectName   string `db:"-"`
	OriginalDate string `db:"-"`
	ObjectLink   string `db:"-"`
	Major        bool   `db:"-"`
}
type MSG_LOG_Action_AddHistory struct {
	Id      int64
	History []*MSG_LOG_History
}

type MSG_LOG_Action_set_read struct {
	ObjectType string
	ObjectID   int32
}
