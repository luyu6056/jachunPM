package handler

import (
	"libraries"
	"protocol"
	"sort"
	"strconv"
	"strings"

	"github.com/luyu6056/gnet"
)

func init() {
	httpHandlerMap["GET"]["/tree/browse"] = get_tree_browse
	httpHandlerMap["POST"]["/tree/manageChild"] = tree_manageChild

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
		//$this->view->root = $this->doc->getLibById($rootID);
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

		// $this->lang->tree->menu      = $this->lang->product->menu;
		// $this->lang->tree->menuOrder = $this->lang->product->menuOrder;
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
		list, err := tree_buildMenuQuery(rootID, typ, startModule, int32(branchID))
		if err != nil {
			return nil, err
		}
		var modules map[int32]*protocol.MSG_PROJECT_tree_cache
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
func tree_buildMenuQuery(rootID int32, typ string, startModule, branch int32) ([]*protocol.MSG_PROJECT_tree_cache, error) {
	list, err := tree_getAllcache()
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
		if m.Deleted || m.Root != rootID || m.Branch != branch {
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
			if typ == "tack" {
				if m.Type != "task" {
					continue
				}
			} else {
				if m.Type != "story" && m.Type != "type" {
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
		if a.Grade > b.Grade {
			return true
		}
		if a.Type > b.Type {
			return true
		}
		if a.Order < b.Order {
			return true
		}
		return false
	})
	//->orderBy("grade desc, type desc, `order`")
	return res, nil
}
func tree_getAllcache() (result []*protocol.MSG_PROJECT_tree_cache, err error) {
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
	return
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
	product := HostConn.GetProductById(rootID)
	if strings.Contains("story|bug|case", viewType) {
		if product.Type != "normal" {
			branches = branch_getPairs(data, 0, product, "noempty")
		}
	}

	list, err := tree_buildMenuQuery(rootID, viewType, 0, 0)
	if err != nil {
		return nil, err
	}
	fullTrees = tree_getDataStructure(list, viewType, nil)

	if len(branches) > 0 {
		var branchTrees []map[string]interface{}
		for _, branchkv := range branches {
			branchID, _ := strconv.Atoi(branchkv.Key)
			list, err := tree_buildMenuQuery(rootID, viewType, 0, int32(branchID))
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
				"children": tree_getDataStructure(list, viewType, nil),
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
func tree_getDataStructure(list []*protocol.MSG_PROJECT_tree_cache, viewType string, keepModules []int32) []map[string]interface{} {
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
		alltree, _ := tree_getAllcache()
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
func tree_getLinePairs() (res []protocol.HtmlKeyValueStr, err error) {
	list, err := tree_getAllcache()
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
func tree_manageChild(data *TemplateData) (action gnet.Action) {
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
