package handler

import (
	"libraries"
	"protocol"
	"reflect"
)

var HostConn *protocol.RpcClient

func Handler(in *protocol.Msg) {
	switch data := in.Data.(type) {
	case *protocol.MSG_LOG_Action_Create:
		action_crate(data, in)
	case *protocol.MSG_LOG_Action_GetByWhereMap:
		action_GetByWhereMap(data, in)
	case *protocol.MSG_LOG_Action_transformActions:
		list, err := action_transformActions(data.Where, data.Order, in)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_LOG_Action_transformActions_result()
		out.List = list
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_LOG_Action_AddHistory:
		action_AddHistory(data, in)
	case *protocol.MSG_LOG_Action_set_read:
		action_read(data, in)
	default:
		libraries.ReleaseLog("未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
	}
}
