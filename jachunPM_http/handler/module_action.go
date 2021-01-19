package handler

import (
	"fmt"
	"html"
	"html/template"
	"libraries"
	"protocol"
	"strings"
)

func actionModelFuncs() {

	global_Funcs["action_printAction"] = func(data *TemplateData, action *protocol.MSG_LOG_Ation, descExt ...string) template.HTML {
		var desc interface{}
		if action.ObjectType == "" || action.Action == "" {
			return template.HTML("")
		}
		objectType := action.ObjectType
		actionType := strings.ToLower(action.Action)

		if desc == nil {
			if objectTypeMap, ok1 := data.Lang[objectType]; ok1 {
				if actionTypeMap, ok2 := objectTypeMap["action"].(map[string]string); ok2 {
					desc = actionTypeMap[actionType]
				} else if actionTypeMap, ok3 := objectTypeMap["action"].(map[string]interface{}); ok3 {
					desc = actionTypeMap[actionType].(map[string]string)
				}

			}
			if desc == nil && data.Lang["action"]["desc"].(map[string]string)[actionType] != "" {
				desc = data.Lang["action"]["desc"].(map[string]string)[actionType]
			} else {
				if action.Extra != "" {
					desc = data.Lang["action"]["desc"].(map[string]string)["extra"]
				} else {
					desc = data.Lang["action"]["desc"].(map[string]string)["common"]
				}
			}
		}
		switch i := desc.(type) {
		case string:
			i = strings.ReplaceAll(i, "$date", action.Date.Format("01-02 15:04"))
			actor := HostConn.GetUserCacheById(action.ActorId)
			name := actor.Realname
			if name == "" {
				name = actor.Account
			}
			i = strings.ReplaceAll(i, "$actor", name)
			i = strings.ReplaceAll(i, "$extra", action.Extra)
			return template.HTML("i")
		case map[string]string:
			s := i["main"]
			s = strings.ReplaceAll(s, "$date", action.Date.Format("01-02 15:04"))
			actor := HostConn.GetUserCacheById(action.ActorId)
			name := actor.Realname
			if name == "" {
				name = actor.Account
			}
			s = strings.ReplaceAll(s, "$actor", name)
			extra := strings.ToLower(action.Extra)
			if extraMap, ok := data.Lang[objectType][i["extra"]].(map[string]string); ok {
				if replacestr, ok2 := extraMap[extra]; ok2 {
					return template.HTML(strings.ReplaceAll(s, "$extra", replacestr))
				}

			}
			return template.HTML(strings.ReplaceAll(s, "$extra", action.Extra))
		}
		return template.HTML("")
	}

	global_Funcs["action_printChanges"] = func(data *TemplateData, objectType string, histories []*protocol.MSG_LOG_History, canChangeTagExt ...bool) template.HTML {
		if len(histories) == 0 {
			return template.HTML("")
		}
		canChangeTag := true
		if len(canChangeTagExt) == 1 {
			canChangeTag = canChangeTagExt[0]
		}

		maxLength := 0
		historiesWithDiff := make([]*protocol.MSG_LOG_History, 0, len(histories))
		historiesWithoutDiff := make([]*protocol.MSG_LOG_History, 0, len(histories))

		for _, history := range histories {

			history.FieldLabel = history.Field
			if data.Lang[objectType] != nil {
				if str, ok := data.Lang[objectType][history.Field].(string); ok {
					history.FieldLabel = str
				}
			}
			if len(history.FieldLabel) > maxLength {
				maxLength = len(history.FieldLabel)
			}
			if history.Diff == "" {
				historiesWithoutDiff = append(historiesWithoutDiff, history)
			} else {
				historiesWithDiff = append(historiesWithDiff, history)
			}

		}

		for _, history := range append(historiesWithoutDiff, historiesWithDiff...) {

			history.FieldLabel += strings.Repeat(data.Lang["action"]["label"].(map[string]interface{})["space"].(string), len(history.FieldLabel)-maxLength)
			if history.Diff != "" {
				history.Diff = history_Diff_srp.Replace(history.Diff)
				if history.Field != "subversion" && history.Field != "git" {
					history.Diff = html.EscapeString(history.Diff)
				}

				history.Diff = history_Diff_Resrp.Replace(history.Diff)

				noTagDiff := ""
				if canChangeTag {
					noTagDiff, _ = libraries.Preg_replace(`/&lt;\/?([a-z][a-z0-9]*)[^\/]*\/?&gt;/Ui`, "", history.Diff)
				}
				return template.HTML(fmt.Sprintf(data.Lang["action"]["desc"].(map[string]string)["diff2"], history.FieldLabel, noTagDiff))
			} else {
				return template.HTML(fmt.Sprintf(data.Lang["action"]["desc"].(map[string]string)["diff1"], history.FieldLabel, history.Old, history.New))
			}
		}
		return template.HTML("")
	}
}

var history_Diff_srp = strings.NewReplacer(`<ins>`, `[ins]`, "</ins>", "[/ins]", "<del>", "[del]", "</del>", "[/del]")
var history_Diff_Resrp = strings.NewReplacer(`[ins]`, `<ins>`, "[/ins]", `</ins>`, "[del]", "<del>", "[/del]", "</del>", "\n", "<br>")
