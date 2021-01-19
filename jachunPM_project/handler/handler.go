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
	default:
		libraries.ReleaseLog("未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
	}
}
