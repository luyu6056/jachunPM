package db

import (
	"jachunPM_oa/config"
	"log"
	"mysql"
	"protocol"
	"time"
)

const (
	TABLE_Attend = "attend"
	TABLE_Attendstat = "attendstat"
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
	errs := db.StoreEngine("MyRocks").Sync2(
		new(Attend),
		new(Attendstat),
	)
	if errs != nil {
		log.Fatalf("数据库启动失败%v", errs)
	}
	db.Regsiter(&protocol.MSG_LOG_Action{}, &protocol.MSG_LOG_transformActions_info{})
	return db
}

type Attend struct{
	Id int32 `db:"auto_increment;pk"`
	Uid int32 `db:index`
	Account string `db:"type:varchar(30)"`
	Date time.Time `db:"index"`
	SignIn string `db:"type:time;not null"`
	SignOut string `db:"type:time;not null"`
	Status string `db:"type:varchar(30)"`
	Ip string `db:"type:varchar(15)"`
	Device string `db:"type:varchar(30)"`
	Client string `db:"type:varchar(20)"`
	ManualIn string `db:"type:time;not null"`
	ManualOut string `db:"type:time;not null"`
	Reason string `db:"type:varchar(30)"`
	Desc string `db:"type:text"`
	ReviewStatus string `db:"type:varchar(30)"`
	ReviewedBy string `db:"type:varchar(30)"`
	ReviewedDate time.Time `db:"not null"`
	EarlyMin int32 `db:"null"`
	LateMin int32 `db:"null"`
}
func (*Attend) TableName() string {
	return TABLE_Attend
}
type Attendstat struct{
	Id int32 `db:"auto_increment;pk"`
	Uid int32 `db:index`
	Account string `db:"type:varchar(30)"`
	Month string `db:"type:varchar(10)"`
	Normal float32 `db:"default(0)"`
	Late float32 `db:"default(0)"`
	Early float32 `db:"default(0)"`
	Absent float32 `db:"default(0)"`
	Trip float32 `db:"default(0)"`
	Egress float32 `db:"default(0)"`
	Lieu float32 `db:"default(0)"`
	PaidLeave float32 `db:"default(0)"`
	UnpaidLeave float32 `db:"default(0)"`
	TimeOvertime float32 `db:"default(0)"`
	RestOvertime float32 `db:"default(0)"`
	HolidayOvertime float32 `db:"default(0)"`
	Deserve float32 `db:"default(0)"`
	Actual float32 `db:"default(0)"`
	Status string `db:"type:varchar(30)"`
}
func (*Attendstat) TableName() string {
	return TABLE_Attendstat
}


