package handler

import (
	"errors"
	"fmt"
	"jachunPM_http/config"
	"jachunPM_http/db"
	"jachunPM_http/js"
	"libraries"
	"math/rand"
	"mysql"
	"protocol"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

func init() {
	httpHandlerMap["GET"]["/search/buildForm"] = get_search_buildForm
	httpHandlerMap["GET"]["/search/saveQuery"] = get_search_saveQuery
	httpHandlerMap["POST"]["/search/saveQuery"] = post_search_saveQuery
	httpHandlerMap["GET"]["/search/deleteQuery"] = get_search_deleteQuery
}

//原代码setSearchParams,相关config[module]["search"]
var searchParamsFunc = map[string]func(*TemplateData) (*searchParam, error){}
var searchFormId = uint32(rand.NewSource(time.Now().Unix()).Int63())

type searchParam struct {
	//额外的参数
	QueryID              int
	ActionURL            string
	Style                string
	OnMenuBar            string
	*config.ConfigSearch //必须包含的
}

func get_search_buildForm(data *TemplateData) (err error) {
	//$module = '', $fields = '', $params = '', $actionURL = '', $queryID = 0)
	var param *searchParam
	module := data.ws.Query("module")
	method := data.ws.Query("method")
	if f, ok := searchParamsFunc[module+"/"+method]; ok {
		param, err = f(data)
	}
	if err != nil {
		return
	}
	if param == nil {
		data.ws.WriteString(js.Alert(data.Lang["search"]["error"].(map[string]string)["notFoundParamsFunc"], module, method))
		return
	}
	queryID, _ := strconv.Atoi(data.ws.Query("queryID"))
	data.Data["actionURL"] = data.ws.Query("actionURL")
	data.Data["style"] = "full"
	data.Data["onMenuBar"] = "no"

	if module == "" && queryID == 0 {
		queryID = param.QueryID
	}
	if module == "" {
		module = param.Module
	}

	if data.Data["actionURL"] == "" {
		data.Data["actionURL"] = param.ActionURL
	}
	data.Data["style"] = param.Style
	data.Data["onMenuBar"] = param.OnMenuBar
	data.Data["searchFields"] = param.Fields
	fieldParams := search_setDefaultParams(data, param.Fields, param.Params)
	data.Data["fieldParams"] = fieldParams
	data.ws.Session().Set(module+"/"+method+"/fieldParams", fieldParams)
	data.Data["module"] = module
	data.Data["method"] = method
	data.Data["groupItems"] = data.Config["search"]["common"]["groupItems"].(int)
	data.Data["groupItems2"] = data.Config["search"]["common"]["groupItems"].(int) + 1
	data.Data["queryID"] = queryID
	var querys []*db.SearchQuery
	if err = HostConn.DB.Table(db.TABLE_SearchQuery).Where("Uid=? and Module=?", data.User.Id, module).Prepare().Select(&querys); err != nil {
		return
	}
	data.Data["queries"] = querys

	data.Data["formId"] = "searchForm-" + strconv.Itoa(int(atomic.AddUint32(&searchFormId, 1)))
	search_initSession(data, module, param.Fields, param.Params)
	templateOut("search.buildForm.html", data)
	return
}
func search_initSession(data *TemplateData, module string, fields []protocol.HtmlKeyValueStr, fieldParams map[string]*config.ConfigSearchParams) {
	formSessionName := module + "Form"
	var queryForm map[string]string
	if ok := data.ws.Session().Get(formSessionName, &queryForm); !ok {
		queryForm = make(map[string]string)
		if fields[0].Key != "" && fields[0].Value != "" {
			fields = append([]protocol.HtmlKeyValueStr{{"", ""}}, fields...)
		}
		for i := 1; i <= data.Config["search"]["common"]["groupItems"].(int)*2 && i < len(fields); i++ {
			operator := "="
			if v, ok := fieldParams[fields[i].Key]; ok && v.Operator != "" {
				operator = v.Operator
			}
			queryForm["field"+strconv.Itoa(i)] = fields[i].Key
			queryForm["andOr"+strconv.Itoa(i)] = "and"
			queryForm["operator"+strconv.Itoa(i)] = operator
			queryForm["value"+strconv.Itoa(i)] = ""
		}
		queryForm["groupAndOr"] = "and"
		data.ws.Session().Set(formSessionName, queryForm)
	}

	data.Data["formSession"] = queryForm
}
func search_setDefaultParams(data *TemplateData, fields []protocol.HtmlKeyValueStr, params map[string]*config.ConfigSearchParams) map[string]*config.ConfigSearchParams {
	var users []protocol.HtmlKeyValueStr
	for _, field := range fields {
		if v, ok := params[field.Key]; ok {
			if len(v.Values) > 0 && v.Values[0].Key != "" && v.Values[len(v.Values)-1].Key != "null" {
				v.Values = append(v.Values, protocol.HtmlKeyValueStr{"null", data.Lang["search"]["null"].(string)})
			}
			if v.ValueExt == "users" {
				if users == nil {
					users, _ = user_getPairs(data, "realname|noclosed")
				}
				v.Values = users
			}
		} else {
			params[field.Key] = &config.ConfigSearchParams{
				Operator: "=",
				Control:  "input",
			}
		}
	}

	return params
}
func post_search_buildQuery(data *TemplateData, queryId int) (querystr string, err error) {
	if queryId > 0 {
		var query *db.SearchQuery
		err = HostConn.DB.Table(db.TABLE_SearchQuery).Where("Id=?", queryId).Prepare().Find(&query)
		if err != nil {
			return "", err
		} else if query != nil {
			return query.Where, err
		}

	}

	var fieldParams map[string]config.ConfigSearchParams
	module := data.ws.Post("module")
	method := data.ws.Post("method")
	if ok := data.ws.Session().Get(module+"/"+method+"/fieldParams", &fieldParams); !ok {
		//尝试从app获取
		module = data.App["moduleName"].(string)
		method = data.App["methodName"].(string)
		if ok = data.ws.Session().Get(module+"/"+method+"/fieldParams", &fieldParams); !ok || fieldParams == nil {
			return "", errors.New(fmt.Sprintf(data.Lang["search"]["error"].(map[string]string)["notFoundParamsFunc"], module, method))
		}
	}

	//
	formSessionName := module + "_whereStr"
	where := bufpool.Get().(*libraries.MsgBuffer)
	condition := bufpool.Get().(*libraries.MsgBuffer)
	defer func() {
		where.Reset()
		condition.Reset()
		bufpool.Put(where)
		bufpool.Put(condition)
	}()
	groupItems := data.Config["search"]["common"]["groupItems"].(int)
	groupAndOr := data.ws.Post("groupAndOr")

	if groupAndOr != "and" && groupAndOr != "or" {
		groupAndOr = "and"
	}
	for i := 1; i <= groupItems*2 && i < len(fieldParams); i++ {

		/* The and or between two groups. */
		if i == 1 {
			where.WriteString("(( 1  ")
		}
		if i == groupItems+1 {
			where.WriteString(" ) ")
			where.WriteString(groupAndOr)
			where.WriteString(" ( 1 ")
		}
		/* Set var names. */
		fieldName := "field" + strconv.Itoa(i)
		andOrName := "andOr" + strconv.Itoa(i)
		operatorName := "operator" + strconv.Itoa(i)
		valueName := "value" + strconv.Itoa(i)
		value := data.ws.Post(valueName)
		field := data.ws.Post(fieldName)

		/* Skip empty values. */
		switch value {
		case "":
			continue
		case "null":
			value = ""
		}

		/* Set and or. */
		andOr := data.ws.Post(andOrName)
		if andOr != "and" && andOr != "or" {
			andOr = "and"
		}
		condition.Reset()
		value = strings.Trim(value, " ")
		//转换value的动态值
		switch value {
		case "$lastMonth":
			now := data.Time
			lastmonthfirst := now.AddDate(0, -1, -now.Day()+1)
			condition.WriteString("between '")
			condition.WriteString(lastmonthfirst.Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString("' and '")
			condition.WriteString(lastmonthfirst.AddDate(0, 1, -1).Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString("'")
		case "$thisMonth":
			now := data.Time
			thismonthfirst := now.AddDate(0, 0, -now.Day()+1)
			condition.WriteString("between '")
			condition.WriteString(thismonthfirst.Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString("' and '")
			condition.WriteString(thismonthfirst.AddDate(0, 1, -1).Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString("'")
		case "$lastWeek":
			now := data.Time
			offset := int(time.Monday - now.Weekday())
			if offset > 0 {
				offset = -6
			}
			lastweekfirst := now.AddDate(0, 0, offset-7)
			condition.WriteString("between '")
			condition.WriteString(lastweekfirst.Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString("' and '")
			condition.WriteString(lastweekfirst.AddDate(0, 0, 6).Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString("'")
		case "$thisWeek":
			now := data.Time
			offset := int(time.Monday - now.Weekday())
			if offset > 0 {
				offset = -6
			}
			thisweekfirst := now.AddDate(0, 0, offset)
			condition.WriteString("between '")
			condition.WriteString(thisweekfirst.Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString("' and '")
			condition.WriteString(thisweekfirst.AddDate(0, 0, 6).Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString("'")
		case "$yesterday":
			now := data.Time
			yesterday := now.AddDate(0, 0, -1)
			condition.WriteString("between '")
			condition.WriteString(yesterday.Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString(" 00:00:00' and '")
			condition.WriteString(yesterday.Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString(" 23:59:59'")
		case "$today":
			now := data.Time
			condition.WriteString("between '")
			condition.WriteString(now.Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString(" 00:00:00' and '")
			condition.WriteString(now.Format(protocol.TIMEFORMAT_MYSQLDATE))
			condition.WriteString(" 23:59:59'")
		case "$@me":
			value = data.User.Account
			break
		default:
			operator := data.ws.Post(operatorName)
			find := false
			for _, v := range data.Lang["search"]["operators"].([]protocol.HtmlKeyValueStr) {
				if v.Key == operator {
					find = true
					break
				}
			}
			if !find {
				operator = "="
			}

			switch operator {
			case "include":
				condition.WriteString(" LIKE ")
				condition.WriteString(mysql.Getvalue("%" + value + "%"))
			case "notinclude":
				condition.WriteString(" NOT LIKE ")
				condition.WriteString(mysql.Getvalue("%" + value + "%"))
			case "belong":
				if field == "module" {
					//$allModules = $this->loadModel('tree')->getAllChildId($value);
					//if($allModules) $condition = helper::dbIN($allModules);
				} else if field == "dept" {
					deptID, _ := strconv.Atoi(value)
					allDepts, err := dept_getAllChildID(int32(deptID))
					if err != nil {

						return "", errors.New(fmt.Sprintf(data.Lang["dept"]["error"].(map[string]string)["ErrDeptInfo"], err))
					}
					mysql.MysqlBuild_in_value(allDepts, condition)
				} else {
					condition.WriteString(" = ")
					mysql.Getvalue(value)
				}
			default:
				if operator == "=" && libraries.Preg_match(`^[0-9]{4}-[0-9]{1,2}-[0-9]{1,2}$`, value) {
					condition.WriteString(" >= '")
					condition.WriteString(value)
					condition.WriteString("' and ")
					condition.WriteString(field)
					condition.WriteString(" <= '")
					condition.WriteString(value)
					condition.WriteString(" 23:59:59'")
				} else {
					condition.WriteString(operator)
					condition.WriteString(" ")
					mysql.Getvalue(value)
				}

			}
		}

		if condition.Len() > 0 {
			where.WriteString(" ")
			where.WriteString(andOr)
			where.WriteString(" ")
			where.WriteString(mysql.Getkey(field))
			where.WriteString(" ")
			where.Write(condition.Bytes())
		}
	}
	where.WriteString(" ))")
	data.ws.Session().Set("post_search_buildQuery_"+formSessionName, where.String())
	return where.String(), nil
}
func get_search_saveQuery(data *TemplateData) (err error) {
	data.Data["module"] = data.ws.Query("module")
	data.Data["onMenuBar"] = data.ws.Query("onMenuBar")
	templateOut("search.saveQuery.html", data)
	return nil
}
func post_search_saveQuery(data *TemplateData) (err error) {

	formSessionName := data.ws.Post("module") + "_whereStr"
	where := data.ws.Session().Load_str("post_search_buildQuery_" + formSessionName)
	insert := &db.SearchQuery{
		Uid:    data.User.Id,
		Title:  data.ws.Post("title"),
		Where:  where,
		Module: data.ws.Post("module"),
	}
	id, err := HostConn.DB.Table(db.TABLE_SearchQuery).Insert(insert)
	if err != nil {
		return
	}
	data.ws.WriteString(js.CloseModal("parent.parent", "", "function(){parent.parent.loadQueries("+strconv.Itoa(int(id))+", 0, '"+data.ws.Post("title")+"')}"))
	return nil
}
func get_search_deleteQuery(data *TemplateData) (err error) {
	HostConn.DB.Table(db.TABLE_SearchQuery).Where("Id=?", data.ws.Query("queryID")).Delete()
	data.ws.WriteString("success")
	return nil
}
