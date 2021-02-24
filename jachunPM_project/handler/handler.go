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
	default:
		libraries.ReleaseLog("未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
	}
}
