package handler

import (
	"fmt"
	"html"
	"html/template"
	"libraries"
	"mysql"
	"protocol"
	"strconv"
	"strings"
	"time"
)

func actionModelFuncs() {

	global_Funcs["action_printAction"] = func(data *TemplateData, action *protocol.MSG_LOG_Action, descExt ...string) template.HTML {
		var desc interface{}
		if action.ObjectType == "" || action.Action == "" {
			return template.HTML("")
		}
		objectType := action.ObjectType
		actionType := strings.ToLower(action.Action)

		if desc == nil {
			if objectTypeMap, ok1 := data.Lang[objectType]; ok1 {
				if actionTypeMap, ok2 := objectTypeMap["action"].(map[string]interface{}); ok2 {
					if actionTypeMap, ok3 := actionTypeMap[actionType].(map[string]interface{}); ok3 {
						desc = actionTypeMap[actionType]
					}

				}
			}
			if desc == nil && data.Lang["action"]["desc"] != nil && data.Lang["action"]["desc"].(map[string]string)[actionType] != "" {
				desc = data.Lang["action"]["desc"].(map[string]string)[actionType]
			} else {
				if action.Extra != "" {
					desc = data.Lang["action"]["desc"].(map[string]string)["extra"]
				} else {
					desc = data.Lang["action"]["desc"].(map[string]string)["common"]
				}
			}

		}
		switch i := desc.(type) {
		case string:
			i = strings.ReplaceAll(i, "$date", action.Date.Format("2006-01-02 15:04"))
			name := ""
			if actor := HostConn.GetUserCacheById(action.ActorId); actor != nil {
				name = actor.Realname
				if name == "" {
					name = actor.Account
				}
			}
			i = strings.ReplaceAll(i, "$actor", name)
			i = strings.ReplaceAll(i, "$extra", action.Extra)
			return template.HTML(i)
		case map[string]string:
			s := i["main"]
			s = strings.ReplaceAll(s, "$date", action.Date.Format("2006-01-02 15:04"))
			actor := HostConn.GetUserCacheById(action.ActorId)
			name := actor.Realname
			if name == "" {
				name = actor.Account
			}
			s = strings.ReplaceAll(s, "$actor", name)
			extra := strings.ToLower(action.Extra)
			if extraMap, ok := data.Lang[objectType][i["extra"]].(map[string]string); ok {
				if replacestr, ok2 := extraMap[extra]; ok2 {
					return template.HTML(strings.ReplaceAll(s, "$extra", replacestr))
				}

			}
			return template.HTML(strings.ReplaceAll(s, "$extra", action.Extra))
		}
		return template.HTML("")
	}

	global_Funcs["action_printChanges"] = func(data *TemplateData, objectType string, histories []*protocol.MSG_LOG_History, canChangeTagExt ...bool) template.HTML {

		if len(histories) == 0 {
			return template.HTML("")
		}
		canChangeTag := true
		if len(canChangeTagExt) == 1 {
			canChangeTag = canChangeTagExt[0]
		}
		var users []protocol.HtmlKeyValueStr
		var ok bool
		if users, ok = data.Data["users"].([]protocol.HtmlKeyValueStr); !ok {
			users, _ = user_getPairs(data, "")
		}
		maxLength := 0
		historiesWithDiff := make([]*protocol.MSG_LOG_History, 0, len(histories))
		historiesWithoutDiff := make([]*protocol.MSG_LOG_History, 0, len(histories))

		for _, history := range histories {
			history.FieldLabel = history.Field
			if data.Lang[objectType] != nil {
				if str, ok := data.Lang[objectType][history.Field].(string); ok {
					history.FieldLabel = str
				}
			}
			if len(history.FieldLabel) > maxLength {
				maxLength = len(history.FieldLabel)
			}
			if history.Diff == "" {
				historiesWithoutDiff = append(historiesWithoutDiff, history)
			} else {
				historiesWithDiff = append(historiesWithDiff, history)
			}

		}
		buf := bufpool.Get().(*libraries.MsgBuffer)
		for _, history := range append(historiesWithoutDiff, historiesWithDiff...) {
			if c := len(history.FieldLabel) - maxLength; c > 0 {
				history.FieldLabel += strings.Repeat(data.Lang["action"]["label"].(map[string]interface{})["space"].(string), c)
			}

			if history.Diff != "" {
				history.Diff = history_Diff_srp.Replace(history.Diff)
				if history.Field != "subversion" && history.Field != "git" {
					history.Diff = html.EscapeString(history.Diff)
				}

				history.Diff = history_Diff_Resrp.Replace(history.Diff)

				noTagDiff := ""
				if canChangeTag {
					noTagDiff, _ = libraries.Preg_replace(`/&lt;\/?([a-z][a-z0-9]*)[^\/]*\/?&gt;/Ui`, "", history.Diff)
				}
				buf.WriteString(fmt.Sprintf(data.Lang["action"]["desc"].(map[string]string)["diff2"], history.FieldLabel, noTagDiff, history.Diff))
			} else {
				if objtypeValue, ok := data.Lang["action"]["descValue"].(map[string]map[string]interface{})[objectType]; ok {
					if value, ok := objtypeValue[strings.ToLower(history.Field)]; ok {
						history.Old = fmt.Sprint(common_getValue(value, history.Old))
						history.New = fmt.Sprint(common_getValue(value, history.New))
					}
				}
				switch strings.ToLower(history.Field) {
				case "mailto":
					var ids []int
					err := libraries.JsonUnmarshalStr(history.Old, &ids)
					if err == nil && len(users) > 0 {
						var names []string
						for _, id := range ids {
							for _, kv := range users {
								if strconv.Itoa(id) == kv.Key {
									names = append(names, kv.Value)
								}
							}
						}
						history.Old = strings.Join(names, ",")
					}
					ids = nil
					err = libraries.JsonUnmarshalStr(history.New, &ids)
					if err == nil && len(users) > 0 {
						var names []string
						for _, id := range ids {
							for _, kv := range users {
								if strconv.Itoa(id) == kv.Key {
									names = append(names, kv.Value)
								}
							}
						}
						history.New = strings.Join(names, ",")
					}
				}
				buf.WriteString(fmt.Sprintf(data.Lang["action"]["desc"].(map[string]string)["diff1"], history.FieldLabel, history.Old, history.New))
			}
		}
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
}

var history_Diff_srp = strings.NewReplacer(`<ins>`, `[ins]`, "</ins>", "[/ins]", "<del>", "[del]", "</del>", "[/del]")
var history_Diff_Resrp = strings.NewReplacer(`[ins]`, `<ins>`, "[/ins]", `</ins>`, "[del]", "<del>", "[/del]", "</del>", "\n", "<br>")

func action_getList(data *TemplateData, objectType string, objectID int32) (actions []*protocol.MSG_LOG_Action, err error) {
	out := protocol.GET_MSG_LOG_Action_GetByWhereMap()
	defer out.Put()
	if objectType == "project" {
		out.Where = map[string]interface{}{
			"ObjectType": []string{"project", "testtask", "build"},
			"Projects":   []interface{}{mysql.WhereOperatorJSONCONTAINS, objectID},
		}
	} else {
		out.Where = map[string]interface{}{
			"ObjectType": objectType,
			"ObjectID":   objectID,
		}
	}
	out.Order = "date,id"
	var result *protocol.MSG_LOG_Action_GetByWhereMap_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	if objectType == "project" {

		/*if err = action_processProjectActions(result.List); err != nil {
			return
		}*/
	}
	for _, action := range result.List {
		if user := HostConn.GetUserCacheById(action.ActorId); user != nil {
			action.Actor = user.Realname
			if action.Actor == "" {
				action.Actor = user.Account
			}
		}
		switch action.Action {
		case "linked2project":
			id, _ := strconv.Atoi(action.Extra)
			if id > 0 {
				if project := HostConn.GetProjectById(int32(id)); project != nil {
					if hasPriv(data, "project", "story") {
						action.Extra = html_a(createLink("project", "story", "projectID="+action.Extra), project.Name)
					} else {
						action.Extra = project.Name
					}
				}
			}
		case "linked2plan", "unlinkedfromplan":
			id, _ := strconv.Atoi(action.Extra)
			if id > 0 {
				getProductPlan := protocol.GET_MSG_PROJECT_productplan_getById()
				getProductPlan.Id = int32(id)

				var result *protocol.MSG_PROJECT_productplan_getById_result
				if err = data.SendMsgWaitResultToDefault(getProductPlan, &result); err != nil {
					return
				}
				if result.Info != nil {
					if hasPriv(data, "productplan", "view") {
						action.Extra = html_a(createLink("productplan", "view", "planID="+action.Extra), result.Info.Title)
					} else {
						action.Extra = result.Info.Title
					}
				}
			}
		case "linked2build", "linked2bug", "unlinkedfrombuild":
			id, _ := strconv.Atoi(action.Extra)
			if id > 0 {
				getBuild := protocol.GET_MSG_PROJECT_build_getById()
				getBuild.Id = int32(id)

				var result *protocol.MSG_PROJECT_build_getById_result
				if err = data.SendMsgWaitResultToDefault(getBuild, &result); err != nil {
					return
				}
				if result.Info != nil {
					if hasPriv(data, "build", "view") {
						action.Extra = html_a(createLink("build", "view", "builID="+action.Extra+"&type="+action.ObjectType), result.Info.Name)
					} else {
						action.Extra = result.Info.Name
					}
				}
			}
		case "linked2release", "unlinkedfromrelease":
			id, _ := strconv.Atoi(action.Extra)
			if id > 0 {
				getRelease := protocol.GET_MSG_PROJECT_release_getById()
				getRelease.Id = int32(id)

				var result *protocol.MSG_PROJECT_release_getById_result
				if err = data.SendMsgWaitResultToDefault(getRelease, &result); err != nil {
					return
				}
				if result.Info != nil {
					if hasPriv(data, "release", "view") {
						action.Extra = html_a(createLink("release", "view", "releaseID="+action.Extra+"&type="+action.ObjectType), result.Info.Name)
					} else {
						action.Extra = result.Info.Name
					}
				}
			}
		case "moved":
			id, _ := strconv.Atoi(action.Extra)
			if id > 0 {
				if project := HostConn.GetProjectById(int32(id)); project != nil {
					if hasPriv(data, "project", "task") {
						action.Extra = html_a(createLink("project", "task", "projectID="+action.Extra), "#"+action.Extra+project.Name)
					} else {
						action.Extra = "#" + action.Extra + project.Name
					}
				}
			}
		case "frombug":
			if hasPriv(data, "bug", "view") {
				action.Extra = html_a(createLink("bug", "view", "bugID="+action.Extra), action.Extra)
			}
		case "unlinkedfromproject":
			id, _ := strconv.Atoi(action.Extra)
			if id > 0 {
				if project := HostConn.GetProjectById(int32(id)); project != nil {
					if hasPriv(data, "project", "story") {
						action.Extra = html_a(createLink("project", "story", "projectID="+action.Extra), "#"+action.Extra+project.Name)
					} else {
						action.Extra = "#" + action.Extra + project.Name
					}
				}
			}
		case "tostory", "linkchildstory", "unlinkchildrenstory", "linkparentstory", "unlinkparentstory", "deletechildrenstory":
			id, _ := strconv.Atoi(action.Extra)
			if id > 0 {
				getStorys := protocol.GET_MSG_PROJECT_product_getStories()
				getStorys.Where = map[string]interface{}{"Id": id}
				var result *protocol.MSG_PROJECT_product_getStories_result
				if err = data.SendMsgWaitResultToDefault(getStorys, &result); err != nil {
					return
				}
				if len(result.List) > 0 {

					if hasPriv(data, "story", "view") {
						action.Extra = html_a(createLink("story", "view", "storyID="+action.Extra), "#"+action.Extra+result.List[0].Title)
					} else {
						action.Extra = "#" + action.Extra + result.List[0].Title
					}
				}
			}
		case "createchildren":
			var ids []int32
			libraries.JsonUnmarshalStr(action.Extra, &ids)
			action.Extra = ""
			if len(ids) > 0 {
				getPairs := protocol.GET_MSG_PROJECT_task_getPairs()
				getPairs.Where = map[string]interface{}{"Id": ids}
				var result *protocol.MSG_PROJECT_task_getPairs_result
				if err = data.SendMsgWaitResultToDefault(getPairs, &result); err != nil {
					return
				}

				if len(result.List) > 0 {
					var strs = make([]string, len(result.List))
					if hasPriv(data, "task", "view") {
						for k, kv := range result.List {
							strs[k] = html_a(createLink("task", "view", "taskID="+kv.Key), "#"+kv.Key+kv.Value)
						}

					} else {
						for k, kv := range result.List {
							strs[k] = "#" + kv.Key + kv.Value
						}
					}
					action.Extra = strings.Join(strs, ",")
				}
			}
		case "createrequirements":
			var ids []int32
			libraries.JsonUnmarshalStr(action.Extra, &ids)
			action.Extra = ""
			if len(ids) > 0 {
				getPairs := protocol.GET_MSG_PROJECT_product_getStoriesMapBySql()
				getPairs.Where = map[string]interface{}{"Id": ids}
				getPairs.Field = "Id,Title"
				var result *protocol.MSG_PROJECT_product_getStoriesMapBySql_result
				if err = data.SendMsgWaitResultToDefault(getPairs, &result); err != nil {
					return
				}

				if len(result.List) > 0 {
					var strs = make([]string, len(result.List))
					if hasPriv(data, "story", "view") {
						for k, v := range result.List {
							strs[k] = html_a(createLink("story", "view", "storyID="+v["Id"]), "#"+v["Id"]+v["Title"])
						}

					} else {
						for k, v := range result.List {
							strs[k] = "#" + v["Id"] + v["Title"]
						}
					}
					action.Extra = strings.Join(strs, ",")
				}
			}
		case "totask", "linkchildtask", "unlinkchildrentask", "linkparenttask", "unlinkparenttask", "deletechildrentask":
			id := action.Extra
			action.Extra = ""
			if id > "0" {
				getPairs := protocol.GET_MSG_PROJECT_task_getPairs()
				getPairs.Where = map[string]interface{}{"Id": id}
				var result *protocol.MSG_PROJECT_task_getPairs_result
				if err = data.SendMsgWaitResultToDefault(getPairs, &result); err != nil {
					return
				}

				if len(result.List) > 0 {
					var strs = make([]string, len(result.List))
					if hasPriv(data, "task", "view") {
						for k, kv := range result.List {
							strs[k] = html_a(createLink("task", "view", "taskID="+kv.Key), "#"+kv.Key+kv.Value)
						}

					} else {
						for k, kv := range result.List {
							strs[k] = "#" + kv.Key + kv.Value
						}
					}
					action.Extra = strings.Join(strs, ",")
				}
			}
		case "buildopened":

			getBuild := protocol.GET_MSG_PROJECT_build_getById()
			getBuild.Id = action.ObjectID

			var result *protocol.MSG_PROJECT_build_getById_result
			if err = data.SendMsgWaitResultToDefault(getBuild, &result); err != nil {
				return
			}
			if result.Info != nil {
				if hasPriv(data, "build", "view") {
					action.Extra = html_a(createLink("build", "view", "builID="+strconv.Itoa(int(action.ObjectID))), "#"+strconv.Itoa(int(action.ObjectID))+" "+result.Info.Name)
				} else {
					action.Extra = "#" + strconv.Itoa(int(action.ObjectID)) + " " + result.Info.Name
				}
			}
		case "testtaskopened", "testtaskstarted", "testtaskclosed":
			getTesttask := protocol.GET_MSG_TEST_testtask_getById()
			getTesttask.Id = action.ObjectID

			var result *protocol.MSG_TEST_testtask_getById_result
			if err = data.SendMsgWaitResultToDefault(getTesttask, &result); err != nil {
				return
			}
			if result.Info != nil {
				if hasPriv(data, "testtask", "view") {
					action.Extra = html_a(createLink("testtask", "view", "testtaskID="+strconv.Itoa(int(action.ObjectID))), "#"+strconv.Itoa(int(action.ObjectID))+" "+result.Info.Name)
				} else {
					action.Extra = "#" + strconv.Itoa(int(action.ObjectID)) + " " + result.Info.Name
				}
			}

		case "fromlib":
			if action.ObjectType == "case" {
				id, _ := strconv.Atoi(action.Extra)
				if id > 0 {
					getTasksuite := protocol.GET_MSG_TEST_testsuite_getById()
					getTasksuite.Id = int32(id)

					var result *protocol.MSG_TEST_testsuite_getById_result
					if err = data.SendMsgWaitResultToDefault(getTasksuite, &result); err != nil {
						return
					}
					if result.Info != nil {
						if hasPriv(data, "caselib", "browse") {
							action.Extra = html_a(createLink("caselib", "browse", "libID="+action.Extra), result.Info.Name)
						} else {
							action.Extra = result.Info.Name
						}
					}
				}
			}
		case "closed":
			if action.ObjectType == "story" {
				extra := strings.Split(action.Extra, ":")
				if len(extra) == 2 {
					id, _ := strconv.Atoi(extra[1])
					if id > 0 {
						getStorys := protocol.GET_MSG_PROJECT_product_getStories()
						getStorys.Where = map[string]interface{}{"Id": id}
						var result *protocol.MSG_PROJECT_product_getStories_result
						if err = data.SendMsgWaitResultToDefault(getStorys, &result); err != nil {
							return
						}
						if len(result.List) > 0 {

							if hasPriv(data, "story", "view") {
								action.AppendLink = html_a(createLink("story", "view", "storyID="+extra[1]), "#"+extra[1]+result.List[0].Title)
							} else {
								action.AppendLink = "#" + extra[1] + result.List[0].Title
							}
						}
					}
					action.Extra = extra[0]

				}
			}
		case "resolved":
			if action.ObjectType == "bug" {
				extra := strings.Split(action.Extra, ":")
				if len(extra) == 2 {
					id, _ := strconv.Atoi(extra[1])
					if id > 0 {
						getBug := protocol.GET_MSG_TEST_bug_getPairs()
						getBug.Where = map[string]interface{}{"Id": id}
						var result *protocol.MSG_TEST_bug_getPairs_result
						if err = data.SendMsgWaitResultToDefault(getBug, &result); err != nil {
							return
						}
						if len(result.List) > 0 {

							if hasPriv(data, "story", "view") {
								action.AppendLink = html_a(createLink("story", "view", "storyID="+extra[1]), "#"+extra[1]+result.List[0].Value)
							} else {
								action.AppendLink = "#" + extra[1] + result.List[0].Value
							}
						}
					}
					action.Extra = extra[0]

				}
			}
		case "opened", "managed":
			if objectType == "project" {
				var ids []int32
				libraries.JsonUnmarshalStr(action.Extra, &ids)
				action.Extra = ""
				if len(ids) > 0 {
					/*var strs []string
					for _, id := range ids {
						if product := HostConn.GetProductById(id); product != nil {
							if hasPriv(data, "product", "browse") {

								strs = append(strs, html_a(createLink("product", "browse", "productID="+strconv.Itoa(int(product.Id)), "#"+strconv.Itoa(int(product.Id)))+" "+product.Name))

							} else {

								strs = append(strs, "#"+strconv.Itoa(int(product.Id))+" "+product.Name)

							}
						}

					}
					action.Extra = fmt.Sprintf(data.Lang["project"]["action"], a) strings.Join(strs, ",")*/
					libraries.ReleaseLog("action记录未处理project opened中extra包含额外信息%+v", action)

				}

			}

		}
		/* }
		elseif($actionName == "finished" and objectType == "todo")
		{
		    action.AppendLink = "";
		    if(strpos(action.Extra, ":")!== false)
		    {
		        list($extra, $id) = explode(":", action.Extra);
		        action.Extra    = strtolower($extra);
		        if($id)
		        {
		            $table = $this->config->objectTables[action.Extra];
		            $field = $this->config->action->objectNameFields[action.Extra];
		            $name  = $this->dao->select($field)->from($table)->where("id")->eq($id)->fetch($field);
		            if($name) action.AppendLink = html::a(helper::createLink(action.Extra, "view", "id=$id"), "#$id " . $name);
		        }
		    }
		}




		if($actionName == "svncommited")
		{
		    foreach(action.History as $history)
		    {
		        if($history->field == "subversion") $history->diff = str_replace("+", "%2B", $history->diff);
		    }
		}
		elseif($actionName == "gitcommited")
		{
		    foreach(action.History as $history)
		    {
		        if($history->field == "git") $history->diff = str_replace("+", "%2B", $history->diff);
		    }
		}*/
	}
	return result.List, nil
}
func action_create(data *TemplateData, objectType string, objectID int32, actionType, comment, extra string, products, projects []int32) {
	//原则上，不应该在http产生actionlog，而是在相应的服务，修改数据库后，session.CommitRollback(func(){进行action_create})
	msg, _ := data.GetMsg()
	msg.ActionCreate(objectType, objectID, actionType, comment, extra, products, projects)
}

const (
	actionGetDynamicParamAll     = 0
	actionGetDynamicParamNotzero = -1
)

//productID与projectID为0时候为all,-1对应notzero
func action_getDynamic(data *TemplateData, account int32, period, orderBy string, pager TempLatePage, productID, projectID int32, ext ...string) (res []*protocol.MSG_LOG_transformActions_info, err error) { //date = "", direction = "next"
	var begin, end string //查询开始结束日期
	period = strings.ToLower(period)
	var condition = make(map[string]interface{})
	if period != "all" {
		now := time.Now()
		switch period {
		case "today":
			begin = now.Format(protocol.TIMEFORMAT_MYSQLDATE)
			end = begin
		case "yesterday":
			begin = now.AddDate(0, 0, -1).Format(protocol.TIMEFORMAT_MYSQLDATE)
			end = begin
		case "twodaysago":
			begin = now.AddDate(0, 0, -2).Format(protocol.TIMEFORMAT_MYSQLDATE)
			end = now.AddDate(0, 0, -1).Format(protocol.TIMEFORMAT_MYSQLDATE)
		case "latest3days":
			begin = now.AddDate(0, 0, -3).Format(protocol.TIMEFORMAT_MYSQLDATE)
			end = now.AddDate(0, 0, -1).Format(protocol.TIMEFORMAT_MYSQLDATE)
		case "thisweek", "lastweek":
			weekBegin := now.AddDate(0, 0, int(time.Monday-now.Weekday()))
			begin = weekBegin.Format(protocol.TIMEFORMAT_MYSQLDATE)
			end = weekBegin.AddDate(0, 0, 6).Format(protocol.TIMEFORMAT_MYSQLDATE)
		case "thismonth":
			monthBegin := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
			begin = monthBegin.Format(protocol.TIMEFORMAT_MYSQLDATE)
			end = monthBegin.AddDate(0, 1, -1).Format(protocol.TIMEFORMAT_MYSQLDATE)
		case "lastmonth":
			monthBegin := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
			begin = monthBegin.AddDate(0, -1, 0).Format(protocol.TIMEFORMAT_MYSQLDATE)
			end = monthBegin.AddDate(0, 0, -1).Format(protocol.TIMEFORMAT_MYSQLDATE)
		}
		condition["Date"] = []interface{}{mysql.WhereOperatorBETWEEN, begin + " 00:00:00", end + " 23:59:59"}
	}
	if len(ext) > 0 && ext[0] != "" { //date
		date := ext[0]
		direction := "<"
		if len(ext) == 2 && ext[1] != "next" {
			direction = ">"
		}
		condition["Date"] = []interface{}{direction, date}
	}
	if account > 0 {
		condition["ActorId"] = account
	}

	if productID > actionGetDynamicParamNotzero {
		if productID == actionGetDynamicParamAll {
			if !data.User.IsAdmin {
				var products []int32
				for id := range data.User.AclProducts {
					products = append(products, id)
				}
				condition["Products"] = []interface{}{mysql.WhereOperatorJSONCONTAINS, products}
			}
		} else {
			condition["Products"] = []interface{}{mysql.WhereOperatorJSONCONTAINS, []int32{productID}}
		}
	} else {
		condition["Products"] = []interface{}{mysql.WhereOperatorRAWNE, "[]"} //products!=[]
	}
	if projectID > actionGetDynamicParamNotzero {
		if projectID == actionGetDynamicParamAll {
			if !data.User.IsAdmin {
				var projects []int32
				for id := range data.User.AclProjects {
					projects = append(projects, id)
				}
				condition["Projects"] = []interface{}{mysql.WhereOperatorJSONCONTAINS, projects}
			}
		} else {
			condition["Projects"] = []interface{}{mysql.WhereOperatorJSONCONTAINS, []int32{projectID}}
		}
	} else {
		condition["Projects"] = []interface{}{mysql.WhereOperatorRAWNE, "[]"} //Projects!=[]
	}

	out := protocol.GET_MSG_LOG_Action_transformActions()
	out.Where = condition
	out.Order = orderBy
	var result *protocol.MSG_LOG_Action_transformActions_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	/*this->loadModel("doc");
	  libs = this->doc->getLibs("all");
	  docs = this->doc->getPrivDocs(array_keys(libs));*/

	//actionCondition = this->getActionCondition();
	//if(is_array(actionCondition)) return array();
	/* Get actions.
	actions = this->dao->select("*")->from(TABLE_ACTION)
	            ->where(1)

	            ->beginIF(docs and !this->app->user->admin)->andWhere("IF(objectType != "doc", "1=1", objectID " . helper::dbIN(docs) . ")")->fi()
	            ->beginIF(libs and !this->app->user->admin)->andWhere("IF(objectType != "doclib", "1=1", objectID " . helper::dbIN(array_keys(libs)) . ") ")->fi()
	            ->beginIF(!empty(actionCondition))->andWhere("(actionCondition)")->fi()
	            ->orderBy(orderBy)
	            ->page(pager)
	            ->fetchAll();

	        if(!actions) return array();

	        this->loadModel("common")->saveQueryCondition(this->dao->get(), "action");
	        return this->transformActions(actions);;*/
	out.Put()
	return result.List, nil
}
