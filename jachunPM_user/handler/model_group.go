package handler

import (
	"jachunPM_user/db"
	"protocol"
	"strconv"
)

func gruop_getPairs(data *protocol.MSG_USER_Group_getPairs, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_Group_getPairs_result()
	var groups []*db.Group
	if err := in.DB.Table(db.TABLE_GROUP).Limit(0).Select(&groups); err != nil {
		in.WriteErr(err)
		return
	}
	for _, g := range groups {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(g.Id)), g.Name})
	}
	in.SendResult(out)
	out.Put()
}
func group_update(data *protocol.MSG_USER_group_update, in *protocol.Msg) {
	var oldGroup *protocol.MSG_USER_Group_cache
	err := in.DB.Table(db.TABLE_GROUP).Where("Id=?", data.Update.Id).Find(&oldGroup)
	if err != nil {
		in.WriteErr(err)
		return
	}
	if change, err := protocol.GetDiffChange(oldGroup, data.Update); err != nil && len(change) == 0 {
		in.WriteErr(err)
		return
	}
	if err = in.DB.Table(db.TABLE_GROUP).Replace(data.Update); err != nil {
		in.WriteErr(err)
		return
	}
	updateUserView(nil, []int32{data.Update.Id}, nil, nil, in)
}
