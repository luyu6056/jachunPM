package config

import (
	"protocol"
)

func config_test_init(local protocol.CountryNo) {
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
}
