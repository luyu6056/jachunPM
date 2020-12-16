package handler

import (
	"html/template"
	"jachunPM_http/js"
	"libraries"
	"protocol"
	"reflect"
	"strconv"

	"github.com/luyu6056/gnet"
)

func init() {
	httpHandlerMap["GET"]["/dept/browse"] = get_dept_browse
	httpHandlerMap["POST"]["/dept/updateOrder"] = post_dept_updateOrder
	httpHandlerMap["POST"]["/dept/manageChild"] = post_dept_manageChild
	httpHandlerMap["GET"]["/dept/delete"] = get_dept_delete
	httpHandlerMap["GET"]["/dept/edit"] = get_dept_edit
	httpHandlerMap["POST"]["/dept/edit"] = post_dept_edit
}
func get_dept_browse(data *TemplateData) gnet.Action {
	ws := data.ws
	deptID, _ := strconv.Atoi(ws.Query("deptID"))
	data.Data["deptID"] = deptID
	msg, err := HostConn.GetMsg()
	if err != nil {
		ws.OutErr(err)
		return gnet.None
	}
	if int32(deptID) > 0 {
		getParents := protocol.GET_MSG_USER_Dept_getParents()
		getParents.Id = int32(deptID)
		res, err := msg.SendMsgWaitResult(0, getParents)
		if r, ok := res.(*protocol.MSG_USER_Dept_getParents_result); ok {
			data.Data["parentDepts"] = r.List
		} else {
			libraries.DebugLog("GET_MSG_USER_Dept_getParents返回结果不对%s", reflect.TypeOf(res).Elem().String())
		}
		getParents.Put()
		if err != nil {
			ws.OutErr(err)
			return gnet.None
		}
	}
	data.Data["depts"], err = dept_getTreeMenu(data, 0)
	if err != nil {
		ws.OutErr(err)
		return gnet.None
	}
	data.Data["sons"], err = dept_getSons(int32(deptID))
	if err != nil {
		ws.OutErr(err)
		return gnet.None
	}
	getDataStructure := protocol.GET_MSG_USER_Dept_getDataStructure()
	getDataStructure.RootDeptID = 0
	res, err := msg.SendMsgWaitResult(0, getDataStructure)
	if r, ok := res.(*protocol.MSG_USER_Dept_getDataStructure_result); ok {
		data.Data["tree"] = r.List
	} else {
		libraries.DebugLog("GET_MSG_USER_Dept_getDataStructure返回结果不对%s", reflect.TypeOf(res).Elem().String())
	}
	getDataStructure.Put()
	if err != nil {
		ws.OutErr(err)
		return gnet.None
	}
	templateOut("dept.browse.html", data, ws)
	return gnet.None
}
func dept_getTreeMenu(data *TemplateData, rootDeptId int32) (template.HTML, error) {
	var deptList []*protocol.MSG_USER_Dept_cache

	var deptInfo protocol.MSG_USER_Dept_cache
	res, err := HostConn.CacheGetPath(protocol.UserServerNo, protocol.PATH_USER_DEPT_CACHE)
	if err != nil {
		return template.HTML(""), err
	}
	buf := bufpool.Get().(*libraries.MsgBuffer)
	for _, b := range res {
		buf.Reset()
		buf.Write(b)
		if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_USER_Dept_cache); ok {
			deptList = append(deptList, v)
		}
	}

	if len(deptList) > 0 {
		protocol.Order_dept(deptList, func(a, b *protocol.MSG_USER_Dept_cache) bool {
			if a.Grade == b.Grade {
				return a.Order < b.Order
			}
			return a.Grade > b.Grade
		})
	}
	if rootDeptId > 0 {
		for i := len(deptList) - 1; i >= 0; i-- {
			has := false
		out:
			for _, a := range deptList[i].Path {
				for _, b := range deptInfo.Path {
					if a == b {
						has = true
						break out
					}
				}
			}
			if !has {
				deptList = append(deptList[:i], deptList[i+1:]...)
			}
		}
	}
	var deptMenu = make(map[int32]*libraries.MsgBuffer)
	for _, dept := range deptList {
		if _buf, ok := deptMenu[dept.Id]; ok {
			buf, ok := deptMenu[dept.Parent]
			if !ok {
				buf = bufpool.Get().(*libraries.MsgBuffer)
				deptMenu[dept.Parent] = buf
			}
			buf.WriteString("<li>")
			dept_createManageLink(data, dept, buf)
			buf.WriteString("<ul>")
			buf.Write(_buf.Bytes())
			buf.WriteString("</ul>\n")
		} else {
			buf, ok := deptMenu[dept.Parent]
			if !ok {
				buf = bufpool.Get().(*libraries.MsgBuffer)
				deptMenu[dept.Parent] = buf
			}
			buf.WriteString("<li>")
			dept_createManageLink(data, dept, buf)
			buf.WriteString("\n")
		}
		deptMenu[dept.Parent].WriteString("</li>\n")
	}
	buf.Reset()
	buf.WriteString("<ul class='tree' data-ride='tree' data-name='tree-dept'>")
	for _, dept := range deptList {
		if _buf, ok := deptMenu[dept.Id]; ok {
			buf.Write(_buf.Bytes())
		}
		dept.Put()
	}
	buf.WriteString("</ul>\n")
	result := buf.String()
	buf.Reset()
	bufpool.Put(buf)
	for _, buf := range deptMenu {
		buf.Reset()
		bufpool.Put(buf)
	}
	return template.HTML(result), err
}
func dept_getSons(deptId int32) (deptList []*protocol.MSG_USER_Dept_cache, err error) {
	res, err := HostConn.CacheGetPath(protocol.UserServerNo, protocol.PATH_USER_DEPT_CACHE)
	if err != nil {
		return nil, err
	}
	buf := bufpool.Get().(*libraries.MsgBuffer)
	for _, b := range res {
		buf.Reset()
		buf.Write(b)
		if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_USER_Dept_cache); ok && v.Parent == deptId {
			deptList = append(deptList, v)
		}
	}
	buf.Reset()
	bufpool.Put(buf)
	if len(deptList) > 0 {
		protocol.Order_dept(deptList, func(a, b *protocol.MSG_USER_Dept_cache) bool { return a.Order < b.Order })
	}
	return
}
func dept_createManageLink(data *TemplateData, dept *protocol.MSG_USER_Dept_cache, buf *libraries.MsgBuffer) {
	buf.WriteString(dept.Name)
	if hasPriv(data.ws, "dept", "edit") {
		buf.WriteString(" ")
		buf.WriteString(html_a(createLink("dept", "edit", "deptid="+strconv.Itoa(int(dept.Id))), data.Lang["common"]["edit"].(string), "", "data-toggle='modal' data-type='ajax'"))
	}
	if hasPriv(data.ws, "dept", "browse") {
		buf.WriteString(" ")
		buf.WriteString(html_a(createLink("dept", "browse", "deptid="+strconv.Itoa(int(dept.Id))), data.Lang["dept"]["manageChild"].(string)))
	}
	if hasPriv(data.ws, "dept", "delete") {
		buf.WriteString(" ")
		buf.WriteString(html_a(createLink("dept", "delete", "deptid="+strconv.Itoa(int(dept.Id))), data.Lang["common"]["delete"].(string), "hiddenwin"))
	}
	if hasPriv(data.ws, "dept", "updateOrder") {
		buf.WriteString(" ")
		buf.WriteString(html_input("orders["+strconv.Itoa(int(dept.Id))+"]", strconv.Itoa(int(dept.Order)), "style='width:30px;text-align:center'"))
	}
	return
}

func post_dept_updateOrder(data *TemplateData) gnet.Action {
	ws := data.ws
	post := ws.GetAllPost()

	res, err := HostConn.CacheGetPath(protocol.UserServerNo, protocol.PATH_USER_DEPT_CACHE)
	if err != nil {
		ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrDeptInfo"], err))
		return gnet.None
	}
	var m = make(map[int32]*protocol.MSG_USER_Dept_cache)
	buf := bufpool.Get().(*libraries.MsgBuffer)
	for _, b := range res {
		buf.Reset()
		buf.Write(b)
		if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_USER_Dept_cache); ok {
			m[v.Id] = v
		}
	}
	buf.Reset()
	bufpool.Put(buf)
	update := protocol.GET_MSG_USER_Dept_update()
	for deptId, order := range post {
		id, err := strconv.Atoi(deptId)
		if err != nil {
			ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrDeptIDType"], deptId) + js.Reload("parent"))
			return gnet.None
		}
		o, err := strconv.Atoi(order[0])
		if err != nil {
			ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrOrderType"], deptId, order) + js.Reload("parent"))
			return gnet.None
		}

		if deptinfo, ok := m[int32(id)]; ok {
			if deptinfo.Order != int8(o) {
				deptinfo.Order = int8(o)
				update.List = append(update.List, deptinfo)
			}

		} else {
			ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrDeptInfoDeptID"], deptId) + js.Reload("parent"))
			return gnet.None
		}
	}

	_, err = HostConn.SendMsgWaitResultToDefault(update)
	update.Put()
	if err != nil {
		ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrUpdate"], err) + js.Reload("parent"))
		return gnet.None
	}
	ws.WriteString(js.Reload("parent"))
	return gnet.None
}
func post_dept_manageChild(data *TemplateData) gnet.Action {
	ws := data.ws
	res, err := HostConn.CacheGetPath(protocol.UserServerNo, protocol.PATH_USER_DEPT_CACHE)
	if err != nil {
		ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrDeptInfo"], err))
		return gnet.None
	}
	var m = make(map[int32]*protocol.MSG_USER_Dept_cache)
	buf := bufpool.Get().(*libraries.MsgBuffer)
	maxorder := int8(0)
	maxid := int32(0)
	for _, b := range res {
		buf.Reset()
		buf.Write(b)
		if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_USER_Dept_cache); ok {
			m[v.Id] = v
			if v.Order > maxorder {
				maxorder = v.Order
			}
			if v.Id > maxid {
				maxid = v.Id
			}
		}
	}
	buf.Reset()
	bufpool.Put(buf)
	grade := int8(1)
	var parentPath []int32
	parentDeptID, _ := strconv.Atoi(ws.Post("parentDeptID"))
	parentDept, ok := m[int32(parentDeptID)]
	if ok {
		grade = parentDept.Grade + 1
		parentPath = parentDept.Path
	}

	update := protocol.GET_MSG_USER_Dept_update()
	for str_id, v := range ws.GetAllPost() {
		if str_id == "0" {
			for _, name := range v {
				if name != "" {

					insert := protocol.GET_MSG_USER_Dept_cache()
					insert.Name = name
					maxorder++
					insert.Order = maxorder
					maxid++
					insert.Id = maxid
					if parentDept != nil {
						insert.Parent = parentDept.Id
						insert.Grade = grade
						insert.Path = append(parentPath, insert.Id)
					} else {
						insert.Grade = 1
						insert.Path = []int32{insert.Id}
					}

					update.List = append(update.List, insert)
				}
			}

		} else {
			id, _ := strconv.Atoi(str_id)
			if deptinfo, ok := m[int32(id)]; !ok {
				//ws.WriteString(js.Alert(fmt.Sprintf(data.Lang["dept"]["error"].(map[string]string)["ErrDeptInfoDeptID"], str_id)) + js.Reload("parent"))
				//return gnet.None
			} else {
				deptinfo.Name = v[0]
				update.List = append(update.List, deptinfo)
			}
		}
	}

	_, err = HostConn.SendMsgWaitResultToDefault(update)
	update.Put()
	if err != nil {
		ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrUpdate"], err) + js.Reload("parent"))
		return gnet.None
	}
	ws.WriteString(js.Reload("parent"))
	return gnet.None
}
func get_dept_delete(data *TemplateData) gnet.Action {
	ws := data.ws
	deptid, _ := strconv.Atoi(ws.Query("deptid"))
	if ws.Query("confirm") != "yes" {
		ws.WriteString(js.Confirm(data.Lang["dept"]["confirmDelete"].(string), createLink(`dept`, `delete`, "deptid="+ws.Query("deptid")+"&confirm=yes"), ""))
		return gnet.None
	}
	out := protocol.GET_MSG_USER_Dept_delete()
	out.DeptId = int32(deptid)
	res, err := HostConn.SendMsgWaitResultToDefault(out)
	out.Put()
	if err != nil {
		ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrUpdate"], err))
		return gnet.None
	}
	if v, ok := res.(*protocol.MSG_USER_Dept_delete_result); ok {
		if v.Result != protocol.Success {
			ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)[v.Result.String()]) + js.Reload("parent"))
			return gnet.None
		}
	}
	ws.WriteString(js.Reload("parent"))
	return gnet.None
}
func get_dept_edit(data *TemplateData) gnet.Action {
	var deptinfo *protocol.MSG_USER_Dept_cache
	deptid, _ := strconv.Atoi(data.ws.Query("deptid"))
	err := HostConn.CacheGet(protocol.UserServerNo, protocol.PATH_USER_DEPT_CACHE, data.ws.Query("deptid"), &deptinfo)
	if err != nil || deptinfo == nil {
		data.ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrDeptInfoDeptID"], deptid) + js.Reload("parent"))
		return gnet.None
	}
	data.Data["dept"] = deptinfo
	out := protocol.GET_MSG_USER_getDeptUserPairs()
	out.DeptId = int32(deptid)
	res, err := HostConn.SendMsgWaitResultToDefault(out)
	out.Put()
	if err != nil {
		data.ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrGetDeptUser"], err) + js.Reload("parent"))
		return gnet.None
	}

	if v, ok := res.(*protocol.MSG_USER_getDeptUserPairs_result); ok {
		var users = []protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"0", ""}}
		for _, v := range v.List {
			users = append(users, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Realname})
		}
		data.Data["users"] = users
	}
	templateOut("dept.edit.html", data, data.ws)
	return gnet.None
}
func post_dept_edit(data *TemplateData) gnet.Action {
	var deptinfo *protocol.MSG_USER_Dept_cache
	deptid, _ := strconv.Atoi(data.ws.Query("deptid"))
	err := HostConn.CacheGet(protocol.UserServerNo, protocol.PATH_USER_DEPT_CACHE, data.ws.Query("deptid"), &deptinfo)
	if err != nil || deptinfo == nil {
		data.ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrDeptInfoDeptID"], deptid) + js.Reload("parent"))
		return gnet.None
	}
	manager, _ := strconv.Atoi(data.ws.Post("manager"))
	if manager != 0 {
		managerUser := getUserCacheById(int32(manager))
		if managerUser == nil {
			data.ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrManager"]) + js.Reload("parent"))
			return gnet.None
		}
		deptinfo.Manager = managerUser.Id
		deptinfo.ManagerName = managerUser.Realname
		if deptinfo.ManagerName == "" {
			deptinfo.ManagerName = managerUser.Account
		}
	} else {
		deptinfo.Manager = 0
		deptinfo.ManagerName = ""
	}

	deptinfo.Name = data.ws.Post("name")
	update := protocol.GET_MSG_USER_Dept_update()
	update.List = []*protocol.MSG_USER_Dept_cache{deptinfo}
	_, err = HostConn.SendMsgWaitResultToDefault(update)
	update.Put()
	if err != nil {
		data.ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrUpdate"], err) + js.Reload("parent"))
		return gnet.None
	}
	data.ws.WriteString(js.Alert(data.Lang["dept"]["successSave"].(string)) + js.Reload("parent"))
	return gnet.None
}
