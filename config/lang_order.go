package config

func getOrder() map[string][]string {
	res := map[string][]string{}
	/* Sort of main menu. */
	res["common"] = []string{
		"my",
		"product",
		"project",
		"qa",
		"ci",
		"doc",
		"report",
		"company",
		"admin",
	}

	/* index menu order. */
	res["index"] = []string{
		"product",
		"project",
	}
	/* my menu order. */
	res["my"] = []string{
		"index",
		"calendar",
		"task",
		"bug",
		"testtask",
		"story",
		"myProject",
		"dynamic",
		"profile",
		"changePassword",
		"score",
	}
	res["todo"] = res["my"]

	/* product menu order. */
	res["product"] = []string{
		"story",
		"plan",
		"release",
		"roadmap",
		"project",
		"dynamic",
		"doc",
		"branch",
		"module",
		"view",
		"create",
		"all",
	}
	res["story"] = res["product"]
	res["productplan"] = res["product"]
	res["release"] = res["product"]
	res["branch"] = res["product"]

	/* project menu order. */
	res["project"] = []string{
		"task",
		"kanban",
		"burn",
		"list",
		"story",
		"qa",
		"doc",
		"team",
		"action",
		"product",
		"view",
	}
	res["task"] = res["project"]
	res["build"] = res["project"]

	/* qa menu order. */
	res["qa"] = []string{
		"product",
		"index",
		"bug",
		"testcase",
		"testtask",
		"report",
		"testsuite",
		"caselib",
	}
	res["bug"] = res["qa"]
	res["testcase"] = res["bug"]
	res["testtask"] = res["testcase"]
	res["testsuite"] = res["testcase"]
	res["caselib"] = res["testcase"]
	res["testreport"] = res["testcase"]
	res["ci"] = []string{
		"review",
		"code",
		"build",
		"jenkins",
		"maintain",
		"rules",
	}
	res["repo"] = res["ci"]
	res["jenkins"] = res["ci"]

	/* doc menu order. */
	res["doc"] = []string{
		"list",
		"product",
		"project",
		"custom",
		"index",
		"create",
	}
	/* report menu order. */
	res["report"] = []string{
		"annual",
		"product",
		"prj",
		"test",
		"staff",
	}
	/* company menu order. */
	res["company"] = []string{
		"browseUser",
		"dept",
		"effort",
		"browseGroup",
		"dynamic",
		"view",
		"addGroup",
		"batchAddUser",
		"addUser",
	}
	res["dept"] = res["company"]
	res["group"] = res["company"]
	res["user"] = res["company"]

	/* admin menu order. */
	res["company"] = []string{
		"index",
		"message",
		"custom",
		"sso",
		"extension",
		"dev",
		"translate",
		"data",
		"safe",
		"system",
	}

	res["build"] = res["project"]
	res["attend"] = res["oa"]
	res["convert"] = res["admin"]
	res["upgrade"] = res["admin"]
	res["action"] = res["admin"]
	res["backup"] = res["admin"]
	res["cron"] = res["admin"]
	res["extension"] = res["admin"]
	res["custom"] = res["admin"]
	res["mail"] = res["admin"]
	res["dev"] = res["admin"]
	res["entry"] = res["admin"]
	res["webhook"] = res["admin"]
	return res
}
func getSubMenuOrder() map[string]map[string][]string {
	res := map[string]map[string][]string{}
	res["admin"] = make(map[string][]string)
	res["admin"]["message"] = []string{
		"mail",
		"sms",
		"webhook",
		"browser",
		"setting",
	}
	res["admin"]["sso"] = []string{
		"ranzhi",
		"libreoffice",
		"ldap",
	}
	res["admin"]["dev"] = []string{
		"api",
		"db",
		"editor",
		"entry",
	}
	res["admin"]["data"] = []string{
		"backup",
		"trash",
	}
	res["admin"]["system"] = []string{
		"cron",
		"buildIndex",
		"timezone",
	}
	return res
}
