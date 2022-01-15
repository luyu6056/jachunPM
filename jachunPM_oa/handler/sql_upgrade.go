package handler

import (
	"fmt"
	"jachunPM_log/db"
	"libraries"
	"protocol"
	"strconv"
	"strings"
	"time"
)

//老格式的表升级为新格式的表
func mysqlUpgrade() {

	out := protocol.GET_MSG_USER_getPairs()
	out.Params = "noletter,account"
	msg, err := HostConn.GetMsg()
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法初始化msg err:%v", err)
		return
	}
	var result *protocol.MSG_USER_getPairs_result
	if err := msg.SendMsgWaitResult(0, out, &result); err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取用户信息 err:%v", err)
		return
	}
	out.Put()
	type Action_old struct {
		Id         int64  `db:"auto_increment;pk"`
		ObjectType string `db:"type:varchar(30)"`
		ObjectID   int32  `db:"default(0)"`
		Product    string
		Project    int32
		Actor      string    `db:"type:varchar(30)"`
		Action     string    `db:"type:varchar(30)"`
		Date       time.Time `db:"not null"`
		Comment    string    `db:"type:text"`
		Extra      string    `db:"type:text"`
		Read       bool
	}
	HostConn.DB.Regsiter(&Action_old{})

	HostConn.DB.Regsiter(&db.History{})
	setp := 5000
	for i := 0; true; i += setp {
		var rows []*Action_old
		err = HostConn.DB.Table("zt_action").Field("Id,ObjectType,ObjectID,Product,Project,Actor,Action,Date,Comment,`Extra`,`Read`").Limit(i, setp).Select(&rows)

		if err != nil {
			libraries.ReleaseLog("mysqlUpgrade无法获取zt_action表 err:%v", err)
		}
		if len(rows) == 0 {
			break
		}
		action_insert := make([]*db.Action, len(rows))
		for k, v := range rows {
			action := &db.Action{}
			action.Id = v.Id
			action.Id = v.Id
			action.ObjectType = v.ObjectType
			action.ObjectID = v.ObjectID
			action.Project = v.Project
			action.Actor = v.Actor
			action.Action = v.Action
			action.Date = v.Date
			action.Comment = v.Comment
			action.Extra = v.Extra
			action.Read = v.Read

			err = HostConn.DB.Table("zt_history").Field("Field,Old,New,Diff").Where("action=?", v.Id).Select(&action.Historys)
			if err != nil {
				libraries.ReleaseLog("mysqlUpgrade无法获取zt_history表 err:%v", err)
			}
			for _, history := range action.Historys {
				if history.Field == "mailto" {
					var ids []int
					for _, account := range strings.Split(history.Old, ",") {
						for _, kv := range result.List {
							if kv.Value == account {
								if id, _ := strconv.Atoi(kv.Key); id > 0 {
									ids = append(ids, id)
								}
							}

						}
					}
					history.Old = fmt.Sprint(ids)
					ids = nil
					for _, account := range strings.Split(history.New, ",") {
						for _, kv := range result.List {
							if kv.Value == account {
								if id, _ := strconv.Atoi(kv.Key); id > 0 {
									ids = append(ids, id)
								}
							}

						}
					}
					history.New = fmt.Sprint(ids)
				}
			}
			for _, s := range strings.Split(v.Product, ",") {
				if s != "" {
					if id, _ := strconv.Atoi(s); id > 0 {
						action.Products = append(action.Products, int32(id))
					}
				}
			}
			for _, kv := range result.List {
				if kv.Value == v.Actor {
					if id, _ := strconv.Atoi(kv.Key); id > 0 {
						action.ActorId = int32(id)
					}
				}

			}

			action_insert[k] = action
		}

		_, err = HostConn.DB.Table(db.TABLE_ACTION).InsertAll(action_insert)
		libraries.DebugLog("插入task %d 条，错误 %v", len(action_insert), err)
	}

}
func init() {
	//go time.AfterFunc(time.Second*5, mysqlUpgrade)
}
