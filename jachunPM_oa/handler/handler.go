package handler

import (
	"protocol"
)

var HostConn *protocol.RpcClient

func Handler(in *protocol.Msg) bool {
	switch data := in.Data.(type) {
	case *protocol.MSG_OA_attend_getByAccount:
		attend_getByAccount(data, in)
	case *protocol.MSG_OA_attend_getAllMonth:
		attend_getAllMonth(data, in)
	case *protocol.MSG_OA_attend_computeStat:
		attend_computeStat(data, in)
	case *protocol.MSG_OA_attend_detail:
		attend_detail(data, in)
	case *protocol.MSG_OA_attend_getWaitAttends:
		attend_getWaitAttends(data, in)
	case *protocol.MSG_OA_attend_getByDate:
		config, err := attend_LoadConfig(in)
		if err != nil {
			in.WriteErr(err)
		} else {
			if attend, err := attend_getByDate(config, data.Date, data.Uid); err != nil {
				in.WriteErr(err)
			} else {
				in.SendResult(attend)
				attend.Put()
			}
		}
	case *protocol.MSG_OA_attend_update:
		attend_update(data, in)
	case *protocol.MSG_OA_attend_getById:
		attend_getById(data, in)
	default:
		return false
	}
	return true
}
