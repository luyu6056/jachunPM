package handler

import (
	"html/template"
	"libraries"
	"protocol"
	"strconv"
	"strings"
	"time"
)

func init() {
	httpHandlerModuleInit["GET"]["attend"] = get_oa_featurebar
	httpHandlerMap["GET"]["/attend/personal"] = get_attend_personal
	httpHandlerMap["GET"]["/attend/company"] = get_attend_company
	httpHandlerMap["GET"]["/attend/settings"] = get_attend_settings
	httpHandlerMap["POST"]["/attend/settings"] = post_attend_settings
}
func attendTemplateFuncs() {
	global_Funcs["attend_printHour"] = func(data *TemplateData, hour float32, status string) string {
		if v, ok := data.Lang["attend"]["statusList"].(map[string]string)[status]; ok {
			status = v
		}
		switch hour {
		case protocol.AttendAM:
			return data.Lang["attend"]["AM"].(string) + status //上午请假
		case protocol.AttendPM:
			return data.Lang["attend"]["PM"].(string) + status //下午请假
		default:
			return status + libraries.I2S(hour) + data.Lang["attend"]["d"].(string)
		}
	}
	global_Funcs["attend_getStatInfoAttend"] = func(stat *protocol.MSG_OA_attend_statInfo, key int64) (res *protocol.MSG_OA_attend_info) {
		if v, ok := stat.Attend[key]; ok {
			res = v
		} else {
			res = &protocol.MSG_OA_attend_info{
				Date: time.Unix(key, 0),
			}
		}
		if v, ok := stat.AbsentDates[key]; res.Status == "" && ok {
			if v.AmAbsent && v.PmAbsent {
				res.Status = "absent"
			} else if v.AmAbsent || v.PmAbsent {
				res.Status = "halfAbsent"
			} else {
				res.Status = "normal"
			}
		}
		if res.SignIn != "" && res.SignIn != "00:00:00" {
			res.SignIn = res.SignIn[:5]
		} else {
			res.SignIn = ""
		}
		if res.SignOut != "" && res.SignOut != "00:00:00" {
			res.SignOut = res.SignOut[:5]
		} else {
			res.SignOut = ""
		}
		return
	}
	global_Funcs["attend_printAttendExtDesc"] = func(data *TemplateData, stat *protocol.MSG_OA_attend_statInfo, key int64) template.HTML {
		if v, ok := stat.AttendExtDesc[key]; ok {
			buf := bufpool.Get().(*libraries.MsgBuffer)
			buf.Reset()
			for _, desc := range v {
				switch desc.Day {
				case protocol.AttendAM:
					buf.WriteString(data.Lang["attend"]["AM"].(string))

				case protocol.AttendPM:
					buf.WriteString(data.Lang["attend"]["PM"].(string))
				default:
					buf.WriteString(strconv.FormatFloat(float64(desc.Day), 'g', -1, 32))
					buf.WriteString(data.Lang["attend"]["d"].(string))
				}
				if v, ok := data.Lang["attend"]["statusList"].(map[string]string)[desc.Type]; ok {
					buf.WriteString(v)
				} else {
					buf.WriteString("attend_printAttendExtDesc未处理type:" + desc.Type)
				}
				buf.WriteString("<br>")
			}
			res := buf.String()
			buf.Reset()
			bufpool.Put(res)
			return template.HTML(res)
		}
		return template.HTML("")
	}
}
func get_oa_featurebar(data *TemplateData) error {
	buf := bufpool.Get().(*libraries.MsgBuffer)
	buf.Reset()
	buf.WriteString("<div id='featurebar'><ul class='nav'>")
	for _, kv := range data.Lang[data.App["moduleName"].(string)]["featurebar"].([]protocol.HtmlKeyValueStr) {
		link := strings.Split(kv.Value, "|")
		name, currentModule, currentMethod, params := link[0], link[1], link[2], link[3]
		if hasPriv(data, currentModule, currentMethod) {
			buf.WriteString("<li id='")
			buf.WriteString(kv.Key)
			buf.WriteString("'")
			if kv.Key == data.App["methodName"].(string) {
				buf.WriteString(" class='active'")
			}
			buf.WriteString(">")
			buf.WriteString(html_a(createLink(currentModule, currentMethod, params), name))
			buf.WriteString("</li>")
		}
	}
	buf.WriteString("</div>")
	data.Data["featurebar"] = template.HTML(buf.String())
	buf.Reset()
	bufpool.Put(buf)
	return nil
}
func get_attend_personal(data *TemplateData) (err error) {
	date, err := time.ParseInLocation("2006-01", data.ws.Query("date"), time.Local)
	if err != nil {
		date, _ = time.ParseInLocation("2006-01", time.Now().Format("2006-01"), time.Local)
	}
	startDate := date.AddDate(0, 0, -1*date.Day()+1)
	endDate := date.AddDate(0, 1, 0)

	data.Data["currentYear"] = date.Format("2006")
	data.Data["currentMonth"] = date.Format("01")
	getAttends := protocol.GET_MSG_OA_attend_getByAccount()
	getAttends.Uid = data.User.Id
	getAttends.StartDate = startDate
	getAttends.EndDate = endDate
	dayNum := endDate.AddDate(0, 0, -1).Day()
	data.Data["weekNum"] = dayNum / 7
	var attends *protocol.MSG_OA_attend_getByAccount_result
	if err = data.SendMsgWaitResultToDefault(getAttends, &attends); err != nil {
		return
	}
	getAllmonth := protocol.GET_MSG_OA_attend_getAllMonth()
	getAllmonth.Uids = []int32{data.User.Id}
	var getAllmonthResult *protocol.MSG_OA_attend_getAllMonth_result
	if err = data.SendMsgWaitResultToDefault(getAllmonth, &getAllmonthResult); err != nil {
		return
	}
	data.Data["yearList"] = getAllmonthResult.List
	data.Data["title"] = data.Lang["attend"]["personal"]
	firstDayIndex := int(time.Monday) //礼拜一是第一天
	data.Data["firstDayIndex"] = strconv.Itoa(firstDayIndex)
	//生成周列表
	startDate = startDate.AddDate(0, 0, -1*int(startDate.Weekday())+firstDayIndex)
	var weekList [][]protocol.HtmlKeyValueInterface
	index := 0
	now := time.Now()
	for endDate.After(startDate) {
		var week []protocol.HtmlKeyValueInterface
		for j := 0; j < 7; j++ {
			var find *protocol.MSG_OA_attend_info
			for _, attend := range attends.List[index:] {
				if attend.Date.After(startDate) || attend.Date.After(now) {
					break
				} else if attend.Date == startDate {
					find = attend
					if (attend.SignIn == "" || attend.SignIn == "00:00:00") && attend.Status != "normal" {
						if v, ok := data.Lang["attend"]["statusList"].(map[string]string)[attend.Status]; ok {
							attend.SignIn = v //显示请假等异常情况
						}
					} else if len(attend.SignIn) > 5 {
						attend.SignIn = attend.SignIn[:5] //显示时:分
					}
					if (attend.SignOut == "" || attend.SignOut == "00:00:00") && attend.Status != "normal" {
						if v, ok := data.Lang["attend"]["statusList"].(map[string]string)[attend.Status]; ok {
							attend.SignOut = v //显示请假等异常情况
						}
					} else if len(attend.SignOut) > 5 {
						attend.SignOut = attend.SignOut[:5] //显示时:分
					}
					index++
					break
				}
			}
			if find == nil {
				find = &protocol.MSG_OA_attend_info{Date: startDate}
			}
			week = append(week, protocol.HtmlKeyValueInterface{strconv.Itoa(int(startDate.Weekday())), find})
			startDate = startDate.AddDate(0, 0, 1)
		}
		weekList = append(weekList, week)
	}

	data.Data["weekList"] = weekList
	templateOut("attend.personal.html", data)
	getAttends.Put()
	attends.Put()
	getAllmonth.Put()
	getAllmonthResult.Put()
	return
}
func get_attend_company(data *TemplateData) (err error) {
	date, err := time.ParseInLocation("2006-01", data.ws.Query("date"), time.Local)
	if err != nil {
		date, _ = time.ParseInLocation("2006-01", time.Now().Format("2006-01"), time.Local)
	}

	startDate := date.AddDate(0, 0, -1*date.Day()+1)
	endDate := date.AddDate(0, 1, 0)
	data.Data["currentYear"] = date.Format("2006")
	data.Data["currentMonth"] = date.Format("01")
	getAllmonth := protocol.GET_MSG_OA_attend_getAllMonth()
	//getAllmonth.Uids = []int32{data.User.Id}
	var getAllmonthResult *protocol.MSG_OA_attend_getAllMonth_result
	if err = data.SendMsgWaitResultToDefault(getAllmonth, &getAllmonthResult); err != nil {
		return
	}
	data.Data["yearList"] = getAllmonthResult.List
	deptList, err := dept_getPairs(data)
	if err != nil {
		return
	}
	users, err := user_getAllcache(data)
	if err != nil {
		return
	}
	out := protocol.GET_MSG_OA_attend_computeStat()
	for i := len(users) - 1; i >= 0; i-- {
		user := users[i]
		if user.Deleted {
			users = append(users[:i], users[i+1:]...)
			continue
		}
		for _, sid := range strings.Split(data.Config["attend"]["custom"]["noAttendDepts"].(string), "") {
			if sid == strconv.Itoa(int(user.Dept)) {
				users = append(users[:i], users[i+1:]...)
				continue
			}
		}
		out.Uids = append(out.Uids, user.Id)
	}
	out.Year = data.Data["currentYear"].(string)
	out.Month = data.Data["currentMonth"].(string)
	var result *protocol.MSG_OA_attend_computeStat_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	statDept := make(map[int32]map[int32]*protocol.MSG_OA_attend_statInfo)
	for _, user := range users {
		if statDept[user.Dept] == nil {
			statDept[user.Dept] = make(map[int32]*protocol.MSG_OA_attend_statInfo)
		}
		if v, ok := result.Stat[user.Id]; ok {
			statDept[user.Dept][user.Id] = v
		} else {
			statDept[user.Dept][user.Id] = protocol.GET_MSG_OA_attend_statInfo()
		}
	}
	var weekDaylist []string
	var dayList []protocol.HtmlKeyValueInterface
	for i := startDate; i.Unix() <= endDate.Unix(); i = i.AddDate(0, 0, 1) {
		w := strconv.Itoa(int(i.Weekday()))
		for _, kv := range data.Lang["datepicker"]["abbrDayNames"].([]protocol.HtmlKeyValueStr) {
			if w == kv.Key {
				weekDaylist = append(weekDaylist, kv.Value)
			}
		}
		dayList = append(dayList, protocol.HtmlKeyValueInterface{strconv.Itoa(int(i.Day())), i.Unix()})
	}
	data.Data["dayList"] = dayList
	data.Data["weekDaylist"] = weekDaylist
	data.Data["date"] = date
	data.Data["statDept"] = statDept
	data.Data["title"] = data.Lang["attend"]["department"].(string)
	data.Data["deptList"] = deptList
	data.Data["users"] = users
	data.Data["company"] = true
	data.Lang["attend"]["abbrStatusList"].(map[string]string)["rest"] = ""
	templateOut("attend.browse.html", data)
	out.Put()
	result.Put()
	getAllmonth.Put()
	getAllmonthResult.Put()
	return
}
func get_attend_settings(data *TemplateData) (err error) {

	data.Data["title"] = data.Lang["attend"]["settings"].(string)
	data.Data["beginDate"] = data.Config["attend"]["custom"]["beginDate"]
	data.Data["notAllowSignInLimit"] = data.Config["attend"]["custom"]["notAllowSignInLimit"]
	data.Data["notAllowSignOutLimit"] = data.Config["attend"]["custom"]["notAllowSignOutLimit"]
	data.Data["signInLimit"] = data.Config["attend"]["custom"]["signInLimit"]
	data.Data["signOutLimit"] = data.Config["attend"]["custom"]["signOutLimit"]
	data.Data["workingDays"] = strings.Split(data.Config["attend"]["custom"]["workingDays"].(string), ",")
	data.Data["workingHours"] = data.Config["attend"]["custom"]["workingHours"]
	data.Data["mustSignOut"] = data.Config["attend"]["custom"]["mustSignOut"]
	data.Data["noAttendUsers"] = data.Config["attend"]["custom"]["noAttendUsers"]
	data.Data["noAttendDepts"] = data.Config["attend"]["custom"]["noAttendDepts"]
	data.Data["halfAbsendMin"] = data.Config["attend"]["custom"]["halfAbsendMin"]
	data.Data["absendMin"] = data.Config["attend"]["custom"]["absendMin"]
	data.Data["ip"] = data.Config["attend"]["custom"]["ip"]
	data.Data["reviewedBy"] = data.Config["attend"]["custom"]["reviewedBy"]
	if data.Data["users"], err = user_getPairs(data, "noempty,noclosed,nodeleted,noforbidden"); err != nil {
		return
	} else {
		data.Data["users"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"", data.Lang["dept"]["manager"].(string)}}, data.Data["users"].([]protocol.HtmlKeyValueStr)...)
	}
	if data.Data["depts"], err = dept_getPairs(data); err != nil {
		return
	}
	templateOut("attend.settings.html", data)

	return
}
func post_attend_settings(data *TemplateData) (err error) {
	out := protocol.GET_MSG_USER_config_savelist()
	for key, value := range data.ws.GetAllPost() {
		config := protocol.GET_MSG_USER_config_save()
		config.Uid = protocol.SYSTEMUSER
		config.Key = key
		config.Section = "custom"
		config.Module = "attend"
		config.Type = "add"
		config.Value = strings.Join(value, ",")
		out.List = append(out.List, config)
	}
	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		data.ajaxResult(false, err.Error())
		return nil
	}
	data.ajaxResult(true, data.Lang["common"]["saveSuccess"], "reload")
	return
}
