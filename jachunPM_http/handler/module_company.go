package handler

import (
	"libraries"
	"protocol"
	"strconv"

	"github.com/luyu6056/gnet"
)

func init() {
	httpHandlerMap["GET"]["/company/browse"] = get_company_browse
}
func get_company_browse(data *TemplateData) gnet.Action {
	ws := data.ws
	param := ws.Query("param")
	TYPE := ws.Query("type")

	deptID := 0
	if TYPE == "" {
		TYPE = "bydept"
	}
	if TYPE == "bydept" {
		deptID, _ = strconv.Atoi(ws.Query("dept"))
		if deptID == 0 && param != "0" {
			deptID, _ = strconv.Atoi(param)
		}
	}
	if deptID > 0 {
		deptinfo, err := dept_getCacheById(int32(deptID))
		if err != nil {
			ws.OutErr(err)
			return gnet.None
		}
		data.Data["dept"] = deptinfo
	}
	/*msg, err := HostConn.GetMsg()
	if err != nil {
		ws.OutErr(err)
		return gnet.None
	}*/
	var err error
	data.Data["deptID"] = strconv.Itoa(deptID)
	data.Data["orderBy"] = ws.Query("orderBy")
	if data.Data["orderBy"].(string) == "" {
		data.Data["orderBy"] = "id"
	}

	data.Data["deptTree"], err = dept_getTreeMenu(data, 0, dept_createMemberLink)
	if err != nil {
		ws.OutErr(err)
		return gnet.None
	}
	getCompanyUser := protocol.GET_MSG_USER_getCompanyUsers()
	getCompanyUser.Type = TYPE
	getCompanyUser.Sort = data.Data["orderBy"].(string)
	getCompanyUser.PerPage = data.Page.PerPage
	getCompanyUser.DeptID = int32(deptID)
	getCompanyUser.Total = data.Page.Total
	getCompanyUser.Page = data.Page.Page
	res, err := HostConn.SendMsgWaitResultToDefault(getCompanyUser)
	if err != nil {
		ws.OutErr(err)
		return gnet.None
	}
	if v, ok := res.(*protocol.MSG_USER_getCompanyUsers_result); ok {
		data.Data["users"] = v.List
		if v.Total > 0 {
			data.Page.Total = v.Total
		}
	}
	data.Data["vars"] = "param=" + param + "&type=" + TYPE + "&orderBy=" + data.Data["orderBy"].(string) + "&recTotal=" + strconv.Itoa(data.Page.Total) + "&recPerPage=" + strconv.Itoa(data.Page.PerPage)
	templateOut("company.browse.html", data, ws)
	return gnet.None
}
func getCompanyInfo() protocol.MSG_USER_Company_cache {
	var c protocol.MSG_USER_Company_cache
	c.Name = "杰骏数码"
	err := HostConn.CacheGet(protocol.UserServerNo, protocol.PATH_USER_COMPANY_CACHE, "1", &c)
	if err != nil {
		libraries.ReleaseLog("获取Company缓存失败%+v", err)
	}
	return c
}
