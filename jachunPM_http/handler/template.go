package handler

import (
	"html/template"
	"jachunPM_http/config"
	"libraries"
	"os"
	"protocol"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

type TemplateData struct {
	App, Data    map[string]interface{}
	Config, Lang map[string]map[string]interface{}
	ws           HttpRequest
	Msg          *protocol.Msg
	Time         time.Time
	Page         struct {
		Total      int
		Page       int
		PerPage    int
		CookieName string
		Param      string
	}
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
	global_data.Config = make(map[string]map[string]interface{})
	global_data.Config["common"] = make(map[string]interface{})
	global_data.Config["common"]["debug"] = config.Config.Debug
	global_data.Config["common"]["webRoot"] = config.Config.Origin + "/"
	global_data.Config["common"]["jsRoot"] = config.Config.Origin + "/js/"
	global_data.Config["common"]["themeRoot"] = config.Config.Origin + "/theme/"
	global_data.Config["common"]["defaultTheme"] = config.Config.Origin + "/theme/default/"
	global_data.Config["common"]["langs"] = []protocol.HtmlKeyValueStr{{string(protocol.ZH_CN), protocol.ZH_CN.String()}}
	global_data.Config["user"] = make(map[string]interface{})
}
func (data *TemplateData) Init(ws HttpRequest) *TemplateData {
	d := &TemplateData{
		Config: data.Config,
		Lang:   config.Lang[protocol.ZH_CN],
		App:    make(map[string]interface{}),
		Data:   make(map[string]interface{}),
		ws:     ws,
		Time:   time.Now(),
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
	} else {
		d.App["moduleName"] = ""
		d.App["methodName"] = ""
	}

	d.App["ClientLangString"] = protocol.CountryNo(d.App["ClientLang"].(string)).String()
	d.App["company"] = getCompanyInfo()
	d.Config["common"]["langTheme"] = global_data.Config["common"]["themeRoot"].(string) + "lang/" + d.App["ClientLang"].(string) + ".css"
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

	if data.Data["title"] == nil {
		data.Data["title"] = data.Lang[data.App["moduleName"].(string)][data.App["methodName"].(string)]
	}
	err := T.ExecuteTemplate(buf, name, data)
	if err != nil {
		ws.OutErr(err)
	} else {
		ws.Write(buf)
	}

}
