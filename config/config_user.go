package config

import (
	"protocol"
)

func config_user_init(local protocol.CountryNo) {
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
}
