package handler

import (
	"errors"
	"jachunPM_project/db"
	"mysql"
	"protocol"
	"strconv"
	"strings"
	"time"
)

func project_getProducts(projectID int32) (products []*db.Product, err error) {
	err = HostConn.DB.Table(db.TABLE_PRODUCT).Where(db.TABLE_PROJECT+".Id = ?", projectID).LeftJoin(db.TABLE_PROJECT).On(db.TABLE_PRODUCT + ".Id = " + db.TABLE_PROJECT + ".Product").Limit(0).Select(&products)
	return
}
func project_setCache(id int32) {
	project := protocol.GET_MSG_PROJECT_project_cache()
	HostConn.DB.Table(db.TABLE_PROJECT).Prepare().Where("Id=?", id).Find(&project)
	if project.Id != 0 {

		out := protocol.GET_MSG_USER_team_getByTypeRoot()
		out.Type = "project"
		out.Root = []int32{project.Id}
		var result *protocol.MSG_USER_team_getByTypeRoot_result
		if err := (&protocol.RpclientSend{HostConn}).SendMsgWaitResultToDefault(nil, out, &result); err == nil {
			project.Teams = result.List
		}
		out.Put()
		HostConn.CacheSet(protocol.PATH_PROJECT_PROJECT_CACHE, strconv.Itoa(int(project.Id)), project, 0)

	}
	project.Put()
}
func project_getBurn(data *protocol.MSG_PROJECT_project_getBurn, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_project_getBurn_result()
	if err := HostConn.DB.Table(db.TABLE_BURN).Where(map[string]interface{}{"Project": data.ProjectIds}).Limit(0).Order("date desc").Select(&out.List); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
func project_linkStory(data *protocol.MSG_PROJECT_project_linkStory, in *protocol.Msg) {
	if len(data.Stories) == 0 {
		in.WriteErr(nil)
		return
	}
	project := HostConn.GetProjectById(data.ProjectID)
	if project == nil {
		in.WriteErr(errors.New(protocol.Err_ProjectNotFound.String()))
		return
	}
	session, err := in.BeginTransaction()
	defer session.Rollback()
	if err != nil {
		in.WriteErr(err)
		return
	}
	var ids []int32
	for _, storyId := range data.Stories {
		find := false
		for _, id := range project.Storys {
			if id == storyId {
				find = true
				break
			}
		}
		if find {
			continue
		}
		var story *db.Story
		if err = session.Table(db.TABLE_STORY).Prepare().Where("Id=?", storyId).Find(&story); err != nil {
			in.WriteErr(err)
			return
		}
		if story == nil {
			in.WriteErr(errors.New(protocol.Err_ProjectStoryNotFound.String()))
			return
		}
		story.Project = project.Id
		if _, err = session.Table(db.TABLE_STORY).Prepare().Where("Id=?", storyId).Update("Project=?", project.Id); err != nil {
			in.WriteErr(err)
			return
		}
		project.Storys = append(project.Storys, storyId)
		if productID, ok := data.Products[storyId]; ok {
			find := false
			for _, id := range project.Products {
				if id == productID {
					find = true
					break
				}
			}
			if !find {
				project.Products = append(project.Products, productID)
			}
		}
		protocol.Order_ascInt32(project.Storys)
		protocol.Order_ascInt32(project.Products)
		ids = append(ids, storyId)

	}

	err = session.Table(db.TABLE_PROJECT).Prepare().Where("Id=?", project.Id).Replace(project)
	if err != nil {
		in.WriteErr(err)
		return
	}
	session.CommitCallback(func() {
		project_setCache(project.Id)
		for _, id := range ids {
			var story *db.Story
			in.DB.Table(db.TABLE_STORY).Prepare().Where("Id=?", id).Find(&story)
			if story != nil {
				story_setStage(id, in)
				in.ActionCreate("story", id, "linked2project", "", strconv.Itoa(int(project.Id)), []int32{story.Product}, project.Id)
			}

		}
	})
	session.Commit()
	in.WriteErr(nil)
}
func project_create(data *protocol.MSG_PROJECT_project_create, in *protocol.Msg) {

	data.Info.Status = "wait"
	if data.Info.Id == 0 {
		c, err := in.DB.Table(db.TABLE_PROJECT).WhereOr(map[string]interface{}{"Name": data.Info.Name, "Code": data.Info.Code}).Count()
		if err != nil {
			in.WriteErr(err)
			return
		}
		if c > 0 {
			in.WriteErr(errors.New(protocol.Err_ProjectNameIsExist.String()))
			return
		}
	}

	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer session.Rollback()
	var id int64
	var oldproject *protocol.MSG_PROJECT_project_cache
	if data.Info.Id > 0 {
		oldproject = HostConn.GetProjectById(data.Info.Id)
		if oldproject == nil {
			in.WriteErr(errors.New(protocol.Err_ProjectNotFound.String()))
			return
		}
		if err = session.Table(db.TABLE_PROJECT).Replace(data.Info); err != nil {
			in.WriteErr(err)
			return
		}

		id = int64(data.Info.Id)
	} else {
		if id, err = session.Table(db.TABLE_PROJECT).Insert(data.Info); err != nil {
			in.WriteErr(err)
			return
		}
	}
	products := data.Info.Products
	session.CommitCallback(func() {
		project_setCache(int32(id))
		project_updateProducts(int32(id))

		if oldproject != nil {
			newproject := HostConn.GetProjectById(int32(id))
			actionID, err := in.ActionCreate("project", int32(id), "edited", "", "", products, int32(id))
			if err == nil {
				in.ActionLogHistory(actionID, oldproject, newproject)
			}
		} else {
			in.ActionCreate("project", int32(id), "opened", "", "", products, int32(id))
		}

	})
	creatorExists := false
	addteam := protocol.GET_MSG_USER_team_addByList()
	updateUserView := protocol.GET_MSG_USER_updateUserView()
	now := time.Now()
	/* Copy team of project. */
	if data.CopyProjectID > 0 {
		out := protocol.GET_MSG_USER_team_getByTypeRoot()
		out.Root = []int32{data.CopyProjectID}
		out.Type = "project"
		var result *protocol.MSG_USER_team_getByTypeRoot_result
		if err = in.SendMsgWaitResult(0, out, &result); err != nil {
			in.WriteErr(err)
			return
		}

		for _, m := range result.List {
			tmp := protocol.GET_MSG_USER_team_info()
			tmp.Root = int32(id)
			tmp.Join = now
			tmp.Days = data.Info.Days
			tmp.Type = "project"
			tmp.Account = m.Account
			tmp.Uid = m.Uid
			tmp.Role = m.Role
			tmp.Limited = m.Limited
			tmp.Hours = m.Hours
			addteam.List = append(addteam.List, tmp)
			if m.Uid == data.Info.OpenedBy {
				creatorExists = true
			}
			updateUserView.UserIds = append(updateUserView.UserIds, m.Uid)
		}

	}
	/* Add the creator to team. */
	if data.CopyProjectID == 0 || !creatorExists {
		tmp := protocol.GET_MSG_USER_team_info()
		tmp.Root = int32(id)
		tmp.Uid = data.Info.OpenedBy
		if user := HostConn.GetUserCacheById(data.Info.OpenedBy); user != nil {
			tmp.Account = user.Account
			tmp.Role = user.Role
			updateUserView.UserIds = append(updateUserView.UserIds, tmp.Uid)
		}

		tmp.Join = now
		tmp.Type = "project"
		tmp.Days = data.Info.Days
		in.LoadConfigToValue("project", "common", "defaultWorkhours", &tmp.Hours)
		addteam.List = append(addteam.List, tmp)

	}
	if err = in.SendMsgWaitResult(0, addteam, nil); err != nil {
		in.WriteErr(err)
		return
	}
	/* Create doc lib.
	   this->app->loadLang("doc");
	   lib = new stdclass();
	   lib->project = projectID;
	   lib->name    = this->lang->doclib->main["project"];
	   lib->type    = "project";
	   lib->main    = "1";
	   lib->acl     = project->acl == "open" ? "open" : "private";
	   this->dao->insert(TABLE_DOCLIB)->data(lib)->exec();*/
	updateUserView.ProjectIds = []int32{int32(id)}
	for _, id := range data.Info.Products {
		updateUserView.ProductIds = append(updateUserView.ProductIds, id)
	}
	for _, gid := range data.Info.Whitelist {
		updateUserView.GroupIds = append(updateUserView.GroupIds, gid)
	}
	if err = in.SendMsgWaitResult(0, updateUserView, nil); err != nil {
		in.WriteErr(err)
		return
	}
	updateUserView.Put()
	addteam.Put()
	session.Commit()
	result := protocol.GET_MSG_PROJECT_project_create_result()
	result.Id = int32(id)
	in.SendResult(result)
	result.Put()

}
func project_updateProducts(projectID int32) {
	return

}
func project_getPairsByIds(data *protocol.MSG_PROJECT_project_getPairsByIds, in *protocol.Msg) {
	var projects []*db.Project
	if err := in.DB.Table(db.TABLE_PROJECT).Field("Id,Name").Where(map[string]interface{}{"Id": data.Ids}).Limit(0).Select(&projects); err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_project_getPairsByIds_result()
	for _, p := range projects {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(p.Id)), p.Name})
	}
	in.SendResult(out)
	out.Put()
}
func project_statRelatedData(data *protocol.MSG_PROJECT_project_statRelatedData, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_project_statRelatedData_result()
	var err error
	defer func() {
		if err != nil {
			in.WriteErr(err)
		} else {
			in.SendResult(out)
		}
		out.Put()
	}()
	if out.StoryCount, err = in.DB.Table(db.TABLE_STORY).Where(map[string]interface{}{"Project": data.ProjectID, "Deleted": false}).Count(); err != nil {
		return
	}
	if out.TaskCount, err = in.DB.Table(db.TABLE_TASK).Where(map[string]interface{}{"Project": data.ProjectID, "Deleted": false, "Parent": 0}).Count(); err != nil {
		return
	}
	getCount := protocol.GET_MSG_TEST_bug_getCountByWhere()
	getCount.Where = map[string]interface{}{"Project": data.ProjectID, "Deleted": false}
	var result *protocol.MSG_TEST_bug_getCountByWhere_result
	if in.SendMsgWaitResult(0, getCount, &result); err != nil {
		in.WriteErr(err)
		return
	}
	out.BugCount = result.Count

}
func project_start(data *protocol.MSG_PROJECT_project_start, in *protocol.Msg) {
	project := HostConn.GetProjectById(data.Id)
	if project == nil {
		in.WriteErr(protocol.Err_ProjectNotFound.Err())
		return
	}
	session, err := in.BeginTransaction()
	defer func() {
		in.WriteErr(err)
		session.Rollback()
	}()
	if err != nil {
		return
	}
	if _, err = session.Table(db.TABLE_PROJECT).Prepare().Where("Id=?", data.Id).Update("Status=?", "doing"); err != nil {
		return
	}
	comment := data.Comment
	session.CommitCallback(func() {
		var newProject *protocol.MSG_PROJECT_project_cache
		in.DB.Table(db.TABLE_PROJECT).Prepare().Where("Id=?", project.Id).Find(&newProject)
		actionID, err := in.ActionCreate("project", project.Id, "Started", comment, "", project.Products, project.Id)
		if err == nil {
			in.ActionLogHistory(actionID, project, newProject)
		}
		project_setCache(project.Id)
	})
	session.Commit()
}
func project_putoff(data *protocol.MSG_PROJECT_project_putoff, in *protocol.Msg) {
	session, err := in.BeginTransaction()
	defer func() {
		session.Rollback()
		in.WriteErr(err)
	}()
	if err != nil {
		return
	}
	oldproject := HostConn.GetProjectById(data.Id)
	if oldproject == nil {
		err = protocol.Err_ProjectNotFound.Err()
		return
	}
	if _, err = session.Table(db.TABLE_PROJECT).Where("Id=?", data.Id).Update(map[string]interface{}{"Begin": data.Begin, "End": data.End, "Days": data.Days}); err != nil {
		return
	}
	id := data.Id
	comment := data.Comment
	session.CommitCallback(func() {
		var project *protocol.MSG_PROJECT_project_cache
		in.DB.Table(db.TABLE_PROJECT).Prepare().Where("Id=?", id).Find(&project)
		actionID, err := in.ActionCreate("project", project.Id, "Delayed", comment, "", project.Products, project.Id)
		if err == nil {
			in.ActionLogHistory(actionID, oldproject, project)
		}
		project_setCache(project.Id)
	})
	session.Commit()
}
func project_suspend(data *protocol.MSG_PROJECT_project_suspend, in *protocol.Msg) {
	session, err := in.BeginTransaction()
	defer func() {
		session.Rollback()
		in.WriteErr(err)
	}()
	if err != nil {
		return
	}
	oldproject := HostConn.GetProjectById(data.Id)
	if oldproject == nil {
		err = protocol.Err_ProjectNotFound.Err()
		return
	}
	if _, err = session.Table(db.TABLE_PROJECT).Where("Id=?", data.Id).Update(map[string]interface{}{"Status": "suspended"}); err != nil {
		return
	}
	id := data.Id
	comment := data.Comment
	session.CommitCallback(func() {
		var project *protocol.MSG_PROJECT_project_cache
		in.DB.Table(db.TABLE_PROJECT).Prepare().Where("Id=?", id).Find(&project)
		actionID, err := in.ActionCreate("project", project.Id, "Suspended", comment, "", project.Products, project.Id)
		if err == nil {
			in.ActionLogHistory(actionID, oldproject, project)
		}
		project_setCache(project.Id)
	})
	session.Commit()
}
func project_activate(data *protocol.MSG_PROJECT_project_activate, in *protocol.Msg) {
	session, err := in.BeginTransaction()
	defer func() {
		session.Rollback()
		in.WriteErr(err)
	}()
	if err != nil {
		return
	}
	oldproject := HostConn.GetProjectById(data.Id)
	if oldproject == nil {
		err = protocol.Err_ProjectNotFound.Err()
		return
	}
	if _, err = session.Table(db.TABLE_PROJECT).Where("Id=?", data.Id).Update(map[string]interface{}{"Status": "doing", "Begin": data.Begin, "End": data.End}); err != nil {
		return
	}
	id := data.Id
	comment := data.Comment
	session.CommitCallback(func() {
		var project *protocol.MSG_PROJECT_project_cache
		in.DB.Table(db.TABLE_PROJECT).Prepare().Where("Id=?", id).Find(&project)
		actionID, err := in.ActionCreate("project", project.Id, "Activated", comment, "", project.Products, project.Id)
		if err == nil {
			in.ActionLogHistory(actionID, oldproject, project)
		}
		project_setCache(project.Id)

	})
	if data.ReadjustTask {
		beginTimeStamp := data.Begin
		var tasks []*db.Task
		if err = session.Table(db.TABLE_TASK).Where(map[string]interface{}{"Deadline": []interface{}{mysql.WhereOperatorNE, "2000-01-01"}, "Status": []string{"wait", "doing"}, "Project": oldproject.Id}).Limit(0).Select(&tasks); err != nil {
			return
		}
		for _, task := range tasks {

			if task.Status == "wait" && task.EstStarted.After(protocol.NORMALTIME) {
				taskDays := task.Deadline.Sub(task.EstStarted)
				taskOffset := task.EstStarted.Sub(oldproject.Begin)

				estStartedTimeStamp := beginTimeStamp.Add(taskOffset)
				estStarted := estStartedTimeStamp
				deadline := estStartedTimeStamp.Add(taskDays)

				if estStarted.After(data.End) {
					estStarted = data.End
				}
				if deadline.After(data.End) {
					deadline = data.End
				}
				if _, err = session.Table(db.TABLE_TASK).Prepare().Where("Id=?", task.Id).Update("estStarted=? and deadline=?", estStarted.Format(protocol.TIMEFORMAT_MYSQLDATE), deadline.Format(protocol.TIMEFORMAT_MYSQLDATE)); err != nil {
					return
				}

			} else {
				taskOffset := task.Deadline.Sub(oldproject.Begin)
				deadline := beginTimeStamp.Add(taskOffset)
				if deadline.After(data.End) {
					deadline = data.End
				}
				if _, err = session.Table(db.TABLE_TASK).Prepare().Where("Id=?", task.Id).Update(" deadline=?", deadline.Format(protocol.TIMEFORMAT_MYSQLDATE)); err != nil {
					return
				}
			}
		}
	}
	session.Commit()
}
func project_close(data *protocol.MSG_PROJECT_project_close, in *protocol.Msg) {
	session, err := in.BeginTransaction()
	defer func() {
		session.Rollback()
		in.WriteErr(err)
	}()
	if err != nil {
		return
	}
	oldproject := HostConn.GetProjectById(data.Id)
	if oldproject == nil {
		err = protocol.Err_ProjectNotFound.Err()
		return
	}
	if _, err = session.Table(db.TABLE_PROJECT).Where("Id=?", data.Id).Update(map[string]interface{}{"Status": "closed", "ClosedBy": in.GetUserID(), "ClosedDate": time.Now()}); err != nil {
		return
	}

	id := data.Id
	comment := data.Comment
	session.CommitCallback(func() {
		var project *protocol.MSG_PROJECT_project_cache
		in.DB.Table(db.TABLE_PROJECT).Prepare().Where("Id=?", id).Find(&project)
		actionID, err := in.ActionCreate("project", project.Id, "Closed", comment, "", project.Products, project.Id)
		if err == nil {
			in.ActionLogHistory(actionID, oldproject, project)
		}
		project_setCache(project.Id)
	})
	session.Commit()
}
func project_delete(data *protocol.MSG_PROJECT_project_delete, in *protocol.Msg) {
	session, err := in.BeginTransaction()
	defer func() {
		session.Rollback()
		in.WriteErr(err)
	}()
	if err != nil {
		return
	}
	oldproject := HostConn.GetProjectById(data.Id)
	if oldproject == nil {
		err = protocol.Err_ProjectNotFound.Err()
		return
	}
	if _, err = session.Table(db.TABLE_PROJECT).Where("Id=?", data.Id).Update(map[string]interface{}{"Deleted": true}); err != nil {
		return
	}

	id := data.Id
	session.CommitCallback(func() {
		project_setCache(id)
	})
	session.Commit()
}
func project_getProjectTasks(data *protocol.MSG_PROJECT_project_getProjectTasks, in *protocol.Msg) {
	where := map[string]interface{}{
		"t1.Deleted": false,
	}
	if data.ProjectID > 0 {
		where["t1.Project"] = data.ProjectID
	}
	if data.ProductID != 0 {
		var trees []*db.Module
		if err := in.DB.Table(db.TABLE_MODULE).Field("Id").Where(map[string]interface{}{"Root": data.ProductID, "Type": "story"}).Limit(0).Select(&trees); err != nil {
			in.WriteErr(err)
			return
		}
		var treeIds []string
		for _, t := range trees {
			treeIds = append(treeIds, strconv.Itoa(int(t.Id)))
		}
		var storys []*db.Story
		if err := in.DB.Table(db.TABLE_MODULE).Field("Id").Where(map[string]interface{}{"Product": data.ProductID}).Limit(0).Select(&storys); err != nil {
			in.WriteErr(err)
			return
		}
		var storyIds []string
		for _, s := range storys {
			storyIds = append(storyIds, strconv.Itoa(int(s.Id)))
		}
		where["productRaw"] = []interface{}{mysql.WhereOperatorRAW, "`t1`.`Module` in (" + strings.Join(treeIds, ",") + ") or `t1`.`Story` in (" + strings.Join(storyIds, ",") + ")"}
	}
	if data.ModuleID != 0 {
		where["t1.Module"] = tree_getAllChildId(data.ModuleID)
	}
	switch data.Type[0] {
	case "all":
		where["t1.Parent"] = []interface{}{mysql.WhereOperatorLT, 1}
	case "myinvolved":
		out := protocol.GET_MSG_USER_team_getByTypeUid()
		out.Type = "task"
		out.Uid = in.GetUserID()
		var result *protocol.MSG_USER_team_getByTypeUid_result
		if err := in.SendMsgWaitResult(0, out, &result); err != nil {
			in.WriteErr(err)
			return
		}
		var ids []string
		for _, t := range result.List {
			ids = append(ids, strconv.Itoa(int(t.Root)))
		}
		if len(ids) == 0 {
			ids = []string{"0"}
		}
		where["myinvolvedRaw"] = []interface{}{mysql.WhereOperatorRAW, "`t1`.`Id` in (" + strings.Join(ids, ",") + ") or `t1`.`AssignedTo` = " + strconv.Itoa(int(in.GetUserID())) + " or `t1`.`Finishedby` = " + strconv.Itoa(int(in.GetUserID()))}
		out.Put()
		result.Put()
	case "undone":
		where["undoneRaw"] = []interface{}{mysql.WhereOperatorRAW, "`t1`.`Status` = 'wait' or `t1`.`Status` = 'doing'"}
		where["t1.Parent"] = []interface{}{mysql.WhereOperatorLT, 1}
	case "needconfirm":
		where["needconfirmRaw"] = []interface{}{mysql.WhereOperatorRAW, "`t2`.`version` > `t1`.`storyVersion` and `t2`.`Status` = 'active'"}
	case "openedBy":
		where["t1.OpenedBy"] = in.GetUserID()
	case "closedBy":
		where["t1.ClosedBy"] = in.GetUserID()
	case "canceledBy":
		where["t1.CanceledBy"] = in.GetUserID()
	case "finishedbyme":
		where["t1.Finishedby"] = in.GetUserID()

		//->andWhere('t1.finishedby', 1)->eq($this->app->user->account)->orWhere('t1.finishedList')->like("%,{$this->app->user->account},%")
	case "delayed":
		where["t1.Deadline"] = []interface{}{mysql.WhereOperatorBETWEEN, []string{"1970-01-01", time.Now().Format(protocol.TIMEFORMAT_MYSQLDATE)}}
		where["t1.Status"] = []string{"wait", "doing"}
	case "assignedtome":
		where["t1.AssignedTo"] = in.GetUserID()
	case "checkedItem":
		where["t1.Id"] = strings.Split(data.Type[1], ",")
	default:
		where["t1.Status"] = data.Type
		where["t1.Parent"] = []interface{}{mysql.WhereOperatorLT, 1}
	}

	var tasks []*protocol.MSG_PROJECT_TASK
	err := in.DB.Table(db.TABLE_TASK).Alias("t1").Field("DISTINCT t1.*, t2.Id AS StoryID, t2.Title AS StoryTitle, t2.Product as Product, t2.Branch as Branch, t2.version AS LatestStoryVersion, t2.Status AS StoryStatus").LeftJoin(db.TABLE_STORY).Alias("t2").On("t1.Story = t2.Id").Where(where).Order(data.OrderBy).Limit((data.Page-1)*data.PerPage, data.PerPage).Select(&tasks)
	if err != nil {
		in.WriteErr(err)
		return
	}
	total := data.Total
	if total == 0 {
		total, err = in.DB.Table(db.TABLE_TASK).Alias("t1").Field("DISTINCT t1.Id").LeftJoin(db.TABLE_STORY).Alias("t2").On("t1.Story = t2.Id").Where(where).Count()
		if err != nil {
			in.WriteErr(err)
			return
		}
	}
	if data.Type[0] != "assignedtome" && data.Type[0] != "finishedbyme" && data.Type[0] != "closedBy" && data.Type[0] != "canceledBy" && data.Type[0] != "checkedItem" {
		tasks, err = project_getProjectTaskOut(in, tasks)
		if err != nil {
			in.WriteErr(err)
			return
		}
	}

	out := protocol.GET_MSG_PROJECT_project_getProjectTasks_result()
	out.List = tasks
	out.Total = total

	in.SendResult(out)
	out.Put()
}

func project_getProjectTasksByWhere(data *protocol.MSG_PROJECT_project_getProjectTasksByWhere, in *protocol.Msg) {
	var tasks []*protocol.MSG_PROJECT_TASK
	err := in.DB.Table(db.TABLE_TASK).Where(data.Where).Order(data.OrderBy).Limit((data.Page-1)*data.PerPage, data.PerPage).Select(&tasks)
	if err != nil {
		in.WriteErr(err)
		return
	}
	total := data.Total
	if total == 0 {
		total, err = in.DB.Table(db.TABLE_TASK).Where(data.Where).Count()
		if err != nil {
			in.WriteErr(err)
			return
		}
	}
	//补全story消息
	var ids []int32
	for _, t := range tasks {
		if t.Story > 0 {
			ids = append(ids, t.Story)
		}
	}
	if len(ids) > 0 {
		var storys []*db.Story
		err = in.DB.Table(db.TABLE_STORY).Where(map[string]interface{}{"Id": ids}).Select(&storys)
		if err != nil {
			in.WriteErr(err)
			return
		}
		for _, t := range tasks {
			if t.Story > 0 {
				for _, s := range storys {
					if s.Id == t.Story {
						t.StoryID = s.Id
						t.StoryTitle = s.Title
						t.Product = s.Product
						t.Branch = s.Branch
						t.LatestStoryVersion = s.Version
						t.StoryStatus = s.Status
					}
				}
			}
		}
	}
	tasks, err = project_getProjectTaskOut(in, tasks)
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_project_getProjectTasks_result()
	out.List = tasks
	out.Total = total
	in.SendResult(out)
	out.Put()
}

//获取子孙任务并返回
func project_getProjectTaskOut(in *protocol.Msg, tasks []*protocol.MSG_PROJECT_TASK) ([]*protocol.MSG_PROJECT_TASK, error) {

	//获取子任务
	var parents, ancestors []int32
	for _, t := range tasks {
		if t.Parent == -1 {
			parents = append(parents, t.Id)
		}
	}
	where := map[string]interface{}{}
	var children []*protocol.MSG_PROJECT_TASK
	if len(parents) > 0 {
		where["t1.Parent"] = parents
		delete(where, "t1.Project")
		if err := in.DB.Table(db.TABLE_TASK).Alias("t1").Field("DISTINCT t1.*, t2.Id AS StoryID, t2.Title AS StoryTitle, t2.Product as Product, t2.Branch as Branch, t2.version AS LatestStoryVersion, t2.Status AS StoryStatus").LeftJoin(db.TABLE_STORY).Alias("t2").On("t1.Story = t2.Id").Where(where).Limit(0).Select(&children); err != nil {
			return nil, err
		}

		for _, child := range children {

			for _, task := range tasks {
				if task.Id == child.Parent {
					task.Children = append(task.Children, child)
					break
				}
			}
			if child.Ancestor == -1 {
				ancestors = append(ancestors, child.Id)
			}
		}

	}
	//获取孙任务
	if len(ancestors) > 0 {
		var grandchildrens []*protocol.MSG_PROJECT_TASK
		where["t1.Parent"] = ancestors
		if err := in.DB.Table(db.TABLE_TASK).Alias("t1").Field("DISTINCT t1.*, t2.Id AS StoryID, t2.Title AS StoryTitle, t2.Product as Product, t2.Branch as Branch, t2.version AS LatestStoryVersion, t2.Status AS StoryStatus").LeftJoin(db.TABLE_STORY).Alias("t2").On("t1.Story = t2.Id").Where(where).Limit(0).Select(&grandchildrens); err != nil {
			return nil, err
		}
		for _, grandchild := range grandchildrens {

			for _, child := range children {
				if child.Id == grandchild.Parent {
					grandchild.Ancestor = child.Parent //修复移动父任务到另一个任务，导致孙任务id没改过来的问题
					child.Grandchildren = append(child.Grandchildren, grandchild)
					break
				}
			}
		}
	}
	task_processTasks(tasks)
	return tasks, nil
}
