package handler

import (
	"config"
	"errors"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"jachunPM_http/js"
	"jachunPM_http/setting"
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
	Page      TempLatePage
	Ajax      bool
}
type TempLatePage struct {
	Total      int
	Page       int
	PerPage    int
	CookieName string
	Param      string
}

var templateLock sync.RWMutex
var global_t = template.New("jachun")
var T *template.Template
var n int
var global_Funcs template.FuncMap = map[string]interface{}{}
var bufpool = sync.Pool{New: func() interface{} {
	return new(libraries.MsgBuffer)
}}
var TemplateDataChan sync.Map

func Init() {
	loadTemplateFuncs()
	loadAlltemplate()
}
func loadTemplateFuncs() {
	commonTemplateFuncs()
	htmlTemplateFuncs()
	hookTemplateFuncs()
	datatableTemplateFuncs()
	storyTemplateFuncs()
	productplanTemplateFuncs()
	fileTemplateFuncs()
	isClickableFuncs()
	actionTemplateFuncs()
	projectTemplateFuncs()
	customTemplateFuncs()
	taskTemplateFuncs()
	groupTemplateFuncs()
	blockTemplateFuncs()
	attendTemplateFuncs()
	global_t.Funcs(global_Funcs)
	copyConfig()
}
func loadAlltemplate() {

	iswatcher := false
	watcher, _ := fsnotify.NewWatcher()
	//iswatcher = err == nil

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
func copyConfig() {
	//异步深度拷贝config和lang
	for _, country := range protocol.AllCountry {
		go func(k protocol.CountryNo) {
			c := make(chan *TemplateData, 20)
			TemplateDataChan.Store(k, c)
			for {
				data := &TemplateData{
					App:    make(map[string]interface{}),
					Data:   make(map[string]interface{}),
					Config: make(map[string]map[string]map[string]interface{}),
					Lang:   make(map[string]map[string]interface{}),
				}

				for kk, vv := range config.Config[k] {
					tmp1 := make(map[string]map[string]interface{})
					for kkk, vvv := range vv {
						tmp2 := make(map[string]interface{})
						for kkkk, vvvv := range vvv {
							res := libraries.CopyMap(reflect.ValueOf(vvvv))
							if res.Kind() == reflect.Invalid {
								fmt.Println("config", k, kk, kkk, kkkk)
							} else {
								tmp2[kkkk] = res.Interface()
							}

						}
						tmp1[kkk] = tmp2
					}
					data.Config[kk] = tmp1
				}
				data.Config = setHttpConfig(data.Config)
				for kk, vv := range config.Lang[k] {
					tmp1 := make(map[string]interface{})
					for kkk, vvv := range vv {
						res := libraries.CopyMap(reflect.ValueOf(vvv))
						if res.Kind() == reflect.Invalid {
							fmt.Println("lang", k, kk, kkk)
						} else {
							tmp1[kkk] = res.Interface()
						}

					}
					data.Lang[kk] = tmp1
				}
				c <- data
			}
		}(country)
	}
}
func templateDataInit(ws HttpRequest, user *protocol.MSG_USER_INFO_cache) (data *TemplateData) {
	clientLang := protocol.DefaultLang
	clientTheme := "default"
	session := ws.Session()
	clientLang = getClientLang(ws)
	clientTheme = session.Load_str("ClientTheme")
	if clientTheme == "" {
		clientTheme = "default"
		session.Set("ClientTheme", "default")
	}
	if v, ok := TemplateDataChan.Load(clientLang); ok {
		data = <-v.(chan *TemplateData)
	}
	data.ws = ws
	data.Time = time.Now()
	if user == nil {
		user = HostConn.GetUserCacheById(session.Load_int32("UserId"))
	}
	data.User = user
	data.getMsg()
	data.App["ClientLang"] = string(clientLang)
	data.App["ClientTheme"] = clientTheme
	names := strings.Split(ws.Path(), "/")
	if len(names) > 2 {
		data.App["moduleName"] = names[1]
		data.App["methodName"] = names[2]
		data.Page.CookieName = "pager" + strings.ToUpper(names[1][:1]) + names[1][1:] + strings.ToUpper(names[2][:1]) + names[2][1:]
		if perPage := ws.Query("recPerPage"); perPage != "" {
			data.Page.PerPage, _ = strconv.Atoi(perPage)
		} else {
			if perPage := ws.Cookie(data.Page.CookieName); perPage != "" {
				data.Page.PerPage, _ = strconv.Atoi(perPage)
			}
		}
		if total := ws.Query("recTotal"); total != "" {
			data.Page.Total, _ = strconv.Atoi(total)
		}
		if page := ws.Query("pageID"); page != "" {
			data.Page.Page, _ = strconv.Atoi(page)
		}
		if data.Page.PerPage <= 0 {
			data.Page.PerPage = 20
		}
		if data.Page.Page < 1 {
			data.Page.Page = 1
		}
		datatableId := data.App["moduleName"].(string) + strings.ToUpper(data.App["methodName"].(string)[:1]) + data.App["methodName"].(string)[1:]
		if v1, ok := data.Config["datatable"][datatableId]; ok {
			if v2, ok := v1["mode"].(string); ok && v2 == "datatable" {
				data.Data["useDatatable"] = true
				data.Data["datatableId"] = datatableId
			}
		}
	} else {
		data.App["moduleName"] = ""
		data.App["methodName"] = ""
	}
	if data.User != nil {
		for module, v := range data.User.Config {
			if data.Config[module] == nil {
				data.Config[module] = make(map[string]map[string]interface{})
			}
			for section, vv := range v {
				if data.Config[module][section] == nil {
					data.Config[module][section] = make(map[string]interface{})
				}
				for key, value := range vv {
					data.Config[module][section][key] = value
				}
			}
		}
	}
	if systemConfig := custom_getSystemConfig(); systemConfig != nil {
		for module, v := range systemConfig {
			if data.Config[module] == nil {
				data.Config[module] = make(map[string]map[string]interface{})
			}
			for section, vv := range v {

				if data.Config[module][section] == nil {
					data.Config[module][section] = make(map[string]interface{})
				}
				for key, value := range vv {
					data.Config[module][section][key] = value
				}
			}
		}
	}
	//判断是不是ajax json请求
	if strings.Contains(data.ws.Header("Accept"), "application/json") || strings.Contains(data.ws.Header("accept"), "application/json") {
		data.Ajax = true
	}
	data.App["isAjaxRequest"] = data.isajax()
	data.App["ClientLangString"] = clientLang.String()
	data.App["company"] = getCompanyInfo()
	data.App["langTheme"] = data.Config["common"]["common"]["themeRoot"].(string) + "lang/" + string(clientLang) + ".css"
	data.App["onlybody"] = strings.ToLower(ws.Query("onlybody")) == "yes"
	return data
}
func (data *TemplateData) onlybody() bool {
	if data.App != nil && data.App["onlybody"].(bool) {
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
		data.outErr(err)
	} else {
		data.ws.Write(buf.Bytes())
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
	if len(locateAndcallback) > 0 && locateAndcallback[0] != "" {
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

var dataErrRedirect = errors.New("HttpServerIsRedirect")
var dataErrAlreadyOut = errors.New("AlreadyOutErrInfo")

func (data *TemplateData) outErr(err error) {
	if err == dataErrRedirect || err == dataErrAlreadyOut {
		return
	}
	if data.ws.Query("ajaxform") == "true" || data.Ajax {
		data.ajaxResult(false, err.Error(), "")
		return
	}
	if data.onlybody() {
		data.ws.WriteString(js.Error(err.Error()))
	} else {

		data.Data["err"] = template.HTML(err.Error())
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
			data.ws.Write(buf.Bytes())
		}
	}

}
func (data *TemplateData) getMsg() (*protocol.Msg, error) {
	if data.Msg == nil {
		msg := &protocol.Msg{DB: &protocol.MsgDB{DB: HostConn.DB}}
		if data.User != nil {
			msg.Uid = data.User.Id
		}
		msg.SetServer(&protocol.RpclientSend{HostConn})
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

//重新封装便捷消息发送，所有的消息都基于msg发出
func (data *TemplateData) SendMsgWaitResultToDefault(out protocol.MSG_DATA, result interface{}, timeout ...time.Duration) (err error) {
	return data.Msg.SendMsgWaitResult(0, out, result, timeout...)
}
func (data *TemplateData) SendMsgToDefault(out protocol.MSG_DATA) (err error) {
	data.Msg.SendMsg(0, out)
	return nil
}
func (data *TemplateData) getCacheProjectById(id int32) *protocol.MSG_PROJECT_project_cache {
	if v, ok := data.Data["project_cache_"+strconv.Itoa(int(id))].(*protocol.MSG_PROJECT_project_cache); ok {
		return v
	}
	if list, ok := data.Data["project_getAll"].([]*protocol.MSG_PROJECT_project_cache); ok {
		for _, p := range list {
			if p.Id == id {
				data.Data["project_cache_"+strconv.Itoa(int(id))] = p
				return p
			}
		}
	}

	project := HostConn.GetProjectById(id)
	data.Data["project_cache_"+strconv.Itoa(int(id))] = project
	return project
}

//执行了BeginTransaction后，data的send消息自动都带上事务
func (data *TemplateData) BeginTransaction() (session *protocol.MsgDBTransaction, err error) {
	return data.Msg.BeginTransaction()
}

func setHttpConfig(config map[string]map[string]map[string]interface{}) map[string]map[string]map[string]interface{} {
	config["common"]["common"]["debug"] = setting.Setting.Debug
	config["common"]["common"]["webRoot"] = setting.Setting.Origin + "/"
	config["common"]["common"]["jsRoot"] = setting.Setting.Origin + "/js/"
	config["common"]["common"]["themeRoot"] = setting.Setting.Origin + "/theme/"
	config["common"]["common"]["defaultTheme"] = setting.Setting.Origin + "/theme/default/"
	config["common"]["common"]["langs"] = []protocol.HtmlKeyValueStr{{string(protocol.ZH_CN), protocol.ZH_CN.String()}}
	config["common"]["common"]["maxUploadSize"] = "4000M"
	return config
}
