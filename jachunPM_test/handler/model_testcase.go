package handler

import (
	"jachunPM_test/db"
	"protocol"
	"strconv"
)

func case_getTaskCasePairs(data *protocol.MSG_TEST_CASE_getTaskCasePairs, in *protocol.Msg) {
	var caseList []*db.Case
	err := in.DB.Table(db.TABLE_CASE).Prepare().Where("story=? and storyVersion=? and Deleted = 0", data.Story, data.StoryVersion).Select(&caseList)
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_TEST_CASE_getTaskCasePairs_result()
	for _, Case := range caseList {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(Case.Id)), Case.Title})
	}
	in.SendResult(out)
	out.Put()
}
