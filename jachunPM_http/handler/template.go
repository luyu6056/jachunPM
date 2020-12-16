package handler

import (
	"html/template"
	"jachunPM_http/config"
	"libraries"
	"os"
	"protocol"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
)

type TemplateData struct {
	Config, App, Data map[string]interface{}
	Lang              map[string]map[string]interface{}
	ws                HttpRequest
}

var templateLock sync.RWMutex
var global_t = template.New("jachun")
var global_data = &TemplateData{}

func init() {
	loadFuncs()
	loadAlltemplate()
	loadTemplateData()
}

var T *template.Template
var n int

func loadAlltemplate() {
	iswatcher := false
	watcher, err := fsnotify.NewWatcher()
	iswatcher = err == nil

	dir, _ := os.Getwd()
	list, _ := libraries.ListDirAll(dir+"/template", "")
	for _, name := range list {
		if iswatcher {
			watcher.Add(name)
		}
		var err error
		global_t, err = global_t.ParseFiles(name)
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
				new_t, err := global_t.ParseFiles(event.Name)
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
func loadTemplateData() {
	global_data.Config = make(map[string]interface{})
	global_data.Config["debug"] = config.Config.Debug
	global_data.Config["webRoot"] = config.Config.Origin + "/"
	global_data.Config["jsRoot"] = config.Config.Origin + "/js/"
	global_data.Config["themeRoot"] = config.Config.Origin + "/theme/"
	global_data.Config["defaultTheme"] = config.Config.Origin + "/theme/default/"
	global_data.Config["langs"] = []protocol.HtmlKeyValueStr{{string(protocol.ZH_CN), protocol.ZH_CN.String()}}

}
func (data *TemplateData) Init(ws HttpRequest) *TemplateData {
	d := &TemplateData{
		Config: data.Config,
		Lang:   config.Lang[protocol.ZH_CN],
		App:    make(map[string]interface{}),
		Data:   make(map[string]interface{}),
		ws:     ws,
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
		if uid := session.Load_str("UserId"); uid != "" {
			var u protocol.MSG_USER_INFO_cache
			err := HostConn.CacheGet(protocol.UserServerNo, protocol.PATH_USER_INFO_CACHE, uid, &u)
			if err == nil {
				d.App["user"] = u
			} else {
				libraries.ReleaseLog("读取user_info缓存错误%s", err)
			}
		}

	}
	d.App["ClientLangString"] = protocol.CountryNo(d.App["ClientLang"].(string)).String()
	d.App["company"] = getCompanyInfo()
	d.Config["langTheme"] = global_data.Config["themeRoot"].(string) + "lang/" + d.App["ClientLang"].(string) + ".css"
	d.App["onlybody"] = ws.Query("onlybody")

	return d
}
func (data *TemplateData) onlybody() bool {
	if data.App != nil && data.App["onlybody"] == "yes" {
		return true
	}
	return false
}
func templateOut(name string, data *TemplateData, ws HttpRequest) {
	templateLock.RLock()
	buf := bufpool.Get().(*libraries.MsgBuffer)
	defer func() {
		buf.Reset()
		bufpool.Put(buf)
		templateLock.RUnlock()
	}()
	data.ws = ws
	data.App["TemplateName"] = name
	names := strings.Split(name, ".")
	data.App["moduleName"] = names[0]
	data.App["methodName"] = names[1]
	if data.Data["title"] == nil {
		data.Data["title"] = data.Lang[names[0]][names[1]]
	}
	err := T.ExecuteTemplate(buf, name, data)
	if err != nil {
		ws.OutErr(err)
	} else {
		ws.Write(buf)
	}

}
