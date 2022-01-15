package handler

import (
	"jachunPM_user/db"
	"protocol"
	"strconv"
)

func team_getByTypeRoot(data *protocol.MSG_USER_team_getByTypeRoot, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_team_getByTypeRoot_result()
	if err := in.DB.Table(db.TABLE_TEAM).Where(map[string]interface{}{
		"Root": data.Root,
		"Type": data.Type,
	}).Limit(0).Select(&out.List); err != nil {
		in.WriteErr(err)
		return
	}

	in.SendResult(out)
	out.Put()
}
func team_addByList(data *protocol.MSG_USER_team_addByList, in *protocol.Msg) {
	success, err := in.DB.Table(db.TABLE_TEAM).InsertAll(data.List)
	in.WriteErr(err)
	if err==nil && success{
		updateTeamUserByteans(data.List,in)
		for _,team:= range data.List{
			in.ActionCreate("team", team.Id, "create", "", "", nil, 0)
		}
	}
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
	if err==nil {
		var teams []*protocol.MSG_USER_team_info
		in.DB.Table(db.TABLE_TEAM).Where(data.Where).Select(&teams)
		updateTeamUserByteans(teams,in)
	}
}
func team_delete(data *protocol.MSG_USER_team_delete, in *protocol.Msg) {
	var teams []*protocol.MSG_USER_team_info
	err := in.DB.Table(db.TABLE_TEAM).Where(data.Where).Select(&teams)
	if err != nil || len(teams) == 0 {
		in.WriteErr(err)
		return
	}
	_, err = in.DB.Table(db.TABLE_TEAM).Where(data.Where).Delete()
	in.WriteErr(err)
	for _, team := range teams {
		in.ActionCreate("team", team.Id, "delete", "", "", nil, 0)
	}
	updateTeamUserByteans(teams,in)
}
func team_getTeams2Import(data *protocol.MSG_USER_team_getTeams2Import, in *protocol.Msg) {
	var teams []*db.Team
	err := in.DB.Table(db.TABLE_TEAM).Field("`Id`,`Root`").Where("Type='project' and Root != ?", data.ProjectId).Order("root desc").Group("root").Select(&teams)
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_USER_team_getTeams2Import_result()
	for _, t := range teams {
		project := HostConn.GetProjectById(t.Root)
		if project != nil && !project.Deleted {
			out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(project.Id)), project.Name})
		}
	}
	in.SendResult(out)
	out.Put()
}
func team_projectManageMembers(data *protocol.MSG_USER_team_projectManageMembers, in *protocol.Msg) {
	var teams []*db.Team
	err := in.DB.Table(db.TABLE_TEAM).Where("Root=? and Type='project'", data.ProjectId).Select(&teams)
	if err != nil {
		in.WriteErr(err)
		return
	}
	var uids []int32
	for i := len(data.Update) - 1; i >= 0; i-- {
		update := data.Update[i]
		uids = append(uids, update.Uid)
		for _, team := range teams {
			if team.Uid == update.Uid {
				team.Limited = update.Limited
				team.Hours = update.Hours
				team.Role = update.Role
				team.Days = update.Days
				data.Update = append(data.Update[:i], data.Update[i+1:]...)
				break
			}
		}
	}
	session, err := in.BeginTransaction()
	defer func() {
		in.WriteErr(err)
		if err != nil {
			session.Rollback()
		} else {
			session.Commit()
		}

	}()
	if err != nil {
		return
	}
	if _, err = in.DB.Table(db.TABLE_TEAM).ReplaceAll(teams); err != nil {
		return
	}
	if _, err = in.DB.Table(db.TABLE_TEAM).InsertAll(data.Update); err != nil {
		return
	}
	updateUserView(uids, nil, nil, []int32{data.ProjectId}, in)
	for _,team:= range teams{
		in.ActionCreate("team", team.Id, "update", "", "", nil, 0)
	}
	for _,team:= range data.Update{
		in.ActionCreate("team", team.Id, "create", "", "", nil, 0)
	}
}
func updateTeamUserByteans(teams []*protocol.MSG_USER_team_info,in *protocol.Msg){
	var uids []int32
	var projects []int32
	for _,team:= range teams{
		findUid:=false
		for _,uid:= range uids{
			if team.Uid==uid{
				findUid=true
				break
			}
		}
		if !findUid{
			uids=append(uids,team.Uid)
		}
		if team.Type=="project"{
			find:= false
			for _,id:= range projects{
				if team.Root==id{
					find=true
					break
				}
			}
			if !find{
				projects=append(projects,team.Root)
			}
		}
	}
	updateUserView(uids,nil,nil,projects,in)

}