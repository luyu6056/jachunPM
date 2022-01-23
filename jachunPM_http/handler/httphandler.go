package handler

import (
	"codec"
	"errors"
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
	URI() string
	Referer() string

	Header(name string) string
	Method() string
	//writer部分
	SetCode(int)
	SetContentType(string)
	SetCookie(name, value string, max_age uint32)
	SetHeader(name, value string)
	StaticHandler() gnet.Action
	Write(*libraries.MsgBuffer) //异步输出，仅可调用一次
	WriteString(string)         //异步输出，仅可调用一次
	Redirect(url string)
	DelSession()
	Close()
	RangeDownload(r codec.HttpIoReader, size int64, name string) //文件下载用
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
	if m, ok := httpHandlerMap[ws.Method()]; ok {
		if f, ok := m[ws.Path()]; ok {
			//检查是否登录
			data := templateDataInit(ws, nil)
			if data.User == nil {
				if !strings.Contains("/user/login|/user/getsalt", ws.Path()) {
					data.ws.Session().Store("referer", data.ws.URI())
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
			defer func() {
				if err := recover(); err != nil {
					data.outErr(errors.New(fmt.Sprintf("执行错误,panic %v", err)))
					fmt.Println(err)
					debug.PrintStack()
				}
				if data.User != nil {
					data.User.Put()
				}
			}()
			moduleName := data.App["moduleName"].(string)
			methodName := data.App["methodName"].(string)
			if data.User != nil && !data.User.IsAdmin && moduleName == "user" && !(methodName == "logout" || methodName == "login") {
				//检查权限
				if !hasPriv(data, moduleName, methodName) {
					if v, ok := data.Lang[moduleName]["common"].(string); ok {
						moduleName = v
					}
					if v, ok := data.Lang[data.App["moduleName"].(string)][methodName].(string); ok {
						methodName = v
					}
					data.outErr(errors.New(fmt.Sprintf(data.Lang["user"]["error"].(map[string]string)["errorDeny"], moduleName, methodName)))
					return gnet.None
				}
				//检查视图
				if v, ok := data.Lang["menugroup"][moduleName].(string); ok {
					moduleName = v
				}
				if !data.User.AclMenu[moduleName] {
					data.outErr(errors.New(fmt.Sprintf(data.Lang["user"]["error"].(map[string]string)["errorView"], data.Lang[moduleName]["common"])))
					return gnet.None
				}

			}

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
