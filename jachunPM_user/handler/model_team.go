package handler

import (
	"jachunPM_user/db"
	"protocol"
)

func team_getByTypeRoot(data *protocol.MSG_USER_team_getByTypeRoot, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_team_getByTypeRoot_result()
	defer out.Put()
	if err := in.DB.Table(db.TABLE_TEAM).Where(map[string]interface{}{
		"Root": data.Root,
		"Type": data.Type,
	}).Limit(0).Select(&out.List); err != nil {
		in.WriteErr(err)
		return
	}


	in.SendResult(out)
}
func team_addByList(data *protocol.MSG_USER_team_addByList, in *protocol.Msg) {
	_, err := in.DB.Table(db.TABLE_TEAM).InsertAll(data.List)
	in.WriteErr(err)
}
func team_getByTypeUid(data *protocol.MSG_USER_team_getByTypeUid, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_team_getByTypeUid_result()
	defer out.Put()
	if err := in.DB.Table(db.TABLE_TEAM).Prepare().Where("`Uid`=? and `Type`=?", data.Uid, data.Type).Limit(0).Select(&out.List); err != nil {
		in.WriteErr(err)
		return
	}

	in.SendResult(out)
}
func team_getByIds(data *protocol.MSG_USER_team_getByIds, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_team_getByIds_result()
	defer out.Put()
	if err := in.DB.Table(db.TABLE_TEAM).Where(map[string]interface{}{"Id": data.Ids}).Limit(0).Select(&out.List); err != nil {
		in.WriteErr(err)
		return
	}

	in.SendResult(out)

}
func team_updateByWhere(data *protocol.MSG_USER_team_updateByWhere, in *protocol.Msg) {
	_, err := in.DB.Table(db.TABLE_TEAM).Where(data.Where).Update(data.Update)
	in.WriteErr(err)
}
func team_delete(data *protocol.MSG_USER_team_delete,in *protocol.Msg){
	var teams []*db.Team
	err:=in.DB.Table(db.TABLE_TEAM).Where(data.Where).Select(&teams)
	if err!=nil || len(teams)==0{
		in.WriteErr(err)
		return
	}
	_,err=in.DB.Table(db.TABLE_TEAM).Where(data.Where).Delete()
	in.WriteErr(err)
	for _,team:=range teams{
		in.ActionCreate("team",team.Id,"deleted","","",nil,nil)
	}
}