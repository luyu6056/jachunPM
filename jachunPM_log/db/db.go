package db

import (
	"jachunPM_log/config"
	"log"
	"mysql"
	"time"
)

const (
	TABLE_ACTION = "action"
)

func Init() *mysql.MysqlDB {
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
	errs := db.StoreEngine("TokuDB").Sync2(
		new(Action),
	)
	if errs != nil {
		log.Fatalf("数据库启动失败%v", errs)
	}
	return db
}

type Action struct {
	Id         int64   `db:"auto_increment;pk"`
	ObjectType string  `db:"type:varchar(30)"`
	ObjectID   int32   `db:"default(0)"`
	Products   []int32 `db:"type:json"`
	Projects   []int32 `db:"type:json"`
	ActorId    int32
	Actor      string    `db:"type:varchar(30)"`
	Action     string    `db:"type:varchar(30)"`
	Date       time.Time `db:"not null"`
	Comment    string    `db:"type:text"`
	Extra      string    `db:"type:text"`
	Read       bool
	Historys   []*History `db:"type:json"`
}

func (*Action) TableName() string {
	return TABLE_ACTION
}

type History struct {
	Field string
	Old   string
	New   string
	Diff  string
}
