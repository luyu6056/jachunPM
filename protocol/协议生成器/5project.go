package main

import (
	"time"
)

type MSG_PROJECT_tree_getLinePairs struct {
}
type MSG_PROJECT_tree_getLinePairs_result struct {
	List []HtmlKeyValueStr
}
type MSG_PROJECT_product_cache struct {
	Id          int32
	Name        string
	Code        string
	Line        int32
	Type        string
	Status      string
	Desc        string
	Branch      []int32
	Plan        []int32
	PO          int32
	QD          int32
	RD          int32
	Acl         string
	Whitelist   []int32
	CreatedBy   int32
	CreatedDate time.Time
	Order       int32
	Deleted     bool
	TimeStamp   time.Time
	Branchs     []*MSG_PROJECT_branch_info `db:"-"`
}
type MSG_PROJECT_product_insert struct {
	DocName string
	Data    *MSG_PROJECT_product_cache
}
type MSG_PROJECT_product_insert_result struct {
	ID int32
}
type MSG_PROJECT_product_update struct {
	Data *MSG_PROJECT_product_cache
}
type MSG_PROJECT_product_getStories struct {
	ProductID  int32
	Branch     int32
	BrowseType string
	ModuleID   int32
	Sort       string
	Uid        int32
	Where      map[string]interface{}
	Page       int
	PerPage    int
	Total      int
}
type MSG_PROJECT_product_getStories_result struct {
	List  []*MSG_PROJECT_story
	Total int
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
	Estimate       float64
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
	ChildStories   []int32
	LinkStories    []int32
	DuplicateStory int32
	Deleted        bool
	Version        int16
	Color          string
	PlanTitle      string                    `db:-`
	Spec           string                    `db:-`
	Verify         string                    `db:-`
	Stages         []*MSG_PROJECT_StoryStage `db:-`
	ExtraStories   []*MSG_PROJECT_story      `db:-`
	Projects       []int32                   `db:-` //具体信息从缓存读取
	Tasks          []*MSG_PROJECT_TASK       `db:-`
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
	TimeStamp time.Time
}
type MSG_PROJECT_tree_getParents struct {
	ModuleID int32
}
type MSG_PROJECT_tree_getParents_result struct {
	List []*MSG_PROJECT_tree_cache
}
type MSG_PROJECT_branch_info struct {
	Id      int32
	Product int32
	Name    string
	Order   int16
	Deleted bool
}
type MSG_PROJECT_tree_manageChild struct {
	RootID         int32
	ViewType       string
	Modules        []*MSG_PROJECT_tree_cache
	ParentModuleID int32
}
type MSG_PROJECT_tree_manageChild_result struct {
	Result ErrCode
	Name   string
}
type MSG_PROJECT_product_getStoriesMapBySql struct {
	Field   string
	Where   map[string]interface{}
	Order   string
	Group   string
	Page    int
	PerPage int
	Total   int
}
type MSG_PROJECT_product_getStoriesMapBySql_result struct {
	List  []map[string]string
	Total int
}
type MSG_PROJECT_tree_updateList struct {
	Modules []*MSG_PROJECT_tree_cache
}
type MSG_PROJECT_tree_delete struct {
	Ids []int32
}
type MSG_PROJECT_productplan_getPairsForStory struct {
	Product int32
	Branch  int32
}
type MSG_PROJECT_productplan_getPairsForStory_result struct {
	List []HtmlKeyValueStr
}
type MSG_PROJECT_productplan_getList struct {
	Ids        []int32
	ProductID  int32
	Branch     int32
	BrowseType string
	Order      string
	Page       int
	PerPage    int
	Total      int
}
type MSG_PROJECT_productplan_getList_result struct {
	List  []map[string]string
	Total int
}
type MSG_PROJECT_productplan_getLast struct {
	ProductId int32
	Branch    int32
}
type MSG_PROJECT_productplan_getLast_result struct {
	Result map[string]string
}
type MSG_PROJECT_product_editBranch struct {
	ProductID int32
	Branchs   []*MSG_PROJECT_branch_info
}
type MSG_PROJECT_product_deleteBranch struct {
	ProductID int32
	BranchID  int32
}
type MSG_PROJECT_product_deleteBranch_result struct {
	Result ErrCode
}
type MSG_PROJECT_productplan_getPairs struct {
	ProductID int32
	BranchID  int32
	Expired   string
	Ids       []int32
}
type MSG_PROJECT_productplan_getPairs_result struct {
	List []HtmlKeyValueStr
}
type MSG_PROJECT_productplan_insertUpdate struct {
	Id       int32 `db:"pk"`
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
	Id     int32
	Result ErrCode
}
type MSG_PROJECT_productplan_delete struct {
	Id      int32
	Product int32
	Branch  int32
}
type MSG_PROJECT_stroy_create struct {
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
	Estimate      float64
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
	Result int32 //小于0为ErrCode,大于0位新增ID
}
type MSG_PROJECT_story_batchGetStoryStage struct {
	Ids []int32
}
type MSG_PROJECT_story_batchGetStoryStage_result struct {
	List map[int32][]HtmlKeyValueStr
}
type MSG_PROJECT_story_getById struct {
	Id      int32
	Version int16
}
type MSG_PROJECT_story_getById_result struct {
	Story *MSG_PROJECT_story
}
type MSG_PROJECT_project_getById struct {
	Id int32
}
type MSG_PROJECT_project_getById_result struct {
	Project *MSG_PROJECT_project_cache
}
type MSG_PROJECT_project_cache struct {
	Id            int32
	IsCat         bool
	CatID         int32
	Type          string
	Parent        int32
	Name          string
	Code          string
	Begin         time.Time
	End           time.Time
	Days          int16
	Status        string
	Statge        int8
	Pri           int8
	Desc          string
	OpenedBy      int32
	OpenedDate    time.Time
	OpenedVersion string
	ClosedBy      int32
	ClosedDate    time.Time
	CanceledBy    int32
	CanceledDate  time.Time
	PO            int32
	PM            int32
	QD            int32
	RD            int32
	Team          string
	Acl           string
	Whitelist     []int32
	Order         int32
	Deleted       bool
	FtpPath       string
	Products      []int32
	Branchs       []int32
	Storys        []int32
	Plans         []int32
	Delay         int64                 `db:"-"` //延期几天
	Hours         map[string]float64    `db:"-"` //时间统计
	Teams         []*MSG_USER_team_info `db:"-"`
}
type MSG_PROJECT_StoryStage struct {
	Story  int32
	Branch int32
	Stage  string
}
type MSG_PROJECT_TASK struct {
	Id             int32
	Ancestor       int32
	Parent         int32
	Project        int32
	Module         int32
	Story          int32
	StoryVersion   int16
	FromBug        int32
	Name           string
	Type           string
	Pri            int8
	Estimate       float64
	Consumed       float64
	Left           float64
	Deadline       time.Time
	Status         string
	Color          string
	Mailto         []int32
	Desc           string
	OpenedBy       int32
	OpenedDate     time.Time
	AssignedTo     int32
	AssignedDate   time.Time
	EstStarted     time.Time
	RealStarted    time.Time
	FinishedBy     int32
	FinishedDate   time.Time
	CanceledBy     int32
	CanceledDate   time.Time
	ClosedBy       int32
	ClosedDate     time.Time
	ClosedReason   string
	LastEditedBy   int32
	LastEditedDate time.Time
	Examine        bool
	ExamineDate    time.Time
	ExamineBy      int32
	Deleted        bool
	Finalfile      bool
	Proofreading   bool
	Team           []int32
	PlaceOrder     bool
	//从stroy拿数据
	StoryID            int32  `db:"-"`
	StoryTitle         string `db:"-"`
	StoryStatus        string `db:"-"`
	LatestStoryVersion int16  `db:"-"`
	Product            int32  `db:"-"`
	Branch             int32  `db:"-"`
	//processTask计算出来
	Progress      int                 `db:"-"`
	Delay         int32               `db:"-"` //延期几天
	Children      []*MSG_PROJECT_TASK `db:"-"`
	Grandchildren []*MSG_PROJECT_TASK `db:"-"`
	NeedConfirm   bool                `db:"-"`
}
type MSG_PROJECT_productplan_getById struct {
	Id int32
}
type MSG_PROJECT_productplan_getById_result struct {
	Info *MSG_PROJECT_productplan
}
type MSG_PROJECT_productplan struct {
	Id       int32
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
type MSG_PROJECT_build struct {
	Id       int32
	Product  int32
	Branch   int32
	Project  int32
	Name     string
	ScmPath  string
	FilePath string
	Date     time.Time
	Stories  []int32
	Bugs     []int32
	Builder  string
	Desc     string
	Deleted  bool
}
type MSG_PROJECT_build_getById struct {
	Id int32
}
type MSG_PROJECT_build_getById_result struct {
	Info *MSG_PROJECT_build
}
type MSG_PROJECT_release struct {
	Id       int32
	Product  int32
	Branch   int32
	Build    int32
	Name     string
	Marker   bool
	Date     time.Time
	Stories  []int32
	Bugs     []int32
	LeftBugs string
	Desc     string
	Status   string
	Deleted  bool
}
type MSG_PROJECT_release_getById struct {
	Id int32
}
type MSG_PROJECT_release_getById_result struct {
	Info *MSG_PROJECT_release
}
type MSG_PROJECT_task_getPairs struct {
	Where map[string]interface{}
}
type MSG_PROJECT_task_getPairs_result struct {
	List []HtmlKeyValueStr
}
type MSG_PROJECT_task_getListByWhereMap struct {
	Where   map[string]interface{}
	Order   string
	Page    int
	PerPage int
	Total   int
}
type MSG_PROJECT_task_getListByWhereMap_result struct {
	List  []*MSG_PROJECT_TASK
	Total int
}
type MSG_PROJECT_project_getBurn struct {
	ProjectIds []int32
}
type MSG_PROJECT_project_getBurn_result struct {
	List []*MSG_PROJECT_project_Burn_info
}
type MSG_PROJECT_project_Burn_info struct {
	Project  int32
	Date     time.Time
	Estimate float64
	Left     float64
	Consumed float64
}
type MSG_PROJECT_story_getPlanStories struct {
	PlanID  int32
	Status  string
	OrderBy string
}
type MSG_PROJECT_story_getPlanStories_result struct {
	List []*MSG_PROJECT_story
}
type MSG_PROJECT_project_linkStory struct {
	ProjectID int32
	Stories   []int32
	Products  map[int32]int32
}

type MSG_PROJECT_branch_getByProducts struct {
	Products     []int32
	AppendBranch []int32
}
type MSG_PROJECT_branch_getByProducts_result struct {
	List map[int32][]HtmlKeyValueStr
}
type MSG_PROJECT_project_create struct {
	CopyProjectID int32
	Info          *MSG_PROJECT_project_cache
}
type MSG_PROJECT_project_create_result struct {
	Id int32
}
type MSG_PROJECT_project_statRelatedData struct {
	ProjectID int32
}
type MSG_PROJECT_project_statRelatedData_result struct {
	StoryCount int
	TaskCount  int
	BugCount   int
}
type MSG_PROJECT_story_getPairsByIds struct {
	Ids []int32
}
type MSG_PROJECT_story_getPairsByIds_result struct {
	List []HtmlKeyValueStr
}
type MSG_PROJECT_product_getPairsByIds struct {
	Ids []int32
}
type MSG_PROJECT_product_getPairsByIds_result struct {
	List []HtmlKeyValueStr
}
type MSG_PROJECT_project_getPairsByIds struct {
	Ids []int32
}
type MSG_PROJECT_project_getPairsByIds_result struct {
	List []HtmlKeyValueStr
}
type MSG_PROJECT_branch_getPairsByIds struct {
	Ids []int32
}
type MSG_PROJECT_branch_getPairsByIds_result struct {
	List []HtmlKeyValueStr
}
type MSG_PROJECT_tree_getPairsByIds struct {
	Ids []int32
}
type MSG_PROJECT_tree_getPairsByIds_result struct {
	List []HtmlKeyValueStr
}
type MSG_PROJECT_project_start struct {
	Id      int32
	Comment string //备注
}
type MSG_PROJECT_project_putoff struct {
	Id      int32
	Begin   time.Time
	End     time.Time
	Days    int16
	Comment string
}
type MSG_PROJECT_project_suspend struct {
	Id      int32
	Comment string
}
type MSG_PROJECT_project_activate struct {
	Id           int32
	Begin        time.Time
	End          time.Time
	Comment      string
	ReadjustTask bool
}
type MSG_PROJECT_project_close struct {
	Id      int32
	Comment string
}
type MSG_PROJECT_project_delete struct {
	Id int32
}
type MSG_PROJECT_project_getProjectTasks struct {
	ProjectID int32
	ProductID int32
	Type      []string
	ModuleID  int32
	OrderBy   string
	Role      string
	Page      int
	PerPage   int
	Total     int
}
type MSG_PROJECT_project_getProjectTasksByWhere struct {
	Where   string
	OrderBy string
	Role    string
	Page    int
	PerPage int
	Total   int
}
type MSG_PROJECT_project_getProjectTasks_result struct {
	List  []*MSG_PROJECT_TASK
	Total int
}
type MSG_PROJECT_tree_getTaskTreeModules struct {
	ProjectID int32
	Parent    bool
	//LinkStory bool 默认true
}
type MSG_PROJECT_tree_getTaskTreeModules_result struct {
	ProjectModules map[int32]int32
}
type MSG_PROJECT_task_getById struct {
	StarTime time.Time
	Id       int32
}

type MSG_PROJECT_task_getById_result struct {
	Info *MSG_PROJECT_TASK
}

type MSG_PROJECT_story_getProjectStoryPairs struct {
	ProjectID    int32
	ProductID    int32
	Branch       int32
	ModuleIdList []int32
	ShortType    bool
}
type MSG_PROJECT_story_getProjectStoryPairs_result struct {
	List []*MSG_PROJECT_TASK
}
type MSG_PROJECT_task_create struct {
	Task *MSG_PROJECT_TASK
}
type MSG_PROJECT_task_create_result struct {
	Id     int32
	Err    ErrCode
	Change ChangeHistory
}
type MSG_PROJECT_TaskEstimate struct {
	Id       int32
	Task     int32
	Date     time.Time
	Estimate float64
	Left     float64
	Consumed float64
	Uid      int32
	Account  string
	Work     string
}
type MSG_PROJECT_task_GetTaskEstimateByTaskId struct {
	TaskId int32
}
type MSG_PROJECT_task_GetTaskEstimateByTaskId_result struct {
	List []*MSG_PROJECT_TaskEstimate
}
type MSG_PROJECT_task_UpdateTaskEstimate struct {
	TaskId int32
	List   []*MSG_PROJECT_TaskEstimate
}
type MSG_PROJECT_task_UpdateTaskEstimate_result struct {
	Changes []*MSG_LOG_History
}

type MSG_PROJECT_story_getProductStories struct {
	Products []int32
	Branches []int32
	ModuleID []int32
	Status   []string
	Sort     string
	Page     int
	PerPage  int
	Total    int
}
type MSG_PROJECT_story_getProductStories_result struct {
	List  []*MSG_PROJECT_story
	Total int
}

type MSG_PROJECT_productplan_getForProducts struct {
	Products []int32
}
type MSG_PROJECT_productplan_getForProducts_result struct {
	List []HtmlKeyValueStr
}

type MSG_PROJECT_task_assignTo struct {
	TaskID     int32
	AssignedTo int32
	Left       float64
	Comment    string
}

type MSG_PROJECT_task_start struct {
	TaskID      int32
	RealStarted time.Time
	Consumed    float64
	Left        float64
	Comment     string
	MethodName  string
}

type MSG_PROJECT_task_start_result struct {
	Changes []*MSG_LOG_History
}
type MSG_PROJECT_task_finish struct {
	TaskID       int32
	FinishedDate time.Time
	Consumed     float64
	Left         float64
	Mailto       []int32
	Comment      string
}

type MSG_PROJECT_task_finish_result struct {
	Changes []*MSG_LOG_History
}

type MSG_PROJECT_task_activate struct {
	TaskID     int32
	AssignedTo int32
	Left       float64
	Mailto     []int32
	Comment    string
}

type MSG_PROJECT_task_pause struct {
	TaskID  int32
	Comment string
}

type MSG_PROJECT_task_internalaudit struct {
	TaskID  int32
	Comment string
	Mailto  []int32
}

type MSG_PROJECT_task_proofreading struct {
	TaskID       int32
	Comment      string
	Proofreading bool
}
type MSG_PROJECT_task_close struct {
	TaskID  int32
	Comment string
}
type MSG_PROJECT_task_getStoryTaskCounts struct {
	Stories   []int32
	ProjectID int32
}
type MSG_PROJECT_task_getStoryTaskCounts_result struct {
	List map[int32]int
}
type MSG_PROJECT_task_examine struct {
	TaskID    int32
	ProjectId int32
	Examine   bool
}
type MSG_PROJECT_task_cancel struct {
	TaskID  int32
	Comment string
}

type MSG_PROJECT_task_delete struct {
	TaskID    int32
	ProjectId int32
}

type MSG_PROJECT_task_placeOrder struct {
	TaskID int32
	Action bool
}

type MSG_PROJECT_getAllprojectProductID struct {
}

type MSG_PROJECT_getAllprojectProductID_result struct {
	ProjectID []int32
	ProductID []int32
}

type MSG_PROJECT_doRawSelect struct {
	Sql string
}
type MSG_PROJECT_doRawSelect_result struct {
	List []map[string]string
}
type MSG_PROJECT_updateCache struct {
	Type string
	Ids  []int32
}
