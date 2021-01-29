package handler

import (
	"jachunPM_http/js"
	"protocol"
	"strconv"
)

func init() {

	httpHandlerMap["GET"]["/my/buildContactLists"] = get_my_buildContactLists
	httpHandlerMap["GET"]["/my/managecontacts"] = get_my_managecontacts
	httpHandlerMap["POST"]["/my/managecontacts"] = post_my_managecontacts
}
func get_my_buildContactLists(data *TemplateData) {
	out := protocol.GET_MSG_USER_getContactLists()
	out.Uid = data.User.Id
	out.Params = "withnote"
	var result *protocol.MSG_USER_getContactLists_result
	err := HostConn.SendMsgWaitResultToDefault(out, &result)
	if err != nil {
		data.OutErr(err)
		return
	}
	data.Data["contactLists"] = result.List
	templateOut("my.buildContactLists.html", data)
	out.Put()
	result.Put()
	return
}
func get_my_managecontacts(data *TemplateData) {
	mode := data.ws.Query("mode")
	if mode == "" {
		mode = "edit"
	}
	listID, _ := strconv.Atoi(data.ws.Query("listID"))
	out := protocol.GET_MSG_USER_getContactLists()
	out.Uid = data.User.Id
	var result *protocol.MSG_USER_getContactLists_result
	err := HostConn.SendMsgWaitResultToDefault(out, &result)
	if err != nil {
		data.OutErr(err)
		return
	}
	lists := result.List
	var globalContacts_result *protocol.MSG_USER_getGlobalContacts_result
	err = HostConn.SendMsgWaitResultToDefault(&protocol.MSG_USER_getGlobalContacts{}, &globalContacts_result)
	if err != nil {
		data.OutErr(err)
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
	err = HostConn.SendMsgWaitResultToDefault(get_list, &get_list_result)

	myContacts := get_list_result.List

	if err != nil {
		data.OutErr(err)
		return
	}
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
		err = HostConn.SendMsgWaitResultToDefault(getContactList, &result)
		if err != nil {
			data.OutErr(err)
			return

		}
		data.Data["list"] = result.Result
	}
	data.Data["mode"] = mode
	data.Data["lists"] = lists
	data.Data["listID"] = listID
	data.Data["users"], err = user_getPairs("noletter|noempty|noclosed|noclosed|nodeleted")
	if err != nil {
		data.OutErr(err)
		return
	}
	data.Data["disabled"] = disabled
	data.Data["globalContacts"] = globalContacts
	templateOut("my.managecontacts.html", data)
	out.Put()
	result.Put()
	return
}
func post_my_managecontacts(data *TemplateData) {
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
		err := HostConn.SendMsgWaitResultToDefault(insert, &result)
		if err != nil {
			data.OutErr(err)
			return
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
		err := HostConn.SendMsgWaitResultToDefault(insert, nil)
		if err != nil {
			data.OutErr(err)
			return
		}
		data.ws.WriteString(js.Location(createLink("my", "manageContacts", "listID="+strconv.Itoa(id)), "parent"))

	default:
		js.Reload("parent")
	}

	return
}
