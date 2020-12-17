package handler

import (
	"fmt"
	"html/template"
	"libraries"
	"protocol"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type moduleMenu struct {
	name   string
	text   string
	module string
	method string
	class  string
	vars   []protocol.HtmlKeyValueStr
	hidden bool
	alias  []string
}

func hasPriv(data *TemplateData, module, method string) bool {
	return true
}
func commonModelFuncs() {

	global_Funcs["strAdd"] = func(str ...interface{}) string {
		buf := bufpool.Get().(*libraries.MsgBuffer)
		for _, s := range str {
			buf.WriteString(fmt.Sprint(s))
		}
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return res
	}
	global_Funcs["string"] = func(i interface{}) string {
		return fmt.Sprint(i)
	}
	global_Funcs["common_hasPriv"] = func(data *TemplateData, module, method string) bool { return hasPriv(data, module, method) }
	//global_Funcs["common_printBreadMenu"] = func(server, name string) bool { return true }
	global_Funcs["loadConfig"] = func(data *TemplateData, server string) int { return 0 }
	global_Funcs["multiply"] = func(a, b interface{}) int64 {
		_a := reflect.ValueOf(a)
		_b := reflect.ValueOf(b)
		return _a.Int() * _b.Int()
	}
	global_Funcs["helper_createLink"] = func(moduleName, methodName string, vars ...string) string {

		return createLink(moduleName, methodName, vars)
	}
	global_Funcs["inlink"] = func(data *TemplateData, methodName string, vars ...string) string {

		return createLink(data.App["moduleName"].(string), methodName, vars)
	}
	global_Funcs["time"] = func() string {
		return strconv.Itoa(int(time.Now().Unix()))
	}
	global_Funcs["strpos"] = func(s, substr string) int {
		return strings.Index(s, substr)
	}
	global_Funcs["strlen"] = func(s string) int {
		return len(s)
	}

	global_Funcs["getTemplateCss"] = func(name string) template.CSS {
		buf := bufpool.Get().(*libraries.MsgBuffer)
		templateLock.RLock()
		defer func() {
			templateLock.RUnlock()
			buf.Reset()
			bufpool.Put(buf)
		}()
		err := T.ExecuteTemplate(buf, strings.Replace(name, ".html", ".css", 1), nil)
		if err != nil {
			libraries.DebugLog("加载%s的css失败,%v", name, err)
		} else {
			res := buf.String()

			return template.CSS(res)
		}
		return template.CSS("")
	}
	global_Funcs["getTemplateJs"] = func(name string) template.JS {
		buf := bufpool.Get().(*libraries.MsgBuffer)
		templateLock.RLock()
		defer func() {
			templateLock.RUnlock()
			buf.Reset()
			bufpool.Put(buf)
		}()
		err := T.ExecuteTemplate(buf, strings.Replace(name, ".html", ".js", 1), nil)
		if err != nil {
			libraries.DebugLog("加载%s的js失败,%v", name, err)
		} else {
			res := buf.String()

			return template.JS(res)
		}
		return template.JS("")
	}
	global_Funcs["printMainmenu"] = func(data *TemplateData) template.HTML {

		moduleName := data.App["moduleName"].(string)
		if v, ok := data.Lang["menugroup"][moduleName]; ok {
			moduleName = v.(string)
		}

		menu := getModuleMenu("common", data)

		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<ul class='nav nav-default'>\n")
		for key, menuItem := range menu {
			if !menuItem.hidden {
				buf.WriteString("<li ")
				if menuItem.name == moduleName {
					buf.WriteString(classActive)
				}
				buf.WriteString("data-id='")
				buf.WriteString(menuItem.name)
				buf.WriteString("'><a href='")
				buf.WriteString(createLink(menuItem.module, menuItem.method, menuItem.vars))
				buf.WriteString("' ")
				if menuItem.name == moduleName {
					buf.WriteString(classActive)
				}
				buf.WriteString(">")
				buf.WriteString(menuItem.text)
				buf.WriteString("</a></li>\n")
				if key != len(menu)-1 {
					for _, v := range data.Lang["common"]["dividerMenu"].([]string) {
						if v == menuItem.name {
							buf.WriteString("<li class='divider'></li>")
							break
						}

					}
				}
			}
		}
		buf.WriteString("</ul>\n")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
	global_Funcs["getValue"] = func(i, key interface{}) interface{} {
		switch r := i.(type) {
		case []protocol.HtmlKeyValueStr:
			if k, ok := key.(string); ok {
				for _, v := range r {
					if v.Key == k {
						return v.Value
					}
				}
				return nil
			}
		}
		r := reflect.ValueOf(i)
		k := reflect.ValueOf(key)
		if r.Type().Kind() == reflect.Map {
			value := r.MapIndex(k)
			if value.Kind() == reflect.Invalid {
				return nil
			}
			return value.Interface()
		} else if r.Type().Kind() == reflect.Slice {
			value := r.Index(int(k.Int()))
			if value.Kind() == reflect.Invalid {
				return nil
			}
			return value.Interface()
		}
		return nil
	}
	global_Funcs["substr"] = func(str string, start, end int) string {
		return str[start:end]
	}
	global_Funcs["common_printModuleMenu"] = func(data *TemplateData, moduleName string) template.HTML {

		if data.Lang[moduleName]["menu"] == nil {
			return template.HTML("<ul></ul>")
		}

		/* get current module and method. */
		currentModule := data.App["moduleName"]
		currentMethod := data.App["methodName"]
		menu := getModuleMenu(moduleName, data)
		isMobile := false

		buf := bufpool.Get().(*libraries.MsgBuffer)
		/* The beginning of the menu. */
		if !isMobile {
			buf.WriteString("<ul class='nav nav-default'>\n")
		}

		if name, ok := data.Lang["menugroup"][moduleName].(string); ok {
			moduleName = name
		}

		/* Cycling to print every sub menu. */
		for _, menuItem := range menu {
			if menuItem.hidden {
				continue
			}
			if isMobile && menuItem.name == "" {
				continue
			}
			for _, v := range data.Lang["common"]["dividerMenu"].([]string) {
				if v == menuItem.name {
					buf.WriteString("<li class='divider'></li>")
					break
				}
			}

			active := ``
			/*for _, s := range menuItem.subModule {
				if s == currentModule {
					active = `active`
					break
				}
			}*/

			if moduleName == currentModule {
				for _, a := range menuItem.alias {
					if a == currentModule {
						active = `active`
						break
					}
				}
			}

			if active == "" && menuItem.module == currentModule {
				if menuItem.method == currentMethod {
					active = `active`
				} else {
					for _, a := range menuItem.alias {
						if a == currentMethod {
							active = `active`
							break
						}
					}
				}
			}

			if isMobile {
				buf.WriteString(html_a(createLink(menuItem.module, menuItem.method, menuItem.vars), menuItem.text, "", "class='"+menuItem.class+" "+active+"'"))
				buf.WriteString("\n")
			} else {
				buf.WriteString("<li class='")
				buf.WriteString(menuItem.class)
				buf.WriteByte(' ')
				buf.WriteString(active)
				buf.WriteString("' data-id='")
				buf.WriteString(menuItem.name)
				buf.WriteString("'>")
				buf.WriteString(html_a(createLink(menuItem.module, menuItem.method, menuItem.vars), menuItem.text, ""))
				//buf.WriteString(subMenu)
				buf.WriteString("</li>\n")
			}
		}

		if !isMobile {
			buf.WriteString("</ul>\n")
		}
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)

	}
	global_Funcs["cookie"] = func(data *TemplateData, key string) string {
		return data.ws.Cookie(key)
	}
	global_Funcs["common_printLink"] = func(data *TemplateData, module, method string, v ...string) template.HTML {
		var (
			vars   string
			label  string
			target string
			misc   string
		)
		if len(v) > 0 {
			vars = v[0]
		}
		if len(v) > 1 {
			label = v[1]
		}
		if len(v) > 2 {
			target = v[2]
		}
		if len(v) > 3 {
			misc = v[3]
		}
		if !hasPriv(data, module, method) {
			return template.HTML("")
		}

		return template.HTML(html_a(createLink(module, method, vars), label, target, misc))
	}
	global_Funcs["common_printOrderLink"] = func(data *TemplateData, fieldName, orderBy, vars, label string, moduleMethod ...string) template.HTML {
		module := data.App["moduleName"].(string)
		method := data.App["methodName"].(string)
		if len(moduleMethod) == 2 {
			module = moduleMethod[0]
			method = moduleMethod[1]
		}
		className := "header"
		order := strings.Split(orderBy, "_")
		order[0] = strings.Trim(order[0], "`")
		if order[0] == fieldName {
			if len(order) > 1 && order[1] == "asc" {
				orderBy = order[0] + "_desc"
				className = "sort-up"
			} else {
				orderBy = order[0] + "_asc"
				className = "sort-down"
			}
		} else {
			orderBy = strings.Trim(fieldName, "`") + "_asc"
			className = "header"
		}
		link := createLink(module, method, fmt.Sprintf(vars, orderBy))

		return template.HTML(html_a(link, label, "", "class='"+className+"'"))
	}
	global_Funcs["appendKeyValueStr"] = func(strs ...string) (res []protocol.HtmlKeyValueStr) {
		for i := 0; i < len(strs); i += 2 {
			if len(strs) > i+1 {
				res = append(res, protocol.HtmlKeyValueStr{strs[i], strs[i+1]})
			}
		}
		return
	}
	global_Funcs["date"] = func(layout string, timestamp int64) string {
		return time.Unix(timestamp, 0).Format(layout)
	}
	global_Funcs["genlist"] = func(n interface{}) []string {
		num, _ := strconv.Atoi(fmt.Sprint(n))
		ret := make([]string, num)
		for i := 0; i < num; i++ {
			ret[i] = strconv.Itoa(i)
		}
		return ret
	}
	global_Funcs["jsonMarshal"] = func(i interface{}) string {
		return libraries.JsonMarshalToString(i)
	}
	global_Funcs["appendStr"] = func(strs ...interface{}) (res []string) {
		for _, s := range strs {
			res = append(res, fmt.Sprint(s))
		}
		return
	}
	global_Funcs["common_printIcon"] = func(data *TemplateData, module, method, vars string, object interface{}, typ, icon string, extvalue ...string) template.HTML { //($target = '', $extraClass = '', $onlyBody = false, $misc = '', $title = '')
		target := ""
		extraClass := ""
		misc := ""
		title := ""
		if len(extvalue) > 0 {
			target = extvalue[0]
			if len(extvalue) > 1 {
				if data.ws.Query("isonlybody") == "yes" {
					if strings.Index(extvalue[1], "showinonlybody") == -1 {
						return template.HTML("")
					}
					extvalue[1] = strings.ReplaceAll(extvalue[1], "iframe", "")
				}
				extraClass = extvalue[1]
			}
			if len(extvalue) > 2 {
				misc = extvalue[2]
			}
			if len(extvalue) > 3 {
				title = extvalue[3]
			}
		}

		switch module {
		case "story":
			if method == "createcase" {
				module = "testcase"
				method = "create"
			}
		case "bug":
			if method == "tostory" {
				module = "story"
				method = "create"
			}
			if method == "createcase" {
				module = "testcase"
				method = "create"
			}
		}
		if !hasPriv(data, module, method) {
			return template.HTML("")
		}
		clickable := true
		if object != nil {
			if f_interface, ok := global_Funcs[reflect.ValueOf(object).Elem().Type().Name()+"_isClickable"]; ok {
				if f, ok := f_interface.(func(*TemplateData, interface{}, string) bool); ok {
					clickable = f(data, object, method)
				}
			}
		}
		link := createLink(module, method, vars)
		if title == "" {
			title = method
			if method == "create" && icon == "copy" {
				method = "copy"
			}
			if icon == "toStory" {
				title = data.Lang["bug"]["toStory"].(string)
			} else if icon == "createBug" {
				title = data.Lang["testtask"]["createBug"].(string)
			} else {
				if str, ok := data.Lang["common"][method].(string); ok {
					title = str
				}
				if str, ok := data.Lang[module][method].(string); ok {
					title = str
				}
			}
		}

		if icon == "" {
			if v, ok := data.Lang["common"]["icons"].(map[string]string)[method]; ok {
				icon = v
			} else {
				icon = method
			}
		}
		for _, v := range []string{"edit", "copy", "report", "export", "delete"} {
			if v == method {
				module = "common"
			}

		}
		class := fmt.Sprintf("icon-%s-%s", module, method)
		if !clickable {
			class += " disabled"
		}
		if icon != "" {
			class += " icon-" + icon
		}
		if clickable {

			if typ == "button" {
				if method != "edit" && method != "copy" && method != "delete" {
					return template.HTML(html_a(link, "<i class='"+class+"'></i> "+"<span class='text'>"+title+"</span>", target, "class='btn btn-link "+extraClass+"' "+misc))
				} else {
					return template.HTML(html_a(link, "<i class='"+class+"'></i>", target, "class='btn btn-link "+extraClass+"' title='"+title+"' "+misc))
				}
			} else {
				return template.HTML(html_a(link, "<i class='"+class+"'></i>", target, "class='btn "+extraClass+"' title='"+title+"' "+misc) + "\n")
			}
		} else {
			if typ == "list" {
				return template.HTML("<button type='button' class='disabled btn " + extraClass + "'><i class='" + class + "' title='" + title + "' " + misc + "></i></button>\n")
			}
		}

		return template.HTML("")
	}
}
func getModuleMenu(module string, data *TemplateData) (menu []moduleMenu) {

	if i, ok := data.Lang[module]["menu"]; ok {
		for _, v := range i.([]protocol.HtmlMenu) {
			l := strings.Split(v.Value["link"], "|")
			if len(l) > 2 {
				menuItem := moduleMenu{
					name:   v.Key,
					hidden: v.Value["hidden"] == "true",
					text:   l[0],
					module: l[1],
					method: l[2],
					class:  v.Value["class"],
				}
				if alias, ok := v.Value["alias"]; ok {
					menuItem.alias = strings.Split(alias, ",")
				}
				/*if subModule, ok := v.Value["subModule"]; ok {
					menuItem.subModule = strings.Split(subModule, ",")
				}*/
				if len(l) == 4 {
					for _, vars := range strings.Split(l[3], "&") {
						s := strings.Split(vars, "=")
						if len(s) == 2 {
							menuItem.vars = append(menuItem.vars, protocol.HtmlKeyValueStr{s[0], fmt.Sprintf(s[1], data.ws.Session().Load_str(s[0]))})
						}

					}
				}
				menu = append(menu, menuItem)
			}
		}
	}
	return
}
