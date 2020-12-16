package server

import (
	"fmt"
	"libraries"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/luyu6056/cache"
)

//ctx不带服务类型应用，应该作为一个通用的ctx

type Context struct { //循环的Context
	In        *libraries.MsgBuffer
	In2       *libraries.MsgBuffer
	Log       []*Err_log
	Sql_build interface{}
	Buf       *libraries.MsgBuffer //辅助out用于序列化
	Conn      *ClientConn
	Conn_m    *sync.Map //rpc用
	//Transaction *mysql.Transaction //sql的事务
}

var ClientId int32

type ClientConn struct {
	Id          int32 //自增的ClientId
	ClientFd    [4]byte
	BeginTime   int64 //连接开始时间
	Session     *cache.Hashvalue
	IP          string
	UserAgent   string
	IsMobile    bool
	Output_data func(*libraries.MsgBuffer)
}

var ServerHand func(*Context)

type Err_log struct {
	Err       string
	Err_func  string
	Err_param string
}

/**
 * 输出
 *
 **/
type OutMsg interface {
	WRITE(buf *libraries.MsgBuffer)
}

func (c *Context) Output_data(msg OutMsg) {
	msg.WRITE(c.Buf)
	c.Conn.Output_data(c.Buf)
}

//暂时打印
func (c *Context) Save_errlog() {
	if len(c.Log) == 0 {
		return
	}

	for _, v := range c.Log {
		libraries.DebugLog("%v", v)
	}

}
func (c *Context) Adderr(err error, param interface{}) {
	if err != nil {
		if c == nil {
			c = &Context{}
		}
		s := libraries.Bytes2str(debug.Stack()[276:])
		if i := strings.Index(s, "bbs/controllers/web.init.0.func1"); i > -1 {
			s = s[:i]
		}
		if err.Error() == "" {
			if len(c.Log) > 0 {
				_, file, line, _ := runtime.Caller(1)
				c.Log = append(c.Log, &Err_log{Err: s + err.Error(), Err_func: fmt.Sprintf("%s,行%d:", file, line), Err_param: libraries.JsonMarshalToString(param)})
			}
		} else {
			_, file, line, _ := runtime.Caller(1)
			c.Log = append(c.Log, &Err_log{Err: s + err.Error(), Err_func: fmt.Sprintf("%s,行%d:", file, line), Err_param: libraries.JsonMarshalToString(param)})
		}

	}
}
