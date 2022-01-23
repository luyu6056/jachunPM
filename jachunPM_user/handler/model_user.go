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
		if user.Deleted {
			value += user.Account
		} else {
			if strings.Index(params, "account") != -1 {
				value += user.Account
			} else {
				if user.Realname == "" {
					value += user.Account
				} else {
					value += user.Realname
				}
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
func user_updateCacheById(uid int32) {
	if user, err := getUserInfoByID(uid);err!=nil{
		libraries.DebugLog("%+v",err)
	}else if user != nil {
		user_setCache(user)
	}
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
	cache.AclMenu = make(map[string]bool)
	cache.LimitedProjects = make(map[int32]bool, len(user.LimitedProjects))
	cache.IsAdmin = user.Id == 1
	for k, v := range user.AclProducts {
		cache.AclProducts[k] = v
	}
	for k, v := range user.AclProjects {
		cache.AclProjects[k] = v
	}
	for k, v := range user.LimitedProjects {
		cache.LimitedProjects[k] = v
	}
	for k, v := range user.AclMenu {
		cache.AclMenu[k] = v
	}
	var configs []*db.Config
	HostConn.DB.Table(db.TABLE_Config).Prepare().Where("Uid=? or Uid=0", user.Id).Select(&configs)
	cache.Config = make(map[string]map[string]map[string]string)
	for _, config := range configs {
		if cache.Config[config.Module] == nil {
			cache.Config[config.Module] = make(map[string]map[string]string)
		}
		if cache.Config[config.Module][config.Section] == nil {
			cache.Config[config.Module][config.Section] = make(map[string]string)
		}
		cache.Config[config.Module][config.Section][config.Key] = config.Value
	}
	cache.Priv = make(map[string]map[string]bool)
	for module, v := range user.Priv {
		if cache.Priv[module] == nil {
			cache.Priv[module] = make(map[string]bool)
		}
		for method := range v {
			cache.Priv[module][method] = true
		}
	}
	HostConn.CacheSet(protocol.PATH_USER_INFO_CACHE, strconv.Itoa(int(user.Id)), cache, 0)
}
func user_insertMap(insert map[string]interface{}) error {
	id, err := HostConn.DB.Table(db.TABLE_USER).Insert(insert)
	if id > 0 {
		user_updateCacheById(int32(id))
	}
	return err
}
func user_getUserInfo(where map[string]interface{}) (users []*db.User, err error) {
	err = HostConn.DB.Table(db.TABLE_USER).Where(where).Limit(0).Select(&users)
	return
}
func updateUserView(uids, groupIds, products, projects []int32, in *protocol.Msg) {
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer func() {
		if err != nil {
			session.Rollback()
		} else {
			session.Commit()
		}
		in.WriteErr(err)
	}()

	if len(uids) == 0 && len(groupIds) == 0 {
		return
	}
	userM := make(map[int32]*db.User)
	var matchIds []int32
	var u []*db.User
	err = in.DB.Table(db.TABLE_USER).Field("Id,`Group`,AclProducts,AclMenu").Limit(0).Select(&u)
	if err != nil {
		return
	}
	var groups []*db.Group
	err = in.DB.Table(db.TABLE_GROUP).Limit(0).Select(&groups)
	if err != nil {
		return
	}
	//获取匹配user
	for _, user := range u {
		for _, id := range uids {
			if user.Id == id {
				userM[user.Id] = user
			}
		}
		for _, id1 := range groupIds {
			for _, id2 := range user.Group {
				if id1 == id2 {
					userM[user.Id] = user
				}
			}
		}
		for _, id := range products {
			if user.AclProducts[id] {
				userM[user.Id] = user
			}
		}
		for _, id := range projects {
			if user.AclProjects[id] {
				userM[user.Id] = user
			}
		}
	}
	projectCache := make(map[int32]*protocol.MSG_PROJECT_project_cache)
	productCache := make(map[int32]*protocol.MSG_PROJECT_product_cache)
	for id, user := range userM {

		if user.AclProjects == nil {
			user.AclProjects = make(map[int32]bool)
		}
		if user.AclProducts == nil {
			user.AclProducts = make(map[int32]bool)
		}
		user.LimitedProjects = make(map[int32]bool)
		//检查product权限，先把受影响的product加上
		for _, productID := range products {
			user.AclProducts[productID] = true
		}
		for productID := range user.AclProducts {
			find := false
			if productCache[productID] == nil {
				productCache[productID] = HostConn.GetProductById(productID)
			}
			if product := productCache[productID]; product != nil {
				//从负责人创建人寻找
				if id == product.CreatedBy || id == product.PO || id == product.QD || id == product.RD {
					find = true
				}

				if !find {
					//从白名单group寻找
					if product.Acl == "custom" {
						for _, whiteID := range product.Whitelist {
							for _, groupID := range user.Group {
								if groupID == whiteID {
									find = true
								}
							}
						}
					}
				}
			}
			//没找到删除
			if !find {
				delete(user.AclProducts, productID)
			}
		}
		//检查project权限
		for _, projectID := range projects {
			user.AclProjects[projectID] = true
		}

		for projectID := range user.AclProjects {
			find := false
			if projectCache[projectID] == nil {
				projectCache[projectID] = HostConn.GetProjectById(projectID)
			}
			if project := projectCache[projectID]; project != nil {
				//从负责人创建人寻找
				if id == project.OpenedBy || id == project.PO || id == project.PM || id == project.QD || id == project.RD {
					find = true
				}
				if !find {
					//从团队寻找
					for _, t := range project.Teams {
						if t.Uid == id {
							if t.Limited == "yes" {
								user.LimitedProjects[project.Id] = true
							}
							find = true
						}
					}
				}
				if !find {
					//从白名单group寻找
					if project.Acl == "custom" {
						for _, whiteID := range project.Whitelist {
							for _, groupID := range user.Group {
								if groupID == whiteID {
									find = true
								}
							}
						}
					}
				}
				//把对应的product权限加上
				if find {
					for _, product := range project.Products {
						user.AclProducts[product] = true
					}
				}
			}
			//没找到删除
			if !find {
				delete(user.AclProjects, projectID)
			}
		}
		//叠加group权限
		user.AclMenu = make(map[string]bool)
		user.Priv = make(map[string]map[string]bool)
		for _, groupID := range user.Group {
			for _, group := range groups {
				if groupID == group.Id {
					for _, projectID := range group.AclProjects {
						user.AclProjects[projectID] = true
					}
					for _, product := range group.AclProducts {
						user.AclProducts[product] = true
					}
					for _, name := range group.Acl {
						user.AclMenu[name] = true
					}
					for module, v := range group.Priv {
						if user.Priv[module] == nil {
							user.Priv[module] = make(map[string]bool)
						}
						for method := range v {
							user.Priv[module][method] = true
						}
					}
				}
			}
		}
		//特殊权限
		if user.Priv["company"] == nil {
			user.Priv["company"] = make(map[string]bool)
		}
		if user.Priv["action"] == nil {
			user.Priv["action"] = make(map[string]bool)
		}
		user.Priv["company"]["dynamic"] = true
		user.Priv["action"]["editcomment"] = true
		user.Priv["action"]["comment"] = true
		//更新

		matchIds = append(matchIds, id)
		_, err = in.DB.Table(db.TABLE_USER).Where("Id=?", user.Id).Update(map[string]interface{}{
			"AclProducts": user.AclProducts,
			"AclProjects": user.AclProjects,
			"AclMenu":     user.AclMenu,
			"Priv":        user.Priv,
		})
		if err != nil {
			return
		}

	}
	session.CommitCallback(func() {
		var users []*db.User
		in.DB.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": matchIds}).Limit(0).Select(&users)
		for _, user := range users {
			user_setCache(user)
		}
		if len(projects) > 0 {
			out := protocol.GET_MSG_PROJECT_updateCache()
			out.Type = "project"
			out.Ids = projects
			if msg, err := HostConn.GetMsg(); err == nil {
				msg.SendMsg(0, out)
			}
			out.Put()
		}
	})

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
func user_getUserqueryByWhere(data *protocol.MSG_USER_user_getUserqueryByWhere, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_user_getUserqueryByWhere_result()
	if err := in.DB.Table(db.TABLE_USERQUERY).Where(data.Where).Limit(0).Select(&out.List); err != nil {
		in.WriteErr(err)
	} else {
		in.SendResult(out)
	}
	return
}

func user_getExportTemplate(data *protocol.MSG_USER_getExportTemplate, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_getExportTemplate_result()
	err := in.DB.Table(db.TABLE_Usertpl).Where("(Uid=? or Public=1) and Type=?", data.Uid, "export"+data.Module).Select(&out.List)
	if err != nil {
		in.WriteErr(err)
	} else {
		in.SendResult(out)
	}
	out.Put()
}
