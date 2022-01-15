package handler

import (
	"errors"
	"jachunPM_http/js"
	"protocol"
	"strconv"
	"strings"
	"time"
)

func init() {

	httpHandlerMap["GET"]["/productplan/browse"] = get_productplan_browse
	httpHandlerMap["GET"]["/productplan/create"] = get_productplan_create
	httpHandlerMap["POST"]["/productplan/create"] = post_productplan_create
	httpHandlerMap["GET"]["/productplan/edit"] = get_productplan_edit
	httpHandlerMap["POST"]["/productplan/edit"] = post_productplan_create
	httpHandlerMap["GET"]["/productplan/delete"] = get_productplan_delete
}
func get_productplan_browse(data *TemplateData) (err error) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	branch, _ := strconv.Atoi(data.ws.Query("branch"))
	browseType := data.ws.Query("browseType")
	if browseType == "" {
		browseType = "all"
	}
	orderBy := data.ws.Query("orderBy")
	if orderBy == "" {
		orderBy = "begin_desc"
	}
	product, err := productplan_commonAction(data, int32(productID), int32(branch))
	if err != nil {
		return
	}
	data.Data["currentProductType"] = product.Type
	data.Data["title"] = product.Name + data.Lang["common"]["colon"].(string) + data.Lang["productplan"]["browse"].(string)
	data.Data["productID"] = productID
	data.Data["browseType"] = browseType
	data.Data["orderBy"] = orderBy
	out := protocol.GET_MSG_PROJECT_productplan_getList()
	out.ProductID = int32(productID)
	out.Branch = int32(branch)
	out.BrowseType = browseType
	out.Order = orderBy
	out.Page = data.Page.Page
	out.PerPage = data.Page.PerPage
	out.Total = data.Page.Total
	var result *protocol.MSG_PROJECT_productplan_getList_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	for _, v := range result.List {
		v["isClickableKey"] = "null"
	}
	data.Data["plans"] = result.List
	templateOut("productplan.browse.html", data)
	return
}
func productplan_commonAction(data *TemplateData, productID, branch int32) (product *protocol.MSG_PROJECT_product_cache, err error) {
	product = HostConn.GetProductById(int32(productID))
	if product == nil {
		data.ws.Redirect(createLink("product", "create", nil))
		return nil, errors.New("")
	}

	data.Data["product"] = product
	data.Data["branch"] = branch
	if product.Type != "normal" {
		data.Data["branches"] = branch_getPairs(data, productID, product, "")
	}
	err = product_setMenu(data, productID, branch, "")
	return
}
func get_productplan_create(data *TemplateData) (err error) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	branch, _ := strconv.Atoi(data.ws.Query("branch"))

	product, err := productplan_commonAction(data, int32(productID), int32(branch))
	if err != nil {
		return
	}
	out := protocol.GET_MSG_PROJECT_productplan_getLast()
	out.ProductId = int32(productID)
	var result *protocol.MSG_PROJECT_productplan_getLast_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	if len(result.Result) > 0 {
		t, err := time.Parse("2006-01-02", result.Result["End"])
		if err != nil {
			return err
		}
		delta := 1
		if int(t.Weekday()) == 5 || int(t.Weekday()) == 6 {
			delta = 8 - int(t.Weekday())
		}
		data.Data["begin"] = t.AddDate(0, 0, delta).Format("2006-01-02")

	} else {
		data.Data["begin"] = ""
	}

	data.Data["title"] = product.Name + data.Lang["common"]["colon"].(string) + data.Lang["productplan"]["create"].(string)
	data.Data["lastPlan"] = result.Result
	templateOut("productplan.create.html", data)
	return
}
func productplan_getPairs(data *TemplateData, product int32, branch int32, expired string) ([]protocol.HtmlKeyValueStr, error) {
	out := protocol.GET_MSG_PROJECT_productplan_getPairs()
	out.ProductID = product
	out.BranchID = branch
	out.Expired = expired
	var result *protocol.MSG_PROJECT_productplan_getPairs_result
	if err := data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return nil, err
	}
	productplan_processFuture(data, result.List)
	return result.List, nil
}
func productplan_processFuture(data *TemplateData, res []protocol.HtmlKeyValueStr) {
	for k, v := range res {
		res[k] = protocol.HtmlKeyValueStr{v.Key, strings.Replace(v.Value, "[2030-01-01 ~ 2030-01-01]", "["+data.Lang["productplan"]["future"].(string)+"]", 1)}
	}
}
func post_productplan_create(data *TemplateData) (e error) {
	if !data.ajaxCheckPost() {
		return
	}
	future := data.ws.Post("future") == "1"
	begin := data.ws.Post("begin")
	end := data.ws.Post("end")
	if future || (begin == "" && end == "") {
		if begin == "" {
			begin = "2030-01-01"
		}
		if end == "" {
			end = "2030-01-01"
		}
	}
	id, _ := strconv.Atoi(data.ws.Query("planID"))
	parent, _ := strconv.Atoi(data.ws.Post("parent"))
	branch, _ := strconv.Atoi(data.ws.Post("branch"))
	productID, _ := strconv.Atoi(data.ws.Post("product"))
	product := HostConn.GetProductById(int32(productID))
	findBranch := branch == 0
	if branch > 0 {
		if product != nil {
			for _, b := range product.Branchs {
				if b.Id == int32(branch) {
					findBranch = true
					break
				}
			}
		}
	}

	if product == nil || !findBranch {
		data.ajaxResult(false, data.Lang["product"]["error"].(map[string]string)["NotFound"])
		return
	}
	insert := protocol.GET_MSG_PROJECT_productplan_insertUpdate()
	insert.Id = int32(id)
	insert.Title = data.ws.Post("title")
	insert.Desc = data.ws.Post("desc")
	insert.Begin, _ = time.Parse(protocol.TIMEFORMAT_MYSQLDATE, begin)
	insert.End, _ = time.Parse(protocol.TIMEFORMAT_MYSQLDATE, end)
	if insert.Begin.Unix() > insert.End.Unix() {
		data.ajaxResult(false, map[string]string{"begin": data.Lang["productplan"]["error"].(map[string]string)["beginGeEnd"]})
		return
	}
	insert.Branch = int32(branch)
	insert.Product = int32(productID)
	insert.Parent = int32(parent)

	desc, newimgids, err := file_descProcessImgURLAnd2Bbcode(data, insert.Desc)
	if err != nil {
		data.ajaxResult(false, map[string]string{"desc": err.Error()})
		return
	}

	insert.Desc = desc
	defer func() {
		if err != nil { //以下使用err来判断图片删除
			file_deleteFromIds(data, newimgids)
		}
	}()
	var result *protocol.MSG_PROJECT_productplan_insertUpdate_result
	err = data.SendMsgWaitResultToDefault(insert, &result)
	if err != nil {
		data.ajaxResult(false, err.Error())
		return
	}
	if result != nil {
		if result.Result != protocol.Success {
			data.ajaxResult(false, data.Lang["productplan"]["error"].(map[string]string)[result.Result.String()])
			return
		}
		insert.Id = result.Id
	}
	file_updateObject(data, newimgids, "productplan", insert.Id)
	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "", "function(){parent.parent.$('a.refresh').click()}"))
	} else {
		data.ajaxResult(true, data.Lang["common"]["saveSuccess"], createLink("productplan", "browse", []interface{}{"productID=", productID, "&branch", branch}))
	}
	insert.Put()
	result.Put()

	return
}
func get_productplan_edit(data *TemplateData) (err error) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	branch, _ := strconv.Atoi(data.ws.Query("branch"))
	planID, _ := strconv.Atoi(data.ws.Query("planID"))
	product, err := productplan_commonAction(data, int32(productID), int32(branch))
	if err != nil {
		return
	}
	out := protocol.GET_MSG_PROJECT_productplan_getList()
	out.Ids = []int32{int32(planID)}
	var result *protocol.MSG_PROJECT_productplan_getList_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	if len(result.List) == 0 {
		return errors.New(data.Lang["productplan"]["error"].(map[string]string)["NotFoundProductPlanInfo"])
	}
	data.Data["plan"] = result.List[0]
	data.Data["product"] = product
	templateOut("productplan.edit.html", data)
	return
}
func get_productplan_delete(data *TemplateData) (e error) {
	confirm := data.ws.Query("confirm")
	planID, _ := strconv.Atoi(data.ws.Query("planID"))
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	branch, _ := strconv.Atoi(data.ws.Query("branch"))
	if confirm != "yes" {
		data.ws.WriteString(js.Confirm(data.Lang["productplan"]["confirmDelete"].(string), createLink("productPlan", "delete", []interface{}{"planID=", planID, "&productID=", productID, "&branch=", branch, "&confirm=yes"}), ""))
		return
	}
	out := protocol.GET_MSG_PROJECT_productplan_delete()
	out.Id = int32(planID)
	out.Product = int32(productID)
	out.Branch = int32(branch)
	err := data.SendMsgWaitResultToDefault(out, nil)
	if data.isajax() {
		if err != nil {
			data.ajaxResult(false, err.Error())
		} else {
			data.ajaxResult(true, "")
		}
	} else {
		if err != nil {
			data.ws.WriteString(js.Alert(err.Error()))
		} else {
			data.ws.WriteString(js.Location(createLink("productplan", "browse", []interface{}{"productID=", productID, "&branch=", branch}), "parent"))
		}
	}
	return
}
func productplanTemplateFuncs() {
}
