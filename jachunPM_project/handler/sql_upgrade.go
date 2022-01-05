package handler

import (
	"jachunPM_project/db"
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
	type Task_old struct {
		Id             int32
		Ancestor       int32
		Parent         int32
		Project        int32
		Module         int32
		Story          int32
		StoryVersion   int16
		FromBug        int32
		Name           string
		Type           string
		Pri            int8
		Estimate       float64
		Consumed       float64
		Left           float64
		Deadline       time.Time
		Status         string
		Color          string
		Mailto         string
		Desc           string
		OpenedBy       string
		OpenedDate     time.Time
		AssignedTo     string
		AssignedDate   time.Time
		EstStarted     time.Time
		RealStarted    time.Time
		FinishedBy     string
		FinishedDate   time.Time
		CanceledBy     string
		CanceledDate   time.Time
		ClosedBy       string
		ClosedDate     time.Time
		ClosedReason   string
		LastEditedBy   string
		LastEditedDate time.Time
		Examine        bool
		ExamineDate    time.Time
		ExamineBy      string
		Deleted        bool
		Finalfile      bool
		Proofreading   bool
		PlaceOrder     bool
	}
	HostConn.DB.Regsiter(&Task_old{})
	var rows []*Task_old
	err = HostConn.DB.Table("zt_task").Field("`Id`,`Ancestor`,`Parent`,`Project`,`Module`,`Story`,`StoryVersion`,`FromBug`,`Name`,`Type`,`Pri`,`Estimate`,`Consumed`,`Left`,`Deadline`,`Status`,`Color`,`Mailto`,`Desc`,`OpenedBy`,`OpenedDate`,`AssignedTo`,`AssignedDate`,`EstStarted`,`RealStarted`,`FinishedBy`,`FinishedDate`,`CanceledBy`,`CanceledDate`,`ClosedBy`,`ClosedDate`,`ClosedReason`,`LastEditedBy`,`LastEditedDate`,`Examine`,`ExamineDate`,`ExamineBy`,`Deleted`,`Finalfile`,`Proofreading`,`PlaceOrder`").Limit(0).Select(&rows)
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取task表 err:%v", err)
	}

	tasks_insert := make([]*db.Task, len(rows))
	for k, v := range rows {
		task := &db.Task{}
		task.Id = v.Id
		task.Ancestor = v.Ancestor
		task.Parent = v.Parent
		task.Project = v.Project
		task.Module = v.Module
		task.Story = v.Story
		task.StoryVersion = v.StoryVersion
		task.FromBug = v.FromBug
		task.Name = v.Name
		task.Type = v.Type
		task.Pri = v.Pri
		task.Estimate = v.Estimate
		task.Consumed = v.Consumed
		task.Left = v.Left
		task.Deadline = v.Deadline
		task.Status = v.Status
		task.Color = v.Color
		for _, account := range strings.Split(v.Mailto, ",") {
			for _, kv := range result.List {
				if kv.Value == account {
					if id, _ := strconv.Atoi(kv.Key); id > 0 {
						task.Mailto = append(task.Mailto, int32(id))
					}
				}
			}
		}
		v.Desc, _ = libraries.Preg_replace(`<img src="{(\d+).png}" alt="" />`, `[img_0]/file/read?fileID=$1[/img]`, v.Desc)
		task.Desc = libraries.Html2bbcode(v.Desc)
		for _, kv := range result.List {
			if kv.Value == v.OpenedBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					task.OpenedBy = int32(id)
				}
			}
			if kv.Value == v.FinishedBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					task.FinishedBy = int32(id)
				}
			}
			if kv.Value == v.CanceledBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					task.CanceledBy = int32(id)
				}
			}
			if kv.Value == v.ClosedBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					task.ClosedBy = int32(id)
				}
			}
			if kv.Value == v.LastEditedBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					task.LastEditedBy = int32(id)
				}
			}
			if kv.Value == v.ExamineBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					task.ExamineBy = int32(id)
				}
			}
			if kv.Value == v.AssignedTo {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					task.AssignedTo = int32(id)
				}
			}

		}
		task.OpenedDate = v.OpenedDate

		task.AssignedDate = v.AssignedDate
		task.EstStarted = v.EstStarted
		task.RealStarted = v.RealStarted
		task.FinishedDate = v.FinishedDate
		task.CanceledDate = v.CanceledDate
		task.ClosedDate = v.ClosedDate
		task.ClosedReason = v.ClosedReason
		task.LastEditedDate = v.LastEditedDate
		task.Examine = v.Examine
		task.ExamineDate = v.ExamineDate
		task.Deleted = v.Deleted
		task.Finalfile = v.Finalfile
		task.Proofreading = v.Proofreading
		task.PlaceOrder = v.PlaceOrder
		tasks_insert[k] = task
	}

	_, err = HostConn.DB.Table(db.TABLE_TASK).InsertAll(tasks_insert)
	libraries.DebugLog("插入task %d 条，错误 %v", len(tasks_insert), err)
}
func init() {
	//go time.AfterFunc(time.Second*5, mysqlUpgrade)
}