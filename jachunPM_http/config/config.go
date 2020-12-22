package config

import (
	"protocol"
)

type ConfigSearchParams struct {
	Operator string
	Control  string
	Values   []protocol.HtmlKeyValueStr
	Class    string
}

/*\['([A-z]+)']\s*= array\('operator' => '([^']+)',\s*'control' => '([^']+)',\s*'values' =>(.+?)(,\s*'class' => '([^']+)')?\);

"$1": ConfigSearchParams{
	Operator: "$2",
	Control:  "$3",
	Values:   $4,
},
*/
func configInit(local protocol.CountryNo) {
	Config[local]["common"] = make(map[string]map[string]interface{})
	Config[local]["common"]["common"] = map[string]interface{}{
		"debug":        Server.Debug,
		"webRoot":      Server.Origin + "/",
		"jsRoot":       Server.Origin + "/js/",
		"themeRoot":    Server.Origin + "/theme/",
		"defaultTheme": Server.Origin + "/theme/default/",
		"langs":        []protocol.HtmlKeyValueStr{{string(protocol.ZH_CN), protocol.ZH_CN.String()}},
	}
	Config[local]["user"] = make(map[string]map[string]interface{})
	Config[local]["user"]["common"] = map[string]interface{}{
		"contactField":     []string{"Mobile", "QQ", "Weixin"},
		"failTimes":        6,
		"lockMinutes":      10,
		"batchCreate":      10,
		"showDeleted":      1,
		"weakPasswordlen":  6,
		"weakPasswordtype": protocol.CONIFG_weakPasswordAny,
	}

	Config[local]["company"] = make(map[string]map[string]interface{})
	Config[local]["company"]["browse"] = map[string]interface{}{
		"search": map[string]interface{}{
			"module": "user",
			"fields": []protocol.HtmlKeyValueStr{
				{"realname", Lang[local]["user"]["realname"].(string)},
				{"email", Lang[local]["user"]["email"].(string)},
				{"dept", Lang[local]["user"]["dept"].(string)},
				{"account", Lang[local]["user"]["account"].(string)},
				{"role", Lang[local]["user"]["role"].(string)},
				{"phone", Lang[local]["user"]["phone"].(string)},
				{"join", Lang[local]["user"]["join"].(string)},
				{"id", Lang[local]["user"]["id"].(string)},
				{"commiter", Lang[local]["user"]["commiter"].(string)},
				{"gender", Lang[local]["user"]["gender"].(string)},
				{"qq", Lang[local]["user"]["QQ"].(string)},
				//{"skype", Lang[local]["user"]["skype"].(string)},
				//{"yahoo", Lang[local]["user"]["yahoo"].(string)},
				//{"gtalk", Lang[local]["user"]["gtalk"].(string)},
				//{"wangwang", Lang[local]["user"]["wangwang"].(string)},
				{"address", Lang[local]["user"]["address"].(string)},
				{"zipcode", Lang[local]["user"]["zipcode"].(string)},
			},
			"params": map[string]ConfigSearchParams{
				"realname": ConfigSearchParams{
					Operator: "include",
					Control:  "input",
				},
				"email": ConfigSearchParams{
					Operator: "include",
					Control:  "input",
					Class:    "",
				},
				"dept": ConfigSearchParams{
					Operator: "belong",
					Control:  "select",
					Class:    "",
				},
				"account": ConfigSearchParams{
					Operator: "include",
					Control:  "input",
					Class:    "",
				},
				"role": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					Values:   Lang[local]["user"]["roleList"].([]protocol.HtmlKeyValueStr),
					Class:    "",
				},
				"phone": ConfigSearchParams{
					Operator: "include",
					Control:  "input",

					Class: "",
				},
				"join": ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "date",
				},
				"id": ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"commiter": ConfigSearchParams{
					Operator: "include",
					Control:  "select",

					Class: "",
				},
				"gender": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					Values:   Lang[local]["user"]["genderList"].([]protocol.HtmlKeyValueStr),
					Class:    "",
				},
				"qq": ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"skype": ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"yahoo": ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"gtalk": ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"wangwang": ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"address": ConfigSearchParams{
					Operator: "include",
					Control:  "input",

					Class: "",
				},
				"zipcode": ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
			},
		},
	}
	Config[local]["search"] = make(map[string]map[string]interface{})
	Config[local]["search"]["common"] = map[string]interface{}{
		"groupItems": 3,
	}
}
