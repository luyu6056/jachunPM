package handler

import (
	"libraries"
	"protocol"
	"reflect"
)

var HostConn *protocol.RpcClient

func Handler(in *protocol.Msg) {

	switch data := in.Data.(type) {
	case *protocol.MSG_USER_GET_LoginSalt:
		getLoginSalt(data, in)
	case *protocol.MSG_USER_CheckPasswd:
		checkPasswd(data, in)
	case *protocol.MSG_USER_Dept_getDataStructure:
		res, err := dept_getDataStructure(data.RootDeptID)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_USER_Dept_getDataStructure_result()
		out.List = res
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_USER_Dept_update:
		in.WriteErr(dept_updateFromCache(data.List))
	case *protocol.MSG_USER_getDeptUserPairs:
		list, err := dept_getDeptUserPairs(data.DeptId)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_USER_getDeptUserPairs_result()
		out.List = list
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_USER_Dept_delete:
		dept_delete(data, in)
	case *protocol.MSG_USER_getCompanyUsers:
		user_getCompanyUsers(data, in)
	case *protocol.MSG_USER_INFO_updateByID:
		update := make(map[string]interface{}, len(data.Update))
		for k, v := range data.Update {
			update[k] = v
		}
		var err error
		if data.UserID != 0 {
			err = user_updateMap(map[string]interface{}{"Id": data.UserID}, update)
		} else {
			err = user_insertMap(update)
		}

		in.WriteErr(err)
	case *protocol.MSG_USER_CheckAccount:
		users, err := user_getUserInfo(map[string]interface{}{"account": data.Account})
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_USER_CheckAccount_result()
		if len(users) == 0 {
			out.Result = protocol.Success
		} else {
			out.Result = protocol.Err_UserAccountIsexist
		}
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_USER_getPairs:
		list, err := user_getPairs(data.Params, data.UsersToAppended, in)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_USER_getPairs_result()
		out.List = list
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_USER_updateUserView:
		updateUserView(data.UserIds,data.GroupIds,data.ProductIds,data.ProjectIds, in)
	case *protocol.MSG_USER_getContactLists:
		user_getContactLists(data, in)
	case *protocol.MSG_USER_getContactListByUid:
		user_getContactListByUid(data, in)
	case *protocol.MSG_USER_getContactListById:
		user_getContactListById(data, in)
	case *protocol.MSG_USER_insertUpdateContactList:
		user_insertUpdateContactList(data, in)
	case *protocol.MSG_USER_getGlobalContacts:
		list, err := user_getGlobalContacts()
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_USER_getGlobalContacts_result()
		for _, v := range list {
			tmp := protocol.GET_MSG_USER_ContactList()
			tmp.Id = v.Id
			tmp.ListName = v.ListName
			tmp.Uid = v.Uid
			tmp.UserList = v.UserList
			tmp.Share = v.Share
			out.Result = append(out.Result, tmp)
		}
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_USER_team_getByTypeRoot:
		team_getByTypeRoot(data, in)
	case *protocol.MSG_USER_team_addByList:
		team_addByList(data, in)
	case *protocol.MSG_USER_Group_getPairs:
		gruop_getPairs(data, in)
	case *protocol.MSG_USER_team_getByTypeUid:
		team_getByTypeUid(data, in)
	case *protocol.MSG_USER_user_getUserqueryByWhere:
		user_getUserqueryByWhere(data, in)
	case *protocol.MSG_USER_team_getByIds:
		team_getByIds(data, in)
	case *protocol.MSG_USER_team_updateByWhere:
		team_updateByWhere(data, in)
	case *protocol.MSG_USER_config_save:
		config_save(data,in)
	case *protocol.MSG_USER_team_delete:
		team_delete(data,in)
	case *protocol.MSG_USER_group_update:
		group_update(data,in)
	default:
		libraries.ReleaseLog("未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
	}
}
