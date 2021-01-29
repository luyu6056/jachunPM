package handler

import (
	"jachunPM_http/js"
	"protocol"
	"strconv"
	"strings"
)

func init() {

	httpHandlerMap["GET"]["/story/create"] = get_story_create
	httpHandlerMap["POST"]["/story/create"] = post_story_create
}
func get_story_create(data *TemplateData) {

	/*extra = str_replace(array(",", " "), array("&", ""), extra)
	  parse_str(extra, output)
	  foreach(output as paramKey => paramValue)
	  {
	      if(isset(this->config->story->fromObjects[paramKey]))
	      {
	          fromObjectIDKey  = paramKey
	          fromObjectID     = paramValue
	          fromObjectName   = this->config->story->fromObjects[fromObjectIDKey]["name"]
	          fromObjectAction = this->config->story->fromObjects[fromObjectIDKey]["action"]
	          break
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
	branch, _ := strconv.Atoi(data.ws.Query("branch"))
	//storyID, _ := strconv.Atoi(data.ws.Query("storyID"))

	var product *protocol.MSG_PROJECT_product_cache
	var products []protocol.HtmlKeyValueStr
	var err error
	if projectID > 0 {
		//$products = $this->product->getProductsByProject($projectID);
		//$product  = $this->product->getById(($productID and array_key_exists($productID, $products)) ? $productID : key($products));
	} else {
		products, err = product_getPairs(data, "noclosed")
		if err != nil {
			data.OutErr(err)
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
	msg, err := data.GetMsg()
	if err != nil {
		data.OutErr(err)
		return
	}
	users, err := user_getPairs("pdfirst|noclosed|nodeleted")
	if err != nil {
		data.OutErr(err)
		return
	}
	moduleOptionMenu, err := tree_getOptionMenu(data, int32(productID), "story", 0, int32(branch))
	if err != nil {
		data.OutErr(err)
		return
	}
	if len(moduleOptionMenu) == 0 {
		data.ws.WriteString(js.Location(createLink("tree", "browse", []interface{}{"productID=", productID, "&view=story"}), ""))
		return
	}

	if err = product_setMenu(data, int32(productID), int32(branch), ""); err != nil {
		data.OutErr(err)
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
	if err = msg.SendMsgWaitResult(0, productplan_getPairsForStory, &plans); err != nil {
		data.OutErr(err)
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
}
func post_story_create(data *TemplateData) {
	if !data.ajaxCheckPost() {
		return
	}
	msg, err := data.GetMsg()
	defer func() {
		if err != nil {
			data.ajaxResult(false, err.Error())
		}
	}()
	if err != nil {
		return
	}
	out := protocol.GET_MSG_PROJECT_stroy_create()
	for key, list := range data.ws.GetAllPost() {
		switch key {
		case "product":
			i, _ := strconv.Atoi(list[0])
			out.Product = int32(i)
		case "branch":
			i, _ := strconv.Atoi(list[0])
			out.Branch = int32(i)
		}
	}
	//response["result"]  = "success";
	//response["message"] = this->lang->saveSuccess;

	/*storyResult = this->story->create(projectID, bugID, from = isset(fromObjectIDKey) ? fromObjectIDKey : "");
	  if(!storyResult or dao::isError())
	  {
	      response["result"]  = "fail";
	      response["message"] = dao::getError();
	      this->send(response);
	  }
	  storyID = storyResult["id"];
	  if(storyResult["status"] == "exists")
	  {
	      response["message"] = sprintf(this->lang->duplicate, this->lang->story->common);
	      if(projectID == 0)
	      {
	          response["locate"] = this->createLink("story", "view", "storyID={storyID}");
	      }
	      else
	      {
	          response["locate"] = this->createLink("project", "story", "projectID=projectID");
	      }
	      this->send(response);
	  }

	  action = bugID == 0 ? "Opened" : "Frombug";
	  extra  = bugID == 0 ? "" : bugID;

	  if(isset(fromObjectID))
	  {
	      action = fromObjectAction;
	      extra  = fromObjectID;
	  }
	  actionID = this->action->create("story", storyID, action, "", extra);

	  if(todoID > 0)
	  {
	      this->dao->update(TABLE_TODO)->set("status")->eq("done")->where("id")->eq(todoID)->exec();
	      this->action->create("todo", todoID, "finished", "", "STORY:storyID");
	  }

	  if(this->post->newStory)
	  {
	      response["message"] = this->lang->story->successSaved . this->lang->story->newStory;
	      response["locate"]  = this->createLink("story", "create", "productID=productID&branch=branch&moduleID=moduleID&story=0&projectID=projectID&bugID=bugID");
	      this->send(response);
	  }

	  moduleID = this->post->module ? this->post->module : 0;
	  response["locate"] = this->createLink("project", "story", "projectID=projectID&branch=&browseType=byModule&moduleID=moduleID");
	  if(projectID == 0) response["locate"] = this->createLink("story", "view", "storyID=storyID");
	  this->send(response);*/
}
