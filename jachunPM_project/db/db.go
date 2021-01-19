package db

import (
	"jachunPM_project/config"
	"log"
	"mysql"
	"time"
)

const (
	TABLE_MODULE  = "module"
	TABLE_PRODUCT = "product"
	TABLE_DOCLIB  = "doclib"
	TABLE_STORY   = "story"
	TABLE_BRANCH  = "branch"
)

func Init() *mysql.MysqlDB {
	var err error
	db, err := mysql.Open(config.Config.MysqlDsn)
	if err != nil {
		log.Fatalf("数据库连接失败 %v", err)
	}
	if config.Config.MysqlMaxConn > 0 {
		db.MaxOpenConns = config.Config.MysqlMaxConn
	}
	errs := db.StoreEngine("Innodb").Sync2(
		new(Module),
		new(Doclib),
		new(Product),
		new(Story),
		new(Branch),
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
	Collector string `db:"type:text"`
	Short     string `db:"type:varchar(30)"`
	Deleted   bool   `db:"default(0)"` // 0=0,1=1,
	TimeStamp int64
}

func (*Module) TableName() string {
	return TABLE_MODULE
}

type Product struct {
	Id          int32  `db:"auto_increment;pk"`
	Name        string `db:"type:varchar(90)"`
	Code        string `db:"type:varchar(45)"`
	Line        int32  `db:"not null"`
	Type        string `db:"default('normal');type:varchar(30)"`
	Status      string `db:"type:varchar(30)"`
	Desc        string `db:"type:text"`
	PO          int32
	QD          int32
	RD          int32
	Acl         string  `db:"type:enum('open','private','custom')"`
	Whitelist   []int32 `db:"type:tinytext"`
	CreatedBy   int32
	CreatedDate int64
	Order       int32
	Deleted     bool
	TimeStamp   int64
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
	Id             int32   `db:"auto_increment;pk"`
	Product        int32   `db:"default(0)"`
	Branch         int32   `db:"default(0)"`        //平台
	Module         int32   `db:"default(0)"`        //模块
	Plan           string  `db:"type:text"`         //计划
	Source         string  `db:"type:varchar(20)"`  //需求来源
	SourceNote     string  `db:"type:varchar(255)"` //备注
	FromBug        int32   `db:"default(0)"`
	Title          string  `db:"type:varchar(255)"`
	Keywords       string  `db:"type:varchar(255)"`
	Pri            int8    `db:"default(3)"` //优先级
	Estimate       float32 `db:""`
	Status         string  `db:"type:enum('','changed','active','draft','closed')"`                                                                                //
	Stage          string  `db:"type:enum('','wait','planned','projected','developing','developed','testing','tested','verified','released','closed');default(1)"` //
	Mailto         string  `db:"type:text"`                                                                                                                        //抄送
	OpenedBy       string  `db:"type:varchar(30)"`
	OpenedDate     time.Time
	AssignedTo     string `db:"type:varchar(30)"` //指派
	AssignedDate   time.Time
	LastEditedBy   string `db:"type:varchar(30)"`
	LastEditedDate time.Time
	ReviewedBy     string `db:"type:varchar(255)"`
	ReviewedDate   time.Time
	ClosedBy       string `db:"type:varchar(30)"`
	ClosedDate     time.Time
	ClosedReason   string `db:"type:varchar(30)"`
	ToBug          int32
	ChildStories   string `db:"type:varchar(255)"`
	LinkStories    string `db:"type:varchar(255)"`
	DuplicateStory int32
	Deleted        bool
	Version        int16 `db:"not null;default(1)"`
	//Color          string  `db:"type:varchar(7)"`
	//Type           string  `db:"type:varchar(30)"`
}

func (*Story) TableName() string {
	return TABLE_STORY
}

type Branch struct {
	Id        int32  `db:"auto_increment;pk"`
	Product   int32  `db:""`
	Name      string `db:"type:varchar(255)"`
	Order     int16  `db:""`
	Deleted   bool
	TimeStamp int64
}

func (*Branch) TableName() string {
	return TABLE_BRANCH
}
