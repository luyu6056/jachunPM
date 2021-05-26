package handler

import (
	"fmt"
	"jachunPM_http/config"
	"libraries"
	"math"
	"protocol"
	"strconv"
	"strings"
)

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

		case "id":
			buf.WriteString(" cell-id'")

		case "deadline":
			if task.Delay > 0 {
				buf.WriteString(" text-center delayed'")
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
		}
		//echo "<td class='" . $class . "'" . $title . ">";
		switch id {
		case "id":
			if canBatchAction {

				buf.WriteString(html_checkbox("taskIDList", []protocol.HtmlKeyValueStr{{strconv.Itoa(int(task.Id)), ""}}))
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
				buf.WriteString(html_a(createLink("project", "browse", "projectID="+strconv.Itoa(int(task.Project))), project.Name))
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
			if task.Team > 0 {
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
				for _, kv := range data.Lang["task"]["statusList"].([]*protocol.HtmlKeyValueStr) {
					if kv.Key == task.Status {
						buf.WriteString(kv.Value)
						break
					}

				}

				buf.WriteString("</span>")
			}
		case "estimate":
			buf.WriteString(fmt.Sprintf("%0.1f", math.Round(task.Estimate*10)/10))
		case "consumed":
			buf.WriteString(fmt.Sprintf("%0.1f", math.Round(task.Consumed*10)/10))

		case "left":
			buf.WriteString(fmt.Sprintf("%0.1f", math.Round(task.Left*10)/10))

		case "progress":
			buf.WriteString(strconv.Itoa(int(task.Progress)))
			buf.WriteString(`%`)
		case "deadline":
			if task.Deadline.After(protocol.ZEROTIME) {
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
				common_printIcon(data, "task", "confirmStoryChange", "taskID="+strconv.Itoa(int(task.Id)), nil, "list", "", "hiddenwin", "btn-wide")
				break
			}
			if task.Status != "pause" {
				common_printIcon(data, "task", "start", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "", "", "iframe", "true")
			} else {
				common_printIcon(data, "task", "restart", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "", "", "iframe", "true")
			}

			common_printIcon(data, "task", "close", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "", "", "iframe", "true")
			common_printIcon(data, "task", "finish", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "", "", "iframe", "true")

			common_printIcon(data, "task", "recordEstimate", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "time", "", "iframe", "true")
			common_printIcon(data, "task", "edit", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "")
			common_printIcon(data, "task", "examine", "taskID="+strconv.Itoa(int(task.Id)), task, "list", "", "", "iframe", "true")
			batchCreateDesc := data.Lang["task"]["children"].(string)
			if task.Parent > 0 {
				batchCreateDesc = data.Lang["task"]["grandchildren"].(string)
			}
			common_printIcon(data, "task", "batchCreate", "project="+strconv.Itoa(int(task.Project))+"&storyID="+strconv.Itoa(int(task.Story))+"&moduleID="+strconv.Itoa(int(task.Module))+"&taskID="+strconv.Itoa(int(task.Id))+"&ifame=0", task, "list", "treemap-alt", "", "", "", "", batchCreateDesc)

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
	if task.AssignedTo == -1 {
		btnClass += " disabled"
	}
	return html_a(createLink("task", "assignTo", []interface{}{"projectID=", task.Project, "&taskID=", task.Id, true}), "<i class='icon icon-hand-right'></i> <span class='"+btnTextClass+"'>"+assignedToText+"</span>", "", "class='"+btnClass+"'")

}
