package handler

import (
	"jachunPM_user/db"
	"protocol"
)

func team_getByTypeRoot(data *protocol.MSG_USER_team_getByTypeRoot, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_team_getByTypeRoot_result()
	if err := in.DB.Table(db.TABLE_TEAM).Prepare().Where("`Root`=? and `Type`=?", data.Root, data.Type).Limit(0).Select(&out.List); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
func team_addByList(data *protocol.MSG_USER_team_addByList, in *protocol.Msg) {
	_, err := in.DB.Table(db.TABLE_TEAM).InsertAll(data.List)
	in.WriteErr(err)
}
func team_getByTypeUid(data *protocol.MSG_USER_team_getByTypeUid, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_team_getByTypeUid_result()
	if err := in.DB.Table(db.TABLE_TEAM).Prepare().Where("`Uid`=? and `Type`=?", data.Uid, data.Type).Limit(0).Select(&out.List); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
