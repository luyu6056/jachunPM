package db

import (
	"jachunPM_commom/config"
	"log"
	"mysql"
	"time"
)

const (
	TABLE_LOG_MSG = "Log_msg"
	TABLE_FILE    = "file"
)

var DB *mysql.MysqlDB

func Init() {
	var err error
	DB, err = mysql.Open(config.Config.MysqlDsn)
	if err != nil {
		log.Fatalf("数据库连接失败 %v", err)
	}
	if config.Config.MysqlMaxConn > 0 {
		DB.MaxOpenConns = config.Config.MysqlMaxConn
	}
	errs := DB.StoreEngine("TokuDB").Sync2(
		new(Log_msg),
	)
	errs = append(errs, DB.StoreEngine("Innodb").Sync2(
		new(File),
	)...)
	if len(errs) > 0 {
		log.Fatalf("数据库启动失败%v", errs)
	}
	go UpdatedbInit()
}

type Log_msg struct {
	Msgno     uint32 `db:"not null;pk"`
	Ttl       uint8  `db:"not null;pk"`
	LocalNo   uint8
	LocalId   uint8
	RemoteNo  uint8
	RemoteId  uint8
	Cmd       string `db:"type:tinytext"`
	Err       string `db:"type:text"`
	Stack     string
	Timestamp time.Time `db:"default(current_timestamp());extra('on update current_timestamp()')"`
}

func (*Log_msg) TableName() string {
	return TABLE_LOG_MSG
}

type File struct {
	Id         int64     `db:"auto_increment;pk"`
	Pathname   string    `db:"type:varchar(255)"`
	Title      string    `db:"type:varchar(255)"`
	Extension  string    `db:"type:varchar(30)"`
	Size       int       `db:"default(0)"`
	ObjectType string    `db:"type:varchar(30)"`
	ObjectID   int32     `db:"not null"`
	AddedBy    string    `db:"type:varchar(30)"`
	AddedDate  time.Time `db:"not null"`
	Deleted    bool      `db:"default(0)"` // 0=0,1=1,
	Type       string    `db:"type:varchar(50)"`
}

func (*File) TableName() string {
	return TABLE_FILE
}
