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
