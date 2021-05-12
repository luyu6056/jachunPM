package handler

import (
	"errors"
	"fmt"
	"html/template"
	"jachunPM_http/js"
	"libraries"
	"protocol"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func init() {
	httpHandlerModuleInit["GET"]["project"] = project_ModuleInit
	httpHandlerMap["GET"]["/project/create"] = get_project_create
	httpHandlerMap["POST"]["/project/create"] = post_project_create
	httpHandlerMap["GET"]["/project/index"] = get_project_index
	httpHandlerMap["GET"]["/project/tips"] = get_project_tips
	httpHandlerMap["GET"]["/project/ajaxGetDropMenu"] = get_project_ajaxGetDropMenu
	httpHandlerMap["GET"]["/project/all"] = get_project_all
	httpHandlerMap["GET"]["/project/view"] = get_project_view
	httpHandlerMap["GET"]["/project/edit"] = get_project_edit
	httpHandlerMap["POST"]["/project/edit"] = post_project_create
}
func project_ModuleInit(data *TemplateData) (err error) {
	projects, err := project_getPairs(data, "nocode")
	if err != nil {
		return
	}
	//if(!$this->projects and $this->methodName != 'index' and $this->methodName != 'create' and $this->app->getViewType() == 'mhtml') $this->locate($this->createLink('project', 'index'));
	if len(projects) == 0 && data.App["methodName"].(string) != "index" && data.App["methodName"].(string) != "create" {
		data.ws.Redirect(createLink("project", "create", nil))
		return dataErrRedirect
	}

	data.Data["projects"] = projects
	return
}
func project_getAll(data *TemplateData) (result []*protocol.MSG_PROJECT_project_cache, err error) {
	if data.Data["project_getAll"] == nil {
		res, err := HostConn.CacheGetPath(protocol.ProjectServerNo, protocol.PATH_PROJECT_PROJECT_CACHE)
		if err != nil {
			return nil, err
		}

		buf := bufpool.Get().(*libraries.MsgBuffer)
		for _, b := range res {
			buf.Reset()
			buf.Write(b)
			if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_PROJECT_project_cache); ok {
				result = append(result, v)
			}
		}
		buf.Reset()
		bufpool.Put(buf)
		data.Data["project_getAll"] = result
	}

	return data.Data["project_getAll"].([]*protocol.MSG_PROJECT_project_cache), nil
}
func project_getPairs(data *TemplateData, mode string) (res []protocol.HtmlKeyValueStr, err error) {
	list, err := project_getAll(data)
	if err != nil {
		return nil, err
	}
	protocol.Order_project(list, func(a, b *protocol.MSG_PROJECT_project_cache) bool {
		if a.Status == "closed" {
			return false
		}
		if a.Order < b.Order {
			return false
		}
		return true
	})
	if strings.Contains(mode, "empty") {
		res = []protocol.HtmlKeyValueStr{{"", ""}}
	}
	for _, p := range list {
		if p.IsCat && (strings.Contains(mode, "noclosed") && p.Status == "done" || p.Status == "closed") {
			continue
		}
		if data.User.Id != 1 && !data.User.AclProjects[p.Id] {
			continue
		}
		res = append(res, protocol.HtmlKeyValueStr{strconv.Itoa(int(p.Id)), p.Name})
	}
	if len(res) == 0 && len(list) > 0 {
		res = []protocol.HtmlKeyValueStr{{strconv.Itoa(int(list[0].Id)), list[0].Name}}
	}
	return
}
func project_getProjectStats(data *TemplateData, status string, productID, branch int32, itemCounts int, orderBy string, projectID ...int32) (projects []*protocol.MSG_PROJECT_project_cache, err error) {
	if status == "" {
		status = "undone"
	}
	if itemCounts < 1 {
		itemCounts = 30
	}
	/* Init vars. */
	if len(projectID) == 1 {
		projects = []*protocol.MSG_PROJECT_project_cache{HostConn.GetProjectById(projectID[0])}
	} else {
		if projects, err = project_getList(data, status, 0, productID, branch); err != nil {
			return
		}
	}

	data.Page.Total = len(projects)
	begin := (data.Page.Page - 1) * data.Page.PerPage
	end := data.Page.Page * data.Page.PerPage
	if begin >= len(projects) {
		return
	}
	if end > len(projects) {
		end = len(projects)
	}
	orderFunc := func(a, b *protocol.MSG_PROJECT_project_cache) bool {
		if a.Order == b.Order {
			return a.Id > b.Id
		}
		return a.Order > b.Order
	}
	switch orderBy {
	case "order_desc", "":
	default:
		libraries.DebugLog("project_getProjectStats未设置排序func %s", orderBy)
	}
	protocol.Order_project(projects, orderFunc)
	projects = projects[begin:end]
	var project_ids = make([]int32, len(projects))
	for k, p := range projects {
		project_ids[k] = p.Id
	}
	hours := map[int32]map[string]float64{}

	/* Get all tasks and compute totalEstimate, totalConsumed, totalLeft, progress according to them. */
	out := protocol.GET_MSG_PROJECT_task_getListByWhereMap()
	out.Where = map[string]interface{}{"Project": project_ids, "Parent": []interface{}{"lt", 1}, "Deleted": false}
	out.Total = -1
	var result *protocol.MSG_PROJECT_task_getListByWhereMap_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	out.Put()
	/* Compute totalEstimate, totalConsumed, totalLeft. */
	for _, task := range result.List {
		if hours[task.Project] == nil {
			hours[task.Project] = map[string]float64{"totalEstimate": 0, "totalConsumed": 0, "totalLeft": 0, "progress": 0}
		}
		hour := hours[task.Project]
		if task.Status != "cancel" {
			hour["totalEstimate"] += task.Estimate
			hour["totalConsumed"] += task.Consumed
		}
		if task.Status != "cancel" && task.Status != "closed" {
			hour["totalLeft"] += task.Left
		}
	}
	result.Put()
	/* Compute totalReal and progress. */
	for k, hour := range hours {
		hour["totalEstimate"], _ = strconv.ParseFloat(fmt.Sprintf("%.1f", hour["totalEstimate"]), 64)
		hour["totalConsumed"], _ = strconv.ParseFloat(fmt.Sprintf("%.1f", hour["totalConsumed"]), 64)
		hour["totalLeft"], _ = strconv.ParseFloat(fmt.Sprintf("%.1f", hour["totalLeft"]), 64)
		hour["totalReal"] = hour["totalConsumed"] + hour["totalLeft"]
		if hour["totalReal"] > 0 {
			hour["progress"], _ = strconv.ParseFloat(fmt.Sprintf("%.3f", hour["totalConsumed"]/hour["totalReal"]), 64)
			hour["progress"] *= 100
		} else {
			hour["progress"] = 0
		}
		hours[k] = hour
	}
	/* Get burndown charts datas.
	getBurn:=protocol.GET_MSG_PROJECT_project_getBurn()
	getBurn.ProjectIds=project_ids
	var getBurnResult *protocol.MSG_PROJECT_project_getBurn_result
	if err=data.SendMsgWaitResultToDefault(getBurn, &getBurnResult);err!=nil{
		return
	}
	getBurn.Put()
	      for _,projectBurn:=range getBurnResult.List{

	      }

	       foreach(burns as projectID => projectBurns)
	       {
	           /* If projectBurns > itemCounts, split it, else call processBurnData() to pad burns.
	           begin = projects[projectID]->begin;
	           end   = projects[projectID]->end;
	           projectBurns = this->processBurnData(projectBurns, itemCounts, begin, end);

	           /* Shorter names.
	           foreach(projectBurns as projectBurn)
	           {
	               projectBurn->name = substr(projectBurn->name, 5);
	               unset(projectBurn->project);
	           }

	           ksort(projectBurns);
	           burns[projectID] = projectBurns;
	       }*/
	/* Process projects. */
	for _, project := range projects {

		/* Judge whether the project is delayed. */
		if project.Status != "done" && project.Status != "closed" && project.Status != "suspended" {
			if delay := (time.Now().Unix() - project.End.Unix()) / 86400; delay > 0 {
				project.Delay = delay
			}

		}

		/* Process the hours. */
		project.Hours = hours[project.Id]
	}
	return
}
func project_getList(data *TemplateData, status string, limit int, productID, branch int32) (list []*protocol.MSG_PROJECT_project_cache, err error) {

	if cache, err := project_getAll(data); err != nil {
		return nil, err
	} else {

		for _, project := range cache {
			if project.Deleted || project.IsCat || (data.User.Id != 1 && !data.User.AclProjects[project.Id]) {
				continue
			}
			if branch > 0 {
				find := false
				for _, id := range project.Branchs {
					if id == branch {
						find = true
						break
					}
				}
				if !find {
					continue
				}
			}
			if productID > 0 {
				find := false
				for _, id := range project.Products {
					if id == productID {
						find = true
						break
					}
				}
				if !find {
					continue
				}
			}
			switch status {
			case "involved":
				//检查team和openedBy
				if project.OpenedBy == data.User.Id {
					list = append(list, project)
				} else {
					for _, teamUser := range project.Teams {
						if teamUser.Uid == data.User.Id {
							list = append(list, project)
							break
						}
					}
				}
			case "undone":
				if project.Status != "done" && project.Status != "closed" {
					list = append(list, project)
				}
			case "all":
				list = append(list, project)
			default:
				for _, s := range strings.Split(status, ",") {
					if project.Status == s {
						list = append(list, project)
					}
				}
			}

		}
		protocol.Order_project(list, nil)

	}
	if limit > 0 && len(list) > limit {
		list = list[:limit]
	}
	return

}
func get_project_create(data *TemplateData) (err error) {
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	copyProjectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	planID, _ := strconv.Atoi(data.ws.Query("planID"))
	confirm := data.ws.Query("confirm") == "yes"
	if projectID > 0 {
		if planID > 0 {
			if confirm {
				return project_linkStories(data)
			} else {
				js.Confirm(data.Lang["project"]["importPlanStory"].(string), createLink("project", "create", "projectID="+strconv.Itoa(projectID)+"&copyProjectID=&planID="+strconv.Itoa(planID)+"&confirm=yes"), createLink("project", "create", "projectID="+strconv.Itoa(projectID)), "parent", "parent")
				return
			}
		}
		data.Data["title"] = data.Lang["project"]["tips"].(string)

		data.Data["tips"] = template.HTML(common_fetch(data, "project", "tips", "projectID="+strconv.Itoa(projectID)))
		data.Data["projectID"] = projectID
		templateOut("project.create.html", data)
		return
	}

	var products []*protocol.MSG_PROJECT_product_cache
	acl := "private"
	var productPlan []protocol.HtmlKeyValueStr
	productPlans := []protocol.HtmlKeyValueInterface{}
	data.Data["name"] = ""
	data.Data["code"] = ""
	data.Data["team"] = ""
	if copyProjectID > 0 {
		if copyProject := HostConn.GetProjectById(int32(copyProjectID)); copyProject != nil {
			data.Data["name"] = copyProject.Name
			data.Data["code"] = copyProject.Code
			data.Data["team"] = copyProject.Team
			acl = copyProject.Acl
			data.Data["whitelist"] = copyProject.Whitelist
			for _, id := range copyProject.Products {
				if product := HostConn.GetProductById(id); product != nil {
					products = append(products, product)
					if pairs, err := productplan_getPairs(data, id, 0, ""); err != nil {
						return err
					} else {
						productPlans = append(productPlans, protocol.HtmlKeyValueInterface{Key: strconv.Itoa(int(id)), Value: pairs})
					}
				}

			}
		}
	}
	data.Data["productPlans"] = productPlans
	if planID > 0 {
		out := protocol.GET_MSG_PROJECT_productplan_getById()
		out.Id = int32(planID)
		var result *protocol.MSG_PROJECT_productplan_getById_result
		if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
			return
		}
		if result.Info == nil {
			return errors.New(protocol.Err_ProjectProductPlanNotFound.String())
		}
		data.Data["plan"] = result.Info
		if product := HostConn.GetProductById(result.Info.Product); product != nil {
			products = append(products, product)
		}
		if productPlan, err = productplan_getPairs(data, result.Info.Product, 0, "unexpired"); err != nil {
			return
		}
		out.Put()
		result.Put()
	}
	data.Data["acl"] = acl
	if id := data.ws.Session().Load_int32("project"); id > 0 {
		projectID = int(id)
	} else if data.Data["projects"] != nil && len(data.Data["projects"].([]protocol.HtmlKeyValueStr)) > 0 {
		projectID, _ = strconv.Atoi(data.Data["projects"].([]protocol.HtmlKeyValueStr)[0].Key)
	}

	if err = project_setMenu(data, int32(projectID), 0, ""); err != nil {
		return
	}
	data.Data["title"] = data.Lang["project"]["create"]
	if data.Data["projects"] == nil {
		data.Data["projects"] = []protocol.HtmlKeyValueStr{{}}
	} else {
		data.Data["projects"] = append([]protocol.HtmlKeyValueStr{{}}, data.Data["projects"].([]protocol.HtmlKeyValueStr)...)
	}

	if data.Data["groups"], err = user_getGroupGetPairs(); err != nil {
		return
	}
	allProducts, err := product_getPairs(data, "noclosed|nocode")
	if err != nil {
		return
	}
	data.Data["allProducts"] = append([]protocol.HtmlKeyValueStr{{"0", ""}}, allProducts...)

	data.Data["projectID"] = projectID
	data.Data["products"] = products
	data.Data["productPlan"] = append([]protocol.HtmlKeyValueStr{{"0", ""}}, productPlan...)

	data.Data["productPlans"] = append([]protocol.HtmlKeyValueInterface{{"0", nil}}, productPlans...)

	data.Data["copyProjectID"] = copyProjectID
	var ids []int32
	for _, p := range products {
		ids = append(ids, p.Id)
	}
	if data.Data["branchGroups"], err = branch_getByProducts(data, ids, "", nil); err != nil {
		return
	}
	templateOut("project.create.html", data)
	return
}
func project_linkStories(data *TemplateData) (err error) {
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	project := HostConn.GetProjectById(int32(projectID))
	if project == nil {
		data.ws.WriteString(js.Alert(data.Lang["project"]["error"].(map[string]string)["NotFount"]))
		return dataErrAlreadyOut
	}

	out := protocol.GET_MSG_PROJECT_project_linkStory()
	count := 0
	for _, planID := range project.Plans {
		out1 := protocol.GET_MSG_PROJECT_story_getPlanStories()
		out1.PlanID = planID
		var result *protocol.MSG_PROJECT_story_getPlanStories_result
		if err = data.SendMsgWaitResultToDefault(out1, &result); err != nil {
			return
		}
		if len(result.List) > 0 {

			for _, story := range result.List {
				if story.Status == "draft" {
					count++
					continue
				}
				out.Stories = append(out.Stories, story.Id)
				out.Products[story.Id] = story.Product
			}

		}
		out1.Put()
		result.Put()
	}
	if len(out.Stories) > 0 {
		out.ProjectID = project.Id
		if err = data.SendMsgToDefault(out); err != nil {
			return
		}
	}
	out.Put()

	if count != 0 {
		data.ws.WriteString(js.Alert(data.Lang["project"]["haveDraft"].(string), count) + js.Location(createLink("project", "create", "projectID="+strconv.Itoa(projectID)), ""))
		return dataErrAlreadyOut
	}
	return
}
func project_setMenu(data *TemplateData, projectID, buildID int32, extra string) error {
	/* Check the privilege. */
	/* project := HostConn.GetProjectById(projectID)

	Unset story, bug, build and testtask if type is ops.
	   if(project and project->type == "ops"){
	       unset(this->lang->project->menu->story);
	       unset(this->lang->project->menu->qa);
	       unset(this->lang->project->subMenu->qa->bug);
	       unset(this->lang->project->subMenu->qa->build);
	       unset(this->lang->project->subMenu->qa->testtask);
	   }*/
	projects, _ := data.Data["projects"].([]protocol.HtmlKeyValueStr)
	if len(projects) > 0 {
		find := false
		for _, kv := range projects {
			if kv.Key == strconv.Itoa(int(projectID)) {
				find = true
				break
			}
		}
		if !find {
			data.ws.WriteString(js.Alert(data.Lang["project"]["error"].(map[string]string)["NotFount"]))
			if strings.Contains(data.ws.Header("Referer"), "/user/login") {
				data.ws.WriteString(js.Location(createLink("project", "index", nil), ""))
			} else {
				data.ws.WriteString(js.Location("back", ""))
			}
			return dataErrAlreadyOut
		}
		if data.User.Id != 1 && !data.User.AclProjects[projectID] {
			data.ws.WriteString(js.Alert(data.Lang["project"]["accessDenied"].(string)))
			if strings.Contains(data.ws.Header("Referer"), "/user/login") {
				data.ws.WriteString(js.Location(createLink("project", "index", nil), ""))
			} else {
				data.ws.WriteString(js.Location("back", ""))
			}
			return dataErrAlreadyOut
		}

	}

	moduleName := data.App["moduleName"].(string)
	methodName := data.App["methodName"].(string)

	selectHtml, err := project_select(data, projectID, buildID, moduleName, methodName, extra)
	if err != nil {
		return err
	}
	label := data.Lang["project"]["index"].(string)
	if moduleName == "project" && methodName == "all" {
		label = data.Lang["project"]["allProjects"].(string)
	}
	if moduleName == "project" && methodName == "create" {
		label = data.Lang["project"]["create"].(string)
	}

	buf := bufpool.Get().(*libraries.MsgBuffer)
	/*isMobile     = this->app->viewType == "mhtml";
	  if(isMobile){
	      buf  = html::a(helper::createLink("project", "index"), this->lang->project->index) . this->lang->colon;
	      buf .= selectHtml;
	  }else{*/
	buf.WriteString("<div class='btn-group angle-btn")
	if methodName == "index" {
		buf.WriteString(" active")
	}
	buf.WriteString("'><div class='btn-group'><button data-toggle='dropdown' type='button' class='btn'>")
	buf.WriteString(label)
	buf.WriteString(" <span class='caret'></span></button><ul class='dropdown-menu'>")
	if hasPriv(data, "project", "index") {
		buf.WriteString("<li>")
		buf.WriteString(html_a(createLink("project", "index", "locate=no"), "<i class='icon icon-home'></i> "+data.Lang["project"]["index"].(string)))
		buf.WriteString("</li>")
	}
	if hasPriv(data, "project", "all") {
		buf.WriteString("<li>")
		buf.WriteString(html_a(createLink("project", "all", "status=all"), "<i class='icon icon-cards-view'></i> "+data.Lang["project"]["allProjects"].(string)))
		buf.WriteString("</li>")
	}
	if hasPriv(data, "project", "create") {
		buf.WriteString("<li>")
		buf.WriteString(html_a(createLink("project", "create", "status=all"), "<i class='icon icon-plus'></i> "+data.Lang["project"]["create"].(string)))
		buf.WriteString("</li>")
	}

	buf.WriteString("</ul></div></div>")

	buf.WriteString(selectHtml)
	//}

	data.Data["modulePageNav"] = template.HTML(buf.String())
	buf.Reset()
	bufpool.Put(buf)
	//if(moduleName != "project") this->lang->moduleName->dividerMenu = this->lang->project->dividerMenu;
	return nil
}
func project_select(data *TemplateData, projectID, buildID int32, currentModule, currentMethod, extra string) (res string, err error) {
	currentProject := HostConn.GetProjectById(projectID)
	if currentProject == nil {
		return
	}

	//isMobile = this->app->viewType == 'mhtml';

	data.ws.SetCookie("lastProject", strconv.Itoa(int(projectID)), 0)

	dropMenuLink := createLink("project", "ajaxGetDropMenu", []interface{}{"objectID=", projectID, "&module=", currentModule, "&method=", currentMethod, "&extra=", extra})
	buf := bufpool.Get().(*libraries.MsgBuffer)

	buf.WriteString("<div class='btn-group angle-btn'><div class='btn-group'><button data-toggle='dropdown' type='button' class='btn btn-limit' id='currentItem' title='")
	buf.WriteString(currentProject.Name)

	buf.WriteString("'>")
	buf.WriteString(currentProject.Name)
	buf.WriteString("<span class='caret'></span></button><div id='dropMenu' class='dropdown-menu search-list' data-ride='searchList' data-url='")
	buf.WriteString(dropMenuLink)
	buf.WriteString("'>")
	buf.WriteString(`<div class="input-control search-box has-icon-left has-icon-right search-example"><input type="search" class="form-control search-input" /><label class="input-control-icon-left search-icon"><i class="icon icon-search"></i></label><a class="input-control-icon-right search-clear-btn"><i class="icon icon-close icon-sm"></i></a></div>`)
	buf.WriteString("</div></div></div>")
	/*if(isMobile) output  = "<a id='currentItem' href=\"javascript:showSearchMenu('project', 'projectID', 'currentModule', 'currentMethod', 'extra')\">{currentProject->name} <span class='icon-caret-down'></span></a><div id='currentItemDropMenu' class='hidden affix enter-from-bottom layer'></div>";*/

	//if(buildID and !isMobile){
	if buildID > 0 {
		data.ws.SetCookie("lastBuild", strconv.Itoa(int(buildID)), 0)
		out := protocol.GET_MSG_PROJECT_build_getById()
		out.Id = buildID
		var result *protocol.MSG_PROJECT_build_getById_result
		if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
			return
		}

		if result.Info != nil {
			dropMenuLink = createLink("build", "ajaxGetProjectBuilds", []interface{}{"projectID=", projectID, "&productID=&varName=dropdownList"})
			buf.WriteString("<div class='btn-group angle-btn'><div class='btn-group'><button data-toggle='dropdown' type='button' class='btn btn-limit' id='currentItem'>")
			buf.WriteString(result.Info.Name)
			buf.WriteString("<span class='caret'></span></button><div id='dropMenu' class='dropdown-menu search-list' data-ride='searchList' data-url='")
			buf.WriteString(dropMenuLink)
			buf.WriteString("'>")
			buf.WriteString(`<div class="input-control search-box has-icon-left has-icon-right search-example"><input type="search" class="form-control search-input" /><label class="input-control-icon-left search-icon"><i class="icon icon-search"></i></label><a class="input-control-icon-right search-clear-btn"><i class="icon icon-close icon-sm"></i></a></div></div></div></div>`)
		}
		out.Put()
		result.Put()
	}
	res = buf.String()
	buf.Reset()
	bufpool.Put(buf)
	return
}
func get_project_index(data *TemplateData) (err error) {
	if data.ws.Query("locate") == "yes" {
		data.ws.Redirect(createLink("project", "task", nil))
		return
	}
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	//if($this->app->viewType != 'mhtml') unset($this->lang->project->menu->index);
	if _, err = project_commonAction(data, int32(projectID)); err != nil {
		return
	}

	data.Data["title"] = data.Lang["project"]["index"]

	templateOut("project.index.html", data)
	return
}
func project_commonAction(data *TemplateData, projectID int32, ext ...string) (project *protocol.MSG_PROJECT_project_cache, err error) {
	var extra string
	if len(ext) > 0 {
		extra = ext[0]
	}
	/* Get projects and products info. */
	if projectID, err = project_saveState(data, projectID); err != nil {
		return
	}
	project = HostConn.GetProjectById(projectID)
	var products []*protocol.MSG_PROJECT_product_cache
	if project != nil {
		for _, id := range project.Products {
			if product := HostConn.GetProductById(id); product != nil {
				products = append(products, product)
			}
		}
	}

	if data.Data["childProjects"], err = project_getChildProjects(data, projectID); err != nil {
		return
	}
	if data.Data["teamMembers"], err = project_getTeamMembers(data, projectID); err != nil {
		return
	}
	/* Set menu. */
	if err = project_setMenu(data, projectID, 0, extra); err != nil {
		return
	}

	/* Assign. */
	data.Data["project"] = project
	data.Data["products"] = products
	return project, nil
}
func project_saveState(data *TemplateData, projectID int32) (int32, error) {
	projects, _ := data.Data["projects"].([]protocol.HtmlKeyValueStr)
	if projectID > 0 {
		data.ws.Session().Set("project", projectID)
	} else {
		if data.ws.Cookie("lastProject") != "" {
			data.ws.Session().Set("project", data.ws.Cookie("lastProject"))
		} else if data.ws.Session().Load_int("project") == 0 {
			if len(projects) > 0 {
				data.ws.Session().Set("project", projects[0].Key)
			}
		}
	}
	var find bool
	for _, kv := range projects {
		if kv.Key == strconv.Itoa(int(projectID)) {
			find = true
			break
		}
	}
	if !find {
		if len(projects) > 0 {
			data.ws.Session().Set("project", projects[0].Key)
		}
		if projectID > 0 {
			data.ws.WriteString(js.Alert(data.Lang["project"]["accessDenied"].(string)))

			if strings.Contains(data.ws.Path(), createLink("user", "login", nil)) {
				js.Location(createLink("my", "index", nil), "self")
			} else {
				js.Location("back", "self")
			}

			return 0, dataErrAlreadyOut
		}
	}
	return data.ws.Session().Load_int32("project"), nil

}
func project_getChildProjects(data *TemplateData, projectID int32) (res []protocol.HtmlKeyValueStr, err error) {
	list, err := project_getAll(data)
	if err != nil {
		return
	}
	protocol.Order_project(list, nil)
	for _, p := range list {
		if p.Parent == projectID {
			res = append(res, protocol.HtmlKeyValueStr{strconv.Itoa(int(p.Id)), p.Name})
		}
	}
	return
}
func project_getTeamMembers(data *TemplateData, projectID int32) ([]*protocol.MSG_USER_team_info, error) {
	out := protocol.GET_MSG_USER_team_getByTypeRoot()
	out.Type = "project"
	out.Root = projectID
	var result *protocol.MSG_USER_team_getByTypeRoot_result
	if err := data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return nil, err
	}
	for _, v := range result.List {
		if user := HostConn.GetUserCacheById(v.Uid); user != nil {
			v.Realname = user.Realname

		}
	}
	return result.List, nil
}
func post_project_create(data *TemplateData) (err error) {
	if !data.ajaxCheckPost() {
		return
	}
	projectID, _ := strconv.Atoi(data.ws.Query("project"))
	copyProjectID, _ := strconv.Atoi(data.ws.Query("copyProjectID"))
	out := protocol.GET_MSG_PROJECT_project_create()
	var project *protocol.MSG_PROJECT_project_cache
	if data.App["methodName"] == "edit" { //修改
		project = HostConn.GetProjectById(int32(projectID))
		if project == nil {
			data.ws.WriteString(js.Alert(data.Lang["project"]["error"].(map[string]string)["NotFount"]))
			return
		}
	} else {
		project = protocol.GET_MSG_PROJECT_project_cache()
	}

	for key, value := range data.ws.GetAllPost() {
		switch key {
		case "name":
			project.Name = value[0]
		case "code":
			project.Code = value[0]
		case "begin":
			project.Begin, err = time.Parse(protocol.TIMEFORMAT_MYSQLDATE, value[0])
			if err != nil {
				data.ajaxResult(false, map[string]string{"begin": data.Lang["project"]["error"].(map[string]string)["beginTime"]}, "")
				return nil
			}
		case "end":
			project.End, err = time.Parse(protocol.TIMEFORMAT_MYSQLDATE, value[0])
			if err != nil {
				data.ajaxResult(false, map[string]string{"begin": data.Lang["project"]["error"].(map[string]string)["endTime"]}, "")
				return nil
			}
		case "team":
			project.Team = value[0]
		case "type":
			project.Type = value[0]
		case "desc":
			project.Desc = libraries.Html2bbcode(value[0])
		case "acl":
			project.Acl = value[0]
		case "whitelist":
			for _, v := range value {
				id, _ := strconv.Atoi(v)
				if id > 0 {
					project.Whitelist = append(project.Whitelist, int32(id))
				}
			}
		}
	}

	if project.Begin.Unix() > project.End.Unix() {
		data.ajaxResult(false, map[string]string{"begin": data.Lang["project"]["error"].(map[string]string)["beginGeEnd"]}, "")
		return nil
	}
	i := 0
	for {
		productkey := "products[" + strconv.Itoa(i) + "]"
		productID, _ := strconv.Atoi(data.ws.Post(productkey))
		if productID == 0 {
			break
		}
		product := HostConn.GetProductById(int32(productID))
		if product == nil {
			data.ajaxResult(false, map[string]string{productkey: data.Lang["project"]["error"].(map[string]string)["CreateNotFountProduct"]}, "")
			return nil
		}
		project.Products = append(project.Products, product.Id)
		branchkey := "branch[" + strconv.Itoa(i) + "]"
		branchID, _ := strconv.Atoi(data.ws.Post(branchkey))
		if branchID > 0 {
			var find bool
			for _, id := range product.Branch {
				if id == int32(branchID) {
					project.Branchs = append(project.Branchs, id)
					find = true
					break
				}
			}
			if !find {
				data.ajaxResult(false, map[string]string{branchkey: data.Lang["project"]["error"].(map[string]string)["CreateNotFountProduct"]}, "")
				return nil
			}
		}

		planskey := "plans[" + strconv.Itoa(int(product.Id)) + "]"
		planID, _ := strconv.Atoi(data.ws.Post(planskey))
		if planID > 0 {
			var find bool
			for _, id := range product.Plan {
				if id == int32(planID) {
					project.Plans = append(project.Plans, id)
					find = true
					break
				}
			}
			if !find {
				data.ajaxResult(false, map[string]string{planskey: data.Lang["project"]["error"].(map[string]string)["CreateNotFountPlan"]}, "")
				return nil
			}
		}

		i++
	}
	project.OpenedBy = data.User.Id
	project.OpenedDate = time.Now()
	if project.Team == "" {
		project.Team = project.Name
	}
	out.Info = project
	out.CopyProjectID = int32(copyProjectID)

	var result *protocol.MSG_PROJECT_project_create_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		if v, ok := data.Lang["project"]["error"].(map[string]string)[err.Error()]; ok {
			data.ajaxResult(false, v, "")
		} else {
			data.ajaxResult(false, err.Error(), "")
		}

		return nil
	}
	if projectID > 0 { //修改
		data.ajaxResult(true, data.Lang["common"]["saveSuccess"], createLink("project", "view", []interface{}{"projectID=", projectID}))
	} else {
		if len(project.Plans) > 0 {
			data.ajaxResult(true, data.Lang["common"]["saveSuccess"], createLink("project", "create", []interface{}{"projectID=", result.Id, "&copyProjectID=&planID=", project.Plans[0], "&confirm=no"}))
		} else {
			data.ajaxResult(true, data.Lang["common"]["saveSuccess"], createLink("project", "create", []interface{}{"projectID=", result.Id}))
		}
	}

	out.Put()
	result.Put()
	return
	/*this->project->updateProducts($projectID);


	  fileType =_POST['fileType'];

	  files =this->loadModel('file')->saveUpload('project',projectID,'','files','labels',$fileType); //2019.1.16 luke 新增fileType

	   if(dao::isError())this->send(array('result' => 'fail', 'message' => dao::getError()));

	  this->loadModel('action')->create('project',projectID, 'opened');

	  planID = reset($_POST['plans']);
	   if(!empty($planID))
	   {
	      this->send(array('result' => 'success', 'message' =>this->lang->saveSuccess, 'locate' => inlink('create', "projectID=$projectID&copyProjectID=&planID=$planID&confirm=no")));
	   }
	   else
	   {
	      this->send(array('result' => 'success', 'message' =>this->lang->saveSuccess, 'locate' => inlink('create', "projectID=$projectID")));
	   }*/
}
func get_project_tips(data *TemplateData) (err error) {
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	data.Data["project"] = HostConn.GetProjectById(int32(projectID))
	data.Data["projectID"] = projectID
	templateOut("project.tips.html", data)
	return
}
func get_project_ajaxGetDropMenu(data *TemplateData) (err error) {
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	module := data.ws.Query("module")
	method := data.ws.Query("method")
	extra := data.ws.Query("extra")
	link := project_getProjectLink(module, method, extra)
	data.Data["projectID"] = projectID
	data.Data["module"] = module
	data.Data["method"] = method
	data.Data["extra"] = extra
	list, err := project_getAll(data)
	if err != nil {
		return
	}
	protocol.Order_project(list, nil)
	myProjectsHtml := bufpool.Get().(*libraries.MsgBuffer)
	normalProjectsHtml := bufpool.Get().(*libraries.MsgBuffer)
	closedProjectsHtml := bufpool.Get().(*libraries.MsgBuffer)
	for _, project := range list {
		if (project.Status != "done") && (project.Status != "closed") {
			if project.PM == data.User.Id {
				myProjectsHtml.WriteString(html_a(fmt.Sprintf(link, project.Id), "<i class='icon icon-folder-outline'></i> "+project.Name, "", "class='text-important' title='"+project.Name+"' data-key='"+project.Code+"'"))

			} else {
				normalProjectsHtml.WriteString(html_a(fmt.Sprintf(link, project.Id), "<i class='icon icon-folder-outline'></i> "+project.Name, "", "title='"+project.Name+"' data-key='"+project.Code+"'"))
			}
		} else {
			closedProjectsHtml.WriteString(html_a(fmt.Sprintf(link, project.Id), "<i class='icon icon-folder-outline'></i> "+project.Name, "", "title='"+project.Name+"' data-key='"+project.Code+"'"))

		}
	}
	data.Data["myProjectsHtml"] = template.HTML(myProjectsHtml.String())
	data.Data["normalProjectsHtml"] = template.HTML(normalProjectsHtml.String())
	data.Data["closedProjectsHtml"] = template.HTML(closedProjectsHtml.String())
	templateOut("project.ajaxGetDropMenu.html", data)
	myProjectsHtml.Reset()
	normalProjectsHtml.Reset()
	closedProjectsHtml.Reset()
	bufpool.Put(myProjectsHtml)
	bufpool.Put(normalProjectsHtml)
	bufpool.Put(closedProjectsHtml)
	return
}
func project_getProjectLink(module, method, extra string) (link string) {

	if module == "task" && (method == "view" || method == "edit" || method == "batchedit") {
		module = "project"
		method = "task"
	}
	if module == "build" && (method == "edit" || method == "view") {
		module = "project"
		method = "build"
	}

	if module == "project" && method == "create" {
		return
	}
	if extra != "" {
		link = createLink(module, method, "projectID=%d&type="+extra)
	} else if module == "project" && (method == "index" || method == "all") {
		link = createLink(module, "task", "projectID=%d")
	} else {
		link = createLink(module, method, "projectID=%d")
	}

	if module == "doc" {
		link = createLink("doc", "objectLibs", "type=project&objectID=%d&from=project")
	}
	return
}
func get_project_all(data *TemplateData) (err error) {
	status := data.ws.Query("status")
	if status == "" {
		status = "undone"
	}
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	orderBy := data.ws.Query("orderBy")
	if orderBy == "" {
		orderBy = "order_desc"
	}
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	if _, ok := data.Data["projects"].([]protocol.HtmlKeyValueStr); ok {
		if project, err := project_commonAction(data, int32(projectID)); err != nil {
			return err
		} else {
			projectID = int(project.Id)
		}

	}
	data.ws.Session().Set("projectList", data.ws.Path())
	data.Data["title"] = data.Lang["project"]["allProject"]
	getProjectStatsStatus := status
	if getProjectStatsStatus == "byproduct" {
		getProjectStatsStatus = "all"
	}
	if data.Data["projectStats"], err = project_getProjectStats(data, getProjectStatsStatus, int32(productID), 0, 30, orderBy); err != nil {
		return
	}
	if products_pairs, err := product_getPairs(data, ""); err != nil {
		return err
	} else {
		data.Data["products"] = append([]protocol.HtmlKeyValueStr{{"0", data.Lang["product"]["select"].(string)}}, products_pairs...)
	}

	data.Data["productID"] = productID
	data.Data["projectID"] = projectID
	data.Data["orderBy"] = orderBy
	if data.Data["users"], err = user_getPairs(data, "noletter"); err != nil {
		return
	}
	data.Data["status"] = status
	templateOut("project.all.html", data)
	return
}
func get_project_view(data *TemplateData) (err error) {

	projectID, _ := strconv.Atoi(data.ws.Query("project"))
	project := HostConn.GetProjectById(int32(projectID))

	if project == nil {
		data.ws.WriteString(js.Error(data.Lang["common"]["notFound"].(string)) + js.Location("back", ""))
		return

	}
	productID := int32(0)
	if len(project.Products) > 0 {
		productID = project.Products[0]
	}
	//补充hours信息
	if list, err := project_getProjectStats(data, "", productID, 0, 1, "", project.Id); err != nil {
		return err
	} else {
		project = list[0]
	}
	msg, _ := data.GetMsg()
	msg.ActionLogHistory(0, project, project)
	products := project_getProducts(data, project.Id)
	productIds := make([]int32, len(products))
	for k, p := range products {
		productIds[k] = p.Id
	}

	/* Set menu. */
	if err = project_setMenu(data, project.Id, 0, ""); err != nil {
		return
	}

	//dateList, interval := project_getDateList(data, project.Begin, project.End, "noweekend", 0, "2006-01-02")
	//chartData = this->project->buildBurnData(projectID, dateList, "noweekend");

	/* Load pager. */

	data.Data["title"] = data.Lang["project"]["view"]
	data.Data["project"] = project
	data.Data["products"] = products
	planGroups := make(map[int32][]protocol.HtmlKeyValueStr)
	for _, p := range products {
		out := protocol.GET_MSG_PROJECT_productplan_getPairs()
		out.ProductID = p.Id

		for _, b := range p.Branchs {
			out.BranchID = b.Id
			var result *protocol.MSG_PROJECT_productplan_getPairs_result
			if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
				return
			}
			planGroups[p.Id] = append(planGroups[p.Id], result.List...)
		}

		out.Put()
	}
	data.Data["planGroups"] = planGroups
	out := protocol.GET_MSG_USER_Group_getPairs()
	var result *protocol.MSG_USER_Group_getPairs_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	data.Data["groups"] = result.List
	if data.Data["actions"], err = action_getList(data, "project", int32(projectID)); err != nil {
		return
	}
	if data.Data["dynamics"], err = action_getDynamic(data, actionGetDynamicParamAll, "all", "date_desc", TempLatePage{PerPage: 30, Page: 1}, actionGetDynamicParamAll, int32(projectID)); err != nil {
		return
	}
	users, err := user_getPairs(data, "noletter")
	if err != nil {
		return
	}
	data.Data["users"] = users
	getTeam := protocol.GET_MSG_USER_team_getByTypeRoot()
	getTeam.Type = "project"
	getTeam.Root = int32(projectID)
	var getTeamResult *protocol.MSG_USER_team_getByTypeRoot_result
	if err = data.SendMsgWaitResultToDefault(getTeam, &getTeamResult); err != nil {
		return
	}
	var maxTeamView = 9
	if project.PM > 0 {
		for _, kv := range users {
			if kv.Key == strconv.Itoa(int(project.PM)) {
				data.Data["PM"] = kv.Value
				maxTeamView--
			}
		}

	}
	if project.PO > 0 {
		for _, kv := range users {
			if kv.Key == strconv.Itoa(int(project.PM)) {
				data.Data["PO"] = kv.Value
				maxTeamView--
			}
		}

	}
	if project.QD > 0 {
		for _, kv := range users {
			if kv.Key == strconv.Itoa(int(project.PM)) {
				data.Data["QD"] = kv.Value
				maxTeamView--
			}
		}

	}
	if project.RD > 0 {
		for _, kv := range users {
			if kv.Key == strconv.Itoa(int(project.PM)) {
				data.Data["RD"] = kv.Value
				maxTeamView--
			}
		}

	}
	for i := len(getTeamResult.List) - 1; i >= 0; i-- {
		t := getTeamResult.List[i]
		if project.PM == t.Uid {
			getTeamResult.List = append(getTeamResult.List[:i], getTeamResult.List[i+1:]...)
			break
		}
		if project.PO == t.Uid {
			getTeamResult.List = append(getTeamResult.List[:i], getTeamResult.List[i+1:]...)
			break
		}
		if project.QD == t.Uid {
			getTeamResult.List = append(getTeamResult.List[:i], getTeamResult.List[i+1:]...)
			break
		}
		if project.RD == t.Uid {
			getTeamResult.List = append(getTeamResult.List[:i], getTeamResult.List[i+1:]...)
			break
		}
	}
	if len(getTeamResult.List) > maxTeamView {
		getTeamResult.List = getTeamResult.List[:maxTeamView]
	}
	data.Data["teamMembers"] = getTeamResult.List
	//data.Data["docLibs"]      = this->loadModel("doc")->getLibsByObject("project", projectID);
	statRelatedData := protocol.GET_MSG_PROJECT_project_statRelatedData()
	statRelatedData.ProjectID = int32(projectID)
	var statRelatedDataResult *protocol.MSG_PROJECT_project_statRelatedData_result
	if err = data.SendMsgWaitResultToDefault(statRelatedData, &statRelatedDataResult); err != nil {
		return
	}
	data.Data["statData"] = statRelatedDataResult
	//data.Data["chartData"]    = chartData;
	data.Data["blockHistory"] = true
	data.Data["progress"] = project.Hours["progress"]
	if link := data.ws.Session().Load_str("projectList"); link != "" {
		data.Data["browseLink"] = link
	} else {
		data.Data["browseLink"] = createLink("project", "browse", []interface{}{"projectID=", projectID})
	}
	templateOut("project.view.html", data)
	out.Put()
	result.Put()
	getTeam.Put()
	getTeamResult.Put()
	return
}
func project_getProducts(data *TemplateData, projectID int32) (res []*protocol.MSG_PROJECT_product_cache) {
	if project := HostConn.GetProjectById(projectID); project != nil {
		for _, id := range project.Products {
			if product := HostConn.GetProductById(id); product != nil {
				res = append(res, product)
			}
		}
	}
	return
}
func project_getDateList(data *TemplateData, begin, end time.Time, typ string, interval float64, format string) ([]string, int) {

	beginWeekDay := begin.Day()
	days := int((end.Unix() - begin.Unix()) / 3600 / 24)
	if typ == "noweekend" {
		allDays := days
		weekDay := beginWeekDay
		for i := 0; i < allDays; i, weekDay = i+1, weekDay+1 {
			weekDay := weekDay % 7
			if (data.Config["project"]["common"]["weekend"].(int) == 2 && weekDay == 6) || weekDay == 0 {
				days--
			}
		}
	}
	if interval == 0 {
		interval = float64(days / data.Config["project"]["common"]["maxBurnDay"].(int))
	}
	var dateList []string
	spaces := int(interval)
	counter := spaces
	weekDay := beginWeekDay
	for date := begin.Unix(); date <= end.Unix(); date, weekDay = date+24*3600, weekDay+1 {
		/* Remove weekend when type is noweekend.*/
		if typ == "noweekend" {
			weekDay = weekDay % 7
			if (data.Config["project"]["common"]["weekend"].(int) == 2 && weekDay == 6) || weekDay == 0 {
				continue
			}
		}

		counter++
		if counter <= spaces {
			continue
		}

		counter = 0
		t := time.Unix(date, 0)
		dateList = append(dateList, t.Format(format))
	}

	return dateList, int(interval)
}
func get_project_edit(data *TemplateData) (err error) {
	projectID, _ := strconv.Atoi(data.ws.Query("project"))
	/* Set menu. */
	if err = project_setMenu(data, int32(projectID), 0, ""); err != nil {
		return
	}

	projects := append([]protocol.HtmlKeyValueStr{{}}, data.Data["projects"].([]protocol.HtmlKeyValueStr)...)
	for k, kv := range projects {
		if kv.Key == strconv.Itoa(projectID) {
			projects = append(projects[k:], projects[:k+1]...)
			break
		}
	}
	project := HostConn.GetProjectById(int32(projectID))
	products, err := product_getPairs(data, "noclosed|nocode")
	if err != nil {
		return
	}
	allProducts := append([]protocol.HtmlKeyValueStr{{"0", ""}}, products...)
	var linkedProducts []*protocol.MSG_PROJECT_product_cache
	var linkedBranches []int32
	for _, productID := range project.Products {
		if product := HostConn.GetProductById(productID); product != nil {
			linkedProducts = append(linkedProducts, product)
			find := false
			for _, kv := range allProducts {
				if kv.Key == strconv.Itoa(int(productID)) {
					find = true
					break
				}
			}
			if !find {
				allProducts = append(allProducts, protocol.HtmlKeyValueStr{strconv.Itoa(int(productID)), product.Name})
			}

			linkedBranches = append(linkedBranches, product.Branch...)
		}
	}
	productPlans := []protocol.HtmlKeyValueInterface{{"0", nil}}
	for _, product := range linkedProducts {
		pairs, err := productplan_getPairs(data, product.Id, 0, "")
		if err != nil {
			return err
		}
		productPlans = append(productPlans, protocol.HtmlKeyValueInterface{strconv.Itoa(int(product.Id)), pairs})

	}

	data.Data["title"] = data.Lang["project"]["edit"].(string) + data.Lang["common"]["colon"].(string) + project.Name

	data.Data["projects"] = projects
	data.Data["project"] = project
	if data.Data["poUsers"], err = user_getPairs(data, "noclosed|nodeleted|pofirst", project.PO); err != nil {
		return
	}
	if data.Data["pmUsers"], err = user_getPairs(data, "noclosed|nodeleted|pmfirst", project.PM); err != nil {
		return
	}
	if data.Data["qdUsers"], err = user_getPairs(data, "noclosed|nodeleted|qdfirst", project.QD); err != nil {
		return
	}
	if data.Data["rdUsers"], err = user_getPairs(data, "noclosed|nodeleted|devfirst", project.RD); err != nil {
		return
	}
	if data.Data["groups"], err = user_getGroupGetPairs(); err != nil {
		return
	}
	data.Data["allProducts"] = allProducts
	data.Data["linkedProducts"] = linkedProducts
	data.Data["productPlans"] = productPlans
	if data.Data["branchGroups"], err = branch_getByProducts(data, project.Products, "", linkedBranches); err != nil {
		return
	}
	templateOut("project.edit.html", data)
	return
}
func projectFuncs() {
	global_Funcs["MSG_PROJECT_project_cache_isClickable"] = func(data *TemplateData, obj interface{}, action string) bool {
		if project, ok := obj.(*protocol.MSG_PROJECT_project_cache); ok {

			if action == "start" {
				return project.Status == "wait"
			}
			if action == "close" {
				return project.Status != "closed"
			}
			if action == "suspend" {
				return project.Status == "wait" || project.Status == "doing"
			}
			if action == "putoff" {
				return project.Status == "wait" || project.Status == "doing"
			}
			if action == "activate" {
				return project.Status == "suspended" || project.Status == "closed"
			}

			return true
		} else {
			libraries.DebugLog("MSG_PROJECT_project_cache_isClickable传入的值类型%v不对", reflect.TypeOf(obj).Elem().Name())
		}
		return true
	}
}
