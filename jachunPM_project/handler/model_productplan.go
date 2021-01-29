package handler

import (
	"jachunPM_project/db"
	"protocol"
	"strconv"
	"time"
)

func productplan_getPairsForStory(data *protocol.MSG_PROJECT_productplan_getPairsForStory, in *protocol.Msg) {
	date := time.Now().Format("2006-01-02")
	where := map[string]interface{}{
		"Deleted": false,
		"Product": data.Product,
		"end":     []interface{}{"ge", date},
	}
	if data.Branch > 0 {
		where["Branch"] = []int32{0, data.Branch}
	}
	var plans []*db.Productplan
	err := in.DB.Table(db.TABLE_PRODUCTPLAN).Field(`Id, CONCAT(title, " [", begin, " ~ ", end, "]") AS Title`).Where(where).Order("Begin desc").Select(&plans)
	if err != nil {
		in.WriteErr(err)
		return
	}
	if len(plans) == 0 {
		where["end"] = []interface{}{"lt", date}
		err = in.DB.Table(db.TABLE_PRODUCTPLAN).Field(`Id, CONCAT(title, " [", begin, " ~ ", end, "]") AS Title`).Where(where).Order("Begin desc").Limit(5).Select(&plans)
		if err != nil {
			in.WriteErr(err)
			return
		}
	}
	result := protocol.GET_MSG_PROJECT_productplan_getPairsForStory_result()
	result.List = append(result.List, protocol.HtmlKeyValueStr{"", ""})
	for _, v := range plans {
		result.List = append(result.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Title})
	}
	in.SendResult(result)
	result.Put()
}
func productplan_getList(data *protocol.MSG_PROJECT_productplan_getList, in *protocol.Msg) {
	if data.Page < 1 {
		data.Page = 1
	}
	var productPlans []map[string]string
	var err error
	if data.Id == 0 {
		where := map[string]interface{}{
			"Product": data.ProductID,
			"Deleted": false,
		}
		if data.Branch > 0 {
			where["branch"] = data.Branch
		}
		switch data.BrowseType {
		case "unexpired":
			where["end"] = []interface{}{"gt", time.Now().Format("2006-01-02")}
		case "overdue":
			where["end"] = []interface{}{"le", time.Now().Format("2006-01-02")}
		}
		productPlans, err = in.DB.Table(db.TABLE_PRODUCTPLAN).Where(where).Limit((data.Page-1)*data.PerPage, data.PerPage).Order(data.Order).SelectMap()
		if err == nil && data.Total == 0 {
			data.Total, err = in.DB.Table(db.TABLE_PRODUCTPLAN).Where(where).Count()
			if err != nil {
				in.WriteErr(err)
				return
			}
		}
	} else {
		productPlans, err = in.DB.Table(db.TABLE_PRODUCTPLAN).Prepare().Where("Id=?", data.Id).SelectMap()
	}
	if err != nil {
		in.WriteErr(err)
		return
	}

	for _, productPlan := range productPlans {
		stories, err := in.DB.Table(db.TABLE_STORY).Field("Estimate").Where("json_contains(`Plan`,'" + productPlan["Id"] + "') and `Deleted` = 0").SelectMap()
		if err != nil {
			in.WriteErr(err)
			return
		}
		productPlan["stories"] = strconv.Itoa(len(stories))
		out := protocol.GET_MSG_TEST_bug_getCount()
		out.Where = map[string]interface{}{
			"Plan":    productPlan["Id"],
			"Deleted": false,
		}
		var result *protocol.MSG_TEST_buf_getCount_result
		if err = in.SendMsgWaitResult(0, out, &result); err != nil {
			in.WriteErr(err)
			return
		}
		productPlan["bugs"] = strconv.Itoa(result.Count)
		hour := 0
		for _, v := range stories {
			h, _ := strconv.Atoi(v["Estimate"])
			hour += h
		}
		productPlan["hour"] = strconv.Itoa(hour)
	}
	out := protocol.GET_MSG_PROJECT_productplan_getList_result()
	out.List = productPlans
	out.Total = data.Total
	in.SendResult(out)
	out.Put()
}
func productplan_getLast(data *protocol.MSG_PROJECT_productplan_getLast, in *protocol.Msg) {
	where := map[string]interface{}{
		"Deleted": false,
		"Product": data.ProductId,
		"end":     []interface{}{"ne", "2030-01-01"},
	}
	if data.Branch > 0 {
		where["Branch"] = data.Branch
	}

	res, err := in.DB.Table(db.TABLE_PRODUCTPLAN).Field("Id,Title,End").Where(where).Order("end desc").FindMap()
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_productplan_getLast_result()
	out.Result = res
	in.SendResult(out)
	out.Put()
}
func productplan_getPairs(data *protocol.MSG_PROJECT_productplan_getPairs, in *protocol.Msg) {
	where := map[string]interface{}{
		"product": data.ProductID,
		"Deleted": false,
	}
	if data.BranchID > 0 {
		where["Branch"] = []int32{0, data.BranchID}
	}
	if data.Expired == "unexpired" {
		where["end"] = []interface{}{"ge", time.Now().Format("2006-01-02")}
	}
	var list []*db.Productplan
	err := in.DB.Table(db.TABLE_PRODUCTPLAN).Where(where).Order("begin desc").Limit(0).Select(&list)
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_productplan_getPairs_result()
	out.List = append(out.List, protocol.HtmlKeyValueStr{})
	var ids []int32
	for _, v := range list {
		ids = append(ids, v.Id)
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Title})
	}
	if data.Expired == "unexpired" {
		if len(list) > 0 {
			where["Id"] = []interface{}{"not in", ids}
		}
		err = in.DB.Table(db.TABLE_PRODUCTPLAN).Field(`Id,CONCAT(title, " [", begin, " ~ ", end, "]") as Title`).Where(where).Order("begin desc").Limit(5).Select(&list)
		where["end"] = []interface{}{"lt", time.Now().Format("2006-01-02")}

		for _, v := range list {
			out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Title})
		}

	}
	in.SendResult(out)
	out.Put()
}
func productplan_insertUpdate(data *protocol.MSG_PROJECT_productplan_insertUpdate, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_productplan_insertUpdate_result()
	if data.Parent > 0 {
		c, err := in.DB.Table(db.TABLE_PRODUCTPLAN).Prepare().Where("Id=?", data.Parent).Count()
		if err != nil {
			in.WriteErr(err)
			return
		} else if c == 0 {
			out.Result = protocol.Err_ProjectProductPlanParentNotFound
			in.SendResult(out)
			return
		}
	}
	if data.Id == 0 {
		if id, err := in.DB.Table(db.TABLE_PRODUCTPLAN).Insert(data); err != nil {
			in.WriteErr(err)
		} else {
			out.Result = protocol.Success
			out.Id = int32(id)
			in.SendResult(out)
		}
	} else {
		err := in.DB.Table(db.TABLE_PRODUCTPLAN).Replace(data)
		if err != nil {
			in.WriteErr(err)
		} else {
			out.Result = protocol.Success
			out.Id = data.Id
			in.SendResult(out)
		}

	}
	out.Put()
}
func productplan_delete(data *protocol.MSG_PROJECT_productplan_delete, in *protocol.Msg) {
	_, err := in.DB.Table(db.TABLE_PRODUCTPLAN).Prepare().Where("Id=? and Product=? and Branch=?", data.Id, data.Product, data.Branch).Update("Deleted = 1")
	in.WriteErr(err)
}
