package handler

import (
	"jachunPM_test/db"
	"protocol"
)

func testtask_getById(data *protocol.MSG_TEST_testtask_getById, in *protocol.Msg) {
	out := protocol.GET_MSG_TEST_testtask_getById_result()
	if err := in.DB.Table(db.TABLE_TESTTASK).Prepare().Where("Id=?", data.Id).Find(&out.Info); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
