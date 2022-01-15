package protocol

import (
	"libraries"
	"net/url"
	"reflect"
	"runtime/debug"
)

//定义一些http渲染模板的结构体
type HtmlKeyValueStr struct {
	Key   string
	Value string
}
type HtmlKeyValueInterface struct {
	Key   string
	Value interface{}
}
type HtmlMenu struct {
	Key   string
	Value map[string]string
}
type HtmlBlock struct {
	Source string
	Grid   int8
	Title  string
	Block  string
	Params  map[string]string
	Height int16
	BlockLink string
	MoreLink string
	ActionLink string
	Id int32
	Order int8
	Module string
}
type HtmlBlockModule struct {
	AvailableBlocks  map[string]string
	MoreLinkList map[string]string
}

func HtmlKeyValueStr2MapStringInterface(in []HtmlKeyValueStr) (out map[string]interface{}) {
	out = make(map[string]interface{}, len(in))
	for _, kv := range in {
		out[kv.Key] = kv.Value
	}
	return
}
func CreateLink(moduleName string, methodName string, vars interface{}) string {
	buf := BufPoolGet()
	buf.WriteString("/")
	buf.WriteString(moduleName)
	buf.WriteString("/")
	buf.WriteString(methodName)
	switch v := vars.(type) {
	case []HtmlKeyValueStr:
		if len(v) > 0 {
			buf.WriteByte('?')
			for _, v := range v {
				buf.WriteString(url.QueryEscape(v.Key))
				buf.WriteByte('=')
				buf.WriteString(v.Value)
				buf.WriteByte('&')
			}
			buf.Truncate(buf.Len() - 1)
		}
	case []string:
		if len(v) > 0 {
			buf.WriteByte('?')
			for _, s := range v {
				buf.WriteString(s)
			}

		}
	case string:
		if v != "" {
			buf.WriteByte('?')
			buf.WriteString(v)
		}

	case nil:
	case []interface{}:
		if len(v) > 0 {
			buf.WriteByte('?')
			for k, s := range v {
				str := libraries.I2S(s)
				if k == len(v)-1 && str == "true" || k == len(v)-1 && str == "false" { //onlybody
					if str == "true" {
						if k == 0 {
							buf.WriteString("onlybody=yes")
						} else {
							buf.WriteString("&onlybody=yes")
						}
					}
				} else {
					buf.WriteString(str)
				}

			}
		}
	default:
		libraries.DebugLog("createLink不识别类型%s\r\n%s", reflect.TypeOf(v).String(), string(debug.Stack()))
	}

	res := buf.String()
	buf.Reset()
	BufPoolPut(buf)
	return res
}
