package handler

import (
	"jachunPM_project/db"
	"libraries"
	"protocol"
	"strconv"
	"time"
)

func story_getProductStories(productID int32, branch int32, modules []int32, status []string, sort string, page int, perpage int, total *int) (list []*protocol.MSG_PROJECT_story, err error) {
	var where = map[string]interface{}{"product": productID, "deleted": false}
	if len(modules) > 0 {
		where["module"] = modules
	}
	if len(status) > 0 && status[0] != "all" {
		where["status"] = status
	}
	if branch > 0 {
		where["branch"] = branch
	}
	err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Order(sort).Limit((page-1)*perpage, perpage).Select(&list)
	if err != nil {
		return
	}
	if *total == 0 {
		*total, err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Count()
	}
	return
}
func story_getByPlan(productID, branch int32, modules []int32, plan, orderBy string, page int, perpage int, total *int) (list []*protocol.MSG_PROJECT_story, err error) {
	return story_getByField(productID, branch, modules, "plan", plan, orderBy, page, perpage, total)
}

func story_getByField(productID, branch int32, modules []int32, fieldName, fieldValue, orderBy string, page int, perpage int, total *int, operators ...string) (list []*protocol.MSG_PROJECT_story, err error) {
	operator := "equal"
	if len(operators) == 1 {
		operator = operators[0]
	}
	if fieldName == "" {
		return
	}
	var where = map[string]interface{}{"product": productID, "deleted": false}
	if branch > 0 {
		where["branch"] = branch
	}
	if len(modules) > 0 {
		where["module"] = modules
	}
	switch operator {
	case "equal":
		where[fieldName] = fieldValue
	case "include":
		where[fieldName] = []string{"like", "%" + fieldValue + "%"}
	}
	err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Order(orderBy).Limit((page-1)*perpage, perpage).Select(&list)
	if err != nil {
		return
	}
	if *total == 0 {
		*total, err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Count()
	}
	return
}
func story_getStoriesMapBySql(data *protocol.MSG_PROJECT_product_getStoriesMapBySql, in *protocol.Msg) {
	var limit []int
	if data.PerPage > 0 {
		limit = []int{(data.Page - 1) * data.PerPage, data.PerPage}
	}
	res, err := HostConn.DB.Table(db.TABLE_STORY).Field(data.Field).Where(data.Where).Order(data.Order).Group(data.Group).Limit(limit...).SelectMap()
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_product_getStoriesMapBySql_result()
	out.List = res
	in.SendResult(out)
	out.Put()
}
func story_create(data *protocol.MSG_PROJECT_stroy_create, in *protocol.Msg) {
	count, err := in.DB.Table(db.TABLE_STORY).Where(map[string]interface{}{"Deleted": false, "Product": data.Product, "Title": data.Title}).Count()
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_stroy_create_result()
	if count > 0 {
		out.Result = int32(protocol.Err_ProjectStoryTitleExists)
		in.SendResult(out)
		return
	}
	insert := new(db.Story)
	insert.Branch = data.Branch
	insert.Color = data.Color
	insert.FromBug = data.FromBug
	insert.Keywords = data.Keywords
	insert.Pri = data.Pri
	insert.Status = "active"
	insert.Stage = "wait"
	insert.Mailto = data.Mailto
	insert.OpenedBy = data.OpenedBy
	insert.Plan = data.Plan
	insert.Product = data.Product
	insert.Module = data.Module
	insert.Source = data.Source
	insert.SourceNote = data.SourceNote
	insert.Title = data.Title
	insert.Version = 1
	if check, err := story_checkForceReview(in, data.OpenedBy); check {
		insert.Status = "draft"
	} else if err != nil {
		in.WriteErr(err)
		return
	}

	if insert.Status == "draft" && data.Plan > 0 {
		insert.Stage = "planned"

	}
	session, err := in.BeginTransaction()
	defer func() {
		if err != nil {
			in.WriteErr(err)
			return
		}
		in.SendResult(out)
		session.Rollback()
	}()
	if err != nil {
		return
	}
	var insertId int64
	insertId, err = session.Table(db.TABLE_STORY).Insert(insert)
	if err != nil {
		return
	} else {

		if _, err = session.Table(db.TABLE_STORYSPEC).Insert(&db.StorySpec{
			Story:   int32(insertId),
			Version: 1,
			Title:   insert.Title,
			Spec:    data.Spec,
			Verify:  data.Verify,
		}); err != nil {
			return
		}

		if data.ProjectID > 0 && insert.Status != "draft" {
			if _, err = session.Table(db.TABLE_PROJECT).Where("Id = ?", data.ProjectID).Update("Story = ?", insertId); err != nil {
				return
			}
		}
		if data.FromBug > 0 {
			bug := protocol.GET_MSG_TEST_bug_updateMapById()
			bug.Id = data.FromBug
			bug.Update = map[string]interface{}{
				"toStory":      insertId,
				"status":       "closed",
				"resolution":   "tostory",
				"resolvedBy":   data.OpenedBy,
				"resolvedDate": time.Now(),
				"closedBy":     data.OpenedBy,
				"closedDate":   time.Now(),
				"assignedTo":   "closed",
				"assignedDate": time.Now(),
			}

			if err = in.SendMsgWaitResult(0, bug, nil); err != nil {
				return
			}
			in.ActionCreate("bug", data.FromBug, "ToStory", "", strconv.Itoa(int(insertId)), []int32{data.Product}, []int32{data.ProjectID})

			//$this->action->create('bug', $bugID, 'Closed');

			/* add files to story from bug. */
			fileupdate := protocol.GET_MSG_FILE_updateMapByWhere()
			fileupdate.Where = map[string]interface{}{
				"ObjectType": "bug",
				"ObjectID":   data.FromBug,
			}
			fileupdate.Update = map[string]interface{}{
				"ObjectType": "story",
				"ObjectID":   insertId,
			}
			if err = in.SendMsgWaitResult(0, fileupdate, nil); err != nil {
				return
			}
		}
		story_setStage(int32(insertId), in)
		out.Result = int32(insertId)
		/* Callback the callable method to process the related data for object that is transfered to story. */
		//if($from && is_callable(array($this, $this->config->story->fromObjects[$from]['callback']))) call_user_func(array($this, $this->config->story->fromObjects[$from]['callback']), $storyID);
		session.CommitCallback(func() {
			action := "Opened"
			extra := ""
			if data.FromBug > 0 {
				action = "Frombug"
				extra = strconv.Itoa(int(data.FromBug))
			}
			/*if(isset(fromObjectID)){
			    action = fromObjectAction;
			    extra  = fromObjectID;
			}*/

			in.ActionCreate("story", int32(insertId), action, "", extra, []int32{data.Product}, []int32{data.ProjectID})
		})
		session.Commit()
	}
}
func story_checkForceReview(in *protocol.Msg, id int32) (forceReview bool, err error) {
	config, err := in.LoadConfig("story")
	if err != nil {
		return
	}
	if i, ok := config["common"]["forceReview"]; ok {
		var list []int32
		libraries.JsonUnmarshal(libraries.JsonMarshal(i), &list)
		for _, v := range list {
			if id == v {
				forceReview = true
				break
			}
		}
	}
	return
}
func story_setStage(storyId int32, in *protocol.Msg) (err error) {
	session, e := in.BeginTransaction()
	if e != nil {
		return e
	}
	defer session.Rollback()
	_, err = session.Table(db.TABLE_STORYSTAGE).Where("story = ?", storyId).Delete()
	if err != nil {
		return
	}
	var story *db.Story
	if err = session.Table(db.TABLE_STORY).Prepare().Where("Id = ?", storyId).Find(&story); err != nil || story == nil {
		return
	}
	product := HostConn.GetProductById(story.Product)
	var projects []*db.Project
	if err = session.Table(db.TABLE_PROJECT).Prepare().Where("JSON_CONTAINS(`Products`,?) and JSON_CONTAINS(`Storys`,?)", product.Id, story.Id).Select(&projects); err != nil {
		return
	}

	hasBranch := product.Type != "normal" && story.Branch == 0
	stages := make(map[int32]string)
	if hasBranch && story.Plan > 0 {
		var plans []*db.Productplan
		if err = session.Table(db.TABLE_PRODUCTPLAN).Prepare().Where("id = ?", story.Plan).Select(&plans); err != nil {
			return
		}
		for _, plan := range plans {
			stages[plan.Branch] = "planned"
		}
	}
	var tasks []*db.Task
	/* If no projects, in plan, stage is planned. No plan, wait. */
	if len(projects) == 0 {
		if story.Plan > 0 {
			if _, err = session.Table(db.TABLE_STORY).Where("id = ?", story.Id).Update(map[string]interface{}{"stage": "planned"}); err != nil {
				return
			}
		} else {
			if _, err = session.Table(db.TABLE_STORY).Where("id = ?", story.Id).Update(map[string]interface{}{"stage": "wait"}); err != nil {
				return
			}
		}

		for branch, stage := range stages {
			if _, err = session.Table(db.TABLE_STORYSTAGE).Insert(&db.StoryStage{Story: storyId, Branch: branch, Stage: stage}); err != nil {
				return
			}

		}
	} else {
		var ids []int32
		for _, p := range projects {
			ids = append(ids, p.Id)
		}

		/* Search related tasks. */

		if err = session.Table(db.TABLE_TASK).Where(map[string]interface{}{
			"project":      ids,
			"story":        storyId,
			"type":         []string{"devel", "test"},
			"status":       []interface{}{"ne", "cancel"},
			"closedReason": []interface{}{"ne", "cancel"},
			"deleted":      false,
		}).Select(tasks); err != nil {
			return
		}

		/* No tasks, then the stage is projected. */
		if len(tasks) == 0 && len(projects) > 0 {

			for branch, _ := range stages {
				if _, err = session.Table(db.TABLE_STORYSTAGE).Insert(&db.StoryStage{Story: storyId, Branch: branch, Stage: "projected"}); err != nil {
					return
				}

			}
			if _, err = session.Table(db.TABLE_STORY).Where("id = ?", story.Id).Update(map[string]interface{}{"stage": "projected"}); err != nil {
				return
			}
		}
	}

	if hasBranch {
		for _, project := range projects {
			for _, id := range project.Branchs {
				stages[id] = "projected"
			}
		}
	}

	/* Cycle all tasks, get counts of every type && every status. */
	branchStatusList := map[int32]map[string]map[string]int{}
	branchDevelTasks := map[int32]int{}
	branchTestTasks := map[int32]int{}

	for _, task := range tasks {

		status := "wait"
		if task.Status != "" {
			if task.Status == "closed" {
				status = "done"
			} else {
				status = task.Status
			}

		}

		var branchs []int32
		for _, p := range projects {
			if p.Id == task.Project {
				branchs = p.Branchs
				break
			}
		}
		for _, branch := range branchs {
			if _, ok := branchStatusList[branch]; !ok {
				branchStatusList[branch] = map[string]map[string]int{
					"devel": map[string]int{"wait": 0, "doing": 0, "done": 0},
					"test":  map[string]int{"wait": 0, "doing": 0, "done": 0},
				}
			}
			if _, ok := branchStatusList[branch][task.Type]; !ok {
				branchStatusList[branch][task.Type] = make(map[string]int)
			}

			branchStatusList[branch][task.Type][status]++
			if task.Type == "devel" {
				branchDevelTasks[branch]++
			} else if task.Type == "test" {
				branchTestTasks[branch]++
			}
		}

	}

	/**
	 * Judge stage according to the devel && test tasks' status.
	 *
	 * 1. one doing devel task, all test tasks waiting, set stage as developing.
	 * 2. all devel tasks done, all test tasks waiting, set stage as developed.
	 * 3. one test task doing, set stage as testing.
	 * 4. all test tasks done, still some devel tasks not done(wait, doing), set stage as testing.
	 * 5. all test tasks done, all devel tasks done, set stage as tested.
	 */
	for branch, statusList := range branchStatusList {
		stage := "projected"
		testTasks := branchTestTasks[branch]
		develTasks := branchDevelTasks[branch]
		if statusList["devel"]["doing"] > 0 && statusList["test"]["wait"] == testTasks {
			stage = "developing"
		}
		if statusList["devel"]["wait"] > 0 && statusList["devel"]["done"] > 0 && statusList["test"]["wait"] == testTasks {
			stage = "developing"
		}
		if statusList["devel"]["done"] == develTasks && develTasks > 0 && statusList["test"]["wait"] == testTasks {
			stage = "developed"
		}
		if statusList["devel"]["done"] == develTasks && develTasks > 0 && statusList["test"]["wait"] > 0 && statusList["test"]["done"] > 0 {
			stage = "testing"
		}
		if statusList["test"]["doing"] > 0 {
			stage = "testing"
		}
		if (statusList["devel"]["wait"] > 0 || statusList["devel"]["doing"] > 0) && statusList["test"]["done"] == testTasks && testTasks > 0 {
			stage = "testing"
		}
		if statusList["devel"]["done"] == develTasks && develTasks > 0 && statusList["test"]["done"] == testTasks && testTasks > 0 {
			stage = "tested"
		}

		stages[branch] = stage
	}
	var releases []*db.Release
	if err = HostConn.DB.Table(db.TABLE_RELEASE).Where("JSON_CONTAINS(stories, ?) and Deleted = 0", storyId).Select(&releases); err != nil {
		return
	}

	for _, r := range releases {
		stages[r.Branch] = "released"
	}
	if len(stages) == 0 {
		return
	}
	if hasBranch {
		var stageList []string
		var stageListConfig []protocol.HtmlKeyValueStr
		if err := in.LoadConfigToValue("story", "common", "stageList", &stageListConfig); err != nil || len(stageListConfig) == 0 {
			return err
		} else {
			for _, kv := range stageListConfig {
				stageList = append(stageList, kv.Key)
			}
		}
		minStage := len(stageList) - 1
		for branch, stage := range stages {
			if _, err = session.Table(db.TABLE_STORYSTAGE).Insert(&db.StoryStage{Story: storyId, Branch: branch, Stage: stage}); err != nil {
				return
			}
			for k, v := range stageList {
				if v == stage {
					if k < minStage {
						minStage = k
					}
					break
				}
			}

		}

		if _, err = session.Table(db.TABLE_STORY).Where("id = ?", story.Id).Update(map[string]interface{}{"stage": stageList[minStage]}); err != nil {
			return
		}
	} else {
		for _, stage := range stages {
			if _, err = session.Table(db.TABLE_STORY).Where("id = ?", story.Id).Update(map[string]interface{}{"stage": stage}); err != nil {
				return
			}
			break
		}
	}
	session.Commit()
	return
}
func story_batchGetStoryStage(data *protocol.MSG_PROJECT_story_batchGetStoryStage, in *protocol.Msg) {
	var stages []*db.StoryStage
	if err := in.DB.Table(db.TABLE_STORYSTAGE).Where(map[string]interface{}{"Story": data.Ids}).Limit(0).Select(&stages); err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_story_batchGetStoryStage_result()
	out.List = make(map[int32][]protocol.HtmlKeyValueStr)
	for _, s := range stages {
		out.List[s.Story] = append(out.List[s.Story], protocol.HtmlKeyValueStr{strconv.Itoa(int(s.Branch)), s.Stage})
	}
	in.SendResult(out)
	out.Put()
}
func story_getBySearch(productID int32, branch int32, projectID int32, where map[string]interface{}, orderBy string, page int, perpage int, total *int) (list []*protocol.MSG_PROJECT_story, err error) {
	var products []*db.Product
	if projectID > 0 {
		products, err = project_getProducts(projectID)
	} else {
		if v, ok := where["Product"]; !ok {
			err = HostConn.DB.Table(db.TABLE_PRODUCT).Field("Id").Where("Deleted = 0").Limit(0).Order("Order asc").Select(&products)
		} else {
			if s, ok := v.(string); ok && s == "all" {
				delete(where, "Product")
			}
		}

	}
	if err != nil {
		return
	}
	var branches = map[int32]int{}
	if len(products) > 0 {
		var ids = make([]int32, len(products))
		for k, product := range products {
			ids[k] = product.Id
			for _, id := range product.Branch {
				branches[id] = 0
			}
		}
		where["Product"] = ids
	}
	if projectID > 0 {
		delete(branches, 0)
		if len(branches) > 0 {
			where["Branch"] = []int32{0}
			for id := range branches {
				where["Branch"] = append(where["Branch"].([]int32), id)
			}
		}

	} else if branch > 0 {
		if v, ok := where["Branch"]; !ok {
			where["Branch"] = []int32{0, branch}
		} else {
			if s, ok := v.(string); ok && s == "all" {
				delete(where, "Branch")
			}
		}
	}
	err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Order(orderBy).Limit((page-1)*perpage, perpage).Select(&list)
	if err != nil {
		return
	}
	if *total == 0 {
		*total, err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Count()
	}
	return
}
func story_getByAssignedTo(productID int32, branch int32, modules []int32, uid int32, orderBy string, page int, perpage int, total *int) (list []*protocol.MSG_PROJECT_story, err error) {
	return story_getByField(productID, branch, modules, "AssignedTo", strconv.Itoa(int(uid)), orderBy, page, perpage, total)
}
func story_getByOpenedBy(productID int32, branch int32, modules []int32, uid int32, orderBy string, page int, perpage int, total *int) (list []*protocol.MSG_PROJECT_story, err error) {
	return story_getByField(productID, branch, modules, "OpenedBy", strconv.Itoa(int(uid)), orderBy, page, perpage, total)
}
func story_getByReviewedBy(productID int32, branch int32, modules []int32, uid int32, orderBy string, page int, perpage int, total *int) (list []*protocol.MSG_PROJECT_story, err error) {
	return story_getByField(productID, branch, modules, "ReviewedBy", strconv.Itoa(int(uid)), orderBy, page, perpage, total)
}
func story_getByClosedBy(productID int32, branch int32, modules []int32, uid int32, orderBy string, page int, perpage int, total *int) (list []*protocol.MSG_PROJECT_story, err error) {
	return story_getByField(productID, branch, modules, "ClosedBy", strconv.Itoa(int(uid)), orderBy, page, perpage, total)
}
func story_getByStatus(productID int32, branch int32, modules []int32, status string, orderBy string, page int, perpage int, total *int) (list []*protocol.MSG_PROJECT_story, err error) {
	return story_getByField(productID, branch, modules, "Status", status, orderBy, page, perpage, total)
}
func story_get2BeClosed(productID int32, branch int32, modules []int32, orderBy string, page int, perpage int, total *int) (list []*protocol.MSG_PROJECT_story, err error) {
	where := map[string]interface{}{
		"Product": productID,
		"Deleted": false,
		"Stage":   []string{"developed", "released"},
		"Status":  []interface{}{"ne", "closed"},
	}
	if len(modules) > 0 {
		where["module"] = modules
	}

	if branch > 0 {
		where["branch"] = branch
	}
	err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Order(orderBy).Limit((page-1)*perpage, perpage).Select(&list)
	if err != nil {
		return
	}
	if *total == 0 {
		*total, err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Count()
	}
	return
}
func story_getById(id int32, version int16) (story *protocol.MSG_PROJECT_story, err error) {
	if err = HostConn.DB.Table(db.TABLE_STORY).Where("Id = ?", id).Find(&story); err != nil {
		return
	}
	if version == 0 {
		version = story.Version
	}
	var storyspec *db.StorySpec
	if err = HostConn.DB.Table(db.TABLE_STORYSPEC).Where("`Story` = ? and `Version` = ?", id, version).Find(&storyspec); err != nil {
		return
	}
	story.Title = storyspec.Title
	story.Spec = storyspec.Spec
	story.Verify = storyspec.Verify
	if story.Plan > 0 {
		var plan *db.Productplan
		if err = HostConn.DB.Table(db.TABLE_PRODUCTPLAN).Where(map[string]interface{}{"Id": story.Plan}).Find(&plan); err != nil {
			return
		}
		story.PlanTitle = plan.Title
	}
	if err = HostConn.DB.Table(db.TABLE_STORYSTAGE).Where("Story=?", story.Id).Select(&story.Stages); err != nil {
		return
	}
	if err = HostConn.DB.Table(db.TABLE_TASK).Prepare().Where("Story=? and Deleted = 0", story.Id).Order("Id desc").Select(&story.Tasks); err != nil {
		return
	}
	var projects []*db.Project
	if err = HostConn.DB.Table(db.TABLE_PROJECT).Field("Id").Where("Story=?", story.Id).Order("order desc").Select(&projects); err != nil {
		return
	}
	for _, p := range projects {
		story.Projects = append(story.Projects, p.Id)
	}
	var idsMap = map[int32]int{}
	if story.DuplicateStory > 0 {
		idsMap[story.DuplicateStory] = 1
	}
	for _, id := range story.LinkStories {
		idsMap[id] = 1
	}
	for _, id := range story.ChildStories {
		idsMap[id] = 1
	}
	if len(idsMap) > 0 {
		var ids []int32
		for id := range idsMap {
			ids = append(ids, id)
		}
		err = HostConn.DB.Table(db.TABLE_STORY).Where(map[string]interface{}{"Id": ids}).Limit(0).Select(&story.ExtraStories)
	}

	return
}
func story_getPlanStories(data *protocol.MSG_PROJECT_story_getPlanStories, in *protocol.Msg) {
	if data.OrderBy == "" {
		data.OrderBy = "id_desc"
	}
	where := map[string]interface{}{
		"Plan":    data.PlanID,
		"Deleted": false,
	}
	if data.Status != "" && data.Status != "all" {
		where["Status"] = data.Status
	}
	out := protocol.GET_MSG_PROJECT_story_getPlanStories_result()
	if err := in.DB.Table(db.TABLE_STORY).Where(where).Limit(0).Select(&out.List); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
func story_getPairsByIds(data *protocol.MSG_PROJECT_story_getPairsByIds, in *protocol.Msg) {
	var storys []*db.Story
	if err := in.DB.Table(db.TABLE_STORY).Field("Id,Title").Where(map[string]interface{}{"Id": data.Ids}).Limit(0).Select(&storys); err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_story_getPairsByIds_result()
	for _, s := range storys {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(s.Id)), s.Title})
	}
	in.SendResult(out)
	out.Put()
}
