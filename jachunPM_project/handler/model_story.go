package handler

import (
	"jachunPM_project/db"
	"protocol"
)

func story_getProductStories(productID int32, branch int32, modules []int32, status []string, sort string, page int, perpage int, total *int) (list []*protocol.MSG_PROJECT_story, err error) {

	var where = map[string]interface{}{"product": productID, "deleted": false}
	if len(modules) > 0 {
		where["module"] = modules
	}
	if len(status) > 0 {
		where["status"] = status
	}
	if branch > 0 {
		where["branch"] = branch
	}
	err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Order(sort).Limit((page-1)*perpage, perpage).Select(&list)
	if err != nil {
		return
	}
	if *total == 0 {
		*total, err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Count()
	}
	return
}
func story_getByPlan(productID, branch int32, modules []int32, plan, orderBy string, page int, perpage int, total *int) (list []*protocol.MSG_PROJECT_story, err error) {
	return story_getByField(productID, branch, modules, "plan", plan, orderBy, page, perpage, total)
}

func story_getByField(productID, branch int32, modules []int32, fieldName, fieldValue, orderBy string, page int, perpage int, total *int, operators ...string) (list []*protocol.MSG_PROJECT_story, err error) {
	operator := "equal"
	if len(operators) == 1 {
		operator = operators[0]
	}
	if fieldName == "" {
		return
	}
	var where = map[string]interface{}{"product": productID, "deleted": false}
	if branch > 0 {
		where["branch"] = branch
	}
	if len(modules) > 0 {
		where["module"] = modules
	}
	switch operator {
	case "equal":
		where[fieldName] = fieldValue
	case "include":
		where[fieldName] = []string{"like", "%" + fieldValue + "%"}
	}
	err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Order(orderBy).Limit((page-1)*perpage, perpage).Select(&list)
	if err != nil {
		return
	}
	if *total == 0 {
		*total, err = HostConn.DB.Table(db.TABLE_STORY).Where(where).Count()
	}
	return
}
