package handler

import (
	"libraries"
	"protocol"
	"reflect"
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
	case *protocol.MSG_PROJECT_tree_getTaskTreeModules:
		tree_getTaskTreeModules(data, in)
	default:
		libraries.ReleaseLog("未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
	}
}
