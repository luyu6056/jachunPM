package handler

import (
	"html/template"
	"jachunPM_http/config"
	"libraries"
	"net/url"
	"protocol"
	"reflect"
	"runtime/debug"
	"strings"
	"sync"
)

const (
	classActive = "class='active'"
)

var global_Funcs template.FuncMap = map[string]interface{}{}
var bufpool = sync.Pool{New: func() interface{} {
	return new(libraries.MsgBuffer)
}}

func init() {
	loadFuncs()
}
func loadFuncs() {
	commonModelFuncs()
	htmlFuncs()
	hookFuncs()
	global_t.Funcs(global_Funcs)
}

func createLink(moduleName string, methodName string, vars interface{}) string {
	buf := bufpool.Get().(*libraries.MsgBuffer)
	buf.WriteString(config.Config.Origin)
	buf.WriteString("/")
	buf.WriteString(moduleName)
	buf.WriteString("/")
	buf.WriteString(methodName)
	switch v := vars.(type) {
	case []protocol.HtmlKeyValueStr:
		if len(v) > 0 {
			buf.WriteByte('?')
			for _, v := range v {
				buf.WriteString(url.QueryEscape(v.Key))
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(v.Value))
			}
		}
	case []string:
		if len(v) > 0 {
			buf.WriteByte('?')
			for _, s := range v {
				buf.WriteString(s)
			}

		}
	case string:
		buf.WriteByte('?')
		buf.WriteString(v)
	case nil:
	default:
		libraries.DebugLog("createLink不识别类型%s\r\n%s", reflect.TypeOf(v).String(), string(debug.Stack()))
	}

	res := buf.String()
	buf.Reset()
	bufpool.Put(buf)
	return res
}
func htmlFuncs() {

	global_Funcs["html_checkBox"] = func(name string, options []protocol.HtmlKeyValueStr, checked string) template.HTML {
		buf := bufpool.Get().(*libraries.MsgBuffer)

		checked = "," + checked + ","
		for _, option := range options {
			key := strings.ReplaceAll(option.Key, "item", "")
			value := option.Value
			buf.WriteString("<div class='checkbox-primary'>")
			buf.WriteString("<input type='checkbox' name='{")
			buf.WriteString(name)
			buf.WriteString("}[]' value='")
			buf.WriteString(key)
			buf.WriteString("' ")
			if strings.Index(checked, ","+key+",") > -1 {
				buf.WriteString(" checked ='checked'")
			}
			buf.WriteString(" id='")
			buf.WriteString(name)
			buf.WriteString(key)
			buf.WriteString("' /> ")
			buf.WriteString("<label for='")
			buf.WriteString(name)
			buf.WriteString(key)
			buf.WriteString("'>")
			buf.WriteString(value)
			buf.WriteString("</label></div>")

		}
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
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
	global_Funcs["html_linkButton"] = func(data *TemplateData, server string) string { return "待处理html_linkButton" }
	global_Funcs["html_hidden"] = func(name string, value ...string) template.HTML {
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<input type='hidden' name='")
		buf.WriteString(name)
		buf.WriteString("' id='")
		buf.WriteString(name)
		buf.WriteString("' value='")
		if len(value) > 0 {
			buf.WriteString(value[0])
		}
		buf.WriteString("' ")
		if len(value) == 2 {
			buf.WriteString(value[1])
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
		buf.WriteString(config.Config.Version)
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
		buf.WriteString(config.Config.Version)
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
	global_Funcs["html_input"] = func(name string, value ...string) template.HTML { // value  attrib
		return template.HTML(html_input(name, value...))
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
	global_Funcs["html_backButton"] = func(data *TemplateData, value ...string) template.HTML { //label = '', misc = '', class = 'btn-wide'
		if data.onlybody() {
			return template.HTML("")
		}
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<a href='javascript:history.go(-1);' class='btn btn-back ")
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
		if len(value) > 0 {
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
	global_Funcs["html_select"] = func(name string, options []protocol.HtmlKeyValueStr, selectedItems []string, attrib string, isappend ...bool) template.HTML {

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
			return template.HTML("")
		}

		buf := bufpool.Get().(*libraries.MsgBuffer)
		/* The begin. */

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
			buf.WriteString("'")
			for _, v := range selectedItems {
				if key == v {
					buf.WriteString(" selected='selected'")
					break
				}
			}
			buf.WriteString(">")
			buf.WriteString(option.Value)
			buf.WriteString("</option>\n")
		}

		buf.WriteString("</select>\n")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
}
func hookFuncs() {
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
	if len(value) == 2 {
		buf.WriteString(value[1])
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
