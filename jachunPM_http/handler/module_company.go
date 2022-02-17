package handler

import (
	"config"
	"libraries"
	"protocol"
	"strconv"
)

func init() {
	httpHandlerMap["GET"]["/company/browse"] = get_company_browse
	httpHandlerMap["POST"]["/company/browse"] = get_company_browse
	httpHandlerMap["GET"]["/company/updateCache"] = get_company_updateCache

}
func get_company_browse(data *TemplateData) (err error) {
	ws := data.ws
	param, _ := strconv.Atoi(ws.Query("param"))
	TYPE := ws.Query("type")

	deptID := 0
	if TYPE == "" {
		TYPE = "bydept"
	}
	if TYPE == "bydept" {
		deptID, _ = strconv.Atoi(ws.Query("dept"))
		if deptID == 0 && param != 0 {
			deptID = param
		}
	}
	if deptID > 0 {
		deptinfo, err := HostConn.GetdeptCacheById(int32(deptID))
		if err != nil {
			return err
		}
		data.Data["dept"] = deptinfo
		param = deptID
	}
	/*msg, err := HostConn.GetMsg()
	if err != nil {
		data.OutErr(err)
		return gnet.None
	}*/
	data.Data["deptID"] = strconv.Itoa(deptID)
	data.Data["orderBy"] = ws.Query("orderBy")
	if data.Data["orderBy"].(string) == "" {
		data.Data["orderBy"] = "id"
	}

	if data.Data["deptTree"], err = dept_getTreeMenu(data, 0, dept_createMemberLink); err != nil {
		return
	}
	getCompanyUser := protocol.GET_MSG_USER_getCompanyUsers()
	getCompanyUser.Type = TYPE
	getCompanyUser.Sort = data.Data["orderBy"].(string)
	getCompanyUser.PerPage = data.Page.PerPage
	getCompanyUser.DeptID = int32(deptID)
	getCompanyUser.Total = data.Page.Total
	if TYPE == "bysearch" {
		if getCompanyUser.Where, err = post_search_buildQuery(data, param); err != nil {
			return
		}
	}

	getCompanyUser.Page = data.Page.Page
	var res *protocol.MSG_USER_getCompanyUsers_result
	if err = data.SendMsgWaitResultToDefault(getCompanyUser, &res); err != nil {
		return
	}
	data.Data["users"] = res.List
	if res.Total > 0 {
		data.Page.Total = res.Total
	}
	data.Data["queryID"] = param
	if TYPE == "bydept" {
		data.Data["queryID"] = 0
	}

	data.Data["vars"] = "param=" + strconv.Itoa(param) + "&type=" + TYPE + "&orderBy=%s" + "&recTotal=" + strconv.Itoa(data.Page.Total) + "&recPerPage=" + strconv.Itoa(data.Page.PerPage)
	templateOut("company.browse.html", data)
	return
}
func getCompanyInfo() protocol.MSG_USER_Company_cache {
	return companyCache
}
func get_company_updateCache(data *TemplateData) error {
	Company_updateCache()
	data.ws.WriteString("ok")
	return nil
}
func Company_updateCache() {
	err := HostConn.CacheGet(protocol.UserServerNo, protocol.PATH_USER_COMPANY_CACHE, "1", &companyCache)
	if err != nil {
		libraries.ReleaseLog("获取Company缓存失败%+v", err)
	}
}

var companyCache protocol.MSG_USER_Company_cache

func init() {

	searchParamsFunc["company/browse"] = func(data *TemplateData) (*searchParam, error) {
		search := &searchParam{
			ConfigSearch: data.Config["company"]["browse"]["search"].(*config.ConfigSearch),
		}

		search.ActionURL = createLink("company", "browse", "param=myQueryID&type=bysearch")
		dept := search.Params["dept"]
		dept.Values = dept.Values[:0]
		dept.Values = append(dept.Values, protocol.HtmlKeyValueStr{"", ""})
		list, _ := dept_getOptionMenu(0)
		dept.Values = append(dept.Values, list...)
		search.Params["dept"] = dept
		data.ws.Session().Store("company/browse", search)
		return search, nil
	}
}
