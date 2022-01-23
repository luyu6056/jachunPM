package handler

import (
	"jachunPM_user/db"
	"protocol"
)

func block_insertUpdate(data *protocol.MSG_USER_block_insertUpdate, in *protocol.Msg) {
	if data.Insert {
		_, err := in.DB.Table(db.TABLE_Block).InsertAll(data.List)
		in.WriteErr(err)
	} else {
		_, err := in.DB.Table(db.TABLE_Block).ReplaceAll(data.List)
		in.WriteErr(err)
	}
}

func block_getList(data *protocol.MSG_USER_block_getList, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_block_getList_result()
	if err := in.DB.Table(db.TABLE_Block).Where("Uid=? and Module=? and Hidden=0", data.Uid, data.Module).Order("order").Select(&out.List); err != nil {
		in.WriteErr(err)
	} else {
		in.SendResult(out)
	}
	out.Put()
}
func block_delectByWhere(data *protocol.MSG_USER_block_delectByWhere, in *protocol.Msg) {
	_, err := in.DB.Table(db.TABLE_Block).Where(data.Where).Delete()
	in.WriteErr(err)
}
