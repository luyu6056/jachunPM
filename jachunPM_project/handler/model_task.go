package handler

import (
	"errors"
	"fmt"
	"jachunPM_project/db"
	"libraries"
	"math"
	"mysql"
	"protocol"
	"strconv"
	"time"
)

func task_getPairs(data *protocol.MSG_PROJECT_task_getPairs, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_task_getPairs_result()
	var list []*db.Task
	if err := in.DB.Table(db.TABLE_TASK).Field("Id,Name").Where(data.Where).Limit(0).Select(&list); err != nil {
		in.WriteErr(err)
		return
	}
	for _, v := range list {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Name})
	}
	in.SendResult(out)
	out.Put()
}
func task_getListByWhereMap(data *protocol.MSG_PROJECT_task_getListByWhereMap, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_task_getListByWhereMap_result()
	sql := in.DB.Table(db.TABLE_TASK).Where(data.Where).Order(data.Order)
	if data.PerPage > 0 {
		sql.Limit((data.Page-1)*data.PerPage, data.PerPage)
	}
	err := sql.Select(&out.List)
	if err != nil {
		in.WriteErr(err)
		return
	}
	out.Total = data.Total
	if data.Total == 0 {
		if out.Total, err = HostConn.DB.Table(db.TABLE_TASK).Where(data.Where).Count(); err != nil {
			in.WriteErr(err)
			return
		}
	}
	in.SendResult(out)
	out.Put()
}
func task_processTasks(tasks []*protocol.MSG_PROJECT_TASK) {
	for _, task := range tasks {
		task_processTasks(task.Children)
		task_processTasks(task.Grandchildren)
		task_processTask(task)
	}
}
func task_processTask(task *protocol.MSG_PROJECT_TASK) {
	/* Delayed or not?. */
	if task.Status != "done" && task.Status != "cancel" && task.Status != "closed" {
		if task.Deadline.After(protocol.NORMALTIME) {
			today, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
			delay := today.Sub(task.Deadline) / time.Second / 86400
			if delay > 0 {
				task.Delay = int32(delay)
			}
		}
	}

	/* Story changed or not. */
	task.NeedConfirm = false
	if task.StoryStatus == "active" && task.LatestStoryVersion > task.StoryVersion {
		task.NeedConfirm = true
	}

	/* Compute task progress. */
	if task.Consumed == 0 && task.Left == 0 {
		task.Progress = 0
	} else if task.Consumed != 0 && task.Left == 0 {
		task.Progress = 100
	} else {
		task.Progress = int(math.Round(task.Consumed / (task.Consumed + task.Left) * 100))
	}

}
func task_getById(data *protocol.MSG_PROJECT_task_getById, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_task_getById_result()
	err := in.DB.Table(db.TABLE_TASK).Alias("t1").LeftJoin(db.TABLE_STORY).Alias("t2").On("t1.story = t2.id").Field("t1.*, t2.Id AS StoryID, t2.Title AS StoryTitle, t2.Product as Product, t2.Branch as Branch, t2.version AS LatestStoryVersion, t2.Status AS StoryStatus").Prepare().Where("t1.Id=?", data.Id).Find(&out.Info)

	defer func() {
		if err != nil {
			in.WriteErr(err)
		} else {
			in.SendResult(out)
		}
		out.Put()
	}()
	if err != nil || out.Info == nil {
		return
	}
	var children []*protocol.MSG_PROJECT_TASK
	if err := in.DB.Table(db.TABLE_TASK).Alias("t1").Field("t1.*, t2.Id AS StoryID, t2.Title AS StoryTitle, t2.Product as Product, t2.Branch as Branch, t2.version AS LatestStoryVersion, t2.Status AS StoryStatus").LeftJoin(db.TABLE_STORY).Alias("t2").On("t1.Story = t2.Id").Prepare().Where("Parent=?", out.Info.Id).Limit(0).Select(&children); err != nil {
		in.WriteErr(err)
		return
	}
	for _, child := range children {
		out.Info.Children = append(out.Info.Children, child)
	}
	task_processTask(out.Info)
}
func task_create(data *protocol.MSG_PROJECT_task_create, in *protocol.Msg) {

	session, err := in.BeginTransaction()
	out := protocol.GET_MSG_PROJECT_task_create_result()
	defer func() {
		if err != nil {
			session.Rollback()
			in.WriteErr(err)
		} else {
			session.Commit()
			in.SendResult(out)
		}
		out.Put()
	}()
	if err != nil {
		return
	}
	if data.Task.Id == 0 { //create
		c, err := in.DB.Table(db.TABLE_TASK).Where("Name=? and Project=? and Deleted=0", data.Task.Name, data.Task.Project).Count()
		if err != nil {
			in.WriteErr(err)
			return
		}
		if c > 0 {
			in.WriteErr(protocol.Err_TaskIsexist.Err())
			return
		}
		var id int64
		if id, err = session.Table(db.TABLE_TASK).Insert(data.Task); err != nil || id == 0 {
			return
		}
		out.Id = int32(id)
	} else { //update
		out.Id = data.Task.Id
		var oldTask *protocol.MSG_PROJECT_TASK
		err = in.DB.Table(db.TABLE_TASK).Prepare().Where("Id=?", data.Task.Id).Find(&oldTask)
		if err != nil {
			return
		}
		if oldTask == nil {
			err = protocol.Err_TaskNotFound.Err()
			return
		}
		data.Task.LastEditedBy = in.GetUserID()
		data.Task.LastEditedDate = time.Now()
		if oldTask.Project != data.Task.Project {
			if _, err = session.Table(db.TABLE_TASK).Where("Parent=? || Ancestor=?", oldTask.Id, oldTask.Id).Update(map[string]interface{}{
				"Project": data.Task.Project,
				"Module":  data.Task.Module,
			}); err != nil {
				return
			}
		}
		var childrenIds []int32
		if data.Task.Parent != oldTask.Parent {
			var children []*db.Task
			//获得当前任务的子任务
			if err = session.Table(db.TABLE_TASK).Field("Id").Where("parent=? and Deleted=0", data.Task.Id).Select(&children); err != nil {
				return
			}

			//检查是否有孙任务
			if len(children) > 0 {
				for _, v := range children {
					childrenIds = append(childrenIds, v.Id)
					if data.Task.Parent==v.Id{
						err = protocol.Err_taskHasAncestors.Err()
						return
					}
				}
				var ancestors []*db.Task
				if err = session.Table(db.TABLE_TASK).Field("Id").Where(map[string]interface{}{"parent": mysql.WhereOperatorIN(childrenIds), "Deleted": false}).Select(&ancestors); err != nil {
					return
				}
				if len(ancestors) > 0 {
					err = protocol.Err_taskHasAncestors.Err()
					return
				}
			}
			data.Task.Ancestor = -1
		}
		if data.Task.Story > 0 {
			var story *protocol.MSG_PROJECT_story
			if story, err = story_getById(data.Task.StoryID, 0); err == nil {
				data.Task.StoryVersion = story.Version
			} else {
				return
			}

		}

		if err = session.Table(db.TABLE_TASK).Replace(data.Task); err != nil {
			return
		}
		if len(childrenIds) > 0 {
			if _, err = session.Table(db.TABLE_TASK).Where(map[string]interface{}{"Id": mysql.WhereOperatorIN(childrenIds)}).Update("Ancestor=?", data.Task.Parent); err != nil {
				return
			}
			if _, err = session.Table(db.TABLE_TASK).Where(map[string]interface{}{"Id": data.Task.Parent}).Update("Ancestor=-1"); err != nil {
				return
			}
		}
		if data.Task.Parent != oldTask.Parent {
			if err = task_updateParentAncestor(data.Task.Id, in); err != nil {
				return
			}
			if oldTask.Parent > 0 {
				if err = task_updateParentAncestor(oldTask.Parent, in); err != nil {
					return
				}
			}
			if err = task_updateParentAncestor(data.Task.Parent, in); err != nil {
				return
			}
		}
		if data.Task.Parent > 0 || oldTask.Parent > 0 {
			if data.Task.Parent > 0 {
				if _, err = session.Table(db.TABLE_TASK).Where("Id=? and Parent=0", data.Task.Parent).Update("Parent=-1"); err != nil {
					return
				}

			}
			if err = task_updateParentStatus(data.Task.Id, in); err != nil {
				return
			}
			if err = task_computeBeginAndEnd(data.Task.Parent, in); err != nil {
				return
			}

		}
		if out.Change, err = protocol.GetDiffChange(oldTask, data.Task); err != nil {
			return
		}
	}

	if err = story_setStage(data.Task.Story, in); err != nil {
		return
	}

}
func task_GetTaskEstimateByTaskId(data *protocol.MSG_PROJECT_task_GetTaskEstimateByTaskId, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_task_GetTaskEstimateByTaskId_result()
	if err := in.DB.Table(db.TABLE_TASKESTIMATE).Where("Task=?", data.TaskId).Select(&out.List); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}

func task_UpdateTaskEstimate(data *protocol.MSG_PROJECT_task_UpdateTaskEstimate, in *protocol.Msg) {
	res := protocol.GET_MSG_PROJECT_task_UpdateTaskEstimate_result()
	if len(data.List) == 0 {
		in.SendResult(res)
		return
	}
	var task *db.Task
	err := in.DB.Table(db.TABLE_TASK).Where("Id=?", data.TaskId).Find(&task)
	if err != nil {
		in.WriteErr(err)
		return
	}

	var earliestTime time.Time
	for _, record := range data.List {
		if earliestTime.IsZero() {
			earliestTime = record.Date
		} else if earliestTime.Unix() > record.Date.Unix() {
			earliestTime = record.Date
		}

	}

	consumed := float64(0)
	left := task.Left
	now := time.Now()
	m, err := in.DB.Table(db.TABLE_TASKESTIMATE).Where("task=?", task.Id).Order("date desc").FindMap()
	if err != nil {
		in.WriteErr(err)
		return
	}
	lastDate, _ := time.Parse("2006-01-02", m["Date"])
	project := HostConn.GetProjectById(task.Project)
	var actionID int64
	for _, estimate := range data.List {
		_, err := task_addTaskEstimate(estimate, in)
		if err != nil {
			in.WriteErr(err)
			return
		}

		consumed += estimate.Consumed
		actionID, err = in.ActionCreate("task", task.Id, "RecordEstimate", estimate.Work, fmt.Sprint(estimate.Consumed), project.Products, project.Id)
		if err != nil {
			in.WriteErr(err)
			return
		}
		if lastDate.IsZero() || lastDate.Unix() <= estimate.Date.Unix() {
			left = float64(estimate.Left)
			lastDate = estimate.Date
		}
	}
	newTask := &db.Task{
		Id:             task.Id,
		Ancestor:       task.Ancestor,
		Parent:         task.Parent,
		Project:        task.Project,
		Module:         task.Module,
		Story:          task.Story,
		StoryVersion:   task.StoryVersion,
		FromBug:        task.FromBug,
		Name:           task.Name,
		Type:           task.Type,
		Pri:            task.Pri,
		Estimate:       task.Estimate,
		Consumed:       task.Consumed + float64(consumed),
		Left:           left,
		Deadline:       task.Deadline,
		Status:         task.Status,
		Color:          task.Color,
		Mailto:         task.Mailto,
		Desc:           task.Desc,
		OpenedBy:       task.OpenedBy,
		OpenedDate:     task.OpenedDate,
		AssignedTo:     task.AssignedTo,
		AssignedDate:   task.AssignedDate,
		EstStarted:     task.EstStarted,
		RealStarted:    task.RealStarted,
		FinishedBy:     task.FinishedBy,
		FinishedDate:   task.FinishedDate,
		CanceledBy:     task.CanceledBy,
		CanceledDate:   task.CanceledDate,
		ClosedBy:       task.ClosedBy,
		ClosedDate:     task.ClosedDate,
		ClosedReason:   task.ClosedReason,
		LastEditedBy:   in.GetUserID(),
		LastEditedDate: now,
		Examine:        task.Examine,
		ExamineDate:    task.ExamineDate,
		ExamineBy:      task.ExamineBy,
		Deleted:        task.Deleted,
		Finalfile:      task.Finalfile,
		Proofreading:   task.Proofreading,
		Team:           task.Team,
		PlaceOrder:     task.PlaceOrder,
	}

	if left == 0 {
		newTask.Status = "done"
		//data.AssignedTo   = task.OpenedBy;
		newTask.AssignedDate = now
		newTask.FinishedBy = in.GetUserID()
		newTask.FinishedDate = now
	} else if task.Status == "wait" {
		newTask.Status = "doing"
		//data.AssignedTo   = this->app->user->account;
		newTask.AssignedDate = now
		newTask.RealStarted = earliestTime
	} else if task.Status == "pause" {
		newTask.Status = "doing"
		//data.AssignedTo   = this->app->user->account;
		newTask.AssignedDate = now
	}

	if len(task.Team) > 0 {
		updateTeam := protocol.GET_MSG_USER_team_updateByWhere()
		updateTeam.Where = map[string]interface{}{
			"Root": task.Id,
			"Type": "task",
			"Uid":  task.AssignedTo,
		}
		updateTeam.Update = map[string]interface{}{
			"Consumed": mysql.UpdateValueRaw("Consumed + " + strconv.FormatFloat(float64(consumed), 'g', -1, 32)),
			"Left":     left,
		}
		if err = task_computeHours4Multiple(task, newTask, in); err != nil {
			in.WriteErr(err)
			return
		}
	}
	if err = in.DB.Table(db.TABLE_TASK).Replace(newTask); err != nil {
		in.WriteErr(err)
		return
	}

	if task.Parent > 0 {
		if err = task_updateParentStatus(task.Id, in); err != nil {
			in.WriteErr(err)
			return
		}
	}
	if task.Story > 0 {
		if err = story_setStage(task.Story, in); err != nil {
			in.WriteErr(err)
			return
		}
	}
	if res.Changes, err = in.ActionLogHistory(actionID, task, newTask); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(res)
	//if(task.Status == "done" and !dao::isError()) this->loadModel('score')->create('task', 'finish', taskID);
}
func task_addTaskEstimate(estimate *protocol.MSG_PROJECT_TaskEstimate, in *protocol.Msg) (int64, error) {
	return in.DB.Table(db.TABLE_TASKESTIMATE).Insert(estimate)
}
func task_computeHours4Multiple(oldTask, newTask *db.Task, in *protocol.Msg) (err error) {
	if oldTask == nil || newTask == nil {
		return errors.New("task_computeHours4Multiple参数错误，不能为nil")
	}
	if len(oldTask.Team) == 0 {
		return
	}

	getTeam := protocol.GET_MSG_USER_team_getByIds()
	getTeam.Ids = oldTask.Team
	var team *protocol.MSG_USER_team_getByIds_result
	if err = in.SendMsgWaitResult(0, getTeam, &team); err != nil {
		return
	}
	now := time.Now()
	if newTask.Status == "" {
		newTask.Status = oldTask.Status
	}

	if oldTask.AssignedTo == 0 {

		newTask.AssignedTo = oldTask.Team[0]
		newTask.AssignedDate = now
	} else {
		var teamAssigendTo *protocol.MSG_USER_team_info
		var nextId int32
		for k, teamMember := range team.List {
			if teamMember.Uid == oldTask.AssignedTo {
				teamAssigendTo = teamMember
				if k != len(team.List)-1 {
					nextId = team.List[k+1].Uid
				}
				break
			}

		}

		if teamAssigendTo.Left == 0 && teamAssigendTo.Consumed != 0 {
			if oldTask.AssignedTo != team.List[len(team.List)-1].Uid {
				newTask.AssignedTo = nextId
			} else {
				newTask.AssignedTo = oldTask.OpenedBy
			}
			newTask.AssignedDate = now
		}
	}

	newTask.Estimate = 0
	newTask.Consumed = 0
	newTask.Left = 0
	for _, teamMember := range team.List {
		newTask.Estimate += teamMember.Estimate
		newTask.Consumed += teamMember.Consumed
		newTask.Left += teamMember.Left
	}

	//if(this->post->status) return currentTask;
	zeroTime, _ := time.Parse("2006-01-02", "0000-01-01")
	if newTask.Consumed == 0 {
		if newTask.Status == "" {
			newTask.Status = "wait"
		}
		newTask.FinishedBy = 0
		newTask.FinishedDate = zeroTime
	}

	if newTask.Consumed > 0 && newTask.Left > 0 {
		newTask.Status = "doing"
		newTask.FinishedBy = 0
		newTask.FinishedDate = zeroTime
	}
	var newAssigendTo, myAssigendTo *protocol.MSG_USER_team_info
	for _, teamMember := range team.List {
		if teamMember.Uid == newTask.AssignedTo {
			newAssigendTo = teamMember
		}
		if teamMember.Uid == in.GetUserID() {
			myAssigendTo = teamMember
		}
	}
	if newTask.Consumed > 0 && newTask.Left == 0 {
		if newAssigendTo != nil && oldTask.AssignedTo != team.List[len(team.List)-1].Uid {
			newTask.Status = "doing"
			newTask.FinishedBy = 0
			newTask.FinishedDate = zeroTime
		} else if oldTask.AssignedTo == team.List[len(team.List)-1].Uid {
			newTask.Status = "doing"
			newTask.FinishedBy = in.GetUserID()
			newTask.FinishedDate = now
		}
	}

	if (oldTask.AssignedTo != newTask.AssignedTo || newTask.Status == "done") && myAssigendTo != nil && myAssigendTo.Left == 0 && !libraries.In_slice(oldTask.FinishedList, in.GetUserID()) {
		newTask.FinishedList = append(newTask.FinishedList, in.GetUserID())
	}
	return nil

}
func task_updateParentStatus(taskID int32, in *protocol.Msg) (err error) {
	var childTask, oldtask, parentTask *db.Task
	err = in.DB.Table(db.TABLE_TASK).Where("Id=?", taskID).Find(&childTask)
	if err != nil {
		return
	}
	parentID := childTask.Parent
	if parentID < 1 {
		return nil
	}

	if err = task_computeWorkingHours(parentID, in); err != nil {
		return
	}
	var childrenStatus = make(map[string]bool)
	status := ""
	if m, e := in.DB.Table(db.TABLE_TASK).Where("Parent=? and Deleted = 0", parentID).Field("id,status").SelectMap(); e != nil {
		return e
	} else {
		for _, row := range m {
			status = row["status"]
			childrenStatus[row["status"]] = true
		}
	}

	if len(childrenStatus) > 1 {
		if childrenStatus["doing"] || childrenStatus["pause"] {
			status = "doing"
		} else if childrenStatus["wait"] {
			status = "wait"
		} else if childrenStatus["done"] {
			status = "done"
		} else if childrenStatus["closed"] {
			status = "closed"
		} else if childrenStatus["cancel"] {
			status = "cancel"
		}
	}
	if err = in.DB.Table(db.TABLE_TASK).Where("Id=?", parentID).Find(&parentTask); err != nil {
		return
	}
	if err = in.DB.Table(db.TABLE_TASK).Where("Id=?", parentID).Find(&oldtask); err != nil {
		return
	}
	if status != "" && parentTask.Status != status {
		now := time.Now()

		oldtask.Status = status
		if status == "done" {
			// task.AssignedTo   = parenttask.OpenedBy;
			oldtask.AssignedDate = now
			oldtask.FinishedBy = in.GetUserID()
			oldtask.FinishedDate = now
		}

		if status == "cancel" {
			//task.AssignedTo   = parenttask.OpenedBy;
			oldtask.AssignedDate = now
			oldtask.FinishedBy = 0
			oldtask.FinishedDate = protocol.ZEROTIME
			oldtask.CanceledBy = in.GetUserID()
			oldtask.CanceledDate = now
		}

		if status == "closed" {
			oldtask.AssignedTo = protocol.CLOSEUSER
			oldtask.AssignedDate = now
			oldtask.ClosedBy = in.GetUserID()
			oldtask.ClosedDate = now
			oldtask.ClosedReason = "done"
		}

		if status == "doing" {
			oldtask.FinishedBy = 0
			oldtask.FinishedDate = protocol.ZEROTIME
			oldtask.ClosedBy = 0
			oldtask.ClosedDate = protocol.ZEROTIME
			oldtask.ClosedReason = ""
		}

		oldtask.LastEditedBy = in.GetUserID()
		oldtask.LastEditedDate = now
		session, e := in.BeginTransaction()
		if e != nil {
			return e
		}
		defer func() {
			if err == nil {
				session.Commit()
			} else {
				session.Rollback()
			}
		}()
		if err = session.Table(db.TABLE_TASK).Replace(oldtask); err != nil {

			return
		}

		action := "Canceled"
		if status == "done" {
			action = "Finished"
		}
		if status == "closed" {
			action = "Closed"
		}
		if status == "pause" {
			action = "Paused"
		}
		if status == "internalaudit" {
			action = "internalaudit"
		}
		if status == "doing" && parentTask.Status == "wait" {
			action = "Started"
		}
		if status == "doing" && parentTask.Status == "pause" {
			action = "Restarted"
		}
		if status == "doing" && parentTask.Status != "wait" && parentTask.Status != "pause" {
			action = "Activated"
		}
		project := HostConn.GetProjectById(oldtask.Project)
		if actionID, e := in.ActionCreate("task", parentID, action, "", "", project.Products, project.Id); e != nil {
			return e
		} else {
			in.ActionLogHistory(actionID, oldtask, parentTask)
		}

		//2020-9-10 递归处理父节点
		if status == "done" && parentTask.Parent > 0 {
			return task_updateParentStatus(parentTask.Id, in)
		}

	}
	return nil
}
func task_computeWorkingHours(taskID int32, in *protocol.Msg) error {
	if taskID == 0 {
		return nil
	}
	var tasks []*db.Task
	err := in.DB.Table(db.TABLE_TASK).Where("Parent=? and status != 'cancel' and Deleted = 0", taskID).Select(&tasks)
	if err != nil || len(tasks) == 0 {
		return err
	}

	var estimate, consumed, left float64
	for _, task := range tasks {
		estimate += task.Estimate
		consumed += task.Consumed
		if task.Status != "closed" {
			left += task.Left
		}
	}

	_, err = in.DB.Table(db.TABLE_TASK).Where("Id=?", taskID).Update(map[string]interface{}{
		"estimate": estimate,
		"consumed": consumed,
		"left":     left,
	})

	return err
}
func task_updateParentAncestor(taskID int32, in *protocol.Msg) (err error) {
	var task *db.Task
	if err = in.DB.Table(db.TABLE_TASK).Prepare().Where("Id=?", taskID).Find(&task); err != nil {
		return
	}

	if task != nil {
		//获取子任务
		var children []*db.Task
		if err = in.DB.Table(db.TABLE_TASK).Field("Id").Where("Parent=?", taskID).Select(&children); err != nil {
			return
		}

		//检查是否有孙任务

		var ancestors []*db.Task
		if len(children) > 0 {
			if task.Parent == 0 {
				task.Parent = -1
			}
			if err = in.DB.Table(db.TABLE_TASK).Field("Id").Where(map[string]interface{}{"Parent": mysql.WhereOperatorIN(children)}).Select(&ancestors); err != nil {
				return
			}

		} else if task.Parent == -1 {
			task.Parent = 0
		}
		if len(ancestors) == 0 && task.Parent <= 0 {
			task.Ancestor = 0
		} else {
			task.Ancestor = -1
		}
		if _, err = in.DB.Table(db.TABLE_TASK).Where("Id=?", taskID).Update(map[string]interface{}{
			"Parent":   task.Parent,
			"Ancestor": task.Ancestor,
		}); err != nil {
			return
		}

	}
	return
}
func task_computeBeginAndEnd(taskID int32, in *protocol.Msg) (err error) {
	var tasks []*db.Task
	if err = in.DB.Table(db.TABLE_TASK).Where("Parent=? and Status!='cancel' and Deleted=0", taskID).Select(&tasks); err != nil {
		return
	}
	if len(tasks) == 0 {
		return
	}
	var earliestEstStarted, earliestRealStarted, latestDeadline int64
	for _, task := range tasks {

		if earliestEstStarted == 0 || (earliestEstStarted > task.EstStarted.Unix()) {
			earliestEstStarted = task.EstStarted.Unix()
		}
		if earliestRealStarted == 0 || (earliestRealStarted > task.RealStarted.Unix()) {
			earliestRealStarted = task.RealStarted.Unix()
		}
		if latestDeadline == 0 || (latestDeadline > task.Deadline.Unix()) {
			latestDeadline = task.Deadline.Unix()
		}
	}
	_, err = in.DB.Table(db.TABLE_TASK).Where("Id=?", taskID).Update(map[string]interface{}{
		"EstStarted":  time.Unix(earliestEstStarted, 0),
		"RealStarted": time.Unix(earliestRealStarted, 0),
		"Deadline":    time.Unix(latestDeadline, 0),
	})
	return
}

func task_assignTo(data *protocol.MSG_PROJECT_task_assignTo, in *protocol.Msg) {
	var oldTask, newTask *db.Task
	if err := in.DB.Table(db.TABLE_TASK).Where("Id=?", data.TaskID).Prepare().Find(&oldTask); err != nil {
		in.WriteErr(err)
		return
	}
	if oldTask.Status != "done" && oldTask.Status != "closed" && data.Left == 0 {
		in.WriteErr(protocol.Err_taskleftnotempty.Err())
		return
	}
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer func() {
		if err != nil {
			session.Rollback()
		} else {
			session.Commit()
		}
		in.WriteErr(err)
	}()
	now := time.Now()
	if err = protocol.CopyObj(oldTask, &newTask); err != nil {
		return
	}
	newTask.Left = data.Left
	newTask.AssignedTo = data.AssignedTo
	newTask.LastEditedBy = in.GetUserID()
	newTask.LastEditedDate = now
	newTask.AssignedDate = now
	if err = session.Table(db.TABLE_TASK).Where("Id=?", data.TaskID).Replace(newTask); err != nil {
		return
	}

	if len(oldTask.Team) > 0 {
		updateTeam := protocol.GET_MSG_USER_team_updateByWhere()
		updateTeam.Where = map[string]interface{}{
			"Root": data.TaskID,
			"Type": "task",
			"Uid":  oldTask.AssignedTo,
		}
		updateTeam.Update = map[string]interface{}{
			"Left": 0,
		}
		if err = in.SendMsgWaitResult(0, updateTeam, nil); err != nil {
			return
		}
		updateTeam.Where["Uid"] = data.AssignedTo
		updateTeam.Update["Left"] = data.Left
		if err = in.SendMsgWaitResult(0, updateTeam, nil); err != nil {
			return
		}
		if err = task_computeHours4Multiple(oldTask, newTask, in); err != nil {
			return
		}
	}
	if oldTask.Parent > 0 {
		if err = task_updateParentStatus(data.TaskID, in); err != nil {
			return
		}
	}
	project := HostConn.GetProjectById(oldTask.Project)
	user := HostConn.GetUserCacheById(data.AssignedTo)
	name := strconv.Itoa(int(data.AssignedTo))
	if user != nil {
		name = user.Realname
		if name == "" {
			name = user.Account
		}
	}
	var actionID int64
	if actionID, err = in.ActionCreate("task", data.TaskID, "Assigned", data.Comment, name, project.Products, project.Id); err != nil {
		return
	}
	_, err = in.ActionLogHistory(actionID, oldTask, newTask)
}

func task_start(data *protocol.MSG_PROJECT_task_start, in *protocol.Msg) {
	var oldtask, newtask *db.Task
	if err := in.DB.Table(db.TABLE_TASK).Where("Id=?", data.TaskID).Prepare().Find(&oldtask); err != nil {
		in.WriteErr(err)
		return
	}
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer func() {
		if err == nil {
			session.Commit()
		} else {
			in.WriteErr(err)
			session.Rollback()
		}
	}()
	if err = protocol.CopyObj(oldtask, &newtask); err != nil {
		return
	}
	now := time.Now()
	newtask.AssignedTo = in.GetUserID()
	newtask.LastEditedBy = in.GetUserID()
	newtask.LastEditedDate = now
	newtask.Status = "doing"
	newtask.RealStarted = data.RealStarted

	if oldtask.AssignedTo != in.GetUserID() {
		newtask.AssignedDate = now
	}
	if len(oldtask.Team) == 0 {
		newtask.Consumed = data.Consumed
		newtask.Left = data.Left
		if data.Left == 0 {
			newtask.Status = "done"
			newtask.FinishedBy = in.GetUserID()
			newtask.FinishedDate = now
		}
	}
	if len(oldtask.Team) > 0 {
		getteam := protocol.GET_MSG_USER_team_getByIds()
		getteam.Ids = oldtask.Team
		var result *protocol.MSG_USER_team_getByIds_result
		if err = in.SendMsgWaitResult(0, getteam, &result); err != nil {
			return
		}
		for _, team := range result.List {
			if team.Id == in.GetUserID() && data.Consumed < team.Consumed {
				err = protocol.Err_taskconsumedSmall.Err()
				return
			}
		}
		getteam.Put()
		result.Put()
		updateTeam := protocol.GET_MSG_USER_team_updateByWhere()
		updateTeam.Where = map[string]interface{}{"Root": data.TaskID, "Type": "task", "Uid": in.GetUserID()}
		updateTeam.Update = map[string]interface{}{
			"Consumed": data.Consumed,
			"Left":     data.Left,
		}
		if err = in.SendMsgWaitResult(0, updateTeam, nil); err != nil {
			return
		}
		if err = task_computeHours4Multiple(oldtask, newtask, in); err != nil {
			return
		}
	} else if data.Consumed < oldtask.Consumed {
		err = protocol.Err_taskconsumedSmall.Err()
		return
	}

	user := HostConn.GetUserCacheById(in.GetUserID())
	estimate := &protocol.MSG_PROJECT_TaskEstimate{
		Task:     data.TaskID,
		Date:     newtask.RealStarted,
		Left:     data.Left,
		Consumed: data.Consumed - oldtask.Consumed,
		Uid:      in.GetUserID(),
		Account:  user.Account,
	}
	if _, err = task_addTaskEstimate(estimate, in); err != nil {
		return
	}
	if err = session.Table(db.TABLE_TASK).Replace(newtask); err != nil {
		return
	}
	if oldtask.Parent > 0 {
		if err = task_updateParentStatus(data.TaskID, in); err != nil {
			return
		}
		if err = task_computeBeginAndEnd(oldtask.Parent, in); err != nil {
			return
		}
	}
	if err = story_setStage(oldtask.Story, in); err != nil {
		return
	}
	var changes protocol.ChangeHistory
	if changes, err = protocol.GetDiffChange(oldtask, newtask); err != nil {
		return
	}
	if data.Comment != "" || len(changes) > 0 {
		act := "Started"
		if data.MethodName == "restart" {
			act = "Restarted"
		}
		if data.Left == 0 {
			act = "Finished"
		}
		project := HostConn.GetProjectById(oldtask.Project)
		var actionID int64
		if actionID, err = in.ActionCreate("task", data.TaskID, act, data.Comment, "", project.Products, project.Id); err != nil {
			return
		}
		if len(changes) > 0 {
			changes.Add(actionID, in)
		}
	}
	out := protocol.GET_MSG_PROJECT_task_start_result()
	out.Changes = changes
	in.SendResult(out)
	out.Put()
}

func task_finish(data *protocol.MSG_PROJECT_task_finish, in *protocol.Msg) {
	var oldtask, newtask *db.Task
	if err := in.DB.Table(db.TABLE_TASK).Where("Id=?", data.TaskID).Prepare().Find(&oldtask); err != nil {
		in.WriteErr(err)
		return
	}
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer func() {
		if err == nil {
			session.Commit()
		} else {
			in.WriteErr(err)
			session.Rollback()
		}
	}()
	if err = protocol.CopyObj(oldtask, &newtask); err != nil {
		return
	}
	now := time.Now()
	newtask.Mailto = data.Mailto
	newtask.LastEditedBy = in.GetUserID()
	newtask.LastEditedDate = now
	if len(newtask.Team) == 0 {
		newtask.FinishedBy = in.GetUserID()
		newtask.FinishedDate = now
		newtask.Status = "done"
		newtask.Left = 0
	}
	newtask.Consumed = data.Consumed

	consumed := newtask.Consumed - oldtask.Consumed

	if len(oldtask.Team) > 0 {
		getteam := protocol.GET_MSG_USER_team_getByIds()
		getteam.Ids = oldtask.Team
		var result *protocol.MSG_USER_team_getByIds_result
		if err = in.SendMsgWaitResult(0, getteam, &result); err != nil {
			return
		}
		for _, team := range result.List {
			if team.Id == in.GetUserID() {
				if data.Consumed < team.Consumed {
					err = protocol.Err_taskconsumedSmall.Err()
					return
				}
				consumed = data.Consumed - team.Consumed
			}
		}
		getteam.Put()
		result.Put()
		updateTeam := protocol.GET_MSG_USER_team_updateByWhere()
		updateTeam.Where = map[string]interface{}{"Root": data.TaskID, "Type": "task", "Uid": in.GetUserID()}
		updateTeam.Update = map[string]interface{}{
			"Consumed": data.Consumed,
			"Left":     data.Left,
		}
		if err = in.SendMsgWaitResult(0, updateTeam, nil); err != nil {
			return
		}
		if err = task_computeHours4Multiple(oldtask, newtask, in); err != nil {
			return
		}
	} else if consumed < 0 {
		err = protocol.Err_taskconsumedSmall.Err()
		return
	}

	if consumed > 0 {
		user := HostConn.GetUserCacheById(in.GetUserID())
		estimate := &protocol.MSG_PROJECT_TaskEstimate{
			Task:     data.TaskID,
			Date:     time.Now(),
			Left:     0,
			Consumed: consumed,
			Uid:      in.GetUserID(),
			Account:  user.Account,
		}

		if _, err = task_addTaskEstimate(estimate, in); err != nil {
			return
		}
	}

	if err = session.Table(db.TABLE_TASK).Replace(newtask); err != nil {
		return
	}
	if oldtask.Parent > 0 {
		if err = task_updateParentStatus(data.TaskID, in); err != nil {
			return
		}
		if err = task_computeBeginAndEnd(oldtask.Parent, in); err != nil {
			return
		}
	}
	if err = story_setStage(oldtask.Story, in); err != nil {
		return
	}
	var changes protocol.ChangeHistory
	if changes, err = protocol.GetDiffChange(oldtask, newtask); err != nil {
		return
	}
	if data.Comment != "" || len(changes) > 0 {
		act := "Finished"
		project := HostConn.GetProjectById(oldtask.Project)
		var actionID int64
		if actionID, err = in.ActionCreate("task", data.TaskID, act, data.Comment, "", project.Products, project.Id); err != nil {
			return
		}
		if len(changes) > 0 {
			changes.Add(actionID, in)
		}
	}

	out := protocol.GET_MSG_PROJECT_task_finish_result()
	out.Changes = changes
	in.SendResult(out)
	out.Put()
}
func task_activate(data *protocol.MSG_PROJECT_task_activate, in *protocol.Msg) {
	var oldtask, newtask *db.Task
	if err := in.DB.Table(db.TABLE_TASK).Where("Id=?", data.TaskID).Prepare().Find(&oldtask); err != nil {
		in.WriteErr(err)
		return
	}
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer func() {
		if err == nil {
			session.Commit()
		} else {
			in.WriteErr(err)
			session.Rollback()
		}
	}()
	if err = protocol.CopyObj(oldtask, &newtask); err != nil {
		return
	}
	newtask.Left = data.Left
	newtask.Mailto = data.Mailto
	newtask.Status = "doing"
	newtask.FinishedBy = 0
	newtask.ClosedBy = 0
	newtask.CanceledBy = 0
	newtask.ClosedReason = ""
	newtask.FinishedDate = protocol.ZEROTIME
	newtask.CanceledDate = protocol.ZEROTIME
	newtask.ClosedDate = protocol.ZEROTIME
	newtask.LastEditedBy = in.GetUserID()
	newtask.LastEditedDate = time.Now()
	newtask.AssignedDate = time.Now()
	newtask.AssignedTo = data.AssignedTo
	if len(oldtask.Team) > 0 {
		updateTeam := protocol.GET_MSG_USER_team_updateByWhere()
		updateTeam.Where = map[string]interface{}{"Root": data.TaskID, "Type": "task", "Uid": data.AssignedTo}
		updateTeam.Update = map[string]interface{}{
			"Left": data.Left,
		}
		if err = in.SendMsgWaitResult(0, updateTeam, nil); err != nil {
			return
		}
		if err = task_computeHours4Multiple(oldtask, newtask, in); err != nil {
			return
		}
	}
	if err = session.Table(db.TABLE_TASK).Replace(newtask); err != nil {
		return
	}
	if oldtask.Parent > 0 {
		if err = task_updateParentStatus(data.TaskID, in); err != nil {
			return
		}
		if err = task_computeBeginAndEnd(oldtask.Parent, in); err != nil {
			return
		}
	}
	if err = story_setStage(oldtask.Story, in); err != nil {
		return
	}
	var changes protocol.ChangeHistory
	if changes, err = protocol.GetDiffChange(oldtask, newtask); err != nil {
		return
	}
	if data.Comment != "" || len(changes) > 0 {
		act := "Activated"
		project := HostConn.GetProjectById(oldtask.Project)
		var actionID int64
		if actionID, err = in.ActionCreate("task", data.TaskID, act, data.Comment, strconv.Itoa(int(data.AssignedTo)), project.Products, project.Id); err != nil {
			return
		}
		if len(changes) > 0 {
			changes.Add(actionID, in)
		}
	}
	in.WriteErr(nil)
}
func task_pause(data *protocol.MSG_PROJECT_task_pause, in *protocol.Msg) {
	var oldtask, newtask *db.Task
	if err := in.DB.Table(db.TABLE_TASK).Where("Id=?", data.TaskID).Prepare().Find(&oldtask); err != nil {
		in.WriteErr(err)
		return
	}
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer func() {
		if err == nil {
			session.Commit()
		} else {
			in.WriteErr(err)
			session.Rollback()
		}
	}()
	if err = protocol.CopyObj(oldtask, &newtask); err != nil {
		return
	}
	newtask.Status = "pause"
	newtask.LastEditedBy = in.GetUserID()
	newtask.LastEditedDate = time.Now()
	if err = session.Table(db.TABLE_TASK).Replace(newtask); err != nil {
		return
	}
	if oldtask.Parent > 0 {
		if err = task_updateParentStatus(data.TaskID, in); err != nil {
			return
		}
		if err = task_computeBeginAndEnd(oldtask.Parent, in); err != nil {
			return
		}
	}
	if err = story_setStage(oldtask.Story, in); err != nil {
		return
	}
	var changes protocol.ChangeHistory
	if changes, err = protocol.GetDiffChange(oldtask, newtask); err != nil {
		return
	}
	if data.Comment != "" || len(changes) > 0 {
		act := "Paused"
		project := HostConn.GetProjectById(oldtask.Project)
		var actionID int64
		if actionID, err = in.ActionCreate("task", data.TaskID, act, data.Comment, "", project.Products, project.Id); err != nil {
			return
		}
		if len(changes) > 0 {
			changes.Add(actionID, in)
		}
	}
	in.WriteErr(nil)
}

func task_internalaudit(data *protocol.MSG_PROJECT_task_internalaudit, in *protocol.Msg) {
	var oldtask, newtask *db.Task
	if err := in.DB.Table(db.TABLE_TASK).Where("Id=?", data.TaskID).Prepare().Find(&oldtask); err != nil {
		in.WriteErr(err)
		return
	}
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer func() {
		if err == nil {
			session.Commit()
		} else {
			in.WriteErr(err)
			session.Rollback()
		}
	}()
	if err = protocol.CopyObj(oldtask, &newtask); err != nil {
		return
	}
	newtask.Status = "internalaudit"
	newtask.LastEditedBy = in.GetUserID()
	newtask.LastEditedDate = time.Now()
	newtask.Mailto = data.Mailto
	if err = session.Table(db.TABLE_TASK).Replace(newtask); err != nil {
		return
	}
	if oldtask.Parent > 0 {
		if err = task_updateParentStatus(data.TaskID, in); err != nil {
			return
		}
		if err = task_computeBeginAndEnd(oldtask.Parent, in); err != nil {
			return
		}
	}
	if err = story_setStage(oldtask.Story, in); err != nil {
		return
	}
	var changes protocol.ChangeHistory
	if changes, err = protocol.GetDiffChange(oldtask, newtask); err != nil {
		return
	}
	if data.Comment != "" || len(changes) > 0 {
		act := "internalaudit"
		project := HostConn.GetProjectById(oldtask.Project)
		var actionID int64
		if actionID, err = in.ActionCreate("task", data.TaskID, act, data.Comment, "", project.Products, project.Id); err != nil {
			return
		}
		if len(changes) > 0 {
			changes.Add(actionID, in)
		}
	}
	in.WriteErr(nil)
}

func task_proofreading(data *protocol.MSG_PROJECT_task_proofreading, in *protocol.Msg) {
	var oldtask, newtask *db.Task
	if err := in.DB.Table(db.TABLE_TASK).Where("Id=?", data.TaskID).Prepare().Find(&oldtask); err != nil {
		in.WriteErr(err)
		return
	}
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer func() {
		if err == nil {
			session.Commit()
		} else {
			in.WriteErr(err)
			session.Rollback()
		}
	}()
	if err = protocol.CopyObj(oldtask, &newtask); err != nil {
		return
	}
	newtask.Proofreading = data.Proofreading
	newtask.LastEditedBy = in.GetUserID()
	newtask.LastEditedDate = time.Now()

	if err = session.Table(db.TABLE_TASK).Replace(newtask); err != nil {
		return
	}
	if oldtask.Parent > 0 {
		if err = task_updateParentStatus(data.TaskID, in); err != nil {
			return
		}
		if err = task_computeBeginAndEnd(oldtask.Parent, in); err != nil {
			return
		}
	}
	if err = story_setStage(oldtask.Story, in); err != nil {
		return
	}
	var changes protocol.ChangeHistory
	if changes, err = protocol.GetDiffChange(oldtask, newtask); err != nil {
		return
	}
	if data.Comment != "" || len(changes) > 0 {
		act := "proofreading"
		project := HostConn.GetProjectById(oldtask.Project)
		var actionID int64
		if actionID, err = in.ActionCreate("task", data.TaskID, act, data.Comment, "", project.Products, project.Id); err != nil {
			return
		}
		if len(changes) > 0 {
			changes.Add(actionID, in)
		}
	}
	in.WriteErr(nil)
}

func task_close(data *protocol.MSG_PROJECT_task_close, in *protocol.Msg) {
	var oldtask, newtask *db.Task
	if err := in.DB.Table(db.TABLE_TASK).Where("Id=?", data.TaskID).Prepare().Find(&oldtask); err != nil {
		in.WriteErr(err)
		return
	}
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer func() {
		if err == nil {
			session.Commit()
		} else {
			in.WriteErr(err)
			session.Rollback()
		}
	}()
	if err = protocol.CopyObj(oldtask, &newtask); err != nil {
		return
	}
	newtask.Status = "closed"
	newtask.AssignedTo = protocol.CLOSEUSER
	newtask.AssignedDate = time.Now()
	newtask.ClosedBy = in.GetUserID()
	newtask.LastEditedBy = in.GetUserID()
	newtask.LastEditedDate = time.Now()
	newtask.ClosedDate = time.Now()
	if oldtask.Status == "done" {
		newtask.ClosedReason = "done"
	} else if oldtask.Status == "cancel" {
		newtask.ClosedReason = "cancel"
	}

	if err = session.Table(db.TABLE_TASK).Replace(newtask); err != nil {
		return
	}
	if oldtask.Parent > 0 {
		if err = task_updateParentStatus(data.TaskID, in); err != nil {
			return
		}
		if err = task_computeBeginAndEnd(oldtask.Parent, in); err != nil {
			return
		}
	}
	if err = story_setStage(oldtask.Story, in); err != nil {
		return
	}
	var changes protocol.ChangeHistory
	if changes, err = protocol.GetDiffChange(oldtask, newtask); err != nil {
		return
	}
	if data.Comment != "" || len(changes) > 0 {
		act := "Closed"
		project := HostConn.GetProjectById(oldtask.Project)
		var actionID int64
		if actionID, err = in.ActionCreate("task", data.TaskID, act, data.Comment, "", project.Products, project.Id); err != nil {
			return
		}
		if len(changes) > 0 {
			changes.Add(actionID, in)
		}
	}

	in.WriteErr(nil)
}

func task_getStoryTaskCounts(data *protocol.MSG_PROJECT_task_getStoryTaskCounts, in *protocol.Msg) {
	where := map[string]interface{}{
		"Story":   data.Stories,
		"Deleted": false,
	}
	if data.ProjectID > 0 {
		where["Project"] = data.ProjectID
	}
	m, err := in.DB.Table(db.TABLE_TASK).Where(where).Field("story, COUNT(*) AS tasks").Group("story").SelectMap()
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_task_getStoryTaskCounts_result()
	out.List = make(map[int32]int)
	for _, row := range m {
		id, _ := strconv.Atoi(row["story"])
		count, _ := strconv.Atoi(row["tasks"])
		out.List[int32(id)] = count
	}
	in.SendResult(out)
	out.Put()
}
func task_examine(data *protocol.MSG_PROJECT_task_examine, in *protocol.Msg) {

	if _, err := in.DB.Table(db.TABLE_TASK).Where("Id=?", data.TaskID).Update(map[string]interface{}{"Examine": data.Examine}); err != nil {
		in.WriteErr(err)
		return
	}

	act := "examine"
	project := HostConn.GetProjectById(data.ProjectId)
	extra:="未通过"
	if data.Examine{
		extra="通过"
	}
	if _, err := in.ActionCreate("task", data.TaskID, act, "", extra, project.Products, project.Id); err != nil {
		in.WriteErr(err)
		return
	}

	in.WriteErr(nil)
}
func task_cancel(data *protocol.MSG_PROJECT_task_cancel, in *protocol.Msg) {

	var oldtask, newtask *db.Task
	if err := in.DB.Table(db.TABLE_TASK).Where("Id=?", data.TaskID).Prepare().Find(&oldtask); err != nil {
		in.WriteErr(err)
		return
	}
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer func() {
		if err == nil {
			session.Commit()
		} else {
			in.WriteErr(err)
			session.Rollback()
		}
	}()
	if err = protocol.CopyObj(oldtask, &newtask); err != nil {
		return
	}
	newtask.Status = "cancel"
	newtask.AssignedTo = oldtask.OpenedBy
	newtask.AssignedDate = time.Now()
	newtask.FinishedBy = 0
	newtask.FinishedDate = protocol.ZEROTIME
	newtask.CanceledDate = time.Now()
	newtask.LastEditedDate = time.Now()
	newtask.CanceledBy=in.GetUserID()
	newtask.LastEditedBy=in.GetUserID()


	if err = session.Table(db.TABLE_TASK).Replace(newtask); err != nil {
		return
	}
	if oldtask.Parent > 0 {
		if err = task_updateParentStatus(data.TaskID, in); err != nil {
			return
		}
		if err = task_computeBeginAndEnd(oldtask.Parent, in); err != nil {
			return
		}
	}
	if err = story_setStage(oldtask.Story, in); err != nil {
		return
	}
	var changes protocol.ChangeHistory
	if changes, err = protocol.GetDiffChange(oldtask, newtask); err != nil {
		return
	}
	if data.Comment != "" || len(changes) > 0 {
		act := "Canceled"
		project := HostConn.GetProjectById(oldtask.Project)
		var actionID int64
		if actionID, err = in.ActionCreate("task", data.TaskID, act, data.Comment, "", project.Products, project.Id); err != nil {
			return
		}
		if len(changes) > 0 {
			changes.Add(actionID, in)
		}
	}

	in.WriteErr(nil)
}
func task_delete(id int32,in *protocol.Msg)(err error){
	if _, err = in.DB.Table(db.TABLE_TASK).Where("Id=?", id).Update(map[string]interface{}{"Deleted":true}); err != nil {
		return
	}
	var tasks []*db.Task
	if err=in.DB.Table(db.TABLE_TASK).Where("Parent=?",id).Prepare().Select(&tasks);err!=nil{
		return
	}
	for _,task:= range tasks{
		if err= task_delete(task.Id,in);err!=nil{
			return
		}
	}
	return nil
}

func task_placeOrder(data *protocol.MSG_PROJECT_task_placeOrder,in *protocol.Msg){
	_,err:=in.DB.Table(db.TABLE_TASK).Where("Id=?",data.TaskID).Update(map[string]interface{}{"PlaceOrder":data.Action})
	in.WriteErr(err)
}