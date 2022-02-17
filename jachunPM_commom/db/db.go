package db

import (
	"jachunPM_commom/setting"
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
	errs := DB.StoreEngine("MyRocks").Sync2(
		new(Log_msg),
		new(ZstdMsg),
	)
	errs = append(errs, DB.StoreEngine("Innodb").Sync2(
		new(File),
	)...)
	if len(errs) > 0 {
		log.Fatalf("数据库启动失败%v", errs)
	}
	go UpdatedbInit()

	/*go func() {
		for {

			var zstd []*ZstdMsg
			DB.Table("zstd").Limit(0).Select(&zstd)
			for _, v := range zstd {

				ioutil.WriteFile("./db/zstd/"+v.Sha, v.Msg, 0666)
			}

		}
	}()*/

}

type Log_msg struct {
	Msgno     uint32 `db:"not null;pk"`
	Ttl       int    `db:"not null;pk"`
	TimeOut   uint16
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
	Id         int64  `db:"auto_increment;pk"`
	Pathname   string `db:"type:varchar(255)"`
	Title      string `db:"type:varchar(255)"`
	Extension  string `db:"type:varchar(30)"`
	Size       int64  `db:"default(0)"`
	ObjectType string `db:"type:varchar(30);index"`
	ObjectID   int32  `db:"not null;index"`
	AddedBy    int32
	AddedDate  time.Time `db:"not null"`
	Deleted    bool      `db:"default(0)"` // 0=0,1=1,
	Type       string    `db:"type:varchar(50)"`
}

func (*File) TableName() string {
	return TABLE_FILE
}

type ZstdMsg struct {
	Sha  string `db:"pk"`
	Cmd  int32  `db:index`
	Name string
	Msg  []byte `db:"type:mediumblob"`
}

func (*ZstdMsg) TableName() string {
	return "zstd"
}
