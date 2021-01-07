package db

import (
	"jachunPM_test/config"
	"log"
	"mysql"
	"time"
)

const (
	TABLE_TESTSUITE = "testsuite"
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
	errs := DB.StoreEngine("Aria").AriaSetting(true, false, false, "PAGE").Sync2(
		new(Testsuite),
	)
	if errs != nil {
		log.Fatalf("数据库启动失败%v", errs)
	}
}

type Testsuite struct {
	Id             int32 `db:"auto_increment;pk"`
	Product        int32
	Name           string `db:"type:varchar(255)"`
	Desc           string `db:"type:text"`
	Type           string `db:"type:varchar(20)"`
	AddedBy        string `db:"type:varchar(30)"`
	AddedDate      time.Time
	LastEditedBy   string `db:"type:varchar(30)"`
	LastEditedDate time.Time
	Deleted        bool
}

func (*Testsuite) TableName() string {
	return TABLE_TESTSUITE
}
