package handler

import (
	"html/template"
	"jachunPM_http/js"
	"libraries"
	"protocol"
	"sort"
	"strconv"
	"strings"

	"github.com/luyu6056/gnet"
)

func init() {

	httpHandlerMap["GET"]["/tree/browse"] = get_tree_browse
	httpHandlerMap["POST"]["/tree/manageChild"] = post_tree_manageChild
	httpHandlerMap["POST"]["/tree/updateOrder"] = post_tree_updateOrder
	httpHandlerMap["GET"]["/tree/edit"] = get_tree_edit
	httpHandlerMap["POST"]["/tree/edit"] = post_tree_edit
	httpHandlerMap["GET"]["/tree/delete"] = get_tree_delete
}
func get_tree_browse(data *TemplateData) (action gnet.Action) {
	rootID, _ := strconv.Atoi(data.ws.Query("rootID"))
	if rootID == 0 {
		rootID, _ = strconv.Atoi(data.ws.Query("productID"))
	}
	currentModuleID, _ := strconv.Atoi(data.ws.Query("currentModuleID"))
	branch, _ := strconv.Atoi(data.ws.Query("branch"))
	viewType := data.ws.Query("view")
	if viewType == "" {
		viewType = data.ws.Query("type")
	}

	msg, err := HostConn.GetMsg()
	defer func() {
		if err != nil {
			data.OutErr(err)
		}
	}()
	if err != nil {
		return
	}

	if viewType != "" {
		data.Data["manageChild"] = "manage" + strings.ToUpper(viewType[:1]) + viewType[1:] + "Child"
	}
	if viewType == "line" {
		data.Data["name"] = data.Lang["tree"]["line"]
	} else if viewType == "doc" {
		data.Data["name"] = data.Lang["tree"]["cate"]
	} else {
		data.Data["name"] = data.Lang["tree"]["name"]
	}

	data.Data["hasBranch"] = false
	if strings.Contains("story|bug|case", viewType) {
		product := HostConn.GetProductById(int32(rootID))
		if product == nil {
			data.ws.Redirect(createLink("product", "create", nil))
			return
		}
		if product.Type != "normal" {
			data.Data["hasBranch"] = true
			branches := []protocol.HtmlKeyValueStr{{"all", data.Lang["branch"]["all"].(string) + data.Lang["product"]["branchName"].(map[string]string)[product.Type]}}
			if currentModuleID > 0 {
				for _, b := range product.Branchs {
					if b.Id == int32(branch) {
						branches = []protocol.HtmlKeyValueStr{{strconv.Itoa(int(b.Id)), b.Name}}
					}
				}
			} else {
				for _, b := range product.Branchs {
					branches = []protocol.HtmlKeyValueStr{{strconv.Itoa(int(b.Id)), b.Name}}
				}
			}
			data.Data["branches"] = branches
		}
		data.Data["root"] = product
	} else if strings.Contains(viewType, "doc") {
		viewType = "doc"
		//data.Data["root = $this->doc->getLibById($rootID);
	} else if strings.Contains(viewType, "caselib") {
		getTestsuite := protocol.GET_MSG_TEST_testsuite_getById()
		getTestsuite.Id = int32(rootID)
		var getTestsuiteResult *protocol.MSG_TEST_testsuite_getById_result
		err = msg.SendMsgWaitResult(0, getTestsuite, &getTestsuiteResult)
		if err != nil {
			return
		}
		data.Data["root"] = getTestsuiteResult.Info
		getTestsuite.Put()
	}
	switch viewType {
	case "story":
		templateLock.Lock()
		data.Lang["menugroup"]["tree"] = "product"
		data.Lang["tree"]["menu"] = data.Lang["product"]["menu"]
		templateLock.Unlock()
		err = product_setMenu(data, int32(rootID), int32(branch), "story")
		if err != nil {
			return
		}
		products, e := product_getPairs(data)
		if e != nil {
			err = e
			return
		}
		for i := len(products) - 1; i >= 0; i-- {
			p := products[i]
			if p.Key == strconv.Itoa(rootID) {
				products = append(products[:i], products[i+1:]...)
				break
			}
		}

		currentProduct := products[0].Key

		data.Data["allProduct"] = products
		data.Data["currentProduct"] = currentProduct
		id, _ := strconv.Atoi(currentProduct)
		data.Data["productModules"], err = tree_getOptionMenu(data, int32(id), "story", 0, 0)
		if err != nil {
			return
		}
		data.Data["title"] = data.Lang["tree"]["manageProduct"]
		/*case "bug":

		     $this->loadModel('bug')->setMenu($this->product->getPairs(), $rootID);
		     $this->lang->tree->menu      = $this->lang->bug->menu;
		     $this->lang->tree->menuOrder = $this->lang->bug->menuOrder;
		     if($this->config->global->flow == 'onlyTest') $this->lang->set('menugroup.tree', 'bug');
		     if($this->config->global->flow != 'onlyTest') $this->lang->set('menugroup.tree', 'qa');

		     $title      = $this->lang->tree->manageBug;
		     $position[] = html::a($this->createLink('bug', 'browse', "product=$rootID"), $product->name);
		     $position[] = $this->lang->tree->manageBug;
		}
		 elseif($viewType == 'case')
		 {
		     $this->loadModel('testcase')->setMenu($this->product->getPairs(), $rootID);
		     $this->lang->tree->menu      = $this->lang->testcase->menu;
		     $this->lang->tree->menuOrder = $this->lang->testcase->menuOrder;
		     if($this->config->global->flow == 'onlyTest') $this->lang->set('menugroup.tree', 'testcase');
		     if($this->config->global->flow != 'onlyTest') $this->lang->set('menugroup.tree', 'qa');

		     $title      = $this->lang->tree->manageCase;
		     $position[] = html::a($this->createLink('testcase', 'browse', "product=$rootID"), $product->name);
		     $position[] = $this->lang->tree->manageCase;
		 }
		 elseif($viewType == 'caselib')
		 {
		     $this->testsuite->setLibMenu($this->testsuite->getLibraries(), $rootID);
		     $this->lang->tree->menu      = $this->lang->testsuite->menu;
		     $this->lang->tree->menuOrder = $this->lang->testsuite->menuOrder;
		     $this->lang->set('menugroup.tree', 'qa');

		     $title      = $this->lang->tree->manageCaseLib;
		     $position[] = html::a($this->createLink('testsuite', 'library', "libID=$rootID"), $lib->name);
		     $position[] = $this->lang->tree->manageCaseLib;
		 }
		 elseif(strpos($viewType, 'doc') !== false)
		 {
		     $type = $lib->product ? 'product' : ($lib->project ? 'project' : 'custom');
		     $this->doc->setMenu($type, $rootID, $currentModuleID);
		     $this->lang->tree->menu      = $this->lang->doc->menu;
		     $this->lang->tree->menuOrder = $this->lang->doc->menuOrder;
		     $this->lang->set('menugroup.tree', 'doc');

		     $title      = $this->lang->tree->manageCustomDoc;
		     $position[] = html::a($this->createLink('doc', 'browse', "libID=$rootID"), $lib->name);
		     $position[] = $this->lang->tree->manageCustomDoc;
		 }*/
	case "line":
		product_setMenu(data, int32(rootID), int32(branch), "line")
		products, e := product_getPairs(data)
		if e != nil {
			data.OutErr(e)
			return
		}
		for i := len(products) - 1; i >= 0; i-- {
			if products[i].Key == strconv.Itoa(int(rootID)) {
				products = append(products[:i], products[i+1:]...)
				break
			}
		}
		if len(products) > 0 {
			currentProduct, _ := strconv.Atoi(products[0].Key)
			data.Data["allProduct"] = products
			data.Data["currentProduct"] = products[0].Key
			data.Data["productModules"], err = tree_getOptionMenu(data, int32(currentProduct), "line", 0, 0)

		}

		data.Data["title"] = data.Lang["tree"]["manageLine"]

	}
	if currentModuleID > 0 {
		getParents := protocol.GET_MSG_PROJECT_tree_getParents()
		getParents.ModuleID = int32(currentModuleID)
		var getParentsResult *protocol.MSG_PROJECT_tree_getParents_result
		err = msg.SendMsgWaitResult(0, getParents, &getParentsResult)
		if err != nil {
			return
		}
		data.Data["parentModules"] = getParentsResult.List
		getParents.Put()
	}
	data.Data["sons"], err = tree_getSons(data, int32(rootID), int32(currentModuleID), viewType, int32(branch))
	if err != nil {
		return
	}
	data.Data["rootID"] = rootID
	data.Data["tree"], err = tree_getProductStructure(data, int32(rootID), viewType)
	data.Data["currentModuleID"] = currentModuleID
	data.Data["viewType"] = viewType
	templateOut("tree.browse.html", data)
	return
}
func tree_getOptionMenu(data *TemplateData, rootID int32, typ string, startModule, branch int32) ([]protocol.HtmlKeyValueStr, error) {

	if typ == "line" {
		rootID = 0
	}
	branches := []protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{strconv.Itoa(int(branch)), ""}}
	if strings.Contains("story|bug|case", typ) {
		product := HostConn.GetProductById(rootID)
		if product != nil && product.Type != "normal" {
			branches = []protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"null", ""}}
			branches = append(branches, branch_getPairs(data, rootID, product, "noempty")...)
			if branch > 0 {
				for i := len(branches) - 1; i >= 0; i++ {
					kv := branches[i]
					if kv.Key == "null" && kv.Key == strconv.Itoa(int(branch)) {
						continue
					}
					branches = append(branches[:i], branches[i+1:]...)
				}
			}
		}
	}

	treeMenu := map[int32]string{}
	for _, branchkv := range branches {
		branchID, _ := strconv.Atoi(branchkv.Key)
		list, err := tree_buildMenuQuery(data, rootID, typ, startModule, int32(branchID))
		if err != nil {
			return nil, err
		}
		var modules = make(map[int32]*protocol.MSG_PROJECT_tree_cache, len(list))
		for _, module := range list {
			modules[module.Id] = module
		}
		param := "/"
		if branch > 0 {
			param = "/branch/"
		}
		for _, module := range modules {
			tree_buildTreeArray(treeMenu, modules, module, param)
		}
	}

	var topMenu []string
	if len(treeMenu) > 0 {
		var ids []int
		for key := range treeMenu {
			ids = append(ids, int(key))
		}
		sort.Ints(ids)
		topMenu = strings.Split(treeMenu[int32(ids[0])], "\n")
	}
	var lastMenu []protocol.HtmlKeyValueStr
	if typ == "bug" || typ == "story" {
		lastMenu = []protocol.HtmlKeyValueStr{{"", ""}}
	} else {
		lastMenu = []protocol.HtmlKeyValueStr{{"", "/"}}
	}

	for _, str := range topMenu {
		menu := strings.Split(str, "|")
		if len(menu) == 2 {
			lastMenu = append(lastMenu, protocol.HtmlKeyValueStr{menu[1], menu[0]})
		}

	}
	return lastMenu, nil
}
func tree_buildMenuQuery(data *TemplateData, rootID int32, typ string, startModule, branch int32) ([]*protocol.MSG_PROJECT_tree_cache, error) {

	list, err := tree_getAllcache(data)
	if err != nil {
		return nil, err
	}
	var startModulePath []int32
	if startModule > 0 {
		module := HostConn.GetTreeById(startModule)
		if module != nil {
			startModulePath = module.Path
		}
	}

	var res []*protocol.MSG_PROJECT_tree_cache
out:
	for _, m := range list {

		if m.Deleted || m.Root != rootID || m.Branch != branch || m.Type != typ {
			continue
		}
		if startModulePath != nil {
			for _, path1 := range startModulePath {
				var find bool
				for _, path2 := range m.Path {
					if path1 == path2 {
						find = true
						break
					}
				}
				if !find {
					continue out
				}
			}
		}

		if typ != "story" {
			if typ == "task" {
				if m.Type != "task" {

					continue
				}
			} else {
				if m.Type != "story" && m.Type != typ {

					continue
				}
			}
		} else {
			if m.Type != "story" {

				continue
			}
		}
		res = append(res, m)
	}
	protocol.Order_tree(res, func(a, b *protocol.MSG_PROJECT_tree_cache) bool {
		if a.Grade == b.Grade {
			if a.Type == b.Type {
				if a.Order == b.Order {
					return a.Id < b.Id
				}
				return a.Order < b.Order
			}
			return a.Type > b.Type
		}
		return a.Grade > b.Grade

	})
	return res, nil
}
func tree_getAllcache(data *TemplateData) (result []*protocol.MSG_PROJECT_tree_cache, err error) {
	if data.Data["tree_getAllcache"] == nil {
		res, err := HostConn.CacheGetPath(protocol.ProjectServerNo, protocol.PATH_PROJECT_TREE_CACHE)
		if err != nil {
			return nil, err
		}

		buf := bufpool.Get().(*libraries.MsgBuffer)
		for _, b := range res {
			buf.Reset()
			buf.Write(b)
			if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_PROJECT_tree_cache); ok {
				result = append(result, v)
			}
		}
		buf.Reset()
		bufpool.Put(buf)
		data.Data["tree_getAllcache"] = result
	}
	return data.Data["tree_getAllcache"].([]*protocol.MSG_PROJECT_tree_cache), nil
}
func tree_buildTreeArray(treeMenu map[int32]string, modules map[int32]*protocol.MSG_PROJECT_tree_cache, module *protocol.MSG_PROJECT_tree_cache, moduleName string) {
	moduleNames := []string{moduleName}
	for _, path := range module.Path {
		if v, ok := modules[path]; ok {
			moduleNames = append(moduleNames, v.Name)
		}

	}
	moduleName = strings.Join(moduleNames, "/") + "|$module->id\n"
	treeMenu[module.Parent] += moduleName + treeMenu[module.Id]
}
func tree_getProductStructure(data *TemplateData, rootID int32, viewType string) (fullTrees []map[string]interface{}, err error) {
	if viewType == "line" {
		rootID = 0
	}
	var branches []protocol.HtmlKeyValueStr
	var product *protocol.MSG_PROJECT_product_cache
	if rootID > 0 {
		product = HostConn.GetProductById(rootID)
		if strings.Contains("story|bug|case", viewType) {
			if product.Type != "normal" {
				branches = branch_getPairs(data, 0, product, "noempty")
			}
		}
	}

	list, err := tree_buildMenuQuery(data, rootID, viewType, 0, 0)
	if err != nil {
		return nil, err
	}
	fullTrees = tree_getDataStructure(data, list, viewType, nil)

	if len(branches) > 0 {
		var branchTrees []map[string]interface{}
		for _, branchkv := range branches {
			branchID, _ := strconv.Atoi(branchkv.Key)
			list, err := tree_buildMenuQuery(data, rootID, viewType, 0, int32(branchID))
			if err != nil {
				return nil, err
			}

			branchTrees = append(branchTrees, map[string]interface{}{
				"branch":   branchID,
				"id":       0,
				"name":     branchkv.Value,
				"root":     rootID,
				"actions":  map[string]bool{"sort": true, "add": false, "edit": false, "delete": false},
				"nodeType": "branch",
				"type":     "branch",
				"children": tree_getDataStructure(data, list, viewType, nil),
			})

		}
		fullTrees = append(fullTrees, map[string]interface{}{
			"name":     data.Lang["product"]["branchName"].(map[string]string)[product.Type],
			"root":     rootID,
			"type":     "branch",
			"actions":  false,
			"children": branchTrees,
		})
	}
	return
}
func tree_getDataStructure(data *TemplateData, list []*protocol.MSG_PROJECT_tree_cache, viewType string, keepModules []int32) []map[string]interface{} {
	parent := map[int32]map[string]interface{}{}
out:
	for _, tree := range list {
		for _, id := range keepModules {
			if id == tree.Id {
				continue out
			}
		}
		module := map[string]interface{}{
			"id":     tree.Id,
			"name":   tree.Name,
			"parent": tree.Parent,
			"root":   tree.Root,
			"type":   tree.Type,
			"branch": tree.Branch,
		}
		/* Ignore useless module for task. */
		if v, ok := parent[tree.Id]; ok {

			module["children"] = v["children"]
			delete(parent, tree.Id)
		}
		if parent[tree.Parent] == nil {
			parent[tree.Parent] = map[string]interface{}{
				"children": []map[string]interface{}{},
			}
		}
		parent[tree.Parent]["children"] = append(parent[tree.Parent]["children"].([]map[string]interface{}), module)
	}
	if viewType == "task" {
		alltree, _ := tree_getAllcache(data)
		for _, tree := range alltree {
			if _, ok := parent[tree.Id]; ok {
				parent[tree.Id]["type"] = tree.Type
			}
		}
	}
	var tree []map[string]interface{}
	for _, module := range parent {
		for _, children := range module["children"].([]map[string]interface{}) {
			if viewType == "task" {
				if v, ok := parent[children["parent"].(int32)]; ok && v["type"] != nil && v["type"].(string) != "task" {
					continue
				}
			}
			if children["parent"].(int32) != 0 {
				continue
			}
			tree = append(tree, children)
		}
	}
	return tree
}
func tree_getLinePairs(data *TemplateData) (res []protocol.HtmlKeyValueStr, err error) {
	list, err := tree_getAllcache(data)
	if err != nil {
		return nil, err
	}
	protocol.Order_tree(list, func(a, b *protocol.MSG_PROJECT_tree_cache) bool {
		return a.Id < b.Id
	})
	for _, v := range list {
		if v.Deleted || v.Type != "line" {
			continue
		}
		res = append(res, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Name})
	}
	return
}
func post_tree_manageChild(data *TemplateData) (action gnet.Action) {
	if !data.ajaxCheckPost() {
		return
	}
	rootID, _ := strconv.Atoi(data.ws.Query("root"))
	viewType := data.ws.Query("viewType")
	if viewType == "line" {
		rootID = 0
	}
	out := protocol.GET_MSG_PROJECT_tree_manageChild()
	out.RootID = int32(rootID)
	out.ViewType = viewType
	parentModuleID, _ := strconv.Atoi(data.ws.Query("parentModuleID"))
	out.ParentModuleID = int32(parentModuleID)
	shorts := data.ws.PostSlice("shorts")
	branchs := data.ws.PostSlice("branch")
	for k, name := range data.ws.PostSlice("modules") {
		if name == "" {
			continue
		}
		tmp := protocol.GET_MSG_PROJECT_tree_cache()
		tmp.Name = name
		if len(shorts) > k {
			tmp.Short = shorts[k]
		}
		if len(branchs) > k {
			b, _ := strconv.Atoi(branchs[k])
			tmp.Branch = int32(b)
		}
		out.Modules = append(out.Modules, tmp)
	}

	var result *protocol.MSG_PROJECT_tree_manageChild_result
	err := HostConn.SendMsgWaitResultToDefault(out, &result)
	if err != nil {
		data.ajaxResult(false, err.Error())
		return
	}
	if result.Name != "" {
		data.ajaxResult(false, map[string]string{result.Name: data.Lang["tree"]["error"].(map[string]string)[result.Result.String()]})
		return
	}
	out.Put()
	result.Put()
	data.ajaxResult(true, data.Lang["common"]["saveSuccess"], "reload")
	return
}
func tree_getSons(data *TemplateData, rootID int32, moduleID int32, viewType string, branch int32) (result []*protocol.MSG_PROJECT_tree_cache, err error) {
	if viewType == "line" {
		rootID = 0
	}
	list, err := tree_getAllcache(data)
	if err != nil {
		return nil, err

	}
	for _, v := range list {
		if v.Root != rootID || v.Parent != moduleID || v.Type != viewType || v.Branch != branch || v.Deleted {
			continue
		}
		result = append(result, v)
	}

	protocol.Order_tree(result, func(a, b *protocol.MSG_PROJECT_tree_cache) bool {
		if a.Order == b.Order {
			return a.Id < b.Id
		}
		return a.Order < b.Order
	})
	return
}
func tree_getTreeMenu(data *TemplateData, rootID int32, viewType string, startModule int32, callback func(data *TemplateData, viewType string, module *protocol.MSG_PROJECT_tree_cache, extra map[string]interface{}) (string, error), extra map[string]interface{}, branch int32) (template.HTML, error) {

	if viewType == "line" {
		rootID = 0
	}
	branches := []protocol.HtmlKeyValueStr{{strconv.Itoa(int(branch)), ""}}
	/*if(branch>0){
	  	branchinfo:=
	      branchName = this->loadModel("branch")->getById(branch)
	      branches   = array("null" => "", branch => branchName)
	      extra      = array("rootID" => rootID, "branch" => branch)
	  }*/
	//manage := reflect2.PtrOf(callback) == reflect2.PtrOf(tree_createManageLink)
	if rootID > 0 {
		product := HostConn.GetProductById(rootID)
		if strings.Contains("story|bug|case", viewType) && branch == 0 {
			if product.Type != "normal" {
				branches = append([]protocol.HtmlKeyValueStr{{"null", ""}}, branch_getPairs(data, rootID, product, "noempty")...)
			}
		}
	}

	buf := bufpool.Get().(*libraries.MsgBuffer)
	defer func() {
		buf.Reset()
		bufpool.Put(buf)
	}()

	firstBranch := true
	for _, kv := range branches {
		branchID, _ := strconv.Atoi(kv.Key)
		branch := kv.Value
		list, err := tree_buildMenuQuery(data, rootID, viewType, startModule, int32(branchID))
		if err != nil {
			return template.HTML(""), err
		}

		treeMenuMap, err := tree_buildTree(data, list, viewType, callback, extra, kv.Key)
		if err != nil {
			return template.HTML(""), err
		}
		if extra != nil && len(treeMenuMap) == 0 {
			continue
		}
		var ids []int
		for id := range treeMenuMap {
			ids = append(ids, int(id))
		}
		sort.Ints(ids)
		var treeMenu []string
		for _, id := range ids {
			treeMenu = append(treeMenu, treeMenuMap[int32(id)])
		}
		//ksort(treeMenu)
		if branchID > 0 && branch > "0" {
			/*linkHtml = (type == "case" and !empty(extra)) ? "<a>" . branch . "</a>" : this->createBranchLink(type, rootID, branchID, branch)
			  linkHtml = manage ? html::a(inlink("browse", "root=$rootID&viewType=$type&currentModuleID=0&branch=$branchID"), branch) : linkHtml
			  if(type == "story" || type == "bug") linkHtml = "<a>" . branch . "</a>"
			  if(firstBranch and product->type != "normal")
			  {
			      linkHtml = "<a>" . this->lang->product->branchName[product->type] . "</a><ul><li>" . linkHtml
			      firstBranch = false
			  }
			  lastMenu .= "<li>$linkHtml<ul>" . @array_shift(treeMenu) . "</ul></li>\n"*/
		} else {
			buf.WriteString(treeMenu[0])

		}
	}

	if !firstBranch {
		buf.WriteString("</li></ul>")
	}
	if buf.Len() > 0 {
		lastMenu := buf.String()
		buf.Reset()
		buf.WriteString("<ul id='modules' class='tree' data-ride='tree' data-name='tree-")
		buf.WriteString(viewType)
		buf.WriteString("'>")
		buf.WriteString(lastMenu)
		buf.WriteString("</ul>\n")

	}
	return template.HTML(buf.String()), nil
}
func tree_buildTree(data *TemplateData, list []*protocol.MSG_PROJECT_tree_cache, viewType string, callback func(data *TemplateData, viewType string, module *protocol.MSG_PROJECT_tree_cache, extra map[string]interface{}) (string, error), extra map[string]interface{}, branch string) (treeMenuMap map[int32]string, err error) {
	treeMenuMap = make(map[int32]string)
	for _, module := range list {
		if (extra["rootID"] != nil && extra["branch"] != nil && branch == "null") || (viewType == "case" && extra["taskID"] != nil) {
			var objects map[int32]bool

			if extra["taskID"] == nil {
				table := data.Config["objectTables"][viewType]
				libraries.ReleaseLog("buildTree需要%s表格的module", table)
				//objects = this->dao->select("module")->from(table)->where("product")->eq((int)extra["rootID"])->andWhere("branch")->eq((int)extra["branch"])->fetchAll("module")
			} else {
				libraries.ReleaseLog("buildTree需要处理task")
				/*objects = this->dao->select("t1.*,t2.module")->from(TABLE_TESTRUN)->alias("t1")
				  ->leftJoin(TABLE_CASE)->alias("t2")->on("t1.case = t2.id")
				  ->where("t1.task")->eq((int)extra)
				  ->fetchAll("module")*/
			}
			list, err := tree_getAllcache(data)
			if err != nil {
				return nil, err
			}
			childModules := make(map[int32]bool)
			for _, tree := range list {
				if viewType == "story" {
					if tree.Type != "stroy" || tree.Root != module.Root || len(tree.Path) == 0 {
						continue
					}

				} else {
					if (tree.Type != "stroy" && tree.Type != viewType) || tree.Root != module.Root || len(tree.Path) == 0 {
						continue
					}
				}
				isChild := true
				for k, v := range module.Path {
					if tree.Path[k] != v {
						isChild = false
						break
					}
				}
				if isChild {
					childModules[tree.Id] = true
				}
			}

			hasObjects := false
			for moduleID := range childModules {
				if _, ok := objects[moduleID]; ok {
					hasObjects = true
					break
				}
			}
			if !hasObjects {
				continue
			}
		}
		if extra["taskID"] == nil {
			extra["branchID"] = branch
		}
		linkHtml, err := callback(data, viewType, module, extra)
		if err != nil {
			return nil, err
		}
		if _, ok := treeMenuMap[module.Id]; ok {

			treeMenuMap[module.Parent] += "<li class='closed'>" + linkHtml
			treeMenuMap[module.Parent] += "<ul>" + treeMenuMap[module.Id] + "</ul>\n"
		} else {

			treeMenuMap[module.Parent] += "<li>" + linkHtml + "\n"
		}
		treeMenuMap[module.Parent] += "</li>\n"
	}
	return
}
func tree_createManageLink(data *TemplateData, viewType string, module *protocol.MSG_PROJECT_tree_cache, extra map[string]interface{}) (string, error) {
	branchID, _ := strconv.Atoi(extra["branchID"].(string))

	out := protocol.GET_MSG_USER_getPairs()
	out.Params = "noletter"

	buf := bufpool.Get().(*libraries.MsgBuffer)
	defer func() {
		buf.Reset()
		bufpool.Put(buf)
	}()
	buf.WriteString(module.Name)
	if module.Type != "story" && strings.Contains("bug,case", viewType) {
		buf.WriteString(`<span style="font-size:smaller;">[`)
		buf.WriteString(strings.ToUpper(viewType[0:1]))
		buf.WriteString(`]</span>`)

	}

	if viewType == "bug" && module.OwnerID > 0 {
		buf.WriteString(`<span class="owner">[`)
		user := HostConn.GetUserCacheById(module.OwnerID)
		if user.Realname != "" {
			buf.WriteString(user.Realname)
		} else {
			buf.WriteString(user.Account)
		}
		buf.WriteString(`]</span>`)
	}
	if viewType != "story" && module.Type == "story" {
		if hasPriv(data, "tree", "edit") && viewType == "bug" {

			buf.WriteString(html_a(createLink("tree", "edit", []interface{}{"module=", module.Id, "&type=", viewType, "&branch=", branchID}), data.Lang["tree"]["edit"].(string), "", `data-toggle="modal" data-type="ajax"`))
		}
		if hasPriv(data, "tree", "browse") {

			buf.WriteString(html_a(createLink("tree", "browse", []interface{}{"root=", module.Root, "&type=", viewType, "&module=", module.Id, "&branch=", branchID}), data.Lang["tree"]["child"].(string)))

		}
	} else {
		if hasPriv(data, "tree", "edit") {
			buf.WriteString(html_a(createLink("tree", "edit", []interface{}{"module=", module.Id, "&type=", viewType, "&branch=", branchID}), data.Lang["tree"]["edit"].(string), "", `data-toggle="modal" data-type="ajax" data-width="500"`))
		}
		if hasPriv(data, "tree", "browse") && strings.Contains(data.Config["tree"]["common"]["noBrowse"].(string), module.Type) {
			buf.WriteString(html_a(createLink("tree", "browse", []interface{}{"root=", module.Root, "&type=", viewType, "&module=", module.Id, "&branch=", branchID}), data.Lang["tree"]["child"].(string)))

		}
		if hasPriv(data, "tree", "delete") {
			buf.WriteString(html_a(createLink("tree", "delete", []interface{}{"root=", module.Root, "&module=", module.Id}), data.Lang["common"]["delete"].(string), "hiddenwin"))
		}
		if hasPriv(data, "tree", "updateorder") {
			buf.WriteString(html_input("orders["+strconv.Itoa(int(module.Id)), strconv.Itoa(int(module.Order)), `class="text-center w-40px inline"`))
		}
	}
	res := buf.String()
	return res, nil
}
func tree_createLineLink(data *TemplateData, viewType string, module *protocol.MSG_PROJECT_tree_cache, extra map[string]interface{}) (string, error) {
	productID, _ := extra["productID"].(int32)
	status, _ := extra["status"].(string)
	return html_a(createLink("product", "all", []interface{}{"productID=", productID, "&line=", module.Id, "&status=", status}), module.Name, "_self", "id='module"+strconv.Itoa(int(module.Id))+"'"), nil
}
func post_tree_updateOrder(data *TemplateData) (action gnet.Action) {
	out := protocol.GET_MSG_PROJECT_tree_updateList()
	list, err := tree_getAllcache(data)
	if err != nil {
		data.ws.WriteString(js.Alert(err.Error()))
		return
	}
	for sid, order := range data.ws.GetAllPost() {
		id, _ := strconv.Atoi(sid)
		for _, m := range list {
			if m.Id == int32(id) {
				o, _ := strconv.Atoi(order[0])
				m.Order = int16(o)
				out.Modules = append(out.Modules, m)
			}
		}
	}
	err = HostConn.SendMsgWaitResultToDefault(out, nil)
	if err != nil {
		data.ws.WriteString(js.Alert(err.Error()))
		return
	}
	data.ws.WriteString(js.Location("reload", "_self"))
	return
}
func get_tree_edit(data *TemplateData) (action gnet.Action) {
	moduleID, _ := strconv.Atoi(data.ws.Query("moduleID"))
	branch, _ := strconv.Atoi(data.ws.Query("branch"))
	viewType := data.ws.Query("type")
	module := HostConn.GetTreeById(int32(moduleID))
	if module == nil {
		data.ws.WriteString(js.Alert(data.Lang["tree"]["error"].(map[string]string)["ModuleNotFound"]) + js.Reload("parent"))
		return
	}
	if module.OwnerID == 0 && module.Root != 0 && module.Type != "task" && viewType != "doc" {
		if product := HostConn.GetProductById(module.Root); product != nil {
			module.OwnerID = product.QD
		}
	}
	var err error
	if viewType == "task" {
		data.Data["optionMenu"], err = tree_getTaskOptionMenu(module.Root, 0, 0)
	} else {
		data.Data["optionMenu"], err = tree_getOptionMenu(data, module.Root, module.Type, 0, int32(branch))
	}
	if err != nil {
		data.ws.WriteString(js.Alert(err.Error()) + js.Reload("parent"))
		return
	}
	data.Data["module"] = module
	data.Data["type"] = viewType
	// data.Data["libs"]   = $this->loadModel('doc')->getLibs($type = 'all', $extra = 'withObject');
	data.Data["branch"] = branch
	userGetPairs := protocol.GET_MSG_USER_getPairs()
	userGetPairs.Params = "noclosed|nodeleted"
	userGetPairs.UsersToAppended = module.OwnerID
	var result *protocol.MSG_USER_getPairs_result
	err = HostConn.SendMsgWaitResultToDefault(userGetPairs, &result)
	if err != nil {
		data.ws.WriteString(js.Alert(err.Error()) + js.Reload("parent"))
		return
	}
	data.Data["users"] = result.List
	data.Data["showProduct"] = strings.Contains("story|bug|case", viewType)
	if data.Data["showProduct"].(bool) {
		data.Data["products"], err = product_getPairs(data)
		if err != nil {
			data.ws.WriteString(js.Alert(err.Error()) + js.Reload("parent"))
			return
		}
	}
	/* Remove self and childs from the $optionMenu. Because it's parent can't be self or childs. */
	childs := tree_getAllChildId(data, int32(moduleID))
	for _, childModuleID := range childs {
		for i := len(data.Data["optionMenu"].([]protocol.HtmlKeyValueStr)) - 1; i >= 0; i-- {
			if data.Data["optionMenu"].([]protocol.HtmlKeyValueStr)[i].Key == strconv.Itoa(int(childModuleID)) {
				data.Data["optionMenu"] = append(data.Data["optionMenu"].([]protocol.HtmlKeyValueStr)[:i], data.Data["optionMenu"].([]protocol.HtmlKeyValueStr)[i+1:]...)
			}
		}
	}
	templateOut("tree.edit.html", data)
	return
}
func tree_getTaskOptionMenu(rootID int32, productID int32, startModule int32) ([]protocol.HtmlKeyValueStr, error) {
	/* If createdVersion <= 4.1, go to getOptionMenu(). */
	/*products       = $this->loadModel('product')->getProductsByProject($rootID);
	  $branchGroups   = $this->loadModel('branch')->getByProducts(array_keys($products));

	  if(!$this->isMergeModule($rootID, 'task') or !$products) return $this->getOptionMenu($rootID, 'task', $startModule);


	  $startModulePath = '';
	  if($startModule > 0)
	  {
	      $startModule = $this->getById($startModule);
	      if($startModule)
	      {
	          $startModulePath = $startModule->path . '%';
	          $modulePaths = explode(",", $startModulePath);
	          $rootModule  = $this->getById($modulePaths[0]);
	          $productID   = $rootModule->root;
	      }
	  }
	  $treeMenu   = array();
	  $lastMenu[] = '/';
	  $projectModules   = $this->getTaskTreeModules($rootID, true);
	  $noProductModules = $this->dao->select('*')->from(TABLE_MODULE)->where("root = '" . (int)$rootID . "' and type = 'task' and parent = 0")->andWhere('deleted')->eq(0)->fetchPairs('id', 'name');


	  $productNum = count($products);
	  foreach(array('product' => $products, 'noProduct' => $noProductModules) as $type => $rootModules)
	  {
	      foreach($rootModules as $id => $rootModule)
	      {
	          if($type == 'product')
	          {
	              $modules = $this->dao->select('*')->from(TABLE_MODULE)->where("((root = '" . (int)$rootID . "' and type = 'task' and parent != 0) OR (root = $id and type = 'story'))")
	                  ->beginIF($startModulePath)->andWhere('path')->like($startModulePath)->fi()
	                  ->andWhere('deleted')->eq(0)
	                  ->orderBy('grade desc, branch, type, `order`')
	                  ->fetchAll('id');
	          }
	          else
	          {
	              $modules = $this->dao->select('*')->from(TABLE_MODULE)->where("root = '" . (int)$rootID . "' and type = 'task' and path like '%,$id,%'")
	                  ->beginIF($startModulePath)->andWhere('path')->like($startModulePath)->fi()
	                  ->andWhere('deleted')->eq(0)
	                  ->orderBy('grade desc, type, `order`')
	                  ->fetchAll('id');
	          }

	          foreach($modules as $module)
	          {
	              $parentModules = explode(',', trim($module->path, ','));
	              if($type == 'product' and isset($noProductModules[$parentModules[0]])) continue;
	              $rootName = ($productNum > 1 and $type == 'product') ? "/$rootModule/" : '/';
	              if($type == 'product' and $module->branch and isset($branchGroups[$id][$module->branch])) $rootName .= $branchGroups[$id][$module->branch] . '/';
	              $this->buildTreeArray($treeMenu, $modules, $module, $rootName);
	          }

	          ksort($treeMenu);
	          $topMenu = @array_shift($treeMenu);
	          $topMenu = explode("\n", trim($topMenu));
	          foreach($topMenu as $menu)
	          {
	              if(!strpos($menu, '|')) continue;
	              list($label, $moduleID) = explode('|', $menu);
	              if(isset($projectModules[$moduleID])) $lastMenu[$moduleID] = $label;
	          }
	          foreach($topMenu as $moduleID => $moduleName)
	          {
	              if(!isset($projectModules[$moduleID])) unset($treeMenu[$moduleID]);
	          }
	      }
	  }
	  return $lastMenu;*/
	return nil, nil
}
func tree_getAllChildId(data *TemplateData, moduleID int32) (res []int32) {

	list, _ := tree_getAllcache(data)
	var module *protocol.MSG_PROJECT_tree_cache
	for _, v := range list {
		if v.Id == moduleID {
			module = v
			break
		}
	}
	if module != nil {
		for _, v := range list {
			find := true
			if !v.Deleted && len(v.Path) > len(module.Path) {
				for k, id := range module.Path {
					if v.Path[k] != id {
						find = false
						break
					}
				}
			} else {
				find = false
			}
			if find {
				res = append(res, v.Id)
			}
		}
	}
	return
}
func post_tree_edit(data *TemplateData) (action gnet.Action) {
	moduleID, _ := strconv.Atoi(data.ws.Query("module"))
	module := HostConn.GetTreeById(int32(moduleID))
	if module == nil {
		data.ws.WriteString(js.Alert(data.Lang["tree"]["error"].(map[string]string)["ModuleNotFound"]) + js.Reload("parent"))
		return
	}
	rootID, err := strconv.Atoi(data.ws.Post("root"))
	if err != nil {
		rootID = -1
	}
	parent, _ := strconv.Atoi(data.ws.Post("parent"))
	out := protocol.GET_MSG_PROJECT_tree_manageChild()
	out.RootID = int32(rootID)
	out.ViewType = data.ws.Query("type")
	out.ParentModuleID = int32(parent)
	module.Name = data.ws.Post("name")
	module.Short = data.ws.Post("short")
	out.Modules = append(out.Modules, module)
	var result *protocol.MSG_PROJECT_tree_manageChild_result
	err = HostConn.SendMsgWaitResultToDefault(out, &result)
	if err != nil {
		data.ws.WriteString(js.Alert(err.Error()) + js.Reload("parent"))
		return
	}
	if result.Result != protocol.Success {
		data.ws.WriteString(js.Alert(data.Lang["tree"]["error"].(map[string]string)[result.Result.String()]) + js.Reload("parent"))
	} else {
		data.ws.WriteString(js.Alert(data.Lang["tree"]["successSave"].(string)) + js.Reload("parent"))
	}
	return
}
func get_tree_delete(data *TemplateData) (action gnet.Action) {
	confirm := data.ws.Query("confirm")
	if confirm != "yes" {
		moduleID, _ := strconv.Atoi(data.ws.Query("moduleID"))
		module := HostConn.GetTreeById(int32(moduleID))
		if module == nil {
			data.ws.WriteString(js.Alert(data.Lang["tree"]["error"].(map[string]string)["ModuleNotFound"]) + js.Reload("parent"))
			return
		}
		confirmLang := data.Lang["tree"]["confirmDelete"].(string)
		if module.Type == "line" {
			confirmLang = data.Lang["tree"]["confirmDeleteLine"].(string)
		}
		data.ws.WriteString(js.Confirm(confirmLang, createLink("tree", "delete", "rootID="+data.ws.Query("rootID")+"&moduleID="+data.ws.Query("moduleID")+"&confirm=yes"), ""))
	} else {
		out := protocol.GET_MSG_PROJECT_tree_delete()
		moduleID, _ := strconv.Atoi(data.ws.Query("moduleID"))
		out.Ids = append(out.Ids, int32(moduleID))
		err := HostConn.SendMsgWaitResultToDefault(out, nil)
		if err != nil {
			data.ws.WriteString(js.Alert(err.Error()) + js.Reload("parent"))
		} else {
			data.ws.WriteString(js.Reload("parent"))
		}
	}
	return
}
