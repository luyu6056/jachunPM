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
	CMD_MSG_PROJECT_product_update = -1139812859
	CMD_MSG_PROJECT_product_getStories = 1365829125
	CMD_MSG_PROJECT_product_getStories_result = 528481285
	CMD_MSG_PROJECT_story = 1713425925
	CMD_MSG_PROJECT_tree_cache = -1861023995
	CMD_MSG_PROJECT_tree_getParents = -2065363451
	CMD_MSG_PROJECT_tree_getParents_result = 40064261
	CMD_MSG_PROJECT_branch_info = 894530565
	CMD_MSG_PROJECT_tree_manageChild = 727801093
	CMD_MSG_PROJECT_tree_manageChild_result = -2003846907
	CMD_MSG_PROJECT_product_getStoriesMapBySql = 2007179013
	CMD_MSG_PROJECT_product_getStoriesMapBySql_result = 1478139909
	CMD_MSG_PROJECT_tree_updateList = -1144133115
	CMD_MSG_PROJECT_tree_delete = 598613509
	CMD_MSG_PROJECT_productplan_getPairsForStory = -670540539
	CMD_MSG_PROJECT_productplan_getPairsForStory_result = 243028229
	CMD_MSG_PROJECT_productplan_getList = -1123948283
	CMD_MSG_PROJECT_productplan_getList_result = 181040389
	CMD_MSG_PROJECT_productplan_getLast = -1290618875
	CMD_MSG_PROJECT_productplan_getLast_result = 421344517
	CMD_MSG_PROJECT_product_editBranch = 2127377925
	CMD_MSG_PROJECT_product_deleteBranch = 1176677381
	CMD_MSG_PROJECT_product_deleteBranch_result = 1034823429
	CMD_MSG_PROJECT_productplan_getPairs = -1661534459
	CMD_MSG_PROJECT_productplan_getPairs_result = -160677883
	CMD_MSG_PROJECT_productplan_insertUpdate = 1228419589
	CMD_MSG_PROJECT_productplan_insertUpdate_result = -1827817723
	CMD_MSG_PROJECT_productplan_delete = -1670634235
	CMD_MSG_PROJECT_stroy_create = -426602747
	CMD_MSG_PROJECT_stroy_create_result = 1356456453
	CMD_MSG_PROJECT_story_batchGetStoryStage = 577900805
	CMD_MSG_PROJECT_story_batchGetStoryStage_result = 3743749
	CMD_MSG_PROJECT_story_getById = -357570299
	CMD_MSG_PROJECT_story_getById_result = 346215685
	CMD_MSG_PROJECT_project_getById = 779733509
	CMD_MSG_PROJECT_project_getById_result = 1679097605
	CMD_MSG_PROJECT_project_cache = 1829698565
	CMD_MSG_PROJECT_StoryStage = -281089787
	CMD_MSG_PROJECT_TASK = 439123205
	CMD_MSG_PROJECT_productplan_getById = 2045178629
	CMD_MSG_PROJECT_productplan_getById_result = -1195990011
	CMD_MSG_PROJECT_productplan = 1422058245
	CMD_MSG_PROJECT_build = 819345413
	CMD_MSG_PROJECT_build_getById = -1798962171
	CMD_MSG_PROJECT_build_getById_result = -242916347
	CMD_MSG_PROJECT_release = -1615845883
	CMD_MSG_PROJECT_release_getById = 1712503301
	CMD_MSG_PROJECT_release_getById_result = -792225787
	CMD_MSG_PROJECT_task_getPairs = -1587037691
	CMD_MSG_PROJECT_task_getPairs_result = -1441671419
	CMD_MSG_PROJECT_task_getListByWhereMap = -1248575483
	CMD_MSG_PROJECT_task_getListByWhereMap_result = -1620079099
	CMD_MSG_PROJECT_project_getBurn = -83610363
	CMD_MSG_PROJECT_project_getBurn_result = 1133272325
	CMD_MSG_PROJECT_project_Burn_info = 1223823365
	CMD_MSG_PROJECT_story_getPlanStories = 842619909
	CMD_MSG_PROJECT_story_getPlanStories_result = -1254780411
	CMD_MSG_PROJECT_project_linkStory = 1700618245
	CMD_MSG_PROJECT_branch_getByProducts = -815775483
	CMD_MSG_PROJECT_branch_getByProducts_result = 484676613
	CMD_MSG_PROJECT_project_create = -568022779
	CMD_MSG_PROJECT_project_create_result = 1720057093
	CMD_MSG_PROJECT_project_statRelatedData = 1630616325
	CMD_MSG_PROJECT_project_statRelatedData_result = 523943941
	CMD_MSG_PROJECT_story_getPairsByIds = 1500885253
	CMD_MSG_PROJECT_story_getPairsByIds_result = 762193157
	CMD_MSG_PROJECT_product_getPairsByIds = -1948818683
	CMD_MSG_PROJECT_product_getPairsByIds_result = 908792837
	CMD_MSG_PROJECT_project_getPairsByIds = -2052331515
	CMD_MSG_PROJECT_project_getPairsByIds_result = 244356101
	CMD_MSG_PROJECT_branch_getPairsByIds = 1861930501
	CMD_MSG_PROJECT_branch_getPairsByIds_result = -679761915
	CMD_MSG_PROJECT_tree_getPairsByIds = 1909900293
	CMD_MSG_PROJECT_tree_getPairsByIds_result = -950315259
	CMD_MSG_PROJECT_project_start = -1288650235
	CMD_MSG_PROJECT_project_putoff = -360695547
	CMD_MSG_PROJECT_project_suspend = -1715112955
	CMD_MSG_PROJECT_project_activate = -1563464187
	CMD_MSG_PROJECT_project_close = 1061688837
	CMD_MSG_PROJECT_project_delete = 1809855749
	CMD_MSG_PROJECT_project_getProjectTasks = -48489211
	CMD_MSG_PROJECT_project_getProjectTasks_result = 1808135429
	CMD_MSG_PROJECT_tree_getTaskTreeModules = 1624184325
	CMD_MSG_PROJECT_tree_getTaskTreeModules_result = 2106113029
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
	if List_len>cap(data.List){
		data.List= make([]HtmlKeyValueStr, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_HtmlKeyValueStr(buf)
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
	Branch []int32
	Plan []int32
	PO int32
	QD int32
	RD int32
	Acl string
	Whitelist []int32
	CreatedBy int32
	CreatedDate int64
	Order int32
	Deleted bool
	TimeStamp time.Time
	Branchs []*MSG_PROJECT_branch_info `db:"-"`
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
	data.Branch = data.Branch[:0]
	data.Plan = data.Plan[:0]
	data.PO = 0
	data.QD = 0
	data.RD = 0
	data.Acl = ``
	data.Whitelist = data.Whitelist[:0]
	data.CreatedBy = 0
	data.CreatedDate = 0
	data.Order = 0
	data.Deleted = false
	data.TimeStamp = time.Unix(0,0)
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
	WRITE_int32(int32(len(data.Branch)), buf)
	for _, v := range data.Branch{
		WRITE_int32(v, buf)
	}
	WRITE_int32(int32(len(data.Plan)), buf)
	for _, v := range data.Plan{
		WRITE_int32(v, buf)
	}
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
	WRITE_int64(data.TimeStamp.UnixNano(), buf)
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
	Branch_len := int(READ_int32(buf))
	if Branch_len>cap(data.Branch){
		data.Branch= make([]int32, Branch_len)
	}else{
		data.Branch = data.Branch[:Branch_len]
	}
	for i := 0; i < Branch_len; i++ {
		data.Branch[i] = READ_int32(buf)
	}
	Plan_len := int(READ_int32(buf))
	if Plan_len>cap(data.Plan){
		data.Plan= make([]int32, Plan_len)
	}else{
		data.Plan = data.Plan[:Plan_len]
	}
	for i := 0; i < Plan_len; i++ {
		data.Plan[i] = READ_int32(buf)
	}
	data.PO = READ_int32(buf)
	data.QD = READ_int32(buf)
	data.RD = READ_int32(buf)
	data.Acl = READ_string(buf)
	Whitelist_len := int(READ_int32(buf))
	if Whitelist_len>cap(data.Whitelist){
		data.Whitelist= make([]int32, Whitelist_len)
	}else{
		data.Whitelist = data.Whitelist[:Whitelist_len]
	}
	for i := 0; i < Whitelist_len; i++ {
		data.Whitelist[i] = READ_int32(buf)
	}
	data.CreatedBy = READ_int32(buf)
	data.CreatedDate = READ_int64(buf)
	data.Order = READ_int32(buf)
	data.Deleted = READ_bool(buf)
	data.TimeStamp = time.Unix(0, READ_int64(buf))
	Branchs_len := int(READ_int32(buf))
	if Branchs_len>cap(data.Branchs){
		data.Branchs= make([]*MSG_PROJECT_branch_info, Branchs_len)
	}else{
		data.Branchs = data.Branchs[:Branchs_len]
	}
	for i := 0; i < Branchs_len; i++ {
		data.Branchs[i] = READ_MSG_PROJECT_branch_info(buf)
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

type MSG_PROJECT_product_update struct {
	QueryID uint32
	Data *MSG_PROJECT_product_cache
}

var pool_MSG_PROJECT_product_update = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_update{} }}

func GET_MSG_PROJECT_product_update() *MSG_PROJECT_product_update {
	return pool_MSG_PROJECT_product_update.Get().(*MSG_PROJECT_product_update)
}

func (data *MSG_PROJECT_product_update) cmd() int32 {
	return CMD_MSG_PROJECT_product_update
}

func (data *MSG_PROJECT_product_update) Put() {
	data.QueryID = 0
	if data.Data != nil {
		data.Data.Put()
		data.Data = nil
	}
	pool_MSG_PROJECT_product_update.Put(data)
}
func (data *MSG_PROJECT_product_update) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_update,buf)
	WRITE_MSG_PROJECT_product_update(data, buf)
}

func WRITE_MSG_PROJECT_product_update(data *MSG_PROJECT_product_update, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	if data.Data == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_PROJECT_product_cache(data.Data, buf)
	}
}

func READ_MSG_PROJECT_product_update(buf *libraries.MsgBuffer) *MSG_PROJECT_product_update {
	data := pool_MSG_PROJECT_product_update.Get().(*MSG_PROJECT_product_update)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_update) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	Data_len := int(READ_int8(buf))
	if Data_len == 1 {
		data.Data = READ_MSG_PROJECT_product_cache(buf)
	}else{
		data.Data = nil
	}

}
func (data *MSG_PROJECT_product_update) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_product_update) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_product_getStories struct {
	QueryID uint32
	ProductID int32
	Branch int32
	BrowseType string
	ModuleID int32
	Sort string
	Uid int32
	Where map[string]interface{}
	Page int
	PerPage int
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
	data.Uid = 0
	data.Where = nil
	data.Page = 0
	data.PerPage = 0
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
	WRITE_int32(data.Uid, buf)
	WRITE_map(data.Where,buf)
	WRITE_int(data.Page, buf)
	WRITE_int(data.PerPage, buf)
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
	data.Uid = READ_int32(buf)
	READ_map(&data.Where,buf)
	data.Page = READ_int(buf)
	data.PerPage = READ_int(buf)
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
	if List_len>cap(data.List){
		data.List= make([]*MSG_PROJECT_story, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_PROJECT_story(buf)
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
	Plan int32
	Source string
	SourceNote string
	FromBug int32
	Title string
	Keywords string
	Pri int8
	Estimate float64
	Status string
	Stage string
	Mailto []int32
	OpenedBy int32
	OpenedDate time.Time
	AssignedTo int32
	AssignedDate time.Time
	LastEditedBy int32
	LastEditedDate time.Time
	ReviewedBy int32
	ReviewedDate time.Time
	ClosedBy int32
	ClosedDate time.Time
	ClosedReason string
	ToBug int32
	ChildStories []int32
	LinkStories []int32
	DuplicateStory int32
	Deleted bool
	Version int16
	Color string
	PlanTitle string `db:-`
	Spec string `db:-`
	Verify string `db:-`
	Stages []*MSG_PROJECT_StoryStage `db:-`
	ExtraStories []*MSG_PROJECT_story `db:-`
	Projects []int32 `db:-`
	Tasks []*MSG_PROJECT_TASK `db:-`
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
	data.Plan = 0
	data.Source = ``
	data.SourceNote = ``
	data.FromBug = 0
	data.Title = ``
	data.Keywords = ``
	data.Pri = 0
	data.Estimate = 0
	data.Status = ``
	data.Stage = ``
	data.Mailto = data.Mailto[:0]
	data.OpenedBy = 0
	data.OpenedDate = time.Unix(0,0)
	data.AssignedTo = 0
	data.AssignedDate = time.Unix(0,0)
	data.LastEditedBy = 0
	data.LastEditedDate = time.Unix(0,0)
	data.ReviewedBy = 0
	data.ReviewedDate = time.Unix(0,0)
	data.ClosedBy = 0
	data.ClosedDate = time.Unix(0,0)
	data.ClosedReason = ``
	data.ToBug = 0
	data.ChildStories = data.ChildStories[:0]
	data.LinkStories = data.LinkStories[:0]
	data.DuplicateStory = 0
	data.Deleted = false
	data.Version = 0
	data.Color = ``
	data.PlanTitle = ``
	data.Spec = ``
	data.Verify = ``
	for _,v := range data.Stages {
		v.Put()
	}
	data.Stages = data.Stages[:0]
	for _,v := range data.ExtraStories {
		v.Put()
	}
	data.ExtraStories = data.ExtraStories[:0]
	data.Projects = data.Projects[:0]
	for _,v := range data.Tasks {
		v.Put()
	}
	data.Tasks = data.Tasks[:0]
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
	WRITE_int32(data.Plan, buf)
	WRITE_string(data.Source, buf)
	WRITE_string(data.SourceNote, buf)
	WRITE_int32(data.FromBug, buf)
	WRITE_string(data.Title, buf)
	WRITE_string(data.Keywords, buf)
	WRITE_int8(data.Pri, buf)
	WRITE_float64(data.Estimate, buf)
	WRITE_string(data.Status, buf)
	WRITE_string(data.Stage, buf)
	WRITE_int32(int32(len(data.Mailto)), buf)
	for _, v := range data.Mailto{
		WRITE_int32(v, buf)
	}
	WRITE_int32(data.OpenedBy, buf)
	WRITE_int64(data.OpenedDate.UnixNano(), buf)
	WRITE_int32(data.AssignedTo, buf)
	WRITE_int64(data.AssignedDate.UnixNano(), buf)
	WRITE_int32(data.LastEditedBy, buf)
	WRITE_int64(data.LastEditedDate.UnixNano(), buf)
	WRITE_int32(data.ReviewedBy, buf)
	WRITE_int64(data.ReviewedDate.UnixNano(), buf)
	WRITE_int32(data.ClosedBy, buf)
	WRITE_int64(data.ClosedDate.UnixNano(), buf)
	WRITE_string(data.ClosedReason, buf)
	WRITE_int32(data.ToBug, buf)
	WRITE_int32(int32(len(data.ChildStories)), buf)
	for _, v := range data.ChildStories{
		WRITE_int32(v, buf)
	}
	WRITE_int32(int32(len(data.LinkStories)), buf)
	for _, v := range data.LinkStories{
		WRITE_int32(v, buf)
	}
	WRITE_int32(data.DuplicateStory, buf)
	WRITE_bool(data.Deleted, buf)
	WRITE_int16(data.Version, buf)
	WRITE_string(data.Color, buf)
	WRITE_string(data.PlanTitle, buf)
	WRITE_string(data.Spec, buf)
	WRITE_string(data.Verify, buf)
	WRITE_int32(int32(len(data.Stages)), buf)
	for _, v := range data.Stages{
		WRITE_MSG_PROJECT_StoryStage(v, buf)
	}
	WRITE_int32(int32(len(data.ExtraStories)), buf)
	for _, v := range data.ExtraStories{
		WRITE_MSG_PROJECT_story(v, buf)
	}
	WRITE_int32(int32(len(data.Projects)), buf)
	for _, v := range data.Projects{
		WRITE_int32(v, buf)
	}
	WRITE_int32(int32(len(data.Tasks)), buf)
	for _, v := range data.Tasks{
		WRITE_MSG_PROJECT_TASK(v, buf)
	}
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
	data.Plan = READ_int32(buf)
	data.Source = READ_string(buf)
	data.SourceNote = READ_string(buf)
	data.FromBug = READ_int32(buf)
	data.Title = READ_string(buf)
	data.Keywords = READ_string(buf)
	data.Pri = READ_int8(buf)
	data.Estimate = READ_float64(buf)
	data.Status = READ_string(buf)
	data.Stage = READ_string(buf)
	Mailto_len := int(READ_int32(buf))
	if Mailto_len>cap(data.Mailto){
		data.Mailto= make([]int32, Mailto_len)
	}else{
		data.Mailto = data.Mailto[:Mailto_len]
	}
	for i := 0; i < Mailto_len; i++ {
		data.Mailto[i] = READ_int32(buf)
	}
	data.OpenedBy = READ_int32(buf)
	data.OpenedDate = time.Unix(0, READ_int64(buf))
	data.AssignedTo = READ_int32(buf)
	data.AssignedDate = time.Unix(0, READ_int64(buf))
	data.LastEditedBy = READ_int32(buf)
	data.LastEditedDate = time.Unix(0, READ_int64(buf))
	data.ReviewedBy = READ_int32(buf)
	data.ReviewedDate = time.Unix(0, READ_int64(buf))
	data.ClosedBy = READ_int32(buf)
	data.ClosedDate = time.Unix(0, READ_int64(buf))
	data.ClosedReason = READ_string(buf)
	data.ToBug = READ_int32(buf)
	ChildStories_len := int(READ_int32(buf))
	if ChildStories_len>cap(data.ChildStories){
		data.ChildStories= make([]int32, ChildStories_len)
	}else{
		data.ChildStories = data.ChildStories[:ChildStories_len]
	}
	for i := 0; i < ChildStories_len; i++ {
		data.ChildStories[i] = READ_int32(buf)
	}
	LinkStories_len := int(READ_int32(buf))
	if LinkStories_len>cap(data.LinkStories){
		data.LinkStories= make([]int32, LinkStories_len)
	}else{
		data.LinkStories = data.LinkStories[:LinkStories_len]
	}
	for i := 0; i < LinkStories_len; i++ {
		data.LinkStories[i] = READ_int32(buf)
	}
	data.DuplicateStory = READ_int32(buf)
	data.Deleted = READ_bool(buf)
	data.Version = READ_int16(buf)
	data.Color = READ_string(buf)
	data.PlanTitle = READ_string(buf)
	data.Spec = READ_string(buf)
	data.Verify = READ_string(buf)
	Stages_len := int(READ_int32(buf))
	if Stages_len>cap(data.Stages){
		data.Stages= make([]*MSG_PROJECT_StoryStage, Stages_len)
	}else{
		data.Stages = data.Stages[:Stages_len]
	}
	for i := 0; i < Stages_len; i++ {
		data.Stages[i] = READ_MSG_PROJECT_StoryStage(buf)
	}
	ExtraStories_len := int(READ_int32(buf))
	if ExtraStories_len>cap(data.ExtraStories){
		data.ExtraStories= make([]*MSG_PROJECT_story, ExtraStories_len)
	}else{
		data.ExtraStories = data.ExtraStories[:ExtraStories_len]
	}
	for i := 0; i < ExtraStories_len; i++ {
		data.ExtraStories[i] = READ_MSG_PROJECT_story(buf)
	}
	Projects_len := int(READ_int32(buf))
	if Projects_len>cap(data.Projects){
		data.Projects= make([]int32, Projects_len)
	}else{
		data.Projects = data.Projects[:Projects_len]
	}
	for i := 0; i < Projects_len; i++ {
		data.Projects[i] = READ_int32(buf)
	}
	Tasks_len := int(READ_int32(buf))
	if Tasks_len>cap(data.Tasks){
		data.Tasks= make([]*MSG_PROJECT_TASK, Tasks_len)
	}else{
		data.Tasks = data.Tasks[:Tasks_len]
	}
	for i := 0; i < Tasks_len; i++ {
		data.Tasks[i] = READ_MSG_PROJECT_TASK(buf)
	}

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
	OwnerID int32
	Collector string
	Short string
	Deleted bool
	TimeStamp time.Time
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
	data.OwnerID = 0
	data.Collector = ``
	data.Short = ``
	data.Deleted = false
	data.TimeStamp = time.Unix(0,0)
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
	WRITE_int32(data.OwnerID, buf)
	WRITE_string(data.Collector, buf)
	WRITE_string(data.Short, buf)
	WRITE_bool(data.Deleted, buf)
	WRITE_int64(data.TimeStamp.UnixNano(), buf)
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
	if Path_len>cap(data.Path){
		data.Path= make([]int32, Path_len)
	}else{
		data.Path = data.Path[:Path_len]
	}
	for i := 0; i < Path_len; i++ {
		data.Path[i] = READ_int32(buf)
	}
	data.Grade = READ_int8(buf)
	data.Order = READ_int16(buf)
	data.Type = READ_string(buf)
	data.Owner = READ_string(buf)
	data.OwnerID = READ_int32(buf)
	data.Collector = READ_string(buf)
	data.Short = READ_string(buf)
	data.Deleted = READ_bool(buf)
	data.TimeStamp = time.Unix(0, READ_int64(buf))

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
	if List_len>cap(data.List){
		data.List= make([]*MSG_PROJECT_tree_cache, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_PROJECT_tree_cache(buf)
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
	if Modules_len>cap(data.Modules){
		data.Modules= make([]*MSG_PROJECT_tree_cache, Modules_len)
	}else{
		data.Modules = data.Modules[:Modules_len]
	}
	for i := 0; i < Modules_len; i++ {
		data.Modules[i] = READ_MSG_PROJECT_tree_cache(buf)
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

type MSG_PROJECT_product_getStoriesMapBySql struct {
	QueryID uint32
	Field string
	Where map[string]interface{}
	Order string
	Group string
	Page int
	PerPage int
	Total int
}

var pool_MSG_PROJECT_product_getStoriesMapBySql = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_getStoriesMapBySql{} }}

func GET_MSG_PROJECT_product_getStoriesMapBySql() *MSG_PROJECT_product_getStoriesMapBySql {
	return pool_MSG_PROJECT_product_getStoriesMapBySql.Get().(*MSG_PROJECT_product_getStoriesMapBySql)
}

func (data *MSG_PROJECT_product_getStoriesMapBySql) cmd() int32 {
	return CMD_MSG_PROJECT_product_getStoriesMapBySql
}

func (data *MSG_PROJECT_product_getStoriesMapBySql) Put() {
	data.QueryID = 0
	data.Field = ``
	data.Where = nil
	data.Order = ``
	data.Group = ``
	data.Page = 0
	data.PerPage = 0
	data.Total = 0
	pool_MSG_PROJECT_product_getStoriesMapBySql.Put(data)
}
func (data *MSG_PROJECT_product_getStoriesMapBySql) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_getStoriesMapBySql,buf)
	WRITE_MSG_PROJECT_product_getStoriesMapBySql(data, buf)
}

func WRITE_MSG_PROJECT_product_getStoriesMapBySql(data *MSG_PROJECT_product_getStoriesMapBySql, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_string(data.Field, buf)
	WRITE_map(data.Where,buf)
	WRITE_string(data.Order, buf)
	WRITE_string(data.Group, buf)
	WRITE_int(data.Page, buf)
	WRITE_int(data.PerPage, buf)
	WRITE_int(data.Total, buf)
}

func READ_MSG_PROJECT_product_getStoriesMapBySql(buf *libraries.MsgBuffer) *MSG_PROJECT_product_getStoriesMapBySql {
	data := pool_MSG_PROJECT_product_getStoriesMapBySql.Get().(*MSG_PROJECT_product_getStoriesMapBySql)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_getStoriesMapBySql) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Field = READ_string(buf)
	READ_map(&data.Where,buf)
	data.Order = READ_string(buf)
	data.Group = READ_string(buf)
	data.Page = READ_int(buf)
	data.PerPage = READ_int(buf)
	data.Total = READ_int(buf)

}
func (data *MSG_PROJECT_product_getStoriesMapBySql) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_product_getStoriesMapBySql) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_product_getStoriesMapBySql_result struct {
	QueryResultID uint32
	List []map[string]string
	Total int
}

var pool_MSG_PROJECT_product_getStoriesMapBySql_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_getStoriesMapBySql_result{} }}

func GET_MSG_PROJECT_product_getStoriesMapBySql_result() *MSG_PROJECT_product_getStoriesMapBySql_result {
	return pool_MSG_PROJECT_product_getStoriesMapBySql_result.Get().(*MSG_PROJECT_product_getStoriesMapBySql_result)
}

func (data *MSG_PROJECT_product_getStoriesMapBySql_result) cmd() int32 {
	return CMD_MSG_PROJECT_product_getStoriesMapBySql_result
}

func (data *MSG_PROJECT_product_getStoriesMapBySql_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	data.Total = 0
	pool_MSG_PROJECT_product_getStoriesMapBySql_result.Put(data)
}
func (data *MSG_PROJECT_product_getStoriesMapBySql_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_getStoriesMapBySql_result,buf)
	WRITE_MSG_PROJECT_product_getStoriesMapBySql_result(data, buf)
}

func WRITE_MSG_PROJECT_product_getStoriesMapBySql_result(data *MSG_PROJECT_product_getStoriesMapBySql_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_any(v, buf)
	}
	WRITE_int(data.Total, buf)
}

func READ_MSG_PROJECT_product_getStoriesMapBySql_result(buf *libraries.MsgBuffer) *MSG_PROJECT_product_getStoriesMapBySql_result {
	data := pool_MSG_PROJECT_product_getStoriesMapBySql_result.Get().(*MSG_PROJECT_product_getStoriesMapBySql_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_getStoriesMapBySql_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	if List_len>cap(data.List){
		data.List= make([]map[string]string, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		READ_any(&data.List[i], buf)
	}
	data.Total = READ_int(buf)

}
func (data *MSG_PROJECT_product_getStoriesMapBySql_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_product_getStoriesMapBySql_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_tree_updateList struct {
	QueryID uint32
	Modules []*MSG_PROJECT_tree_cache
}

var pool_MSG_PROJECT_tree_updateList = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_updateList{} }}

func GET_MSG_PROJECT_tree_updateList() *MSG_PROJECT_tree_updateList {
	return pool_MSG_PROJECT_tree_updateList.Get().(*MSG_PROJECT_tree_updateList)
}

func (data *MSG_PROJECT_tree_updateList) cmd() int32 {
	return CMD_MSG_PROJECT_tree_updateList
}

func (data *MSG_PROJECT_tree_updateList) Put() {
	data.QueryID = 0
	for _,v := range data.Modules {
		v.Put()
	}
	data.Modules = data.Modules[:0]
	pool_MSG_PROJECT_tree_updateList.Put(data)
}
func (data *MSG_PROJECT_tree_updateList) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_updateList,buf)
	WRITE_MSG_PROJECT_tree_updateList(data, buf)
}

func WRITE_MSG_PROJECT_tree_updateList(data *MSG_PROJECT_tree_updateList, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.Modules)), buf)
	for _, v := range data.Modules{
		WRITE_MSG_PROJECT_tree_cache(v, buf)
	}
}

func READ_MSG_PROJECT_tree_updateList(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_updateList {
	data := pool_MSG_PROJECT_tree_updateList.Get().(*MSG_PROJECT_tree_updateList)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_updateList) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	Modules_len := int(READ_int32(buf))
	if Modules_len>cap(data.Modules){
		data.Modules= make([]*MSG_PROJECT_tree_cache, Modules_len)
	}else{
		data.Modules = data.Modules[:Modules_len]
	}
	for i := 0; i < Modules_len; i++ {
		data.Modules[i] = READ_MSG_PROJECT_tree_cache(buf)
	}

}
func (data *MSG_PROJECT_tree_updateList) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_tree_updateList) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_tree_delete struct {
	QueryID uint32
	Ids []int32
}

var pool_MSG_PROJECT_tree_delete = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_delete{} }}

func GET_MSG_PROJECT_tree_delete() *MSG_PROJECT_tree_delete {
	return pool_MSG_PROJECT_tree_delete.Get().(*MSG_PROJECT_tree_delete)
}

func (data *MSG_PROJECT_tree_delete) cmd() int32 {
	return CMD_MSG_PROJECT_tree_delete
}

func (data *MSG_PROJECT_tree_delete) Put() {
	data.QueryID = 0
	data.Ids = data.Ids[:0]
	pool_MSG_PROJECT_tree_delete.Put(data)
}
func (data *MSG_PROJECT_tree_delete) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_delete,buf)
	WRITE_MSG_PROJECT_tree_delete(data, buf)
}

func WRITE_MSG_PROJECT_tree_delete(data *MSG_PROJECT_tree_delete, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_PROJECT_tree_delete(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_delete {
	data := pool_MSG_PROJECT_tree_delete.Get().(*MSG_PROJECT_tree_delete)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_delete) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	Ids_len := int(READ_int32(buf))
	if Ids_len>cap(data.Ids){
		data.Ids= make([]int32, Ids_len)
	}else{
		data.Ids = data.Ids[:Ids_len]
	}
	for i := 0; i < Ids_len; i++ {
		data.Ids[i] = READ_int32(buf)
	}

}
func (data *MSG_PROJECT_tree_delete) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_tree_delete) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_productplan_getPairsForStory struct {
	QueryID uint32
	Product int32
	Branch int32
}

var pool_MSG_PROJECT_productplan_getPairsForStory = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_getPairsForStory{} }}

func GET_MSG_PROJECT_productplan_getPairsForStory() *MSG_PROJECT_productplan_getPairsForStory {
	return pool_MSG_PROJECT_productplan_getPairsForStory.Get().(*MSG_PROJECT_productplan_getPairsForStory)
}

func (data *MSG_PROJECT_productplan_getPairsForStory) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_getPairsForStory
}

func (data *MSG_PROJECT_productplan_getPairsForStory) Put() {
	data.QueryID = 0
	data.Product = 0
	data.Branch = 0
	pool_MSG_PROJECT_productplan_getPairsForStory.Put(data)
}
func (data *MSG_PROJECT_productplan_getPairsForStory) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_getPairsForStory,buf)
	WRITE_MSG_PROJECT_productplan_getPairsForStory(data, buf)
}

func WRITE_MSG_PROJECT_productplan_getPairsForStory(data *MSG_PROJECT_productplan_getPairsForStory, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Branch, buf)
}

func READ_MSG_PROJECT_productplan_getPairsForStory(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_getPairsForStory {
	data := pool_MSG_PROJECT_productplan_getPairsForStory.Get().(*MSG_PROJECT_productplan_getPairsForStory)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_getPairsForStory) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Product = READ_int32(buf)
	data.Branch = READ_int32(buf)

}
func (data *MSG_PROJECT_productplan_getPairsForStory) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_productplan_getPairsForStory) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_productplan_getPairsForStory_result struct {
	QueryResultID uint32
	List []HtmlKeyValueStr
}

var pool_MSG_PROJECT_productplan_getPairsForStory_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_getPairsForStory_result{} }}

func GET_MSG_PROJECT_productplan_getPairsForStory_result() *MSG_PROJECT_productplan_getPairsForStory_result {
	return pool_MSG_PROJECT_productplan_getPairsForStory_result.Get().(*MSG_PROJECT_productplan_getPairsForStory_result)
}

func (data *MSG_PROJECT_productplan_getPairsForStory_result) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_getPairsForStory_result
}

func (data *MSG_PROJECT_productplan_getPairsForStory_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	pool_MSG_PROJECT_productplan_getPairsForStory_result.Put(data)
}
func (data *MSG_PROJECT_productplan_getPairsForStory_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_getPairsForStory_result,buf)
	WRITE_MSG_PROJECT_productplan_getPairsForStory_result(data, buf)
}

func WRITE_MSG_PROJECT_productplan_getPairsForStory_result(data *MSG_PROJECT_productplan_getPairsForStory_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_PROJECT_productplan_getPairsForStory_result(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_getPairsForStory_result {
	data := pool_MSG_PROJECT_productplan_getPairsForStory_result.Get().(*MSG_PROJECT_productplan_getPairsForStory_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_getPairsForStory_result) read(buf *libraries.MsgBuffer) {
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
func (data *MSG_PROJECT_productplan_getPairsForStory_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_productplan_getPairsForStory_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_productplan_getList struct {
	QueryID uint32
	Ids []int32
	ProductID int32
	Branch int32
	BrowseType string
	Order string
	Page int
	PerPage int
	Total int
}

var pool_MSG_PROJECT_productplan_getList = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_getList{} }}

func GET_MSG_PROJECT_productplan_getList() *MSG_PROJECT_productplan_getList {
	return pool_MSG_PROJECT_productplan_getList.Get().(*MSG_PROJECT_productplan_getList)
}

func (data *MSG_PROJECT_productplan_getList) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_getList
}

func (data *MSG_PROJECT_productplan_getList) Put() {
	data.QueryID = 0
	data.Ids = data.Ids[:0]
	data.ProductID = 0
	data.Branch = 0
	data.BrowseType = ``
	data.Order = ``
	data.Page = 0
	data.PerPage = 0
	data.Total = 0
	pool_MSG_PROJECT_productplan_getList.Put(data)
}
func (data *MSG_PROJECT_productplan_getList) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_getList,buf)
	WRITE_MSG_PROJECT_productplan_getList(data, buf)
}

func WRITE_MSG_PROJECT_productplan_getList(data *MSG_PROJECT_productplan_getList, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int32(v, buf)
	}
	WRITE_int32(data.ProductID, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_string(data.BrowseType, buf)
	WRITE_string(data.Order, buf)
	WRITE_int(data.Page, buf)
	WRITE_int(data.PerPage, buf)
	WRITE_int(data.Total, buf)
}

func READ_MSG_PROJECT_productplan_getList(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_getList {
	data := pool_MSG_PROJECT_productplan_getList.Get().(*MSG_PROJECT_productplan_getList)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_getList) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	Ids_len := int(READ_int32(buf))
	if Ids_len>cap(data.Ids){
		data.Ids= make([]int32, Ids_len)
	}else{
		data.Ids = data.Ids[:Ids_len]
	}
	for i := 0; i < Ids_len; i++ {
		data.Ids[i] = READ_int32(buf)
	}
	data.ProductID = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.BrowseType = READ_string(buf)
	data.Order = READ_string(buf)
	data.Page = READ_int(buf)
	data.PerPage = READ_int(buf)
	data.Total = READ_int(buf)

}
func (data *MSG_PROJECT_productplan_getList) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_productplan_getList) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_productplan_getList_result struct {
	QueryResultID uint32
	List []map[string]string
	Total int
}

var pool_MSG_PROJECT_productplan_getList_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_getList_result{} }}

func GET_MSG_PROJECT_productplan_getList_result() *MSG_PROJECT_productplan_getList_result {
	return pool_MSG_PROJECT_productplan_getList_result.Get().(*MSG_PROJECT_productplan_getList_result)
}

func (data *MSG_PROJECT_productplan_getList_result) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_getList_result
}

func (data *MSG_PROJECT_productplan_getList_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	data.Total = 0
	pool_MSG_PROJECT_productplan_getList_result.Put(data)
}
func (data *MSG_PROJECT_productplan_getList_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_getList_result,buf)
	WRITE_MSG_PROJECT_productplan_getList_result(data, buf)
}

func WRITE_MSG_PROJECT_productplan_getList_result(data *MSG_PROJECT_productplan_getList_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_any(v, buf)
	}
	WRITE_int(data.Total, buf)
}

func READ_MSG_PROJECT_productplan_getList_result(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_getList_result {
	data := pool_MSG_PROJECT_productplan_getList_result.Get().(*MSG_PROJECT_productplan_getList_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_getList_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	if List_len>cap(data.List){
		data.List= make([]map[string]string, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		READ_any(&data.List[i], buf)
	}
	data.Total = READ_int(buf)

}
func (data *MSG_PROJECT_productplan_getList_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_productplan_getList_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_productplan_getLast struct {
	QueryID uint32
	ProductId int32
	Branch int32
}

var pool_MSG_PROJECT_productplan_getLast = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_getLast{} }}

func GET_MSG_PROJECT_productplan_getLast() *MSG_PROJECT_productplan_getLast {
	return pool_MSG_PROJECT_productplan_getLast.Get().(*MSG_PROJECT_productplan_getLast)
}

func (data *MSG_PROJECT_productplan_getLast) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_getLast
}

func (data *MSG_PROJECT_productplan_getLast) Put() {
	data.QueryID = 0
	data.ProductId = 0
	data.Branch = 0
	pool_MSG_PROJECT_productplan_getLast.Put(data)
}
func (data *MSG_PROJECT_productplan_getLast) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_getLast,buf)
	WRITE_MSG_PROJECT_productplan_getLast(data, buf)
}

func WRITE_MSG_PROJECT_productplan_getLast(data *MSG_PROJECT_productplan_getLast, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.ProductId, buf)
	WRITE_int32(data.Branch, buf)
}

func READ_MSG_PROJECT_productplan_getLast(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_getLast {
	data := pool_MSG_PROJECT_productplan_getLast.Get().(*MSG_PROJECT_productplan_getLast)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_getLast) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.ProductId = READ_int32(buf)
	data.Branch = READ_int32(buf)

}
func (data *MSG_PROJECT_productplan_getLast) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_productplan_getLast) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_productplan_getLast_result struct {
	QueryResultID uint32
	Result map[string]string
}

var pool_MSG_PROJECT_productplan_getLast_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_getLast_result{} }}

func GET_MSG_PROJECT_productplan_getLast_result() *MSG_PROJECT_productplan_getLast_result {
	return pool_MSG_PROJECT_productplan_getLast_result.Get().(*MSG_PROJECT_productplan_getLast_result)
}

func (data *MSG_PROJECT_productplan_getLast_result) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_getLast_result
}

func (data *MSG_PROJECT_productplan_getLast_result) Put() {
	data.QueryResultID = 0
	data.Result = nil
	pool_MSG_PROJECT_productplan_getLast_result.Put(data)
}
func (data *MSG_PROJECT_productplan_getLast_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_getLast_result,buf)
	WRITE_MSG_PROJECT_productplan_getLast_result(data, buf)
}

func WRITE_MSG_PROJECT_productplan_getLast_result(data *MSG_PROJECT_productplan_getLast_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_map(data.Result,buf)
}

func READ_MSG_PROJECT_productplan_getLast_result(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_getLast_result {
	data := pool_MSG_PROJECT_productplan_getLast_result.Get().(*MSG_PROJECT_productplan_getLast_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_getLast_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	READ_map(&data.Result,buf)

}
func (data *MSG_PROJECT_productplan_getLast_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_productplan_getLast_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_product_editBranch struct {
	QueryID uint32
	ProductID int32
	Branchs []*MSG_PROJECT_branch_info
}

var pool_MSG_PROJECT_product_editBranch = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_editBranch{} }}

func GET_MSG_PROJECT_product_editBranch() *MSG_PROJECT_product_editBranch {
	return pool_MSG_PROJECT_product_editBranch.Get().(*MSG_PROJECT_product_editBranch)
}

func (data *MSG_PROJECT_product_editBranch) cmd() int32 {
	return CMD_MSG_PROJECT_product_editBranch
}

func (data *MSG_PROJECT_product_editBranch) Put() {
	data.QueryID = 0
	data.ProductID = 0
	for _,v := range data.Branchs {
		v.Put()
	}
	data.Branchs = data.Branchs[:0]
	pool_MSG_PROJECT_product_editBranch.Put(data)
}
func (data *MSG_PROJECT_product_editBranch) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_editBranch,buf)
	WRITE_MSG_PROJECT_product_editBranch(data, buf)
}

func WRITE_MSG_PROJECT_product_editBranch(data *MSG_PROJECT_product_editBranch, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.ProductID, buf)
	WRITE_int32(int32(len(data.Branchs)), buf)
	for _, v := range data.Branchs{
		WRITE_MSG_PROJECT_branch_info(v, buf)
	}
}

func READ_MSG_PROJECT_product_editBranch(buf *libraries.MsgBuffer) *MSG_PROJECT_product_editBranch {
	data := pool_MSG_PROJECT_product_editBranch.Get().(*MSG_PROJECT_product_editBranch)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_editBranch) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.ProductID = READ_int32(buf)
	Branchs_len := int(READ_int32(buf))
	if Branchs_len>cap(data.Branchs){
		data.Branchs= make([]*MSG_PROJECT_branch_info, Branchs_len)
	}else{
		data.Branchs = data.Branchs[:Branchs_len]
	}
	for i := 0; i < Branchs_len; i++ {
		data.Branchs[i] = READ_MSG_PROJECT_branch_info(buf)
	}

}
func (data *MSG_PROJECT_product_editBranch) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_product_editBranch) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_product_deleteBranch struct {
	QueryID uint32
	ProductID int32
	BranchID int32
}

var pool_MSG_PROJECT_product_deleteBranch = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_deleteBranch{} }}

func GET_MSG_PROJECT_product_deleteBranch() *MSG_PROJECT_product_deleteBranch {
	return pool_MSG_PROJECT_product_deleteBranch.Get().(*MSG_PROJECT_product_deleteBranch)
}

func (data *MSG_PROJECT_product_deleteBranch) cmd() int32 {
	return CMD_MSG_PROJECT_product_deleteBranch
}

func (data *MSG_PROJECT_product_deleteBranch) Put() {
	data.QueryID = 0
	data.ProductID = 0
	data.BranchID = 0
	pool_MSG_PROJECT_product_deleteBranch.Put(data)
}
func (data *MSG_PROJECT_product_deleteBranch) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_deleteBranch,buf)
	WRITE_MSG_PROJECT_product_deleteBranch(data, buf)
}

func WRITE_MSG_PROJECT_product_deleteBranch(data *MSG_PROJECT_product_deleteBranch, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.ProductID, buf)
	WRITE_int32(data.BranchID, buf)
}

func READ_MSG_PROJECT_product_deleteBranch(buf *libraries.MsgBuffer) *MSG_PROJECT_product_deleteBranch {
	data := pool_MSG_PROJECT_product_deleteBranch.Get().(*MSG_PROJECT_product_deleteBranch)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_deleteBranch) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.ProductID = READ_int32(buf)
	data.BranchID = READ_int32(buf)

}
func (data *MSG_PROJECT_product_deleteBranch) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_product_deleteBranch) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_product_deleteBranch_result struct {
	QueryResultID uint32
	Result ErrCode
}

var pool_MSG_PROJECT_product_deleteBranch_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_deleteBranch_result{} }}

func GET_MSG_PROJECT_product_deleteBranch_result() *MSG_PROJECT_product_deleteBranch_result {
	return pool_MSG_PROJECT_product_deleteBranch_result.Get().(*MSG_PROJECT_product_deleteBranch_result)
}

func (data *MSG_PROJECT_product_deleteBranch_result) cmd() int32 {
	return CMD_MSG_PROJECT_product_deleteBranch_result
}

func (data *MSG_PROJECT_product_deleteBranch_result) Put() {
	data.QueryResultID = 0
	data.Result = 0
	pool_MSG_PROJECT_product_deleteBranch_result.Put(data)
}
func (data *MSG_PROJECT_product_deleteBranch_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_deleteBranch_result,buf)
	WRITE_MSG_PROJECT_product_deleteBranch_result(data, buf)
}

func WRITE_MSG_PROJECT_product_deleteBranch_result(data *MSG_PROJECT_product_deleteBranch_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_ErrCode(data.Result, buf)
}

func READ_MSG_PROJECT_product_deleteBranch_result(buf *libraries.MsgBuffer) *MSG_PROJECT_product_deleteBranch_result {
	data := pool_MSG_PROJECT_product_deleteBranch_result.Get().(*MSG_PROJECT_product_deleteBranch_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_deleteBranch_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Result = READ_ErrCode(buf)

}
func (data *MSG_PROJECT_product_deleteBranch_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_product_deleteBranch_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_productplan_getPairs struct {
	QueryID uint32
	ProductID int32
	BranchID int32
	Expired string
	Ids []int32
}

var pool_MSG_PROJECT_productplan_getPairs = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_getPairs{} }}

func GET_MSG_PROJECT_productplan_getPairs() *MSG_PROJECT_productplan_getPairs {
	return pool_MSG_PROJECT_productplan_getPairs.Get().(*MSG_PROJECT_productplan_getPairs)
}

func (data *MSG_PROJECT_productplan_getPairs) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_getPairs
}

func (data *MSG_PROJECT_productplan_getPairs) Put() {
	data.QueryID = 0
	data.ProductID = 0
	data.BranchID = 0
	data.Expired = ``
	data.Ids = data.Ids[:0]
	pool_MSG_PROJECT_productplan_getPairs.Put(data)
}
func (data *MSG_PROJECT_productplan_getPairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_getPairs,buf)
	WRITE_MSG_PROJECT_productplan_getPairs(data, buf)
}

func WRITE_MSG_PROJECT_productplan_getPairs(data *MSG_PROJECT_productplan_getPairs, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.ProductID, buf)
	WRITE_int32(data.BranchID, buf)
	WRITE_string(data.Expired, buf)
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_PROJECT_productplan_getPairs(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_getPairs {
	data := pool_MSG_PROJECT_productplan_getPairs.Get().(*MSG_PROJECT_productplan_getPairs)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_getPairs) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.ProductID = READ_int32(buf)
	data.BranchID = READ_int32(buf)
	data.Expired = READ_string(buf)
	Ids_len := int(READ_int32(buf))
	if Ids_len>cap(data.Ids){
		data.Ids= make([]int32, Ids_len)
	}else{
		data.Ids = data.Ids[:Ids_len]
	}
	for i := 0; i < Ids_len; i++ {
		data.Ids[i] = READ_int32(buf)
	}

}
func (data *MSG_PROJECT_productplan_getPairs) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_productplan_getPairs) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_productplan_getPairs_result struct {
	QueryResultID uint32
	List []HtmlKeyValueStr
}

var pool_MSG_PROJECT_productplan_getPairs_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_getPairs_result{} }}

func GET_MSG_PROJECT_productplan_getPairs_result() *MSG_PROJECT_productplan_getPairs_result {
	return pool_MSG_PROJECT_productplan_getPairs_result.Get().(*MSG_PROJECT_productplan_getPairs_result)
}

func (data *MSG_PROJECT_productplan_getPairs_result) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_getPairs_result
}

func (data *MSG_PROJECT_productplan_getPairs_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	pool_MSG_PROJECT_productplan_getPairs_result.Put(data)
}
func (data *MSG_PROJECT_productplan_getPairs_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_getPairs_result,buf)
	WRITE_MSG_PROJECT_productplan_getPairs_result(data, buf)
}

func WRITE_MSG_PROJECT_productplan_getPairs_result(data *MSG_PROJECT_productplan_getPairs_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_PROJECT_productplan_getPairs_result(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_getPairs_result {
	data := pool_MSG_PROJECT_productplan_getPairs_result.Get().(*MSG_PROJECT_productplan_getPairs_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_getPairs_result) read(buf *libraries.MsgBuffer) {
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
func (data *MSG_PROJECT_productplan_getPairs_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_productplan_getPairs_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_productplan_insertUpdate struct {
	QueryID uint32 `db:"-"`
	Id int32 `db:"pk"`
	Product int32
	Branch int32
	Parent int32
	Projects []int32
	Title string
	Desc string
	Begin time.Time
	End time.Time
	Order string
	Deleted bool
}

var pool_MSG_PROJECT_productplan_insertUpdate = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_insertUpdate{} }}

func GET_MSG_PROJECT_productplan_insertUpdate() *MSG_PROJECT_productplan_insertUpdate {
	return pool_MSG_PROJECT_productplan_insertUpdate.Get().(*MSG_PROJECT_productplan_insertUpdate)
}

func (data *MSG_PROJECT_productplan_insertUpdate) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_insertUpdate
}

func (data *MSG_PROJECT_productplan_insertUpdate) Put() {
	data.QueryID = 0
	data.Id = 0
	data.Product = 0
	data.Branch = 0
	data.Parent = 0
	data.Projects = data.Projects[:0]
	data.Title = ``
	data.Desc = ``
	data.Begin = time.Unix(0,0)
	data.End = time.Unix(0,0)
	data.Order = ``
	data.Deleted = false
	pool_MSG_PROJECT_productplan_insertUpdate.Put(data)
}
func (data *MSG_PROJECT_productplan_insertUpdate) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_insertUpdate,buf)
	WRITE_MSG_PROJECT_productplan_insertUpdate(data, buf)
}

func WRITE_MSG_PROJECT_productplan_insertUpdate(data *MSG_PROJECT_productplan_insertUpdate, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_int32(data.Parent, buf)
	WRITE_int32(int32(len(data.Projects)), buf)
	for _, v := range data.Projects{
		WRITE_int32(v, buf)
	}
	WRITE_string(data.Title, buf)
	WRITE_string(data.Desc, buf)
	WRITE_int64(data.Begin.UnixNano(), buf)
	WRITE_int64(data.End.UnixNano(), buf)
	WRITE_string(data.Order, buf)
	WRITE_bool(data.Deleted, buf)
}

func READ_MSG_PROJECT_productplan_insertUpdate(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_insertUpdate {
	data := pool_MSG_PROJECT_productplan_insertUpdate.Get().(*MSG_PROJECT_productplan_insertUpdate)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_insertUpdate) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.Parent = READ_int32(buf)
	Projects_len := int(READ_int32(buf))
	if Projects_len>cap(data.Projects){
		data.Projects= make([]int32, Projects_len)
	}else{
		data.Projects = data.Projects[:Projects_len]
	}
	for i := 0; i < Projects_len; i++ {
		data.Projects[i] = READ_int32(buf)
	}
	data.Title = READ_string(buf)
	data.Desc = READ_string(buf)
	data.Begin = time.Unix(0, READ_int64(buf))
	data.End = time.Unix(0, READ_int64(buf))
	data.Order = READ_string(buf)
	data.Deleted = READ_bool(buf)

}
func (data *MSG_PROJECT_productplan_insertUpdate) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_productplan_insertUpdate) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_productplan_insertUpdate_result struct {
	QueryResultID uint32
	Id int32
	Result ErrCode
}

var pool_MSG_PROJECT_productplan_insertUpdate_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_insertUpdate_result{} }}

func GET_MSG_PROJECT_productplan_insertUpdate_result() *MSG_PROJECT_productplan_insertUpdate_result {
	return pool_MSG_PROJECT_productplan_insertUpdate_result.Get().(*MSG_PROJECT_productplan_insertUpdate_result)
}

func (data *MSG_PROJECT_productplan_insertUpdate_result) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_insertUpdate_result
}

func (data *MSG_PROJECT_productplan_insertUpdate_result) Put() {
	data.QueryResultID = 0
	data.Id = 0
	data.Result = 0
	pool_MSG_PROJECT_productplan_insertUpdate_result.Put(data)
}
func (data *MSG_PROJECT_productplan_insertUpdate_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_insertUpdate_result,buf)
	WRITE_MSG_PROJECT_productplan_insertUpdate_result(data, buf)
}

func WRITE_MSG_PROJECT_productplan_insertUpdate_result(data *MSG_PROJECT_productplan_insertUpdate_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(data.Id, buf)
	WRITE_ErrCode(data.Result, buf)
}

func READ_MSG_PROJECT_productplan_insertUpdate_result(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_insertUpdate_result {
	data := pool_MSG_PROJECT_productplan_insertUpdate_result.Get().(*MSG_PROJECT_productplan_insertUpdate_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_insertUpdate_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Id = READ_int32(buf)
	data.Result = READ_ErrCode(buf)

}
func (data *MSG_PROJECT_productplan_insertUpdate_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_productplan_insertUpdate_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_productplan_delete struct {
	QueryID uint32
	Id int32
	Product int32
	Branch int32
}

var pool_MSG_PROJECT_productplan_delete = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_delete{} }}

func GET_MSG_PROJECT_productplan_delete() *MSG_PROJECT_productplan_delete {
	return pool_MSG_PROJECT_productplan_delete.Get().(*MSG_PROJECT_productplan_delete)
}

func (data *MSG_PROJECT_productplan_delete) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_delete
}

func (data *MSG_PROJECT_productplan_delete) Put() {
	data.QueryID = 0
	data.Id = 0
	data.Product = 0
	data.Branch = 0
	pool_MSG_PROJECT_productplan_delete.Put(data)
}
func (data *MSG_PROJECT_productplan_delete) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_delete,buf)
	WRITE_MSG_PROJECT_productplan_delete(data, buf)
}

func WRITE_MSG_PROJECT_productplan_delete(data *MSG_PROJECT_productplan_delete, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Branch, buf)
}

func READ_MSG_PROJECT_productplan_delete(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_delete {
	data := pool_MSG_PROJECT_productplan_delete.Get().(*MSG_PROJECT_productplan_delete)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_delete) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Branch = READ_int32(buf)

}
func (data *MSG_PROJECT_productplan_delete) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_productplan_delete) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_stroy_create struct {
	QueryID uint32
	Product int32
	Branch int32
	Module int32
	Plan int32
	Source string
	SourceNote string
	AssignedTo int32
	Title string
	Color string
	Pri int8
	Estimate float64
	Spec string
	Verify string
	Mailto []int32
	Keywords string
	NeedNotReview bool
	FromBug int32
	OpenedBy int32
	ProjectID int32
}

var pool_MSG_PROJECT_stroy_create = sync.Pool{New: func() interface{} { return &MSG_PROJECT_stroy_create{} }}

func GET_MSG_PROJECT_stroy_create() *MSG_PROJECT_stroy_create {
	return pool_MSG_PROJECT_stroy_create.Get().(*MSG_PROJECT_stroy_create)
}

func (data *MSG_PROJECT_stroy_create) cmd() int32 {
	return CMD_MSG_PROJECT_stroy_create
}

func (data *MSG_PROJECT_stroy_create) Put() {
	data.QueryID = 0
	data.Product = 0
	data.Branch = 0
	data.Module = 0
	data.Plan = 0
	data.Source = ``
	data.SourceNote = ``
	data.AssignedTo = 0
	data.Title = ``
	data.Color = ``
	data.Pri = 0
	data.Estimate = 0
	data.Spec = ``
	data.Verify = ``
	data.Mailto = data.Mailto[:0]
	data.Keywords = ``
	data.NeedNotReview = false
	data.FromBug = 0
	data.OpenedBy = 0
	data.ProjectID = 0
	pool_MSG_PROJECT_stroy_create.Put(data)
}
func (data *MSG_PROJECT_stroy_create) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_stroy_create,buf)
	WRITE_MSG_PROJECT_stroy_create(data, buf)
}

func WRITE_MSG_PROJECT_stroy_create(data *MSG_PROJECT_stroy_create, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_int32(data.Module, buf)
	WRITE_int32(data.Plan, buf)
	WRITE_string(data.Source, buf)
	WRITE_string(data.SourceNote, buf)
	WRITE_int32(data.AssignedTo, buf)
	WRITE_string(data.Title, buf)
	WRITE_string(data.Color, buf)
	WRITE_int8(data.Pri, buf)
	WRITE_float64(data.Estimate, buf)
	WRITE_string(data.Spec, buf)
	WRITE_string(data.Verify, buf)
	WRITE_int32(int32(len(data.Mailto)), buf)
	for _, v := range data.Mailto{
		WRITE_int32(v, buf)
	}
	WRITE_string(data.Keywords, buf)
	WRITE_bool(data.NeedNotReview, buf)
	WRITE_int32(data.FromBug, buf)
	WRITE_int32(data.OpenedBy, buf)
	WRITE_int32(data.ProjectID, buf)
}

func READ_MSG_PROJECT_stroy_create(buf *libraries.MsgBuffer) *MSG_PROJECT_stroy_create {
	data := pool_MSG_PROJECT_stroy_create.Get().(*MSG_PROJECT_stroy_create)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_stroy_create) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Product = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.Module = READ_int32(buf)
	data.Plan = READ_int32(buf)
	data.Source = READ_string(buf)
	data.SourceNote = READ_string(buf)
	data.AssignedTo = READ_int32(buf)
	data.Title = READ_string(buf)
	data.Color = READ_string(buf)
	data.Pri = READ_int8(buf)
	data.Estimate = READ_float64(buf)
	data.Spec = READ_string(buf)
	data.Verify = READ_string(buf)
	Mailto_len := int(READ_int32(buf))
	if Mailto_len>cap(data.Mailto){
		data.Mailto= make([]int32, Mailto_len)
	}else{
		data.Mailto = data.Mailto[:Mailto_len]
	}
	for i := 0; i < Mailto_len; i++ {
		data.Mailto[i] = READ_int32(buf)
	}
	data.Keywords = READ_string(buf)
	data.NeedNotReview = READ_bool(buf)
	data.FromBug = READ_int32(buf)
	data.OpenedBy = READ_int32(buf)
	data.ProjectID = READ_int32(buf)

}
func (data *MSG_PROJECT_stroy_create) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_stroy_create) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_stroy_create_result struct {
	QueryResultID uint32
	Result int32
}

var pool_MSG_PROJECT_stroy_create_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_stroy_create_result{} }}

func GET_MSG_PROJECT_stroy_create_result() *MSG_PROJECT_stroy_create_result {
	return pool_MSG_PROJECT_stroy_create_result.Get().(*MSG_PROJECT_stroy_create_result)
}

func (data *MSG_PROJECT_stroy_create_result) cmd() int32 {
	return CMD_MSG_PROJECT_stroy_create_result
}

func (data *MSG_PROJECT_stroy_create_result) Put() {
	data.QueryResultID = 0
	data.Result = 0
	pool_MSG_PROJECT_stroy_create_result.Put(data)
}
func (data *MSG_PROJECT_stroy_create_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_stroy_create_result,buf)
	WRITE_MSG_PROJECT_stroy_create_result(data, buf)
}

func WRITE_MSG_PROJECT_stroy_create_result(data *MSG_PROJECT_stroy_create_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(data.Result, buf)
}

func READ_MSG_PROJECT_stroy_create_result(buf *libraries.MsgBuffer) *MSG_PROJECT_stroy_create_result {
	data := pool_MSG_PROJECT_stroy_create_result.Get().(*MSG_PROJECT_stroy_create_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_stroy_create_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Result = READ_int32(buf)

}
func (data *MSG_PROJECT_stroy_create_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_stroy_create_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_story_batchGetStoryStage struct {
	QueryID uint32
	Ids []int32
}

var pool_MSG_PROJECT_story_batchGetStoryStage = sync.Pool{New: func() interface{} { return &MSG_PROJECT_story_batchGetStoryStage{} }}

func GET_MSG_PROJECT_story_batchGetStoryStage() *MSG_PROJECT_story_batchGetStoryStage {
	return pool_MSG_PROJECT_story_batchGetStoryStage.Get().(*MSG_PROJECT_story_batchGetStoryStage)
}

func (data *MSG_PROJECT_story_batchGetStoryStage) cmd() int32 {
	return CMD_MSG_PROJECT_story_batchGetStoryStage
}

func (data *MSG_PROJECT_story_batchGetStoryStage) Put() {
	data.QueryID = 0
	data.Ids = data.Ids[:0]
	pool_MSG_PROJECT_story_batchGetStoryStage.Put(data)
}
func (data *MSG_PROJECT_story_batchGetStoryStage) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_story_batchGetStoryStage,buf)
	WRITE_MSG_PROJECT_story_batchGetStoryStage(data, buf)
}

func WRITE_MSG_PROJECT_story_batchGetStoryStage(data *MSG_PROJECT_story_batchGetStoryStage, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_PROJECT_story_batchGetStoryStage(buf *libraries.MsgBuffer) *MSG_PROJECT_story_batchGetStoryStage {
	data := pool_MSG_PROJECT_story_batchGetStoryStage.Get().(*MSG_PROJECT_story_batchGetStoryStage)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_story_batchGetStoryStage) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	Ids_len := int(READ_int32(buf))
	if Ids_len>cap(data.Ids){
		data.Ids= make([]int32, Ids_len)
	}else{
		data.Ids = data.Ids[:Ids_len]
	}
	for i := 0; i < Ids_len; i++ {
		data.Ids[i] = READ_int32(buf)
	}

}
func (data *MSG_PROJECT_story_batchGetStoryStage) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_story_batchGetStoryStage) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_story_batchGetStoryStage_result struct {
	QueryResultID uint32
	List map[int32][]HtmlKeyValueStr
}

var pool_MSG_PROJECT_story_batchGetStoryStage_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_story_batchGetStoryStage_result{} }}

func GET_MSG_PROJECT_story_batchGetStoryStage_result() *MSG_PROJECT_story_batchGetStoryStage_result {
	return pool_MSG_PROJECT_story_batchGetStoryStage_result.Get().(*MSG_PROJECT_story_batchGetStoryStage_result)
}

func (data *MSG_PROJECT_story_batchGetStoryStage_result) cmd() int32 {
	return CMD_MSG_PROJECT_story_batchGetStoryStage_result
}

func (data *MSG_PROJECT_story_batchGetStoryStage_result) Put() {
	data.QueryResultID = 0
	data.List = nil
	pool_MSG_PROJECT_story_batchGetStoryStage_result.Put(data)
}
func (data *MSG_PROJECT_story_batchGetStoryStage_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_story_batchGetStoryStage_result,buf)
	WRITE_MSG_PROJECT_story_batchGetStoryStage_result(data, buf)
}

func WRITE_MSG_PROJECT_story_batchGetStoryStage_result(data *MSG_PROJECT_story_batchGetStoryStage_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_map(data.List,buf)
}

func READ_MSG_PROJECT_story_batchGetStoryStage_result(buf *libraries.MsgBuffer) *MSG_PROJECT_story_batchGetStoryStage_result {
	data := pool_MSG_PROJECT_story_batchGetStoryStage_result.Get().(*MSG_PROJECT_story_batchGetStoryStage_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_story_batchGetStoryStage_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	READ_map(&data.List,buf)

}
func (data *MSG_PROJECT_story_batchGetStoryStage_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_story_batchGetStoryStage_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_story_getById struct {
	QueryID uint32
	Id int32
	Version int16
}

var pool_MSG_PROJECT_story_getById = sync.Pool{New: func() interface{} { return &MSG_PROJECT_story_getById{} }}

func GET_MSG_PROJECT_story_getById() *MSG_PROJECT_story_getById {
	return pool_MSG_PROJECT_story_getById.Get().(*MSG_PROJECT_story_getById)
}

func (data *MSG_PROJECT_story_getById) cmd() int32 {
	return CMD_MSG_PROJECT_story_getById
}

func (data *MSG_PROJECT_story_getById) Put() {
	data.QueryID = 0
	data.Id = 0
	data.Version = 0
	pool_MSG_PROJECT_story_getById.Put(data)
}
func (data *MSG_PROJECT_story_getById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_story_getById,buf)
	WRITE_MSG_PROJECT_story_getById(data, buf)
}

func WRITE_MSG_PROJECT_story_getById(data *MSG_PROJECT_story_getById, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
	WRITE_int16(data.Version, buf)
}

func READ_MSG_PROJECT_story_getById(buf *libraries.MsgBuffer) *MSG_PROJECT_story_getById {
	data := pool_MSG_PROJECT_story_getById.Get().(*MSG_PROJECT_story_getById)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_story_getById) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)
	data.Version = READ_int16(buf)

}
func (data *MSG_PROJECT_story_getById) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_story_getById) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_story_getById_result struct {
	QueryResultID uint32
	Story *MSG_PROJECT_story
}

var pool_MSG_PROJECT_story_getById_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_story_getById_result{} }}

func GET_MSG_PROJECT_story_getById_result() *MSG_PROJECT_story_getById_result {
	return pool_MSG_PROJECT_story_getById_result.Get().(*MSG_PROJECT_story_getById_result)
}

func (data *MSG_PROJECT_story_getById_result) cmd() int32 {
	return CMD_MSG_PROJECT_story_getById_result
}

func (data *MSG_PROJECT_story_getById_result) Put() {
	data.QueryResultID = 0
	if data.Story != nil {
		data.Story.Put()
		data.Story = nil
	}
	pool_MSG_PROJECT_story_getById_result.Put(data)
}
func (data *MSG_PROJECT_story_getById_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_story_getById_result,buf)
	WRITE_MSG_PROJECT_story_getById_result(data, buf)
}

func WRITE_MSG_PROJECT_story_getById_result(data *MSG_PROJECT_story_getById_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	if data.Story == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_PROJECT_story(data.Story, buf)
	}
}

func READ_MSG_PROJECT_story_getById_result(buf *libraries.MsgBuffer) *MSG_PROJECT_story_getById_result {
	data := pool_MSG_PROJECT_story_getById_result.Get().(*MSG_PROJECT_story_getById_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_story_getById_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	Story_len := int(READ_int8(buf))
	if Story_len == 1 {
		data.Story = READ_MSG_PROJECT_story(buf)
	}else{
		data.Story = nil
	}

}
func (data *MSG_PROJECT_story_getById_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_story_getById_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_project_getById struct {
	QueryID uint32
	Id int32
}

var pool_MSG_PROJECT_project_getById = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_getById{} }}

func GET_MSG_PROJECT_project_getById() *MSG_PROJECT_project_getById {
	return pool_MSG_PROJECT_project_getById.Get().(*MSG_PROJECT_project_getById)
}

func (data *MSG_PROJECT_project_getById) cmd() int32 {
	return CMD_MSG_PROJECT_project_getById
}

func (data *MSG_PROJECT_project_getById) Put() {
	data.QueryID = 0
	data.Id = 0
	pool_MSG_PROJECT_project_getById.Put(data)
}
func (data *MSG_PROJECT_project_getById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_getById,buf)
	WRITE_MSG_PROJECT_project_getById(data, buf)
}

func WRITE_MSG_PROJECT_project_getById(data *MSG_PROJECT_project_getById, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
}

func READ_MSG_PROJECT_project_getById(buf *libraries.MsgBuffer) *MSG_PROJECT_project_getById {
	data := pool_MSG_PROJECT_project_getById.Get().(*MSG_PROJECT_project_getById)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_getById) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)

}
func (data *MSG_PROJECT_project_getById) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_getById) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_getById_result struct {
	QueryResultID uint32
	Project *MSG_PROJECT_project_cache
}

var pool_MSG_PROJECT_project_getById_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_getById_result{} }}

func GET_MSG_PROJECT_project_getById_result() *MSG_PROJECT_project_getById_result {
	return pool_MSG_PROJECT_project_getById_result.Get().(*MSG_PROJECT_project_getById_result)
}

func (data *MSG_PROJECT_project_getById_result) cmd() int32 {
	return CMD_MSG_PROJECT_project_getById_result
}

func (data *MSG_PROJECT_project_getById_result) Put() {
	data.QueryResultID = 0
	if data.Project != nil {
		data.Project.Put()
		data.Project = nil
	}
	pool_MSG_PROJECT_project_getById_result.Put(data)
}
func (data *MSG_PROJECT_project_getById_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_getById_result,buf)
	WRITE_MSG_PROJECT_project_getById_result(data, buf)
}

func WRITE_MSG_PROJECT_project_getById_result(data *MSG_PROJECT_project_getById_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	if data.Project == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_PROJECT_project_cache(data.Project, buf)
	}
}

func READ_MSG_PROJECT_project_getById_result(buf *libraries.MsgBuffer) *MSG_PROJECT_project_getById_result {
	data := pool_MSG_PROJECT_project_getById_result.Get().(*MSG_PROJECT_project_getById_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_getById_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	Project_len := int(READ_int8(buf))
	if Project_len == 1 {
		data.Project = READ_MSG_PROJECT_project_cache(buf)
	}else{
		data.Project = nil
	}

}
func (data *MSG_PROJECT_project_getById_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_project_getById_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_project_cache struct {
	Id int32
	IsCat bool
	CatID int32
	Type string
	Parent int32
	Name string
	Code string
	Begin time.Time
	End time.Time
	Days int16
	Status string
	Statge int8
	Pri int8
	Desc string
	OpenedBy int32
	OpenedDate time.Time
	OpenedVersion string
	ClosedBy int32
	ClosedDate time.Time
	CanceledBy int32
	CanceledDate time.Time
	PO int32
	PM int32
	QD int32
	RD int32
	Team string
	Acl string
	Whitelist []int32
	Order int32
	Deleted bool
	FtpPath string
	Products []int32
	Branchs []int32
	Storys []int32
	Plans []int32
	Delay int64 `db:"-"`
	Hours map[string]float64 `db:"-"`
	Teams []*MSG_USER_team_info `db:"-"`
}

var pool_MSG_PROJECT_project_cache = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_cache{} }}

func GET_MSG_PROJECT_project_cache() *MSG_PROJECT_project_cache {
	return pool_MSG_PROJECT_project_cache.Get().(*MSG_PROJECT_project_cache)
}

func (data *MSG_PROJECT_project_cache) cmd() int32 {
	return CMD_MSG_PROJECT_project_cache
}

func (data *MSG_PROJECT_project_cache) Put() {
	data.Id = 0
	data.IsCat = false
	data.CatID = 0
	data.Type = ``
	data.Parent = 0
	data.Name = ``
	data.Code = ``
	data.Begin = time.Unix(0,0)
	data.End = time.Unix(0,0)
	data.Days = 0
	data.Status = ``
	data.Statge = 0
	data.Pri = 0
	data.Desc = ``
	data.OpenedBy = 0
	data.OpenedDate = time.Unix(0,0)
	data.OpenedVersion = ``
	data.ClosedBy = 0
	data.ClosedDate = time.Unix(0,0)
	data.CanceledBy = 0
	data.CanceledDate = time.Unix(0,0)
	data.PO = 0
	data.PM = 0
	data.QD = 0
	data.RD = 0
	data.Team = ``
	data.Acl = ``
	data.Whitelist = data.Whitelist[:0]
	data.Order = 0
	data.Deleted = false
	data.FtpPath = ``
	data.Products = data.Products[:0]
	data.Branchs = data.Branchs[:0]
	data.Storys = data.Storys[:0]
	data.Plans = data.Plans[:0]
	data.Delay = 0
	data.Hours = nil
	for _,v := range data.Teams {
		v.Put()
	}
	data.Teams = data.Teams[:0]
	pool_MSG_PROJECT_project_cache.Put(data)
}
func (data *MSG_PROJECT_project_cache) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_cache,buf)
	WRITE_MSG_PROJECT_project_cache(data, buf)
}

func WRITE_MSG_PROJECT_project_cache(data *MSG_PROJECT_project_cache, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_bool(data.IsCat, buf)
	WRITE_int32(data.CatID, buf)
	WRITE_string(data.Type, buf)
	WRITE_int32(data.Parent, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Code, buf)
	WRITE_int64(data.Begin.UnixNano(), buf)
	WRITE_int64(data.End.UnixNano(), buf)
	WRITE_int16(data.Days, buf)
	WRITE_string(data.Status, buf)
	WRITE_int8(data.Statge, buf)
	WRITE_int8(data.Pri, buf)
	WRITE_string(data.Desc, buf)
	WRITE_int32(data.OpenedBy, buf)
	WRITE_int64(data.OpenedDate.UnixNano(), buf)
	WRITE_string(data.OpenedVersion, buf)
	WRITE_int32(data.ClosedBy, buf)
	WRITE_int64(data.ClosedDate.UnixNano(), buf)
	WRITE_int32(data.CanceledBy, buf)
	WRITE_int64(data.CanceledDate.UnixNano(), buf)
	WRITE_int32(data.PO, buf)
	WRITE_int32(data.PM, buf)
	WRITE_int32(data.QD, buf)
	WRITE_int32(data.RD, buf)
	WRITE_string(data.Team, buf)
	WRITE_string(data.Acl, buf)
	WRITE_int32(int32(len(data.Whitelist)), buf)
	for _, v := range data.Whitelist{
		WRITE_int32(v, buf)
	}
	WRITE_int32(data.Order, buf)
	WRITE_bool(data.Deleted, buf)
	WRITE_string(data.FtpPath, buf)
	WRITE_int32(int32(len(data.Products)), buf)
	for _, v := range data.Products{
		WRITE_int32(v, buf)
	}
	WRITE_int32(int32(len(data.Branchs)), buf)
	for _, v := range data.Branchs{
		WRITE_int32(v, buf)
	}
	WRITE_int32(int32(len(data.Storys)), buf)
	for _, v := range data.Storys{
		WRITE_int32(v, buf)
	}
	WRITE_int32(int32(len(data.Plans)), buf)
	for _, v := range data.Plans{
		WRITE_int32(v, buf)
	}
	WRITE_int64(data.Delay, buf)
	WRITE_map(data.Hours,buf)
	WRITE_int32(int32(len(data.Teams)), buf)
	for _, v := range data.Teams{
		WRITE_MSG_USER_team_info(v, buf)
	}
}

func READ_MSG_PROJECT_project_cache(buf *libraries.MsgBuffer) *MSG_PROJECT_project_cache {
	data := pool_MSG_PROJECT_project_cache.Get().(*MSG_PROJECT_project_cache)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_cache) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.IsCat = READ_bool(buf)
	data.CatID = READ_int32(buf)
	data.Type = READ_string(buf)
	data.Parent = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Code = READ_string(buf)
	data.Begin = time.Unix(0, READ_int64(buf))
	data.End = time.Unix(0, READ_int64(buf))
	data.Days = READ_int16(buf)
	data.Status = READ_string(buf)
	data.Statge = READ_int8(buf)
	data.Pri = READ_int8(buf)
	data.Desc = READ_string(buf)
	data.OpenedBy = READ_int32(buf)
	data.OpenedDate = time.Unix(0, READ_int64(buf))
	data.OpenedVersion = READ_string(buf)
	data.ClosedBy = READ_int32(buf)
	data.ClosedDate = time.Unix(0, READ_int64(buf))
	data.CanceledBy = READ_int32(buf)
	data.CanceledDate = time.Unix(0, READ_int64(buf))
	data.PO = READ_int32(buf)
	data.PM = READ_int32(buf)
	data.QD = READ_int32(buf)
	data.RD = READ_int32(buf)
	data.Team = READ_string(buf)
	data.Acl = READ_string(buf)
	Whitelist_len := int(READ_int32(buf))
	if Whitelist_len>cap(data.Whitelist){
		data.Whitelist= make([]int32, Whitelist_len)
	}else{
		data.Whitelist = data.Whitelist[:Whitelist_len]
	}
	for i := 0; i < Whitelist_len; i++ {
		data.Whitelist[i] = READ_int32(buf)
	}
	data.Order = READ_int32(buf)
	data.Deleted = READ_bool(buf)
	data.FtpPath = READ_string(buf)
	Products_len := int(READ_int32(buf))
	if Products_len>cap(data.Products){
		data.Products= make([]int32, Products_len)
	}else{
		data.Products = data.Products[:Products_len]
	}
	for i := 0; i < Products_len; i++ {
		data.Products[i] = READ_int32(buf)
	}
	Branchs_len := int(READ_int32(buf))
	if Branchs_len>cap(data.Branchs){
		data.Branchs= make([]int32, Branchs_len)
	}else{
		data.Branchs = data.Branchs[:Branchs_len]
	}
	for i := 0; i < Branchs_len; i++ {
		data.Branchs[i] = READ_int32(buf)
	}
	Storys_len := int(READ_int32(buf))
	if Storys_len>cap(data.Storys){
		data.Storys= make([]int32, Storys_len)
	}else{
		data.Storys = data.Storys[:Storys_len]
	}
	for i := 0; i < Storys_len; i++ {
		data.Storys[i] = READ_int32(buf)
	}
	Plans_len := int(READ_int32(buf))
	if Plans_len>cap(data.Plans){
		data.Plans= make([]int32, Plans_len)
	}else{
		data.Plans = data.Plans[:Plans_len]
	}
	for i := 0; i < Plans_len; i++ {
		data.Plans[i] = READ_int32(buf)
	}
	data.Delay = READ_int64(buf)
	READ_map(&data.Hours,buf)
	Teams_len := int(READ_int32(buf))
	if Teams_len>cap(data.Teams){
		data.Teams= make([]*MSG_USER_team_info, Teams_len)
	}else{
		data.Teams = data.Teams[:Teams_len]
	}
	for i := 0; i < Teams_len; i++ {
		data.Teams[i] = READ_MSG_USER_team_info(buf)
	}

}

type MSG_PROJECT_StoryStage struct {
	Story int32
	Branch int32
	Stage string
}

var pool_MSG_PROJECT_StoryStage = sync.Pool{New: func() interface{} { return &MSG_PROJECT_StoryStage{} }}

func GET_MSG_PROJECT_StoryStage() *MSG_PROJECT_StoryStage {
	return pool_MSG_PROJECT_StoryStage.Get().(*MSG_PROJECT_StoryStage)
}

func (data *MSG_PROJECT_StoryStage) cmd() int32 {
	return CMD_MSG_PROJECT_StoryStage
}

func (data *MSG_PROJECT_StoryStage) Put() {
	data.Story = 0
	data.Branch = 0
	data.Stage = ``
	pool_MSG_PROJECT_StoryStage.Put(data)
}
func (data *MSG_PROJECT_StoryStage) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_StoryStage,buf)
	WRITE_MSG_PROJECT_StoryStage(data, buf)
}

func WRITE_MSG_PROJECT_StoryStage(data *MSG_PROJECT_StoryStage, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Story, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_string(data.Stage, buf)
}

func READ_MSG_PROJECT_StoryStage(buf *libraries.MsgBuffer) *MSG_PROJECT_StoryStage {
	data := pool_MSG_PROJECT_StoryStage.Get().(*MSG_PROJECT_StoryStage)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_StoryStage) read(buf *libraries.MsgBuffer) {
	data.Story = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.Stage = READ_string(buf)

}

type MSG_PROJECT_TASK struct {
	Id int32
	Ancestor int32
	Parent int32
	Project int32
	Module int32
	Story int32
	StoryVersion int16
	FromBug int32
	Name string
	Type string
	Pri int8
	Estimate float64
	Consumed float64
	Left float64
	Deadline time.Time
	Status string
	Color string
	Mailto []int32
	Desc string
	OpenedBy int32
	OpenedDate time.Time
	AssignedTo int32
	AssignedDate time.Time
	EstStarted time.Time
	RealStarted time.Time
	FinishedBy int32
	FinishedDate time.Time
	FinishedList string
	CanceledBy int32
	CanceledDate time.Time
	ClosedBy int32
	ClosedDate time.Time
	ClosedReason string
	LastEditedBy int32
	LastEditedDate time.Time
	Examine bool
	ExamineDate time.Time
	ExamineBy int32
	Deleted bool
	Finalfile bool
	Proofreading bool
	Team int32
	PlaceOrder bool
	StoryID int32
	StoryTitle string
	StoryStatus string
	LatestStoryVersion int16
	Product int32
	Branch int32
	Progress int `db:"-"`
	Delay int64 `db:"-"`
	Children []*MSG_PROJECT_TASK `db:"-"`
	Grandchildren []*MSG_PROJECT_TASK `db:"-"`
}

var pool_MSG_PROJECT_TASK = sync.Pool{New: func() interface{} { return &MSG_PROJECT_TASK{} }}

func GET_MSG_PROJECT_TASK() *MSG_PROJECT_TASK {
	return pool_MSG_PROJECT_TASK.Get().(*MSG_PROJECT_TASK)
}

func (data *MSG_PROJECT_TASK) cmd() int32 {
	return CMD_MSG_PROJECT_TASK
}

func (data *MSG_PROJECT_TASK) Put() {
	data.Id = 0
	data.Ancestor = 0
	data.Parent = 0
	data.Project = 0
	data.Module = 0
	data.Story = 0
	data.StoryVersion = 0
	data.FromBug = 0
	data.Name = ``
	data.Type = ``
	data.Pri = 0
	data.Estimate = 0
	data.Consumed = 0
	data.Left = 0
	data.Deadline = time.Unix(0,0)
	data.Status = ``
	data.Color = ``
	data.Mailto = data.Mailto[:0]
	data.Desc = ``
	data.OpenedBy = 0
	data.OpenedDate = time.Unix(0,0)
	data.AssignedTo = 0
	data.AssignedDate = time.Unix(0,0)
	data.EstStarted = time.Unix(0,0)
	data.RealStarted = time.Unix(0,0)
	data.FinishedBy = 0
	data.FinishedDate = time.Unix(0,0)
	data.FinishedList = ``
	data.CanceledBy = 0
	data.CanceledDate = time.Unix(0,0)
	data.ClosedBy = 0
	data.ClosedDate = time.Unix(0,0)
	data.ClosedReason = ``
	data.LastEditedBy = 0
	data.LastEditedDate = time.Unix(0,0)
	data.Examine = false
	data.ExamineDate = time.Unix(0,0)
	data.ExamineBy = 0
	data.Deleted = false
	data.Finalfile = false
	data.Proofreading = false
	data.Team = 0
	data.PlaceOrder = false
	data.StoryID = 0
	data.StoryTitle = ``
	data.StoryStatus = ``
	data.LatestStoryVersion = 0
	data.Product = 0
	data.Branch = 0
	data.Progress = 0
	data.Delay = 0
	for _,v := range data.Children {
		v.Put()
	}
	data.Children = data.Children[:0]
	for _,v := range data.Grandchildren {
		v.Put()
	}
	data.Grandchildren = data.Grandchildren[:0]
	pool_MSG_PROJECT_TASK.Put(data)
}
func (data *MSG_PROJECT_TASK) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_TASK,buf)
	WRITE_MSG_PROJECT_TASK(data, buf)
}

func WRITE_MSG_PROJECT_TASK(data *MSG_PROJECT_TASK, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Ancestor, buf)
	WRITE_int32(data.Parent, buf)
	WRITE_int32(data.Project, buf)
	WRITE_int32(data.Module, buf)
	WRITE_int32(data.Story, buf)
	WRITE_int16(data.StoryVersion, buf)
	WRITE_int32(data.FromBug, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Type, buf)
	WRITE_int8(data.Pri, buf)
	WRITE_float64(data.Estimate, buf)
	WRITE_float64(data.Consumed, buf)
	WRITE_float64(data.Left, buf)
	WRITE_int64(data.Deadline.UnixNano(), buf)
	WRITE_string(data.Status, buf)
	WRITE_string(data.Color, buf)
	WRITE_int32(int32(len(data.Mailto)), buf)
	for _, v := range data.Mailto{
		WRITE_int32(v, buf)
	}
	WRITE_string(data.Desc, buf)
	WRITE_int32(data.OpenedBy, buf)
	WRITE_int64(data.OpenedDate.UnixNano(), buf)
	WRITE_int32(data.AssignedTo, buf)
	WRITE_int64(data.AssignedDate.UnixNano(), buf)
	WRITE_int64(data.EstStarted.UnixNano(), buf)
	WRITE_int64(data.RealStarted.UnixNano(), buf)
	WRITE_int32(data.FinishedBy, buf)
	WRITE_int64(data.FinishedDate.UnixNano(), buf)
	WRITE_string(data.FinishedList, buf)
	WRITE_int32(data.CanceledBy, buf)
	WRITE_int64(data.CanceledDate.UnixNano(), buf)
	WRITE_int32(data.ClosedBy, buf)
	WRITE_int64(data.ClosedDate.UnixNano(), buf)
	WRITE_string(data.ClosedReason, buf)
	WRITE_int32(data.LastEditedBy, buf)
	WRITE_int64(data.LastEditedDate.UnixNano(), buf)
	WRITE_bool(data.Examine, buf)
	WRITE_int64(data.ExamineDate.UnixNano(), buf)
	WRITE_int32(data.ExamineBy, buf)
	WRITE_bool(data.Deleted, buf)
	WRITE_bool(data.Finalfile, buf)
	WRITE_bool(data.Proofreading, buf)
	WRITE_int32(data.Team, buf)
	WRITE_bool(data.PlaceOrder, buf)
	WRITE_int32(data.StoryID, buf)
	WRITE_string(data.StoryTitle, buf)
	WRITE_string(data.StoryStatus, buf)
	WRITE_int16(data.LatestStoryVersion, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_int(data.Progress, buf)
	WRITE_int64(data.Delay, buf)
	WRITE_int32(int32(len(data.Children)), buf)
	for _, v := range data.Children{
		WRITE_MSG_PROJECT_TASK(v, buf)
	}
	WRITE_int32(int32(len(data.Grandchildren)), buf)
	for _, v := range data.Grandchildren{
		WRITE_MSG_PROJECT_TASK(v, buf)
	}
}

func READ_MSG_PROJECT_TASK(buf *libraries.MsgBuffer) *MSG_PROJECT_TASK {
	data := pool_MSG_PROJECT_TASK.Get().(*MSG_PROJECT_TASK)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_TASK) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Ancestor = READ_int32(buf)
	data.Parent = READ_int32(buf)
	data.Project = READ_int32(buf)
	data.Module = READ_int32(buf)
	data.Story = READ_int32(buf)
	data.StoryVersion = READ_int16(buf)
	data.FromBug = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Type = READ_string(buf)
	data.Pri = READ_int8(buf)
	data.Estimate = READ_float64(buf)
	data.Consumed = READ_float64(buf)
	data.Left = READ_float64(buf)
	data.Deadline = time.Unix(0, READ_int64(buf))
	data.Status = READ_string(buf)
	data.Color = READ_string(buf)
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
	data.OpenedBy = READ_int32(buf)
	data.OpenedDate = time.Unix(0, READ_int64(buf))
	data.AssignedTo = READ_int32(buf)
	data.AssignedDate = time.Unix(0, READ_int64(buf))
	data.EstStarted = time.Unix(0, READ_int64(buf))
	data.RealStarted = time.Unix(0, READ_int64(buf))
	data.FinishedBy = READ_int32(buf)
	data.FinishedDate = time.Unix(0, READ_int64(buf))
	data.FinishedList = READ_string(buf)
	data.CanceledBy = READ_int32(buf)
	data.CanceledDate = time.Unix(0, READ_int64(buf))
	data.ClosedBy = READ_int32(buf)
	data.ClosedDate = time.Unix(0, READ_int64(buf))
	data.ClosedReason = READ_string(buf)
	data.LastEditedBy = READ_int32(buf)
	data.LastEditedDate = time.Unix(0, READ_int64(buf))
	data.Examine = READ_bool(buf)
	data.ExamineDate = time.Unix(0, READ_int64(buf))
	data.ExamineBy = READ_int32(buf)
	data.Deleted = READ_bool(buf)
	data.Finalfile = READ_bool(buf)
	data.Proofreading = READ_bool(buf)
	data.Team = READ_int32(buf)
	data.PlaceOrder = READ_bool(buf)
	data.StoryID = READ_int32(buf)
	data.StoryTitle = READ_string(buf)
	data.StoryStatus = READ_string(buf)
	data.LatestStoryVersion = READ_int16(buf)
	data.Product = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.Progress = READ_int(buf)
	data.Delay = READ_int64(buf)
	Children_len := int(READ_int32(buf))
	if Children_len>cap(data.Children){
		data.Children= make([]*MSG_PROJECT_TASK, Children_len)
	}else{
		data.Children = data.Children[:Children_len]
	}
	for i := 0; i < Children_len; i++ {
		data.Children[i] = READ_MSG_PROJECT_TASK(buf)
	}
	Grandchildren_len := int(READ_int32(buf))
	if Grandchildren_len>cap(data.Grandchildren){
		data.Grandchildren= make([]*MSG_PROJECT_TASK, Grandchildren_len)
	}else{
		data.Grandchildren = data.Grandchildren[:Grandchildren_len]
	}
	for i := 0; i < Grandchildren_len; i++ {
		data.Grandchildren[i] = READ_MSG_PROJECT_TASK(buf)
	}

}

type MSG_PROJECT_productplan_getById struct {
	QueryID uint32
	Id int32
}

var pool_MSG_PROJECT_productplan_getById = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_getById{} }}

func GET_MSG_PROJECT_productplan_getById() *MSG_PROJECT_productplan_getById {
	return pool_MSG_PROJECT_productplan_getById.Get().(*MSG_PROJECT_productplan_getById)
}

func (data *MSG_PROJECT_productplan_getById) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_getById
}

func (data *MSG_PROJECT_productplan_getById) Put() {
	data.QueryID = 0
	data.Id = 0
	pool_MSG_PROJECT_productplan_getById.Put(data)
}
func (data *MSG_PROJECT_productplan_getById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_getById,buf)
	WRITE_MSG_PROJECT_productplan_getById(data, buf)
}

func WRITE_MSG_PROJECT_productplan_getById(data *MSG_PROJECT_productplan_getById, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
}

func READ_MSG_PROJECT_productplan_getById(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_getById {
	data := pool_MSG_PROJECT_productplan_getById.Get().(*MSG_PROJECT_productplan_getById)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_getById) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)

}
func (data *MSG_PROJECT_productplan_getById) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_productplan_getById) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_productplan_getById_result struct {
	QueryResultID uint32
	Info *MSG_PROJECT_productplan
}

var pool_MSG_PROJECT_productplan_getById_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan_getById_result{} }}

func GET_MSG_PROJECT_productplan_getById_result() *MSG_PROJECT_productplan_getById_result {
	return pool_MSG_PROJECT_productplan_getById_result.Get().(*MSG_PROJECT_productplan_getById_result)
}

func (data *MSG_PROJECT_productplan_getById_result) cmd() int32 {
	return CMD_MSG_PROJECT_productplan_getById_result
}

func (data *MSG_PROJECT_productplan_getById_result) Put() {
	data.QueryResultID = 0
	if data.Info != nil {
		data.Info.Put()
		data.Info = nil
	}
	pool_MSG_PROJECT_productplan_getById_result.Put(data)
}
func (data *MSG_PROJECT_productplan_getById_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan_getById_result,buf)
	WRITE_MSG_PROJECT_productplan_getById_result(data, buf)
}

func WRITE_MSG_PROJECT_productplan_getById_result(data *MSG_PROJECT_productplan_getById_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	if data.Info == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_PROJECT_productplan(data.Info, buf)
	}
}

func READ_MSG_PROJECT_productplan_getById_result(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan_getById_result {
	data := pool_MSG_PROJECT_productplan_getById_result.Get().(*MSG_PROJECT_productplan_getById_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan_getById_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_PROJECT_productplan(buf)
	}else{
		data.Info = nil
	}

}
func (data *MSG_PROJECT_productplan_getById_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_productplan_getById_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_productplan struct {
	Id int32
	Product int32
	Branch int32
	Parent int32
	Projects []int32
	Title string
	Desc string
	Begin time.Time
	End time.Time
	Order string
	Deleted bool
}

var pool_MSG_PROJECT_productplan = sync.Pool{New: func() interface{} { return &MSG_PROJECT_productplan{} }}

func GET_MSG_PROJECT_productplan() *MSG_PROJECT_productplan {
	return pool_MSG_PROJECT_productplan.Get().(*MSG_PROJECT_productplan)
}

func (data *MSG_PROJECT_productplan) cmd() int32 {
	return CMD_MSG_PROJECT_productplan
}

func (data *MSG_PROJECT_productplan) Put() {
	data.Id = 0
	data.Product = 0
	data.Branch = 0
	data.Parent = 0
	data.Projects = data.Projects[:0]
	data.Title = ``
	data.Desc = ``
	data.Begin = time.Unix(0,0)
	data.End = time.Unix(0,0)
	data.Order = ``
	data.Deleted = false
	pool_MSG_PROJECT_productplan.Put(data)
}
func (data *MSG_PROJECT_productplan) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_productplan,buf)
	WRITE_MSG_PROJECT_productplan(data, buf)
}

func WRITE_MSG_PROJECT_productplan(data *MSG_PROJECT_productplan, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_int32(data.Parent, buf)
	WRITE_int32(int32(len(data.Projects)), buf)
	for _, v := range data.Projects{
		WRITE_int32(v, buf)
	}
	WRITE_string(data.Title, buf)
	WRITE_string(data.Desc, buf)
	WRITE_int64(data.Begin.UnixNano(), buf)
	WRITE_int64(data.End.UnixNano(), buf)
	WRITE_string(data.Order, buf)
	WRITE_bool(data.Deleted, buf)
}

func READ_MSG_PROJECT_productplan(buf *libraries.MsgBuffer) *MSG_PROJECT_productplan {
	data := pool_MSG_PROJECT_productplan.Get().(*MSG_PROJECT_productplan)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_productplan) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.Parent = READ_int32(buf)
	Projects_len := int(READ_int32(buf))
	if Projects_len>cap(data.Projects){
		data.Projects= make([]int32, Projects_len)
	}else{
		data.Projects = data.Projects[:Projects_len]
	}
	for i := 0; i < Projects_len; i++ {
		data.Projects[i] = READ_int32(buf)
	}
	data.Title = READ_string(buf)
	data.Desc = READ_string(buf)
	data.Begin = time.Unix(0, READ_int64(buf))
	data.End = time.Unix(0, READ_int64(buf))
	data.Order = READ_string(buf)
	data.Deleted = READ_bool(buf)

}

type MSG_PROJECT_build struct {
	Id int32
	Product int32
	Branch int32
	Project int32
	Name string
	ScmPath string
	FilePath string
	Date time.Time
	Stories []int32
	Bugs []int32
	Builder string
	Desc string
	Deleted bool
}

var pool_MSG_PROJECT_build = sync.Pool{New: func() interface{} { return &MSG_PROJECT_build{} }}

func GET_MSG_PROJECT_build() *MSG_PROJECT_build {
	return pool_MSG_PROJECT_build.Get().(*MSG_PROJECT_build)
}

func (data *MSG_PROJECT_build) cmd() int32 {
	return CMD_MSG_PROJECT_build
}

func (data *MSG_PROJECT_build) Put() {
	data.Id = 0
	data.Product = 0
	data.Branch = 0
	data.Project = 0
	data.Name = ``
	data.ScmPath = ``
	data.FilePath = ``
	data.Date = time.Unix(0,0)
	data.Stories = data.Stories[:0]
	data.Bugs = data.Bugs[:0]
	data.Builder = ``
	data.Desc = ``
	data.Deleted = false
	pool_MSG_PROJECT_build.Put(data)
}
func (data *MSG_PROJECT_build) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_build,buf)
	WRITE_MSG_PROJECT_build(data, buf)
}

func WRITE_MSG_PROJECT_build(data *MSG_PROJECT_build, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_int32(data.Project, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.ScmPath, buf)
	WRITE_string(data.FilePath, buf)
	WRITE_int64(data.Date.UnixNano(), buf)
	WRITE_int32(int32(len(data.Stories)), buf)
	for _, v := range data.Stories{
		WRITE_int32(v, buf)
	}
	WRITE_int32(int32(len(data.Bugs)), buf)
	for _, v := range data.Bugs{
		WRITE_int32(v, buf)
	}
	WRITE_string(data.Builder, buf)
	WRITE_string(data.Desc, buf)
	WRITE_bool(data.Deleted, buf)
}

func READ_MSG_PROJECT_build(buf *libraries.MsgBuffer) *MSG_PROJECT_build {
	data := pool_MSG_PROJECT_build.Get().(*MSG_PROJECT_build)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_build) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.Project = READ_int32(buf)
	data.Name = READ_string(buf)
	data.ScmPath = READ_string(buf)
	data.FilePath = READ_string(buf)
	data.Date = time.Unix(0, READ_int64(buf))
	Stories_len := int(READ_int32(buf))
	if Stories_len>cap(data.Stories){
		data.Stories= make([]int32, Stories_len)
	}else{
		data.Stories = data.Stories[:Stories_len]
	}
	for i := 0; i < Stories_len; i++ {
		data.Stories[i] = READ_int32(buf)
	}
	Bugs_len := int(READ_int32(buf))
	if Bugs_len>cap(data.Bugs){
		data.Bugs= make([]int32, Bugs_len)
	}else{
		data.Bugs = data.Bugs[:Bugs_len]
	}
	for i := 0; i < Bugs_len; i++ {
		data.Bugs[i] = READ_int32(buf)
	}
	data.Builder = READ_string(buf)
	data.Desc = READ_string(buf)
	data.Deleted = READ_bool(buf)

}

type MSG_PROJECT_build_getById struct {
	QueryID uint32
	Id int32
}

var pool_MSG_PROJECT_build_getById = sync.Pool{New: func() interface{} { return &MSG_PROJECT_build_getById{} }}

func GET_MSG_PROJECT_build_getById() *MSG_PROJECT_build_getById {
	return pool_MSG_PROJECT_build_getById.Get().(*MSG_PROJECT_build_getById)
}

func (data *MSG_PROJECT_build_getById) cmd() int32 {
	return CMD_MSG_PROJECT_build_getById
}

func (data *MSG_PROJECT_build_getById) Put() {
	data.QueryID = 0
	data.Id = 0
	pool_MSG_PROJECT_build_getById.Put(data)
}
func (data *MSG_PROJECT_build_getById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_build_getById,buf)
	WRITE_MSG_PROJECT_build_getById(data, buf)
}

func WRITE_MSG_PROJECT_build_getById(data *MSG_PROJECT_build_getById, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
}

func READ_MSG_PROJECT_build_getById(buf *libraries.MsgBuffer) *MSG_PROJECT_build_getById {
	data := pool_MSG_PROJECT_build_getById.Get().(*MSG_PROJECT_build_getById)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_build_getById) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)

}
func (data *MSG_PROJECT_build_getById) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_build_getById) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_build_getById_result struct {
	QueryResultID uint32
	Info *MSG_PROJECT_build
}

var pool_MSG_PROJECT_build_getById_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_build_getById_result{} }}

func GET_MSG_PROJECT_build_getById_result() *MSG_PROJECT_build_getById_result {
	return pool_MSG_PROJECT_build_getById_result.Get().(*MSG_PROJECT_build_getById_result)
}

func (data *MSG_PROJECT_build_getById_result) cmd() int32 {
	return CMD_MSG_PROJECT_build_getById_result
}

func (data *MSG_PROJECT_build_getById_result) Put() {
	data.QueryResultID = 0
	if data.Info != nil {
		data.Info.Put()
		data.Info = nil
	}
	pool_MSG_PROJECT_build_getById_result.Put(data)
}
func (data *MSG_PROJECT_build_getById_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_build_getById_result,buf)
	WRITE_MSG_PROJECT_build_getById_result(data, buf)
}

func WRITE_MSG_PROJECT_build_getById_result(data *MSG_PROJECT_build_getById_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	if data.Info == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_PROJECT_build(data.Info, buf)
	}
}

func READ_MSG_PROJECT_build_getById_result(buf *libraries.MsgBuffer) *MSG_PROJECT_build_getById_result {
	data := pool_MSG_PROJECT_build_getById_result.Get().(*MSG_PROJECT_build_getById_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_build_getById_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_PROJECT_build(buf)
	}else{
		data.Info = nil
	}

}
func (data *MSG_PROJECT_build_getById_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_build_getById_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_release struct {
	Id int32
	Product int32
	Branch int32
	Build int32
	Name string
	Marker bool
	Date time.Time
	Stories []int32
	Bugs []int32
	LeftBugs string
	Desc string
	Status string
	Deleted bool
}

var pool_MSG_PROJECT_release = sync.Pool{New: func() interface{} { return &MSG_PROJECT_release{} }}

func GET_MSG_PROJECT_release() *MSG_PROJECT_release {
	return pool_MSG_PROJECT_release.Get().(*MSG_PROJECT_release)
}

func (data *MSG_PROJECT_release) cmd() int32 {
	return CMD_MSG_PROJECT_release
}

func (data *MSG_PROJECT_release) Put() {
	data.Id = 0
	data.Product = 0
	data.Branch = 0
	data.Build = 0
	data.Name = ``
	data.Marker = false
	data.Date = time.Unix(0,0)
	data.Stories = data.Stories[:0]
	data.Bugs = data.Bugs[:0]
	data.LeftBugs = ``
	data.Desc = ``
	data.Status = ``
	data.Deleted = false
	pool_MSG_PROJECT_release.Put(data)
}
func (data *MSG_PROJECT_release) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_release,buf)
	WRITE_MSG_PROJECT_release(data, buf)
}

func WRITE_MSG_PROJECT_release(data *MSG_PROJECT_release, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Product, buf)
	WRITE_int32(data.Branch, buf)
	WRITE_int32(data.Build, buf)
	WRITE_string(data.Name, buf)
	WRITE_bool(data.Marker, buf)
	WRITE_int64(data.Date.UnixNano(), buf)
	WRITE_int32(int32(len(data.Stories)), buf)
	for _, v := range data.Stories{
		WRITE_int32(v, buf)
	}
	WRITE_int32(int32(len(data.Bugs)), buf)
	for _, v := range data.Bugs{
		WRITE_int32(v, buf)
	}
	WRITE_string(data.LeftBugs, buf)
	WRITE_string(data.Desc, buf)
	WRITE_string(data.Status, buf)
	WRITE_bool(data.Deleted, buf)
}

func READ_MSG_PROJECT_release(buf *libraries.MsgBuffer) *MSG_PROJECT_release {
	data := pool_MSG_PROJECT_release.Get().(*MSG_PROJECT_release)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_release) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Product = READ_int32(buf)
	data.Branch = READ_int32(buf)
	data.Build = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Marker = READ_bool(buf)
	data.Date = time.Unix(0, READ_int64(buf))
	Stories_len := int(READ_int32(buf))
	if Stories_len>cap(data.Stories){
		data.Stories= make([]int32, Stories_len)
	}else{
		data.Stories = data.Stories[:Stories_len]
	}
	for i := 0; i < Stories_len; i++ {
		data.Stories[i] = READ_int32(buf)
	}
	Bugs_len := int(READ_int32(buf))
	if Bugs_len>cap(data.Bugs){
		data.Bugs= make([]int32, Bugs_len)
	}else{
		data.Bugs = data.Bugs[:Bugs_len]
	}
	for i := 0; i < Bugs_len; i++ {
		data.Bugs[i] = READ_int32(buf)
	}
	data.LeftBugs = READ_string(buf)
	data.Desc = READ_string(buf)
	data.Status = READ_string(buf)
	data.Deleted = READ_bool(buf)

}

type MSG_PROJECT_release_getById struct {
	QueryID uint32
	Id int32
}

var pool_MSG_PROJECT_release_getById = sync.Pool{New: func() interface{} { return &MSG_PROJECT_release_getById{} }}

func GET_MSG_PROJECT_release_getById() *MSG_PROJECT_release_getById {
	return pool_MSG_PROJECT_release_getById.Get().(*MSG_PROJECT_release_getById)
}

func (data *MSG_PROJECT_release_getById) cmd() int32 {
	return CMD_MSG_PROJECT_release_getById
}

func (data *MSG_PROJECT_release_getById) Put() {
	data.QueryID = 0
	data.Id = 0
	pool_MSG_PROJECT_release_getById.Put(data)
}
func (data *MSG_PROJECT_release_getById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_release_getById,buf)
	WRITE_MSG_PROJECT_release_getById(data, buf)
}

func WRITE_MSG_PROJECT_release_getById(data *MSG_PROJECT_release_getById, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
}

func READ_MSG_PROJECT_release_getById(buf *libraries.MsgBuffer) *MSG_PROJECT_release_getById {
	data := pool_MSG_PROJECT_release_getById.Get().(*MSG_PROJECT_release_getById)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_release_getById) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)

}
func (data *MSG_PROJECT_release_getById) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_release_getById) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_release_getById_result struct {
	QueryResultID uint32
	Info *MSG_PROJECT_release
}

var pool_MSG_PROJECT_release_getById_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_release_getById_result{} }}

func GET_MSG_PROJECT_release_getById_result() *MSG_PROJECT_release_getById_result {
	return pool_MSG_PROJECT_release_getById_result.Get().(*MSG_PROJECT_release_getById_result)
}

func (data *MSG_PROJECT_release_getById_result) cmd() int32 {
	return CMD_MSG_PROJECT_release_getById_result
}

func (data *MSG_PROJECT_release_getById_result) Put() {
	data.QueryResultID = 0
	if data.Info != nil {
		data.Info.Put()
		data.Info = nil
	}
	pool_MSG_PROJECT_release_getById_result.Put(data)
}
func (data *MSG_PROJECT_release_getById_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_release_getById_result,buf)
	WRITE_MSG_PROJECT_release_getById_result(data, buf)
}

func WRITE_MSG_PROJECT_release_getById_result(data *MSG_PROJECT_release_getById_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	if data.Info == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_PROJECT_release(data.Info, buf)
	}
}

func READ_MSG_PROJECT_release_getById_result(buf *libraries.MsgBuffer) *MSG_PROJECT_release_getById_result {
	data := pool_MSG_PROJECT_release_getById_result.Get().(*MSG_PROJECT_release_getById_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_release_getById_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_PROJECT_release(buf)
	}else{
		data.Info = nil
	}

}
func (data *MSG_PROJECT_release_getById_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_release_getById_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_task_getPairs struct {
	QueryID uint32
	Where map[string]interface{}
}

var pool_MSG_PROJECT_task_getPairs = sync.Pool{New: func() interface{} { return &MSG_PROJECT_task_getPairs{} }}

func GET_MSG_PROJECT_task_getPairs() *MSG_PROJECT_task_getPairs {
	return pool_MSG_PROJECT_task_getPairs.Get().(*MSG_PROJECT_task_getPairs)
}

func (data *MSG_PROJECT_task_getPairs) cmd() int32 {
	return CMD_MSG_PROJECT_task_getPairs
}

func (data *MSG_PROJECT_task_getPairs) Put() {
	data.QueryID = 0
	data.Where = nil
	pool_MSG_PROJECT_task_getPairs.Put(data)
}
func (data *MSG_PROJECT_task_getPairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_task_getPairs,buf)
	WRITE_MSG_PROJECT_task_getPairs(data, buf)
}

func WRITE_MSG_PROJECT_task_getPairs(data *MSG_PROJECT_task_getPairs, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_map(data.Where,buf)
}

func READ_MSG_PROJECT_task_getPairs(buf *libraries.MsgBuffer) *MSG_PROJECT_task_getPairs {
	data := pool_MSG_PROJECT_task_getPairs.Get().(*MSG_PROJECT_task_getPairs)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_task_getPairs) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	READ_map(&data.Where,buf)

}
func (data *MSG_PROJECT_task_getPairs) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_task_getPairs) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_task_getPairs_result struct {
	QueryResultID uint32
	List []HtmlKeyValueStr
}

var pool_MSG_PROJECT_task_getPairs_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_task_getPairs_result{} }}

func GET_MSG_PROJECT_task_getPairs_result() *MSG_PROJECT_task_getPairs_result {
	return pool_MSG_PROJECT_task_getPairs_result.Get().(*MSG_PROJECT_task_getPairs_result)
}

func (data *MSG_PROJECT_task_getPairs_result) cmd() int32 {
	return CMD_MSG_PROJECT_task_getPairs_result
}

func (data *MSG_PROJECT_task_getPairs_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	pool_MSG_PROJECT_task_getPairs_result.Put(data)
}
func (data *MSG_PROJECT_task_getPairs_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_task_getPairs_result,buf)
	WRITE_MSG_PROJECT_task_getPairs_result(data, buf)
}

func WRITE_MSG_PROJECT_task_getPairs_result(data *MSG_PROJECT_task_getPairs_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_PROJECT_task_getPairs_result(buf *libraries.MsgBuffer) *MSG_PROJECT_task_getPairs_result {
	data := pool_MSG_PROJECT_task_getPairs_result.Get().(*MSG_PROJECT_task_getPairs_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_task_getPairs_result) read(buf *libraries.MsgBuffer) {
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
func (data *MSG_PROJECT_task_getPairs_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_task_getPairs_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_task_getListByWhereMap struct {
	QueryID uint32
	Where map[string]interface{}
	Order string
	Page int
	PerPage int
	Total int
}

var pool_MSG_PROJECT_task_getListByWhereMap = sync.Pool{New: func() interface{} { return &MSG_PROJECT_task_getListByWhereMap{} }}

func GET_MSG_PROJECT_task_getListByWhereMap() *MSG_PROJECT_task_getListByWhereMap {
	return pool_MSG_PROJECT_task_getListByWhereMap.Get().(*MSG_PROJECT_task_getListByWhereMap)
}

func (data *MSG_PROJECT_task_getListByWhereMap) cmd() int32 {
	return CMD_MSG_PROJECT_task_getListByWhereMap
}

func (data *MSG_PROJECT_task_getListByWhereMap) Put() {
	data.QueryID = 0
	data.Where = nil
	data.Order = ``
	data.Page = 0
	data.PerPage = 0
	data.Total = 0
	pool_MSG_PROJECT_task_getListByWhereMap.Put(data)
}
func (data *MSG_PROJECT_task_getListByWhereMap) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_task_getListByWhereMap,buf)
	WRITE_MSG_PROJECT_task_getListByWhereMap(data, buf)
}

func WRITE_MSG_PROJECT_task_getListByWhereMap(data *MSG_PROJECT_task_getListByWhereMap, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_map(data.Where,buf)
	WRITE_string(data.Order, buf)
	WRITE_int(data.Page, buf)
	WRITE_int(data.PerPage, buf)
	WRITE_int(data.Total, buf)
}

func READ_MSG_PROJECT_task_getListByWhereMap(buf *libraries.MsgBuffer) *MSG_PROJECT_task_getListByWhereMap {
	data := pool_MSG_PROJECT_task_getListByWhereMap.Get().(*MSG_PROJECT_task_getListByWhereMap)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_task_getListByWhereMap) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	READ_map(&data.Where,buf)
	data.Order = READ_string(buf)
	data.Page = READ_int(buf)
	data.PerPage = READ_int(buf)
	data.Total = READ_int(buf)

}
func (data *MSG_PROJECT_task_getListByWhereMap) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_task_getListByWhereMap) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_task_getListByWhereMap_result struct {
	QueryResultID uint32
	List []*MSG_PROJECT_TASK
	Total int
}

var pool_MSG_PROJECT_task_getListByWhereMap_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_task_getListByWhereMap_result{} }}

func GET_MSG_PROJECT_task_getListByWhereMap_result() *MSG_PROJECT_task_getListByWhereMap_result {
	return pool_MSG_PROJECT_task_getListByWhereMap_result.Get().(*MSG_PROJECT_task_getListByWhereMap_result)
}

func (data *MSG_PROJECT_task_getListByWhereMap_result) cmd() int32 {
	return CMD_MSG_PROJECT_task_getListByWhereMap_result
}

func (data *MSG_PROJECT_task_getListByWhereMap_result) Put() {
	data.QueryResultID = 0
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	data.Total = 0
	pool_MSG_PROJECT_task_getListByWhereMap_result.Put(data)
}
func (data *MSG_PROJECT_task_getListByWhereMap_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_task_getListByWhereMap_result,buf)
	WRITE_MSG_PROJECT_task_getListByWhereMap_result(data, buf)
}

func WRITE_MSG_PROJECT_task_getListByWhereMap_result(data *MSG_PROJECT_task_getListByWhereMap_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_MSG_PROJECT_TASK(v, buf)
	}
	WRITE_int(data.Total, buf)
}

func READ_MSG_PROJECT_task_getListByWhereMap_result(buf *libraries.MsgBuffer) *MSG_PROJECT_task_getListByWhereMap_result {
	data := pool_MSG_PROJECT_task_getListByWhereMap_result.Get().(*MSG_PROJECT_task_getListByWhereMap_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_task_getListByWhereMap_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	if List_len>cap(data.List){
		data.List= make([]*MSG_PROJECT_TASK, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_PROJECT_TASK(buf)
	}
	data.Total = READ_int(buf)

}
func (data *MSG_PROJECT_task_getListByWhereMap_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_task_getListByWhereMap_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_project_getBurn struct {
	QueryID uint32
	ProjectIds []int32
}

var pool_MSG_PROJECT_project_getBurn = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_getBurn{} }}

func GET_MSG_PROJECT_project_getBurn() *MSG_PROJECT_project_getBurn {
	return pool_MSG_PROJECT_project_getBurn.Get().(*MSG_PROJECT_project_getBurn)
}

func (data *MSG_PROJECT_project_getBurn) cmd() int32 {
	return CMD_MSG_PROJECT_project_getBurn
}

func (data *MSG_PROJECT_project_getBurn) Put() {
	data.QueryID = 0
	data.ProjectIds = data.ProjectIds[:0]
	pool_MSG_PROJECT_project_getBurn.Put(data)
}
func (data *MSG_PROJECT_project_getBurn) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_getBurn,buf)
	WRITE_MSG_PROJECT_project_getBurn(data, buf)
}

func WRITE_MSG_PROJECT_project_getBurn(data *MSG_PROJECT_project_getBurn, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.ProjectIds)), buf)
	for _, v := range data.ProjectIds{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_PROJECT_project_getBurn(buf *libraries.MsgBuffer) *MSG_PROJECT_project_getBurn {
	data := pool_MSG_PROJECT_project_getBurn.Get().(*MSG_PROJECT_project_getBurn)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_getBurn) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	ProjectIds_len := int(READ_int32(buf))
	if ProjectIds_len>cap(data.ProjectIds){
		data.ProjectIds= make([]int32, ProjectIds_len)
	}else{
		data.ProjectIds = data.ProjectIds[:ProjectIds_len]
	}
	for i := 0; i < ProjectIds_len; i++ {
		data.ProjectIds[i] = READ_int32(buf)
	}

}
func (data *MSG_PROJECT_project_getBurn) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_getBurn) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_getBurn_result struct {
	QueryResultID uint32
	List []*MSG_PROJECT_project_Burn_info
}

var pool_MSG_PROJECT_project_getBurn_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_getBurn_result{} }}

func GET_MSG_PROJECT_project_getBurn_result() *MSG_PROJECT_project_getBurn_result {
	return pool_MSG_PROJECT_project_getBurn_result.Get().(*MSG_PROJECT_project_getBurn_result)
}

func (data *MSG_PROJECT_project_getBurn_result) cmd() int32 {
	return CMD_MSG_PROJECT_project_getBurn_result
}

func (data *MSG_PROJECT_project_getBurn_result) Put() {
	data.QueryResultID = 0
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_PROJECT_project_getBurn_result.Put(data)
}
func (data *MSG_PROJECT_project_getBurn_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_getBurn_result,buf)
	WRITE_MSG_PROJECT_project_getBurn_result(data, buf)
}

func WRITE_MSG_PROJECT_project_getBurn_result(data *MSG_PROJECT_project_getBurn_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_MSG_PROJECT_project_Burn_info(v, buf)
	}
}

func READ_MSG_PROJECT_project_getBurn_result(buf *libraries.MsgBuffer) *MSG_PROJECT_project_getBurn_result {
	data := pool_MSG_PROJECT_project_getBurn_result.Get().(*MSG_PROJECT_project_getBurn_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_getBurn_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	if List_len>cap(data.List){
		data.List= make([]*MSG_PROJECT_project_Burn_info, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_PROJECT_project_Burn_info(buf)
	}

}
func (data *MSG_PROJECT_project_getBurn_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_project_getBurn_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_project_Burn_info struct {
	Project int32
	Date time.Time
	Estimate float64
	Left float64
	Consumed float64
}

var pool_MSG_PROJECT_project_Burn_info = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_Burn_info{} }}

func GET_MSG_PROJECT_project_Burn_info() *MSG_PROJECT_project_Burn_info {
	return pool_MSG_PROJECT_project_Burn_info.Get().(*MSG_PROJECT_project_Burn_info)
}

func (data *MSG_PROJECT_project_Burn_info) cmd() int32 {
	return CMD_MSG_PROJECT_project_Burn_info
}

func (data *MSG_PROJECT_project_Burn_info) Put() {
	data.Project = 0
	data.Date = time.Unix(0,0)
	data.Estimate = 0
	data.Left = 0
	data.Consumed = 0
	pool_MSG_PROJECT_project_Burn_info.Put(data)
}
func (data *MSG_PROJECT_project_Burn_info) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_Burn_info,buf)
	WRITE_MSG_PROJECT_project_Burn_info(data, buf)
}

func WRITE_MSG_PROJECT_project_Burn_info(data *MSG_PROJECT_project_Burn_info, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Project, buf)
	WRITE_int64(data.Date.UnixNano(), buf)
	WRITE_float64(data.Estimate, buf)
	WRITE_float64(data.Left, buf)
	WRITE_float64(data.Consumed, buf)
}

func READ_MSG_PROJECT_project_Burn_info(buf *libraries.MsgBuffer) *MSG_PROJECT_project_Burn_info {
	data := pool_MSG_PROJECT_project_Burn_info.Get().(*MSG_PROJECT_project_Burn_info)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_Burn_info) read(buf *libraries.MsgBuffer) {
	data.Project = READ_int32(buf)
	data.Date = time.Unix(0, READ_int64(buf))
	data.Estimate = READ_float64(buf)
	data.Left = READ_float64(buf)
	data.Consumed = READ_float64(buf)

}

type MSG_PROJECT_story_getPlanStories struct {
	QueryID uint32
	PlanID int32
	Status string
	OrderBy string
}

var pool_MSG_PROJECT_story_getPlanStories = sync.Pool{New: func() interface{} { return &MSG_PROJECT_story_getPlanStories{} }}

func GET_MSG_PROJECT_story_getPlanStories() *MSG_PROJECT_story_getPlanStories {
	return pool_MSG_PROJECT_story_getPlanStories.Get().(*MSG_PROJECT_story_getPlanStories)
}

func (data *MSG_PROJECT_story_getPlanStories) cmd() int32 {
	return CMD_MSG_PROJECT_story_getPlanStories
}

func (data *MSG_PROJECT_story_getPlanStories) Put() {
	data.QueryID = 0
	data.PlanID = 0
	data.Status = ``
	data.OrderBy = ``
	pool_MSG_PROJECT_story_getPlanStories.Put(data)
}
func (data *MSG_PROJECT_story_getPlanStories) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_story_getPlanStories,buf)
	WRITE_MSG_PROJECT_story_getPlanStories(data, buf)
}

func WRITE_MSG_PROJECT_story_getPlanStories(data *MSG_PROJECT_story_getPlanStories, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.PlanID, buf)
	WRITE_string(data.Status, buf)
	WRITE_string(data.OrderBy, buf)
}

func READ_MSG_PROJECT_story_getPlanStories(buf *libraries.MsgBuffer) *MSG_PROJECT_story_getPlanStories {
	data := pool_MSG_PROJECT_story_getPlanStories.Get().(*MSG_PROJECT_story_getPlanStories)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_story_getPlanStories) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.PlanID = READ_int32(buf)
	data.Status = READ_string(buf)
	data.OrderBy = READ_string(buf)

}
func (data *MSG_PROJECT_story_getPlanStories) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_story_getPlanStories) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_story_getPlanStories_result struct {
	QueryResultID uint32
	List []*MSG_PROJECT_story
}

var pool_MSG_PROJECT_story_getPlanStories_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_story_getPlanStories_result{} }}

func GET_MSG_PROJECT_story_getPlanStories_result() *MSG_PROJECT_story_getPlanStories_result {
	return pool_MSG_PROJECT_story_getPlanStories_result.Get().(*MSG_PROJECT_story_getPlanStories_result)
}

func (data *MSG_PROJECT_story_getPlanStories_result) cmd() int32 {
	return CMD_MSG_PROJECT_story_getPlanStories_result
}

func (data *MSG_PROJECT_story_getPlanStories_result) Put() {
	data.QueryResultID = 0
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_PROJECT_story_getPlanStories_result.Put(data)
}
func (data *MSG_PROJECT_story_getPlanStories_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_story_getPlanStories_result,buf)
	WRITE_MSG_PROJECT_story_getPlanStories_result(data, buf)
}

func WRITE_MSG_PROJECT_story_getPlanStories_result(data *MSG_PROJECT_story_getPlanStories_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_MSG_PROJECT_story(v, buf)
	}
}

func READ_MSG_PROJECT_story_getPlanStories_result(buf *libraries.MsgBuffer) *MSG_PROJECT_story_getPlanStories_result {
	data := pool_MSG_PROJECT_story_getPlanStories_result.Get().(*MSG_PROJECT_story_getPlanStories_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_story_getPlanStories_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	if List_len>cap(data.List){
		data.List= make([]*MSG_PROJECT_story, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_PROJECT_story(buf)
	}

}
func (data *MSG_PROJECT_story_getPlanStories_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_story_getPlanStories_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_project_linkStory struct {
	QueryID uint32
	ProjectID int32
	Stories []int32
	Products map[int32]int32
}

var pool_MSG_PROJECT_project_linkStory = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_linkStory{} }}

func GET_MSG_PROJECT_project_linkStory() *MSG_PROJECT_project_linkStory {
	return pool_MSG_PROJECT_project_linkStory.Get().(*MSG_PROJECT_project_linkStory)
}

func (data *MSG_PROJECT_project_linkStory) cmd() int32 {
	return CMD_MSG_PROJECT_project_linkStory
}

func (data *MSG_PROJECT_project_linkStory) Put() {
	data.QueryID = 0
	data.ProjectID = 0
	data.Stories = data.Stories[:0]
	data.Products = nil
	pool_MSG_PROJECT_project_linkStory.Put(data)
}
func (data *MSG_PROJECT_project_linkStory) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_linkStory,buf)
	WRITE_MSG_PROJECT_project_linkStory(data, buf)
}

func WRITE_MSG_PROJECT_project_linkStory(data *MSG_PROJECT_project_linkStory, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.ProjectID, buf)
	WRITE_int32(int32(len(data.Stories)), buf)
	for _, v := range data.Stories{
		WRITE_int32(v, buf)
	}
	WRITE_map(data.Products,buf)
}

func READ_MSG_PROJECT_project_linkStory(buf *libraries.MsgBuffer) *MSG_PROJECT_project_linkStory {
	data := pool_MSG_PROJECT_project_linkStory.Get().(*MSG_PROJECT_project_linkStory)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_linkStory) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.ProjectID = READ_int32(buf)
	Stories_len := int(READ_int32(buf))
	if Stories_len>cap(data.Stories){
		data.Stories= make([]int32, Stories_len)
	}else{
		data.Stories = data.Stories[:Stories_len]
	}
	for i := 0; i < Stories_len; i++ {
		data.Stories[i] = READ_int32(buf)
	}
	READ_map(&data.Products,buf)

}
func (data *MSG_PROJECT_project_linkStory) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_linkStory) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_branch_getByProducts struct {
	QueryID uint32
	Products []int32
	AppendBranch []int32
}

var pool_MSG_PROJECT_branch_getByProducts = sync.Pool{New: func() interface{} { return &MSG_PROJECT_branch_getByProducts{} }}

func GET_MSG_PROJECT_branch_getByProducts() *MSG_PROJECT_branch_getByProducts {
	return pool_MSG_PROJECT_branch_getByProducts.Get().(*MSG_PROJECT_branch_getByProducts)
}

func (data *MSG_PROJECT_branch_getByProducts) cmd() int32 {
	return CMD_MSG_PROJECT_branch_getByProducts
}

func (data *MSG_PROJECT_branch_getByProducts) Put() {
	data.QueryID = 0
	data.Products = data.Products[:0]
	data.AppendBranch = data.AppendBranch[:0]
	pool_MSG_PROJECT_branch_getByProducts.Put(data)
}
func (data *MSG_PROJECT_branch_getByProducts) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_branch_getByProducts,buf)
	WRITE_MSG_PROJECT_branch_getByProducts(data, buf)
}

func WRITE_MSG_PROJECT_branch_getByProducts(data *MSG_PROJECT_branch_getByProducts, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.Products)), buf)
	for _, v := range data.Products{
		WRITE_int32(v, buf)
	}
	WRITE_int32(int32(len(data.AppendBranch)), buf)
	for _, v := range data.AppendBranch{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_PROJECT_branch_getByProducts(buf *libraries.MsgBuffer) *MSG_PROJECT_branch_getByProducts {
	data := pool_MSG_PROJECT_branch_getByProducts.Get().(*MSG_PROJECT_branch_getByProducts)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_branch_getByProducts) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	Products_len := int(READ_int32(buf))
	if Products_len>cap(data.Products){
		data.Products= make([]int32, Products_len)
	}else{
		data.Products = data.Products[:Products_len]
	}
	for i := 0; i < Products_len; i++ {
		data.Products[i] = READ_int32(buf)
	}
	AppendBranch_len := int(READ_int32(buf))
	if AppendBranch_len>cap(data.AppendBranch){
		data.AppendBranch= make([]int32, AppendBranch_len)
	}else{
		data.AppendBranch = data.AppendBranch[:AppendBranch_len]
	}
	for i := 0; i < AppendBranch_len; i++ {
		data.AppendBranch[i] = READ_int32(buf)
	}

}
func (data *MSG_PROJECT_branch_getByProducts) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_branch_getByProducts) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_branch_getByProducts_result struct {
	QueryResultID uint32
	List map[int32][]HtmlKeyValueStr
}

var pool_MSG_PROJECT_branch_getByProducts_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_branch_getByProducts_result{} }}

func GET_MSG_PROJECT_branch_getByProducts_result() *MSG_PROJECT_branch_getByProducts_result {
	return pool_MSG_PROJECT_branch_getByProducts_result.Get().(*MSG_PROJECT_branch_getByProducts_result)
}

func (data *MSG_PROJECT_branch_getByProducts_result) cmd() int32 {
	return CMD_MSG_PROJECT_branch_getByProducts_result
}

func (data *MSG_PROJECT_branch_getByProducts_result) Put() {
	data.QueryResultID = 0
	data.List = nil
	pool_MSG_PROJECT_branch_getByProducts_result.Put(data)
}
func (data *MSG_PROJECT_branch_getByProducts_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_branch_getByProducts_result,buf)
	WRITE_MSG_PROJECT_branch_getByProducts_result(data, buf)
}

func WRITE_MSG_PROJECT_branch_getByProducts_result(data *MSG_PROJECT_branch_getByProducts_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_map(data.List,buf)
}

func READ_MSG_PROJECT_branch_getByProducts_result(buf *libraries.MsgBuffer) *MSG_PROJECT_branch_getByProducts_result {
	data := pool_MSG_PROJECT_branch_getByProducts_result.Get().(*MSG_PROJECT_branch_getByProducts_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_branch_getByProducts_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	READ_map(&data.List,buf)

}
func (data *MSG_PROJECT_branch_getByProducts_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_branch_getByProducts_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_project_create struct {
	QueryID uint32
	CopyProjectID int32
	Info *MSG_PROJECT_project_cache
}

var pool_MSG_PROJECT_project_create = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_create{} }}

func GET_MSG_PROJECT_project_create() *MSG_PROJECT_project_create {
	return pool_MSG_PROJECT_project_create.Get().(*MSG_PROJECT_project_create)
}

func (data *MSG_PROJECT_project_create) cmd() int32 {
	return CMD_MSG_PROJECT_project_create
}

func (data *MSG_PROJECT_project_create) Put() {
	data.QueryID = 0
	data.CopyProjectID = 0
	if data.Info != nil {
		data.Info.Put()
		data.Info = nil
	}
	pool_MSG_PROJECT_project_create.Put(data)
}
func (data *MSG_PROJECT_project_create) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_create,buf)
	WRITE_MSG_PROJECT_project_create(data, buf)
}

func WRITE_MSG_PROJECT_project_create(data *MSG_PROJECT_project_create, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.CopyProjectID, buf)
	if data.Info == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_PROJECT_project_cache(data.Info, buf)
	}
}

func READ_MSG_PROJECT_project_create(buf *libraries.MsgBuffer) *MSG_PROJECT_project_create {
	data := pool_MSG_PROJECT_project_create.Get().(*MSG_PROJECT_project_create)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_create) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.CopyProjectID = READ_int32(buf)
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_PROJECT_project_cache(buf)
	}else{
		data.Info = nil
	}

}
func (data *MSG_PROJECT_project_create) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_create) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_create_result struct {
	QueryResultID uint32
	Id int32
}

var pool_MSG_PROJECT_project_create_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_create_result{} }}

func GET_MSG_PROJECT_project_create_result() *MSG_PROJECT_project_create_result {
	return pool_MSG_PROJECT_project_create_result.Get().(*MSG_PROJECT_project_create_result)
}

func (data *MSG_PROJECT_project_create_result) cmd() int32 {
	return CMD_MSG_PROJECT_project_create_result
}

func (data *MSG_PROJECT_project_create_result) Put() {
	data.QueryResultID = 0
	data.Id = 0
	pool_MSG_PROJECT_project_create_result.Put(data)
}
func (data *MSG_PROJECT_project_create_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_create_result,buf)
	WRITE_MSG_PROJECT_project_create_result(data, buf)
}

func WRITE_MSG_PROJECT_project_create_result(data *MSG_PROJECT_project_create_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(data.Id, buf)
}

func READ_MSG_PROJECT_project_create_result(buf *libraries.MsgBuffer) *MSG_PROJECT_project_create_result {
	data := pool_MSG_PROJECT_project_create_result.Get().(*MSG_PROJECT_project_create_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_create_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Id = READ_int32(buf)

}
func (data *MSG_PROJECT_project_create_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_project_create_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_project_statRelatedData struct {
	QueryID uint32
	ProjectID int32
}

var pool_MSG_PROJECT_project_statRelatedData = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_statRelatedData{} }}

func GET_MSG_PROJECT_project_statRelatedData() *MSG_PROJECT_project_statRelatedData {
	return pool_MSG_PROJECT_project_statRelatedData.Get().(*MSG_PROJECT_project_statRelatedData)
}

func (data *MSG_PROJECT_project_statRelatedData) cmd() int32 {
	return CMD_MSG_PROJECT_project_statRelatedData
}

func (data *MSG_PROJECT_project_statRelatedData) Put() {
	data.QueryID = 0
	data.ProjectID = 0
	pool_MSG_PROJECT_project_statRelatedData.Put(data)
}
func (data *MSG_PROJECT_project_statRelatedData) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_statRelatedData,buf)
	WRITE_MSG_PROJECT_project_statRelatedData(data, buf)
}

func WRITE_MSG_PROJECT_project_statRelatedData(data *MSG_PROJECT_project_statRelatedData, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.ProjectID, buf)
}

func READ_MSG_PROJECT_project_statRelatedData(buf *libraries.MsgBuffer) *MSG_PROJECT_project_statRelatedData {
	data := pool_MSG_PROJECT_project_statRelatedData.Get().(*MSG_PROJECT_project_statRelatedData)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_statRelatedData) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.ProjectID = READ_int32(buf)

}
func (data *MSG_PROJECT_project_statRelatedData) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_statRelatedData) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_statRelatedData_result struct {
	QueryResultID uint32
	StoryCount int
	TaskCount int
	BugCount int
}

var pool_MSG_PROJECT_project_statRelatedData_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_statRelatedData_result{} }}

func GET_MSG_PROJECT_project_statRelatedData_result() *MSG_PROJECT_project_statRelatedData_result {
	return pool_MSG_PROJECT_project_statRelatedData_result.Get().(*MSG_PROJECT_project_statRelatedData_result)
}

func (data *MSG_PROJECT_project_statRelatedData_result) cmd() int32 {
	return CMD_MSG_PROJECT_project_statRelatedData_result
}

func (data *MSG_PROJECT_project_statRelatedData_result) Put() {
	data.QueryResultID = 0
	data.StoryCount = 0
	data.TaskCount = 0
	data.BugCount = 0
	pool_MSG_PROJECT_project_statRelatedData_result.Put(data)
}
func (data *MSG_PROJECT_project_statRelatedData_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_statRelatedData_result,buf)
	WRITE_MSG_PROJECT_project_statRelatedData_result(data, buf)
}

func WRITE_MSG_PROJECT_project_statRelatedData_result(data *MSG_PROJECT_project_statRelatedData_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int(data.StoryCount, buf)
	WRITE_int(data.TaskCount, buf)
	WRITE_int(data.BugCount, buf)
}

func READ_MSG_PROJECT_project_statRelatedData_result(buf *libraries.MsgBuffer) *MSG_PROJECT_project_statRelatedData_result {
	data := pool_MSG_PROJECT_project_statRelatedData_result.Get().(*MSG_PROJECT_project_statRelatedData_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_statRelatedData_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.StoryCount = READ_int(buf)
	data.TaskCount = READ_int(buf)
	data.BugCount = READ_int(buf)

}
func (data *MSG_PROJECT_project_statRelatedData_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_project_statRelatedData_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_story_getPairsByIds struct {
	QueryID uint32
	Ids []int32
}

var pool_MSG_PROJECT_story_getPairsByIds = sync.Pool{New: func() interface{} { return &MSG_PROJECT_story_getPairsByIds{} }}

func GET_MSG_PROJECT_story_getPairsByIds() *MSG_PROJECT_story_getPairsByIds {
	return pool_MSG_PROJECT_story_getPairsByIds.Get().(*MSG_PROJECT_story_getPairsByIds)
}

func (data *MSG_PROJECT_story_getPairsByIds) cmd() int32 {
	return CMD_MSG_PROJECT_story_getPairsByIds
}

func (data *MSG_PROJECT_story_getPairsByIds) Put() {
	data.QueryID = 0
	data.Ids = data.Ids[:0]
	pool_MSG_PROJECT_story_getPairsByIds.Put(data)
}
func (data *MSG_PROJECT_story_getPairsByIds) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_story_getPairsByIds,buf)
	WRITE_MSG_PROJECT_story_getPairsByIds(data, buf)
}

func WRITE_MSG_PROJECT_story_getPairsByIds(data *MSG_PROJECT_story_getPairsByIds, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_PROJECT_story_getPairsByIds(buf *libraries.MsgBuffer) *MSG_PROJECT_story_getPairsByIds {
	data := pool_MSG_PROJECT_story_getPairsByIds.Get().(*MSG_PROJECT_story_getPairsByIds)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_story_getPairsByIds) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	Ids_len := int(READ_int32(buf))
	if Ids_len>cap(data.Ids){
		data.Ids= make([]int32, Ids_len)
	}else{
		data.Ids = data.Ids[:Ids_len]
	}
	for i := 0; i < Ids_len; i++ {
		data.Ids[i] = READ_int32(buf)
	}

}
func (data *MSG_PROJECT_story_getPairsByIds) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_story_getPairsByIds) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_story_getPairsByIds_result struct {
	QueryResultID uint32
	List []HtmlKeyValueStr
}

var pool_MSG_PROJECT_story_getPairsByIds_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_story_getPairsByIds_result{} }}

func GET_MSG_PROJECT_story_getPairsByIds_result() *MSG_PROJECT_story_getPairsByIds_result {
	return pool_MSG_PROJECT_story_getPairsByIds_result.Get().(*MSG_PROJECT_story_getPairsByIds_result)
}

func (data *MSG_PROJECT_story_getPairsByIds_result) cmd() int32 {
	return CMD_MSG_PROJECT_story_getPairsByIds_result
}

func (data *MSG_PROJECT_story_getPairsByIds_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	pool_MSG_PROJECT_story_getPairsByIds_result.Put(data)
}
func (data *MSG_PROJECT_story_getPairsByIds_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_story_getPairsByIds_result,buf)
	WRITE_MSG_PROJECT_story_getPairsByIds_result(data, buf)
}

func WRITE_MSG_PROJECT_story_getPairsByIds_result(data *MSG_PROJECT_story_getPairsByIds_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_PROJECT_story_getPairsByIds_result(buf *libraries.MsgBuffer) *MSG_PROJECT_story_getPairsByIds_result {
	data := pool_MSG_PROJECT_story_getPairsByIds_result.Get().(*MSG_PROJECT_story_getPairsByIds_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_story_getPairsByIds_result) read(buf *libraries.MsgBuffer) {
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
func (data *MSG_PROJECT_story_getPairsByIds_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_story_getPairsByIds_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_product_getPairsByIds struct {
	QueryID uint32
	Ids []int32
}

var pool_MSG_PROJECT_product_getPairsByIds = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_getPairsByIds{} }}

func GET_MSG_PROJECT_product_getPairsByIds() *MSG_PROJECT_product_getPairsByIds {
	return pool_MSG_PROJECT_product_getPairsByIds.Get().(*MSG_PROJECT_product_getPairsByIds)
}

func (data *MSG_PROJECT_product_getPairsByIds) cmd() int32 {
	return CMD_MSG_PROJECT_product_getPairsByIds
}

func (data *MSG_PROJECT_product_getPairsByIds) Put() {
	data.QueryID = 0
	data.Ids = data.Ids[:0]
	pool_MSG_PROJECT_product_getPairsByIds.Put(data)
}
func (data *MSG_PROJECT_product_getPairsByIds) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_getPairsByIds,buf)
	WRITE_MSG_PROJECT_product_getPairsByIds(data, buf)
}

func WRITE_MSG_PROJECT_product_getPairsByIds(data *MSG_PROJECT_product_getPairsByIds, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_PROJECT_product_getPairsByIds(buf *libraries.MsgBuffer) *MSG_PROJECT_product_getPairsByIds {
	data := pool_MSG_PROJECT_product_getPairsByIds.Get().(*MSG_PROJECT_product_getPairsByIds)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_getPairsByIds) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	Ids_len := int(READ_int32(buf))
	if Ids_len>cap(data.Ids){
		data.Ids= make([]int32, Ids_len)
	}else{
		data.Ids = data.Ids[:Ids_len]
	}
	for i := 0; i < Ids_len; i++ {
		data.Ids[i] = READ_int32(buf)
	}

}
func (data *MSG_PROJECT_product_getPairsByIds) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_product_getPairsByIds) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_product_getPairsByIds_result struct {
	QueryResultID uint32
	List []HtmlKeyValueStr
}

var pool_MSG_PROJECT_product_getPairsByIds_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_product_getPairsByIds_result{} }}

func GET_MSG_PROJECT_product_getPairsByIds_result() *MSG_PROJECT_product_getPairsByIds_result {
	return pool_MSG_PROJECT_product_getPairsByIds_result.Get().(*MSG_PROJECT_product_getPairsByIds_result)
}

func (data *MSG_PROJECT_product_getPairsByIds_result) cmd() int32 {
	return CMD_MSG_PROJECT_product_getPairsByIds_result
}

func (data *MSG_PROJECT_product_getPairsByIds_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	pool_MSG_PROJECT_product_getPairsByIds_result.Put(data)
}
func (data *MSG_PROJECT_product_getPairsByIds_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_product_getPairsByIds_result,buf)
	WRITE_MSG_PROJECT_product_getPairsByIds_result(data, buf)
}

func WRITE_MSG_PROJECT_product_getPairsByIds_result(data *MSG_PROJECT_product_getPairsByIds_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_PROJECT_product_getPairsByIds_result(buf *libraries.MsgBuffer) *MSG_PROJECT_product_getPairsByIds_result {
	data := pool_MSG_PROJECT_product_getPairsByIds_result.Get().(*MSG_PROJECT_product_getPairsByIds_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_product_getPairsByIds_result) read(buf *libraries.MsgBuffer) {
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
func (data *MSG_PROJECT_product_getPairsByIds_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_product_getPairsByIds_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_project_getPairsByIds struct {
	QueryID uint32
	Ids []int32
}

var pool_MSG_PROJECT_project_getPairsByIds = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_getPairsByIds{} }}

func GET_MSG_PROJECT_project_getPairsByIds() *MSG_PROJECT_project_getPairsByIds {
	return pool_MSG_PROJECT_project_getPairsByIds.Get().(*MSG_PROJECT_project_getPairsByIds)
}

func (data *MSG_PROJECT_project_getPairsByIds) cmd() int32 {
	return CMD_MSG_PROJECT_project_getPairsByIds
}

func (data *MSG_PROJECT_project_getPairsByIds) Put() {
	data.QueryID = 0
	data.Ids = data.Ids[:0]
	pool_MSG_PROJECT_project_getPairsByIds.Put(data)
}
func (data *MSG_PROJECT_project_getPairsByIds) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_getPairsByIds,buf)
	WRITE_MSG_PROJECT_project_getPairsByIds(data, buf)
}

func WRITE_MSG_PROJECT_project_getPairsByIds(data *MSG_PROJECT_project_getPairsByIds, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_PROJECT_project_getPairsByIds(buf *libraries.MsgBuffer) *MSG_PROJECT_project_getPairsByIds {
	data := pool_MSG_PROJECT_project_getPairsByIds.Get().(*MSG_PROJECT_project_getPairsByIds)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_getPairsByIds) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	Ids_len := int(READ_int32(buf))
	if Ids_len>cap(data.Ids){
		data.Ids= make([]int32, Ids_len)
	}else{
		data.Ids = data.Ids[:Ids_len]
	}
	for i := 0; i < Ids_len; i++ {
		data.Ids[i] = READ_int32(buf)
	}

}
func (data *MSG_PROJECT_project_getPairsByIds) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_getPairsByIds) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_getPairsByIds_result struct {
	QueryResultID uint32
	List []HtmlKeyValueStr
}

var pool_MSG_PROJECT_project_getPairsByIds_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_getPairsByIds_result{} }}

func GET_MSG_PROJECT_project_getPairsByIds_result() *MSG_PROJECT_project_getPairsByIds_result {
	return pool_MSG_PROJECT_project_getPairsByIds_result.Get().(*MSG_PROJECT_project_getPairsByIds_result)
}

func (data *MSG_PROJECT_project_getPairsByIds_result) cmd() int32 {
	return CMD_MSG_PROJECT_project_getPairsByIds_result
}

func (data *MSG_PROJECT_project_getPairsByIds_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	pool_MSG_PROJECT_project_getPairsByIds_result.Put(data)
}
func (data *MSG_PROJECT_project_getPairsByIds_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_getPairsByIds_result,buf)
	WRITE_MSG_PROJECT_project_getPairsByIds_result(data, buf)
}

func WRITE_MSG_PROJECT_project_getPairsByIds_result(data *MSG_PROJECT_project_getPairsByIds_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_PROJECT_project_getPairsByIds_result(buf *libraries.MsgBuffer) *MSG_PROJECT_project_getPairsByIds_result {
	data := pool_MSG_PROJECT_project_getPairsByIds_result.Get().(*MSG_PROJECT_project_getPairsByIds_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_getPairsByIds_result) read(buf *libraries.MsgBuffer) {
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
func (data *MSG_PROJECT_project_getPairsByIds_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_project_getPairsByIds_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_branch_getPairsByIds struct {
	QueryID uint32
	Ids []int32
}

var pool_MSG_PROJECT_branch_getPairsByIds = sync.Pool{New: func() interface{} { return &MSG_PROJECT_branch_getPairsByIds{} }}

func GET_MSG_PROJECT_branch_getPairsByIds() *MSG_PROJECT_branch_getPairsByIds {
	return pool_MSG_PROJECT_branch_getPairsByIds.Get().(*MSG_PROJECT_branch_getPairsByIds)
}

func (data *MSG_PROJECT_branch_getPairsByIds) cmd() int32 {
	return CMD_MSG_PROJECT_branch_getPairsByIds
}

func (data *MSG_PROJECT_branch_getPairsByIds) Put() {
	data.QueryID = 0
	data.Ids = data.Ids[:0]
	pool_MSG_PROJECT_branch_getPairsByIds.Put(data)
}
func (data *MSG_PROJECT_branch_getPairsByIds) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_branch_getPairsByIds,buf)
	WRITE_MSG_PROJECT_branch_getPairsByIds(data, buf)
}

func WRITE_MSG_PROJECT_branch_getPairsByIds(data *MSG_PROJECT_branch_getPairsByIds, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_PROJECT_branch_getPairsByIds(buf *libraries.MsgBuffer) *MSG_PROJECT_branch_getPairsByIds {
	data := pool_MSG_PROJECT_branch_getPairsByIds.Get().(*MSG_PROJECT_branch_getPairsByIds)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_branch_getPairsByIds) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	Ids_len := int(READ_int32(buf))
	if Ids_len>cap(data.Ids){
		data.Ids= make([]int32, Ids_len)
	}else{
		data.Ids = data.Ids[:Ids_len]
	}
	for i := 0; i < Ids_len; i++ {
		data.Ids[i] = READ_int32(buf)
	}

}
func (data *MSG_PROJECT_branch_getPairsByIds) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_branch_getPairsByIds) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_branch_getPairsByIds_result struct {
	QueryResultID uint32
	List []HtmlKeyValueStr
}

var pool_MSG_PROJECT_branch_getPairsByIds_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_branch_getPairsByIds_result{} }}

func GET_MSG_PROJECT_branch_getPairsByIds_result() *MSG_PROJECT_branch_getPairsByIds_result {
	return pool_MSG_PROJECT_branch_getPairsByIds_result.Get().(*MSG_PROJECT_branch_getPairsByIds_result)
}

func (data *MSG_PROJECT_branch_getPairsByIds_result) cmd() int32 {
	return CMD_MSG_PROJECT_branch_getPairsByIds_result
}

func (data *MSG_PROJECT_branch_getPairsByIds_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	pool_MSG_PROJECT_branch_getPairsByIds_result.Put(data)
}
func (data *MSG_PROJECT_branch_getPairsByIds_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_branch_getPairsByIds_result,buf)
	WRITE_MSG_PROJECT_branch_getPairsByIds_result(data, buf)
}

func WRITE_MSG_PROJECT_branch_getPairsByIds_result(data *MSG_PROJECT_branch_getPairsByIds_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_PROJECT_branch_getPairsByIds_result(buf *libraries.MsgBuffer) *MSG_PROJECT_branch_getPairsByIds_result {
	data := pool_MSG_PROJECT_branch_getPairsByIds_result.Get().(*MSG_PROJECT_branch_getPairsByIds_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_branch_getPairsByIds_result) read(buf *libraries.MsgBuffer) {
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
func (data *MSG_PROJECT_branch_getPairsByIds_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_branch_getPairsByIds_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_tree_getPairsByIds struct {
	QueryID uint32
	Ids []int32
}

var pool_MSG_PROJECT_tree_getPairsByIds = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_getPairsByIds{} }}

func GET_MSG_PROJECT_tree_getPairsByIds() *MSG_PROJECT_tree_getPairsByIds {
	return pool_MSG_PROJECT_tree_getPairsByIds.Get().(*MSG_PROJECT_tree_getPairsByIds)
}

func (data *MSG_PROJECT_tree_getPairsByIds) cmd() int32 {
	return CMD_MSG_PROJECT_tree_getPairsByIds
}

func (data *MSG_PROJECT_tree_getPairsByIds) Put() {
	data.QueryID = 0
	data.Ids = data.Ids[:0]
	pool_MSG_PROJECT_tree_getPairsByIds.Put(data)
}
func (data *MSG_PROJECT_tree_getPairsByIds) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_getPairsByIds,buf)
	WRITE_MSG_PROJECT_tree_getPairsByIds(data, buf)
}

func WRITE_MSG_PROJECT_tree_getPairsByIds(data *MSG_PROJECT_tree_getPairsByIds, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_PROJECT_tree_getPairsByIds(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_getPairsByIds {
	data := pool_MSG_PROJECT_tree_getPairsByIds.Get().(*MSG_PROJECT_tree_getPairsByIds)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_getPairsByIds) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	Ids_len := int(READ_int32(buf))
	if Ids_len>cap(data.Ids){
		data.Ids= make([]int32, Ids_len)
	}else{
		data.Ids = data.Ids[:Ids_len]
	}
	for i := 0; i < Ids_len; i++ {
		data.Ids[i] = READ_int32(buf)
	}

}
func (data *MSG_PROJECT_tree_getPairsByIds) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_tree_getPairsByIds) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_tree_getPairsByIds_result struct {
	QueryResultID uint32
	List []HtmlKeyValueStr
}

var pool_MSG_PROJECT_tree_getPairsByIds_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_getPairsByIds_result{} }}

func GET_MSG_PROJECT_tree_getPairsByIds_result() *MSG_PROJECT_tree_getPairsByIds_result {
	return pool_MSG_PROJECT_tree_getPairsByIds_result.Get().(*MSG_PROJECT_tree_getPairsByIds_result)
}

func (data *MSG_PROJECT_tree_getPairsByIds_result) cmd() int32 {
	return CMD_MSG_PROJECT_tree_getPairsByIds_result
}

func (data *MSG_PROJECT_tree_getPairsByIds_result) Put() {
	data.QueryResultID = 0
	data.List = data.List[:0]
	pool_MSG_PROJECT_tree_getPairsByIds_result.Put(data)
}
func (data *MSG_PROJECT_tree_getPairsByIds_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_getPairsByIds_result,buf)
	WRITE_MSG_PROJECT_tree_getPairsByIds_result(data, buf)
}

func WRITE_MSG_PROJECT_tree_getPairsByIds_result(data *MSG_PROJECT_tree_getPairsByIds_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_PROJECT_tree_getPairsByIds_result(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_getPairsByIds_result {
	data := pool_MSG_PROJECT_tree_getPairsByIds_result.Get().(*MSG_PROJECT_tree_getPairsByIds_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_getPairsByIds_result) read(buf *libraries.MsgBuffer) {
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
func (data *MSG_PROJECT_tree_getPairsByIds_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_tree_getPairsByIds_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_project_start struct {
	QueryID uint32
	Id int32
	Comment string
}

var pool_MSG_PROJECT_project_start = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_start{} }}

func GET_MSG_PROJECT_project_start() *MSG_PROJECT_project_start {
	return pool_MSG_PROJECT_project_start.Get().(*MSG_PROJECT_project_start)
}

func (data *MSG_PROJECT_project_start) cmd() int32 {
	return CMD_MSG_PROJECT_project_start
}

func (data *MSG_PROJECT_project_start) Put() {
	data.QueryID = 0
	data.Id = 0
	data.Comment = ``
	pool_MSG_PROJECT_project_start.Put(data)
}
func (data *MSG_PROJECT_project_start) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_start,buf)
	WRITE_MSG_PROJECT_project_start(data, buf)
}

func WRITE_MSG_PROJECT_project_start(data *MSG_PROJECT_project_start, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
	WRITE_string(data.Comment, buf)
}

func READ_MSG_PROJECT_project_start(buf *libraries.MsgBuffer) *MSG_PROJECT_project_start {
	data := pool_MSG_PROJECT_project_start.Get().(*MSG_PROJECT_project_start)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_start) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)
	data.Comment = READ_string(buf)

}
func (data *MSG_PROJECT_project_start) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_start) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_putoff struct {
	QueryID uint32
	Id int32
	Begin time.Time
	End time.Time
	Days int16
	Comment string
}

var pool_MSG_PROJECT_project_putoff = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_putoff{} }}

func GET_MSG_PROJECT_project_putoff() *MSG_PROJECT_project_putoff {
	return pool_MSG_PROJECT_project_putoff.Get().(*MSG_PROJECT_project_putoff)
}

func (data *MSG_PROJECT_project_putoff) cmd() int32 {
	return CMD_MSG_PROJECT_project_putoff
}

func (data *MSG_PROJECT_project_putoff) Put() {
	data.QueryID = 0
	data.Id = 0
	data.Begin = time.Unix(0,0)
	data.End = time.Unix(0,0)
	data.Days = 0
	data.Comment = ``
	pool_MSG_PROJECT_project_putoff.Put(data)
}
func (data *MSG_PROJECT_project_putoff) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_putoff,buf)
	WRITE_MSG_PROJECT_project_putoff(data, buf)
}

func WRITE_MSG_PROJECT_project_putoff(data *MSG_PROJECT_project_putoff, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
	WRITE_int64(data.Begin.UnixNano(), buf)
	WRITE_int64(data.End.UnixNano(), buf)
	WRITE_int16(data.Days, buf)
	WRITE_string(data.Comment, buf)
}

func READ_MSG_PROJECT_project_putoff(buf *libraries.MsgBuffer) *MSG_PROJECT_project_putoff {
	data := pool_MSG_PROJECT_project_putoff.Get().(*MSG_PROJECT_project_putoff)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_putoff) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)
	data.Begin = time.Unix(0, READ_int64(buf))
	data.End = time.Unix(0, READ_int64(buf))
	data.Days = READ_int16(buf)
	data.Comment = READ_string(buf)

}
func (data *MSG_PROJECT_project_putoff) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_putoff) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_suspend struct {
	QueryID uint32
	Id int32
	Comment string
}

var pool_MSG_PROJECT_project_suspend = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_suspend{} }}

func GET_MSG_PROJECT_project_suspend() *MSG_PROJECT_project_suspend {
	return pool_MSG_PROJECT_project_suspend.Get().(*MSG_PROJECT_project_suspend)
}

func (data *MSG_PROJECT_project_suspend) cmd() int32 {
	return CMD_MSG_PROJECT_project_suspend
}

func (data *MSG_PROJECT_project_suspend) Put() {
	data.QueryID = 0
	data.Id = 0
	data.Comment = ``
	pool_MSG_PROJECT_project_suspend.Put(data)
}
func (data *MSG_PROJECT_project_suspend) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_suspend,buf)
	WRITE_MSG_PROJECT_project_suspend(data, buf)
}

func WRITE_MSG_PROJECT_project_suspend(data *MSG_PROJECT_project_suspend, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
	WRITE_string(data.Comment, buf)
}

func READ_MSG_PROJECT_project_suspend(buf *libraries.MsgBuffer) *MSG_PROJECT_project_suspend {
	data := pool_MSG_PROJECT_project_suspend.Get().(*MSG_PROJECT_project_suspend)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_suspend) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)
	data.Comment = READ_string(buf)

}
func (data *MSG_PROJECT_project_suspend) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_suspend) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_activate struct {
	QueryID uint32
	Id int32
	Begin time.Time
	End time.Time
	Comment string
	ReadjustTask bool
}

var pool_MSG_PROJECT_project_activate = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_activate{} }}

func GET_MSG_PROJECT_project_activate() *MSG_PROJECT_project_activate {
	return pool_MSG_PROJECT_project_activate.Get().(*MSG_PROJECT_project_activate)
}

func (data *MSG_PROJECT_project_activate) cmd() int32 {
	return CMD_MSG_PROJECT_project_activate
}

func (data *MSG_PROJECT_project_activate) Put() {
	data.QueryID = 0
	data.Id = 0
	data.Begin = time.Unix(0,0)
	data.End = time.Unix(0,0)
	data.Comment = ``
	data.ReadjustTask = false
	pool_MSG_PROJECT_project_activate.Put(data)
}
func (data *MSG_PROJECT_project_activate) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_activate,buf)
	WRITE_MSG_PROJECT_project_activate(data, buf)
}

func WRITE_MSG_PROJECT_project_activate(data *MSG_PROJECT_project_activate, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
	WRITE_int64(data.Begin.UnixNano(), buf)
	WRITE_int64(data.End.UnixNano(), buf)
	WRITE_string(data.Comment, buf)
	WRITE_bool(data.ReadjustTask, buf)
}

func READ_MSG_PROJECT_project_activate(buf *libraries.MsgBuffer) *MSG_PROJECT_project_activate {
	data := pool_MSG_PROJECT_project_activate.Get().(*MSG_PROJECT_project_activate)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_activate) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)
	data.Begin = time.Unix(0, READ_int64(buf))
	data.End = time.Unix(0, READ_int64(buf))
	data.Comment = READ_string(buf)
	data.ReadjustTask = READ_bool(buf)

}
func (data *MSG_PROJECT_project_activate) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_activate) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_close struct {
	QueryID uint32
	Id int32
	Comment string
}

var pool_MSG_PROJECT_project_close = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_close{} }}

func GET_MSG_PROJECT_project_close() *MSG_PROJECT_project_close {
	return pool_MSG_PROJECT_project_close.Get().(*MSG_PROJECT_project_close)
}

func (data *MSG_PROJECT_project_close) cmd() int32 {
	return CMD_MSG_PROJECT_project_close
}

func (data *MSG_PROJECT_project_close) Put() {
	data.QueryID = 0
	data.Id = 0
	data.Comment = ``
	pool_MSG_PROJECT_project_close.Put(data)
}
func (data *MSG_PROJECT_project_close) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_close,buf)
	WRITE_MSG_PROJECT_project_close(data, buf)
}

func WRITE_MSG_PROJECT_project_close(data *MSG_PROJECT_project_close, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
	WRITE_string(data.Comment, buf)
}

func READ_MSG_PROJECT_project_close(buf *libraries.MsgBuffer) *MSG_PROJECT_project_close {
	data := pool_MSG_PROJECT_project_close.Get().(*MSG_PROJECT_project_close)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_close) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)
	data.Comment = READ_string(buf)

}
func (data *MSG_PROJECT_project_close) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_close) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_delete struct {
	QueryID uint32
	Id int32
}

var pool_MSG_PROJECT_project_delete = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_delete{} }}

func GET_MSG_PROJECT_project_delete() *MSG_PROJECT_project_delete {
	return pool_MSG_PROJECT_project_delete.Get().(*MSG_PROJECT_project_delete)
}

func (data *MSG_PROJECT_project_delete) cmd() int32 {
	return CMD_MSG_PROJECT_project_delete
}

func (data *MSG_PROJECT_project_delete) Put() {
	data.QueryID = 0
	data.Id = 0
	pool_MSG_PROJECT_project_delete.Put(data)
}
func (data *MSG_PROJECT_project_delete) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_delete,buf)
	WRITE_MSG_PROJECT_project_delete(data, buf)
}

func WRITE_MSG_PROJECT_project_delete(data *MSG_PROJECT_project_delete, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
}

func READ_MSG_PROJECT_project_delete(buf *libraries.MsgBuffer) *MSG_PROJECT_project_delete {
	data := pool_MSG_PROJECT_project_delete.Get().(*MSG_PROJECT_project_delete)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_delete) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)

}
func (data *MSG_PROJECT_project_delete) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_delete) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_getProjectTasks struct {
	QueryID uint32
	ProjectID int32
	ProductID int32
	Type []string
	ModuleID int32
	OrderBy string
	Role string
	Page int
	PerPage int
	Total int
}

var pool_MSG_PROJECT_project_getProjectTasks = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_getProjectTasks{} }}

func GET_MSG_PROJECT_project_getProjectTasks() *MSG_PROJECT_project_getProjectTasks {
	return pool_MSG_PROJECT_project_getProjectTasks.Get().(*MSG_PROJECT_project_getProjectTasks)
}

func (data *MSG_PROJECT_project_getProjectTasks) cmd() int32 {
	return CMD_MSG_PROJECT_project_getProjectTasks
}

func (data *MSG_PROJECT_project_getProjectTasks) Put() {
	data.QueryID = 0
	data.ProjectID = 0
	data.ProductID = 0
	data.Type = data.Type[:0]
	data.ModuleID = 0
	data.OrderBy = ``
	data.Role = ``
	data.Page = 0
	data.PerPage = 0
	data.Total = 0
	pool_MSG_PROJECT_project_getProjectTasks.Put(data)
}
func (data *MSG_PROJECT_project_getProjectTasks) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_getProjectTasks,buf)
	WRITE_MSG_PROJECT_project_getProjectTasks(data, buf)
}

func WRITE_MSG_PROJECT_project_getProjectTasks(data *MSG_PROJECT_project_getProjectTasks, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.ProjectID, buf)
	WRITE_int32(data.ProductID, buf)
	WRITE_int32(int32(len(data.Type)), buf)
	for _, v := range data.Type{
		WRITE_string(v, buf)
	}
	WRITE_int32(data.ModuleID, buf)
	WRITE_string(data.OrderBy, buf)
	WRITE_string(data.Role, buf)
	WRITE_int(data.Page, buf)
	WRITE_int(data.PerPage, buf)
	WRITE_int(data.Total, buf)
}

func READ_MSG_PROJECT_project_getProjectTasks(buf *libraries.MsgBuffer) *MSG_PROJECT_project_getProjectTasks {
	data := pool_MSG_PROJECT_project_getProjectTasks.Get().(*MSG_PROJECT_project_getProjectTasks)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_getProjectTasks) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.ProjectID = READ_int32(buf)
	data.ProductID = READ_int32(buf)
	Type_len := int(READ_int32(buf))
	if Type_len>cap(data.Type){
		data.Type= make([]string, Type_len)
	}else{
		data.Type = data.Type[:Type_len]
	}
	for i := 0; i < Type_len; i++ {
		data.Type[i] = READ_string(buf)
	}
	data.ModuleID = READ_int32(buf)
	data.OrderBy = READ_string(buf)
	data.Role = READ_string(buf)
	data.Page = READ_int(buf)
	data.PerPage = READ_int(buf)
	data.Total = READ_int(buf)

}
func (data *MSG_PROJECT_project_getProjectTasks) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_project_getProjectTasks) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_project_getProjectTasks_result struct {
	QueryResultID uint32
	List []*MSG_PROJECT_TASK
	Total int
}

var pool_MSG_PROJECT_project_getProjectTasks_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_project_getProjectTasks_result{} }}

func GET_MSG_PROJECT_project_getProjectTasks_result() *MSG_PROJECT_project_getProjectTasks_result {
	return pool_MSG_PROJECT_project_getProjectTasks_result.Get().(*MSG_PROJECT_project_getProjectTasks_result)
}

func (data *MSG_PROJECT_project_getProjectTasks_result) cmd() int32 {
	return CMD_MSG_PROJECT_project_getProjectTasks_result
}

func (data *MSG_PROJECT_project_getProjectTasks_result) Put() {
	data.QueryResultID = 0
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	data.Total = 0
	pool_MSG_PROJECT_project_getProjectTasks_result.Put(data)
}
func (data *MSG_PROJECT_project_getProjectTasks_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_project_getProjectTasks_result,buf)
	WRITE_MSG_PROJECT_project_getProjectTasks_result(data, buf)
}

func WRITE_MSG_PROJECT_project_getProjectTasks_result(data *MSG_PROJECT_project_getProjectTasks_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_MSG_PROJECT_TASK(v, buf)
	}
	WRITE_int(data.Total, buf)
}

func READ_MSG_PROJECT_project_getProjectTasks_result(buf *libraries.MsgBuffer) *MSG_PROJECT_project_getProjectTasks_result {
	data := pool_MSG_PROJECT_project_getProjectTasks_result.Get().(*MSG_PROJECT_project_getProjectTasks_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_project_getProjectTasks_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	if List_len>cap(data.List){
		data.List= make([]*MSG_PROJECT_TASK, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_PROJECT_TASK(buf)
	}
	data.Total = READ_int(buf)

}
func (data *MSG_PROJECT_project_getProjectTasks_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_project_getProjectTasks_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_PROJECT_tree_getTaskTreeModules struct {
	QueryID uint32
	ProjectID int32
	Parent bool
}

var pool_MSG_PROJECT_tree_getTaskTreeModules = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_getTaskTreeModules{} }}

func GET_MSG_PROJECT_tree_getTaskTreeModules() *MSG_PROJECT_tree_getTaskTreeModules {
	return pool_MSG_PROJECT_tree_getTaskTreeModules.Get().(*MSG_PROJECT_tree_getTaskTreeModules)
}

func (data *MSG_PROJECT_tree_getTaskTreeModules) cmd() int32 {
	return CMD_MSG_PROJECT_tree_getTaskTreeModules
}

func (data *MSG_PROJECT_tree_getTaskTreeModules) Put() {
	data.QueryID = 0
	data.ProjectID = 0
	data.Parent = false
	pool_MSG_PROJECT_tree_getTaskTreeModules.Put(data)
}
func (data *MSG_PROJECT_tree_getTaskTreeModules) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_getTaskTreeModules,buf)
	WRITE_MSG_PROJECT_tree_getTaskTreeModules(data, buf)
}

func WRITE_MSG_PROJECT_tree_getTaskTreeModules(data *MSG_PROJECT_tree_getTaskTreeModules, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.ProjectID, buf)
	WRITE_bool(data.Parent, buf)
}

func READ_MSG_PROJECT_tree_getTaskTreeModules(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_getTaskTreeModules {
	data := pool_MSG_PROJECT_tree_getTaskTreeModules.Get().(*MSG_PROJECT_tree_getTaskTreeModules)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_getTaskTreeModules) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.ProjectID = READ_int32(buf)
	data.Parent = READ_bool(buf)

}
func (data *MSG_PROJECT_tree_getTaskTreeModules) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_PROJECT_tree_getTaskTreeModules) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_PROJECT_tree_getTaskTreeModules_result struct {
	QueryResultID uint32
	ProjectModules map[int32]int32
}

var pool_MSG_PROJECT_tree_getTaskTreeModules_result = sync.Pool{New: func() interface{} { return &MSG_PROJECT_tree_getTaskTreeModules_result{} }}

func GET_MSG_PROJECT_tree_getTaskTreeModules_result() *MSG_PROJECT_tree_getTaskTreeModules_result {
	return pool_MSG_PROJECT_tree_getTaskTreeModules_result.Get().(*MSG_PROJECT_tree_getTaskTreeModules_result)
}

func (data *MSG_PROJECT_tree_getTaskTreeModules_result) cmd() int32 {
	return CMD_MSG_PROJECT_tree_getTaskTreeModules_result
}

func (data *MSG_PROJECT_tree_getTaskTreeModules_result) Put() {
	data.QueryResultID = 0
	data.ProjectModules = nil
	pool_MSG_PROJECT_tree_getTaskTreeModules_result.Put(data)
}
func (data *MSG_PROJECT_tree_getTaskTreeModules_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_PROJECT_tree_getTaskTreeModules_result,buf)
	WRITE_MSG_PROJECT_tree_getTaskTreeModules_result(data, buf)
}

func WRITE_MSG_PROJECT_tree_getTaskTreeModules_result(data *MSG_PROJECT_tree_getTaskTreeModules_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_map(data.ProjectModules,buf)
}

func READ_MSG_PROJECT_tree_getTaskTreeModules_result(buf *libraries.MsgBuffer) *MSG_PROJECT_tree_getTaskTreeModules_result {
	data := pool_MSG_PROJECT_tree_getTaskTreeModules_result.Get().(*MSG_PROJECT_tree_getTaskTreeModules_result)
	data.read(buf)
	return data
}

func (data *MSG_PROJECT_tree_getTaskTreeModules_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	READ_map(&data.ProjectModules,buf)

}
func (data *MSG_PROJECT_tree_getTaskTreeModules_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_PROJECT_tree_getTaskTreeModules_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

