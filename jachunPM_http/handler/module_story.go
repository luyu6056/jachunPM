package handler

import (
	"jachunPM_http/js"
	"protocol"
	"strconv"

	"github.com/luyu6056/gnet"
)

func init() {

	httpHandlerMap["GET"]["/story/create"] = get_story_create
}
func get_story_create(data *TemplateData) (action gnet.Action) {

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
			productID, _ = strconv.Atoi(products[0].Key)
			product = HostConn.GetProductById(int32(productID))
		}

	}
	if len(products) == 0 {
		data.ws.WriteString(js.Location(createLink("product", "create", nil), ""))
		return
	}
	msg, err := HostConn.GetMsg()
	if err != nil {
		data.OutErr(err)
		return
	}
	user_getPairs := protocol.GET_MSG_USER_getPairs()
	user_getPairs.Params = "pdfirst|noclosed|nodeleted"
	var result *protocol.MSG_USER_getPairs_result
	err = msg.SendMsgWaitResult(0, user_getPairs, &result)
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

	err = product_setMenu(data, int32(productID), int32(branch), "")
	if err != nil {
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
	var customFields = make(map[string]string, len(data.Config["story"]["list"]["customCreateFields"].([]string)))
	for _, field := range data.Config["story"]["list"]["customCreateFields"].([]string) {
		customFields[field] = data.Lang["story"][field].(string)
	}
	data.Data["customFields"] = customFields

	data.Data["showFields"] = data.Config["story"]["custom"]["createFields"]

	data.Data["title"] = product.Name + data.Lang["common"]["colon"].(string) + data.Lang["story"]["create"].(string)
	/*data.Data["position"][]       = html::a(this->createLink("product", "browse", "product=productID&branch=branch"), product->name)
	  data.Data["position"][]       = this->lang->story->common
	  data.Data["position"][]       = this->lang->story->create*/
	data.Data["products"] = products
	data.Data["users"] = result.List
	user_getPairs.Put()
	result.Put()
	data.Data["moduleID"] = moduleID
	data.Data["moduleOptionMenu"] = moduleOptionMenu
	//data.Data["plans"]            = this->loadModel("productplan")->getPairsForStory(productID, branch)
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
	if data.User.Id == product.PO || projectID > 0 || !data.Config["story"]["common"]["needReview"].(bool) {
		data.Data["needReview"] = "checked='checked'"
	}

	templateOut("story.create.html", data)
	return
}
