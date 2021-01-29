package db

import (
	"jachunPM_test/config"
	"log"
	"mysql"
	"time"
)

const (
	TABLE_BUG = "bug"
)

func Init() *mysql.MysqlDB {
	db, err := mysql.Open(config.Config.MysqlDsn)
	if err != nil {
		log.Fatalf("数据库连接失败 %v", err)
	}
	if config.Config.MysqlMaxConn > 0 {
		db.MaxOpenConns = config.Config.MysqlMaxConn
	}
	errs := db.StoreEngine("TokuDB").Sync2(
		new(Bug),
	)
	if errs != nil {
		log.Fatalf("数据库启动失败%v", errs)
	}
	return db
}

type Bug struct {
	Id             int32     `db:"not null;auto_increment;pk"`
	Product        int32     `db:"default(0)"`
	Branch         int32     `db:"default(0)"`
	Module         int32     `db:"default(0)"`
	Project        int32     `db:"default(0)"`
	Plan           int32     `db:"default(0)"`
	Story          int32     `db:"default(0)"`
	StoryVersion   int16     `db:"not null;default(1)"`
	Task           int32     `db:"default(0)"`
	ToTask         int32     `db:"default(0)"`
	ToStory        int32     `db:"not null;default(0)"`
	Title          string    `db:"type:varchar(255)"`
	Keywords       string    `db:"type:varchar(255)"`
	Severity       int8      `db:"not null;default(0)"`
	Pri            int8      `db:""`
	Type           string    `db:"type:varchar(30)"`
	Os             string    `db:"type:varchar(30)"`
	Browser        string    `db:"type:varchar(30)"`
	Hardware       string    `db:"type:varchar(30)"`
	Found          string    `db:"type:varchar(30)"`
	Steps          string    `db:"type:text"`
	Status         int8      `db:"default(0)"` // 0=active,1=resolved,2=closed,
	Color          string    `db:"type:varchar(7)"`
	Confirmed      int8      `db:"not null;default(0)"`
	ActivatedCount int16     `db:"not null"`
	ActivatedDate  time.Time `db:"not null"`
	Mailto         string    `db:"type:text"`
	OpenedBy       string    `db:"type:varchar(30)"`
	OpenedDate     time.Time `db:"not null"`
	OpenedBuild    string    `db:"type:varchar(255)"`
	AssignedTo     string    `db:"type:varchar(30)"`
	AssignedDate   time.Time `db:"not null"`
	Deadline       time.Time `db:"not null"`
	ResolvedBy     string    `db:"type:varchar(30)"`
	Resolution     string    `db:"type:varchar(30)"`
	ResolvedBuild  string    `db:"type:varchar(30)"`
	ResolvedDate   time.Time `db:"not null"`
	ClosedBy       string    `db:"type:varchar(30)"`
	ClosedDate     time.Time `db:"not null"`
	DuplicateBug   int32     `db:""`
	LinkBug        string    `db:"type:varchar(255)"`
	Case           int32     `db:""`
	CaseVersion    int16     `db:"not null;default(1)"`
	Result         int32     `db:""`
	Testtask       int32     `db:""`
	LastEditedBy   string    `db:"type:varchar(30)"`
	LastEditedDate time.Time `db:"not null"`
	Deleted        bool
}

func (Bug) TableName() string {
	return TABLE_BUG
}
