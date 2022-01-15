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
		"search": &ConfigSearch{
			Module: "user",
			Fields: []protocol.HtmlKeyValueStr{
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
				{"attendNo", Lang[local]["user"]["attendNo"].(string)},
			},
			Params: map[string]*ConfigSearchParams{
				"realname": &ConfigSearchParams{
					Operator: "include",
					Control:  "input",
				},
				"email": &ConfigSearchParams{
					Operator: "include",
					Control:  "input",
					Class:    "",
				},
				"dept": &ConfigSearchParams{
					Operator: "belong",
					Control:  "select",
					Class:    "",
				},
				"account": &ConfigSearchParams{
					Operator: "include",
					Control:  "input",
					Class:    "",
				},
				"role": &ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					Values:   Lang[local]["user"]["roleList"].([]protocol.HtmlKeyValueStr),
					Class:    "",
				},
				"phone": &ConfigSearchParams{
					Operator: "include",
					Control:  "input",

					Class: "",
				},
				"join": &ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "date",
				},
				"id": &ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"commiter": &ConfigSearchParams{
					Operator: "include",
					Control:  "select",

					Class: "",
				},
				"gender": &ConfigSearchParams{
					Operator: "=",
					Control:  "select",
					Values:   Lang[local]["user"]["genderList"].([]protocol.HtmlKeyValueStr),
					Class:    "",
				},
				"qq": &ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"skype": &ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"yahoo": &ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"gtalk": &ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"wangwang": &ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"address": &ConfigSearchParams{
					Operator: "include",
					Control:  "input",

					Class: "",
				},
				"zipcode": &ConfigSearchParams{
					Operator: "=",
					Control:  "input",

					Class: "",
				},
				"attendNo": &ConfigSearchParams{
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

	Config[local]["product"]["common"]["search"] = &ConfigSearch{
		Module: "story",
		Fields: []protocol.HtmlKeyValueStr{
			{"title", Lang[local]["story"]["title"].(string)},
			{"id", Lang[local]["story"]["id"].(string)},
			{"keywords", Lang[local]["story"]["keywords"].(string)},
			{"stage", Lang[local]["story"]["stage"].(string)},
			{"status", Lang[local]["story"]["status"].(string)},
			{"pri", Lang[local]["story"]["pri"].(string)},

			{"product", Lang[local]["story"]["product"].(string)},
			{"branch", ""},
			{"module", Lang[local]["story"]["module"].(string)},
			{"plan", Lang[local]["story"]["plan"].(string)},
			{"estimate", Lang[local]["story"]["estimate"].(string)},

			{"source", Lang[local]["story"]["source"].(string)},
			{"sourceNote", Lang[local]["story"]["sourceNote"].(string)},
			{"fromBug", Lang[local]["story"]["fromBug"].(string)},

			{"openedBy", Lang[local]["story"]["openedBy"].(string)},
			{"reviewedBy", Lang[local]["story"]["reviewedBy"].(string)},
			{"assignedTo", Lang[local]["story"]["assignedTo"].(string)},
			{"closedBy", Lang[local]["story"]["closedBy"].(string)},
			{"lastEditedBy", Lang[local]["story"]["lastEditedBy"].(string)},

			{"mailto", Lang[local]["story"]["mailto"].(string)},

			{"closedReason", Lang[local]["story"]["closedReason"].(string)},
			{"version", Lang[local]["story"]["version"].(string)},

			{"openedDate", Lang[local]["story"]["openedDate"].(string)},
			{"reviewedDate", Lang[local]["story"]["reviewedDate"].(string)},
			{"assignedDate", Lang[local]["story"]["assignedDate"].(string)},
			{"closedDate", Lang[local]["story"]["closedDate"].(string)},
			{"lastEditedDate", Lang[local]["story"]["lastEditedDate"].(string)},
		},
		Params: map[string]*ConfigSearchParams{
			"title": &ConfigSearchParams{
				Operator: "include",
				Control:  "input",
			},
			"keywords": &ConfigSearchParams{
				Operator: "include",
				Control:  "input",
			},
			"status": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				Values:   Lang[local]["story"]["statusList"].([]protocol.HtmlKeyValueStr),
			},
			"stage": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				Values:   Lang[local]["story"]["stageList"].([]protocol.HtmlKeyValueStr),
			},
			"pri": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				Values:   Lang[local]["story"]["priList"].([]protocol.HtmlKeyValueStr),
			},
			"product": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
			},
			"branch": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
			},
			"module": &ConfigSearchParams{
				Operator: "belong",
				Control:  "select",
			},
			"plan": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
			},
			"estimate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
			},

			"source": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				Values:   Lang[local]["story"]["sourceList"].([]protocol.HtmlKeyValueStr),
			},
			"sourceNote": &ConfigSearchParams{
				Operator: "include",
				Control:  "input",
			},
			"fromBug": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
			},

			"openedBy": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				ValueExt: "users",
			},
			"reviewedBy": &ConfigSearchParams{
				Operator: "include",
				Control:  "select",
				ValueExt: "users",
			},
			"assignedTo": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				ValueExt: "users",
			},
			"closedBy": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				ValueExt: "users",
			},
			"lastEditedBy": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				ValueExt: "users",
			},

			"mailto": &ConfigSearchParams{
				Operator: "include",
				Control:  "select",
				ValueExt: "users",
			},

			"closedReason": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				Values:   Lang[local]["story"]["reasonList"].([]protocol.HtmlKeyValueStr),
			},
			"version": &ConfigSearchParams{
				Operator: ">=",
				Control:  "input",
			},

			"openedDate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
			},
			"reviewedDate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
			},
			"assignedDate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
			},
			"closedDate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
			},
			"lastEditedDate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
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
		"customCreateFields":      "source,verify,pri,estimate,mailto,keywords",
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
	Config[local]["project"] = make(map[string]map[string]interface{})
	Config[local]["project"]["common"] = make(map[string]interface{})
	Config[local]["project"]["common"]["defaultWorkhours"] = float64(7.0)
	Config[local]["project"]["common"]["orderBy"] = "isDone,status,order_desc"
	Config[local]["project"]["common"]["maxBurnDay"] = 31
	Config[local]["project"]["common"]["weekend"] = 2
	Config[local]["project"]["list"] = make(map[string]interface{})

	Config[local]["project"]["list"]["exportFields"] = []string{"id", "name", "code", "PM", "end", "status", "totalEstimate", "totalConsumed", "totalLeft", "progress"}

	Config[local]["project"]["create"] = make(map[string]interface{})
	Config[local]["project"]["edit"] = make(map[string]interface{})

	Config[local]["project"]["create"]["requiredFields"] = "name,code,begin,end"
	Config[local]["project"]["edit"]["requiredFields"] = "name,code,begin,end"

	Config[local]["project"]["common"]["customBatchEditFields"] = "days,type,teamname,status,desc,PO,QD,PM,RD"
	Config[local]["project"]["custom"] = make(map[string]interface{})
	Config[local]["project"]["custom"]["batchEditFields"] = []string{"days", "status", "PM"}

	Config[local]["project"]["editor"] = make(map[string]interface{})
	Config[local]["project"]["editor"]["create"] = map[string]string{"id": "desc", "tools": "simpleTools"}
	Config[local]["project"]["editor"]["edit"] = map[string]string{"id": "desc", "tools": "simpleTools"}
	Config[local]["project"]["editor"]["putoff"] = map[string]string{"id": "comment", "tools": "simpleTools"}
	Config[local]["project"]["editor"]["activate"] = map[string]string{"id": "comment", "tools": "simpleTools"}
	Config[local]["project"]["editor"]["close"] = map[string]string{"id": "comment", "tools": "simpleTools"}
	Config[local]["project"]["editor"]["start"] = map[string]string{"id": "comment", "tools": "simpleTools"}
	Config[local]["project"]["editor"]["suspend"] = map[string]string{"id": "comment", "tools": "simpleTools"}
	Config[local]["project"]["editor"]["tree"] = map[string]string{"id": "comment", "tools": "simpleTools"}

	Config[local]["project"]["common"]["search"] = &ConfigSearch{
		Module: "task",
		Fields: []protocol.HtmlKeyValueStr{
			{"name", Lang[local]["task"]["name"].(string)},
			{"id", Lang[local]["task"]["id"].(string)},
			{"status", Lang[local]["task"]["status"].(string)},
			{"desc", Lang[local]["task"]["desc"].(string)},
			{"assignedTo", Lang[local]["task"]["assignedTo"].(string)},
			{"pri", Lang[local]["task"]["pri"].(string)},
			{"project", Lang[local]["task"]["project"].(string)},
			{"module", Lang[local]["task"]["module"].(string)},
			{"estimate", Lang[local]["task"]["estimate"].(string)},
			{"left", Lang[local]["task"]["left"].(string)},
			{"consumed", Lang[local]["task"]["consumed"].(string)},
			{"type", Lang[local]["task"]["type"].(string)},
			{"fromBug", Lang[local]["task"]["fromBug"].(string)},
			{"closedReason", Lang[local]["task"]["closedReason"].(string)},
			{"openedBy", Lang[local]["task"]["openedBy"].(string)},
			{"finishedBy", Lang[local]["task"]["finishedBy"].(string)},
			{"closedBy", Lang[local]["task"]["closedBy"].(string)},
			{"canceledBy", Lang[local]["task"]["canceledBy"].(string)},
			{"lastEditedBy", Lang[local]["task"]["lastEditedBy"].(string)},
			{"parent", Lang[local]["task"]["parent"].(string)},
			{"proofreading", Lang[local]["task"]["proofreading"].(string)},
			{"examine", Lang[local]["task"]["examine"].(string)},
			{"mailto", Lang[local]["task"]["mailto"].(string)},
			{"finishedList", Lang[local]["task"]["finishedList"].(string)},
			{"openedDate", Lang[local]["task"]["openedDate"].(string)},
			{"deadline", Lang[local]["task"]["deadline"].(string)},
			{"estStarted", Lang[local]["task"]["estStarted"].(string)},
			{"realStarted", Lang[local]["task"]["realStarted"].(string)},
			{"placeOrder", Lang[local]["task"]["placeOrder"].(string)},
			{"assignedDate", Lang[local]["task"]["assignedDate"].(string)},
			{"finishedDate", Lang[local]["task"]["finishedDate"].(string)},
			{"closedDate", Lang[local]["task"]["closedDate"].(string)},
			{"canceledDate", Lang[local]["task"]["canceledDate"].(string)},
			{"lastEditedDate", Lang[local]["task"]["lastEditedDate"].(string)},
		},
		Params: map[string]*ConfigSearchParams{
			"name": &ConfigSearchParams{
				Operator: "include",
				Control:  "input",
				Values:   nil,
			},
			"status": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				Values:   Lang[local]["task"]["statusList"].([]protocol.HtmlKeyValueStr),
			},
			"desc": &ConfigSearchParams{
				Operator: "include",
				Control:  "input",
				Values:   nil,
			},
			"assignedTo": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				ValueExt: "users",
			},
			"pri": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				Values:   append([]protocol.HtmlKeyValueStr{{}}, Lang[local]["task"]["priList"].([]protocol.HtmlKeyValueStr)...),
			},

			"project": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				Values:   nil,
			},
			"module": &ConfigSearchParams{
				Operator: "belong",
				Control:  "select",
				Values:   nil,
			},
			"estimate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Values:   nil,
			},
			"left": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Values:   nil,
			},
			"consumed": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Values:   nil,
			},
			"type": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				Values:   Lang[local]["task"]["typeList"].([]protocol.HtmlKeyValueStr),
			},
			"fromBug": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Values:   Lang[local]["task"]["typeList"].([]protocol.HtmlKeyValueStr),
			},
			"closedReason": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				Values:   Lang[local]["task"]["reasonList"].([]protocol.HtmlKeyValueStr),
			},
			"openedBy": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				ValueExt: "users",
			},
			"finishedBy": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				ValueExt: "users",
			},
			"closedBy": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				ValueExt: "users",
			},
			"cancelBy": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				ValueExt: "users",
			},
			"lastEditedBy": &ConfigSearchParams{
				Operator: "=",
				Control:  "select",
				ValueExt: "users",
			},

			"mailto": &ConfigSearchParams{
				Operator: "include",
				Control:  "select",
				ValueExt: "users",
			},
			"finishedList": &ConfigSearchParams{
				Operator: "include",
				Control:  "select",
				ValueExt: "users",
			},

			"openedDate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Class:    "date",
			},
			"deadline": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Class:    "date",
			},
			"estStarted": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Class:    "date",
			},
			"realStarted": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Class:    "date",
			},
			"assignedDate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Class:    "date",
			},
			"finishedDate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Class:    "date",
			},
			"closedDate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Class:    "date",
			},
			"canceledDate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Class:    "date",
			},
			"lastEditedDate": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Class:    "date",
			},
			"proofreading": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Values:   nil,
			},
			"activatedCount": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Values:   nil,
			},
			"examine": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Values:   nil,
			},
			"placeOrder": &ConfigSearchParams{
				Operator: "=",
				Control:  "input",
				Values:   nil,
			},
		},
	}
	Config[local]["printKanban"] = make(map[string]map[string]interface{})
	Config[local]["printKanban"]["col"] = make(map[string]interface{})
	Config[local]["printKanban"]["col"]["story"] = 1
	Config[local]["printKanban"]["col"]["wait"] = 2
	Config[local]["printKanban"]["col"]["doing"] = 3
	Config[local]["printKanban"]["col"]["done"] = 4
	Config[local]["printKanban"]["col"]["closed"] = 5
	Config[local]["project"]["kanbanSetting"] = make(map[string]interface{})
	Config[local]["project"]["kanbanSetting"]["colorList"] = map[string]interface{}{
		"wait":   "#7EC5FF",
		"doing":  "#0991FF",
		"pause":  "#fdc137",
		"done":   "#0BD986",
		"cancel": "#CBD0DB",
		"closed": "#838A9D",
	}
	Config[local]["action"] = make(map[string]map[string]interface{})
	Config[local]["action"]["common"] = map[string]interface{}{
		"commonImgSize": 870,
	}
	Config[local]["action"]["objectTypes"] = make(map[string]interface{})
	for k, v := range Lang[local]["action"]["objectTypes"].(map[string]string) {
		Config[local]["action"]["objectTypes"][k] = v
	}
	Config[local]["action"]["majorList"] = map[string]interface{}{
		"task":    []string{"assigned", "finished", "activated"},
		"bug":     []string{"assigned", "resolved"},
		"release": []string{"opened"},
		"build":   []string{"opened"},
	}
	Config[local]["action"]["label"] = make(map[string]interface{})
	for k, v := range Lang[local]["action"]["label"].(map[string]interface{}) {
		Config[local]["action"]["label"][k] = v
	}

	Config[local]["task"] = make(map[string]map[string]interface{})
	Config[local]["task"]["common"] = map[string]interface{}{
		"batchCreate": 5,
	}
	Config[local]["task"]["create"] = map[string]interface{}{
		"requiredFields": "name,type",
	}
	Config[local]["task"]["edit"] = map[string]interface{}{
		"requiredFields": Config[local]["task"]["create"]["requiredFields"].(string),
	}

	Config[local]["task"]["finish"] = map[string]interface{}{
		"requiredFields": "consumed",
	}
	Config[local]["task"]["activate"] = map[string]interface{}{
		"requiredFields": "left",
	}

	Config[local]["task"]["editor"] = map[string]interface{}{
		"create":      map[string]interface{}{"id": []string{"desc"}, "tools": "simpleTools"},
		"edit":        map[string]interface{}{"id": []string{"desc", "comment"}, "tools": "simpleTools"},
		"view":        map[string]interface{}{"id": []string{"comment", "lastComment"}, "tools": "simpleTools"},
		"assignTo":    map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"start":       map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"restart":     map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"finish":      map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"close":       map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"activate":    map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"cancel":      map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"pause":       map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"finishall":   map[string]interface{}{"id": []string{"comment"}, "tools": "simpleTools"},
		"batchCreate": map[string]interface{}{"id": []string{"desc0", "desc1", "desc2", "desc3", "desc4"}, "tools": "simpleTools"},
	}
	Config[local]["task"]["common"]["exportFields"] = []string{"id", "project", "module", "story", "name", "desc", "type", "pri", "estStarted", "realStarted", "deadline", "status", "estimate", "consumed", "left", "mailto", "progress", "openedBy", "openedDate", "assignedTo", "assignedDate", "finishedBy", "proofreading", "finishedDate", "canceledBy", "canceledDate", "closedBy", "closedDate", "closedReason", "lastEditedBy", "lastEditedDate", "files"}

	Config[local]["task"]["common"]["customCreateFields"] = []string{"story", "estStarted", "deadline", "mailto", "pri", "estimate"}
	Config[local]["task"]["common"]["customBatchCreateFields"] = []string{"module", "story", "assignedTo", "estimate", "estStarted", "deadline", "desc", "pri"}
	Config[local]["task"]["common"]["customBatchEditFields"] = []string{"module", "assignedTo", "status", "pri", "estimate", "record", "left", "estStarted", "deadline", "finishedBy", "canceledBy", "closedBy", "closedReason"}

	Config[local]["task"]["custom"] = map[string]interface{}{
		"createFields":      Config[local]["task"]["common"]["customCreateFields"],
		"batchCreateFields": "module,story,assignedTo,estimate,desc,pri",
		"batchEditFields":   "module,assignedTo,status,pri,estimate,record,left,finishedBy,closedBy,closedReason",
	}

	Config[local]["task"]["datatable"] = map[string]interface{}{
		"defaultField": []string{"id", "pri", "project", "name", "status", "assignedTo", "finishedBy", "estimate", "consumed", "left", "progress", "finalfile", "examine", "deadline", "actions"},
		"fieldList": map[string]map[string]string{
			"id": map[string]string{
				"title":    "idAB",
				"fixed":    "left",
				"width":    "100",
				"required": "yes",
			},
			"pri": map[string]string{
				"title":    "priAB",
				"fixed":    "left",
				"width":    "50",
				"required": "no",
			},
			"project": map[string]string{
				"title":    "project",
				"fixed":    "left",
				"width":    "auto",
				"required": "no",
			},

			"name": map[string]string{
				"title":    "name",
				"fixed":    "left",
				"width":    "auto",
				"required": "yes",
			},
			"type": map[string]string{
				"title":    "type",
				"fixed":    "no",
				"width":    "80",
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
				"width":    "60",
				"required": "no",
			},
			"consumed": map[string]string{
				"title":    "consumedAB",
				"fixed":    "no",
				"width":    "60",
				"required": "no",
			},
			"left": map[string]string{
				"title":    "leftAB",
				"fixed":    "no",
				"width":    "60",
				"required": "no",
			},
			"progress": map[string]string{
				"title":    "progress",
				"fixed":    "no",
				"width":    "50",
				"required": "no",
				"sort":     "no",
			},
			"deadline": map[string]string{
				"title":    "deadlineAB",
				"fixed":    "no",
				"width":    "60",
				"required": "no",
			},
			"openedBy": map[string]string{
				"title":    "openedByAB",
				"fixed":    "no",
				"width":    "70",
				"required": "no",
			},
			"openedDate": map[string]string{
				"title":    "openedDate",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"estStarted": map[string]string{
				"title":    "estStarted",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"realStarted": map[string]string{
				"title":    "realStarted",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"placeOrder": map[string]string{
				"title":    "placeOrder",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"assignedTo": map[string]string{
				"title":    "assignedTo",
				"fixed":    "no",
				"width":    "100",
				"required": "no",
			},
			"assignedDate": map[string]string{
				"title":    "assignedDate",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"finishedBy": map[string]string{
				"title":    "finishedByAB",
				"fixed":    "no",
				"width":    "80",
				"required": "no",
			},
			"finishedDate": map[string]string{
				"title":    "finishedDateAB",
				"fixed":    "no",
				"width":    "90",
				"required": "no",
			},
			"canceledBy": map[string]string{
				"title":    "canceledBy",
				"fixed":    "no",
				"width":    "80",
				"required": "no",
			},
			"canceledDate": map[string]string{
				"title":    "canceledDate",
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
				"width":    "80",
				"required": "no",
			},
			"finalfile": map[string]string{
				"title":    "finalfile",
				"fixed":    "no",
				"width":    "80",
				"required": "no",
			},
			"examine": map[string]string{
				"title":    "examine",
				"fixed":    "no",
				"width":    "80",
				"required": "no",
			},
			"proofreading": map[string]string{
				"title":    "proofreading",
				"fixed":    "no",
				"width":    "80",
				"required": "no",
			},
			"story": map[string]string{
				"title":    "storyAB",
				"fixed":    "no",
				"width":    "40",
				"required": "no",
				"name":     Lang[local]["task"]["story"].(string),
			},
			"mailto": map[string]string{
				"title":    "mailto",
				"fixed":    "no",
				"width":    "100",
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
			"actions": map[string]string{
				"title":    "actions",
				"fixed":    "right",
				"width":    "190",
				"required": "yes",
			},
		},
	}
	Config[local]["testtask"] = make(map[string]map[string]interface{})
	Config[local]["testtask"]["create"] = map[string]interface{}{
		"requiredFields": "project,build,begin,end,name",
	}
	Config[local]["testtask"]["edit"] = map[string]interface{}{
		"requiredFields": "project,build,begin,end,name",
	}

	Config[local]["testtask"]["editor"] = map[string]interface{}{
		"create": map[string]interface{}{
			"id":    []string{"desc"},
			"tools": "simpleTools",
		},
		"edit": map[string]interface{}{
			"id":    []string{"desc", "report", "comment"},
			"tools": "simpleTools",
		},
		"view": map[string]interface{}{
			"id":    []string{"comment"},
			"tools": "simpleTools",
		},
		"start": map[string]interface{}{
			"id":    []string{"report", "comment"},
			"tools": "simpleTools",
		},
		"block": map[string]interface{}{
			"id":    []string{"comment"},
			"tools": "simpleTools",
		},
		"activate": map[string]interface{}{
			"id":    []string{"comment"},
			"tools": "simpleTools",
		},
	}
	Config[local]["testtask"]["datatable"] = map[string]interface{}{
		"defaultField": []string{"id", "pri", "title", "type", "assignedTo", "lastRunner", "lastRunDate", "lastRunResult", "status", "bugs", "results", "stepNumber", "actions"},
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
}
