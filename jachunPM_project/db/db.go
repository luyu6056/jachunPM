package db

import (
	"jachunPM_project/config"
	"log"
	"mysql"
)

const (
	TABLE_MODULE = "module"
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
		new(Module),
	)
	if errs != nil {
		log.Fatalf("数据库启动失败%v", errs)
	}
}

type Module struct {
	Id        int32  `db:"auto_increment;pk"`
	Root      int32  `db:"default(0)"`
	Branch    int32  `db:"default(0)"`
	Name      string `db:"type:varchar(60)"`
	Parent    int32  `db:"default(0)"`
	Path      string `db:"type:varchar(255)"`
	Grade     int8   `db:"default(0)"`
	Order     int16  `db:"default(0)"`
	Type      string `db:"type:varchar(30)"`
	Owner     string `db:"type:varchar(30)"`
	Collector string `db:"type:text"`
	Short     string `db:"type:varchar(30)"`
	Deleted   int8   `db:"default(0)"` // 0=0,1=1,
}

func (*Module) TableName() string {
	return TABLE_MODULE
}
