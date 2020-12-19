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
	fields := "account, realname, deleted"
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
	var conditions map[string]interface{}
	if strings.Index(params, "nodeleted") > -1 || !userconfig["showDeleted"].(bool) {
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
		if user.Deleted && strings.Index(params, "realname") == -1 {
			res = append(res, protocol.HtmlKeyValueStr{user.Account, user.Account})
		} else {
			if user.Realname == "" {
				res = append(res, protocol.HtmlKeyValueStr{user.Account, user.Account})
			} else {
				res = append(res, protocol.HtmlKeyValueStr{user.Account, user.Realname})
			}
		}

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
	err := db.DB.Table(db.TABLE_USER).Where(where).Order("deleted asc,"+data.Sort).Limit(data.PerPage*(data.Page-1), data.Page*data.PerPage).Select(&out.List)
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
