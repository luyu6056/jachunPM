package handler

import (
	"errors"
	"fmt"
	"jachunPM_project/db"
	"protocol"
	"reflect"
	"strconv"
	"time"
)

func product_insert(data *protocol.MSG_PROJECT_product_insert, in *protocol.Msg) {

	session, err := in.BeginTransaction()
	defer session.EndTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	insert := &db.Product{
		Name:        data.Data.Name,
		Code:        data.Data.Code,
		Line:        data.Data.Line,
		Type:        data.Data.Type,
		Status:      data.Data.Status,
		Desc:        data.Data.Desc,
		PO:          data.Data.PO,
		QD:          data.Data.QD,
		RD:          data.Data.RD,
		Acl:         data.Data.Acl,
		Whitelist:   data.Data.Whitelist,
		CreatedBy:   data.Data.CreatedBy,
		CreatedDate: data.Data.CreatedDate,
		Order:       data.Data.Order,
		Deleted:     data.Data.Deleted,
		TimeStamp:   time.Now().Unix(),
	}
	id, err := session.Table(db.TABLE_PRODUCT).Insert(insert)
	if err != nil {
		in.WriteErr(err)
		return
	}
	if id > 0 {

		data.Data.Id = int32(id)
		session.Table(db.TABLE_PRODUCT).Where("Id=?", data.Data.Id).Update("Order=?", id*5)
		data.Data.Order = data.Data.Id * 5
		doc := &db.Doclib{
			Product: data.Data.Id,
			Name:    data.DocName,
			Type:    "product",
			Main:    true,
			Acl:     "private",
		}
		if data.Data.Acl == "open" {
			doc.Acl = "open"
		}
		_, err := session.Table(db.TABLE_DOCLIB).Insert(doc)
		if err != nil {
			in.WriteErr(err)
			return
		}
		out := protocol.GET_MSG_USER_updateUserView()
		out.ProductId = data.Data.Id
		if data.Data.PO > 0 {
			out.UserIds = append(out.UserIds, data.Data.PO)
		}
		if data.Data.RD > 0 {
			out.UserIds = append(out.UserIds, data.Data.RD)
		}
		if data.Data.QD > 0 {
			out.UserIds = append(out.UserIds, data.Data.QD)
		}
		out.GroupIds = data.Data.Whitelist
		err = in.SendMsgWaitResult(0, out, nil)
		if err != nil {
			in.WriteErr(err)
			return
		}
		session.CommitCallback(func() {
			product_setCache(int32(id))
		})
		session.Commit()

		out1 := protocol.GET_MSG_PROJECT_product_insert_result()
		out1.ID = data.Data.Id
		in.SendResult(out1)
		out.Put()
		out1.Put()

	} else {
		in.WriteErr(errors.New("error insert id"))
	}
}
func product_setCache(id int32) {
	product := protocol.GET_MSG_PROJECT_product_cache()
	HostConn.DB.Table(db.TABLE_PRODUCT).Prepare().Where("Id=?", id).Find(&product)
	if product.Id != 0 {
		HostConn.DB.Table(db.TABLE_BRANCH).Prepare().Where("Product=?", id).Order("`Order` asc").Limit(0).Select(&product.Branchs)
		var ids = make([]int32, len(product.Branchs))
		for k, b := range product.Branchs {
			ids[k] = b.Id
		}
		HostConn.DB.Table(db.TABLE_PRODUCT).Prepare().Where("Id=?", id).Update(map[string]interface{}{"Branch": ids})
		HostConn.CacheSet(protocol.PATH_PROJECT_PRODUCT_CACHE, strconv.Itoa(int(product.Id)), product, 0)

	}
	product.Put()
}
func product_getStories(data *protocol.MSG_PROJECT_product_getStories, in *protocol.Msg) {
	modules := tree_getAllChildId(data.ModuleID)
	var err error
	defer func() {
		if err != nil {
			in.WriteErr(err)
		}
	}()
	var list []*protocol.MSG_PROJECT_story
	switch data.BrowseType {
	case "unclosed":
		var statusList []protocol.HtmlKeyValueStr
		err = in.LoadConfigToValue("story", "common", "statusList", &statusList)
		if err != nil {
			return
		}

		var unclosedStatus []string
		for _, kv := range statusList {
			if kv.Key == "closed" {
				continue
			}
			unclosedStatus = append(unclosedStatus, kv.Key)
		}

		list, err = story_getProductStories(data.ProductID, data.Branch, modules, unclosedStatus, data.Sort, data.Page, data.PerPage, &data.Total)
	case "unplan":
		list, err = story_getByPlan(data.ProductID, data.Branch, modules, "", data.Sort, data.Page, data.PerPage, &data.Total)
	case "allstory":
		list, err = story_getProductStories(data.ProductID, data.Branch, modules, []string{"all"}, data.Sort, data.Page, data.PerPage, &data.Total)
	case "bymodule":
		list, err = story_getProductStories(data.ProductID, data.Branch, modules, []string{"all"}, data.Sort, data.Page, data.PerPage, &data.Total)
	case "bysearch":
		list, err = story_getBySearch(data.ProductID, data.Branch, 0, data.Where, data.Sort, data.Page, data.PerPage, &data.Total)
	}
	if err != nil {
		return
	}
	//获取planTitle
	var planids = make([]int32, len(list))
	for k, story := range list {
		planids[k] = story.Plan
	}
	var plans []*db.Productplan
	if err = in.DB.Table(db.TABLE_PRODUCTPLAN).Where(map[string]interface{}{"Id": planids}).Select(&plans); err != nil {
		return
	}
	for _, story := range list {
		for _, plan := range plans {
			if story.Plan == plan.Id {
				story.PlanTitle = plan.Title
			}
		}
	}

	out := protocol.GET_MSG_PROJECT_product_getStories_result()
	out.Total = data.Total
	out.List = list
	in.SendResult(out)
	out.Put()
	/*

	  if($browseType == 'bysearch')     $stories = $this->story->getBySearch($productID, $queryID, $sort, $pager, '', $branch);
	  if($browseType == 'assignedtome') $stories = $this->story->getByAssignedTo($productID, $branch, $modules, $this->app->user->account, $sort, $pager);
	  if($browseType == 'openedbyme')   $stories = $this->story->getByOpenedBy($productID, $branch, $modules, $this->app->user->account, $sort, $pager);
	  if($browseType == 'reviewedbyme') $stories = $this->story->getByReviewedBy($productID, $branch, $modules, $this->app->user->account, $sort, $pager);
	  if($browseType == 'closedbyme')   $stories = $this->story->getByClosedBy($productID, $branch, $modules, $this->app->user->account, $sort, $pager);
	  if($browseType == 'draftstory')   $stories = $this->story->getByStatus($productID, $branch, $modules, 'draft', $sort, $pager);
	  if($browseType == 'activestory')  $stories = $this->story->getByStatus($productID, $branch, $modules, 'active', $sort, $pager);
	  if($browseType == 'changedstory') $stories = $this->story->getByStatus($productID, $branch, $modules, 'changed', $sort, $pager);
	  if($browseType == 'willclose')    $stories = $this->story->get2BeClosed($productID, $branch, $modules, $sort, $pager);
	  if($browseType == 'closedstory')  $stories = $this->story->getByStatus($productID, $branch, $modules, 'closed', $sort, $pager);*/
}
func product_update(data *protocol.MSG_PROJECT_product_update, in *protocol.Msg) {
	var oldProduct *db.Product
	err := in.DB.Table(db.TABLE_PRODUCT).Prepare().Where("Id=?", data.Data.Id).Find(&oldProduct)
	if oldProduct == nil {
		in.WriteErr(nil)
		return
	}
	if err != nil {
		in.WriteErr(err)
		return
	}
	update := make(map[string]interface{})
	r1 := reflect.ValueOf(oldProduct).Elem()
	r2 := reflect.ValueOf(data.Data).Elem()
	for i := 0; i < r1.Type().NumField(); i++ {
		f1 := r1.Field(i)
		f2 := r2.FieldByName(r1.Type().Field(i).Name)
		if f2.Kind() != reflect.Invalid {
			i1 := fmt.Sprint(f1.Interface())
			i2 := fmt.Sprint(f2.Interface())
			if i1 != i2 {
				update[r1.Type().Field(i).Name] = f2.Interface()
			}
		}
	}

	if len(update) == 0 {
		in.WriteErr(nil)
		return
	}
	session, err := in.BeginTransaction()
	defer session.EndTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}

	update["TimeStamp"] = time.Now().Unix()
	_, err = session.Table(db.TABLE_PRODUCT).Where("Id=?", data.Data.Id).Update(update)
	if err != nil {
		in.WriteErr(err)
		return
	}
	if data.Data.Acl != oldProduct.Acl {
		_, err = session.Table(db.TABLE_DOCLIB).Where(map[string]interface{}{"product": oldProduct.Id}).Update(map[string]interface{}{"Acl": data.Data.Acl})
		if err != nil {
			in.WriteErr(err)
			return
		}
	}
	/*
	   if(!dao::isError())
	   {


	       $this->file->updateObjectID($this->post->uid, $productID, 'product');
	       if($product->acl != $oldProduct->acl or $product->whitelist != $oldProduct->whitelist) $this->loadModel('user')->updateUserView($productID, 'product');
	       return common::createChanges($oldProduct, $product);
	   }*/
	session.CommitCallback(func() {
		product_setCache(oldProduct.Id)
	})
	session.Commit()
	in.WriteErr(nil)
}
func product_editBranch(data *protocol.MSG_PROJECT_product_editBranch, in *protocol.Msg) {
	product := HostConn.GetProductById(data.ProductID)
	if product != nil {
		session, err := in.BeginTransaction()
		if err != nil {
			in.WriteErr(err)
		} else {
			defer session.EndTransaction()
			var insert []*protocol.MSG_PROJECT_branch_info
			for _, branch := range data.Branchs {
				if branch.Id > 0 {

					_, err = session.Table(db.TABLE_BRANCH).Prepare().Where("Id=?", branch.Id).Update("Name=?,TimeStamp=?,`Order`=?", branch.Name, time.Now().Unix(), branch.Order)
					if err != nil {
						in.WriteErr(err)
						return
					}
				} else {
					insert = append(insert, branch)
				}
			}
			if len(insert) > 0 {
				_, err = session.Table(db.TABLE_BRANCH).InsertAll(insert)
				if err != nil {
					in.WriteErr(err)
					return
				}
			}
		}
		session.CommitCallback(func() {
			product_setCache(data.ProductID)
		})
		session.Commit()
	}
	in.WriteErr(nil)
}
func product_deleteBranch(data *protocol.MSG_PROJECT_product_deleteBranch, in *protocol.Msg) {
	count, err := in.DB.Table(db.TABLE_PROJECT).Where("Branch=" + strconv.Itoa(int(data.BranchID)) + " and Deleted = 0").Count()
	out := protocol.GET_MSG_PROJECT_product_deleteBranch_result()
	out.Result = protocol.Success
	defer func() {
		if err != nil {
			in.WriteErr(err)
			return
		}

		in.SendResult(out)
		out.Put()
	}()
	if err != nil || count != 0 {
		out.Result = protocol.Err_ProjectBranchCanNotDelete_PROJECT
		return
	}
	count, err = in.DB.Table(db.TABLE_MODULE).Where("Branch=" + strconv.Itoa(int(data.BranchID)) + " and Deleted = 0").Count()
	if err != nil || count != 0 {
		out.Result = protocol.Err_ProjectBranchCanNotDelete_MODULE
		return
	}
	count, err = in.DB.Table(db.TABLE_STORY).Where("Branch=" + strconv.Itoa(int(data.BranchID)) + " and Deleted = 0").Count()
	if err != nil || count != 0 {
		out.Result = protocol.Err_ProjectBranchCanNotDelete_STORY
		return
	}
	count, err = in.DB.Table(db.TABLE_PRODUCTPLAN).Where("Branch=" + strconv.Itoa(int(data.BranchID)) + " and Deleted = 0").Count()
	if err != nil || count != 0 {
		out.Result = protocol.Err_ProjectBranchCanNotDelete_PRODUCTPLAN
		return
	}
	check := protocol.GET_MSG_TEST_product_deleteBranch_check()
	check.BranchID = data.BranchID
	var result *protocol.MSG_TEST_product_deleteBranch_result
	err = in.SendMsgWaitResult(0, check, &result)
	if err != nil {
		return
	}
	out.Result = result.Result
	check.Put()
	result.Put()
	if out.Result == protocol.Success {
		_, err = in.DB.Table(db.TABLE_BRANCH).Where("Id = " + strconv.Itoa(int(data.BranchID))).Update(map[string]interface{}{"Deleted": true, "TimeStamp": time.Now().Unix()})
		product_setCache(data.ProductID)
	}
	return
}
