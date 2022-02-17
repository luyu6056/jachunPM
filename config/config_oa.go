package config

import (
	"protocol"
)

func config_oa_init(local protocol.CountryNo) {
	Config[local]["attend"] = make(map[string]map[string]interface{})
	Config[local]["attend"]["custom"] = map[string]interface{}{
		"beginDate":            "",
		"notAllowSignInLimit":  "07:00",
		"signInLimit":          "09:30",
		"signOutLimit":         "18:00",
		"mustSignOut":          "no",
		"notAllowSignOutLimit": "07:00",
		"workingHours":         "8",
		"halfAbsendMin":        "60",
		"absendMin":            "180",
		"reviewedBy":           "",
		"noAttendUsers":        "",
		"noAttendDepts":        "",
		"workingDays":          "1,2,3,4,5",
	}
	Config[local]["attend"]["require"] = map[string]interface{}{
		"review": "comment",
	}

	Config[local]["attend"]["editor"] = map[string]interface{}{
		"review": map[string]interface{}{"id": []string{"comment"}, "tools": "simple"},
	}
	Config[local]["attend"]["list"] = map[string]interface{}{
		"exportFields": "dept, realname, date, dayName, status, signIn, signOut, ip, desc",
	}
}
