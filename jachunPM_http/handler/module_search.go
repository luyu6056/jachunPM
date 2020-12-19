package handler

import (
	"fmt"
	"jachunPM_http/config"
	"jachunPM_http/js"
	"libraries"
	"math/rand"
	"mysql"
	"protocol"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/luyu6056/gnet"
)

func init() {
	httpHandlerMap["GET"]["/search/buildForm"] = get_search_buildForm
	httpHandlerMap["POST"]["/search/buildQuery"] = post_search_buildQuery
}

var searchParamsFunc = map[string]func(*TemplateData) map[string]interface{}{}

var searchFormId = uint32(rand.NewSource(time.Now().Unix()).Int63())

func get_search_buildForm(data *TemplateData) gnet.Action {
	//$module = '', $fields = '', $params = '', $actionURL = '', $queryID = 0)
	var param map[string]interface{}
	module := data.ws.Query("module")
	method := data.ws.Query("method")
	if f, ok := searchParamsFunc[module+"/"+method]; ok {
		param = f(data)
	}
	if param == nil {
		data.ws.WriteString(js.Alert(data.Lang["search"]["error"].(map[string]string)["notFoundParamsFunc"], module, method))
		return gnet.None
	}
	queryID, _ := strconv.Atoi(data.ws.Query("queryID"))
	data.Data["actionURL"] = data.ws.Query("actionURL")
	data.Data["style"] = "full"
	data.Data["onMenuBar"] = "no"

	if module == "" && queryID == 0 {
		queryID, _ = param["queryID"].(int)
	}
	if module == "" {
		module, _ = param["module"].(string)
	}

	if data.Data["actionURL"] == "" {
		data.Data["actionURL"], _ = param["actionURL"].(string)
	}
	data.Data["style"], _ = param["style"].(string)
	data.Data["onMenuBar"], _ = param["onMenuBar"].(string)
	data.Data["searchFields"] = param["fields"].([]protocol.HtmlKeyValueStr)
	data.Data["fieldParams"] = search_setDefaultParams(data, param["fields"].([]protocol.HtmlKeyValueStr), param["params"].(map[string]config.ConfigSearchParams))

	data.Data["module"] = module
	data.Data["method"] = method
	data.Data["groupItems"] = data.Config["search"]["common"]["groupItems"].(int)
	data.Data["groupItems2"] = data.Config["search"]["common"]["groupItems"].(int) + 1
	data.Data["queryID"] = queryID
	// $this->view->queries      = $this->search->getQueryPairs($module);

	data.Data["formId"] = "searchForm-" + strconv.Itoa(int(atomic.AddUint32(&searchFormId, 1)))
	search_initSession(data, module, param["fields"].([]protocol.HtmlKeyValueStr), param["params"].(map[string]config.ConfigSearchParams))
	templateOut("search.buildForm.html", data)
	return gnet.None
}
func search_initSession(data *TemplateData, module string, fields []protocol.HtmlKeyValueStr, fieldParams map[string]config.ConfigSearchParams) {
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
func search_setDefaultParams(data *TemplateData, fields []protocol.HtmlKeyValueStr, params map[string]config.ConfigSearchParams) map[string]config.ConfigSearchParams {

	for _, field := range fields {
		if v, ok := params[field.Key]; ok {
			if len(v.Values) > 0 && v.Values[0].Key != "" && v.Values[len(v.Values)-1].Key != "null" {
				v.Values = append(v.Values, protocol.HtmlKeyValueStr{"null", data.Lang["search"]["null"].(string)})
			}
		}
	}

	return params
}
func post_search_buildQuery(data *TemplateData) (action gnet.Action) {

	var param map[string]interface{}
	module := data.ws.Post("module")
	method := data.ws.Post("method")
	if ok := data.ws.Session().Get(module+"/"+method, &param); !ok {
		data.ws.WriteString(js.Alert(data.Lang["search"]["error"].(map[string]string)["notFoundParamsFunc"], module, method))
		return gnet.None
	}
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
	fieldParams := param["params"].(map[string]config.ConfigSearchParams)
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
			value = data.App["user"].(*protocol.MSG_USER_INFO_cache).Account
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
						data.ws.WriteString(js.Alert(data.Lang["dept"]["error"].(map[string]string)["ErrDeptInfo"], err))
						return
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
	fmt.Println(where.String())
	data.ws.Session().Store(module+"/"+method+"/Query", where.String())
	data.ws.Session().Store(module+"/"+method+"/From", data.ws.GetAllPost())
	data.ws.WriteString(js.Location(data.ws.Post("actionURL"), "parent"))
	return
}
