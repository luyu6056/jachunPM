package handler

import (
	"config"
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
		Project:    data.Project,
		Comment:    libraries.Html2bbcode(data.Comment),
	}
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
func action_transformActions(data *protocol.MSG_LOG_Action_transformActions, in *protocol.Msg) {
	out := protocol.GET_MSG_LOG_Action_transformActions_result()
	var err error
	defer func() {
		if err != nil {
			in.WriteErr(err)
		} else {
			in.SendResult(out)
			out.Put()
		}
	}()
	if err = HostConn.DB.Table(db.TABLE_ACTION).Prepare().WhereOr(data.Where).Order(data.Order).Limit((data.Page-1)*data.PerPage, data.Page*data.PerPage).Select(&out.List); err != nil {
		return
	}
	if data.Total == 0 {
		if data.Total, err = HostConn.DB.Table(db.TABLE_ACTION).WhereOr(data.Where).Order(data.Order).Count(); err != nil {
			return
		}
	}
	objectTypes := make(map[string][]int32)
	for _, action := range out.List {
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
			out := protocol.GET_MSG_PROJECT_task_getPairs()
			out.Where = map[string]interface{}{"Id": ids}
			var result *protocol.MSG_PROJECT_task_getPairs_result
			if err = in.SendMsgWaitResult(0, out, &result); err != nil {
				return
			}
			objectNames[typename] = result.List
			out.Put()
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
	config := config.Config[in.Lang]["action"]

	for i := len(out.List) - 1; i >= 0; i-- {

		action := out.List[i]
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
				out.List = append(out.List[:i], out.List[i+1:]...)
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
}
func action_AddHistory(data *protocol.MSG_LOG_Action_AddHistory, in *protocol.Msg) {
	var action *protocol.MSG_LOG_Action
	err := in.DB.Table(db.TABLE_ACTION).Prepare().Where("Id=?", data.Id).Find(&action)
	if err != nil {
		libraries.ReleaseLog("增加history失败%+v", err)
	}
	if action != nil {
		action.Historys = append(action.Historys, data.History...)
		_, err = in.DB.Table(db.TABLE_ACTION).Where("Id=?", data.Id).Update("Historys = ?", action.Historys)
		if err != nil {
			libraries.ReleaseLog("增加history失败%+v", err)
		}
	}

}
func action_read(data *protocol.MSG_LOG_Action_set_read, in *protocol.Msg) {
	_, err := in.DB.Table(db.TABLE_ACTION).Prepare().Where("'objectType' = ? and 'objectID' = ? and 'read' = 0", data.ObjectType, data.ObjectID).Update(map[string]interface{}{"read": 1})
	if err != nil {
		libraries.ReleaseLog("修改action为read失败，data:%+v ,err:%+v", data, err)
	}
}
