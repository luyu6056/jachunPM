package handler

import (
	"jachunPM_test/db"
	"protocol"
	"strconv"
)

func bug_getPairs(data *protocol.MSG_TEST_bug_getPairs, in *protocol.Msg) {
	out := protocol.GET_MSG_TEST_bug_getPairs_result()
	var list []*db.Bug
	if err := in.DB.Table(db.TABLE_BUG).Field("Id,Title").Where(data.Where).Limit(0).Select(&list); err != nil {
		in.WriteErr(err)
		return
	}
	for _, v := range list {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Title})
	}
	in.SendResult(out)
	out.Put()
}
func bug_getCountByWhere(data *protocol.MSG_TEST_bug_getCountByWhere, in *protocol.Msg) {
	out := protocol.GET_MSG_TEST_bug_getCountByWhere_result()
	var err error
	if out.Count, err = in.DB.Table(db.TABLE_BUG).Where(data.Where).Count(); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
