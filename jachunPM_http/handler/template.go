package handler

import (
	"errors"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"jachunPM_http/config"
	"libraries"
	"mysql"
	"os"
	"path/filepath"
	"protocol"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/fsnotify/fsnotify"
	"github.com/luyu6056/reflect2"
)

type TemplateData struct {
	App, Data map[string]interface{}
	User      *protocol.MSG_USER_INFO_cache
	Config    map[string]map[string]map[string]interface{}
	Lang      map[string]map[string]interface{}
	ws        HttpRequest
	Msg       *protocol.Msg
	Time      time.Time
	Page      struct {
		Total      int
		Page       int
		PerPage    int
		CookieName string
		Param      string
	}
}

var templateLock sync.RWMutex
var global_t = template.New("jachun")
var T *template.Template
var n int
var global_Funcs template.FuncMap = map[string]interface{}{}
var bufpool = sync.Pool{New: func() interface{} {
	return new(libraries.MsgBuffer)
}}

func init() {
	loadFuncs()
	loadAlltemplate()

}
func loadFuncs() {
	commonModelFuncs()
	htmlFuncs()
	hookFuncs()
	datatableFuncs()
	storyFuncs()
	isClickableFuncs()
	actionModelFuncs()
	global_t.Funcs(global_Funcs)
}
func loadAlltemplate() {
	iswatcher := false
	watcher, err := fsnotify.NewWatcher()
	iswatcher = err == nil

	dir, _ := os.Getwd()
	list, _ := libraries.ListDirAll(dir+"/template", "")
	importTemplate := func(filename string) (*template.Template, error) {
		//js文件进行特殊转义
		if filename[len(filename)-3:] == ".js" {
			b, err := ioutil.ReadFile(filename)
			if err != nil {
				return nil, err
			}
			newt := global_t.New(filepath.Base(filename))
			return newt.Parse(html.EscapeString(string(b)))
		}
		return global_t.ParseFiles(filename)
	}
	for _, name := range list {
		if iswatcher {
			watcher.Add(name)
		}
		var err error
		global_t, err = importTemplate(name)
		if err != nil {
			panic(err)
		}

	}
	T, _ = global_t.Clone()
	go func() {

		for iswatcher {
			select {
			case event := <-watcher.Events:
				templateLock.Lock()
				new_t, err := importTemplate(event.Name)
				if err == nil {
					global_t = new_t
					T, _ = global_t.Clone()
				} else {
					libraries.ReleaseLog("热更新渲染模板%s错误，%v", event.Name, err)
				}
				templateLock.Unlock()

			case err := <-watcher.Errors:
				libraries.ReleaseLog("error:%v", err)
			}
		}
	}()
}

func templateDataInit(ws HttpRequest) *TemplateData {
	d := &TemplateData{
		App:  make(map[string]interface{}),
		Data: make(map[string]interface{}),
		ws:   ws,
		Time: time.Now(),
	}
	if ws.Cookie("sessionID") == "" {
		d.App["ClientLang"] = string(protocol.DefaultLang)
		d.App["ClientTheme"] = "default"
	} else {
		session := ws.Session()
		d.App["ClientLang"] = string(getClientLang(ws))
		d.App["ClientTheme"] = session.Load_str("ClientTheme")
		if d.App["ClientTheme"] == "" {
			d.App["ClientTheme"] = "default"
			session.Set("ClientTheme", "default")
		}
		if uid := session.Load_int32("UserId"); uid > 0 {
			d.User = HostConn.GetUserCacheById(uid)
		}
	}
	d.Config = config.Config[protocol.CountryNo(d.App["ClientLang"].(string))]
	d.Lang = config.Lang[protocol.CountryNo(d.App["ClientLang"].(string))]
	names := strings.Split(ws.Path(), "/")
	if len(names) > 2 {
		d.App["moduleName"] = names[1]
		d.App["methodName"] = names[2]
		d.Page.CookieName = "pager" + strings.ToUpper(names[1][:1]) + names[1][1:] + strings.ToUpper(names[2][:1]) + names[2][1:]
		if perPage := ws.Query("recPerPage"); perPage != "" {
			d.Page.PerPage, _ = strconv.Atoi(perPage)
		} else {
			if perPage := ws.Cookie(d.Page.CookieName); perPage != "" {
				d.Page.PerPage, _ = strconv.Atoi(perPage)
			}
		}
		if total := ws.Query("recTotal"); total != "" {
			d.Page.Total, _ = strconv.Atoi(total)
		}
		if page := ws.Query("pageID"); page != "" {
			d.Page.Page, _ = strconv.Atoi(page)
		}
		if d.Page.PerPage <= 0 {
			d.Page.PerPage = 20
		}
		if d.Page.Page < 1 {
			d.Page.Page = 1
		}
		datatableId := d.App["moduleName"].(string) + strings.ToUpper(d.App["methodName"].(string)[:1]) + d.App["methodName"].(string)[1:]
		if v1, ok := d.Config["datatable"][datatableId]; ok {
			if v2, ok := v1["mode"].(string); ok && v2 == "datatable" {
				d.Data["useDatatable"] = true
				d.Data["datatableId"] = datatableId
			}
		}
	} else {
		d.App["moduleName"] = ""
		d.App["methodName"] = ""
	}

	d.App["ClientLangString"] = protocol.CountryNo(d.App["ClientLang"].(string)).String()
	d.App["company"] = getCompanyInfo()
	d.Config["common"]["common"]["langTheme"] = d.Config["common"]["common"]["themeRoot"].(string) + "lang/" + d.App["ClientLang"].(string) + ".css"
	d.App["onlybody"] = ws.Query("onlybody")
	return d
}
func (data *TemplateData) onlybody() bool {
	if data.App != nil && data.App["onlybody"] == "yes" {
		return true
	}
	return false
}
func templateOut(name string, data *TemplateData) {
	templateLock.RLock()
	buf := bufpool.Get().(*libraries.MsgBuffer)
	defer func() {
		buf.Reset()
		bufpool.Put(buf)
		templateLock.RUnlock()
	}()
	data.App["TemplateName"] = name

	if data.Data["title"] == nil {
		data.Data["title"] = data.Lang[data.App["moduleName"].(string)][data.App["methodName"].(string)]
	}
	err := T.ExecuteTemplate(buf, name, data)
	if err != nil {
		data.OutErr(err)
	} else {
		data.ws.Write(buf)
	}

}

//ajaxForm组件返回
func (data *TemplateData) ajaxResult(result bool, message interface{}, locateAndcallback ...string) {
	buf := bufpool.Get().(*libraries.MsgBuffer)
	if result {
		buf.WriteString(`{"result":"success","message":`)
	} else {
		buf.WriteString(`{"result":"fail","message":`)
	}
	buf.Write(libraries.JsonMarshal(message))
	if len(locateAndcallback) > 0 {
		buf.WriteString(`,"locate":"`)
		buf.WriteString(strings.ReplaceAll(locateAndcallback[0], `"`, `\"`))
		buf.WriteByte('"')
	}
	if len(locateAndcallback) > 1 {
		buf.WriteString(`,"callback":"`)
		buf.WriteString(strings.ReplaceAll(locateAndcallback[1], `"`, `\"`))
		buf.WriteByte('"')
	}
	buf.WriteString("}")
	data.ws.WriteString(buf.String())
	buf.Reset()
	bufpool.Put(buf)
}
func (data *TemplateData) OutErr(err error) {
	data.Data["err"] = err.Error()
	templateLock.RLock()
	buf := bufpool.Get().(*libraries.MsgBuffer)
	defer func() {
		buf.Reset()
		bufpool.Put(buf)
		templateLock.RUnlock()
	}()
	name := "error.html"
	data.App["TemplateName"] = name
	data.Data["title"] = "无法访问"
	e := T.ExecuteTemplate(buf, name, data)
	if e != nil {
		libraries.ReleaseLog("%+v", e)
	} else {
		data.ws.Write(buf)
	}
}
func (data *TemplateData) GetMsg() (*protocol.Msg, error) {
	if data.Msg == nil {
		out := protocol.GET_MSG_COMMON_GET_Msgno()
		defer out.Put()
		var resdata *protocol.MSG_COMMON_GET_Msgno_result
		err := HostConn.SendMsgWaitResultToDefault(out, &resdata)
		if err != nil {
			return nil, err
		}
		msg := &protocol.Msg{Msgno: resdata.Msgno, DB: &protocol.MsgDB{DB: HostConn.DB}}
		msg.SetServer(HostConn)
		data.Msg = msg
	}

	return data.Msg, nil
}
func (data *TemplateData) isajax() bool {
	return data.ws.Header("X-Requested-With") == "XMLHttpRequest"
}

//从post获取值赋值到struct，忽略错误，请在ajaxCheckPost里面提前规范好输入
func (data *TemplateData) SetValueFromPost(i interface{}) {
	t := reflect.TypeOf(i).Elem()
	ref_ptr := uintptr(reflect2.PtrOf(i))
	for i := 0; i < t.NumField(); i++ {
		if field := t.Field(i); field.Type.Kind() != reflect.Invalid {
			value := data.ws.PostSlice(strings.ToLower(field.Name))
			if len(value) == 0 {
				value = data.ws.PostSlice(field.Name)
			}
			if len(value) == 0 {
				continue
			}
			if err := setValueUnsafe(ref_ptr+field.Offset, field.Type, value); err != nil {
				libraries.DebugLog("Name %s,err %v", field.Name, err)
			}
		}
	}
}
func setValueUnsafe(ptr uintptr, t reflect.Type, value []string) error {
	switch t.Kind() {
	case reflect.Int:
		ii, _ := strconv.Atoi(value[0])
		*((*int)(unsafe.Pointer(ptr))) = ii
	case reflect.Int8:
		ii, _ := strconv.Atoi(value[0])
		*((*int8)(unsafe.Pointer(ptr))) = int8(ii)
	case reflect.Int16:
		ii, _ := strconv.Atoi(value[0])
		*((*int16)(unsafe.Pointer(ptr))) = int16(ii)
	case reflect.Int32:
		ii, _ := strconv.Atoi(value[0])
		*((*int32)(unsafe.Pointer(ptr))) = int32(ii)
	case reflect.Int64:
		ii, _ := strconv.Atoi(value[0])
		*((*int64)(unsafe.Pointer(ptr))) = int64(ii)
	case reflect.Uint:
		ii, _ := strconv.Atoi(value[0])
		*((*uint)(unsafe.Pointer(ptr))) = uint(ii)
	case reflect.Uint8:
		ii, _ := strconv.Atoi(value[0])
		*((*uint8)(unsafe.Pointer(ptr))) = uint8(ii)
	case reflect.Uint16:
		ii, _ := strconv.Atoi(value[0])
		*((*uint16)(unsafe.Pointer(ptr))) = uint16(ii)
	case reflect.Uint32:
		ii, _ := strconv.Atoi(value[0])
		*((*uint32)(unsafe.Pointer(ptr))) = uint32(ii)
	case reflect.Uint64:
		ii, _ := strconv.Atoi(value[0])
		*((*uint64)(unsafe.Pointer(ptr))) = uint64(ii)
	case reflect.String:
		*((*string)(unsafe.Pointer(ptr))) = value[0]
	case reflect.Float32:
		ii, _ := strconv.ParseFloat(value[0], 32)
		*((*float32)(unsafe.Pointer(ptr))) = float32(ii)
	case reflect.Float64:
		ii, _ := strconv.ParseFloat(value[0], 64)
		*((*float64)(unsafe.Pointer(ptr))) = float64(ii)
	case reflect.Slice:
		header := (*mysql.SliceHeader)(unsafe.Pointer(ptr))
		if uintptr(header.Data) == 0 || header.Cap < len(value) {
			header.Data = unsafe.Pointer(reflect.MakeSlice(t, len(value), len(value)).Pointer())
			header.Len = len(value)
			header.Cap = len(value)
		} else {
			header.Len = len(value)
		}
		memberT := t.Elem()
		for i := 0; i < len(value); i++ {
			if err := setValueUnsafe(uintptr(header.Data)+memberT.Size()*uintptr(i), memberT, []string{value[i]}); err != nil {
				return err
			}
		}
	default:
		return errors.New(fmt.Sprintf("未处理的Kind %v", t.Kind()))
	}
	return nil
}
