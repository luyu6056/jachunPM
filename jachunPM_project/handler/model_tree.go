package handler

import (
	"jachunPM_project/db"
	"protocol"
	"strconv"
)

func tree_getLinePairs(data *protocol.MSG_PROJECT_tree_getLinePairs, in *protocol.Msg) {
	var list []*db.Module
	err := db.DB.Table(db.TABLE_MODULE).Field("Id,Name").Where("type = 'line' and deleted = 0").Select(&list)
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_tree_getLinePairs_result()
	for _, v := range list {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Name})
	}
	in.SendResult(out)
	out.Put()
}
