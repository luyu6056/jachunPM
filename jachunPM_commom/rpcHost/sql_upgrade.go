package rpcHost

import (
	"jachunPM_commom/db"
	"libraries"
	"protocol"
	"strconv"
	"time"
)

//老格式的表升级为新格式的表
func mysqlUpgrade() {

	out := protocol.GET_MSG_USER_getPairs()
	out.Params = "noletter,account"
	msg := GetMsg()

	var result *protocol.MSG_USER_getPairs_result
	if err := msg.SendMsgWaitResult(0, out, &result); err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取用户信息 err:%v", err)
		return
	}
	msg.DB.Table(db.TABLE_FILE).Delete()
	out.Put()
	type FILE_old struct {
		Id         int64  `db:"auto_increment;pk"`
		Pathname   string `db:"type:varchar(255)"`
		Title      string `db:"type:varchar(255)"`
		Extension  string `db:"type:varchar(30)"`
		Size       int64  `db:"default(0)"`
		ObjectType string `db:"type:varchar(30);index"`
		ObjectID   int32  `db:"not null;index"`
		AddedBy    string
		AddedDate  time.Time `db:"not null"`
		Deleted    bool      `db:"default(0)"` // 0=0,1=1,
		Type       string    `db:"type:varchar(50)"`
	}
	db.DB.Regsiter(&FILE_old{})

	var rows []*FILE_old
	err := db.DB.Table("zt_file").Field("`Id`,`Pathname`,`Title`,`Extension`,`Size`,`ObjectType`,`ObjectID`,`AddedBy`,`AddedDate`,`Deleted`,`Type`").Limit(0).Select(&rows)
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取file表 err:%v", err)
	}

	file_insert := make([]*db.File, len(rows))
	for k, v := range rows {
		file := &db.File{
			Id:         v.Id,
			Pathname:   v.Pathname,
			Title:      v.Title,
			Extension:  v.Extension,
			Size:       v.Size,
			ObjectType: v.ObjectType,
			ObjectID:   v.ObjectID,
			AddedDate:  v.AddedDate,
			Deleted:    v.Deleted,
			Type:       v.Type,
		}
		for _, kv := range result.List {
			if kv.Value == v.AddedBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					file.AddedBy = int32(id)
				}
			}
		}
		file_insert[k] = file
	}
	for i := 0; i < len(file_insert); i += 100 {
		en := i + 100
		if en > len(file_insert) {
			en = len(file_insert)
		}

		_, err = db.DB.Table(db.TABLE_FILE).InsertAll(file_insert[i:en])
		libraries.DebugLog("插入file %d-%d 条，错误 %v", i, en, err)
	}

}
func init() {
	//go time.AfterFunc(time.Second*5, mysqlUpgrade)
}
