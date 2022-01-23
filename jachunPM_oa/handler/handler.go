package handler

import (
	"libraries"
	"protocol"
)

var HostConn *protocol.RpcClient

func Handler(in *protocol.Msg) {
	switch data := in.Data.(type) {
	case *protocol.MSG_OA_attend_getByAccount:
		attend_getByAccount(data, in)
	case *protocol.MSG_OA_attend_getAllMonth:
		attend_getAllMonth(data, in)
	case *protocol.MSG_OA_attend_computeStat:
		attend_computeStat(data, in)
	default:
		if v, ok := protocol.CmdToName[in.Cmd]; ok {
			libraries.ReleaseLog("未设置消息CMD%s处理", v)
		} else {
			libraries.ReleaseLog("未设置消息CMD%d处理", in.Cmd)
		}

	}
}
