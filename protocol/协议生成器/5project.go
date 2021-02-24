package main

import (
	"time"
)

type MSG_PROJECT_tree_getLinePairs struct {
	QueryID uint32
}
type MSG_PROJECT_tree_getLinePairs_result struct {
	QueryResultID uint32
	List          []HtmlKeyValueStr
}
type MSG_PROJECT_product_cache struct {
	Id          int32
	Name        string
	Code        string
	Line        int32
	Type        string
	Status      string
	Desc        string
	PO          int32
	QD          int32
	RD          int32
	Acl         string
	Whitelist   []int32
	CreatedBy   int32
	CreatedDate int64
	Order       int32
	Deleted     bool
	TimeStamp   int64
	Branchs     []*MSG_PROJECT_branch_info
}
type MSG_PROJECT_product_insert struct {
	QueryID uint32
	DocName string
	Data    *MSG_PROJECT_product_cache
}
type MSG_PROJECT_product_insert_result struct {
	QueryResultID uint32
	ID            int32
}
type MSG_PROJECT_product_update struct {
	QueryID uint32
	Data    *MSG_PROJECT_product_cache
}
type MSG_PROJECT_product_getStories struct {
	QueryID    uint32
	ProductID  int32
	Branch     int32
	BrowseType string
	ModuleID   int32
	Sort       string
	Page       int
	PerPage    int
	Where      map[string]interface{}
	Total      int
}
type MSG_PROJECT_product_getStories_result struct {
	QueryResultID uint32
	List          []*MSG_PROJECT_story
	Total         int
}
type MSG_PROJECT_story struct {
	Id             int32
	Product        int32
	Branch         int32
	Module         int32
	Plan           int32
	Source         string
	SourceNote     string
	FromBug        int32
	Title          string
	Keywords       string
	Pri            int8
	Estimate       float32
	Status         string
	Stage          string
	Mailto         []int32
	OpenedBy       int32
	OpenedDate     time.Time
	AssignedTo     int32
	AssignedDate   time.Time
	LastEditedBy   int32
	LastEditedDate time.Time
	ReviewedBy     int32
	ReviewedDate   time.Time
	ClosedBy       int32
	ClosedDate     time.Time
	ClosedReason   string
	ToBug          int32
	ChildStories   string
	LinkStories    string
	DuplicateStory int32
	Deleted        bool
	Version        int16
	Color          string
	PlanTitle      string `db:-`
}
type MSG_PROJECT_tree_cache struct {
	Id        int32
	Root      int32
	Branch    int32
	Name      string
	Parent    int32
	Path      []int32
	Grade     int8
	Order     int16
	Type      string
	Owner     string
	OwnerID   int32
	Collector string
	Short     string
	Deleted   bool
	TimeStamp int64
}
type MSG_PROJECT_tree_getParents struct {
	QueryID  uint32
	ModuleID int32
}
type MSG_PROJECT_tree_getParents_result struct {
	QueryResultID uint32
	List          []*MSG_PROJECT_tree_cache
}
type MSG_PROJECT_branch_info struct {
	Id      int32
	Product int32
	Name    string
	Order   int16
	Deleted bool
}
type MSG_PROJECT_tree_manageChild struct {
	QueryID        uint32
	RootID         int32
	ViewType       string
	Modules        []*MSG_PROJECT_tree_cache
	ParentModuleID int32
}
type MSG_PROJECT_tree_manageChild_result struct {
	QueryResultID uint32
	Result        ErrCode
	Name          string
}
type MSG_PROJECT_product_getStoriesMapBySql struct {
	QueryID uint32
	Field   string
	Where   map[string]interface{}
	Order   string
	Group   string
	Page    int
	PerPage int
	Total   int
}
type MSG_PROJECT_product_getStoriesMapBySql_result struct {
	QueryResultID uint32
	List          []map[string]string
	Total         int
}
type MSG_PROJECT_tree_updateList struct {
	QueryID uint32
	Modules []*MSG_PROJECT_tree_cache
}
type MSG_PROJECT_tree_delete struct {
	QueryID uint32
	Ids     []int32
}
type MSG_PROJECT_productplan_getPairsForStory struct {
	QueryID uint32
	Product int32
	Branch  int32
}
type MSG_PROJECT_productplan_getPairsForStory_result struct {
	QueryResultID uint32
	List          []HtmlKeyValueStr
}
type MSG_PROJECT_productplan_getList struct {
	QueryID    uint32
	Id         int32
	ProductID  int32
	Branch     int32
	BrowseType string
	Order      string
	Page       int
	PerPage    int
	Total      int
}
type MSG_PROJECT_productplan_getList_result struct {
	QueryResultID uint32
	List          []map[string]string
	Total         int
}
type MSG_PROJECT_productplan_getLast struct {
	QueryID   uint32
	ProductId int32
	Branch    int32
}
type MSG_PROJECT_productplan_getLast_result struct {
	QueryResultID uint32
	Result        map[string]string
}
type MSG_PROJECT_product_editBranch struct {
	QueryID   uint32
	ProductID int32
	Branchs   []*MSG_PROJECT_branch_info
}
type MSG_PROJECT_product_deleteBranch struct {
	QueryID   uint32
	ProductID int32
	BranchID  int32
}
type MSG_PROJECT_product_deleteBranch_result struct {
	QueryResultID uint32
	Result        ErrCode
}
type MSG_PROJECT_productplan_getPairs struct {
	QueryID   uint32
	ProductID int32
	BranchID  int32
	Expired   string
}
type MSG_PROJECT_productplan_getPairs_result struct {
	QueryResultID uint32
	List          []HtmlKeyValueStr
}
type MSG_PROJECT_productplan_insertUpdate struct {
	QueryID  uint32 `db:"-"`
	Id       int32  `db:"pk"`
	Product  int32
	Branch   int32
	Parent   int32
	Projects []int32
	Title    string
	Desc     string
	Begin    time.Time
	End      time.Time
	Order    string
	Deleted  bool
}

type MSG_PROJECT_productplan_insertUpdate_result struct {
	QueryResultID uint32
	Id            int32
	Result        ErrCode
}
type MSG_PROJECT_productplan_delete struct {
	QueryID uint32
	Id      int32
	Product int32
	Branch  int32
}
type MSG_PROJECT_stroy_create struct {
	QueryID       uint32
	Product       int32
	Branch        int32
	Module        int32
	Plan          int32
	Source        string //来源
	SourceNote    string //来源备注
	AssignedTo    int32  //评审
	Title         string
	Color         string
	Pri           int8
	Estimate      float32
	Spec          string
	Verify        string
	Mailto        []int32
	Keywords      string
	NeedNotReview bool
	FromBug       int32
	OpenedBy      int32
	ProjectID     int32
}
type MSG_PROJECT_stroy_create_result struct {
	QueryResultID uint32
	Result        int32 //小于0为ErrCode,大于0位新增ID
}
type MSG_PROJECT_story_batchGetStoryStage struct {
	QueryID uint32
	Ids     []int32
}
type MSG_PROJECT_story_batchGetStoryStage_result struct {
	QueryResultID uint32
	List          map[int32][]HtmlKeyValueStr
}
