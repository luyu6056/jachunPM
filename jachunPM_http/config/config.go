package config

import (
	"protocol"
)

type ConfigSearchParams struct {
	Operator string
	Control  string
	Values   []protocol.HtmlKeyValueStr
	ValueExt string
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
	Config[local]["product"] = make(map[string]map[string]interface{})
	Config[local]["product"]["common"] = map[string]interface{}{
		"orderBy":               "isClosed,order_desc",
		"customBatchEditFields": "line,PO,QD,RD,status,type,desc",
	}
	Config[local]["product"]["custom"] = map[string]interface{}{
		"batchEditFields": "line,PO,QD,RD,status",
	}
	Config[local]["product"]["list"] = map[string]interface{}{
		"exportFields": "id,name,line,activeStories,changedStories,draftStories,closedStories,plans,releases,bugs,unResolvedBugs,assignToNullBugs",
	}

	Config[local]["product"]["search"] = map[string]interface{}{
		"module": "story",
		"fields": map[string]interface{}{
			"title":    Lang[local]["story"]["title"].(string),
			"id":       Lang[local]["story"]["id"].(string),
			"keywords": Lang[local]["story"]["keywords"].(string),
			"stage":    Lang[local]["story"]["stage"].(string),
			"status":   Lang[local]["story"]["status"].(string),
			"pri":      Lang[local]["story"]["pri"].(string),

			"product":  Lang[local]["story"]["product"].(string),
			"branch":   "",
			"module":   Lang[local]["story"]["module"].(string),
			"plan":     Lang[local]["story"]["plan"].(string),
			"estimate": Lang[local]["story"]["estimate"].(string),

			"source":     Lang[local]["story"]["source"].(string),
			"sourceNote": Lang[local]["story"]["sourceNote"].(string),
			"fromBug":    Lang[local]["story"]["fromBug"].(string),

			"openedBy":     Lang[local]["story"]["openedBy"].(string),
			"reviewedBy":   Lang[local]["story"]["reviewedBy"].(string),
			"assignedTo":   Lang[local]["story"]["assignedTo"].(string),
			"closedBy":     Lang[local]["story"]["closedBy"].(string),
			"lastEditedBy": Lang[local]["story"]["lastEditedBy"].(string),

			"mailto": Lang[local]["story"]["mailto"].(string),

			"closedReason": Lang[local]["story"]["closedReason"].(string),
			"version":      Lang[local]["story"]["version"].(string),

			"openedDate":     Lang[local]["story"]["openedDate"].(string),
			"reviewedDate":   Lang[local]["story"]["reviewedDate"].(string),
			"assignedDate":   Lang[local]["story"]["assignedDate"].(string),
			"closedDate":     Lang[local]["story"]["closedDate"].(string),
			"lastEditedDate": Lang[local]["story"]["lastEditedDate"].(string),
			"params": map[string]ConfigSearchParams{
				"title": ConfigSearchParams{
					Operator: "include",
					Control:  "input",
				},
				"keywords": ConfigSearchParams{
					Operator: "include",
					Control:  "input",
				},
				"status": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					Values:   Lang[local]["story"]["statusList"].([]protocol.HtmlKeyValueStr),
				},
				"stage": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					Values:   Lang[local]["story"]["stageList"].([]protocol.HtmlKeyValueStr),
				},
				"pri": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					Values:   Lang[local]["story"]["priList"].([]protocol.HtmlKeyValueStr),
				},
				"product": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
				},
				"branch": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
				},
				"module": ConfigSearchParams{
					Operator: "belong",
					Control:  "select",
				},
				"plan": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
				},
				"estimate": ConfigSearchParams{
					Operator: "=",
					Control:  "input",
				},

				"source": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					Values:   Lang[local]["story"]["sourceList"].([]protocol.HtmlKeyValueStr),
				},
				"sourceNote": ConfigSearchParams{
					Operator: "include",
					Control:  "input",
				},
				"fromBug": ConfigSearchParams{
					Operator: "=",
					Control:  "input",
				},

				"openedBy": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					ValueExt: "users",
				},
				"reviewedBy": ConfigSearchParams{
					Operator: "include",
					Control:  "select",
					ValueExt: "users",
				},
				"assignedTo": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					ValueExt: "users",
				},
				"closedBy": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					ValueExt: "users",
				},
				"lastEditedBy": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					ValueExt: "users",
				},

				"mailto": ConfigSearchParams{
					Operator: "include",
					Control:  "select",
					ValueExt: "users",
				},

				"closedReason": ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					Values:   Lang[local]["story"]["reasonList"].([]protocol.HtmlKeyValueStr),
				},
				"version": ConfigSearchParams{
					Operator: ">=",
					Control:  "input",
				},

				"openedDate": ConfigSearchParams{
					Operator: "=",
					Control:  "input",
				},
				"reviewedDate": ConfigSearchParams{
					Operator: "=",
					Control:  "input",
				},
				"assignedDate": ConfigSearchParams{
					Operator: "=",
					Control:  "input",
				},
				"closedDate": ConfigSearchParams{
					Operator: "=",
					Control:  "input",
				},
				"lastEditedDate": ConfigSearchParams{
					Operator: "=",
					Control:  "input",
				},
			},
		},
	}
	Config[local]["product"]["create"] = map[string]interface{}{
		"requiredFields": "name",
	}
	Config[local]["product"]["edit"] = map[string]interface{}{
		"requiredFields": "name",
	}
	Config[local]["product"]["editor"] = map[string]interface{}{
		"create": map[string]interface{}{
			"id":    []string{"desc"},
			"tools": "simpleTools",
		},
		"edit": map[string]interface{}{
			"id":    []string{"desc"},
			"tools": "simpleTools",
		},
		"close": map[string]interface{}{
			"id":    []string{"comment"},
			"tools": "simpleTools",
		},
	}
	Config[local]["product"]["report"] = map[string]interface{}{
		"stageLabels":   []string{"wait", "planed", "released"},
		"planLabels":    []string{""},
		"projectLabels": []string{""},
	}
	Config[local]["datatable"] = make(map[string]map[string]interface{})
	Config[local]["datatable"]["moduleAlias"] = map[string]interface{}{
		"product-browse": "story",
		"project-task":   "task",
		"testtask-cases": "testcase",
		"my-task":        "task",
	}
	Config[local]["story"] = make(map[string]map[string]interface{})
	Config[local]["story"]["statusList"] = protocol.HtmlKeyValueStr2MapStringInterface(Lang[protocol.DefaultLang]["story"]["statusList"].([]protocol.HtmlKeyValueStr))
}
