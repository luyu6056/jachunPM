package handler

import (
	"jachunPM_user/db"
	"libraries"
	"protocol"
	"strconv"
	"strings"
)

func getLoginSalt(data *protocol.MSG_USER_GET_LoginSalt, msg *protocol.Msg) {

	var user *db.User
	err := msg.DB.Table(db.TABLE_USER).WhereOr(map[string]interface{}{"Account": data.Name, "Realname": data.Name, "Mobile": data.Name}).Find(&user)
	if err != nil {
		msg.WriteErr(err)
		return
	}
	res := protocol.GET_MSG_USER_GET_LoginSalt_result()
	if user == nil {
		res.Salt = libraries.SHA256_S(data.Name)
	} else {
		res.Salt = user.Salt
	}
	msg.SendResult(res)
	res.Put()
}
func checkPasswd(data *protocol.MSG_USER_CheckPasswd, msg *protocol.Msg) {
	var user *db.User
	var err error
	if data.UserId > 0 {
		user, err = getUserInfoByID(data.UserId)
	} else {
		err = msg.DB.Table(db.TABLE_USER).WhereOr(map[string]interface{}{"Account": data.Name, "Realname": data.Name, "Mobile": data.Name}).Find(&user)
	}

	if err != nil {
		msg.WriteErr(err)
		return
	}
	res := protocol.GET_MSG_USER_CheckPasswd_result()
	if user != nil && libraries.SHA256_S(user.Password+strconv.FormatInt(data.Rand, 10)) == data.Passwd {
		res.UserId = user.Id
		res.Result = protocol.Success
	} else {
		res.Result = protocol.Err_Password
	}
	msg.SendResult(res)
	res.Put()
}
func getUserInfoByID(uid int32) (userinfo *db.User, err error) {
	err = HostConn.DB.Table(db.TABLE_USER).Prepare().Where("Id=?", uid).Find(&userinfo)
	return
}
func getUserInfoByIDS(ids []int32) (users []*db.User, err error) {
	userlist := []*db.User{}
	err = HostConn.DB.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": ids}).Select(&userlist)
	return userlist, err

}
func user_getPairs(params string, usersToAppended int32, in *protocol.Msg) ([]protocol.HtmlKeyValueStr, error) {
	fields := "Id, Account, Realname, Deleted"
	if strings.Index(params, "pofirst") > -1 {
		fields += ", INSTR(',pd,po,', role) AS roleOrder"
	}
	if strings.Index(params, "pdfirst") > -1 {
		fields += ", INSTR(',po,pd,', role) AS roleOrder"
	}
	if strings.Index(params, "qafirst") > -1 {
		fields += ", INSTR(',qd,qa,', role) AS roleOrder"
	}
	if strings.Index(params, "qdfirst") > -1 {
		fields += ", INSTR(',qa,qd,', role) AS roleOrder"
	}
	if strings.Index(params, "pmfirst") > -1 {
		fields += ", INSTR(',td,pm,', role) AS roleOrder"
	}
	if strings.Index(params, "devfirst") > -1 {
		fields += ", INSTR(',td,pm,qd,qa,dev,', role) AS roleOrder"
	}
	orderBy := "account"
	if strings.Index(params, "first") > -1 {
		orderBy = "roleOrder DESC, account"
	}
	userconfig, err := in.LoadConfig("user")
	if err != nil {
		return nil, err
	}
	var userList []*db.User
	var conditions = make(map[string]interface{})
	if strings.Index(params, "nodeleted") > -1 || !userconfig["common"]["showDeleted"].(bool) {
		conditions["Deleted"] = false
	}
	if usersToAppended > 0 {
		conditions["Id"] = usersToAppended
	}

	err = HostConn.DB.Table(db.TABLE_USER).Field(fields).WhereOr(conditions).Order(orderBy).Select(&userList)
	if err != nil {
		return nil, err
	}
	res := []protocol.HtmlKeyValueStr{}
	/* Cycle the user records to append the first letter of his account. */
	for _, user := range userList {

		var value string
		if !strings.Contains(params, "noletter") {
			value = strings.ToUpper(user.Account[:1]) + ":"
		}
		if user.Deleted && strings.Index(params, "realname") == -1 {
			value += user.Account
		} else {
			if user.Realname == "" {
				value += user.Account
			} else {
				value += user.Realname
			}
		}
		res = append(res, protocol.HtmlKeyValueStr{strconv.Itoa(int(user.Id)), value})
	}

	if strings.Index(params, "noempty") == -1 {
		res = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"", ""}}, res...)
	}
	if strings.Index(params, "noclosed") == -1 {
		res = append(res, protocol.HtmlKeyValueStr{"closed", "Closed"})
	}
	return res, nil
}
func user_getCompanyUsers(data *protocol.MSG_USER_getCompanyUsers, in *protocol.Msg) {
	if data.Page < 1 {
		data.Page = 1
	}
	if data.PerPage < 1 {
		data.PerPage = 1
	}
	out := protocol.GET_MSG_USER_getCompanyUsers_result()
	defer out.Put()
	out.Total = data.Total
	if data.Total > 0 && (data.Page-1)*data.PerPage > data.Total {
		in.SendResult(out)
		return
	}
	data.Sort = strings.Replace(data.Sort, "_", " ", 1)
	var where interface{}
	if data.Type == "bydept" {
		if data.DeptID > 0 {
			ids, err := dept_getAllChildID(data.DeptID)
			if err != nil {
				in.WriteErr(err)
				return
			}
			where = map[string]interface{}{"dept": ids}
		}
	} else {
		if data.Where == "" {
			data.Where = "1 = 1"
		}
		where = data.Where
	}
	err := in.DB.Table(db.TABLE_USER).Where(where).Order("deleted asc,"+data.Sort).Limit(data.PerPage*(data.Page-1), data.PerPage).Select(&out.List)
	if err != nil {
		in.WriteErr(err)
		return
	}
	if data.Total <= 0 {
		out.Total, err = in.DB.Table(db.TABLE_USER).Where(where).Count()
		if err != nil {
			in.WriteErr(err)
			return
		}
	}
	in.SendResult(out)
	return
}
func user_updateMap(where map[string]interface{}, update map[string]interface{}) error {
	res, err := HostConn.DB.Table(db.TABLE_USER).Where(where).Update(update)
	if res {
		var users []*db.User
		HostConn.DB.Table(db.TABLE_USER).Where(where).Limit(0).Select(&users)
		for _, user := range users {
			user_setCache(user)
		}
	}
	return err
}
func user_setCache(user *db.User) {
	cache := protocol.GET_MSG_USER_INFO_cache()
	cache.Account = user.Account
	cache.AttendNo = user.AttendNo
	cache.ClientLang = user.ClientLang
	cache.Commiter = user.Commiter
	cache.Dept = user.Dept
	cache.Email = user.Email
	cache.Fails = user.Fails
	cache.Gender = user.Gender
	cache.Id = user.Id
	cache.Ip = user.Ip
	cache.Join = user.Join.Unix()
	cache.Last = user.Last.Unix()
	cache.Locked = user.Locked.Unix()
	cache.Mobile = user.Mobile
	cache.Realname = user.Realname
	cache.Role = user.Role
	cache.Visits = user.Visits
	cache.Deleted = user.Deleted
	cache.QQ = user.QQ
	cache.Group = user.Group
	cache.Weixin = user.Weixin
	cache.Address = user.Address
	cache.AclProducts = make(map[int32]bool, len(user.AclProducts))
	cache.AclProjects = make(map[int32]bool, len(user.AclProducts))
	cache.IsAdmin = user.Id == 1
	for k, v := range user.AclProducts {
		cache.AclProducts[k] = v
	}
	for k, v := range user.AclProjects {
		cache.AclProjects[k] = v
	}
	HostConn.CacheSet(protocol.PATH_USER_INFO_CACHE, strconv.Itoa(int(user.Id)), cache, 0)
	cache.Put()
}
func user_insertMap(insert map[string]interface{}) error {
	id, err := HostConn.DB.Table(db.TABLE_USER).Insert(insert)
	if id > 0 {
		user, _ := getUserInfoByID(int32(id))
		user_setCache(user)
	}
	return err
}
func user_getUserInfo(where map[string]interface{}) (users []*db.User, err error) {
	err = HostConn.DB.Table(db.TABLE_USER).Where(where).Limit(0).Select(&users)
	return
}
func updateUserView(data *protocol.MSG_USER_updateUserView, in *protocol.Msg) {
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer func() {
		session.Rollback()
		in.WriteErr(err)
	}()
	if (len(data.ProductIds) == 0 && len(data.ProjectIds) == 0) || (len(data.UserIds) == 0 && len(data.GroupIds) == 0) {
		return
	}

	var users []*db.User
	var matchIds []int32
	err = session.Table(db.TABLE_USER).Limit(0).Select(&users)
	if err != nil {
		return
	}
	if len(data.UserIds) > 0 {
		for _, user := range users {
			for _, id := range data.UserIds {
				if user.Id == id {
					matchIds = append(matchIds, user.Id)
					break
				}
			}
		}
	} else if len(data.GroupIds) > 0 {
		for _, user := range users {
		out:
			for _, id := range data.UserIds {
				for _, groupid := range user.Group {
					if id == groupid {
						matchIds = append(matchIds, user.Id)
						break out
					}
				}
			}
		}
	}
	var productIds, projectIds []int32
	//把需要删除权限的找出来
	for i := len(users) - 1; i >= 0; i-- {
		user := users[i]
		for _, productid := range data.ProductIds {
			if _, ok := user.AclProducts[productid]; ok {
				productIds = append(productIds, user.Id)
				delete(user.AclProducts, productid)
				break
			}

		}
		for _, projectid := range data.ProjectIds {
			if _, ok := user.AclProjects[projectid]; ok {
				projectIds = append(projectIds, user.Id)
				delete(user.AclProjects, projectid)
				break
			}
		}

	}
	//先把权限删除
	if len(productIds) > 0 {
		for _, productid := range data.ProductIds {
			_, err = session.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": productIds}).Update(map[string]interface{}{"AclProducts": []string{"exp", "json_remove(AclProducts, '$." + strconv.Itoa(int(productid)) + "' )"}})
			if err != nil {
				return
			}
		}

	}
	if len(projectIds) > 0 {
		for _, projectid := range data.ProjectIds {
			_, err = session.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": projectIds}).Update(map[string]interface{}{"AclProjects": []string{"exp", "json_remove(AclProjects, '$." + strconv.Itoa(int(projectid)) + "' )"}})
			if err != nil {
				return
			}
		}
	}
	//增加权限
	if len(matchIds) > 0 {
		for _, productid := range data.ProductIds {
			_, err = session.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": matchIds, "AclProducts": nil}).Update(map[string]interface{}{"AclProducts": `{}`})
			if err != nil {
				return
			}
			_, err = session.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": matchIds}).Update(map[string]interface{}{"AclProducts": []string{"exp", "json_set(AclProducts, '$." + strconv.Itoa(int(productid)) + "','true' )"}})
			if err != nil {
				return
			}
		}
		for _, projectid := range data.ProjectIds {
			_, err = session.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": matchIds, "AclProjects": nil}).Update(map[string]interface{}{"AclProjects": `{}`})
			if err != nil {
				return
			}
			_, err = session.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": matchIds}).Update(map[string]interface{}{"AclProjects": []string{"exp", "json_set(AclProjects, '$." + strconv.Itoa(int(projectid)) + "','true' )"}})
			if err != nil {
				return
			}
		}
	}
	session.CommitCallback(func() {
		var users []*db.User
		in.DB.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": matchIds}).Limit(0).Select(&users)
		for _, user := range users {
			user_setCache(user)
		}
	})
	session.Commit()
}
func user_getGlobalContacts() (globalContacts []*db.Usercontact, err error) {
	err = HostConn.DB.Table(db.TABLE_USERCONTACT).Prepare().Where(map[string]interface{}{"Share": true}).Limit(0).Select(&globalContacts)
	return
}
func user_getContactLists(data *protocol.MSG_USER_getContactLists, in *protocol.Msg) {
	var contacts []*db.Usercontact
	err := in.DB.Table(db.TABLE_USERCONTACT).Prepare().Where("Uid = ?", data.Uid).Limit(0).Select(&contacts)
	if err != nil {
		in.WriteErr(err)
		return
	}
	globalContacts, err := user_getGlobalContacts()
	if err != nil {
		in.WriteErr(err)
		return
	}
	for _, c1 := range globalContacts {
		find := false
		for _, c2 := range contacts {
			if c1.Id == c2.Id {
				find = true
				break
			}
		}
		if !find {
			contacts = append(contacts, c1)
		}
	}

	out := protocol.GET_MSG_USER_getContactLists_result()
	if len(contacts) != 0 {
		if strings.Contains(data.Params, "withempty") {
			out.List = []protocol.HtmlKeyValueStr{{}}
		}
		if strings.Contains(data.Params, "withnote") {
			userConfig, _ := in.LoadConfig("user")
			if userConfig != nil && userConfig["contacts"] != nil && userConfig["contacts"]["common"] != nil {
				out.List = []protocol.HtmlKeyValueStr{{userConfig["contacts"]["common"].(string), ""}}
			}
		}
		for _, c1 := range contacts {
			out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(c1.Id)), c1.ListName})
		}
	}
	in.SendResult(out)
	out.Put()
}
func user_getContactListByUid(data *protocol.MSG_USER_getContactListByUid, in *protocol.Msg) {
	var contacts []*db.Usercontact
	err := in.DB.Table(db.TABLE_USERCONTACT).Prepare().Where("Uid = ?", data.Uid).Limit(0).Select(&contacts)
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_USER_getContactListByUid_result()
	for _, c := range contacts {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(c.Id)), c.ListName})
	}
	in.SendResult(out)
	out.Put()
}

func user_getContactListById(data *protocol.MSG_USER_getContactListById, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_getContactListById_result()
	out.Result = protocol.GET_MSG_USER_ContactList()
	err := in.DB.Table(db.TABLE_USERCONTACT).Prepare().Where("Id = ?", data.Id).Limit(0).Select(&out.Result)
	if err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
func user_insertUpdateContactList(data *protocol.MSG_USER_insertUpdateContactList, in *protocol.Msg) {

	if data.Insert.Id == 0 {
		id, err := in.DB.Table(db.TABLE_USERCONTACT).Insert(data.Insert)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_USER_insertUpdateContactList_result()
		out.Id = int32(id)
		in.SendResult(out)
		out.Put()
	} else {
		err := in.DB.Table(db.TABLE_USERCONTACT).Replace(data.Insert)
		in.WriteErr(err)
		if err != nil {
			return
		}
	}

}
