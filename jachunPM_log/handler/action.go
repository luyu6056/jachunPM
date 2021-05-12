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

func action_crate(data *protocol.MSG_LOG_Action_Create, in *protocol.Msg) {
	if data.ActionType == "commented" && data.Comment == "" {
		return
	}

	insert := &db.Action{
		ObjectType: strings.ReplaceAll(data.ObjectType, "`", ""),
		ObjectID:   data.ObjectID,
		ActorId:    data.ActorId,
		Action:     strings.ToLower(data.ActionType),
		Date:       time.Now(),
		Extra:      data.Extra,
		Products:   data.Products,
		Projects:   data.Projects,
		Comment:    libraries.Html2bbcode(data.Comment),
	}
	libraries.DebugLog("%+v", data)
	if user := HostConn.GetUserCacheById(data.ActorId); user != nil {
		insert.Actor = user.Account
	}
	id, err := in.DB.Table(db.TABLE_ACTION).Insert(insert)
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_LOG_Action_Create_result()
	out.ActionId = id
	in.SendResult(out)
	out.Put()
	return
}
func action_GetByWhereMap(data *protocol.MSG_LOG_Action_GetByWhereMap, in *protocol.Msg) {
	out := protocol.GET_MSG_LOG_Action_GetByWhereMap_result()
	if err := in.DB.Table(db.TABLE_ACTION).Where(data.Where).Order(data.Order).Limit(0).Select(&out.List); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
func action_transformActions(where map[string]interface{}, order string, in *protocol.Msg) (res []*protocol.MSG_LOG_transformActions_info, err error) {
	if err = HostConn.DB.Table(db.TABLE_ACTION).WhereOr(where).Order(order).Limit(0).Select(&res); err != nil {
		return
	}
	objectTypes := make(map[string][]int32)
	for _, action := range res {
		action.ObjectName = "action_transformActions需要处理ObjectType:" + action.ObjectType
		objectTypes[action.ObjectType] = append(objectTypes[action.ObjectType], action.ObjectID)
	}
	objectNames := make(map[string][]protocol.HtmlKeyValueStr)
	for typename, ids := range objectTypes {
		switch typename {
		case "story":
			out := protocol.GET_MSG_PROJECT_story_getPairsByIds()
			out.Ids = ids
			var result *protocol.MSG_PROJECT_story_getPairsByIds_result
			if err = in.SendMsgWaitResult(0, out, &result); err != nil {
				return
			}
			objectNames[typename] = result.List
			out.Put()
		case "product":
			out := protocol.GET_MSG_PROJECT_product_getPairsByIds()
			out.Ids = ids
			var result *protocol.MSG_PROJECT_product_getPairsByIds_result
			if err = in.SendMsgWaitResult(0, out, &result); err != nil {
				return
			}
			objectNames[typename] = result.List
			libraries.DebugLog("%+v", ids)
			out.Put()
		case "productplan":
			out := protocol.GET_MSG_PROJECT_productplan_getPairs()
			out.Ids = ids
			var result *protocol.MSG_PROJECT_product_getPairsByIds_result
			if err = in.SendMsgWaitResult(0, out, &result); err != nil {
				return
			}
			objectNames[typename] = result.List
			out.Put()
		case "release":
		case "project":
			out := protocol.GET_MSG_PROJECT_project_getPairsByIds()
			out.Ids = ids
			var result *protocol.MSG_PROJECT_project_getPairsByIds_result
			if err = in.SendMsgWaitResult(0, out, &result); err != nil {
				return
			}
			objectNames[typename] = result.List
			out.Put()
		case "task":
		case "build":
		case "bug":
		case "case":
		case "testcase":
		case "testtask":
		case "testsuite":
		case "testreport":
		case "user":
			out := protocol.GET_MSG_USER_getPairs()
			var result *protocol.MSG_USER_getPairs_result
			if err = in.SendMsgWaitResult(0, out, &result); err != nil {
				return
			}
			for _, kv := range result.List {
				for _, id := range ids {
					if kv.Key == strconv.Itoa(int(id)) {
						objectNames[typename] = append(objectNames[typename], kv)
					}
				}
			}
			out.Put()
		case "doc":
		case "doclib":
		case "todo":
		case "custom":
		case "branch":
			out := protocol.GET_MSG_PROJECT_branch_getPairsByIds()
			out.Ids = ids
			var result *protocol.MSG_PROJECT_branch_getPairsByIds_result
			if err = in.SendMsgWaitResult(0, out, &result); err != nil {
				return
			}
			objectNames[typename] = result.List
			out.Put()
		case "module":
			out := protocol.GET_MSG_PROJECT_tree_getPairsByIds()
			out.Ids = ids
			var result *protocol.MSG_PROJECT_tree_getPairsByIds_result
			if err = in.SendMsgWaitResult(0, out, &result); err != nil {
				return
			}
			objectNames[typename] = result.List
			out.Put()
		case "caselib":
		case "entry":
		case "webhook":
		case "attend":
		case "attendstat":
		case "holiday":
		case "leave":
		case "overtime":
		case "lieu":
		case "trip":
		}
	}
	config, err := in.LoadConfig("action")
	if err != nil {
		return
	}
	for i := len(res) - 1; i >= 0; i-- {
		action := res[i]
		/* Add name field to the actions. */
		if v, ok := objectNames[action.ObjectType]; ok {
			for _, kv := range v {
				if kv.Key == strconv.Itoa(int(action.ObjectID)) {
					action.ObjectName = kv.Value
				}
			}
		}
		action.ActionLabel = action.Action
		if s, ok := config["label"][action.Action].(string); ok {
			action.ActionLabel = s
		}
		action.ObjectLabel = action.ObjectType
		if v, ok := config["label"][action.ObjectType]; ok {
			if s, ok := v.(string); ok {
				action.ObjectLabel = s
			} else {
				if m, ok := v.(map[string]string); ok {
					if v, ok := m[action.Action]; ok {
						action.ObjectLabel = v
					}
				}
			}
		}
		/* Other actions, create a link. */
		label := strings.Split(action.ObjectLabel, "|")
		if len(label) == 4 {
			objectLabel, moduleName, methodName, vars := label[0], label[1], label[2], label[3]
			if !in.HasPriv(moduleName, methodName) {
				res = append(res[:i], res[i+1:]...)
				continue
			}

			if action.ObjectType == "user" {
				action.ObjectLink = protocol.CreateLink(moduleName, methodName, fmt.Sprintf(vars, action.ObjectName))
			} else {
				action.ObjectLink = protocol.CreateLink(moduleName, methodName, fmt.Sprintf(vars, action.ObjectID))
			}
			action.ObjectLabel = objectLabel
		} else {
			if v, ok := config["objectTypes"][action.ObjectLabel].(string); ok {
				action.ObjectLabel = v
			}

		}
		if list, ok := config["majorList"][action.ObjectType].([]string); ok {
			for _, v := range list {
				if v == action.Action {
					action.Major = true
					break
				}
			}
		}
	}
	return
}
func action_AddHistory(data *protocol.MSG_LOG_Action_AddHistory, in *protocol.Msg) {
	in.DB.Table(db.TABLE_ACTION).Prepare().Where("Id=?", data.Id).Update("History = JSON_ARRAY_APPEND(History, '$', ?)", data.History)
}
