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
	QueryID    uint32
	ObjectType string
	ObjectID   int32
	ActionType string //操作类型
	Comment    string //信息
	Extra      string //额外信息
	ActorId    int32  //操作者
	Products   []int32
	Projects   []int32
}
type MSG_LOG_Action_Create_result struct {
	QueryResultID uint32
	ActionId      int64
}
type MSG_LOG_Action_GetByWhereMap struct {
	QueryID uint32
	Where   map[string]interface{}
	Order   string
}
type MSG_LOG_Action_GetByID struct {
	QueryID uint32
	Id      int64
}
type MSG_LOG_Action_GetByID_result struct {
	QueryResultID uint32
	Info          *MSG_LOG_Action
}

type MSG_LOG_Action_GetByWhereMap_result struct {
	QueryResultID uint32
	List          []*MSG_LOG_Action
}
type MSG_LOG_Action_transformActions struct {
	QueryID uint32
	Where   map[string]interface{}
	Order   string
}
type MSG_LOG_Action_transformActions_result struct {
	QueryResultID uint32
	List          []*MSG_LOG_transformActions_info
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
	QueryID uint32
	Id      int64
	History []*MSG_LOG_History
}
