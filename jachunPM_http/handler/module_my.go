package handler

import (
	"html/template"
	"jachunPM_http/config"
	"jachunPM_http/js"
	"libraries"
	"protocol"
	"strconv"
	"time"
)

func init() {

	httpHandlerMap["GET"]["/my/buildContactLists"] = get_my_buildContactLists
	httpHandlerMap["GET"]["/my/managecontacts"] = get_my_managecontacts
	httpHandlerMap["POST"]["/my/managecontacts"] = post_my_managecontacts
	httpHandlerMap["GET"]["/my/index"] = get_my_index
	httpHandlerMap["GET"]["/my/task"] = get_my_task
	httpHandlerMap["GET"]["/my/project"] = get_my_project

}
func get_my_buildContactLists(data *TemplateData) (err error) {
	out := protocol.GET_MSG_USER_getContactLists()
	out.Uid = data.User.Id
	out.Params = "withnote"
	var result *protocol.MSG_USER_getContactLists_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	data.Data["contactLists"] = result.List
	templateOut("my.buildContactLists.html", data)
	out.Put()
	result.Put()
	return
}
func get_my_managecontacts(data *TemplateData) (err error) {
	mode := data.ws.Query("mode")
	if mode == "" {
		mode = "edit"
	}
	listID, _ := strconv.Atoi(data.ws.Query("listID"))
	out := protocol.GET_MSG_USER_getContactLists()
	out.Uid = data.User.Id
	var result *protocol.MSG_USER_getContactLists_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	lists := result.List
	var globalContacts_result *protocol.MSG_USER_getGlobalContacts_result
	if err = data.SendMsgWaitResultToDefault(&protocol.MSG_USER_getGlobalContacts{}, &globalContacts_result); err != nil {
		return
	}
	globalContacts := make(map[int32]bool, len(globalContacts_result.Result))
	for _, c := range globalContacts_result.Result {
		globalContacts[c.Id] = true
	}
	disabled := make(map[int32]bool, len(globalContacts))
	for k, v := range globalContacts {
		disabled[k] = v
	}
	get_list := protocol.GET_MSG_USER_getContactListByUid()
	get_list.Uid = data.User.Id
	var get_list_result *protocol.MSG_USER_getContactListByUid_result
	if err = data.SendMsgWaitResultToDefault(get_list, &get_list_result); err != nil {
		return
	}
	myContacts := get_list_result.List

	if len(myContacts) > 0 && len(globalContacts) > 0 {
		for _, kv := range myContacts {
			id, _ := strconv.Atoi(kv.Key)
			delete(disabled, int32(id))
		}
	}
	if listID == 0 && len(lists) > 0 {
		listID, _ = strconv.Atoi(lists[0].Key)
	}
	if listID == 0 {
		mode = "new"
	}

	if mode == "new" {
		data.Data["title"] = data.Lang["my"]["common"].(string) + data.Lang["common"]["colon"].(string) + data.Lang["user"]["contacts"].(map[string]string)["createList"]
		data.Data["list"] = protocol.GET_MSG_USER_ContactList()
	} else {
		data.Data["title"] = data.Lang["my"]["common"].(string) + data.Lang["common"]["colon"].(string) + data.Lang["user"]["contacts"].(map[string]string)["manage"]
		getContactList := protocol.GET_MSG_USER_getContactListById()
		getContactList.Id = int32(listID)
		var result *protocol.MSG_USER_getContactListById_result
		if err = data.SendMsgWaitResultToDefault(getContactList, &result); err != nil {
			return

		}
		data.Data["list"] = result.Result
	}
	data.Data["mode"] = mode
	data.Data["lists"] = lists
	data.Data["listID"] = listID
	if data.Data["users"], err = user_getPairs(data, "noletter|noempty|noclosed|noclosed|nodeleted"); err != nil {
		return
	}
	data.Data["disabled"] = disabled
	data.Data["globalContacts"] = globalContacts
	templateOut("my.managecontacts.html", data)
	out.Put()
	result.Put()
	return
}
func post_my_managecontacts(data *TemplateData) (err error) {
	mode := data.ws.Post("mode")
	ListName := ""
	var users []int32
	var share bool
	var id int
	for k, v := range data.ws.GetAllPost() {
		switch k {
		case "newList", "listName":
			ListName = v[0]
		case "users":
			for _, sid := range v {
				id, _ := strconv.Atoi(sid)
				if id > 0 {
					users = append(users, int32(id))
				}
			}
		case "share":
			share = v[0] == "1"
		case "listID":
			id, _ = strconv.Atoi(v[0])
		}

	}
	if ListName == "" {
		data.ws.WriteString(js.Alert(data.Lang["my"]["error"].(map[string]string)["managecontactsEmptyListName"]))
		return
	}
	if len(users) == 0 {
		data.ws.WriteString(js.Alert(data.Lang["my"]["error"].(map[string]string)["managecontactsEmptyUsers"]))
		return
	}
	switch mode {
	case "new":

		insert := protocol.GET_MSG_USER_insertUpdateContactList()
		insert.Insert = protocol.GET_MSG_USER_ContactList()
		insert.Insert.ListName = ListName
		insert.Insert.UserList = users
		insert.Insert.Uid = data.User.Id
		insert.Insert.Share = share
		var result *protocol.MSG_USER_insertUpdateContactList_result
		if err := data.SendMsgWaitResultToDefault(insert, &result); err != nil {
			data.ws.WriteString(js.Alert(err.Error()))
			return nil
		}
		if data.onlybody() {
			data.ws.WriteString(js.CloseModal("parent.parent", "", "function(){parent.parent.ajaxGetContacts('#mailto')}"))
			return
		}
		data.ws.WriteString(js.Location(createLink("my", "manageContacts", "listID="+strconv.Itoa(int(result.Id))), "parent"))

	case "edit":
		if id <= 0 {
			data.ws.WriteString(js.Alert(data.Lang["my"]["error"].(map[string]string)["managecontactsErrorID"]))
			return
		}
		insert := protocol.GET_MSG_USER_insertUpdateContactList()
		insert.Insert = protocol.GET_MSG_USER_ContactList()
		insert.Insert = protocol.GET_MSG_USER_ContactList()
		insert.Insert.ListName = ListName
		insert.Insert.UserList = users
		insert.Insert.Uid = data.User.Id
		insert.Insert.Share = share
		insert.Insert.Id = int32(id)
		if err := data.SendMsgWaitResultToDefault(insert, nil); err != nil {
			data.ws.WriteString(js.Alert(err.Error()))
			return nil
		}
		data.ws.WriteString(js.Location(createLink("my", "manageContacts", "listID="+strconv.Itoa(id)), "parent"))

	default:
		js.Reload("parent")
	}

	return
}
func get_my_index(data *TemplateData) (err error) {
	data.Data["title"] = data.Lang["my"]["common"]
	templateOut("my.index.html", data)
	return nil
}
func get_my_task(data *TemplateData) (err error) {
	typ := data.ws.Query("type")
	if typ == "" {
		typ = "assignedtome"
	}
	orderBy := data.ws.Query("orderBy")
	orderBy = "id_desc"

	data.Data["title"] = data.Lang["my"]["common"].(string) + data.Lang["common"]["colon"].(string) + data.Lang["my"]["task"].(string)
	data.Data["tabID"] = "task"
	//走MSG_PROJECT_project_getProjectTasks
	out := protocol.GET_MSG_PROJECT_project_getProjectTasks()
	out.Type = []string{typ}
	out.Total = data.Page.Total
	out.Page = data.Page.Page
	out.PerPage = data.Page.PerPage
	out.OrderBy = orderBy
	var result *protocol.MSG_PROJECT_project_getProjectTasks_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	data.Page.Total = result.Total
	data.Data["type"] = typ

	data.Data["orderBy"] = orderBy
	if data.Data["users"], err = user_getPairs(data, "noletter"); err != nil {
		return
	}
	data.Data["customFields"] = datatable_getSetting(data, "my", "")
	data.Data["vars"] = "type=" + typ + "&orderBy=%s&recTotal=" + strconv.Itoa(data.Page.Total) + "&recPerPage=" + strconv.Itoa(data.Page.PerPage) + "&pageID=" + strconv.Itoa(data.Page.Page)
	//分段渲染
	users, err := user_getPairs(data, "noletter")
	if err != nil {
		return
	}

	cellMode := "datatable"
	if config, ok := data.Config["datatable"]["projectTask"]; ok {
		if mode, ok := config["mode"].(string); ok && mode == "datatable" {
			data.Data["useDatatable"] = true
		} else {
			data.Data["useDatatable"] = false
			cellMode = "table"
		}
	} else {
		data.Data["useDatatable"] = false
		cellMode = "table"
	}
	customFields := datatable_getSetting(data, "project", "task")
	data.Data["customFields"] = customFields
	buf := bufpool.Get().(*libraries.MsgBuffer)

	for _, task := range result.List {

		buf.WriteString("<tr data-id='")
		buf.WriteString(strconv.Itoa(int(task.Id)))
		buf.WriteString("' data-status='")
		buf.WriteString(task.Status)
		buf.WriteString("' data-estimate='")
		buf.WriteString(strconv.Itoa(int(task.Estimate)))
		buf.WriteString("' data-consumed='")
		buf.WriteString(strconv.Itoa(int(task.Consumed)))
		buf.WriteString("' data-left='")
		buf.WriteString(strconv.Itoa(int(task.Left)))
		buf.WriteString("'>")
		for _, field := range data.Data["customFields"].([]*config.ConfigDatatable) {
			buf.WriteString(task_printCell(data, field, task, users, typ, nil, nil, cellMode, false, 0))
		}
		buf.WriteString("</tr>")
		if len(task.Children) > 0 {
			for i, child := range task.Children {
				n++
				buf.WriteString("<tr class='table-children")

				if i == 0 {
					buf.WriteString(" table-child-top")
				}
				if i == len(task.Children)-1 {
					buf.WriteString(" table-child-bottom")
				}

				buf.WriteString("parent-")
				buf.WriteString(strconv.Itoa(int(task.Id)))
				buf.WriteString("' data-id='")
				buf.WriteString(strconv.Itoa(int(child.Id)))
				buf.WriteString("' data-status='")
				buf.WriteString(child.Status)
				buf.WriteString("' data-estimate='")
				buf.WriteString(strconv.Itoa(int(child.Estimate)))
				buf.WriteString("' data-consumed='")
				buf.WriteString(strconv.Itoa(int(child.Consumed)))
				buf.WriteString("' data-left='")
				buf.WriteString(strconv.Itoa(int(child.Left)))
				buf.WriteString("'>\r\n")
				var end_flag1, endflag2 int
				for _, field := range data.Data["customFields"].([]*config.ConfigDatatable) {
					end_flag1 = 0
					if i == len(task.Children)-1 {
						end_flag1 = 1
					}
					buf.WriteString(task_printCell(data, field, child, users, typ, nil, nil, cellMode, true, end_flag1))
				}
				buf.WriteString("</tr>\r\n")
				if len(child.Grandchildren) > 0 {
					for k, grandchild := range child.Grandchildren {
						n++
						buf.WriteString("<tr class='table-children")
						if k == 0 {
							buf.WriteString(" table-child-top")
						}
						if k == len(child.Grandchildren)-1 {
							buf.WriteString(" table-child-bottom")
						}

						buf.WriteString("parent-")
						buf.WriteString(strconv.Itoa(int(child.Id)))
						buf.WriteString("' data-id='")
						buf.WriteString(strconv.Itoa(int(grandchild.Id)))
						buf.WriteString("' data-status='")
						buf.WriteString(grandchild.Status)
						buf.WriteString("' data-estimate='")
						buf.WriteString(strconv.Itoa(int(grandchild.Estimate)))
						buf.WriteString("' data-consumed='")
						buf.WriteString(strconv.Itoa(int(grandchild.Consumed)))
						buf.WriteString("' data-left='")
						buf.WriteString(strconv.Itoa(int(grandchild.Left)))
						buf.WriteString("'>\r\n")
						for _, field := range data.Data["customFields"].([]*config.ConfigDatatable) {
							endflag2 = 0
							if k == len(child.Grandchildren)-1 {
								endflag2 = 2
							}
							buf.WriteString(task_printCell(data, field, grandchild, users, typ, nil, nil, cellMode, true, end_flag1|endflag2))
						}
						buf.WriteString("</tr>\r\n")

					}
				}

			}
		}

	}
	data.Data["taskCell"] = template.HTML(buf.String())
	templateOut("my.task.html", data)
	result.Put()
	out.Put()
	return
}
func get_my_project(data *TemplateData) (err error) {
	data.Data["title"] = data.Lang["my"]["common"].(string) + data.Lang["common"]["colon"].(string) + data.Lang["my"]["myProject"].(string)
	data.Data["tabID"] = "project"

	var projects = make(map[int32]*protocol.MSG_PROJECT_project_cache)
	var teams []*protocol.MSG_USER_team_info
	if allproject, err := project_getAll(data); err != nil {
		return err
	} else {
		for _, project := range allproject {
			for _, team := range project.Teams {
				if team.Uid == data.User.Id {
					teams = append(teams, team)
					if project.Status != "done" && project.Status != "closed" && project.Status != "suspended" {
						if time.Now().After(project.End) {
							project.Delay = int64(time.Now().Sub(project.End) / time.Second / 86400)
						}
					}
					projects[project.Id] = project
					break
				}
			}
		}
	}
	data.Data["teams"] = teams
	data.Data["projects"] = projects
	templateOut("my.project.html", data)
	return
}
