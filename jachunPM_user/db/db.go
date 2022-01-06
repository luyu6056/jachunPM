package db

import (
	"jachunPM_user/config"
	"log"
	"mysql"
	"protocol"
	"time"
)

const (
	TABLE_USER        = "user"
	TABLE_COMPANY     = "company"
	TABLE_DEPT        = "dept"
	TABLE_GROUP       = "group"
	TABLE_USERCONTACT = "usercontact"
	TABLE_TEAM        = "team"
	TABLE_USERQUERY   = "userquery"
	TABLE_Config      = "config"
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
	errs := db.StoreEngine("Innodb").Sync2(
		new(User),
		new(Company),
		new(Dept),
		new(Group),
		new(Usercontact),
		new(Team),
		new(Userquery),
		new(Config),
	)
	if errs != nil {
		log.Fatalf("数据库启动失败%v", errs)
	}
	db.Regsiter(&protocol.MSG_USER_team_info{}, &protocol.MSG_USER_INFO_cache{}, &protocol.MSG_USER_Group_cache{})
	return db
}

type User struct {
	Id          int32  `db:"auto_increment;pk"`
	Dept        int32  `db:"default(0)"`
	Account     string `db:"type:varchar(30)"`
	Salt        string `db:"type:varchar(64)"`
	Password    string `db:"type:varchar(64)"`
	Role        string `db:"type:varchar(10)"`
	Realname    string `db:"type:varchar(100)"`
	Group       []int32
	Commiter    string    `db:"type:varchar(100)"`
	Gender      int8      `db:"default(0)"` // 1男，0女
	Email       string    `db:"type:varchar(90)"`
	QQ          int64     `db:"type:varchar(20)"`
	Mobile      string    `db:"type:varchar(11)"`
	Phone       string    `db:"type:varchar(20)"`
	Weixin      string    `db:"type:varchar(90)"`
	Dingding    string    `db:"type:varchar(90)"`
	Address     string    `db:"type:varchar(120)"`
	Zipcode     string    `db:"type:varchar(10)"`
	Join        time.Time `db:"type:date;default('0000-00-00')"`
	Visits      int32     `db:"default(0)"`            //访问次数
	Ip          string    `db:"type:varchar(15)"`      //上次登录ip
	Last        time.Time `db:"default('0000-00-00')"` //上次登录时间
	Fails       int8      `db:"not null;default(0)"`   //密码错误次数
	Locked      time.Time `db:"not null;default('0000-00-00 00:00:00')"`
	Deleted     bool
	ClientLang  string          `db:"default('zh-cn');type:varchar(10)"`
	AclMenu     map[string]bool //允许访问的视图
	AclProducts map[int32]bool  //允许访问的产品
	AclProjects map[int32]bool  //允许访问的项目
	AttendNo    int32           `db:"null"` //打卡机编号
	//Birthday     time.Time `db:"not null;default('0000-00-00')"`
	//Skype        string    `db:"type:varchar(90)"`
	//Yahoo        string    `db:"type:varchar(90)"`
	//Gtalk        string    `db:"type:varchar(90)"`
	//Wangwang     string    `db:"type:varchar(90)"`
	//Slack        string    `db:"type:varchar(90)"`
	//Whatsapp     string    `db:"type:varchar(90)"`
	//Feedback     int8      `db:"default(0)"` // 0=0,1=1,
	//Ranzhi       string    `db:"type:varchar(30)"`
	//Ldap         string    `db:"type:varchar(30)"`
	//Score        int32     `db:"not null;default(0)"`
	//ScoreLevel   int32     `db:"not null;default(0)"`
	//ClientStatus int8      `db:"default(3)"` // 0=online,1=away,2=busy,3=offline,
	//Status       int8      `db:"default(3)"` // 0=online,1=away,2=busy,3=offline,
	TimeStamp time.Time `db:"default(current_timestamp());extra('on update current_timestamp()')"` //更新时间戳
}

func (*User) TableName() string {
	return TABLE_USER
}

type Company struct {
	Id        int32    `db:"auto_increment;pk"`
	Name      string   `db:"type:varchar(120)"`
	Phone     string   `db:"type:varchar(20)"`
	Fax       string   `db:"type:varchar(20)"`
	Address   string   `db:"type:varchar(120)"`
	Zipcode   string   `db:"type:varchar(10)"`
	Website   string   `db:"type:varchar(120)"`
	Backyard  string   `db:"type:varchar(120)"`
	Admins    []string `db:"type:varchar(255)"`
	Deleted   bool
	TimeStamp time.Time `db:"default(current_timestamp());extra('on update current_timestamp()')"` //更新时间戳
}

func (*Company) TableName() string {
	return TABLE_COMPANY
}

type Dept struct {
	Id        int32     `db:"auto_increment;pk"`
	Name      string    `db:"type:varchar(60)"`
	Parent    int32     `db:"default(0)"`
	Path      []int32   `db:"type:varchar(255)"`
	Grade     int8      `db:"default(0)"`
	Order     int8      `db:"default(0)"`
	Manager   int32     //负责人
	TimeStamp time.Time `db:"default(current_timestamp());extra('on update current_timestamp()')"` //更新时间戳
}

func (*Dept) TableName() string {
	return TABLE_DEPT
}

type Group struct {
	Id          int32                      `db:"auto_increment;pk"`
	Name        string                     `db:"type:varchar(30)"`
	Role        string                     `db:"type:varchar(30)"`
	Desc        string                     `db:"type:tinytext"`
	Acl         []string                   `db:"type:tinytext"`
	AclProducts []int32                    //额外允许访问的产品，-1为允许查看所有，group权限大于其他权限
	AclProjects []int32                    //额外允许访问的项目，-1为允许查看所有，group权限大于其他权限
	Developer   int8                       `db:"default(1)"` // 0=0,1=1,
	Priv        map[string]map[string]bool //访问权限map[Module][Method]
	TimeStamp   time.Time                  `db:"default(current_timestamp());extra('on update current_timestamp()')"` //更新时间戳

}

func (*Group) TableName() string {
	return TABLE_GROUP
}

type Usercontact struct {
	Id       int32   `db:"auto_increment;pk"`
	Uid      int32   `db:"type:varchar(30);index"`
	ListName string  `db:"type:varchar(60)"`
	UserList []int32 `db:"type:json"`
	Share    bool    `db:index`
}

func (*Usercontact) TableName() string {
	return TABLE_USERCONTACT
}

type Team struct {
	Id       int32     `db:"auto_increment;pk"`
	Root     int32     `db:"index"`
	Type     string    `db:"index"`
	Uid      int32     `db:"index"`
	Account  string    `db:"type:varchar(30)"`
	Role     string    `db:"type:varchar(30)"`
	Limited  string    `db:"default('no');type:varchar(8)"`
	Join     time.Time `db:"type:date"`
	Days     int16
	Hours    float64
	Estimate float64
	Consumed float64
	Left     float64
	Order    int8 `db:"not null;default(0)"`
}

func (*Team) TableName() string {
	return TABLE_TEAM
}

type Userquery struct {
	Id       int32  `db:"auto_increment;pk"`
	Uid      int32  `db:"index"`
	Module   string `db:"type:varchar(30)"`
	Title    string `db:"type:varchar(90)"`
	Form     string
	Sql      string
	Shortcut bool
}

func (*Userquery) TableName() string {
	return TABLE_USERQUERY
}

type Config struct {
	Id      int32  `db:"auto_increment;pk"`
	Uid     int32  `db:"index"`
	Module  string `db:"type:varchar(30)"`
	Key     string `db:"type:varchar(90)"`
	Value   string
	Section string
}

func (*Config) TableName() string {
	return TABLE_Config
}
