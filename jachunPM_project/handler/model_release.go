package handler

import (
	"jachunPM_project/db"
	"protocol"
)

func release_getById(data *protocol.MSG_PROJECT_release_getById, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_release_getById_result()
	if err := in.DB.Table(db.TABLE_RELEASE).Prepare().Where("Id=?", data.Id).Find(&out.Info); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
