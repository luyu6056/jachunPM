package protocol

import (
	"sync"
	"libraries"
	"time"
)

const (
	CMD_MSG_PROJECT_tree_getLinePairs = -1108380155
	CMD_MSG_PROJECT_tree_getLinePairs_result = -1262905851
	CMD_MSG_PROJECT_product_cache = -553153019
	CMD_MSG_PROJECT_product_insert = -504988411
	CMD_MSG_PROJECT_product_insert_result = -686336507
	CMD_MSG_PROJECT_product_getStories = 1365829125
	CMD_MSG_PROJECT_product_getStories_result = 528481285
	CMD_MSG_PROJECT_story = 1713425925
	CMD_MSG_PROJECT_tree_cache = -1861023995
	CMD_MSG_PROJECT_tree_getParents = -2065363451
	CMD_MSG_PROJECT_tree_getParents_result = 40064261
	CMD_MSG_PROJECT_branch_info = 894530565
	CMD_MSG_PROJECT_tree_manageChild = 727801093
	CMD_MSG_PROJECT_tree_manageChild_result = -2003846907
)

type MSG_PROJECT_tree_getLinePairs struct {
	QueryID uint32
}

var pool_MSG_PROJECT_tree_getLinePairs = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_getLinePairs{} }}

func GET_MSG_PROJECT_tree_getLinePairs() *MSG_PROJECT_tree_getLinePairs {
	return pool_MSG_PROJECT_tree_getLinePairs.Get().(*MSG_PROJECT_tree_getLinePairs)
}

func (data *MSG_PROJECT_tree_getLinePairs) cmd() int32 {
	return CMD_MSG_PROJECT_tree_getLinePairs
}

func (data *MSG_PROJECT_tree_getLinePairs) Put() {
	data.QueryID = 0
	pool_MSG_PROJECT_tree_getLinePairs.Put(data)
}
func (data *MSG_PROJECT_tree_getLinePairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_getLinePairs,buf)
	WRITE_MSG_PROJECT_tree_getLinePairs(data, buf)
}

func WRITE_MSG_PROJECT_tree_getLinePairs(data *MSG_PROJECT_tree_getLinePairs, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
}

func READ_MSG_PROJECT_tree_getLinePairs(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_getLinePairs {
	data := pool_MSG_PROJECT_tree_getLinePairs.Get().(*MSG_PROJECT_tree_getLinePairs)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_getLinePairs) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)

}
func (data *MSG_PROJECT_tree_getLinePairs) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_tree_getLinePairs) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_tree_getLinePairs_result struct {
	QueryResultID uint32
	List []HtmlKeyValueStr
}

var pool_MSG_PROJECT_tree_getLinePairs_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_getLinePairs_result{} }}

func GET_MSG_PROJECT_tree_getLinePairs_result() *MSG_PROJECT_tree_getLinePairs_result {
	return pool_MSG_PROJECT_tree_getLinePairs_result.Get().(*MSG_PROJECT_tree_getLinePairs_result)
}

func (data *MSG_PROJECT_tree_getLinePairs_result) cmd() int32 {
	return CMD_MSG_PROJECT_tree_getLinePairs_result
}

func (data *MSG_PROJECT_tree_getLinePairs_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	pool_MSG_PROJECT_tree_getLinePairs_result.Put(data)
}
func (data *MSG_PROJECT_tree_getLinePairs_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_getLinePairs_result,buf)
	WRITE_MSG_PROJECT_tree_getLinePairs_result(data, buf)
}

func WRITE_MSG_PROJECT_tree_getLinePairs_result(data *MSG_PROJECT_tree_getLinePairs_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_PROJECT_tree_getLinePairs_result(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_getLinePairs_result {
	data := pool_MSG_PROJECT_tree_getLinePairs_result.Get().(*MSG_PROJECT_tree_getLinePairs_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_getLinePairs_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	for i := 0; i < List_len; i++ {
		data.List = append(data.List, READ_HtmlKeyValueStr(buf))
	}

}
func (data *MSG_PROJECT_tree_getLinePairs_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_tree_getLinePairs_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_product_cache struct {
	Id int32
	Name string
	Code string
	Line int32
	Type string
	Status string
	Desc string
	PO int32
	QD int32
	RD int32
	Acl string
	Whitelist []int32
	CreatedBy int32
	CreatedDate int64
	Order int32
	Deleted bool
	TimeStamp int64
	Branchs []*MSG_PROJECT_branch_info
}

var pool_MSG_PROJECT_product_cache = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_cache{} }}

func GET_MSG_PROJECT_product_cache() *MSG_PROJECT_product_cache {
	return pool_MSG_PROJECT_product_cache.Get().(*MSG_PROJECT_product_cache)
}

func (data *MSG_PROJECT_product_cache) cmd() int32 {
	return CMD_MSG_PROJECT_product_cache
}

func (data *MSG_PROJECT_product_cache) Put() {
	data.Id = 0
	data.Name = ``
	data.Code = ``
	data.Line = 0
	data.Type = ``
	data.Status = ``
	data.Desc = ``
	data.PO = 0
	data.QD = 0
	data.RD = 0
	data.Acl = ``
	data.Whitelist = data.Whitelist[:0]
	data.CreatedBy = 0
	data.CreatedDate = 0
	data.Order = 0
	data.Deleted = false
	data.TimeStamp = 0
	for _,v := range data.Branchs {
		v.Put()
	}
	data.Branchs = data.Branchs[:0]
	pool_MSG_PROJECT_product_cache.Put(data)
}
func (data *MSG_PROJECT_product_cache) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_cache,buf)
	WRITE_MSG_PROJECT_product_cache(data, buf)
}

func WRITE_MSG_PROJECT_product_cache(data *MSG_PROJECT_product_cache, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Code, buf)
	WRITE_int32(data.Line, buf)
	WRITE_string(data.Type, buf)
	WRITE_string(data.Status, buf)
	WRITE_string(data.Desc, buf)
	WRITE_int32(data.PO, buf)
	WRITE_int32(data.QD, buf)
	WRITE_int32(data.RD, buf)
	WRITE_string(data.Acl, buf)
	WRITE_int32(int32(len(data.Whitelist)), buf)
	for _, v := range data.Whitelist{
		WRITE_int32(v, buf)
	}
	WRITE_int32(data.CreatedBy, buf)
	WRITE_int64(data.CreatedDate, buf)
	WRITE_int32(data.Order, buf)
	WRITE_bool(data.Deleted, buf)
	WRITE_int64(data.TimeStamp, buf)
	WRITE_int32(int32(len(data.Branchs)), buf)
	for _, v := range data.Branchs{
		WRITE_MSG_PROJECT_branch_info(v, buf)
	}
}

func READ_MSG_PROJECT_product_cache(buf *libraries.MsgBuffer) *MSG_PROJECT_product_cache {
	data := pool_MSG_PROJECT_product_cache.Get().(*MSG_PROJECT_product_cache)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_cache) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Code = READ_string(buf)
	data.Line = READ_int32(buf)
	data.Type = READ_string(buf)
	data.Status = READ_string(buf)
	data.Desc = READ_string(buf)
	data.PO = READ_int32(buf)
	data.QD = READ_int32(buf)
	data.RD = READ_int32(buf)
	data.Acl = READ_string(buf)
	Whitelist_len := int(READ_int32(buf))
	for i := 0; i < Whitelist_len; i++ {
		data.Whitelist = append(data.Whitelist, READ_int32(buf))
	}
	data.CreatedBy = READ_int32(buf)
	data.CreatedDate = READ_int64(buf)
	data.Order = READ_int32(buf)
	data.Deleted = READ_bool(buf)
	data.TimeStamp = READ_int64(buf)
	Branchs_len := int(READ_int32(buf))
	for i := 0; i < Branchs_len; i++ {
		data.Branchs = append(data.Branchs, READ_MSG_PROJECT_branch_info(buf))
	}

}

type MSG_PROJECT_product_insert struct {
	QueryID uint32
	DocName string
	Data *MSG_PROJECT_product_cache
}

var pool_MSG_PROJECT_product_insert = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_insert{} }}

func GET_MSG_PROJECT_product_insert() *MSG_PROJECT_product_insert {
	return pool_MSG_PROJECT_product_insert.Get().(*MSG_PROJECT_product_insert)
}

func (data *MSG_PROJECT_product_insert) cmd() int32 {
	return CMD_MSG_PROJECT_product_insert
}

func (data *MSG_PROJECT_product_insert) Put() {
	data.QueryID = 0
	data.DocName = ``
	if data.Data != nil {
		data.Data.Put()
		data.Data = nil
	}
	pool_MSG_PROJECT_product_insert.Put(data)
}
func (data *MSG_PROJECT_product_insert) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_insert,buf)
	WRITE_MSG_PROJECT_product_insert(data, buf)
}

func WRITE_MSG_PROJECT_product_insert(data *MSG_PROJECT_product_insert, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_string(data.DocName, buf)
	if data.Data == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_PROJECT_product_cache(data.Data, buf)
	}
}

func READ_MSG_PROJECT_product_insert(buf *libraries.MsgBuffer) *MSG_PROJECT_product_insert {
	data := pool_MSG_PROJECT_product_insert.Get().(*MSG_PROJECT_product_insert)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_insert) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.DocName = READ_string(buf)
	Data_len := int(READ_int8(buf))
	if Data_len == 1 {
		data.Data = READ_MSG_PROJECT_product_cache(buf)
	}else{
		data.Data = nil
	}

}
func (data *MSG_PROJECT_product_insert) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_product_insert) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_product_insert_result struct {
	QueryResultID uint32
	ID int32
}

var pool_MSG_PROJECT_product_insert_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_insert_result{} }}

func GET_MSG_PROJECT_product_insert_result() *MSG_PROJECT_product_insert_result {
	return pool_MSG_PROJECT_product_insert_result.Get().(*MSG_PROJECT_product_insert_result)
}

func (data *MSG_PROJECT_product_insert_result) cmd() int32 {
	return CMD_MSG_PROJECT_product_insert_result
}

func (data *MSG_PROJECT_product_insert_result) Put() {
	data.QueryResultID = 0
	data.ID = 0
	pool_MSG_PROJECT_product_insert_result.Put(data)
}
func (data *MSG_PROJECT_product_insert_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_insert_result,buf)
	WRITE_MSG_PROJECT_product_insert_result(data, buf)
}

func WRITE_MSG_PROJECT_product_insert_result(data *MSG_PROJECT_product_insert_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(data.ID, buf)
}

func READ_MSG_PROJECT_product_insert_result(buf *libraries.MsgBuffer) *MSG_PROJECT_product_insert_result {
	data := pool_MSG_PROJECT_product_insert_result.Get().(*MSG_PROJECT_product_insert_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_insert_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.ID = READ_int32(buf)

}
func (data *MSG_PROJECT_product_insert_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_product_insert_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_product_getStories struct {
	QueryID uint32
	ProductID int32
	Branch int32
	BrowseType string
	ModuleID int32
	Sort string
	Page int
	PerPage int
	Where string
	Total int
}

var pool_MSG_PROJECT_product_getStories = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_getStories{} }}

func GET_MSG_PROJECT_product_getStories() *MSG_PROJECT_product_getStories {
	return pool_MSG_PROJECT_product_getStories.Get().(*MSG_PROJECT_product_getStories)
}

func (data *MSG_PROJECT_product_getStories) cmd() int32 {
	return CMD_MSG_PROJECT_product_getStories
}

func (data *MSG_PROJECT_product_getStories) Put() {
	data.QueryID = 0
	data.ProductID = 0
	data.Branch = 0
	data.BrowseType = ``
	data.ModuleID = 0
	data.Sort = ``
	data.Page = 0
	data.PerPage = 0
	data.Where = ``
	data.Total = 0
	pool_MSG_PROJECT_product_getStories.Put(data)
}
func (data *MSG_PROJECT_product_getStories) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_getStories,buf)
	WRITE_MSG_PROJECT_product_getStories(data, buf)
}

func WRITE_MSG_PROJECT_product_getStories(data *MSG_PROJECT_product_getStories, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.ProductID, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_string(data.BrowseType, buf)
	WRITE_int32(data.ModuleID, buf)
	WRITE_string(data.Sort, buf)
	WRITE_int(data.Page, buf)
	WRITE_int(data.PerPage, buf)
	WRITE_string(data.Where, buf)
	WRITE_int(data.Total, buf)
}

func READ_MSG_PROJECT_product_getStories(buf *libraries.MsgBuffer) *MSG_PROJECT_product_getStories {
	data := pool_MSG_PROJECT_product_getStories.Get().(*MSG_PROJECT_product_getStories)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_getStories) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.ProductID = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.BrowseType = READ_string(buf)
	data.ModuleID = READ_int32(buf)
	data.Sort = READ_string(buf)
	data.Page = READ_int(buf)
	data.PerPage = READ_int(buf)
	data.Where = READ_string(buf)
	data.Total = READ_int(buf)

}
func (data *MSG_PROJECT_product_getStories) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_product_getStories) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_product_getStories_result struct {
	QueryResultID uint32
	List []*MSG_PROJECT_story
	Total int
}

var pool_MSG_PROJECT_product_getStories_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_getStories_result{} }}

func GET_MSG_PROJECT_product_getStories_result() *MSG_PROJECT_product_getStories_result {
	return pool_MSG_PROJECT_product_getStories_result.Get().(*MSG_PROJECT_product_getStories_result)
}

func (data *MSG_PROJECT_product_getStories_result) cmd() int32 {
	return CMD_MSG_PROJECT_product_getStories_result
}

func (data *MSG_PROJECT_product_getStories_result) Put() {
	data.QueryResultID = 0
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	data.Total = 0
	pool_MSG_PROJECT_product_getStories_result.Put(data)
}
func (data *MSG_PROJECT_product_getStories_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_getStories_result,buf)
	WRITE_MSG_PROJECT_product_getStories_result(data, buf)
}

func WRITE_MSG_PROJECT_product_getStories_result(data *MSG_PROJECT_product_getStories_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_MSG_PROJECT_story(v, buf)
	}
	WRITE_int(data.Total, buf)
}

func READ_MSG_PROJECT_product_getStories_result(buf *libraries.MsgBuffer) *MSG_PROJECT_product_getStories_result {
	data := pool_MSG_PROJECT_product_getStories_result.Get().(*MSG_PROJECT_product_getStories_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_getStories_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	for i := 0; i < List_len; i++ {
		data.List = append(data.List, READ_MSG_PROJECT_story(buf))
	}
	data.Total = READ_int(buf)

}
func (data *MSG_PROJECT_product_getStories_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_product_getStories_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_story struct {
	Id int32
	Product int32
	Branch int32
	Module int32
	Plan string
	Source string
	SourceNote string
	FromBug int32
	Title string
	Keywords string
	Pri int8
	Estimate float32
	Status string
	Stage string
	Mailto string
	OpenedBy string
	OpenedDate time.Time
	AssignedTo string
	AssignedDate time.Time
	LastEditedBy string
	LastEditedDate time.Time
	ReviewedBy string
	ReviewedDate time.Time
	ClosedBy string
	ClosedDate time.Time
	ClosedReason string
	ToBug int32
	ChildStories string
	LinkStories string
	DuplicateStory int32
	Deleted bool
	Version int16
}

var pool_MSG_PROJECT_story = sync.Pool{New: func() interface{} { return &MSG_PROJECT_story{} }}

func GET_MSG_PROJECT_story() *MSG_PROJECT_story {
	return pool_MSG_PROJECT_story.Get().(*MSG_PROJECT_story)
}

func (data *MSG_PROJECT_story) cmd() int32 {
	return CMD_MSG_PROJECT_story
}

func (data *MSG_PROJECT_story) Put() {
	data.Id = 0
	data.Product = 0
	data.Branch = 0
	data.Module = 0
	data.Plan = ``
	data.Source = ``
	data.SourceNote = ``
	data.FromBug = 0
	data.Title = ``
	data.Keywords = ``
	data.Pri = 0
	data.Status = ``
	data.Stage = ``
	data.Mailto = ``
	data.OpenedBy = ``
	data.OpenedDate = time.Unix(0,0)
	data.AssignedTo = ``
	data.AssignedDate = time.Unix(0,0)
	data.LastEditedBy = ``
	data.LastEditedDate = time.Unix(0,0)
	data.ReviewedBy = ``
	data.ReviewedDate = time.Unix(0,0)
	data.ClosedBy = ``
	data.ClosedDate = time.Unix(0,0)
	data.ClosedReason = ``
	data.ToBug = 0
	data.ChildStories = ``
	data.LinkStories = ``
	data.DuplicateStory = 0
	data.Deleted = false
	data.Version = 0
	pool_MSG_PROJECT_story.Put(data)
}
func (data *MSG_PROJECT_story) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_story,buf)
	WRITE_MSG_PROJECT_story(data, buf)
}

func WRITE_MSG_PROJECT_story(data *MSG_PROJECT_story, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_int32(data.Module, buf)
	WRITE_string(data.Plan, buf)
	WRITE_string(data.Source, buf)
	WRITE_string(data.SourceNote, buf)
	WRITE_int32(data.FromBug, buf)
	WRITE_string(data.Title, buf)
	WRITE_string(data.Keywords, buf)
	WRITE_int8(data.Pri, buf)
	WRITE_float32(data.Estimate, buf)
	WRITE_string(data.Status, buf)
	WRITE_string(data.Stage, buf)
	WRITE_string(data.Mailto, buf)
	WRITE_string(data.OpenedBy, buf)
	WRITE_int64(data.OpenedDate.UnixNano(), buf)
	WRITE_string(data.AssignedTo, buf)
	WRITE_int64(data.AssignedDate.UnixNano(), buf)
	WRITE_string(data.LastEditedBy, buf)
	WRITE_int64(data.LastEditedDate.UnixNano(), buf)
	WRITE_string(data.ReviewedBy, buf)
	WRITE_int64(data.ReviewedDate.UnixNano(), buf)
	WRITE_string(data.ClosedBy, buf)
	WRITE_int64(data.ClosedDate.UnixNano(), buf)
	WRITE_string(data.ClosedReason, buf)
	WRITE_int32(data.ToBug, buf)
	WRITE_string(data.ChildStories, buf)
	WRITE_string(data.LinkStories, buf)
	WRITE_int32(data.DuplicateStory, buf)
	WRITE_bool(data.Deleted, buf)
	WRITE_int16(data.Version, buf)
}

func READ_MSG_PROJECT_story(buf *libraries.MsgBuffer) *MSG_PROJECT_story {
	data := pool_MSG_PROJECT_story.Get().(*MSG_PROJECT_story)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_story) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.Module = READ_int32(buf)
	data.Plan = READ_string(buf)
	data.Source = READ_string(buf)
	data.SourceNote = READ_string(buf)
	data.FromBug = READ_int32(buf)
	data.Title = READ_string(buf)
	data.Keywords = READ_string(buf)
	data.Pri = READ_int8(buf)
	data.Estimate = READ_float32(buf)
	data.Status = READ_string(buf)
	data.Stage = READ_string(buf)
	data.Mailto = READ_string(buf)
	data.OpenedBy = READ_string(buf)
	data.OpenedDate = time.Unix(0, READ_int64(buf))
	data.AssignedTo = READ_string(buf)
	data.AssignedDate = time.Unix(0, READ_int64(buf))
	data.LastEditedBy = READ_string(buf)
	data.LastEditedDate = time.Unix(0, READ_int64(buf))
	data.ReviewedBy = READ_string(buf)
	data.ReviewedDate = time.Unix(0, READ_int64(buf))
	data.ClosedBy = READ_string(buf)
	data.ClosedDate = time.Unix(0, READ_int64(buf))
	data.ClosedReason = READ_string(buf)
	data.ToBug = READ_int32(buf)
	data.ChildStories = READ_string(buf)
	data.LinkStories = READ_string(buf)
	data.DuplicateStory = READ_int32(buf)
	data.Deleted = READ_bool(buf)
	data.Version = READ_int16(buf)

}

type MSG_PROJECT_tree_cache struct {
	Id int32
	Root int32
	Branch int32
	Name string
	Parent int32
	Path []int32
	Grade int8
	Order int16
	Type string
	Owner string
	Collector string
	Short string
	Deleted bool
}

var pool_MSG_PROJECT_tree_cache = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_cache{} }}

func GET_MSG_PROJECT_tree_cache() *MSG_PROJECT_tree_cache {
	return pool_MSG_PROJECT_tree_cache.Get().(*MSG_PROJECT_tree_cache)
}

func (data *MSG_PROJECT_tree_cache) cmd() int32 {
	return CMD_MSG_PROJECT_tree_cache
}

func (data *MSG_PROJECT_tree_cache) Put() {
	data.Id = 0
	data.Root = 0
	data.Branch = 0
	data.Name = ``
	data.Parent = 0
	data.Path = data.Path[:0]
	data.Grade = 0
	data.Order = 0
	data.Type = ``
	data.Owner = ``
	data.Collector = ``
	data.Short = ``
	data.Deleted = false
	pool_MSG_PROJECT_tree_cache.Put(data)
}
func (data *MSG_PROJECT_tree_cache) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_cache,buf)
	WRITE_MSG_PROJECT_tree_cache(data, buf)
}

func WRITE_MSG_PROJECT_tree_cache(data *MSG_PROJECT_tree_cache, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Root, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_string(data.Name, buf)
	WRITE_int32(data.Parent, buf)
	WRITE_int32(int32(len(data.Path)), buf)
	for _, v := range data.Path{
		WRITE_int32(v, buf)
	}
	WRITE_int8(data.Grade, buf)
	WRITE_int16(data.Order, buf)
	WRITE_string(data.Type, buf)
	WRITE_string(data.Owner, buf)
	WRITE_string(data.Collector, buf)
	WRITE_string(data.Short, buf)
	WRITE_bool(data.Deleted, buf)
}

func READ_MSG_PROJECT_tree_cache(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_cache {
	data := pool_MSG_PROJECT_tree_cache.Get().(*MSG_PROJECT_tree_cache)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_cache) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Root = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Parent = READ_int32(buf)
	Path_len := int(READ_int32(buf))
	for i := 0; i < Path_len; i++ {
		data.Path = append(data.Path, READ_int32(buf))
	}
	data.Grade = READ_int8(buf)
	data.Order = READ_int16(buf)
	data.Type = READ_string(buf)
	data.Owner = READ_string(buf)
	data.Collector = READ_string(buf)
	data.Short = READ_string(buf)
	data.Deleted = READ_bool(buf)

}

type MSG_PROJECT_tree_getParents struct {
	QueryID uint32
	ModuleID int32
}

var pool_MSG_PROJECT_tree_getParents = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_getParents{} }}

func GET_MSG_PROJECT_tree_getParents() *MSG_PROJECT_tree_getParents {
	return pool_MSG_PROJECT_tree_getParents.Get().(*MSG_PROJECT_tree_getParents)
}

func (data *MSG_PROJECT_tree_getParents) cmd() int32 {
	return CMD_MSG_PROJECT_tree_getParents
}

func (data *MSG_PROJECT_tree_getParents) Put() {
	data.QueryID = 0
	data.ModuleID = 0
	pool_MSG_PROJECT_tree_getParents.Put(data)
}
func (data *MSG_PROJECT_tree_getParents) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_getParents,buf)
	WRITE_MSG_PROJECT_tree_getParents(data, buf)
}

func WRITE_MSG_PROJECT_tree_getParents(data *MSG_PROJECT_tree_getParents, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.ModuleID, buf)
}

func READ_MSG_PROJECT_tree_getParents(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_getParents {
	data := pool_MSG_PROJECT_tree_getParents.Get().(*MSG_PROJECT_tree_getParents)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_getParents) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.ModuleID = READ_int32(buf)

}
func (data *MSG_PROJECT_tree_getParents) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_tree_getParents) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_tree_getParents_result struct {
	QueryResultID uint32
	List []*MSG_PROJECT_tree_cache
}

var pool_MSG_PROJECT_tree_getParents_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_getParents_result{} }}

func GET_MSG_PROJECT_tree_getParents_result() *MSG_PROJECT_tree_getParents_result {
	return pool_MSG_PROJECT_tree_getParents_result.Get().(*MSG_PROJECT_tree_getParents_result)
}

func (data *MSG_PROJECT_tree_getParents_result) cmd() int32 {
	return CMD_MSG_PROJECT_tree_getParents_result
}

func (data *MSG_PROJECT_tree_getParents_result) Put() {
	data.QueryResultID = 0
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_PROJECT_tree_getParents_result.Put(data)
}
func (data *MSG_PROJECT_tree_getParents_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_getParents_result,buf)
	WRITE_MSG_PROJECT_tree_getParents_result(data, buf)
}

func WRITE_MSG_PROJECT_tree_getParents_result(data *MSG_PROJECT_tree_getParents_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_MSG_PROJECT_tree_cache(v, buf)
	}
}

func READ_MSG_PROJECT_tree_getParents_result(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_getParents_result {
	data := pool_MSG_PROJECT_tree_getParents_result.Get().(*MSG_PROJECT_tree_getParents_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_getParents_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	for i := 0; i < List_len; i++ {
		data.List = append(data.List, READ_MSG_PROJECT_tree_cache(buf))
	}

}
func (data *MSG_PROJECT_tree_getParents_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_tree_getParents_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_branch_info struct {
	Id int32
	Product int32
	Name string
	Order int16
	Deleted bool
}

var pool_MSG_PROJECT_branch_info = sync.Pool{New: func() interface{} { return &MSG_PROJECT_branch_info{} }}

func GET_MSG_PROJECT_branch_info() *MSG_PROJECT_branch_info {
	return pool_MSG_PROJECT_branch_info.Get().(*MSG_PROJECT_branch_info)
}

func (data *MSG_PROJECT_branch_info) cmd() int32 {
	return CMD_MSG_PROJECT_branch_info
}

func (data *MSG_PROJECT_branch_info) Put() {
	data.Id = 0
	data.Product = 0
	data.Name = ``
	data.Order = 0
	data.Deleted = false
	pool_MSG_PROJECT_branch_info.Put(data)
}
func (data *MSG_PROJECT_branch_info) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_branch_info,buf)
	WRITE_MSG_PROJECT_branch_info(data, buf)
}

func WRITE_MSG_PROJECT_branch_info(data *MSG_PROJECT_branch_info, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Product, buf)
	WRITE_string(data.Name, buf)
	WRITE_int16(data.Order, buf)
	WRITE_bool(data.Deleted, buf)
}

func READ_MSG_PROJECT_branch_info(buf *libraries.MsgBuffer) *MSG_PROJECT_branch_info {
	data := pool_MSG_PROJECT_branch_info.Get().(*MSG_PROJECT_branch_info)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_branch_info) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Order = READ_int16(buf)
	data.Deleted = READ_bool(buf)

}

type MSG_PROJECT_tree_manageChild struct {
	QueryID uint32
	RootID int32
	ViewType string
	Modules []*MSG_PROJECT_tree_cache
	ParentModuleID int32
}

var pool_MSG_PROJECT_tree_manageChild = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_manageChild{} }}

func GET_MSG_PROJECT_tree_manageChild() *MSG_PROJECT_tree_manageChild {
	return pool_MSG_PROJECT_tree_manageChild.Get().(*MSG_PROJECT_tree_manageChild)
}

func (data *MSG_PROJECT_tree_manageChild) cmd() int32 {
	return CMD_MSG_PROJECT_tree_manageChild
}

func (data *MSG_PROJECT_tree_manageChild) Put() {
	data.QueryID = 0
	data.RootID = 0
	data.ViewType = ``
	for _,v := range data.Modules {
		v.Put()
	}
	data.Modules = data.Modules[:0]
	data.ParentModuleID = 0
	pool_MSG_PROJECT_tree_manageChild.Put(data)
}
func (data *MSG_PROJECT_tree_manageChild) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_manageChild,buf)
	WRITE_MSG_PROJECT_tree_manageChild(data, buf)
}

func WRITE_MSG_PROJECT_tree_manageChild(data *MSG_PROJECT_tree_manageChild, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.RootID, buf)
	WRITE_string(data.ViewType, buf)
	WRITE_int32(int32(len(data.Modules)), buf)
	for _, v := range data.Modules{
		WRITE_MSG_PROJECT_tree_cache(v, buf)
	}
	WRITE_int32(data.ParentModuleID, buf)
}

func READ_MSG_PROJECT_tree_manageChild(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_manageChild {
	data := pool_MSG_PROJECT_tree_manageChild.Get().(*MSG_PROJECT_tree_manageChild)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_manageChild) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.RootID = READ_int32(buf)
	data.ViewType = READ_string(buf)
	Modules_len := int(READ_int32(buf))
	for i := 0; i < Modules_len; i++ {
		data.Modules = append(data.Modules, READ_MSG_PROJECT_tree_cache(buf))
	}
	data.ParentModuleID = READ_int32(buf)

}
func (data *MSG_PROJECT_tree_manageChild) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_tree_manageChild) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_tree_manageChild_result struct {
	QueryResultID uint32
	Result ErrCode
	Name string
}

var pool_MSG_PROJECT_tree_manageChild_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_manageChild_result{} }}

func GET_MSG_PROJECT_tree_manageChild_result() *MSG_PROJECT_tree_manageChild_result {
	return pool_MSG_PROJECT_tree_manageChild_result.Get().(*MSG_PROJECT_tree_manageChild_result)
}

func (data *MSG_PROJECT_tree_manageChild_result) cmd() int32 {
	return CMD_MSG_PROJECT_tree_manageChild_result
}

func (data *MSG_PROJECT_tree_manageChild_result) Put() {
	data.QueryResultID = 0
	data.Result = 0
	data.Name = ``
	pool_MSG_PROJECT_tree_manageChild_result.Put(data)
}
func (data *MSG_PROJECT_tree_manageChild_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_manageChild_result,buf)
	WRITE_MSG_PROJECT_tree_manageChild_result(data, buf)
}

func WRITE_MSG_PROJECT_tree_manageChild_result(data *MSG_PROJECT_tree_manageChild_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_ErrCode(data.Result, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_PROJECT_tree_manageChild_result(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_manageChild_result {
	data := pool_MSG_PROJECT_tree_manageChild_result.Get().(*MSG_PROJECT_tree_manageChild_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_manageChild_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Result = READ_ErrCode(buf)
	data.Name = READ_string(buf)

}
func (data *MSG_PROJECT_tree_manageChild_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_tree_manageChild_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

