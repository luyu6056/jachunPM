package protocol

import (
	"libraries"
	"sync"
)

const (
	MaxOutLen       = 1<<32 - 1
	MsgHeadLen      = 4 + 1 + 2 + 2 + 3 //不包含cmd
	MaxMsgLen       = 1<<24 - 1 + MsgHeadLen
	MaxMsgNum       = 127
	MaxMsgTtl       = 100 //目前只允许查询100次
	Rpcmsgnum       = 200 //缓冲消息数
	MaxReconnectNum = 0   //重试次数,0无限制
	//满足任一条件开启压缩
	CompressMinLen          = 8000      //最小压缩长度
	CompressMinNum          = 5         //最小压缩消息数
	DefaultWindowSize       = 20        //默认窗口值，为msg消息数量，窗口值允许为负值
	RpcTickDefaultTime      = 1         //单位秒
	MsgTimeOut              = 300       //单位秒
	SessionTempExpires      = 8 * 3600  //临时session
	SessionKeepLoginExpires = 7 * 86400 //keepLogin的session
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
	MaxServerNoNum  = 7
	HostServerNo    = 0
	CommomServerNo  = 0 //hostServer别名
	FileServerNo    = 1
	HttpServerNo    = 2
	LogServerNo     = 3
	UserServerNo    = 4
	ProjectServerNo = 5
	TestServerNo    = 6
	OaServerNo      = 7
	MailServerNo    = 8
)

type CountryNo string

const (
	//EN          CountryNo = "en"
	ZH_CN       CountryNo = "zh-cn"
	DefaultLang CountryNo = ZH_CN
)

var countryStr = map[CountryNo]string{
	//EN:    "English",
	ZH_CN: "简体中文",
}

func (c CountryNo) String() string {
	return countryStr[c]
}

//cache Path定义
const (
	PATH_USER_INFO_CACHE    = "user_cache"
	PATH_USER_COMPANY_CACHE = "company_cache"
	PATH_USER_DEPT_CACHE    = "dept_cache"
	PATH_CONFIG_CACHE       = "config_cache"
)

//格式化时间定义
const (
	TIMEFORMAT_MYSQLTIME = "2006-01-02 15:04:05"
	TIMEFORMAT_MYSQLDATE = "2006-01-02"
)
