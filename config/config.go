package config

import (
	"protocol"
)

var Lang = make(map[protocol.CountryNo]map[string]map[string]interface{}) //map[name][key]=value
var Config = make(map[protocol.CountryNo]map[string]map[string]map[string]interface{})

func init() {

	//Lang[protocol.EN] = make(map[string]map[string]interface{})
	Lang[protocol.ZH_CN] = make(map[string]map[string]interface{})
	Config[protocol.ZH_CN] = make(map[string]map[string]map[string]interface{})
	//LangEnInit()
	LangZH_CNInit()
	configInit(protocol.ZH_CN)
}

type ConfigSearchParams struct {
	Operator string
	Control  string
	Values   []protocol.HtmlKeyValueStr
	ValueExt string
	Class    string
}
type ConfigSearch struct {
	Module string
	Fields []protocol.HtmlKeyValueStr
	Params map[string]*ConfigSearchParams
}

/*\['([A-z]+)']\s*= array\('operator' => '([^']+)',\s*'control' => '([^']+)',\s*'values' =>(.+?)(,\s*'class' => '([^']+)')?\);

"$1": &ConfigSearchParams{
	Operator: "$2",
	Control:  "$3",
	Values:   $4,
},
*/
func configInit(local protocol.CountryNo) {
	Config[local]["common"] = make(map[string]map[string]interface{})
	Config[local]["my"] = make(map[string]map[string]interface{})
	Config[local]["common"]["common"] = map[string]interface{}{}

	Config[local]["search"] = make(map[string]map[string]interface{})
	Config[local]["search"]["common"] = map[string]interface{}{
		"groupItems": 3,
	}

	Config[local]["datatable"] = make(map[string]map[string]interface{})
	Config[local]["datatable"]["moduleAlias"] = map[string]interface{}{
		"product-browse": "story",
		"project-task":   "task",
		"testtask-cases": "testcase",
		"my-task":        "task",
	}

	Config[local]["charsets"] = make(map[string]map[string]interface{})
	Config[local]["charsets"]["common"] = map[string]interface{}{
		"zh-cn": []protocol.HtmlKeyValueStr{
			{"utf-8", "UTF-8"},
			{"gbk", "GBK"},
		},
		"zh-tw": []protocol.HtmlKeyValueStr{
			{"utf-8", "UTF-8"},
			{"big5", "BIG5"},
		},
		"en": []protocol.HtmlKeyValueStr{
			{"utf-8", "UTF-8"},
			{"GBK", "GBK"},
		},
	}
	config_oa_init(local)
	config_log_init(local)
	config_project_init(local)
	config_user_init(local)
	config_test_init(local)
}
