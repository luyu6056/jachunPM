package db

import (
	"jachunPM_project/config"
	"log"
	"mysql"
)

const (
	TABLE_MODULE  = "module"
	TABLE_PRODUCT = "product"
	TABLE_DOCLIB  = "doclib"
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
		new(Doclib),
		new(Product),
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
	Deleted   bool   `db:"default(0)"` // 0=0,1=1,
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
	Acl         string  `db:"type:varchar(30)"` // 0=open,1=private,2=custom,
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
