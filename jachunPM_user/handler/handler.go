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
	case *protocol.MSG_USER_Dept_delete:
		dept_delete(data, in)
	case *protocol.MSG_USER_getCompanyUsers:
		user_getCompanyUsers(data, in)
	default:
		libraries.ReleaseLog("未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
	}
}
