package protocol

import (
	"libraries"
	"sync"
	"time"
)

const (
	MaxOutLen               = 2<<31 - 1
	MsgHeadLen              = 4 + 2 + 2 + 2 + 4 + 4 + 4 + 3 //不包含cmd
	MaxMsgLen               = 2<<24 - 1 - MsgHeadLen
	MaxMsgNum               = 127
	MaxMsgTtl               = 1000      //目前只允许查询1000次
	Rpcmsgnum               = 200       //缓冲消息数
	MaxReconnectNum         = 0         //重试次数,0无限制
	CompressMinNum          = 5         //最小压缩消息数
	CompressMinLen          = 4096      //最小压缩长度
	DefaultWindowSize       = 20        //默认窗口值，为msg消息数量，窗口值允许为负值
	RpcTickDefaultTime      = 1         //单位秒
	MsgTimeOut              = 60        //单位秒
	SessionTempExpires      = 8 * 3600  //临时session
	SessionKeepLoginExpires = 7 * 86400 //keepLogin的session
)

var ZEROTIME, _ = time.Parse("2006-01-02", "0000-01-01")
var NORMALTIME, _ = time.Parse("2006-01-02", "2010-01-01")

const (
	ADMINUSER  = 1
	CLOSEUSER  = -1
	SYSTEMUSER = -2 //在config用到
)

var buf_pool = sync.Pool{New: func() interface{} {
	return new(libraries.MsgBuffer)
}}

func BufPoolGet() *libraries.MsgBuffer {
	return buf_pool.Get().(*libraries.MsgBuffer)
}
func BufPoolPut(buf *libraries.MsgBuffer) {
	buf.Reset()
	buf_pool.Put(buf)
}

//定义服务序号No,允许0-127,redis将注册到对应的原始No+128
type ServerNo uint8

const (
	MaxServerNoNum = 8

	HostServerNo    = 1
	FileServerNo    = 1 //file集成在网关内
	HttpServerNo    = 2
	LogServerNo     = 3
	UserServerNo    = 4
	ProjectServerNo = 5
	TestServerNo    = 6
	OaServerNo      = 7
	MailServerNo    = 8
)

var serverNoToStr = map[ServerNo]string{
	HostServerNo:    "host",
	HttpServerNo:    "http",
	LogServerNo:     "log",
	UserServerNo:    "user",
	ProjectServerNo: "project",
	TestServerNo:    "test",
	OaServerNo:      "oa",
	MailServerNo:    "mail",
}

func (n ServerNo) String() string {
	if n > 127 {
		return serverNoToStr[n-127] + "_cache"
	}
	return serverNoToStr[n]
}

//多语言设置
type CountryNo string

const (
	//EN          CountryNo = "en"
	ZH_CN       CountryNo = "zh-cn"
	DefaultLang CountryNo = ZH_CN
)

var AllCountry = []CountryNo{ZH_CN}

var CountryToStr = map[CountryNo]string{
	//EN:    "English",
	ZH_CN: "简体中文",
}

func (c CountryNo) String() string {
	return CountryToStr[c]
}

//cache Path定义
const (
	PATH_USER_INFO_CACHE       = "user_cache"
	PATH_USER_COMPANY_CACHE    = "company_cache"
	PATH_USER_DEPT_CACHE       = "dept_cache"
	PATH_CONFIG_CACHE          = "config_cache"
	PATH_USER_GROUP_CACHE      = "group_cache"
	PATH_PROJECT_PRODUCT_CACHE = "product_cache"
	PATH_PROJECT_TREE_CACHE    = "tree_cache"
	PATH_PROJECT_PROJECT_CACHE = "project_cache"
)

//格式化时间定义
const (
	TIMEFORMAT_MYSQLTIME = "2006-01-02 15:04:05"
	TIMEFORMAT_MYSQLDATE = "2006-01-02"
)

//config相关杂项
const (
	CONIFG_weakPasswordAny        = 0
	CONIFG_weakPasswordLowerUpper = 1 << 1 //包含大小写
	CONIFG_weakPasswordSpecial    = 1 << 2 //包含特殊字符
	CONIFG_weakPasswordNum        = 1 << 3 //包含数字
)

var CMD_NO_CHECK_TTL = map[int32]int{
	CMD_MSG_FILE_RangeDown:        0,
	CMD_MSG_FILE_RangeDown_result: 0,
}

const (
	AttendAM = -1
	AttendPM = -2
)
