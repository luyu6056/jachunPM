package handler

import (
	"config"
	"errors"
	"fmt"
	"html/template"
	"jachunPM_http/js"
	"libraries"
	"math"
	"mysql"
	"protocol"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func init() {

	httpHandlerModuleInit["GET"]["task"] = task_commonAction
	httpHandlerModuleInit["POST"]["task"] = task_commonAction
	httpHandlerMap["GET"]["/task/create"] = get_task_create
	httpHandlerMap["POST"]["/task/create"] = post_task_create
	httpHandlerMap["GET"]["/task/assignTo"] = get_task_assignTo
	httpHandlerMap["POST"]["/task/assignTo"] = post_task_assignTo
	httpHandlerMap["GET"]["/task/view"] = get_task_view
	httpHandlerMap["GET"]["/task/edit"] = get_task_edit
	httpHandlerMap["POST"]["/task/edit"] = post_task_edit
	httpHandlerMap["GET"]["/task/recordEstimate"] = get_task_recordEstimate
	httpHandlerMap["POST"]["/task/recordEstimate"] = post_task_recordEstimate
	httpHandlerMap["GET"]["/task/start"] = get_task_start
	httpHandlerMap["POST"]["/task/start"] = post_task_start
	httpHandlerMap["GET"]["/task/finish"] = get_task_finish
	httpHandlerMap["POST"]["/task/finish"] = post_task_finish
	httpHandlerMap["GET"]["/task/activate"] = get_task_activate
	httpHandlerMap["POST"]["/task/activate"] = post_task_activate
	httpHandlerMap["GET"]["/task/pause"] = get_task_pause
	httpHandlerMap["POST"]["/task/pause"] = post_task_pause
	httpHandlerMap["GET"]["/task/restart"] = get_task_start
	httpHandlerMap["POST"]["/task/restart"] = post_task_start
	httpHandlerMap["GET"]["/task/internalaudit"] = get_task_internalaudit
	httpHandlerMap["POST"]["/task/internalaudit"] = post_task_internalaudit
	httpHandlerMap["GET"]["/task/proofreading"] = get_task_proofreading
	httpHandlerMap["POST"]["/task/proofreading"] = post_task_proofreading
	httpHandlerMap["GET"]["/task/close"] = get_task_close
	httpHandlerMap["POST"]["/task/close"] = post_task_close
	httpHandlerMap["GET"]["/task/batchCreate"] = get_task_batchCreate
	httpHandlerMap["POST"]["/task/batchCreate"] = post_task_batchCreate
	httpHandlerMap["GET"]["/task/examine"] = get_task_examine
	httpHandlerMap["POST"]["/task/examine"] = post_task_examine
	httpHandlerMap["GET"]["/task/cancel"] = get_task_cancel
	httpHandlerMap["POST"]["/task/cancel"] = post_task_cancel
	httpHandlerMap["POST"]["/task/batchCancel"] = post_task_batchCancel
	httpHandlerMap["POST"]["/task/batchClose"] = post_task_batchClose
	httpHandlerMap["GET"]["/task/delete"] = get_task_delete
	httpHandlerMap["POST"]["/task/batchEdit"] = post_task_batchEdit
	httpHandlerMap["GET"]["/task/batchEdit"] = post_task_batchEdit
	httpHandlerMap["POST"]["/task/batchexamine"] = post_task_batchexamine
	httpHandlerMap["POST"]["/task/batchexaminec"] = post_task_batchexaminec
	httpHandlerMap["POST"]["/task/batchproofreading"] = post_task_batchproofreading
	httpHandlerMap["POST"]["/task/batchproofreadingc"] = post_task_batchproofreadingc
	httpHandlerMap["POST"]["/task/finishall"] = post_task_finishall
	httpHandlerMap["POST"]["/task/exportfinish"] = post_task_exportfinish
	httpHandlerMap["POST"]["/task/placeOrder"] = post_task_placeOrder
	httpHandlerMap["POST"]["/task/batchAssignTo"] = post_task_batchAssignTo
	httpHandlerMap["GET"]["/task/export"] = get_task_export
	httpHandlerMap["POST"]["/task/export"] = post_task_export
}
func taskTemplateFuncs() {
	global_Funcs["MSG_PROJECT_TASK_isClickable"] = func(data *TemplateData, obj interface{}, action string) bool {
		if task, ok := obj.(*protocol.MSG_PROJECT_TASK); ok {
			switch action {

			case "start":
				if len(task.Children) > 0 {
					return false
				}
				return task.Status == "wait"
			case "recordEstimate":
				if len(task.Children) > 0 {
					return false
				}
			case "finish":
				if len(task.Children) > 0 {
					return false
				}
				return task.Status != "done" && task.Status != "closed" && task.Status != "cancel"
			case "cancel":
				if len(task.Children) > 0 {
					return false
				}
				return task.Status != "done" && task.Status != "closed" && task.Status != "cancel"
			case "pause":
				if len(task.Children) > 0 {
					return false
				}
				return task.Status == "doing"
			case "internalaudit":
				if len(task.Children) > 0 {
					return false
				}
				return task.Status == "doing"
			case "proofreading":
				if len(task.Children) > 0 {
					return false
				}
				return task.Status == "done" && task.Finalfile
			case "activate":
				if len(task.Children) > 0 {
					return false
				}
				return task.Status == "done" || task.Status == "closed" || task.Status == "cancel"
			case "assignto":
				if len(task.Children) > 0 {
					return false
				}
				return task.Status != "closed" && task.Status != "cancel"
			case "close":
				if len(task.Children) > 0 {
					return false
				}
				return task.Status == "done" || task.Status == "cancel"
			case "batchcreate":

				if task.Ancestor > 0 {
					return false
				}

			case "examine":
				return task.Status != "wait"
			case "restart":
				return task.Status == "pause"
			}
		} else {
			libraries.DebugLog("MSG_PROJECT_project_cache_isClickable传入的值类型%v不对", reflect.TypeOf(obj).Elem().Name())
		}
		return true
	}
}
func task_commonAction(data *TemplateData) (err error) {

	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	switch data.App["methodName"] {
	case "create", "batchCreate", "batchCancel", "batchClose", "batchEdit", "batchexamine", "batchexaminec", "batchproofreading", "batchproofreadingc", "finishall", "exportfinish", "placeOrder", "batchAssignTo", "export":
		projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
		project := data.getCacheProjectById(int32(projectID))
		if project == nil {
			return errors.New(data.Lang["project"]["error"].(map[string]string)["NotFound"])
		}
		if !data.User.IsAdmin && !data.User.AclProjects[project.Id] {
			return errors.New(data.Lang["project"]["accessDenied"].(string))
		}
		data.Data["project"] = project
	default:

		task, err := task_getByID(data, int32(taskID))
		if err != nil {
			return err
		}
		if task == nil {
			return errors.New(data.Lang["task"]["error"].(map[string]string)["notFoundTask"])
		}
		project := data.getCacheProjectById(task.Project)
		if !data.User.IsAdmin && !data.User.AclProjects[project.Id] {
			return errors.New(data.Lang["project"]["accessDenied"].(string))
		}
		data.Data["project"] = project
		data.Data["task"] = task
	}

	projectID := data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Id

	//校验数字
	if data.ws.Method() == "POST" {
		var consumedRequired, leftRequired bool
		if requiredFields, ok := data.Config["task"][data.App["methodName"].(string)]["requiredFields"].(string); ok {
			for _, str := range strings.Split(requiredFields, ",") {
				if str == "consumed" {
					consumedRequired = true
				}
				if str == "left" {
					leftRequired = true
				}
			}
		}

		if consumed, err := strconv.ParseFloat(data.ws.Post("consumed"), 64); (consumedRequired && err != nil) || consumed < 0 {
			return errors.New(data.Lang["task"]["error"].(map[string]string)["consumedNumber"])
		}

		if left, err := strconv.ParseFloat(data.ws.Post("left"), 64); (leftRequired && err != nil) || left < 0 {
			return errors.New(data.Lang["task"]["error"].(map[string]string)["leftNumber"])
		}
	}

	projects, err := project_getPairs(data, "nocode")
	if err != nil {
		return
	}
	data.Data["projects"] = projects
	if err = project_setMenu(data, projectID, 0, ""); err != nil {
		return
	}

	if taskID > 0 {
		if data.Data["actions"], err = action_getList(data, "task", int32(taskID)); err != nil {
			return
		}
	}

	if data.Data["members"], err = project_getTeamMemberPairs(data, projectID, "nodeleted"); err != nil {
		return
	}
	return nil
}
func task_printCell(data *TemplateData, col *config.ConfigDatatable, task *protocol.MSG_PROJECT_TASK, users []protocol.HtmlKeyValueStr, browseType string, branchGroups map[int32][]protocol.HtmlKeyValueStr, modulePairs []protocol.HtmlKeyValueStr, mode string, child bool, end int) string {
	buf := bufpool.Get().(*libraries.MsgBuffer)
	defer func() {
		buf.Reset()
		bufpool.Put(buf)
	}()
	canBatchEdit := hasPriv(data, "task", "batchEdit", task)
	canBatchClose := (hasPriv(data, "task", "batchClose", task) && strings.ToLower(browseType) != "closedBy")
	canBatchCancel := hasPriv(data, "task", "batchCancel", task)
	canBatchChangeModule := hasPriv(data, "task", "batchChangeModule", task)
	canBatchAssignTo := hasPriv(data, "task", "batchAssignTo", task)

	canBatchAction := canBatchEdit || canBatchClose || canBatchCancel || canBatchChangeModule || canBatchAssignTo
	storyChanged := (task.StoryStatus == "active" && task.LatestStoryVersion > task.StoryVersion)
	id := col.Id
	if col.Show {
		buf.WriteString("<td class='c-")
		buf.WriteString(col.Id)
		switch id {
		case "status":
			buf.WriteString(" task-")
			buf.WriteString(task.Status)
			buf.WriteString("'")
		case "id":
			buf.WriteString(" cell-id'")

		case "deadline":
			if task.Delay > 0 {
				buf.WriteString(" text-center delayed'")
			} else {
				buf.WriteString("'")
			}
		case "assignedTo":
			buf.WriteString(" has-btn text-left'")
		case "progress":
			buf.WriteString(" text-right'")
		case "placeOrder":
			buf.WriteString(" text-center'")
		case "name":
			buf.WriteString(" text-left")
			if len(task.Children) > 0 {
				buf.WriteString(" has-child")
			}
			buf.WriteString("' title='")
			buf.WriteString(task.Name)
			buf.WriteString("'")

		case "story":

			buf.WriteString("' title='")
			buf.WriteString(task.StoryTitle)
			buf.WriteString("'")
		default:
			buf.WriteString("'")
		}
		buf.WriteString(">")
		//echo "<td class='" . class . "'" . title . ">";
		switch id {
		case "id":
			if canBatchAction {

				buf.WriteString(html_checkbox("taskIDList", []protocol.HtmlKeyValueStr{{strconv.Itoa(int(task.Id)), ""}}, "", "", "block"))
				buf.WriteString(html_a(createLink("task", "view", "taskID="+strconv.Itoa(int(task.Id))), fmt.Sprintf("%03d", task.Id)))
			} else {
				buf.WriteString(fmt.Sprintf("%03d", task.Id))
			}
		case "pri":
			buf.WriteString("<span class='label-pri label-pri-")
			buf.WriteString(strconv.Itoa(int(task.Pri)))
			buf.WriteString("' title='")
			if s, ok := common_getValue(data.Lang["task"]["priList"], task.Pri).(string); ok {
				buf.WriteString(s)
			}
			buf.WriteString("'>")
			if s, ok := common_getValue(data.Lang["task"]["priList"], task.Pri).(string); ok {
				buf.WriteString(s)
			}
			buf.WriteString("</span>")
		case "project":
			if project := data.getCacheProjectById(task.Project); project != nil {
				buf.WriteString(html_a(createLink("project", "task", "projectID="+strconv.Itoa(int(task.Project))), project.Name))
			}

		case "name":
			if v, ok := branchGroups[task.Product]; ok {
				for _, kv := range v {
					if kv.Key == strconv.Itoa(int(task.Branch)) {
						buf.WriteString("<span class='label label-info label-outline'>")
						buf.WriteString(kv.Value)
						buf.WriteString("</span>")
						break
					}
				}

			}
			if len(task.Children) == 0 && task.Module > 0 {
				for _, kv := range modulePairs {

					if kv.Key == strconv.Itoa(int(task.Module)) {
						buf.WriteString("<span class='label label-gray label-badge'>")
						buf.WriteString(kv.Value)
						buf.WriteString("</span>")
						break
					}
				}
			}
			if task.Ancestor > 0 {

				switch end {
				case 1:
					buf.WriteString("<span style='margin-left:28px;font-size:30px;position:absolute;'>├</span> ")
				case 2:
					buf.WriteString("<span style='margin-left:3px;font-size:30px;position:absolute;'>│</span><span style='margin-left:28px;font-size:30px;position:absolute;'>└</span> ")
				case 3:
					buf.WriteString("<span style='margin-left:28px;font-size:30px;position:absolute;'>└</span> ")
				default:
					buf.WriteString("<span style='margin-left:3px;font-size:30px;position:absolute;'>│</span><span style='margin-left:28px;font-size:30px;position:absolute;'>├</span> ")
				}
				buf.WriteString("<span  style='margin-left:54px'> </span> ")
			} else if task.Parent > 0 {
				if end > 0 {
					buf.WriteString("<span style='font-size:30px;position:absolute;'>└</span> ")
				} else {
					buf.WriteString("<span style='font-size:30px;position:absolute;'>├</span> ")
				}
				buf.WriteString("<span  style='margin-left:24px'> </span> ")
			}
			if len(task.Team) > 0 {
				buf.WriteString("<span class='label label-badge label-light'>")
				buf.WriteString(data.Lang["task"]["multipleAB"].(string))
				buf.WriteString("</span> ")
			}
			if hasPriv(data, "task", "view") {
				buf.WriteString(html_a(createLink("task", "view", "taskID="+strconv.Itoa(int(task.Id))), task.Name, "", "style='color: task.Color'"))
			} else {
				buf.WriteString("<span style='color: task.Color'>")
				buf.WriteString(task.Name)
				buf.WriteString("</span>")
			}
			if len(task.Children) > 0 {
				buf.WriteString(`<a class="task-toggle" data-id="`)
				buf.WriteString(strconv.Itoa(int(task.Id)))
				buf.WriteString(`"><i class="icon icon-angle-double-right"></i></a>`)
			}
			if len(task.Grandchildren) > 0 {
				buf.WriteString(`<a class="task-toggle" data-id="`)
				buf.WriteString(strconv.Itoa(int(task.Id)))
				buf.WriteString(`"><i class="icon icon-angle-double-right"></i></a>`)
			}
			if task.FromBug > 0 {
				buf.WriteString(html_a(createLink("bug", "view", "id="+strconv.Itoa(int(task.FromBug))), "[BUG#"+strconv.Itoa(int(task.FromBug))+"]", "_blank", "class='bug'"))
			}
		case "type":
			for _, kv := range data.Lang["task"]["typeList"].([]*protocol.HtmlKeyValueStr) {
				if kv.Key == task.Type {
					buf.WriteString(kv.Value)
					break
				}

			}

		case "status":
			if storyChanged {
				buf.WriteString("<span class='status-story status-changed'>")
				buf.WriteString(data.Lang["story"]["changed"].(string))
				buf.WriteString("</span>")
			} else {
				buf.WriteString("<span class='status-task status-")
				buf.WriteString(task.Status)
				buf.WriteString("'> ")
				for _, kv := range data.Lang["task"]["statusList"].([]protocol.HtmlKeyValueStr) {
					if kv.Key == task.Status {
						buf.WriteString(kv.Value)
						break
					}

				}

				buf.WriteString("</span>")
			}
		case "estimate":
			buf.WriteString(strconv.Itoa(int(math.Round(task.Estimate*10) / 10)))
		case "consumed":
			buf.WriteString(strconv.Itoa(int(math.Round(task.Consumed*10) / 10)))

		case "left":
			buf.WriteString(strconv.Itoa(int(math.Round(task.Left*10) / 10)))

		case "progress":
			buf.WriteString(strconv.Itoa(int(task.Progress)))
			buf.WriteString(`%`)
		case "deadline":
			if task.Deadline.After(protocol.NORMALTIME) {
				buf.WriteString(task.Deadline.Format("01-02"))
			}

		case "openedBy":
			for _, kv := range users {
				if kv.Key == strconv.Itoa(int(task.OpenedBy)) {
					buf.WriteString(kv.Value)
					break
				}
			}

		case "openedDate":
			buf.WriteString(task.Deadline.Format("01-02 15:04"))
		case "estStarted":
			buf.WriteString(task.EstStarted.Format("01-02 15:04"))
		case "realStarted":
			buf.WriteString(task.RealStarted.Format("01-02 15:04"))
		case "assignedTo":
			buf.WriteString(tack_printAssignedHtml(data, task, users))
		case "assignedDate":
			buf.WriteString(task.AssignedDate.Format("01-02 15:04"))

		case "finishedBy":
			for _, kv := range users {
				if kv.Key == strconv.Itoa(int(task.FinishedBy)) {
					buf.WriteString(kv.Value)
					break
				}
			}

		case "finishedDate":
			buf.WriteString(task.FinishedDate.Format("01-02 15:04"))

		case "canceledBy":
			for _, kv := range users {
				if kv.Key == strconv.Itoa(int(task.CanceledBy)) {
					buf.WriteString(kv.Value)
					break
				}
			}

		case "canceledDate":
			buf.WriteString(task.CanceledDate.Format("01-02 15:04"))

		case "closedBy":
			for _, kv := range users {
				if kv.Key == strconv.Itoa(int(task.ClosedBy)) {
					buf.WriteString(kv.Value)
					break
				}
			}

		case "closedDate":
			buf.WriteString(task.ClosedDate.Format("01-02 15:04"))

		case "closedReason":
			for _, kv := range data.Lang["task"]["reasonList"].([]*protocol.HtmlKeyValueStr) {
				if kv.Key == task.ClosedReason {
					buf.WriteString(kv.Value)
					break
				}

			}
		case "finalfile":
			if task.Finalfile {
				buf.WriteString("<span class='status-story status-done'>已传</span>")
			} else {
				buf.WriteString("未传")
			}
		case "proofreading":
			if task.Proofreading {
				buf.WriteString("<span class='status-story status-done'>已对</span>")
			} else {
				buf.WriteString("未对")
			}

		case "examine":
			if task.Examine {
				buf.WriteString("<span class='status-story status-done'>通过</span>")
			} else {
				buf.WriteString("未通过")
			}

		case "placeOrder":
			if task.PlaceOrder {
				buf.WriteString("<span class='status-story text-center status-done'>是</span>")
			} else {
				buf.WriteString("否")
			}

		case "story":
			if task.StoryID > 0 {
				if hasPriv(data, "story", "view") {
					buf.WriteString(html_a(createLink("story", "view", []interface{}{"storyid=", task.StoryID, true}), "<i class='icon icon-"+data.Lang["icons"]["story"].(string)+"'></i>", "", "class='iframe' title='"+task.StoryTitle+"'"))
				} else {
					buf.WriteString("<i class='icon icon-")
					buf.WriteString(data.Lang["icons"]["story"].(string))
					buf.WriteString("' title='")
					buf.WriteString(task.StoryTitle)
					buf.WriteString("'></i>")
				}
			}
		case "mailto":

			for _, id := range task.Mailto {
				for _, kv := range users {
					if kv.Key == strconv.Itoa(int(id)) {
						buf.WriteString(kv.Value)
						buf.WriteString(" &nbsp;")
						break
					}
				}

			}
		case "lastEditedBy":
			for _, kv := range users {
				if kv.Key == strconv.Itoa(int(task.LastEditedBy)) {
					buf.WriteString(kv.Value)
					break
				}
			}

		case "lastEditedDate":
			buf.WriteString(task.LastEditedDate.Format("01-02 15:04"))

		case "actions":
			if storyChanged {
				buf.WriteString(common_printIcon(data, "task", "confirmStoryChange", "taskID="+strconv.Itoa(int(task.Id)), nil, "list", "", "hiddenwin", "btn-wide"))
				break
			}
			if task.Status != "pause" {
				buf.WriteString(common_printIcon(data, "task", "start", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "", "", "iframe", "true"))
			} else {
				buf.WriteString(common_printIcon(data, "task", "restart", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "", "", "iframe", "true"))
			}

			buf.WriteString(common_printIcon(data, "task", "close", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "", "", "iframe", "true"))
			buf.WriteString(common_printIcon(data, "task", "finish", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "", "", "iframe", "true"))

			buf.WriteString(common_printIcon(data, "task", "recordEstimate", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "time", "", "iframe", "true"))
			buf.WriteString(common_printIcon(data, "task", "edit", "taskID="+strconv.Itoa(int(task.Id)), task, "list", ""))
			buf.WriteString(common_printIcon(data, "task", "examine", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "", "", "iframe", "true"))
			batchCreateDesc := data.Lang["task"]["children"].(string)
			if task.Parent > 0 {
				batchCreateDesc = data.Lang["task"]["grandchildren"].(string)
			}
			buf.WriteString(common_printIcon(data, "task", "batchCreate", "projectID="+strconv.Itoa(int(task.Project))+"&storyID="+strconv.Itoa(int(task.Story))+"&moduleID="+strconv.Itoa(int(task.Module))+"&taskID="+strconv.Itoa(int(task.Id))+"&ifame=0", task, "list", "treemap-alt", "", "", "", "", batchCreateDesc))

		}
		buf.WriteString("</td>\r\n")
	}
	res := buf.String()
	return res
}
func tack_printAssignedHtml(data *TemplateData, task *protocol.MSG_PROJECT_TASK, users []protocol.HtmlKeyValueStr) string {
	btnTextClass := "text-primary"
	assignedToText := data.Lang["task"]["noAssigned"].(string)
	for _, kv := range users {
		if kv.Key == strconv.Itoa(int(task.AssignedTo)) {
			assignedToText = kv.Value
		}
	}

	if task.AssignedTo == data.User.Id {
		btnTextClass = "text-red"
	}
	if !hasPriv(data, "task", "assignTo") {
		return "<span style='padding-left: 21px' class='" + btnTextClass + "'>" + assignedToText + "</span>"
	}
	btnClass := "iframe btn btn-icon-left btn-sm"
	if task.AssignedTo == -1 || (data.Config["task"]["custom"]["allowParentAssignTo"].(string) == "false" && task.AssignedTo == 0 && len(task.Children) > 0) {
		btnClass += " disabled"
	}
	return html_a(createLink("task", "assignTo", []interface{}{"projectID=", task.Project, "&taskID=", task.Id, true}), "<i class='icon icon-hand-right'></i> <span class='"+btnTextClass+"'>"+assignedToText+"</span>", "", "class='"+btnClass+"'")

}
func get_task_assignTo(data *TemplateData) (err error) {
	task := data.Data["task"].(*protocol.MSG_PROJECT_TASK)

	/* Compute next assignedTo. */

	if user := HostConn.GetUserCacheById(task.AssignedTo); user != nil {
		if user.Realname != "" {
			data.Data["AssignedToRealName"] = user.Realname
		} else {
			data.Data["AssignedToRealName"] = user.Account
		}
	} else {
		data.Data["AssignedToRealName"] = "ID:" + strconv.Itoa(int(task.AssignedTo))
	}

	data.Data["title"] = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["assign"].(string)

	data.Data["task"] = task
	if data.Data["users"], err = user_getPairs(data, ""); err != nil {
		return
	}

	templateOut("task.assignto.html", data)
	return nil
}
func post_task_assignTo(data *TemplateData) (err error) {
	assignedTo, _ := strconv.Atoi(data.ws.Post("assignedTo"))
	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))

	if assignedTo != 0 && (assignedTo < 0 || HostConn.GetUserCacheById(int32(assignedTo)) == nil) {
		data.ws.WriteString(js.Error(data.Lang["task"]["error"].(map[string]string)["AssignedToNotFoundUser"]))
		return
	}
	out := protocol.GET_MSG_PROJECT_task_assignTo()
	out.TaskID = int32(taskID)
	out.Left, _ = strconv.ParseFloat(data.ws.Post("left"), 64)
	out.AssignedTo = int32(assignedTo)
	out.Comment = data.ws.Post("comment")
	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		e := err.Error()
		if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
			e = v
		}
		data.ws.WriteString(js.Error(e))
		return dataErrAlreadyOut
	}
	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
	} else {
		data.ws.WriteString(js.Location(createLink("task", "view", "taskID="+strconv.Itoa(taskID))))
	}
	return
}
func task_getByID(data *TemplateData, taskID int32) (*protocol.MSG_PROJECT_TASK, error) {
	out := protocol.GET_MSG_PROJECT_task_getById()
	out.Id = taskID
	var result *protocol.MSG_PROJECT_task_getById_result
	if err := data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return nil, err
	}
	out.Put()
	return result.Info, nil

}

func get_task_view(data *TemplateData) (err error) {

	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	task, err := task_getByID(data, int32(taskID))
	if err != nil {
		return
	}

	if task == nil {
		data.ws.WriteString(js.Error(data.Lang["common"]["notFound"].(string)) + js.Location("back"))
		return
	}
	if task.Parent > 0 {
		parent, err := task_getByID(data, int32(task.Parent))
		if parent != nil {
			data.Data["parentName"] = html_a(createLink("task", "view", []interface{}{"taskID=", parent.Id}), parent.Name)
		}
		if err != nil {
			libraries.ReleaseLog("无法获取parent task id%d", task.Parent)
		}
	}
	if task.Ancestor > 0 {
		ancestor, err := task_getByID(data, int32(task.Ancestor))
		if ancestor != nil {
			data.Data["parentName"] = html_a(createLink("task", "view", []interface{}{"taskID=", ancestor.Id}), ancestor.Name) + "/" + data.Data["parentName"].(string)
		}
		if err != nil {
			libraries.ReleaseLog("无法获取parent task id%d", task.Parent)
		}
	}
	if task.FromBug > 0 {
		getbug := protocol.GET_MSG_TEST_bug_getById()
		getbug.Id = task.FromBug
		var result *protocol.MSG_TEST_bug_getById_result
		err := data.SendMsgWaitResultToDefault(getbug, &result)
		getbug.Put()
		if err != nil {
			return err
		}
		/*task->bugSteps = ''
		  if(bug)
		  {
		      task->bugSteps = this->loadModel('file')->setImgSize(bug->steps);
		      foreach(bug->files as file) task->files[] = file;
		  }*/
		data.Data["fromBug"] = result.Info
	} else if task.Story > 0 {
		getStory := protocol.GET_MSG_PROJECT_story_getById()
		getStory.Id = task.Story
		var result *protocol.MSG_PROJECT_story_getById_result
		err := data.SendMsgWaitResultToDefault(getStory, &result)
		getStory.Put()
		if err != nil {
			return err
		}
		/*task->storySpec     = empty(story) ? '' : this->loadModel('file')->setImgSize(story->spec);
		  task->storyVerify   = empty(story) ? '' : this->loadModel('file')->setImgSize(story->verify);*/
		//data.Data["story"]
		if data.Data["storyFiles"], err = file_getByObject(data, "story", result.Story.Id); err != nil {
			return err
		}
		result.Put()
	}
	//if(task->team) this->lang->task->assign = this->lang->task->transfer;

	/* Update action. */
	setAction := protocol.GET_MSG_LOG_Action_set_read()
	setAction.ObjectType = "task"
	setAction.ObjectID = task.Id
	data.SendMsgToDefault(setAction)
	setAction.Put()
	/* Set menu. */

	if err := project_setMenu(data, task.Project, 0, ""); err != nil {
		return err
	}
	project := data.Data["project"].(*protocol.MSG_PROJECT_project_cache)
	//var_dump(task);die;
	data.Data["title"] = fmt.Sprintf("TASK#%d %s / %s", task.Id, task.Name, project.Name)
	data.Data["project"] = project
	data.Data["task"] = task
	if data.Data["users"], err = user_getPairs(data, "noletter"); err != nil {
		return
	}
	//data.Data["preAndNext"]  = this->loadModel('common')->getPreAndNextObject('task', taskID);
	if task.Module > 0 {
		cache, err := tree_getAllcache(data)
		if err != nil {
			return err
		}
		var path []int32
		for _, module := range cache {
			if module.Id == task.Module {
				path = module.Path
				break
			}
		}

		var modulePath []*protocol.MSG_PROJECT_tree_cache
		var productName string
		if len(path) > 0 {
			for _, module := range cache {
				//获取product
				if module.Id == path[0] {
					if module.Type == "story" && module.Root > 0 {
						for _, m := range cache {
							if m.Id == module.Root {
								productName = module.Name
								break
							}
						}

					}
				}
				//获取modulePath
				for _, id := range path {
					if module.Id == id {
						modulePath = append(modulePath, module)
					}
				}
			}
			protocol.Order_tree(modulePath, func(a, b *protocol.MSG_PROJECT_tree_cache) bool {
				return a.Grade < b.Grade
			})
		}

		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf2 := bufpool.Get().(*libraries.MsgBuffer)
		if len(modulePath) == 0 {
			buf.WriteByte('/')
			buf2.WriteByte('/')
		} else {
			if productName != "" {
				buf2.WriteString(productName)
				buf2.WriteByte('/')
				buf.WriteString(productName)
				buf.WriteString(data.Lang["common"]["arrow"].(string))
			}
			for k, module := range modulePath {
				buf2.WriteString(module.Name)
				if link := common_printLink(data, "project", "task", "projectID="+strconv.Itoa(int(task.Project))+"&browseType=byModule&param="+strconv.Itoa(int(module.Id)), module.Name); link == "" {
					buf.WriteString(module.Name)
				} else {
					buf.WriteString(link)
				}
				if k < len(modulePath)-1 {
					buf2.WriteByte('/')
					buf.WriteString(data.Lang["common"]["arrow"].(string))
				}
			}
		}
		data.Data["modulePath"] = template.HTML(buf.String())
		data.Data["moduleTitle"] = buf2.String()
		buf.Reset()
		buf2.Reset()
		bufpool.Put(buf)
		bufpool.Put(buf2)
	} else {
		data.Data["modulePath"] = "/"
	}
	data.Data["checkObject"] = map[string]interface{}{"Project": task.Project}
	if task.Story > 0 {
		out := protocol.GET_MSG_TEST_CASE_getTaskCasePairs()
		out.Story = task.Story
		out.StoryVersion = task.StoryVersion
		var result *protocol.MSG_TEST_CASE_getTaskCasePairs_result
		if err := data.SendMsgWaitResultToDefault(out, &result); err != nil {
			return err
		}
		data.Data["Cases"] = result.List
		out.Put()
	}
	data.Data["actionFormLink"] = createLink("action", "comment", []interface{}{"objectType=task&objectID=", task.Id})
	if files, err := file_getByObject(data, "task", task.Id); err != nil {
		return err
	} else {
		var ProcessFile, FeedbackFile, SourceFile, FinalFile []*protocol.MSG_FILE_getByID_result
		for _, file := range files {
			switch file.Type {
			case "processFile":
				ProcessFile = append(ProcessFile, file)
			case "feedbackFile":
				FeedbackFile = append(FeedbackFile, file)
			case "sourceFile":
				SourceFile = append(SourceFile, file)
			case "FinalFile":
				FinalFile = append(FinalFile, file)
			}
		}
		data.Data["ProcessFile"] = ProcessFile
		data.Data["FeedbackFile"] = FeedbackFile
		data.Data["SourceFile"] = SourceFile
		data.Data["FinalFile"] = FinalFile
	}
	if user := HostConn.GetUserCacheById(task.AssignedTo); user != nil {
		if user.Realname != "" {
			data.Data["AssignedToRealName"] = user.Realname
		} else {
			data.Data["AssignedToRealName"] = user.Account
		}
	}
	if len(task.Team) > 0 {
		out := protocol.GET_MSG_USER_team_getByIds()
		out.Ids = task.Team
		var result *protocol.MSG_USER_team_getByIds_result
		if err := data.SendMsgWaitResultToDefault(out, &result); err != nil {
			return err
		}
		data.Data["taskTeams"] = result.List
		out.Put()
	}
	if data.Data["parentName"] != nil {
		data.Data["taskName"] = data.Data["parentName"].(string) + "/" + task.Name
	} else {
		data.Data["taskName"] = task.Name
	}

	templateOut("task.view.html", data)
	return
}

func get_task_create(data *TemplateData) (err error) {
	//已屏蔽团队限制用户功能，后续看需求再加

	/*task = new stdClass();
	  task->module     = moduleID;
	  task->assignedTo = "";
	  task->name       = "";
	  task->story      = storyID;
	  task->type       = "";
	  task->pri        = "3";
	  task->estimate   = "";
	  task->desc       = "";
	  task->estStarted = "";
	  task->deadline   = "";
	  task->mailto     = "";
	  task->color      = "";
	  if(taskID > 0)
	  {
	      task      = this->task->getByID(taskID);
	      projectID = task->project;
	  }

	  if(todoID > 0)
	  {
	      todo = this->loadModel("todo")->getById(todoID);
	      task->name = todo->name;
	      task->pri  = todo->pri;
	      task->desc = todo->desc;
	  }*/

	moduleID, _ := strconv.Atoi(data.ws.Query("moduleID"))
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	var task *protocol.MSG_PROJECT_TASK
	data.Data["EstStarted"] = ""
	data.Data["Deadline"] = ""
	if taskID > 0 {
		out := protocol.GET_MSG_PROJECT_task_getById()
		out.Id = int32(taskID)
		var result *protocol.MSG_PROJECT_task_getById_result
		if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
			return
		}
		out.Put()

		if result != nil && result.Info != nil {
			task = result.Info
			projectID = int(task.Project)
			if task.EstStarted.After(protocol.NORMALTIME) {
				data.Data["EstStarted"] = task.EstStarted.Format("2006-01-02")
			}
			if task.Deadline.After(protocol.NORMALTIME) {
				data.Data["Deadline"] = task.Deadline.Format("2006-01-02")
			}
		}
	}

	project := data.getCacheProjectById(int32(projectID))
	if project == nil {
		return errors.New(data.Lang["project"]["error"].(map[string]string)["NotFound"])
	}

	if err = project_setMenu(data, int32(projectID), 0, ""); err != nil {
		return
	}

	if data.Data["users"], err = user_getPairs(data, "noclosed|nodeleted"); err != nil {
		return
	}
	moduleIdList, err := dept_getAllChildID(int32(moduleID))
	if err != nil {

		return
	}
	if data.Data["stories"], err = story_getProjectStoryPairs(data, int32(projectID), 0, 0, moduleIdList, ""); err != nil {
		return
	}

	if data.Data["moduleOptionMenu"], err = tree_getTaskOptionMenu(data, int32(projectID), 0, 0); err != nil {
		return
	}

	data.Data["title"] = project.Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["create"].(string)

	/* Set Custom*/
	var customFields []protocol.HtmlKeyValueStr
	for _, field := range data.Config["task"]["common"]["customCreateFields"].([]string) {
		if !(field == "story" && project.Type == "ops") {
			customFields = append(customFields, protocol.HtmlKeyValueStr{field, data.Lang["task"][field].(string)})
		}

		switch field {
		case "story":
			data.Data["hasStoryField"] = true
		case "pri":
			data.Data["hasPriField"] = true
		case "estimate":
			data.Data["hasEstimateField"] = true
		case "estStarted":
			data.Data["hasEstStartedField"] = true
		case "deadline":
			data.Data["hasDeadlineField"] = true
		case "mailto":
			data.Data["hasMailtoField"] = true
		}
	}
	data.Data["showFields"] = data.Config["task"]["custom"]["createFields"]
	hasCustomPri := false
	for _, kv := range data.Lang["task"]["priList"].([]protocol.HtmlKeyValueStr) {
		if kv.Key != "" && kv.Key != kv.Value {
			hasCustomPri = true
			break
		}
	}

	var priList []protocol.HtmlKeyValueStr = make([]protocol.HtmlKeyValueStr, len(data.Lang["task"]["priList"].([]protocol.HtmlKeyValueStr)))
	copy(priList, data.Lang["task"]["priList"].([]protocol.HtmlKeyValueStr))
	data.Data["priList"] = priList
	data.Data["hasCustomPri"] = hasCustomPri

	data.Data["customFields"] = customFields

	data.Data["project"] = project
	if task == nil {
		task = &protocol.MSG_PROJECT_TASK{Pri: 3}
	}
	data.Data["task"] = task
	data.Data["customLink"] = createLink("custom", "ajaxSaveCustomFields", "module=task&section=custom&key=createFields")

	templateOut("task.create.html", data)
	return
}
func post_task_create(data *TemplateData) (err error) {
	if !data.ajaxCheckPost() {
		return
	}
	moduleID, _ := strconv.Atoi(data.ws.Query("moduleID"))
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	//taskID, _ := strconv.Atoi(data.ws.Query("taskID"))

	var task = &protocol.MSG_PROJECT_TASK{
		Project:    int32(projectID),
		Status:     "wait",
		OpenedBy:   data.User.Id,
		OpenedDate: time.Now(),
	}
	if _, ok := task_createEdit(data, task); !ok {
		return
	}
	project := data.Data["project"].(*protocol.MSG_PROJECT_project_cache)
	data.Msg.ActionCreate("task", task.Id, "Opened", "", "", project.Products, project.Id)

	if data.onlybody() {
		data.ajaxResult(true, "", "reload", "parent")
		return dataErrAlreadyOut
	}

	/* Locate the browser. */
	switch data.ws.Post("after") {
	case "continueAdding":
		data.ajaxResult(true, data.Lang["task"]["successSaved"].(string)+common_getValue(data.Lang["task"]["afterChoices"], "continueAdding").(string),
			createLink("task", "create", []interface{}{"projectID=", projectID, "&storyID=", task.Story, "&moduleID=", moduleID}))
		return dataErrAlreadyOut
	case "toTaskList":
		moduleID = int(task.Module)
		data.ajaxResult(true, data.Lang["task"]["successSaved"].(string)+common_getValue(data.Lang["task"]["afterChoices"], "toTaskList").(string),
			createLink("project", "task", []interface{}{"projectID=", projectID, "&browseType=byModule&param=", moduleID}))
		return dataErrAlreadyOut
	case "toStoryList":
		storyLink := data.ws.Session().Load_str("storyList")
		if storyLink == "" {
			storyLink = createLink("project", "story", "projectID="+data.ws.Query("projectID"))
		}
		data.ajaxResult(true, data.Lang["task"]["successSaved"].(string)+common_getValue(data.Lang["task"]["afterChoices"], "toStoryList").(string),
			storyLink)
		return dataErrAlreadyOut
	}
	data.ajaxResult(true, data.Lang["task"]["successSaved"].(string), createLink("project", "browse", "projectID="+data.ws.Query("projectID")+"&tab=task"))
	return dataErrAlreadyOut

}
func task_createEdit(data *TemplateData, task *protocol.MSG_PROJECT_TASK) (protocol.ChangeHistory, bool) {

	for key, value := range data.ws.GetAllPost() {
		switch key {
		case "type":
			task.Type = value[0]
		case "module":
			id, _ := strconv.Atoi(value[0])
			task.Module = int32(id)
		case "color":
			task.Color = value[0]
		case "name":
			task.Name = value[0]
		case "pri":
			pri, _ := strconv.Atoi(value[0])
			task.Pri = int8(pri)
		case "desc":
			task.Desc = value[0]
		case "estStarted":
			if value[0] != "" {
				task.EstStarted, _ = time.ParseInLocation(protocol.TIMEFORMAT_MYSQLDATE, value[0], time.Local)

			}

		case "deadline":
			if value[0] != "" {
				task.Deadline, _ = time.ParseInLocation(protocol.TIMEFORMAT_MYSQLDATE, value[0], time.Local)

			}
		case "mailto":
			for _, v := range value {
				id, _ := strconv.Atoi(v)
				if id > 0 {
					task.Mailto = append(task.Mailto, int32(id))
				}
			}

		case "assignedTo":
			id, _ := strconv.Atoi(value[0])
			if id > 0 {
				task.AssignedTo = int32(id)
			}
		case "consumed":
			newConsumed, _ := strconv.ParseFloat(value[0], 64)
			if task.Consumed > newConsumed {
				data.ajaxResult(false, map[string]string{"consumed": data.Lang["task"]["error"].(map[string]string)["consumedSmall"]}, "")
				return nil, false
			}
			task.Consumed = newConsumed
		case "project":
			newProjectID, _ := strconv.Atoi(value[0])

			task.Project = int32(newProjectID)
		case "parent":
			parent, _ := strconv.Atoi(value[0])
			task.Parent = int32(parent)
		case "realStarted":
			if value[0] != "" {
				task.RealStarted, _ = time.ParseInLocation(protocol.TIMEFORMAT_MYSQLDATE, value[0], time.Local)
			}

		case "estimate":
			task.Estimate, _ = strconv.ParseFloat(value[0], 64)
		case "left":
			task.Left, _ = strconv.ParseFloat(value[0], 64)
		case "finishedBy":
			id, _ := strconv.Atoi(value[0])
			task.FinishedBy = int32(id)
		case "finishedDate":
			if value[0] != "" {
				task.FinishedDate, _ = time.ParseInLocation(protocol.TIMEFORMAT_MYSQLDATE, value[0], time.Local)
			}
		case "canceledBy":
			id, _ := strconv.Atoi(value[0])
			task.CanceledBy = int32(id)
		case "canceledDate":
			if value[0] != "" {
				task.CanceledDate, _ = time.ParseInLocation(protocol.TIMEFORMAT_MYSQLDATE, value[0], time.Local)
			}
		case "closedBy":
			id, _ := strconv.Atoi(value[0])
			task.ClosedBy = int32(id)
		case "closedDate":
			if value[0] != "" {
				task.ClosedDate, _ = time.ParseInLocation(protocol.TIMEFORMAT_MYSQLDATE, value[0], time.Local)
			}
		case "closedReason":
			task.ClosedReason = value[0]
		case "multiple":
			libraries.ReleaseLog("task团队功能未开发")
		case "uploadFileTmpName":
			for k := range value {
				if data.ws.PostSlice("fileType")[k] == "FinalFile" {
					task.Finalfile = true
				}
			}
		}
	}
	if task.EstStarted.After(task.Deadline) {
		data.ajaxResult(false, map[string]string{"estStarted": data.Lang["task"]["error"].(map[string]string)["deadlineGtEstStarted"]}, "")
		return nil, false
	}
	if task.Status != "wait" && len(task.Team) == 0 && task.Left == 0 && task.Status != "cancel" && task.Status != "closed" {
		task.Status = "done"
	} else if task.Status == "wait" || task.Status == "doing" {
		task.FinishedBy = 0
		task.FinishedDate = protocol.ZEROTIME
		task.CanceledBy = 0
		task.CanceledDate = protocol.ZEROTIME
		task.ClosedBy = 0
		task.ClosedDate = protocol.ZEROTIME
		task.ClosedReason = ""
	}
	if task.Status == "done" {
		if task.Consumed < 0 {
			data.ajaxResult(false, map[string]string{"consumed": data.Lang["task"]["error"].(map[string]string)["consumed"]}, "")
			return nil, false
		}
		if task.ClosedReason != "" {
			task.ClosedReason = "done"
		}
		task.CanceledBy = 0
		task.CanceledDate = protocol.ZEROTIME
	}
	if task.Status == "closed" && task.ClosedReason == "" {
		data.ajaxResult(false, map[string]string{"closedReason": data.Lang["task"]["error"].(map[string]string)["closedReasonNotempty"]}, "")
		return nil, false
	} else if task.ClosedReason == "cancel" {
		task.FinishedBy = 0
		task.FinishedDate = protocol.ZEROTIME
	}

	return do_task_create(data, task)
}
func do_task_create(data *TemplateData, task *protocol.MSG_PROJECT_TASK) (protocol.ChangeHistory, bool) {
	projectID := int(task.Project)
	project := data.getCacheProjectById(int32(projectID))
	if project == nil {
		data.ajaxResult(false, data.Lang["project"]["error"].(map[string]string)["NotFound"], "")
		return nil, false
	}
	product := &protocol.MSG_PROJECT_product_cache{}
	if len(project.Products) > 0 {
		product = HostConn.GetProductById(project.Products[0])
		if product == nil {
			data.ajaxResult(false, data.Lang["product"]["error"].(map[string]string)["NotFound"], "")
			return nil, false
		}
	}

	var uploaderr error
	var newimgids []int64
	//oldDesc := task.Desc
	task.Desc, newimgids, uploaderr = file_descProcessImgURLAnd2Bbcode(data, task.Desc)
	if uploaderr != nil {
		data.ajaxResult(false, uploaderr.Error(), "")
		return nil, false
	}
	session, err := data.Msg.BeginTransaction()
	defer func() {
		if err != nil { //以下使用err来判断图片删除
			deleteimg := protocol.GET_MSG_FILE_DeleteByID()
			for _, id := range newimgids {
				deleteimg.FileID = id
				data.Msg.SendMsg(0, deleteimg)
			}
			deleteimg.Put()
			session.Rollback()
		}

	}()
	if err != nil {
		data.ajaxResult(false, err.Error(), "")
		return nil, false
	}

	out := protocol.GET_MSG_PROJECT_task_create()
	out.Task = task
	var result *protocol.MSG_PROJECT_task_create_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		e := err.Error()
		if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
			e = v
		}
		data.ajaxResult(false, e)
		return nil, false
	}
	if result.Err == protocol.Err_TaskIsexist {
		data.ajaxResult(false, fmt.Sprintf(data.Lang["common"]["duplicate"].(string), data.Lang["task"]["common"]), createLink("task", "view", "taskID="+strconv.Itoa(int(result.Id))))
		return nil, false
	}

	//上传文件
	uploadFileTmp := protocol.GET_MSG_FILE_updateTmp()
	for k, name := range data.ws.PostSlice("uploadFileTmpName") {
		upload := protocol.GET_MSG_FILE_upload()
		upload.AddBy = data.User.Id
		upload.Name = file_getRealName(name)
		upload.Type = data.ws.PostSlice("fileType")[k]
		upload.Code = product.Code
		upload.Title = file_getTitleName(name)
		if upload.Code == "" {
			upload.Code = product.Name
		}
		if project.Code == "" {
			upload.Code += "/" + project.Name
		} else {
			upload.Code += "/" + project.Code
		}
		upload.Code = strings.TrimLeft(upload.Code, "/")
		upload.ObjectID = result.Id
		upload.ObjectType = "task"
		uploadFileTmp.Files = append(uploadFileTmp.Files, upload)
	}
	if err = data.SendMsgWaitResultToDefault(uploadFileTmp, nil); err != nil {
		libraries.ReleaseLog("上传失败%v", err)
		data.ajaxResult(false, data.Lang["file"]["error"].(map[string]string)["updateTmp"])
		return nil, false
	}
	task.Id = result.Id
	session.Commit()
	return result.Change, true
}
func get_task_edit(data *TemplateData) (err error) {
	var task = data.Data["task"].(*protocol.MSG_PROJECT_TASK)
	var projectID int32
	projectID = int32(task.Project)

	if task == nil {
		return errors.New(data.Lang["task"]["error"].(map[string]string)["notFoundTask"])
	}
	noclosedProjects, err := project_getPairs(data, "noclosed,nocode")
	if err != nil {
		return
	}
	for k, kv := range noclosedProjects {
		if kv.Key == strconv.Itoa(int(projectID)) {
			kv.Value = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Name
			noclosedProjects[k] = kv
			break
		}
	}
	getParentTaskPairs := protocol.GET_MSG_PROJECT_task_getPairs()
	getParentTaskPairs.Where = map[string]interface{}{
		"Deleted":  false,
		"Ancestor": []interface{}{"le", 0},
		"Status":   mysql.WhereOperatorNOTIN([]string{"cancel", "closed"}),
		"Project":  projectID,
	}

	var getParentTaskPairsResult *protocol.MSG_PROJECT_task_getPairs_result
	if err = data.SendMsgWaitResultToDefault(getParentTaskPairs, &getParentTaskPairsResult); err != nil {
		return
	}
	protocol.Order_htmlkvStr(getParentTaskPairsResult.List, func(a, b protocol.HtmlKeyValueStr) bool {
		return a.Key > b.Key
	})
	for k, kv := range getParentTaskPairsResult.List {
		if kv.Key == strconv.Itoa(int(task.Id)) {
			getParentTaskPairsResult.List = append(getParentTaskPairsResult.List[:k], getParentTaskPairsResult.List[k+1:]...)
			break
		}
	}

	if task.Parent == -1 {
		data.Data["tasks"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"-1", "无"}}, getParentTaskPairsResult.List...)
	} else {
		data.Data["tasks"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"0", "无"}}, getParentTaskPairsResult.List...)
	}

	if task.AssignedTo > 0 {
		assignedUser := HostConn.GetUserCacheById(task.AssignedTo)
		var foundAssignedTo bool
		for _, kv := range data.Data["members"].([]protocol.HtmlKeyValueStr) {
			if kv.Key == strconv.Itoa(int(task.AssignedTo)) {
				foundAssignedTo = true
				break
			}
		}
		if !foundAssignedTo && assignedUser != nil {
			name := assignedUser.Realname
			if name == "" {
				name = assignedUser.Account
			}
			data.Data["members"] = append(data.Data["members"].([]protocol.HtmlKeyValueStr), protocol.HtmlKeyValueStr{strconv.Itoa(int(assignedUser.Id)), name})
		}
	}
	data.Data["task"] = task
	data.Data["title"] = data.Lang["task"]["edit"].(string) + "TASK" + data.Lang["common"]["colon"].(string) + task.Name
	if data.Data["stories"], err = story_getProjectStoryPairs(data, projectID, 0, 0, nil, ""); err != nil {
		return
	}
	if data.Data["users"], err = user_getPairs(data, "nodeleted", task.OpenedBy, task.CanceledBy, task.ClosedBy); err != nil {
		return
	}
	if data.Data["modules"], err = tree_getTaskOptionMenu(data, projectID, 0, 0); err != nil {
		return
	}
	if data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Acl == "private" {
		data.Data["MailtoUser"] = data.Data["members"].([]protocol.HtmlKeyValueStr)
	} else {
		data.Data["MailtoUser"] = data.Data["users"].([]protocol.HtmlKeyValueStr)
	}
	templateOut("task.edit.html", data)
	return
}
func get_task_recordEstimate(data *TemplateData) (err error) {
	var task = data.Data["task"].(*protocol.MSG_PROJECT_TASK)
	data.ws.Session().Set("estimateList", data.ws.URI())
	out := protocol.GET_MSG_PROJECT_task_GetTaskEstimateByTaskId()
	out.TaskId = task.Id
	var result *protocol.MSG_PROJECT_task_GetTaskEstimateByTaskId_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}

	data.Data["estimates"] = result.List
	data.Data["title"] = data.Lang["task"]["record"]
	if len(task.Team) > 0 && task.Team[len(task.Team)-1] != task.AssignedTo {
		data.Data["confirmRecord"] = data.Lang["task"]["confirmTransfer"]
	} else {
		data.Data["confirmRecord"] = data.Lang["task"]["confirmRecord"]
	}
	templateOut("task.recordEstimate.html", data)
	return
}
func post_task_recordEstimate(data *TemplateData) (err error) {
	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	session, err := data.BeginTransaction()
	if err != nil {
		data.ws.WriteString(js.Error(err.Error()))
		return
	}
	defer session.Rollback()
	update := protocol.GET_MSG_PROJECT_task_UpdateTaskEstimate()
	update.TaskId = int32(taskID)
	for i := 1; i < 3; i++ {
		key := strconv.Itoa(i)
		date := data.ws.Post("dates[" + key + "]")
		if date == "" {
			continue
		}
		estimate := protocol.GET_MSG_PROJECT_TaskEstimate()
		estimate.Task = int32(taskID)
		estimate.Uid = data.User.Id
		estimate.Account = data.User.Account
		estimate.Date, err = time.ParseInLocation("2006-01-02", date, time.Local)
		left, err := strconv.ParseFloat(data.ws.Post("left["+key+"]"), 64)
		if err != nil {
			data.ws.WriteString(js.Alert(data.Lang["task"]["error"].(map[string]string)["leftNumber"]))
			return err
		}
		consumed, err := strconv.ParseFloat(data.ws.Post("consumed["+key+"]"), 64)
		if err != nil {
			data.ws.WriteString(js.Alert(data.Lang["task"]["error"].(map[string]string)["consumedNumber"]))
			return err
		}
		estimate.Left = left
		estimate.Consumed = consumed
		estimate.Work = data.ws.Post("work[" + key + "]")
		if err != nil {
			data.ws.WriteString(js.Error(data.Lang["task"]["error"].(map[string]string)["dateError"]))
			return err
		}
		update.List = append(update.List, estimate)
	}
	if len(update.List) == 0 {
		if data.onlybody() {
			data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
		} else {
			data.ws.WriteString(js.Location(createLink("task", "view", []interface{}{"taskID=", taskID}), "parent"))
		}
		return
	}
	var result *protocol.MSG_PROJECT_task_UpdateTaskEstimate_result
	if err = data.SendMsgWaitResultToDefault(update, &result); err != nil {
		data.ws.WriteString(js.Error(err.Error()))
		return
	}
	update.Put()

	/* Remind whether to update status of the bug, if task which from that bug has been finished. */
	task, _ := task_getByID(data, int32(taskID))
	needUpdate, err := task_needUpdateBugStatus(data, task)
	if err != nil {
		data.ws.WriteString(js.Error(err.Error()))
		return
	}
	if len(result.Changes) > 0 && needUpdate {
		for _, change := range result.Changes {

			if change.Field == "status" && change.New == "done" {
				confirmURL := createLink("bug", "view", []interface{}{"id=", task.FromBug})
				cancelURL := createLink("task", "view", []interface{}{"taskID=", taskID})
				data.ws.WriteString(js.Confirm(fmt.Sprintf(data.Lang["task"]["remindBug"].(string), task.FromBug), confirmURL, cancelURL, "parent", "parent.parent"))
				return
			}
		}
	}
	session.Commit()
	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
	} else {
		data.ws.WriteString(js.Location(createLink("task", "view", []interface{}{"taskID=", taskID}), "parent"))
	}
	return
}

func task_needUpdateBugStatus(data *TemplateData, task *protocol.MSG_PROJECT_TASK) (bool, error) {
	if task.FromBug == 0 {
		return false, nil
	}
	out := protocol.GET_MSG_TEST_bug_getById()
	out.Id = task.FromBug
	var result *protocol.MSG_TEST_bug_getById_result
	if err := data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return false, err
	}
	res := true
	if result.Info.Status == "resolved" {
		res = false
	}
	result.Put()
	return res, nil
}

func post_task_edit(data *TemplateData) (err error) {
	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	comment := data.ws.Query("comment") != ""
	task := data.Data["task"].(*protocol.MSG_PROJECT_TASK)
	if strconv.FormatInt(task.LastEditedDate.Unix(), 10) != data.ws.Post("lastEditedDate") {
		data.ajaxResult(false, data.Lang["error"]["editedByOther"].(string), "")
		return nil
	}
	var change protocol.ChangeHistory
	if !comment {
		var ok bool
		if change, ok = task_createEdit(data, task); !ok {

			return
		}
	}

	project := data.Data["project"].(*protocol.MSG_PROJECT_project_cache)
	action := "Commented"
	if len(change) > 0 {
		action = "Edited"
	}
	actionid, err := data.Msg.ActionCreate("task", task.Id, action, data.ws.Post("comment"), "", project.Products, project.Id)
	if err != nil {
		return err
	}
	if len(change) > 0 {
		change.Add(actionid, data.Msg)
		/*if result.Info.FromBug > 0 {
			for _, c := range change {
				if c.Field == "Status" {
					confirmURL := createLink("bug", "view", []interface{}{"id=", result.Info.FromBug})
					cancelURL := data.ws.Referer()
					if cancelURL == "" {
						cancelURL = data.ws.Header("REFERER")
					}
					data.ajaxResult(true, nil, js.Confirm(fmt.Sprintf(data.Lang["task"]["remindBug"].(string), result.Info.FromBug), confirmURL, cancelURL, "parent", "parent"))
					return
				}
			}
		}*/
	}
	data.ajaxResult(true, nil, createLink("task", "view", []interface{}{"taskID=", taskID}), "parent")

	return
}
func get_task_start(data *TemplateData) (err error) {
	data.Data["title"] = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["start"].(string)
	task := data.Data["task"].(*protocol.MSG_PROJECT_TASK)
	if task.RealStarted.After(protocol.NORMALTIME) {
		data.Data["RealStarted"] = task.RealStarted.Format("2006-01-02")
	} else {
		data.Data["RealStarted"] = time.Now().Format("2006-01-02")
	}
	data.Data["Consumed"] = task.Consumed
	data.Data["Left"] = task.Left
	if len(task.Team) > 0 {
		out := protocol.GET_MSG_USER_team_getByIds()
		out.Ids = task.Team
		var result *protocol.MSG_USER_team_getByIds_result
		if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
			return
		}
		for _, team := range result.List {
			if team.Uid == task.AssignedTo {
				data.Data["Consumed"] = team.Consumed
				data.Data["Left"] = team.Left
			}
		}
		out.Put()
		result.Put()
	}
	data.Data["users"], err = user_getPairs(data, "noletter")
	if data.App["methodName"] == "restart" {
		data.Lang["task"]["start"] = data.Lang["task"]["restart"]
	}
	templateOut("task.start.html", data)
	return
}
func post_task_start(data *TemplateData) (err error) {
	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	out := protocol.GET_MSG_PROJECT_task_start()
	out.TaskID = int32(taskID)
	out.Consumed, _ = strconv.ParseFloat(data.ws.Post("consumed"), 64)
	if out.RealStarted, err = time.ParseInLocation("2006-01-02", data.ws.Post("realStarted"), time.Local); err != nil {
		data.ws.WriteString(js.Error(data.Lang["task"]["error"].(map[string]string)["RealStartedErr"]))
		return dataErrAlreadyOut
	}

	out.Left, _ = strconv.ParseFloat(data.ws.Post("left"), 64)
	out.Comment = data.ws.Post("comment")
	out.MethodName = data.App["methodName"].(string)
	var result *protocol.MSG_PROJECT_task_start_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		e := err.Error()
		if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
			e = v
		}
		data.ws.WriteString(js.Error(e))
		return dataErrAlreadyOut
	}
	task, _ := task_getByID(data, int32(taskID))
	if data.App["methodName"] == "start" && task != nil && len(result.Changes) > 0 {
		need, _ := task_needUpdateBugStatus(data, task)
		if need {
			for _, change := range result.Changes {
				if change.Field == "status" && change.New == "done" {
					confirmURL := createLink("bug", "view", []interface{}{"id=", task.FromBug})
					cancelURL := createLink("task", "view", "taskID="+data.ws.Query("taskID"))
					data.ws.WriteString(js.Confirm(fmt.Sprintf(data.Lang["task"]["remindBug"].(string), task.FromBug), confirmURL, cancelURL, "parent", "parent.parent"))
					return

				}
			}
		}
	}
	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
	} else {
		data.ws.WriteString(js.Location(createLink("task", "view", "taskID="+data.ws.Query("taskID")), "parent"))
	}
	return
}
func get_task_finish(data *TemplateData) (err error) {
	task := data.Data["task"].(*protocol.MSG_PROJECT_TASK)

	data.Data["nextBy"] = task.OpenedBy
	if len(task.Team) > 0 {
		out := protocol.GET_MSG_USER_team_getByIds()
		out.Ids = task.Team
		var result *protocol.MSG_USER_team_getByIds_result
		if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
			return
		}
		var memberPairs = []protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{}}
		for k, team := range result.List {
			if team.Uid == task.AssignedTo {
				if k+1 < len(result.List) {
					data.Data["nextBy"] = result.List[k+1].Uid

				}
				data.Data["myConsumed"] = team.Consumed
			}
			name := team.Realname
			if name == "" {
				name = team.Account
			}
			memberPairs = append(memberPairs, protocol.HtmlKeyValueStr{strconv.Itoa(int(team.Uid)), name})
		}

		if result.List[len(result.List)-1].Uid != task.AssignedTo {
			data.Data["members"] = memberPairs
		}
	}
	data.Data["title"] = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["finish"].(string)
	templateOut("task.finish.html", data)
	return
}
func post_task_finish(data *TemplateData) (err error) {
	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	out := protocol.GET_MSG_PROJECT_task_finish()
	out.TaskID = int32(taskID)
	if out.Consumed, err = strconv.ParseFloat(data.ws.Post("consumed"), 64); err != nil {
		data.ws.WriteString(js.Error(data.Lang["task"]["error"].(map[string]string)["consumedNumber"]))
		return dataErrAlreadyOut
	}
	if out.FinishedDate, err = time.ParseInLocation("2006-01-02", data.ws.Post("finishedDate"), time.Local); err != nil {
		data.ws.WriteString(js.Error(data.Lang["task"]["error"].(map[string]string)["finishedDateErr"]))
		return dataErrAlreadyOut
	}
	for _, v := range data.ws.PostSlice("mailto") {
		id, _ := strconv.Atoi(v)
		if user := HostConn.GetUserCacheById(int32(id)); user != nil {
			out.Mailto = append(out.Mailto, user.Id)
		}
	}
	out.Comment = data.ws.Post("comment")
	var result *protocol.MSG_PROJECT_task_finish_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		e := err.Error()
		if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
			e = v
		}
		data.ws.WriteString(js.Error(e))
		return dataErrAlreadyOut
	}
	task, _ := task_getByID(data, int32(taskID))
	if task != nil && len(result.Changes) > 0 {
		need, _ := task_needUpdateBugStatus(data, task)
		if need {
			for _, change := range result.Changes {
				if change.Field == "status" && change.New == "done" {
					confirmURL := createLink("bug", "view", []interface{}{"id=", task.FromBug})
					cancelURL := createLink("task", "view", "taskID="+data.ws.Query("taskID"))
					data.ws.WriteString(js.Confirm(fmt.Sprintf(data.Lang["task"]["remindBug"].(string), task.FromBug), confirmURL, cancelURL, "parent", "parent.parent"))
					return

				}
			}
		}
	}
	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
	} else {
		data.ws.WriteString(js.Location(createLink("task", "view", "taskID="+data.ws.Query("taskID")), "parent"))
	}
	return
}

func get_task_activate(data *TemplateData) (err error) {
	members := data.Data["members"].([]protocol.HtmlKeyValueStr)
	find := false
	task := data.Data["task"].(*protocol.MSG_PROJECT_TASK)
	for _, kv := range members {
		if kv.Key == strconv.Itoa(int(task.FinishedBy)) {
			find = true
			break
		}
	}
	if !find {
		if user := HostConn.GetUserCacheById(task.FinishedBy); user != nil {
			name := user.Realname
			if name == "" {
				name = user.Account
			}
			members = append(members, protocol.HtmlKeyValueStr{strconv.Itoa(int(user.Id)), name})
		}
	}
	data.Data["members"] = members
	data.Data["title"] = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["activate"].(string)
	if data.Data["user"], err = user_getPairs(data, "noletter"); err != nil {
		return
	}
	templateOut("task.activate.html", data)
	return
}

func post_task_activate(data *TemplateData) (err error) {
	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	out := protocol.GET_MSG_PROJECT_task_activate()
	out.TaskID = int32(taskID)
	out.Left, _ = strconv.ParseFloat(data.ws.Post("left"), 64)
	for _, v := range data.ws.PostSlice("mailto") {
		id, _ := strconv.Atoi(v)
		if user := HostConn.GetUserCacheById(int32(id)); user != nil {
			out.Mailto = append(out.Mailto, user.Id)
		}
	}
	assignedTo, _ := strconv.Atoi(data.ws.Post("assignedTo"))

	if HostConn.GetUserCacheById(int32(assignedTo)) == nil {
		data.ws.WriteString(js.Error(data.Lang["task"]["error"].(map[string]string)["AssignedToNotFoundUser"]))
		return
	}
	out.AssignedTo = int32(assignedTo)
	out.Comment = data.ws.Post("comment")

	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		e := err.Error()
		if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
			e = v
		}
		data.ws.WriteString(js.Error(e))
		return dataErrAlreadyOut
	}

	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
	} else {
		data.ws.WriteString(js.Location(createLink("task", "view", "taskID="+data.ws.Query("taskID")), "parent"))
	}
	return
}

func get_task_pause(data *TemplateData) (err error) {
	data.Data["title"] = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["pause"].(string)
	if data.Data["user"], err = user_getPairs(data, "noletter"); err != nil {
		return
	}
	templateOut("task.pause.html", data)
	return
}

func post_task_pause(data *TemplateData) (err error) {
	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	out := protocol.GET_MSG_PROJECT_task_pause()
	out.TaskID = int32(taskID)
	out.Comment = data.ws.Post("comment")
	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		e := err.Error()
		if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
			e = v
		}
		data.ws.WriteString(js.Error(e))
		return dataErrAlreadyOut
	}
	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
	} else {
		data.ws.WriteString(js.Location(createLink("task", "view", "taskID="+data.ws.Query("taskID")), "parent"))
	}
	return
}
func get_task_internalaudit(data *TemplateData) (err error) {
	task := data.Data["task"].(*protocol.MSG_PROJECT_TASK)
	assignedUser := HostConn.GetUserCacheById(task.AssignedTo)
	var foundAssignedTo bool
	for _, kv := range data.Data["members"].([]protocol.HtmlKeyValueStr) {
		if kv.Key == strconv.Itoa(int(task.AssignedTo)) {
			foundAssignedTo = true
			break
		}
	}
	if !foundAssignedTo && assignedUser != nil {
		name := assignedUser.Realname
		if name == "" {
			name = assignedUser.Account
		}
		data.Data["members"] = append(data.Data["members"].([]protocol.HtmlKeyValueStr), protocol.HtmlKeyValueStr{strconv.Itoa(int(assignedUser.Id)), name})
	}
	data.Data["title"] = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["internalaudit"].(string)
	if data.Data["user"], err = user_getPairs(data, "noletter"); err != nil {
		return
	}
	templateOut("task.internalaudit.html", data)
	return
}

func post_task_internalaudit(data *TemplateData) (err error) {
	task := data.Data["task"].(*protocol.MSG_PROJECT_TASK)
	session, err := data.Msg.BeginTransaction()
	defer func() {
		session.Rollback()
	}()
	out := protocol.GET_MSG_PROJECT_task_internalaudit()
	out.TaskID = task.Id
	out.Comment = data.ws.Post("comment")
	for _, v := range data.ws.PostSlice("mailto") {
		id, _ := strconv.Atoi(v)
		if user := HostConn.GetUserCacheById(int32(id)); user != nil {
			out.Mailto = append(out.Mailto, user.Id)
		}
	}
	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		e := err.Error()
		if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
			e = v
		}
		data.ajaxResult(false, e)
		return dataErrAlreadyOut
	}
	project := data.Data["project"].(*protocol.MSG_PROJECT_project_cache)
	product := HostConn.GetProductById(task.Product)
	//上传文件
	uploadFileTmp := protocol.GET_MSG_FILE_updateTmp()
	for k, name := range data.ws.PostSlice("uploadFileTmpName") {
		upload := protocol.GET_MSG_FILE_upload()
		upload.AddBy = data.User.Id
		upload.Name = file_getRealName(name)
		upload.Type = data.ws.PostSlice("fileType")[k]
		upload.Code = product.Code
		upload.Title = file_getTitleName(name)
		if upload.Code == "" {
			upload.Code = product.Name
		}
		if project.Code == "" {
			upload.Code += "/" + project.Name
		} else {
			upload.Code += "/" + project.Code
		}
		upload.ObjectID = task.Id
		upload.ObjectType = "task"
		uploadFileTmp.Files = append(uploadFileTmp.Files, upload)
	}
	if err = data.SendMsgWaitResultToDefault(uploadFileTmp, nil); err != nil {
		data.ajaxResult(false, data.Lang["file"]["error"].(map[string]string)["updateTmp"])
		return dataErrAlreadyOut
	}
	session.Commit()
	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
	} else {
		data.ws.WriteString(js.Location(createLink("task", "view", "taskID="+data.ws.Query("taskID")), "parent"))
	}
	return
}

func get_task_proofreading(data *TemplateData) (err error) {
	data.Data["title"] = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["proofreading"].(string)
	templateOut("task.proofreading.html", data)
	return
}

func post_task_proofreading(data *TemplateData) (err error) {

	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	out := protocol.GET_MSG_PROJECT_task_proofreading()
	out.TaskID = int32(taskID)
	out.Proofreading = data.ws.Post("proofreading") == "true"
	out.Comment = data.ws.Post("comment")

	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		e := err.Error()
		if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
			e = v
		}
		data.ws.WriteString(js.Error(e))
		return dataErrAlreadyOut
	}

	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
	} else {
		data.ws.WriteString(js.Location(createLink("task", "view", "taskID="+data.ws.Query("taskID")), "parent"))
	}
	return
}

func get_task_close(data *TemplateData) (err error) {

	data.Data["title"] = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["close"].(string)
	if data.Data["user"], err = user_getPairs(data, "noletter"); err != nil {
		return
	}
	templateOut("task.close.html", data)
	return
}

func post_task_close(data *TemplateData) (err error) {

	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	out := protocol.GET_MSG_PROJECT_task_close()
	out.TaskID = int32(taskID)
	out.Comment = data.ws.Post("comment")
	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		e := err.Error()
		if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
			e = v
		}
		data.ws.WriteString(js.Error(e))
		return dataErrAlreadyOut
	}

	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
	} else {
		data.ws.WriteString(js.Location(createLink("task", "view", "taskID="+data.ws.Query("taskID")), "parent"))
	}
	return
}
func get_task_batchCreate(data *TemplateData) (err error) {
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	storyID, _ := strconv.Atoi(data.ws.Query("storyID"))
	moduleID, _ := strconv.Atoi(data.ws.Query("moduleID"))
	project := data.Data["project"].(*protocol.MSG_PROJECT_project_cache)
	if err = project_setMenu(data, int32(projectID), 0, ""); err != nil {
		return
	}

	visibleFields := map[string]bool{}
	requiredFields := map[string]string{}
	for _, field := range strings.Split(data.Config["task"]["custom"]["batchCreateFields"].(string), ",") {
		visibleFields[field] = true
	}
	for _, field := range strings.Split(data.Config["task"]["create"]["requiredFields"].(string), ",") {
		requiredFields[field] = " required"
		for _, field2 := range data.Config["task"]["common"]["customBatchCreateFields"].([]string) {
			if field2 == field {
				visibleFields[field] = true
			}
		}
	}
	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	if taskID > 0 {
		if task, err := task_getByID(data, int32(taskID)); err != nil {
			return err
		} else {
			data.Data["parentTitle"] = task.Name
		}
	}

	data.Data["title"] = project.Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["batchCreate"].(string)

	data.Data["project"] = project
	if data.Data["stories"], err = story_getProjectStoryPairs(data, int32(projectID), 0, 0, nil, "short"); err != nil {
		return
	}
	if data.Data["modules"], err = tree_getTaskOptionMenu(data, int32(projectID), 0, 0); err != nil {
		return
	}
	data.Data["modules"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"ditto", data.Lang["task"]["ditto"].(string)}}, data.Data["modules"].([]protocol.HtmlKeyValueStr)...)
	data.Data["parent"] = taskID
	data.Data["storyID"] = storyID
	data.Data["module"] = moduleID
	if storyID > 0 {
		out := protocol.GET_MSG_PROJECT_story_getById()
		out.Id = int32(storyID)
		var result *protocol.MSG_PROJECT_story_getById_result
		if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
			return
		}
		data.Data["story"] = result.Story
		out.Put()
		if data.onlybody() {
			delete(visibleFields, "story")
		}
		data.Data["module"] = result.Story.Module
	}
	data.Data["colspan"] = len(visibleFields) + 3
	data.Data["visibleFields"] = visibleFields
	data.Data["requiredFields"] = requiredFields
	if len(data.Data["stories"].([]protocol.HtmlKeyValueStr)) > 0 {
		out := protocol.GET_MSG_PROJECT_task_getStoryTaskCounts()
		out.ProjectID = int32(projectID)
		for _, kv := range data.Data["stories"].([]protocol.HtmlKeyValueStr) {
			id, _ := strconv.Atoi(kv.Key)
			out.Stories = append(out.Stories, int32(id))
		}
		var result *protocol.MSG_PROJECT_task_getStoryTaskCounts_result
		if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
			return
		}
		data.Data["storyTasks"] = result.List
	} else {
		data.Data["storyTasks"] = make(map[int32]int)
	}
	data.Data["stories"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"ditto", data.Lang["task"]["ditto"].(string)}}, data.Data["stories"].([]protocol.HtmlKeyValueStr)...)
	data.Data["members"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"ditto", data.Lang["task"]["ditto"].(string)}}, data.Data["members"].([]protocol.HtmlKeyValueStr)...)

	data.Data["moduleID"] = moduleID

	data.Data["customLink"] = createLink("custom", "ajaxSaveCustomFields", "module=task&section=custom&key=batchCreateFields")
	data.Lang["task"]["typeList"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"ditto", data.Lang["task"]["ditto"].(string)}}, data.Lang["task"]["typeList"].([]protocol.HtmlKeyValueStr)...)
	templateOut("task.batchcreate.html", data)
	return nil
}
func post_task_batchCreate(data *TemplateData) (err error) {
	var parentID, ancestorID int32
	id, _ := strconv.Atoi(data.ws.Post("parent[0]"))
	parentID = int32(id)
	var parentTask *protocol.MSG_PROJECT_TASK
	if parentID > 0 {
		if parentTask, err = task_getByID(data, parentID); err != nil {
			return err
		} else {
			ancestorID = parentTask.Parent
		}
	}
	project := data.Data["project"].(*protocol.MSG_PROJECT_project_cache)
	var tasks []*protocol.MSG_PROJECT_TASK
	for i := 0; i < data.Config["task"]["common"]["batchCreate"].(int); i++ {
		task := &protocol.MSG_PROJECT_TASK{Name: data.ws.Post("name[" + strconv.Itoa(i) + "]")}

		if task.Name == "" {
			continue
		}
		story := data.ws.Post("story[" + strconv.Itoa(i) + "]")
		module := data.ws.Post("module[" + strconv.Itoa(i) + "]")
		task.Type = data.ws.Post("type[" + strconv.Itoa(i) + "]")
		assignedTo := data.ws.Post("assignedTo[" + strconv.Itoa(i) + "]")
		if story == "ditto" {
			task.Story = tasks[i-1].Story
		} else {
			storyID, _ := strconv.Atoi(story)
			task.Story = int32(storyID)
		}
		if module == "ditto" {
			task.Module = tasks[i-1].Module
		} else {
			moduleID, _ := strconv.Atoi(module)
			task.Module = int32(moduleID)
		}
		if task.Type == "ditto" {
			task.Type = tasks[i-1].Type
		} else if task.Type == "" {
			return errors.New(data.Lang["task"]["error"].(map[string]string)["errorTaskType"])
		}
		if assignedTo == "ditto" {
			task.AssignedTo = tasks[i-1].AssignedTo
		} else {
			id, _ := strconv.Atoi(assignedTo)
			if user := HostConn.GetUserCacheById(int32(id)); user != nil {
				task.AssignedTo = int32(id)
				task.AssignedDate = time.Now()
			}

		}
		task.Color = data.ws.Post("color[" + strconv.Itoa(i) + "]")
		task.Desc = data.ws.Post("color[" + strconv.Itoa(i) + "]")
		pri, _ := strconv.Atoi(data.ws.Post("pri[" + strconv.Itoa(i) + "]"))
		task.Pri = int8(pri)
		estimate := data.ws.Post("estimate[" + strconv.Itoa(i) + "]")
		task.Estimate, err = strconv.ParseFloat(estimate, 64)
		if task.Estimate < 0 || (err != nil && estimate != "") {
			return errors.New(data.Lang["task"]["error"].(map[string]string)["estimateNumber"])
		}
		task.Left = task.Estimate
		task.Project = project.Id
		task.Ancestor = ancestorID
		if estStarted := data.ws.Post("estStarted[" + strconv.Itoa(i) + "]"); estStarted == "" {
			if parentTask != nil {
				task.EstStarted = parentTask.EstStarted
			} else {
				task.EstStarted = protocol.ZEROTIME
			}
		} else {
			if task.EstStarted, err = time.ParseInLocation("2006-01-02", estStarted, time.Local); err != nil {

				return errors.New(data.Lang["task"]["error"].(map[string]string)["estStarted"])
			}
		}
		if deadline := data.ws.Post("deadline[" + strconv.Itoa(i) + "]"); deadline == "" {
			if parentTask != nil {
				task.Deadline = parentTask.Deadline
			} else {
				task.Deadline = protocol.ZEROTIME
			}
		} else {
			if task.Deadline, err = time.ParseInLocation("2006-01-02", deadline, time.Local); err != nil {

				return errors.New(data.Lang["task"]["error"].(map[string]string)["deadline"])
			}
		}
		task.Status = "wait"
		task.OpenedBy = data.User.Id
		task.OpenedDate = time.Now()
		task.Parent = parentID
		tasks = append(tasks, task)
	}
	session, err := data.BeginTransaction()
	if err != nil {

		return
	}
	for _, task := range tasks {

		if _, ok := do_task_create(data, task); !ok {
			session.Rollback()
			return nil
		}
	}
	session.Commit()
	if data.ws.Query("iframe") != "" {
		data.ajaxResult(true, "", "parent")
	} else {
		storyLink := data.ws.Session().Load_str("storyList")
		if storyLink == "" {
			storyLink = createLink("project", "story", "projectID="+data.ws.Query("projectID"))
		}
		data.ajaxResult(true, data.Lang["common"]["saveSuccess"].(string), storyLink)
	}
	return nil
}

func get_task_examine(data *TemplateData) (err error) {

	data.Data["title"] = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["examine"].(string)
	if data.Data["user"], err = user_getPairs(data, "noletter"); err != nil {
		return
	}
	templateOut("task.examine.html", data)
	return
}

func post_task_examine(data *TemplateData) (err error) {

	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	out := protocol.GET_MSG_PROJECT_task_examine()
	out.TaskID = int32(taskID)
	out.Examine = data.ws.Post("examine") == "true"
	out.ProjectId = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Id
	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		e := err.Error()
		if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
			e = v
		}
		data.ws.WriteString(js.Error(e))
		return dataErrAlreadyOut
	}

	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
	} else {
		data.ws.WriteString(js.Location(createLink("task", "view", "taskID="+data.ws.Query("taskID")), "parent"))
	}
	return
}
func get_task_cancel(data *TemplateData) (err error) {

	data.Data["title"] = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["cancel"].(string)
	if data.Data["user"], err = user_getPairs(data, "noletter"); err != nil {
		return
	}
	templateOut("task.cancel.html", data)
	return
}

func post_task_cancel(data *TemplateData) (err error) {

	taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
	out := protocol.GET_MSG_PROJECT_task_cancel()
	out.TaskID = int32(taskID)
	out.Comment = data.ws.Post("comment")

	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		e := err.Error()
		if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
			e = v
		}
		data.ws.WriteString(js.Error(e))
		return dataErrAlreadyOut
	}

	if data.onlybody() {
		data.ws.WriteString(js.CloseModal("parent.parent", "this", ""))
	} else {
		data.ws.WriteString(js.Location(createLink("task", "view", "taskID="+data.ws.Query("taskID")), "parent"))
	}
	return
}
func get_task_delete(data *TemplateData) (err error) {
	if data.ws.Query("confirm") == "yes" {
		taskID, _ := strconv.Atoi(data.ws.Query("taskID"))
		out := protocol.GET_MSG_PROJECT_task_delete()
		out.TaskID = int32(taskID)
		out.ProjectId = data.Data["project"].(*protocol.MSG_PROJECT_project_cache).Id
		if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
			e := err.Error()
			if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
				e = v
			}
			data.ws.WriteString(js.Error(e))
			return dataErrAlreadyOut
		}

		data.ws.WriteString(js.Location(data.ws.Session().Load_str("taskList"), "parent"))

	} else {
		data.ws.WriteString(js.Confirm(data.Lang["task"]["confirmDelete"].(string), createLink("task", "delete", []interface{}{"projectID=", data.ws.Query("projectID"), "&taskID=", data.ws.Query("taskID"), "&confirm=yes"}), ""))
	}

	return
}
func post_task_batchCancel(data *TemplateData) error {
	out := protocol.GET_MSG_PROJECT_task_cancel()
	session, err := data.Msg.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	for _, idstr := range data.ws.PostSlice("taskIDList") {
		id, e := strconv.Atoi(idstr)
		if e == nil {
			out.TaskID = int32(id)
			if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
				data.ws.WriteString(js.Reload("parent"))
				return dataErrAlreadyOut
			}
		}

	}
	data.ws.WriteString(js.Reload("parent"))
	session.Commit()
	return nil
}
func post_task_batchClose(data *TemplateData) error {
	out := protocol.GET_MSG_PROJECT_task_close()
	session, err := data.Msg.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	for _, idstr := range data.ws.PostSlice("taskIDList") {
		id, e := strconv.Atoi(idstr)
		if e == nil {
			out.TaskID = int32(id)
			if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
				data.ws.WriteString(js.Reload("parent"))
				return dataErrAlreadyOut
			}
		}

	}
	data.ws.WriteString(js.Reload("parent"))
	session.Commit()
	return nil
}
func post_task_batchEdit(data *TemplateData) (err error) {
	project := data.Data["project"].(*protocol.MSG_PROJECT_project_cache)
	if project != nil {

		/* Set modules and members. */
		modules, err := tree_getTaskOptionMenu(data, project.Id, 0, 0)
		if err != nil {
			return err
		}
		modules = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"ditto", data.Lang["task"]["ditto"].(string)}}, modules...)
		members, err := project_getTeamMemberPairs(data, project.Id, "nodeleted")
		if err != nil {
			return err
		}
		members = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{}, protocol.HtmlKeyValueStr{"ditto", data.Lang["task"]["ditto"].(string)}}, members...)
		members = append(members, protocol.HtmlKeyValueStr{"close", "Closed"})
		data.Data["title"] = project.Name + data.Lang["common"]["colon"].(string) + data.Lang["task"]["batchEdit"].(string)
		data.Data["modules"] = modules
		data.Data["members"] = members
		data.Data["projectName"] = project.Name
	} else {
		data.Data["projectName"] = ""

		/* The tasks of my.

		this->lang->task->menu = this->lang->my->menu;
		this->lang->set("menugroup.task", "my");
		this->lang->task->menuOrder = this->lang->my->menuOrder;
		this->loadModel("my")->setMenu();
		data.Data["position[] = html::a(this->createLink("my", "task"), this->lang->my->task);
		data.Data["title      = this->lang->task->batchEdit;
		data.Data["users      = this->loadModel("user")->getPairs("noletter");
		*/
	}
	/* Get edited tasks. */

	var customFields []protocol.HtmlKeyValueStr
	for _, field := range data.Config["task"]["common"]["customBatchEditFields"].([]string) {
		name, ok := data.Lang["task"][field].(string)
		if !ok {
			return errors.New(`无法获取Lang["task"]["` + field + `"]的值`)
		}
		customFields = append(customFields, protocol.HtmlKeyValueStr{field, name})
	}
	data.Data["customFields"] = customFields
	data.Data["showFields"] = strings.Split(data.Config["task"]["custom"]["batchEditFields"].(string), ",")
	visibleFields := map[string]bool{}
	requiredFields := map[string]string{}
	for _, field := range strings.Split(data.Config["task"]["custom"]["batchEditFields"].(string), ",") {
		visibleFields[field] = true
	}
	for _, field := range strings.Split(data.Config["task"]["edit"]["requiredFields"].(string), ",") {
		requiredFields[field] = " required"
		for _, field2 := range data.Config["task"]["common"]["customBatchCreateFields"].([]string) {
			if field2 == field {
				visibleFields[field] = true
			}
		}
	}
	data.Data["visibleFields"] = visibleFields
	data.Data["requiredFields"] = requiredFields
	/* Assign. */

	data.Data["projectID"] = data.ws.Query("projectID")
	data.Data["priList"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"0", ""}, protocol.HtmlKeyValueStr{"ditto", data.Lang["task"]["ditto"].(string)}}, data.Lang["task"]["priList"].([]protocol.HtmlKeyValueStr)...)
	data.Data["statusList"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"0", ""}, protocol.HtmlKeyValueStr{"ditto", data.Lang["task"]["ditto"].(string)}}, data.Lang["task"]["statusList"].([]protocol.HtmlKeyValueStr)...)
	data.Data["typeList"] = append([]protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"0", ""}, protocol.HtmlKeyValueStr{"ditto", data.Lang["task"]["ditto"].(string)}}, data.Lang["task"]["typeList"].([]protocol.HtmlKeyValueStr)...)
	data.Data["taskIDList"] = data.ws.PostSlice("taskIDList")
	getTasks := protocol.GET_MSG_PROJECT_task_getListByWhereMap()
	getTasks.Where = map[string]interface{}{"Id": data.ws.PostSlice("taskIDList")}
	getTasks.PerPage = 99999999
	getTasks.Page = 1
	var getTasksResult *protocol.MSG_PROJECT_task_getListByWhereMap_result
	if err = data.SendMsgWaitResultToDefault(getTasks, &getTasksResult); err != nil {
		return
	}
	data.Data["tasks"] = getTasksResult.List
	getTeams := protocol.GET_MSG_USER_team_getByTypeRoot()
	getTeams.Type = "task"
	var disableHour = make(map[int32]string)
	for _, task := range getTasksResult.List {
		getTeams.Root = append(getTeams.Root, task.Id)
		disableHour[task.Id] = ""
	}
	var getTeamsResult *protocol.MSG_USER_team_getByTypeRoot_result
	if err = data.SendMsgWaitResultToDefault(getTeams, &getTeamsResult); err != nil {
		return err
	}
	data.Data["customLink"] = createLink("custom", "ajaxSaveCustomFields", "module=task&section=custom&key=batchEditFields")

	for _, team := range getTeamsResult.List {
		disableHour[team.Root] = "disabled='disabled'"
	}
	data.Data["disableHour"] = disableHour
	templateOut("task.batchEdit.html", data)

	getTeams.Put()
	getTasksResult.Put()
	getTasks.Put()
	getTasksResult.Put()
	return
}
func post_task_batchexamine(data *TemplateData) error {
	out := protocol.GET_MSG_PROJECT_task_examine()
	session, err := data.Msg.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	for _, idstr := range data.ws.PostSlice("taskIDList") {
		id, e := strconv.Atoi(idstr)
		if e == nil {
			out.TaskID = int32(id)
			out.Examine = true
			if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
				data.ws.WriteString(js.Reload("parent"))
				return dataErrAlreadyOut
			}
		}

	}
	data.ws.WriteString(js.Reload("parent"))
	session.Commit()
	return nil
}
func post_task_batchexaminec(data *TemplateData) error {
	out := protocol.GET_MSG_PROJECT_task_examine()
	session, err := data.Msg.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	for _, idstr := range data.ws.PostSlice("taskIDList") {
		id, e := strconv.Atoi(idstr)
		if e == nil {
			out.TaskID = int32(id)
			out.Examine = false
			if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
				data.ws.WriteString(js.Reload("parent"))
				return dataErrAlreadyOut
			}
		}

	}
	data.ws.WriteString(js.Reload("parent"))
	session.Commit()
	return nil
}
func post_task_batchproofreading(data *TemplateData) error {
	out := protocol.GET_MSG_PROJECT_task_proofreading()
	session, err := data.Msg.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	for _, idstr := range data.ws.PostSlice("taskIDList") {
		id, e := strconv.Atoi(idstr)
		if e == nil {
			out.TaskID = int32(id)
			out.Proofreading = true
			if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
				data.ws.WriteString(js.Reload("parent"))
				return dataErrAlreadyOut
			}
		}

	}
	data.ws.WriteString(js.Reload("parent"))
	session.Commit()
	return nil
}
func post_task_batchproofreadingc(data *TemplateData) error {
	out := protocol.GET_MSG_PROJECT_task_proofreading()
	session, err := data.Msg.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	for _, idstr := range data.ws.PostSlice("taskIDList") {
		id, e := strconv.Atoi(idstr)
		if e == nil {
			out.TaskID = int32(id)
			out.Proofreading = false
			if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
				data.ws.WriteString(js.Reload("parent"))
				return dataErrAlreadyOut
			}
		}

	}
	data.ws.WriteString(js.Reload("parent"))
	session.Commit()
	return nil
}
func post_task_finishall(data *TemplateData) error {
	out := protocol.GET_MSG_PROJECT_task_finish()
	session, err := data.Msg.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	for _, idstr := range data.ws.PostSlice("taskIDList") {
		id, e := strconv.Atoi(idstr)
		if e == nil {
			out.TaskID = int32(id)
			out.FinishedDate = time.Now()
			if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
				data.ws.WriteString(js.Reload("parent"))
				return dataErrAlreadyOut
			}
		}

	}
	data.ws.WriteString(js.Reload("parent"))
	session.Commit()
	return nil
}
func post_task_exportfinish(data *TemplateData) error {
	out := protocol.GET_MSG_FILE_getByWhere()
	out.Where = map[string]interface{}{
		"ObjectType": "task",
		"ObjectID":   data.ws.PostSlice("taskIDList"),
		"Type":       "FinalFile",
	}
	out.PerPage = 999999999
	out.Page = 1
	out.Total = 1
	var result *protocol.MSG_FILE_getByWhere_result
	if err := data.SendMsgWaitResultToDefault(out, &result); err != nil {
		data.ws.WriteString(js.Error(err.Error()))
		return nil
	}
	getDownload := protocol.GET_MSG_FILE_download_byIds()
	for _, file := range result.List {
		getDownload.Ids = append(getDownload.Ids, file.FileID)
	}

	var newId *protocol.MSG_FILE_download_byIds_result
	if err := data.SendMsgWaitResultToDefault(getDownload, &newId, 300*time.Second); err != nil {
		data.ws.WriteString(js.Error(err.Error()))
		return nil
	}
	data.ws.Redirect(createLink("file", "download", "fileID="+strconv.FormatInt(newId.FileID, 10)))
	getDownload.Put()
	newId.Put()
	out.Put()
	result.Put()
	return nil
}
func post_task_placeOrder(data *TemplateData) (err error) {
	out := protocol.GET_MSG_PROJECT_task_placeOrder()
	session, err := data.Msg.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	for _, idstr := range data.ws.PostSlice("taskIDList") {
		id, e := strconv.Atoi(idstr)
		if e == nil {
			out.TaskID = int32(id)
			out.Action = data.ws.Query("action") == "1"
			if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
				data.ws.WriteString(js.Reload("parent"))
				return dataErrAlreadyOut
			}
		}

	}
	data.ws.WriteString(js.Reload("parent"))
	session.Commit()
	return nil
}
func post_task_batchAssignTo(data *TemplateData) error {
	assignedTo, _ := strconv.Atoi(data.ws.Post("assignedTo"))
	if user := HostConn.GetUserCacheById(int32(assignedTo)); user == nil {
		data.ws.WriteString(js.Error(data.Lang["task"]["error"].(map[string]string)["AssignedToNotFoundUser"]))
		return nil
	}
	out := protocol.GET_MSG_PROJECT_task_assignTo()
	session, err := data.Msg.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	for _, idstr := range data.ws.PostSlice("taskIDList") {
		id, e := strconv.Atoi(idstr)
		if e == nil {
			out.TaskID = int32(id)
			out.AssignedTo = int32(assignedTo)
			if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
				e := err.Error()
				if v, ok := data.Lang["task"]["error"].(map[string]string)[e]; ok {
					e = v
				}
				data.ws.WriteString(js.Error("任务ID" + idstr + "错误" + e))
				return dataErrAlreadyOut
			}
		}

	}
	data.ws.WriteString(js.Reload("parent"))
	session.Commit()
	return nil
}
func task_getUserTasks(data *TemplateData, uid int32, typ string, pager *TempLatePage, orderBy string) ([]*protocol.MSG_PROJECT_TASK, error) {
	if typ == "" {
		typ = "assignedTo"
	}
	typ = strings.ToUpper(typ[:1]) + typ[1:]
	if typ != "all" {
		if field, ok := reflect.TypeOf(protocol.MSG_PROJECT_TASK{}).FieldByName(typ); !ok || field.Type.Kind() != reflect.Int32 {
			return nil, nil
		}
	}

	out := protocol.GET_MSG_PROJECT_project_getProjectTasksByWhere()
	out.Where = "Deleted=0"
	if typ == "AssignedTo" {
		out.Where += " and Status!='closed'"
	}
	if typ != "all" {
		out.Where += " and " + typ + "=" + strconv.Itoa(int(uid))
	}
	if pager == nil {
		pager = &TempLatePage{
			Total:   -1,
			Page:    1,
			PerPage: 999999999,
		}
	}
	out.PerPage = pager.PerPage
	out.Page = pager.Page
	out.Total = pager.Total
	out.OrderBy = orderBy
	var result *protocol.MSG_PROJECT_project_getProjectTasks_result
	if err := data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return nil, err
	}
	out.Put()
	pager.Total = result.Total
	return result.List, nil
}
func get_task_export(data *TemplateData) (err error) {
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	typ := data.ws.Query("projectID")
	project := data.getCacheProjectById(int32(projectID))
	allExportFields := data.Config["task"]["common"]["exportFields"].([]string)
	if project.Type == "ops" {
		for i := len(allExportFields) - 1; i >= 0; i-- {
			if allExportFields[i] == "story" {
				allExportFields = append(allExportFields[:i], allExportFields[i+1:]...)
				break
			}
		}
	}
	var browseType string

	fileName := data.Lang["task"]["common"].(string)

	if v, ok := data.Lang["project"]["featureBar"].(map[string][]protocol.HtmlKeyValueStr)["task"]; ok {
		for _, kv := range v {
			if kv.Key == typ {
				browseType = kv.Value
			}
		}

	} else {
		for _, kv := range data.Lang["project"]["statusSelects"].([]protocol.HtmlKeyValueStr) {
			if kv.Key == typ {
				browseType = kv.Value
			}
		}
	}

	data.Data["fileName"] = project.Name + data.Lang["common"]["dash"].(string) + browseType + fileName
	data.Data["allExportFields"] = allExportFields
	data.Data["customExport"] = true
	data.Data["orderBy"] = data.ws.Query("orderBy")
	data.Data["type"] = typ
	data.Data["projectID"] = projectID
	templateOut("file.export.html", data)
	return
}
func post_task_export(data *TemplateData) error {
	taskLang := data.Lang["task"]
	exportType := data.ws.Post("exportType")
	/* Create field lists. */
	//fields = this->post->exportFields ? this->post->exportFields : explode(",", allExportFields);
	projectID, _ := strconv.Atoi(data.ws.Query("projectID"))
	project := data.getCacheProjectById(int32(projectID))
	fields := data.Config["task"]["common"]["exportFields"].([]string)
	if project.Type == "ops" {
		for i := len(fields) - 1; i >= 0; i-- {
			if fields[i] == "story" {
				fields = append(fields[:i], fields[i+1:]...)
				break
			}
		}
	}
	var fieldkv []protocol.HtmlKeyValueStr
	for _, fieldName := range fields {
		name := fieldName
		if v, ok := taskLang[fieldName].(string); ok {
			name = v
		}
		fieldkv = append(fieldkv, protocol.HtmlKeyValueStr{fieldName, name})
	}
	var result *protocol.MSG_PROJECT_project_getProjectTasks_result
	if exportType == "selected" {
		out := protocol.GET_MSG_PROJECT_project_getProjectTasks()
		out.Type = []string{"checkedItem", data.ws.Cookie("checkedItem")}
		out.OrderBy = data.ws.Query("orderBy")
		out.Page = 1
		out.PerPage = 99999999
		out.Total = -1
		if err := data.SendMsgWaitResultToDefault(out, &result); err != nil {
			data.ws.WriteString(js.Error(err.Error()))
			return nil
		}
	} else {
		var msg1 *protocol.MSG_PROJECT_project_getProjectTasks
		var msg2 *protocol.MSG_PROJECT_project_getProjectTasksByWhere
		if ok := data.ws.Session().Get("project_task_msg", &msg1); ok {
			msg1.Page = 1
			msg1.PerPage = 99999999
			msg1.Total = -1
			if err := data.SendMsgWaitResultToDefault(msg1, &result); err != nil {
				data.ws.WriteString(js.Error(err.Error()))
				return nil
			}
		} else if ok := data.ws.Session().Get("project_task_msg", &msg2); ok {
			msg2.Page = 1
			msg2.PerPage = 99999999
			msg2.Total = -1
			if err := data.SendMsgWaitResultToDefault(msg2, &result); err != nil {
				data.ws.WriteString(js.Error(err.Error()))
				return nil
			}
		}
	}
	if result == nil || len(result.List) == 0 {
		data.ws.WriteString(js.Error(taskLang["error"].(map[string]string)["taskExportNotFoundTasks"]))
		return nil
	}
	/* Get tasks.*/

	/* Get users and projects.*/
	users, err := user_getPairs(data, "noletter")
	if err != nil {
		data.ws.WriteString(js.Error(err.Error()))
		return nil
	}
	projects, err := project_getPairs(data, "all|nocode")
	if err != nil {
		data.ws.WriteString(js.Error(err.Error()))
		return nil
	}
	relatedModules, err := tree_getTaskOptionMenu(data, int32(projectID), 0, 0)
	if err != nil {
		data.ws.WriteString(js.Error(err.Error()))
		return nil
	}
	var taskids []int32
	for _, task := range result.List {
		taskids = append(taskids, task.Id)
	}
	getFile := protocol.GET_MSG_FILE_getByWhere()
	getFile.Where = map[string]interface{}{
		"ObjectType": "task",
		"ObjectID":   taskids,
		"Deleted":    0,
	}
	getFile.Page = 1
	getFile.PerPage = 99999999
	getFile.Total = -1
	var getFileResult *protocol.MSG_FILE_getByWhere_result
	if err := data.SendMsgWaitResultToDefault(getFile, &getFileResult); err != nil {
		data.ws.WriteString(js.Error(err.Error()))
		return nil
	}
	var fileMap = make(map[int32][]*protocol.MSG_FILE_getByID_result)
	for _, file := range getFileResult.List {
		fileMap[file.ObjectID] = append(fileMap[file.ObjectID], file)
	}
	var exportMap []map[string]string
	for _, task := range result.List {
		if task.Consumed == 0 && task.Left == 0 {
			task.Progress = 0
		} else if task.Consumed != 0 && task.Left == 0 {
			task.Progress = 100
		} else {
			task.Progress = int(task.Consumed / (task.Consumed + task.Left) * 100)
		}
		rowMap := make(map[string]string, len(fields))
		for _, field := range fields {
			switch field {
			case "desc":
				rowMap[field] = libraries.Bbcode2html(task.Desc, true, false, false, false, true, false)
			case "project":
				for _, kv := range projects {
					if kv.Key == strconv.Itoa(int(task.Project)) {
						rowMap[field] = fmt.Sprintf("%s(#%d)", kv.Value, task.Project)
						break
					}
				}
			case "type":
				for _, kv := range taskLang["typeList"].([]protocol.HtmlKeyValueStr) {
					if kv.Key == task.Type {
						rowMap[field] = kv.Value
						break
					}
				}
			case "pri":
				for _, kv := range taskLang["priList"].([]protocol.HtmlKeyValueStr) {
					if kv.Key == strconv.Itoa(int(task.Pri)) {
						rowMap[field] = kv.Value
						break
					}
				}
			case "status":
				for _, kv := range taskLang["statusList"].([]protocol.HtmlKeyValueStr) {
					if kv.Key == task.Status {
						rowMap[field] = kv.Value
						break
					}
				}
			case "closedReason":
				for _, kv := range taskLang["reasonList"].([]protocol.HtmlKeyValueStr) {
					if kv.Key == task.ClosedReason {
						rowMap[field] = kv.Value
						break
					}
				}
			case "module":
				if task.Module > 0 {
					for _, kv := range relatedModules {
						if kv.Key == strconv.Itoa(int(task.Module)) {
							rowMap[field] = fmt.Sprintf("%s(#%d)", kv.Value, task.Module)
							break
						}
					}
				}
			case "openedBy":
				for _, kv := range users {
					if kv.Key == strconv.Itoa(int(task.OpenedBy)) {
						rowMap[field] = kv.Value
						break
					}
				}
			case "assignedTo":
				for _, kv := range users {
					if kv.Key == strconv.Itoa(int(task.AssignedTo)) {
						rowMap[field] = kv.Value
						break
					}
				}
			case "finishedBy":
				for _, kv := range users {
					if kv.Key == strconv.Itoa(int(task.FinishedBy)) {
						rowMap[field] = kv.Value
						break
					}
				}
			case "canceledBy":
				for _, kv := range users {
					if kv.Key == strconv.Itoa(int(task.CanceledBy)) {
						rowMap[field] = kv.Value
						break
					}
				}
			case "closedBy":
				for _, kv := range users {
					if kv.Key == strconv.Itoa(int(task.ClosedBy)) {
						rowMap[field] = kv.Value
						break
					}
				}
			case "lastEditedBy":
				for _, kv := range users {
					if kv.Key == strconv.Itoa(int(task.LastEditedBy)) {
						rowMap[field] = kv.Value
						break
					}
				}
			case "openedDate":
				rowMap[field] = task.OpenedDate.Format(protocol.TIMEFORMAT_MYSQLDATE)
			case "assignedDate":
				if task.AssignedDate.After(protocol.NORMALTIME) {
					rowMap[field] = task.AssignedDate.Format(protocol.TIMEFORMAT_MYSQLDATE)
				}
			case "finishedDate":
				if task.FinishedDate.After(protocol.NORMALTIME) {
					rowMap[field] = task.FinishedDate.Format(protocol.TIMEFORMAT_MYSQLDATE)
				}
			case "canceledDate":
				if task.CanceledDate.After(protocol.NORMALTIME) {
					rowMap[field] = task.CanceledDate.Format(protocol.TIMEFORMAT_MYSQLDATE)
				}
			case "closedDate":
				if task.ClosedDate.After(protocol.NORMALTIME) {
					rowMap[field] = task.ClosedDate.Format(protocol.TIMEFORMAT_MYSQLDATE)
				}
			case "lastEditedDate":
				if task.LastEditedDate.After(protocol.NORMALTIME) {
					rowMap[field] = task.LastEditedDate.Format(protocol.TIMEFORMAT_MYSQLDATE)
				}
			case "proofreading":
				if task.Proofreading {
					rowMap[field] = "已对"
				} else {
					rowMap[field] = "未对"
				}
			case "id":
				rowMap[field] = strconv.Itoa(int(task.Id))
			case "story":
				if task.Story > 0 {
					rowMap[field] = fmt.Sprintf("%s(#%d)", task.StoryTitle, task.Story)
				}

			case "name":
				rowMap[field] = task.Name
			case "estStarted":
				if task.EstStarted.After(protocol.NORMALTIME) {
					rowMap[field] = task.EstStarted.Format(protocol.TIMEFORMAT_MYSQLDATE)
				}
			case "realStarted":
				if task.RealStarted.After(protocol.NORMALTIME) {
					rowMap[field] = task.RealStarted.Format(protocol.TIMEFORMAT_MYSQLDATE)
				}
			case "deadline":
				rowMap[field] = task.Deadline.Format(protocol.TIMEFORMAT_MYSQLDATE)
			case "estimate":
				rowMap[field] = strconv.FormatFloat(task.Estimate, 'g', -1, 64)
			case "consumed":
				rowMap[field] = strconv.FormatFloat(task.Consumed, 'g', -1, 64)
			case "left":
				rowMap[field] = strconv.FormatFloat(task.Left, 'g', -1, 64)
			case "mailto":
				var str []string
				for _, id := range task.Mailto {
					for _, kv := range users {
						if kv.Key == strconv.Itoa(int(id)) {
							str = append(str, kv.Value)
							break
						}
					}
				}
				rowMap[field] = strings.Join(str, ",")
			case "progress":
				rowMap[field] = strconv.Itoa(int(task.Progress)) + "%"
			case "files":
				var files []string
				for _, file := range fileMap[task.Id] {
					ext := ""
					if i := strings.LastIndex(file.Name, "."); i > -1 {
						ext = file.Name[i+1:]
					}
					switch strings.ToLower(ext) {
					case "webp", "jpg", "jpeg", "png", "bmp":
						files = append(files, html_a(createLink("file", "read", fmt.Sprintf("fileID=%d", file.FileID)), file.Name))
					default:
						files = append(files, html_a(createLink("file", "download", fmt.Sprintf("fileID=%d", file.FileID)), file.Name))
					}
				}
				rowMap[field] = strings.Join(files, "||||")
			default:
				libraries.ReleaseLog("导出列名称" + field + "未处理")
			}

		}
		exportMap = append(exportMap, rowMap)
	}
	switch data.ws.Post("fileType") {
	case "xlsx":
		if err := file_export2xlsx(data, data.ws.Post("fileName"), fieldkv, exportMap); err != nil {
			data.ws.WriteString(js.Error(err.Error()))
		}
	default:
		data.ws.WriteString(js.Error("未处理导出格式" + data.ws.Post("fileType")))
	}
	return nil
}
