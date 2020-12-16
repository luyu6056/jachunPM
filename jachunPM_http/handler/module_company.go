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
		deptID, _ = strconv.Atoi(param)
	}
	data.Data["deptID"] = strconv.Itoa(deptID)
	data.Data["orderBy"] = ws.Query("orderBy")
	if data.Data["orderBy"].(string) == "" {
		data.Data["orderBy"] = "id"
	}
	data.Data["vars"] = "param=" + param + "&type=" + TYPE + "&orderBy=%s&recTotal={$pager->recTotal}&recPerPage={$pager->recPerPage}"
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
