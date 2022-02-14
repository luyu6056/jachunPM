package handler

import (
	"protocol"
)

var HostConn *protocol.RpcClient

func Handler(in *protocol.Msg) bool {
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
		return false
	}
	return true
}
