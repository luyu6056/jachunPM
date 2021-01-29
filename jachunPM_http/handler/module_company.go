package handler

import (
	"jachunPM_http/config"
	"libraries"
	"protocol"
	"strconv"
)

func init() {
	httpHandlerMap["GET"]["/company/browse"] = get_company_browse
	httpHandlerMap["POST"]["/company/browse"] = get_company_browse
}
func get_company_browse(data *TemplateData) {
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
			data.OutErr(err)
			return
		}
		data.Data["dept"] = deptinfo
		param = strconv.Itoa(deptID)
	}
	/*msg, err := HostConn.GetMsg()
	if err != nil {
		data.OutErr(err)
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
		data.OutErr(err)
		return
	}
	getCompanyUser := protocol.GET_MSG_USER_getCompanyUsers()
	getCompanyUser.Type = TYPE
	getCompanyUser.Sort = data.Data["orderBy"].(string)
	getCompanyUser.PerPage = data.Page.PerPage
	getCompanyUser.DeptID = int32(deptID)
	getCompanyUser.Total = data.Page.Total
	if TYPE == "bysearch" {
		getCompanyUser.Where, err = post_search_buildQuery(data)
		if err != nil {
			data.OutErr(err)
			return
		}
	}

	getCompanyUser.Page = data.Page.Page
	var res *protocol.MSG_USER_getCompanyUsers_result
	if err = HostConn.SendMsgWaitResultToDefault(getCompanyUser, &res); err != nil {
		data.OutErr(err)
		return
	}
	data.Data["users"] = res.List
	if res.Total > 0 {
		data.Page.Total = res.Total
	}
	data.Data["queryID"], _ = strconv.Atoi(param)
	if TYPE == "bydept" {
		data.Data["queryID"] = 0
	}

	data.Data["vars"] = "param=" + param + "&type=" + TYPE + "&orderBy=%s" + "&recTotal=" + strconv.Itoa(data.Page.Total) + "&recPerPage=" + strconv.Itoa(data.Page.PerPage)
	templateOut("company.browse.html", data)
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
func init() {
	searchParamsFunc["company/browse"] = func(data *TemplateData) map[string]interface{} {
		search := data.Config["company"]["browse"]["search"].(map[string]interface{})
		search["actionURL"] = createLink("company", "browse", "param=myQueryID&type=bysearch")
		dept := search["params"].(map[string]config.ConfigSearchParams)["dept"]
		dept.Values = dept.Values[:0]
		dept.Values = append(dept.Values, protocol.HtmlKeyValueStr{"", ""})
		list, _ := dept_getOptionMenu(0)
		dept.Values = append(dept.Values, list...)
		search["params"].(map[string]config.ConfigSearchParams)["dept"] = dept
		data.ws.Session().Store("company/browse", search)
		return search
	}
}
