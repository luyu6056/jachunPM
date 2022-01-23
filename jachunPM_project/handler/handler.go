package handler

import (
	"jachunPM_project/db"
	"libraries"
	"protocol"
)

var HostConn *protocol.RpcClient

func Handler(in *protocol.Msg) {

	switch data := in.Data.(type) {
	case *protocol.MSG_PROJECT_tree_getLinePairs:
		tree_getLinePairs(data, in)
	case *protocol.MSG_PROJECT_product_insert:
		product_insert(data, in)
	case *protocol.MSG_PROJECT_product_getStories:
		product_getStories(data, in)
	case *protocol.MSG_PROJECT_tree_manageChild:
		tree_manageChild(data, in)
	case *protocol.MSG_PROJECT_product_getStoriesMapBySql:
		story_getStoriesMapBySql(data, in)
	case *protocol.MSG_PROJECT_tree_updateList:
		tree_updateList(data, in)
	case *protocol.MSG_PROJECT_tree_delete:
		tree_delete(data, in)
	case *protocol.MSG_PROJECT_product_update:
		product_update(data, in)
	case *protocol.MSG_PROJECT_productplan_getPairsForStory:
		productplan_getPairsForStory(data, in)
	case *protocol.MSG_PROJECT_productplan_getList:
		productplan_getList(data, in)
	case *protocol.MSG_PROJECT_productplan_getLast:
		productplan_getLast(data, in)
	case *protocol.MSG_PROJECT_product_editBranch:
		product_editBranch(data, in)
	case *protocol.MSG_PROJECT_product_deleteBranch:
		product_deleteBranch(data, in)
	case *protocol.MSG_PROJECT_productplan_getPairs:
		productplan_getPairs(data, in)
	case *protocol.MSG_PROJECT_productplan_insertUpdate:
		productplan_insertUpdate(data, in)
	case *protocol.MSG_PROJECT_productplan_delete:
		productplan_delete(data, in)
	case *protocol.MSG_PROJECT_stroy_create:
		story_create(data, in)
	case *protocol.MSG_PROJECT_story_batchGetStoryStage:
		story_batchGetStoryStage(data, in)
	case *protocol.MSG_PROJECT_story_getById:
		story, err := story_getById(data.Id, data.Version)
		if err != nil {
			in.WriteErr(err)
		} else {
			out := protocol.GET_MSG_PROJECT_story_getById_result()
			out.Story = story
			in.SendResult(out)
			out.Put()
		}

	case *protocol.MSG_PROJECT_tree_getParents:
		list, err := tree_getParents(data.ModuleID)
		if err != nil {
			in.WriteErr(err)
		} else {
			out := protocol.GET_MSG_PROJECT_tree_getParents_result()
			out.List = list
			in.SendResult(out)
			out.Put()
		}
	case *protocol.MSG_PROJECT_productplan_getById:
		productplan_getById(data, in)
	case *protocol.MSG_PROJECT_build_getById:
		build_getById(data, in)
	case *protocol.MSG_PROJECT_release_getById:
		release_getById(data, in)
	case *protocol.MSG_PROJECT_task_getPairs:
		task_getPairs(data, in)
	case *protocol.MSG_PROJECT_task_getListByWhereMap:
		task_getListByWhereMap(data, in)
	case *protocol.MSG_PROJECT_project_getBurn:
		project_getBurn(data, in)
	case *protocol.MSG_PROJECT_story_getPlanStories:
		story_getPlanStories(data, in)
	case *protocol.MSG_PROJECT_project_linkStory:
		project_linkStory(data, in)
	case *protocol.MSG_PROJECT_branch_getByProducts:
		branch_getByProducts(data, in)
	case *protocol.MSG_PROJECT_project_create:
		project_create(data, in)
	case *protocol.MSG_PROJECT_story_getPairsByIds:
		story_getPairsByIds(data, in)
	case *protocol.MSG_PROJECT_product_getPairsByIds:
		product_getPairsByIds(data, in)
	case *protocol.MSG_PROJECT_project_getPairsByIds:
		project_getPairsByIds(data, in)
	case *protocol.MSG_PROJECT_branch_getPairsByIds:
		branch_getPairsByIds(data, in)
	case *protocol.MSG_PROJECT_tree_getPairsByIds:
		tree_getPairsByIds(data, in)
	case *protocol.MSG_PROJECT_project_statRelatedData:
		project_statRelatedData(data, in)
	case *protocol.MSG_PROJECT_project_start:
		project_start(data, in)
	case *protocol.MSG_PROJECT_project_putoff:
		project_putoff(data, in)
	case *protocol.MSG_PROJECT_project_suspend:
		project_suspend(data, in)
	case *protocol.MSG_PROJECT_project_activate:
		project_activate(data, in)
	case *protocol.MSG_PROJECT_project_close:
		project_close(data, in)
	case *protocol.MSG_PROJECT_project_delete:
		project_delete(data, in)
	case *protocol.MSG_PROJECT_project_getProjectTasks:
		project_getProjectTasks(data, in)
	case *protocol.MSG_PROJECT_project_getProjectTasksByWhere:
		project_getProjectTasksByWhere(data, in)
	case *protocol.MSG_PROJECT_tree_getTaskTreeModules:
		tree_getTaskTreeModules(data, in)
	case *protocol.MSG_PROJECT_task_getById:
		task_getById(data, in)
	case *protocol.MSG_PROJECT_story_getProjectStoryPairs:
		story_getProjectStoryPairs(data, in)
	case *protocol.MSG_PROJECT_task_create:
		task_create(data, in)
	case *protocol.MSG_PROJECT_task_GetTaskEstimateByTaskId:
		task_GetTaskEstimateByTaskId(data, in)
	case *protocol.MSG_PROJECT_task_UpdateTaskEstimate:
		task_UpdateTaskEstimate(data, in)
	case *protocol.MSG_PROJECT_story_getProductStories:
		var out = protocol.GET_MSG_PROJECT_story_getProductStories_result()
		var err error
		out.List, err = story_getProductStories(data.Products, data.Branches, data.ModuleID, data.Status, data.Sort, data.Page, data.PerPage, &data.Total)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out.Total = data.Total
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_PROJECT_productplan_getForProducts:
		list, err := productplan_getForProducts(data.Products, in)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_PROJECT_productplan_getForProducts_result()
		out.List = list
		in.SendResult(out)
		out.Put()
	case *protocol.MSG_PROJECT_task_assignTo:
		task_assignTo(data, in)
	case *protocol.MSG_PROJECT_task_start:
		task_start(data, in)
	case *protocol.MSG_PROJECT_task_finish:
		task_finish(data, in)
	case *protocol.MSG_PROJECT_task_activate:
		task_activate(data, in)
	case *protocol.MSG_PROJECT_task_pause:
		task_pause(data, in)
	case *protocol.MSG_PROJECT_task_internalaudit:
		task_internalaudit(data, in)
	case *protocol.MSG_PROJECT_task_proofreading:
		task_proofreading(data, in)
	case *protocol.MSG_PROJECT_task_close:
		task_close(data, in)
	case *protocol.MSG_PROJECT_task_getStoryTaskCounts:
		task_getStoryTaskCounts(data, in)
	case *protocol.MSG_PROJECT_task_examine:
		task_examine(data, in)
	case *protocol.MSG_PROJECT_task_cancel:
		task_cancel(data, in)
	case *protocol.MSG_PROJECT_task_delete:
		in.WriteErr(task_delete(data.TaskID, in))
	case *protocol.MSG_PROJECT_task_placeOrder:
		task_placeOrder(data, in)
	case *protocol.MSG_PROJECT_getAllprojectProductID:
		var project []*db.Project
		var product []*db.Product
		err := in.DB.Table(db.TABLE_PROJECT).Limit(0).Select(&project)
		if err != nil {
			in.WriteErr(err)
			return
		}
		err = in.DB.Table(db.TABLE_PRODUCTPLAN).Limit(0).Select(&product)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_PROJECT_getAllprojectProductID_result()
		for _, v := range project {
			out.ProjectID = append(out.ProjectID, v.Id)
		}
		for _, v := range product {
			out.ProductID = append(out.ProductID, v.Id)
		}
		in.SendResult(out)
	case *protocol.MSG_PROJECT_doRawSelect:
		out := protocol.GET_MSG_PROJECT_doRawSelect_result()
		var err error
		if out.List, err = in.DB.Raw(data.Sql).SelectMap(); err != nil {
			in.WriteErr(err)
		} else {
			in.SendResult(out)
		}
	case *protocol.MSG_PROJECT_updateCache:
		switch data.Type {
		case "project":
			for _, id := range data.Ids {
				project_setCache(id)
			}
		case "product":
			for _, id := range data.Ids {
				product_setCache(id)
			}
		}
	default:
		if v, ok := protocol.CmdToName[in.Cmd]; ok {
			libraries.ReleaseLog("未设置消息CMD%s处理", v)
		} else {
			libraries.ReleaseLog("未设置消息CMD%d处理", in.Cmd)
		}
	}
}
