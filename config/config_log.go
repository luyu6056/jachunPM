package config

import (
	"protocol"
)

func config_log_init(local protocol.CountryNo) {
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
}
