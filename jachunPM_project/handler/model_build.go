package handler

import (
	"jachunPM_project/db"
	"protocol"
)

func build_getById(data *protocol.MSG_PROJECT_build_getById, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_build_getById_result()
	if err := in.DB.Table(db.TABLE_BUILD).Prepare().Where("Id=?", data.Id).Find(&out.Info); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
