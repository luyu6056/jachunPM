package db

import (
	"jachunPM_http/setting"
	"log"
	"mysql"
)

const (
	TABLE_SearchQuery = "searchQuery"
)

var DB *mysql.MysqlDB

func Init() {
	var err error
	DB, err = mysql.Open(setting.Setting.MysqlDsn)
	if err != nil {
		log.Fatalf("数据库连接失败 %v", err)
	}
	if setting.Setting.MysqlMaxConn > 0 {
		DB.SetMaxOpenConns(setting.Setting.MysqlMaxConn)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("数据库启动失败 %v", err)
	}

	errs := DB.StoreEngine("Innodb").Sync2(
		new(SearchQuery),
	)
	if len(errs) > 0 {
		log.Fatalf("数据库启动失败%v", errs)
	}

}

type SearchQuery struct {
	Id     int32  `db:"auto_increment;pk"`
	Uid    int32  `db:"index"`
	Module string `db:"index"`
	Title  string
	Where  string
}

func (*SearchQuery) TableName() string {
	return TABLE_SearchQuery
}
