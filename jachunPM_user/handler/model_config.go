package handler

import (
	"jachunPM_user/db"
	"protocol"
)

func config_save(data *protocol.MSG_USER_config_save, in *protocol.Msg) {
	var custom db.Config
	err := in.DB.Table(db.TABLE_Config).Where("Uid=? and Module=? and `Key`=?",data.Uid,data.Module,data.Key).Find(&custom)
	if err != nil {
		in.WriteErr(err)
		return
	}
	if data.Type == "add" {
		custom.Uid = data.Uid
		custom.Module = data.Module
		custom.Key = data.Key
		custom.Value = data.Value
		custom.Section = data.Section
		err = in.DB.Table(db.TABLE_Config).Replace(custom)
	} else {
		_, err = in.DB.Table(db.TABLE_Config).Where("Uid=? and Module=? and `Key`=?", data.Uid, data.Module, data.Key).Delete()
	}
	if err == nil {
		user_updateCacheById(data.Uid)
	}
	in.WriteErr(err)
}
