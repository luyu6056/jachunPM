package handler

import (
	"libraries"
	"protocol"
	"reflect"
	"strconv"
	"strings"
)

func init() {
	httpHandlerMap["GET"]["/group/browse"] = get_group_browse
	httpHandlerMap["GET"]["/group/manageView"] = get_group_manageView
	httpHandlerMap["POST"]["/group/manageView"] = post_group_manageView
	httpHandlerMap["GET"]["/group/managePriv"] = get_group_managePriv
}
func groupTemplateFuncs() {
	global_Funcs["MSG_USER_Group_cache_isClickable"] = func(data *TemplateData, obj interface{}, action string) bool {
		if group, ok := obj.(*protocol.MSG_USER_Group_cache); ok {
			switch action {
			case "manageview", "manageView":
				if group.Role == "limited" {
					return false
				}
			case "copy":
				if group.Role == "limited" {
					return false
				}
			}
		} else {
			libraries.DebugLog("MSG_USER_Group_cache_isClickable传入的值类型%v不对", reflect.TypeOf(obj).Elem().Name())
		}
		return true
	}
}
func get_group_browse(data *TemplateData) (err error) {
	companyID, _ := strconv.Atoi(data.ws.Query("companyID"))

	if companyID == 0 {
		companyID = int(data.App["company"].(protocol.MSG_USER_Company_cache).Id)
	}

	groups, err := group_getList(data, int32(companyID))
	if err != nil {
		return err
	}
	var groupUsers = make(map[int32][]protocol.HtmlKeyValueStr)
	for _, group := range groups {
		groupUsers[group.Id], err = group_getUserPairs(data, group.Id)
		if err != nil {
			return
		}
	}

	data.Data["title"] = data.Lang["company"]["orgView"].(string) + data.Lang["common"]["colon"].(string) + data.Lang["group"]["browse"].(string)
	data.Data["groups"] = groups
	data.Data["groupUsers"] = groupUsers
	data.Lang["group"]["managepriv"] = data.Lang["group"]["managePrivByGroup"]
	templateOut("group.browse.html", data)
	return nil
}
func group_getAllList(data *TemplateData) (result []*protocol.MSG_USER_Group_cache, err error) {
	if data.Data["group_getAllList"] == nil {
		res, err := HostConn.CacheGetPath(protocol.UserServerNo, protocol.PATH_USER_GROUP_CACHE)
		if err != nil {
			return nil, err
		}

		buf := bufpool.Get().(*libraries.MsgBuffer)
		for _, b := range res {
			buf.Reset()
			buf.Write(b)
			if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_USER_Group_cache); ok {
				result = append(result, v)
			}
		}
		buf.Reset()
		bufpool.Put(buf)
		protocol.Order_group(result, nil)
		data.Data["group_getAllList"] = result
	}
	return data.Data["group_getAllList"].([]*protocol.MSG_USER_Group_cache), nil
}
func group_getList(data *TemplateData, companyID int32) (result []*protocol.MSG_USER_Group_cache, err error) {
	return group_getAllList(data)
}
func group_getById(data *TemplateData, groupID int32) (result *protocol.MSG_USER_Group_cache, err error) {
	list, err := group_getAllList(data)
	if err != nil {
		return nil, err
	}
	for _, group := range list {
		if group.Id == groupID {
			return group, nil
		}
	}
	return nil, nil
}
func group_getUserPairs(data *TemplateData, groupID int32) (res []protocol.HtmlKeyValueStr, err error) {
	users, err := user_getAllcache(data)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		for _, id := range user.Group {
			if id == groupID && !user.Deleted {
				name := user.Realname
				if name == "" {
					name = user.Account
				}
				res = append(res, protocol.HtmlKeyValueStr{strconv.Itoa(int(user.Id)), name})
			}
		}
	}
	return
}

func get_group_manageView(data *TemplateData) (err error) {
	groupID, _ := strconv.Atoi(data.ws.Query("groupID"))
	group, err := group_getById(data, int32(groupID))
	if err != nil {
		return
	}
	data.Data["title"] = data.Lang["company"]["common"].(string) + data.Lang["common"]["colon"].(string) + group.Name + data.Lang["common"]["colon"].(string) + data.Lang["group"]["manageView"].(string)
	data.Data["group"] = group
	if products, err := product_getAll(data); err != nil {
		return err
	} else {
		data.Data["products"] = append([]*protocol.MSG_PROJECT_product_cache{&protocol.MSG_PROJECT_product_cache{Id: -1, Name: "all"}}, products...)
	}
	if projects, err := project_getAll(data); err != nil {
		return err
	} else {
		data.Data["projects"] = append([]*protocol.MSG_PROJECT_project_cache{&protocol.MSG_PROJECT_project_cache{Id: -1, Name: "all"}}, projects...)
	}
	var menus []protocol.HtmlKeyValueInterface
	for _, menu := range data.Lang["common"]["menu"].([]protocol.HtmlMenu) {
		if menu.Key == "my" {
			continue
		}
		name := ""
		if v, ok := menu.Value["link"]; ok {
			if i := strings.Index(v, "|"); i > -1 {
				name = v[:i]
			}
		}
		if len(group.Acl) == 0 {
			menus = append(menus, protocol.HtmlKeyValueInterface{menu.Key, protocol.HtmlKeyValueStr{name, "checked"}})
		} else {
			checked := ""
			for _, key := range group.Acl {
				if key == menu.Key {
					checked = "checked"
					break
				}
			}
			menus = append(menus, protocol.HtmlKeyValueInterface{menu.Key, protocol.HtmlKeyValueStr{name, checked}})
		}
	}
	data.Data["menus"] = menus
	templateOut("group.manageView.html", data)
	return
}
func post_group_manageView(data *TemplateData) (err error) {
	groupID, _ := strconv.Atoi(data.ws.Query("groupID"))
	group, err := group_getById(data, int32(groupID))
	if err != nil {
		return
	}
	group.Acl = group.Acl[:0]
	group.AclProjects = group.AclProjects[:0]
	group.AclProducts = group.AclProducts[:0]
	if data.ws.Post("allchecker") == "" {
		for _, s := range data.ws.PostSlice("Acl") {
			group.Acl = append(group.Acl, s)
		}
	}
	for _, s := range data.ws.PostSlice("AclProducts") {
		id, err := strconv.Atoi(s)
		if err == nil {
			group.AclProducts = append(group.AclProducts, int32(id))
		}
	}
	for _, s := range data.ws.PostSlice("AclProjects") {
		id, err := strconv.Atoi(s)
		if err == nil {
			group.AclProjects = append(group.AclProjects, int32(id))
		}
	}
	out := protocol.GET_MSG_USER_group_update()
	out.Update = group
	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		data.ajaxResult(false, err.Error())
		return nil
	}
	data.ajaxResult(true, data.Lang["common"]["saveSuccess"], createLink("group", "browse", ""))
	return nil
}
func get_group_managePriv(data *TemplateData) (err error) {
	typ := data.ws.Query("type")
	groupID, _ := strconv.Atoi(data.ws.Query("param"))
	menu := data.ws.Query("menu")
	data.Data["type"] = typ

	if typ == "byGroup" {

		group, err := group_getById(data, int32(groupID))
		if err != nil {
			return err
		}
		var commonMenu []protocol.HtmlKeyValueStr
		for _, menu := range data.Lang["common"]["menu"].([]protocol.HtmlMenu) {
			name := ""
			if v, ok := menu.Value["link"]; ok {
				if i := strings.Index(v, "|"); i > -1 {
					name = v[:i]
				}
			}
			commonMenu = append(commonMenu, protocol.HtmlKeyValueStr{menu.Key, name})
		}
		var resource []protocol.HtmlKeyValueInterface
		for _, moduleName := range data.Lang["common"]["moduleOrder"].([]string) {
			if list, ok := data.Lang["resource"][moduleName].([]protocol.HtmlKeyValueStr); ok {
				if !group_checkModule(data, menu, moduleName) {
					continue
				}
				var kvi []protocol.HtmlKeyValueInterface
				for _, kv := range list {
					tmp := make(map[string]string)
					if v, ok := data.Lang[moduleName]; ok {
						tmp["name"], _ = v[kv.Value].(string)
						tmp["pri"] = ""
						if group.Priv[moduleName][kv.Key] {
							tmp["pri"] = kv.Key
						}
					}
					kvi = append(kvi, protocol.HtmlKeyValueInterface{kv.Key, tmp})
				}
				resource = append(resource, protocol.HtmlKeyValueInterface{moduleName, kvi})
			}
		}
		data.Data["resource"] = resource
		data.Data["commonMenu"] = commonMenu
		data.Data["title"] = data.Lang["company"]["common"].(string) + data.Lang["common"]["colon"].(string) + group.Name + data.Lang["common"]["colon"].(string) + data.Lang["group"]["managePriv"].(string)

		/* Join changelog when be equal or greater than this version.*/
		data.Data["group"] = group
		data.Data["groupID"] = groupID
		data.Data["menu"] = menu
		data.Data["params"] = "type=byGroup&param=" + data.ws.Query("param") + "&menu=%s"

	} else if typ == "byModule" {

		data.Data["title"] = data.Lang["company"]["common"].(string) + data.Lang["common"]["colon"].(string) + data.Lang["group"]["managePriv"].(string)

		/*foreach(this->lang->resource as module => moduleActions)
		  {
		      modules[module] = this->lang->module->common;
		      if(module == "caselib") module = "testsuite";
		      foreach(moduleActions as action)
		      {
		          actions[module][action] = this->lang->module->action;
		      }
		  }
		  data.Data["groups"]  = this->group->getPairs();
		  data.Data["modules"] = modules;
		  data.Data["actions"] = actions;*/
	}
	templateOut("group.managePriv.html", data)
	return nil
}
func group_checkModule(data *TemplateData, menu, moduleName string) bool {
	if menu != "" {
		if menu == "other" {
			if data.Lang["menugroup"][moduleName] != nil || data.Lang[moduleName]["menu"] != nil {
				return false
			}
		} else if data.Lang["menugroup"][moduleName] == nil || (menu != moduleName && data.Lang["menugroup"][moduleName] != nil && data.Lang["menugroup"][moduleName].(string) != menu) {
			return false

		}
	}
	return true
}
