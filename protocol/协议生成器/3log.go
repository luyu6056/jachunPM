package main

import "time"

type MSG_LOG_Ation struct {
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
	historys   []*MSG_LOG_History
}
type MSG_LOG_History struct {
	Field      string
	Old        string
	New        string
	Diff       string
	FieldLabel string
}
