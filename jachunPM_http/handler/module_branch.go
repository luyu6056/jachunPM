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
	httpHandlerMap["GET"]["/branch/ajaxGetBranches"] = get_branch_ajaxGetBranches
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
func get_branch_ajaxGetDropMenu(data *TemplateData) (err error) {
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
	return
}
func get_branch_manage(data *TemplateData) (err error) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	if err = product_setMenu(data, int32(productID), 0, ""); err != nil {
		return
	}
	data.Data["title"] = data.Lang["branch"]["manage"]
	data.Data["branches"] = branch_getPairs(data, int32(productID), nil, "noempty")
	data.Data["product"] = HostConn.GetProductById(int32(productID))
	templateOut("branch.manage.html", data)
	return
}
func post_branch_manage(data *TemplateData) (err error) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	product := HostConn.GetProductById(int32(productID))
	if product == nil {
		return errors.New(data.Lang["product"]["error"].(map[string]string)["NotFound"])
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
		if err := data.SendMsgWaitResultToDefault(out, nil); err != nil {
			data.ws.WriteString(js.Alert(err.Error()))
			return nil
		}
	}
	data.ws.WriteString(js.Alert(data.Lang["common"]["saveSuccess"].(string)) + js.Reload("parent"))
	return
}
func get_branch_delete(data *TemplateData) (err error) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	branchID, _ := strconv.Atoi(data.ws.Query("branchID"))
	product := HostConn.GetProductById(int32(productID))
	if product == nil {
		data.ws.WriteString(js.Alert(data.Lang["product"]["error"].(map[string]string)["NotFound"]))
		return nil
	}
	confirm := data.ws.Query("confirm")
	if confirm != "yes" {
		data.ws.WriteString(js.Confirm(strings.Replace(data.Lang["branch"]["confirmDelete"].(string), "@branch@", data.Lang["product"]["branchName"].(map[string]string)[product.Type], 1), createLink("branch", "delete", []interface{}{"branchID=", branchID, "&productID=", productID, "&confirm=yes"}), ""))
		return nil

	}
	out := protocol.GET_MSG_PROJECT_product_deleteBranch()
	out.BranchID = int32(branchID)
	out.ProductID = int32(productID)
	var result *protocol.MSG_PROJECT_product_deleteBranch_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		data.ws.WriteString(js.Alert(data.Lang["branch"]["error"].(map[string]string)["ErrDelete"], err))
		return nil
	}
	if result.Result != protocol.Success {
		data.ws.WriteString(js.Alert(strings.Replace(data.Lang["branch"]["canNotDelete"].(map[string]string)[result.Result.String()], "@branch@", data.Lang["product"]["branchName"].(map[string]string)[product.Type], 1)))
	} else {
		data.ws.WriteString(js.Reload("parent"))
	}
	out.Put()
	return
}
func post_branch_sort(data *TemplateData) (err error) {
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
		if err := data.SendMsgWaitResultToDefault(out, nil); err != nil {
			data.ws.WriteString(js.Alert(err.Error()))
			return nil
		}
	}
	return
}
func branch_getByProducts(data *TemplateData, products []int32, params string, appendBranch []int32) (res map[int32][]protocol.HtmlKeyValueStr, err error) {
	out := protocol.GET_MSG_PROJECT_branch_getByProducts()
	out.Products = products
	out.AppendBranch = appendBranch
	var result *protocol.MSG_PROJECT_branch_getByProducts_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	out.Put()
	res = make(map[int32][]protocol.HtmlKeyValueStr)
	if !strings.Contains(params, "noempty") {
		for id, kv := range result.List {
			if kv[0].Key != "0" {
				if product := HostConn.GetProductById(int32(id)); product != nil {
					var name = data.Lang["product"]["branchName"].(map[string]string)[product.Type]
					if name != "" {
						kv = append([]protocol.HtmlKeyValueStr{{"0", data.Lang["branch"]["all"].(string) + name}}, kv...)
					}
				}
			}
			res[id] = kv
		}
	} else {
		for id, kv := range result.List {
			res[id] = kv
		}
	}
	return
}
func get_branch_ajaxGetBranches(data *TemplateData) (err error) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	product := HostConn.GetProductById(int32(productID))
	if product == nil || product.Type == "normal" {
		return
	}
	branches := branch_getPairs(data, int32(productID), product, "")
	oldBranch, _ := strconv.Atoi(data.ws.Query("oldBranch"))
	if oldBranch > 0 {
		find := []protocol.HtmlKeyValueStr{{strconv.Itoa(oldBranch), ""}}
		for _, kv := range branches {
			if kv.Key == find[0].Key {
				find[0].Value = kv.Value
				break
			}
		}
		branches = find
	}
	data.ws.WriteString(html_select("branch", branches, "", "class='form-control' onchange='loadBranch(this)'"))
	return
}
