package handler

import (
	"errors"
	"fmt"
	"html/template"
	"jachunPM_http/config"
	"jachunPM_http/js"
	"libraries"
	"protocol"
	"reflect"
	"strconv"
	"strings"
)

func init() {

	httpHandlerMap["GET"]["/story/create"] = get_story_create
	httpHandlerMap["POST"]["/story/create"] = post_story_create
	httpHandlerMap["GET"]["/story/view"] = get_story_view
	httpHandlerMap["GET"]["/story/ajaxGetProjectStories"] = get_story_ajaxGetProjectStories

}
func get_story_create(data *TemplateData) (err error) {
	/*var fromObjectIDKey,fromObjectID,fromObjectName,fromObjectAction string
	for key,value := range data.ws.GetAllQuery{
		if v1,ok:=data.Config["story"]["fromObjects"];ok{
			if v2,ok:=v1[key].(map[string]string);ok{
				fromObjectIDKey=key
				fromObjectID=value
				fromObjectName=v2["name"]
				fromObjectAction=v2["action"]
				break
			}
		}
	}
	 if(isset(fromObjectID))
	  {
	      fromObject = this->loadModel(fromObjectName)->getById(fromObjectID)
	      if(!fromObject) die(js::error(this->lang->notFound) . js::locate("back", "parent"))

	      data.Data["fromObjectIDKey"] = fromObjectID
	      data.Data["fromObjectName"]  = fromObject
	  }*/

	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	moduleID, _ := strconv.Atoi(data.ws.Query("moduleID"))
	planID, _ := strconv.Atoi(data.ws.Query("planID"))
	//storyID, _ := strconv.Atoi(data.ws.Query("storyID"))
	branch, _ := strconv.Atoi(data.ws.Query("storyID"))
	var product *protocol.MSG_PROJECT_product_cache
	var products []protocol.HtmlKeyValueStr
	if projectID > 0 {
		//$products = $this->product->getProductsByProject($projectID);
		//$product  = $this->product->getById(($productID and array_key_exists($productID, $products)) ? $productID : key($products));
	} else {
		products, err = product_getPairs(data, "noclosed")
		if err != nil {
			return
		}
		if productID == 0 && len(products) > 0 {
			id, _ := strconv.Atoi(products[0].Key)
			product = HostConn.GetProductById(int32(id))
		} else {
			product = HostConn.GetProductById(int32(productID))
		}

	}
	if len(products) == 0 {
		data.ws.WriteString(js.Location(createLink("product", "create", nil), ""))
		return
	}

	users, err := user_getPairs(data, "pdfirst|noclosed|nodeleted")
	if err != nil {
		return
	}
	moduleOptionMenu, err := tree_getOptionMenu(data, int32(productID), "story", 0, int32(branch))
	if err != nil {
		return
	}
	if len(moduleOptionMenu) == 0 {
		data.ws.WriteString(js.Location(createLink("tree", "browse", []interface{}{"productID=", productID, "&view=story"}), ""))
		return
	}

	if err = product_setMenu(data, int32(productID), int32(branch), ""); err != nil {
		return
	}
	var (
		source     = ""
		sourceNote = ""
		pri        = 0
		estimate   = ""
		title      = ""
		spec       = ""
		verify     = ""
		keywords   = ""
		mailto     = ""
		color      = ""
	)

	/*if(storyID > 0){
	      story      = this->story->getByID(storyID)
	      planID     = story->plan
	      source     = story->source
	      sourceNote = story->sourceNote
	      color      = story->color
	      pri        = story->pri
	      productID  = story->product
	      moduleID   = story->module
	      estimate   = story->estimate
	      title      = story->title
	      spec       = htmlspecialchars(story->spec)
	      verify     = htmlspecialchars(story->verify)
	      keywords   = story->keywords
	      mailto     = story->mailto
	  }

	  if(bugID > 0){
	      oldBug    = this->loadModel("bug")->getById(bugID)
	      productID = oldBug->product
	      source    = "bug"
	      title     = oldBug->title
	      keywords  = oldBug->keywords
	      spec      = oldBug->steps
	      pri       = oldBug->pri
	      if(strpos(oldBug->mailto, oldBug->openedBy) === false)
	      {
	          mailto = oldBug->mailto . oldBug->openedBy . ","
	      }
	      else
	      {
	          mailto = oldBug->mailto
	      }
	  }

	  if(todoID > 0){
	      todo   = this->loadModel("todo")->getById(todoID)
	      source = "todo"
	      title  = todo->name
	      spec   = todo->desc
	      pri    = todo->pri
	  }


	  if(isset(fromObject)){
	      if(isset(this->config->story->fromObjects[fromObjectIDKey]["source"]))
	      {
	          sourceField = this->config->story->fromObjects[fromObjectIDKey]["source"]
	          sourceUser  = this->loadModel("user")->getById(fromObject->{sourceField})
	          source      = sourceUser->role
	          sourceNote  = sourceUser->realname
	      }
	      else
	      {
	          source      = fromObjectName
	          sourceNote  = fromObjectID
	      }

	      foreach(this->config->story->fromObjects[fromObjectIDKey]["fields"] as storyField => fromObjectField)
	      {
	          storyField = fromObject->{fromObjectField}
	      }
	  }*/
	var customFields = make([]protocol.HtmlKeyValueStr, len(data.Config["story"]["list"]["customCreateFields"].([]string)))
	for k, field := range data.Config["story"]["list"]["customCreateFields"].([]string) {
		customFields[k] = protocol.HtmlKeyValueStr{Key: field, Value: data.Lang["story"][field].(string)}
	}
	data.Data["customFields"] = customFields

	data.Data["showFields"] = strings.Join(data.Config["story"]["custom"]["createFields"].([]string), ",")

	data.Data["title"] = product.Name + data.Lang["common"]["colon"].(string) + data.Lang["story"]["create"].(string)
	/*data.Data["position"][]       = html::a(this->createLink("product", "browse", "product=productID&branch=branch"), product->name)
	  data.Data["position"][]       = this->lang->story->common
	  data.Data["position"][]       = this->lang->story->create*/
	data.Data["products"] = products
	data.Data["users"] = users
	data.Data["moduleID"] = moduleID
	data.Data["moduleOptionMenu"] = moduleOptionMenu
	productplan_getPairsForStory := protocol.GET_MSG_PROJECT_productplan_getPairsForStory()
	productplan_getPairsForStory.Product = int32(productID)
	productplan_getPairsForStory.Branch = int32(branch)
	var plans *protocol.MSG_PROJECT_productplan_getPairsForStory_result
	if err = data.SendMsgWaitResultToDefault(productplan_getPairsForStory, &plans); err != nil {
		return
	}
	productplan_processFuture(data, plans.List)
	data.Data["plans"] = plans.List
	data.Data["planID"] = planID
	data.Data["source"] = source
	data.Data["sourceNote"] = sourceNote
	data.Data["color"] = color
	data.Data["pri"] = pri
	data.Data["branch"] = branch
	if product.Type != "normal" {
		data.Data["branches"] = branch_getPairs(data, product.Id, product)
	}

	data.Data["productID"] = productID
	data.Data["product"] = product
	data.Data["projectID"] = projectID
	data.Data["estimate"] = estimate
	data.Data["storyTitle"] = title
	data.Data["spec"] = spec
	data.Data["verify"] = verify
	data.Data["keywords"] = keywords
	data.Data["mailto"] = mailto
	data.Data["customLink"] = createLink("custom", "ajaxSaveCustomFields", "module=story&section=custom&key=createFields")
	if data.User.Id == product.PO || projectID > 0 || !data.Config["story"]["common"]["needReview"].(bool) {
		data.Data["needReview"] = "checked='checked'"
	}

	templateOut("story.create.html", data)
	productplan_getPairsForStory.Put()
	plans.Put()
	return
}
func post_story_create(data *TemplateData) (e error) {
	if !data.ajaxCheckPost() {
		return
	}
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	bugID, _ := strconv.Atoi(data.ws.Query("bugID"))
	var err error
	defer func() {
		if err != nil {
			data.ajaxResult(false, err.Error())
		}
	}()
	if projectID > 0 {
		project := data.getCacheProjectById(int32(projectID))
		if project == nil {
			err = errors.New(data.Lang["project"]["error"].(map[string]string)["NotFound"])
			return
		}
	}

	out := protocol.GET_MSG_PROJECT_stroy_create()
	data.SetValueFromPost(out)
	out.FromBug = int32(bugID)
	out.ProjectID = int32(projectID)
	for i := len(out.Mailto) - 1; i >= 0; i-- {
		if user := HostConn.GetUserCacheById(out.Mailto[i]); user == nil {
			out.Mailto = append(out.Mailto[:i], out.Mailto[i+1:]...)
		}
	}
	out.OpenedBy = data.User.Id

	var result *protocol.MSG_PROJECT_stroy_create_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}

	storyID := result.Result
	result.Put()
	defer out.Put()
	if storyID == int32(protocol.Err_ProjectStoryTitleExists) {

		data.ajaxResult(false, map[string]string{"title": fmt.Sprintf(data.Lang["common"]["duplicate"].(string), data.Lang["story"]["common"])})
		return
	}

	/*if(todoID > 0){//代办
	    this->dao->update(TABLE_TODO)->set("status")->eq("done")->where("id")->eq(todoID)->exec();
	    this->action->create("todo", todoID, "finished", "", "STORY:storyID");
	}*/

	if data.ws.Post("newStory") != "" {

		data.ajaxResult(true, data.Lang["story"]["successSaved"].(string)+data.Lang["story"]["newStory"].(string), createLink("story", "create", []interface{}{"productID=", out.Product, "&branch=", out.Branch, "&moduleID=", out.Module, "&story=0&projectID=", projectID, "&bugID=", bugID}))
		return
	}
	var locate string
	if projectID == 0 {
		locate = createLink("story", "view", "storyID="+strconv.Itoa(int(storyID)))
	} else {
		locate = createLink("project", "story", "projectID="+strconv.Itoa(int(projectID))+"&branch=&browseType=byModule&moduleID="+strconv.Itoa(int(out.Module)))
	}
	data.ajaxResult(true, data.Lang["common"]["saveSuccess"].(string), locate)
	return
}
func storyTemplateFuncs() {
	global_Funcs["story_printCell"] = func(data *TemplateData, col *config.ConfigDatatable, story *protocol.MSG_PROJECT_story, users []protocol.HtmlKeyValueStr, branches []protocol.HtmlKeyValueStr, storyStages map[int32][]protocol.HtmlKeyValueStr, modulePairs []protocol.HtmlKeyValueStr, storyTasks, storyBugs, storyCases map[int32]int, mode string) template.HTML { //mode = 'datatable'
		if mode == "" {
			mode = "datatable"
		}
		canView := hasPriv(data, "story", "view")
		storyLink := createLink("story", "view", "storyID="+strconv.Itoa(int(story.Id)))
		buf := bufpool.Get().(*libraries.MsgBuffer)
		if col.Show {
			buf.WriteString("<td class='c-")
			buf.WriteString(col.Id)
			switch col.Id {
			case "id":
				buf.WriteString("'>")
				buf.WriteString(html_checkbox("storyIDList", []protocol.HtmlKeyValueStr{{strconv.Itoa(int(story.Id)), ""}}, "", "", "block"))
				buf.WriteString(html_a(createLink("story", "view", "storyID="+strconv.Itoa(int(story.Id))), fmt.Sprintf("%03d", story.Id)))

			case "pri":
				var pri = strconv.Itoa(int(story.Pri))
				buf.WriteString("'>")
				buf.WriteString("<span class='label-pri label-pri-")
				buf.WriteString(pri)
				buf.WriteString("' title='")
				for _, kv := range data.Lang["story"]["priList"].([]protocol.HtmlKeyValueStr) {
					if kv.Key == pri {
						pri = kv.Value
						break
					}
				}
				buf.WriteString(pri)
				buf.WriteString("'>")
				buf.WriteString(pri)
				buf.WriteString("</span>")
			case "title":
				buf.WriteString("' title='")
				buf.WriteString(story.Title)
				buf.WriteString("'>")
				if story.Branch > 0 {
					for _, branch := range branches {
						if branch.Key == strconv.Itoa(int(story.Branch)) {
							buf.WriteString("<span class='label label-outline label-badge'>")
							buf.WriteString(branch.Value)
							buf.WriteString("</span> ")
							break
						}
					}
				}
				if story.Module > 0 {
					for _, module := range modulePairs {
						if module.Key == strconv.Itoa(int(story.Module)) {
							buf.WriteString("<span class='label label-gray label-badge'>")
							buf.WriteString(module.Value)
							buf.WriteString("</span> ")
							break
						}
					}
				}
				if canView {
					buf.WriteString(html_a(storyLink, story.Title, "", "style='color: "+story.Color+"'"))
				} else {
					buf.WriteString("<span style='color: ")
					buf.WriteString(story.Color)
					buf.WriteString("'>")
					buf.WriteString(story.Title)
					buf.WriteString("</span>")
				}

			case "plan":
				buf.WriteString(" text-ellipsis")
				buf.WriteString("' title='")
				buf.WriteString(story.PlanTitle)
				buf.WriteString("'>")
				buf.WriteString(story.PlanTitle)
			case "branch":
				buf.WriteString("'>")
				for _, branch := range branches {
					if branch.Key == strconv.Itoa(int(story.Branch)) {
						buf.WriteString(branch.Value)
						break
					}
				}
			case "keywords":
				buf.WriteString("'>")
				buf.WriteString(story.Keywords)
			case "source":
				buf.WriteString("'>")
				source := story.Source
				for _, kv := range data.Lang["story"]["sourceList"].([]protocol.HtmlKeyValueStr) {
					if kv.Key == source {
						source = kv.Value
						break
					}
				}
				buf.WriteString(source)

			case "sourceNote":
				buf.WriteString(" text-ellipsis")
				buf.WriteString("' title='")
				buf.WriteString(story.SourceNote)
				buf.WriteString("'>")
				buf.WriteString(story.SourceNote)
			case "status":
				buf.WriteString("'>")
				buf.WriteString("<span class='status-")
				buf.WriteString(story.Status)
				buf.WriteString("'>")

				for _, kv := range data.Lang["story"]["statusList"].([]protocol.HtmlKeyValueStr) {
					if kv.Key == story.Status {
						buf.WriteString(kv.Value)
						break
					}
				}
				buf.WriteString("</span>")

			case "estimate":
				buf.WriteString("'>")
				buf.WriteString(fmt.Sprintf("%.1f", story.Estimate))

			case "stage":
				buf.WriteString("'>")
				if v, ok := storyStages[story.Id]; ok {
					buf.WriteString("<div class='dropdown dropdown-hover'>")
					for _, kv := range data.Lang["story"]["stageList"].([]protocol.HtmlKeyValueStr) {
						if kv.Key == story.Stage {
							buf.WriteString(kv.Value)
							break
						}
					}
					buf.WriteString("<span class='caret'></span><ul class='dropdown-menu pull-right'>")

					for _, storyStage := range v {
						buf.WriteString("<li class='text-ellipsis'>")
						for _, branch := range branches {
							if branch.Key == storyStage.Key {
								buf.WriteString(branch.Value)
								break
							}
						}
						buf.WriteString(": ")
						for _, kv := range data.Lang["story"]["stageList"].([]protocol.HtmlKeyValueStr) {
							if kv.Key == storyStage.Value {
								buf.WriteString(kv.Value)
								break
							}
						}
						buf.WriteString("</li>")
					}

					buf.WriteString("</ul></div>")

				} else {

					for _, kv := range data.Lang["story"]["stageList"].([]protocol.HtmlKeyValueStr) {
						if kv.Key == story.Stage {
							buf.WriteString(kv.Value)
							break
						}
					}
				}
			case "taskCount":
				buf.WriteString("'>")
				if n, ok := storyTasks[story.Id]; ok {
					buf.WriteString(html_a(createLink("story", "tasks", "storyID="+strconv.Itoa(int(story.Id))), strconv.Itoa(n), "", "class='iframe'"))
				} else {
					buf.WriteString("0")
				}
			case "bugCount":
				buf.WriteString("'>")
				if n, ok := storyBugs[story.Id]; ok {
					buf.WriteString(html_a(createLink("story", "bugs", "storyID="+strconv.Itoa(int(story.Id))), strconv.Itoa(n), "", "class='iframe'"))
				} else {
					buf.WriteString("0")
				}
			case "caseCount":
				buf.WriteString("'>")
				if n, ok := storyCases[story.Id]; ok {
					buf.WriteString(html_a(createLink("story", "cases", "storyID="+strconv.Itoa(int(story.Id))), strconv.Itoa(n), "", "class='iframe'"))
				} else {
					buf.WriteString("0")
				}
			case "openedBy":
				var username string
				for _, user := range users {
					if user.Key == strconv.Itoa(int(story.OpenedBy)) {
						username = user.Value
						break
					}
				}
				buf.WriteString("' title='")
				buf.WriteString(username)
				buf.WriteString("'>")
				buf.WriteString(username)

			case "openedDate":
				buf.WriteString("'>")
				buf.WriteString(story.OpenedDate.Format("01-02 15:04"))
			case "assignedTo":
				if story.AssignedTo == data.User.Id {
					buf.WriteString(" red")
				}
				var username string
				for _, user := range users {
					if user.Key == strconv.Itoa(int(story.AssignedTo)) {
						username = user.Value
						break
					}
				}
				buf.WriteString("' title='")
				buf.WriteString(username)
				buf.WriteString("'>")
				buf.WriteString("<span style='padding-left:10px;' ")
				if story.AssignedTo == data.User.Id {
					buf.WriteString("class='text-red' ")
				}
				buf.WriteString(">")
				buf.WriteString(username)
				buf.WriteString("</span>")
			case "assignedDate":
				buf.WriteString("'>")
				buf.WriteString(story.AssignedDate.Format("01-02 15:04"))
			case "reviewedBy":
				buf.WriteString("'>")
				for _, user := range users {
					if user.Key == strconv.Itoa(int(story.ReviewedBy)) {
						buf.WriteString(user.Value)
						break
					}
				}
			case "reviewedDate":
				buf.WriteString("'>")
				buf.WriteString(story.ReviewedDate.Format("01-02 15:04"))
			case "closedBy":
				buf.WriteString("'>")
				for _, user := range users {
					if user.Key == strconv.Itoa(int(story.ClosedBy)) {
						buf.WriteString(user.Value)
						break
					}
				}
			case "closedDate":
				buf.WriteString("'>")
				buf.WriteString(story.ClosedDate.Format("01-02 15:04"))
			case "closedReason":
				buf.WriteString("'>")
				reason := story.ClosedReason
				for _, kv := range data.Lang["story"]["reasonList"].([]protocol.HtmlKeyValueStr) {
					if kv.Key == reason {
						reason = kv.Value
						break
					}
				}
				buf.WriteString(reason)

			case "lastEditedBy":
				buf.WriteString("'>")
				for _, user := range users {
					if user.Key == strconv.Itoa(int(story.LastEditedBy)) {
						buf.WriteString(user.Value)
						break
					}
				}
			case "lastEditedDate":
				buf.WriteString("'>")
				buf.WriteString(story.LastEditedDate.Format("01-02 15:04"))
			case "mailto":
				buf.WriteString("'>")

				for _, id := range story.Mailto {
					for _, user := range users {
						if user.Key == strconv.Itoa(int(id)) {
							buf.WriteString(user.Value)
							buf.WriteString("&nbsp;")
							break
						}
					}

				}
			case "version":
				buf.WriteString("'>")
				buf.WriteString(strconv.Itoa(int(story.Version)))

			case "actions":
				buf.WriteString("'>")
				vars := "story=" + strconv.Itoa(int(story.Id))
				buf.WriteString(common_printIcon(data, "story", "change", vars, story, "list", "fork"))
				buf.WriteString(common_printIcon(data, "story", "review", vars, story, "list", "glasses"))
				buf.WriteString(common_printIcon(data, "story", "close", vars, story, "list", "off", "", "iframe", "true"))
				buf.WriteString(common_printIcon(data, "story", "edit", vars, story, "list", ""))
				buf.WriteString(common_printIcon(data, "story", "createCase", "productID="+strconv.Itoa(int(story.Product))+"&branch="+strconv.Itoa(int(story.Branch))+"&module=0&from=&param=0&"+vars, story, "list", "sitemap"))

			}
			buf.WriteString("</td>")
		}
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
	global_Funcs["MSG_PROJECT_story_isClickable"] = func(data *TemplateData, obj interface{}, action string) bool {
		if story, ok := obj.(*protocol.MSG_PROJECT_story); ok {
			if action == "change" {
				return story.Status != "closed"
			}
			if action == "review" {
				return story.Status == "draft" || story.Status == "changed"
			}
			if action == "close" {
				return story.Status != "closed"
			}
			if action == "activate" {
				return story.Status == "closed"
			}

		} else {
			libraries.DebugLog("MSG_PROJECT_story_isClickable传入的值类型%v不对", reflect.TypeOf(obj).Elem().Name())
		}

		return true

	}
}
func get_story_view(data *TemplateData) (err error) {
	storyID, _ := strconv.Atoi(data.ws.Query("storyID"))
	version, _ := strconv.Atoi(data.ws.Query("version"))
	param, _ := strconv.Atoi(data.ws.Query("param"))
	from := data.ws.Query("storyID")
	if from == "" {
		from = "product"
	}

	if err != nil {
		return
	}
	story_getById := protocol.GET_MSG_PROJECT_story_getById()
	story_getById.Id = int32(storyID)
	story_getById.Version = int16(version)
	var result *protocol.MSG_PROJECT_story_getById_result
	if err = data.SendMsgWaitResultToDefault(story_getById, &result); err != nil {
		return
	} else if result.Story == nil {
		data.ws.WriteString(js.Alert(data.Lang["common"]["notFound"].(string)) + js.Location("back", "self"))
		return
	}
	if files, err := file_getByObject(data, "story", int32(storyID)); err != nil {
		return err
	} else {
		data.Data["Files"] = files
	}
	if err = product_setMenu(data, result.Story.Product, result.Story.Branch, ""); err != nil {
		return
	}
	product := HostConn.GetProductById(result.Story.Product)
	if data.Data["projects"], err = project_getPairs(data, "nocode"); err != nil {
		return
	}
	if result.Story.Plan > 0 {
		getplan := protocol.GET_MSG_PROJECT_productplan_getList()
		getplan.Ids = []int32{result.Story.Plan}
		getplan.Total = 1
		getplan.Page = 1
		getplan.PerPage = 1
		var getplanResult *protocol.MSG_PROJECT_productplan_getList_result
		if err = data.SendMsgWaitResultToDefault(getplan, &getplanResult); err != nil {
			return
		}
		if len(getplanResult.List) > 0 {
			data.Data["plan"] = getplanResult.List[0]
		}

	}

	//bugs         = this->dao->select("id,title")->from(TABLE_BUG)->where("story")->eq(storyID)->andWhere("deleted")->eq(0)->fetchAll();
	//fromBug      = this->dao->select("id,title")->from(TABLE_BUG)->where("toStory")->eq(storyID)->fetch();
	//cases        = this->dao->select("id,title")->from(TABLE_CASE)->where("story")->eq(storyID)->andWhere("deleted")->eq(0)->fetchAll();
	//this->view->bugs       = bugs;
	//this->view->fromBug    = fromBug;
	//this->view->cases      = cases;
	if result.Story.Module > 0 {
		getParents := protocol.GET_MSG_PROJECT_tree_getParents()
		getParents.ModuleID = result.Story.Module
		var getParentsResult *protocol.MSG_PROJECT_tree_getParents_result
		if err = data.SendMsgWaitResultToDefault(getParents, &getParentsResult); err != nil {
			return

		}
		data.Data["modulePath"] = getParentsResult.List
		if data.Data["treeOption"], err = tree_getOptionMenu(data, result.Story.Product, "story", result.Story.Module, result.Story.Branch); err != nil {
			return
		}
	}

	if data.Data["users"], err = user_getPairs(data, "noletter"); err != nil {
		return
	}
	/* Set the menu. */

	if from == "project" && param > 0 {
		if project := data.getCacheProjectById(int32(param)); project != nil && project.Status == "done" {
			from = ""
		}
	}

	data.Data["title"] = fmt.Sprintf("STORY #%d %s - %s", result.Story.Id, result.Story.Title, product.Name)
	data.Data["product"] = product
	if product.Type != "normal" {
		data.Data["branches"] = branch_getPairs(data, product.Id, product)
	}
	data.Data["story"] = result.Story

	if data.Data["actions"], err = action_getList(data, "story", int32(storyID)); err != nil {
		return
	}
	if version == 0 {
		data.Data["version"] = result.Story.Version
	} else {
		data.Data["version"] = version
	}
	//this->view->preAndNext = this->loadModel("common")->getPreAndNextObject("story", storyID);
	data.Data["from"] = from
	data.Data["param"] = param
	data.Data["actionFormLink"] = createLink("action", "comment", []interface{}{"objectType=story&objectID=", result.Story.Id})
	templateOut("story.view.html", data)
	return
}
func story_getProjectStoryPairs(data *TemplateData, projectID, productID, branch int32, moduleIdList []int32, Type string) (list []protocol.HtmlKeyValueStr, err error) {
	out := protocol.GET_MSG_PROJECT_story_getProjectStoryPairs()
	out.ProductID = productID
	out.ProjectID = projectID
	out.Branch = branch
	out.ModuleIdList = moduleIdList

	var result *protocol.MSG_PROJECT_story_getProjectStoryPairs_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}

	out.Put()
	result.Put()
	return
}
func story_formatStories(data *TemplateData, stories []*protocol.MSG_PROJECT_TASK, Type string, limit int) (list []protocol.HtmlKeyValueStr) {
	i := 0
	buf := bufpool.Get().(*libraries.MsgBuffer)
	for _, story := range stories {
		i++
		buf.WriteString(strconv.Itoa(int(story.Id)))
		buf.WriteByte(':')
		buf.WriteString(story.StoryTitle)
		buf.WriteByte(' ')
		if Type == "short" {
			buf.WriteString("[p")
			buf.WriteString(common_getValue(data.Lang["story"]["priList"], story.Pri).(string))
			buf.WriteString(", ")
			buf.WriteString(fmt.Sprintf("%.2f", story.Estimate))
			buf.WriteString("h]")

		} else {
			buf.WriteString("(")
			buf.WriteString(data.Lang["story"]["pri"].(string))
			buf.WriteByte(':')
			buf.WriteString(common_getValue(data.Lang["story"]["priList"], story.Pri).(string))
			buf.WriteByte(',')
			buf.WriteString(data.Lang["story"]["estimate"].(string))
			buf.WriteString(fmt.Sprintf("%.2f", story.Estimate))
			buf.WriteByte(')')

		}
		list = append(list, protocol.HtmlKeyValueStr{strconv.Itoa(int(story.Id)), buf.String()})
		buf.Reset()
		if limit > 0 && i > limit {
			list = append(list, protocol.HtmlKeyValueStr{"showmore", data.Lang["common"]["more"].(string) + data.Lang["common"]["ellipsis"].(string)})

			break
		}
	}
	return
}
func get_story_ajaxGetProjectStories(data *TemplateData) (err error) {

	id, _ := strconv.Atoi(data.ws.Query("moduleID"))
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	branch, _ := strconv.Atoi(data.ws.Query("branch"))
	Type := data.ws.Query("type")
	if Type == "" {
		Type = "full"
	}
	var moduleIdList []int32
	if id > 0 {
		moduleID, err := tree_getStoryModule(data, int32(id))
		if err != nil {
			return err
		}
		moduleIdList = tree_getAllChildId(data, moduleID)
	}
	stories, err := story_getProjectStoryPairs(data, int32(projectID), int32(productID), int32(branch), moduleIdList, Type)
	if err != nil {
		return
	}
	storyName := "story"
	if data.ws.Query("number") != "" {
		storyName += "[" + data.ws.Query("number") + "]"
	}
	data.ws.WriteString(html_select(storyName, stories, data.ws.Query("storyID"), "class=form-control onchange=setStoryRelated("+data.ws.Query("number")+");"))
	return
}
func story_getUserStories(data *TemplateData, uid int32, typ, orderby string, page *TempLatePage) ([]map[string]string, error) {
	out := protocol.GET_MSG_PROJECT_product_getStoriesMapBySql()
	out.PerPage = page.PerPage
	out.Total = page.Total
	out.Page = page.Page
	out.Order = orderby
	out.Field = "*"
	out.Where = map[string]interface{}{"Deleted": false}
	if typ != "all" {
		out.Where[typ] = uid
	}
	var result *protocol.MSG_PROJECT_product_getStoriesMapBySql_result
	if err := data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return nil, err
	}
	page.Total = result.Total
	out.Put()
	return result.List, nil
}
