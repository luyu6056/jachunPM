package db

import (
	"jachunPM_oa/setting"
	"log"
	"mysql"
	"protocol"
	"time"
)

const (
	TABLE_ATTEND     = "attend"
	TABLE_ATTENDSTAT = "attendstat"
	TABLE_HOLIDAY    = "holiday"
	TABLE_LEAVE      = "leave"
	//TABLE_LIEU         = "lieu"
	TABLE_OVERTIME     = "overtime"
	TABLE_OVERTIMEBASE = "overtimebase"
	TABLE_TRIP         = "trip"
)

func Init() *mysql.MysqlDB {
	db, err := mysql.Open(setting.Setting.MysqlDsn)
	if err != nil {
		log.Fatalf("数据库连接失败 %v", err)
	}
	if setting.Setting.MysqlMaxConn > 0 {
		db.SetMaxOpenConns(setting.Setting.MysqlMaxConn)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("数据库启动失败 %v", err)
	}
	errs := db.StoreEngine("MyRocks").Sync2(
		new(Attend),
		new(Attendstat),
		new(Leave),
		//new(Lieu),
		new(Holiday),
		new(Overtime),
		new(Trip),
		new(OvertimeBase),
	)
	if errs != nil {
		log.Fatalf("数据库启动失败%v", errs)
	}
	db.Regsiter(&protocol.MSG_OA_attend_info{}, &protocol.MSG_LOG_transformActions_info{})
	return db
}

type Attend struct {
	Id           int32     `db:"auto_increment;pk"`
	Uid          int32     `db:"index"`
	Account      string    `db:"type:varchar(30);index"`
	Date         time.Time `db:"index"`
	SignIn       string    `db:"type:time;not null"`
	SignOut      string    `db:"type:time;not null"`
	Status       string    `db:"type:varchar(20);index"`
	Ip           string    `db:"type:varchar(15)"`
	Device       string    `db:"type:varchar(30)"`
	Client       string    `db:"type:varchar(20)"`
	ManualIn     string    `db:"type:time;not null"`
	ManualOut    string    `db:"type:time;not null"`
	Reason       string    `db:"type:varchar(30)"`
	Desc         string    `db:"type:text"`
	ReviewStatus string    `db:"type:varchar(30);index"`
	ReviewedBy   int32
	RejectDesc   string    `db:"type:varchar(255)"`
	ReviewedDate time.Time `db:"not null"`
	EarlyMin     int32     `db:"null"`
	LateMin      int32     `db:"null"`
}

func (*Attend) TableName() string {
	return TABLE_ATTEND
}

type Attendstat struct {
	Id              int32   `db:"auto_increment;pk"`
	Uid             int32   `db:"index"`
	Account         string  `db:"type:varchar(30)"`
	Month           string  `db:"type:varchar(10)"`
	Normal          float32 `db:"default(0)"`
	Late            float32 `db:"default(0)"`
	Early           float32 `db:"default(0)"`
	Absent          float32 `db:"default(0)"`
	Trip            float32 `db:"default(0)"`
	Egress          float32 `db:"default(0)"`
	Lieu            float32 `db:"default(0)"`
	PaidLeave       float32 `db:"default(0)"`
	UnpaidLeave     float32 `db:"default(0)"`
	TimeOvertime    float32 `db:"default(0)"`
	RestOvertime    float32 `db:"default(0)"`
	HolidayOvertime float32 `db:"default(0)"`
	Deserve         float32 `db:"default(0)"`
	Actual          float32 `db:"default(0)"`
	Status          string  `db:"type:varchar(30)"`
}

func (*Attendstat) TableName() string {
	return TABLE_ATTENDSTAT
}

type Holiday struct {
	Year   int32     `db:"index"`
	Date   time.Time `db:"type:date;pk"`
	Option string    `db:"type:varchar(10)"`
	Uid    int32
	Name   string
	//TimeStamp time.Time `db:"default(current_timestamp());extra('on update current_timestamp()')"`
}

func (*Holiday) TableName() string {
	return TABLE_HOLIDAY
}

type Leave struct {
	Id           int32     `db:"auto_increment;pk"`
	Year         int16     `db:"index"`
	Begin        time.Time `db:"index"`
	End          time.Time `db:"index"`
	Start        int8      `db:"not null"`
	Finish       int8      `db:"not null"`
	Day          float32   `db:"default(0)"`
	BackDate     time.Time `db:"not null"`
	Type         string    `db:"type:varchar(30)"`
	Desc         string    `db:"type:text"`
	Status       string    `db:"type:varchar(20);index"`
	CreatedBy    int32     `db:"index"`
	CreatedDate  time.Time `db:"not null"`
	ReviewedBy   int32     `db:"type:varchar(30)"`
	ReviewedDate time.Time `db:"not null"`
	//Level         int8      `db:"not null"`
	//AssignedTo    int32     `db:"type:varchar(30)"`
	//Reviewers     string    `db:"type:text"`
	//BackReviewers string    `db:"type:text"`
	//Noticeleader  string    `db:"type:varchar(255)"`
}

func (*Leave) TableName() string {
	return TABLE_LEAVE
}

/*type Lieu struct {
	Id           int32     `db:"auto_increment;pk"`
	Year         int16     `db:"index"`
	Begin        time.Time `db:"index"`
	End          time.Time `db:"index"`
	Start        int8      `db:"not null"`
	Finish       int8      `db:"not null"`
	Day          float32   `db:"default(0)"`
	Overtime     string    `db:"type:varchar(255)"`
	Trip         string    `db:"type:varchar(255)"`
	Desc         string    `db:"type:text"`
	Status       string    `db:"type:varchar(20);index"`
	CreatedBy    int32     `db:"index"`
	CreatedDate  time.Time `db:"not null"`
	ReviewedBy   int32     `db:"type:varchar(30)"`
	ReviewedDate time.Time `db:"not null"`
	Level        int8      `db:"not null"`
	AssignedTo   int32     `db:"type:varchar(30)"`
	Reviewers    string    `db:"type:text"`
	Noticeleader string    `db:"type:varchar(255)"`
}

func (*Lieu) TableName() string {
	return TABLE_LIEU
}*/

type Overtime struct {
	Id     int32     `db:"auto_increment;pk"`
	Year   int16     `db:"index"`
	Begin  time.Time `db:"index"`
	End    time.Time `db:"index"`
	Start  string    `db:"type:time;not null"`
	Finish string    `db:"type:time;not null"`
	Day    float32   `db:"default(0)"`
	//Leave        string    `db:"type:varchar(255)"`
	Type         string    `db:"type:varchar(30)"`
	Desc         string    `db:"type:text"`
	Status       string    `db:"type:varchar(20);index"`
	RejectReason string    `db:"type:varchar(100)"`
	CreatedBy    int32     `db:"index"`
	CreatedDate  time.Time `db:"not null"`
	ReviewedBy   int32     `db:"type:varchar(30)"`
	ReviewedDate time.Time `db:"not null"`
	//Level        int8      `db:"not null"`
	//AssignedTo   int32     `db:"type:varchar(30)"`
	//Reviewers    string    `db:"type:text"`
}
type OvertimeBase struct {
	Uid       int32 `db:"pk"`
	Account   string
	OffsetDay float64
}

func (*OvertimeBase) TableName() string {
	return TABLE_OVERTIMEBASE
}
func (*Overtime) TableName() string {
	return TABLE_OVERTIME
}

type Trip struct {
	Id           int32     `db:"auto_increment;pk"`
	Type         string    `db:"type:varchar(20)"` // 0=trip,1=egress,
	Customers    string    `db:"type:varchar(20)"`
	Name         string    `db:"type:varchar(30)"`
	Desc         string    `db:"type:text"`
	Status       string    `db:"type:varchar(20);index"`
	Year         int16     `db:"index"`
	Begin        time.Time `db:"index"`
	End          time.Time `db:"index"`
	Start        int8      `db:"not null"`
	Finish       int8      `db:"not null"`
	Day          float32   `db:"default(0)"`
	From         string    `db:"type:varchar(50)"`
	To           string    `db:"type:varchar(50)"`
	CreatedBy    int32     `db:"index"`
	CreatedDate  time.Time `db:"not null"`
	ReviewedBy   int32     `db:"type:varchar(30)"`
	ReviewedDate time.Time `db:"null"`
}

func (*Trip) TableName() string {
	return TABLE_TRIP
}
