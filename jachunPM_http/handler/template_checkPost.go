package handler

import (
	"jachunPM_http/config"
	"protocol"
	"strconv"
	"time"
)

type checkType int

const (
	checkTypeNone    checkType = 1 << iota
	checkTypeRequire           //不能为空
	checkTypeInt
	checkTypeFloat
	checkTypePositive //正数
	checkTypeZero     //包含0
	checkTypeNegative //负数
	checkTypeUserId   //允许0或者空,不允许错误的值
	checkTypeDate     //2006-01-02

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
	"/productplan/create": map[string]interface{}{
		"title": checkTypeRequire,
		"begin": checkTypeDate,
		"end":   checkTypeDate,
	},
	"/productplan/edit": map[string]interface{}{
		"title": checkTypeRequire,
		"begin": checkTypeDate,
		"end":   checkTypeDate,
	},
	"/story/create": map[string]interface{}{
		"title":      checkTypeRequire,
		"pri":        config.Lang[protocol.DefaultLang]["story"]["pri"].([]protocol.HtmlKeyValueStr),
		"estimate":   checkTypeFloat | checkTypePositive,
		"assignedTo": checkTypeUserId | checkTypePositive,
		"mailto":     checkTypeUserId,
	},
}

//post请求检查
func (data *TemplateData) ajaxCheckPost() bool {
	if check, ok := checkinfo[data.ws.Path()]; ok {
		for key, i := range check {
			reskey, err := func() (string, string) {
				list := data.ws.PostSlice(key)
				require := false
				if typ, ok := i.(checkType); ok && typ&checkTypeRequire == checkTypeRequire {
					require = true
					if len(list) == 0 {
						return key, data.Lang["error"]["checkTypeRequire"].(string)
					}
				}

				for _, v := range list {

					switch typ := i.(type) {
					case checkType:
						if require {
							if v == "" {
								return key, data.Lang["error"]["checkTypeRequire"].(string)
							}
						} else if v == "" {
							continue
						}
						switch typ {
						case checkTypeRequire:
							//上面已处理
						case checkTypeNum:
							_, err := strconv.Atoi(v)
							if err != nil {
								return key, data.Lang["error"]["checkTypeNum"].(string)
							}

						case checkTypeUserId:
							id, _ := strconv.Atoi(v)
							if HostConn.GetUserCacheById(int32(id)) == nil {
								return key, data.Lang["error"]["checkTypeUserId"].(string)
							}

						case checkTypeDate:
							_, err := time.Parse(protocol.TIMEFORMAT_MYSQLDATE, v)
							if err != nil {
								return key, data.Lang["error"]["checkTypeDate"].(string)
							}

						}
					case []protocol.HtmlKeyValueStr:
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

					case func() ([]protocol.HtmlKeyValueStr, error):
						checklist, _ := typ()
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
