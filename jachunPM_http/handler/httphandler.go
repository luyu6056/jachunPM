package handler

import (
	"fmt"
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
}

var httpHandlerMap = map[string]map[string]func(data *TemplateData) gnet.Action{
	"GET":  make(map[string]func(data *TemplateData) gnet.Action),
	"POST": make(map[string]func(data *TemplateData) gnet.Action),
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
					ws.Redirect(createLink("user", "login", nil))
					return gnet.None
				}
			}
			//检查权限
			action := f(data)
			if data.User != nil {
				data.User.Put()
			}
			return action
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
