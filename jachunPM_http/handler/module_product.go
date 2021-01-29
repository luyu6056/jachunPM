package handler

import (
	"errors"
	"fmt"
	"html/template"
	"jachunPM_http/js"
	"libraries"
	"protocol"
	"strconv"
	"strings"
	"time"
)

func init() {
	httpHandlerMap["GET"]["/product/index"] = get_product_index
	httpHandlerMap["GET"]["/product/create"] = get_product_create
	httpHandlerMap["POST"]["/product/create"] = post_product_create
	httpHandlerMap["GET"]["/product/browse"] = get_product_browse
	httpHandlerMap["GET"]["/product/ajaxGetDropMenu"] = get_product_ajaxGetDropMenu
	httpHandlerMap["GET"]["/product/all"] = get_product_all
	httpHandlerMap["GET"]["/product/view"] = get_product_view
	httpHandlerMap["GET"]["/product/edit"] = get_product_edit
	httpHandlerMap["POST"]["/product/edit"] = post_product_edit
	httpHandlerMap["GET"]["/product/ajaxGetPlans"] = get_product_ajaxGetPlans
}
func get_product_index(data *TemplateData) {

	if data.ws.Query("locate") == "yes" {
		data.ws.Redirect(createLink("product", "browse", nil))
		return
	}
	//if($this->app->getViewType() != 'mhtml') unset($this->lang->product->menu->index);
	id, _ := strconv.Atoi(data.ws.Query("productID"))
	productID, branch, err := product_saveState(data, int32(id))
	if err != nil {
		data.OutErr(err)
		return
	}
	err = product_setMenu(data, productID, branch, "")
	if err != nil {
		data.OutErr(err)
		return
	}
	templateOut("product.index.html", data)
	return
}
func get_product_create(data *TemplateData) {

	productID, branch, err := product_saveState(data, 0)
	if err != nil {
		data.OutErr(err)
		return
	}
	err = product_setMenu(data, productID, branch, "")
	data.Data["groups"], _ = user_getGroupOptionMenu()
	msg, err := data.GetMsg()
	if err != nil {
		data.OutErr(err)
		return
	}
	data.Data["poUsers"], err = user_getPairs("nodeleted|pofirst|noclosed")
	if err != nil {
		data.OutErr(err)
		return
	}
	data.Data["qdUsers"], err = user_getPairs("nodeleted|qdfirst|noclosed")
	if err != nil {
		data.OutErr(err)
		return
	}
	data.Data["rdUsers"], err = user_getPairs("nodeleted|devfirst|noclosed")
	if err != nil {
		data.OutErr(err)
		return
	}

	var productTypeList []protocol.HtmlKeyValueStr
	for _, v := range data.Lang["product"]["typeList"].([]protocol.HtmlKeyValueStr) {
		tip, _ := data.Lang["product"]["typeTips"].(map[string]string)[v.Key]
		productTypeList = append(productTypeList, protocol.HtmlKeyValueStr{v.Key, v.Value + tip})

	}
	data.Data["productTypeList"] = productTypeList
	getLinePairs := protocol.GET_MSG_PROJECT_tree_getLinePairs()
	var res3 *protocol.MSG_PROJECT_tree_getLinePairs_result
	if err = msg.SendMsgWaitResult(0, getLinePairs, &res3); err != nil {
		data.OutErr(err)
		return
	}
	res3.List = append([]protocol.HtmlKeyValueStr{{"", ""}}, res3.List...)
	data.Data["lines"] = res3.List
	data.Data["rootID"] = productID
	templateOut("product.create.html", data)
	res3.Put()
	return
}
func post_product_create(data *TemplateData) {
	if !data.ajaxCheckPost() {
		return
	}
	msg, err := data.GetMsg()
	if err != nil {
		data.ajaxResult(false, err.Error())
		return
	}
	out := protocol.GET_MSG_PROJECT_product_insert()
	insert := protocol.GET_MSG_PROJECT_product_cache()
	insert.Status = "normal"
	insert.CreatedBy = data.User.Id
	insert.CreatedDate = time.Now().Unix()

	for key, v := range data.ws.GetAllPost() {
		switch key {
		case "acl":
			insert.Acl = v[0]
		case "whitelist":
			for _, sid := range v {
				id, _ := strconv.Atoi(sid)
				if id > 0 {
					insert.Whitelist = append(insert.Whitelist, int32(id))
				}
			}
		case "name":
			insert.Name = v[0]
		case "code":
			insert.Code = v[0]
		case "line":
			id, _ := strconv.Atoi(v[0])
			insert.Line = int32(id)
		case "PO":
			id, _ := strconv.Atoi(v[0])
			insert.PO = int32(id)
		case "QD":
			id, _ := strconv.Atoi(v[0])
			insert.QD = int32(id)
		case "RD":
			id, _ := strconv.Atoi(v[0])
			insert.RD = int32(id)
		case "type":
			insert.Type = v[0]
		case "desc":
			insert.Desc = v[0]

		}
	}
	desc, newimgids, err := file_descProcessImgURLAnd2Bbcode(data, insert.Desc)
	if err != nil {
		data.ajaxResult(false, map[string]string{"desc": err.Error()})
		return
	}
	insert.Desc = desc
	defer func() {
		if err != nil { //以下使用err来判断图片删除
			file_deleteFromIds(newimgids)
		}
	}()
	out.Data = insert
	out.DocName = data.Lang["doclib"]["main"].(map[string]string)["product"]
	var res *protocol.MSG_PROJECT_product_insert_result
	err = msg.SendMsgWaitResult(0, out, &res)
	if err != nil {
		data.ajaxResult(false, err.Error())
		return
	}
	locate := createLink("product", "browse", []string{"productID=", strconv.Itoa(int(res.ID))})
	data.ajaxResult(true, data.Lang["common"]["saveSuccess"], locate)
	file_updateObject(newimgids, "product", res.ID)
	out.Put()
	res.Put()
	return
}
func get_product_browse(data *TemplateData) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	branch, _ := strconv.Atoi(data.ws.Query("branch"))
	moduleID, _ := strconv.Atoi(data.ws.Query("param"))
	browseType := data.ws.Query("browseType")
	orderBy := data.ws.Query("orderBy")
	if orderBy == "" {
		orderBy = "id_desc"
	}
	if browseType == "" {
		browseType = "unclosed"
	}
	msg, err := data.GetMsg()
	if err != nil {
		data.OutErr(err)
		return
	}
	err = product_setMenu(data, int32(productID), int32(branch), "")
	if err != nil {
		data.OutErr(err)
		return
	}
	getStories := protocol.GET_MSG_PROJECT_product_getStories()
	getStories.ProductID = int32(productID)
	getStories.Branch = int32(branch)
	getStories.BrowseType = browseType
	getStories.ModuleID = int32(moduleID)
	getStories.Sort = orderBy
	getStories.Page = data.Page.Page
	getStories.PerPage = data.Page.PerPage
	getStories.Total = data.Page.Total
	var stories *protocol.MSG_PROJECT_product_getStories_result
	err = msg.SendMsgWaitResult(0, getStories, &stories)
	if err != nil {
		data.OutErr(err)
		return
	}
	if moduleID > 0 {
		if module := HostConn.GetTreeById(int32(moduleID)); module != nil {
			data.Data["moduleName"] = module.Name
		}

	}
	if data.Data["moduleName"] == nil {
		data.Data["moduleName"] = data.Lang["tree"]["all"].(string)
	}

	data.Data["stories"] = stories.List
	data.Data["productID"] = productID
	data.Data["branch"] = branch
	data.Data["moduleID"] = moduleID
	data.Data["browseType"] = browseType
	templateOut("product.browse.html", data)
	return
}
func product_getAll(data *TemplateData) (result []*protocol.MSG_PROJECT_product_cache, err error) {
	if data.Data["product_getAll"] == nil {
		res, err := HostConn.CacheGetPath(protocol.ProjectServerNo, protocol.PATH_PROJECT_PRODUCT_CACHE)
		if err != nil {
			return nil, err
		}

		buf := bufpool.Get().(*libraries.MsgBuffer)
		for _, b := range res {
			buf.Reset()
			buf.Write(b)
			if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_PROJECT_product_cache); ok {
				result = append(result, v)
			}
		}
		buf.Reset()
		bufpool.Put(buf)
		data.Data["product_getAll"] = result
	}

	return data.Data["product_getAll"].([]*protocol.MSG_PROJECT_product_cache), nil
}
func product_getPairs(data *TemplateData, mode ...string) (res []protocol.HtmlKeyValueStr, err error) {
	list, err := product_getAll(data)
	if err != nil {
		return nil, err
	}
	protocol.Order_product(list, func(a, b *protocol.MSG_PROJECT_product_cache) bool {
		if a.Status == "close" {
			return false
		}
		if a.Order < b.Order {
			return false
		}
		return true
	})
	for _, p := range list {
		if len(mode) == 1 && mode[0] == "noclosed" && p.Status == "closed" {
			continue
		}
		if data.User.Id != 1 && !data.User.AclProducts[p.Id] {
			continue
		}
		res = append(res, protocol.HtmlKeyValueStr{strconv.Itoa(int(p.Id)), p.Name})
	}
	return
}
func product_setMenu(data *TemplateData, productID, branch int32, extra string) error {
	products, err := product_getPairs(data)
	if err != nil {
		return err
	}
	var find bool
	productIDStr := strconv.Itoa(int(productID))
	for _, p := range products {
		if p.Key == productIDStr {
			find = true
			break
		}
	}
	if !find {
		return errors.New(data.Lang["product"]["error"].(map[string]string)["NotFound"])
	}
	currentModule := data.App["moduleName"].(string)
	currentMethod := data.App["methodName"].(string)

	if currentModule == "story" {
		if currentMethod != `create` && currentMethod != `batchcreate` {
			currentModule = `product`
		}
		if currentMethod == `view` {
			currentMethod = `browse`
		}
	}
	if currentMethod == `report` {
		currentMethod = `browse`
	}

	//selectHtml = this->select(products, productID, currentModule, currentMethod, extra, branch, module, moduleType);

	label := data.Lang["product"]["index"].(string)
	if currentModule == `product` && currentMethod == `all` {
		label = data.Lang["product"]["all"].(string)
	}
	if currentModule == `product` && currentMethod == `create` {
		label = data.Lang["product"]["create"].(string)
	}

	/*isMobile = this->app->viewType == `mhtml`;
	  if(isMobile)
	  {
	      pageNav  = html::a(helper::createLink(`product`, `index`), this->lang->product->index) . this->lang->colon;
	      pageNav .= selectHtml;
	  }
	  else
	  {*/
	buf := bufpool.Get().(*libraries.MsgBuffer)
	defer func() {
		buf.Reset()
		bufpool.Put(buf)
	}()
	selectHtml := func() error {
		//public function select($products, $productID, $currentModule, $currentMethod, $extra = '', $branch = 0, $module = 0, $moduleType = '')
		//$isMobile = $this->app->viewType == 'mhtml';
		data.ws.SetCookie("lastProduct", productIDStr, protocol.SessionKeepLoginExpires)

		currentProduct := HostConn.GetProductById(int32(productID))
		if currentProduct == nil {
			return errors.New(data.Lang["product"]["error"].(map[string]string)["NotFound"])
		}
		data.ws.Session().Set("currentProductType", currentProduct.Type)
		buf.WriteString(`<div class='btn-group angle-btn'><div class='btn-group'><button data-toggle='dropdown' type='button' class='btn btn-limit' id='currentItem' title='`)
		buf.WriteString(currentProduct.Name)
		buf.WriteString("'>")
		buf.WriteString(currentProduct.Name)
		buf.WriteString(`<span class='caret'></span></button><div id='dropMenu' class='dropdown-menu search-list' data-ride='searchList' data-url='`)
		buf.WriteString(createLink("product", "ajaxGetDropMenu", []interface{}{"productID=", productID, "&module=", currentModule, "&method=", currentMethod, "&extra=", extra}))
		buf.WriteString(`'><div class='input-control search-box has-icon-left has-icon-right search-example'><input type='search' class='form-control search-input' /><label class='input-control-icon-left search-icon'><i class='icon icon-search'></i></label><a class='input-control-icon-right search-clear-btn'><i class='icon icon-close icon-sm'></i></a></div></div></div>`)

		//if($isMobile) $output = "<a id='currentItem' href=\"javascript:showSearchMenu('product', '$productID', '$currentModule', '$currentMethod', '$extra')\">{$currentProduct->name} <span class='icon-caret-down'></span></a><div id='currentItemDropMenu' class='hidden affix enter-from-bottom layer'></div>";

		//if($currentProduct->type == 'normal') unset($this->lang->product->menu->branch);
		if currentProduct.Type != "normal" {
			branches := branch_getPairs(data, 0, currentProduct)
			branchName := branches[0].Value
			if branch > 0 {
				for _, kv := range branches {
					if kv.Key == strconv.Itoa(int(branch)) {
						branchName = kv.Value
						break
					}
				}
			}

			if true { //!$isMobile){
				buf.WriteString(`<div class='btn-group'><button id='currentBranch' data-toggle='dropdown' type='button' class='btn btn-limit'>`)
				buf.WriteString(branchName)
				buf.WriteString(`<span class='caret'></span></button><div id='dropMenu' class='dropdown-menu search-list' data-ride='searchList' data-url='`)
				buf.WriteString(createLink("branch", "ajaxGetDropMenu", []interface{}{"productID=", productID, "&module=", currentModule, "&method=", currentMethod, "&extra=", extra}))
				buf.WriteString(`'><div class='input-control search-box has-icon-left has-icon-right search-example'><input type='search' class='form-control search-input' /><label class='input-control-icon-left search-icon'><i class='icon icon-search'></i></label><a class='input-control-icon-right search-clear-btn'><i class='icon icon-close icon-sm'></i></a></div></div></div>`)
			} else {
				buf.WriteString("<a id='currentBranch' href=\"javascript:showSearchMenu('branch', '")
				buf.WriteString(strconv.Itoa(int(productID)))
				buf.WriteString("', '")
				buf.WriteString(currentModule)
				buf.WriteString("', '")
				buf.WriteString(currentMethod)
				buf.WriteString("', '")
				buf.WriteString(extra)
				buf.WriteString("')\">")
				buf.WriteString(branchName)
				buf.WriteString("<span class='icon-caret-down'></span></a><div id='currentBranchDropMenu' class='hidden affix enter-from-bottom layer'></div>")
			}
		}

		if true { //!$isMobile)
			buf.WriteString("</div>")
		}
		return nil
	}
	buf.WriteString(`<div class="btn-group angle-btn`)
	if currentMethod == `index` {
		buf.WriteString(" active")
	}
	buf.WriteString(`"><div class="btn-group"><button data-toggle="dropdown" type="button" class="btn">`)
	buf.WriteString(label)
	buf.WriteString(` <span class="caret"></span></button><ul class="dropdown-menu">`)
	if hasPriv(data, `product`, `index`) {
		buf.WriteString(`<li>`)
		buf.WriteString(html_a(createLink(`product`, `index`, `locate=no`), `<i class="icon icon-home"></i> `+data.Lang["product"]["index"].(string)))
		buf.WriteString(`</li>`)
	}
	if hasPriv(data, `product`, `all`) {
		buf.WriteString(`<li>`)
		buf.WriteString(html_a(createLink(`product`, `all`, nil), `<i class="icon icon-cards-view"></i> `+data.Lang["product"]["all"].(string)))
		buf.WriteString(`</li>`)
	}
	if hasPriv(data, `product`, `create`) {
		buf.WriteString(`<li>`)
		buf.WriteString(html_a(createLink(`product`, `create`, nil), `<i class="icon icon-plus"></i> `+data.Lang["product"]["create"].(string)))
		buf.WriteString(`</li>`)
	}
	buf.WriteString(`</ul></div></div>`)
	if err = selectHtml(); err != nil {
		return err
	}

	data.Data["modulePageNav"] = template.HTML(buf.String())
	if data.App["menuReplace"] == nil {
		data.App["menuReplace"] = make(map[string]string)
	}
	data.App["menuReplace"].(map[string]string)["productID"] = strconv.Itoa(int(productID))
	data.App["menuReplace"].(map[string]string)["branch"] = strconv.Itoa(int(branch))
	return nil
}
func product_saveState(data *TemplateData, id int32) (productID int32, preBranch int32, err error) {
	products, err := product_getPairs(data, "nocode")
	if err != nil {
		return 0, 0, err
	}
	if id > 0 {
		data.ws.Session().Set("product", id)
	} else {
		if data.ws.Cookie("lastProduct") > "0" {
			data.ws.Session().Set("product", data.ws.Cookie("lastProduct"))
		}
		if _, ok := data.ws.Session().Load("product"); !ok && len(products) > 0 {
			data.ws.Session().Set("product", products[0].Key)
		}
	}
	if len(products) > 0 {
		productstr := data.ws.Session().Load_str("product")
		var find bool
		for _, p := range products {
			if p.Key == productstr {
				find = true
				break
			}
		}
		if !find {
			data.ws.Session().Set("product", products[0].Key)
			if id > 0 {
				return 0, 0, errors.New(data.Lang["product"]["accessDenied"].(string))
			}
		}
	}
	preBranchID, _ := strconv.Atoi(data.ws.Cookie("preBranch"))
	if data.ws.Cookie("preProductID") != strconv.Itoa(int(id)) {
		data.ws.SetCookie("preBranch", "0", protocol.SessionKeepLoginExpires)
		preBranchID = 0
	}
	return data.ws.Session().Load_int32("product"), int32(preBranchID), nil
}
func get_product_ajaxGetDropMenu(data *TemplateData) {
	method := data.ws.Query("method")
	module := data.ws.Query("module")
	productID, _ := strconv.Atoi(data.ws.Query("objectID"))
	link := product_getProductLink(module, method, data.ws.Query("extra"), false)
	data.Data["productID"] = productID
	data.Data["module"] = module
	data.Data["method"] = method
	data.Data["extra"] = data.ws.Query("extra")
	products, err := product_getAll(data)
	if err != nil {
		data.OutErr(err)
		return
	}
	lines, err := tree_getLinePairs(data)
	if err != nil {
		data.OutErr(err)
		return
	}

	for _, kv := range lines {
		for _, product := range products {
			if strconv.Itoa(int(product.Line)) == kv.Key {
				product.Name = kv.Value + "/" + product.Name
				product.Order = -1
			}
		}
	}
	protocol.Order_product(products, func(a, b *protocol.MSG_PROJECT_product_cache) bool {
		if a.Status == "close" {
			return false
		}
		if a.Order > b.Order {
			return false
		}
		if a.Id > b.Id {
			return false
		}
		return true
	})
	myProductsHtml := bufpool.Get().(*libraries.MsgBuffer)
	normalProductsHtml := bufpool.Get().(*libraries.MsgBuffer)
	closedProductsHtml := bufpool.Get().(*libraries.MsgBuffer)

	for _, product := range products {

		if data.User.Id != 1 && !data.User.AclProducts[product.Id] {
			continue
		}
		if product.Status == "normal" && product.PO == data.User.Id {
			if product.Type != "platform" && module == "branch" && method == "manage" {
				myProductsHtml.WriteString(html_a(fmt.Sprintf(link, productID), "<i class='icon icon-cube'></i> "+product.Name, "", "class='text-important' title='"+product.Name+"'"))
			} else {
				myProductsHtml.WriteString(html_a(fmt.Sprintf(link, product.Id), "<i class='icon icon-cube'></i> "+product.Name, "", "class='text-important' title='"+product.Name+"'"))
			}
		} else if product.Status == "normal" && (product.PO != data.User.Id) {
			if product.Type != "platform" && module == "branch" && method == "manage" {
				normalProductsHtml.WriteString(html_a(fmt.Sprintf(link, productID), "<i class='icon icon-cube'></i> "+product.Name, "", "title='"+product.Name+"'"))
			} else {
				normalProductsHtml.WriteString(html_a(fmt.Sprintf(link, product.Id), "<i class='icon icon-cube'></i> "+product.Name, "", "title='"+product.Name+"'"))
			}
		} else if product.Status == "closed" {
			if product.Type != "platform" && module == "branch" && method == "manage" {
				closedProductsHtml.WriteString(html_a(fmt.Sprintf(link, productID), "<i class='icon icon-cube'></i> "+product.Name, "", "title='"+product.Name+"' class='closed'"))
			} else {
				closedProductsHtml.WriteString(html_a(fmt.Sprintf(link, product.Id), "<i class='icon icon-cube'></i> "+product.Name, "", "title='"+product.Name+"' class='closed'"))
			}
		}
	}
	data.Data["myProductsHtml"] = template.HTML(myProductsHtml.String())
	data.Data["normalProductsHtml"] = template.HTML(normalProductsHtml.String())
	data.Data["closedProductsHtml"] = template.HTML(closedProductsHtml.String())
	myProductsHtml.Reset()
	normalProductsHtml.Reset()
	closedProductsHtml.Reset()
	bufpool.Put(myProductsHtml)
	bufpool.Put(normalProductsHtml)
	bufpool.Put(closedProductsHtml)
	data.Data["products"] = products
	templateOut("product.ajaxGetDropMenu.html", data)
	return
}
func product_getProductLink(module, method, extra string, branch bool) string {

	switch module {
	case "product", "roadmap", "bug", "testcase", "testtask", "story", "qa", "testsuite", "testreport", "build":
		switch {
		case module == "product" && method == "project":
			if branch {
				return createLink(module, method, "status=all&productID=%d&branch=%s")
			} else {
				return createLink(module, method, "status=all&productID=%d")
			}
		case (module == "product" && (method == "dynamic" || method == "doc" || method == "view")):
			return createLink(module, method, "productID=%d")

		case (module == "qa" && method == "index"):
			if branch {
				return createLink("bug", "browse", "productID=%d&branch=%s")
			} else {
				return createLink("bug", "browse", "productID=%d")
			}

		case (module == "product" && (method == "browse" || method == "index" || method == "all")):

			if branch {
				return createLink(module, "browse", "productID=%d&branch=%s")
			} else {
				return createLink(module, "browse", "productID=%d")
			}

		default:
			if branch {
				return createLink(module, method, "productID=%d&branch=%s")
			} else {
				return createLink(module, method, "productID=%d")
			}

		}

	case "productplan", "release":

		if method != "browse" && method != "create" {
			method = "browse"
		}
		if branch {
			return createLink(module, method, "productID=%d&branch=%s")
		} else {
			return createLink(module, method, "productID=%d")
		}

	case "tree":
		if branch {
			return createLink(module, method, "productID=%d&type="+extra+"&currentModuleID=0&branch=%s")
		} else {
			return createLink(module, method, "productID=%d&type="+extra+"&currentModuleID=0")
		}

	case "branch":
		if branch {
			return createLink(module, method, "productID=%d&branch=%s")
		} else {
			return createLink(module, method, "productID=%d")
		}

	case "doc":
		return createLink("doc", "objectLibs", "type=product&objectID=%d&from=product")

	}

	return ""
}
func get_product_all(data *TemplateData) {

	//this->session->set("productList", this->app->getURI(true))
	id, _ := strconv.Atoi(data.ws.Query("productID"))
	productID, branch, err := product_saveState(data, int32(id))
	err = product_setMenu(data, productID, branch, "")
	if err != nil {
		data.OutErr(err)
		return
	}
	line, _ := strconv.Atoi(data.ws.Query("line"))
	orderBy := data.ws.Query("orderBy")
	status := data.ws.Query("status")
	if status == "" {
		status = "noclosed"
	}
	/* Load pager and get tasks. */
	//this->app->loadClass("pager", static = true)
	//pager = new pager(recTotal, recPerPage, pageID)

	data.Data["title"] = data.Lang["product"]["allProduct"]

	data.Data["productStats"], err = product_getStats(data, orderBy, status, int32(line))
	data.Data["lineTree"], err = tree_getTreeMenu(data, 0, "line", 0, tree_createLineLink, map[string]interface{}{"productID": int32(productID), "status": status}, 0)
	if err != nil {
		data.OutErr(err)
		return
	}
	lines, err := tree_getLinePairs(data)
	if err != nil {
		data.OutErr(err)
		return
	}
	data.Data["lines"] = append([]protocol.HtmlKeyValueStr{{"", ""}}, lines...)
	data.Data["productID"] = productID
	data.Data["line"] = line
	data.Data["status"] = status
	data.Data["orderBy"] = orderBy
	templateOut("product.all.html", data)
	return
}
func product_getStats(data *TemplateData, orderBy string, status string, line int32, rootID ...int32) (result []map[string]interface{}, err error) {
	order := func(a, b *protocol.MSG_PROJECT_product_cache) bool {
		if a.Order == b.Order {
			return a.Id < b.Id
		}
		return a.Order < b.Order
	}
	orders := strings.Split(orderBy, "_")
	if len(orders) == 2 {
		switch orders[0] {
		case "id":
			if orders[1] == "asc" {
				order = func(a, b *protocol.MSG_PROJECT_product_cache) bool {
					return a.Id < b.Id
				}
			} else {
				order = func(a, b *protocol.MSG_PROJECT_product_cache) bool {
					return a.Id > b.Id
				}
			}

		case "name":
			if orders[1] == "asc" {
				order = func(a, b *protocol.MSG_PROJECT_product_cache) bool {
					return a.Name < b.Name
				}
			} else {
				order = func(a, b *protocol.MSG_PROJECT_product_cache) bool {
					return a.Name > b.Name
				}
			}
		case "line":
			if orders[1] == "asc" {
				order = func(a, b *protocol.MSG_PROJECT_product_cache) bool {
					return a.Line < b.Line
				}
			} else {
				order = func(a, b *protocol.MSG_PROJECT_product_cache) bool {
					return a.Line > b.Line
				}
			}
		}
	}

	if status == "" {
		status = "noclosed"
	}
	id := int32(0)
	if len(rootID) == 1 {
		id = rootID[0]
	}
	products, err := product_getList(data, order, status, 0, line, id)
	if err != nil {
		return
	}
	if data.Page.Total == 0 {
		data.Page.Total = len(products)
	}
	if (data.Page.Page-1)*data.Page.PerPage >= len(products) {
		return
	}
	end := (data.Page.Page) * data.Page.PerPage
	if end > len(products) {
		end = len(products)
	}
	products = products[(data.Page.Page-1)*data.Page.PerPage : end]
	var ids = make([]int32, len(products))
	for k, p := range products {
		ids[k] = p.Id
	}
	getstories := protocol.GET_MSG_PROJECT_product_getStoriesMapBySql()
	getstories.Field = "product, status, count(status) AS count"
	getstories.Where = map[string]interface{}{
		"Deleted": false,
		"Product": ids,
	}
	getstories.Group = "product, status"
	var getstories_result *protocol.MSG_PROJECT_product_getStoriesMapBySql_result
	HostConn.SendMsgWaitResultToDefault(getstories, &getstories_result)
	stories := map[int32]map[string]string{}
	for _, v := range getstories_result.List {
		id, _ := strconv.Atoi(v["product"])
		if stories[int32(id)] == nil {
			stories[int32(id)] = make(map[string]string)
		}
		stories[int32(id)][v["status"]] = v["count"]
	}
	var langstatus []string
	for _, kv := range data.Lang["story"]["statusList"].([]protocol.HtmlKeyValueStr) {
		if kv.Key != "" {
			langstatus = append(langstatus, kv.Key)
		}
	}
	for _, id := range ids {
		if _, ok := stories[id]; !ok {
			stories[id] = make(map[string]string)
			for _, key := range langstatus {
				stories[id][key] = "0"
			}
		} else {
			for _, key := range langstatus {
				if _, ok := stories[id][key]; !ok {
					stories[id][key] = "0"
				}
			}
		}
	}
	/*
	  plans = this->dao->select("product, count(*) AS count")
	      ->from(TABLE_PRODUCTPLAN)
	      ->where("deleted")->eq(0)
	      ->andWhere("product")->in(array_keys(products))
	      ->andWhere("end")->gt(helper::now())
	      ->groupBy("product")
	      ->fetchPairs()

	  releases = this->dao->select("product, count(*) AS count")
	      ->from(TABLE_RELEASE)
	      ->where("deleted")->eq(0)
	      ->andWhere("product")->in(array_keys(products))
	      ->groupBy("product")
	      ->fetchPairs()

	  bugs = this->dao->select("product,count(*) AS conut")
	      ->from(TABLE_BUG)
	      ->where("deleted")->eq(0)
	      ->andWhere("product")->in(array_keys(products))
	      ->groupBy("product")
	      ->fetchPairs()
	  unResolved = this->dao->select("product,count(*) AS count")
	      ->from(TABLE_BUG)
	      ->where("deleted")->eq(0)
	      ->andwhere("status")->eq("active")
	      ->andWhere("product")->in(array_keys(products))
	      ->groupBy("product")
	      ->fetchPairs()
	  assignToNull = this->dao->select("product,count(*) AS count")
	      ->from(TABLE_BUG)
	      ->where("deleted")->eq(0)
	      ->andwhere("assignedTo")->eq("")
	      ->andWhere("product")->in(array_keys(products))
	      ->groupBy("product")
	      ->fetchPairs()

	  stats = array()
	  foreach(products as key => product)
	  {
	      product->stories  = stories[product->id]
	      product->plans    = isset(plans[product->id])    ? plans[product->id]    : 0
	      product->releases = isset(releases[product->id]) ? releases[product->id] : 0

	      product->bugs         = isset(bugs[product->id]) ? bugs[product->id] : 0
	      product->unResolved   = isset(unResolved[product->id]) ? unResolved[product->id] : 0
	      product->assignToNull = isset(assignToNull[product->id]) ? assignToNull[product->id] : 0
	      stats[] = product
	  }

	  return stats*/
	result = make([]map[string]interface{}, len(products))
	for k, product := range products {
		tmp := make(map[string]interface{})
		tmp["Id"] = product.Id
		tmp["Name"] = product.Name
		tmp["Code"] = product.Code
		tmp["Line"] = product.Line
		tmp["Type"] = product.Type
		tmp["Status"] = product.Status
		tmp["Desc"] = product.Desc
		tmp["PO"] = product.PO
		tmp["QD"] = product.QD
		tmp["RD"] = product.RD
		tmp["Acl"] = product.Acl
		tmp["Whitelist"] = product.Whitelist
		tmp["CreatedBy"] = product.CreatedBy
		tmp["CreatedDate"] = product.CreatedDate
		tmp["Order"] = product.Order
		tmp["stories"] = stories[product.Id]
		tmp["isClickableKey"] = "MSG_PROJECT_product_cache_map_isClickable"
		result[k] = tmp
	}
	return
}
func product_getList(data *TemplateData, order func(a, b *protocol.MSG_PROJECT_product_cache) bool, status string, limit int, line int32, rootID int32) (result []*protocol.MSG_PROJECT_product_cache, err error) {
	var list []*protocol.MSG_PROJECT_product_cache
	if rootID == 0 {
		list, err = product_getAll(data)
		if err != nil {
			return
		}
	} else {
		list = []*protocol.MSG_PROJECT_product_cache{HostConn.GetProductById(rootID)}
	}

	for _, v := range list {
		if v.Deleted || (line > 0 && v.Line != line) || (data.User.Id != 1 && data.User.AclProducts[v.Id] == false) {
			continue
		}
		switch status {
		case "noclosed":
			if v.Status == "closed" {
				continue
			}
		case "involved":
			if v.PO != data.User.Id && v.QD != data.User.Id && v.RD != data.User.Id && v.CreatedBy != data.User.Id {
				continue
			}
		default:
			if status != "all" {
				if v.Status != status {
					continue
				}
			}
		}
		result = append(result, v)
	}
	protocol.Order_product(result, order)
	/*protocol.Order_product(result, func(a, b *protocol.MSG_PROJECT_product_cache) bool {
		if a.Order == b.Order {
			return a.Desc < b.Desc
		}
		return a.Order < b.Order
	})*/
	if limit > 0 && limit > len(result) {
		result = result[:limit]
	}

	return
}
func get_product_view(data *TemplateData) {
	id, _ := strconv.Atoi(data.ws.Query("product"))
	productID, branch, err := product_saveState(data, int32(id))
	err = product_setMenu(data, productID, branch, "")
	if err != nil {
		data.OutErr(err)
		return
	}
	list, err := product_getStats(data, "", "all", 0, productID)
	if err != nil {
		data.OutErr(err)
		return
	} else if len(list) == 0 {
		data.OutErr(errors.New(data.Lang["product"]["error"].(map[string]string)["NotFound"]))
		return
	}

	data.Data["product"] = list[0]
	data.Data["users"], err = user_getPairs("noletter")
	if err != nil {
		data.OutErr(err)
		return
	}
	lines, err := tree_getLinePairs(data)
	if err != nil {
		data.OutErr(err)
		return
	}
	data.Data["lines"] = append([]protocol.HtmlKeyValueStr{{"", ""}}, lines...)
	data.Data["blockHistory"] = true
	templateOut("product.view.html", data)
	return
}
func get_product_edit(data *TemplateData) {
	id, _ := strconv.Atoi(data.ws.Query("product"))
	productID, branch, err := product_saveState(data, int32(id))
	err = product_setMenu(data, productID, branch, "")
	if err != nil {
		data.OutErr(err)
		return
	}
	product := HostConn.GetProductById(productID)
	if product == nil {
		data.OutErr(errors.New(data.Lang["product"]["error"].(map[string]string)["NotFound"]))
		return
	}

	data.Data["product"] = product
	lines, err := tree_getLinePairs(data)
	if err != nil {
		data.OutErr(err)
		return
	}
	user, err := user_getPairs("nodeleted")
	if err != nil {
		data.OutErr(err)
		return
	}
	data.Data["poUsers"] = user
	data.Data["qdUsers"] = user
	data.Data["rdUsers"] = user
	data.Data["lines"] = append([]protocol.HtmlKeyValueStr{{"", ""}}, lines...)
	data.Data["groups"], _ = user_getGroupOptionMenu()
	templateOut("product.edit.html", data)
	return
}
func post_product_edit(data *TemplateData) {
	if !data.ajaxCheckPost() {
		return
	}
	id, _ := strconv.Atoi(data.ws.Query("product"))
	product := HostConn.GetProductById(int32(id))
	if product == nil {
		data.ajaxResult(false, data.Lang["product"]["error"].(map[string]string)["NotFound"])
		return
	}
	msg, err := data.GetMsg()
	if err != nil {
		data.ajaxResult(false, err.Error())
		return
	}
	out := protocol.GET_MSG_PROJECT_product_update()
	product.Id = product.Id
	product.Status = "normal"
	product.CreatedBy = data.User.Id

	for key, v := range data.ws.GetAllPost() {
		switch key {
		case "acl":
			product.Acl = v[0]
		case "whitelist":
			for _, sid := range v {
				id, _ := strconv.Atoi(sid)
				if id > 0 {
					product.Whitelist = append(product.Whitelist, int32(id))
				}
			}
		case "name":
			product.Name = v[0]
		case "code":
			product.Code = v[0]
		case "line":
			id, _ := strconv.Atoi(v[0])
			product.Line = int32(id)
		case "PO":
			id, _ := strconv.Atoi(v[0])
			product.PO = int32(id)
		case "QD":
			id, _ := strconv.Atoi(v[0])
			product.QD = int32(id)
		case "RD":
			id, _ := strconv.Atoi(v[0])
			product.RD = int32(id)
		case "type":
			product.Type = v[0]
		case "desc":
			product.Desc = v[0]

		}
	}
	m, _ := libraries.Preg_match_result(`<img src="/file/tmpimg\?fileID=(\d+)&amp;t=([^"]+)" alt="" \/>`, product.Desc, -1)
	var uploaderr error
	var newimgids []int64
	for _, match := range m {
		b, ok := file_getTempFile(match[1])
		if ok {
			upload := protocol.GET_MSG_FILE_upload()
			upload.AddBy = data.User.Id
			upload.Data = b
			upload.Name = time.Now().Format("20060102") + "_" + match[1] + "." + match[2]
			var res *protocol.MSG_FILE_upload_result
			uploaderr = msg.SendMsgWaitResult(0, upload, &res)
			if uploaderr == nil {
				newimgids = append(newimgids, res.FileID)
				product.Desc = strings.ReplaceAll(product.Desc, match[0], `<img src="/file/read?fileID=`+strconv.FormatInt(res.FileID, 10)+` alt="" />`)
			}
			res.Put()
			if uploaderr != nil {
				deleteimg := protocol.GET_MSG_FILE_DeleteByID()
				for _, id := range newimgids {
					deleteimg.FileID = id
					msg.SendMsg(0, deleteimg)
				}
				deleteimg.Put()
				data.ajaxResult(false, map[string]string{"desc": fmt.Sprintf(data.Lang["file"]["imguploadFail"].(string), uploaderr)})
				return
			}
		} else {
			product.Desc = strings.ReplaceAll(product.Desc, match[0], "")
		}

	}
	product.Desc = libraries.Html2bbcode(product.Desc)
	defer func() {
		if err != nil { //以下使用err来判断图片删除
			deleteimg := protocol.GET_MSG_FILE_DeleteByID()
			for _, id := range newimgids {
				deleteimg.FileID = id
				msg.SendMsg(0, deleteimg)
			}
			deleteimg.Put()
		}
	}()
	out.Data = product
	if err = msg.SendMsgWaitResult(0, out, nil); err != nil {
		data.ajaxResult(false, err.Error())
		return
	}
	locate := createLink("product", "view", []string{"productID=", strconv.Itoa(int(product.Id))})
	data.ajaxResult(true, data.Lang["common"]["saveSuccess"], locate)
	out.Put()
	return
}
func get_product_ajaxGetPlans(data *TemplateData) {
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	branch, _ := strconv.Atoi(data.ws.Query("branch"))
	planID, _ := strconv.Atoi(data.ws.Query("planID"))
	fieldID, _ := strconv.Atoi(data.ws.Query("fieldID"))
	needCreate := data.ws.Query("needCreate") == "true"
	plans, err := productplan_getPairs(data, int32(productID), int32(branch), data.ws.Query("expired"))
	if err != nil {
		data.ws.WriteString(js.Alert(err.Error()))
		return
	}
	field := "plan"
	if fieldID > 0 {
		field = "plans[" + strconv.Itoa(fieldID) + "]"
	}
	buf := bufpool.Get().(*libraries.MsgBuffer)
	buf.WriteString(html_select(field, plans, planID, "class='form-control chosen'"))

	if len(plans) == 1 && needCreate {
		buf.WriteString("<div class='input-group-btn'>")
		buf.WriteString(html_a(createLink("productplan", "create", []interface{}{"productID=", productID, "&branch=", branch, true}), "<i class='icon icon-plus'></i>", "", "class='btn btn-icon' data-toggle='modal' data-type='iframe' data-width='95%' title='"+data.Lang["productplan"]["create"].(string)+"'"))
		buf.WriteString("</div><div class='input-group-btn'>")
		buf.WriteString(html_a("javascript:void(0)", "<i class='icon icon-refresh'></i>", "", "class='btn btn-icon refresh' data-toggle='tooltip' title='"+data.Lang["common"]["refresh"].(string)+"' onclick='loadProductPlans("+strconv.Itoa(productID)+")'"))
		buf.WriteString("</div>")
	}
	data.ws.WriteString(buf.String())
	buf.Reset()
	bufpool.Put(buf)
}
