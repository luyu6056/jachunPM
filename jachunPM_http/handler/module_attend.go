package handler

import (
	"errors"
	"html/template"
	"libraries"
	"protocol"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func init() {
	httpHandlerModuleInit["GET"]["attend"] = get_oa_featurebar
	httpHandlerModuleInit["POST"]["attend"] = get_oa_featurebar
	httpHandlerMap["GET"]["/attend/personal"] = get_attend_personal
	httpHandlerMap["GET"]["/attend/company"] = get_attend_company
	httpHandlerMap["GET"]["/attend/settings"] = get_attend_settings
	httpHandlerMap["POST"]["/attend/settings"] = post_attend_settings
	httpHandlerMap["GET"]["/attend/detail"] = get_attend_detail
	httpHandlerMap["POST"]["/attend/detail"] = get_attend_detail
	httpHandlerMap["GET"]["/attend/browsereview"] = get_attend_browseReview
	httpHandlerMap["GET"]["/attend/edit"] = get_attend_edit
	httpHandlerMap["POST"]["/attend/edit"] = post_attend_edit
	httpHandlerMap["GET"]["/attend/review"] = get_attend_review
	httpHandlerMap["POST"]["/attend/cancel"] = post_attend_cancel
	httpHandlerMap["GET"]["/attend/reject"] = get_attend_reject
	httpHandlerMap["POST"]["/attend/reject"] = post_attend_reject
	httpHandlerMap["GET"]["/attend/exportDetail"] = get_attend_exportDetail
	httpHandlerMap["POST"]["/attend/exportDetail"] = post_attend_exportDetail
	httpHandlerMap["GET"]["/attend/stat"] = get_attend_stat
	httpHandlerMap["GET"]["/attend/zkteco"] = get_attend_zkteco
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
	global_Funcs["attend_detail_printStatus"] = func(data *TemplateData, attend *protocol.MSG_OA_attend_detail_info) template.HTML {
		var str []string
		for status, hour := range attend.HoursList {
			if v, ok := data.Lang["attend"]["statusList"].(map[string]string)[status]; ok {
				status = v
			}
			switch hour {
			case protocol.AttendAM:
				str = append(str, data.Lang["attend"]["AM"].(string)+status) //上午请假
			case protocol.AttendPM:
				str = append(str, data.Lang["attend"]["PM"].(string)+status) //下午请假
			default:
				str = append(str, status+libraries.I2S(hour)+data.Lang["attend"]["d"].(string))
			}
		}
		if attend.Status == "late" || attend.Status == "early" || attend.Status == "both" {
			if attend.LateMin > 0 {
				str = append(str, data.Lang["attend"]["statusList"].(map[string]string)["late"]+strconv.Itoa(int(attend.LateMin))+data.Lang["attend"]["m"].(string))
			}
			if attend.EarlyMin > 0 {
				str = append(str, data.Lang["attend"]["statusList"].(map[string]string)["early"]+strconv.Itoa(int(attend.EarlyMin))+data.Lang["attend"]["m"].(string))
			}
		}

		if len(str) == 0 {
			if v, ok := data.Lang["attend"]["statusList"].(map[string]string)[attend.Status]; ok {
				return template.HTML(v)
			}
		}
		return template.HTML(strings.Join(str, "<br>"))
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
	beginDate := date.AddDate(0, 0, -1*date.Day()+1)
	endDate := date.AddDate(0, 1, 0)

	data.Data["currentYear"] = date.Format("2006")
	data.Data["currentMonth"] = date.Format("01")
	getAttends := protocol.GET_MSG_OA_attend_getByAccount()
	getAttends.Uids = []int32{data.User.Id}
	getAttends.BeginDate = beginDate
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
	beginDate = beginDate.AddDate(0, 0, -1*int(beginDate.Weekday())+firstDayIndex)
	var weekList [][]protocol.HtmlKeyValueInterface
	index := 0
	now := time.Now()
	for endDate.After(beginDate) {
		var week []protocol.HtmlKeyValueInterface
		for j := 0; j < 7; j++ {
			var find *protocol.MSG_OA_attend_info
			for _, attend := range attends.List[index:] {
				if attend.Date.After(beginDate) || attend.Date.After(now) {
					break
				} else if attend.Date == beginDate {
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
				find = &protocol.MSG_OA_attend_info{Date: beginDate}
			}
			week = append(week, protocol.HtmlKeyValueInterface{strconv.Itoa(int(beginDate.Weekday())), find})
			beginDate = beginDate.AddDate(0, 0, 1)
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

	beginDate := date.AddDate(0, 0, -1*date.Day()+1)
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
f:
	for i := len(users) - 1; i >= 0; i-- {
		user := users[i]
		if user.Deleted {
			users = append(users[:i], users[i+1:]...)
			continue f
		}
		for _, deptid := range strings.Split(data.Config["attend"]["custom"]["noAttendDepts"].(string), ",") {
			if deptid == strconv.Itoa(int(user.Dept)) {
				users = append(users[:i], users[i+1:]...)
				continue f
			}
		}
		for _, id := range strings.Split(data.Config["attend"]["custom"]["noAttendUsers"].(string), ",") {
			if id == strconv.Itoa(protocol.DEPTManager) {
				if dept, err := dept_getAll(data); err != nil {
					return err
				} else {
					for _, d := range dept {
						if d.Manager == user.Id {
							users = append(users[:i], users[i+1:]...)
							continue f
						}
					}
				}
			}
			if id == strconv.Itoa(int(user.Id)) {
				users = append(users[:i], users[i+1:]...)
				continue f
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
	for i := beginDate; i.Unix() <= endDate.Unix(); i = i.AddDate(0, 0, 1) {
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
	data.Data["noAttendUsers"] = strings.Split(data.Config["attend"]["custom"]["noAttendUsers"].(string), ",")
	data.Data["noAttendDepts"] = strings.Split(data.Config["attend"]["custom"]["noAttendDepts"].(string), ",")
	data.Data["halfAbsendMin"] = data.Config["attend"]["custom"]["halfAbsendMin"]
	data.Data["absendMin"] = data.Config["attend"]["custom"]["absendMin"]
	data.Data["ip"] = data.Config["attend"]["custom"]["ip"]
	data.Data["reviewedBy"] = data.Config["attend"]["custom"]["reviewedBy"]
	if data.Data["users"], err = user_getPairs(data, "noempty,noclosed,nodeleted,noforbidden"); err != nil {
		return
	} else {
		data.Data["users"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{strconv.Itoa(protocol.DEPTManager), data.Lang["dept"]["manager"].(string)}}, data.Data["users"].([]protocol.HtmlKeyValueStr)...)
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
func get_attend_detail(data *TemplateData) (err error) {
	deptID, _ := strconv.Atoi(data.ws.Post("deptID"))
	userID, _ := strconv.Atoi(data.ws.Post("userID"))
	beginDate, err1 := time.ParseInLocation("2006-01-02", data.ws.Post("begin"), time.Local)
	endDate, err2 := time.ParseInLocation("2006-01-02", data.ws.Post("end"), time.Local)
	date, err := time.ParseInLocation("2006-01", data.ws.Query("date"), time.Local)
	if err != nil {
		date, _ = time.ParseInLocation("2006-01", time.Now().Format("2006-01"), time.Local)
	}
	if err1 != nil || err2 != nil {
		beginDate = date.AddDate(0, 0, -1*date.Day()+1)
		endDate = date.AddDate(0, 1, 0)
	}

	data.Data["currentYear"] = beginDate.Format("2006")
	data.Data["currentMonth"] = beginDate.Format("01")
	if deptList, err := dept_getPairs(data); err != nil {
		return err
	} else {
		data.Data["deptList"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{}}, deptList...)
	}
	data.ws.Session().Set("attendDeptID", deptID)
	data.ws.Session().Set("attendUserID", userID)
	var userIds []int32

	if alluser, err := user_getAllcache(data); err != nil {
		return err
	} else if userID == 0 {

		protocol.Order_user(alluser, func(a, b *protocol.MSG_USER_INFO_cache) bool {
			if a.Dept == b.Dept {
				nameA := a.Realname
				nameB := b.Realname
				if nameA == "" {
					nameA = a.Account
				}
				if nameB == "" {
					nameB = b.Account
				}
				return protocol.Order_Pinyin(nameA, nameB, true)
			}
			return a.Dept < b.Dept
		})

		var userList = []protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{}}
	f:
		for _, user := range alluser {
			for _, id := range strings.Split(data.Config["attend"]["custom"]["noAttendUsers"].(string), ",") {
				if id == strconv.Itoa(protocol.DEPTManager) {
					if dept, err := dept_getAll(data); err != nil {
						return err
					} else {
						for _, d := range dept {
							if d.Manager == user.Id {
								continue f
							}
						}
					}
				}
				if id == strconv.Itoa(int(user.Id)) {
					continue f
				}
			}
			if (deptID == 0 || deptID == int(user.Dept)) && user.Dept != 0 {
				name := user.Realname
				if name == "" {
					name = user.Account
				}
				userList = append(userList, protocol.HtmlKeyValueStr{strconv.Itoa(int(user.Id)), name})
				userIds = append(userIds, user.Id)
			}
		}
		protocol.Order_htmlkvStr(userList, func(a, b protocol.HtmlKeyValueStr) bool {
			return protocol.Order_Pinyin(a.Value, b.Value, true)
		})
		data.Data["userList"] = userList
	} else {
		userIds = []int32{int32(userID)}
	}
	out := protocol.GET_MSG_OA_attend_detail()
	out.BeginDate = beginDate
	out.EndDate = endDate
	out.User = userIds
	var result *protocol.MSG_OA_attend_detail_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	data.Data["title"] = data.Lang["attend"]["detail"]
	data.Data["deptID"] = deptID
	data.Data["userID"] = userID
	data.Data["date"] = date
	data.Data["beginDate"] = beginDate
	data.Data["endDate"] = endDate
	data.Data["attends"] = result.List
	data.Lang["attend"]["abbrStatusList"].(map[string]string)["rest"] = ""
	var monthlist []string
	end := 12
	if beginDate.Format("2006") == time.Now().Format("2006") {
		end = int(time.Now().Month())
	}
	for i := 1; i <= end; i++ {
		if i < 10 {
			monthlist = append(monthlist, "0"+strconv.Itoa(i))
		} else {
			monthlist = append(monthlist, strconv.Itoa(i))
		}

	}
	data.Data["monthlist"] = monthlist
	templateOut("attend.detail.html", data)
	out.Put()
	result.Put()
	return
}
func get_attend_browseReview(data *TemplateData) (err error) {
	deptList, err := dept_getPairs(data)
	if err != nil {
		return
	}
	alluser, err := user_getAllcache(data)
	if err != nil {
		return
	}
	reviewedBy, _ := strconv.Atoi(data.Config["attend"]["custom"]["reviewedBy"].(string))
	out := protocol.GET_MSG_OA_attend_getWaitAttends()
	users := make(map[int32]*protocol.MSG_USER_INFO_cache)
	for _, u := range alluser {
		users[u.Id] = u
		out.Users = append(out.Users, u.Id)
	}
	if !(data.User.IsAdmin || reviewedBy == int(data.User.Id)) {
		out.Users = out.Users[:0]
		if managedDepts, err := dept_getDeptManagedByMe(data.User.Id); err != nil {
			return err
		} else {
			for _, deptid := range managedDepts {
				for _, u := range alluser {
					if u.Dept == deptid {
						out.Users = append(out.Users, u.Id)
					}
				}
			}
		}
	}
	if len(out.Users) > 0 {
		var result *protocol.MSG_OA_attend_getWaitAttends_result
		if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
			return
		}
		data.Data["attends"] = result.List
	}
	data.Data["title"] = data.Lang["attend"]["review"]
	data.Data["users"] = users
	data.Data["deptList"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"0", "/"}}, deptList...)
	templateOut("attend.browsereview.html", data)
	out.Put()
	return
}
func get_attend_edit(data *TemplateData) (err error) {
	date, err := time.ParseInLocation("2006-01-02", data.ws.Query("date"), time.Local)
	if err != nil {
		date, _ = time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	}
	out := protocol.GET_MSG_OA_attend_getByDate()
	out.Date = date
	out.Uid = data.User.Id

	var attend *protocol.MSG_OA_attend_info
	if err = data.SendMsgWaitResultToDefault(out, &attend); err != nil {
		return
	}
	if data.Data["user"], err = user_getPairs(data, "noletter"); err != nil {
		return
	}
	data.Data["title"] = data.Lang["attend"]["edit"]
	data.Data["attend"] = attend
	data.Data["date"] = date
	templateOut("attend.edit.html", data)
	out.Put()
	return
}
func post_attend_edit(data *TemplateData) (err error) {
	date, err := time.ParseInLocation("2006-01-02", data.ws.Query("date"), time.Local)
	if err != nil {
		date, _ = time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	}
	out := protocol.GET_MSG_OA_attend_update()
	out.Uid = data.User.Id
	manualIn := data.ws.Post("manualIn")
	manualOut := data.ws.Post("manualOut")
	if strings.Count(manualIn, ":") == 1 {
		manualIn += ":00"
	}
	if strings.Count(manualOut, ":") == 1 {
		manualOut += ":00"
	}
	if out.ManualIn, err = time.ParseInLocation("2006-01-02 15:04:05", date.Format("2006-01-02 ")+manualIn, time.Local); err != nil {
		data.ajaxResult(false, data.Lang["attend"]["error"].(map[string]string)["manualIn"])
		return nil
	}
	if out.ManualOut, err = time.ParseInLocation("2006-01-02 15:04:05", date.Format("2006-01-02 ")+manualOut, time.Local); err != nil {
		data.ajaxResult(false, data.Lang["attend"]["error"].(map[string]string)["manualOut"])
		return nil
	}
	out.Desc = data.ws.Post("desc")
	out.Date = date
	out.ReviewStatus = "wait"
	out.Reason = "normal"
	if err := data.SendMsgWaitResultToDefault(out, nil); err != nil {
		data.ajaxResult(false, err.Error())
		return nil
	}
	data.ajaxResult(true, data.Lang["common"]["saveSuccess"], createLink("attend", "personal", nil))
	return
}
func get_attend_review(data *TemplateData) (err error) {
	attendID, _ := strconv.Atoi(data.ws.Query("attendID"))
	out := protocol.GET_MSG_OA_attend_getById()
	out.Id = int32(attendID)
	var result *protocol.MSG_OA_attend_getbyId_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	reviewedBy, _ := strconv.Atoi(data.Config["attend"]["custom"]["reviewedBy"].(string))
	if !(data.User.IsAdmin || reviewedBy == int(data.User.Id)) {
		if managedDepts, err := dept_getDeptManagedByMe(data.User.Id); err != nil {
			return err
		} else {
			find := false
			user := HostConn.GetUserCacheById(result.Info.Uid)
			if user == nil {
				data.ajaxResult(false, data.Lang["attend"]["error"].(map[string]string)["reviewNotAllow"])
				return nil
			}
			for _, deptid := range managedDepts {
				if user.Dept == deptid {
					find = true
					break
				}
			}
			if !find {
				data.ajaxResult(false, data.Lang["attend"]["error"].(map[string]string)["reviewNotAllow"])
				return nil
			}
		}
	}
	update := protocol.GET_MSG_OA_attend_update()
	update.Uid = data.User.Id
	update.ManualIn, _ = time.ParseInLocation("2006-01-02 15:04:05", result.Info.Date.Format("2006-01-02 ")+result.Info.ManualIn, time.Local)
	update.ManualOut, _ = time.ParseInLocation("2006-01-02 15:04:05", result.Info.Date.Format("2006-01-02 ")+result.Info.ManualOut, time.Local)
	update.Desc = result.Info.Desc
	update.Date = result.Info.Date
	update.ReviewStatus = data.ws.Query("status")
	update.Reason = result.Info.Reason
	update.ReviewedBy = data.User.Id
	libraries.DebugLog("%+v", update)
	if err := data.SendMsgWaitResultToDefault(update, nil); err != nil {
		data.ajaxResult(false, err.Error())
		return nil
	}
	data.ajaxResult(true, data.Lang["common"]["saveSuccess"], createLink("attend", "browsereview", nil))
	return
}
func post_attend_cancel(data *TemplateData) (err error) {
	out := protocol.GET_MSG_OA_attend_update()
	out.Uid = data.User.Id
	date, err := time.ParseInLocation("2006-01-02", data.ws.Query("date"), time.Local)
	if err != nil {
		date, _ = time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	}
	out.ManualIn = protocol.ZEROTIME
	out.ManualOut = protocol.ZEROTIME
	out.Desc = data.ws.Post("desc")
	out.Date = date
	out.ReviewStatus = ""
	out.Reason = ""
	if err := data.SendMsgWaitResultToDefault(out, nil); err != nil {
		data.ajaxResult(false, err.Error())
		return nil
	}
	data.ajaxResult(true, data.Lang["attend"]["cancel"], createLink("attend", "personal", nil))
	return

}
func get_attend_reject(data *TemplateData) (err error) {
	data.Data["title"] = data.Lang["attend"]["review"]
	data.Data["attendID"] = data.ws.Query("attendID")
	templateOut("attend.reject.html", data)
	return
}
func post_attend_reject(data *TemplateData) (err error) {
	attendID, _ := strconv.Atoi(data.ws.Query("attendID"))
	out := protocol.GET_MSG_OA_attend_getById()
	out.Id = int32(attendID)
	var result *protocol.MSG_OA_attend_getbyId_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		data.ajaxResult(false, err.Error())
		return nil
	}
	reviewedBy, _ := strconv.Atoi(data.Config["attend"]["custom"]["reviewedBy"].(string))
	if !(data.User.IsAdmin || reviewedBy == int(data.User.Id)) {
		if managedDepts, err := dept_getDeptManagedByMe(data.User.Id); err != nil {
			return err
		} else {
			find := false
			user := HostConn.GetUserCacheById(result.Info.Uid)
			if user == nil {
				data.ajaxResult(false, data.Lang["attend"]["error"].(map[string]string)["reviewNotAllow"])
				return nil
			}
			for _, deptid := range managedDepts {
				if user.Dept == deptid {
					find = true
					break
				}
			}
			if !find {
				data.ajaxResult(false, data.Lang["attend"]["error"].(map[string]string)["reviewNotAllow"])
				return nil
			}
		}
	}
	update := protocol.GET_MSG_OA_attend_update()
	update.Uid = data.User.Id
	update.ManualIn, _ = time.ParseInLocation("2006-01-02 15:04:05", result.Info.Date.Format("2006-01-02 ")+result.Info.ManualIn, time.Local)
	update.ManualOut, _ = time.ParseInLocation("2006-01-02 15:04:05", result.Info.Date.Format("2006-01-02 ")+result.Info.ManualOut, time.Local)
	update.Desc = result.Info.Desc
	update.Date = result.Info.Date
	update.ReviewStatus = "reject"
	update.RejectDesc = data.ws.Post("comment")
	update.Reason = result.Info.Reason
	update.ReviewedBy = data.User.Id
	libraries.DebugLog("%+v", update)
	if err := data.SendMsgWaitResultToDefault(update, nil); err != nil {
		data.ajaxResult(false, err.Error())
		return nil
	}
	data.ajaxResult(true, data.Lang["common"]["saveSuccess"], createLink("attend", "browsereview", nil))
	return
}
func get_attend_exportDetail(data *TemplateData) (err error) {
	begin, err := time.ParseInLocation("2006-01-02", data.ws.Query("begin"), time.Local)
	if err != nil {
		return errors.New(data.Lang["attend"]["error"].(map[string]string)["beginDate"])
	}
	end, err := time.ParseInLocation("2006-01-02", data.ws.Query("end"), time.Local)
	if err != nil {
		return errors.New(data.Lang["attend"]["error"].(map[string]string)["endDate"])
	}

	deptID := data.ws.Session().Load_int32("attendDeptID")
	userID := data.ws.Session().Load_int32("attendUserID")

	fileName := ""
	if deptID > 0 {

		if dept, _ := HostConn.GetdeptCacheById(deptID); dept != nil {
			fileName = dept.Name + " - "
		}
	}
	if userID > 0 {

		if user := HostConn.GetUserCacheById(userID); user != nil {
			fileName += user.Realname + " - "
		}
	}
	fileName += begin.Format("2006-01-02") + data.Lang["attend"]["to"].(string) + end.Format("2006-01-02") + data.Lang["attend"]["detail"].(string)
	data.Data["fileName"] = fileName
	templateOut("attend.export.html", data)
	return nil
}
func post_attend_exportDetail(data *TemplateData) (err error) {
	begin, err := time.ParseInLocation("2006-01-02", data.ws.Query("begin"), time.Local)
	if err != nil {
		return errors.New(data.Lang["attend"]["error"].(map[string]string)["beginDate"])
	}
	end, err := time.ParseInLocation("2006-01-02", data.ws.Query("end"), time.Local)
	if err != nil {
		return errors.New(data.Lang["attend"]["error"].(map[string]string)["endDate"])
	}
	var fields []protocol.HtmlKeyValueStr
	for _, field := range strings.Split(data.Config["attend"]["list"]["exportFields"].(string), ",") {
		field = strings.Trim(field, " ")
		name := ""
		switch field {
		case "dept":
			name = data.Lang["user"]["dept"].(string)
		case "realname":
			name = data.Lang["user"]["realname"].(string)
		default:
			if v, ok := data.Lang["attend"][field]; ok {
				name, _ = v.(string)
			}
		}
		fields = append(fields, protocol.HtmlKeyValueStr{field, name})
	}
	deptID := data.ws.Session().Load_int32("attendDeptID")
	userID := data.ws.Session().Load_int32("attendUserID")
	var userIds []int32

	if alluser, err := user_getAllcache(data); err != nil {
		return err
	} else if userID == 0 {
		protocol.Order_user(alluser, func(a, b *protocol.MSG_USER_INFO_cache) bool {
			if a.Dept == b.Dept {
				nameA := a.Realname
				nameB := b.Realname
				if nameA == "" {
					nameA = a.Account
				}
				if nameB == "" {
					nameB = b.Account
				}
				return protocol.Order_Pinyin(nameA, nameB, true)
			}
			return a.Dept < b.Dept
		})

	f:
		for _, user := range alluser {
			for _, id := range strings.Split(data.Config["attend"]["custom"]["noAttendUsers"].(string), ",") {
				if id == strconv.Itoa(protocol.DEPTManager) {
					if dept, err := dept_getAll(data); err != nil {
						return err
					} else {
						for _, d := range dept {
							if d.Manager == user.Id {
								continue f
							}
						}
					}
				}
				if id == strconv.Itoa(int(user.Id)) {
					continue f
				}
			}
			if (deptID == 0 || deptID == user.Dept) && user.Dept != 0 {
				name := user.Realname
				if name == "" {
					name = user.Account
				}
				userIds = append(userIds, user.Id)
			}
		}

	} else {
		userIds = []int32{userID}
	}
	out := protocol.GET_MSG_OA_attend_detail()
	out.BeginDate = begin
	out.EndDate = end
	out.User = userIds
	var result *protocol.MSG_OA_attend_detail_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	var values []map[string]string
	for _, v := range result.List {
		r := reflect.ValueOf(v).Elem()
		tmp := make(map[string]string)
		for _, field := range fields {
			key := strings.ToUpper(field.Key[:1]) + field.Key[1:]
			if v := r.FieldByName(key); v.Kind() != reflect.Invalid {
				tmp[field.Key] = libraries.I2S(v.Interface())
			}
		}
		values = append(values, tmp)
	}
	return file_export2xlsx(data, data.ws.Post("fileName"), fields, values)
}
func get_attend_stat(data *TemplateData) (err error) {
	mode := data.ws.Query("mode")
	month, err := time.ParseInLocation("2006-01-02", data.ws.Query("month")+"-01-01", time.Local)
	if err != nil {
		month, err = time.ParseInLocation("2006-01-02", data.ws.Query("month")+"-01", time.Local)
		if err != nil {
			month = time.Now()
		}
	}

	if data.Data["users"], err = user_getPairs(data, "noclosed,noempty,nodeleted,noforbidden"); err != nil {
		return err
	}
	out := protocol.GET_MSG_OA_attend_getStat()
	out.Month = month

	var result *protocol.MSG_OA_attend_getStat_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	if len(result.List) > 0 {
		if mode == "" {
			mode = "view"
		}
	} else {
		mode = "edit"

	}
	getAllmonth := protocol.GET_MSG_OA_attend_getAllMonth()
	getAllmonth.Uids = []int32{data.User.Id}
	var getAllmonthResult *protocol.MSG_OA_attend_getAllMonth_result
	if err = data.SendMsgWaitResultToDefault(getAllmonth, &getAllmonthResult); err != nil {
		return
	}
	data.Data["yearList"] = getAllmonthResult.List
	data.Data["title"] = data.Lang["attend"]["stat"].(string)
	checkWaitReviews := protocol.GET_MSG_OA_attend_checkWaitReviews()
	checkWaitReviews.Month = month
	var checkResult *protocol.MSG_OA_attend_checkWaitReviews_result
	if err = data.SendMsgWaitResultToDefault(checkWaitReviews, &checkResult); err != nil {
		return
	}
	data.Data["waitReviews"] = checkResult.WaitReviews
	data.Data["mode"] = mode
	data.Data["stat"] = result.List
	data.Data["month"] = month
	data.Data["currentYear"] = month.Format("2006")
	data.Data["currentMonth"] = month.Format("01")
	templateOut("attend.export.html", data)
	out.Put()
	result.Put()
	getAllmonth.Put()
	getAllmonthResult.Put()
	checkWaitReviews.Put()
	checkResult.Put()
	return
}
func get_attend_zkteco(data *TemplateData) (err error) {
	data.Data["title"] = data.Lang["attend"]["zkteco"]

	templateOut("attend.zkteco.html", data)
	return
}
