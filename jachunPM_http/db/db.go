package db

import (
	"jachunPM_user/config"
	"log"
	"mysql"
	"time"
)

const (
	TABLE_LOG_MSG = "Log_msg"
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
	if errs != nil {
		log.Fatalf("数据库启动失败%v", errs)
	}
}

type Log_msg struct {
	Msgno     uint32 `db:"not null;pk"`
	Ttl       uint8  `db:"not null;pk"`
	LocalNo   uint8
	LocalId   uint8
	RemoteNo  uint8
	RemoteId  uint8
	Cmd       int32
	Timestamp time.Time `db:"default(current_timestamp());extra('on update current_timestamp()')"`
}

func (*Log_msg) TableName() string {
	return TABLE_LOG_MSG
}
