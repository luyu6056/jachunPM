package handler

import (
	"libraries"
	"protocol"
)

var HostConn *protocol.RpcClient

func Handler(in *protocol.Msg) {
	switch data := in.Data.(type) {
	case *protocol.MSG_LOG_Action_Create:
		action_crate(data, in)
	case *protocol.MSG_LOG_Action_GetByWhereMap:
		action_GetByWhereMap(data, in)
	case *protocol.MSG_LOG_Action_transformActions:
		action_transformActions(data, in)
	case *protocol.MSG_LOG_Action_AddHistory:
		action_AddHistory(data, in)
	case *protocol.MSG_LOG_Action_set_read:
		action_read(data, in)
	default:
		if v, ok := protocol.CmdToName[in.Cmd]; ok {
			libraries.ReleaseLog("未设置消息CMD%s处理", v)
		} else {
			libraries.ReleaseLog("未设置消息CMD%d处理", in.Cmd)
		}
	}
}
