package handler

import (
	"fmt"
	"jachunPM_http/js"
	"libraries"
	"protocol"
	"runtime/debug"
	"strings"

	"github.com/luyu6056/cache"
	"github.com/luyu6056/gnet"
)

type HttpRequest interface {
	Body() []byte
	RemoteAddr() string
	IP() string
	Path() string
	Query(key string) string
	Post(key string) string
	PostSlice(key string) []string
	GetAllPost() map[string][]string
	GetAllQuery() map[string][]string
	AddQuery(name, value string)
	Cookie(key string) string
	Session() *cache.Hashvalue

	Header(name string) string
	Method() string
	//writer部分
	SetCode(int)
	SetContentType(string)
	SetCookie(name, value string, max_age uint32)
	SetHeader(name, value string)
	StaticHandler() gnet.Action
	Write(*libraries.MsgBuffer)
	WriteString(string)
	Redirect(url string)
	DelSession()
	Close()
}

var httpHandlerMap = map[string]map[string]func(data *TemplateData) error{
	"GET":  make(map[string]func(data *TemplateData) error),
	"POST": make(map[string]func(data *TemplateData) error),
}
var httpHandlerModuleInit = map[string]map[string]func(data *TemplateData) error{
	"GET":  make(map[string]func(data *TemplateData) error),
	"POST": make(map[string]func(data *TemplateData) error),
}

func HttpHandler(ws HttpRequest) gnet.Action {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	if m, ok := httpHandlerMap[ws.Method()]; ok {
		if f, ok := m[ws.Path()]; ok {
			//检查是否登录
			data := templateDataInit(ws)
			if data.User == nil {
				if !strings.Contains("/user/login|/user/getsalt", ws.Path()) {
					if strings.Contains(ws.Path(), "onlyBody=yes") {
						ws.WriteString(js.Location(createLink("user", "login", nil), "parent"))
						return gnet.None
					}
					if strings.Contains(data.ws.Header("X-Requested-With"), "XMLHttpRequest") || strings.Contains(data.ws.Header("x-requested-with"), "XMLHttpRequest") {
						data.ajaxResult(false, data.Lang["user"]["relogin"], createLink("user", "login", nil))
						return gnet.None
					}
					ws.Redirect(createLink("user", "login", nil))
					return gnet.None
				}
			}
			//检查权限
			//执行moduleInit
			if init, ok := httpHandlerModuleInit[ws.Method()][data.App["moduleName"].(string)]; ok {
				if err := init(data); err != nil {
					data.outErr(err)
					return gnet.None
				}
			}

			//执行路由
			if err := f(data); err != nil {
				data.outErr(err)
			}
			if data.User != nil {
				data.User.Put()
			}
			return gnet.None
		}
	}
	return ws.StaticHandler()

}

func getClientLang(ws HttpRequest) protocol.CountryNo {
	if ws.Cookie("sessionID") == "" {
		return protocol.DefaultLang
	}
	session := ws.Session()
	client := protocol.CountryNo(session.Load_str("ClientLang"))
	if client.String() == "" {
		client = protocol.DefaultLang
	}
	return client
}
