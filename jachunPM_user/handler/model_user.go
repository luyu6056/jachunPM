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
	err := db.DB.Table(db.TABLE_USER).WhereOr(map[string]interface{}{"Account": data.Name, "Realname": data.Name, "Mobile": data.Name}).Find(&user)
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
		err = db.DB.Table(db.TABLE_USER).WhereOr(map[string]interface{}{"Account": data.Name, "Realname": data.Name, "Mobile": data.Name}).Find(&user)
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
	err = db.DB.Table(db.TABLE_USER).Prepare().Where("Id=?", uid).Find(&userinfo)
	return
}
func getUserInfoByIDS(ids []int32) (users []*db.User, err error) {
	userlist := []*db.User{}
	err = db.DB.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": ids}).Select(&userlist)
	return userlist, err

}
func user_getPairs(params string, usersToAppended int32) ([]protocol.HtmlKeyValueStr, error) {
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
	userconfig, err := HostConn.LoadConfig("user")
	if err != nil {
		return nil, err
	}
	var userList []*db.User
	var conditions = make(map[string]interface{})
	if strings.Index(params, "nodeleted") > -1 || !userconfig["common"]["showDeleted"].(bool) {
		conditions["Deleted"] = false
	}
	if usersToAppended > 0 {
		conditions["account"] = usersToAppended
	}

	err = db.DB.Table(db.TABLE_USER).Field(fields).WhereOr(conditions).Order(orderBy).Select(&userList)
	if err != nil {
		return nil, err
	}
	res := []protocol.HtmlKeyValueStr{}
	/* Cycle the user records to append the first letter of his account. */
	for _, user := range userList {
		//firstLetter = ucfirst(substr(account, 0, 1)) . ":";
		//if((strpos(params, "noletter") !== false) or (isset(this->config->isINT) and this->config->isINT)) firstLetter =  "";
		value := strings.ToUpper(user.Account[:1]) + ":"
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
	err := db.DB.Table(db.TABLE_USER).Where(where).Order("deleted asc,"+data.Sort).Limit(data.PerPage*(data.Page-1), data.PerPage).Select(&out.List)
	if err != nil {
		in.WriteErr(err)
		return
	}
	if data.Total <= 0 {
		out.Total, err = db.DB.Table(db.TABLE_USER).Where(where).Count()
		if err != nil {
			in.WriteErr(err)
			return
		}
	}
	in.SendResult(out)
	return
}
func user_updateMap(where map[string]interface{}, update map[string]interface{}) error {
	res, err := db.DB.Table(db.TABLE_USER).Where(where).Update(update)
	if res {
		var users []*db.User
		db.DB.Table(db.TABLE_USER).Where(where).Limit(0).Select(&users)
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
	id, err := db.DB.Table(db.TABLE_USER).Insert(insert)
	if id > 0 {
		user, _ := getUserInfoByID(int32(id))
		user_setCache(user)
	}
	return err
}
func user_getUserInfo(where map[string]interface{}) (users []*db.User, err error) {
	err = db.DB.Table(db.TABLE_USER).Where(where).Limit(0).Select(&users)
	return
}
func updateUserView(data *protocol.MSG_USER_updateUserView, in *protocol.Msg) {
	var err error
	defer func() {
		in.WriteErr(err)
	}()
	if (data.ProductId == 0 && data.ProjectId == 0) || (len(data.UserIds) == 0 && len(data.GroupIds) == 0) {
		return
	}

	var users []*db.User
	var matchIds []int32
	err = db.DB.Table(db.TABLE_USER).Limit(0).Select(&users)
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
		if data.ProductId > 0 {
			if _, ok := user.AclProducts[data.ProductId]; ok {
				productIds = append(productIds, user.Id)
				delete(user.AclProducts, data.ProductId)
				break
			}

		}
		if data.ProjectId > 0 {
			if _, ok := user.AclProjects[data.ProjectId]; ok {
				projectIds = append(projectIds, user.Id)
				delete(user.AclProjects, data.ProjectId)
				break
			}
		}

	}
	//先把权限删除
	if len(productIds) > 0 && data.ProductId > 0 {
		_, err = db.DB.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": productIds}).Update(map[string]interface{}{"AclProducts": []string{"exp", "json_remove(AclProducts, '$." + strconv.Itoa(int(data.ProductId)) + "' )"}})
		if err != nil {
			return
		}

	}
	if len(projectIds) > 0 && data.ProjectId > 0 {
		_, err = db.DB.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": projectIds}).Update(map[string]interface{}{"AclProjects": []string{"exp", "json_remove(AclProjects, '$." + strconv.Itoa(int(data.ProjectId)) + "' )"}})
		if err != nil {
			return
		}
	}
	//增加权限
	if len(matchIds) > 0 {
		if data.ProductId > 0 {
			_, err = db.DB.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": matchIds, "AclProducts": nil}).Update(map[string]interface{}{"AclProducts": `{}`})
			if err != nil {
				return
			}
			_, err = db.DB.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": matchIds}).Update(map[string]interface{}{"AclProducts": []string{"exp", "json_set(AclProducts, '$." + strconv.Itoa(int(data.ProductId)) + "','true' )"}})
			if err != nil {
				return
			}
		}
		if data.ProjectId > 0 {
			_, err = db.DB.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": matchIds, "AclProjects": nil}).Update(map[string]interface{}{"AclProjects": `{}`})
			if err != nil {
				return
			}
			_, err = db.DB.Table(db.TABLE_USER).Where(map[string]interface{}{"Id": matchIds}).Update(map[string]interface{}{"AclProjects": []string{"exp", "json_set(AclProjects, '$." + strconv.Itoa(int(data.ProjectId)) + "','true' )"}})
			if err != nil {
				return
			}
		}
	}
}
