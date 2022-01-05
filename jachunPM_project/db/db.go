package db

import (
	"jachunPM_project/config"
	"log"
	"mysql"
	"protocol"
	"time"
)

const (
	TABLE_MODULE       = "module"
	TABLE_PRODUCT      = "product"
	TABLE_DOCLIB       = "doclib"
	TABLE_STORY        = "story"
	TABLE_BRANCH       = "branch"
	TABLE_PRODUCTPLAN  = "productplan"
	TABLE_PROJECT      = "project"
	TABLE_STORYSPEC    = "storyspec"
	TABLE_STORYSTAGE   = "storystage"
	TABLE_TASK         = "task"
	TABLE_RELEASE      = "release"
	TABLE_BUILD        = "build"
	TABLE_BURN         = "burn"
	TABLE_TASKESTIMATE = "taskestimate"
)

func Init() *mysql.MysqlDB {
	var err error
	db, err := mysql.Open(config.Config.MysqlDsn)
	if err != nil {
		log.Fatalf("数据库连接失败 %v", err)
	}
	if config.Config.MysqlMaxConn > 0 {
		db.SetMaxOpenConns(config.Config.MysqlMaxConn)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("数据库启动失败 %v", err)
	}
	errs := db.StoreEngine("Innodb").Sync2(
		new(Module),
		new(Doclib),
		new(Product),
		new(Story),
		new(Branch),
		new(Productplan),
		new(Project),
		new(StorySpec),
		new(StoryStage),
		new(Task),
		new(Release),
		new(Build),
		new(Burn),
		new(TaskEstimate),
	)
	db.Regsiter(&protocol.MSG_PROJECT_product_cache{},
		&protocol.MSG_PROJECT_tree_cache{},
		&protocol.MSG_PROJECT_project_cache{},
		&protocol.MSG_PROJECT_TASK{},
		&protocol.MSG_PROJECT_TaskEstimate{},
		&protocol.MSG_PROJECT_story{},
	)

	if errs != nil {
		log.Fatalf("数据库启动失败%v", errs)
	}
	return db
}

type Module struct {
	Id        int32   `db:"auto_increment;pk"`
	Root      int32   `db:"default(0)"`
	Branch    int32   `db:"default(0)"`
	Name      string  `db:"type:varchar(60)"`
	Parent    int32   `db:"default(0)"`
	Path      []int32 `db:"type:json"`
	Grade     int8    `db:"default(0)"`
	Order     int16   `db:"default(0)"`
	Type      string  `db:"type:varchar(30)"`
	Owner     string  `db:"type:varchar(30)"`
	OwnerID   int32
	Collector string    `db:"type:text"`
	Short     string    `db:"type:varchar(30)"`
	Deleted   bool      `db:"default(0)"` // 0=0,1=1,
	TimeStamp time.Time `db:"default(current_timestamp());extra('on update current_timestamp()')"`
}

func (*Module) TableName() string {
	return TABLE_MODULE
}

type Product struct {
	Id          int32   `db:"auto_increment;pk"`
	Name        string  `db:"type:varchar(90)"`
	Code        string  `db:"type:varchar(45)"`
	Line        int32   `db:"not null"`
	Type        string  `db:"default('normal');type:varchar(30)"`
	Status      string  `db:"type:varchar(30)"`
	Desc        string  `db:"type:text"`
	Branch      []int32 //product包含的branch
	Plan        []int32 //product包含的plan
	PO          int32
	QD          int32
	RD          int32
	Acl         string `db:"type:enum('open','private','custom')"`
	Whitelist   []int32
	CreatedBy   int32
	CreatedDate time.Time
	Order       int32
	Deleted     bool
	TimeStamp   time.Time `db:"default(current_timestamp());extra('on update current_timestamp()')"`
}

func (*Product) TableName() string {
	return TABLE_PRODUCT
}

type Doclib struct {
	Id        int16  `db:"auto_increment;pk"`
	Type      string `db:"type:varchar(30)"`
	Product   int32
	Project   int32
	Name      string  `db:"type:varchar(60)"`
	Acl       string  `db:"default('open');type:varchar(10)"`
	Groups    []int32 `db:"type:varchar(255)"`
	Users     []int32 `db:"type:text"`
	Main      bool    `db:"default(0)"`
	Collector string  `db:"type:text"`
	Order     int8    `db:""`
	Deleted   bool    `db:"default(0)"`
}

func (*Doclib) TableName() string {
	return TABLE_DOCLIB
}

type Story struct {
	Id             int32  `db:"auto_increment;pk"`
	Product        int32  `db:"default(0)"`
	Project        int32  `db:"index"`
	Branch         int32  `db:"default(0)"`        //平台
	Module         int32  `db:"default(0)"`        //模块
	Plan           int32  `db:"type:text"`         //计划
	Source         string `db:"type:varchar(20)"`  //需求来源
	SourceNote     string `db:"type:varchar(255)"` //备注
	FromBug        int32  `db:"default(0)"`
	Title          string `db:"type:varchar(255)"`
	Keywords       string `db:"type:varchar(255)"`
	Pri            int8   `db:"default(3)"` //优先级
	Estimate       float64
	Status         string  `db:"type:enum('','changed','active','draft','closed')"`                                                                                //
	Stage          string  `db:"type:enum('','wait','planned','projected','developing','developed','testing','tested','verified','released','closed');default(1)"` //
	Mailto         []int32 `db:"type:json"`                                                                                                                        //抄送
	OpenedBy       int32   `db:"type:varchar(30)"`
	OpenedDate     time.Time
	AssignedTo     int32 `db:"type:varchar(30)"` //指派
	AssignedDate   time.Time
	LastEditedBy   int32 `db:"type:varchar(30)"`
	LastEditedDate time.Time
	ReviewedBy     int32 `db:"type:varchar(255)"`
	ReviewedDate   time.Time
	ClosedBy       int32 `db:"type:varchar(30)"`
	ClosedDate     time.Time
	ClosedReason   string `db:"type:varchar(30)"`
	ToBug          int32
	ChildStories   []int32 `db:"type:json"`
	LinkStories    []int32 `db:"type:json"`
	DuplicateStory int32
	Deleted        bool
	Version        int16  `db:"not null;default(1)"`
	Color          string `db:"type:varchar(7)"`
	//Type           string  `db:"type:varchar(30)"`
}

func (*Story) TableName() string {
	return TABLE_STORY
}

type Branch struct {
	Id        int32  `db:"auto_increment;pk"`
	Product   int32  `db:"index"`
	Name      string `db:"type:varchar(255)"`
	Order     int16  `db:""`
	Deleted   bool
	TimeStamp time.Time `db:"default(current_timestamp());extra('on update current_timestamp()')"`
}

func (*Branch) TableName() string {
	return TABLE_BRANCH
}

type Productplan struct {
	Id       int32 `db:"auto_increment;pk"`
	Product  int32 `db:"index"`
	Branch   int32 `db:"index"`
	Parent   int32
	Projects []int32
	Title    string    `db:"type:varchar(255)"`
	Desc     string    `db:"type:text"`
	Begin    time.Time `db:"index;type:date"`
	End      time.Time `db:"index;type:date"`
	Order    string    `db:"type:text"`
	Deleted  bool
}

func (*Productplan) TableName() string {
	return TABLE_PRODUCTPLAN
}

type Project struct {
	Id            int32 `db:"auto_increment;pk"`
	IsCat         bool
	CatID         int32
	Type          string    `db:"default('sprint');type:varchar(20)"`
	Parent        int32     `db:"default(0)"`
	Name          string    `db:"type:varchar(90)"`
	Code          string    `db:"type:varchar(45)"`
	Begin         time.Time `db:"not null;type:date"`
	End           time.Time `db:"not null;type:date"`
	Days          int16
	Status        string `db:"type:varchar(10)"`
	Statge        int8   `db:"default(0)"` // 0=1,1=2,2=3,3=4,4=5,
	Pri           int8   `db:"default(0)"` // 0=1,1=2,2=3,3=4,
	Desc          string `db:"type:text"`
	OpenedBy      int32  `db:"default(0)"`
	OpenedDate    time.Time
	OpenedVersion string `db:"type:varchar(20)"`
	ClosedBy      int32
	ClosedDate    time.Time
	CanceledBy    int32
	CanceledDate  time.Time
	PO            int32
	PM            int32
	QD            int32
	RD            int32
	Team          string
	Acl           string  `db:"default(0)"` // 0=open,1=private,2=custom,
	Whitelist     []int32 `db:"type:json"`
	Order         int32
	Deleted       bool
	FtpPath       string    `db:"type:varchar(255)"`
	Products      []int32   `db:"index"`
	Branchs       []int32   `db:"index"`
	Storys        []int32   `db:"index"`
	Plans         []int32   `db:"index"`
	TimeStamp     time.Time `db:"default(current_timestamp());extra('on update current_timestamp()')"`
}

func (*Project) TableName() string {
	return TABLE_PROJECT
}

type StorySpec struct {
	Story   int32  `db:"pk;index"`
	Version int16  `db:"pk;index"`
	Title   string `db:"type:varchar(255)"`
	Spec    string `db:"type:text"`
	Verify  string `db:"type:text"`
}

func (*StorySpec) TableName() string {
	return TABLE_STORYSPEC
}

type StoryStage struct {
	Story  int32  `db:"pk"`
	Branch int32  `db:"pk"`
	Stage  string `db:"type:varchar(50)"`
}

func (*StoryStage) TableName() string {
	return TABLE_STORYSTAGE
}

type Task struct {
	Id             int32     `db:"auto_increment;pk"`
	Ancestor       int32     `db:"index"`
	Parent         int32     `db:"not null;default(0);index"`
	Project        int32     `db:"default(0);index"`
	Module         int32     `db:"default(0)"`
	Story          int32     `db:"default(0);index"`
	StoryVersion   int16     `db:"not null;default(1)"`
	FromBug        int32     `db:"default(0)"`
	Name           string    `db:"type:varchar(255)"`
	Type           string    `db:"type:varchar(20)"`
	Pri            int8      `db:"default(0)"`
	Estimate       float64   `db:""`
	Consumed       float64   `db:""`
	Left           float64   `db:""`
	Deadline       time.Time `db:"type:date"`
	Status         string    `db:"type:varchar(32)"`
	Color          string    `db:"type:varchar(7)"`
	Mailto         []int32   `db:"type:json"`
	Desc           string    `db:"type:text"`
	OpenedBy       int32     `db:"type:varchar(30)"`
	OpenedDate     time.Time `db:"not null"`
	AssignedTo     int32     `db:"type:varchar(30);index"`
	AssignedDate   time.Time `db:"not null"`
	EstStarted     time.Time `db:"type:date"`
	RealStarted    time.Time `db:"type:date"`
	FinishedBy     int32     `db:"type:varchar(30)"`
	FinishedDate   time.Time `db:"not null"`
	FinishedList   []int32   `db:""`
	CanceledBy     int32     `db:"type:varchar(30)"`
	CanceledDate   time.Time `db:"not null"`
	ClosedBy       int32     `db:"type:varchar(30)"`
	ClosedDate     time.Time `db:"not null"`
	ClosedReason   string    `db:"type:varchar(30)"`
	LastEditedBy   int32     `db:"type:varchar(30)"`
	LastEditedDate time.Time `db:"not null"`
	Examine        bool      `db:"not null;default(0);index"`
	ExamineDate    time.Time `db:"not null"`
	ExamineBy      int32     `db:"type:varchar(30)"`
	Deleted        bool
	Finalfile      bool `db:"default('0');type:varchar(3)"`
	Proofreading   bool
	Team           []int32
	PlaceOrder     bool
}

func (*Task) TableName() string {
	return TABLE_TASK
}

type Release struct {
	Id       int32  `db:"auto_increment;pk"`
	Product  int32  `db:"default(0)"`
	Branch   int32  `db:"default(0)"`
	Build    int32  `db:""`
	Name     string `db:"type:varchar(30)"`
	Marker   bool
	Date     time.Time `db:"not null"`
	Stories  []int32   `db:"type:json"`
	Bugs     []int32   `db:"type:json"`
	LeftBugs string    `db:"type:text"`
	Desc     string    `db:"type:text"`
	Status   string    `db:"default('normal');type:varchar(20)"`
	Deleted  bool
}

func (*Release) TableName() string {
	return TABLE_RELEASE
}

type Build struct {
	Id       int32  `db:"auto_increment;pk"`
	Product  int32  `db:"default(0)"`
	Branch   int32  `db:"default(0)"`
	Project  int32  `db:"index"`
	Name     string `db:"type:varchar(150)"`
	ScmPath  string `db:"type:varchar(255)"`
	FilePath string `db:"type:varchar(255)"`
	Date     time.Time
	Stories  []int32 `db:"type:json"`
	Bugs     []int32 `db:"type:json"`
	Builder  string  `db:"type:varchar(30)"`
	Desc     string  `db:"type:text"`
	Deleted  bool
}

func (*Build) TableName() string {
	return TABLE_BUILD
}

type Burn struct {
	Project  int32     `db:"pk"`
	Date     time.Time `db:"pk;type:date"`
	Estimate float64
	Left     float64
	Consumed float64
}

func (*Burn) TableName() string {
	return TABLE_BURN
}

type TaskEstimate struct {
	Id       int32     `db:"auto_increment;pk"`
	Task     int32     `db:"index"`
	Date     time.Time `db:"not null"`
	Estimate float64
	Left     float64 `db:"default(0)"`
	Consumed float64 `db:""`
	Uid      int32
	Account  string
	Work     string `db:"type:text"`
}

func (*TaskEstimate) TableName() string {
	return TABLE_TASKESTIMATE
}
