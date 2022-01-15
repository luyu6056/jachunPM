package handler

import (
	"fmt"
	"html"
	"html/template"
	"libraries"
	"protocol"
	"reflect"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/luyu6056/cache"
)

type moduleMenu struct {
	Name      string
	Text      string
	Module    string
	Method    string
	Class     string
	Vars      []protocol.HtmlKeyValueStr
	Hidden    bool
	Alias     []string
	SubModule []string
	Link      map[string]string
}
type commonPreAndNext struct {
	pre  interface{}
	next interface{}
}

var uuidSeed uint64
var commoncache = cache.Hget("common", "global")

func hasPriv(data *TemplateData, module, method string, obj ...interface{}) bool {

	if data.User.IsAdmin {
		return true
	}
	if len(obj) > 0 && obj[0] != nil && !hasPrivObj(obj[0], data.User) {
		return false
	}
	menu := module
	if v, ok := data.Lang["menugroup"][module].(string); ok {
		menu = v
	}
	if menu != "qa" && (data.Lang[menu] == nil || data.Lang[menu]["menu"] == nil) {
		return true
	} else if menu == "my" || menu == "index" || module == "true" {
		return true
	}
	if data.User.Priv[module] == nil {
		return false
	}
	return data.User.Priv[module][method]
}

type eface struct {
	typ, val unsafe.Pointer
}

func hasPrivObj(obj interface{}, user *protocol.MSG_USER_INFO_cache) bool {
	uid := user.Id
	ptr := (*eface)(unsafe.Pointer(&obj)).val
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
		t = t.Elem()
		if f, ok := t.FieldByName("OpenedBy"); ok {
			if *(*int32)(unsafe.Pointer(uintptr(ptr) + f.Offset)) == uid {
				return true
			}
		}
		if f, ok := t.FieldByName("AddedBy"); ok {
			if *(*int32)(unsafe.Pointer(uintptr(ptr) + f.Offset)) == uid {
				return true
			}
		}
		if f, ok := t.FieldByName("Uid"); ok {
			if *(*int32)(unsafe.Pointer(uintptr(ptr) + f.Offset)) == uid {
				return true
			}
		}
		if f, ok := t.FieldByName("AssignedTo"); ok {
			if *(*int32)(unsafe.Pointer(uintptr(ptr) + f.Offset)) == uid {
				return true
			}
		}
		if f, ok := t.FieldByName("FinishedBy"); ok {
			if *(*int32)(unsafe.Pointer(uintptr(ptr) + f.Offset)) == uid {
				return true
			}
		}
		if f, ok := t.FieldByName("CanceledBy"); ok {
			if *(*int32)(unsafe.Pointer(uintptr(ptr) + f.Offset)) == uid {
				return true
			}
		}
		if f, ok := t.FieldByName("ClosedBy"); ok {
			if *(*int32)(unsafe.Pointer(uintptr(ptr) + f.Offset)) == uid {
				return true
			}
		}
		if f, ok := t.FieldByName("LastEditedBy"); ok {
			if *(*int32)(unsafe.Pointer(uintptr(ptr) + f.Offset)) == uid {
				return true
			}
		}
	} else if t.Kind() == reflect.Map {
		v := reflect.ValueOf(obj)
		r := v.MapRange()
		for r.Next() {
			switch r.Key().String() {
			case "Project":
				if user.LimitedProjects[(r.Value().Interface().(int32))] {
					return false
				}
			default:
				libraries.ReleaseLog("hasPrivObj未处理map的key%v,值%+v", r.Key(), r.Value())
			}
		}
	} else {
		panic("啥玩意")
		libraries.ReleaseLog("hasPrivObj未处理类型%v,值%+v", t.Kind(), obj)
	}
	return true
}
func commonTemplateFuncs() {
	global_Funcs["log"] = func(i interface{}) string {
		libraries.DebugLog("%+v", i)
		return fmt.Sprint(i)
	}
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
		return libraries.I2S(i)
	}
	global_Funcs["common_hasPriv"] = func(data *TemplateData, module, method string, obj ...interface{}) bool {
		return hasPriv(data, module, method, obj...)
	}
	//global_Funcs["common_printBreadMenu"] = func(server, name string) bool { return true }
	global_Funcs["loadConfig"] = func(data *TemplateData, server string) int { return 0 }
	global_Funcs["multiply"] = func(a, b interface{}) int64 {
		_a := reflect.ValueOf(a)
		_b := reflect.ValueOf(b)
		return _a.Int() * _b.Int()
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

	global_Funcs["getTemplateCss"] = func(data *TemplateData, name string) template.CSS {
		buf := bufpool.Get().(*libraries.MsgBuffer)
		templateLock.RLock()
		defer func() {
			templateLock.RUnlock()
			buf.Reset()
			bufpool.Put(buf)
		}()
		s := strings.Split(name, ".")
		T.ExecuteTemplate(buf, s[0]+".common.css", nil)
		T.ExecuteTemplate(buf, strings.Replace(name, ".html", ".css", 1), nil)
		T.ExecuteTemplate(buf, strings.Replace(name, ".html", "."+data.App["ClientLang"].(string)+".css", 1), nil)
		return template.CSS(buf.String())
	}
	global_Funcs["getTemplateJs"] = func(name string) template.JS {
		buf := bufpool.Get().(*libraries.MsgBuffer)
		templateLock.RLock()
		defer func() {
			templateLock.RUnlock()
			buf.Reset()
			bufpool.Put(buf)
		}()
		s := strings.Split(name, ".")
		T.ExecuteTemplate(buf, s[0]+".common.js", nil)
		err := T.ExecuteTemplate(buf, strings.Replace(name, ".html", ".js", 1), nil)
		if err != nil {
			libraries.DebugLog("加载%s的js失败,%v", name, err)
		}
		return template.JS(html.UnescapeString(buf.String()))
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
			if !menuItem.Hidden {
				buf.WriteString("<li ")
				if menuItem.Name == moduleName {
					buf.WriteString(classActive)
				}
				buf.WriteString("data-id='")
				buf.WriteString(menuItem.Name)
				buf.WriteString("'><a href='")
				buf.WriteString(createLink(menuItem.Module, menuItem.Method, menuItem.Vars))
				buf.WriteString("' ")
				if menuItem.Name == moduleName {
					buf.WriteString(classActive)
				}
				buf.WriteString(">")
				buf.WriteString(menuItem.Text)
				buf.WriteString("</a></li>\n")
				if key != len(menu)-1 {
					for _, v := range data.Lang["common"]["dividerMenu"].([]string) {
						if v == menuItem.Name {
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
		return common_getValue(i, key)
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
			if menuItem.Hidden {
				continue
			}
			if isMobile && menuItem.Name == "" {
				continue
			}
			if dividerMenu, ok := data.Lang[moduleName]["dividerMenu"].([]string); ok {
				for _, v := range dividerMenu {
					if v == menuItem.Name {
						buf.WriteString("<li class='divider'></li>")
						break
					}
				}
			}

			active := ``
			for _, s := range menuItem.SubModule {
				if s == currentModule {
					active = `active`
					break
				}
			}

			if moduleName == currentModule {
				for _, a := range menuItem.Alias {
					if a == currentModule {
						active = `active`
						break
					}
				}
			}

			if active == "" && menuItem.Module == currentModule {
				if menuItem.Method == currentMethod {
					active = `active`
				} else {
					for _, a := range menuItem.Alias {
						if a == currentMethod {
							active = `active`
							break
						}
					}
				}
			}

			if isMobile {
				buf.WriteString(html_a(createLink(menuItem.Module, menuItem.Method, menuItem.Vars), menuItem.Text, "", "class='"+menuItem.Class+" "+active+"'"))
				buf.WriteString("\n")
			} else {
				buf.WriteString("<li class='")
				buf.WriteString(menuItem.Class)
				buf.WriteByte(' ')
				buf.WriteString(active)
				buf.WriteString("' data-id='")
				buf.WriteString(menuItem.Name)
				buf.WriteString("'>")
				buf.WriteString(html_a(createLink(menuItem.Module, menuItem.Method, menuItem.Vars), menuItem.Text, ""))
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
		return template.HTML(common_printLink(data, module, method, v...))
	}
	global_Funcs["common_printOrderLink"] = func(data *TemplateData, fieldName, orderBy, vars, label string, moduleMethod ...string) template.HTML {
		return template.HTML(common_printOrderLink(data, fieldName, orderBy, vars, label, moduleMethod...))
	}
	global_Funcs["appendKeyValueStr"] = func(strs ...interface{}) (res []protocol.HtmlKeyValueStr) {
		for i := 0; i < len(strs); i += 2 {
			if len(strs) > i+1 {
				res = append(res, protocol.HtmlKeyValueStr{libraries.I2S(strs[i]), libraries.I2S(strs[i+1])})
			}
		}
		return
	}
	global_Funcs["mergeKeyValueStr"] = func(kvstrs ...[]protocol.HtmlKeyValueStr) (res []protocol.HtmlKeyValueStr) {
		if len(kvstrs) > 0 {
			res = make([]protocol.HtmlKeyValueStr, len(kvstrs[0]))
			copy(res, kvstrs[0])
			for i := 1; i < len(kvstrs); i++ {
				res = append(res, kvstrs[i]...)
			}
		}
		return res
	}
	//格式化输出时间戳，允许不输入timestamp，则为当前时间.timestamp可以是time.Time和int int64 int32,string
	global_Funcs["date"] = func(layout string, timestamp ...interface{}) (res string) {
		var t time.Time
		if len(timestamp) > 0 {
			switch v := timestamp[0].(type) {
			case int:
				t = time.Unix(int64(v), 0)
			case int32:
				t = time.Unix(int64(v), 0)
			case int64:
				t = time.Unix(v, 0)
			case time.Time:
				t = v
			case string:
				if v == "" || v == "today" || v == "now" {
					t = time.Now()
				} else {
					i, _ := strconv.Atoi(v)
					t = time.Unix(int64(i), 0)
				}
			}

		} else {
			t = time.Now()
		}
		if len(timestamp) > 1 {
			if v, ok := timestamp[1].(string); ok && (strings.ToLower(v) == "normaltime" || strings.ToLower(v) == "normal") {
				if !t.After(protocol.NORMALTIME) {
					return ""
				}
			}
		}
		if layout == "Unix" {
			return strconv.FormatInt(t.Unix(), 10)
		}
		if layout == "today" {
			return t.Format("2006-01-02")
		}
		return t.Format(layout)
	}
	//num正值+负值-,如(5 -2)或(5,2,-1)，均可输出[5,4]
	global_Funcs["genlist"] = func(star, num interface{}, setpExt ...int) []int {
		n, _ := strconv.Atoi(fmt.Sprint(num))
		s, _ := strconv.Atoi(fmt.Sprint(star))
		setp := 1
		if len(setpExt) > 0 && setpExt[0] > 1 {
			setp = setpExt[0]
		}
		ret := make([]int, n)
		if n > 0 {
			for i := 0; i < n; i++ {
				ret[i] = s + i*setp
			}
		} else {
			for i := 0; i < n*-1; i++ {
				ret[i] = s - i*setp
			}
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
		return template.HTML(common_printIcon(data, module, method, vars, object, typ, icon, extvalue...))
	}
	global_Funcs["json_marshal"] = func(i interface{}) string {
		return libraries.JsonMarshalToString(i)
	}
	global_Funcs["str2js"] = func(s string) template.JS {
		return template.JS(s)
	}
	global_Funcs["strings_split"] = func(s, sep string) []string {
		return strings.Split(s, sep)
	}
	global_Funcs["rem"] = func(i, k int) int {
		return i % k
	}
	global_Funcs["fetch"] = func(oldData *TemplateData, module, method string, varstr ...string) template.HTML {
		return template.HTML(common_fetch(oldData, module, method, varstr...))
	}
	global_Funcs["generateUid"] = func() string {
		id := commoncache.INCRBY("generateUid", 1)
		return strconv.FormatUint(uint64(id), 10)
	}
	global_Funcs["common_printCommentIcon"] = func(data *TemplateData, commentFormLink string, object interface{}) template.HTML {

		if !hasPriv(data, "action", "comment", object) {
			return template.HTML("")
		}
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString(string(global_Funcs["html_commonButton"].(func(label string, value ...string) template.HTML)("<i class='icon icon-chat-line'></i> "+data.Lang["action"]["create"].(string), "", "btn btn-link pull-right btn-comment")))
		buf.WriteString(`<div class="modal fade modal-comment">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal"><i class="icon icon-close"></i></button>
        <h4 class="modal-title">`)
		buf.WriteString(data.Lang["action"]["create"].(string))
		buf.WriteString(`</h4>
      </div>
      <div class="modal-body">
        <form class="load-indicator" action="`)
		buf.WriteString(commentFormLink)
		buf.WriteString(`" target='hiddenwin' method='post'>
          <div class="form-group">
            <textarea id='comment' name='comment' class="form-control" rows="8" autofocus="autofocus"></textarea>
          </div>
          <div class="form-group form-actions text-center">
            <button type="submit" class="btn btn-primary btn-wide">`)
		buf.WriteString(data.Lang["common"]["save"].(string))
		buf.WriteString(`</button><button type="button" class="btn btn-wide" data-dismiss="modal">`)
		buf.WriteString(data.Lang["common"]["close"].(string))
		buf.WriteString("</button></div></form></div></div></div></div>")
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)
	}
	global_Funcs["intsum"] = func(a, b int) int {
		return a + b
	}
	global_Funcs["intsub"] = func(a, b int) int {
		return a - b
	}
	global_Funcs["strings_replace"] = func(a, b, c string) string {
		return strings.ReplaceAll(a, b, c)
	}
	global_Funcs["common_printBack"] = func(data *TemplateData, link string, classExt ...string) template.HTML {

		if data.onlybody() {
			return template.HTML("")
		}
		class := "btn"
		if len(classExt) == 1 {
			class = classExt[0]
		}

		title := data.Lang["common"]["goback"].(string) + data.Lang["common"]["backShortcutKey"].(string)
		return template.HTML(html_a(link, "<i class='icon-goback icon-back'></i> "+data.Lang["common"]["goback"].(string), "", "id='back' class='"+class+"' title='"+title+"'"))
	}
	global_Funcs["session"] = func(data *TemplateData, key string) string {
		return data.ws.Session().Load_str(key)
	}
	global_Funcs["printPreAndNext"] = func(data *TemplateData, ext ...interface{}) template.HTML { //preAndNext , linkTemplate
		return template.HTML("") //暂不做
		/*if data.onlybody() {
			return template.HTML("")
		}
		var preAndNext *commonPreAndNext
		var linkTemplate string
		if len(ext) > 0 {
			preAndNext, _ = ext[0].(*commonPreAndNext)
		}
		if len(ext) > 1 {
			linkTemplate, _ = ext[1].(string)
		}
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.WriteString("<nav class='container'>")
		if(isset($preAndNext->pre) and $preAndNext->pre)
		  {
		      $id = (isset($_SESSION['testcaseOnlyCondition']) and !$_SESSION['testcaseOnlyCondition'] and $app->getModuleName() == 'testcase' and isset($preAndNext->pre->case)) ? 'case' : 'id';
		      $title = isset($preAndNext->pre->title) ? $preAndNext->pre->title : $preAndNext->pre->name;
		      $title = '#' . $preAndNext->pre->$id . ' ' . $title . ' ' . $lang->preShortcutKey;
		      $link  = $linkTemplate ? sprintf($linkTemplate, $preAndNext->pre->$id) : inLink('view', "ID={$preAndNext->pre->$id}");
		      echo html::a($link, '<i class="icon-pre icon-chevron-left"></i>', '', "id='prevPage' class='btn' title='{$title}'");
		  }
		  if(isset($preAndNext->next) and $preAndNext->next)
		  {
		      $id = (isset($_SESSION['testcaseOnlyCondition']) and !$_SESSION['testcaseOnlyCondition'] and $app->getModuleName() == 'testcase' and isset($preAndNext->next->case)) ? 'case' : 'id';
		      $title = isset($preAndNext->next->title) ? $preAndNext->next->title : $preAndNext->next->name;
		      $title = '#' . $preAndNext->next->$id . ' ' . $title . ' ' . $lang->nextShortcutKey;
		      $link  = $linkTemplate ? sprintf($linkTemplate, $preAndNext->next->$id) : inLink('view', "ID={$preAndNext->next->$id}");
		      echo html::a($link, '<i class="icon-pre icon-chevron-right"></i>', '', "id='nextPage' class='btn' title='$title'");
		  }
		  echo '</nav>';
		res := buf.String()
		buf.Reset()
		bufpool.Put(buf)
		return template.HTML(res)*/
	}
	global_Funcs["helper_workDays"] = func(begin, end interface{}) template.HTML {
		return template.HTML("待处理")
	}
	global_Funcs["fileSize"] = func(size int64) string {
		if size < 1024 {
			return strconv.FormatInt(size, 10) + "B"
		} else if size < 1024*1024 {
			return fmt.Sprintf("%0.2fK", float64(size)/1024)
		} else if size < 1024*1024*1024 {
			return fmt.Sprintf("%0.2fM", float64(size)/1024/1024)
		}
		return fmt.Sprintf("%0.2fG", float64(size)/1024/1024/1024)
	}
	global_Funcs["sprintf"] = func(format string, a ...interface{}) string {
		return fmt.Sprintf(format, a...)
	}
	global_Funcs["join"] = func(elems []string, sep string) string {
		return strings.Join(elems, sep)
	}
	global_Funcs["cycle"] = func(data *TemplateData, a, b string) string {
		if data.Data["cycle_n"] == nil {
			data.Data["cycle_n"] = 0
		}
		data.Data["cycle_n"] = data.Data["cycle_n"].(int) + 1
		if data.Data["cycle_n"].(int)%2 == 1 {
			return a
		}
		return b
	}
	global_Funcs["uniqid"] = func() uint64 {
		id := atomic.AddUint64(&uuidSeed, 1)
		return id
	}
}

func getModuleMenu(module string, data *TemplateData) (menu []moduleMenu) {
	if v, ok := data.Data["getModuleMenu_"+module].([]moduleMenu); ok {
		return v
	}
	if i, ok := data.Lang[module]["menu"]; ok {
		for _, v := range i.([]protocol.HtmlMenu) {
			l := strings.Split(v.Value["link"], "|")
			if len(l) > 2 {
				menuItem := moduleMenu{
					Name:   v.Key,
					Hidden: v.Value["hidden"] == "true",
					Text:   l[0],
					Module: l[1],
					Method: l[2],
					Class:  v.Value["class"],
				}
				//允许访问视图
				if module == "common" {
					if v.Key != "my" {
						find := false
						if data.User.IsAdmin {
							find = true
						} else {
							find = data.User.AclMenu[v.Key]
						}
						if !find {
							continue
						}
					}
				} else {
					if !data.User.IsAdmin && (data.User.Priv[menuItem.Module] == nil || !data.User.Priv[menuItem.Module][menuItem.Method]) {
						continue
					}
				}
				if alias, ok := v.Value["alias"]; ok {
					menuItem.Alias = strings.Split(alias, ",")
				}
				if subModule, ok := v.Value["subModule"]; ok {
					menuItem.SubModule = strings.Split(subModule, ",")
				}
				if len(l) == 4 {
					for _, vars := range strings.Split(l[3], "&") {
						s := strings.Split(vars, "=")
						if len(s) == 2 {
							value := s[1]
							if replace, ok := data.App["menuReplace"].(map[string]string); ok {
								if r, ok := replace[s[0]]; ok {
									value = fmt.Sprintf(value, r)
								}
							}
							menuItem.Vars = append(menuItem.Vars, protocol.HtmlKeyValueStr{s[0], value})
						}

					}
				}
				if module == "common" {
					_menu := getModuleMenu(v.Key, data)
					if len(_menu) == 0 {
						continue
					}
					data.Data["getModuleMenu_"+v.Key] = _menu
				}
				menu = append(menu, menuItem)

			}
		}
	}
	return
}
func common_printOrderLink(data *TemplateData, fieldName, orderBy, vars, label string, moduleMethod ...string) string {
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

	return html_a(link, label, "", "class='"+className+"'")
}
func common_printIcon(data *TemplateData, module, method, vars string, object interface{}, typ, icon string, extvalue ...string) string { //($target = '', $extraClass = '', $onlyBody = false, $misc = '', $title = '')
	target := ""
	extraClass := ""
	misc := ""
	title := ""
	onlyBody := false
	if len(extvalue) > 0 {
		target = extvalue[0]
		if len(extvalue) > 1 {
			if data.ws.Query("isonlybody") == "yes" {
				if strings.Index(extvalue[1], "showinonlybody") == -1 {
					return ""
				}
				extvalue[1] = strings.ReplaceAll(extvalue[1], "iframe", "")
			}
			extraClass = extvalue[1]
		}
		if len(extvalue) > 2 {
			onlyBody = extvalue[2] == "true"
		}
		if len(extvalue) > 3 {
			misc = extvalue[3]
		}
		if len(extvalue) > 4 {
			title = extvalue[4]
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
		return ""
	}
	clickable := true
	if object != nil {
		key := ""
		r := reflect.ValueOf(object)
		switch r.Kind() {
		case reflect.Ptr:
			key = r.Elem().Type().Name()
		case reflect.Map:
			if k := r.MapIndex(reflect.ValueOf("isClickableKey")); k.Kind() == reflect.String {
				key = k.String()
			}
		case reflect.String:
			if r.String() == "" {
				key = "null"
			}
		}
		if f_interface, ok := global_Funcs[key+"_isClickable"]; ok {
			if f, ok := f_interface.(func(*TemplateData, interface{}, string) bool); ok {
				clickable = f(data, object, method)
			}
		} else {
			libraries.DebugLog("%s页面common_printIcon找不到 %s 的方法", data.ws.Path(), key+"_isClickable")
		}
	}
	link := createLink(module, method, []interface{}{vars, onlyBody})
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
			} else {
				if str, ok := data.Lang[module][method].(string); ok {
					title = str
				} else {
					if v, ok := data.Lang[module][method].(map[string]interface{}); ok {
						if str, ok := v["common"].(string); ok {
							title = str
						}
					}
					if v, ok := data.Lang[module][method].(map[string]string); ok {
						if str, ok := v["common"]; ok {
							title = str
						}
					}
				}
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
				return html_a(link, "<i class='"+class+"'></i> "+"<span class='text'>"+title+"</span>", target, "class='btn btn-link "+extraClass+"' "+misc)
			} else {
				return html_a(link, "<i class='"+class+"'></i>", target, "class='btn btn-link "+extraClass+"' title='"+title+"' "+misc)
			}
		} else {
			return html_a(link, "<i class='"+class+"'></i>", target, "class='btn "+extraClass+"' title='"+title+"' "+misc) + "\n"
		}
	} else {
		if typ == "list" {
			return "<button type='button' class='disabled btn " + extraClass + "'><i class='" + class + "' title='" + title + "' " + misc + "></i></button>\n"
		}
	}

	return ""
}
func common_fetch(oldData *TemplateData, module, method string, varstr ...string) string {
	path := "/" + module + "/" + method
	data := getFetchInterface(oldData.ws, path, oldData.User)
	if !hasPriv(data, module, method) {
		moduleName := module
		methodName := method
		if v, ok := data.Lang[module]["common"].(string); ok {
			moduleName = v
		}
		if v, ok := data.Lang[module][methodName].(string); ok {
			methodName = v
		}
		return fmt.Sprintf(data.Lang["user"]["error"].(map[string]string)["errorDeny"], moduleName, methodName)
	}
	if f, ok := httpHandlerMap["GET"][path]; ok {
		if len(varstr) > 0 {
			if len(varstr) == 1 {
				for _, vars := range strings.Split(varstr[0], "&") {
					s := strings.Split(vars, "=")
					if len(s) == 2 {
						data.ws.AddQuery(s[0], s[1])
					}
				}
			} else {
				for _, vars := range varstr {
					s := strings.Split(vars, "=")
					if len(s) == 2 {
						data.ws.AddQuery(s[0], s[1])
					}
				}
			}

		}
		var res string
		if err := f(data); err != nil {
			res = err.Error()
		} else {
			res = string(data.ws.(*CommonFetch).OutBuffer())
		}

		putFetchInterface(data.ws.(*CommonFetch))
		return res
	} else {
		putFetchInterface(data.ws.(*CommonFetch))
	}
	return "没有找到GET " + path + "方法"

}
func common_getValue(i, key interface{}) interface{} {
	switch r := i.(type) {
	case []protocol.HtmlKeyValueStr:
		k := libraries.I2S(key)
		for _, v := range r {
			if v.Key == k {
				return v.Value
			}
		}
		return nil
	case []protocol.HtmlKeyValueInterface:
		k := libraries.I2S(key)
		for _, v := range r {
			if v.Key == k {
				return v.Value
			}
		}
		return nil
	}
	r := reflect.ValueOf(i)
	for r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	if r.Kind() == reflect.Invalid {
		return nil
	}
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
	} else if r.Type().Kind() == reflect.Struct {
		value := r.FieldByName(libraries.I2S(key))
		if value.Kind() != reflect.Invalid {
			return value.Interface()
		}
	}
	return nil
}
func common_printLink(data *TemplateData, module, method string, v ...string) string {
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
		return ""
	}

	return html_a(createLink(module, method, vars), label, target, misc)
}
