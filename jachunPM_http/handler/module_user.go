package handler

import (
	"encoding/hex"
	"errors"
	"fmt"
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
	httpHandlerMap["GET"]["/user/edit"] = get_user_edit
	httpHandlerMap["POST"]["/user/edit"] = post_user_edit
	httpHandlerMap["POST"]["/user/create"] = post_user_edit
	httpHandlerMap["GET"]["/user/delete"] = get_user_delete
	httpHandlerMap["POST"]["/user/delete"] = post_user_delete
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
	deptList, err := dept_getOptionMenu(0)
	if err != nil {
		data.ws.OutErr(errors.New(data.Lang["dept"]["err"].(map[string]string)[protocol.Err_DeptNotFound.String()]))
		return
	}

	roleGroup := make(map[string]string)
	for _, v := range data.Lang["user"]["roleList"].([]protocol.HtmlKeyValueStr) {
		roleGroup[v.Key] = v.Key
	}
	data.Data["roleGroup"] = roleGroup
	data.Data["groupList"], _ = user_getGroupOptionMenu()
	data.Data["deptID"] = data.ws.Query("dept")
	data.Data["depts"] = deptList
	out := protocol.GET_MSG_USER_GET_LoginSalt()
	out.Name = data.App["user"].(protocol.MSG_USER_INFO_cache).Account
	res, err := HostConn.SendMsgWaitResultToDefault(out)
	out.Put()
	if err == nil {
		if resdata, ok := res.(*protocol.MSG_USER_GET_LoginSalt_result); ok {
			session := data.ws.Session()
			r := rand.Int63()
			session.Set("edit_rand", r)
			data.Data["salt"] = resdata.Salt
			data.Data["rand"] = strconv.Itoa(int(r))
		} else {
			err = errors.New("请求salt返回消息错误，返回" + reflect.TypeOf(res).Elem().String())
		}
	}
	if err != nil {
		data.ws.OutErr(err)
	}
	templateOut("user.create.html", data)
	return
}
func get_user_edit(data *TemplateData) (action gnet.Action) {
	userID, _ := strconv.Atoi(data.ws.Query("userID"))
	userInfo := getUserCacheById(int32(userID))
	if userInfo == nil {
		data.ws.OutErr(errors.New(data.Lang["user"]["error"].(map[string]string)[protocol.Err_UserInfoNotFound.String()]))
		return
	}
	deptList, err := dept_getOptionMenu(0)
	if err != nil {
		data.ws.OutErr(errors.New(data.Lang["dept"]["err"].(map[string]string)[protocol.Err_DeptNotFound.String()]))
		return
	}
	data.Data["groups"], _ = user_getGroupOptionMenu()
	data.Data["userGroups"] = userInfo.Group
	data.Data["depts"] = deptList
	data.Data["user"] = userInfo
	out := protocol.GET_MSG_USER_GET_LoginSalt()
	out.Name = data.App["user"].(protocol.MSG_USER_INFO_cache).Account
	res, err := HostConn.SendMsgWaitResultToDefault(out)
	out.Put()
	if err == nil {
		if resdata, ok := res.(*protocol.MSG_USER_GET_LoginSalt_result); ok {
			session := data.ws.Session()
			r := rand.Int63()
			session.Set("edit_rand", r)
			data.Data["salt"] = resdata.Salt
			data.Data["rand"] = strconv.Itoa(int(r))
		} else {
			err = errors.New("请求salt返回消息错误，返回" + reflect.TypeOf(res).Elem().String())
		}
	}
	if err != nil {
		data.ws.OutErr(err)
	}
	templateOut("user.edit.html", data)
	return
}
func user_getGroupOptionMenu() (optionList []protocol.HtmlKeyValueStr, err error) {
	optionList = []protocol.HtmlKeyValueStr{{"", ""}}
	res, err := HostConn.CacheGetPath(protocol.UserServerNo, protocol.PATH_USER_GROUP_CACHE)
	if err != nil {
		return optionList, err
	}
	var groupMap = make(map[int32]*protocol.MSG_USER_Group_cache)
	var ids []int32
	buf := bufpool.Get().(*libraries.MsgBuffer)
	for _, b := range res {
		buf.Reset()
		buf.Write(b)
		if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_USER_Group_cache); ok {
			groupMap[v.Id] = v
			ids = append(ids, v.Id)
		}
	}
	libraries.SortInt32(ids)
	for _, id := range ids {
		v := groupMap[id]
		optionList = append(optionList, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Name})
		v.Put()
	}
	buf.Reset()
	bufpool.Put(buf)
	return
}
func user_getGroupListByIds(ids []int32) (res []*protocol.MSG_USER_Group_cache, err error) {
	for _, id := range ids {
		var group *protocol.MSG_USER_Group_cache
		err = HostConn.CacheGet(protocol.UserServerNo, protocol.PATH_USER_GROUP_CACHE, strconv.Itoa(int(id)), &group)
		if err != nil {
			return
		}
		res = append(res, group)
	}
	return
}
func post_user_edit(data *TemplateData) (action gnet.Action) {
	userID, _ := strconv.Atoi(data.ws.Post("userID"))
	if userID < 0 {
		data.ajaxResult(false, map[string]string{"account": data.Lang["user"]["error"].(map[string]string)["NotFoundUserInfo"]}, createLink("company", "browse", nil))
		return
	}
	password1 := data.ws.Post("password1")
	password2 := data.ws.Post("password2")
	if password1 != password2 {
		data.ajaxResult(false, map[string]string{"password1": data.Lang["user"]["error"].(map[string]string)["passwordsame"], "password2": data.Lang["user"]["error"].(map[string]string)["passwordsame"]}, "")
		return
	}
	msg, err := HostConn.GetMsg()
	if err != nil {
		data.ajaxResult(false, map[string]string{"verifyPassword": fmt.Sprintf(data.Lang["common"]["error"].(map[string]string)["ErrGetMsg"], err)}, "")
		return
	}
	out := protocol.GET_MSG_USER_CheckPasswd()
	out.UserId = data.App["user"].(protocol.MSG_USER_INFO_cache).Id
	out.Passwd = data.ws.Post("verifyPassword")
	session := data.ws.Session()
	out.Rand = session.Load_int64("edit_rand")
	res, err := msg.SendMsgWaitResult(0, out)
	out.Put()
	if err == nil {
		if resdata, ok := res.(*protocol.MSG_USER_CheckPasswd_result); ok {
			if resdata.Result == protocol.Success {
				update := make(map[string]string)
				for k, v := range data.ws.GetAllPost() {
					if k == "oldaccount" || k == "password1" || k == "password2" || k == "verifyPassword" || k == "passwordStrength" || k == "userID" {
						continue
					}
					if k == "groups" {
						var groups []int
						for _, vv := range v {
							id, _ := strconv.Atoi(vv)
							groups = append(groups, id)
						}
						update["group"] = libraries.JsonMarshalToString(groups)
						continue
					}
					update[k] = v[0]
				}
				if len(password1) > 0 {
					aesdata, _ := hex.DecodeString(password1)
					newpwd := string(libraries.AesCFBDecrypt(aesdata, []byte(data.ws.Post("verifyPassword")), []byte("jachunPM")))
					if !user_checkNewpasswd(newpwd, data) {
						return
					}
					salt := libraries.SHA256_S(strconv.Itoa(rand.Int()))
					update["salt"] = salt
					update["password"] = libraries.SHA256_S(newpwd + salt)
				}
				if update["account"] != "" {
					checkaccount := protocol.GET_MSG_USER_CheckAccount()
					checkaccount.Account = update["account"]
					res, err := msg.SendMsgWaitResult(0, checkaccount)
					result, ok := res.(*protocol.MSG_USER_CheckAccount_result)
					if result.Result != protocol.Success {
						data.ajaxResult(false, map[string]string{"account": data.Lang["user"]["error"].(map[string]string)[result.Result.String()]}, "")
						return
					}
					if err != nil || !ok {
						data.ajaxResult(false, map[string]string{"account": fmt.Sprintf(data.Lang["user"]["error"].(map[string]string)["ErrCheckaccount"], err)}, "")
						return
					}
					checkaccount.Put()
				}

				outupdate := protocol.GET_MSG_USER_INFO_updateByID()
				outupdate.UserID = int32(userID)
				outupdate.Update = update
				_, err := msg.SendMsgWaitResult(0, outupdate)
				if err != nil {
					data.ajaxResult(false, map[string]string{"verifyPassword": fmt.Sprintf(data.Lang["user"]["error"].(map[string]string)["ErrUpdate"], err)}, "")
					return
				}
				outupdate.Put()
				session.Delete("edit_rand")
				data.ajaxResult(true, data.Lang["common"]["saveSuccess"], createLink("company", "browse", nil))
				return
			} else {
				data.ajaxResult(false, map[string]string{"verifyPassword": data.Lang["user"]["error"].(map[string]string)[resdata.Result.String()]}, "")
				return
			}
		} else {
			err = errors.New("login登录返回消息错误，返回" + reflect.TypeOf(res).Elem().String())
		}
	}
	if err != nil {
		data.ajaxResult(false, err.Error(), "")
	}
	return
}
func user_checkNewpasswd(newpwd string, data *TemplateData) bool {
	if len(newpwd) < data.Config["user"]["common"]["weakPasswordlen"].(int) {
		data.ajaxResult(false, map[string]string{"password1": fmt.Sprintf(data.Lang["user"]["error"].(map[string]string)["weakPasswordlen"], data.Config["user"]["common"]["weakPasswordlen"]), "password2": fmt.Sprintf(data.Lang["user"]["error"].(map[string]string)["weakPasswordlen"], data.Config["user"]["common"]["weakPasswordlen"])}, "")
		return false
	}

	if data.Config["user"]["common"]["weakPasswordtype"].(int)&protocol.CONIFG_weakPasswordNum == protocol.CONIFG_weakPasswordNum {
		if !libraries.Preg_match(`\d+`, newpwd) {
			data.ajaxResult(false, map[string]string{"password1": data.Lang["user"]["error"].(map[string]string)["weakPasswordNum"], "password2": data.Lang["user"]["error"].(map[string]string)["weakPasswordNum"]}, "")
			return false
		}
	}
	if data.Config["user"]["common"]["weakPasswordtype"].(int)&protocol.CONIFG_weakPasswordLowerUpper == protocol.CONIFG_weakPasswordLowerUpper {
		if !libraries.Preg_match(`[a-z]+`, newpwd) || !libraries.Preg_match(`[A-Z]+`, newpwd) {
			data.ajaxResult(false, map[string]string{"password1": data.Lang["user"]["error"].(map[string]string)["weakPasswordLowerUpper"], "password2": data.Lang["user"]["error"].(map[string]string)["weakPasswordLowerUpper"]}, "")
			return false
		}
	}
	if data.Config["user"]["common"]["weakPasswordtype"].(int)&protocol.CONIFG_weakPasswordSpecial == protocol.CONIFG_weakPasswordSpecial {
		if !libraries.Preg_match(`[^\w]+`, newpwd) {
			data.ajaxResult(false, map[string]string{"password1": data.Lang["user"]["error"].(map[string]string)["weakPasswordSpecial"], "password2": data.Lang["user"]["error"].(map[string]string)["weakPasswordSpecial"]}, "")
			return false
		}
	}
	return true
}
func get_user_delete(data *TemplateData) (action gnet.Action) {
	out := protocol.GET_MSG_USER_GET_LoginSalt()
	out.Name = data.App["user"].(protocol.MSG_USER_INFO_cache).Account
	res, err := HostConn.SendMsgWaitResultToDefault(out)
	out.Put()
	if err == nil {
		if resdata, ok := res.(*protocol.MSG_USER_GET_LoginSalt_result); ok {
			session := data.ws.Session()
			r := rand.Int63()
			session.Set("delete_rand", r)
			data.Data["salt"] = resdata.Salt
			data.Data["rand"] = strconv.Itoa(int(r))
		} else {
			err = errors.New("请求salt返回消息错误，返回" + reflect.TypeOf(res).Elem().String())
		}
	}
	if err != nil {
		data.ws.OutErr(err)
	}
	templateOut("user.delete.html", data)
	return
}
func post_user_delete(data *TemplateData) (action gnet.Action) {
	out := protocol.GET_MSG_USER_CheckPasswd()
	out.UserId = data.App["user"].(protocol.MSG_USER_INFO_cache).Id
	out.Passwd = data.ws.Post("verifyPassword")
	session := data.ws.Session()
	out.Rand = session.Load_int64("delete_rand")
	deleteId, _ := strconv.Atoi(data.ws.Query("userID"))
	out.DeleteID = int32(deleteId)
	res, err := HostConn.SendMsgWaitResultToDefault(out)
	out.Put()
	if err == nil {
		if resdata, ok := res.(*protocol.MSG_USER_CheckPasswd_result); ok {
			if resdata.Result == protocol.Success {
				session.Delete("delete_rand")
				data.ajaxResult(true, "", "", `top.closetrigger()`)
				return
			}
		}
	}

	return
}
