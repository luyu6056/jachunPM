package handler

import (
	"jachunPM_user/db"
	"libraries"
	"protocol"
)

func config_save(data *protocol.MSG_USER_config_save, in *protocol.Msg) {
	var custom db.Config
	err := in.DB.Table(db.TABLE_Config).Where("Uid=? and Module=? and `Key`=?", data.Uid, data.Module, data.Key).Find(&custom)
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

func config_get(data *protocol.MSG_USER_config_get,in *protocol.Msg){
	var configs []*db.Config
	if err:=in.DB.Table(db.TABLE_Config).Where("Uid=? and Module=?",data.Uid,data.Module).Select(&configs);err!=nil{
		in.WriteErr(err)
		return
	}

	out:=protocol.GET_MSG_USER_config_get_result()
	out.Config=make(map[string]map[string]string)
	configCache,err:=in.LoadConfig(data.Module)
	if err==nil{
		for key,value:= range configCache{
			if out.Config[key]==nil{
				out.Config[key]=make(map[string]string)
			}
			for k,v:= range value{
				if str,ok:=v.(string);ok{
					out.Config[key][k]=str
				}else{
					out.Config[key][k]=libraries.I2S(v)
				}
			}
		}
	}
	for _,config:= range configs{
		if out.Config[config.Section]==nil{
			out.Config[config.Section]=make(map[string]string)
		}
		out.Config[config.Section][config.Key]=config.Value
	}
	in.SendResult(out)
	out.Put()
}
func config_savelist(datas *protocol.MSG_USER_config_savelist, in *protocol.Msg) {
	session,err:=in.BeginTransaction()
	var ids []int32
	defer func(){
		if err!=nil{
			session.Rollback()
		}else{
			session.Commit()
			for _,id:= range ids{
				user_updateCacheById(id)
			}
		}
		in.WriteErr(err)
	}()
	for _,data:= range datas.List{
		var custom db.Config
		if err = session.Table(db.TABLE_Config).Where("Uid=? and Module=? and `Key`=?", data.Uid, data.Module, data.Key).Find(&custom);err != nil {
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
		if err != nil {
			return
		}
		find:=false
		for _,id:= range ids{
			if id==data.Uid{
				find=true
			}
		}
		if !find{
			ids=append(ids,data.Uid)
		}
	}


}