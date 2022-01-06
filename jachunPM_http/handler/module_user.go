package handler

import (
	"encoding/hex"
	"errors"
	"fmt"
	"jachunPM_http/config"
	"jachunPM_http/js"
	"libraries"
	"math/rand"
	"protocol"
	"strconv"
	"strings"
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
	httpHandlerMap["GET"]["/user/restore"] = get_user_restore
	httpHandlerMap["GET"]["/user/ajaxGetContactUsers"] = get_user_ajaxGetContactUsers
}
func get_user_login(data *TemplateData) (err error) {
	//检查是否登录
	ws := data.ws
	if data.User != nil {
		ws.Redirect(createLink("company", "browse", nil))
		return
	}
	ws.Session()
	data.Data["keepLogin"] = ""
	data.Data["referer"] = ws.Header("Referer")
	data.Data["title"] = data.Lang["user"]["login"].(string)
	templateOut("user.login.html", data)
	return
}
func get_user_logout(data *TemplateData) (err error) {
	data.ws.DelSession()
	data.ws.Redirect(createLink("user", "login", nil))
	return
}
func get_user_getsalt(data *TemplateData) (e error) {
	ws := data.ws
	name := strings.Trim(ws.Post("account"), " ")
	if name == "" {
		ws.WriteString(`{"error":"` + data.Lang["user"]["error"].(map[string]string)["loginFailed"] + `"}`)
		return
	}
	out := protocol.GET_MSG_USER_GET_LoginSalt()
	out.Name = name
	var resdata *protocol.MSG_USER_GET_LoginSalt_result
	err := data.SendMsgWaitResultToDefault(out, &resdata)
	out.Put()
	if err == nil {
		session := ws.Session()
		r := rand.Int63()
		session.Set("login_rand", r)
		ws.WriteString(`{"salt":"` + resdata.Salt + `","rand":"` + strconv.Itoa(int(r)) + `"}`)
	} else {
		libraries.ReleaseLog("login请求salt失败%v", err)
		ws.WriteString(`{"error":"` + err.Error() + `"}`)
	}
	resdata.Put()
	return
}
func post_user_login(data *TemplateData) (e error) {
	ws := data.ws

	name := strings.Trim(ws.Post("account"), " ")
	if name == "" {
		ws.WriteString(`{"error":"` + config.Lang[getClientLang(ws)]["user"]["loginFailed"].(string) + `"}`)
		return
	}
	out := protocol.GET_MSG_USER_CheckPasswd()
	out.Name = name
	out.Passwd = ws.Post("password")
	session := ws.Session()
	out.Rand = session.Load_int64("login_rand")
	session.Delete("login_rand")
	var resdata *protocol.MSG_USER_CheckPasswd_result
	err := data.SendMsgWaitResultToDefault(out, &resdata)
	out.Put()
	if err == nil {
		if resdata.Result == protocol.Success {

			session.Set("UserId", resdata.UserId)
			keepLogin := ws.Post("keepLogin")
			if keepLogin == "1" {
				session.Expire(protocol.SessionKeepLoginExpires)
				ws.SetCookie("sessionID", session.Load_str("sessionID"), protocol.SessionKeepLoginExpires)
			}
			referer := ws.Post("referer")
			if strings.Index(referer, config.Server.Origin) == -1 {
				referer = createLink("company", "browse", nil)
			}
			ws.WriteString(`{"locate":"` + referer + `"}`)
			return
		} else {
			err = errors.New(data.Lang["user"]["error"].(map[string]string)[resdata.Result.String()])
		}
	}
	ws.WriteString(`{"error":"` + err.Error() + `"}`)

	resdata.Put()
	return
}

func get_user_create(data *TemplateData) (err error) {
	deptList, err := dept_getOptionMenu(0)
	if err != nil {
		return errors.New(data.Lang["dept"]["err"].(map[string]string)[protocol.Err_DeptNotFound.String()])
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
	out.Name = data.User.Account
	var resdata *protocol.MSG_USER_GET_LoginSalt_result
	err = data.SendMsgWaitResultToDefault(out, &resdata)
	out.Put()
	if err == nil {
		session := data.ws.Session()
		r := rand.Int63()
		session.Set("edit_rand", r)
		data.Data["salt"] = resdata.Salt
		data.Data["rand"] = strconv.Itoa(int(r))
	} else {
		return
	}
	resdata.Put()
	templateOut("user.create.html", data)
	return
}
func get_user_edit(data *TemplateData) (err error) {
	userID, _ := strconv.Atoi(data.ws.Query("userID"))
	userInfo := HostConn.GetUserCacheById(int32(userID))
	if userInfo == nil {
		return errors.New(data.Lang["user"]["error"].(map[string]string)[protocol.Err_UserInfoNotFound.String()])
	}
	deptList, err := dept_getOptionMenu(0)
	if err != nil {
		return errors.New(data.Lang["dept"]["err"].(map[string]string)[protocol.Err_DeptNotFound.String()])
	}
	data.Data["groups"], _ = user_getGroupOptionMenu()
	data.Data["userGroups"] = userInfo.Group
	data.Data["depts"] = deptList
	data.Data["user"] = userInfo
	out := protocol.GET_MSG_USER_GET_LoginSalt()
	out.Name = data.User.Account
	var resdata *protocol.MSG_USER_GET_LoginSalt_result
	err = data.SendMsgWaitResultToDefault(out, &resdata)
	out.Put()
	if err == nil {
		session := data.ws.Session()
		r := rand.Int63()
		session.Set("edit_rand", r)
		data.Data["salt"] = resdata.Salt
		data.Data["rand"] = strconv.Itoa(int(r))
	} else {
		return
	}
	resdata.Put()
	templateOut("user.edit.html", data)
	return
}
func user_getGroupOptionMenu() (optionList []protocol.HtmlKeyValueStr, err error) {
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
func post_user_edit(data *TemplateData) (e error) {
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
	msg, err := data.GetMsg()
	if err != nil {
		data.ajaxResult(false, map[string]string{"verifyPassword": fmt.Sprintf(data.Lang["common"]["error"].(map[string]string)["ErrGetMsg"], err)}, "")
		return
	}
	out := protocol.GET_MSG_USER_CheckPasswd()
	out.UserId = data.User.Id
	out.Passwd = data.ws.Post("verifyPassword")
	session := data.ws.Session()
	out.Rand = session.Load_int64("edit_rand")
	var resdata *protocol.MSG_USER_CheckPasswd_result
	err = msg.SendMsgWaitResult(0, out, &resdata)
	out.Put()
	if err == nil {
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
				var result *protocol.MSG_USER_CheckAccount_result
				err = msg.SendMsgWaitResult(0, checkaccount, &result)
				if err != nil {
					data.ajaxResult(false, map[string]string{"account": fmt.Sprintf(data.Lang["user"]["error"].(map[string]string)["ErrCheckaccount"], err)}, "")
					return
				}
				if result.Result != protocol.Success {
					data.ajaxResult(false, map[string]string{"account": data.Lang["user"]["error"].(map[string]string)[result.Result.String()]}, "")
					return
				}
				checkaccount.Put()
				result.Put()
			}

			outupdate := protocol.GET_MSG_USER_INFO_updateByID()
			outupdate.UserID = int32(userID)
			outupdate.Update = update
			err := msg.SendMsgWaitResult(0, outupdate, nil)
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
		resdata.Put()
	} else {
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
func get_user_delete(data *TemplateData) (err error) {
	out := protocol.GET_MSG_USER_GET_LoginSalt()
	out.Name = data.User.Account
	var resdata *protocol.MSG_USER_GET_LoginSalt_result
	err = data.SendMsgWaitResultToDefault(out, &resdata)
	out.Put()
	if err == nil {
		session := data.ws.Session()
		r := rand.Int63()
		session.Set("delete_rand", r)
		data.Data["salt"] = resdata.Salt
		data.Data["rand"] = strconv.Itoa(int(r))
		resdata.Put()
	} else {
		return
	}
	templateOut("user.delete.html", data)
	return
}
func post_user_delete(data *TemplateData) (e error) {
	msg, err := data.GetMsg()
	if err != nil {
		data.ajaxResult(false, map[string]string{"verifyPassword": fmt.Sprintf(data.Lang["common"]["error"].(map[string]string)["ErrGetMsg"], err)}, "")
		return
	}
	out := protocol.GET_MSG_USER_CheckPasswd()
	out.UserId = data.User.Id
	out.Passwd = data.ws.Post("verifyPassword")
	session := data.ws.Session()
	out.Rand = session.Load_int64("delete_rand")
	deleteId, _ := strconv.Atoi(data.ws.Query("userID"))
	var resdata *protocol.MSG_USER_CheckPasswd_result
	err = msg.SendMsgWaitResult(0, out, &resdata)
	out.Put()
	if err == nil {
		if resdata.Result == protocol.Success {
			outupdate := protocol.GET_MSG_USER_INFO_updateByID()
			outupdate.UserID = int32(deleteId)
			outupdate.Update = map[string]string{
				"Deleted": "1",
			}
			err := msg.SendMsgWaitResult(0, outupdate, nil)
			if err != nil {
				data.ajaxResult(false, map[string]string{"verifyPassword": fmt.Sprintf(data.Lang["user"]["error"].(map[string]string)["ErrUpdate"], err)}, "")
				return
			}
			outupdate.Put()
			session.Delete("delete_rand")
			data.ajaxResult(true, "", "", `top.closetrigger(true)`)
			return
		}
		resdata.Put()
	} else {
		data.ajaxResult(false, err.Error())
	}

	return
}
func get_user_restore(data *TemplateData) (err error) {
	userID, _ := strconv.Atoi(data.ws.Query("userID"))
	outupdate := protocol.GET_MSG_USER_INFO_updateByID()
	outupdate.UserID = int32(userID)
	outupdate.Update = map[string]string{
		"Deleted": "0",
	}
	if err = data.SendMsgWaitResultToDefault(outupdate, nil); err != nil {
		return
	}
	outupdate.Put()
	data.ws.WriteString(js.Location("back", "_self"))
	return
}
func user_getPairs(data *TemplateData, params string, usersToAppended ...int32) ([]protocol.HtmlKeyValueStr, error) {
	out := protocol.GET_MSG_USER_getPairs()
	out.Params = params
	if len(usersToAppended) == 1 {
		out.UsersToAppended = usersToAppended[0]
	}
	var result *protocol.MSG_USER_getPairs_result
	if err := data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return nil, err
	}
	out.Put()
	return result.List, nil
}
func get_user_ajaxGetContactUsers(data *TemplateData) (e error) {
	users, err := user_getPairs(data, "devfirst|nodeleted")
	if err != nil {
		data.ws.WriteString(js.Alert(data.Lang["user"]["error"].(map[string]string)["NotFoundUserInfo"]))
		return
	}
	listID, _ := strconv.Atoi(data.ws.Query("listID"))
	if listID <= 0 {
		data.ws.WriteString(html_select("mailto", users, "", "class='form-control' multiple data-placeholder='"+data.Lang["common"]["chooseUsersToMail"].(string)+"'"))
		return
	}
	out := protocol.GET_MSG_USER_getContactListById()
	out.Id = int32(listID)
	var result *protocol.MSG_USER_getContactListById_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		data.ws.WriteString(js.Alert(err.Error()))
		return
	}
	data.ws.WriteString(html_select("mailto", users, result.Result.UserList, "class='form-control' multiple data-placeholder='"+data.Lang["common"]["chooseUsersToMail"].(string)+"'"))
	return
}
func user_getGroupGetPairs() (list []protocol.HtmlKeyValueStr, err error) {
	res, err := HostConn.CacheGetPath(protocol.UserServerNo, protocol.PATH_USER_GROUP_CACHE)
	if err != nil {
		return nil, err
	}

	var groups []*protocol.MSG_USER_Group_cache
	buf := bufpool.Get().(*libraries.MsgBuffer)
	for _, b := range res {
		buf.Reset()
		buf.Write(b)
		if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_USER_Group_cache); ok {
			groups = append(groups, v)
		}
	}
	protocol.Order_group(groups, nil)
	for _, v := range groups {
		list = append(list, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Name})
	}
	return
}

func user_getAllcache(data *TemplateData) (result []*protocol.MSG_USER_INFO_cache, err error) {
	if data.Data["user_getAllcache"] == nil {
		res, err := HostConn.CacheGetPath(protocol.UserServerNo, protocol.PATH_USER_INFO_CACHE)
		if err != nil {
			return nil, err
		}

		buf := bufpool.Get().(*libraries.MsgBuffer)
		for _, b := range res {
			buf.Reset()
			buf.Write(b)
			if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_USER_INFO_cache); ok {
				result = append(result, v)
			}
		}
		buf.Reset()
		bufpool.Put(buf)
		protocol.Order_user(result, nil)
		data.Data["user_getAllcache"] = result
	}
	return data.Data["user_getAllcache"].([]*protocol.MSG_USER_INFO_cache), nil
}
