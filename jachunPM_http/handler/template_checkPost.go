package handler

import (
	"jachunPM_http/config"
	"protocol"
	"strconv"
)

type checkType int

const (
	checkTypeNone    = 1 << iota
	checkTypeRequire //不能为空
	checkTypeNum
	checkTypeUserId //允许0或者空,不允许错误的值

)

var checkinfo = map[string]map[string]interface{}{ //目前接受checkType和[]protocol.htmlKeyValStr,func() ([]protocol.HtmlKeyValueStr, error)
	"/product/create": map[string]interface{}{
		"name":      checkTypeRequire,
		"code":      checkTypeRequire,
		"PO":        checkTypeUserId,
		"QD":        checkTypeUserId,
		"RD":        checkTypeUserId,
		"whitelist": user_getGroupOptionMenu,
		"type":      config.Lang[protocol.DefaultLang]["product"]["typeList"].([]protocol.HtmlKeyValueStr),
		"acl":       config.Lang[protocol.DefaultLang]["product"]["aclList"].([]protocol.HtmlKeyValueStr),
	},
	"/product/edit": map[string]interface{}{
		"name":      checkTypeRequire,
		"code":      checkTypeRequire,
		"PO":        checkTypeUserId,
		"QD":        checkTypeUserId,
		"RD":        checkTypeUserId,
		"whitelist": user_getGroupOptionMenu,
		"type":      config.Lang[protocol.DefaultLang]["product"]["typeList"].([]protocol.HtmlKeyValueStr),
		"acl":       config.Lang[protocol.DefaultLang]["product"]["aclList"].([]protocol.HtmlKeyValueStr),
	},
}

//post请求检查
func (data *TemplateData) ajaxCheckPost() bool {
	if check, ok := checkinfo[data.ws.Path()]; ok {
		for key, i := range check {
			reskey, err := func() (string, string) {
				list := data.ws.PostSlice(key)
				switch typ := i.(type) {
				case checkType:
					switch typ {
					case checkTypeRequire:
						for _, v := range list {
							if v == "" {
								return key, data.Lang["error"]["checkTypeRequire"].(string)
							}
						}
					case checkTypeNum:
						for _, v := range list {
							_, err := strconv.Atoi(v)
							if err != nil {
								return key, data.Lang["error"]["checkTypeNum"].(string)
							}
						}
					case checkTypeUserId | checkTypeRequire: //必须是有效的userid且不能为空
						for _, v := range list {
							if v == "0" || v == "" {
								return key, data.Lang["error"]["checkTypeRequire"].(string)
							}
						}
						fallthrough
					case checkTypeUserId:
						for _, v := range list {
							if v == "0" || v == "" {
								continue
							}
							id, _ := strconv.Atoi(v)
							if HostConn.GetUserCacheById(int32(id)) == nil {
								return key, data.Lang["error"]["checkTypeUserId"].(string)
							}
						}
					}
				case []protocol.HtmlKeyValueStr:
					for _, v := range list {
						find := false
						for _, kv := range typ {
							if kv.Key == v {
								find = true
								break
							}
						}
						if !find {
							return key, data.Lang["error"]["checkHtmlKeyValueStr"].(string)
						}
					}
				case func() ([]protocol.HtmlKeyValueStr, error):
					checklist, _ := typ()
					for _, v := range list {
						find := false
						for _, kv := range checklist {
							if kv.Key == v {
								find = true
								break
							}
						}
						if !find {
							return key, data.Lang["error"]["checkHtmlKeyValueStr"].(string)
						}
					}
				}

				return "", ""
			}()
			if reskey != "" {
				data.ajaxResult(false, map[string]string{reskey: err})
				return false
			}
		}
	}
	return true
}
