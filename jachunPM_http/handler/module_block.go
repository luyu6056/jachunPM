package handler

import (
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

	httpHandlerMap["GET"]["/block/dashboard"] = get_block_dashboard
	httpHandlerMap["GET"]["/block/getblockdata"] = get_block_getblockdata
	httpHandlerMap["GET"]["/block/ajaxReset"] = get_block_ajaxReset

}
func blockTemplateFuncs() {
	global_Funcs["block_printBlock"] = func(data *TemplateData, block protocol.HtmlBlock, module string) template.HTML {
		switch {
		case block.Block == "html":
			if v, ok := block.Params["html"]; !ok {
				return template.HTML("<div class='empty-tip'>" + data.Lang["block"]["emptyTip"].(string) + "</div>")
			} else {
				return template.HTML("<div class='panel-body'><div class='article-content'>" + libraries.Bbcode2html(v, true, true, false, false, true, true) + "</div></div>")
			}
		case block.Source != "":
			data.ws.Session().Set("temp_block", block)
			return template.HTML(common_fetch(data, "block", "getblockdata"))

		case block.Block == "dynamic":
			return template.HTML(common_fetch(data, "block", "dynamic"))

		case block.Block == "flowchart":
			return template.HTML(common_fetch(data, "block", "flowchart"))

		case block.Block == "assigntome":
			data.ws.Session().Set("temp_block", block)
			return template.HTML(common_fetch(data, "block", "printAssignToMeBlock"))

		case block.Block == "welcome":
			return template.HTML(common_fetch(data, "block", "welcome"))
		}
		return template.HTML("")
	}
}
func get_block_dashboard(data *TemplateData) (err error) {
	module := data.ws.Query("module")
	data.ws.Session().Set("blockModule", module)
	getBlock := protocol.GET_MSG_USER_block_getList()
	getBlock.Module = module
	getBlock.Uid = data.User.Id
	var result *protocol.MSG_USER_block_getList_result
	if err = data.SendMsgWaitResultToDefault(getBlock, &result); err != nil {
		return
	}
	inited := false
	if v, ok := data.Config[module]["common"]["blockInited"].(string); ok {
		inited = v == "1"
	}
	if len(result.List) == 0 && !inited {
		if err = block_initBlock(data, module); err != nil {
			return
		}
		data.ws.WriteString(js.Reload())
		return
	}
	//var blocks []protocol.HtmlBlock
	var shortBlocks, longBlocks []protocol.HtmlBlock

	for _, block := range result.List {
		if block.Source != "" && block.Source != "my" && !data.User.IsAdmin && !data.User.AclMenu[block.Source] {
			continue
		}
		hBlock := protocol.HtmlBlock{
			Id:     block.Id,
			Source: block.Source,
			Grid:   block.Grid,
			Title:  block.Title,
			Block:  block.Block,
			Height: block.Height,
			Params: make(map[string]string),
			Order:  block.Order,
			Module: block.Module,
		}
		libraries.JsonUnmarshalStr(block.Params, &hBlock.Params)
		blockID := block.Block
		source := "common"
		if block.Source != "" {
			source = block.Source
		}

		hBlock.BlockLink = createLink("block", "printBlock", []interface{}{"id=", block.Id, "&module=", block.Module})
		hBlock.MoreLink = ""
		if v, ok := data.Lang["block"]["modules"].(map[string]protocol.HtmlBlockModule)[source]; ok && v.MoreLinkList != nil && v.MoreLinkList[blockID] != "" {
			if s := strings.Split(fmt.Sprintf(v.MoreLinkList[blockID], hBlock.Params["type"]), "|"); len(s) == 3 {
				hBlock.MoreLink = createLink(s[0], s[1], s[2])
			}

		} else if block.Block == "dynamic" {
			hBlock.MoreLink = createLink("company", "dynamic", nil)
		}

		hBlock.ActionLink = ""
		if block.Block == "overview" {
			if module == "product" && hasPriv(data, "product", "create") {
				hBlock.ActionLink = html_a(createLink("product", "create", nil), "<i class='icon icon-sm icon-plus'></i> "+data.Lang["product"]["create"].(string), "", "class='btn btn-primary'")
			}
			if module == "project" && hasPriv(data, "project", "create") {
				hBlock.ActionLink = html_a(createLink("project", "create", nil), "<i class='icon icon-sm icon-plus'></i> "+data.Lang["project"]["create"].(string), "", "class='btn btn-primary'")
			}
			if module == "qa" && hasPriv(data, "testcase", "create") {

				hBlock.ActionLink = html_a(createLink("testcase", "create", "productID="), "<i class='icon icon-sm icon-plus'></i> "+data.Lang["testcase"]["create"].(string), "", "class='btn btn-primary'")
			}
		}
		if block.Grid >= 6 {
			longBlocks = append(longBlocks, hBlock)
		} else {
			shortBlocks = append(shortBlocks, hBlock)
		}
		//blocks = append(blocks, hBlock)
	}
	/*公共相关
	notice := block_getBytitle()
	if notice != nil {
		if notice.Params["html"] == "" {
			html = "<div class='empty-tip'>" + data.Lang["block"]["emptyTip"].(string) + "</div>"
		} else {
			html = "<div class='panel-body'><div class='article-content'>" + notice.Params["html"] + "</div></div>"
		}
	}*/
	data.Data["longBlocks"] = longBlocks
	data.Data["shortBlocks"] = shortBlocks
	data.Data["module"] = module
	dropmenu := protocol.BufPoolGet()
	dropmenu.WriteString("<div class='btn-group'><button type='button' class='btn dropdown-toggle' data-toggle='dropdown' style='padding-bottom:7px;'>")
	dropmenu.WriteString(data.Lang["block"]["common"].(string))
	dropmenu.WriteString(" <span class='caret'></span></button><ul class='dropdown-menu pull-right' role='menu'><li>")
	dropmenu.WriteString(html_a(createLink("block", "admin", "id=0&module="+module), "<i class='icon icon-plus'></i> "+data.Lang["block"]["createBlock"].(string), "", "data-toggle='modal' data-type='ajax' data-width='700' data-title='"+data.Lang["block"]["createBlock"].(string)+"'"))
	dropmenu.WriteString("</li><li>")
	dropmenu.WriteString(html_a(createLink("block", "ajaxReset", "module="+module), "<i class='icon icon-refresh'></i>"+data.Lang["block"]["reset"].(string), "hiddenwin"))
	dropmenu.WriteString("</li></ul></div>")
	data.Data["dropmenu"] = dropmenu.String()
	templateOut("block.dashboard.html", data)
	dropmenu.Reset()
	protocol.BufPoolPut(dropmenu)
	return nil
}
func block_initBlock(data *TemplateData, module string) (err error) {

	blocks := data.Lang["block"]["default"].(map[string][]protocol.HtmlBlock)[module]
	configSave := protocol.GET_MSG_USER_config_save()
	configSave.Module = module
	configSave.Section = "common"
	configSave.Key = "blockInited"
	configSave.Value = "1"
	configSave.Uid = data.User.Id
	configSave.Type = "add"
	if err = data.SendMsgWaitResultToDefault(configSave, nil); err != nil {
		return
	}
	configSave.Put()
	/* Mark this app has init. */
	insert := protocol.GET_MSG_USER_block_insertUpdate()
	insert.Insert = true
	for key, block := range blocks {
		tmp := &protocol.MSG_USER_Block_info{
			Uid:    data.User.Id,
			Module: module,
			Title:  block.Title,
			Source: block.Source,
			Block:  block.Block,
			Params: libraries.JsonMarshalToString(block.Params),
			Order:  int8(key),
			Grid:   block.Grid,
			Height: block.Height,
			Hidden: false,
		}
		if tmp.Source == "" {
			tmp.Source = module
		}
		insert.List = append(insert.List, tmp)
	}
	if err = data.SendMsgWaitResultToDefault(insert, nil); err != nil {
		return
	}
	return
}
func get_block_getblockdata(data *TemplateData) (err error) {
	var block protocol.HtmlBlock
	data.ws.Session().Get("temp_block", &block)
	if block.Block == "" {
		return
	}
	data.Data["longBlock"] = block.Grid >= 6
	key := "print" + strings.ToUpper(block.Block[:1]) + block.Block[1:] + "Block"
	switch key {
	case "printStatisticBlock":
		switch block.Source {
		case "project":
			return printProjectStatisticBlock(data)
		case "product":
			return printproductStatisticBlock(data)
		default:
			libraries.DebugLog("printStatisticBlock未处理" + block.Source)
			data.ws.WriteString("printStatisticBlock未处理" + block.Source)
		}
	case "printListBlock":
		switch block.Source {
		case "project":
			return printProjectBlock(data)
		case "product":
			return printProductBlock(data)
		default:
			libraries.DebugLog("printListBlock未处理" + block.Source)
			data.ws.WriteString("printListBlock未处理" + block.Source)
		}
	case "printOverviewBlock":
		switch block.Source {
		case "project":
			return printProjectOverviewBlock(data)
		case "product":
			return printProductOverviewBlock(data)
		default:
			libraries.DebugLog("printOverviewBlock未处理" + block.Source)
			data.ws.WriteString("printOverviewBlock未处理" + block.Source)
		}
	case "printTaskBlock":
		if libraries.Preg_match("/[^a-zA-Z0-9_]/", block.Params["type"]) {
			return
		}
		data.ws.Session().Set("taskList", data.ws.Referer())
		data.ws.Session().Set("storyList", data.ws.Referer())
		num, _ := strconv.Atoi(block.Params["num"])
		if data.Data["tasks"], err = task_getUserTasks(data, data.User.Id, block.Params["type"], &TempLatePage{Total: -1, Page: 1, PerPage: num}, block.Params["orderBy"]); err != nil {
			return err
		}
		templateOut("block.taskblock.html", data)
	case "printWelcomeBlock":
		return printWelcomeBlock(data)
	case "printStoryBlock":
		return printStoryBlock(data)
	default:
		libraries.DebugLog("get_block_getblockdata未处理" + key)
		data.ws.WriteString("get_block_getblockdata未处理" + key)
	}
	return

}
func printProjectStatisticBlock(data *TemplateData) (err error) {
	var block protocol.HtmlBlock
	data.ws.Session().Get("temp_block", &block)

	if libraries.Preg_match("/[^a-zA-Z0-9_]/", block.Params["type"]) {
		return
	}
	//status := block.Params["type"]
	num, _ := strconv.Atoi(block.Params["num"])
	var projects []*protocol.MSG_PROJECT_project_cache
	data.Data["projects"] = projects
	/* Get projects. */
	if allproject, err := project_getAll(data); err == nil && num > 0 {
		var mineProjects, otherProjects, closedProjects []*protocol.MSG_PROJECT_project_cache
		for _, project := range allproject {
			if !data.User.IsAdmin && !data.User.AclProjects[project.Id] || project.Deleted {
				continue
			}
			if project.Status != "done" && project.Status != "closed" && project.PM == data.User.Id {
				mineProjects = append(mineProjects, project)
			} else if project.Status != "done" && project.Status != "closed" && project.PM != data.User.Id {
				otherProjects = append(otherProjects, project)
			} else if project.Status == "done" || project.Status == "closed" {
				closedProjects = append(closedProjects, project)
			}
		}
		projects = append(mineProjects, append(otherProjects, closedProjects...)...)
		if len(projects) > num {
			projects = projects[:num]
		}

	} else {
		return err
	}
	if len(projects) == 0 {
		return
	}
	var projectMap = make(map[string]map[string]interface{})
	var projectIdList []string

	for _, p := range projects {
		id := strconv.Itoa(int(p.Id))
		projectMap[id] = make(map[string]interface{})
		projectMap[id]["Name"] = p.Name
		projectIdList = append(projectIdList, id)
	}
	/* Get tasks. */
	yesterday := time.Now().AddDate(0, 0, -1).Format(protocol.TIMEFORMAT_MYSQLDATE)
	gettask := protocol.GET_MSG_PROJECT_doRawSelect()
	gettask.Sql = fmt.Sprintf("select project, count(id) as totalTasks, count(status in (\"wait\",\"doing\",\"pause\") or null) as undoneTasks, count(finishedDate like \"%s%%\" or null) as yesterdayFinished, sum(if(status != \"cancel\", estimate, 0)) as totalEstimate, sum(consumed) as totalConsumed, sum(if(status != \"cancel\", `left`, 0)) as totalLeft from %s where project in(%s) and deleted=0 and parent<1 group by project", yesterday, "task", strings.Join(projectIdList, ","))
	var gettaskResult *protocol.MSG_PROJECT_doRawSelect_result
	if err = data.SendMsgWaitResultToDefault(gettask, &gettaskResult); err != nil {
		return
	}
	for _, row := range gettaskResult.List {

		for key, value := range row {
			if key == "project" {
				continue
			}
			projectMap[row["project"]][key] = value
		}
		projectMap[row["project"]]["undoneTasks"], _ = strconv.Atoi(row["undoneTasks"])
		projectMap[row["project"]]["totalTasks"], _ = strconv.Atoi(row["totalTasks"])
		projectMap[row["project"]]["totalEstimate"] = libraries.Round(row["totalEstimate"], 2)
		projectMap[row["project"]]["totalConsumed"] = libraries.Round(row["totalConsumed"], 2)
		projectMap[row["project"]]["totalLeft"] = libraries.Round(row["totalLeft"], 2)
	}
	gettask.Sql = fmt.Sprintf("select project, count(status) as totalStories, count(status != \"closed\" or null) as unclosedStories, count(stage = \"released\" or null) as releasedStories from %s where project in(%s) and deleted=0 group by project", "story", strings.Join(projectIdList, ","))
	/* Get stories. */
	if err = data.SendMsgWaitResultToDefault(gettask, &gettaskResult); err != nil {

		return
	}
	for _, row := range gettaskResult.List {
		for key, value := range row {
			if key == "project" {
				continue
			}
			projectMap[row["project"]][key] = value
		}
		projectMap[row["project"]]["totalStories"], _ = strconv.Atoi(row["totalStories"])
		projectMap[row["project"]]["unclosedStories"], _ = strconv.Atoi(row["unclosedStories"])
	}

	/* Get bugs. */
	getbug := protocol.GET_MSG_TEST_doRawSelect()
	getbug.Sql = fmt.Sprintf("select project, status, count(status) as totalBugs, count(status = \"active\" or null) as activeBugs, count(resolvedDate like \"%s%%\" or null) as yesterdayResolved from %s where project in(%s) and deleted=0 group by project", yesterday, "bug", strings.Join(projectIdList, ","))
	var getbufResult *protocol.MSG_TEST_doRawSelect_result
	if err = data.SendMsgWaitResultToDefault(getbug, &getbufResult); err != nil {

		return
	}
	for _, row := range getbufResult.List {
		for key, value := range row {
			if key == "project" {
				continue
			}
			projectMap[row["project"]][key] = value
		}
		projectMap[row["project"]]["totalBugs"], _ = strconv.Atoi(row["totalBugs"])
		projectMap[row["project"]]["activeBugs"], _ = strconv.Atoi(row["activeBugs"])
	}

	for key, project := range projectMap {
		if _, ok := project["totalTasks"]; !ok {
			project["totalTasks"] = 0
			project["undoneTasks"] = 0
			project["yesterdayFinished"] = 0
			project["totalEstimate"] = 0
			project["totalConsumed"] = float64(0)
			project["totalLeft"] = float64(0)
		}

		if _, ok := project["totalBugs"]; !ok {
			project["totalBugs"] = 0
			project["activeBugs"] = 0
			project["yesterdayResolved"] = 0
		}

		if _, ok := project["totalStories"]; !ok {
			project["totalStories"] = 0
			project["unclosedStories"] = 0
			project["releasedStories"] = 0
		}

		project["progress"] = 0
		if project["totalConsumed"].(float64) > 0 || project["totalLeft"].(float64) > 0 {
			project["progress"] = libraries.Round(project["totalConsumed"].(float64)/(project["totalConsumed"].(float64)+project["totalLeft"].(float64)), 2) * 100
		}
		project["taskProgress"] = 0
		if project["totalTasks"].(int) > 0 {
			project["taskProgress"] = libraries.Round(float64(project["totalTasks"].(int)-project["undoneTasks"].(int))/float64(project["totalTasks"].(int)), 2) * 100
		}
		project["storyProgress"] = 0
		if project["totalStories"].(int) > 0 {
			project["storyProgress"] = libraries.Round(float64(project["totalStories"].(int)-project["unclosedStories"].(int))/float64(project["totalStories"].(int)), 2) * 100
		}
		project["bugProgress"] = 0
		if project["totalBugs"].(int) > 0 {
			project["bugProgress"] = libraries.Round(float64(project["totalBugs"].(int)-project["activeBugs"].(int))/float64(project["totalBugs"].(int)), 2) * 100
		}
		projectMap[key] = project
	}
	data.Data["projectIdList"] = projectIdList
	data.Data["projects"] = projectMap
	templateOut("block.projectstatisticblock.html", data)
	gettask.Put()
	gettaskResult.Put()
	getbug.Put()
	getbufResult.Put()
	return
}
func printProjectBlock(data *TemplateData) (err error) {
	var block protocol.HtmlBlock
	data.ws.Session().Get("temp_block", &block)
	if libraries.Preg_match("/[^a-zA-Z0-9_]/", block.Params["type"]) {
		return
	}
	typ := block.Params["type"]
	num, _ := strconv.Atoi(block.Params["num"])
	if typ == "" {
		typ = "all"
	}
	data.Page.PerPage = num
	data.Page.Page = 1
	data.Data["projectStats"], err = project_getProjectStats(data, typ, 0, 0, 30, "order_desc")
	templateOut("block.projectblock.html", data)
	return
}
func printProjectOverviewBlock(data *TemplateData) (err error) {
	projects, err := project_getAll(data)
	if err != nil {
		return
	}
	total := 0
	overview := make(map[string]int)
	for _, project := range projects {
		if !data.User.IsAdmin && !data.User.AclProjects[project.Id] || project.Deleted {
			continue
		}
		overview[project.Status]++
		total++
	}

	overviewPercent := make(map[string]string)
	for _, kv := range data.Lang["project"]["statusList"].([]protocol.HtmlKeyValueStr) {
		if total > 0 {
			overviewPercent[kv.Key] = strconv.FormatFloat(libraries.Round(float64(overview[kv.Key])/float64(total), 2)*100, 'f', 2, 64)
		} else {
			overviewPercent[kv.Key] = "0%"
		}
	}

	data.Data["total"] = total
	data.Data["overview"] = overview
	data.Data["overviewPercent"] = overviewPercent
	templateOut("block.projectoverviewblock.html", data)
	return nil
}
func printproductStatisticBlock(data *TemplateData) (err error) {
	var block protocol.HtmlBlock
	data.ws.Session().Get("temp_block", &block)
	if libraries.Preg_match("/[^a-zA-Z0-9_]/", block.Params["type"]) {
		return
	}
	status := block.Params["type"]
	num, _ := strconv.Atoi(block.Params["num"])
	products, err := block_getProducts(data, status, num)
	if err != nil {
		return
	}
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1).Format(protocol.TIMEFORMAT_MYSQLDATE)
	var stories = make(map[int32]map[string]int)
	var plans = make(map[int32]map[string]int)
	var projects = make(map[int32]map[string]int)
	var releases = make(map[int32]map[string]int)
	var lastReleases = make(map[int32]int)
	rawSelect := protocol.GET_MSG_PROJECT_doRawSelect()
	var result *protocol.MSG_PROJECT_doRawSelect_result
	if len(products) > 0 {
		var ids []string
		for _, product := range products {
			ids = append(ids, strconv.Itoa(int(product.Id)))
		}

		rawSelect.Sql = fmt.Sprintf("select product, stage, COUNT(status) AS count from %s where product in(%s) and deleted=0 group by product,stage", "story", strings.Join(ids, ","))

		if err = data.SendMsgWaitResultToDefault(rawSelect, &result); err != nil {
			return
		}
		for _, row := range result.List {
			id, _ := strconv.Atoi(row["product"])
			if stories[int32(id)] == nil {
				stories[int32(id)] = make(map[string]int)
			}
			stories[int32(id)][row["stage"]], _ = strconv.Atoi(row["count"])
		}
		rawSelect.Sql = fmt.Sprintf("select product, end from %s where product in(%s) and deleted=0", "productplan", strings.Join(ids, ","))
		result.List = result.List[:0]
		if err = data.SendMsgWaitResultToDefault(rawSelect, &result); err != nil {
			return
		}
		for _, row := range result.List {
			id, _ := strconv.Atoi(row["product"])
			if plans[int32(id)] == nil {
				plans[int32(id)] = make(map[string]int)
			}
			end, err := time.ParseInLocation(protocol.TIMEFORMAT_MYSQLTIME, row["end"], time.Local)
			if err != nil {
				return err
			}
			if end.After(now) {
				plans[int32(id)]["unexpired"]++
			} else {
				plans[int32(id)]["expired"]++
			}
			plans[int32(id)]["totalPlan"]++
		}
		for _, product := range products {
			if plans[product.Id] == nil {
				plans[product.Id] = map[string]int{
					"totalPlan":     0,
					"unexpiredRate": 0,
				}
			} else {
				plans[product.Id]["unexpiredRate"] = int(float64(plans[product.Id]["unexpired"]) / float64(plans[product.Id]["totalPlan"]) * 100)
			}
		}
		allprojects, err := project_getAll(data)
		if err != nil {
			return err
		}
		for _, project := range allprojects {
			if project.Deleted {
				continue
			}
			for _, id := range project.Products {
				for _, product := range products {
					if product.Id == id {
						if projects[product.Id] == nil {
							projects[product.Id] = make(map[string]int)
						}
						projects[product.Id]["all"]++
						if project.Status == "doing" {
							projects[product.Id]["doing"]++
						} else if project.Status == "done" || project.Status == "closed" {
							projects[product.Id]["done"]++
						} else if project.Status != "done" && project.Status != "closed" && project.Status != "suspended" && now.After(project.End) {
							projects[product.Id]["delay"]++
						}
						break
					}
				}
			}
		}
		for _, product := range products {
			if projects[product.Id] == nil {
				projects[product.Id] = make(map[string]int)
			} else {
				projects[product.Id]["doingRate"] = int(float64(projects[product.Id]["doing"]) / float64(projects[product.Id]["all"]) * 100)
			}
		}
		rawSelect.Sql = fmt.Sprintf("select product, status, COUNT(*) AS count from %s where product in(%s) and deleted=0 group by product,status", "`release`", strings.Join(ids, ","))
		result.List = result.List[:0]
		if err = data.SendMsgWaitResultToDefault(rawSelect, &result); err != nil {
			return err
		}
		for _, row := range result.List {
			id, _ := strconv.Atoi(row["product"])
			if releases[int32(id)] == nil {
				releases[int32(id)] = make(map[string]int)
			}
			releases[int32(id)][row["status"]], _ = strconv.Atoi(row["count"])
		}
		for _, product := range products {
			if releases[product.Id] == nil {
				releases[product.Id] = make(map[string]int)
			} else {
				releases[product.Id]["totalRelease"] = len(releases[product.Id])
				releases[product.Id]["normalRate"] = int(float64(releases[product.Id]["normal"]) / float64(releases[product.Id]["totalRelease"]) * 100)
			}
		}
		rawSelect.Sql = fmt.Sprintf("select product, COUNT(*) AS count from %s where product in(%s) and deleted=0 and `date`='%s' group by product", "`release`", strings.Join(ids, ","), yesterday)
		result.List = result.List[:0]
		if err = data.SendMsgWaitResultToDefault(rawSelect, &result); err != nil {
			return err
		}
		for _, row := range result.List {
			id, _ := strconv.Atoi(row["product"])
			lastReleases[int32(id)], _ = strconv.Atoi(row["count"])
		}
	}
	data.Data["products"] = products
	data.Data["stories"] = stories
	data.Data["plans"] = plans
	data.Data["projects"] = projects
	data.Data["releases"] = releases
	data.Data["lastRelease"] = lastReleases
	templateOut("block.productstatisticblock.html", data)
	rawSelect.Put()
	result.Put()
	return
}
func block_getProducts(data *TemplateData, status string, num int) ([]*protocol.MSG_PROJECT_product_cache, error) {
	products, err := product_getList(data, nil, status, 0, 0, 0)
	if len(products) == 0 || err != nil || num == 0 {
		return products, err
	}

	lines, err := tree_getLinePairs(data)
	if err != nil {
		return nil, err
	}
	order := int32(len(products)) * 10
	for _, l := range lines {
		for k, product := range products {
			if strconv.Itoa(int(product.Line)) == l.Key {
				product.Name = l.Value + "/" + product.Name
				product.Order = order*2 - int32(k)
			}
		}
	}
	protocol.Order_product(products, nil)
	var mineProducts, otherProducts, closedProducts []*protocol.MSG_PROJECT_product_cache
	for _, product := range products {
		if !data.User.IsAdmin && !data.User.AclProducts[product.Id] {
			continue
		}
		if product.Status == "normal" && product.PO == data.User.Id {
			mineProducts = append(mineProducts, product)
		} else if product.Status == "normal" && product.PO != data.User.Id {
			otherProducts = append(otherProducts, product)

		} else if product.Status == "closed" {
			closedProducts = append(closedProducts, product)

		}
	}
	products = append(mineProducts, append(otherProducts, closedProducts...)...)
	if num < len(products) {
		products = products[:num]
	}
	return products, nil
}

func printProductOverviewBlock(data *TemplateData) (err error) {
	products, err := product_getList(data, nil, "", 0, 0, 0)
	if err != nil {
		return err
	}
	var normal, closed int
	for _, product := range products {
		if !data.User.IsAdmin && !data.User.AclProducts[product.Id] {
			continue
		}
		if product.Status == "normal" {
			normal++
		} else if product.Status == "closed" {
			closed++
		}
	}
	total := normal + closed
	data.Data["total"] = total
	data.Data["normal"] = normal
	data.Data["closed"] = closed
	data.Data["normalPercent"] = 0
	if total > 0 {
		data.Data["normalPercent"] = int(float64(normal) / float64(total) * 100)
	}
	templateOut("block.productoverviewblock.html", data)
	return
}
func printWelcomeBlock(data *TemplateData) (err error) {
	today := time.Now().Format(protocol.TIMEFORMAT_MYSQLDATE)
	projectRawSelect := protocol.GET_MSG_PROJECT_doRawSelect()
	testRawSelect := protocol.GET_MSG_TEST_doRawSelect()
	result := &protocol.MSG_PROJECT_doRawSelect_result{
		List: make([]map[string]string, 0, 1),
	}
	testRawSelectresult := &protocol.MSG_TEST_doRawSelect_result{
		List: make([]map[string]string, 0, 1),
	}
	//获取task
	projectRawSelect.Sql = fmt.Sprintf("select count(*) AS count from %s where assignedTo=%d and deleted=0", "`task`", data.User.Id)
	result.List = result.List[:0]
	if err = data.SendMsgWaitResultToDefault(projectRawSelect, &result); err != nil {
		return err
	}
	data.Data["tasks"] = 0
	if len(result.List) == 1 {
		data.Data["tasks"], _ = strconv.Atoi(result.List[0]["count"])
	}

	//bug
	testRawSelect.Sql = fmt.Sprintf("select count(*) AS count from %s where assignedTo=%d and deleted=0", "`bug`", data.User.Id)
	testRawSelectresult.List = testRawSelectresult.List[:0]
	if err = data.SendMsgWaitResultToDefault(testRawSelect, &testRawSelectresult); err != nil {
		return err
	}
	data.Data["bugs"] = 0
	if len(testRawSelectresult.List) == 1 {
		data.Data["bugs"], _ = strconv.Atoi(testRawSelectresult.List[0]["count"])
	}

	//story
	projectRawSelect.Sql = fmt.Sprintf("select count(*) AS count from %s where assignedTo=%d and deleted=0", "`story`", data.User.Id)
	result.List = result.List[:0]
	if err = data.SendMsgWaitResultToDefault(projectRawSelect, &result); err != nil {
		return err
	}
	data.Data["stories"] = 0
	if len(result.List) == 1 {
		data.Data["stories"], _ = strconv.Atoi(result.List[0]["count"])
	}

	//project
	allprojects, err := project_getAll(data)
	if err != nil {
		return err
	}
	count := 0
	delay := 0
	for _, project := range allprojects {
		if !data.User.IsAdmin && !data.User.AclProjects[project.Id] || project.Deleted {
			continue
		}
		if project.Status == "wait" || project.Status == "doing" {
			count++
		}
		if time.Now().After(project.End) {
			delay++
		}
	}
	data.Data["projects"] = count
	data.Data["delayProject"] = delay

	//product
	products, err := product_getAll(data)
	if err != nil {
		return err
	}
	count = 0
	for _, product := range products {
		if !data.User.IsAdmin && !data.User.AclProducts[product.Id] || product.Deleted {
			continue
		}
		if product.Status != "closed" {
			count++
		}
	}
	data.Data["products"] = count

	//delayTask
	projectRawSelect.Sql = fmt.Sprintf("select count(*) AS count from %s where assignedTo=%d and deleted=0 and status in ('wait','doing') and deadline > '%s' and deadline < '%s'", "`task`", data.User.Id, protocol.NORMALTIME, today)
	result.List = result.List[:0]
	if err = data.SendMsgWaitResultToDefault(projectRawSelect, &result); err != nil {
		return err
	}
	data.Data["delayTask"] = 0
	if len(result.List) == 1 {
		data.Data["delayTask"], _ = strconv.Atoi(result.List[0]["count"])
	}

	//delayBug
	testRawSelect.Sql = fmt.Sprintf("select count(*) AS count from %s where assignedTo=%d and deleted=0 and status = 'active' and deadline > '%s' and deadline < '%s'", "`bug`", data.User.Id)
	testRawSelectresult.List = testRawSelectresult.List[:0]
	if err = data.SendMsgWaitResultToDefault(testRawSelect, &testRawSelectresult); err != nil {
		return err
	}
	data.Data["delayBug"] = 0
	if len(testRawSelectresult.List) == 1 {
		data.Data["delayBug"], _ = strconv.Atoi(testRawSelectresult.List[0]["count"])
	}
	welcomeType := "19:00"
	now := time.Now().Format("04:05")
	for _, kv := range data.Lang["block"]["welcomeList"].([]protocol.HtmlKeyValueStr) {
		if now > kv.Key {
			welcomeType = kv.Key
		}
	}
	data.Data["welcomeType"] = welcomeType
	templateOut("block.welcome.html", data)
	projectRawSelect.Put()
	testRawSelect.Put()
	result.Put()
	testRawSelectresult.Put()
	return nil
}
func get_block_ajaxReset(data *TemplateData) (err error) {
	if data.ws.Query("confirm") != "yes" {
		data.ws.WriteString(js.Confirm(data.Lang["block"]["confirmReset"].(string), createLink("block", "ajaxReset", "module="+data.ws.Query("module")+"&confirm=yes"), ""))
		return
	}
	session, e := data.BeginTransaction()
	defer func() {
		if e == nil {
			session.Commit()
		} else {
			data.ws.WriteString(js.Error(e.Error()))
			session.Rollback()
		}
	}()
	delectConfig := protocol.GET_MSG_USER_config_save()
	delectConfig.Type = "delete"
	delectConfig.Uid = data.User.Id
	delectConfig.Key = "blockInited"
	delectConfig.Module = data.ws.Query("module")
	delectConfig.Section = "common"
	if e = data.SendMsgWaitResultToDefault(delectConfig, nil); e != nil {
		return
	}
	if e = data.SendMsgWaitResultToDefault(&protocol.MSG_USER_block_delectByWhere{Where: map[string]interface{}{"Module": data.ws.Query("module"), "Uid": data.User.Id}}, nil); e != nil {
		return
	}
	data.ws.WriteString(js.Reload("parent"))
	delectConfig.Put()
	return
}
func printProductBlock(data *TemplateData) (err error) {
	var block protocol.HtmlBlock
	data.ws.Session().Get("temp_block", &block)
	if libraries.Preg_match("/[^a-zA-Z0-9_]/", block.Params["type"]) {
		return
	}
	typ := block.Params["type"]
	num, _ := strconv.Atoi(block.Params["num"])
	data.Page.PerPage = num
	data.Page.Page = 1
	data.Page.Total = -1
	productStats, err := product_getStats(data, "order_desc", typ, 0)
	if err != nil {
		return
	}
	var projects []protocol.HtmlKeyValueStr
	if allproject, err := project_getAll(data); err == nil && len(productStats) > 0 {
		for _, project := range allproject {
			if !data.User.IsAdmin && !data.User.AclProjects[project.Id] || project.Deleted {
				continue
			}
			for _, id := range project.Products {
				for _, product := range productStats {
					if id == product["Id"].(int32) {
						projects = append(projects, protocol.HtmlKeyValueStr{strconv.Itoa(int(id)), project.Name})
					}
				}
			}

		}
	}
	data.Data["productStats"] = productStats
	data.Data["projects"] = projects
	templateOut("block.productblock.html", data)
	return nil
}
func printStoryBlock(data *TemplateData) (err error) {
	var block protocol.HtmlBlock
	data.ws.Session().Get("temp_block", &block)
	if libraries.Preg_match("/[^a-zA-Z0-9_]/", block.Params["type"]) {
		return
	}
	typ := block.Params["type"]
	num, _ := strconv.Atoi(block.Params["num"])
	if typ == "" {
		typ = "assignedTo"
	}
	orderBy := block.Params["orderBy"]
	if orderBy == "" {
		orderBy = "id_asc"
	}
	if data.Data["stories"], err = story_getUserStories(data, data.User.Id, typ, orderBy, &TempLatePage{Total: -1, Page: 1, PerPage: num}); err != nil {
		return
	}
	templateOut("block.storyblock.html", data)
	return nil
}
