package handler

import (
	"jachunPM_project/db"
	"protocol"
	"strconv"
)

func task_getPairs(data *protocol.MSG_PROJECT_task_getPairs, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_task_getPairs_result()
	var list []*db.Task
	if err := in.DB.Table(db.TABLE_BUILD).Field("Id,Name").Where(data.Where).Limit(0).Select(&list); err != nil {
		in.WriteErr(err)
		return
	}
	for _, v := range list {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Name})
	}
	in.SendResult(out)
	out.Put()
}
func task_getListByWhereMap(data *protocol.MSG_PROJECT_task_getListByWhereMap, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_task_getListByWhereMap_result()
	sql := in.DB.Table(db.TABLE_TASK).Where(data.Where).Order(data.Order)
	if data.PerPage > 0 {
		sql.Limit((data.Page-1)*data.PerPage, data.PerPage)
	}
	err := sql.Select(&out.List)
	if err != nil {
		in.WriteErr(err)
		return
	}
	out.Total = data.Total
	if data.Total == 0 {
		if out.Total, err = HostConn.DB.Table(db.TABLE_TASK).Where(data.Where).Count(); err != nil {
			in.WriteErr(err)
			return
		}
	}
	in.SendResult(out)
	out.Put()
}
