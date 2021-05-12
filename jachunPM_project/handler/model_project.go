package handler

import (
	"errors"
	"jachunPM_project/db"
	"protocol"
	"strconv"
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
		out.Root = project.Id
		var result *protocol.MSG_USER_team_getByTypeRoot_result
		if err := (&protocol.RpclientSend{HostConn}).SendMsgWaitResultToDefault(out, &result); err == nil {
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
	defer session.EndTransaction()
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
			in.WriteErr(errors.New(protocol.Err_ProjectStoryNotFount.String()))
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
				in.ActionCreate("story", id, "linked2project", "", strconv.Itoa(int(project.Id)), []int32{story.Product}, []int32{project.Id})
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
	defer session.EndTransaction()
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
			session.Rollback()
			return
		}

		id = int64(data.Info.Id)
	} else {
		if id, err = session.Table(db.TABLE_PROJECT).Insert(data.Info); err != nil {
			in.WriteErr(err)
			session.Rollback()
			return
		}
	}
	products := data.Info.Products
	session.CommitCallback(func() {
		project_setCache(int32(id))
		project_updateProducts(int32(id))

		if oldproject != nil {
			newproject := HostConn.GetProjectById(int32(id))
			actionID, err := in.ActionCreate("project", int32(id), "edited", "", "", products, []int32{int32(id)})
			if err == nil {
				in.ActionLogHistory(actionID, oldproject, newproject)
			}
		} else {
			in.ActionCreate("project", int32(id), "opened", "", "", products, []int32{int32(id)})
		}

	})
	creatorExists := false
	addteam := protocol.GET_MSG_USER_team_addByList()
	updateUserView := protocol.GET_MSG_USER_updateUserView()
	now := time.Now()
	/* Copy team of project. */
	if data.CopyProjectID > 0 {
		out := protocol.GET_MSG_USER_team_getByTypeRoot()
		out.Root = data.CopyProjectID
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
		session.Rollback()
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
		session.Rollback()
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
