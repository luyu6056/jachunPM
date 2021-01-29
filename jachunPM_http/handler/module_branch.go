package handler

import (
	"errors"
	"jachunPM_http/js"
	"protocol"
	"strconv"
	"strings"
)

func init() {
	httpHandlerMap["GET"]["/branch/ajaxGetDropMenu"] = get_branch_ajaxGetDropMenu
	httpHandlerMap["GET"]["/branch/manage"] = get_branch_manage
	httpHandlerMap["POST"]["/branch/manage"] = post_branch_manage
	httpHandlerMap["GET"]["/branch/delete"] = get_branch_delete
	httpHandlerMap["POST"]["/branch/sort"] = post_branch_sort
}
func branch_getPairs(data *TemplateData, productID int32, productInfo *protocol.MSG_PROJECT_product_cache, params ...string) []protocol.HtmlKeyValueStr {
	if productInfo == nil {
		productInfo = HostConn.GetProductById(productID)
	}
	var res []protocol.HtmlKeyValueStr
	for _, b := range productInfo.Branchs {
		if !b.Deleted {
			res = append(res, protocol.HtmlKeyValueStr{strconv.Itoa(int(b.Id)), b.Name})
		}

	}
	if len(params) == 0 || !strings.Contains(params[0], "noempty") {
		if productInfo == nil || productInfo.Type == "normal" {
			return nil
		}
		res = append([]protocol.HtmlKeyValueStr{{"0", data.Lang["branch"]["all"].(string) + data.Lang["product"]["branchName"].(map[string]string)[productInfo.Type]}}, res...)
	}
	return res
}
func get_branch_ajaxGetDropMenu(data *TemplateData) {
	module := data.ws.Query("module")
	method := data.ws.Query("method")
	extra := data.ws.Query("extra")
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	data.Data["productID"] = productID
	data.Data["link"] = product_getProductLink(module, method, extra, true)
	data.Data["module"] = module
	data.Data["method"] = method
	data.Data["extra"] = extra
	data.Data["branches"] = branch_getPairs(data, int32(productID), nil)
	templateOut("branch.ajaxGetDropMenu.html", data)
}
func get_branch_manage(data *TemplateData) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	if err := product_setMenu(data, int32(productID), 0, ""); err != nil {
		data.OutErr(err)
		return
	}
	data.Data["title"] = data.Lang["branch"]["manage"]
	data.Data["branches"] = branch_getPairs(data, int32(productID), nil, "noempty")
	data.Data["product"] = HostConn.GetProductById(int32(productID))
	templateOut("branch.manage.html", data)
}
func post_branch_manage(data *TemplateData) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	product := HostConn.GetProductById(int32(productID))
	if product == nil {
		data.OutErr(errors.New(data.Lang["product"]["error"].(map[string]string)["NotFound"]))
		return
	}
	out := protocol.GET_MSG_PROJECT_product_editBranch()
	out.ProductID = int32(productID)
	for key, v := range data.ws.GetAllPost() {
		if key == "newbranch" {
			for _, name := range v {
				if name != "" {
					tmp := protocol.GET_MSG_PROJECT_branch_info()
					tmp.Name = name
					tmp.Product = int32(productID)
					out.Branchs = append(out.Branchs, tmp)
				}
			}
		} else {
			id, _ := strconv.Atoi(key)
			for _, branch := range product.Branchs {
				if branch.Id == int32(id) {
					tmp := protocol.GET_MSG_PROJECT_branch_info()
					tmp.Name = v[0]
					tmp.Id = branch.Id
					tmp.Order = branch.Order
					out.Branchs = append(out.Branchs, tmp)
				}
			}

		}

	}

	if len(out.Branchs) > 0 {
		if err := HostConn.SendMsgWaitResultToDefault(out, nil); err != nil {
			data.ws.WriteString(js.Alert(err.Error()))
			return
		}
	}
	data.ws.WriteString(js.Alert(data.Lang["common"]["saveSuccess"].(string)) + js.Reload("parent"))
}
func get_branch_delete(data *TemplateData) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	branchID, _ := strconv.Atoi(data.ws.Query("branchID"))
	product := HostConn.GetProductById(int32(productID))
	if product == nil {
		data.ws.WriteString(js.Alert(data.Lang["product"]["error"].(map[string]string)["NotFound"]))
		return
	}
	confirm := data.ws.Query("confirm")
	if confirm != "yes" {
		data.ws.WriteString(js.Confirm(strings.Replace(data.Lang["branch"]["confirmDelete"].(string), "@branch@", data.Lang["product"]["branchName"].(map[string]string)[product.Type], 1), createLink("branch", "delete", []interface{}{"branchID=", branchID, "&productID=", productID, "&confirm=yes"}), ""))
		return

	}
	out := protocol.GET_MSG_PROJECT_product_deleteBranch()
	out.BranchID = int32(branchID)
	out.ProductID = int32(productID)
	var result *protocol.MSG_PROJECT_product_deleteBranch_result
	err := HostConn.SendMsgWaitResultToDefault(out, &result)
	if err != nil {
		data.ws.WriteString(js.Alert(data.Lang["branch"]["error"].(map[string]string)["ErrDelete"], err))
		return
	}
	if result.Result != protocol.Success {
		data.ws.WriteString(js.Alert(strings.Replace(data.Lang["branch"]["canNotDelete"].(map[string]string)[result.Result.String()], "@branch@", data.Lang["product"]["branchName"].(map[string]string)[product.Type], 1)))
	} else {
		data.ws.WriteString(js.Reload("parent"))
	}
	out.Put()
}
func post_branch_sort(data *TemplateData) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	product := HostConn.GetProductById(int32(productID))
	if product == nil {
		data.ws.WriteString(js.Alert(data.Lang["product"]["error"].(map[string]string)["NotFound"]))
		return
	}
	out := protocol.GET_MSG_PROJECT_product_editBranch()
	out.ProductID = int32(productID)
	sort := strings.Split(strings.Trim(data.ws.Post("branches"), ","), ",")
	for k, sid := range sort {
		id, _ := strconv.Atoi(sid)
		for _, b := range product.Branchs {
			if b.Id == int32(id) {
				tmp := protocol.GET_MSG_PROJECT_branch_info()
				tmp.Name = b.Name
				tmp.Id = b.Id
				tmp.Order = int16(k)
				out.Branchs = append(out.Branchs, tmp)
				break
			}
		}

	}

	if len(out.Branchs) > 0 {
		if err := HostConn.SendMsgWaitResultToDefault(out, nil); err != nil {
			data.ws.WriteString(js.Alert(err.Error()))
			return
		}
	}

}
