package handler

import (
	"errors"
	"jachunPM_http/config"
	"libraries"
	"math/rand"
	"protocol"
	"reflect"
	"strconv"
	"strings"

	"github.com/luyu6056/gnet"
)

func init() {
	httpHandlerMap["POST"]["/user/getsalt"] = get_user_getsalt
	httpHandlerMap["GET"]["/user/login"] = get_user_login
	httpHandlerMap["POST"]["/user/login"] = post_user_login
	httpHandlerMap["GET"]["/"] = get_user_login
	httpHandlerMap["GET"]["/user/logout"] = get_user_logout
	httpHandlerMap["GET"]["/user/create"] = get_user_create
}
func get_user_login(data *TemplateData) gnet.Action {
	//检查是否登录
	ws := data.ws
	if data.App["user"] != nil {
		ws.Redirect(createLink("company", "browse", nil))
		return gnet.None
	}
	ws.Session()
	data.Data["keepLogin"] = ""
	data.Data["referer"] = ws.Header("Referer")
	data.Data["title"] = data.Lang["user"]["common"].(string) + data.Lang["common"]["colon"].(string) + data.Lang["user"]["todo"].(string)
	templateOut("user.login.html", data)
	return gnet.None
}
func get_user_logout(data *TemplateData) gnet.Action {
	data.ws.DelSession()
	data.ws.Redirect(createLink("user", "login", nil))
	return gnet.None
}
func get_user_getsalt(data *TemplateData) gnet.Action {
	ws := data.ws
	name := strings.Trim(ws.Post("account"), " ")
	if name == "" {
		ws.WriteString(`{"error":"` + config.Lang[getClientLang(ws)]["user"]["loginFailed"].(string) + `"}`)
		return gnet.None
	}
	out := protocol.GET_MSG_USER_GET_LoginSalt()
	out.Name = name
	res, err := HostConn.SendMsgWaitResultToDefault(out)
	out.Put()
	if err == nil {
		if resdata, ok := res.(*protocol.MSG_USER_GET_LoginSalt_result); ok {
			session := ws.Session()
			r := rand.Int63()
			session.Set("login_rand", r)
			ws.WriteString(`{"salt":"` + resdata.Salt + `","rand":"` + strconv.Itoa(int(r)) + `"}`)
		} else {
			err = errors.New("login请求salt返回消息错误，返回" + reflect.TypeOf(res).Elem().String())
		}
	}
	if err != nil {
		libraries.ReleaseLog("login请求salt失败%v", err)
		ws.WriteString(`{"error":"` + err.Error() + `"}`)
	}

	return gnet.None
}
func post_user_login(data *TemplateData) gnet.Action {
	ws := data.ws

	name := strings.Trim(ws.Post("account"), " ")
	if name == "" {
		ws.WriteString(`{"error":"` + config.Lang[getClientLang(ws)]["user"]["loginFailed"].(string) + `"}`)
		return gnet.None
	}
	out := protocol.GET_MSG_USER_CheckPasswd()
	out.Name = name
	out.Passwd = ws.Post("password")
	session := ws.Session()
	out.Rand = session.Load_int64("login_rand")
	session.Delete("login_rand")
	res, err := HostConn.SendMsgWaitResultToDefault(out)
	out.Put()
	if err == nil {
		if resdata, ok := res.(*protocol.MSG_USER_CheckPasswd_result); ok {
			if resdata.Result == protocol.Success {

				session.Set("UserId", resdata.UserId)
				keepLogin := ws.Post("keepLogin")
				if keepLogin == "1" {
					session.Expire(protocol.SessionKeepLoginExpires)
				} else {
					session.Expire(protocol.SessionTempExpires)
				}
				referer := ws.Post("referer")
				if strings.Index(referer, config.Server.Origin) == -1 {
					referer = createLink("company", "browse", nil)
				}
				ws.WriteString(`{"locate":"` + referer + `"}`)
				return gnet.None
			} else {
				err = errors.New(data.Lang["user"]["error"].(map[string]string)[resdata.Result.String()])
			}
		} else {
			err = errors.New("login登录返回消息错误，返回" + reflect.TypeOf(res).Elem().String())
		}
	}
	if err != nil {
		ws.WriteString(`{"error":"` + err.Error() + `"}`)
	}
	return gnet.None
}
func getUserCacheById(id int32) (user *protocol.MSG_USER_INFO_cache) {
	err := HostConn.CacheGet(protocol.UserServerNo, protocol.PATH_USER_INFO_CACHE, strconv.Itoa(int(id)), &user)
	if err != nil {
		libraries.DebugLog("获取user缓存失败%+v", err)
	}
	return
}
func get_user_create(data *TemplateData) (action gnet.Action) {

	templateOut("user.create.html", data)
	return
}
