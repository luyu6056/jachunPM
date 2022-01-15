package handler

import (
	"html/template"
	"jachunPM_http/config"
	"jachunPM_http/js"
	"libraries"
	"protocol"
	"strconv"
	"strings"
)

func init() {
	httpHandlerMap["GET"]["/datatable/ajaxCustom"] = datatable_ajaxCustom
	httpHandlerMap["POST"]["/datatable/ajaxSave"] = datatable_ajaxSave
	httpHandlerMap["GET"]["/datatable/ajaxReset"] = datatable_ajaxReset

}
func datatable_getSetting(data *TemplateData, module string, method string) (setting []*config.ConfigDatatable) {
	datatableId := module
	if method != "" {
		datatableId += strings.ToUpper(method[:1]) + method[1:]
	}

	mode := "table"
	if v1, ok := data.Config["datatable"][datatableId]; ok {
		if v2, ok := v1["mode"].(string); ok {
			mode = v2
		}
	}

	key := "tablecols"
	if mode == "datatable" {
		key = "cols"
	}
	if v1, ok := data.Config["datatable"]["moduleAlias"]; ok {
		if v2, ok := v1[module+"-"+method].(string); ok {
			module = v2
		}
	}

	if v1, ok := data.Config["datatable"][datatableId]; ok {
		if v2, ok := v1[key].(string); ok {
			libraries.JsonUnmarshalStr(v2, &setting)

		}
	}
	fieldList := datatable_getFieldList(data, module)
	if len(setting) == 0 {

		if v1, ok := data.Config[module]["datatable"]; ok {
			if v2, ok := v1["defaultField"].([]string); ok {
				for k, value := range v2 {
					set := &config.ConfigDatatable{
						Id:    value,
						Order: k + 1,
						Show:  true,
						Width: fieldList[value]["width"],
						Fixed: fieldList[value]["fixed"],
						Title: fieldList[value]["title"],
						Sort:  "yes",
					}
					if v, ok := fieldList[value]["sort"]; ok {
						set.Sort = v
					}
					if v, ok := fieldList[value]["name"]; ok {
						set.Name = v
					}
					setting = append(setting, set)
				}
			}
		}

	} else {
		for i := len(setting) - 1; i >= 0; i-- {
			set := setting[i]

			if data.ws.Session().Load_str("currentProductType") == "normal" && set.Id == "branch" {
				setting = append(setting[:i], setting[i+1:]...)
				continue
			}
			if set.Id == "actions" {
				set.Width = fieldList[set.Id]["width"]
			}
			set.Title = fieldList[set.Id]["title"]
			if v, ok := fieldList[set.Id]["sort"]; ok {
				set.Sort = v
			}
		}
	}
	data.Data["mode"] = mode
	return setting
}
func datatable_sortCols(list []*config.ConfigDatatable) {
	f := func(a, b *config.ConfigDatatable) bool {
		return a.Order < b.Order
	}
	max_len := len(list)
	tmp := make([]*config.ConfigDatatable, max_len)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if f(list[i+1], list[i]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if f(list[i+2], list[i]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if f(list[i+3], list[i+1]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if f(list[i+2], list[i]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r int
	step = 4
	for step < max_len {
		step <<= 1
		for i := 0; i < max_len; i += step {
			l, r, max = i, i+step/2, i+step
			if max > max_len {
				max = max_len
			}
			for index := i; index < max; index++ {
				if l == step/2+i || (r < max && f(list[r], list[l])) {
					tmp[index] = list[r]
					r++
				} else {
					tmp[index] = list[l]
					l++
				}
			}
		}
		if step < max_len {
			for i := 0; i < max_len; i += step {
				l, r, max = i, i+step/2, i+step
				if max > max_len {
					max = max_len
				}
				for index := i; index < max; index++ {
					if l == step/2+i || (r < max && f(tmp[r], tmp[l])) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		} else {
			copy(list, tmp)
		}
	}

}

func datatable_setFixedFieldWidth(setting []*config.ConfigDatatable) map[string]int {
	widths := make(map[string]int)
	widths["leftWidth"] = 30
	widths["rightWidth"] = 0
	hasLeftAuto := false
	hasRightAuto := false
	for _, value := range setting {
		if value.Fixed != "no" {
			if value.Fixed == "left" && value.Width == "auto" {
				hasLeftAuto = true
			}
			if value.Fixed == "right" && value.Width == "auto" {
				hasRightAuto = true
			}
			widthKey := value.Fixed + "Width"
			add, _ := strconv.Atoi(strings.Trim(value.Width, "px"))
			widths[widthKey] += add
		}
	}
	if widths["leftWidth"] <= 550 && hasLeftAuto {
		widths["leftWidth"] = 550
	}
	if widths["rightWidth"] <= 0 && hasRightAuto {
		widths["rightWidth"] = 140
	}

	return widths

}
func datatableTemplateFuncs() {
	global_Funcs["datatable_printHead"] = func(data *TemplateData, col *config.ConfigDatatable, orderBy, vars string, checkBoxExt ...bool) template.HTML {
		checkBox := true
		if len(checkBoxExt) > 0 {
			checkBox = checkBoxExt[0]
		}
		buf := bufpool.Get().(*libraries.MsgBuffer)
		if col.Show {
			buf.WriteString("<th data-flex='")
			if col.Fixed == "no" {
				buf.WriteString("true")
			} else {
				buf.WriteString("false")
			}
			buf.WriteString("' ")
			if i, err := strconv.Atoi(col.Width); err == nil {
				if col.Id == "id" && i < 90 {
					buf.WriteString("data-width='90px' style='width:90px'")
				} else {
					buf.WriteString("data-width='")
					buf.WriteString(col.Width)
					buf.WriteString("px' style='width:")
					buf.WriteString(col.Width)
					buf.WriteString("px'")
				}

			} else {

				buf.WriteString("data-width='")
				buf.WriteString(col.Width)
				buf.WriteString("' style='width:")
				buf.WriteString(col.Width)
				buf.WriteString("'")
			}
			buf.WriteString("  class='c-")
			buf.WriteString(col.Id)
			buf.WriteString("")
			if col.Id == "actions" {
				buf.WriteString(" text-center")
			}
			buf.WriteString("' ")
			if col.Name != "" {
				buf.WriteString("title='")
				buf.WriteString(col.Name)
				buf.WriteString("'")
			} else if col.Title != "" {
				buf.WriteString("title='")
				buf.WriteString(col.Title)
				buf.WriteString("'")
			}
			buf.WriteString(">")
			if col.Id == "actions" {
				buf.WriteString(data.Lang["common"]["actions"].(string))
			} else if col.Sort == "no" {
				buf.WriteString(col.Title)
			} else {
				if col.Id == "id" && checkBox {
					buf.WriteString("<div class='checkbox-primary check-all' title='")
					buf.WriteString(data.Lang["common"]["selectAll"].(string))
					buf.WriteString("'><label></label></div>")
				}
				buf.WriteString(common_printOrderLink(data, col.Id, orderBy, vars, col.Title))
			}
			buf.WriteString("</th>")
		}
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
}
func datatable_ajaxCustom(data *TemplateData) (err error) {
	module := data.ws.Query("module")
	method := data.ws.Query("method")
	target := module + strings.ToUpper(method[0:1]) + method[1:]
	mode := "table"
	if v, ok := data.Config["datatable"][target]; ok {
		if m, ok := v["mode"].(string); ok {
			mode = m
		}
	}

	key := "tablecols"
	if mode == "datatable" {
		key = "cols"
	}

	if module == "testtask" {

		data.Config["testcase"]["datatable"]["defaultField"] = data.Config["testtask"]["datatable"]["defaultField"]
		data.Config["testcase"]["datatable"]["fieldList"] = map[string]map[string]string{
			"actions": map[string]string{
				"width": "100",
			},
		}
	}
	if module == "testcase" {
		delete(data.Config["testcase"]["datatable"]["fieldList"].(map[string]map[string]string), "assignedTo")

	}

	data.Data["module"] = module
	data.Data["method"] = method
	data.Data["mode"] = mode
	if v, ok := data.Config["datatable"]["moduleAlias"][module+"-"+method].(string); ok {
		module = v
	}

	setting, _ := data.Config["datatable"][target][key].(string)
	if setting == "" {
		setting = libraries.JsonMarshalToString(data.Config[module]["datatable"]["defaultField"])
	}

	data.Data["cols"] = datatable_getFieldList(data, module)
	data.Data["setting"] = setting
	templateOut("datatable.ajaxcustom.html", data)
	return nil
}
func datatable_getFieldList(data *TemplateData, module string) map[string]map[string]string {
	if fieldList, ok := data.Config[module]["datatable"]["fieldList"].(map[string]map[string]string); ok {
		if data.ws.Session().Load_str("currentProductType") == "normal" {
			delete(fieldList, "branch")
		}
		for field, items := range fieldList {
			if field == "branch" {
				if data.ws.Session().Load_str("currentProductType") == "branch" {
					items["title"] = data.Lang["datatable"]["branch"].(string)
				}
				if data.ws.Session().Load_str("currentProductType") == "platform" {
					items["title"] = data.Lang["datatable"]["platform"].(string)
				}
			}
			var title string

			if v, ok := data.Lang[module]; ok {
				title, _ = v[items["title"]].(string)
			}
			if title == "" {
				title, _ = data.Lang["common"][items["title"]].(string)
			}
			items["title"] = title
			fieldList[field] = items
		}
		return fieldList
	}
	return nil

}
func datatable_ajaxSave(data *TemplateData) (err error) {
	if data.User.Id == 0 {
		data.ajaxResult(true, nil)
		return nil
	}
	out := protocol.GET_MSG_USER_config_save()
	out.Uid = data.User.Id
	out.Module = "datatable"
	out.Key = data.ws.Post("name")
	target := data.ws.Post("target")
	if i := strings.Index(target, "/"); i > 0 {
		target = target[:i] + strings.ToUpper(target[i+1:i+2]) + target[i+2:]
	}
	out.Section = target
	out.Type = "add"
	if data.ws.Post("global") == "1" {
		out.Uid = 0

	}
	out.Value = data.ws.Post("value")
	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		return
	}
	data.ajaxResult(true, nil)
	return nil
}
func datatable_ajaxReset(data *TemplateData) (err error) {
	if data.ws.Query("confirm") != "yes" {
		data.ws.WriteString(js.Confirm(data.Lang["datatable"]["confirmReset"].(string), createLink("datatable", "ajaxReset", []interface{}{"module=", data.ws.Query("module"), "&method=", data.ws.Query("method"), "&system=", data.ws.Query("system"), "&confirm=yes"}), ""))
		return
	}
	if data.User.Id == 0 {
		data.ws.WriteString(js.Reload("parent"))
		return nil
	}
	out := protocol.GET_MSG_USER_config_save()
	out.Uid = data.User.Id
	out.Module = "datatable"
	out.Key = "tablecols"
	target := data.ws.Query("method")
	target = data.ws.Query("module") + strings.ToUpper(target[0:1]) + target[1:]
	if v, ok := data.Config["datatable"][target]; ok {
		if m, ok := v["mode"].(string); ok {
			out.Key = m
		}
	}
	out.Section = target
	out.Type = "delete"
	if data.ws.Query("system") == "1" {
		out.Uid = 0

	}
	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		return
	}
	data.ws.WriteString(js.Reload("parent"))
	return nil
}
