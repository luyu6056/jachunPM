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
	Config[local]["my"] = make(map[string]map[string]interface{})
	Config[local]["common"]["common"] = map[string]interface{}{
		"debug":         Server.Debug,
		"webRoot":       Server.Origin + "/",
		"jsRoot":        Server.Origin + "/js/",
		"themeRoot":     Server.Origin + "/theme/",
		"defaultTheme":  Server.Origin + "/theme/default/",
		"langs":         []protocol.HtmlKeyValueStr{{string(protocol.ZH_CN), protocol.ZH_CN.String()}},
		"maxUploadSize": "4000M",
	}
	Config[local]["user"] = make(map[string]map[string]interface{})
	Config[local]["user"]["common"] = map[string]interface{}{
		"contactField":     []string{"Mobile", "QQ", "Weixin"},
		"failTimes":        6,
		"lockMinutes":      10,
		"batchCreate":      10,
		"showDeleted":      true,
		"weakPasswordlen":  6,
		"weakPasswordtype": protocol.CONIFG_weakPasswordAny,
	}
	if contacts, ok := Lang[local]["user"]["contacts"].(map[string]string); ok {
		Config[local]["user"]["contacts"] = make(map[string]interface{}, len(contacts))
		for k, v := range contacts {
			Config[local]["user"]["contacts"][k] = v
		}
	} else {
		Config[local]["user"]["contacts"] = map[string]interface{}{}
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

	Config[local]["tree"] = make(map[string]map[string]interface{})
	Config[local]["tree"]["common"] = map[string]interface{}{
		"noBrowse": ",productdoc,projectdoc,",
	}
	Config[local]["story"]["common"] = map[string]interface{}{
		"batchCreate":      3,
		"affectedFixedNum": 7,
		"needReview":       true,
		"statusList":       Lang[protocol.DefaultLang]["story"]["statusList"].([]protocol.HtmlKeyValueStr),
		"stageList":        Lang[protocol.DefaultLang]["story"]["stageList"].([]protocol.HtmlKeyValueStr),
	}

	Config[local]["story"]["batchClose"] = map[string]interface{}{
		"columns": 10,
	}
	Config[local]["story"]["create"] = map[string]interface{}{
		"requiredFields": "title",
	}
	Config[local]["story"]["edit"] = map[string]interface{}{}
	Config[local]["story"]["change"] = map[string]interface{}{
		"requiredFields": "title",
	}
	Config[local]["story"]["close"] = map[string]interface{}{
		"requiredFields": "closedReason",
	}
	Config[local]["story"]["review"] = map[string]interface{}{
		"requiredFields": "assignedTo,reviewedBy",
	}
	Config[local]["story"]["editor"] = map[string]interface{}{
		"create":   map[string]interface{}{"id": []string{"spec", "verify"}, "tools": "simpleTools"},
		"change":   map[string]interface{}{"id": []string{"spec", "verify", "comment"}, "tools": "simpleTools"},
		"edit":     map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"view":     map[string]interface{}{"id": []string{"comment", "lastComment"}, "tools": "simpleTools"},
		"close":    map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"review":   map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"activate": map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
	}

	Config[local]["story"]["list"] = map[string]interface{}{
		"exportFields": `id, product, branch, module, plan, source, sourceNote, title, spec, verify, keywords,
    pri, estimate, status, stage, taskCountAB, bugCountAB, caseCountAB,
    openedBy, openedDate, assignedTo, assignedDate, mailto,
    reviewedBy, reviewedDate,
    closedBy, closedDate, closedReason,
    lastEditedBy, lastEditedDate,
    childStories, linkStories, duplicateStory, files`,
		"customCreateFields":      []string{"source", "verify", "pri", "estimate", "mailto", "keywords"},
		"customBatchCreateFields": "plan,spec,source,verify,pri,estimate,review,keywords",
		"customBatchEditFields":   "branch,plan,estimate,pri,assignedTo,source,stage,closedBy,closedReason,keywords",
	}
	Config[local]["story"]["custom"] = map[string]interface{}{
		"createFields":      Config[local]["story"]["list"]["customCreateFields"],
		"batchCreateFields": "module,plan,spec,pri,estimate,review",
		"batchEditFields":   "branch,module,plan,estimate,pri,source,stage,closedBy,closedReason",
	}
	Config[local]["story"]["datatable"] = map[string]interface{}{
		"defaultField": []string{"id", "pri", "title", "plan", "openedBy", "assignedTo", "estimate", "status", "stage", "taskCount", "actions"},
		"fieldList": map[string]map[string]string{
			"id": map[string]string{
				"title":    "idAB",
				"fixed":    "left",
				"width":    "60",
				"required": "yes",
			},
			"pri": map[string]string{
				"title":    "priAB",
				"fixed":    "left",
				"width":    "50",
				"required": "no",
			},
			"title": map[string]string{
				"title":    "title",
				"fixed":    "left",
				"width":    "auto",
				"required": "yes",
			},

			"branch": map[string]string{
				"title":    "branch",
				"fixed":    "no",
				"width":    "100",
				"required": "no",
			},
			"keywords": map[string]string{
				"title":    "keywords",
				"fixed":    "no",
				"width":    "100",
				"required": "no",
			},
			"plan": map[string]string{
				"title":    "planAB",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"source": map[string]string{
				"title":    "source",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"sourceNote": map[string]string{
				"title":    "sourceNote",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"status": map[string]string{
				"title":    "statusAB",
				"fixed":    "no",
				"width":    "80",
				"required": "no",
			},

			"estimate": map[string]string{
				"title":    "estimateAB",
				"fixed":    "no",
				"width":    "65",
				"required": "no",
			},
			"stage": map[string]string{
				"title":    "stageAB",
				"fixed":    "no",
				"width":    "70",
				"required": "no",
			},
			"openedBy": map[string]string{
				"title":    "openedByAB",
				"fixed":    "no",
				"width":    "80",
				"required": "no",
			},
			"openedDate": map[string]string{
				"title":    "openedDate",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"assignedTo": map[string]string{
				"title":    "assignedToAB",
				"fixed":    "no",
				"width":    "80",
				"required": "no",
			},
			"assignedDate": map[string]string{
				"title":    "assignedDate",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"reviewedBy": map[string]string{
				"title":    "reviewedBy",
				"fixed":    "no",
				"width":    "80",
				"required": "no",
			},
			"reviewedDate": map[string]string{
				"title":    "reviewedDate",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"closedBy": map[string]string{
				"title":    "closedBy",
				"fixed":    "no",
				"width":    "80",
				"required": "no",
			},
			"closedDate": map[string]string{
				"title":    "closedDate",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"closedReason": map[string]string{
				"title":    "closedReason",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"lastEditedBy": map[string]string{
				"title":    "lastEditedBy",
				"fixed":    "no",
				"width":    "80",
				"required": "no",
			},
			"lastEditedDate": map[string]string{
				"title":    "lastEditedDate",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"mailto": map[string]string{
				"title":    "mailto",
				"fixed":    "no",
				"width":    "100",
				"required": "no",
			},
			"version": map[string]string{
				"title":    "version",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"taskCount": map[string]string{
				"title":    "T",
				"fixed":    "no",
				"width":    "30",
				"required": "no",
				"sort":     "no",
				"name":     Lang[local]["story"]["taskCount"].(string),
			},
			"bugCount": map[string]string{
				"title":    "B",
				"fixed":    "no",
				"width":    "30",
				"required": "no",
				"sort":     "no",
				"name":     Lang[local]["story"]["bugCount"].(string),
			},
			"caseCount": map[string]string{
				"title":    "C",
				"fixed":    "no",
				"width":    "30",
				"required": "no",
				"sort":     "no",
				"name":     Lang[local]["story"]["caseCount"].(string),
			},
			"actions": map[string]string{
				"title":    "actions",
				"fixed":    "right",
				"width":    "150",
				"required": "yes",
			},
		},
	}
	Config[local]["productplan"] = make(map[string]map[string]interface{})
	Config[local]["productplan"]["editor"] = map[string]interface{}{
		"create": map[string]interface{}{
			"id":    []string{"desc"},
			"tools": "simpleTools",
		},
		"edit": map[string]interface{}{
			"id":    []string{"desc"},
			"tools": "simpleTools",
		},
	}

}
