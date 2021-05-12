package handler

import (
	"jachunPM_test/db"
	"protocol"
)

func testsuite_getById(data *protocol.MSG_TEST_testsuite_getById, in *protocol.Msg) {
	out := protocol.GET_MSG_TEST_testsuite_getById_result()
	if err := in.DB.Table(db.TABLE_TESTSUITE).Prepare().Where("Id=?", data.Id).Find(&out.Info); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
