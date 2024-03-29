package handler

import (
	"fmt"
	"html/template"
	"jachunPM_http/setting"
	"libraries"
	"protocol"
	"reflect"
	"strconv"
	"strings"
)

const (
	classActive = "class='active'"
)

func createLink(moduleName string, methodName string, vars interface{}) string {
	return setting.Setting.Origin + protocol.CreateLink(moduleName, methodName, vars)
}
func htmlTemplateFuncs() {
	global_Funcs["helper_createLink"] = func(moduleName, methodName string, vars ...interface{}) string {

		return createLink(moduleName, methodName, vars)
	}

	global_Funcs["html_submitButton"] = func(label, class, misc string, data *TemplateData) template.HTML {
		if label == "" {
			label, _ = data.Lang["common"]["save"].(string)
		}
		if strings.Index(misc, "data-loading") == -1 {
			if loading, ok := data.Lang["common"]["loading"].(string); ok {
				misc += "data-loading='" + loading + "'"
			}

		}
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<button type='submit' id='submit' class='")
		buf.WriteString(class)
		buf.WriteString("' ")
		buf.WriteString(misc)
		buf.WriteByte('>')
		buf.WriteString(label)
		buf.WriteString("</button>")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
	global_Funcs["html_linkButton"] = func(data *TemplateData, label, link string, ext ...string) template.HTML { //class='btn', $misc = '', $target = 'self'
		if data.App["onlybody"].(bool) && data.Lang["common"]["goback"].(string) == label {
			return template.HTML("")
		}
		link = processOnlyBodyParam(data, link, false)
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<button type='button' class='")
		if len(ext) > 0 {
			buf.WriteString(ext[0])
		} else {
			buf.WriteString("btn")
		}
		buf.WriteString("' ")
		if len(ext) > 1 {
			buf.WriteString(ext[1])
			buf.WriteString(" ")
		}
		buf.WriteString("onclick='")
		if len(ext) > 2 {
			buf.WriteString(ext[2])
		} else {
			buf.WriteString("self")
		}
		buf.WriteString(".location.href=\"")
		buf.WriteString(link)
		buf.WriteString("\"'>")
		buf.WriteString(label)
		buf.WriteString("</button>")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
	global_Funcs["html_hidden"] = func(name string, value ...interface{}) template.HTML {
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<input type='hidden' name='")
		buf.WriteString(name)
		buf.WriteString("' id='")
		buf.WriteString(name)
		buf.WriteString("' value='")
		if len(value) > 0 {
			buf.WriteString(libraries.I2S(value[0]))
		}
		buf.WriteString("' ")
		if len(value) == 2 {
			buf.WriteString(libraries.I2S(value[1]))
		}
		buf.WriteString(" />\n")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
	global_Funcs["html_a"] = func(href string, value ...string) template.HTML {

		return template.HTML(html_a(href, value...))

	}

	global_Funcs["js_import"] = func(url string) template.HTML {

		mark := "?"
		if strings.Index(url, "?") > -1 {
			mark = "&"
		}
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<script src='")
		buf.WriteString(url)
		buf.WriteString(mark)
		buf.WriteString("v=")
		buf.WriteString(setting.Setting.Version)
		buf.WriteString("'></script>")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)

	}
	global_Funcs["css_import"] = func(url string, attrib ...string) template.HTML {

		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<link rel='stylesheet' href='")
		buf.WriteString(url)

		buf.WriteString("?v=")
		buf.WriteString(setting.Setting.Version)
		buf.WriteString("' type='text/css' media='screen'")
		if len(attrib) == 1 {
			buf.WriteString(" ")
			buf.WriteString(attrib[0])
		}
		buf.WriteString(" />")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)

	}
	global_Funcs["html_input"] = func(name string, value ...interface{}) template.HTML { // value  attrib
		var strValue []string
		for _, v := range value {
			strValue = append(strValue, libraries.I2S(v))
		}
		return template.HTML(html_input(name, strValue...))
	}

	global_Funcs["html_submitButton"] = func(data *TemplateData, value ...string) template.HTML { //label,class,misc

		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString(" <button type='submit' id='submit' class='")
		if len(value) > 1 {
			buf.WriteString(value[1]) //class
		} else {
			buf.WriteString("btn btn-primary")
		}
		buf.WriteString("' ")
		if len(value) > 2 {
			buf.WriteString(value[2]) //misc
		}
		if len(value) < 3 || strings.Index(value[2], "data-loading") == -1 {
			buf.WriteString(" data-loading='")
			buf.WriteString(data.Lang["common"]["loading"].(string))
			buf.WriteString(`'`)
		}
		buf.WriteByte('>')
		if len(value) > 0 {
			buf.WriteString(value[0]) //label
		} else {
			buf.WriteString(data.Lang["common"]["save"].(string))
		}
		buf.WriteString("</button>")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
	global_Funcs["html_backButton"] = func(data *TemplateData, value ...string) template.HTML { //label = '', misc = '', class = 'btn-wide',url
		if data.onlybody() {
			return template.HTML("")
		}
		buf := bufpool.Get().(*libraries.MsgBuffer)
		if len(value) > 3 {
			buf.WriteString("<a href='")
			buf.WriteString(value[3])
			buf.WriteString("' class='btn btn-back ")
		} else {
			buf.WriteString("<a href='javascript:history.go(-1);' class='btn btn-back ")
		}
		if len(value) > 2 {
			buf.WriteString(value[2]) //class
		} else {
			buf.WriteString("btn-wide")
		}
		buf.WriteString("' ")
		if len(value) > 1 {
			buf.WriteString(value[1]) //misc
		}
		buf.WriteByte('>')
		if len(value) > 0 && value[0] != "" {
			buf.WriteString(value[0]) //label
		} else {
			buf.WriteString(data.Lang["common"]["goback"].(string))
		}
		buf.WriteString("</a>")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
	global_Funcs["html_select"] = func(name string, options interface{}, selectedItem interface{}, attrib string, isappend ...bool) template.HTML {
		return template.HTML(html_select(name, options, selectedItem, attrib, isappend...))
	}
	global_Funcs["pager_show"] = func(data *TemplateData, align, typ string) template.HTML { //($align = 'right', $type = 'full')

		if typ == "pagerjs" {
			if data.Page.Total == 0 {
				return template.HTML("<div class='pull-right'>" + data.Lang["pager"]["noRecord"].(string) + "</div>")
			}
			var vars []protocol.HtmlKeyValueStr

			for key, value := range data.ws.GetAllQuery() {
				if key == "recPerPage" || key == "pageID" {
					continue
				}
				vars = append(vars, protocol.HtmlKeyValueStr{key, value[0]})
			}

			vars = append(vars, protocol.HtmlKeyValueStr{"recPerPage", "{recPerPage}"})
			vars = append(vars, protocol.HtmlKeyValueStr{"pageID", "{page}"})
			buf := bufpool.Get().(*libraries.MsgBuffer)
			buf.WriteString("<ul class='pager' data-page-cookie='")
			buf.WriteString(data.Page.CookieName)
			buf.WriteString("' data-ride='pager' data-rec-total='")
			buf.WriteString(strconv.Itoa(data.Page.Total))
			buf.WriteString("' data-rec-per-page='")
			buf.WriteString(strconv.Itoa(data.Page.PerPage))
			buf.WriteString("' data-page='")
			buf.WriteString(strconv.Itoa(data.Page.Page))
			buf.WriteString("' data-link-creator='")
			buf.WriteString(createLink(data.App["moduleName"].(string), data.App["methodName"].(string), vars))
			buf.WriteString("'></ul>")
			res := buf.String()
			buf.Reset()
			bufpool.Put(buf)
			return template.HTML(res)
		} else {
			//parent::show($align, $type);
		}
		return template.HTML("")
	}
	global_Funcs["html_commonButton"] = func(label string, value ...string) template.HTML { // $misc = '', $class = 'btn', $icon = '')
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<button type='button'")
		if len(value) > 1 {
			buf.WriteString(" class='")
			buf.WriteString(value[1]) //class
			buf.WriteString("'")
		} else {
			buf.WriteString(" class='btn'")
		}
		if len(value) > 0 {
			buf.WriteString(" ")
			buf.WriteString(value[0]) //misc
		}
		buf.WriteString(">")
		if len(value) > 2 {
			buf.WriteString("<i class='icon-")
			buf.WriteString(value[2]) //icon
			buf.WriteString("'></i>")
		}
		buf.WriteString(label)
		buf.WriteString("</button>")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
	global_Funcs["html_radio"] = func(name string, options []protocol.HtmlKeyValueStr, checked_i interface{}, value ...string) template.HTML { //( $attrib = '', $type = 'inline')
		if len(options) == 0 {
			return template.HTML("")
		}
		var isBlock bool
		if len(value) == 2 {
			isBlock = value[1] == "block"
		}
		checked := libraries.I2S(checked_i)
		buf := bufpool.Get().(*libraries.MsgBuffer)
		for _, option := range options {
			if isBlock {
				buf.WriteString("<div class='radio'><label>")
			} else {
				buf.WriteString("<label class='radio-inline'>")
			}
			buf.WriteString("<input type='radio' name='")
			buf.WriteString(name)
			buf.WriteString("' value='")
			buf.WriteString(option.Key)
			buf.WriteString("' ")
			if option.Key == checked {
				buf.WriteString(" checked ='checked' ")
			}
			if len(value) > 0 {
				buf.WriteString(value[0])
			}
			buf.WriteString(" id='")
			buf.WriteString(name)
			buf.WriteString(option.Key)
			buf.WriteString("' /> ")
			buf.WriteString(option.Value)
			if isBlock {
				buf.WriteString("</label></div>")
			} else {
				buf.WriteString("</label>")
			}
		}
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
	global_Funcs["html_password"] = func(name string, value ...string) template.HTML { //($value = "", $attrib = "")
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<input type='password' name='")
		buf.WriteString(name)
		buf.WriteString("' id='")
		buf.WriteString(name)
		if len(value) > 0 {
			buf.WriteString("' value='")
			buf.WriteString(value[0])
		}
		buf.WriteString("' ")
		if len(value) > 1 {
			buf.WriteString(value[1])
		}
		buf.WriteString("/>\n")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)

	}
	global_Funcs["html_textarea"] = func(name string, value ...string) template.HTML { //($value = "", $attrib = "")
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<textarea name='")
		buf.WriteString(name)
		buf.WriteString("' id='")
		buf.WriteString(name)
		buf.WriteString("' ")
		if len(value) > 1 {
			buf.WriteString(value[1])
		}
		buf.WriteString(">")
		if len(value) > 0 {
			buf.WriteString(value[0])
		}
		buf.WriteString("</textarea>\n")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)

	}
	global_Funcs["html_checkbox"] = func(name string, options []protocol.HtmlKeyValueStr, value ...interface{}) template.HTML { //$checked = "", $attrib = "", $type = 'inline'){
		return template.HTML(html_checkbox(name, options, value...))
	}
	global_Funcs["bbcode2html"] = func(code string, param ...bool) interface{} { //参数1 是否显示图片，参数2是否输出template.HTML
		allowimgcode := true
		if len(param) > 0 {
			allowimgcode = param[0]
		}
		isHtml := true
		if len(param) > 1 {
			isHtml = param[1]
		}
		if isHtml {
			return template.HTML(libraries.Bbcode2html(code, true, false, false, false, allowimgcode, false))
		}
		return libraries.Bbcode2html(code, true, false, false, false, allowimgcode, false)
	}
	global_Funcs["html_icon"] = func(name string, classExt ...string) template.HTML {
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<i class='")
		buf.WriteString("icon-")
		buf.WriteString(name)
		if len(classExt) == 1 {
			buf.WriteString(" ")
			buf.WriteString(classExt[0])
		}
		buf.WriteString("'></i>")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)

	}
	global_Funcs["toHtml"] = func(i interface{}) template.HTML {
		return template.HTML(fmt.Sprint(i))
	}
	global_Funcs["tohtml"] = func(i interface{}) template.HTML {
		return template.HTML(fmt.Sprint(i))
	}
}
func hookTemplateFuncs() {
	global_Funcs["importHeaderHook"] = func(data *TemplateData) (res template.HTML) {
		templateLock.RLock()
		buf := bufpool.Get().(*libraries.MsgBuffer)
		defer func() {
			buf.Reset()
			bufpool.Put(buf)
			templateLock.RUnlock()
		}()
		moduleName := data.App["moduleName"].(string)
		if v, ok := data.Lang["menugroup"][moduleName]; ok {
			moduleName = v.(string)
		}
		name := "header." + moduleName + ".html"
		err := T.ExecuteTemplate(buf, name, data)
		if err == nil {
			res = template.HTML(buf.String())
		}
		buf.Reset()
		return res
	}
	global_Funcs["importFooterHook"] = func(data *TemplateData) (res template.HTML) {
		templateLock.RLock()
		buf := bufpool.Get().(*libraries.MsgBuffer)
		defer func() {
			buf.Reset()
			bufpool.Put(buf)
			templateLock.RUnlock()
		}()
		moduleName := data.App["moduleName"].(string)
		if v, ok := data.Lang["menugroup"][moduleName]; ok {
			moduleName = v.(string)
		}
		name := "footer." + moduleName + ".html"
		err := T.ExecuteTemplate(buf, name, data)
		if err == nil {
			res = template.HTML(buf.String())
		}
		buf.Reset()
		return res
	}
}
func html_a(href string, value ...string) string {
	buf := bufpool.Get().(*libraries.MsgBuffer)
	buf.WriteString("<a href='")
	buf.WriteString(href)
	buf.WriteString("' ")
	if len(value) > 2 {
		for i := 2; i < len(value); i++ {
			buf.WriteString(value[i])
		}
	}
	if len(value) > 1 && value[1] != "_self" && value[1] != "" {
		buf.WriteString(" target='" + value[1])
		buf.WriteString("'")
	}
	buf.WriteString(">")
	if len(value) > 0 {
		buf.WriteString(value[0])
	} else {
		buf.WriteString(href)
	}
	buf.WriteString("</a>")
	res := buf.String()
	buf.Reset()
	bufpool.Put(buf)
	return res
}
func html_input(name string, value ...string) string { // value  attrib
	buf := bufpool.Get().(*libraries.MsgBuffer)
	buf.WriteString("<input type='text' name='")
	buf.WriteString(name)
	buf.WriteString("' ")
	if len(value) < 2 || strings.Index(value[1], "id=") == -1 {
		buf.WriteString("id='")
		buf.WriteString(name)
		buf.WriteString("' ")
	}
	buf.WriteString("value='")
	if len(value) > 0 {
		buf.WriteString(strings.ReplaceAll(value[0], "'", "&#039;")) //value
	}
	buf.WriteString("' ")
	if len(value) > 1 {
		buf.WriteString(value[1]) //attrib
	}
	buf.WriteString(" />")
	res := buf.String()
	buf.Reset()
	bufpool.Put(buf)
	return res
}
func html_select(name string, optionsI interface{}, selectedItem interface{}, attrib string, isappend ...bool) string {
	var selectedItems []string
	r := reflect.ValueOf(selectedItem)
	if r.Kind() == reflect.Slice {
		for i := 0; i < r.Len(); i++ {
			selectedItems = append(selectedItems, libraries.I2S(r.Index(i).Interface()))
		}
	} else {
		selectedItems = []string{libraries.I2S(selectedItem)}
	}
	var options []protocol.HtmlKeyValueStr
	switch o := optionsI.(type) {
	case []protocol.HtmlKeyValueStr:
		options = o
	case []*protocol.MSG_PROJECT_project_cache:
		for _, project := range o {
			options = append(options, protocol.HtmlKeyValueStr{strconv.Itoa(int(project.Id)), project.Name})
		}
	case []*protocol.MSG_PROJECT_product_cache:
		for _, product := range o {
			options = append(options, protocol.HtmlKeyValueStr{strconv.Itoa(int(product.Id)), product.Name})
		}
	case []*protocol.MSG_PROJECT_TASK:
		for _, task := range o {
			options = append(options, protocol.HtmlKeyValueStr{strconv.Itoa(int(task.Id)), task.Name})
		}
	case []*protocol.MSG_USER_INFO_cache:
		for _, user := range o {
			name := user.Realname
			if name == "" {
				name = user.Account
			}
			options = append(options, protocol.HtmlKeyValueStr{strconv.Itoa(int(user.Id)), name})
		}
	default:
		return "html_select无法处理options类型" + reflect.TypeOf(optionsI).String()
	}
	if len(isappend) > 0 && isappend[0] {
		for _, item := range selectedItems {
			find := false
			for _, v := range options {
				if v.Key == item {
					find = true
					break
				}
			}
			if !find {
				options = append(options, protocol.HtmlKeyValueStr{item, item})
			}
		}

	}

	if len(options) == 0 {
		return ""
	}

	buf := bufpool.Get().(*libraries.MsgBuffer)
	buf.WriteString("<select name='")
	buf.WriteString(name)
	buf.WriteString("' ")
	if strings.Index(attrib, "id=") == -1 {
		buf.WriteString("id='")
		if strings.Index(name, "[") > -1 {
			buf.WriteString(strings.Trim(strings.ReplaceAll(strings.ReplaceAll(name, "[", ""), "]", ""), " "))
		} else {
			buf.WriteString(name)
		}
		buf.WriteString("'")
	}
	buf.WriteString(" ")
	buf.WriteString(attrib)
	buf.WriteString(">\n")
	for _, option := range options {
		key := strings.ReplaceAll(option.Key, "item", "")
		buf.WriteString("<option value='")
		buf.WriteString(key)
		buf.WriteString("' data-keys='")
		buf.WriteString(key)
		buf.WriteString("'")
		for _, v := range selectedItems {
			if key == v {
				buf.WriteString(" selected='selected'")
				break
			}
		}
		buf.WriteString(" title='")
		buf.WriteString(option.Value)
		buf.WriteString("'>")
		buf.WriteString(option.Value)
		buf.WriteString("</option>\n")
	}

	buf.WriteString("</select>\n")
	res := buf.String()
	buf.Reset()
	bufpool.Put(buf)
	return res
}
func isClickableFuncs() {
	global_Funcs["null_isClickable"] = func(data *TemplateData, obj interface{}, action string) bool { return true }
	global_Funcs["MSG_USER_INFO_cache_isClickable"] = func(data *TemplateData, obj interface{}, action string) bool {
		v := obj.(*protocol.MSG_USER_INFO_cache)
		lockMinutes, _ := data.Config["user"]["common"]["lockMinutes"].(int)
		if action == "unlock" && data.Time.Unix()-v.Locked >= int64(lockMinutes)*60 {
			return false
		}
		return true
	}
	global_Funcs["MSG_PROJECT_product_cache_isClickable"] = func(data *TemplateData, obj interface{}, action string) bool {
		v := obj.(*protocol.MSG_PROJECT_product_cache)
		if action == "close" {
			return v.Status != "closed"
		}
		return true
	}
	global_Funcs["MSG_PROJECT_product_cache_map_isClickable"] = func(data *TemplateData, obj interface{}, action string) bool {
		v := obj.(map[string]interface{})
		if action == "close" {
			return v["Status"].(string) != "closed"
		}
		return true
	}
	global_Funcs["MSG_PROJECT_TASK_isClickable"] = func(data *TemplateData, obj interface{}, action string) bool {
		task := obj.(*protocol.MSG_PROJECT_TASK)
		if len(task.Children) > 0 {
			if action == "start" || action == "recordEstimate" || action == "finish" || action == "cancel" || action == "pause" || action == "internalaudit" || action == "proofreading" || action == "activate" || action == "assignto" || action == "close" {
				return false
			}
		}

		switch action {
		case "start":
			return task.Status == "wait"
		case "examine":
			return task.Status != "wait"
		case "restart":
			return task.Status == "pause"
		case "pause", "internalaudit":
			return task.Status == "doing"
		case "assignto":
			return task.Status != "closed" && task.Status != "cancel"
		case "close":
			return task.Status == "done" || task.Status == "cancel"
		case "activate":
			return task.Status == "done" || task.Status == "closed" || task.Status == "cancel"
		case "proofreading":
			return task.Status == "done" && task.Finalfile
		case "finish":
			return task.Status != "done" && task.Status != "closed" && task.Status != "cancel"
		case "cancel":
			return task.Status != "done" && task.Status != "closed" && task.Status != "cancel"
		case "batchcreate":
			if task.Ancestor > 0 {
				return false
			}
		}

		return true
	}

}
func html_checkbox(name string, options []protocol.HtmlKeyValueStr, value ...interface{}) string { //$checked = "", $attrib = "", $type = 'inline')
	if len(options) == 0 {
		return ""
	}
	var checked []string
	if len(value) > 0 {
		r := reflect.ValueOf(value[0])
		if r.Kind() == reflect.Slice {
			for i := 0; i < r.Len(); i++ {
				checked = append(checked, libraries.I2S(r.Index(i).Interface()))
			}
		} else {
			checked = []string{libraries.I2S(value[0])}
		}
	}
	var isBlock bool
	if len(value) == 3 {
		isBlock = value[2].(string) == "block"
	}

	buf := bufpool.Get().(*libraries.MsgBuffer)
	for _, option := range options {
		key := strings.ReplaceAll(option.Key, "item", "")
		if isBlock {
			buf.WriteString("<div class='checkbox-primary'>")
		} else {
			buf.WriteString("<div class='checkbox-primary checkbox-inline'>")
		}
		buf.WriteString("<input type='checkbox' name='")
		buf.WriteString(name)
		buf.WriteString("' value='")
		buf.WriteString(key)
		buf.WriteString("' ")
		for _, c := range checked {
			if c == key {
				buf.WriteString("checked ='checked' ")
				break
			}
		}
		if len(value) > 1 {
			buf.WriteString(libraries.I2S(value[1])) //$attrib
		}
		buf.WriteString(" id='")
		buf.WriteString(name)
		buf.WriteString(key)
		buf.WriteString("' />")
		buf.WriteString("<label for='")
		buf.WriteString(name)
		buf.WriteString(key)
		buf.WriteString("'>")
		buf.WriteString(option.Value)
		buf.WriteString("</label></div>")
	}
	res := buf.String()
	buf.Reset()
	bufpool.Put(buf)
	return res
}
func processOnlyBodyParam(data *TemplateData, link string, onlyBody bool) string {
	if !onlyBody && !data.App["onlybody"].(bool) {
		return link
	}
	if !strings.Contains(link, "onlybody=yes") {
		if strings.Contains(link, "?") {
			link += "&onlybody=yes"
		} else {
			link += "?onlybody=yes"
		}
	}
	return link
}
