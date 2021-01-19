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
	session, err := in.DB.BeginTransaction()
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
			var newProduct *protocol.MSG_PROJECT_product_cache
			in.DB.Table(db.TABLE_PRODUCT).Prepare().Where("Id=?", id).Find(&newProduct)
			if newProduct != nil {
				product_setCache(newProduct)
			}
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
func product_setCache(product *protocol.MSG_PROJECT_product_cache) {
	HostConn.CacheSet(protocol.PATH_PROJECT_PRODUCT_CACHE, strconv.Itoa(int(product.Id)), product, 0)
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
		var config map[string]map[string]interface{}
		config, err = HostConn.LoadConfig("story")
		if err != nil {
			return
		}
		var unclosedStatus []string
		for k, _ := range config["statusList"] {
			if k == "closed" {
				continue
			}
			unclosedStatus = append(unclosedStatus, k)
		}

		list, err = story_getProductStories(data.ProductID, data.Branch, modules, unclosedStatus, data.Sort, data.Page, data.PerPage, &data.Total)
		if err != nil {
			return
		}
	case "unplan":
		list, err = story_getByPlan(data.ProductID, data.Branch, modules, "", data.Sort, data.Page, data.PerPage, &data.Total)
	}
	out := protocol.GET_MSG_PROJECT_product_getStories_result()
	out.Total = data.Total
	out.List = list
	in.SendResult(out)
	out.Put()
	/*if($browseType == 'unplan')       $stories = $this->story->getByPlan($productID, $queryID, $modules, '', $sort, $pager);
	  if($browseType == 'allstory')     $stories = $this->story->getProductStories($productID, $branch, $modules, 'all', $sort, $pager);
	  if($browseType == 'bymodule')     $stories = $this->story->getProductStories($productID, $branch, $modules, 'all', $sort, $pager);
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
	session, err := in.DB.BeginTransaction()
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
		var newProduct *protocol.MSG_PROJECT_product_cache
		in.DB.Table(db.TABLE_PRODUCT).Prepare().Where("Id=?", oldProduct.Id).Find(&newProduct)
		if newProduct != nil {
			product_setCache(newProduct)
		}
	})
	session.Commit()
	in.WriteErr(nil)
}
