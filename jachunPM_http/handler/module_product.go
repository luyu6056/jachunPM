package handler

import (
	"errors"
	"fmt"
	"html/template"
	"libraries"
	"protocol"
	"strconv"
	"strings"
	"time"

	"github.com/luyu6056/gnet"
)

func init() {
	httpHandlerMap["GET"]["/product/index"] = get_product_index
	httpHandlerMap["GET"]["/product/create"] = get_product_create
	httpHandlerMap["POST"]["/product/create"] = post_product_create
	httpHandlerMap["GET"]["/product/browse"] = get_product_browse
	httpHandlerMap["GET"]["/product/ajaxGetDropMenu"] = get_product_ajaxGetDropMenu
}
func get_product_index(data *TemplateData) (action gnet.Action) {

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
func get_product_create(data *TemplateData) (action gnet.Action) {
	//$rootID = key($this->products);
	//if($this->session->product) $rootID = $this->session->product;
	//$this->product->setMenu($this->products, $rootID);
	data.Data["groups"], _ = user_getGroupOptionMenu()
	msg, err := HostConn.GetMsg()
	if err != nil {
		data.OutErr(err)
		return
	}
	getuser := protocol.GET_MSG_USER_getPairs()
	getuser.Params = "nodeleted|pofirst|noclosed"
	var res *protocol.MSG_USER_getPairs_result
	err = msg.SendMsgWaitResult(0, getuser, &res)
	if err != nil {
		data.OutErr(err)
		return
	}
	data.Data["poUsers"] = res.List
	getuser.Params = "nodeleted|qdfirst|noclosed"
	var res1 *protocol.MSG_USER_getPairs_result
	err = msg.SendMsgWaitResult(0, getuser, &res1)
	if err != nil {
		data.OutErr(err)
		return
	}
	data.Data["qdUsers"] = res1.List
	getuser.Params = "nodeleted|devfirst|noclosed"
	var res2 *protocol.MSG_USER_getPairs_result
	err = msg.SendMsgWaitResult(0, getuser, &res2)
	if err != nil {
		data.OutErr(err)
		return
	}
	data.Data["rdUsers"] = res2.List
	getuser.Put()
	var productTypeList []protocol.HtmlKeyValueStr
	for _, v := range data.Lang["product"]["typeList"].([]protocol.HtmlKeyValueStr) {
		tip, _ := data.Lang["product"]["typeTips"].(map[string]string)[v.Key]
		productTypeList = append(productTypeList, protocol.HtmlKeyValueStr{v.Key, v.Value + tip})

	}
	data.Data["productTypeList"] = productTypeList
	getLinePairs := protocol.GET_MSG_PROJECT_tree_getLinePairs()
	var res3 *protocol.MSG_PROJECT_tree_getLinePairs_result
	err = msg.SendMsgWaitResult(0, getLinePairs, &res3)
	if err != nil {
		data.OutErr(err)
		return
	}
	res3.List = append([]protocol.HtmlKeyValueStr{{"", ""}}, res3.List...)
	data.Data["lines"] = res3.List
	templateOut("product.create.html", data)
	res.Put()
	res1.Put()
	res2.Put()
	res3.Put()
	return
}
func post_product_create(data *TemplateData) (action gnet.Action) {
	if !data.ajaxCheckPost() {
		return
	}
	msg, err := HostConn.GetMsg()
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
	m, _ := libraries.Preg_match_result(`<img src="/file/tmpimg\?fileID=(\d+)&amp;t=([^"]+)" alt="" \/>`, insert.Desc, -1)
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
				insert.Desc = strings.ReplaceAll(insert.Desc, match[0], `<img src="/file/read?fileID=`+strconv.FormatInt(res.FileID, 10)+` alt="" />`)
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
			insert.Desc = strings.ReplaceAll(insert.Desc, match[0], "")
		}

	}
	insert.Desc = libraries.Html2bbcode(insert.Desc)
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
	out.Put()
	res.Put()
	return
}
func get_product_browse(data *TemplateData) (action gnet.Action) {
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
	msg, err := HostConn.GetMsg()
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
func product_getAll() (result []*protocol.MSG_PROJECT_product_cache, err error) {
	res, err := HostConn.CacheGetPath(protocol.ProjectServerNo, protocol.PATH_PROJECT_PRODUCT_CACHE)
	if err != nil {
		return
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

	return
}
func product_getPairs(data *TemplateData, mode ...string) (res []protocol.HtmlKeyValueStr, err error) {
	list, err := product_getAll()
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
	selectHtml := func() {
		//public function select($products, $productID, $currentModule, $currentMethod, $extra = '', $branch = 0, $module = 0, $moduleType = '')
		//$isMobile = $this->app->viewType == 'mhtml';
		data.ws.SetCookie("lastProduct", productIDStr, protocol.SessionKeepLoginExpires)

		currentProduct := HostConn.GetProductById(int32(productID))
		data.ws.Session().Set("currentProductType", currentProduct.Type)
		buf.WriteString(`<div class='btn-group angle-btn'><div class='btn-group'><button data-toggle='dropdown' type='button' class='btn btn-limit' id='currentItem' title='`)
		buf.WriteString(currentProduct.Name)
		buf.WriteString("'>")
		buf.WriteString(currentProduct.Name)
		buf.WriteString(`<span class='caret'></span></button><div id='dropMenu' class='dropdown-menu search-list' data-ride='searchList' data-url='`)
		buf.WriteString(createLink("product", "ajaxGetDropMenu", []interface{}{"objectID=", productID, "&module=", currentModule, "&method=", currentMethod, "&extra=", extra}))
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
				buf.WriteString(createLink("branch", "ajaxGetDropMenu", []interface{}{"objectID=", productID, "&module=", currentModule, "&method=", currentMethod, "&extra=", extra}))
				buf.WriteString(`><div class='input-control search-box has-icon-left has-icon-right search-example'><input type='search' class='form-control search-input' /><label class='input-control-icon-left search-icon'><i class='icon icon-search'></i></label><a class='input-control-icon-right search-clear-btn'><i class='icon icon-close icon-sm'></i></a></div></div></div>`)
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
	selectHtml()

	data.Data["modulePageNav"] = template.HTML(buf.String())
	if data.App["menuReplace"] == nil {
		data.App["menuReplace"] = make(map[string]string)
	}
	data.App["menuReplace"].(map[string]string)["productID"] = strconv.Itoa(int(productID))
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
func get_product_ajaxGetDropMenu(data *TemplateData) (action gnet.Action) {
	method := data.ws.Query("method")
	module := data.ws.Query("module")
	productID, _ := strconv.Atoi(data.ws.Query("objectID"))
	link := product_getProductLink(module, method, data.ws.Query("extra"), 0)
	data.Data["productID"] = productID
	data.Data["module"] = module
	data.Data["method"] = method
	data.Data["extra"] = data.ws.Query("extra")
	products, err := product_getAll()
	if err != nil {
		data.OutErr(err)
		return
	}
	lines, err := tree_getLinePairs()
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
func product_getProductLink(module, method, extra string, branch int32) string {

	switch module {
	case "product", "roadmap", "bug", "testcase", "testtask", "story", "qa", "testsuite", "testreport", "build":
		switch {
		case module == "product" && method == "project":
			if branch > 0 {
				return createLink(module, method, "status=all&productID=%d&branch=%d")
			} else {
				return createLink(module, method, "status=all&productID=%d")
			}
		case (module == "product" && (method == "dynamic" || method == "doc" || method == "view")):
			return createLink(module, method, "productID=%d")

		case (module == "qa" && method == "index"):
			if branch > 0 {
				return createLink("bug", "browse", "productID=%d&branch=%d")
			} else {
				return createLink("bug", "browse", "productID=%d")
			}

		case (module == "product" && (method == "browse" || method == "index" || method == "all")):

			if branch > 0 {
				return createLink(module, "browse", "productID=%d&branch=%d")
			} else {
				return createLink(module, "browse", "productID=%d")
			}

		default:
			if branch > 0 {
				return createLink(module, method, "productID=%d&branch=%d")
			} else {
				return createLink(module, method, "productID=%d")
			}

		}

	case "productplan", "release":

		if method != "browse" && method != "create" {
			method = "browse"
		}
		if branch > 0 {
			return createLink(module, method, "productID=%d&branch=%d")
		} else {
			return createLink(module, method, "productID=%d")
		}

	case "tree":
		if branch > 0 {
			return createLink(module, method, "productID=%d&type="+extra+"&currentModuleID=0&branch=%d")
		} else {
			return createLink(module, method, "productID=%d&type="+extra+"&currentModuleID=0")
		}

	case "branch":

		return createLink(module, method, "productID=%d")

	case "doc":
		return createLink("doc", "objectLibs", "type=product&objectID=%d&from=product")

	}

	return ""
}
