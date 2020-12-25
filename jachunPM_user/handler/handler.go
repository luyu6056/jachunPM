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
		list, err := user_getPairs(data.Params, data.UsersToAppended)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_USER_getPairs_result()
		out.List = list
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_USER_updateUserView:
		updateUserView(data, in)
	default:
		libraries.ReleaseLog("未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
	}
}
