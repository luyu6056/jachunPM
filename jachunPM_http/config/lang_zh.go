package config

import (
	"fmt"
	"html/template"
	"protocol"
)

func LangZH_CNInit() {
	Lang[protocol.ZH_CN]["common"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["common"]["productCommon"] = "产品"
	Lang[protocol.ZH_CN]["common"]["projectCommon"] = "项目"
	Lang[protocol.ZH_CN]["common"]["storyCommon"] = "需求"
	Lang[protocol.ZH_CN]["common"]["backhome"] = "返回"
	Lang[protocol.ZH_CN]["common"]["textHasMoreItems"] = "还有 {0} 项..."
	Lang[protocol.ZH_CN]["common"]["textNetworkError"] = "网络错误"
	Lang[protocol.ZH_CN]["common"]["today"] = "今天"
	Lang[protocol.ZH_CN]["common"]["searchAB"] = "搜索"
	Lang[protocol.ZH_CN]["common"]["deleteing"] = "删除中"
	Lang[protocol.ZH_CN]["common"]["confirmDelete"] = "您确定要执行删除操作吗？"
	Lang[protocol.ZH_CN]["common"]["exportIcon"] = "<i class='icon-upload-alt'> </i>"
	Lang[protocol.ZH_CN]["common"]["minus"] = " - "
	Lang[protocol.ZH_CN]["common"]["detail"] = "详情"
	Lang[protocol.ZH_CN]["common"]["month"] = "月"
	Lang[protocol.ZH_CN]["common"]["create"] = "新建"

	Lang[protocol.ZH_CN]["common"]["pasteImgUploading"] = "正在上传图片，请稍后..."
	Lang[protocol.ZH_CN]["common"]["pasteImgFail"] = "贴图失败，请稍后重试。"
	Lang[protocol.ZH_CN]["common"]["noticePasteImg"] = "可以在编辑器直接贴图。"
	Lang[protocol.ZH_CN]["common"]["chooseUsersToMail"] = "选择要发信通知的用户..."
	Lang[protocol.ZH_CN]["common"]["searchMore"] = "搜索此关键字的更多结果："
	Lang[protocol.ZH_CN]["common"]["noResultsMatch"] = "没有匹配结果"
	Lang[protocol.ZH_CN]["common"]["importAndInsert"] = "全新插入"
	Lang[protocol.ZH_CN]["common"]["importAndCover"] = "覆盖"
	Lang[protocol.ZH_CN]["common"]["importConfirm"] = "导入确认"
	Lang[protocol.ZH_CN]["common"]["noticeImport"] = "导入数据中，含有已经存在系统的数据，请确认这些数据要覆盖或者全新插入。"
	Lang[protocol.ZH_CN]["common"]["pasteTextInfo"] = "粘贴文本到文本域中，每行文字作为一条数据的标题。"

	Lang[protocol.ZH_CN]["common"]["clientHelp"] = "客户端使用说明"
	Lang[protocol.ZH_CN]["common"]["downloadClient"] = "下载客户端"
	Lang[protocol.ZH_CN]["common"]["downNotify"] = "下载桌面提醒"
	Lang[protocol.ZH_CN]["common"]["theme"] = "主题"
	Lang[protocol.ZH_CN]["common"]["lang"] = "Language"
	Lang[protocol.ZH_CN]["common"]["importEncodeList['utf-8']"] = "UTF-8"
	Lang[protocol.ZH_CN]["common"]["searchTips"] = `` //"编号(ctrl+g)"替换``
	Lang[protocol.ZH_CN]["common"]["dividerMenu"] = []string{"qa", "report"}
	Lang[protocol.ZH_CN]["common"]["common"] = "公有模块"
	Lang[protocol.ZH_CN]["common"]["typeAB"] = "类型"
	Lang[protocol.ZH_CN]["common"]["assignedToAB"] = "指派"
	Lang[protocol.ZH_CN]["common"]["openedByAB"] = "创建"
	Lang[protocol.ZH_CN]["common"]["statusAB"] = "状态"
	Lang[protocol.ZH_CN]["common"]["priAB"] = "P"
	Lang[protocol.ZH_CN]["common"]["idAB"] = "ID"
	Lang[protocol.ZH_CN]["common"]["workingHour"] = "工时"
	Lang[protocol.ZH_CN]["common"]["year"] = "年"
	Lang[protocol.ZH_CN]["common"]["future"] = "未来"
	Lang[protocol.ZH_CN]["common"]["selectedItems"] = "已选择 <strong>{0}</strong> 项"
	Lang[protocol.ZH_CN]["common"]["showAll"] = "[[全部显示]]"
	Lang[protocol.ZH_CN]["common"]["notPage"] = "抱歉，您访问的功能正在开发中！"
	Lang[protocol.ZH_CN]["common"]["notFound"] = "抱歉，您访问的对象并不存在！"
	Lang[protocol.ZH_CN]["common"]["loading"] = "稍候..."
	Lang[protocol.ZH_CN]["common"]["selectReverse"] = "反选"
	Lang[protocol.ZH_CN]["common"]["selectAll"] = "全选"
	Lang[protocol.ZH_CN]["common"]["select"] = "选择"
	Lang[protocol.ZH_CN]["common"]["backShortcutKey"] = "[快捷键:Alt+↑]"
	Lang[protocol.ZH_CN]["common"]["nextShortcutKey"] = "[快捷键:→]"
	Lang[protocol.ZH_CN]["common"]["preShortcutKey"] = "[快捷键:←]"
	Lang[protocol.ZH_CN]["common"]["tutorialConfirm"] = "检测到你尚未退出新手教程模式，是否现在退出？"
	Lang[protocol.ZH_CN]["common"]["lineNumber"] = "行号"
	Lang[protocol.ZH_CN]["common"]["customField"] = "自定义表单项"
	Lang[protocol.ZH_CN]["common"]["customMenu"] = "自定义导航"
	Lang[protocol.ZH_CN]["common"]["manualUrl"] = "需要修改链接"
	Lang[protocol.ZH_CN]["common"]["manual"] = "手册"
	Lang[protocol.ZH_CN]["common"]["changeLog"] = "修改日志"
	Lang[protocol.ZH_CN]["common"]["noviceTutorial"] = "新手教程"
	Lang[protocol.ZH_CN]["common"]["homepage"] = "设为模块首页"
	Lang[protocol.ZH_CN]["common"]["fold"] = "-"
	Lang[protocol.ZH_CN]["common"]["unfold"] = "+"
	Lang[protocol.ZH_CN]["common"]["ipLimited"] = "<html><head><meta http-equiv='Content-Type' content='text/html; charset=utf-8' /></head><body>抱歉，管理员限制当前IP登录，请联系管理员解除限制。</body></html>"
	Lang[protocol.ZH_CN]["common"]["duplicate"] = "已有相同标题的%s"
	Lang[protocol.ZH_CN]["common"]["repairTable"] = "数据库表可能损坏，请用其他工具检查修复。"
	Lang[protocol.ZH_CN]["common"]["timeout"] = "连接超时，请检查网络环境，或重试！"
	Lang[protocol.ZH_CN]["common"]["uploadImages"] = "多图上传 "
	Lang[protocol.ZH_CN]["common"]["pasteText"] = "多项录入"
	Lang[protocol.ZH_CN]["common"]["files"] = "附件 "
	Lang[protocol.ZH_CN]["common"]["addFiles"] = "上传了附件 "
	Lang[protocol.ZH_CN]["common"]["fail"] = "失败"
	Lang[protocol.ZH_CN]["common"]["collapse"] = "收起"
	Lang[protocol.ZH_CN]["common"]["expand"] = "展开全部"
	Lang[protocol.ZH_CN]["common"]["switchDisplay"] = "切换显示"
	Lang[protocol.ZH_CN]["common"]["reverse"] = "切换顺序"
	Lang[protocol.ZH_CN]["common"]["attatch"] = "附件"
	Lang[protocol.ZH_CN]["common"]["history"] = "历史记录"
	Lang[protocol.ZH_CN]["common"]["comment"] = "备注"
	Lang[protocol.ZH_CN]["common"]["restore"] = "恢复默认"
	Lang[protocol.ZH_CN]["common"]["actions"] = "操作"
	Lang[protocol.ZH_CN]["common"]["retrack"] = "收起"
	Lang[protocol.ZH_CN]["common"]["fullscreen"] = "全屏"
	Lang[protocol.ZH_CN]["common"]["noData"] = "暂无"
	Lang[protocol.ZH_CN]["common"]["required"] = `必填` //"必填"替换`必填`
	Lang[protocol.ZH_CN]["common"]["sort"] = "排序"
	Lang[protocol.ZH_CN]["common"]["trunk"] = "主干"
	Lang[protocol.ZH_CN]["common"]["public"] = "公共"
	Lang[protocol.ZH_CN]["common"]["customConfig"] = "自定义"
	Lang[protocol.ZH_CN]["common"]["day"] = `日` //"天"替换`日`
	Lang[protocol.ZH_CN]["common"]["more"] = "更多"
	Lang[protocol.ZH_CN]["common"]["goPC"] = "PC版"
	Lang[protocol.ZH_CN]["common"]["goback"] = "返回"
	Lang[protocol.ZH_CN]["common"]["preview"] = "查看"
	Lang[protocol.ZH_CN]["common"]["confirm"] = "确认"
	Lang[protocol.ZH_CN]["common"]["saveSuccess"] = `保存成功`
	Lang[protocol.ZH_CN]["common"]["save"] = "保存"
	Lang[protocol.ZH_CN]["common"]["submitting"] = "稍候..."
	Lang[protocol.ZH_CN]["common"]["setFileName"] = "文件名："
	Lang[protocol.ZH_CN]["common"]["export"] = "导出"
	Lang[protocol.ZH_CN]["common"]["import"] = "导入"
	Lang[protocol.ZH_CN]["common"]["unlink"] = "移除"
	Lang[protocol.ZH_CN]["common"]["close"] = "关闭"
	Lang[protocol.ZH_CN]["common"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["common"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["common"]["refresh"] = "刷新"
	Lang[protocol.ZH_CN]["common"]["cancel"] = "取消"
	Lang[protocol.ZH_CN]["common"]["reset"] = "重填"
	Lang[protocol.ZH_CN]["common"]["runInfo"] = "<div class='row'><div class='u-1 a-center' id='debugbar'>时间: %s 毫秒, 内存: %s KB, 查询: %s.  </div></div>"
	Lang[protocol.ZH_CN]["common"]["changePassword"] = "更改密码"
	Lang[protocol.ZH_CN]["common"]["profile"] = "个人档案"
	Lang[protocol.ZH_CN]["common"]["help"] = "帮助"
	Lang[protocol.ZH_CN]["common"]["login"] = "登录"
	Lang[protocol.ZH_CN]["common"]["logout"] = "退出"
	Lang[protocol.ZH_CN]["common"]["welcome"] = "%s项目管理系统"
	Lang[protocol.ZH_CN]["common"]["logoImg"] = "zt-logo.png"
	Lang[protocol.ZH_CN]["common"]["jachunPM"] = "杰骏"
	Lang[protocol.ZH_CN]["common"]["dash"] = "-"
	Lang[protocol.ZH_CN]["common"]["percent"] = "%"
	Lang[protocol.ZH_CN]["common"]["ellipsis"] = "…"
	Lang[protocol.ZH_CN]["common"]["null"] = "空"
	Lang[protocol.ZH_CN]["common"]["downArrow"] = "↓"
	Lang[protocol.ZH_CN]["common"]["at"] = " 于 "
	Lang[protocol.ZH_CN]["common"]["dot"] = "。"
	Lang[protocol.ZH_CN]["common"]["comma"] = "，"
	Lang[protocol.ZH_CN]["common"]["colon"] = "-"

	Lang[protocol.ZH_CN]["common"]["arrow"] = template.HTML(`&nbsp;<i class="icon-angle-right"></i>&nbsp;`)
	Lang[protocol.ZH_CN]["common"]["menu"] = []protocol.HtmlMenu{
		{"my", map[string]string{`link`: "<span> 我的地盘</span>|my|index"}},
		{"product", map[string]string{`link`: Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "|product|index|locate=no"}},
		{"project", map[string]string{`link`: Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "|project|index|locate=no"}},
		{"qa", map[string]string{`link`: "测试|qa|index"}},
		{"oa", map[string]string{"link": "办公|attend|personal"}},
		//{"doc", map[string]string{`link`: "文档|doc|index"}},
		//{"report", map[string]string{`link`: "统计|report|index"}},
		{"company", map[string]string{`link`: "组织|company|browse"}},
		{"admin", map[string]string{`link`: "后台|admin|index"}},
	}
	Lang[protocol.ZH_CN]["common"]["searchObjects"] = map[string]string{
		"bug":         "Bug",
		"story":       Lang[protocol.ZH_CN]["common"]["storyCommon"].(string),
		"task":        "任务",
		"testcase":    "用例",
		"project":     Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"product":     Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"user":        "用户",
		"build":       "版本",
		"release":     "发布",
		"productplan": Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "计划",
		"testtask":    "测试单",
		"doc":         "文档",
		"caselib":     "用例库",
		"testreport":  "测试报告",
		"all":         "全部",
	}
	Lang[protocol.ZH_CN]["common"]["importEncodeList"] = map[string]string{
		"gbk":  "GBK",
		"big5": "BIG5",
	}
	Lang[protocol.ZH_CN]["common"]["exportFileTypeList"] = []protocol.HtmlKeyValueStr{
		{"xlsx", "xlsx"},
		//{"csv", "csv"},
		//{"xml", "xml"},
		//{"html", "html"},
		//{"xls", "xls"},
	}
	Lang[protocol.ZH_CN]["common"]["exportTypeList"] = []protocol.HtmlKeyValueStr{
		{"all", "全部记录"},
		{"selected", "选中记录"},
	}
	Lang[protocol.ZH_CN]["common"]["themes"] = map[string]string{
		"default":    "深邃蓝（默认）",
		"green":      "叶兰绿",
		"red":        "赤诚红",
		"purple":     "玉烟紫",
		"pink":       "芙蕖粉",
		"blackberry": "露莓黑",
		"classic":    "经典蓝",
	}
	Lang[protocol.ZH_CN]["index"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["index"]["notice"] = "公告"
	Lang[protocol.ZH_CN]["index"]["index"] = "首页"
	Lang[protocol.ZH_CN]["index"]["common"] = "首页"
	Lang[protocol.ZH_CN]["index"]["menu"] = []protocol.HtmlMenu{
		{"product", map[string]string{`link`: "浏览" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "|product|browse"}},
		{"project", map[string]string{`link`: "浏览" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "|project|browse"}},
	}
	Lang[protocol.ZH_CN]["my"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["my"]["availableDays"] = "可用天数"
	Lang[protocol.ZH_CN]["my"]["applicationDays"] = "申请天数"
	Lang[protocol.ZH_CN]["my"]["reviewModule"] = "模块"
	Lang[protocol.ZH_CN]["my"]["review"] = "审批"
	Lang[protocol.ZH_CN]["my"]["effort"] = "我的日志"
	Lang[protocol.ZH_CN]["my"]["noTodo"] = "暂时没有待办。"
	Lang[protocol.ZH_CN]["my"]["scoreRule"] = "积分规则"
	Lang[protocol.ZH_CN]["my"]["score"] = "我的积分"
	Lang[protocol.ZH_CN]["my"]["limited"] = "受限操作(只能编辑与自己相关的内容)"
	Lang[protocol.ZH_CN]["my"]["shareContacts"] = "共享联系人列表"
	Lang[protocol.ZH_CN]["my"]["deleteContacts"] = "删除联系人"
	Lang[protocol.ZH_CN]["my"]["manageContacts"] = "维护联系人"
	Lang[protocol.ZH_CN]["my"]["unbind"] = "解除ZDOO绑定"
	Lang[protocol.ZH_CN]["my"]["changePassword"] = "修改密码"
	Lang[protocol.ZH_CN]["my"]["editProfile"] = "修改档案"
	Lang[protocol.ZH_CN]["my"]["dynamic"] = "我的动态"
	Lang[protocol.ZH_CN]["my"]["profile"] = "我的档案"
	Lang[protocol.ZH_CN]["my"]["myProject"] = "我的" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["my"]["story"] = "我的" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string)
	Lang[protocol.ZH_CN]["my"]["testCase"] = "我的用例"
	Lang[protocol.ZH_CN]["my"]["testTask"] = "我的版本"
	Lang[protocol.ZH_CN]["my"]["bug"] = "我的Bug"
	Lang[protocol.ZH_CN]["my"]["task"] = "我的任务"
	Lang[protocol.ZH_CN]["my"]["calendar"] = "日程"
	Lang[protocol.ZH_CN]["my"]["todo"] = "我的待办"
	Lang[protocol.ZH_CN]["my"]["index"] = "首页"
	Lang[protocol.ZH_CN]["my"]["common"] = "我的地盘"
	Lang[protocol.ZH_CN]["my"]["dividerMenu"] = []string{"task", "myProject", "profile"}
	Lang[protocol.ZH_CN]["my"]["menu"] = []protocol.HtmlMenu{
		{"index", map[string]string{`link`: "首页|my|index"}},
		{"calendar", map[string]string{`link`: `日程|my|calendar|`, `subModule`: `todo,effort`, `alias`: `todo`, `class`: `dropdown dropdown-hover`}},
		{"task", map[string]string{`link`: `任务|my|task|`, `subModule`: `task`}},
		{"bug", map[string]string{`link`: `Bug|my|bug|`, `subModule`: `bug`}},
		{"testtask", map[string]string{`link`: `测试|my|testtask|`, `subModule`: `testcase,testtask`, `alias`: `testcase`}},
		{"story", map[string]string{`link`: Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "|my|story|", `subModule`: `story`}},
		{"myProject", map[string]string{`link`: Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "|my|project|"}},
		//{"dynamic", map[string]string{`link`: "动态|my|dynamic|"}},
		{"profile", map[string]string{`link`: `档案|my|profile`, `alias`: `editprofile`}},
		{"changePassword", map[string]string{`link`: "密码|my|changepassword"}},
		{"manageContacts", map[string]string{`link`: "联系人|my|managecontacts"}},
		//{"score", map[string]string{`link`: `积分|my|score`, `subModule`: `score`}},
		{"review", map[string]string{`link`: "审批|my|review|type=all"}},
	}
	Lang[protocol.ZH_CN]["todo"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["todo"]["noAssignedTo"] = "被指派人不能为空"
	Lang[protocol.ZH_CN]["todo"]["noTodo"] = "该类型没有待办事务"
	Lang[protocol.ZH_CN]["todo"]["lblClickCreate"] = "点击添加待办"
	Lang[protocol.ZH_CN]["todo"]["lblBeforeDays"] = "提前%s天生成待办"
	Lang[protocol.ZH_CN]["todo"]["lblDisableDate"] = "暂时不设定时间"
	Lang[protocol.ZH_CN]["todo"]["thisIsPrivate"] = "这是一条私人事务。:)"
	Lang[protocol.ZH_CN]["todo"]["confirmDelete"] = "您确定要删除这条待办吗？"
	Lang[protocol.ZH_CN]["todo"]["confirmStory"] = "该Todo关联的是Story #%s，需要修改它吗？"
	Lang[protocol.ZH_CN]["todo"]["confirmTask"] = "该Todo关联的是Task #%s，需要修改它吗？"
	Lang[protocol.ZH_CN]["todo"]["confirmBug"] = "该Todo关联的是Bug #%s，需要修改它吗？"

	Lang[protocol.ZH_CN]["todo"]["beforeDays"] = "<span class='input-group-addon'>提前</span>%s<span class='input-group-addon'>天生成待办</span>"
	Lang[protocol.ZH_CN]["todo"]["every"] = "间隔"
	Lang[protocol.ZH_CN]["todo"]["deadline"] = "过期时间"
	Lang[protocol.ZH_CN]["todo"]["cycleMonth"] = "月"
	Lang[protocol.ZH_CN]["todo"]["cycleWeek"] = "周"
	Lang[protocol.ZH_CN]["todo"]["cycleDay"] = "天"
	Lang[protocol.ZH_CN]["todo"]["private"] = "私人事务"
	Lang[protocol.ZH_CN]["todo"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["todo"]["status"] = "状态"
	Lang[protocol.ZH_CN]["todo"]["name"] = "待办名称"
	Lang[protocol.ZH_CN]["todo"]["pri"] = "优先级"
	Lang[protocol.ZH_CN]["todo"]["type"] = "类型"
	Lang[protocol.ZH_CN]["todo"]["idvalue"] = "关联编号"
	Lang[protocol.ZH_CN]["todo"]["beginAndEnd"] = "起止时间"
	Lang[protocol.ZH_CN]["todo"]["endAB"] = "结束"
	Lang[protocol.ZH_CN]["todo"]["beginAB"] = "开始"
	Lang[protocol.ZH_CN]["todo"]["end"] = "结束"
	Lang[protocol.ZH_CN]["todo"]["begin"] = "开始"
	Lang[protocol.ZH_CN]["todo"]["date"] = "日期"
	Lang[protocol.ZH_CN]["todo"]["account"] = "所有者"
	Lang[protocol.ZH_CN]["todo"]["id"] = "编号"
	Lang[protocol.ZH_CN]["todo"]["cycleConfig"] = "周期设置"
	Lang[protocol.ZH_CN]["todo"]["cycle"] = "周期"
	Lang[protocol.ZH_CN]["todo"]["legendBasic"] = "基本信息"
	Lang[protocol.ZH_CN]["todo"]["import"] = "导入"
	Lang[protocol.ZH_CN]["todo"]["import2Today"] = "导入到今天"
	Lang[protocol.ZH_CN]["todo"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["todo"]["export"] = "导出"
	Lang[protocol.ZH_CN]["todo"]["batchFinish"] = "批量完成"
	Lang[protocol.ZH_CN]["todo"]["finish"] = "完成待办"
	Lang[protocol.ZH_CN]["todo"]["view"] = "待办详情"
	Lang[protocol.ZH_CN]["todo"]["batchEdit"] = "批量编辑"
	Lang[protocol.ZH_CN]["todo"]["batchClose"] = "批量关闭"
	Lang[protocol.ZH_CN]["todo"]["close"] = "关闭"
	Lang[protocol.ZH_CN]["todo"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["todo"]["batchCreate"] = "批量添加"
	Lang[protocol.ZH_CN]["todo"]["activate"] = "激活"
	Lang[protocol.ZH_CN]["todo"]["assignTo"] = "指派给"
	Lang[protocol.ZH_CN]["todo"]["createCycle"] = "创建周期待办"
	Lang[protocol.ZH_CN]["todo"]["create"] = "添加待办"
	Lang[protocol.ZH_CN]["todo"]["index"] = "待办一览"
	Lang[protocol.ZH_CN]["todo"]["common"] = "待办"
	Lang[protocol.ZH_CN]["todo"]["menu"] = Lang[protocol.ZH_CN]["my"]["menu"]
	Lang[protocol.ZH_CN]["score"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["score"]["menu"] = Lang[protocol.ZH_CN]["my"]["menu"]
	Lang[protocol.ZH_CN]["product"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["product"]["noMatched"] = `找不到包含"%s"的` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["product"]["noProduct"] = "暂时没有产品。"
	Lang[protocol.ZH_CN]["product"]["noModule"] = template.HTML("<div>您现在还没有模块信息</div><div>请维护产品模块</div>")
	Lang[protocol.ZH_CN]["product"]["checkedSummary"] = "选中 <strong>%total%</strong> 个需求，预计 <strong>%estimate%</strong> 个工时，用例覆盖率 <strong>%rate%</strong>。"
	Lang[protocol.ZH_CN]["product"]["storySummary"] = "本页共 <strong>%s</strong> 个需求，预计 <strong>%s</strong> 个工时，用例覆盖率 <strong>%s</strong>。"
	Lang[protocol.ZH_CN]["product"]["allProductsOfProject"] = `全部关联` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["product"]["allProduct"] = `全部` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["product"]["allStory"] = "所有"
	Lang[protocol.ZH_CN]["product"]["unplan"] = "未计划"
	Lang[protocol.ZH_CN]["product"]["unclosed"] = "未关闭"
	Lang[protocol.ZH_CN]["product"]["closedStory"] = "已关闭"
	Lang[protocol.ZH_CN]["product"]["willClose"] = "待关闭"
	Lang[protocol.ZH_CN]["product"]["changedStory"] = "已变更"
	Lang[protocol.ZH_CN]["product"]["activeStory"] = "激活"
	Lang[protocol.ZH_CN]["product"]["draftStory"] = "草稿"
	Lang[protocol.ZH_CN]["product"]["closedByMe"] = "我关闭"
	Lang[protocol.ZH_CN]["product"]["reviewedByMe"] = "我评审"
	Lang[protocol.ZH_CN]["product"]["openedByMe"] = "我创建"
	Lang[protocol.ZH_CN]["product"]["assignedToMe"] = "指给我"
	Lang[protocol.ZH_CN]["product"]["searchStory"] = "搜索"
	Lang[protocol.ZH_CN]["product"]["iterationView"] = "查看详情"
	Lang[protocol.ZH_CN]["product"]["iterationInfo"] = "迭代 %s 次"
	Lang[protocol.ZH_CN]["product"]["iteration"] = "版本迭代"
	Lang[protocol.ZH_CN]["product"]["plan"] = "计划"
	Lang[protocol.ZH_CN]["product"]["latestDynamic"] = "最新动态"
	Lang[protocol.ZH_CN]["product"]["maintain"] = "维护中"
	Lang[protocol.ZH_CN]["product"]["allRelease"] = "所有发布"
	Lang[protocol.ZH_CN]["product"]["release"] = "发布"
	Lang[protocol.ZH_CN]["product"]["qa"] = "测试"
	Lang[protocol.ZH_CN]["product"]["branch"] = "所属%s"
	Lang[protocol.ZH_CN]["product"]["whitelist"] = "分组白名单"
	Lang[protocol.ZH_CN]["product"]["acl"] = "访问控制"
	Lang[protocol.ZH_CN]["product"]["RD"] = "发布负责人"
	Lang[protocol.ZH_CN]["product"]["QD"] = "测试负责人"
	Lang[protocol.ZH_CN]["product"]["PO"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "负责人"
	Lang[protocol.ZH_CN]["product"]["manager"] = "负责人"
	Lang[protocol.ZH_CN]["product"]["desc"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "描述"
	Lang[protocol.ZH_CN]["product"]["status"] = "状态"
	Lang[protocol.ZH_CN]["product"]["type"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "类型"
	Lang[protocol.ZH_CN]["product"]["order"] = "排序"
	Lang[protocol.ZH_CN]["product"]["line"] = "产品线"
	Lang[protocol.ZH_CN]["product"]["code"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "代号"
	Lang[protocol.ZH_CN]["product"]["name"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "名称"
	Lang[protocol.ZH_CN]["product"]["id"] = "编号"
	Lang[protocol.ZH_CN]["product"]["accessDenied"] = "您无权访问该" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["product"]["errorNoProduct"] = "还没有创建" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "！"
	Lang[protocol.ZH_CN]["product"]["confirmDelete"] = " 您确定删除该" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "吗？"
	Lang[protocol.ZH_CN]["product"]["assignToNullBugs"] = "未指派Bug"
	Lang[protocol.ZH_CN]["product"]["unResolvedBugs"] = "未解决Bug"
	Lang[protocol.ZH_CN]["product"]["closedStories"] = "已关闭需求"
	Lang[protocol.ZH_CN]["product"]["draftStories"] = "草稿需求"
	Lang[protocol.ZH_CN]["product"]["changedStories"] = "已变更需求"
	Lang[protocol.ZH_CN]["product"]["activeStories"] = "激活需求"
	Lang[protocol.ZH_CN]["product"]["currentProject"] = `当前` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["product"]["build"] = "版本列表"
	Lang[protocol.ZH_CN]["product"]["project"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `列表`
	Lang[protocol.ZH_CN]["product"]["doc"] = "文档列表"
	Lang[protocol.ZH_CN]["product"]["roadmap"] = "路线图"
	Lang[protocol.ZH_CN]["product"]["builds"] = "BUILD数"
	Lang[protocol.ZH_CN]["product"]["cases"] = "用例数"
	Lang[protocol.ZH_CN]["product"]["projects"] = "关联" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "数"
	Lang[protocol.ZH_CN]["product"]["bugs"] = "相关Bug"
	Lang[protocol.ZH_CN]["product"]["docs"] = "文档数"
	Lang[protocol.ZH_CN]["product"]["releases"] = "发布数"
	Lang[protocol.ZH_CN]["product"]["plans"] = "计划数"
	Lang[protocol.ZH_CN]["product"]["otherInfo"] = "其他信息"
	Lang[protocol.ZH_CN]["product"]["basicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["product"]["export"] = "导出数据"
	Lang[protocol.ZH_CN]["product"]["all"] = "所有" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["product"]["updateOrder"] = "排序"
	Lang[protocol.ZH_CN]["product"]["closed"] = "已关闭"
	Lang[protocol.ZH_CN]["product"]["other"] = "其他："
	Lang[protocol.ZH_CN]["product"]["mine"] = "我负责："
	Lang[protocol.ZH_CN]["product"]["select"] = "请选择" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["product"]["close"] = "关闭"
	Lang[protocol.ZH_CN]["product"]["deleted"] = "已删除"
	Lang[protocol.ZH_CN]["product"]["delete"] = "删除" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["product"]["create"] = "添加" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["product"]["batchEdit"] = "批量编辑"
	Lang[protocol.ZH_CN]["product"]["edit"] = "编辑" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["product"]["view"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "概况"
	Lang[protocol.ZH_CN]["product"]["dynamic"] = "动态"
	Lang[protocol.ZH_CN]["product"]["browse"] = "需求列表"
	Lang[protocol.ZH_CN]["product"]["index"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `主页`
	Lang[protocol.ZH_CN]["product"]["common"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `视图`
	Lang[protocol.ZH_CN]["product"]["dividerMenu"] = []string{"plan", "project", "doc"}
	Lang[protocol.ZH_CN]["product"]["menu"] = []protocol.HtmlMenu{
		{"story", map[string]string{`link`: Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "|product|browse|productID=%s&branch=%s", `alias`: `batchedit`, `subModule`: `story`}},
		{"plan", map[string]string{`link`: `计划|productplan|browse|productID=%s&branch=%s`, `subModule`: `productplan`}},
		{"release", map[string]string{`link`: `发布|release|browse|productID=%s&branch=%s`, `subModule`: `release`}},
		{"roadmap", map[string]string{`link`: "路线图|product|roadmap|productID=%s&branch=%s"}},
		{"project", map[string]string{`link`: Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "|product|project|status=all&productID=%s&branch=%s"}},
		//{"dynamic", map[string]string{`link`: "动态|product|dynamic|productID=%s&branch=%s"}},
		{"doc", map[string]string{`link`: `文档|doc|objectLibs|type=product&objectID=%s&from=product`, `subModule`: `doc`}},
		//{"branch", map[string]string{`link`: "@branch@|branch|manage|productID=%s&branch=%s"}},
		{"module", map[string]string{`link`: "模块|tree|browse|productID=%s&branch=%s&&view=story"}},
		{"view", map[string]string{`link`: `概况|product|view|productID=%s&branch=%s`, `alias`: `edit`}},
	}
	Lang[protocol.ZH_CN]["branch"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["branch"]["canNotDelete"] = map[string]interface{}{
		"BranchCanNotDeletePROJECTHasData":     "该@branch@下 项目有数据，不能删除！",
		"BranchCanNotDeleteSTORYHasData":       "该@branch@下 需求有数据，不能删除！",
		"BranchCanNotDeleteMODULEHasData":      "该@branch@下 模块有数据，不能删除！",
		"BranchCanNotDeletePRODUCTPLANHasData": "该@branch@下 产品计划有数据，不能删除！",
		"BranchCanNotDeleteBUGHasData":         "该@branch@下 BUG有数据，不能删除！",
		"BranchCanNotDeleteCASEHasData":        "该@branch@下 测试有数据，不能删除！",
		"BranchCanNotDeleteRELEASEHasData":     "该@branch@下 产品发布 有数据，不能删除！",
		"BranchCanNotDeleteBUILDHasData":       "该@branch@下 测试版本有数据，不能删除！",
	}
	Lang[protocol.ZH_CN]["branch"]["confirmDelete"] = "是否删除该@branch@？"
	Lang[protocol.ZH_CN]["branch"]["order"] = "排序"
	Lang[protocol.ZH_CN]["branch"]["name"] = "名称"
	Lang[protocol.ZH_CN]["branch"]["product"] = "所属产品"
	Lang[protocol.ZH_CN]["branch"]["id"] = "编号"
	Lang[protocol.ZH_CN]["branch"]["all"] = "所有"
	Lang[protocol.ZH_CN]["branch"]["manageTitle"] = "%s管理"
	Lang[protocol.ZH_CN]["branch"]["add"] = "添加"
	Lang[protocol.ZH_CN]["branch"]["delete"] = "分支删除"
	Lang[protocol.ZH_CN]["branch"]["sort"] = "排序"
	Lang[protocol.ZH_CN]["branch"]["manage"] = "分支管理"
	Lang[protocol.ZH_CN]["branch"]["common"] = "分支"
	Lang[protocol.ZH_CN]["branch"]["menu"] = Lang[protocol.ZH_CN]["product"]["menu"]
	Lang[protocol.ZH_CN]["story"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["story"]["ignoreChangeStage"] = "需求 %s 为草稿状态，没有修改其阶段。"
	Lang[protocol.ZH_CN]["story"]["noStory"] = "暂时没有需求。"
	Lang[protocol.ZH_CN]["story"]["mustChoosePreVersion"] = "必须选择回溯的版本"
	Lang[protocol.ZH_CN]["story"]["mustChooseResult"] = "必须选择评审结果"
	Lang[protocol.ZH_CN]["story"]["errorEmptyChildStory"] = "『细分需求』不能为空。"
	Lang[protocol.ZH_CN]["story"]["confirmDelete"] = "您确认删除该需求吗?"
	Lang[protocol.ZH_CN]["story"]["successSaved"] = "需求成功添加，"
	Lang[protocol.ZH_CN]["story"]["needNotReview"] = "不需要评审"
	Lang[protocol.ZH_CN]["story"]["specTemplate"] = "建议参考的模板：作为一名<某种类型的用户>，我希望<达成某些目的>，这样可以<开发的价值>。"
	Lang[protocol.ZH_CN]["story"]["affectedCases"] = "影响的用例"
	Lang[protocol.ZH_CN]["story"]["affectedBugs"] = "影响的Bug"
	Lang[protocol.ZH_CN]["story"]["affectedProjects"] = `影响的` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["story"]["checkAffection"] = "影响范围"
	Lang[protocol.ZH_CN]["story"]["lblTBC"] = "任务Bug用例"
	Lang[protocol.ZH_CN]["story"]["lblClose"] = "关闭需求"
	Lang[protocol.ZH_CN]["story"]["lblActivate"] = "激活需求"
	Lang[protocol.ZH_CN]["story"]["lblReview"] = "评审需求"
	Lang[protocol.ZH_CN]["story"]["lblChange"] = "变更需求"
	Lang[protocol.ZH_CN]["story"]["legendMisc"] = "其他相关"
	Lang[protocol.ZH_CN]["story"]["legendVerify"] = "验收标准"
	Lang[protocol.ZH_CN]["story"]["legendSpec"] = "需求描述"
	Lang[protocol.ZH_CN]["story"]["legendChildStories"] = "细分需求"
	Lang[protocol.ZH_CN]["story"]["legendLinkStories"] = "相关需求"
	Lang[protocol.ZH_CN]["story"]["legendCases"] = "相关用例"
	Lang[protocol.ZH_CN]["story"]["legendFromBug"] = "来源Bug"
	Lang[protocol.ZH_CN]["story"]["legendBugs"] = "相关Bug"
	Lang[protocol.ZH_CN]["story"]["legendProjectAndTask"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `任务`
	Lang[protocol.ZH_CN]["story"]["legendAttatch"] = "附件"
	Lang[protocol.ZH_CN]["story"]["legendMailto"] = "抄送给"
	Lang[protocol.ZH_CN]["story"]["legendRelated"] = "相关信息"
	Lang[protocol.ZH_CN]["story"]["legendLifeTime"] = "需求的一生"
	Lang[protocol.ZH_CN]["story"]["legendBasicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["story"]["dittoNotice"] = "该需求与上一需求不属于同一产品！"
	Lang[protocol.ZH_CN]["story"]["ditto"] = "同上"
	Lang[protocol.ZH_CN]["story"]["deleted"] = "已删除"
	Lang[protocol.ZH_CN]["story"]["unclosed"] = "未关闭"
	Lang[protocol.ZH_CN]["story"]["allStories"] = "所有需求"
	Lang[protocol.ZH_CN]["story"]["total"] = "总需求"
	Lang[protocol.ZH_CN]["story"]["copy"] = "复制需求"
	Lang[protocol.ZH_CN]["story"]["files"] = "附件"
	Lang[protocol.ZH_CN]["story"]["colorTag"] = "颜色标签"
	Lang[protocol.ZH_CN]["story"]["newStory"] = "继续添加需求"
	Lang[protocol.ZH_CN]["story"]["keywords"] = "关键词"
	Lang[protocol.ZH_CN]["story"]["preVersion"] = "之前版本"
	Lang[protocol.ZH_CN]["story"]["reviewResult"] = "评审结果"
	Lang[protocol.ZH_CN]["story"]["duplicateStory"] = "重复需求"
	Lang[protocol.ZH_CN]["story"]["childStories"] = "细分需求"
	Lang[protocol.ZH_CN]["story"]["linkStories"] = "相关需求"
	Lang[protocol.ZH_CN]["story"]["comment"] = "备注"
	Lang[protocol.ZH_CN]["story"]["planAB"] = "计划"
	Lang[protocol.ZH_CN]["story"]["plan"] = "所属计划"
	Lang[protocol.ZH_CN]["story"]["version"] = "版本号"
	Lang[protocol.ZH_CN]["story"]["reviewedDate"] = "评审时间"
	Lang[protocol.ZH_CN]["story"]["reviewedBy"] = "由谁评审"
	Lang[protocol.ZH_CN]["story"]["rejectedReason"] = "拒绝原因"
	Lang[protocol.ZH_CN]["story"]["closedReason"] = "关闭原因"
	Lang[protocol.ZH_CN]["story"]["closedDate"] = "关闭日期"
	Lang[protocol.ZH_CN]["story"]["closedBy"] = "由谁关闭"
	Lang[protocol.ZH_CN]["story"]["lastEditedDate"] = "最后修改日期"
	Lang[protocol.ZH_CN]["story"]["lastEditedBy"] = "最后修改"
	Lang[protocol.ZH_CN]["story"]["assignedDate"] = "指派日期"
	Lang[protocol.ZH_CN]["story"]["assignedTo"] = "指派给"
	Lang[protocol.ZH_CN]["story"]["openedDate"] = "创建日期"
	Lang[protocol.ZH_CN]["story"]["openedBy"] = "由谁创建"
	Lang[protocol.ZH_CN]["story"]["mailto"] = "抄送给"
	Lang[protocol.ZH_CN]["story"]["stageAB"] = "阶段"
	Lang[protocol.ZH_CN]["story"]["stage"] = "所处阶段"
	Lang[protocol.ZH_CN]["story"]["status"] = "当前状态"
	Lang[protocol.ZH_CN]["story"]["hour"] = "小时"
	Lang[protocol.ZH_CN]["story"]["estimateAB"] = "预计"
	Lang[protocol.ZH_CN]["story"]["estimate"] = "预计工时"
	Lang[protocol.ZH_CN]["story"]["pri"] = "优先级"
	Lang[protocol.ZH_CN]["story"]["verify"] = "验收标准"
	Lang[protocol.ZH_CN]["story"]["spec"] = "需求描述"
	Lang[protocol.ZH_CN]["story"]["title"] = "需求名称"
	Lang[protocol.ZH_CN]["story"]["fromBug"] = "来源Bug"
	Lang[protocol.ZH_CN]["story"]["sourceNote"] = "来源备注"
	Lang[protocol.ZH_CN]["story"]["source"] = "需求来源"
	Lang[protocol.ZH_CN]["story"]["moduleAB"] = "更改模块"
	Lang[protocol.ZH_CN]["story"]["module"] = "所属模块"
	Lang[protocol.ZH_CN]["story"]["branch"] = "分支/平台"
	Lang[protocol.ZH_CN]["story"]["product"] = "所属" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["story"]["id"] = "编号"
	Lang[protocol.ZH_CN]["story"]["common"] = "需求"
	Lang[protocol.ZH_CN]["story"]["viewAll"] = "查看全部"
	Lang[protocol.ZH_CN]["story"]["batchChangeModule"] = "批量修改模块"
	Lang[protocol.ZH_CN]["story"]["batchAssignTo"] = "批量指派"
	Lang[protocol.ZH_CN]["story"]["batchChangeStage"] = "批量修改阶段"
	Lang[protocol.ZH_CN]["story"]["batchChangeBranch"] = "批量修改分支"
	Lang[protocol.ZH_CN]["story"]["batchChangePlan"] = "批量修改计划"
	Lang[protocol.ZH_CN]["story"]["copyTitle"] = "同需求名称"
	Lang[protocol.ZH_CN]["story"]["reportChart"] = "统计报表"
	Lang[protocol.ZH_CN]["story"]["zeroTask"] = "只列零任务需求"
	Lang[protocol.ZH_CN]["story"]["zeroCase"] = "零用例需求"
	Lang[protocol.ZH_CN]["story"]["exportfile"] = "导出文件"
	Lang[protocol.ZH_CN]["story"]["export"] = "导出数据"
	Lang[protocol.ZH_CN]["story"]["unlinkStory"] = "移除相关需求"
	Lang[protocol.ZH_CN]["story"]["linkStory"] = "关联需求"
	Lang[protocol.ZH_CN]["story"]["caseCountAB"] = "C"
	Lang[protocol.ZH_CN]["story"]["bugCountAB"] = "B"
	Lang[protocol.ZH_CN]["story"]["taskCountAB"] = "T"
	Lang[protocol.ZH_CN]["story"]["caseCount"] = "用例数"
	Lang[protocol.ZH_CN]["story"]["bugCount"] = "Bug数"
	Lang[protocol.ZH_CN]["story"]["taskCount"] = "任务数"
	Lang[protocol.ZH_CN]["story"]["cases"] = "相关用例"
	Lang[protocol.ZH_CN]["story"]["bugs"] = "相关Bug"
	Lang[protocol.ZH_CN]["story"]["tasks"] = "相关任务"
	Lang[protocol.ZH_CN]["story"]["setting"] = "设置"
	Lang[protocol.ZH_CN]["story"]["view"] = "需求详情"
	Lang[protocol.ZH_CN]["story"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["story"]["activate"] = "激活"
	Lang[protocol.ZH_CN]["story"]["batchClose"] = "批量关闭"
	Lang[protocol.ZH_CN]["story"]["close"] = "关闭"
	Lang[protocol.ZH_CN]["story"]["subdivide"] = "细分"
	Lang[protocol.ZH_CN]["story"]["batchEdit"] = "批量编辑"
	Lang[protocol.ZH_CN]["story"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["story"]["batchReview"] = "批量评审"
	Lang[protocol.ZH_CN]["story"]["review"] = "评审"
	Lang[protocol.ZH_CN]["story"]["changed"] = "需求变更"
	Lang[protocol.ZH_CN]["story"]["change"] = "变更"
	Lang[protocol.ZH_CN]["story"]["batchCreate"] = "批量创建"
	Lang[protocol.ZH_CN]["story"]["create"] = "提需求"
	Lang[protocol.ZH_CN]["story"]["menu"] = Lang[protocol.ZH_CN]["product"]["menu"]
	Lang[protocol.ZH_CN]["productplan"] = make(map[string]interface{})

	Lang[protocol.ZH_CN]["productplan"]["project"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["productplan"]["hour"] = "工时"
	Lang[protocol.ZH_CN]["productplan"]["bugs"] = "Bug数"
	Lang[protocol.ZH_CN]["productplan"]["stories"] = "需求数"
	Lang[protocol.ZH_CN]["productplan"]["future"] = "待定"
	Lang[protocol.ZH_CN]["productplan"]["last"] = "上次计划"
	Lang[protocol.ZH_CN]["productplan"]["end"] = "结束日期"
	Lang[protocol.ZH_CN]["productplan"]["begin"] = "开始日期"
	Lang[protocol.ZH_CN]["productplan"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["productplan"]["title"] = "名称"
	Lang[protocol.ZH_CN]["productplan"]["branch"] = "平台/分支"
	Lang[protocol.ZH_CN]["productplan"]["product"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["productplan"]["id"] = "编号"
	Lang[protocol.ZH_CN]["productplan"]["noPlan"] = "暂时没有计划。"
	Lang[protocol.ZH_CN]["productplan"]["confirmUnlinkBug"] = "您确认移除该Bug吗？"
	Lang[protocol.ZH_CN]["productplan"]["confirmUnlinkStory"] = "您确认移除该需求吗？"
	Lang[protocol.ZH_CN]["productplan"]["confirmDelete"] = "您确认删除该计划吗？"
	Lang[protocol.ZH_CN]["productplan"]["all"] = "所有计划"
	Lang[protocol.ZH_CN]["productplan"]["unexpired"] = "未过期计划"
	Lang[protocol.ZH_CN]["productplan"]["unlinkedBugs"] = "未关联Bug"
	Lang[protocol.ZH_CN]["productplan"]["linkedBugs"] = "Bug"
	Lang[protocol.ZH_CN]["productplan"]["batchUnlinkBug"] = "批量移除Bug"
	Lang[protocol.ZH_CN]["productplan"]["unlinkBug"] = "移除Bug"
	Lang[protocol.ZH_CN]["productplan"]["linkBug"] = "关联Bug"
	Lang[protocol.ZH_CN]["productplan"]["updateOrder"] = "排序"
	Lang[protocol.ZH_CN]["productplan"]["unlinkedStories"] = "未关联需求"
	Lang[protocol.ZH_CN]["productplan"]["linkedStories"] = "需求"
	Lang[protocol.ZH_CN]["productplan"]["batchUnlinkStory"] = "批量移除需求"
	Lang[protocol.ZH_CN]["productplan"]["unlinkStory"] = "移除需求"
	Lang[protocol.ZH_CN]["productplan"]["linkStory"] = "关联需求"
	Lang[protocol.ZH_CN]["productplan"]["batchUnlink"] = "批量移除"
	Lang[protocol.ZH_CN]["productplan"]["batchEdit"] = "批量编辑"
	Lang[protocol.ZH_CN]["productplan"]["basicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["productplan"]["bugSummary"] = "本页共 <strong>%s</strong> 个Bug"
	Lang[protocol.ZH_CN]["productplan"]["view"] = "计划详情"
	Lang[protocol.ZH_CN]["productplan"]["delete"] = "删除计划"
	Lang[protocol.ZH_CN]["productplan"]["edit"] = "编辑计划"
	Lang[protocol.ZH_CN]["productplan"]["create"] = "创建计划"
	Lang[protocol.ZH_CN]["productplan"]["index"] = "计划列表"
	Lang[protocol.ZH_CN]["productplan"]["browse"] = "浏览计划"
	Lang[protocol.ZH_CN]["productplan"]["common"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `计划`
	Lang[protocol.ZH_CN]["productplan"]["menu"] = Lang[protocol.ZH_CN]["product"]["menu"]
	Lang[protocol.ZH_CN]["release"] = make(map[string]interface{})

	Lang[protocol.ZH_CN]["release"]["scmPath"] = "版本库地址："
	Lang[protocol.ZH_CN]["release"]["filePath"] = "下载地址："
	Lang[protocol.ZH_CN]["release"]["yesterday"] = "昨日发布"
	Lang[protocol.ZH_CN]["release"]["export"] = "导出HTML"
	Lang[protocol.ZH_CN]["release"]["createdBugs"] = "本次共遗留 %s 个Bug"
	Lang[protocol.ZH_CN]["release"]["resolvedBugs"] = "本次共解决 %s 个Bug"
	Lang[protocol.ZH_CN]["release"]["finishStories"] = "本次共完成 %s 个需求"
	Lang[protocol.ZH_CN]["release"]["generatedBugs"] = "遗留的Bug"
	Lang[protocol.ZH_CN]["release"]["leftBugs"] = "遗留的Bug"
	Lang[protocol.ZH_CN]["release"]["bugs"] = "解决的Bug"
	Lang[protocol.ZH_CN]["release"]["stories"] = "完成的需求"
	Lang[protocol.ZH_CN]["release"]["unlinkBug"] = "移除Bug"
	Lang[protocol.ZH_CN]["release"]["unlinkStory"] = "移除需求"
	Lang[protocol.ZH_CN]["release"]["last"] = "上次发布"
	Lang[protocol.ZH_CN]["release"]["status"] = "状态"
	Lang[protocol.ZH_CN]["release"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["release"]["date"] = "发布日期"
	Lang[protocol.ZH_CN]["release"]["marker"] = "里程碑"
	Lang[protocol.ZH_CN]["release"]["name"] = "发布名称"
	Lang[protocol.ZH_CN]["release"]["build"] = "版本"
	Lang[protocol.ZH_CN]["release"]["branch"] = "平台/分支"
	Lang[protocol.ZH_CN]["release"]["product"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["release"]["id"] = "ID"
	Lang[protocol.ZH_CN]["release"]["basicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["release"]["noRelease"] = "暂时没有发布。"
	Lang[protocol.ZH_CN]["release"]["existBuild"] = "『版本』已经有『%s』这条记录了。您可以更改『发布名称』或者选择一个『版本』。"
	Lang[protocol.ZH_CN]["release"]["confirmUnlinkBug"] = "您确认移除该Bug吗？"
	Lang[protocol.ZH_CN]["release"]["confirmUnlinkStory"] = "您确认移除该需求吗？"
	Lang[protocol.ZH_CN]["release"]["confirmDelete"] = "您确认删除该发布吗？"
	Lang[protocol.ZH_CN]["release"]["batchUnlinkBug"] = "批量移除Bug"
	Lang[protocol.ZH_CN]["release"]["batchUnlinkStory"] = "批量移除需求"
	Lang[protocol.ZH_CN]["release"]["batchUnlink"] = "批量移除"
	Lang[protocol.ZH_CN]["release"]["changeStatus"] = "修改状态"
	Lang[protocol.ZH_CN]["release"]["browse"] = "浏览发布"
	Lang[protocol.ZH_CN]["release"]["view"] = "发布详情"
	Lang[protocol.ZH_CN]["release"]["deleted"] = "已删除"
	Lang[protocol.ZH_CN]["release"]["delete"] = "删除发布"
	Lang[protocol.ZH_CN]["release"]["linkBug"] = "关联Bug"
	Lang[protocol.ZH_CN]["release"]["linkStory"] = "关联需求"
	Lang[protocol.ZH_CN]["release"]["edit"] = "编辑发布"
	Lang[protocol.ZH_CN]["release"]["create"] = "创建发布"
	Lang[protocol.ZH_CN]["release"]["common"] = "发布"
	Lang[protocol.ZH_CN]["release"]["menu"] = Lang[protocol.ZH_CN]["product"]["menu"]
	Lang[protocol.ZH_CN]["project"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["project"]["ganttchart"] = "甘特图"
	Lang[protocol.ZH_CN]["project"]["viewrelation"] = "查看任务关系"
	Lang[protocol.ZH_CN]["project"]["deleterelation"] = "删除任务关系"
	Lang[protocol.ZH_CN]["project"]["maintainRelation"] = "维护任务关系"
	Lang[protocol.ZH_CN]["project"]["editrelation"] = "维护任务关系"
	Lang[protocol.ZH_CN]["project"]["setLaneFields"] = "设置看板列字段"
	Lang[protocol.ZH_CN]["project"]["computeTaskEffort"] = "更新工时"
	Lang[protocol.ZH_CN]["project"]["noStart"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "还未开始!"
	Lang[protocol.ZH_CN]["project"]["taskEffort"] = "工时明细"
	Lang[protocol.ZH_CN]["project"]["effortAction"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "日志"
	Lang[protocol.ZH_CN]["project"]["started"] = "开始了"
	Lang[protocol.ZH_CN]["project"]["finished"] = "完成了"
	Lang[protocol.ZH_CN]["project"]["month"] = "本月"
	Lang[protocol.ZH_CN]["project"]["effortCalendar"] = "日志日历"
	Lang[protocol.ZH_CN]["project"]["calendar"] = "日历"
	Lang[protocol.ZH_CN]["project"]["ftpPath"] = "ftp地址"
	Lang[protocol.ZH_CN]["project"]["showFile"] = "显示规范文件"
	Lang[protocol.ZH_CN]["project"]["specFile"] = "规范文件"
	Lang[protocol.ZH_CN]["project"]["kanbanColsColor"] = "看板列自定义颜色"
	Lang[protocol.ZH_CN]["project"]["kanbanShowOption"] = "显示折叠信息"
	Lang[protocol.ZH_CN]["project"]["kanbanHideCols"] = "看板隐藏已关闭、已取消列"
	Lang[protocol.ZH_CN]["project"]["bugList"] = "Bug列表"
	Lang[protocol.ZH_CN]["project"]["printKanban"] = "打印看板"
	Lang[protocol.ZH_CN]["project"]["resetKanban"] = "恢复默认"
	Lang[protocol.ZH_CN]["project"]["kanbanSetting"] = "看板设置"
	Lang[protocol.ZH_CN]["project"]["kanban"] = "看板"
	Lang[protocol.ZH_CN]["project"]["fixFirstWithLeft"] = "修改剩余工时"
	Lang[protocol.ZH_CN]["project"]["interval"] = "间隔"
	Lang[protocol.ZH_CN]["project"]["withweekend"] = "显示周末"
	Lang[protocol.ZH_CN]["project"]["noweekend"] = "去除周末"
	Lang[protocol.ZH_CN]["project"]["goback"] = "返回任务列表"
	Lang[protocol.ZH_CN]["project"]["createTask"] = "创建任务"
	Lang[protocol.ZH_CN]["project"]["setTeam"] = "设置团队"
	Lang[protocol.ZH_CN]["project"]["afterInfo"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "添加成功，您现在可以进行以下操作："
	Lang[protocol.ZH_CN]["project"]["tips"] = "提示"
	Lang[protocol.ZH_CN]["project"]["accessDenied"] = "您无权访问该" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "！"
	Lang[protocol.ZH_CN]["project"]["errorSameProducts"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "不能关联多个相同的" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "。"
	Lang[protocol.ZH_CN]["project"]["errorNoLinkedProducts"] = "该" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "没有关联的" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "，系统将转到" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "关联页面"
	Lang[protocol.ZH_CN]["project"]["confirmUnlinkStory"] = "您确定从该" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "中移除该需求吗？"
	Lang[protocol.ZH_CN]["project"]["confirmUnlinkMember"] = "您确定从该" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "中移除该用户吗？"
	Lang[protocol.ZH_CN]["project"]["confirmDelete"] = "您确定删除" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "[%s]吗？"
	Lang[protocol.ZH_CN]["project"]["noMembers"] = "暂时没有团队成员。"
	Lang[protocol.ZH_CN]["project"]["noProject"] = "暂时没有项目。"
	Lang[protocol.ZH_CN]["project"]["byUser"] = "按用户"
	Lang[protocol.ZH_CN]["project"]["byPeriod"] = "按时间段"
	Lang[protocol.ZH_CN]["project"]["cancelCopy"] = "取消复制"
	Lang[protocol.ZH_CN]["project"]["copyFromProject"] = "复制自" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + " <strong>%s</strong>"
	Lang[protocol.ZH_CN]["project"]["copyNoProject"] = "没有可用的" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "来复制"
	Lang[protocol.ZH_CN]["project"]["copyTeamTitle"] = "选择一个" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "团队来复制"
	Lang[protocol.ZH_CN]["project"]["copyTitle"] = "请选择一个" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "来复制"
	Lang[protocol.ZH_CN]["project"]["noMatched"] = "找不到包含`%s`的" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["copyFromTeam"] = "复制自" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "团队： <strong>%s</strong>"
	Lang[protocol.ZH_CN]["project"]["copyTeam"] = "复制团队"
	Lang[protocol.ZH_CN]["project"]["selectDeptTitle"] = "选择一个部门的成员"
	Lang[protocol.ZH_CN]["project"]["selectDept"] = "选择部门"
	Lang[protocol.ZH_CN]["project"]["doneProjects"] = "已结束"
	Lang[protocol.ZH_CN]["project"]["haveDraft"] = "有%s条草稿状态的需求无法关联到该" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["productStories"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "关联的需求是" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "需求的子集，并且只有评审通过的需求才能关联。请<a href=`%s`>关联需求</a>。"
	Lang[protocol.ZH_CN]["project"]["whyNoStories"] = "看起来没有需求可以关联。请检查下" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "关联的" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "中有没有需求，而且要确保它们已经审核通过。"

	Lang[protocol.ZH_CN]["project"]["batchWBS"] = "批量分解"
	Lang[protocol.ZH_CN]["project"]["wbs"] = "分解任务"
	Lang[protocol.ZH_CN]["project"]["groupSummaryAB"] = "<div>总任务 <strong>%s : </strong><span class='text-muted'>未开始</span> %s &nbsp; <span class='text-muted'>进行中</span> %s</div><div>总预计 <strong>%s : </strong><span class='text-muted'>已消耗</span> %s &nbsp; <span class='text-muted'>剩余</span> %s</div>"
	Lang[protocol.ZH_CN]["project"]["timeSummary"] = `<div class="table-col"><div class="clearfix segments"><div class="segment"><div class="segment-title">总预计</div><div class="segment-value">%s</div></div><div class="segment"><div class="segment-title">已消耗</div><div class="segment-value text-red">%s</div></div><div class="segment"><div class="segment-title">剩余</div><div class="segment-value">%s</div></div></div></div>`
	Lang[protocol.ZH_CN]["project"]["countSummary"] = `<div class="table-col"><div class="clearfix segments"><div class="segment"><div class="segment-title">总任务</div><div class="segment-value">%s</div></div><div class="segment"><div class="segment-title">进行中</div><div class="segment-value"><span class="label label-dot label-primary"></span> %s</div></div><div class="segment"><div class="segment-title">未开始</div><div class="segment-value"><span class="label label-dot label-primary muted"></span> %s</div></div></div></div>`
	Lang[protocol.ZH_CN]["project"]["memberHours"] = `<div class="table-col"><div class="clearfix segments"><div class="segment"><div class="segment-title">%s可用工时</div><div class="segment-value">%s</div></div></div></div>`
	Lang[protocol.ZH_CN]["project"]["memberHoursAB"] = "<div>%s有 <strong>%s</strong> 工时</div>"
	Lang[protocol.ZH_CN]["project"]["checkedSummary"] = "选中 <strong>%total%</strong> 个任务，未开始 <strong>%wait%</strong>，进行中 <strong>%doing%</strong>，总预计 <strong>%estimate%</strong> 工时，已消耗 <strong>%consumed%</strong> 工时，剩余 <strong>%left%</strong> 工时。"
	Lang[protocol.ZH_CN]["project"]["taskSummary"] = "本页共 <strong>%d</strong> 个任务，未开始 <strong>%d</strong>，进行中 <strong>%d</strong>，总预计 <strong>%0.0f</strong> 工时，已消耗 <strong>%0.0f</strong> 工时，剩余 <strong>%0.0f</strong> 工时。"
	Lang[protocol.ZH_CN]["project"]["stats"] = "可用工时 <strong>%s</strong> 工时，总共预计 <strong>%s</strong> 工时，已经消耗 <strong>%s</strong> 工时，预计剩余 <strong>%s</strong> 工时"
	Lang[protocol.ZH_CN]["project"]["lblStats"] = "工时统计"
	Lang[protocol.ZH_CN]["project"]["beginAndEnd"] = "起止时间"
	Lang[protocol.ZH_CN]["project"]["selectProject"] = "请选择" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["linkStoryByPlanTips"] = "此操作会将所选计划下面的需求全部关联到此项目中"
	Lang[protocol.ZH_CN]["project"]["aboveAllProject"] = "以上所有" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["aboveAllProduct"] = "以上所有" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["allProject"] = "所有" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["byQuery"] = "搜索"
	Lang[protocol.ZH_CN]["project"]["groups['']"] = "分组查看"
	Lang[protocol.ZH_CN]["project"]["myInvolved"] = "由我参与"
	Lang[protocol.ZH_CN]["project"]["assignedToMe"] = "指派给我"
	Lang[protocol.ZH_CN]["project"]["allTasks"] = "所有"
	Lang[protocol.ZH_CN]["project"]["viewAll"] = "查看所有"
	Lang[protocol.ZH_CN]["project"]["iterationInfo"] = "迭代%s次"
	Lang[protocol.ZH_CN]["project"]["iteration"] = "版本迭代"
	Lang[protocol.ZH_CN]["project"]["importPlanStory"] = "创建" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "成功！\n是否导入计划关联的相关需求？"
	Lang[protocol.ZH_CN]["project"]["storySort"] = "需求排序"
	Lang[protocol.ZH_CN]["project"]["storyKanban"] = "需求看板"
	Lang[protocol.ZH_CN]["project"]["treeOnlyStory"] = "树状图只看需求"
	Lang[protocol.ZH_CN]["project"]["treeOnlyTask"] = "树状图只看任务"
	Lang[protocol.ZH_CN]["project"]["treeStory"] = "只看需求"
	Lang[protocol.ZH_CN]["project"]["treeTask"] = "只看任务"
	Lang[protocol.ZH_CN]["project"]["tree"] = "树状图"
	Lang[protocol.ZH_CN]["project"]["updateOrder"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "排序"
	Lang[protocol.ZH_CN]["project"]["importBug"] = "导入Bug"
	Lang[protocol.ZH_CN]["project"]["importPlanStories"] = "按计划关联需求"
	Lang[protocol.ZH_CN]["project"]["importTask"] = "转入任务"
	Lang[protocol.ZH_CN]["project"]["batchUnlinkStory"] = "批量移除需求"
	Lang[protocol.ZH_CN]["project"]["unlinkStory"] = "移除需求"
	Lang[protocol.ZH_CN]["project"]["unlinkMember"] = "移除成员"
	Lang[protocol.ZH_CN]["project"]["manageMembers"] = "团队管理"
	Lang[protocol.ZH_CN]["project"]["batchEdit"] = "批量编辑"
	Lang[protocol.ZH_CN]["project"]["edit"] = "编辑" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["browse"] = "浏览" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["delete"] = "删除" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["copy"] = "复制" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["create"] = "添加" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["view"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "概况"
	Lang[protocol.ZH_CN]["project"]["unlinkedProducts"] = "未关联"
	Lang[protocol.ZH_CN]["project"]["linkedProducts"] = "已关联"
	Lang[protocol.ZH_CN]["project"]["unlinkStoryTasks"] = "未关联需求任务"
	Lang[protocol.ZH_CN]["project"]["linkPlan"] = "关联计划"
	Lang[protocol.ZH_CN]["project"]["linkStoryByPlan"] = "按照计划关联"
	Lang[protocol.ZH_CN]["project"]["linkStory"] = `关联需求` //"关联需求"替换`关联需求`
	Lang[protocol.ZH_CN]["project"]["manageProducts"] = `关联` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["doclib"] = "文档库列表"
	Lang[protocol.ZH_CN]["project"]["doc"] = "文档列表"
	Lang[protocol.ZH_CN]["project"]["team"] = "团队成员"
	Lang[protocol.ZH_CN]["project"]["fixFirst"] = "修改首天工时"
	Lang[protocol.ZH_CN]["project"]["burnData"] = "燃尽图数据"
	Lang[protocol.ZH_CN]["project"]["computeBurn"] = "更新燃尽图"
	Lang[protocol.ZH_CN]["project"]["burn"] = "燃尽图"
	Lang[protocol.ZH_CN]["project"]["testtask"] = "测试任务"
	Lang[protocol.ZH_CN]["project"]["build"] = "所有版本"
	Lang[protocol.ZH_CN]["project"]["latestDynamic"] = "最新动态"
	Lang[protocol.ZH_CN]["project"]["dynamic"] = "动态"
	Lang[protocol.ZH_CN]["project"]["bug"] = "Bug列表"
	Lang[protocol.ZH_CN]["project"]["story"] = "需求列表"
	Lang[protocol.ZH_CN]["project"]["groupTask"] = "分组浏览任务"
	Lang[protocol.ZH_CN]["project"]["task"] = "任务列表"
	Lang[protocol.ZH_CN]["project"]["index"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "主页"
	Lang[protocol.ZH_CN]["project"]["otherInfo"] = "其他信息"
	Lang[protocol.ZH_CN]["project"]["basicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["project"]["exportfile"] = "导出文件"
	Lang[protocol.ZH_CN]["project"]["export"] = "导出"
	Lang[protocol.ZH_CN]["project"]["close"] = "关闭"
	Lang[protocol.ZH_CN]["project"]["suspend"] = "挂起"
	Lang[protocol.ZH_CN]["project"]["putoff"] = "延期"
	Lang[protocol.ZH_CN]["project"]["activate"] = "激活"
	Lang[protocol.ZH_CN]["project"]["start"] = "开始"
	Lang[protocol.ZH_CN]["project"]["relatedMember"] = "相关成员"
	Lang[protocol.ZH_CN]["project"]["effort"] = `日志` //"日志"替换`日志`
	Lang[protocol.ZH_CN]["project"]["readjustTask"] = "顺延任务的起止时间"
	Lang[protocol.ZH_CN]["project"]["readjustTime"] = "调整项目起止时间"
	Lang[protocol.ZH_CN]["project"]["delayed"] = "已延期"
	Lang[protocol.ZH_CN]["project"]["deleted"] = "已删除"
	Lang[protocol.ZH_CN]["project"]["other"] = "其他："
	Lang[protocol.ZH_CN]["project"]["mine"] = "我负责："
	Lang[protocol.ZH_CN]["project"]["typeDesc"] = "运维" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "没有需求、bug、版本、测试功能，同时禁用燃尽图。"
	Lang[protocol.ZH_CN]["project"]["unclosed"] = "未关闭"
	Lang[protocol.ZH_CN]["project"]["undone"] = "未完成"
	Lang[protocol.ZH_CN]["project"]["all"] = "所有"
	Lang[protocol.ZH_CN]["project"]["createStory"] = "添加需求"
	Lang[protocol.ZH_CN]["project"]["noProduct"] = "无" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["viewBug"] = "查看bug"
	Lang[protocol.ZH_CN]["project"]["hours"] = "预计 %s 消耗 %s 剩余 %s"
	Lang[protocol.ZH_CN]["project"]["finalfile"] = "最终文件"
	Lang[protocol.ZH_CN]["project"]["progress"] = "进度"
	Lang[protocol.ZH_CN]["project"]["totalLeft"] = "剩余"
	Lang[protocol.ZH_CN]["project"]["totalConsumed"] = "消耗"
	Lang[protocol.ZH_CN]["project"]["totalEstimate"] = "预计"
	Lang[protocol.ZH_CN]["project"]["whitelist"] = "分组白名单"
	Lang[protocol.ZH_CN]["project"]["products"] = `相关` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["product"] = Lang[protocol.ZH_CN]["project"]["products"]
	Lang[protocol.ZH_CN]["project"]["order"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `排序`
	Lang[protocol.ZH_CN]["project"]["teamname"] = "团队名称"
	Lang[protocol.ZH_CN]["project"]["acl"] = "访问控制"
	Lang[protocol.ZH_CN]["project"]["release"] = "发布"
	Lang[protocol.ZH_CN]["project"]["qa"] = "测试"
	Lang[protocol.ZH_CN]["project"]["RD"] = "发布负责人"
	Lang[protocol.ZH_CN]["project"]["QD"] = "测试负责人"
	Lang[protocol.ZH_CN]["project"]["PM"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `负责人`
	Lang[protocol.ZH_CN]["project"]["PO"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `负责人`
	Lang[protocol.ZH_CN]["project"]["owner"] = "负责人"
	Lang[protocol.ZH_CN]["project"]["desc"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `描述`
	Lang[protocol.ZH_CN]["project"]["status"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `状态`
	Lang[protocol.ZH_CN]["project"]["totalDays"] = "可用工日"
	Lang[protocol.ZH_CN]["project"]["totalHours"] = "可用工时"
	Lang[protocol.ZH_CN]["project"]["workHour"] = "工时"
	Lang[protocol.ZH_CN]["project"]["day"] = "天"
	Lang[protocol.ZH_CN]["project"]["days"] = "可用工作日"
	Lang[protocol.ZH_CN]["project"]["to"] = "至"
	Lang[protocol.ZH_CN]["project"]["dateRange"] = "起始日期"
	Lang[protocol.ZH_CN]["project"]["end"] = `截止日期`   //"结束日期"替换`截止日期`
	Lang[protocol.ZH_CN]["project"]["begin"] = `开始日期` //"开始日期"替换`开始日期`
	Lang[protocol.ZH_CN]["project"]["code"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `代号`
	Lang[protocol.ZH_CN]["project"]["name"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `名称`
	Lang[protocol.ZH_CN]["project"]["type"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `类型`
	Lang[protocol.ZH_CN]["project"]["allProjects"] = `所有` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["project"]["common"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `视图`
	Lang[protocol.ZH_CN]["project"]["dividerMenu"] = []string{"story", "team", "product"}
	Lang[protocol.ZH_CN]["project"]["menu"] = []protocol.HtmlMenu{
		{"task", map[string]string{`link`: `任务|project|task|projectID=%s`, `subModule`: `task,tree`, `alias`: `importtask,importbug`}},
		{"kanban", map[string]string{`link`: `看板|project|kanban|projectID=%s`}},
		{"burn", map[string]string{`link`: `燃尽图|project|burn|projectID=%s`}},
		{"list", map[string]string{`link`: `更多|project|grouptask|projectID=%s`, `alias`: `grouptask,tree,gantt,maintainrelation,relation,calendar`, `class`: `dropdown dropdown-hover`}},
		{"story", map[string]string{`link`: Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "|project|story|projectID=%s", `subModule`: `story`, `alias`: `linkstory,storykanban`}},
		{"qa", map[string]string{`link`: `测试|project|bug|projectID=%s`, `subModule`: `bug,build,testtask`, `alias`: `build,testtask`, `class`: `dropdown dropdown-hover`}},
		{"doc", map[string]string{`link`: `文档|doc|objectLibs|type=project&objectID=%s&from=project`, `subModule`: `doc`}},
		{"action", map[string]string{`link`: `动态|project|dynamic|projectID=%s`, `subModule`: `dynamic`, `class`: `dropdown dropdown-hover`}},
		{"product", map[string]string{`link`: Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "|project|manageproducts|projectID=%s"}},
		{"team", map[string]string{`link`: `团队|project|team|projectID=%s`, `alias`: `managemembers`}},
		{"view", map[string]string{`link`: `概况|project|view|projectID=%s`, `alias`: `edit,start,suspend,putoff,close`}},
		{"effort", map[string]string{`link`: `日志|project|effortcalendar|projectID=%s`, `alias`: `effort`}},
	}
	Lang[protocol.ZH_CN]["project"]["subMenu"] = []protocol.HtmlKeyValueInterface{
		{"list", []protocol.HtmlMenu{
			{"groupTask", map[string]string{`link`: "分组视图|project|groupTask|projectID=%s"}},
			{"tree", map[string]string{`link`: "树状图|project|tree|projectID=%s"}},
			{"taskeffort", map[string]string{`link`: "工时明细表|project|taskeffort|projectID=%s"}},
			{"gantt", map[string]string{`link`: `甘特图|project|gantt|projectID=%s`, `alias`: `maintainrelation,relation`}},
			{"calendar", map[string]string{`link`: `任务日历|project|calendar|projectID=%s`, `alias`: `calendar`}},
		}},
		{"qa", []protocol.HtmlMenu{
			{"bug", map[string]string{`link`: "Bug|project|bug|projectID=%s"}},
			{"build", map[string]string{`link`: `版本|project|build|projectID=%s`, `subModule`: `build`}},
			{"testtask", map[string]string{`link`: `测试单|project|testtask|projectID=%s`, `subModule`: `testreport,testtask`}},
		}},
	}
	Lang[protocol.ZH_CN]["task"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["task"]["grandchildren"] = "孙任务"
	Lang[protocol.ZH_CN]["task"]["exportfinish"] = "导出最终文件"
	Lang[protocol.ZH_CN]["task"]["createDenied"] = "你不能在该项目添加任务"
	Lang[protocol.ZH_CN]["task"]["noTask"] = "暂时没有任务。"
	Lang[protocol.ZH_CN]["task"]["deniedNotice"] = "当前任务只有%s才可以%s。"
	Lang[protocol.ZH_CN]["task"]["commentActions"] = "%s. %s, 由 <strong>%s</strong> 添加备注。"
	Lang[protocol.ZH_CN]["task"]["noticeSaveRecord"] = "您有尚未保存的工时记录，请先将其保存。"
	Lang[protocol.ZH_CN]["task"]["noticeLinkStory"] = "没有可关联的相关需求，您可以为当前项目%s，然后%s"
	Lang[protocol.ZH_CN]["task"]["confirmTransfer"] = `"当前剩余"为0，任务将被转交，您确定吗？`
	Lang[protocol.ZH_CN]["task"]["confirmRecord"] = "\"剩余\"为0，任务将标记为\"已完成\"，您确定吗？"
	Lang[protocol.ZH_CN]["task"]["confirmFinish"] = "\"预计剩余\"为0，确认将任务状态改为\"已完成\"吗？"
	Lang[protocol.ZH_CN]["task"]["confirmChangeProject"] = "修改" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "会导致相应的所属模块、相关需求和指派人发生变化，确定吗？"
	Lang[protocol.ZH_CN]["task"]["remindBug"] = "该任务为Bug转化得到，是否更新Bug:%d ?"
	Lang[protocol.ZH_CN]["task"]["delayWarning"] = " <strong class=\"text-danger\"> 延期%v天 </strong>"
	Lang[protocol.ZH_CN]["task"]["successSaved"] = "成功添加，"
	Lang[protocol.ZH_CN]["task"]["afterSubmit"] = "添加之后"
	Lang[protocol.ZH_CN]["task"]["copyStoryTitle"] = "同需求"
	Lang[protocol.ZH_CN]["task"]["confirmDeleteEstimate"] = "您确定要删除这个记录吗？"
	Lang[protocol.ZH_CN]["task"]["confirmDelete"] = "您确定要删除这个任务吗？"
	Lang[protocol.ZH_CN]["task"]["legendDesc"] = "任务描述"
	Lang[protocol.ZH_CN]["task"]["legendLife"] = "任务的一生"
	Lang[protocol.ZH_CN]["task"]["legendEffort"] = "工时信息"
	Lang[protocol.ZH_CN]["task"]["legendBasic"] = "基本信息"

	Lang[protocol.ZH_CN]["task"]["allTasks"] = "总任务"
	Lang[protocol.ZH_CN]["task"]["yesterdayFinished"] = "昨日完成"
	Lang[protocol.ZH_CN]["task"]["noClosed"] = "未关闭"
	Lang[protocol.ZH_CN]["task"]["noFinished"] = "未完成"
	Lang[protocol.ZH_CN]["task"]["noAssigned"] = "未指派"
	Lang[protocol.ZH_CN]["task"]["noStory"] = "无需求"
	Lang[protocol.ZH_CN]["task"]["selectAllUser"] = "全部"
	Lang[protocol.ZH_CN]["task"]["dittoNotice"] = "该任务与上一任务不属于同一项目！"
	Lang[protocol.ZH_CN]["task"]["ditto"] = "同上"
	Lang[protocol.ZH_CN]["task"]["cancelPlaceOrder"] = "取消系统已下单"
	Lang[protocol.ZH_CN]["task"]["setPlaceOrder"] = "设置系统已下单"
	Lang[protocol.ZH_CN]["task"]["placeOrder"] = "系统已下单"
	Lang[protocol.ZH_CN]["task"]["finalfile"] = "最终文件"
	Lang[protocol.ZH_CN]["task"]["lblHour"] = "(h)"
	Lang[protocol.ZH_CN]["task"]["lblPri"] = "P"
	Lang[protocol.ZH_CN]["task"]["parentAB"] = "父"
	Lang[protocol.ZH_CN]["task"]["parent"] = "父任务"
	Lang[protocol.ZH_CN]["task"]["grandchildrenAB"] = "孙"
	Lang[protocol.ZH_CN]["task"]["childrenAB"] = "子"
	Lang[protocol.ZH_CN]["task"]["children"] = "子任务"
	Lang[protocol.ZH_CN]["task"]["transferTo"] = "转交给"
	Lang[protocol.ZH_CN]["task"]["transfer"] = "转交"
	Lang[protocol.ZH_CN]["task"]["team"] = "团队"
	Lang[protocol.ZH_CN]["task"]["multipleAB"] = "多人"
	Lang[protocol.ZH_CN]["task"]["multiple"] = "多人任务"
	Lang[protocol.ZH_CN]["task"]["hasConsumed"] = "已消耗"
	Lang[protocol.ZH_CN]["task"]["files"] = "图片附件"
	Lang[protocol.ZH_CN]["task"]["colorTag"] = "颜色标签"
	Lang[protocol.ZH_CN]["task"]["deleteEstimate"] = "删除工时"
	Lang[protocol.ZH_CN]["task"]["editEstimate"] = "编辑工时"
	Lang[protocol.ZH_CN]["task"]["recordEstimate"] = "工时"
	Lang[protocol.ZH_CN]["task"]["lastEdited"] = "最后编辑"
	Lang[protocol.ZH_CN]["task"]["lastEditedDate"] = "最后修改日期"
	Lang[protocol.ZH_CN]["task"]["lastEditedBy"] = "最后修改"
	Lang[protocol.ZH_CN]["task"]["closedReason"] = "关闭原因"
	Lang[protocol.ZH_CN]["task"]["closedDate"] = "关闭时间"
	Lang[protocol.ZH_CN]["task"]["closedBy"] = "由谁关闭"
	Lang[protocol.ZH_CN]["task"]["canceledDate"] = "取消时间"
	Lang[protocol.ZH_CN]["task"]["canceledBy"] = "由谁取消"
	Lang[protocol.ZH_CN]["task"]["finishedList"] = "完成者列表"
	Lang[protocol.ZH_CN]["task"]["finishedDateAB"] = "完成时间"
	Lang[protocol.ZH_CN]["task"]["finishedDate"] = "完成时间"
	Lang[protocol.ZH_CN]["task"]["finishedByAB"] = "完成者"
	Lang[protocol.ZH_CN]["task"]["finishedBy"] = "由谁完成"
	Lang[protocol.ZH_CN]["task"]["openedDateAB"] = "创建"
	Lang[protocol.ZH_CN]["task"]["openedDate"] = "创建日期"
	Lang[protocol.ZH_CN]["task"]["openedBy"] = "由谁创建"
	Lang[protocol.ZH_CN]["task"]["assignedDate"] = "指派日期"
	Lang[protocol.ZH_CN]["task"]["assignedToAB"] = "指派给"
	Lang[protocol.ZH_CN]["task"]["assignedTo"] = "指派给"
	Lang[protocol.ZH_CN]["task"]["batchAssignTo"] = "批量指派"
	Lang[protocol.ZH_CN]["task"]["assign"] = "指派"
	Lang[protocol.ZH_CN]["task"]["assignTo"] = Lang[protocol.ZH_CN]["task"]["assign"]
	Lang[protocol.ZH_CN]["task"]["desc"] = "任务描述"
	Lang[protocol.ZH_CN]["task"]["status"] = "任务状态"
	Lang[protocol.ZH_CN]["task"]["deadlineAB"] = "截止"
	Lang[protocol.ZH_CN]["task"]["deadline"] = "截止日期"
	Lang[protocol.ZH_CN]["task"]["date"] = "日期"
	Lang[protocol.ZH_CN]["task"]["examineDate"] = "审核日期"
	Lang[protocol.ZH_CN]["task"]["examineBy"] = "审核人"
	Lang[protocol.ZH_CN]["task"]["realStarted"] = "实际开始"
	Lang[protocol.ZH_CN]["task"]["estStarted"] = "预计开始"
	Lang[protocol.ZH_CN]["task"]["datePlan"] = "日程规划"
	Lang[protocol.ZH_CN]["task"]["leftThisTime"] = "剩余"
	Lang[protocol.ZH_CN]["task"]["consumedThisTime"] = "工时"
	Lang[protocol.ZH_CN]["task"]["hour"] = "小时"
	Lang[protocol.ZH_CN]["task"]["consumedAB"] = "消耗"
	Lang[protocol.ZH_CN]["task"]["myConsumed"] = "我的总消耗"
	Lang[protocol.ZH_CN]["task"]["consumed"] = "总消耗"
	Lang[protocol.ZH_CN]["task"]["leftAB"] = "剩余"
	Lang[protocol.ZH_CN]["task"]["left"] = "预计剩余"
	Lang[protocol.ZH_CN]["task"]["estimateAB"] = "预计"
	Lang[protocol.ZH_CN]["task"]["estimate"] = "最初预计"
	Lang[protocol.ZH_CN]["task"]["mailto"] = "抄送给"
	Lang[protocol.ZH_CN]["task"]["pri"] = "优先级"
	Lang[protocol.ZH_CN]["task"]["type"] = "任务类型"
	Lang[protocol.ZH_CN]["task"]["name"] = "任务名称"
	Lang[protocol.ZH_CN]["task"]["storyVerify"] = "验收标准"
	Lang[protocol.ZH_CN]["task"]["storySpec"] = "需求描述"
	Lang[protocol.ZH_CN]["task"]["storyAB"] = "需求"
	Lang[protocol.ZH_CN]["task"]["story"] = "相关需求"
	Lang[protocol.ZH_CN]["task"]["moduleAB"] = "模块"
	Lang[protocol.ZH_CN]["task"]["module"] = "所属模块"
	Lang[protocol.ZH_CN]["task"]["project"] = `所属` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["task"]["id"] = "编号"
	Lang[protocol.ZH_CN]["task"]["common"] = "任务"
	Lang[protocol.ZH_CN]["task"]["waitTask"] = "未开始的任务"
	Lang[protocol.ZH_CN]["task"]["copy"] = "复制任务"
	Lang[protocol.ZH_CN]["task"]["progressTips"] = "已消耗/(已消耗+剩余)"
	Lang[protocol.ZH_CN]["task"]["progress"] = "进度"
	Lang[protocol.ZH_CN]["task"]["confirmStoryChange"] = "确认需求变动"
	Lang[protocol.ZH_CN]["task"]["case"] = "相关用例"
	Lang[protocol.ZH_CN]["task"]["fromBug"] = "来源Bug"
	Lang[protocol.ZH_CN]["task"]["reportChart"] = "报表统计"
	Lang[protocol.ZH_CN]["task"]["exportfile"] = "导出文件"
	Lang[protocol.ZH_CN]["task"]["export"] = "导出数据"
	Lang[protocol.ZH_CN]["task"]["activate"] = "激活"
	Lang[protocol.ZH_CN]["task"]["cancel"] = "取消"
	Lang[protocol.ZH_CN]["task"]["close"] = "关闭"
	Lang[protocol.ZH_CN]["task"]["pause"] = "暂停"
	Lang[protocol.ZH_CN]["task"]["finish"] = "完成"
	Lang[protocol.ZH_CN]["task"]["restart"] = "继续"
	Lang[protocol.ZH_CN]["task"]["start"] = "开始"
	Lang[protocol.ZH_CN]["task"]["record"] = "工时"
	Lang[protocol.ZH_CN]["task"]["logEfforts"] = "记录工时"
	Lang[protocol.ZH_CN]["task"]["view"] = "查看任务"
	Lang[protocol.ZH_CN]["task"]["delayed"] = "延期"
	Lang[protocol.ZH_CN]["task"]["deleted"] = "已删除"
	Lang[protocol.ZH_CN]["task"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["task"]["proofreading2"] = "已对单"
	Lang[protocol.ZH_CN]["task"]["proofreading"] = "对单完成"
	Lang[protocol.ZH_CN]["task"]["internalaudit"] = "内审"
	Lang[protocol.ZH_CN]["task"]["examine"] = `审核` //"任务审核"替换`审核`
	Lang[protocol.ZH_CN]["task"]["finishall"] = "一键完成"
	Lang[protocol.ZH_CN]["task"]["batchproofreadingc"] = "取消对单"
	Lang[protocol.ZH_CN]["task"]["batchproofreading"] = "一键对单"
	Lang[protocol.ZH_CN]["task"]["batchexaminec"] = "取消审核"
	Lang[protocol.ZH_CN]["task"]["batchexamine"] = "一键审核"
	Lang[protocol.ZH_CN]["task"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["task"]["batchCancel"] = "批量取消"
	Lang[protocol.ZH_CN]["task"]["batchClose"] = "批量关闭"
	Lang[protocol.ZH_CN]["task"]["batchChangeModule"] = "批量修改模块"
	Lang[protocol.ZH_CN]["task"]["batchEdit"] = "批量编辑"
	Lang[protocol.ZH_CN]["task"]["batchCreateChildren"] = "批量建子任务"
	Lang[protocol.ZH_CN]["task"]["batchCreate"] = "批量创建"
	Lang[protocol.ZH_CN]["task"]["create"] = "建任务"
	Lang[protocol.ZH_CN]["task"]["index"] = "任务一览"
	Lang[protocol.ZH_CN]["task"]["menu"] = Lang[protocol.ZH_CN]["project"]["menu"]
	Lang[protocol.ZH_CN]["build"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["build"]["createdBugs"] = " 本次共产生 %s 个Bug"
	Lang[protocol.ZH_CN]["build"]["resolvedBugs"] = " 本次共解决 %s 个Bug"
	Lang[protocol.ZH_CN]["build"]["finishStories"] = " 本次共完成 %s 个需求"
	Lang[protocol.ZH_CN]["build"]["noBuild"] = "暂时没有版本。"
	Lang[protocol.ZH_CN]["build"]["noProduct"] = " <span style=`color:red`>该" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "没有关联" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "，无法创建版本，请先<a href=`%s`>关联" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "</a></span>"
	Lang[protocol.ZH_CN]["build"]["generatedBugs"] = "产生的Bug"
	Lang[protocol.ZH_CN]["build"]["bugs"] = "解决的Bug"
	Lang[protocol.ZH_CN]["build"]["stories"] = "完成的需求"
	Lang[protocol.ZH_CN]["build"]["unlinkBug"] = "移除Bug"
	Lang[protocol.ZH_CN]["build"]["unlinkStory"] = "移除需求"
	Lang[protocol.ZH_CN]["build"]["packageType"] = "包类型"
	Lang[protocol.ZH_CN]["build"]["last"] = "上个版本"
	Lang[protocol.ZH_CN]["build"]["files"] = "上传发行包"
	Lang[protocol.ZH_CN]["build"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["build"]["filePath"] = "下载地址"
	Lang[protocol.ZH_CN]["build"]["scmPath"] = "源代码地址"
	Lang[protocol.ZH_CN]["build"]["builder"] = "构建者"
	Lang[protocol.ZH_CN]["build"]["date"] = "打包日期"
	Lang[protocol.ZH_CN]["build"]["name"] = "名称编号"
	Lang[protocol.ZH_CN]["build"]["project"] = `所属` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["build"]["branch"] = "平台/分支"
	Lang[protocol.ZH_CN]["build"]["product"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["build"]["id"] = "ID"
	Lang[protocol.ZH_CN]["build"]["basicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["build"]["confirmUnlinkBug"] = "您确认移除该Bug吗？"
	Lang[protocol.ZH_CN]["build"]["confirmUnlinkStory"] = "您确认移除该需求吗？"
	Lang[protocol.ZH_CN]["build"]["confirmDelete"] = "您确认删除该版本吗？"
	Lang[protocol.ZH_CN]["build"]["batchUnlinkBug"] = "批量移除Bug"
	Lang[protocol.ZH_CN]["build"]["batchUnlinkStory"] = "批量移除需求"
	Lang[protocol.ZH_CN]["build"]["batchUnlink"] = "批量移除"
	Lang[protocol.ZH_CN]["build"]["view"] = "版本详情"
	Lang[protocol.ZH_CN]["build"]["deleted"] = "已删除"
	Lang[protocol.ZH_CN]["build"]["delete"] = "删除版本"
	Lang[protocol.ZH_CN]["build"]["linkBug"] = "关联Bug"
	Lang[protocol.ZH_CN]["build"]["linkStory"] = "关联需求"
	Lang[protocol.ZH_CN]["build"]["edit"] = "编辑版本"
	Lang[protocol.ZH_CN]["build"]["create"] = "创建版本"
	Lang[protocol.ZH_CN]["build"]["common"] = "版本"
	Lang[protocol.ZH_CN]["build"]["menu"] = Lang[protocol.ZH_CN]["project"]["menu"]
	Lang[protocol.ZH_CN]["qa"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["qa"]["index"] = "测试主页"
	Lang[protocol.ZH_CN]["qa"]["common"] = "测试视图"
	Lang[protocol.ZH_CN]["qa"]["menu"] = []protocol.HtmlMenu{
		{"bug", map[string]string{`link`: `Bug|bug|browse|productID=%s&branch=%s`}},
		{"testcase", map[string]string{`link`: `用例|testcase|browse|productID=%s&branch=%s`, `class`: `dropdown dropdown-hover`}},
		{"testtask", map[string]string{`link`: `测试单|testtask|browse|productID=%s&branch=%s`}},
		{"testsuite", map[string]string{`link`: `套件|testsuite|browse|productID=%s&branch=%s`}},
		{"report", map[string]string{`link`: `报告|testreport|browse|productID=%s&branch=%s`}},
		{"caselib", map[string]string{`link`: `用例库|caselib|browse`}},
	}
	Lang[protocol.ZH_CN]["qa"]["subMenu"] = []protocol.HtmlKeyValueInterface{
		{"testcase", []protocol.HtmlMenu{
			{"feature", map[string]string{`link`: `功能测试|testcase|browse|productID=%s&branch=%s`, `alias`: `view,create,batchcreate,edit,batchedit,showimport,groupcase,importfromlib`, `subModule`: `tree,story`}},
			{"unit", map[string]string{`link`: `单元测试|testtask|browseUnits|productID=%s&branch=%s`}},
		}},
	}
	Lang[protocol.ZH_CN]["bug"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["bug"]["tplExpect"] = "<p>[期望]</p><br/>"
	Lang[protocol.ZH_CN]["bug"]["tplResult"] = "<p>[结果]</p><br/>"
	Lang[protocol.ZH_CN]["bug"]["tplStep"] = "<p>[步骤]</p><br/>"
	Lang[protocol.ZH_CN]["bug"]["confirmDeleteTemplate"] = "您确认要删除该模板吗？"
	Lang[protocol.ZH_CN]["bug"]["applyTemplate"] = "应用模板"
	Lang[protocol.ZH_CN]["bug"]["skipClose"] = "Bug %s 不是已解决状态，不能关闭。"
	Lang[protocol.ZH_CN]["bug"]["remindTask"] = "该Bug已经转化为任务，是否更新任务(编号:%s)状态 ?"
	Lang[protocol.ZH_CN]["bug"]["setTemplateTitle"] = "请输入bug模板标题"
	Lang[protocol.ZH_CN]["bug"]["confirmDelete"] = "您确认要删除该Bug吗？"
	Lang[protocol.ZH_CN]["bug"]["confirmChangeProduct"] = "修改" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "会导致相应的" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "、需求和任务发生变化，确定吗？"
	Lang[protocol.ZH_CN]["bug"]["summary"] = "本页共 <strong>%s</strong> 个Bug，未解决 <strong>%s</strong>。"
	Lang[protocol.ZH_CN]["bug"]["buttonConfirm"] = "确认"
	Lang[protocol.ZH_CN]["bug"]["legendRelated"] = "其他信息"
	Lang[protocol.ZH_CN]["bug"]["legendMisc"] = "其他相关"
	Lang[protocol.ZH_CN]["bug"]["legendLife"] = "Bug的一生"
	Lang[protocol.ZH_CN]["bug"]["legendComment"] = "备注"
	Lang[protocol.ZH_CN]["bug"]["legendSteps"] = "重现步骤"
	Lang[protocol.ZH_CN]["bug"]["lblSystemBrowserAndHardware"] = "系统/浏览器"
	Lang[protocol.ZH_CN]["bug"]["lblTypeAndSeverity"] = "类型/严重程度"
	Lang[protocol.ZH_CN]["bug"]["legendPrjStoryTask"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `/需求/任务`
	Lang[protocol.ZH_CN]["bug"]["legendAttatch"] = "附件"
	Lang[protocol.ZH_CN]["bug"]["legendBasicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["bug"]["createBuild"] = "新建"
	Lang[protocol.ZH_CN]["bug"]["allBuilds"] = "所有"
	Lang[protocol.ZH_CN]["bug"]["allUsers"] = "加载所有用户"
	Lang[protocol.ZH_CN]["bug"]["lblResolved"] = "由谁解决"
	Lang[protocol.ZH_CN]["bug"]["lblLastEdited"] = "最后修改"
	Lang[protocol.ZH_CN]["bug"]["lblMailto"] = "抄送给"
	Lang[protocol.ZH_CN]["bug"]["lblAssignedTo"] = "当前指派"
	Lang[protocol.ZH_CN]["bug"]["noModule"] = "<div>您现在还没有模块信息</div><div>请维护测试模块</div>"
	Lang[protocol.ZH_CN]["bug"]["noBug"] = "暂时没有Bug。"
	Lang[protocol.ZH_CN]["bug"]["noAssigned"] = "未指派"
	Lang[protocol.ZH_CN]["bug"]["dittoNotice"] = "该bug与上一bug不属于同一产品！"
	Lang[protocol.ZH_CN]["bug"]["ditto"] = "同上"
	Lang[protocol.ZH_CN]["bug"]["resolvedByMeAB"] = "由我解决"
	Lang[protocol.ZH_CN]["bug"]["openedByMeAB"] = "由我创建"
	Lang[protocol.ZH_CN]["bug"]["assignToMeAB"] = "指派给我"
	Lang[protocol.ZH_CN]["bug"]["yesterdayClosed"] = "昨天关闭"
	Lang[protocol.ZH_CN]["bug"]["yesterdayConfirmed"] = "昨天确认"
	Lang[protocol.ZH_CN]["bug"]["yesterdayResolved"] = "昨天解决"
	Lang[protocol.ZH_CN]["bug"]["my"] = "我的"
	Lang[protocol.ZH_CN]["bug"]["allProduct"] = `所有` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["bug"]["needConfirm"] = "需求变动"
	Lang[protocol.ZH_CN]["bug"]["byQuery"] = "搜索"
	Lang[protocol.ZH_CN]["bug"]["allBugs"] = "所有"
	Lang[protocol.ZH_CN]["bug"]["overdueBugs"] = "过期Bug"
	Lang[protocol.ZH_CN]["bug"]["postponedBugs"] = "被延期"
	Lang[protocol.ZH_CN]["bug"]["longLifeBugs"] = "久未处理"
	Lang[protocol.ZH_CN]["bug"]["unconfirmed"] = "未确认"
	Lang[protocol.ZH_CN]["bug"]["unclosed"] = "未关闭"
	Lang[protocol.ZH_CN]["bug"]["toClosed"] = "待关闭"
	Lang[protocol.ZH_CN]["bug"]["unResolved"] = "未解决"
	Lang[protocol.ZH_CN]["bug"]["assignToNull"] = "未指派"
	Lang[protocol.ZH_CN]["bug"]["closedByMe"] = "由我关闭"
	Lang[protocol.ZH_CN]["bug"]["resolvedByMe"] = "由我解决"
	Lang[protocol.ZH_CN]["bug"]["openedByMe"] = "由我创建"
	Lang[protocol.ZH_CN]["bug"]["assignToMe"] = "指派给我"
	Lang[protocol.ZH_CN]["bug"]["search"] = "搜索"
	Lang[protocol.ZH_CN]["bug"]["copy"] = "复制Bug"
	Lang[protocol.ZH_CN]["bug"]["confirmStoryChange"] = "确认需求变动"
	Lang[protocol.ZH_CN]["bug"]["deleteTemplate"] = "删除模板"
	Lang[protocol.ZH_CN]["bug"]["setPublic"] = "设为公共模板"
	Lang[protocol.ZH_CN]["bug"]["saveTemplate"] = "保存模板"
	Lang[protocol.ZH_CN]["bug"]["deleted"] = "已删除"
	Lang[protocol.ZH_CN]["bug"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["bug"]["export"] = "导出数据"
	Lang[protocol.ZH_CN]["bug"]["reportChart"] = "报表统计"
	Lang[protocol.ZH_CN]["bug"]["batchActivate"] = "批量激活"
	Lang[protocol.ZH_CN]["bug"]["activate"] = "激活"
	Lang[protocol.ZH_CN]["bug"]["close"] = "关闭"
	Lang[protocol.ZH_CN]["bug"]["batchResolve"] = "批量解决"
	Lang[protocol.ZH_CN]["bug"]["resolve"] = "解决"
	Lang[protocol.ZH_CN]["bug"]["view"] = "Bug详情"
	Lang[protocol.ZH_CN]["bug"]["browse"] = "Bug列表"
	Lang[protocol.ZH_CN]["bug"]["batchAssignTo"] = "批量指派"
	Lang[protocol.ZH_CN]["bug"]["assignTo"] = "指派"
	Lang[protocol.ZH_CN]["bug"]["batchClose"] = "批量关闭"
	Lang[protocol.ZH_CN]["bug"]["batchChangeBranch"] = "批量修改分支"
	Lang[protocol.ZH_CN]["bug"]["batchChangeModule"] = "批量修改模块"
	Lang[protocol.ZH_CN]["bug"]["batchEdit"] = "批量编辑"
	Lang[protocol.ZH_CN]["bug"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["bug"]["batchConfirm"] = "批量确认"
	Lang[protocol.ZH_CN]["bug"]["confirmBug"] = "确认"
	Lang[protocol.ZH_CN]["bug"]["batchCreate"] = "批量提Bug"
	Lang[protocol.ZH_CN]["bug"]["create"] = "提Bug"
	Lang[protocol.ZH_CN]["bug"]["index"] = "首页"
	Lang[protocol.ZH_CN]["bug"]["colorTag"] = "颜色标签"
	Lang[protocol.ZH_CN]["bug"]["toCase"] = "生成用例"
	Lang[protocol.ZH_CN]["bug"]["fromCase"] = "来源用例"
	Lang[protocol.ZH_CN]["bug"]["lastEditedDate"] = "修改日期"
	Lang[protocol.ZH_CN]["bug"]["lastEditedDateAB"] = "修改日期"
	Lang[protocol.ZH_CN]["bug"]["lastEditedByAB"] = "修改者"
	Lang[protocol.ZH_CN]["bug"]["keywords"] = "关键词"
	Lang[protocol.ZH_CN]["bug"]["files"] = "附件"
	Lang[protocol.ZH_CN]["bug"]["case"] = "相关用例"
	Lang[protocol.ZH_CN]["bug"]["unlinkBug"] = "移除相关Bug"
	Lang[protocol.ZH_CN]["bug"]["linkBugs"] = "关联相关Bug"
	Lang[protocol.ZH_CN]["bug"]["linkBug"] = "相关Bug"
	Lang[protocol.ZH_CN]["bug"]["lastEditedBy"] = "最后修改者"
	Lang[protocol.ZH_CN]["bug"]["duplicateBug"] = "重复ID"
	Lang[protocol.ZH_CN]["bug"]["closedDate"] = "关闭日期"
	Lang[protocol.ZH_CN]["bug"]["closedBy"] = "由谁关闭"
	Lang[protocol.ZH_CN]["bug"]["plan"] = "所属计划"
	Lang[protocol.ZH_CN]["bug"]["deadline"] = "截止日期"
	Lang[protocol.ZH_CN]["bug"]["resolvedDateAB"] = "解决日期"
	Lang[protocol.ZH_CN]["bug"]["resolvedDate"] = "解决日期"
	Lang[protocol.ZH_CN]["bug"]["resolvedBuild"] = "解决版本"
	Lang[protocol.ZH_CN]["bug"]["resolutionAB"] = "方案"
	Lang[protocol.ZH_CN]["bug"]["resolution"] = "解决方案"
	Lang[protocol.ZH_CN]["bug"]["resolvedByAB"] = "解决"
	Lang[protocol.ZH_CN]["bug"]["resolvedBy"] = "解决者"
	Lang[protocol.ZH_CN]["bug"]["assignedDate"] = "指派日期"
	Lang[protocol.ZH_CN]["bug"]["assignedToAB"] = "指派给"
	Lang[protocol.ZH_CN]["bug"]["assignBug"] = "指派给"
	Lang[protocol.ZH_CN]["bug"]["assignedTo"] = "指派给"
	Lang[protocol.ZH_CN]["bug"]["openedBuild"] = "影响版本"
	Lang[protocol.ZH_CN]["bug"]["openedDateAB"] = "创建日期"
	Lang[protocol.ZH_CN]["bug"]["openedDate"] = "创建日期"
	Lang[protocol.ZH_CN]["bug"]["openedBy"] = "由谁创建"
	Lang[protocol.ZH_CN]["bug"]["mailto"] = "抄送给"
	Lang[protocol.ZH_CN]["bug"]["toStory"] = "转需求"
	Lang[protocol.ZH_CN]["bug"]["toTask"] = "转任务"
	Lang[protocol.ZH_CN]["bug"]["confirmed"] = "是否确认"
	Lang[protocol.ZH_CN]["bug"]["activatedDate"] = "激活日期"
	Lang[protocol.ZH_CN]["bug"]["activatedCountAB"] = "激活次数"
	Lang[protocol.ZH_CN]["bug"]["activatedCount"] = "激活次数"
	Lang[protocol.ZH_CN]["bug"]["statusAB"] = "状态"
	Lang[protocol.ZH_CN]["bug"]["status"] = "Bug状态"
	Lang[protocol.ZH_CN]["bug"]["steps"] = "重现步骤"
	Lang[protocol.ZH_CN]["bug"]["browser"] = "浏览器"
	Lang[protocol.ZH_CN]["bug"]["os"] = "操作系统"
	Lang[protocol.ZH_CN]["bug"]["type"] = "Bug类型"
	Lang[protocol.ZH_CN]["bug"]["pri"] = "优先级"
	Lang[protocol.ZH_CN]["bug"]["severityAB"] = "级别"
	Lang[protocol.ZH_CN]["bug"]["severity"] = "严重程度"
	Lang[protocol.ZH_CN]["bug"]["title"] = "Bug标题"
	Lang[protocol.ZH_CN]["bug"]["task"] = "相关任务"
	Lang[protocol.ZH_CN]["bug"]["story"] = "相关需求"
	Lang[protocol.ZH_CN]["bug"]["project"] = `所属` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["bug"]["moduleAB"] = "模块"
	Lang[protocol.ZH_CN]["bug"]["module"] = "所属模块"
	Lang[protocol.ZH_CN]["bug"]["productplan"] = "所属计划"
	Lang[protocol.ZH_CN]["bug"]["branch"] = "分支/平台"
	Lang[protocol.ZH_CN]["bug"]["product"] = `所属` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["bug"]["id"] = "Bug编号"
	Lang[protocol.ZH_CN]["bug"]["common"] = "Bug"
	Lang[protocol.ZH_CN]["bug"]["subMenu"] = Lang[protocol.ZH_CN]["qa"]["subMenu"]
	Lang[protocol.ZH_CN]["bug"]["menu"] = []protocol.HtmlMenu{
		{"bug", map[string]string{`link`: `Bug|bug|browse|productID=%s&branch=%s`, `alias`: `view,create,batchcreate,edit,resolve,close,activate,report,batchedit,batchactivate,confirmbug,assignto`, `subModule`: `tree`}},
		{"testcase", map[string]string{`link`: `用例|testcase|browse|productID=%s&branch=%s`, `class`: `dropdown dropdown-hover`}},
		{"testtask", map[string]string{`link`: `测试单|testtask|browse|productID=%s&branch=%s`}},
		{"testsuite", map[string]string{`link`: `套件|testsuite|browse|productID=%s&branch=%s`}},
		{"report", map[string]string{`link`: `报告|testreport|browse|productID=%s&branch=%s`}},
		{"caselib", map[string]string{`link`: `用例库|caselib|browse`}},
	}
	Lang[protocol.ZH_CN]["testcase"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["testcase"]["selectLib"] = "请选择库"
	Lang[protocol.ZH_CN]["testcase"]["searchStories"] = "键入来搜索需求"
	Lang[protocol.ZH_CN]["testcase"]["noCase"] = "暂时没有用例。"
	Lang[protocol.ZH_CN]["testcase"]["noModule"] = "<div>您现在还没有模块信息</div><div>请维护测试模块</div>"
	Lang[protocol.ZH_CN]["testcase"]["mustChooseResult"] = "必须选择评审结果"
	Lang[protocol.ZH_CN]["testcase"]["noLibrary"] = "现在还没有公共库，请先创建！"
	Lang[protocol.ZH_CN]["testcase"]["noRequire"] = "%s行的“%s”是必填字段，不能为空"
	Lang[protocol.ZH_CN]["testcase"]["noFunction"] = "不存在iconv和mb_convert_encoding转码方法，不能将数据转成想要的编码！"
	Lang[protocol.ZH_CN]["testcase"]["errorEncode"] = "无数据，请选择正确的编码重新上传！"
	Lang[protocol.ZH_CN]["testcase"]["buttonToList"] = "返回"
	Lang[protocol.ZH_CN]["testcase"]["dittoNotice"] = "该用例与上一用例不属于同一产品！"
	Lang[protocol.ZH_CN]["testcase"]["ditto"] = "同上"
	Lang[protocol.ZH_CN]["testcase"]["confirmBatchDelete"] = "您确认要批量删除这些测试用例吗？"
	Lang[protocol.ZH_CN]["testcase"]["confirmDelete"] = "您确认要删除该测试用例吗？"
	Lang[protocol.ZH_CN]["testcase"]["summary"] = "本页共 <strong>%s</strong> 个用例，已执行<strong>%s</strong>个。"
	Lang[protocol.ZH_CN]["testcase"]["legendComment"] = "备注"
	Lang[protocol.ZH_CN]["testcase"]["legendOpenAndEdit"] = "创建编辑"
	Lang[protocol.ZH_CN]["testcase"]["legendLinkBugs"] = "相关Bug"
	Lang[protocol.ZH_CN]["testcase"]["legendAttatch"] = "附件"
	Lang[protocol.ZH_CN]["testcase"]["legendBasicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["testcase"]["lblStatusValue"] = "状态可选值列表"
	Lang[protocol.ZH_CN]["testcase"]["lblStageValue"] = "阶段可选值列表"
	Lang[protocol.ZH_CN]["testcase"]["lblTypeValue"] = "类型可选值列表"
	Lang[protocol.ZH_CN]["testcase"]["lblLastEdited"] = "最后编辑"
	Lang[protocol.ZH_CN]["testcase"]["lblStory"] = "相关需求"
	Lang[protocol.ZH_CN]["testcase"]["unexecuted"] = "未执行"
	Lang[protocol.ZH_CN]["testcase"]["bySearch"] = "搜索"
	Lang[protocol.ZH_CN]["testcase"]["needConfirm"] = "需求变动"
	Lang[protocol.ZH_CN]["testcase"]["allTestcases"] = "所有用例"
	Lang[protocol.ZH_CN]["testcase"]["allCases"] = "所有"
	Lang[protocol.ZH_CN]["testcase"]["openedByMe"] = "我建的用例"
	Lang[protocol.ZH_CN]["testcase"]["assignToMe"] = "给我的用例"
	Lang[protocol.ZH_CN]["testcase"]["insertAfter"] = "之后添加"
	Lang[protocol.ZH_CN]["testcase"]["insertBefore"] = "之前添加"
	Lang[protocol.ZH_CN]["testcase"]["deleteStep"] = "删除"
	Lang[protocol.ZH_CN]["testcase"]["num"] = "用例记录数："
	Lang[protocol.ZH_CN]["testcase"]["new"] = "新增"
	Lang[protocol.ZH_CN]["testcase"]["viewAll"] = "查看所有"
	Lang[protocol.ZH_CN]["testcase"]["stepChild"] = "子步骤"
	Lang[protocol.ZH_CN]["testcase"]["step"] = "步骤"
	Lang[protocol.ZH_CN]["testcase"]["groupName"] = "分组名称"
	Lang[protocol.ZH_CN]["testcase"]["group"] = "分组"
	Lang[protocol.ZH_CN]["testcase"]["copy"] = "复制用例"
	Lang[protocol.ZH_CN]["testcase"]["confirmStoryChange"] = "确认需求变动"
	Lang[protocol.ZH_CN]["testcase"]["confirmChange"] = "确认用例变动"
	Lang[protocol.ZH_CN]["testcase"]["reportChart"] = "报表统计"
	Lang[protocol.ZH_CN]["testcase"]["export"] = "导出数据"
	Lang[protocol.ZH_CN]["testcase"]["exportTemplet"] = "导出模板"
	Lang[protocol.ZH_CN]["testcase"]["showImport"] = "显示导入内容"
	Lang[protocol.ZH_CN]["testcase"]["importFromLib"] = "从用例库中导入"
	Lang[protocol.ZH_CN]["testcase"]["importFile"] = "导入CSV"
	Lang[protocol.ZH_CN]["testcase"]["import"] = "导入"
	Lang[protocol.ZH_CN]["testcase"]["groupCase"] = "分组浏览用例"
	Lang[protocol.ZH_CN]["testcase"]["browse"] = "用例列表"
	Lang[protocol.ZH_CN]["testcase"]["batchCaseTypeChange"] = "批量修改类型"
	Lang[protocol.ZH_CN]["testcase"]["batchConfirmStoryChange"] = "批量确认变更"
	Lang[protocol.ZH_CN]["testcase"]["batchDelete"] = "批量删除 "
	Lang[protocol.ZH_CN]["testcase"]["batchChangeBranch"] = "批量修改分支"
	Lang[protocol.ZH_CN]["testcase"]["batchChangeModule"] = "批量修改模块"
	Lang[protocol.ZH_CN]["testcase"]["batchEdit"] = "批量编辑 "
	Lang[protocol.ZH_CN]["testcase"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["testcase"]["batchReview"] = "批量评审"
	Lang[protocol.ZH_CN]["testcase"]["review"] = "评审"
	Lang[protocol.ZH_CN]["testcase"]["view"] = "用例详情"
	Lang[protocol.ZH_CN]["testcase"]["delete"] = "删除" //"删除用例"替换"删除"
	Lang[protocol.ZH_CN]["testcase"]["batchCreate"] = "批量建用例"
	Lang[protocol.ZH_CN]["testcase"]["create"] = "建用例"
	Lang[protocol.ZH_CN]["testcase"]["index"] = "用例管理首页"
	Lang[protocol.ZH_CN]["testcase"]["common"] = "用例"
	Lang[protocol.ZH_CN]["testcase"]["stepVersion"] = "版本"
	Lang[protocol.ZH_CN]["testcase"]["stepExpect"] = "预期"
	Lang[protocol.ZH_CN]["testcase"]["stepDesc"] = "步骤"
	Lang[protocol.ZH_CN]["testcase"]["stepID"] = "编号"
	Lang[protocol.ZH_CN]["testcase"]["fromCase"] = "来源用例"
	Lang[protocol.ZH_CN]["testcase"]["fromModule"] = "来源模块"
	Lang[protocol.ZH_CN]["testcase"]["createBug"] = "转Bug"
	Lang[protocol.ZH_CN]["testcase"]["stepNumberAB"] = "S"
	Lang[protocol.ZH_CN]["testcase"]["stepNumber"] = "用例步骤数"
	Lang[protocol.ZH_CN]["testcase"]["resultsAB"] = "R"
	Lang[protocol.ZH_CN]["testcase"]["results"] = "执行结果数"
	Lang[protocol.ZH_CN]["testcase"]["bugsAB"] = "B"
	Lang[protocol.ZH_CN]["testcase"]["bugs"] = "产生Bug数"
	Lang[protocol.ZH_CN]["testcase"]["changed"] = "用例变更"
	Lang[protocol.ZH_CN]["testcase"]["toBug"] = "生成Bug"
	Lang[protocol.ZH_CN]["testcase"]["fromBug"] = "来源Bug"
	Lang[protocol.ZH_CN]["testcase"]["allProduct"] = "所有" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["testcase"]["expect"] = "预期"
	Lang[protocol.ZH_CN]["testcase"]["desc"] = "步骤"
	Lang[protocol.ZH_CN]["testcase"]["lastRunResult"] = "结果"
	Lang[protocol.ZH_CN]["testcase"]["colorTag"] = "颜色标签"
	Lang[protocol.ZH_CN]["testcase"]["assignedTo"] = "指派给"
	Lang[protocol.ZH_CN]["testcase"]["lastRunDate"] = "执行时间"
	Lang[protocol.ZH_CN]["testcase"]["lastRunner"] = "执行人"
	Lang[protocol.ZH_CN]["testcase"]["version"] = "用例版本"
	Lang[protocol.ZH_CN]["testcase"]["lastEditedDate"] = "修改日期"
	Lang[protocol.ZH_CN]["testcase"]["lastEditedDateAB"] = "修改日期"
	Lang[protocol.ZH_CN]["testcase"]["lastEditedByAB"] = "修改者"
	Lang[protocol.ZH_CN]["testcase"]["forceNotReview"] = "不需要评审"
	Lang[protocol.ZH_CN]["testcase"]["reviewResultAB"] = "结果"
	Lang[protocol.ZH_CN]["testcase"]["reviewedDateAB"] = "日期"
	Lang[protocol.ZH_CN]["testcase"]["reviewedByAB"] = "评审人"
	Lang[protocol.ZH_CN]["testcase"]["reviewResult"] = "评审结果"
	Lang[protocol.ZH_CN]["testcase"]["reviewedDate"] = "评审时间"
	Lang[protocol.ZH_CN]["testcase"]["reviewedBy"] = "由谁评审"
	Lang[protocol.ZH_CN]["testcase"]["stage"] = "适用阶段"
	Lang[protocol.ZH_CN]["testcase"]["unlinkCase"] = "移除相关用例"
	Lang[protocol.ZH_CN]["testcase"]["linkCases"] = "关联相关用例"
	Lang[protocol.ZH_CN]["testcase"]["linkCase"] = "相关用例"
	Lang[protocol.ZH_CN]["testcase"]["files"] = "附件"
	Lang[protocol.ZH_CN]["testcase"]["keywords"] = "关键词"
	Lang[protocol.ZH_CN]["testcase"]["real"] = "实际情况"
	Lang[protocol.ZH_CN]["testcase"]["result"] = "测试结果"
	Lang[protocol.ZH_CN]["testcase"]["lastEditedBy"] = "最后修改者"
	Lang[protocol.ZH_CN]["testcase"]["openedDate"] = "创建日期"
	Lang[protocol.ZH_CN]["testcase"]["openedBy"] = "由谁创建"
	Lang[protocol.ZH_CN]["testcase"]["steps"] = "用例步骤"
	Lang[protocol.ZH_CN]["testcase"]["status"] = "用例状态"
	Lang[protocol.ZH_CN]["testcase"]["type"] = "用例类型"
	Lang[protocol.ZH_CN]["testcase"]["pri"] = "优先级"
	Lang[protocol.ZH_CN]["testcase"]["precondition"] = "前置条件"
	Lang[protocol.ZH_CN]["testcase"]["title"] = "用例标题"
	Lang[protocol.ZH_CN]["testcase"]["story"] = "相关需求"
	Lang[protocol.ZH_CN]["testcase"]["moduleAB"] = "模块"
	Lang[protocol.ZH_CN]["testcase"]["branch"] = "分支/平台"
	Lang[protocol.ZH_CN]["testcase"]["lib"] = "所属库"
	Lang[protocol.ZH_CN]["testcase"]["module"] = "所属模块"
	Lang[protocol.ZH_CN]["testcase"]["product"] = "所属" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["testcase"]["id"] = "用例编号"
	Lang[protocol.ZH_CN]["testcase"]["subMenu"] = Lang[protocol.ZH_CN]["qa"]["subMenu"]
	Lang[protocol.ZH_CN]["testcase"]["menu"] = []protocol.HtmlMenu{
		{"bug", map[string]string{`link`: `Bug|bug|browse|productID=%s&branch=%s`}},
		{"testcase", map[string]string{`link`: `用例|testcase|browse|productID=%s&branch=%s`, `alias`: `view,create,batchcreate,edit,batchedit,showimport,groupcase,importfromlib`, `subModule`: `tree,story`, `class`: `dropdown dropdown-hover`}},
		{"testtask", map[string]string{`link`: `测试单|testtask|browse|productID=%s&branch=%s`}},
		{"testsuite", map[string]string{`link`: `套件|testsuite|browse|productID=%s&branch=%s`}},
		{"report", map[string]string{`link`: `报告|testreport|browse|productID=%s&branch=%s`}},
		{"caselib", map[string]string{`link`: `用例库|caselib|browse`}},
	}
	Lang[protocol.ZH_CN]["testtask"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["testtask"]["unexecuted"] = "未执行"
	Lang[protocol.ZH_CN]["testtask"]["lblResults"] = "执行结果"
	Lang[protocol.ZH_CN]["testtask"]["lblRunCase"] = "执行用例"
	Lang[protocol.ZH_CN]["testtask"]["lblUnlinkCase"] = "移除用例"
	Lang[protocol.ZH_CN]["testtask"]["lblCases"] = "用例列表"
	Lang[protocol.ZH_CN]["testtask"]["allCases"] = "所有用例"
	Lang[protocol.ZH_CN]["testtask"]["assignedToMe"] = "指派给我"
	Lang[protocol.ZH_CN]["testtask"]["checkLinked"] = "请检查测试单的产品是否与项目相关联"
	Lang[protocol.ZH_CN]["testtask"]["noTesttask"] = "暂时没有测试版本。"
	Lang[protocol.ZH_CN]["testtask"]["noticeNoOther"] = "该产品还没有其他测试版本"
	Lang[protocol.ZH_CN]["testtask"]["confirmUnlinkCase"] = "您确认要移除该用例吗？"
	Lang[protocol.ZH_CN]["testtask"]["confirmDelete"] = "您确认要删除该版本吗？"
	Lang[protocol.ZH_CN]["testtask"]["showFail"] = `失败<span class="text-danger">%s</span>次`
	Lang[protocol.ZH_CN]["testtask"]["showResult"] = `共执行<span class="text-info">%s</span>次`
	Lang[protocol.ZH_CN]["testtask"]["fail"] = "失败"
	Lang[protocol.ZH_CN]["testtask"]["pass"] = "通过"
	Lang[protocol.ZH_CN]["testtask"]["passAll"] = "全部通过"
	Lang[protocol.ZH_CN]["testtask"]["linkBySuite"] = "按套件关联"
	Lang[protocol.ZH_CN]["testtask"]["linkByBug"] = "按Bug关联"
	Lang[protocol.ZH_CN]["testtask"]["linkByStory"] = "按需求关联"
	Lang[protocol.ZH_CN]["testtask"]["linkByBuild"] = "复制版本"
	Lang[protocol.ZH_CN]["testtask"]["unlinkedCases"] = "未关联"
	Lang[protocol.ZH_CN]["testtask"]["legendBasicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["testtask"]["legendReport"] = "测试总结"
	Lang[protocol.ZH_CN]["testtask"]["legendDesc"] = "版本描述"
	Lang[protocol.ZH_CN]["testtask"]["to"] = "至"
	Lang[protocol.ZH_CN]["testtask"]["beginAndEnd"] = "起止时间"
	Lang[protocol.ZH_CN]["testtask"]["date"] = "测试时间"
	Lang[protocol.ZH_CN]["testtask"]["lastRunDate"] = "最后执行时间"
	Lang[protocol.ZH_CN]["testtask"]["lastRunner"] = "最后执行人"
	Lang[protocol.ZH_CN]["testtask"]["stepResults"] = "步骤结果"
	Lang[protocol.ZH_CN]["testtask"]["caseResult"] = "测试结果"
	Lang[protocol.ZH_CN]["testtask"]["version"] = "版本"
	Lang[protocol.ZH_CN]["testtask"]["case"] = "用例"
	Lang[protocol.ZH_CN]["testtask"]["files"] = "上传附件"
	Lang[protocol.ZH_CN]["testtask"]["reportField"] = "测试总结"
	Lang[protocol.ZH_CN]["testtask"]["lastRunResult"] = "结果"
	Lang[protocol.ZH_CN]["testtask"]["lastRunTime"] = "执行时间"
	Lang[protocol.ZH_CN]["testtask"]["lastRunAccount"] = "执行人"
	Lang[protocol.ZH_CN]["testtask"]["linkVersion"] = "版本"
	Lang[protocol.ZH_CN]["testtask"]["assignedTo"] = "指派给"
	Lang[protocol.ZH_CN]["testtask"]["status"] = "当前状态"
	Lang[protocol.ZH_CN]["testtask"]["mailto"] = "抄送给"
	Lang[protocol.ZH_CN]["testtask"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["testtask"]["end"] = "结束日期"
	Lang[protocol.ZH_CN]["testtask"]["begin"] = "开始日期"
	Lang[protocol.ZH_CN]["testtask"]["name"] = "名称"
	Lang[protocol.ZH_CN]["testtask"]["pri"] = "优先级"
	Lang[protocol.ZH_CN]["testtask"]["owner"] = "负责人"
	Lang[protocol.ZH_CN]["testtask"]["build"] = "版本"
	Lang[protocol.ZH_CN]["testtask"]["project"] = `所属` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["testtask"]["product"] = `所属` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["testtask"]["common"] = "测试单"
	Lang[protocol.ZH_CN]["testtask"]["id"] = "编号"
	Lang[protocol.ZH_CN]["testtask"]["expandAll"] = "全部展开"
	Lang[protocol.ZH_CN]["testtask"]["collapseAll"] = "全部折叠"
	Lang[protocol.ZH_CN]["testtask"]["allTasks"] = "所有测试"
	Lang[protocol.ZH_CN]["testtask"]["all"] = "全部" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["testtask"]["totalStatus"] = "全部"
	Lang[protocol.ZH_CN]["testtask"]["done"] = "已测版本"
	Lang[protocol.ZH_CN]["testtask"]["blocked"] = "被阻塞版本"
	Lang[protocol.ZH_CN]["testtask"]["testing"] = "测试中版本"
	Lang[protocol.ZH_CN]["testtask"]["activate"] = "激活"
	Lang[protocol.ZH_CN]["testtask"]["block"] = "阻塞"
	Lang[protocol.ZH_CN]["testtask"]["wait"] = "待测版本"
	Lang[protocol.ZH_CN]["testtask"]["close"] = "关闭"
	Lang[protocol.ZH_CN]["testtask"]["start"] = "开始"
	Lang[protocol.ZH_CN]["testtask"]["next"] = "下一个"
	Lang[protocol.ZH_CN]["testtask"]["pre"] = "上一个"
	Lang[protocol.ZH_CN]["testtask"]["groupCase"] = "分组浏览用例"
	Lang[protocol.ZH_CN]["testtask"]["cases"] = "用例"
	Lang[protocol.ZH_CN]["testtask"]["assign"] = "指派"
	Lang[protocol.ZH_CN]["testtask"]["createBug"] = "提Bug"
	Lang[protocol.ZH_CN]["testtask"]["results"] = "结果"
	Lang[protocol.ZH_CN]["testtask"]["batchRun"] = "批量执行"
	Lang[protocol.ZH_CN]["testtask"]["runCase"] = "执行"
	Lang[protocol.ZH_CN]["testtask"]["batchAssign"] = "批量指派"
	Lang[protocol.ZH_CN]["testtask"]["batchUnlinkCases"] = "批量移除用例"
	Lang[protocol.ZH_CN]["testtask"]["unlinkCase"] = "移除"
	Lang[protocol.ZH_CN]["testtask"]["selectVersion"] = "选择版本"
	Lang[protocol.ZH_CN]["testtask"]["linkCase"] = "关联用例"
	Lang[protocol.ZH_CN]["testtask"]["browse"] = "版本列表"
	Lang[protocol.ZH_CN]["testtask"]["edit"] = "编辑测试单"
	Lang[protocol.ZH_CN]["testtask"]["view"] = "概况"
	Lang[protocol.ZH_CN]["testtask"]["delete"] = "删除版本"
	Lang[protocol.ZH_CN]["testtask"]["reportChart"] = "报表统计"
	Lang[protocol.ZH_CN]["testtask"]["create"] = "提交测试"
	Lang[protocol.ZH_CN]["testtask"]["index"] = "版本首页"
	Lang[protocol.ZH_CN]["testtask"]["subMenu"] = Lang[protocol.ZH_CN]["qa"]["subMenu"]
	Lang[protocol.ZH_CN]["testtask"]["menu"] = []protocol.HtmlMenu{
		{"bug", map[string]string{`link`: `Bug|bug|browse|productID=%s&branch=%s`}},
		{"testcase", map[string]string{`link`: `用例|testcase|browse|productID=%s&branch=%s`, `class`: `dropdown dropdown-hover`}},
		{"testtask", map[string]string{`link`: `测试单|testtask|browse|productID=%s&branch=%s`, `subModule`: `testtask`, `alias`: `view,create,edit,linkcase,cases,start,close,batchrun,groupcase,report`}},
		{"testsuite", map[string]string{`link`: `套件|testsuite|browse|productID=%s&branch=%s`}},
		{"report", map[string]string{`link`: `报告|testreport|browse|productID=%s&branch=%s`}},
		{"caselib", map[string]string{`link`: `用例库|caselib|browse`}},
	}
	Lang[protocol.ZH_CN]["testsuite"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["testsuite"]["libView"] = "查看库概况"
	Lang[protocol.ZH_CN]["testsuite"]["createCase"] = "创建用例"
	Lang[protocol.ZH_CN]["testsuite"]["library"] = "浏览库用例"
	Lang[protocol.ZH_CN]["testsuite"]["editLib"] = "编辑库"
	Lang[protocol.ZH_CN]["testsuite"]["createLib"] = "创建库"
	Lang[protocol.ZH_CN]["testsuite"]["lblUnlinkCase"] = "移除用例"
	Lang[protocol.ZH_CN]["testsuite"]["lblCases"] = "用例列表"
	Lang[protocol.ZH_CN]["testsuite"]["noTestsuite"] = "暂时没有套件。"
	Lang[protocol.ZH_CN]["testsuite"]["noModule"] = "<div>您现在还没有模块信息</div><div>请维护用例库模块</div>"
	Lang[protocol.ZH_CN]["testsuite"]["noticeNone"] = "您还没有创建套件"
	Lang[protocol.ZH_CN]["testsuite"]["confirmUnlinkCase"] = "您确认要移除该用例吗？"
	Lang[protocol.ZH_CN]["testsuite"]["libraryDelete"] = "您确认要删除该用例库吗？"
	Lang[protocol.ZH_CN]["testsuite"]["confirmDelete"] = "您确认要删除该套件吗？"
	Lang[protocol.ZH_CN]["testsuite"]["unlinkedCases"] = "未关联"
	Lang[protocol.ZH_CN]["testsuite"]["legendBasicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["testsuite"]["legendDesc"] = "描述"
	Lang[protocol.ZH_CN]["testsuite"]["addedDate"] = "创建时间"
	Lang[protocol.ZH_CN]["testsuite"]["addedBy"] = "由谁创建"
	Lang[protocol.ZH_CN]["testsuite"]["author"] = "访问权限"
	Lang[protocol.ZH_CN]["testsuite"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["testsuite"]["name"] = "名称"
	Lang[protocol.ZH_CN]["testsuite"]["product"] = `所属` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["testsuite"]["common"] = "套件"
	Lang[protocol.ZH_CN]["testsuite"]["successSaved"] = "保存成功"
	Lang[protocol.ZH_CN]["testsuite"]["showImport"] = "显示导入数据"
	Lang[protocol.ZH_CN]["testsuite"]["import"] = "导入"
	Lang[protocol.ZH_CN]["testsuite"]["batchCreateCase"] = "批量创建用例"
	Lang[protocol.ZH_CN]["testsuite"]["exportTemplet"] = "导出模板"
	Lang[protocol.ZH_CN]["testsuite"]["deleted"] = "已删除"
	Lang[protocol.ZH_CN]["testsuite"]["batchUnlinkCases"] = "批量移除用例"
	Lang[protocol.ZH_CN]["testsuite"]["unlinkCase"] = "移除"
	Lang[protocol.ZH_CN]["testsuite"]["linkVersion"] = "版本"
	Lang[protocol.ZH_CN]["testsuite"]["linkCase"] = "关联用例"
	Lang[protocol.ZH_CN]["testsuite"]["browse"] = "套件列表"
	Lang[protocol.ZH_CN]["testsuite"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["testsuite"]["view"] = "概况"
	Lang[protocol.ZH_CN]["testsuite"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["testsuite"]["create"] = "建套件"
	Lang[protocol.ZH_CN]["testsuite"]["index"] = "套件首页"
	Lang[protocol.ZH_CN]["testsuite"]["subMenu"] = Lang[protocol.ZH_CN]["qa"]["subMenu"]
	Lang[protocol.ZH_CN]["testsuite"]["menu"] = []protocol.HtmlMenu{
		{"bug", map[string]string{`link`: `Bug|bug|browse|productID=%s&branch=%s`}},
		{"testcase", map[string]string{`link`: `用例|testcase|browse|productID=%s&branch=%s`, `class`: `dropdown dropdown-hover`}},
		{"testtask", map[string]string{`link`: `测试单|testtask|browse|productID=%s&branch=%s`}},
		{"testsuite", map[string]string{`link`: `套件|testsuite|browse|productID=%s&branch=%s`, `alias`: `view,create,edit,linkcase`}},
		{"report", map[string]string{`link`: `报告|testreport|browse|productID=%s&branch=%s`}},
		{"caselib", map[string]string{`link`: `用例库|caselib|browse`}},
	}
	Lang[protocol.ZH_CN]["testreport"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["testreport"]["bugSummary"] = "共发现<strong>%s</strong>个Bug <a data-toggle='tooltip' class='text-warning' title='{$lang->testreport->foundBugTip}'><i class='icon-help'></i></a>，遗留<strong>%s</strong>个Bug <a data-toggle='tooltip' class='text-warning' title='{$lang->testreport->legacyBugTip}'><i class='icon-help'></i></a>。用例执行产生<strong>%s</strong>个Bug <a data-toggle='tooltip' class='text-warning' title='{$lang->testreport->fromCaseBugTip}'><i class='icon-help'></i></a>。有效Bug率（方案为已解决或延期 / 状态为已解决或已关闭）：<strong>%s</strong>，用例发现Bug率（用例创建的Bug / 发现Bug数）：<strong>%s</strong>"
	Lang[protocol.ZH_CN]["testreport"]["moreProduct"] = "只能对同一个产品生成测试报告。"
	Lang[protocol.ZH_CN]["testreport"]["errorTrunk"] = "主干版本不能创建测试报告，请修改关联版本！"
	Lang[protocol.ZH_CN]["testreport"]["fromCaseBugTip"] = "测试时间范围内，用例执行失败后创建的Bug。"
	Lang[protocol.ZH_CN]["testreport"]["legacyBugTip"] = "Bug状态是激活，或Bug的解决时间在测试结束时间之后。"
	Lang[protocol.ZH_CN]["testreport"]["foundBugTip"] = "影响版本在测试轮次内，并且创建时间在测试时间范围内产生的Bug数。"
	Lang[protocol.ZH_CN]["testreport"]["noReport"] = "报表还没有生成，请稍候查看。"
	Lang[protocol.ZH_CN]["testreport"]["confirmDelete"] = "是否删除该报告？"
	Lang[protocol.ZH_CN]["testreport"]["buildSummary"] = "共测试了<strong>%s</strong>个版本。"
	Lang[protocol.ZH_CN]["testreport"]["caseSummary"] = "共有<strong>%s</strong>个用例，共执行<strong>%s</strong>个用例，产生了<strong>%s</strong>个结果，失败的用例有<strong>%s</strong>个。"
	Lang[protocol.ZH_CN]["testreport"]["bugCreateByCaseRate"] = "用例发现Bug率 (用例创建的Bug / 时间区间中新增的Bug)"
	Lang[protocol.ZH_CN]["testreport"]["bugConfirmedRate"] = "有效Bug率 (方案为已解决或延期 / 状态为已解决或已关闭)"
	Lang[protocol.ZH_CN]["testreport"]["bugModuleGroups"] = "Bug模块分布"
	Lang[protocol.ZH_CN]["testreport"]["bugResolutionGroups"] = "Bug解决方案分布"
	Lang[protocol.ZH_CN]["testreport"]["bugResolvedByGroups"] = "Bug解决者分布"
	Lang[protocol.ZH_CN]["testreport"]["bugOpenedByGroups"] = "Bug创建者分布"
	Lang[protocol.ZH_CN]["testreport"]["bugStatusGroups"] = "Bug状态分布"
	Lang[protocol.ZH_CN]["testreport"]["bugTypeGroups"] = "Bug类型别分布"
	Lang[protocol.ZH_CN]["testreport"]["bugSeverityGroups"] = "Bug严重级别分布"
	Lang[protocol.ZH_CN]["testreport"]["legendMore"] = "更多功能"
	Lang[protocol.ZH_CN]["testreport"]["legendComment"] = "总结"
	Lang[protocol.ZH_CN]["testreport"]["legendReport"] = "报表"
	Lang[protocol.ZH_CN]["testreport"]["legendLegacyBugs"] = "遗留的Bug"
	Lang[protocol.ZH_CN]["testreport"]["legendCase"] = "关联的用例"
	Lang[protocol.ZH_CN]["testreport"]["legendBuild"] = "测试轮次"
	Lang[protocol.ZH_CN]["testreport"]["legendStoryAndBug"] = "测试范围"
	Lang[protocol.ZH_CN]["testreport"]["legendBasic"] = "基本信息"
	Lang[protocol.ZH_CN]["testreport"]["deleted"] = "已删除"
	Lang[protocol.ZH_CN]["testreport"]["all"] = "所有报告"
	Lang[protocol.ZH_CN]["testreport"]["none"] = "无"
	Lang[protocol.ZH_CN]["testreport"]["value"] = "值"
	Lang[protocol.ZH_CN]["testreport"]["profile"] = "概况"
	Lang[protocol.ZH_CN]["testreport"]["objectID"] = "所属对象"
	Lang[protocol.ZH_CN]["testreport"]["createdDate"] = "创建时间"
	Lang[protocol.ZH_CN]["testreport"]["legacyBugs"] = `遗留的Bug` //"遗留的Bug"替换`遗留的Bug`
	Lang[protocol.ZH_CN]["testreport"]["report"] = "总结"
	Lang[protocol.ZH_CN]["testreport"]["bugInfo"] = "Bug分布"
	Lang[protocol.ZH_CN]["testreport"]["cases"] = "用例"
	Lang[protocol.ZH_CN]["testreport"]["goal"] = "项目目标"
	Lang[protocol.ZH_CN]["testreport"]["builds"] = "版本信息"
	Lang[protocol.ZH_CN]["testreport"]["bugs"] = "测试的Bug"
	Lang[protocol.ZH_CN]["testreport"]["stories"] = "测试的需求"
	Lang[protocol.ZH_CN]["testreport"]["end"] = "结束时间"
	Lang[protocol.ZH_CN]["testreport"]["begin"] = "开始时间"
	Lang[protocol.ZH_CN]["testreport"]["members"] = "参与人员"
	Lang[protocol.ZH_CN]["testreport"]["owner"] = "负责人"
	Lang[protocol.ZH_CN]["testreport"]["startEnd"] = "起止时间"
	Lang[protocol.ZH_CN]["testreport"]["testtask"] = "测试版本"
	Lang[protocol.ZH_CN]["testreport"]["tasks"] = Lang[protocol.ZH_CN]["testreport"]["testtask"]
	Lang[protocol.ZH_CN]["testreport"]["project"] = "所属项目"
	Lang[protocol.ZH_CN]["testreport"]["storyTitle"] = "需求标题"
	Lang[protocol.ZH_CN]["testreport"]["bugTitle"] = "Bug 标题"
	Lang[protocol.ZH_CN]["testreport"]["title"] = "标题"
	Lang[protocol.ZH_CN]["testreport"]["recreate"] = "重新生成报告"
	Lang[protocol.ZH_CN]["testreport"]["view"] = "报告详情"
	Lang[protocol.ZH_CN]["testreport"]["export"] = "导出报告"
	Lang[protocol.ZH_CN]["testreport"]["delete"] = "删除报告"
	Lang[protocol.ZH_CN]["testreport"]["edit"] = "编辑报告"
	Lang[protocol.ZH_CN]["testreport"]["create"] = "创建报告"
	Lang[protocol.ZH_CN]["testreport"]["browse"] = "报告列表"
	Lang[protocol.ZH_CN]["testreport"]["common"] = "测试报告"
	Lang[protocol.ZH_CN]["testreport"]["subMenu"] = Lang[protocol.ZH_CN]["qa"]["subMenu"]
	Lang[protocol.ZH_CN]["testreport"]["menu"] = []protocol.HtmlMenu{
		{"bug", map[string]string{`link`: `Bug|bug|browse|productID=%s&branch=%s`}},
		{"testcase", map[string]string{`link`: `用例|testcase|browse|productID=%s&branch=%s`, `class`: `dropdown dropdown-hover`}},
		{"testtask", map[string]string{`link`: `测试单|testtask|browse|productID=%s&branch=%s`}},
		{"testsuite", map[string]string{`link`: `套件|testsuite|browse|productID=%s&branch=%s`}},
		{"report", map[string]string{`link`: `报告|testreport|browse|productID=%s&branch=%s`, `alias`: `view,create,edit`}},
		{"caselib", map[string]string{`link`: `用例库|caselib|browse`}},
	}
	Lang[protocol.ZH_CN]["caselib"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["caselib"]["noModule"] = "<div>您现在还没有模块信息</div><div>请维护用例库模块</div>"
	Lang[protocol.ZH_CN]["caselib"]["libraryDelete"] = "您确认要删除该用例库吗？"
	Lang[protocol.ZH_CN]["caselib"]["legendDesc"] = "描述"
	Lang[protocol.ZH_CN]["caselib"]["lastEditedDate"] = "最后编辑时间"
	Lang[protocol.ZH_CN]["caselib"]["lastEditedBy"] = "最后编辑人"
	Lang[protocol.ZH_CN]["caselib"]["addedDate"] = "创建时间"
	Lang[protocol.ZH_CN]["caselib"]["addedBy"] = "由谁创建"
	Lang[protocol.ZH_CN]["caselib"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["caselib"]["name"] = "名称"
	Lang[protocol.ZH_CN]["caselib"]["id"] = "编号"
	Lang[protocol.ZH_CN]["caselib"]["showImport"] = "显示导入数据"
	Lang[protocol.ZH_CN]["caselib"]["importAction"] = "导入用例"
	Lang[protocol.ZH_CN]["caselib"]["import"] = "导入"
	Lang[protocol.ZH_CN]["caselib"]["batchCreateCase"] = "批量创建用例"
	Lang[protocol.ZH_CN]["caselib"]["exportTemplet"] = "导出模板"
	Lang[protocol.ZH_CN]["caselib"]["deleted"] = "已删除"
	Lang[protocol.ZH_CN]["caselib"]["linkVersion"] = "版本"
	Lang[protocol.ZH_CN]["caselib"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["caselib"]["createCase"] = "创建用例"
	Lang[protocol.ZH_CN]["caselib"]["view"] = "查看库概况"
	Lang[protocol.ZH_CN]["caselib"]["browse"] = "浏览库用例"
	Lang[protocol.ZH_CN]["caselib"]["edit"] = "编辑用例库"
	Lang[protocol.ZH_CN]["caselib"]["create"] = "创建用例库"
	Lang[protocol.ZH_CN]["caselib"]["index"] = "用例库首页"
	Lang[protocol.ZH_CN]["caselib"]["all"] = `所有用例库`    //"所有用例库"替换`所有用例库`
	Lang[protocol.ZH_CN]["caselib"]["common"] = `公共用例库` //"用例库"替换`公共用例库`
	Lang[protocol.ZH_CN]["caselib"]["subMenu"] = Lang[protocol.ZH_CN]["qa"]["subMenu"]
	Lang[protocol.ZH_CN]["caselib"]["menu"] = []protocol.HtmlMenu{
		{"bug", map[string]string{`link`: `Bug|bug|browse|`}},
		{"testcase", map[string]string{`link`: `用例|testcase|browse|`, `class`: `dropdown dropdown-hover`}},
		{"testtask", map[string]string{`link`: `测试单|testtask|browse|`}},
		{"testsuite", map[string]string{`link`: `套件|testsuite|browse|`}},
		{"report", map[string]string{`link`: `报告|testreport|browse|`}},
		{"caselib", map[string]string{`link`: `用例库|caselib|browse|libID=%s`, `alias`: `create,createcase,view,edit,batchcreatecase,showimport`, `subModule`: `tree,testcase`}},
	}
	Lang[protocol.ZH_CN]["ci"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["ci"]["menu"] = []protocol.HtmlMenu{
		{"code", map[string]string{`link`: `代码|repo|browse|repoID=%s`, `alias`: `diff,view,revision,log,blame,showsynccomment`}},
		{"build", map[string]string{`link`: `构建|job|browse`, `subModule`: `compile,job`}},
		{"jenkins", map[string]string{`link`: `Jenkins|jenkins|browse`, `alias`: `create,edit`}},
		{"maintain", map[string]string{`link`: `版本库|repo|maintain`, `alias`: `create,edit`}},
		{"rules", map[string]string{`link`: `指令|repo|setrules`}},
		{"review", map[string]string{`link`: `评审|repo|review|repoID=%s`, `subModule`: `bug`}},
	}
	Lang[protocol.ZH_CN]["repo"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["repo"]["menu"] = Lang[protocol.ZH_CN]["ci"]["menu"]
	Lang[protocol.ZH_CN]["jenkins"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["jenkins"]["menu"] = Lang[protocol.ZH_CN]["ci"]["menu"]
	Lang[protocol.ZH_CN]["compile"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["compile"]["menu"] = Lang[protocol.ZH_CN]["ci"]["menu"]
	Lang[protocol.ZH_CN]["job"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["job"]["menu"] = Lang[protocol.ZH_CN]["ci"]["menu"]
	Lang[protocol.ZH_CN]["doc"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["doc"]["noCollectedDoc"] = "您还没有收藏任何文档。"
	Lang[protocol.ZH_CN]["doc"]["noOpenedDoc"] = "您还没有创建任何文档。"
	Lang[protocol.ZH_CN]["doc"]["noEditedDoc"] = "您还没有编辑任何文档。"
	Lang[protocol.ZH_CN]["doc"]["noSearchedDoc"] = "没有搜索到任何文档。"
	Lang[protocol.ZH_CN]["doc"]["noDoc"] = "暂时没有文档。"
	Lang[protocol.ZH_CN]["doc"]["versionNotFound"] = "该版本文档不存在"
	Lang[protocol.ZH_CN]["doc"]["accessDenied"] = "您没有权限访问！"
	Lang[protocol.ZH_CN]["doc"]["errorMainSysLib"] = "该系统文档库不能删除！"
	Lang[protocol.ZH_CN]["doc"]["errorEmptyProject"] = "没有" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "，无法创建文档"
	Lang[protocol.ZH_CN]["doc"]["errorEmptyProduct"] = "没有" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "，无法创建文档"
	Lang[protocol.ZH_CN]["doc"]["errorEditSystemDoc"] = "系统文档库无需修改。"
	Lang[protocol.ZH_CN]["doc"]["confirmDeleteLib"] = "您确定删除该文档库吗？"
	Lang[protocol.ZH_CN]["doc"]["confirmDelete"] = "您确定删除该文档吗？"
	Lang[protocol.ZH_CN]["doc"]["customShowLibs"] = "文档库显示设置"
	Lang[protocol.ZH_CN]["doc"]["browseType"] = "浏览方式"
	Lang[protocol.ZH_CN]["doc"]["allProject"] = `所有` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["doc"]["allProduct"] = `所有` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["doc"]["search"] = "搜索"
	Lang[protocol.ZH_CN]["doc"]["removeMenu"] = "从菜单栏移除"
	Lang[protocol.ZH_CN]["doc"]["fixedMenu"] = "固定到菜单栏"
	Lang[protocol.ZH_CN]["doc"]["deleteLib"] = "删除文档库"
	Lang[protocol.ZH_CN]["doc"]["editLib"] = "编辑文档库"
	Lang[protocol.ZH_CN]["doc"]["showFiles"] = "附件库"
	Lang[protocol.ZH_CN]["doc"]["objectLibs"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "/" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "库列表"
	Lang[protocol.ZH_CN]["doc"]["allLibs"] = "文档库列表"
	Lang[protocol.ZH_CN]["doc"]["createLib"] = "创建文档库"
	Lang[protocol.ZH_CN]["doc"]["customAB"] = "自定义库"
	Lang[protocol.ZH_CN]["doc"]["custom"] = "自定义文档库"
	Lang[protocol.ZH_CN]["doc"]["libType"] = "文档库类型"
	Lang[protocol.ZH_CN]["doc"]["libName"] = "文档库名称"
	Lang[protocol.ZH_CN]["doc"]["cancelCollection"] = "取消收藏"
	Lang[protocol.ZH_CN]["doc"]["collect"] = "收藏"
	Lang[protocol.ZH_CN]["doc"]["childType"] = "子分类"
	Lang[protocol.ZH_CN]["doc"]["addType"] = "增加分类"
	Lang[protocol.ZH_CN]["doc"]["deleteType"] = "删除分类"
	Lang[protocol.ZH_CN]["doc"]["editType"] = "编辑分类"
	Lang[protocol.ZH_CN]["doc"]["manageType"] = "维护分类"
	Lang[protocol.ZH_CN]["doc"]["sort"] = "排序"
	Lang[protocol.ZH_CN]["doc"]["diff"] = "对比"
	Lang[protocol.ZH_CN]["doc"]["view"] = "文档详情"
	Lang[protocol.ZH_CN]["doc"]["browse"] = "文档列表"
	Lang[protocol.ZH_CN]["doc"]["delete"] = "删除文档"
	Lang[protocol.ZH_CN]["doc"]["edit"] = "编辑文档"
	Lang[protocol.ZH_CN]["doc"]["create"] = "创建文档"
	Lang[protocol.ZH_CN]["doc"]["index"] = "文档主页"
	Lang[protocol.ZH_CN]["doc"]["myCollection"] = "我的收藏"
	Lang[protocol.ZH_CN]["doc"]["myDoc"] = "我的文档"
	Lang[protocol.ZH_CN]["doc"]["pastEdited"] = "往日更新"
	Lang[protocol.ZH_CN]["doc"]["todayEdited"] = "今日更新"
	Lang[protocol.ZH_CN]["doc"]["orderByVisit"] = "最近访问"
	Lang[protocol.ZH_CN]["doc"]["orderByEdit"] = "最近更新"
	Lang[protocol.ZH_CN]["doc"]["orderByOpen"] = "最近添加"
	Lang[protocol.ZH_CN]["doc"]["openedByMe"] = "由我创建"
	Lang[protocol.ZH_CN]["doc"]["allDoc"] = "所有文档"
	Lang[protocol.ZH_CN]["doc"]["fast"] = "快速访问"
	Lang[protocol.ZH_CN]["doc"]["searchDoc"] = "搜索"
	Lang[protocol.ZH_CN]["doc"]["moduleDoc"] = "按模块浏览"
	Lang[protocol.ZH_CN]["doc"]["retrack"] = "收起"
	Lang[protocol.ZH_CN]["doc"]["fullscreen"] = "全屏"
	Lang[protocol.ZH_CN]["doc"]["searchResult"] = "搜索结果"
	Lang[protocol.ZH_CN]["doc"]["num"] = "文档数量"
	Lang[protocol.ZH_CN]["doc"]["item"] = "项"
	Lang[protocol.ZH_CN]["doc"]["users"] = "用户"
	Lang[protocol.ZH_CN]["doc"]["groups"] = "分组"
	Lang[protocol.ZH_CN]["doc"]["acl"] = "权限"
	Lang[protocol.ZH_CN]["doc"]["download"] = "下载"
	Lang[protocol.ZH_CN]["doc"]["size"] = "大小"
	Lang[protocol.ZH_CN]["doc"]["extension"] = "类型"
	Lang[protocol.ZH_CN]["doc"]["filePath"] = "地址"
	Lang[protocol.ZH_CN]["doc"]["fileTitle"] = "附件名"
	Lang[protocol.ZH_CN]["doc"]["separator"] = "<i class='icon-angle-right'></i>"
	Lang[protocol.ZH_CN]["doc"]["contentType"] = "文档格式"
	Lang[protocol.ZH_CN]["doc"]["whiteList"] = "白名单"
	Lang[protocol.ZH_CN]["doc"]["fileObject"] = "所属对象"
	Lang[protocol.ZH_CN]["doc"]["deleted"] = "已删除"
	Lang[protocol.ZH_CN]["doc"]["basicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["doc"]["version"] = "版本号"
	Lang[protocol.ZH_CN]["doc"]["editedDate"] = "编辑时间"
	Lang[protocol.ZH_CN]["doc"]["editedBy"] = "由谁编辑"
	Lang[protocol.ZH_CN]["doc"]["addedDate"] = "添加时间"
	Lang[protocol.ZH_CN]["doc"]["addedBy"] = "由谁添加"
	Lang[protocol.ZH_CN]["doc"]["files"] = "附件"
	Lang[protocol.ZH_CN]["doc"]["url"] = "文档URL"
	Lang[protocol.ZH_CN]["doc"]["keywords"] = "关键字"
	Lang[protocol.ZH_CN]["doc"]["content"] = "文档正文"
	Lang[protocol.ZH_CN]["doc"]["type"] = "文档类型"
	Lang[protocol.ZH_CN]["doc"]["comment"] = "文档备注"
	Lang[protocol.ZH_CN]["doc"]["digest"] = "文档摘要"
	Lang[protocol.ZH_CN]["doc"]["title"] = "文档标题"
	Lang[protocol.ZH_CN]["doc"]["module"] = "所属分类"
	Lang[protocol.ZH_CN]["doc"]["lib"] = "所属文档库"
	Lang[protocol.ZH_CN]["doc"]["project"] = `所属` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["doc"]["product"] = `所属` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["doc"]["id"] = "文档编号"
	Lang[protocol.ZH_CN]["doc"]["common"] = "文档"
	Lang[protocol.ZH_CN]["report"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["report"]["plan"] = "计划"
	Lang[protocol.ZH_CN]["report"]["day"] = "天"
	Lang[protocol.ZH_CN]["report"]["delay"] = "延期"
	Lang[protocol.ZH_CN]["report"]["bugStatus"] = "Bug状态"
	Lang[protocol.ZH_CN]["report"]["bugType"] = "Bug类型"
	Lang[protocol.ZH_CN]["report"]["severity"] = "严重程度"
	Lang[protocol.ZH_CN]["report"]["buildTitle"] = "测试版本"
	Lang[protocol.ZH_CN]["report"]["module"] = "模块名称"
	Lang[protocol.ZH_CN]["report"]["taskAssignedDate"] = "任务指派时间"
	Lang[protocol.ZH_CN]["report"]["bugAssignedDate"] = "Bug指派日期"
	Lang[protocol.ZH_CN]["report"]["bugResolvedDate"] = "Bug解决日期"
	Lang[protocol.ZH_CN]["report"]["userConsumed"] = "用户总消耗"
	Lang[protocol.ZH_CN]["report"]["projectConsumed"] = "项目总消耗"
	Lang[protocol.ZH_CN]["report"]["taskConsumed"] = "任务总消耗"
	Lang[protocol.ZH_CN]["report"]["taskFinishedDate"] = "任务完成时间"
	Lang[protocol.ZH_CN]["report"]["bugAssignSummary"] = "Bug指派汇总表"
	Lang[protocol.ZH_CN]["report"]["workAssignSummary"] = "任务指派汇总表"
	Lang[protocol.ZH_CN]["report"]["roadmap"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `路线图表`
	Lang[protocol.ZH_CN]["report"]["bugSummary"] = "Bug解决汇总表"
	Lang[protocol.ZH_CN]["report"]["workSummary"] = "任务完成汇总表"
	Lang[protocol.ZH_CN]["report"]["build"] = "版本统计表"
	Lang[protocol.ZH_CN]["report"]["casesrun"] = "用例执行统计表"
	Lang[protocol.ZH_CN]["report"]["testcase"] = "用例统计表"
	Lang[protocol.ZH_CN]["report"]["export"] = "导出报表"
	Lang[protocol.ZH_CN]["report"]["exportName"] = "%s报表"
	Lang[protocol.ZH_CN]["report"]["errorNoChart"] = "还没有报表数据！"
	Lang[protocol.ZH_CN]["report"]["errorExportChart"] = "该浏览器不支持Canvas图像导出功能，请换其他浏览器。"
	Lang[protocol.ZH_CN]["report"]["reportExport"] = "统计导出"
	Lang[protocol.ZH_CN]["report"]["show"] = "显示报表"
	Lang[protocol.ZH_CN]["report"]["saveReport"] = "保存报表"
	Lang[protocol.ZH_CN]["report"]["editReportAction"] = "编辑报表"
	Lang[protocol.ZH_CN]["report"]["editReport"] = "编辑"
	Lang[protocol.ZH_CN]["report"]["deleteReport"] = "删除报表"
	Lang[protocol.ZH_CN]["report"]["browseReport"] = "浏览保存报表"
	Lang[protocol.ZH_CN]["report"]["useReportAction"] = "设计报表"
	Lang[protocol.ZH_CN]["report"]["null"] = "空"
	Lang[protocol.ZH_CN]["report"]["crystalExport"] = "水晶报表导出"
	Lang[protocol.ZH_CN]["report"]["deadline"] = "截止日期"
	Lang[protocol.ZH_CN]["report"]["testTaskName"] = "版本名称"
	Lang[protocol.ZH_CN]["report"]["todoName"] = "待办名称"
	Lang[protocol.ZH_CN]["report"]["taskName"] = "任务名称"
	Lang[protocol.ZH_CN]["report"]["bugTitle"] = "Bug标题"
	Lang[protocol.ZH_CN]["report"]["idAB"] = "ID"
	Lang[protocol.ZH_CN]["report"]["overduePlan"] = "过期计划"
	Lang[protocol.ZH_CN]["report"]["closedProduct"] = `关闭` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["report"]["conditions"] = "筛选条件："
	Lang[protocol.ZH_CN]["report"]["diffDays"] = "工作日天数"
	Lang[protocol.ZH_CN]["report"]["workday"] = "每天工时"
	Lang[protocol.ZH_CN]["report"]["unplanned"] = "未计划"
	Lang[protocol.ZH_CN]["report"]["validRateTips"] = "方案为已解决或延期/状态为已解决或已关闭"
	Lang[protocol.ZH_CN]["report"]["validRate"] = "有效率"
	Lang[protocol.ZH_CN]["report"]["manhourTotal"] = "总工时"
	Lang[protocol.ZH_CN]["report"]["taskTotal"] = "总任务数"
	Lang[protocol.ZH_CN]["report"]["to"] = "至"
	Lang[protocol.ZH_CN]["report"]["total"] = "总计"
	Lang[protocol.ZH_CN]["report"]["deviationRate"] = "偏差率"
	Lang[protocol.ZH_CN]["report"]["deviation"] = "偏差"
	Lang[protocol.ZH_CN]["report"]["remain"] = "剩余工时"
	Lang[protocol.ZH_CN]["report"]["consumed"] = "总消耗"
	Lang[protocol.ZH_CN]["report"]["estimate"] = "总预计"
	Lang[protocol.ZH_CN]["report"]["task"] = "任务数"
	Lang[protocol.ZH_CN]["report"]["bugTotal"] = "Bug"
	Lang[protocol.ZH_CN]["report"]["user"] = "姓名"
	Lang[protocol.ZH_CN]["report"]["product"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "名称"
	Lang[protocol.ZH_CN]["report"]["project"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "名称"
	Lang[protocol.ZH_CN]["report"]["id"] = "编号"
	Lang[protocol.ZH_CN]["report"]["deviationChart"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `偏差曲线`
	Lang[protocol.ZH_CN]["report"]["dept"] = "部门"
	Lang[protocol.ZH_CN]["report"]["beginAndEnd"] = "起止时间"
	Lang[protocol.ZH_CN]["report"]["bugOpenedDate"] = "Bug创建时间"
	Lang[protocol.ZH_CN]["report"]["workloadAB"] = "工作负载"
	Lang[protocol.ZH_CN]["report"]["workload"] = "员工负载表"
	Lang[protocol.ZH_CN]["report"]["bugAssign"] = "Bug指派表"
	Lang[protocol.ZH_CN]["report"]["bugCreate"] = "Bug创建表"
	Lang[protocol.ZH_CN]["report"]["productSummary"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `汇总表`
	Lang[protocol.ZH_CN]["report"]["projectDeviation"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `偏差报表`
	Lang[protocol.ZH_CN]["report"]["singleColor"] = []string{"F6BD0F"}
	Lang[protocol.ZH_CN]["report"]["colors"] = []string{"AFD8F8", "F6BD0F", "8BBA00", "FF8E46", "008E8E", "D64646", "8E468E", "588526", "B3AA00", "008ED6", "9D080D", "A186BE"}
	Lang[protocol.ZH_CN]["report"]["annual"] = "年度总结"
	Lang[protocol.ZH_CN]["report"]["query"] = "查询"
	Lang[protocol.ZH_CN]["report"]["undefined"] = "未设定"
	Lang[protocol.ZH_CN]["report"]["percent"] = "百分比"
	Lang[protocol.ZH_CN]["report"]["value"] = "值"
	Lang[protocol.ZH_CN]["report"]["item"] = "条目"
	Lang[protocol.ZH_CN]["report"]["list"] = "统计报表"
	Lang[protocol.ZH_CN]["report"]["index"] = "统计首页"
	Lang[protocol.ZH_CN]["report"]["common"] = "统计视图"
	Lang[protocol.ZH_CN]["report"]["menu"] = []protocol.HtmlMenu{
		{"annual", map[string]string{`link`: `年度总结|report|annualData`, `target`: `_blank`}},
		{"product", map[string]string{`link`: "roadmap"}},
		{"prj", map[string]string{`link`: Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `|report|projectdeviation`}},
		{"test", map[string]string{`link`: `测试|report|bugcreate`, `alias`: `bugassign,testcase,build,casesrun,storylinkedbug`}},
		{"staff", map[string]string{`link`: `组织|report|workload`, `alias`: `worksummary,bugsummary,workassignsummary,bugassignsummary`}},
		{"custom", map[string]string{`link`: `自定义|report|browsereport`, `alias`: `custom`}},
	}
	Lang[protocol.ZH_CN]["report"]["notice"] = map[string]string{
		"help": "注：统计报表的数据来源于列表页面的检索结果，生成统计报表前请先在列表页面进行检索。比如列表页面我们检索的是%tab%，那么报表就是基于之前检索的%tab%的结果集进行统计。",
	}
	Lang[protocol.ZH_CN]["company"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["company"]["allEffort"] = "查看所有人日志"
	Lang[protocol.ZH_CN]["company"]["companyEffort"] = "组织日志"
	Lang[protocol.ZH_CN]["company"]["currentDept"] = "当前部门"
	Lang[protocol.ZH_CN]["company"]["allTodo"] = "查看所有人待办"
	Lang[protocol.ZH_CN]["company"]["showAll"] = "显示部门所有成员"
	Lang[protocol.ZH_CN]["company"]["effortList"] = "日志日历列表"
	Lang[protocol.ZH_CN]["company"]["todoList"] = "组织待办列表"
	Lang[protocol.ZH_CN]["company"]["companyTodo"] = "组织待办"
	Lang[protocol.ZH_CN]["company"]["endDate"] = `结束`   //"结束"替换`结束`
	Lang[protocol.ZH_CN]["company"]["beginDate"] = `开始` //"开始"替换`开始`
	Lang[protocol.ZH_CN]["company"]["todoCalendar"] = "待办日历"
	Lang[protocol.ZH_CN]["company"]["effortCalendar"] = "日志日历"
	Lang[protocol.ZH_CN]["company"]["dept"] = `部门` //"部门"替换`部门`
	Lang[protocol.ZH_CN]["company"]["to"] = "至"
	Lang[protocol.ZH_CN]["company"]["allDept"] = "所有"
	Lang[protocol.ZH_CN]["company"]["date"] = "日期"
	Lang[protocol.ZH_CN]["company"]["selectDept"] = "请选择部门"
	Lang[protocol.ZH_CN]["company"]["todo"] = "待办"
	Lang[protocol.ZH_CN]["company"]["calendar"] = "日志"
	Lang[protocol.ZH_CN]["company"]["user"] = `用户` //"用户"替换`用户`
	Lang[protocol.ZH_CN]["company"]["project"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["company"]["product"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["company"]["admins"] = "管理员"
	Lang[protocol.ZH_CN]["company"]["guest"] = "匿名登录"
	Lang[protocol.ZH_CN]["company"]["backyard"] = "内网"
	Lang[protocol.ZH_CN]["company"]["website"] = "官网"
	Lang[protocol.ZH_CN]["company"]["zipcode"] = "邮政编码"
	Lang[protocol.ZH_CN]["company"]["address"] = "通讯地址"
	Lang[protocol.ZH_CN]["company"]["fax"] = "传真"
	Lang[protocol.ZH_CN]["company"]["phone"] = "联系电话"
	Lang[protocol.ZH_CN]["company"]["name"] = "公司名称"
	Lang[protocol.ZH_CN]["company"]["orgView"] = "组织视图"
	Lang[protocol.ZH_CN]["company"]["dynamic"] = "组织动态"
	Lang[protocol.ZH_CN]["company"]["browse"] = "用户列表"
	Lang[protocol.ZH_CN]["company"]["view"] = "公司信息"
	Lang[protocol.ZH_CN]["company"]["edit"] = "编辑公司"
	Lang[protocol.ZH_CN]["company"]["index"] = "组织视图首页"
	Lang[protocol.ZH_CN]["company"]["common"] = "组织视图"
	Lang[protocol.ZH_CN]["company"]["menu"] = []protocol.HtmlMenu{
		{"browseUser", map[string]string{`link`: `用户|company|browse`, `subModule`: `user`}},
		{"dept", map[string]string{`link`: `部门|dept|browse`, `subModule`: `dept`}},
		{"browseGroup", map[string]string{`link`: `权限|group|browse`, `subModule`: `group`}},
		{"dynamic", map[string]string{`link`: "动态|company|dynamic|"}},
		{"view", map[string]string{`link`: `公司|company|view`}},
		{"todo", map[string]string{`link`: "待办|company|todo|"}},
		{"effort", map[string]string{`link`: `日志|company|calendar|`, `alias`: `effort`}},
	}
	Lang[protocol.ZH_CN]["dept"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["dept"]["successSave"] = " 修改成功。"
	Lang[protocol.ZH_CN]["dept"]["confirmDelete"] = " 您确定删除该部门吗？"
	Lang[protocol.ZH_CN]["dept"]["dragAndSort"] = "拖动排序"
	Lang[protocol.ZH_CN]["dept"]["add"] = "添加部门"
	Lang[protocol.ZH_CN]["dept"]["updateOrder"] = "更新排序"
	Lang[protocol.ZH_CN]["dept"]["manage"] = "维护部门"
	Lang[protocol.ZH_CN]["dept"]["browse"] = "部门维护"
	Lang[protocol.ZH_CN]["dept"]["name"] = "部门名称"
	Lang[protocol.ZH_CN]["dept"]["manager"] = `部门经理` //"负责人"替换`部门经理`
	Lang[protocol.ZH_CN]["dept"]["parent"] = "上级部门"
	Lang[protocol.ZH_CN]["dept"]["delete"] = "删除部门"
	Lang[protocol.ZH_CN]["dept"]["edit"] = "编辑部门"
	Lang[protocol.ZH_CN]["dept"]["manageChild"] = "下级部门"
	Lang[protocol.ZH_CN]["dept"]["common"] = "部门结构"
	Lang[protocol.ZH_CN]["dept"]["menu"] = Lang[protocol.ZH_CN]["company"]["menu"]
	Lang[protocol.ZH_CN]["group"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["group"]["all"] = "所有权限"
	Lang[protocol.ZH_CN]["group"]["other"] = "其他模块"
	Lang[protocol.ZH_CN]["group"]["outside"] = "组外用户"
	Lang[protocol.ZH_CN]["group"]["inside"] = "组内用户"
	Lang[protocol.ZH_CN]["group"]["option"] = "选项"
	Lang[protocol.ZH_CN]["group"]["priv"] = "权限"
	Lang[protocol.ZH_CN]["group"]["method"] = "方法"
	Lang[protocol.ZH_CN]["group"]["module"] = "模块"
	Lang[protocol.ZH_CN]["group"]["users"] = "用户列表"
	Lang[protocol.ZH_CN]["group"]["acl"] = "权限"
	Lang[protocol.ZH_CN]["group"]["role"] = "角色"
	Lang[protocol.ZH_CN]["group"]["desc"] = "分组描述"
	Lang[protocol.ZH_CN]["group"]["name"] = "分组名称"
	Lang[protocol.ZH_CN]["group"]["id"] = "编号"
	Lang[protocol.ZH_CN]["group"]["noticeVisit"] = "空代表访问没有访问限制"
	Lang[protocol.ZH_CN]["group"]["projectList"] = `额外允许访问` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["group"]["productList"] = `额外允许访问` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["group"]["viewList"] = "允许访问视图"
	Lang[protocol.ZH_CN]["group"]["errorNotSaved"] = "没有保存，请确认选择了权限数据。"
	Lang[protocol.ZH_CN]["group"]["successSaved"] = "成功保存"
	Lang[protocol.ZH_CN]["group"]["confirmDelete"] = "您确定删除该用户分组吗？"
	Lang[protocol.ZH_CN]["group"]["manageMember"] = "成员维护"
	Lang[protocol.ZH_CN]["group"]["byModuleTips"] = `<span class="tips">（可以按住Shift或者Ctrl键进行多选）</span>`
	Lang[protocol.ZH_CN]["group"]["managePrivByModule"] = "按模块分配权限"
	Lang[protocol.ZH_CN]["group"]["managePrivByGroup"] = "权限维护"
	Lang[protocol.ZH_CN]["group"]["managePriv"] = "权限维护"
	Lang[protocol.ZH_CN]["group"]["manageView"] = "视图维护"
	Lang[protocol.ZH_CN]["group"]["delete"] = "删除分组"
	Lang[protocol.ZH_CN]["group"]["copy"] = "复制分组"
	Lang[protocol.ZH_CN]["group"]["edit"] = "编辑分组"
	Lang[protocol.ZH_CN]["group"]["create"] = "新增分组"
	Lang[protocol.ZH_CN]["group"]["browse"] = "浏览分组"
	Lang[protocol.ZH_CN]["group"]["common"] = "权限分组"
	Lang[protocol.ZH_CN]["group"]["error"] = map[string]string{
		"GroupNotFound": "没有找到分组信息，请返回重试",
	}
	Lang[protocol.ZH_CN]["group"]["menu"] = Lang[protocol.ZH_CN]["company"]["menu"]
	Lang[protocol.ZH_CN]["user"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["user"]["restore"] = "还原用户"
	Lang[protocol.ZH_CN]["user"]["allLDAP"] = "所有"
	Lang[protocol.ZH_CN]["user"]["link"] = "关联本地账号"
	Lang[protocol.ZH_CN]["user"]["importLDAP"] = "从LDAP导入用户"
	Lang[protocol.ZH_CN]["user"]["type"] = "用户类型"
	Lang[protocol.ZH_CN]["user"]["feedbackNotice"] = "非研发用户可以使用待办、日志、办公、文档、反馈功能，而且只能使用这些功能。"
	Lang[protocol.ZH_CN]["user"]["developerAB"] = "研发"
	Lang[protocol.ZH_CN]["user"]["feedbackAB"] = "非研发"
	Lang[protocol.ZH_CN]["user"]["feedback"] = "非研发用户"
	Lang[protocol.ZH_CN]["user"]["select"] = "请选择用户"
	Lang[protocol.ZH_CN]["user"]["effort"] = "日志"
	Lang[protocol.ZH_CN]["user"]["effortcalendar"] = "日志日历"
	Lang[protocol.ZH_CN]["user"]["todocalendar"] = "待办日历"
	Lang[protocol.ZH_CN]["user"]["calendar"] = "日历"
	Lang[protocol.ZH_CN]["user"]["noticeReset"] = "请联系管理员重置密码"
	Lang[protocol.ZH_CN]["user"]["resetSuccess"] = "重置密码成功，请用新密码登录。"
	Lang[protocol.ZH_CN]["user"]["resetFail"] = "重置密码失败，检查用户名是否存在！"

	Lang[protocol.ZH_CN]["user"]["caseByHim"] = "%s建的用例"
	Lang[protocol.ZH_CN]["user"]["case2Him"] = "给%s的用例"
	Lang[protocol.ZH_CN]["user"]["testTask2Him"] = "%s负责的版本"
	Lang[protocol.ZH_CN]["user"]["canceledBy"] = "由%s取消"
	Lang[protocol.ZH_CN]["user"]["reviewedBy"] = "由%s评审"
	Lang[protocol.ZH_CN]["user"]["closedBy"] = "由%s关闭"
	Lang[protocol.ZH_CN]["user"]["resolvedBy"] = "由%s解决"
	Lang[protocol.ZH_CN]["user"]["finishedBy"] = "由%s完成"
	Lang[protocol.ZH_CN]["user"]["assignedTo"] = "指派给%s"
	Lang[protocol.ZH_CN]["user"]["openedBy"] = "由%s创建"
	Lang[protocol.ZH_CN]["user"]["dynamic"] = "动态"

	Lang[protocol.ZH_CN]["user"]["todo"] = "待办"
	Lang[protocol.ZH_CN]["user"]["schedule"] = "日程"
	Lang[protocol.ZH_CN]["user"]["testCase"] = "测试用例"
	Lang[protocol.ZH_CN]["user"]["testTask"] = "测试单"
	Lang[protocol.ZH_CN]["user"]["test"] = "测试"
	Lang[protocol.ZH_CN]["user"]["bug"] = "Bug"
	Lang[protocol.ZH_CN]["user"]["task"] = "任务"
	Lang[protocol.ZH_CN]["user"]["project"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["user"]["profile"] = "档案"
	Lang[protocol.ZH_CN]["user"]["tplContentNotEmpty"] = "模板内容不能为空!"
	Lang[protocol.ZH_CN]["user"]["setPublicTemplate"] = "设为公共模板"
	Lang[protocol.ZH_CN]["user"]["confirmDeleteTemplate"] = "您确认要删除该模板吗？"
	Lang[protocol.ZH_CN]["user"]["applyTemplate"] = "应用模板"
	Lang[protocol.ZH_CN]["user"]["setTemplateTitle"] = "请输入模板标题"
	Lang[protocol.ZH_CN]["user"]["deleteTemplate"] = "删除模板"
	Lang[protocol.ZH_CN]["user"]["setPublic"] = "设为公共模板"
	Lang[protocol.ZH_CN]["user"]["saveTemplate"] = "保存模板"
	Lang[protocol.ZH_CN]["user"]["search"] = "搜索"
	Lang[protocol.ZH_CN]["user"]["deleted"] = "(已删除)"
	Lang[protocol.ZH_CN]["user"]["goback"] = "返回前一页"
	Lang[protocol.ZH_CN]["user"]["asGuest"] = "游客访问"
	Lang[protocol.ZH_CN]["user"]["relogin"] = "重新登录"
	Lang[protocol.ZH_CN]["user"]["confirmUnbind"] = "您确定解除该用户跟ZDOO的绑定吗？"
	Lang[protocol.ZH_CN]["user"]["confirmUnlock"] = "您确定解除该用户的锁定状态吗？"
	Lang[protocol.ZH_CN]["user"]["confirmDelete"] = "您确定删除该用户吗？"
	Lang[protocol.ZH_CN]["user"]["deny"] = "访问受限"
	Lang[protocol.ZH_CN]["user"]["editProfile"] = "修改档案"
	Lang[protocol.ZH_CN]["user"]["mobileLogin"] = "手机访问"
	Lang[protocol.ZH_CN]["user"]["login"] = "用户登录"
	Lang[protocol.ZH_CN]["user"]["unbind"] = "解除ZDOO绑定"
	Lang[protocol.ZH_CN]["user"]["delete"] = "删除用户"
	Lang[protocol.ZH_CN]["user"]["unlock"] = "解锁用户"
	Lang[protocol.ZH_CN]["user"]["batchEdit"] = "批量编辑"
	Lang[protocol.ZH_CN]["user"]["edit"] = "编辑用户"
	Lang[protocol.ZH_CN]["user"]["batchCreate"] = "批量添加用户"
	Lang[protocol.ZH_CN]["user"]["create"] = "添加用户"
	Lang[protocol.ZH_CN]["user"]["view"] = "用户详情"
	Lang[protocol.ZH_CN]["user"]["index"] = "用户视图首页"
	Lang[protocol.ZH_CN]["user"]["legendContribution"] = "个人贡献"
	Lang[protocol.ZH_CN]["user"]["legendBasic"] = "基本资料"
	Lang[protocol.ZH_CN]["user"]["attendNo"] = "考勤机号码"
	Lang[protocol.ZH_CN]["user"]["score"] = "积分"
	Lang[protocol.ZH_CN]["user"]["resetPassword"] = "忘记密码"
	Lang[protocol.ZH_CN]["user"]["verifyPassword"] = "您的密码"
	Lang[protocol.ZH_CN]["user"]["newPassword"] = "新密码"
	Lang[protocol.ZH_CN]["user"]["originalPassword"] = "原密码"
	Lang[protocol.ZH_CN]["user"]["ditto"] = "同上"
	Lang[protocol.ZH_CN]["user"]["ranzhi"] = "ZDOO账号"
	Lang[protocol.ZH_CN]["user"]["last"] = "最后登录"
	Lang[protocol.ZH_CN]["user"]["ip"] = "最后IP"
	Lang[protocol.ZH_CN]["user"]["visits"] = "访问次数"
	Lang[protocol.ZH_CN]["user"]["join"] = "入职日期"
	Lang[protocol.ZH_CN]["user"]["zipcode"] = "邮编"
	Lang[protocol.ZH_CN]["user"]["address"] = "通讯地址"
	Lang[protocol.ZH_CN]["user"]["whatsapp"] = "WhatsApp"
	Lang[protocol.ZH_CN]["user"]["slack"] = "Slack"
	Lang[protocol.ZH_CN]["user"]["dingding"] = "钉钉"
	Lang[protocol.ZH_CN]["user"]["Weixin"] = "微信"
	Lang[protocol.ZH_CN]["user"]["phone"] = "电话"
	Lang[protocol.ZH_CN]["user"]["Mobile"] = "手机"
	Lang[protocol.ZH_CN]["user"]["QQ"] = "QQ"
	Lang[protocol.ZH_CN]["user"]["skype"] = "Skype"
	Lang[protocol.ZH_CN]["user"]["contactInfo"] = "联系信息"
	Lang[protocol.ZH_CN]["user"]["verify"] = "安全验证"
	Lang[protocol.ZH_CN]["user"]["accountInfo"] = "账号信息"
	Lang[protocol.ZH_CN]["user"]["basicInfo"] = "基本信息"
	Lang[protocol.ZH_CN]["user"]["email"] = "邮箱"
	Lang[protocol.ZH_CN]["user"]["gender"] = "性别"
	Lang[protocol.ZH_CN]["user"]["birthyear"] = "出生年"
	Lang[protocol.ZH_CN]["user"]["commiter"] = "源代码账号"
	Lang[protocol.ZH_CN]["user"]["nickname"] = "昵称"
	Lang[protocol.ZH_CN]["user"]["realname"] = "真实姓名"
	Lang[protocol.ZH_CN]["user"]["group"] = "权限分组"
	Lang[protocol.ZH_CN]["user"]["role"] = "职位"
	Lang[protocol.ZH_CN]["user"]["password2"] = "请重复密码"
	Lang[protocol.ZH_CN]["user"]["password"] = "密码"
	Lang[protocol.ZH_CN]["user"]["account"] = "用户名"
	Lang[protocol.ZH_CN]["user"]["dept"] = "所属部门"
	Lang[protocol.ZH_CN]["user"]["company"] = "所属公司"
	Lang[protocol.ZH_CN]["user"]["id"] = "用户编号"
	Lang[protocol.ZH_CN]["user"]["common"] = "用户"
	Lang[protocol.ZH_CN]["user"]["menu"] = []protocol.HtmlMenu{
		{"browseUser", map[string]string{`link`: `用户|company|browse`, `subModule`: `user`}},
		{"dept", map[string]string{`link`: `部门|dept|browse`, `subModule`: `dept`}},
		{"browseGroup", map[string]string{`link`: `权限|group|browse`, `subModule`: `group`}},
		//{"dynamic", map[string]string{`link`: "动态|company|dynamic|"}},
		{"view", map[string]string{`link`: `公司|company|view`}},
		{"effort", map[string]string{`link`: `日志|company|effort`, "subModule": `effort`}},
	}
	Lang[protocol.ZH_CN]["admin"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["admin"]["days"] = "日志保存天数"
	Lang[protocol.ZH_CN]["admin"]["setting"] = "设置"
	Lang[protocol.ZH_CN]["admin"]["log"] = "日志"
	Lang[protocol.ZH_CN]["admin"]["api"] = "接口"
	Lang[protocol.ZH_CN]["admin"]["getCaptcha"] = "获取验证码"
	Lang[protocol.ZH_CN]["admin"]["captcha"] = "验证码"
	Lang[protocol.ZH_CN]["admin"]["ztCompany"] = "认证公司"
	Lang[protocol.ZH_CN]["admin"]["certifyEmail"] = "认证邮箱"
	Lang[protocol.ZH_CN]["admin"]["certifyMobile"] = "认证手机"
	Lang[protocol.ZH_CN]["admin"]["checkWeak"] = "弱口令检查"
	Lang[protocol.ZH_CN]["admin"]["safeIndex"] = "安全"
	Lang[protocol.ZH_CN]["admin"]["sso"] = "然之集成"
	Lang[protocol.ZH_CN]["admin"]["checkDB"] = "检查数据库"
	Lang[protocol.ZH_CN]["admin"]["index"] = "后台管理首页"
	Lang[protocol.ZH_CN]["admin"]["common"] = "后台管理"
	Lang[protocol.ZH_CN]["admin"]["menu"] = []protocol.HtmlMenu{
		{"index", map[string]string{`link`: `首页|admin|index`, `alias`: `register,certifytemail,certifyztmobile,ztcompany`}},
		{"message", map[string]string{`link`: `通知|message|index`, `subModule`: `message,mail,webhook,sms`}},
		{"custom", map[string]string{`link`: `自定义|custom|set`, `subModule`: `custom`}},
		{"sso", map[string]string{`link`: `集成|admin|sso`, `subModule`: ``}},
		{"extension", map[string]string{`link`: `插件|extension|browse`, `subModule`: `extension`}},
		{"dev", map[string]string{`link`: `二次开发|dev|api`, `alias`: `db`, `subModule`: `dev,entry`}},
		{"translate", map[string]string{`link`: `翻译|dev|translate`}},
		{"data", map[string]string{`link`: `数据|backup|index`, `subModule`: `backup,action`}},
		{"safe", map[string]string{`link`: `安全|admin|safe`, `alias`: `checkweak`}},
		{"system", map[string]string{`link`: `系统|cron|index`, `subModule`: `cron`}},
	}
	Lang[protocol.ZH_CN]["admin"]["subMenu"] = []protocol.HtmlKeyValueInterface{
		{"message", []protocol.HtmlMenu{
			{"mail", map[string]string{`link`: `邮件|mail|index`, `subModule`: `mail`}},
			{"webhook", map[string]string{`link`: `Webhook|webhook|browse`, `subModule`: `webhook`}},
			{"browser", map[string]string{`link`: `浏览器|message|browser`}},
			{"setting", map[string]string{`link`: `设置|message|setting`}},
			{"sms", map[string]string{`link`: "短信|sms|index"}},
		}},
		{"sso", []protocol.HtmlMenu{
			{"ranzhi", map[string]string{`link`: "ZDOO|admin|sso"}},
			{"ldap", map[string]string{`link`: `LDAP|ldap|set`, `subModule`: `ldap`}},
		}},
		{"dev", []protocol.HtmlMenu{
			{"api", map[string]string{`link`: `API|dev|api`}},
			{"db", map[string]string{`link`: `数据库|dev|db`}},
			{"editor", map[string]string{`link`: `编辑器|dev|editor`}},
			{"entry", map[string]string{`link`: `应用|entry|browse`, `subModule`: `entry`}},
		}},
		{"data", []protocol.HtmlMenu{
			{"backup", map[string]string{`link`: `备份|backup|index`, `subModule`: `backup`}},
			{"trash", map[string]string{`link`: "回收站|action|trash"}},
		}},
		{"system", []protocol.HtmlMenu{
			{"cron", map[string]string{`link`: `定时|cron|index`, `subModule`: `cron`}},
		}},
	}
	Lang[protocol.ZH_CN]["convert"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["convert"]["aimType"] = "问题类型转换"
	Lang[protocol.ZH_CN]["convert"]["setParam"] = "请设置转换参数"
	Lang[protocol.ZH_CN]["convert"]["errorCopyFailed"] = "文件 %s 拷贝失败"
	Lang[protocol.ZH_CN]["convert"]["errorReleaseExists"] = "发布 %s 已存在"
	Lang[protocol.ZH_CN]["convert"]["errorBuildExists"] = "Build %s 已存在"
	Lang[protocol.ZH_CN]["convert"]["errorGroupExists"] = "分组 %s 已存在"
	Lang[protocol.ZH_CN]["convert"]["errorUserExists"] = "用户 %s 已存在"
	Lang[protocol.ZH_CN]["convert"]["errorFileNotExits"] = "文件 %s 不存在"
	Lang[protocol.ZH_CN]["convert"]["info"] = "转换信息"
	Lang[protocol.ZH_CN]["convert"]["count"] = "转换数量"
	Lang[protocol.ZH_CN]["convert"]["item"] = "转换项"
	Lang[protocol.ZH_CN]["convert"]["execute"] = "执行转换"
	Lang[protocol.ZH_CN]["convert"]["checkPath"] = "安装路径"
	Lang[protocol.ZH_CN]["convert"]["checkTable"] = "表"
	Lang[protocol.ZH_CN]["convert"]["checkDB"] = "数据库"
	Lang[protocol.ZH_CN]["convert"]["installPath"] = "%s安装的根目录"
	Lang[protocol.ZH_CN]["convert"]["dbPrefix"] = "%s表前缀"
	Lang[protocol.ZH_CN]["convert"]["dbCharset"] = "%s数据库编码"
	Lang[protocol.ZH_CN]["convert"]["dbName"] = "%s使用的库"
	Lang[protocol.ZH_CN]["convert"]["dbPassword"] = "数据库密码"
	Lang[protocol.ZH_CN]["convert"]["dbUser"] = "数据库用户名"
	Lang[protocol.ZH_CN]["convert"]["dbPort"] = "服务器端口"
	Lang[protocol.ZH_CN]["convert"]["dbHost"] = "数据库服务器"
	Lang[protocol.ZH_CN]["convert"]["fail"] = `<span class="text-danger"><i class="icon-remove-sign"></i> 检查失败</span>`
	Lang[protocol.ZH_CN]["convert"]["ok"] = `<span class="text-success"><i class="icon-check-sign"></i> 检查通过</span>`
	Lang[protocol.ZH_CN]["convert"]["checkConfig"] = "检查配置"
	Lang[protocol.ZH_CN]["convert"]["setting"] = "设置"

	Lang[protocol.ZH_CN]["convert"]["questionTypeOfRedmine"] = "Redmine中问题类型"
	Lang[protocol.ZH_CN]["convert"]["direction"] = "请选择" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "问题转换方向"
	Lang[protocol.ZH_CN]["convert"]["mustSelectSource"] = "必须选择一个来源。"
	Lang[protocol.ZH_CN]["convert"]["selectSource"] = "选择来源系统及版本"
	Lang[protocol.ZH_CN]["convert"]["convertBugFree"] = "转换BugFree"
	Lang[protocol.ZH_CN]["convert"]["convertRedmine"] = "转换Redmine"
	Lang[protocol.ZH_CN]["convert"]["checkRedmine"] = "检查Redmine"
	Lang[protocol.ZH_CN]["convert"]["checkBugFree"] = "检查Bugfree"
	Lang[protocol.ZH_CN]["convert"]["setRedmine"] = "Redmine配置"
	Lang[protocol.ZH_CN]["convert"]["setBugfree"] = "Bugfree配置"
	Lang[protocol.ZH_CN]["convert"]["setConfig"] = "来源系统配置"
	Lang[protocol.ZH_CN]["convert"]["start"] = "开始转换"
	Lang[protocol.ZH_CN]["convert"]["index"] = "首页"
	Lang[protocol.ZH_CN]["convert"]["common"] = "从其他系统导入"
	Lang[protocol.ZH_CN]["convert"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]
	Lang[protocol.ZH_CN]["upgrade"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["upgrade"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]
	Lang[protocol.ZH_CN]["action"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["action"]["undelete"] = "还原"
	Lang[protocol.ZH_CN]["action"]["trashTips"] = "提示：为了保证系统的完整性，项目管理系统的删除都是标记删除。"
	Lang[protocol.ZH_CN]["action"]["trash"] = "回收站"
	Lang[protocol.ZH_CN]["action"]["textDiff"] = "文本格式"
	Lang[protocol.ZH_CN]["action"]["product"] = Lang[protocol.ZH_CN]["common"]["productCommon"]
	Lang[protocol.ZH_CN]["action"]["project"] = Lang[protocol.ZH_CN]["common"]["projectCommon"]
	Lang[protocol.ZH_CN]["action"]["original"] = "原始格式"
	Lang[protocol.ZH_CN]["action"]["objectType"] = "对象类型"
	Lang[protocol.ZH_CN]["action"]["objectName"] = "对象名称"
	Lang[protocol.ZH_CN]["action"]["objectID"] = "对象ID"
	Lang[protocol.ZH_CN]["action"]["noDynamic"] = "暂时没有动态。"
	Lang[protocol.ZH_CN]["action"]["needEdit"] = "要还原%s的名称或代号已经存在，请编辑更改。"
	Lang[protocol.ZH_CN]["action"]["historyEdit"] = "历史记录编辑不能为空。"
	Lang[protocol.ZH_CN]["action"]["hideOne"] = "隐藏"
	Lang[protocol.ZH_CN]["action"]["hideAll"] = "全部隐藏"
	Lang[protocol.ZH_CN]["action"]["editComment"] = "修改备注"
	Lang[protocol.ZH_CN]["action"]["date"] = "日期"
	Lang[protocol.ZH_CN]["action"]["create"] = "添加备注"
	Lang[protocol.ZH_CN]["action"]["confirmHideAll"] = "您确定要全部隐藏这些记录吗？"
	Lang[protocol.ZH_CN]["action"]["common"] = "系统日志"
	Lang[protocol.ZH_CN]["action"]["comment"] = "备注"
	Lang[protocol.ZH_CN]["action"]["actor"] = "操作者"
	Lang[protocol.ZH_CN]["action"]["actionID"] = "记录ID"
	Lang[protocol.ZH_CN]["action"]["action"] = "动作"
	Lang[protocol.ZH_CN]["action"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]

	Lang[protocol.ZH_CN]["action"]["desc"] = map[string]string{
		"activated":         `$date, 由 <strong>$actor</strong> 激活。` + "\n",
		"asked":             `$date, 由 <strong>$actor</strong> 追问。` + "\n",
		"assigned":          `$date, 由 <strong>$actor</strong> 指派给 <strong>$extra</strong>。` + "\n",
		"batchexamine":      `$date 由 <strong>$actor</strong> 审核`,
		"blocked":           `$date, 由 <strong>$actor</strong> 阻塞。` + "\n",
		"bugconfirmed":      `$date, 由 <strong>$actor</strong> 确认Bug。` + "\n",
		"canceled":          `$date, 由 <strong>$actor</strong> 取消。` + "\n",
		"caseconfirmed":     `$date, 由 <strong>$actor</strong> 确认用例变动，最新版本为<strong>#$extra</strong>。` + "\n",
		"changed":           `$date, 由 <strong>$actor</strong> 变更。` + "\n",
		"closed":            `$date, 由 <strong>$actor</strong> 关闭。` + "\n",
		"commented":         `$date, 由 <strong>$actor</strong> 添加备注。` + "\n",
		"commited":          `$date, 由 <strong>$actor</strong> 提交。` + "\n",
		"common":            `$date, <strong>$action</strong> by <strong>$actor</strong>。` + "\n",
		"confirmed":         `$date, 由 <strong>$actor</strong> 确认需求变动，最新版本为<strong>#$extra</strong>。` + "\n",
		"created":           `$date, 由 <strong>$actor</strong> 创建。` + "\n",
		"delayed":           `$date, 由 <strong>$actor</strong> 延期。` + "\n",
		"deleted":           `$date, 由 <strong>$actor</strong> 删除。` + "\n",
		"deletedfile":       `$date, 由 <strong>$actor</strong> 删除了附件：<strong><i>$extra</i></strong>。` + "\n",
		"deleteestimate":    `$date, 由 <strong>$actor</strong> 删除工时。`,
		"diff1":             `修改了 <strong><i>%s</i></strong>，旧值为 "%s"，新值为 "%s"。<br />` + "\n",
		"diff2":             `修改了 <strong><i>%s</i></strong>，区别为：` + "\n" + "<blockquote class='textdiff'>%s</blockquote>" + "\n<blockquote class='original'>%s</blockquote>",
		"diff3":             `将文件名 %s 改为 %s 。` + "\n",
		"edited":            `$date, 由 <strong>$actor</strong> 编辑。` + "\n",
		"editestimate":      `$date, 由 <strong>$actor</strong> 编辑工时。`,
		"editfile":          `$date, 由 <strong>$actor</strong> 编辑了附件：<strong><i>$extra</i></strong>。` + "\n",
		"erased":            `$date, 由 <strong>$actor</strong> 删除。` + "\n",
		"examine":           `$date, 由 <strong>$actor</strong> 审核，结果：<strong>$extra</strong>`,
		"executehooks":      `$date, 由 <strong>$actor</strong> 执行扩展动作。` + "\n",
		"extra":             `$date, <strong>$action</strong> as <strong>$extra</strong> by <strong>$actor</strong>。` + "\n",
		"finished":          `$date, 由 <strong>$actor</strong> 完成。` + "\n",
		"frombug":           `$date, 由 <strong>$actor</strong> Bug转化而来，Bug编号为 <strong>$extra</strong>。`,
		"fromfeedback":      `$date, 由 <strong>$actor</strong> 从<strong>反馈</strong>转化而来，反馈编号为 <strong>$extra</strong>。` + "\n",
		"gitcommited":       `$date, 由 <strong>$actor</strong> 提交代码，版本为<strong>#$extra</strong>。` + "\n",
		"hidden":            `$date, 由 <strong>$actor</strong> 隐藏。` + "\n",
		"internalaudit":     `$date, 由 <strong>$actor</strong> 提交内审。` + "\n",
		"linkcases":         `$date, 由 <strong>$actor</strong> 关联用例` + "\n",
		"linkcomponent":     `$date, 由 <strong>$actor</strong> 关联组件` + "\n",
		"linked":            `$date, 由 <strong>$actor</strong> 关联 <strong>$extra</strong>。` + "\n",
		"linked2bug":        `$date 由 <strong>$actor</strong> 关联到版本 <strong>$extra</strong>`,
		"linkedto":          `$date, 由 <strong>$actor</strong> 关联到 <strong>$extra</strong>。` + "\n",
		"linkhost":          `$date, 由 <strong>$actor</strong> 关联主机。` + "\n",
		"linkrelatedcase":   `$date, 由 <strong>$actor</strong> 关联相关用例 <strong>$extra</strong>。` + "\n",
		"linkservice":       `$date, 由 <strong>$actor</strong> 关联服务。` + "\n",
		"moved":             `$date, 由 <strong>$actor</strong> 移动，之前为 "$extra"。` + "\n",
		"offline":           `$date, 由 <strong>$actor</strong> 下架。` + "\n",
		"online":            `$date, 由 <strong>$actor</strong> 上架。` + "\n",
		"opened":            `$date, 由 <strong>$actor</strong> 创建。` + "\n",
		"paused":            `$date, 由 <strong>$actor</strong> 暂停。` + "\n",
		"proofreading":      `$date, 由 <strong>$actor</strong> 对单完成。`,
		"recordEstimate":    `$date, 由 <strong>$actor</strong> 记录工时，消耗 <strong>$extra</strong> 小时。`,
		"replied":           `$date, 由 <strong>$actor</strong> 回复。` + "\n",
		"repocreated":       `$date, 由 <strong>$actor</strong> 评审创建：$extra。` + "\n",
		"reported":          `$date, 由 <strong>$actor</strong> 销假。` + "\n",
		"restarted":         `$date, 由 <strong>$actor</strong> 继续。` + "\n",
		"revoked":           `$date, 由 <strong>$actor</strong> 撤销。` + "\n",
		"started":           `$date, 由 <strong>$actor</strong> 启动。` + "\n",
		"suspended":         `$date, 由 <strong>$actor</strong> 挂起。` + "\n",
		"svncommited":       `$date, 由 <strong>$actor</strong> 提交代码，版本为<strong>#$extra</strong>。` + "\n",
		"tobug":             `$date, 由 <strong>$actor</strong> 转为Bug <strong>$extra</strong>。` + "\n",
		"tostory":           `$date, 由 <strong>$actor</strong> 转为` + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + ` <strong>$extra</strong>。` + "\n",
		"totask":            `$date, 由 <strong>$actor</strong> 转为任务 <strong>$extra</strong>。` + "\n",
		"totodo":            `$date, 由 <strong>$actor</strong> 转待办 <strong>$extra</strong>。`,
		"undeleted":         `$date, 由 <strong>$actor</strong> 还原。` + "\n",
		"unlinked":          `$date, 由 <strong>$actor</strong> 移除 <strong>$extra</strong>。` + "\n",
		"unlinkedfrom":      `$date, 由 <strong>$actor</strong> 从 <strong>$extra</strong> 移除。` + "\n",
		"unlinkrelatedcase": `$date, 由 <strong>$actor</strong> 移除相关用例 <strong>$extra</strong>。` + "\n",
		"userdefined":       `$date, 由 <strong>$actor</strong> $extra。` + "\n",
		"verified":          `$date, 由 <strong>$actor</strong> 验收。` + "\n",
		"workflowAction":    `$date, 由 <strong>$actor</strong> %s。` + "\n",
		"batchproofreading": `$date, 由 <strong>$actor</strong> 对单完成` + "\n",
	}
	Lang[protocol.ZH_CN]["backup"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["backup"]["restoreTip"] = "还原功能只还原附件和数据库，如果需要还原代码，可以手动还原。"
	Lang[protocol.ZH_CN]["backup"]["holdDays"] = "备份保留最近 %s 天"
	Lang[protocol.ZH_CN]["backup"]["confirmRestore"] = "是否还原该备份？"
	Lang[protocol.ZH_CN]["backup"]["confirmDelete"] = "是否删除备份？"
	Lang[protocol.ZH_CN]["backup"]["progressCode"] = "<p>SQL备份完成</p><p>附件备份完成</p><p>代码备份中，已备份%s</p>"
	Lang[protocol.ZH_CN]["backup"]["progressAttach"] = "<p>SQL备份完成</p><p>附件备份中，已备份%s</p>"
	Lang[protocol.ZH_CN]["backup"]["progressSQL"] = "<p>SQL备份中，已备份%s</p>"
	Lang[protocol.ZH_CN]["backup"]["waitting"] = `<span id="backupType"></span>正在进行中，请稍候...`
	Lang[protocol.ZH_CN]["backup"]["settingDir"] = "备份目录"
	Lang[protocol.ZH_CN]["backup"]["setting"] = "设置"
	Lang[protocol.ZH_CN]["backup"]["size"] = "大小"
	Lang[protocol.ZH_CN]["backup"]["files"] = "备份文件"
	Lang[protocol.ZH_CN]["backup"]["time"] = "备份时间"

	Lang[protocol.ZH_CN]["backup"]["changeAB"] = "修改"
	Lang[protocol.ZH_CN]["backup"]["change"] = "修改保留时间"
	Lang[protocol.ZH_CN]["backup"]["restore"] = "还原"
	Lang[protocol.ZH_CN]["backup"]["backup"] = "备份"
	Lang[protocol.ZH_CN]["backup"]["delete"] = "删除备份"
	Lang[protocol.ZH_CN]["backup"]["history"] = "备份历史"
	Lang[protocol.ZH_CN]["backup"]["index"] = "备份首页"
	Lang[protocol.ZH_CN]["backup"]["common"] = "备份"
	Lang[protocol.ZH_CN]["backup"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]
	Lang[protocol.ZH_CN]["cron"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["cron"]["confirmOpen"] = `<p>是否开启该功能？<a href="%s" target='hiddenwin'><strong>打开计划任务</strong></a></p>`
	Lang[protocol.ZH_CN]["cron"]["introduction"] = "<p>计划任务功能可以定时执行诸如更新燃尽图、备份等操作，免除自己布置计划任务。</p><p>该功能还有待完善，所以默认关闭该功能。</p>"
	Lang[protocol.ZH_CN]["cron"]["confirmTurnon"] = "是否关闭计划任务？"
	Lang[protocol.ZH_CN]["cron"]["confirmDelete"] = "是否删除该计划任务？"
	Lang[protocol.ZH_CN]["cron"]["lastTime"] = "最后执行"
	Lang[protocol.ZH_CN]["cron"]["remark"] = "备注"
	Lang[protocol.ZH_CN]["cron"]["type"] = "任务类型"
	Lang[protocol.ZH_CN]["cron"]["status"] = "状态"
	Lang[protocol.ZH_CN]["cron"]["command"] = "命令"
	Lang[protocol.ZH_CN]["cron"]["dow"] = "周"
	Lang[protocol.ZH_CN]["cron"]["mon"] = "月"
	Lang[protocol.ZH_CN]["cron"]["dom"] = "天"
	Lang[protocol.ZH_CN]["cron"]["h"] = "小时"
	Lang[protocol.ZH_CN]["cron"]["m"] = "分"
	Lang[protocol.ZH_CN]["cron"]["openProcess"] = "重启"
	Lang[protocol.ZH_CN]["cron"]["turnon"] = "打开/关闭"
	Lang[protocol.ZH_CN]["cron"]["toggle"] = "激活/禁用"
	Lang[protocol.ZH_CN]["cron"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["cron"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["cron"]["create"] = "添加"
	Lang[protocol.ZH_CN]["cron"]["list"] = "任务列表"
	Lang[protocol.ZH_CN]["cron"]["index"] = "首页"
	Lang[protocol.ZH_CN]["cron"]["common"] = "计划任务"
	Lang[protocol.ZH_CN]["cron"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]

	Lang[protocol.ZH_CN]["custom"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["custom"]["saveFail"] = "保存失败！"
	Lang[protocol.ZH_CN]["custom"]["menuTip"] = "点击显示或隐藏导航条目，拖拽来更改显示顺序。"
	Lang[protocol.ZH_CN]["custom"]["weekend"] = "休息日"
	Lang[protocol.ZH_CN]["custom"]["workingHours"] = "每天可用工时"
	Lang[protocol.ZH_CN]["custom"]["forceNotReview"] = "不需要评审"
	Lang[protocol.ZH_CN]["custom"]["forceReview"] = "强制评审"
	Lang[protocol.ZH_CN]["custom"]["storyReview"] = "评审流程"
	Lang[protocol.ZH_CN]["custom"]["confirmRestore"] = "是否要恢复默认配置？"
	Lang[protocol.ZH_CN]["custom"]["allLang"] = "适用所有语言"
	Lang[protocol.ZH_CN]["custom"]["currentLang"] = "适用当前语言"
	Lang[protocol.ZH_CN]["custom"]["scoreReset"] = "重置积分"
	Lang[protocol.ZH_CN]["custom"]["score"] = "积分"
	Lang[protocol.ZH_CN]["custom"]["required"] = "必填项"
	Lang[protocol.ZH_CN]["custom"]["setPublic"] = "设为公共"
	Lang[protocol.ZH_CN]["custom"]["lang"] = "所属语言"
	Lang[protocol.ZH_CN]["custom"]["section"] = "附加部分"
	Lang[protocol.ZH_CN]["custom"]["module"] = "模块"
	Lang[protocol.ZH_CN]["custom"]["owner"] = "所有者"
	Lang[protocol.ZH_CN]["custom"]["branch"] = "多分支"
	Lang[protocol.ZH_CN]["custom"]["select"] = "请选择流程："
	Lang[protocol.ZH_CN]["custom"]["working"] = "工作方式"
	Lang[protocol.ZH_CN]["custom"]["flow"] = "流程"
	Lang[protocol.ZH_CN]["custom"]["value"] = "值"
	Lang[protocol.ZH_CN]["custom"]["key"] = "键"
	Lang[protocol.ZH_CN]["custom"]["restore"] = "恢复默认"
	Lang[protocol.ZH_CN]["custom"]["set"] = "自定义配置"
	Lang[protocol.ZH_CN]["custom"]["index"] = "首页"
	Lang[protocol.ZH_CN]["custom"]["common"] = "自定义"
	Lang[protocol.ZH_CN]["custom"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]
	Lang[protocol.ZH_CN]["mail"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["mail"]["addressWhiteList"] = "为防止邮件被屏蔽，请在邮件服务器里面将发信邮箱设为白名单"
	Lang[protocol.ZH_CN]["mail"]["closeSendCloud"] = "关闭SendCloud"
	Lang[protocol.ZH_CN]["mail"]["sendCloudSuccess"] = "操作成功"
	Lang[protocol.ZH_CN]["mail"]["sendCloudFail"] = "操作失败，原因："
	Lang[protocol.ZH_CN]["mail"]["centifyFail"] = "验证失败，可能密钥已经修改。请重新绑定！"
	Lang[protocol.ZH_CN]["mail"]["connectFail"] = "无法连接网站。"
	Lang[protocol.ZH_CN]["mail"]["needConfigure"] = "无法找到邮件配置信息，请先配置邮件发送参数。"
	Lang[protocol.ZH_CN]["mail"]["sendmailTips"] = "提示：系统不会为当前操作者发信。"
	Lang[protocol.ZH_CN]["mail"]["confirmDelete"] = "是否删除邮件？"
	Lang[protocol.ZH_CN]["mail"]["successSended"] = "成功发送！"
	Lang[protocol.ZH_CN]["mail"]["testContent"] = "邮箱设置成功"
	Lang[protocol.ZH_CN]["mail"]["testSubject"] = "测试邮件"
	Lang[protocol.ZH_CN]["mail"]["successSaved"] = "配置信息已经成功保存。"
	Lang[protocol.ZH_CN]["mail"]["nextStep"] = "下一步"
	Lang[protocol.ZH_CN]["mail"]["inputFromEmail"] = "请输入发信邮箱："
	Lang[protocol.ZH_CN]["mail"]["noticeResend"] = "已经重新发信！"
	Lang[protocol.ZH_CN]["mail"]["more"] = "更多..."
	Lang[protocol.ZH_CN]["mail"]["failReason"] = "失败原因"
	Lang[protocol.ZH_CN]["mail"]["status"] = "状态"
	Lang[protocol.ZH_CN]["mail"]["sendTime"] = "发送时间"
	Lang[protocol.ZH_CN]["mail"]["createdDate"] = "创建时间"
	Lang[protocol.ZH_CN]["mail"]["createdBy"] = "发送者"
	Lang[protocol.ZH_CN]["mail"]["subject"] = "主题"
	Lang[protocol.ZH_CN]["mail"]["ccList"] = "抄送给"
	Lang[protocol.ZH_CN]["mail"]["toList"] = "收信人"
	Lang[protocol.ZH_CN]["mail"]["remove"] = "移除"
	Lang[protocol.ZH_CN]["mail"]["sync"] = "同步"
	Lang[protocol.ZH_CN]["mail"]["unsyncUser"] = "未同步"
	Lang[protocol.ZH_CN]["mail"]["syncedUser"] = "已经同步"
	Lang[protocol.ZH_CN]["mail"]["smtp"] = "SMTP发信"
	Lang[protocol.ZH_CN]["mail"]["selectMTA"] = "请选择发信方式："
	Lang[protocol.ZH_CN]["mail"]["license"] = "云发信使用须知"
	Lang[protocol.ZH_CN]["mail"]["secretKey"] = "secretKey"
	Lang[protocol.ZH_CN]["mail"]["accessKey"] = "accessKey"
	Lang[protocol.ZH_CN]["mail"]["charset"] = "编码"
	Lang[protocol.ZH_CN]["mail"]["debug"] = "调试级别"
	Lang[protocol.ZH_CN]["mail"]["secure"] = "是否加密"
	Lang[protocol.ZH_CN]["mail"]["password"] = "smtp密码"
	Lang[protocol.ZH_CN]["mail"]["username"] = "smtp帐号"
	Lang[protocol.ZH_CN]["mail"]["auth"] = "是否需要验证"
	Lang[protocol.ZH_CN]["mail"]["port"] = "smtp端口号"
	Lang[protocol.ZH_CN]["mail"]["host"] = "smtp服务器"
	Lang[protocol.ZH_CN]["mail"]["domain"] = "域名"
	Lang[protocol.ZH_CN]["mail"]["fromName"] = "发信人"
	Lang[protocol.ZH_CN]["mail"]["fromAddress"] = "发信邮箱"
	Lang[protocol.ZH_CN]["mail"]["async"] = "异步发送"
	Lang[protocol.ZH_CN]["mail"]["turnon"] = "是否打开"
	Lang[protocol.ZH_CN]["mail"]["disagree"] = "不同意"
	Lang[protocol.ZH_CN]["mail"]["agreeLicense"] = "同意"
	Lang[protocol.ZH_CN]["mail"]["sendcloudUser"] = "同步联系人"
	Lang[protocol.ZH_CN]["mail"]["batchDelete"] = "批量删除"
	Lang[protocol.ZH_CN]["mail"]["sendCloud"] = "Notice发信"
	Lang[protocol.ZH_CN]["mail"]["gmail"] = "GMAIL发信"
	Lang[protocol.ZH_CN]["mail"]["ztCloud"] = "云发信"
	Lang[protocol.ZH_CN]["mail"]["delete"] = "删除邮件"
	Lang[protocol.ZH_CN]["mail"]["browse"] = "邮件列表"
	Lang[protocol.ZH_CN]["mail"]["resend"] = "重发"
	Lang[protocol.ZH_CN]["mail"]["reset"] = "重置"
	Lang[protocol.ZH_CN]["mail"]["test"] = "测试发信"
	Lang[protocol.ZH_CN]["mail"]["save"] = "保存"
	Lang[protocol.ZH_CN]["mail"]["edit"] = "编辑配置"
	Lang[protocol.ZH_CN]["mail"]["detect"] = "检测"
	Lang[protocol.ZH_CN]["mail"]["index"] = "首页"
	Lang[protocol.ZH_CN]["mail"]["common"] = "发信配置"
	Lang[protocol.ZH_CN]["mail"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]
	Lang[protocol.ZH_CN]["dev"] = make(map[string]interface{})

	Lang[protocol.ZH_CN]["dev"]["paramMailto"] = "填写帐号，多个账号用','分隔。"
	Lang[protocol.ZH_CN]["dev"]["paramColor"] = "颜色格式：#RGB，如：#3da7f5"
	Lang[protocol.ZH_CN]["dev"]["paramDate"] = "日期格式：YY-mm-dd，如：2019-01-01"
	Lang[protocol.ZH_CN]["dev"]["paramRange"] = "取值范围：%s"
	Lang[protocol.ZH_CN]["dev"]["post"] = "POST参数"
	Lang[protocol.ZH_CN]["dev"]["noParams"] = "无参数"
	Lang[protocol.ZH_CN]["dev"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["dev"]["type"] = "类型"
	Lang[protocol.ZH_CN]["dev"]["params"] = "参数列表"
	Lang[protocol.ZH_CN]["dev"]["moduleList"] = "模块列表"
	Lang[protocol.ZH_CN]["dev"]["dbList"] = "数据库"
	Lang[protocol.ZH_CN]["dev"]["translate"] = "翻译"
	Lang[protocol.ZH_CN]["dev"]["editor"] = "编辑器"
	Lang[protocol.ZH_CN]["dev"]["db"] = "数据库"
	Lang[protocol.ZH_CN]["dev"]["api"] = "API"
	Lang[protocol.ZH_CN]["dev"]["common"] = "二次开发"
	Lang[protocol.ZH_CN]["dev"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]
	Lang[protocol.ZH_CN]["entry"] = make(map[string]interface{})

	Lang[protocol.ZH_CN]["entry"]["notify"] = "消息通知"
	Lang[protocol.ZH_CN]["entry"]["help"] = "使用说明"
	Lang[protocol.ZH_CN]["entry"]["confirmDelete"] = "您确认要删除该应用吗？"
	Lang[protocol.ZH_CN]["entry"]["url"] = "请求地址"
	Lang[protocol.ZH_CN]["entry"]["date"] = "请求时间"
	Lang[protocol.ZH_CN]["entry"]["editedDate"] = "编辑时间"
	Lang[protocol.ZH_CN]["entry"]["editedby"] = "最后编辑"
	Lang[protocol.ZH_CN]["entry"]["createdDate"] = "创建时间"
	Lang[protocol.ZH_CN]["entry"]["createdBy"] = "由谁创建"
	Lang[protocol.ZH_CN]["entry"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["entry"]["ip"] = "IP"
	Lang[protocol.ZH_CN]["entry"]["key"] = "密钥"
	Lang[protocol.ZH_CN]["entry"]["freePasswd"] = "免密登录"
	Lang[protocol.ZH_CN]["entry"]["code"] = "代号"
	Lang[protocol.ZH_CN]["entry"]["account"] = "账号"
	Lang[protocol.ZH_CN]["entry"]["name"] = "名称"
	Lang[protocol.ZH_CN]["entry"]["id"] = "ID"
	Lang[protocol.ZH_CN]["entry"]["createKey"] = "重新生成密钥"
	Lang[protocol.ZH_CN]["entry"]["delete"] = "删除应用"
	Lang[protocol.ZH_CN]["entry"]["edit"] = "编辑应用"
	Lang[protocol.ZH_CN]["entry"]["create"] = "添加应用"
	Lang[protocol.ZH_CN]["entry"]["browse"] = "浏览应用"
	Lang[protocol.ZH_CN]["entry"]["setting"] = "设置"
	Lang[protocol.ZH_CN]["entry"]["log"] = "日志"
	Lang[protocol.ZH_CN]["entry"]["webhook"] = "Webhook"
	Lang[protocol.ZH_CN]["entry"]["api"] = "接口"
	Lang[protocol.ZH_CN]["entry"]["list"] = "应用列表"
	Lang[protocol.ZH_CN]["entry"]["common"] = "应用"
	Lang[protocol.ZH_CN]["entry"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]
	Lang[protocol.ZH_CN]["webhook"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["webhook"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]
	Lang[protocol.ZH_CN]["message"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["message"]["setting"] = "设置"
	Lang[protocol.ZH_CN]["message"]["index"] = "首页"
	Lang[protocol.ZH_CN]["message"]["common"] = "消息"
	Lang[protocol.ZH_CN]["message"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]
	Lang[protocol.ZH_CN]["menugroup"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["menugroup"]["sms"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["trip"] = "oa"
	Lang[protocol.ZH_CN]["menugroup"]["holiday"] = "oa"
	Lang[protocol.ZH_CN]["menugroup"]["lieu"] = "oa"
	Lang[protocol.ZH_CN]["menugroup"]["overtime"] = "oa"
	Lang[protocol.ZH_CN]["menugroup"]["makeup"] = "oa"
	Lang[protocol.ZH_CN]["menugroup"]["leave"] = "oa"
	Lang[protocol.ZH_CN]["menugroup"]["attend"] = "oa"
	Lang[protocol.ZH_CN]["menugroup"]["ldap"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["effort"] = "my"
	Lang[protocol.ZH_CN]["menugroup"]["job"] = "ci"
	Lang[protocol.ZH_CN]["menugroup"]["compile"] = "ci"
	Lang[protocol.ZH_CN]["menugroup"]["jenkins"] = "ci"
	Lang[protocol.ZH_CN]["menugroup"]["repo"] = "ci"
	Lang[protocol.ZH_CN]["menugroup"]["message"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["webhook"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["entry"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["dev"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["mail"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["custom"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["extension"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["cron"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["backup"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["action"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["score"] = "my"
	Lang[protocol.ZH_CN]["menugroup"]["todo"] = "my"
	Lang[protocol.ZH_CN]["menugroup"]["dept"] = "company"
	Lang[protocol.ZH_CN]["menugroup"]["people"] = "company"
	Lang[protocol.ZH_CN]["menugroup"]["doclib"] = "doc"
	Lang[protocol.ZH_CN]["menugroup"]["testreport"] = "qa"
	Lang[protocol.ZH_CN]["menugroup"]["caselib"] = "qa"
	Lang[protocol.ZH_CN]["menugroup"]["testsuite"] = "qa"
	Lang[protocol.ZH_CN]["menugroup"]["testtask"] = "qa"
	Lang[protocol.ZH_CN]["menugroup"]["case"] = "qa"
	Lang[protocol.ZH_CN]["menugroup"]["testcase"] = "qa"
	Lang[protocol.ZH_CN]["menugroup"]["bug"] = "qa"
	Lang[protocol.ZH_CN]["menugroup"]["group"] = "company"
	Lang[protocol.ZH_CN]["menugroup"]["user"] = "company"
	Lang[protocol.ZH_CN]["menugroup"]["upgrade"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["convert"] = "admin"
	Lang[protocol.ZH_CN]["menugroup"]["build"] = "project"
	Lang[protocol.ZH_CN]["menugroup"]["task"] = "project"
	Lang[protocol.ZH_CN]["menugroup"]["productplan"] = "product"
	Lang[protocol.ZH_CN]["menugroup"]["branch"] = "product"
	Lang[protocol.ZH_CN]["menugroup"]["story"] = "product"
	Lang[protocol.ZH_CN]["menugroup"]["release"] = "product"
	Lang[protocol.ZH_CN]["error"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["error"]["noCurlExt"] = "服务器未安装Curl模块。"
	Lang[protocol.ZH_CN]["error"]["tutorialData"] = "新手模式下不会插入数据，请退出新手模式操作"
	Lang[protocol.ZH_CN]["error"]["editedByOther"] = "该记录可能已经被改动。请刷新页面重新编辑！"
	Lang[protocol.ZH_CN]["error"]["noData"] = "没有数据"
	Lang[protocol.ZH_CN]["error"]["pasteImg"] = "您的浏览器不支持粘贴图片！"
	Lang[protocol.ZH_CN]["error"]["accessDenied"] = "您没有访问权限"
	Lang[protocol.ZH_CN]["error"]["account"] = "『%s』只能是字母和数字的组合三位以上。"
	Lang[protocol.ZH_CN]["error"]["code"] = "『%s』应当为字母或数字的组合。"
	Lang[protocol.ZH_CN]["error"]["datetime"] = "『%s』应当为合法的日期。"
	Lang[protocol.ZH_CN]["error"]["date"] = "『%s』应当为合法的日期。"
	Lang[protocol.ZH_CN]["error"]["URL"] = "『%s』应当为合法的URL。"
	Lang[protocol.ZH_CN]["error"]["email"] = "『%s』应当为合法的EMAIL。"
	Lang[protocol.ZH_CN]["error"]["float"] = "『%s』应当是数字，可以是小数。"
	Lang[protocol.ZH_CN]["error"]["int"] = []string{"『%s』应当是数字。", "『%s』应当介于『%s-%s』之间。"}
	Lang[protocol.ZH_CN]["error"]["equal"] = "『%s』必须为『%s』。"
	Lang[protocol.ZH_CN]["error"]["empty"] = "『%s』必须为空。"
	Lang[protocol.ZH_CN]["error"]["notempty"] = "『%s』不能为空。"
	Lang[protocol.ZH_CN]["error"]["ge"] = "『%s』应当不小于『%s』。"
	Lang[protocol.ZH_CN]["error"]["gt"] = "『%s』应当大于『%s』。"
	Lang[protocol.ZH_CN]["error"]["unique"] = "『%s』已经有『%s』这条记录了。如果您确定该记录已删除，请到后台-数据-回收站还原。"
	Lang[protocol.ZH_CN]["error"]["reg"] = "『%s』不符合格式，应当为:『%s』。"
	Lang[protocol.ZH_CN]["error"]["length"] = []string{"『%s』长度错误，应当为『%d』", "『%s』长度应当不超过『%d』，且大于『%d』。"}
	Lang[protocol.ZH_CN]["error"]["companyNotFound"] = "您访问的域名 %s 没有对应的公司。"
	Lang[protocol.ZH_CN]["pager"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["pager"]["pageOfTotal"] = "第 <strong>{page}</strong>/<strong>{totalPage}</strong> 页"
	Lang[protocol.ZH_CN]["pager"]["itemsRange"] = "第 <strong>{start}</strong> ~ <strong>{end}</strong> 项"
	Lang[protocol.ZH_CN]["pager"]["pageSize"] = "每页 <strong>{recPerPage}</strong> 项"
	Lang[protocol.ZH_CN]["pager"]["totalCount"] = "共 <strong>{recTotal}</strong> 项"
	Lang[protocol.ZH_CN]["pager"]["totalPage"] = "共 <strong>{totalPage}</strong> 页"
	Lang[protocol.ZH_CN]["pager"]["pageOf"] = "第 <strong>{page}</strong> 页"
	Lang[protocol.ZH_CN]["pager"]["goto"] = "跳转"
	Lang[protocol.ZH_CN]["pager"]["lastPage"] = "最后一页"
	Lang[protocol.ZH_CN]["pager"]["firstPage"] = "第一页"
	Lang[protocol.ZH_CN]["pager"]["pageOfText"] = "第 {0} 页"
	Lang[protocol.ZH_CN]["pager"]["summery"] = "第 <strong>%s-%s</strong> 项，共 <strong>%s</strong> 项"
	Lang[protocol.ZH_CN]["pager"]["nextPage"] = "下一页"
	Lang[protocol.ZH_CN]["pager"]["previousPage"] = "上一页"
	Lang[protocol.ZH_CN]["pager"]["locate"] = "GO!"
	Lang[protocol.ZH_CN]["pager"]["last"] = "<i class='icon-step-forward' title='末页'></i>"
	Lang[protocol.ZH_CN]["pager"]["next"] = "<i class='icon-play' title='下一页'></i>"
	Lang[protocol.ZH_CN]["pager"]["pre"] = "<i class='icon-play icon-flip-horizontal' title='上一页'></i>"
	Lang[protocol.ZH_CN]["pager"]["first"] = "<i class='icon-step-backward' title='首页'></i>"
	Lang[protocol.ZH_CN]["pager"]["recPerPage"] = "每页 <strong>%s</strong> 条"
	Lang[protocol.ZH_CN]["pager"]["digest"] = "共 <strong>%s</strong> 条记录，%s <strong>%s/%s</strong> &nbsp; "
	Lang[protocol.ZH_CN]["pager"]["noRecord"] = "暂时没有记录"
	Lang[protocol.ZH_CN]["colorPicker"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["colorPicker"]["errorTip"] = "不是有效的颜色值"
	Lang[protocol.ZH_CN]["datepicker"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["datepicker"]["monthNames"] = []string{`一月`, `二月`, `三月`, `四月`, `五月`, `六月`, `七月`, `八月`, `九月`, `十月`, `十一月`, `十二月`}
	Lang[protocol.ZH_CN]["datepicker"]["abbrDayNames"] = []protocol.HtmlKeyValueStr{{"0", "日"}, {"1", "一"}, {"2", "二"}, {"3", "三"}, {"4", "四"}, {"5", "五"}, {"6", "六"}}
	Lang[protocol.ZH_CN]["datepicker"]["dayNames"] = []string{`星期日`, `星期一`, `星期二`, `星期三`, `星期四`, `星期五`, `星期六`}
	Lang[protocol.ZH_CN]["datepicker"]["dpText"] = map[string]string{
		"TEXT_OR":          "或 ",
		"TEXT_PREV_YEAR":   "去年",
		"TEXT_PREV_MONTH":  "上月",
		"TEXT_PREV_WEEK":   "上周",
		"TEXT_YESTERDAY":   "昨天",
		"TEXT_THIS_MONTH":  "本月",
		"TEXT_THIS_WEEK":   "本周",
		"TEXT_TODAY":       "今天",
		"TEXT_NEXT_YEAR":   "明年",
		"TEXT_NEXT_MONTH":  "下月",
		"TEXT_CLOSE":       "关闭",
		"TEXT_DATE":        "选择时间段",
		"TEXT_CHOOSE_DATE": "选择日期",
		"TEXT_WEEK_MONDAY": "本周一",
		"TEXT_WEEK_SUNDAY": "本周日",
		"TEXT_MONTH_BEGIN": "本月初",
		"TEXT_MONTH_END":   "本月末",
	}
	Lang[protocol.ZH_CN]["common"]["icons"] = map[string]string{
		"todo":               "check",
		"product":            "cube",
		"bug":                "bug",
		"task":               "check-sign",
		"tasks":              "tasks",
		"project":            "stack",
		"doc":                "file-text",
		"doclib":             "folder-close",
		"story":              "lightbulb",
		"release":            "tags",
		"roadmap":            "code-fork",
		"plan":               "flag",
		"dynamic":            "volume-up",
		"build":              "tag",
		"test":               "check",
		"testtask":           "check",
		"group":              "group",
		"team":               "group",
		"company":            "sitemap",
		"user":               "user",
		"dept":               "sitemap",
		"tree":               "sitemap",
		"usecase":            "sitemap",
		"testcase":           "sitemap",
		"result":             "list-alt",
		"mail":               "envelope",
		"trash":              "trash",
		"extension":          "th-large",
		"app":                "th-large",
		"results":            "list-alt",
		"create":             "plus",
		"post":               "edit",
		"batchCreate":        "plus-sign",
		"batchEdit":          "edit-sign",
		"batchClose":         "off",
		"edit":               "edit",
		"delete":             "close",
		"copy":               "copy",
		"report":             "bar-chart",
		"export":             "export",
		"import":             "import",
		"finish":             "checked",
		"resolve":            "check",
		"start":              "play",
		"restart":            "play",
		"run":                "play",
		"runCase":            "play",
		"batchRun":           "play-sign",
		"assign":             "hand-right",
		"assignTo":           "hand-right",
		"change":             "fork",
		"link":               "link",
		"close":              "off",
		"activate":           "magic",
		"review":             "glasses",
		"confirm":            "search",
		"confirmBug":         "search",
		"putoff":             "calendar",
		"suspend":            "pause",
		"pause":              "pause",
		"cancel":             "ban-circle",
		"recordEstimate":     "time",
		"customFields":       "cogs",
		"manage":             "cog",
		"unlock":             "unlock-alt",
		"confirmStoryChange": "search",
		"score":              "tint",
		"report-file":        "file-powerpoint",
	}
	Lang[protocol.ZH_CN]["my"]["subMenu"] = []protocol.HtmlKeyValueInterface{
		{"calendar", []protocol.HtmlMenu{
			{"effort", map[string]string{`link`: "日志|effort|calendar|"}},
		}},
		{"calendar", []protocol.HtmlMenu{
			{"todo", map[string]string{`link`: "待办|todo|calendar|"}},
		}},
		{"calendar", []protocol.HtmlMenu{
			{"todo", map[string]string{`link`: "待办|my|todo|"}},
			{"effort", map[string]string{`link`: "日志|my|effort|"}},
		}},
	}
	Lang[protocol.ZH_CN]["my"]["menuOrder"] = []string{"effort"}
	Lang[protocol.ZH_CN]["company"]["menuOrder"] = []string{"effort", "todo", "effort"}
	Lang[protocol.ZH_CN]["project"]["menuOrder"] = []string{"effort"}
	Lang[protocol.ZH_CN]["user"]["menuOrder"] = []string{"effort"}
	Lang[protocol.ZH_CN]["effort"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["effort"]["menu"] = Lang[protocol.ZH_CN]["my"]["menu"]
	Lang[protocol.ZH_CN]["excel"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["excel"]["errorTitle"] = "输入有误"
	Lang[protocol.ZH_CN]["excel"]["error"] = "您输入的值不在下拉框列表内。"
	Lang[protocol.ZH_CN]["excel"]["fileField"] = "附件"
	Lang[protocol.ZH_CN]["excel"]["insert"] = "全新插入"
	Lang[protocol.ZH_CN]["excel"]["canNotRead"] = "不能解析该文件"
	Lang[protocol.ZH_CN]["excel"]["noData"] = "没有数据"
	Lang[protocol.ZH_CN]["ldap"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["ldap"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]
	Lang[protocol.ZH_CN]["oa"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["oa"]["menu"] = []protocol.HtmlMenu{
		{"attend", map[string]string{`link`: `考勤|attend|personal`, `subModule`: `attend`}},
		{"leave", map[string]string{`link`: `请假|leave|personal`, `alias`: `browse`, `subModule`: `leave`}},
		{"overtime", map[string]string{`link`: `加班|overtime|personal`, `subModule`: `overtime`}},
		{"holiday", map[string]string{`link`: `节假日|holiday|browse`, `subModule`: `holiday`}},
		{"trip", map[string]string{`link`: `外出|trip|personal`, `subModule`: `trip`}},
		{"review", map[string]string{`link`: `审批|my|review|type=all&orderBy=status&from=oa`}},
	}
	Lang[protocol.ZH_CN]["oa"]["menuOrder"] = []string{"attend", "leave", "makeup", "overtime", "lieu", "holiday", "trip", "review"}
	Lang[protocol.ZH_CN]["attend"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["attend"]["allAbsent"] = "旷工一天迟到早退分钟数"
	Lang[protocol.ZH_CN]["attend"]["halfAbsent"] = "旷工半天迟到早退分钟数"
	Lang[protocol.ZH_CN]["attend"]["total"] = "共"
	Lang[protocol.ZH_CN]["attend"]["excelday"] = "日"
	Lang[protocol.ZH_CN]["attend"]["notSignOut"] = "下班缺卡次数"
	Lang[protocol.ZH_CN]["attend"]["notSignIn"] = "上班缺卡次数"
	Lang[protocol.ZH_CN]["attend"]["earlyMin"] = "早退时长(分钟)"
	Lang[protocol.ZH_CN]["attend"]["lateMin"] = "迟到时长(分钟)"
	Lang[protocol.ZH_CN]["attend"]["workmin"] = "工作时长(分钟)"
	Lang[protocol.ZH_CN]["attend"]["attendanceDays"] = "出勤天数"
	Lang[protocol.ZH_CN]["attend"]["role"] = "职位"
	Lang[protocol.ZH_CN]["attend"]["dept"] = "部门"
	Lang[protocol.ZH_CN]["attend"]["realname"] = "姓名"
	Lang[protocol.ZH_CN]["attend"]["confirmEarly"] = "还未到最晚签退时间，是否签退？"
	Lang[protocol.ZH_CN]["attend"]["editAction"] = "补录"
	Lang[protocol.ZH_CN]["attend"]["detailAction"] = "考勤明细"
	Lang[protocol.ZH_CN]["attend"]["reportAction"] = "统计考勤"
	Lang[protocol.ZH_CN]["attend"]["exportAction"] = "导出考勤"
	Lang[protocol.ZH_CN]["attend"]["saveStatAction"] = "保存考勤统计"
	Lang[protocol.ZH_CN]["attend"]["signOutLimitError"] = "当前时间不允许打卡下班，请在%s之后打卡下班"
	Lang[protocol.ZH_CN]["attend"]["notAllowSignOutLimitError"] = "不允许在%s之后打卡下班"
	Lang[protocol.ZH_CN]["attend"]["notAllowSignInLimitError"] = "不允许在%s之前打卡上班"
	Lang[protocol.ZH_CN]["attend"]["notAllowSignOutLimit"] = "最晚打卡下班"
	Lang[protocol.ZH_CN]["attend"]["notAllowSignInLimit"] = "最早打卡上班"
	Lang[protocol.ZH_CN]["attend"]["cancel"] = "取消补录"
	Lang[protocol.ZH_CN]["attend"]["waitReviews"] = "<strong>%s</strong> 存在未审批的记录，请审批之后再进行统计。"
	Lang[protocol.ZH_CN]["attend"]["signInClientError"] = "签到失败！已设置只能通过 %s 签到。"
	Lang[protocol.ZH_CN]["attend"]["s"] = "秒"
	Lang[protocol.ZH_CN]["attend"]["m"] = "分钟"
	Lang[protocol.ZH_CN]["attend"]["h"] = "小时"
	Lang[protocol.ZH_CN]["attend"]["d"] = "天"
	Lang[protocol.ZH_CN]["attend"]["setDept"] = "部门设置"
	Lang[protocol.ZH_CN]["attend"]["setManager"] = "部门经理设置"
	Lang[protocol.ZH_CN]["attend"]["personalSettings"] = "个人考勤设置"
	Lang[protocol.ZH_CN]["attend"]["weeks"] = []string{`第一周`, `第二周`, `第三周`, `第四周`, `第五周`, `第六周`}
	Lang[protocol.ZH_CN]["attend"]["reviewSuccess"] = "审核成功"
	Lang[protocol.ZH_CN]["attend"]["nodata"] = "没有选择数据"
	Lang[protocol.ZH_CN]["attend"]["denied"] = "信息访问受限"
	Lang[protocol.ZH_CN]["attend"]["mustSignOut"] = "必须签退"
	Lang[protocol.ZH_CN]["attend"]["workingHours"] = "每天工作工时"
	Lang[protocol.ZH_CN]["attend"]["workingDays"] = "每周工作天数"
	Lang[protocol.ZH_CN]["attend"]["signOutLimit"] = "最早打卡下班"
	Lang[protocol.ZH_CN]["attend"]["signInLimit"] = "最晚打卡上班"
	Lang[protocol.ZH_CN]["attend"]["outFail"] = "签退失败"
	Lang[protocol.ZH_CN]["attend"]["outSuccess"] = "签退成功"
	Lang[protocol.ZH_CN]["attend"]["inFail"] = "签到失败"
	Lang[protocol.ZH_CN]["attend"]["inSuccess"] = "签到成功"
	Lang[protocol.ZH_CN]["attend"]["actualDays"] = "实际出勤天数"
	Lang[protocol.ZH_CN]["attend"]["deserveDays"] = "应出勤天数"
	Lang[protocol.ZH_CN]["attend"]["reviewedDate"] = "审核时间"
	Lang[protocol.ZH_CN]["attend"]["reviewedBy"] = "审核人"
	Lang[protocol.ZH_CN]["attend"]["reviewStatus"] = "补录状态"
	Lang[protocol.ZH_CN]["attend"]["reason"] = "原因"
	Lang[protocol.ZH_CN]["attend"]["manualOut"] = "签退时间"
	Lang[protocol.ZH_CN]["attend"]["manualIn"] = "签到时间"
	Lang[protocol.ZH_CN]["attend"]["search"] = "搜索"
	Lang[protocol.ZH_CN]["attend"]["end"] = "截至"
	Lang[protocol.ZH_CN]["attend"]["begin"] = "开始"
	Lang[protocol.ZH_CN]["attend"]["user"] = "用户"
	Lang[protocol.ZH_CN]["attend"]["rejectReason"] = "拒绝原因"
	Lang[protocol.ZH_CN]["attend"]["signInClient"] = "签到途径"
	Lang[protocol.ZH_CN]["attend"]["noAttendUsers"] = "无需考勤者"
	Lang[protocol.ZH_CN]["attend"]["noAttendDepts"] = "无需考勤的部门"
	Lang[protocol.ZH_CN]["attend"]["ipList"] = "IP列表"
	Lang[protocol.ZH_CN]["attend"]["PM"] = "下午"
	Lang[protocol.ZH_CN]["attend"]["AM"] = "上午"
	Lang[protocol.ZH_CN]["attend"]["report"] = "考勤表"
	Lang[protocol.ZH_CN]["attend"]["dayName"] = "星期"
	Lang[protocol.ZH_CN]["attend"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["attend"]["device"] = "设备"
	Lang[protocol.ZH_CN]["attend"]["ip"] = "IP"
	Lang[protocol.ZH_CN]["attend"]["status"] = "状态"
	Lang[protocol.ZH_CN]["attend"]["signOut"] = "签退"
	Lang[protocol.ZH_CN]["attend"]["signIn"] = "签到"
	Lang[protocol.ZH_CN]["attend"]["account"] = "用户"
	Lang[protocol.ZH_CN]["attend"]["date"] = "日期"
	Lang[protocol.ZH_CN]["attend"]["id"] = "编号"
	Lang[protocol.ZH_CN]["attend"]["clockOut"] = "下班打卡"
	Lang[protocol.ZH_CN]["attend"]["clockIn"] = "上班打卡"
	Lang[protocol.ZH_CN]["attend"]["clockinout"] = "首页上班打卡"
	Lang[protocol.ZH_CN]["attend"]["batchPass"] = "批量通过"
	Lang[protocol.ZH_CN]["attend"]["batchReview"] = "批量审核"
	Lang[protocol.ZH_CN]["attend"]["browseReview"] = "补录列表"
	Lang[protocol.ZH_CN]["attend"]["exportDetail"] = "导出考勤明细"
	Lang[protocol.ZH_CN]["attend"]["exportStat"] = "导出考勤统计表"
	Lang[protocol.ZH_CN]["attend"]["saveStat"] = "保存考勤统计"
	Lang[protocol.ZH_CN]["attend"]["stat"] = "统计"
	Lang[protocol.ZH_CN]["attend"]["export"] = "导出"
	Lang[protocol.ZH_CN]["attend"]["settings"] = `公司考勤设置` //"设置"替换`公司考勤设置`
	Lang[protocol.ZH_CN]["attend"]["review"] = `审核补录`     //"补录审核"替换`审核补录`
	Lang[protocol.ZH_CN]["attend"]["overtimed"] = "已加班"
	Lang[protocol.ZH_CN]["attend"]["overtime"] = "加班"
	Lang[protocol.ZH_CN]["attend"]["egressed"] = "已外出"
	Lang[protocol.ZH_CN]["attend"]["egress"] = "外出"
	Lang[protocol.ZH_CN]["attend"]["triped"] = "已出差"
	Lang[protocol.ZH_CN]["attend"]["trip"] = "出差"
	Lang[protocol.ZH_CN]["attend"]["lieud"] = "已调休"
	Lang[protocol.ZH_CN]["attend"]["lieu"] = "调休"
	Lang[protocol.ZH_CN]["attend"]["makeuped"] = "已补班"
	Lang[protocol.ZH_CN]["attend"]["makeup"] = "补班"
	Lang[protocol.ZH_CN]["attend"]["leaved"] = "已请假"
	Lang[protocol.ZH_CN]["attend"]["leave"] = "请假"
	Lang[protocol.ZH_CN]["attend"]["edited"] = "补录审核"
	Lang[protocol.ZH_CN]["attend"]["edit"] = "补录"
	Lang[protocol.ZH_CN]["attend"]["detail"] = "考勤明细"
	Lang[protocol.ZH_CN]["attend"]["company"] = `公司考勤`    //"公司考勤"替换`公司考勤`
	Lang[protocol.ZH_CN]["attend"]["department"] = `部门考勤` //"部门考勤"替换`部门考勤`
	Lang[protocol.ZH_CN]["attend"]["personal"] = "我的考勤"
	Lang[protocol.ZH_CN]["attend"]["common"] = "考勤"
	Lang[protocol.ZH_CN]["attend"]["menu"] = Lang[protocol.ZH_CN]["oa"]["menu"]
	Lang[protocol.ZH_CN]["leave"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["leave"]["lieulist"] = "设置查看加班调休详情"
	Lang[protocol.ZH_CN]["leave"]["setReviewerAction"] = "设置审批人"
	Lang[protocol.ZH_CN]["leave"]["backAction"] = "销假"
	Lang[protocol.ZH_CN]["leave"]["viewAction"] = "请假详情"
	Lang[protocol.ZH_CN]["leave"]["deleteAction"] = "删除请假"
	Lang[protocol.ZH_CN]["leave"]["editAction"] = "编辑请假"
	Lang[protocol.ZH_CN]["leave"]["createAction"] = "申请请假"
	Lang[protocol.ZH_CN]["leave"]["exportAction"] = "导出请假记录"
	Lang[protocol.ZH_CN]["leave"]["reviewAction"] = "审核请假"
	Lang[protocol.ZH_CN]["leave"]["companyAction"] = "所有请假"
	Lang[protocol.ZH_CN]["leave"]["switchstatus"] = "提交或撤销"
	Lang[protocol.ZH_CN]["leave"]["getDays"] = "获取时间（申请请假必选）"
	Lang[protocol.ZH_CN]["leave"]["entryDay"] = "入职时间"
	Lang[protocol.ZH_CN]["leave"]["lieurealday"] = "实际天数"
	Lang[protocol.ZH_CN]["leave"]["leaderTipError"] = "请告知 %s 你请假的时间与工作安排"
	Lang[protocol.ZH_CN]["leave"]["leaderTip"] = "我已告知以下人员 ，并且安排好后续工作"
	Lang[protocol.ZH_CN]["leave"]["lieuoffsetday"] = "额外天数"
	Lang[protocol.ZH_CN]["leave"]["lieuday"] = "可调休天数"
	Lang[protocol.ZH_CN]["leave"]["lieuTip"] = "可用调休 %s 天"
	Lang[protocol.ZH_CN]["leave"]["annualTip"] = "可用年假 %s 天"
	Lang[protocol.ZH_CN]["leave"]["daysTip2"] = "请手动减去节假日"
	Lang[protocol.ZH_CN]["leave"]["daysTip"] = "天"
	Lang[protocol.ZH_CN]["leave"]["hoursTip"] = "小时"
	Lang[protocol.ZH_CN]["leave"]["lieudenied 	"] = "可申请调休时间不够，请删除待审核、拒绝记录，现在可申请%s天"
	Lang[protocol.ZH_CN]["leave"]["annualdenied"] = "可申请年假时间不够，请删除待审核、拒绝记录，现在可申请%s天"
	Lang[protocol.ZH_CN]["leave"]["reviewSuccess"] = "审核成功"
	Lang[protocol.ZH_CN]["leave"]["nodata"] = "没有选择数据"
	Lang[protocol.ZH_CN]["leave"]["wrongBackDate"] = "报到时间应该大于开始时间"
	Lang[protocol.ZH_CN]["leave"]["wrongEnd"] = "结束时间应该大于开始时间"
	Lang[protocol.ZH_CN]["leave"]["sameMonth"] = "不支持跨月份请假"
	Lang[protocol.ZH_CN]["leave"]["unique"] = "%s 已经存在请假记录"
	Lang[protocol.ZH_CN]["leave"]["denied"] = "信息访问受限"
	Lang[protocol.ZH_CN]["leave"]["notExist"] = "记录不存在"
	Lang[protocol.ZH_CN]["leave"]["unpaid"] = "非带薪假"
	Lang[protocol.ZH_CN]["leave"]["paid"] = "带薪假"
	Lang[protocol.ZH_CN]["leave"]["totallieu"] = "个人调休总天数"
	Lang[protocol.ZH_CN]["leave"]["totalAnnual"] = "个人年假总天数"
	Lang[protocol.ZH_CN]["leave"]["dateRange"] = "起止时间"
	Lang[protocol.ZH_CN]["leave"]["account"] = "用户"
	Lang[protocol.ZH_CN]["leave"]["rejectReason"] = "拒绝原因"
	Lang[protocol.ZH_CN]["leave"]["time"] = "时间"
	Lang[protocol.ZH_CN]["leave"]["date"] = "日期"
	Lang[protocol.ZH_CN]["leave"]["reviewedDate"] = "审核时间"
	Lang[protocol.ZH_CN]["leave"]["reviewedBy"] = "审核者"
	Lang[protocol.ZH_CN]["leave"]["createdDate"] = "申请时间"
	Lang[protocol.ZH_CN]["leave"]["createdBy"] = "申请者"
	Lang[protocol.ZH_CN]["leave"]["status"] = "状态"
	Lang[protocol.ZH_CN]["leave"]["desc"] = "事由"
	Lang[protocol.ZH_CN]["leave"]["type"] = "类型"
	Lang[protocol.ZH_CN]["leave"]["backDate"] = "报到时间"
	Lang[protocol.ZH_CN]["leave"]["hours"] = "总时长"
	Lang[protocol.ZH_CN]["leave"]["finish"] = "结束时间"
	Lang[protocol.ZH_CN]["leave"]["start"] = "开始时间"
	Lang[protocol.ZH_CN]["leave"]["end"] = "结束"
	Lang[protocol.ZH_CN]["leave"]["begin"] = "开始"
	Lang[protocol.ZH_CN]["leave"]["year"] = "年"
	Lang[protocol.ZH_CN]["leave"]["id"] = "编号"
	Lang[protocol.ZH_CN]["leave"]["batchPass"] = "批量通过"
	Lang[protocol.ZH_CN]["leave"]["batchReview"] = "批量审核"
	Lang[protocol.ZH_CN]["leave"]["personallieu 	"] = "个人加班调休"
	Lang[protocol.ZH_CN]["leave"]["personalAnnual"] = "个人年假"
	Lang[protocol.ZH_CN]["leave"]["setReviewer"] = "审批人"
	Lang[protocol.ZH_CN]["leave"]["company"] = "所有请假"
	Lang[protocol.ZH_CN]["leave"]["browseReview"] = "审核列表"
	Lang[protocol.ZH_CN]["leave"]["personal"] = "我的请假"
	Lang[protocol.ZH_CN]["leave"]["personalieu"] = "设置个人调休"
	Lang[protocol.ZH_CN]["leave"]["personalannual"] = "设置个人年假"
	Lang[protocol.ZH_CN]["leave"]["reviewBack"] = "审核销假"
	Lang[protocol.ZH_CN]["leave"]["export"] = "导出请假记录"
	Lang[protocol.ZH_CN]["leave"]["back"] = "销假"
	Lang[protocol.ZH_CN]["leave"]["commit"] = "提交"
	Lang[protocol.ZH_CN]["leave"]["cancel"] = "撤销"
	Lang[protocol.ZH_CN]["leave"]["review"] = "审核"
	Lang[protocol.ZH_CN]["leave"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["leave"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["leave"]["create"] = "申请请假"
	Lang[protocol.ZH_CN]["leave"]["view"] = "详情"
	Lang[protocol.ZH_CN]["leave"]["browse"] = "请假列表"
	Lang[protocol.ZH_CN]["leave"]["common"] = "请假"
	Lang[protocol.ZH_CN]["leave"]["menu"] = Lang[protocol.ZH_CN]["oa"]["menu"]
	Lang[protocol.ZH_CN]["makeup"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["makeup"]["menu"] = Lang[protocol.ZH_CN]["oa"]["menu"]
	Lang[protocol.ZH_CN]["overtime"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["overtime"]["setReviewerAction"] = "加班设置"
	Lang[protocol.ZH_CN]["overtime"]["viewAction"] = "加班详情"
	Lang[protocol.ZH_CN]["overtime"]["deleteAction"] = "删除加班"
	Lang[protocol.ZH_CN]["overtime"]["editAction"] = "编辑加班"
	Lang[protocol.ZH_CN]["overtime"]["createAction"] = "申请加班"
	Lang[protocol.ZH_CN]["overtime"]["exportAction"] = "导出加班记录"
	Lang[protocol.ZH_CN]["overtime"]["reviewAction"] = "审核加班"
	Lang[protocol.ZH_CN]["overtime"]["companyAction"] = "所有加班"
	Lang[protocol.ZH_CN]["overtime"]["switchstatus"] = "提交或撤销"
	Lang[protocol.ZH_CN]["overtime"]["daysTip"] = "天"
	Lang[protocol.ZH_CN]["overtime"]["hoursTip"] = "小时"
	Lang[protocol.ZH_CN]["overtime"]["reviewSuccess"] = "审核成功"
	Lang[protocol.ZH_CN]["overtime"]["nodata"] = "没有选择数据"
	Lang[protocol.ZH_CN]["overtime"]["wrongEnd"] = "结束时间应该大于开始时间"
	Lang[protocol.ZH_CN]["overtime"]["sameMonth"] = "不支持跨月份加班"
	Lang[protocol.ZH_CN]["overtime"]["unique"] = "%s 已经存在加班记录"
	Lang[protocol.ZH_CN]["overtime"]["denied"] = "信息访问受限"
	Lang[protocol.ZH_CN]["overtime"]["notExist"] = "记录不存在"
	Lang[protocol.ZH_CN]["overtime"]["rejectReason"] = "拒绝原因"
	Lang[protocol.ZH_CN]["overtime"]["time"] = "时间"
	Lang[protocol.ZH_CN]["overtime"]["date"] = "日期"
	Lang[protocol.ZH_CN]["overtime"]["reviewedDate"] = "审核时间"
	Lang[protocol.ZH_CN]["overtime"]["reviewedBy"] = "审核者"
	Lang[protocol.ZH_CN]["overtime"]["createdDate"] = "申请时间"
	Lang[protocol.ZH_CN]["overtime"]["createdBy"] = "申请者"
	Lang[protocol.ZH_CN]["overtime"]["status"] = "状态"
	Lang[protocol.ZH_CN]["overtime"]["desc"] = "事由"
	Lang[protocol.ZH_CN]["overtime"]["type"] = "类型"
	Lang[protocol.ZH_CN]["overtime"]["leave"] = "请假记录"
	Lang[protocol.ZH_CN]["overtime"]["hours"] = "总时长"
	Lang[protocol.ZH_CN]["overtime"]["finish"] = "结束时间"
	Lang[protocol.ZH_CN]["overtime"]["start"] = "开始时间"
	Lang[protocol.ZH_CN]["overtime"]["end"] = "结束"
	Lang[protocol.ZH_CN]["overtime"]["begin"] = "开始"
	Lang[protocol.ZH_CN]["overtime"]["year"] = "年"
	Lang[protocol.ZH_CN]["overtime"]["id"] = "编号"
	Lang[protocol.ZH_CN]["overtime"]["batchPass"] = "批量通过"
	Lang[protocol.ZH_CN]["overtime"]["batchReview"] = "批量审核"
	Lang[protocol.ZH_CN]["overtime"]["setReviewer"] = "加班设置"
	Lang[protocol.ZH_CN]["overtime"]["company"] = "所有加班"
	Lang[protocol.ZH_CN]["overtime"]["browseReview"] = "审核列表"
	Lang[protocol.ZH_CN]["overtime"]["personal"] = "我的加班"
	Lang[protocol.ZH_CN]["overtime"]["export"] = "导出加班记录"
	Lang[protocol.ZH_CN]["overtime"]["commit"] = "提交"
	Lang[protocol.ZH_CN]["overtime"]["cancel"] = "撤销"
	Lang[protocol.ZH_CN]["overtime"]["review"] = "审核"
	Lang[protocol.ZH_CN]["overtime"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["overtime"]["view"] = "详情"
	Lang[protocol.ZH_CN]["overtime"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["overtime"]["create"] = "申请加班"
	Lang[protocol.ZH_CN]["overtime"]["browse"] = "加班列表"
	Lang[protocol.ZH_CN]["overtime"]["common"] = "加班"
	Lang[protocol.ZH_CN]["overtime"]["menu"] = Lang[protocol.ZH_CN]["oa"]["menu"]
	Lang[protocol.ZH_CN]["lieu"] = make(map[string]interface{})

	Lang[protocol.ZH_CN]["lieu"]["setReviewerAction"] = "调休设置"
	Lang[protocol.ZH_CN]["lieu"]["viewAction"] = "调休详情"
	Lang[protocol.ZH_CN]["lieu"]["deleteAction"] = "删除调休"
	Lang[protocol.ZH_CN]["lieu"]["editAction"] = "编辑调休"
	Lang[protocol.ZH_CN]["lieu"]["createAction"] = "申请调休"
	Lang[protocol.ZH_CN]["lieu"]["companyAction"] = "所有调休"
	Lang[protocol.ZH_CN]["lieu"]["reviewAction"] = "审核调休"
	Lang[protocol.ZH_CN]["lieu"]["browseReviewAction"] = "调休审核列表"
	Lang[protocol.ZH_CN]["lieu"]["switchstatus"] = "提交或撤销"
	Lang[protocol.ZH_CN]["lieu"]["daysTip"] = "天"
	Lang[protocol.ZH_CN]["lieu"]["hoursTip"] = "小时"
	Lang[protocol.ZH_CN]["lieu"]["bothEmpty"] = `<strong>加班记录</strong>不能为空` //"<strong>加班记录</strong>和<strong>出差记录</strong>不能同时为空"替换`<strong>加班记录</strong>不能为空`
	Lang[protocol.ZH_CN]["lieu"]["nobccomp"] = "请安装bcmath扩展"
	Lang[protocol.ZH_CN]["lieu"]["wrongHours"] = `加班总时长 <strong>%s</strong> 小时，调休时长不能超过总时长。` //"加班和出差总时长 <strong>%s</strong> 小时，调休时长不能超过总时长。"替换`加班总时长 <strong>%s</strong> 小时，调休时长不能超过总时长。`
	Lang[protocol.ZH_CN]["lieu"]["reviewSuccess"] = "审核成功"
	Lang[protocol.ZH_CN]["lieu"]["nodata"] = "没有选择数据"
	Lang[protocol.ZH_CN]["lieu"]["wrongEnd"] = "结束时间应该大于开始时间"
	Lang[protocol.ZH_CN]["lieu"]["sameMonth"] = "不支持跨月份调休"
	Lang[protocol.ZH_CN]["lieu"]["unique"] = "%s 已经存在调休记录"
	Lang[protocol.ZH_CN]["lieu"]["denied"] = "信息访问受限"
	Lang[protocol.ZH_CN]["lieu"]["checkHours"] = "调休时长检测"
	Lang[protocol.ZH_CN]["lieu"]["notExist"] = "记录不存在"
	Lang[protocol.ZH_CN]["lieu"]["rejectReason"] = "拒绝原因"
	Lang[protocol.ZH_CN]["lieu"]["time"] = "时间"
	Lang[protocol.ZH_CN]["lieu"]["date"] = "日期"
	Lang[protocol.ZH_CN]["lieu"]["reviewedDate"] = "审核时间"
	Lang[protocol.ZH_CN]["lieu"]["reviewedBy"] = "审核者"
	Lang[protocol.ZH_CN]["lieu"]["createdDate"] = "申请时间"
	Lang[protocol.ZH_CN]["lieu"]["createdBy"] = "申请者"
	Lang[protocol.ZH_CN]["lieu"]["desc"] = "事由"
	Lang[protocol.ZH_CN]["lieu"]["status"] = "状态"
	Lang[protocol.ZH_CN]["lieu"]["trip"] = "出差记录"
	Lang[protocol.ZH_CN]["lieu"]["overtime"] = "加班记录"
	Lang[protocol.ZH_CN]["lieu"]["hours"] = "总时长"
	Lang[protocol.ZH_CN]["lieu"]["finish"] = "结束时间"
	Lang[protocol.ZH_CN]["lieu"]["start"] = "开始时间"
	Lang[protocol.ZH_CN]["lieu"]["end"] = "结束"
	Lang[protocol.ZH_CN]["lieu"]["begin"] = "开始"
	Lang[protocol.ZH_CN]["lieu"]["year"] = "年"
	Lang[protocol.ZH_CN]["lieu"]["id"] = "编号"
	Lang[protocol.ZH_CN]["lieu"]["mydays"] = "可调休天数"
	Lang[protocol.ZH_CN]["lieu"]["batchPass"] = "批量通过"
	Lang[protocol.ZH_CN]["lieu"]["batchReview"] = "批量审核"
	Lang[protocol.ZH_CN]["lieu"]["setReviewer"] = "调休设置"
	Lang[protocol.ZH_CN]["lieu"]["company"] = "所有调休"
	Lang[protocol.ZH_CN]["lieu"]["browseReview"] = "审核列表"
	Lang[protocol.ZH_CN]["lieu"]["personal"] = "我的调休"
	Lang[protocol.ZH_CN]["lieu"]["commit"] = "提交"
	Lang[protocol.ZH_CN]["lieu"]["cancel"] = "撤销"
	Lang[protocol.ZH_CN]["lieu"]["review"] = "审核"
	Lang[protocol.ZH_CN]["lieu"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["lieu"]["view"] = "详情"
	Lang[protocol.ZH_CN]["lieu"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["lieu"]["create"] = "申请调休"
	Lang[protocol.ZH_CN]["lieu"]["browse"] = "调休列表"
	Lang[protocol.ZH_CN]["lieu"]["common"] = "调休"
	Lang[protocol.ZH_CN]["lieu"]["menu"] = Lang[protocol.ZH_CN]["oa"]["menu"]
	Lang[protocol.ZH_CN]["holiday"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["holiday"]["deleteAction"] = "删除节假日"
	Lang[protocol.ZH_CN]["holiday"]["editAction"] = "编辑节假日"
	Lang[protocol.ZH_CN]["holiday"]["createAction"] = "创建节假日"
	Lang[protocol.ZH_CN]["holiday"]["actor"] = "操作人"
	Lang[protocol.ZH_CN]["holiday"]["holiday"] = "假期"
	Lang[protocol.ZH_CN]["holiday"]["date"] = "日期"
	Lang[protocol.ZH_CN]["holiday"]["end"] = "结束日期"
	Lang[protocol.ZH_CN]["holiday"]["begin"] = "开始日期"
	Lang[protocol.ZH_CN]["holiday"]["option"] = "类型"
	Lang[protocol.ZH_CN]["holiday"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["holiday"]["name"] = "名称"
	Lang[protocol.ZH_CN]["holiday"]["id"] = "编号"
	Lang[protocol.ZH_CN]["holiday"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["holiday"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["holiday"]["create"] = "新建"
	Lang[protocol.ZH_CN]["holiday"]["browse"] = `浏览节假日` //"浏览"替换`浏览节假日`
	Lang[protocol.ZH_CN]["holiday"]["common"] = "节假日"
	Lang[protocol.ZH_CN]["holiday"]["menu"] = Lang[protocol.ZH_CN]["oa"]["menu"]
	Lang[protocol.ZH_CN]["trip"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["trip"]["setReviewer"] = "外出设置"
	Lang[protocol.ZH_CN]["trip"]["reviewedBy"] = "审核者"
	Lang[protocol.ZH_CN]["trip"]["type"] = "类型"
	Lang[protocol.ZH_CN]["trip"]["sameMonth"] = "不支持跨月份出差"
	Lang[protocol.ZH_CN]["trip"]["wrongEnd"] = "结束时间应该大于开始时间"
	Lang[protocol.ZH_CN]["trip"]["unique"] = "%s 已经存在出差记录"
	Lang[protocol.ZH_CN]["trip"]["denied"] = "信息访问受限"
	Lang[protocol.ZH_CN]["trip"]["time"] = "时间"
	Lang[protocol.ZH_CN]["trip"]["date"] = "日期"
	Lang[protocol.ZH_CN]["trip"]["createdDate"] = "创建时间"
	Lang[protocol.ZH_CN]["trip"]["createdBy"] = "创建者"
	Lang[protocol.ZH_CN]["trip"]["desc"] = "事由"
	Lang[protocol.ZH_CN]["trip"]["to"] = "目的地"
	Lang[protocol.ZH_CN]["trip"]["from"] = "出发城市"
	Lang[protocol.ZH_CN]["trip"]["end"] = "结束"
	Lang[protocol.ZH_CN]["trip"]["begin"] = "开始"
	Lang[protocol.ZH_CN]["trip"]["name"] = "名称"
	Lang[protocol.ZH_CN]["trip"]["customer"] = "客户 / 供应商"
	Lang[protocol.ZH_CN]["trip"]["id"] = "编号"
	Lang[protocol.ZH_CN]["trip"]["setReviewerAction"] = "外出设置"
	Lang[protocol.ZH_CN]["trip"]["viewAction"] = "外出详情"
	Lang[protocol.ZH_CN]["trip"]["deleteAction"] = "删除外出"
	Lang[protocol.ZH_CN]["trip"]["editAction"] = "编辑外出"
	Lang[protocol.ZH_CN]["trip"]["createAction"] = "申请外出"
	Lang[protocol.ZH_CN]["trip"]["browseReview"] = "外出审核"
	Lang[protocol.ZH_CN]["trip"]["company"] = "所有外出"
	Lang[protocol.ZH_CN]["trip"]["department"] = "部门"
	Lang[protocol.ZH_CN]["trip"]["personal"] = "我的外出"
	Lang[protocol.ZH_CN]["trip"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["trip"]["view"] = "外出详情"
	Lang[protocol.ZH_CN]["trip"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["trip"]["create"] = "新建外出"
	Lang[protocol.ZH_CN]["trip"]["browse"] = "外出列表"
	Lang[protocol.ZH_CN]["trip"]["common"] = "外出"
	Lang[protocol.ZH_CN]["trip"]["menu"] = Lang[protocol.ZH_CN]["oa"]["menu"]
	Lang[protocol.ZH_CN]["attend"]["featurebar"] = []protocol.HtmlKeyValueStr{
		{"personal", "我的考勤|attend|personal|"},
		{"company", "公司考勤|attend|company|"},
		{"detail", "考勤明细|attend|detail|"},
		{"browsereview", "补录审核|attend|browsereview|"},
		{"stat", "统计|attend|stat|"},
		{"settings", `设置|attend|settings|`},
	}
	Lang[protocol.ZH_CN]["leave"]["featurebar"] = []protocol.HtmlKeyValueStr{
		{"personal", "我的请假|leave|personal|"},
		{"browseReview", "我的审核|leave|browsereview|"},
		{"company", "所有请假|leave|company|"},
		{"setReviewer", "设置|leave|setReviewer|"},
	}
	Lang[protocol.ZH_CN]["makeup"]["featurebar"] = []protocol.HtmlKeyValueStr{
		{"personal", "我的补班|makeup|personal|"},
		{"browseReview", "我的审核|makeup|browsereview|"},
		{"company", "所有补班|makeup|company|"},
		{"setReviewer", "设置|makeup|setReviewer|"},
	}
	Lang[protocol.ZH_CN]["overtime"]["featurebar"] = []protocol.HtmlKeyValueStr{
		{"personal", "我的加班|overtime|personal|"},
		{"browseReview", "我的审核|overtime|browsereview|"},
		{"company", "所有加班|overtime|company|"},
		{"setReviewer", "设置|overtime|setReviewer|"},
	}
	Lang[protocol.ZH_CN]["lieu"]["featurebar"] = []protocol.HtmlKeyValueStr{
		{"personal", "我的调休|lieu|personal|"},
		{"browseReview", "我的审核|lieu|browsereview|"},
		{"company", "所有调休|lieu|company|"},
		{"setReviewer", "设置|lieu|setReviewer|"},
	}
	Lang[protocol.ZH_CN]["holiday"]["featurebar"] = []protocol.HtmlKeyValueStr{
		{"browse", "所有|holiday|browse|"},
	}
	Lang[protocol.ZH_CN]["trip"]["featurebar"] = []protocol.HtmlKeyValueStr{
		{"personal", "我的外出|trip|personal|"},
		{"company", "所有外出|trip|company|"},
		{"browseReview", "审批|trip|browseReview|"},
		{"setReviewer", "设置|trip|setReviewer|"},
	}
	Lang[protocol.ZH_CN]["ci"]["menuOrder"] = []string{"review"}
	Lang[protocol.ZH_CN]["repo"]["menuOrder"] = []string{"review"}
	Lang[protocol.ZH_CN]["jenkins"]["menuOrder"] = []string{"review"}
	Lang[protocol.ZH_CN]["report"]["methodOrder"] = []string{"testcase", "build", "workSummary"}
	Lang[protocol.ZH_CN]["sms"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["sms"]["menu"] = Lang[protocol.ZH_CN]["admin"]["menu"]
	Lang[protocol.ZH_CN]["my"]["subMenu"] = []protocol.HtmlKeyValueInterface{}
	Lang[protocol.ZH_CN]["action"]["periods"] = map[string]string{
		"all":       "所有",
		"lastMonth": "上月",
		"lastWeek":  "上周",
		"thisMonth": "本月",
		"thisWeek":  "本周",
		"today":     "今天",
		"yesterday": "昨天",
	}
	Lang[protocol.ZH_CN]["action"]["label"] = map[string]interface{}{
		"activated": "激活了",
		"asked":     "追问了",
		"assigned":  "指派了",
		"attend": map[string]string{
			"commited": "考勤|attend|browsereview|",
			"reviewed": "考勤|attend|personal|",
		},
		"batchexamine":      "批量审核了",
		"batchproofreading": "批量对单完成",
		"blocked":           "阻塞了",
		"bug":               "Bug|bug|view|bugID=%s",
		"bugconfirmed":      "确认了",
		"build":             "版本|build|view|buildID=%s",
		"canceled":          "取消了",
		"case":              "用例|testcase|view|caseID=%s",
		"caselib":           "用例库|testsuite|libview|libID=%s",
		"changed":           "变更了",
		"changestatus":      "修改状态",
		"closed":            "关闭了",
		"commented":         "评论了",
		"commited":          "提交了",
		"confirmed":         "确认了需求",
		"created":           "创建",
		"delayed":           "延期",
		"deleted":           "删除了",
		"deletedfile":       "删除附件",
		"deleteestimate":    "删除了工时",
		"deploy":            "上线计划|deploy|view|id=%s",
		"deploystep":        "上线步骤|deploy|viewStep|id=%s",
		"doc":               "文档|doc|view|docID=%s",
		"doclib":            "文档库|doc|browse|libID=%s",
		"edited":            "编辑了",
		"editestimate":      "编辑了工时",
		"editfile":          "编辑附件",
		"effort":            "日志|effort|view|effortID=%s",
		"entry":             "应用|entry|browse|",
		"erased":            "删除了",
		"examine":           "审核了",
		"feedback":          "反馈|feedback|view|id=%s",
		"finished":          "完成了",
		"frombug":           "转需求",
		"fromfeedback":      "由反馈创建",
		"fromlib":           "从用例库导入",
		"gitcommited":       "提交代码",
		"hidden":            "隐藏了",
		"holiday":           "放假安排|holiday|browse|",
		"host":              "主机|host|view|id=%s",
		"internalaudit":     "提交内审",
		"leave": map[string]string{
			"commited": "请假|leave|browsereview|",
			"created":  "请假|leave|browsereview|",
			"reported": "销假|leave|browsereview|",
			"reviewed": "请假|leave|personal|",
			"revoked":  "请假|leave|browsereview|",
		},
		"lieu": map[string]string{
			"commited": "调休|lieu|browsereview|",
			"created":  "调休|lieu|browsereview|",
			"reviewed": "调休|lieu|personal|",
			"revoked":  "调休|lieu|browsereview|",
		},
		"linkcases":        "用例关联到",
		"linkcomponent":    "组件关联到",
		"linked2bug":       "关联了",
		"linked2build":     "关联了",
		"linked2plan":      "关联计划",
		"linked2project":   "关联" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"linked2release":   "关联发布",
		"linkhost":         "主机关联到",
		"linkrelatedbug":   "关联了相关Bug",
		"linkrelatedcase":  "关联了相关用例",
		"linkrelatedstory": "关联了相关需求",
		"linkservice":      "服务关联到",
		"login":            "登录系统",
		"logout":           "退出登录",
		"makeup": map[string]string{
			"commited": "补班|makeup|browsereview|",
			"created":  "补班|makeup|browsereview|",
			"reviewed": "补班|makeup|personal|",
			"revoked":  "补班|makeup|browsereview|",
		},
		"marked":  "编辑了",
		"moved":   "移动了",
		"offline": "下架了",
		"online":  "上架了",
		"opened":  "创建",
		"overtime": map[string]string{
			"commited": "加班|overtime|browsereview|",
			"created":  "加班|overtime|browsereview|",
			"reviewed": "加班|overtime|personal|",
			"revoked":  "加班|overtime|browsereview|",
		},
		"paused":         "暂停了",
		"product":        Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "|product|view|productID=%s",
		"productplan":    "计划|productplan|view|productID=%s",
		"project":        Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "|project|view|projectID=%s",
		"proofreading":   "对单完成",
		"recordEstimate": "记录了工时",
		"release":        "发布|release|view|productID=%s",
		"replied":        "回复了",
		"repocreated":    "创建评审",
		"reported":       "销假了",
		"resolved":       "解决了",
		"restarted":      "继续了",
		"reviewed":       "评审了",
		"revoked":        "撤销了",
		"serverroom":     "机房|serverroom|browse|",
		"service":        "服务|service|view|id=%s",
		"space":          " ",
		"started":        "开始了",
		"story":          "需求|story|view|storyID=%s",
		"subdividestory": "细分了需求",
		"suspended":      "挂起",
		"svncommited":    "提交代码",
		"task":           "任务|task|view|taskID=%s",
		"testreport":     "报告|testreport|view|report=%s",
		"testsuite":      "测试套件|testsuite|view|suiteID=%s",
		"testtask":       "测试单|testtask|view|caseID=%s",
		"tobug":          "转bug",
		"todo":           "待办|todo|view|todoID=%s",
		"tostory":        "转需求",
		"totask":         "转任务",
		"totodo":         "转待办",
		"trip": map[string]string{
			"commited": "外出|trip|browsereview|",
			"created":  "外出|trip|browsereview|",
			"reviewed": "外出|trip|personal|",
			"revoked":  "外出|trip|browsereview|",
		},
		"undeleted":           "还原了",
		"unlinkchildstory":    "移除了细分需求",
		"unlinkedfrombuild":   "移除版本",
		"unlinkedfromplan":    "移除计划",
		"unlinkedfromproject": "移除" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"unlinkedfromrelease": "移除发布",
		"unlinkrelatedbug":    "移除了相关Bug",
		"unlinkrelatedcase":   "移除了相关用例",
		"unlinkrelatedstory":  "移除了相关需求",
		"user":                "用户|user|view|account=%s",
		"verified":            "验收了",
		"webhook":             "Webhook|webhook|browse|",
	}
	Lang[protocol.ZH_CN]["action"]["search"] = map[string]interface{}{
		"label": map[string]string{
			"":                    "",
			"activated":           Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["activated"].(string),
			"assigned":            Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["assigned"].(string),
			"blocked":             Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["blocked"].(string),
			"bugconfirmed":        Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["bugconfirmed"].(string),
			"canceled":            Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["canceled"].(string),
			"changed":             Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["changed"].(string),
			"changestatus":        Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["changestatus"].(string),
			"closed":              Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["closed"].(string),
			"commented":           Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["commented"].(string),
			"confirmed":           Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["confirmed"].(string),
			"created":             Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["created"].(string),
			"deleted":             Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["deleted"].(string),
			"deletedfile":         Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["deletedfile"].(string),
			"edited":              Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["edited"].(string),
			"editestimate":        Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["editestimate"].(string),
			"editfile":            Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["editfile"].(string),
			"erased":              Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["erased"].(string),
			"finished":            Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["finished"].(string),
			"frombug":             Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["frombug"].(string),
			"gitcommited":         Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["gitcommited"].(string),
			"hidden":              Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["hidden"].(string),
			"linked2plan":         Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["linked2plan"].(string),
			"linked2project":      Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["linked2project"].(string),
			"login":               Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["login"].(string),
			"logout":              Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["logout"].(string),
			"marked":              Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["marked"].(string),
			"moved":               Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["moved"].(string),
			"opened":              Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["opened"].(string),
			"paused":              Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["paused"].(string),
			"recordEstimate":      Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["recordEstimate"].(string),
			"resolved":            Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["resolved"].(string),
			"restarted":           Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["restarted"].(string),
			"reviewed":            Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["reviewed"].(string),
			"started":             Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["started"].(string),
			"svncommited":         Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["svncommited"].(string),
			"tostory":             Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["tostory"].(string),
			"totask":              Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["totask"].(string),
			"undeleted":           Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["undeleted"].(string),
			"unlinkedfromplan":    Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["unlinkedfromplan"].(string),
			"unlinkedfromproject": Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["unlinkedfromproject"].(string),
			"verified":            Lang[protocol.ZH_CN]["action"]["label"].(map[string]interface{})["verified"].(string),
		},
		"objectTypeList": map[string]string{
			"":            "",
			"branch":      "分支",
			"bug":         "Bug",
			"build":       "版本",
			"case":        "用例",
			"caselib":     "公共库",
			"caseresult":  "用例结果",
			"doc":         "文档",
			"doclib":      "文档库",
			"productplan": "计划",
			"release":     "发布",
			"stepresult":  "用例步骤",
			"story":       "需求",
			"task":        "任务",
			"testreport":  "报告",
			"testsuite":   "套件",
			"testtask":    "测试单",
			"todo":        "待办",
			"user":        "用户",
			"product":     Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
			"project":     Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		},
	}
	Lang[protocol.ZH_CN]["action"]["dynamic"] = map[string]string{
		"all":        "所有",
		"hidden":     "已隐藏",
		"lastMonth":  "上月",
		"lastWeek":   "上周",
		"search":     "搜索",
		"thisMonth":  "本月",
		"thisWeek":   "本周",
		"today":      "今天",
		"twoDaysAgo": "前天",
		"yesterday":  "昨天",
	}
	Lang[protocol.ZH_CN]["action"]["dynamicAction"] = map[string]interface{}{
		"attend": map[string]string{
			"commited": "提交考勤",
			"reviewed": "审核考勤",
		},
		"bug": map[string]string{
			"fromfeedback": "由反馈创建Bug",
			"repocreated":  "创建代码评审",
		},
		"deploy": map[string]string{
			"activated": "激活上线计划",
			"finished":  "完成上线计划",
			"commented": "上线计划备注",
			"created":   "创建上线计划",
			"deleted":   "删除上线计划",
			"edited":    "编辑上线计划",
			"hidden":    "隐藏上线计划",
			"undeleted": "还原上线计划",
		},
		"deploystep": map[string]string{
			"assigned":  "指派上线步骤",
			"finished":  "完成上线步骤",
			"created":   "创建上线步骤",
			"deleted":   "删除上线步骤",
			"edited":    "编辑上线步骤",
			"hidden":    "隐藏上线步骤",
			"undeleted": "还原上线步骤",
		},
		"effort": map[string]string{
			"created":   "创建日志",
			"deleted":   "删除日志",
			"edited":    "编辑日志",
			"hidden":    "隐藏日志",
			"undeleted": "还原日志",
		},
		"feedback": map[string]string{
			"asked":     "追问反馈",
			"assigned":  "指派反馈",
			"closed":    "关闭反馈",
			"commented": "评论反馈",
			"edited":    "编辑反馈",
			"opened":    "创建反馈",
			"replied":   "回复反馈",
			"reviewed":  "审核反馈",
			"tobug":     "反馈转Bug",
			"totask":    "反馈转任务",
			"totodo":    "反馈转待办",
			"tostory":   "反馈转" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string),
		},
		"holiday": map[string]string{
			"created": "创建节假日",
		},
		"host": map[string]string{
			"created":   "创建主机",
			"deleted":   "删除主机",
			"edited":    "编辑主机",
			"hidden":    "隐藏主机",
			"offline":   "下架主机",
			"online":    "上架主机",
			"undeleted": "还原主机",
		},
		"leave": map[string]string{
			"commited": "提交请假",
			"created":  "创建请假",
			"edited":   "编辑请假",
			"reported": "销假",
			"reviewed": "审核请假",
			"revoked":  "撤销请假",
		},
		"lieu": map[string]string{
			"commited": "提交调休",
			"created":  "创建调休",
			"edited":   "编辑调休",
			"reviewed": "审核调休",
			"revoked":  "撤销调休",
		},
		"makeup": map[string]string{
			"commited": "提交补班",
			"created":  "创建补班",
			"edited":   "编辑补班",
			"reviewed": "审核补班",
			"revoked":  "撤销补班",
		},
		"overtime": map[string]string{
			"commited": "提交加班",
			"created":  "创建加班",
			"edited":   "编辑加班",
			"reviewed": "审核加班",
			"revoked":  "撤销加班",
		},
		"serverroom": map[string]string{
			"created":   "创建机房",
			"deleted":   "删除机房",
			"edited":    "编辑机房",
			"hidden":    "隐藏机房",
			"undeleted": "还原机房",
		},
		"service": map[string]string{
			"created":   "创建服务",
			"deleted":   "删除服务",
			"edited":    "编辑服务",
			"hidden":    "隐藏服务",
			"undeleted": "还原服务",
		},
		"task": map[string]string{
			"fromfeedback": "由反馈创建任务",
			"gitcommited":  "git提交",
		},
		"todo": map[string]string{
			"fromfeedback": "由反馈创建待办",
		},
		"story": map[string]string{
			"fromfeedback": "由反馈创建" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string),
		},
	}
	Lang[protocol.ZH_CN]["action"]["history"] = map[string]string{
		"action": "关联日志",
		"diff":   "不同",
		"field":  "字段",
		"new":    "新值",
		"old":    "旧值",
	}
	Lang[protocol.ZH_CN]["action"]["objectTypes"] = map[string]string{
		"branch":      "分支",
		"bug":         "Bug",
		"build":       "版本",
		"case":        "用例",
		"caselib":     "用例库",
		"caseresult":  "用例结果",
		"deploy":      "上线计划",
		"deploystep":  "上线步骤",
		"doc":         "文档",
		"doclib":      "文档库",
		"effort":      "日志",
		"entry":       "应用",
		"feedback":    "反馈",
		"holiday":     "放假安排",
		"host":        "主机",
		"leave":       "请假",
		"lieu":        "调休",
		"makeup":      "补班",
		"module":      "模块",
		"overtime":    "加班",
		"product":     Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"productplan": "计划",
		"project":     Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"release":     "发布",
		"report":      "报表",
		"serverroom":  "机房",
		"service":     "服务",
		"stepresult":  "用例步骤",
		"story":       "需求",
		"task":        "任务",
		"testreport":  "报告",
		"testsuite":   "套件",
		"testtask":    "测试单",
		"todo":        "待办",
		"trip":        "外出",
		"user":        "用户",
		"webhook":     "Webhook",
	}
	Lang[protocol.ZH_CN]["admin"]["info"] = map[string]string{
		"version": "当前系统的版本是%s，",
		"links":   "您可以访问以下链接：",
		"log":     "超出存天数的日志会被删除，需要开启计划任务。",
	}
	Lang[protocol.ZH_CN]["admin"]["notice"] = map[string]string{
		"ignore": "不再提示",
		"int":    "『%s』应当是正整数。",
	}
	Lang[protocol.ZH_CN]["admin"]["register"] = map[string]string{
		"common":     "注册新帐号绑定",
		"click":      "点击此处",
		"lblAccount": "请设置您的用户名，英文字母和数字的组合，三位以上。",
		"lblPasswd":  "请设置您的密码。数字和字母的组合，六位以上。",
		"submit":     "登记",
		"bind":       "绑定已有帐号",
		"success":    "登记账户成功",
	}
	Lang[protocol.ZH_CN]["admin"]["bind"] = map[string]string{
		"caption": "关联社区帐号",
		"success": "关联账户成功",
	}
	Lang[protocol.ZH_CN]["admin"]["safe"] = map[string]interface{}{
		"common":                   "安全策略",
		"set":                      "密码安全设置",
		"password":                 "密码安全",
		"weak":                     "常用弱口令",
		"reason":                   "类型",
		"checkWeak":                "弱口令扫描",
		"changeWeak":               "修改弱口令密码",
		"modifyPasswordFirstLogin": "首次登录修改密码",
		"modeList":                 []string{"不检查", "中", "强"},
		"modeRuleList":             []string{"6位以上，包含大小写字母，数字。", "10位以上，包含字母，数字，特殊字符。"},
		"reasonList": map[string]string{
			"weak":     "常用弱口令",
			"account":  "与帐号相同",
			"mobile":   "与手机相同",
			"phone":    "与电话相同",
			"birthday": "与生日相同",
		},
		"modifyPasswordList": []string{"必须修改", "不强制"},
		"noticeMode":         "系统会在登录、创建和修改用户、修改密码的时候检查用户口令。",
		"noticeStrong":       "密码长度越长，含有大写字母或数字或特殊符号越多，密码字母越不重复，安全度越强！",
	}
	Lang[protocol.ZH_CN]["administration"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["administration"]["sign"] = "首页显示"
	Lang[protocol.ZH_CN]["administration"]["common"] = "行政管理"
	Lang[protocol.ZH_CN]["api"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["api"]["post"] = "POST方式调试请参照页面表单"
	Lang[protocol.ZH_CN]["api"]["noParam"] = "GET方式调试不需要输入参数，"
	Lang[protocol.ZH_CN]["api"]["data"] = "内容"
	Lang[protocol.ZH_CN]["api"]["status"] = "状态"
	Lang[protocol.ZH_CN]["api"]["result"] = "返回结果"
	Lang[protocol.ZH_CN]["api"]["url"] = "请求地址"
	Lang[protocol.ZH_CN]["api"]["submit"] = "提交"
	Lang[protocol.ZH_CN]["api"]["debug"] = "调试"
	Lang[protocol.ZH_CN]["api"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["api"]["startLine"] = "%s,%s行"
	Lang[protocol.ZH_CN]["api"]["position"] = "位置"
	Lang[protocol.ZH_CN]["api"]["sql"] = "SQL查询接口"
	Lang[protocol.ZH_CN]["api"]["getModel"] = "超级model调用接口"
	Lang[protocol.ZH_CN]["api"]["common"] = "API接口"
	Lang[protocol.ZH_CN]["api"]["error"] = map[string]string{
		"onlySelect": "SQL查询接口只允许SELECT查询",
	}
	Lang[protocol.ZH_CN]["attend"]["clientList"] = map[string]string{
		"all":     "所有",
		"desktop": "电脑版",
	}
	Lang[protocol.ZH_CN]["attend"]["statusList"] = map[string]string{
		"normal":     "正常",
		"late":       "迟到",
		"early":      "早退",
		"both":       "迟到+早退",
		"absent":     "旷工",
		"leave":      "请假",
		"makeup":     "补班",
		"overtime":   "加班",
		"lieu":       "调休",
		"trip":       "出差",
		"egress":     "外出",
		"rest":       "休息日",
		"halfAbsent": "旷工半天",
		"allAbsent":  "旷工一天",
	}
	Lang[protocol.ZH_CN]["attend"]["abbrStatusList"] = map[string]string{
		"normal":   "√",
		"late":     "迟",
		"early":    "早",
		"both":     "迟+早",
		"absent":   "旷",
		"leave":    "假",
		"makeup":   "补",
		"overtime": "加",
		"lieu":     "调",
		"trip":     "差",
		"egress":   "出",
		"rest":     "休",
	}
	Lang[protocol.ZH_CN]["attend"]["markStatusList"] = map[string]string{
		"normal":   "√",
		"late":     "=",
		"early":    ">",
		"both":     "=>",
		"absent":   "x",
		"leave":    "!",
		"makeup":   "↑",
		"overtime": "+",
		"lieu":     "↓",
		"trip":     "$",
		"egress":   "#",
		"rest":     "~",
	}
	Lang[protocol.ZH_CN]["attend"]["reasonList"] = map[string]string{
		"normal":   "准点上下班",
		"leave":    "请假",
		"makeup":   "补班",
		"overtime": "加班",
		"lieu":     "调休",
		"trip":     "出差",
		"egress":   "外出",
	}
	Lang[protocol.ZH_CN]["attend"]["reviewStatusList"] = map[string]string{
		"wait":   "待审核",
		"pass":   "通过",
		"reject": "拒绝",
	}
	Lang[protocol.ZH_CN]["attend"]["workingDaysList"] = []protocol.HtmlKeyValueStr{
		{"1", "周一"},
		{"2", "周二"},
		{"3", "周三"},
		{"4", "周四"},
		{"5", "周五"},
		{"6", "周六"},
		{"7", "周日"},
	}
	Lang[protocol.ZH_CN]["attend"]["mustSignOutList"] = map[string]string{
		"yes": "需要",
		"no":  "不需要",
	}
	Lang[protocol.ZH_CN]["attend"]["notice"] = map[string]string{
		"today":    "<p>您今天的考勤状态为：%s，<a href='%s' %s>去补录</a>。</p>",
		"yestoday": "<p>您昨天的考勤状态为：%s，<a href='%s' %s>去补录</a>。</p>",
		"absent":   "没有记录",
	}
	Lang[protocol.ZH_CN]["attend"]["confirmReview"] = map[string]string{
		"pass":   "您确定要执行通过操作吗？",
		"reject": "您确定要执行拒绝操作吗？",
	}
	Lang[protocol.ZH_CN]["attend"]["beginDate"] = map[string]string{
		"company":  "公司开始考勤日期",
		"personal": "个人开始考勤日期",
	}
	Lang[protocol.ZH_CN]["attend"]["note"] = map[string]string{
		"ip":           "允许签到的ip，多个ip用逗号隔开。支持IP段，如192.168.1.*",
		"allip":        "无限制",
		"IPDenied":     "签到IP受限，无法签到",
		"beginDate":    "设置开始考勤的日期，在该日期之前不记录考勤状态。如果不设置则根据实际数据记录考勤状态。默认使用公司开始考勤日期计算考勤状态，如果设置了个人开始考勤日期则使用个人日期。",
		"signInClient": "设置为所有时可以通过任意途径签到，否则只能通过选定的访问方式签到。",
	}
	Lang[protocol.ZH_CN]["attend"]["tripTypeList"] = map[string]string{
		"trip":   "出差",
		"egress": "外出",
	}
	Lang[protocol.ZH_CN]["backup"]["settingList"] = map[string]string{
		"nofile": "不备份附件和代码",
		"nozip":  "只拷贝文件，不压缩",
		"nosafe": "不需要防下载PHP文件头",
	}
	Lang[protocol.ZH_CN]["backup"]["success"] = map[string]string{
		"backup":  "备份成功！",
		"restore": "还原成功！",
	}
	Lang[protocol.ZH_CN]["backup"]["error"] = map[string]string{
		"noCreateDir": "备份目录不存在，也无法创建该目录",
		"noWritable":  "<code>%s</code> 不可写！请检查该目录权限，否则无法备份。",
		"noDelete":    "文件 %s 无法删除，修改权限或手工删除。",
		"restoreSQL":  "数据库还原失败，错误：%s",
		"restoreFile": "附件还原失败，错误：%s",
		"backupFile":  "附件备份失败，错误：%s",
		"backupCode":  "代码备份失败，错误：%s",
	}
	Lang[protocol.ZH_CN]["block"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["block"]["default"] = map[string][]protocol.HtmlBlock{
		"product": []protocol.HtmlBlock{
			protocol.HtmlBlock{
				Title: Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "统计",
				Block: `statistic`,
				Grid:  8,
				Params: map[string]string{
					`type`: `all`,
					`num`:  `20`,
				},
			},
			protocol.HtmlBlock{
				Title: Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `总览`,
				Block: `overview`,
				Grid:  4,
			},
			protocol.HtmlBlock{
				Title: `未关闭的` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
				Block: `list`,
				Grid:  8,
				Params: map[string]string{
					`num`:  `15`,
					`type`: `noclosed`,
				},
			},
			protocol.HtmlBlock{
				Title: `指派给我的需求`,
				Block: `story`,
				Grid:  4,
				Params: map[string]string{
					`num`:     `15`,
					`orderBy`: `id_desc`,
					`type`:    `assignedTo`,
				},
			},
		},
		"project": []protocol.HtmlBlock{
			protocol.HtmlBlock{
				Title: Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `统计`,
				Block: `statistic`,
				Grid:  8,
				Params: map[string]string{
					`type`: `all`,
					`num`:  `20`,
				},
			}, protocol.HtmlBlock{
				Title: Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `总览`,
				Block: `overview`,
				Grid:  4,
			}, protocol.HtmlBlock{
				Title: `进行中的` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
				Block: `list`,
				Grid:  8,
				Params: map[string]string{
					`num`:     `15`,
					`orderBy`: `id_desc`,
					`type`:    `undone`,
				},
			}, protocol.HtmlBlock{
				Title: `指派给我的任务`,
				Block: `task`,
				Grid:  4,
				Params: map[string]string{
					`num`:     `15`,
					`orderBy`: `id_desc`,
					`type`:    `assignedTo`,
				},
			},
		},
		"qa": []protocol.HtmlBlock{
			protocol.HtmlBlock{
				Title: `测试统计`,
				Block: `statistic`,
				Grid:  8,
				Params: map[string]string{
					`type`: `noclosed`,
					`num`:  `20`,
				},
			},
			protocol.HtmlBlock{
				Title: `指派给我的bug`,
				Block: `bug`,
				Grid:  4,
			},
			protocol.HtmlBlock{
				Title: `指派给我的用例`,
				Block: `case`,
				Grid:  4,
				Params: map[string]string{
					`num`:     `15`,
					`orderBy`: `id_desc`,
					`type`:    `assigntome`,
				},
			},
			protocol.HtmlBlock{
				Title: `待测版本列表`,
				Block: `testtask`,
				Grid:  4,
				Params: map[string]string{
					`num`:     `15`,
					`orderBy`: `id_desc`,
					`type`:    `wait`,
				},
			},
		},
		"my": []protocol.HtmlBlock{
			protocol.HtmlBlock{
				Title:  `上下班打卡`,
				Block:  `clockinout`,
				Grid:   8,
				Source: "oa",
			},
			protocol.HtmlBlock{
				Title: `欢迎`,
				Block: `welcome`,
				Grid:  8,
			},
			/*protocol.HtmlBlock{
				Title: `最新动态`,
				Block: `dynamic`,
				Grid:  4,
			},
			protocol.HtmlBlock{
				Title: `流程图`,
				Block: `flowchart`,
				Grid:  8,
			},*/
			protocol.HtmlBlock{
				Title:  `我的待办`,
				Block:  `list`,
				Grid:   4,
				Source: "todo",
				Params: map[string]string{
					`num`: `20`,
				},
			},
			protocol.HtmlBlock{
				Title:  Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `统计`,
				Block:  `statistic`,
				Grid:   8,
				Source: "project",
				Params: map[string]string{
					`type`: `all`,
					`num`:  `20`,
				},
			}, protocol.HtmlBlock{
				Title:  Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `总览`,
				Block:  `overview`,
				Grid:   4,
				Source: "project",
			},
			protocol.HtmlBlock{
				Title:  Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "统计",
				Block:  `statistic`,
				Grid:   8,
				Source: "product",
				Params: map[string]string{
					`type`: `all`,
					`num`:  `20`,
				},
			},
			protocol.HtmlBlock{
				Title:  Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `总览`,
				Block:  `overview`,
				Grid:   4,
				Source: "product",
			},
			protocol.HtmlBlock{
				Title: `指派给我的任务`,
				Block: `task`,
				Grid:  4,
				Params: map[string]string{
					`num`:     `15`,
					`orderBy`: `id_desc`,
					`type`:    `assignedTo`,
				},
			},
		},
	}
	Lang[protocol.ZH_CN]["block"]["flowchart"] = [][]string{
		[]string{`管理员`, `维护公司`, `添加用户`, `维护权限`},
		[]string{Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `经理`, `创建` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string), `维护模块`, `维护计划`, `维护需求`, `创建发布`},
		[]string{Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `经理`, `创建` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string), `维护团队`, `关联` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string), `关联需求`, `分解任务`},
		[]string{`研发人员`, `领取任务和Bug`, `更新状态`, `完成任务和Bug`},
		[]string{`测试人员`, `撰写用例`, `执行用例`, `提交Bug`, `验证Bug`, `关闭Bug`},
	}
	Lang[protocol.ZH_CN]["block"]["welcomeList"] = []protocol.HtmlKeyValueStr{
		{"19:00", "%s，晚上好！"},
		{"13:30", "%s，下午好！"},
		{"11:30", "%s，中午好！"},
		{"06:00", "%s，早上好！"},
	}
	Lang[protocol.ZH_CN]["block"]["welcome"] = "欢迎总览"
	Lang[protocol.ZH_CN]["block"]["type"] = "类型"
	Lang[protocol.ZH_CN]["block"]["todoNum"] = "待办数"
	Lang[protocol.ZH_CN]["block"]["title"] = "区块名称"
	Lang[protocol.ZH_CN]["block"]["taskNum"] = "任务数"
	Lang[protocol.ZH_CN]["block"]["style"] = "外观"
	Lang[protocol.ZH_CN]["block"]["source"] = "来源模块"
	Lang[protocol.ZH_CN]["block"]["role"] = "角色"
	Lang[protocol.ZH_CN]["block"]["reset"] = "恢复默认"
	Lang[protocol.ZH_CN]["block"]["remove"] = "移除"
	Lang[protocol.ZH_CN]["block"]["refresh"] = "刷新"
	Lang[protocol.ZH_CN]["block"]["ordersSaved"] = "排序已保存"
	Lang[protocol.ZH_CN]["block"]["orderBy"] = "排序"
	Lang[protocol.ZH_CN]["block"]["order"] = "排序"
	Lang[protocol.ZH_CN]["block"]["num"] = "数量"
	Lang[protocol.ZH_CN]["block"]["noticeNewBlock"] = "10.0版本以后各个视图主页提供了全新的视图，您要启用新的视图布局吗？"
	Lang[protocol.ZH_CN]["block"]["noData"] = "当前统计类型下暂无数据"
	Lang[protocol.ZH_CN]["block"]["name"] = "区块名称"
	Lang[protocol.ZH_CN]["block"]["myTask"] = "我的任务"
	Lang[protocol.ZH_CN]["block"]["myStory"] = "我的需求"
	Lang[protocol.ZH_CN]["block"]["myProject"] = `进行中的` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)
	Lang[protocol.ZH_CN]["block"]["myProduct"] = `未关闭的` + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["block"]["myBug"] = "我的BUG"
	Lang[protocol.ZH_CN]["block"]["module"] = "所属模块"
	Lang[protocol.ZH_CN]["block"]["leftToday"] = "今天剩余工作总计"
	Lang[protocol.ZH_CN]["block"]["lblTesttask"] = "查看测试详情"
	Lang[protocol.ZH_CN]["block"]["lblNum"] = "条数"
	Lang[protocol.ZH_CN]["block"]["lblModule"] = "模块"
	Lang[protocol.ZH_CN]["block"]["lblHtml"] = "HTML内容"
	Lang[protocol.ZH_CN]["block"]["lblFlowchart"] = "流程图"
	Lang[protocol.ZH_CN]["block"]["lblBlock"] = "区块"
	Lang[protocol.ZH_CN]["block"]["hidden"] = "隐藏"
	Lang[protocol.ZH_CN]["block"]["height"] = "高度"
	Lang[protocol.ZH_CN]["block"]["grid"] = "位置"
	Lang[protocol.ZH_CN]["block"]["emptyTip"] = "暂无信息"
	Lang[protocol.ZH_CN]["block"]["editBlock"] = "编辑区块"
	Lang[protocol.ZH_CN]["block"]["dynamicInfo"] = "<span class='timeline-tag'>%s</span> <span class='timeline-text'>%s <em>%s</em> %s <a href='%s' title='%s'>%s</a></span>"
	Lang[protocol.ZH_CN]["block"]["dynamic"] = "最新动态"
	Lang[protocol.ZH_CN]["block"]["delayed"] = "已延期"
	Lang[protocol.ZH_CN]["block"]["createBlock"] = "添加区块"
	Lang[protocol.ZH_CN]["block"]["confirmReset"] = "是否恢复默认布局"
	Lang[protocol.ZH_CN]["block"]["confirmRemoveBlock"] = "确定移除区块吗？"
	Lang[protocol.ZH_CN]["block"]["confirmClose"] = "确定永久关闭该区块吗？关闭后所有人都将无法使用该区块，可以在后台自定义中打开。"
	Lang[protocol.ZH_CN]["block"]["common"] = "区块"
	Lang[protocol.ZH_CN]["block"]["color"] = "颜色"
	Lang[protocol.ZH_CN]["block"]["closeForever"] = "永久关闭"
	Lang[protocol.ZH_CN]["block"]["bugNum"] = "Bug数"
	Lang[protocol.ZH_CN]["block"]["block"] = "来源区块"
	Lang[protocol.ZH_CN]["block"]["assignToMe"] = "指派给我"
	Lang[protocol.ZH_CN]["block"]["account"] = "所属用户"
	Lang[protocol.ZH_CN]["block"]["availableBlocks"] = map[string]string{
		"bug":      "我的Bug",
		"build":    "版本列表",
		"case":     "我的用例",
		"plan":     "计划列表",
		"product":  Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "列表",
		"project":  Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "列表",
		"release":  "发布列表",
		"story":    "我的需求",
		"task":     "我的任务",
		"testtask": "测试版本列表",
		"todo":     "我的待办",
	}
	Lang[protocol.ZH_CN]["block"]["gridOptions"] = []string{"右侧", `左侧`}
	Lang[protocol.ZH_CN]["block"]["moduleList"] = map[string]string{
		"attend":  "办公",
		"product": Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"project": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"qa":      "测试",
		"todo":    "待办",
	}
	Lang[protocol.ZH_CN]["block"]["modules"] = map[string]protocol.HtmlBlockModule{
		"attend": protocol.HtmlBlockModule{
			AvailableBlocks: map[string]string{
				"clockinout": "上下班打卡",
			},
		},
		"common": protocol.HtmlBlockModule{
			MoreLinkList: map[string]string{
				"dynamic": "company|dynamic|",
			},
		},
		"product": protocol.HtmlBlockModule{
			AvailableBlocks: map[string]string{
				"list":      Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "列表",
				"overview":  Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "总览",
				"plan":      "计划列表",
				"release":   "发布列表",
				"statistic": Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "统计",
				"story":     "需求列表",
			},
			MoreLinkList: map[string]string{
				"list":  "product|all|product=&line=0&status=%s",
				"story": "my|story|type=%s",
			},
		},
		"project": protocol.HtmlBlockModule{
			AvailableBlocks: map[string]string{
				"build":     "版本列表",
				"list":      Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "列表",
				"overview":  Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "总览",
				"statistic": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "统计",
				"task":      "任务列表",
			},
			MoreLinkList: map[string]string{
				"list": "project|all|status=%s&project=",
				"task": "my|task|type=%s",
			},
		},
		"qa": protocol.HtmlBlockModule{
			AvailableBlocks: map[string]string{
				"bug":       "Bug列表",
				"case":      "用例列表",
				"statistic": "测试统计",
				"testtask":  "版本列表",
			},
			MoreLinkList: map[string]string{
				"bug":      "my|bug|type=%s",
				"case":     "my|testcase|type=%s",
				"testtask": "testtask|browse|type=%s",
			},
		},
		"todo": protocol.HtmlBlockModule{
			AvailableBlocks: map[string]string{
				"list": "待办列表",
			},
			MoreLinkList: map[string]string{
				"list": "my|todo|type=all",
			},
		},
	}

	Lang[protocol.ZH_CN]["block"]["orderByList"] = map[string]interface{}{
		"bug": map[string]string{
			"id_asc":        "ID 递增",
			"status_desc":   "状态倒序",
			"status_asc":    "状态正序",
			"stage_desc":    "阶段倒序",
			"stage_asc":     "阶段正序",
			"pri_desc":      "优先级递减",
			"pri_asc":       "优先级递增",
			"id_desc":       "ID 递减",
			"severity_asc":  "级别递增",
			"severity_desc": "级别递减",
		},
		"case": map[string]string{
			"id_asc": "ID 递增",
		},
		"product": map[string]string{
			"id_asc": "ID 递增",
		},
		"project": map[string]string{
			"id_asc": "ID 递增",
		},
		"story": map[string]string{
			"id_asc": "ID 递增",
		},
		"task": map[string]string{
			"deadline_asc":  "截止日期递增",
			"deadline_desc": "截止日期递减",
			"estimate_asc":  "预计时间递增",
			"estimate_desc": "预计时间递减",
			"id_asc":        "ID 递增",
			"id_desc":       "ID 递减",
			"pri_asc":       "优先级递增",
			"pri_desc":      "优先级递减",
			"status_asc":    "状态正序",
			"status_desc":   "状态倒序",
		},
	}
	Lang[protocol.ZH_CN]["block"]["params"] = map[string]string{
		"name":  "参数名称",
		"value": "参数值",
	}
	Lang[protocol.ZH_CN]["block"]["typeList"] = map[string]interface{}{
		"bug": map[string]string{
			"assignedTo": "指派给我",
			"reviewedBy": "由我评审",
			"openedBy":   "由我创建",
			"closedBy":   "由我关闭",
			"resolvedBy": "由我解决",
		},
		"case": map[string]string{
			"assigntome": "指派给我",
			"openedbyme": "由我创建",
		},
		"product": map[string]string{
			"all":        "全部",
			"wait":       "待测版本",
			"done":       "已测版本",
			"doing":      "测试中版本",
			"blocked":    "阻塞版本",
			"openedBy":   "由我创建",
			"finishedBy": "由我完成",
			"closedBy":   "由我关闭",
			"canceledBy": "由我取消",
			"assignedTo": "指派给我",
			"undone":     "未完成",
			"involved":   "我参与的",
			"closed":     "已关闭",
			"noclosed":   "未关闭",
		},
		"project": map[string]string{
			"all": "全部",
		},
		"story": map[string]string{
			"assignedTo": "指派给我",
		},
		"task": map[string]string{
			"all": "全部",
		},
		"testtask": map[string]string{
			"all": "全部",
		},
	}
	Lang[protocol.ZH_CN]["bug"]["severityList"] = []string{"1", `2`, `3`, `4`}
	Lang[protocol.ZH_CN]["bug"]["priList"] = []string{"", `1`, `2`, `3`, `4`}
	Lang[protocol.ZH_CN]["bug"]["osList"] = map[string]string{
		"":        "",
		"all":     "全部",
		"windows": "Windows",
		"win8":    "Windows 8",
		"win7":    "Windows 7",
		"vista":   "Windows Vista",
		"winxp":   "Windows XP",
		"win2012": "Windows 2012",
		"win2008": "Windows 2008",
		"win2003": "Windows 2003",
		"win2000": "Windows 2000",
		"android": "Android",
		"ios":     "IOS",
		"wp8":     "WP8",
		"wp7":     "WP7",
		"symbian": "Symbian",
		"linux":   "Linux",
		"freebsd": "FreeBSD",
		"osx":     "OS X",
		"unix":    "Unix",
		"others":  "其他",
	}
	Lang[protocol.ZH_CN]["bug"]["browserList"] = map[string]string{
		"":         "",
		"all":      "全部",
		"ie":       "IE系列",
		"ie11":     "IE11",
		"ie10":     "IE10",
		"ie9":      "IE9",
		"ie8":      "IE8",
		"ie7":      "IE7",
		"ie6":      "IE6",
		"chrome":   "chrome",
		"firefox":  "firefox系列",
		"firefox4": "firefox4",
		"firefox3": "firefox3",
		"firefox2": "firefox2",
		"opera":    "opera系列",
		"oprea11":  "opera11",
		"oprea10":  "opera10",
		"opera9":   "opera9",
		"safari":   "safari",
		"maxthon":  "傲游",
		"uc":       "UC",
		"other":    "其他",
	}
	Lang[protocol.ZH_CN]["bug"]["typeList"] = map[string]string{
		"":             "",
		"codeerror":    "代码错误",
		"interface":    "界面优化",
		"config":       "配置相关",
		"install":      "安装部署",
		"security":     "安全相关",
		"performance":  "性能问题",
		"standard":     "标准规范",
		"automation":   "测试脚本",
		"others":       "其他",
		"designchange": "设计变更",
		"newfeature":   "新增需求",
		"designdefect": "设计缺陷",
		"trackthings":  "事务跟踪",
	}
	Lang[protocol.ZH_CN]["bug"]["statusList"] = map[string]string{
		"":         "",
		"active":   "激活",
		"resolved": "已解决",
		"closed":   "已关闭",
	}
	Lang[protocol.ZH_CN]["bug"]["confirmedList"] = []string{"已确认", `未确认`}
	Lang[protocol.ZH_CN]["bug"]["resolutionList"] = map[string]string{

		"":           "",
		"bydesign":   "设计如此",
		"duplicate":  "重复Bug",
		"external":   "外部原因",
		"fixed":      "已解决",
		"notrepro":   "无法重现",
		"postponed":  "延期处理",
		"willnotfix": "不予解决",
		"tostory":    "转为需求",
	}
	Lang[protocol.ZH_CN]["bug"]["report"] = map[string]interface{}{
		"common": "报表",
		"select": "请选择报表类型",
		"create": "生成报表",
		"charts": map[string]string{
			"bugsPerProject":        Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "Bug数量",
			"bugsPerBuild":          "版本Bug数量",
			"bugsPerModule":         "模块Bug数量",
			"openedBugsPerDay":      "每天新增Bug数",
			"resolvedBugsPerDay":    "每天解决Bug数",
			"closedBugsPerDay":      "每天关闭的Bug数",
			"openedBugsPerUser":     "每人提交的Bug数",
			"resolvedBugsPerUser":   "每人解决的Bug数",
			"closedBugsPerUser":     "每人关闭的Bug数",
			"bugsPerSeverity":       "Bug严重程度统计",
			"bugsPerResolution":     "Bug解决方案统计",
			"bugsPerStatus":         "Bug状态统计",
			"bugsPerActivatedCount": "Bug激活次数统计",
			"bugsPerPri":            "Bug优先级统计",
			"bugsPerType":           "Bug类型统计",
			"bugsPerAssignedTo":     "指派给统计",
		},
		"options": map[string]string{
			"type":   "pie",
			"width":  "500",
			"height": "140",
		},
		"bugsPerProject": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
			},
		},
		"bugsPerBuild": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "版本",
			},
		},
		"bugsPerModule": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "模块",
			},
		},
		"openedBugsPerDay": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "日期",
				"type":      "bar",
			},
			"type": "bar",
		},
		"resolvedBugsPerDay": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "日期",
				"type":      "bar",
			},
			"type": "bar",
		},
		"closedBugsPerDay": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "日期",
				"type":      "bar",
			},
			"type": "bar",
		},
		"openedBugsPerUser": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "用户",
			},
		},
		"resolvedBugsPerUser": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "用户",
			},
		},
		"closedBugsPerUser": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "用户",
			},
		},
		"bugsPerSeverity": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "严重程度",
			},
		},
		"bugsPerResolution": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "解决方案",
			},
		},
		"bugsPerStatus": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "状态",
			},
		},
		"bugsPerActivatedCount": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "激活次数",
			},
		},
		"bugsPerPri": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "优先级",
			},
		},
		"bugsPerType": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "类型",
			},
		},
		"bugsPerAssignedTo": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "指派给",
			},
		},
		"bugLiveDays": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "处理时间",
			},
		},
		"bugHistories": map[string]interface{}{
			"graph": map[string]string{
				"xAxisName": "处理步骤",
			},
		},
	}
	Lang[protocol.ZH_CN]["bug"]["action"] = map[string]interface{}{
		"resolved": map[string]string{
			`main`:  `$date, 由 <strong>$actor</strong> 解决，方案为 <strong>$extra</strong> $appendLink。`,
			`extra`: `resolutionList`,
		},
		"tostory": map[string]string{
			"main": "$date, 由 <strong>$actor</strong> 转为<strong>需求</strong>，编号为 <strong>$extra</strong>。",
		},
		"totask": map[string]string{
			"main": "$date, 由 <strong>$actor</strong> 导入为<strong>任务</strong>，编号为 <strong>$extra</strong>。",
		},
		"linked2plan": map[string]string{
			"main": "$date, 由 <strong>$actor</strong> 关联到计划 <strong>$extra</strong>。",
		},
		"unlinkedfromplan": map[string]string{
			"main": "$date, 由 <strong>$actor</strong> 从计划 <strong>$extra</strong> 移除。",
		},
		"linked2build": map[string]string{
			"main": "$date, 由 <strong>$actor</strong> 关联到版本 <strong>$extra</strong>。",
		},
		"unlinkedfrombuild": map[string]string{
			"main": "$date, 由 <strong>$actor</strong> 从版本 <strong>$extra</strong> 移除。",
		},
		"linked2release": map[string]string{
			"main": "$date, 由 <strong>$actor</strong> 关联到发布 <strong>$extra</strong>。",
		},
		"unlinkedfromrelease": map[string]string{
			"main": "$date, 由 <strong>$actor</strong> 从发布 <strong>$extra</strong> 移除。",
		},
		"linkrelatedbug": map[string]string{
			"main": "$date, 由 <strong>$actor</strong> 关联相关Bug <strong>$extra</strong>。",
		},
		"unlinkrelatedbug": map[string]string{
			"main": "$date, 由 <strong>$actor</strong> 移除相关Bug <strong>$extra</strong>。",
		},
	}
	Lang[protocol.ZH_CN]["bug"]["placeholder"] = map[string]string{
		"chooseBuilds": "选择相关版本...",
		"newBuildName": "新版本名称",
	}
	Lang[protocol.ZH_CN]["bug"]["featureBar"] = map[string]map[string]string{
		"browse": map[string]string{
			"all":          Lang[protocol.ZH_CN]["bug"]["allBugs"].(string),
			"unclosed":     Lang[protocol.ZH_CN]["bug"]["unclosed"].(string),
			"openedbyme":   Lang[protocol.ZH_CN]["bug"]["openedByMe"].(string),
			"assigntome":   Lang[protocol.ZH_CN]["bug"]["assignToMe"].(string),
			"resolvedbyme": Lang[protocol.ZH_CN]["bug"]["resolvedByMe"].(string),
			"toclosed":     Lang[protocol.ZH_CN]["bug"]["toClosed"].(string),
			"unresolved":   Lang[protocol.ZH_CN]["bug"]["unResolved"].(string),
			"more":         Lang[protocol.ZH_CN]["common"]["more"].(string),
		},
	}
	Lang[protocol.ZH_CN]["bug"]["moreSelects"] = map[string]string{
		"unconfirmed":   Lang[protocol.ZH_CN]["bug"]["unconfirmed"].(string),
		"assigntonull":  Lang[protocol.ZH_CN]["bug"]["assignToNull"].(string),
		"longlifebugs":  Lang[protocol.ZH_CN]["bug"]["longLifeBugs"].(string),
		"postponedbugs": Lang[protocol.ZH_CN]["bug"]["postponedBugs"].(string),
		"overduebugs":   Lang[protocol.ZH_CN]["bug"]["overdueBugs"].(string),
		"needconfirm":   Lang[protocol.ZH_CN]["bug"]["needConfirm"].(string),
	}
	Lang[protocol.ZH_CN]["build"]["placeholder"] = map[string]string{
		"scmPath":  " 软件源代码库，如Subversion、Git库地址",
		"filePath": " 该版本软件包下载存储地址",
	}
	Lang[protocol.ZH_CN]["build"]["action"] = map[string]string{
		"buildopened": `$date, 由 <strong>$actor</strong> 创建版本 <strong>$extra</strong>。` + "\n",
	}
	Lang[protocol.ZH_CN]["company"]["guestOptions"] = []string{"不允许", `允许`}
	Lang[protocol.ZH_CN]["company"]["effort"] = map[string]string{
		"selectDate":    "日期",
		"projectSelect": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"productSelect": Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"userSelect":    "用户",
		"view":          "查看",
		"common":        "组织日志",
		"timeStat":      "本页总消耗%01.1f小时",
	}
	Lang[protocol.ZH_CN]["convert"]["directionList"] = map[string]string{
		"bug":   "Bug",
		"task":  "任务",
		"story": "需求",
	}
	Lang[protocol.ZH_CN]["convert"]["sourceList"] = map[string]interface{}{
		"BugFree": map[string]string{
			`bugfree_1`: `1.x`,
			`bugfree_2`: `2.x`,
		},
		"Redmine": map[string]string{`Redmine_1.1`: `1.1`},
	}
	Lang[protocol.ZH_CN]["convert"]["bugfree"] = map[string]string{
		"users":    "用户",
		"projects": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"modules":  "模块",
		"bugs":     "Bug",
		"cases":    "测试用例",
		"results":  "测试结果",
		"actions":  "历史记录",
		"files":    "附件",
	}
	Lang[protocol.ZH_CN]["convert"]["redmine"] = map[string]string{
		"users":        "用户",
		"groups":       "用户分组",
		"products":     Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"projects":     Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"stories":      "需求",
		"tasks":        "任务",
		"bugs":         "Bug",
		"productPlans": Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "计划",
		"teams":        "团队",
		"releases":     "发布",
		"builds":       "Build",
		"docLibs":      "文档库",
		"docs":         "文档",
		"files":        "附件",
	}
	Lang[protocol.ZH_CN]["convert"]["statusType"] = map[string]string{
		"bug":   "状态类型转换(Bug状态)",
		"story": "状态类型转换(Story状态)",
		"task":  "状态类型转换(Task状态)",
	}
	Lang[protocol.ZH_CN]["convert"]["priType"] = map[string]string{
		"bug":   "优先级类型转换(Bug状态)",
		"story": "优先级类型转换(Story状态)",
		"task":  "优先级类型转换(Task状态)",
	}
	Lang[protocol.ZH_CN]["convert"]["issue"] = map[string]string{
		"redmine": "Redmine",
		"goto":    "转换为",
	}
	Lang[protocol.ZH_CN]["cron"]["turnonList"] = []string{"打开", `关闭`}
	Lang[protocol.ZH_CN]["cron"]["statusList"] = map[string]string{
		"normal":  "正常",
		"running": "运行中",
		"stop":    "停止",
	}
	Lang[protocol.ZH_CN]["cron"]["typeList"] = map[string]string{
		"pm":     "http自调用",
		"system": "操作系统命令",
	}
	Lang[protocol.ZH_CN]["cron"]["toggleList"] = map[string]string{
		"start": "激活",
		"stop":  "禁用",
	}
	Lang[protocol.ZH_CN]["cron"]["notice"] = map[string]string{
		"m":   `取值范围:0-59，"*"代表取值范围内的数字，"/"代表"每"， "-"代表数字范围。`,
		"h":   "取值范围:0-23",
		"dom": "取值范围:1-31",
		"mon": "取值范围:1-12",
		"dow": "取值范围:0-6",
		"help": `注：如果服务器重启，或者发现计划任务没有正常工作，那么计划任务已经停止工作。需要手动点击【重启】按钮，或者一分钟后刷新页面，来开启计划任务。如果任务列表中第一条记录的最后执行时间改变，说明任务开启成功。",
		"errorRule":""%s" 填写的不是合法的值`,
	}
	Lang[protocol.ZH_CN]["custom"]["object"] = map[string]string{
		"story":    "需求",
		"task":     "任务",
		"bug":      "Bug",
		"testcase": "用例",
		"testtask": "版本",
		"todo":     "待办",
		"user":     "用户",
		"block":    "区块",
	}
	Lang[protocol.ZH_CN]["custom"]["story"] = map[string]interface{}{
		"fields": map[string]string{
			"priList":          "优先级",
			"sourceList":       "来源",
			"reasonList":       "关闭原因",
			"stageList":        "阶段",
			"statusList":       "状态",
			"reviewResultList": "评审结果",
			"review":           "评审流程",
		},
	}
	Lang[protocol.ZH_CN]["custom"]["task"] = map[string]interface{}{
		"fields": map[string]string{
			"priList":    "优先级",
			"typeList":   "类型",
			"reasonList": "关闭原因",
			"statusList": "状态",
			"hours":      "工时",
		},
	}
	Lang[protocol.ZH_CN]["custom"]["bug"] = map[string]interface{}{
		"fields": map[string]string{
			"priList":        "优先级",
			"severityList":   "严重程度",
			"osList":         "操作系统",
			"browserList":    "浏览器",
			"typeList":       "类型",
			"resolutionList": "解决方案",
			"statusList":     "状态",
			"longlife":       "久未处理天数",
		},
	}
	Lang[protocol.ZH_CN]["custom"]["testcase"] = map[string]interface{}{
		"fields": map[string]string{
			"priList":    "优先级",
			"typeList":   "类型",
			"stageList":  "阶段",
			"resultList": "执行结果",
			"statusList": "状态",
			"review":     "评审流程",
		},
	}
	Lang[protocol.ZH_CN]["custom"]["testtask"] = map[string]interface{}{
		"fields": map[string]string{
			"priList":    "优先级",
			"statusList": "状态",
		},
	}
	Lang[protocol.ZH_CN]["custom"]["todo"] = map[string]interface{}{
		"fields": map[string]string{
			"priList":    "优先级",
			"typeList":   "类型",
			"statusList": "状态",
		},
	}
	Lang[protocol.ZH_CN]["custom"]["user"] = map[string]interface{}{
		"fields": map[string]string{
			"roleList":   "职位",
			"statusList": "状态",
			"deleted":    "列出已删除用户",
		},
	}
	Lang[protocol.ZH_CN]["custom"]["block"] = map[string]interface{}{
		"fields": map[string]string{
			"closed": "关闭的区块",
		},
	}
	Lang[protocol.ZH_CN]["custom"]["notice"] = map[string]interface{}{
		"canNotAdd":        "该项参与运算，不提供自定义添加功能",
		"forceReview":      "指定人提交的%s必须评审。",
		"forceNotReview":   "指定人提交的%s不需要评审。",
		"longlife":         "Bug列表页面的久未处理标签中，列出设置天数之前未处理的Bug。",
		"invalidNumberKey": "键值应为不大于255的数字",
		"invalidStringKey": "键值应当为小写英文字母、数字或下划线的组合",
		"indexPage": map[string]string{
			"product": "从8.2版本起增加了产品主页视图，是否默认进入产品主页？",
			"project": "从8.2版本起增加了项目主页视图，是否默认进入项目主页？",
			"qa":      "从8.2版本起增加了测试主页视图，是否默认进入测试主页？",
		},
		"invalidStrlen": map[string]string{
			"ten":        "键的长度必须小于10个字符！",
			"twenty":     "键的长度必须小于20个字符！",
			"thirty":     "键的长度必须小于30个字符！",
			"twoHundred": "键的长度必须小于225个字符！",
		},
	}
	Lang[protocol.ZH_CN]["custom"]["reviewList"] = []string{"开启", `关闭`}
	Lang[protocol.ZH_CN]["custom"]["deletedList"] = []string{"列出", `不列出`}
	Lang[protocol.ZH_CN]["custom"]["weekendList"] = []string{"双休", `单休`}
	Lang[protocol.ZH_CN]["custom"]["productProject"] = map[string]interface{}{
		"relation": map[string]string{
			"0_0": "产品 - 项目",
			"0_1": "产品 - 迭代",
			"1_1": "项目 - 迭代",
		},
		"notice": "请根据实际情况选择适合自己团队的概念。",
	}
	Lang[protocol.ZH_CN]["custom"]["workingList"] = map[string]string{
		"full":      "完整研发管理工具",
		"onlyTest":  "测试管理工具",
		"onlyStory": "需求管理工具",
		"onlyTask":  "任务管理工具",
	}
	Lang[protocol.ZH_CN]["custom"]["scoreStatus"] = []string{"关闭", `开启`}
	Lang[protocol.ZH_CN]["custom"]["moduleName"] = map[string]string{
		"product":     Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"productplan": "计划",
		"project":     Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
	}
	Lang[protocol.ZH_CN]["datatable"] = make(map[string]interface{})

	Lang[protocol.ZH_CN]["datatable"]["showModule"] = "列表页是否显示模块名"
	Lang[protocol.ZH_CN]["datatable"]["platform"] = "平台"
	Lang[protocol.ZH_CN]["datatable"]["branch"] = "分支"
	Lang[protocol.ZH_CN]["datatable"]["resetGlobal"] = "全局恢复默认"
	Lang[protocol.ZH_CN]["datatable"]["setGlobal"] = "全局"
	Lang[protocol.ZH_CN]["datatable"]["confirmReset"] = "是否恢复默认设置？"
	Lang[protocol.ZH_CN]["datatable"]["required"] = "必选"
	Lang[protocol.ZH_CN]["datatable"]["switchToDatatable"] = "切换到高级表格"
	Lang[protocol.ZH_CN]["datatable"]["switchToTable"] = "切换到简单表格"
	Lang[protocol.ZH_CN]["datatable"]["customTip"] = "勾选需要显示的列，拖动列名进行排序。"
	Lang[protocol.ZH_CN]["datatable"]["custom"] = "自定义列"
	Lang[protocol.ZH_CN]["datatable"]["reset"] = "恢复默认"
	Lang[protocol.ZH_CN]["datatable"]["hide"] = "隐藏"
	Lang[protocol.ZH_CN]["datatable"]["show"] = "显示"
	Lang[protocol.ZH_CN]["datatable"]["width"] = "宽度"
	Lang[protocol.ZH_CN]["datatable"]["common"] = "数据表格"
	Lang[protocol.ZH_CN]["datatable"]["showModuleList"] = []protocol.HtmlKeyValueStr{
		{"[]", "不显示"},
		{"base", "只显示一级模块"},
		{"end", "只显示最后一级模块"},
	}

	Lang[protocol.ZH_CN]["dev"]["fields"] = map[string]string{
		"id":     "序号",
		"name":   "字段",
		"desc":   "描述",
		"type":   "类型",
		"length": "长度",
		"null":   "是否可空",
	}
	Lang[protocol.ZH_CN]["dev"]["tableList"] = map[string]string{
		"action":           "系统日志",
		"bug":              "Bug",
		"build":            "版本",
		"burn":             "燃尽图",
		"case":             "测试用例",
		"casestep":         "用例步骤",
		"company":          "公司",
		"config":           "配置",
		"custom":           "自定义",
		"dept":             "部门",
		"doc":              "文档",
		"doclib":           "文档库",
		"effort":           "日志",
		"extension":        "插件",
		"file":             "附件",
		"group":            "用户组",
		"grouppriv":        "组权限",
		"history":          "操作历史",
		"lang":             "语言定义",
		"module":           "模块",
		"product":          Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"productplan":      Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "计划",
		"project":          Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"projectproduct":   Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"projectstory":     Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string),
		"story":            Lang[protocol.ZH_CN]["common"]["storyCommon"].(string),
		"release":          "发布",
		"storyspec":        "{$lang->common->storyCommon}描述",
		"task":             "任务",
		"taskestimate":     "任务预计",
		"team":             "团队",
		"testresult":       "测试结果",
		"testrun":          "测试执行",
		"testtask":         "测试版本",
		"todo":             "待办",
		"user":             "用户",
		"usercontact":      "用户联系人",
		"usergroup":        "用户组",
		"userquery":        "用户查询",
		"usertpl":          "用户模板",
		"admin":            "后台管理",
		"api":              "API接口",
		"backup":           "备份",
		"common":           "公有模块",
		"convert":          "导入",
		"dev":              "二次开发",
		"git":              "GIT",
		"index":            "首页",
		"install":          "安装",
		"mail":             "邮箱",
		"misc":             "杂项",
		"my":               "我的地盘",
		"qa":               "测试",
		"report":           "统计",
		"search":           "搜索",
		"sso":              "单点登录",
		"svn":              "SVN",
		"testcase":         "测试用例",
		"testreport":       "测试报告",
		"testsuite":        "测试套件",
		"tree":             "模块关系",
		"upgrade":          "更新",
		"cron":             "定时任务",
		"datatable":        "数据表格",
		"block":            "区块",
		"branch":           "平台/分支",
		"doccontent":       "文档内容",
		"storystage":       "{$lang->common->storyCommon}阶段",
		"tutorial":         "新手教程",
		"suitecase":        "套件用例",
		"score":            "积分",
		"entry":            "应用",
		"webhook":          "WebHook",
		"log":              "接口日志",
		"message":          "消息",
		"notify":           "通知",
		"userview":         "可访问权限",
		"im_chatuser":      "客户端用户",
		"im_message":       "客户端消息",
		"im_messagestatus": "客户端状态",
		"repo":             "代码",
		"repohistory":      "版本历史",
		"repofiles":        "代码文件",
		"repobranch":       "代码分支",
		"faq":              "Faq",
		"feedbackview":     "反馈视图",
		"feedback":         "反馈",
		"feedbackproduct":  "反馈权限",
		"relationoftasks":  "任务关系",
		"ldap":             "LDAP",
		"makeup":           "补班",
		"attend":           "考勤",
		"holiday":          "节假日",
		"leave":            "请假",
		"lieu":             "调休",
		"overtime":         "加班",
		"attendstat":       "考勤统计",
		"trip":             "出差",
		"searchdict":       "搜索字典",
		"searchindex":      "搜索索引",
		"sms":              "短信配置",
		"ops":              "运维",
		"host":             "主机",
		"asset":            "资产",
		"serverroom":       "机房",
		"service":          "服务",
		"deploy":           "上线",
		"deploystep":       "上线步骤",
		"deployproduct":    "上线产品",
		"deployscope":      "上线范围",
	}
	Lang[protocol.ZH_CN]["dev"]["groupList"] = map[string]string{
		"my":       "我的地盘",
		"product":  Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"project":  Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"qa":       "测试",
		"doc":      "文档",
		"report":   "统计",
		"company":  "组织",
		"repo":     "代码",
		"api":      "API",
		"message":  "消息",
		"feedback": "反馈",
		"oa":       "办公",
		"search":   "搜索",
		"ops":      "运维",
	}
	Lang[protocol.ZH_CN]["dev"]["endGroupList"] = map[string]string{
		"admin":  "后台",
		"system": "系统",
		"other":  "其他",
	}
	Lang[protocol.ZH_CN]["doc"]["libTypeList"] = map[string]string{
		"product": Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "文档库",
		"project": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "文档库",
		"custom":  "自定义文档库",
	}
	Lang[protocol.ZH_CN]["doc"]["libIconList"] = map[string]string{
		"product": "icon-cube",
		"project": "icon-stack",
		"custom":  "icon-folder-o",
	}
	Lang[protocol.ZH_CN]["doc"]["systemLibs"] = map[string]string{
		"product": Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"project": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
	}
	Lang[protocol.ZH_CN]["doc"]["aclList"] = map[string]string{
		"open":    "公开",
		"custom":  "自定义",
		"private": "私有",
	}
	Lang[protocol.ZH_CN]["doc"]["types"] = map[string]string{
		"text": "文档",
		"url":  "链接",
	}
	Lang[protocol.ZH_CN]["doc"]["contentTypeList"] = map[string]string{
		"html":     "HTML",
		"markdown": "MarkDown",
	}
	Lang[protocol.ZH_CN]["doc"]["browseTypeList"] = map[string]string{
		"list": "列表",
		"grid": "目录",
	}
	Lang[protocol.ZH_CN]["doc"]["fastMenuList"] = map[string]string{
		"byediteddate":  "最近更新",
		"openedbyme":    "我的文档",
		"collectedbyme": "我的收藏",
	}
	Lang[protocol.ZH_CN]["doc"]["fastMenuIconList"] = map[string]string{
		"byediteddate":  "icon-folder-upload",
		"openedbyme":    "icon-folder-account",
		"collectedbyme": "icon-folder-star",
	}
	Lang[protocol.ZH_CN]["doc"]["customObjectLibs"] = map[string]string{
		"files":       "显示附件库",
		"customFiles": "显示自定义文档库",
	}
	Lang[protocol.ZH_CN]["doc"]["customShowLibsList"] = map[string]string{
		"zero":     "显示空文档的库",
		"unclosed": "只显示未关闭的项目",
	}
	Lang[protocol.ZH_CN]["doc"]["placeholder"] = map[string]string{
		"url": "相应的链接地址",
	}
	Lang[protocol.ZH_CN]["doclib"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["doclib"]["product"] = Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `库`
	Lang[protocol.ZH_CN]["doclib"]["project"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `库`
	Lang[protocol.ZH_CN]["doclib"]["select"] = "选择文档库"
	Lang[protocol.ZH_CN]["doclib"]["all"] = "所有文档库"
	Lang[protocol.ZH_CN]["doclib"]["files"] = "附件库"
	Lang[protocol.ZH_CN]["doclib"]["user"] = "用户"
	Lang[protocol.ZH_CN]["doclib"]["group"] = "分组"
	Lang[protocol.ZH_CN]["doclib"]["control"] = "访问控制"
	Lang[protocol.ZH_CN]["doclib"]["name"] = "文档库名称"
	Lang[protocol.ZH_CN]["doclib"]["main"] = map[string]string{
		"product": Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "主库",
		"project": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "主库",
	}
	Lang[protocol.ZH_CN]["doclib"]["tabList"] = map[string]string{
		"product": Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"project": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"custom":  "自定义",
	}
	Lang[protocol.ZH_CN]["doclib"]["nameList"] = map[string]string{
		"custom": "自定义文档库名称",
	}
	Lang[protocol.ZH_CN]["editor"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["editor"]["translate['model.php']"] = "model"
	Lang[protocol.ZH_CN]["editor"]["translate['control.php']"] = "control"
	Lang[protocol.ZH_CN]["editor"]["translate['config.php']"] = "config"
	Lang[protocol.ZH_CN]["editor"]["emptyFileName"] = "请写入一个文件名！"
	Lang[protocol.ZH_CN]["editor"]["notDelete"] = "无法删除，请检查权限！"
	Lang[protocol.ZH_CN]["editor"]["notWritable"] = "无法写入，可能没有权限。请尝试执行 chmod 777 -R "
	Lang[protocol.ZH_CN]["editor"]["repeatPage"] = "已经有此页面。是否覆盖？"
	Lang[protocol.ZH_CN]["editor"]["repeatFile"] = "文件名重复"
	Lang[protocol.ZH_CN]["editor"]["extendConfirm"] = "是否要重用原来代码？"
	Lang[protocol.ZH_CN]["editor"]["deleteConfirm"] = "是否要删除？"
	Lang[protocol.ZH_CN]["editor"]["pageName"] = "页面名称："
	Lang[protocol.ZH_CN]["editor"]["examplePHP"] = "(例如：***.php)"
	Lang[protocol.ZH_CN]["editor"]["exampleCss"] = "(例如：***.css)"
	Lang[protocol.ZH_CN]["editor"]["exampleJs"] = "(例如：***.js)"
	Lang[protocol.ZH_CN]["editor"]["exampleHook"] = "(例如：***.html.hook.php)"
	Lang[protocol.ZH_CN]["editor"]["isOverride"] = "覆盖重复文件"
	Lang[protocol.ZH_CN]["editor"]["fileName"] = "文件名："
	Lang[protocol.ZH_CN]["editor"]["sourceFile"] = "源文件："
	Lang[protocol.ZH_CN]["editor"]["filePath"] = "扩展："
	Lang[protocol.ZH_CN]["editor"]["moduleList"] = "模块列表"
	Lang[protocol.ZH_CN]["editor"]["delete"] = "删除页面"
	Lang[protocol.ZH_CN]["editor"]["save"] = "保存页面"
	Lang[protocol.ZH_CN]["editor"]["edit"] = "编辑扩展"
	Lang[protocol.ZH_CN]["editor"]["override"] = "覆盖"
	Lang[protocol.ZH_CN]["editor"]["newPage"] = "新增页面"
	Lang[protocol.ZH_CN]["editor"]["newExtend"] = "新增扩展"
	Lang[protocol.ZH_CN]["editor"]["newHook"] = "新增钩子"
	Lang[protocol.ZH_CN]["editor"]["newConfig"] = "新增配置"
	Lang[protocol.ZH_CN]["editor"]["newLang"] = "新增语言"
	Lang[protocol.ZH_CN]["editor"]["extend"] = "扩展"
	Lang[protocol.ZH_CN]["editor"]["newMethod"] = "新增方法"
	Lang[protocol.ZH_CN]["editor"]["index"] = "首页"
	Lang[protocol.ZH_CN]["editor"]["api"] = `API`
	Lang[protocol.ZH_CN]["editor"]["common"] = "编辑器"
	Lang[protocol.ZH_CN]["editor"]["modules"] = map[string]string{
		"action":      "系统日志",
		"admin":       "后台管理",
		"api":         "API接口",
		"bug":         "Bug管理",
		"build":       "Build",
		"common":      "公有模块",
		"company":     "组织视图",
		"convert":     "从其他系统导入",
		"dept":        "部门结构",
		"doc":         "文档视图",
		"extension":   "插件管理",
		"file":        "附件",
		"group":       "权限分组",
		"index":       "首页",
		"install":     "安装",
		"misc":        "杂项",
		"my":          "我的地盘",
		"product":     Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "视图",
		"productplan": "计划",
		"project":     Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "视图",
		"qa":          "测试视图",
		"release":     "发布",
		"report":      "报表",
		"search":      "搜索",
		"story":       "需求",
		"task":        "任务",
		"testcase":    "用例管理",
		"testtask":    "测试任务",
		"todo":        "待办",
		"tree":        "模块维护",
		"upgrade":     "升级",
		"user":        "用户",
	}
	Lang[protocol.ZH_CN]["entry"]["note"] = map[string]string{
		"name":    "授权应用名称",
		"code":    "授权应用代号，必须为字母或数字的组合",
		"ip":      "允许访问API的应用ip，多个ip用逗号隔开。支持IP段，如192.168.1.*",
		"allIP":   "无限制",
		"account": "授权应用账号",
	}
	Lang[protocol.ZH_CN]["entry"]["freePasswdList"] = []string{"开启", `关闭`}
	Lang[protocol.ZH_CN]["entry"]["errmsg"] = map[string]string{
		"PARAM_CODE_MISSING":    "缺少code参数",
		"PARAM_TOKEN_MISSING":   "缺少token参数",
		"SESSION_CODE_MISSING":  "缺少session code",
		"EMPTY_KEY":             "应用未设置密钥",
		"INVALID_TOKEN":         "无效的token参数",
		"SESSION_VERIFY_FAILED": "session验证失败",
		"IP_DENIED":             "该IP被限制访问",
		"ACCOUNT_UNBOUND":       "未绑定用户",
		"INVALID_ACCOUNT":       "用户不存在",
		"EMPTY_ENTRY":           "应用不存在",
		"CALLED_TIME":           "Token已失效",
		"ERROR_TIMESTAMP":       "错误的时间戳。",
	}
	Lang[protocol.ZH_CN]["file"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["file"]["collaboraFail"] = "Collabora Online 访问失败，可以尝试重启，或检查配置问题，请确保 %s/hosting/discovery 访问正常。"
	Lang[protocol.ZH_CN]["file"]["officeBusy"] = "Office转换在使用中，请稍候使用。"
	Lang[protocol.ZH_CN]["file"]["onlySupportXLSX"] = "只支持XLSX格式导入，请转换XLSX格式再导入。"
	Lang[protocol.ZH_CN]["file"]["fileNotFound"] = "未找到该文件，可能物理文件已被删除！"
	Lang[protocol.ZH_CN]["file"]["errorExtract"] = "解压缩失败！可能文件已经损坏，或压缩包里含有非法上传文件。"
	Lang[protocol.ZH_CN]["file"]["errorSuffix"] = "压缩包格式错误，只能上传zip压缩包！"
	Lang[protocol.ZH_CN]["file"]["dangerFile"] = " 您选择的文件存在安全风险，系统将不予上传。"
	Lang[protocol.ZH_CN]["file"]["errorFileMove"] = " 文件上传失败，移动文件时出错"
	Lang[protocol.ZH_CN]["file"]["errorFileFormate"] = " 文件上传失败，文件格式不在规定范围内"
	Lang[protocol.ZH_CN]["file"]["errorFileUpload"] = " 文件上传失败，文件大小可能超出限制"
	Lang[protocol.ZH_CN]["file"]["errorFileSize"] = " 文件大小已经超过%s，可能不能成功上传！"
	Lang[protocol.ZH_CN]["file"]["confirmDelete"] = " 您确定删除该附件吗？"
	Lang[protocol.ZH_CN]["file"]["errorCanNotWrite"] = "<span class='text-red'>文件夹 '%s' 不可写,请改变文件夹的权限。在linux中输入指令: <span class='code'>sudo chmod -R 777 %s</span></span>"
	Lang[protocol.ZH_CN]["file"]["errorNotExists"] = "<span class='text-red'>文件夹 '%s' 不存在</span>"
	Lang[protocol.ZH_CN]["file"]["importSummary"] = "本次导入共有<strong id='allCount'>%s</strong>条记录，每页导入%s条，需要导入<strong id='times'>%s</strong>次"
	Lang[protocol.ZH_CN]["file"]["importPager"] = "共有<strong>%s</strong>条记录，当前第<strong>%s</strong>页，共有<strong>%s</strong>页"
	Lang[protocol.ZH_CN]["file"]["saveAndNext"] = "保存并跳转下一页"
	Lang[protocol.ZH_CN]["file"]["uploadImagesExplain"] = `注：请上传"jpg, jpeg, gif, png"格式的图片，程序会以文件名作为标题，以图片作为内容。` //"注：请上传"jpg, jpeg, gif, png"格式的图片，程序会以文件名作为标题，以图片作为内容。"替换`注：请上传"jpg, jpeg, gif, png"格式的图片，程序会以文件名作为标题，以图片作为内容。`
	Lang[protocol.ZH_CN]["file"]["childTaskTips"] = "任务名称前有'>'标记的为子任务"
	Lang[protocol.ZH_CN]["file"]["dragFile"] = "请拖拽文件到此处"
	Lang[protocol.ZH_CN]["file"]["extra"] = "备注"
	Lang[protocol.ZH_CN]["file"]["downloads"] = "下载次数"
	Lang[protocol.ZH_CN]["file"]["addedDate"] = "添加时间"
	Lang[protocol.ZH_CN]["file"]["addedBy"] = "由谁添加"
	Lang[protocol.ZH_CN]["file"]["encoding"] = "编码"
	Lang[protocol.ZH_CN]["file"]["size"] = "大小"
	Lang[protocol.ZH_CN]["file"]["extension"] = "文件类型"
	Lang[protocol.ZH_CN]["file"]["untitled"] = "未命名"
	Lang[protocol.ZH_CN]["file"]["fileName"] = "文件名"
	Lang[protocol.ZH_CN]["file"]["title"] = "标题"
	Lang[protocol.ZH_CN]["file"]["pathname"] = "路径"
	Lang[protocol.ZH_CN]["file"]["uploadSuccess"] = "上传成功"
	Lang[protocol.ZH_CN]["file"]["beginUpload"] = "开始上传"
	Lang[protocol.ZH_CN]["file"]["addFile"] = "添加文件"
	Lang[protocol.ZH_CN]["file"]["preview"] = "预览"
	Lang[protocol.ZH_CN]["file"]["setExportTPL"] = "设置"
	Lang[protocol.ZH_CN]["file"]["defaultTPL"] = "默认模板"
	Lang[protocol.ZH_CN]["file"]["exportRange"] = "要导出的数据"
	Lang[protocol.ZH_CN]["file"]["exportFields"] = "要导出字段"
	Lang[protocol.ZH_CN]["file"]["setPublic"] = "设置公共模板"
	Lang[protocol.ZH_CN]["file"]["tplTitleAB"] = "模板名称"
	Lang[protocol.ZH_CN]["file"]["tplTitle"] = "模板名称"
	Lang[protocol.ZH_CN]["file"]["applyTemplate"] = "应用模板"
	Lang[protocol.ZH_CN]["file"]["maxUploadSize"] = "（不超过%s）"
	Lang[protocol.ZH_CN]["file"]["label"] = "标题："
	Lang[protocol.ZH_CN]["file"]["delete"] = "删除附件"
	Lang[protocol.ZH_CN]["file"]["inputFileName"] = "请输入附件名称"
	Lang[protocol.ZH_CN]["file"]["edit"] = "重命名"
	Lang[protocol.ZH_CN]["file"]["uploadDate"] = "上传时间："
	Lang[protocol.ZH_CN]["file"]["download"] = "下载附件"
	Lang[protocol.ZH_CN]["file"]["uploadImages"] = "多图上传"
	Lang[protocol.ZH_CN]["file"]["common"] = "附件"
	Lang[protocol.ZH_CN]["file"]["imguploadFail"] = "图片上传失败，建议刷新重试,错误%v"
	Lang[protocol.ZH_CN]["file"]["processFile"] = []protocol.HtmlKeyValueStr{
		{"processFile", "过程文件"},
	}
	Lang[protocol.ZH_CN]["file"]["typeChoices"] = []protocol.HtmlKeyValueStr{

		{"processFile", "过程文件"},
		{"sourceFile", "源文件"},
		{"feedbackFile", "反馈文件"},
	}
	Lang[protocol.ZH_CN]["file"]["FinalFile"] = []protocol.HtmlKeyValueStr{
		{"FinalFile", "最终文件"},
	}
	Lang[protocol.ZH_CN]["file"]["feedbackFile"] = []protocol.HtmlKeyValueStr{
		{"feedbackFile", "反馈文件"},
	}
	Lang[protocol.ZH_CN]["file"]["specFile"] = []protocol.HtmlKeyValueStr{
		{"modelFile", "模型规范文件"},
		{"animationFile", "动画规范文件"},
		{"specialFile", "特效规范文件"},
		{"bindingFile", "绑定规范文件"},
		{"twodFile", "2D规范文件"},
		{"otherFile", "其他规范文件"},
	}
	Lang[protocol.ZH_CN]["file"]["typeTOName"] = []protocol.HtmlKeyValueStr{
		{"processFile", "过程文件"},
		{"FinalFile", "最终文件"},
	}
	Lang[protocol.ZH_CN]["excel"]["title"] = []protocol.HtmlKeyValueStr{
		{"testcase", "用例"},
		{"bug", "Bug"},
		{"task", "任务"},
		{"story", Lang[protocol.ZH_CN]["common"]["storyCommon"].(string)},
		{"caselib", "用例库"},
		{"sysValue", "系统数据"},
		{"tree", "树状图"},
		{"feedback", "反馈"},
	}
	Lang[protocol.ZH_CN]["excel"]["help"] = []protocol.HtmlKeyValueStr{
		{"testcase", "添加用例时，每个用例步骤在新行用数字 + ‘.’来标记。同样的，预期也是用数字 + ‘.’与步骤对应。“用例标题”和“用例类型”是必填字段，如果不填导入时会忽略该条数据。"},
		{"bug", "添加Bug时，“标题”是必填字段，如果不填导入时会忽略该条数据。"},
		{"task", "添加任务时，“任务名称”和“任务类型”是必填字段，如果不填导入时会忽略该条数据；如需添加子任务，请在任务名称前用“>”标记。"},
		{"multiple", "如需添加多人任务，请在“最初预计”列里面，按照“用户名:年-月-日”格式添加，多个用户之间用换行分隔。用户名在“系统数据”工作表的F列查看。"},
		{"story", "添加" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "时，“" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "名称”是必填字段，如果不填导入时会忽略该条数据。"},
	}
	Lang[protocol.ZH_CN]["word"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["word"]["more"] = "更多请点击"
	Lang[protocol.ZH_CN]["word"]["headNotice"] = "该文件由系统自动导出"
	Lang[protocol.ZH_CN]["word"]["fileField"] = "附件"
	Lang[protocol.ZH_CN]["word"]["notice"] = map[string]string{
		"noexport": "目前没有该模块的导出功能。",
	}
	Lang[protocol.ZH_CN]["word"]["kanbanColor"] = map[string]string{
		"wait":   "#EBEDF7",
		"doing":  "#FBE7E9",
		"pause":  "#FFF4E5",
		"done":   "#E8F3DB",
		"cancel": "#EBEBEC",
		"closed": "#D7D7D7",
	}
	Lang[protocol.ZH_CN]["git"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["git"]["apiSync"] = "接口：同步git日志"
	Lang[protocol.ZH_CN]["git"]["diff"] = "比较源代码"
	Lang[protocol.ZH_CN]["git"]["cat"] = "查看源代码"
	Lang[protocol.ZH_CN]["git"]["common"] = "Git"
	Lang[protocol.ZH_CN]["group"]["copyOptions"] = map[string]string{
		"copyPriv": "复制权限",
		"copyUser": "复制用户",
	}
	Lang[protocol.ZH_CN]["holiday"]["typeList"] = map[string]string{
		"RE": "假期",
		"OT": "补班",
	}
	Lang[protocol.ZH_CN]["leave"]["typeList"] = map[string]string{
		"affairs":   "事假",
		"sick":      "病假",
		"annual":    "年假",
		"lieu":      "调休",
		"marry":     "婚/丧假",
		"maternity": "产假",
	}
	Lang[protocol.ZH_CN]["leave"]["statusList"] = map[string]string{
		"draft":   "草稿",
		"wait":    "等待审核",
		"doing":   "审核中",
		"pass":    "通过",
		"reject":  "拒绝",
		"back":    "撤销",
		"restore": "还原",
	}
	Lang[protocol.ZH_CN]["leave"]["confirmReview"] = map[string]string{
		"pass":   "您确定要执行通过操作吗？",
		"reject": "您确定要执行拒绝操作吗？",
		"back":   "您确定要取消此次请假吗？",
	}
	Lang[protocol.ZH_CN]["leave"]["reviewStatusList"] = map[string]string{
		"pass":   "通过",
		"reject": "拒绝",
	}
	Lang[protocol.ZH_CN]["leave"]["settings"] = map[string]string{
		"setReviewer":    "审批人|leave|setreviewer",
		"personalAnnual": "个人年假|leave|personalannual",
		"personallieu":   "个人加班调休|leave|personalieu",
	}
	Lang[protocol.ZH_CN]["leave"]["startfinishList"] = []string{"上午", `下午`}
	Lang[protocol.ZH_CN]["leave"]["leaderList"] = []string{"组长"}
	Lang[protocol.ZH_CN]["lieu"]["statusList"] = map[string]string{
		"draft":  "草稿",
		"wait":   "等待审核",
		"doing":  "审核中",
		"pass":   "通过",
		"reject": "拒绝",
	}
	Lang[protocol.ZH_CN]["lieu"]["confirmReview"] = map[string]string{
		"pass":   "您确定要执行通过操作吗？",
		"reject": "您确定要执行拒绝操作吗？",
	}
	Lang[protocol.ZH_CN]["lieu"]["checkHoursList"] = []string{"不检测调休时长", `调休时长不能超过加班时长`}
	Lang[protocol.ZH_CN]["lieu"]["reviewStatusList"] = map[string]string{
		"pass":   "通过",
		"reject": "拒绝",
	}
	Lang[protocol.ZH_CN]["lieu"]["startfinishList"] = []string{"上午", `下午`}
	Lang[protocol.ZH_CN]["mail"]["statusList"] = map[string]string{
		"sended": "成功",
		"fail":   "失败",
	}
	Lang[protocol.ZH_CN]["mail"]["turnonList"] = []string{"打开", `关闭`}
	Lang[protocol.ZH_CN]["mail"]["asyncList"] = []string{"是", `否`}
	Lang[protocol.ZH_CN]["mail"]["debugList"] = []string{"关闭", `一般`, `较高`}
	Lang[protocol.ZH_CN]["mail"]["authList"] = []string{"需要", `不需要`}
	Lang[protocol.ZH_CN]["mail"]["secureList"] = map[string]string{
		"":    "不加密",
		"ssl": "ssl",
		"tls": "tls",
	}
	Lang[protocol.ZH_CN]["mail"]["placeholder"] = map[string]string{
		"password": "有些邮箱需要填写单独申请的授权码，具体请到邮箱相关设置查询。",
	}
	Lang[protocol.ZH_CN]["message"]["typeList"] = map[string]string{
		"mail":    "邮件",
		"message": "浏览器通知",
		"webhook": "Webhook",
	}
	Lang[protocol.ZH_CN]["my"]["taskMenu"] = map[string]string{
		"assignedToMe": "指派给我",
		"openedByMe":   "由我创建",
		"finishedByMe": "由我完成",
		"closedByMe":   "由我关闭",
		"canceledByMe": "由我取消",
	}
	Lang[protocol.ZH_CN]["my"]["storyMenu"] = map[string]string{
		"assignedToMe": "指派给我",
		"openedByMe":   "由我创建",
		"reviewedByMe": "由我评审",
		"closedByMe":   "由我关闭",
	}
	Lang[protocol.ZH_CN]["my"]["home"] = map[string]string{
		"latest":        "最新动态",
		"action":        "%s, %s <em>%s</em> %s <a href='%s'>%s</a>。",
		"projects":      Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"products":      Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"createProject": "添加" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"createProduct": "添加" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"noProductsTip": "这里还没有" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "。",
	}
	Lang[protocol.ZH_CN]["my"]["form"] = map[string]string{
		"lblBasic":   "基本信息",
		"lblContact": "联系信息",
		"lblAccount": "帐号信息",
	}
	Lang[protocol.ZH_CN]["my"]["calendarMenu"] = map[string]interface{}{
		"todo": map[string]string{
			`link`:      `待办|my|todo|`,
			`subModule`: `todo`,
		},
		"effort": map[string]string{`link`: `日志|my|effort|`, `subModule`: `effort`},
	}
	Lang[protocol.ZH_CN]["my"]["reviewTypeList"] = map[string]string{
		"all":      "所有",
		"attend":   "考勤",
		"leave":    "请假",
		"makeup":   "补班",
		"overtime": "加班",
		"lieu":     "调休",
		"trip":     "外出",
	}
	Lang[protocol.ZH_CN]["my"]["webMenu"] = map[string]interface{}{
		"todo": map[string]string{
			`link`:      `待办|my|todo|`,
			`subModule`: `todo`,
		},
		"effort": map[string]string{`link`: `日志|my|effort|`, `subModule`: `effort`},
	}
	Lang[protocol.ZH_CN]["notice"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["notice"]["jumping"] = " 3秒钟后页面将自动跳转 <a href='%s' class='btn btn-primary btn-xs'>立即跳转</a> "
	Lang[protocol.ZH_CN]["notice"]["common"] = "错误"
	Lang[protocol.ZH_CN]["notice"]["typeList"] = map[string]string{
		"notFound":      "您要访问的内容没有找到，请检查地址是否正确",
		"accessLimited": "信息访问受限",
	}
	Lang[protocol.ZH_CN]["overtime"]["typeList"] = map[string]string{
		"rest":    "休息日加班",
		"holiday": "节假日加班",
	}
	Lang[protocol.ZH_CN]["overtime"]["statusList"] = map[string]string{
		"draft":  "草稿",
		"wait":   "等待审核",
		"doing":  "审核中",
		"pass":   "通过",
		"reject": "拒绝",
	}
	Lang[protocol.ZH_CN]["overtime"]["confirmReview"] = map[string]string{
		"pass":   "您确定要执行通过操作吗？",
		"reject": "您确定要执行拒绝操作吗？",
	}
	Lang[protocol.ZH_CN]["overtime"]["reviewStatusList"] = map[string]string{
		"pass":   "通过",
		"reject": "拒绝",
	}
	Lang[protocol.ZH_CN]["product"]["typeList"] = []protocol.HtmlKeyValueStr{
		{"", ""},
		{"normal", "正常"},
		{"branch", "多分支"},
		{"platform", "多平台"},
	}
	Lang[protocol.ZH_CN]["product"]["typeTips"] = map[string]string{
		"branch":   "(适用于客户定制场景)",
		"platform": "(适用于跨平台应用开发，比如ios、安卓、pc端等)",
	}
	Lang[protocol.ZH_CN]["product"]["branchName"] = map[string]string{
		"normal":   "",
		"branch":   "分支",
		"platform": "平台",
	}
	Lang[protocol.ZH_CN]["product"]["statusList"] = []protocol.HtmlKeyValueStr{
		{"", ""},
		{"normal", "正常"},
		{"closed", "结束"},
	}
	Lang[protocol.ZH_CN]["product"]["aclList"] = []protocol.HtmlKeyValueStr{
		{"private", "私有" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "(只有" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "相关负责人和" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "团队成员才能访问)"},
		{"custom", "自定义白名单(团队成员和白名单的成员可以访问)"},
	}
	Lang[protocol.ZH_CN]["product"]["featureBar"] = map[string][]protocol.HtmlKeyValueStr{
		"browse": []protocol.HtmlKeyValueStr{
			{"allstory", Lang[protocol.ZH_CN]["product"]["allStory"].(string)},
			{"unclosed", Lang[protocol.ZH_CN]["product"]["unclosed"].(string)},
			{"assignedtome", Lang[protocol.ZH_CN]["product"]["assignedToMe"].(string)},
			{"openedbyme", Lang[protocol.ZH_CN]["product"]["openedByMe"].(string)},
			{"reviewedbyme", Lang[protocol.ZH_CN]["product"]["reviewedByMe"].(string)},
			{"draftstory", Lang[protocol.ZH_CN]["product"]["draftStory"].(string)},
			{"more", Lang[protocol.ZH_CN]["common"]["more"].(string)},
		},
		"all": []protocol.HtmlKeyValueStr{
			{"noclosed", Lang[protocol.ZH_CN]["product"]["unclosed"].(string)},
			{"closed", "结束"},
			{"all", Lang[protocol.ZH_CN]["product"]["allProduct"].(string)},
		},
	}
	Lang[protocol.ZH_CN]["product"]["moreSelects"] = []protocol.HtmlKeyValueStr{
		{"closedbyme", Lang[protocol.ZH_CN]["product"]["closedByMe"].(string)},
		{"activestory", Lang[protocol.ZH_CN]["product"]["activeStory"].(string)},
		{"changedstory", Lang[protocol.ZH_CN]["product"]["changedStory"].(string)},
		{"willclose", Lang[protocol.ZH_CN]["product"]["willClose"].(string)},
		{"closedstory", Lang[protocol.ZH_CN]["product"]["closedStory"].(string)},
	}
	Lang[protocol.ZH_CN]["productplan"]["endList"] = []protocol.HtmlKeyValueStr{
		{"7", "一星期"},
		{"14", "两星期"},
		{"31", "一个月"},
		{"62", "两个月"},
		{"93", "三个月"},
		{"186", "半年"},
		{"365", "一年"},
	}
	Lang[protocol.ZH_CN]["productplan"]["featureBar"] = map[string]interface{}{
		"browse": map[string]string{
			"all":       "全部",
			"unexpired": "未过期",
			"overdue":   "已过期",
		},
	}
	Lang[protocol.ZH_CN]["project"]["typeList"] = []protocol.HtmlKeyValueStr{
		{"sprint", "短期" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)},
		{"waterfall", "长期" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)},
		{"ops", "运维" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string)},
	}
	Lang[protocol.ZH_CN]["project"]["endList"] = []protocol.HtmlKeyValueStr{
		{"7", "一星期"},
		{"14", "两星期"},
		{"31", "一个月"},
		{"62", "两个月"},
		{"93", "三个月"},
		{"186", "半年"},
		{"365", "一年"},
	}
	Lang[protocol.ZH_CN]["team"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["team"]["limited"] = "受限用户"
	Lang[protocol.ZH_CN]["team"]["totalHours"] = "总计"
	Lang[protocol.ZH_CN]["team"]["days"] = "可用工日"
	Lang[protocol.ZH_CN]["team"]["hours"] = "可用工时/天"
	Lang[protocol.ZH_CN]["team"]["join"] = "加盟日"
	Lang[protocol.ZH_CN]["team"]["role"] = "角色"
	Lang[protocol.ZH_CN]["team"]["account"] = "用户"
	Lang[protocol.ZH_CN]["team"]["limitedList"] = []protocol.HtmlKeyValueStr{
		{"no", "否"},
		{"yes", "是"},
	}
	Lang[protocol.ZH_CN]["project"]["statusList"] = []protocol.HtmlKeyValueStr{

		{"wait", "未开始"},
		{"doing", "进行中"},
		{"suspended", "已挂起"},
		{"closed", "已关闭"},
	}
	Lang[protocol.ZH_CN]["project"]["aclList"] = []protocol.HtmlKeyValueStr{
		{"private", "私有" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "(只有" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "团队成员才能访问)"},
		{"custom", "自定义白名单(团队成员和白名单的成员可以访问)"},
	}
	Lang[protocol.ZH_CN]["project"]["statusSelects"] = []protocol.HtmlKeyValueStr{
		{"wait", "未开始"},
		{"doing", "进行中"},
		{"finishedbyme", "我完成"},
		{"done", "已完成"},
		{"closed", "已关闭"},
		{"cancel", "已取消"},
		{"", "更多"},
	}
	Lang[protocol.ZH_CN]["project"]["groups"] = map[string]string{
		"story":      "需求分组",
		"status":     "状态分组",
		"pri":        "优先级分组",
		"assignedTo": "指派给分组",
		"finishedBy": "完成者分组",
		"closedBy":   "关闭者分组",
		"type":       "类型分组",
	}
	Lang[protocol.ZH_CN]["project"]["groupFilter"] = map[string]interface{}{
		"story": map[string]string{
			"linked": "已关联需求的任务",
			"all":    Lang[protocol.ZH_CN]["project"]["all"].(string),
		},
		"pri": map[string]string{
			"noset": "未设置",
			"all":   Lang[protocol.ZH_CN]["project"]["all"].(string),
		},
		"assignedTo": map[string]string{
			"undone": "未完成",
			"all":    Lang[protocol.ZH_CN]["project"]["all"].(string),
		},
	}
	Lang[protocol.ZH_CN]["project"]["placeholder"] = map[string]string{
		"code":      "团队内部的简称",
		"totalLeft": "项目开始时的总预计工时",
	}
	Lang[protocol.ZH_CN]["project"]["selectGroup"] = map[string]string{
		"done": "(已结束)",
	}
	Lang[protocol.ZH_CN]["project"]["orderList"] = map[string]string{
		"order_asc":  "需求排序正序",
		"order_desc": "需求排序倒序",
		"pri_asc":    "需求优先级正序",
		"pri_desc":   "需求优先级倒序",
		"stage_asc":  "需求阶段正序",
		"stage_desc": "需求阶段倒序",
	}
	Lang[protocol.ZH_CN]["kanbanSetting"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["kanbanSetting"]["subStatusTips"] = "请分别设置任务看板和Bug看板下的显示列以及颜色。"
	Lang[protocol.ZH_CN]["kanbanSetting"]["emptyColumns"] = "看板显示列不能为空。"
	Lang[protocol.ZH_CN]["kanbanSetting"]["subStatus"] = "看板显示列"
	Lang[protocol.ZH_CN]["kanbanSetting"]["mode"] = "看板模式"
	Lang[protocol.ZH_CN]["kanbanSetting"]["laneField"] = "看板列字段"
	Lang[protocol.ZH_CN]["kanbanSetting"]["noticeReset"] = "是否恢复看板默认设置？"
	Lang[protocol.ZH_CN]["kanbanSetting"]["optionList"] = []string{"隐藏", `显示`}
	Lang[protocol.ZH_CN]["printKanban"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["printKanban"]["taskStatus"] = "状态"
	Lang[protocol.ZH_CN]["printKanban"]["print"] = "打印"
	Lang[protocol.ZH_CN]["printKanban"]["content"] = "内容"
	Lang[protocol.ZH_CN]["printKanban"]["common"] = "看板打印"
	Lang[protocol.ZH_CN]["printKanban"]["typeList"] = map[string]string{
		"all":       "全部",
		"increment": "增量",
	}
	Lang[protocol.ZH_CN]["project"]["featureBar"] = map[string][]protocol.HtmlKeyValueStr{
		"task": []protocol.HtmlKeyValueStr{
			{"all", Lang[protocol.ZH_CN]["project"]["allTasks"].(string)},
			{"unclosed", Lang[protocol.ZH_CN]["project"]["unclosed"].(string)},
			{"assignedtome", Lang[protocol.ZH_CN]["project"]["assignedToMe"].(string)},
			{"myinvolved", Lang[protocol.ZH_CN]["project"]["myInvolved"].(string)},
			{"delayed", "已延期"},
			{"needconfirm", "需求变更"},
			{"status", "更多"},
		},
	}
	Lang[protocol.ZH_CN]["project"]["treeLevel"] = map[string]string{
		"all":   "全部展开",
		"root":  "全部折叠",
		"story": "只看需求",
		"task":  "只看任务",
	}
	Lang[protocol.ZH_CN]["kanbanSetting"]["laneFields"] = map[string]string{
		"status":    "状态",
		"subStatus": "子状态",
	}
	Lang[protocol.ZH_CN]["kanbanSetting"]["modeList"] = map[string]string{
		"task": "任务看板",
		"bug":  "Bug看板",
	}
	Lang[protocol.ZH_CN]["project"]["gantt"] = map[string]interface{}{
		"common":                   "甘特图",
		"id":                       "编号",
		"pretask":                  "条件任务",
		"condition":                "条件动作",
		"task":                     "任务",
		"action":                   "动作",
		"type":                     "关系类型",
		"exportImg":                "导出图片",
		"exportPDF":                "导出 PDF",
		"exporting":                "正在导出……",
		"exportFail":               "导出失败。",
		"createRelationOfTasks":    "创建任务关系",
		"newCreateRelationOfTasks": "新增任务关系",
		"editRelationOfTasks":      "维护任务关系",
		"relationOfTasks":          "查看任务关系",
		"relation":                 "任务关系",
		"showCriticalPath":         "显示关键路径",
		"hideCriticalPath":         "隐藏关键路径",
		"fullScreen":               "全屏",
		"zooming": map[string]string{
			"day":   "天",
			"week":  "周",
			"month": "月",
		},
		"assignTo":  "指派给",
		"duration":  "持续天数",
		"comp":      "进度",
		"startDate": "开始日期",
		"endDate":   "结束日期",
		"days":      " 天",
		"format":    "查看格式",
		"preTaskStatus": map[string]string{
			"":      "",
			"end":   "完成后",
			"begin": "开始后",
		},
		"taskActions": map[string]string{
			"":      "",
			"begin": "才能开始",
			"end":   "才能完成",
		},
		"color": []string{"bbb", "ff5d5d", "ff9800", "16a8f8", "00da88"},
		"browseType": map[string]string{
			"type":       "按任务类型分组",
			"module":     "按模块分组",
			"assignedTo": "按指派给分组",
			"story":      "按{$lang->common->storyCommon}分组",
		},
		"confirmDelete": "确认要删除此任务关系吗？",
		"tmpNotWrite":   "不可写",
		"warning": map[string]string{
			"noEditSame":     "已有的编号%s前后任务不能相同！",
			"noEditRepeat":   "已有的编号%s与已有的编号%s任务关系之间重复！",
			"noEditContrary": "已有的编号%s与已有的编号%s任务关系之间有矛盾！",
			"noRepeat":       "已有的编号%s与新增的编号%s任务关系之间重复！",
			"noContrary":     "已有的编号%s与新增的编号%s任务关系之间有矛盾！",
			"noNewSame":      "新增的编号%s前后任务不能相同！",
			"noNewRepeat":    "新增的编号%s与新增的编号%s任务关系之间重复！",
			"noNewContrary":  "新增的编号%s与新增的编号%s任务关系之间有矛盾！",
		},
	}
	Lang[protocol.ZH_CN]["project"]["webMenu"] = map[string]interface{}{
		"task": map[string]string{
			`link`:      "任务|project|task|projectID=%s",
			`subModule`: `task`,
			`alias`:     `importtask,importbug,tree`,
		},
		"story":    "{$lang->common->storyCommon}|project|story|projectID=%s",
		"bug":      "Bug|project|bug|projectID=%s",
		"build":    map[string]string{`link`: "版本|project|build|projectID=%s", `subModule`: `build`},
		"testtask": "测试单|project|testtask|projectID=%s",
		"team":     "团队|project|team|projectID=%s",
		//"action":   "动态|project|dynamic|projectID=%s",
		"view": "概况|project|view|projectID=%s",
		"all":  "所有项目|project|all|",
	}
	Lang[protocol.ZH_CN]["project"]["webMenuOrder"] = []string{"task", "story", "bug", "build", "testtask", "team", "action", "view", "all"}
	Lang[protocol.ZH_CN]["project"]["charts"] = map[string]interface{}{
		"burn": map[string]map[string]string{
			"graph": map[string]string{
				"caption":      "燃尽图",
				"xAxisName":    "日期",
				"yAxisName":    "HOUR",
				"baseFontSize": "12",
				"formatNumber": "0",
				"animation":    "0",
				"rotateNames":  "1",
				"showValues":   "0",
				"reference":    "参考",
				"actuality":    "实际",
			},
		},
	}
	Lang[protocol.ZH_CN]["release"]["exportTypeList"] = map[string]string{
		"all":     "所有",
		"story":   "需求",
		"bug":     "Bug",
		"leftbug": "遗留Bug",
	}
	Lang[protocol.ZH_CN]["release"]["statusList"] = map[string]string{
		"":          "",
		"normal":    "正常",
		"terminate": "停止维护",
	}
	Lang[protocol.ZH_CN]["release"]["changeStatusList"] = map[string]string{
		"normal":    "激活",
		"terminate": "停止维护",
	}
	Lang[protocol.ZH_CN]["release"]["action"] = map[string]interface{}{
		"changestatus": map[string]string{
			`main`:  `$date, 由 <strong>$actor</strong> $extra。`,
			`extra`: `changeStatusList`,
		},
	}
	Lang[protocol.ZH_CN]["report"]["assign"] = map[string]string{
		"noassign": "未指派",
		"assign":   "已指派",
	}
	Lang[protocol.ZH_CN]["report"]["typeList"] = map[string]string{
		"default": "默认",
		"pie":     "饼图",
		"bar":     "柱状图",
		"line":    "折线图",
	}
	Lang[protocol.ZH_CN]["report"]["mailTitle"] = map[string]string{
		"begin":    "提醒：您有",
		"bug":      " Bug(%s),",
		"task":     " 任务(%s),",
		"todo":     " 待办(%s),",
		"testTask": " 测试版本(%s),",
	}
	Lang[protocol.ZH_CN]["report"]["annualData"] = map[string]string{
		"title":             "%s年工作内容统计一览表 —— %s",
		"baseInfo":          "基本数据信息",
		"logins":            "累计登录次数",
		"actions":           "累计动态数",
		"efforts":           "累计日志数",
		"consumed":          "累计工时数",
		"foundBugs":         "累计创建Bug数",
		"createdCases":      "累计创建用例数",
		"involvedProducts":  "累计参与" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "数",
		"createdPlans":      "累计创建计划数",
		"createdStories":    "累计创建" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "数",
		"productOverview":   Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "创建" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "数及占比",
		"qaOverview":        Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "创建Bug数及占比",
		"projectOverview":   "参与" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "概览",
		"doneProject":       "已完成的" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"doingProject":      "正在进行的" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"suspendProject":    "已挂起的" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"projectName":       Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "名称",
		"finishedStory":     "完成" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "数",
		"finishedTask":      "完成任务数",
		"foundBug":          "创建Bug数",
		"resolvedBug":       "解决Bug数",
		"productName":       Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "名称",
		"planCount":         "计划数",
		"storyCount":        Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "数",
		"qaData":            "累计创建Bug数和创建用例数",
		"totalCreatedBug":   "累计创建Bug数",
		"totalCreatedCase":  "累计创建用例数",
		"devData":           "完成任务数和解决Bug数",
		"totalFinishedTask": "完成任务数",
		"totalResolvedBug":  "解决Bug数",
		"totalConsumed":     "累计工时",
		"poData":            "所创建" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "数和对应的优先级及状态",
		"totalStoryPri":     "创建" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "优先级分布",
		"totalStoryStage":   "创建" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "阶段分布",
		"qaStatistics":      "月创建Bug数和创建用例数",
		"poStatistics":      "月创建" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "数",
		"devStatistics":     "月完成任务数及累计工时和解决Bug数",
		"unit":              "个",
	}
	Lang[protocol.ZH_CN]["crystal"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["crystal"]["emptyName"] = "报表名称不能为空。"
	Lang[protocol.ZH_CN]["crystal"]["noStep"] = "请先查询出数据再保存报表！"
	Lang[protocol.ZH_CN]["crystal"]["noSumAppend"] = "第%s行没有选择求和字段"
	Lang[protocol.ZH_CN]["crystal"]["errorSql"] = "SQL语句有错！错误："
	Lang[protocol.ZH_CN]["crystal"]["confirmDelete"] = "是否删除此报表？"
	Lang[protocol.ZH_CN]["crystal"]["errorNoReport"] = "不存在此报表！"
	Lang[protocol.ZH_CN]["crystal"]["errorSave"] = "该SQL的变量信息不能为空！"
	Lang[protocol.ZH_CN]["crystal"]["noticeShowName"] = "变量%s的显示名称没有设置"
	Lang[protocol.ZH_CN]["crystal"]["noticeRequestType"] = "变量%s的输入方式没有设置"
	Lang[protocol.ZH_CN]["crystal"]["noticeVarName"] = "变量名称没有设置"
	Lang[protocol.ZH_CN]["crystal"]["noticeResult"] = "共有%s条数据，显示了%s条"
	Lang[protocol.ZH_CN]["crystal"]["notResult"] = "没有查询数据."
	Lang[protocol.ZH_CN]["crystal"]["noticeBlack"] = "SQL中含有禁用SQL关键字 %s"
	Lang[protocol.ZH_CN]["crystal"]["noticeSelect"] = "SQL语句只能是查询语句"
	Lang[protocol.ZH_CN]["crystal"]["codePlaceholder"] = "报表的唯一代号"
	Lang[protocol.ZH_CN]["crystal"]["sumPlaceholder"] = "选择求和字段"
	Lang[protocol.ZH_CN]["crystal"]["sqlPlaceholder"] = "直接写入一句SQL查询语句，只能进行查询操作，禁止其他SQL操作"

	Lang[protocol.ZH_CN]["crystal"]["default"] = "默认值"
	Lang[protocol.ZH_CN]["crystal"]["desc"] = "描述"
	Lang[protocol.ZH_CN]["crystal"]["lang"] = "语言"
	Lang[protocol.ZH_CN]["crystal"]["fieldValue"] = "显示名"
	Lang[protocol.ZH_CN]["crystal"]["fieldName"] = "字段名"
	Lang[protocol.ZH_CN]["crystal"]["all"] = "所有"
	Lang[protocol.ZH_CN]["crystal"]["code"] = "代号"
	Lang[protocol.ZH_CN]["crystal"]["id"] = "编号"
	Lang[protocol.ZH_CN]["crystal"]["module"] = "所属类目"
	Lang[protocol.ZH_CN]["crystal"]["name"] = "报表名称"
	Lang[protocol.ZH_CN]["crystal"]["showName"] = "显示名称"
	Lang[protocol.ZH_CN]["crystal"]["varName"] = "变量名称"
	Lang[protocol.ZH_CN]["crystal"]["requestType"] = "输入方式"
	Lang[protocol.ZH_CN]["crystal"]["total"] = "总计"
	Lang[protocol.ZH_CN]["crystal"]["isUser"] = "显示用户真实姓名"
	Lang[protocol.ZH_CN]["crystal"]["percentAB"] = "百分比"
	Lang[protocol.ZH_CN]["crystal"]["showAlone"] = "独占一列"
	Lang[protocol.ZH_CN]["crystal"]["contrast"] = "对比"
	Lang[protocol.ZH_CN]["crystal"]["percent"] = "显示百分比"
	Lang[protocol.ZH_CN]["crystal"]["reportTotal"] = "显示汇总"
	Lang[protocol.ZH_CN]["crystal"]["reportType"] = "统计方式"
	Lang[protocol.ZH_CN]["crystal"]["reportField"] = "要统计的字段"
	Lang[protocol.ZH_CN]["crystal"]["group2"] = "第二分组字段"
	Lang[protocol.ZH_CN]["crystal"]["group1"] = "第一分组字段"
	Lang[protocol.ZH_CN]["crystal"]["statistics"] = "统计字段"
	Lang[protocol.ZH_CN]["crystal"]["group"] = "分组字段"
	Lang[protocol.ZH_CN]["crystal"]["result"] = "查询结果"
	Lang[protocol.ZH_CN]["crystal"]["params"] = "报表条件"
	Lang[protocol.ZH_CN]["crystal"]["condition"] = "报表设计"
	Lang[protocol.ZH_CN]["crystal"]["query"] = "查询"
	Lang[protocol.ZH_CN]["crystal"]["sql"] = "查询语句"
	Lang[protocol.ZH_CN]["crystal"]["saveAs"] = "另存为"
	Lang[protocol.ZH_CN]["crystal"]["custom"] = "新增报表"
	Lang[protocol.ZH_CN]["crystal"]["addLang"] = "设置字段名"
	Lang[protocol.ZH_CN]["crystal"]["addVar"] = "添加变量"
	Lang[protocol.ZH_CN]["crystal"]["use"] = "设计"
	Lang[protocol.ZH_CN]["crystal"]["browse"] = "已保存报表"
	Lang[protocol.ZH_CN]["crystal"]["setVar"] = "设置变量"
	Lang[protocol.ZH_CN]["crystal"]["common"] = "水晶报表"
	Lang[protocol.ZH_CN]["report"]["useReport"] = Lang[protocol.ZH_CN]["crystal"]["use"]
	Lang[protocol.ZH_CN]["report"]["custom"] = Lang[protocol.ZH_CN]["crystal"]["custom"]
	Lang[protocol.ZH_CN]["crystal"]["reportTypeList"] = map[string]string{
		"count": "计数",
		"sum":   "求和",
	}
	Lang[protocol.ZH_CN]["crystal"]["requestTypeList"] = map[string]string{
		"input":  "文本框",
		"date":   "日期",
		"select": "下拉菜单",
	}
	Lang[protocol.ZH_CN]["crystal"]["selectList"] = map[string]string{
		"project.status": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `状态列表`,
		"user":           "用户列表",
		"product":        Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "列表",
		"project":        Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "列表",
		"dept":           "部门列表",
	}
	Lang[protocol.ZH_CN]["crystal"]["moduleList"] = map[string]string{
		"":        "",
		"product": Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		"project": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		"test":    "测试",
		"staff":   "组织",
	}
	Lang[protocol.ZH_CN]["report"]["case"] = map[string]string{
		"total":    "总用例数",
		"run":      "总执行数",
		"passRate": "用例通过率",
		"name":     "名称",
	}
	Lang[protocol.ZH_CN]["report"]["bugTypeList"] = map[string]string{
		"codeerror":   "代码",
		"interface":   "界面",
		"config":      "配置",
		"install":     "安装",
		"security":    "安全",
		"performance": "性能",
		"standard":    "标准",
		"automation":  "脚本",
		"others":      "其他",
	}
	Lang[protocol.ZH_CN]["report"]["bug"] = map[string]string{
		"total":  "总计",
		"title":  "Bug标题",
		"status": "状态",
		"story":  Lang[protocol.ZH_CN]["common"]["storyCommon"].(string),
	}
	Lang[protocol.ZH_CN]["report"]["storyLinkedBug"] = Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + `关联Bug汇总表`
	Lang[protocol.ZH_CN]["reportList"] = map[string]interface{}{}
	Lang[protocol.ZH_CN]["reportList"]["project"] = []string{Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + `偏差报表|report|projectdeviation`}
	Lang[protocol.ZH_CN]["reportList"]["product"] = []string{Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `汇总表|report|productsummary`, Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + `路线图表|report|roadmap`}
	Lang[protocol.ZH_CN]["reportList"]["test"] = []string{`Bug创建表|report|bugcreate`, `Bug指派表|report|bugassign`, `用例统计表|report|testcase`, `用例执行统计表|report|casesrun`, `版本统计表|report|build`, Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + `关联Bug汇总表|report|storylinkedbug`}
	Lang[protocol.ZH_CN]["reportList"]["staff"] = []string{"员工负载表|report|workload", "任务完成汇总表|report|worksummary", "任务指派汇总表|report|workAssignSummary", "Bug解决汇总表|report|bugsummary", "Bug指派汇总表|report|bugAssignSummary"}
	Lang[protocol.ZH_CN]["search"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["search"]["null"] = "空"
	Lang[protocol.ZH_CN]["search"]["sql"] = "SQL条件"
	Lang[protocol.ZH_CN]["search"]["form"] = "表单字段"
	Lang[protocol.ZH_CN]["search"]["title"] = "名称"
	Lang[protocol.ZH_CN]["search"]["module"] = "模块"
	Lang[protocol.ZH_CN]["search"]["account"] = "用户名"
	Lang[protocol.ZH_CN]["search"]["custom"] = "自定义"
	Lang[protocol.ZH_CN]["search"]["onMenuBar"] = "显示在菜单栏"
	Lang[protocol.ZH_CN]["search"]["shortcut"] = Lang[protocol.ZH_CN]["search"]["onMenuBar"]
	Lang[protocol.ZH_CN]["search"]["noQuery"] = "还没有保存查询！"
	Lang[protocol.ZH_CN]["search"]["me"] = "自己"
	Lang[protocol.ZH_CN]["search"]["select"] = "需求/任务筛选"
	Lang[protocol.ZH_CN]["search"]["setQueryTitle"] = "请输入查询标题（保存之前请先查询）："
	Lang[protocol.ZH_CN]["search"]["deleteQuery"] = "删除查询"
	Lang[protocol.ZH_CN]["search"]["savedQuery"] = "已保存的检索历史"
	Lang[protocol.ZH_CN]["search"]["buildQuery"] = "执行搜索"
	Lang[protocol.ZH_CN]["search"]["buildForm"] = "搜索表单"
	Lang[protocol.ZH_CN]["search"]["group2"] = "第二组"
	Lang[protocol.ZH_CN]["search"]["group1"] = "第一组"
	Lang[protocol.ZH_CN]["search"]["myQuery"] = "我的查询"
	Lang[protocol.ZH_CN]["search"]["saveQuery"] = "保存"
	Lang[protocol.ZH_CN]["search"]["reset"] = "重置"
	Lang[protocol.ZH_CN]["search"]["common"] = "搜索"
	Lang[protocol.ZH_CN]["search"]["operators"] = []protocol.HtmlKeyValueStr{
		{"=", "="},
		{"!=", "!="},
		{">", ">"},
		{">=", ">="},
		{"<", "<"},
		{"<=", "<="},
		{"include", "包含"},
		{"between", "介于"},
		{"notinclude", "不包含"},
		{"belong", "从属于"},
	}
	Lang[protocol.ZH_CN]["search"]["andor"] = []protocol.HtmlKeyValueStr{
		{"and", "并且"},
		{"or", "或者"},
	}
	Lang[protocol.ZH_CN]["search"]["error"] = map[string]string{
		"notFoundParamsFunc": "没有找到%s/%s的搜索参数初始化方法，请返回首页刷新重试",
	}
	Lang[protocol.ZH_CN]["story"]["useList"] = []string{"不使用", `使用`}
	Lang[protocol.ZH_CN]["story"]["statusList"] = []protocol.HtmlKeyValueStr{
		{"", ""},
		{"draft", "草稿"},
		{"active", "激活"},
		{"closed", "已关闭"},
		{"changed", "已变更"},
	}
	Lang[protocol.ZH_CN]["story"]["stageList"] = []protocol.HtmlKeyValueStr{
		{"", ""},
		{"wait", "未开始"},
		{"planned", "已计划"},
		{"projected", "已立项"},
		{"developing", "研发中"},
		{"developed", "研发完毕"},
		{"testing", "测试中"},
		{"tested", "测试完毕"},
		{"verified", "已验收"},
		{"released", "已发布"},
		{"closed", "已关闭"},
	}
	Lang[protocol.ZH_CN]["story"]["reasonList"] = []protocol.HtmlKeyValueStr{
		{"", ""},
		{"done", "已完成"},
		{"subdivided", "已细分"},
		{"duplicate", "重复"},
		{"postponed", "延期"},
		{"willnotdo", "不做"},
		{"cancel", "已取消"},
		{"bydesign", "设计如此"},
	}
	Lang[protocol.ZH_CN]["story"]["reviewResultList"] = []protocol.HtmlKeyValueStr{
		{"", ""},
		{"pass", "确认通过"},
		{"revert", "撤销变更"},
		{"clarify", "有待明确"},
		{"reject", "拒绝"},
	}
	Lang[protocol.ZH_CN]["story"]["reviewList"] = []string{"否", `是`}
	Lang[protocol.ZH_CN]["story"]["sourceList"] = []protocol.HtmlKeyValueStr{
		{"", ""},
		{"customer", "客户"},
		{"user", "用户"},
		{"po", Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "经理"},
		{"market", "市场"},
		{"service", "客服"},
		{"operation", "运营"},
		{"support", "技术支持"},
		{"competitor", "竞争对手"},
		{"partner", "合作伙伴"},
		{"dev", "开发人员"},
		{"tester", "测试人员"},
		{"bug", "Bug"},
		{"other", "其他"},
	}
	Lang[protocol.ZH_CN]["story"]["priList"] = []protocol.HtmlKeyValueStr{{"", ""}, {"1", "1"}, {"2", "2"}, {"3", "3"}, {"4", "4"}}
	Lang[protocol.ZH_CN]["story"]["form"] = map[string]string{
		"area":     "该需求所属范围",
		"desc":     "描述及标准，什么需求？如何验收？",
		"resource": "资源分配，有谁完成？需要多少时间？",
		"file":     "附件，如果该需求有相关文件，请点此上传。",
	}
	Lang[protocol.ZH_CN]["story"]["action"] = map[string]interface{}{
		"reviewed": map[string]string{
			`main`:  `$date, 由 <strong>$actor</strong> 记录评审结果，结果为 <strong>$extra</strong>。`,
			`extra`: `reviewResultList`,
		},
		"closed":              map[string]string{`main`: `$date, 由 <strong>$actor</strong> 关闭，原因为 <strong>$extra</strong> $appendLink。`, `extra`: `reasonList`},
		"linked2plan":         map[string]string{`main`: `$date, 由 <strong>$actor</strong> 关联到计划 <strong>$extra</strong>。`},
		"unlinkedfromplan":    map[string]string{`main`: `$date, 由 <strong>$actor</strong> 从计划 <strong>$extra</strong> 移除。`},
		"linked2project":      map[string]string{`main`: `$date, 由 <strong>$actor</strong> 关联到` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + ` <strong>$extra</strong>。`},
		"unlinkedfromproject": map[string]string{`main`: `$date, 由 <strong>$actor</strong> 从` + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + ` <strong>$extra</strong> 移除。`},
		"linked2build":        map[string]string{`main`: `$date, 由 <strong>$actor</strong> 关联到版本 <strong>$extra</strong>。`},
		"unlinkedfrombuild":   map[string]string{`main`: `$date, 由 <strong>$actor</strong> 从版本 <strong>$extra</strong> 移除。`},
		"linked2release":      map[string]string{`main`: `$date, 由 <strong>$actor</strong> 关联到发布 <strong>$extra</strong>。`},
		"unlinkedfromrelease": map[string]string{`main`: `$date, 由 <strong>$actor</strong> 从发布 <strong>$extra</strong> 移除。`},
		"linkrelatedstory":    map[string]string{`main`: `$date, 由 <strong>$actor</strong> 关联相关需求 <strong>$extra</strong>。`},
		"subdividestory":      map[string]string{`main`: `$date, 由 <strong>$actor</strong> 细分为需求 <strong>$extra</strong>。`},
		"unlinkrelatedstory":  map[string]string{`main`: `$date, 由 <strong>$actor</strong> 移除相关需求 <strong>$extra</strong>。`},
		"unlinkchildstory":    map[string]string{`main`: `$date, 由 <strong>$actor</strong> 移除细分需求 <strong>$extra</strong>。`},
	}
	Lang[protocol.ZH_CN]["story"]["report"] = map[string]interface{}{
		"common": "报表",
		"select": "请选择报表类型",
		"create": "生成报表",
		"value":  "需求数",
		"charts": map[string]string{
			"storysPerProduct":      Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "需求数量",
			"storysPerModule":       "模块需求数量",
			"storysPerSource":       "需求来源统计",
			"storysPerPlan":         "计划进行统计",
			"storysPerStatus":       "状态进行统计",
			"storysPerStage":        "所处阶段进行统计",
			"storysPerPri":          "优先级进行统计",
			"storysPerEstimate":     "预计工时进行统计",
			"storysPerOpenedBy":     "由谁创建来进行统计",
			"storysPerAssignedTo":   "当前指派来进行统计",
			"storysPerClosedReason": "关闭原因来进行统计",
			"storysPerChange":       "变更次数来进行统计",
		},
		"options": map[string]string{
			"type":   "pie",
			"width":  "500",
			"height": "140",
		},
		"storysPerModule": map[string]string{
			"item": "模块",
		},
		"storysPerSource": map[string]string{
			"item": "来源",
		},
		"storysPerPlan": map[string]string{
			"item": "计划",
		},
		"storysPerStatus": map[string]string{
			"item": "状态",
		},
		"storysPerStage": map[string]interface{}{
			"item": "阶段",
			"graph": map[string]string{
				"xAxisName": "所处阶段",
			},
		},
		"storysPerPri": map[string]string{
			"item": "优先级",
		},
		"storysPerOpenedBy": map[string]interface{}{
			"item": "用户",
			"graph": map[string]string{
				"xAxisName": "由谁创建",
			},
		},
		"storysPerAssignedTo": map[string]interface{}{
			"item": "用户",
			"graph": map[string]string{
				"xAxisName": "当前指派",
			},
		},
		"storysPerClosedReason": map[string]interface{}{
			"item": "原因",
			"graph": map[string]string{
				"xAxisName": "关闭原因",
			},
		},
		"storysPerEstimate": map[string]string{
			"item": "预计工时",
		},
		"storysPerChange": map[string]string{
			"item": "变更次数",
		},
		"storysPerProduct": map[string]interface{}{
			"item": Lang[protocol.ZH_CN]["common"]["productCommon"].(string),
		},
	}
	Lang[protocol.ZH_CN]["story"]["chosen"] = map[string]string{
		"reviewedBy": "选择评审人...",
	}
	Lang[protocol.ZH_CN]["story"]["notice"] = map[string]string{
		"closed": "您选择的需求已经被关闭了！",
	}
	Lang[protocol.ZH_CN]["story"]["placeholder"] = map[string]string{
		"estimate": Lang[protocol.ZH_CN]["story"]["hour"].(string),
	}
	Lang[protocol.ZH_CN]["svn"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["svn"]["apiSync"] = "接口：同步svn日志"
	Lang[protocol.ZH_CN]["svn"]["diff"] = "比较源代码"
	Lang[protocol.ZH_CN]["svn"]["cat"] = "查看源代码"
	Lang[protocol.ZH_CN]["svn"]["common"] = "Subversion"

	Lang[protocol.ZH_CN]["task"]["examinePass"] = []protocol.HtmlKeyValueStr{{"true", "通过"}, {"false", "未通过"}}

	Lang[protocol.ZH_CN]["task"]["proofreadingPass"] = []protocol.HtmlKeyValueStr{{"true", "是"}, {"false", "否"}}

	Lang[protocol.ZH_CN]["task"]["statusList"] = []protocol.HtmlKeyValueStr{
		{"", ""},
		{"wait", "未开始"},
		{"doing", "进行中"},
		{"done", "已完成"},
		{"pause", "已暂停"},
		{"cancel", "已取消"},
		{"closed", "已关闭"},
		{"internalaudit", "内审中"},
	}

	Lang[protocol.ZH_CN]["task"]["typeList"] = []protocol.HtmlKeyValueStr{
		{"", ""},
		{"character2d", "2D_人物"},
		{"scene2d", "2D_场景"},
		{"ui2d", "2D_UI"},
		{"animation2d", "2D_动画"},
		{"cutChart2d", "2D_切图"},
		{"specialEffects2d", "2D_特效"},
		{"model3d", "3D_模型"},
		{"binding3d", "3D_绑定"},
		{"animation3d", "3D_动画"},
		{"specialEffects3d", "3D_特效"},
		{"rendering3d", "3D_渲染"},
		{"misc", "其他"},
	}

	Lang[protocol.ZH_CN]["task"]["priList"] = []protocol.HtmlKeyValueStr{{"1", "1"}, {`2`, "2"}, {`3`, "3"}, {"4", `4`}}

	Lang[protocol.ZH_CN]["task"]["reasonList"] = []protocol.HtmlKeyValueStr{
		{"", ""},
		{"done", "已完成"},
		{"cancel", "已取消"},
	}

	Lang[protocol.ZH_CN]["task"]["afterChoices"] = []protocol.HtmlKeyValueStr{
		{"continueAdding", "继续为该需求添加任务"},
		{"toTaskList", "返回任务列表"},
		{"toStoryList", "返回需求列表"},
	}

	Lang[protocol.ZH_CN]["task"]["error"] = map[string]string{
		"consumedNumber":          "\"已经消耗\"必须为数字,不能为负数",
		"leftNumber":              "\"剩余\"必须为数字,不能为负数",
		"estimateNumber":          "\"预计剩余\"必须为数字",
		"consumedSmall":           "\"已经消耗\"必须大于之前消耗",
		"consumedThisTime":        "请填写\"工时\"",
		"left":                    "请填写\"剩余\"",
		"work":                    "\"备注\"必须小于%d个字符",
		"skipClose":               "任务：%s 不是“已完成”或“已取消”状态，确定要关闭吗？",
		"consumed":                "任务：%s工时不能小于0，忽略该任务工时的改动",
		"assignedTo":              "当前状态的多人任务不能指派给任务团队外的成员。",
		"assignedToNotFoundUser":  "未找到指派人",
		"notFoundTask":            "没有找到任务信息",
		"estStarted":              "日程规划起始时间不对",
		"deadline":                "日程规划结束时间不对",
		"deadlineGtEstStarted":    "起始时间不能大于结束时间",
		"taskHasAncestors":        "当前任务包含孙任务，无法操作",
		"dateError":               "日期格式不正确",
		"newProject":              "找不到新的所属项目，请重新选择",
		"closedReasonNotempty":    "关闭原因不能为空",
		"taskleftnotempty":        fmt.Sprintf(Lang[protocol.ZH_CN]["error"]["notempty"].(string), Lang[protocol.ZH_CN]["task"]["left"].(string)),
		"RealStartedErr":          "实际开始填写错误,必须为日期",
		"finishedDateErr":         "完成时间填写错误,必须为日期",
		"TaskIsexist":             "已有相同名字的任务，无法创建",
		"errorTaskType":           "任务类型不能为空",
		"taskCanSetChildToParent": "不能将当前任务的子任务设置为当前任务的父任务",
		"taskExportNotFoundTasks": "找不到需要导出的任务，请刷新再试",
	}

	Lang[protocol.ZH_CN]["task"]["report"] = map[string]interface{}{
		"common": "报表",
		"select": "请选择报表类型",
		"create": "生成报表",
		"value":  "任务数",

		"charts": map[string]string{
			"tasksPerProject":      Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "任务数统计",
			"tasksPerModule":       "模块任务数统计",
			"tasksPerAssignedTo":   "指派给统计",
			"tasksPerType":         "任务类型统计",
			"tasksPerPri":          "优先级统计",
			"tasksPerStatus":       "任务状态统计",
			"tasksPerDeadline":     "截止日期统计",
			"tasksPerEstimate":     "预计时间统计",
			"tasksPerLeft":         "剩余时间统计",
			"tasksPerConsumed":     "消耗时间统计",
			"tasksPerFinishedBy":   "由谁完成统计",
			"tasksPerClosedReason": "关闭原因统计",
			"finishedTasksPerDay":  "每天完成统计",
		},

		"options": map[string]string{
			"type":   "pie",
			"width":  "500",
			"height": "140",
		},

		"tasksPerModule": map[string]string{
			"item": "模块",
		},

		"tasksPerAssignedTo": map[string]string{
			"item": "用户",
		},

		"tasksPerType": map[string]string{
			"item": "类型",
		},

		"tasksPerPri": map[string]string{
			"item": "优先级",
		},

		"tasksPerStatus": map[string]string{
			"item": "状态",
		},

		"tasksPerDeadline": map[string]string{
			"item": "日期",
			"type": "bar",
		},

		"tasksPerEstimate": map[string]interface{}{
			"item": "预计",
			"graph": map[string]string{
				"xAxisName": "时间",
			},
		},

		"tasksPerLeft": map[string]interface{}{
			"item": "剩余",
			"graph": map[string]string{
				"xAxisName": "时间",
			},
		},

		"tasksPerConsumed": map[string]interface{}{
			"item": "消耗",
			"graph": map[string]string{
				"xAxisName": "时间",
			},
		},

		"tasksPerFinishedBy": map[string]string{
			"item": "用户",
		},

		"tasksPerClosedReason": map[string]interface{}{
			"item": "原因",
			"graph": map[string]string{
				"xAxisName": "关闭原因",
			},
		},

		"finishedTasksPerDay": map[string]string{
			"item": "日期",
		},
		"tasksPerProject": map[string]string{
			"item": Lang[protocol.ZH_CN]["common"]["projectCommon"].(string),
		},
	}

	Lang[protocol.ZH_CN]["testcase"]["reviewList"] = []string{"否", `是`}

	Lang[protocol.ZH_CN]["testcase"]["priList"] = []string{"", `3`, `1`, `2`, `4`}

	Lang[protocol.ZH_CN]["testcase"]["typeList"] = map[string]string{
		"":            "",
		"feature":     "功能测试",
		"performance": "性能测试",
		"config":      "配置相关",
		"install":     "安装部署",
		"security":    "安全相关",
		"interface":   "接口测试",
		"other":       "其他",
	}

	Lang[protocol.ZH_CN]["testcase"]["stageList"] = map[string]string{
		"":           "",
		"unittest":   "单元测试阶段",
		"feature":    "功能测试阶段",
		"intergrate": "集成测试阶段",
		"system":     "系统测试阶段",
		"smoke":      "冒烟测试阶段",
		"bvt":        "版本验证阶段",
	}

	Lang[protocol.ZH_CN]["testcase"]["reviewResultList"] = map[string]string{
		"":        "",
		"pass":    "确认通过",
		"clarify": "继续完善",
	}

	Lang[protocol.ZH_CN]["testcase"]["groups"] = map[string]string{
		"":           "分组查看",
		"story":      "需求分组",
		"assignedTo": "指派分组",
	}

	Lang[protocol.ZH_CN]["testcase"]["statusList"] = map[string]string{
		"":            "",
		"wait":        "待评审",
		"normal":      "正常",
		"blocked":     "被阻塞",
		"investigate": "研究中",
	}

	Lang[protocol.ZH_CN]["testcase"]["resultList"] = map[string]string{
		"n/a":     "忽略",
		"pass":    "通过",
		"fail":    "失败",
		"blocked": "阻塞",
	}

	Lang[protocol.ZH_CN]["testcase"]["action"] = map[string]interface{}{
		"fromlib": map[string]string{
			`main`: `$date, 由 <strong>$actor</strong> 从用例库 <strong>$extra</strong>导入。`,
		},
		"reviewed": map[string]string{`main`: `$date, 由 <strong>$actor</strong> 记录评审结果，结果为 <strong>$extra</strong>。`, `extra`: `reviewResultList`},
	}

	Lang[protocol.ZH_CN]["testcase"]["featureBar"] = map[string]interface{}{
		"browse": map[string]string{
			"wait":        "待评审",
			"group":       "分组查看",
			"suite":       "套件",
			"zerocase":    "零用例需求",
			"all":         Lang[protocol.ZH_CN]["testcase"]["allCases"].(string),
			"needconfirm": Lang[protocol.ZH_CN]["testcase"]["needConfirm"].(string),
		},
	}
	Lang[protocol.ZH_CN]["testcase"]["featureBar"].(map[string]interface{})["groupcase"] = Lang[protocol.ZH_CN]["testcase"]["featureBar"].(map[string]interface{})["browse"]
	Lang[protocol.ZH_CN]["testsuite"]["authorList"] = map[string]string{
		"private": "私有",
		"public":  "公开",
	}

	Lang[protocol.ZH_CN]["testtask"]["statusList"] = map[string]string{
		"wait":    "未开始",
		"doing":   "进行中",
		"done":    "已完成",
		"blocked": "被阻塞",
	}

	Lang[protocol.ZH_CN]["testtask"]["priList"] = []string{"", `3`, `1`, `2`, `4`}

	Lang[protocol.ZH_CN]["testtask"]["placeholder"] = map[string]string{
		"begin": "开始日期",
		"end":   "结束日期",
	}

	Lang[protocol.ZH_CN]["testtask"]["mail"] = map[string]interface{}{
		"create": map[string]string{
			"title": "%s创建了版本 #%s:%s",
		},

		"edit": map[string]string{
			"title": "%s编辑了版本 #%s:%s",
		},

		"close": map[string]string{
			"title": "%s关闭了版本 #%s:%s",
		},
	}

	Lang[protocol.ZH_CN]["testtask"]["report"] = map[string]interface{}{
		"common": "报表",
		"select": "请选择报表类型",
		"create": "生成报表",

		"charts": map[string]string{
			"testTaskPerRunResult": "用例结果统计",
			"testTaskPerType":      "用例类型统计",
			"testTaskPerModule":    "用例模块统计",
			"testTaskPerRunner":    "用例执行人统计",
			"bugSeverityGroups":    "Bug严重级别分布",
			"bugStatusGroups":      "Bug状态分布",
			"bugOpenedByGroups":    "Bug创建者分布",
			"bugResolvedByGroups":  "Bug解决者分布",
			"bugResolutionGroups":  "Bug解决方案分布",
			"bugModuleGroups":      "Bug模块分布",
		},

		"options": map[string]string{
			"type":   "pie",
			"width":  "500",
			"height": "140",
		},
	}

	Lang[protocol.ZH_CN]["todo"]["reasonList"] = map[string]string{
		"story": "转需求",
		"task":  "转任务",
		"bug":   "转Bug",
		"done":  "完成",
	}

	Lang[protocol.ZH_CN]["todo"]["statusList"] = map[string]string{
		"wait":   "未开始",
		"doing":  "进行中",
		"done":   "已完成",
		"closed": "已关闭",
	}

	Lang[protocol.ZH_CN]["todo"]["priList"] = []string{"一般", `最高`, `较高`, `最低`, ``}

	Lang[protocol.ZH_CN]["todo"]["typeList"] = map[string]string{
		"custom": "自定义",
		"cycle":  "周期",
		"bug":    "Bug",
		"task":   Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "任务",
		"story":  Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "需求",
	}

	Lang[protocol.ZH_CN]["todo"]["periods"] = map[string]string{
		"all":      "所有待办",
		"thisYear": "本年",
		"future":   "待定",
		"before":   "未完",
		"cycle":    "周期",
	}

	Lang[protocol.ZH_CN]["todo"]["action"] = map[string]interface{}{
		"finished": map[string]string{
			`main`:  `$date, 由 <strong>$actor</strong> $extra。$appendLink`,
			`extra`: `reasonList`,
		},
		"marked": map[string]string{`main`: `$date, 由 <strong>$actor</strong> 标记为<strong>$extra</strong>。`, `extra`: `statusList`},
	}
	Lang[protocol.ZH_CN]["testtask"]["action"] = map[string]string{
		"testtaskopened":  "$date, 由 <strong>$actor</strong> 创建版本 <strong>$extra</strong>。\n",
		"testtaskstarted": "$date, 由 <strong>$actor</strong> 启动版本 <strong>$extra</strong>。\n",
		"testtaskclosed":  "$date, 由 <strong>$actor</strong> 完成版本 <strong>$extra</strong>。\n",
	}
	Lang[protocol.ZH_CN]["todo"]["dayNames"] = []string{`星期日`, `星期一`, `星期二`, `星期三`, `星期四`, `星期五`, `星期六`}
	Lang[protocol.ZH_CN]["tree"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["tree"]["product"] = "所属" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string)
	Lang[protocol.ZH_CN]["tree"]["projectDoc"] = Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "文档"
	Lang[protocol.ZH_CN]["tree"]["all"] = "所有模块"
	Lang[protocol.ZH_CN]["tree"]["short"] = "简称"
	Lang[protocol.ZH_CN]["tree"]["order"] = "排序"
	Lang[protocol.ZH_CN]["tree"]["owner"] = "负责人"
	Lang[protocol.ZH_CN]["tree"]["lineChild"] = "子产品线"
	Lang[protocol.ZH_CN]["tree"]["child"] = "子模块"
	Lang[protocol.ZH_CN]["tree"]["parentCate"] = "上级分类"
	Lang[protocol.ZH_CN]["tree"]["parent"] = "上级模块"
	Lang[protocol.ZH_CN]["tree"]["type"] = "类型"
	Lang[protocol.ZH_CN]["tree"]["path"] = "路径"
	Lang[protocol.ZH_CN]["tree"]["branch"] = "平台/分支"
	Lang[protocol.ZH_CN]["tree"]["root"] = "所属根"
	Lang[protocol.ZH_CN]["tree"]["cate"] = "分类名称"
	Lang[protocol.ZH_CN]["tree"]["line"] = "产品线名称"
	Lang[protocol.ZH_CN]["tree"]["name"] = "模块名称"
	Lang[protocol.ZH_CN]["tree"]["module"] = "模块"

	Lang[protocol.ZH_CN]["tree"]["successFixed"] = "成功修正数据！"
	Lang[protocol.ZH_CN]["tree"]["successSave"] = "成功保存"
	Lang[protocol.ZH_CN]["tree"]["confirmRoot"] = "模块的所属" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "修改，会关联修改该模块下的需求、Bug、用例的所属" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "，以及" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "和" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "的关联关系。该操作比较危险，请谨慎操作。是否确认修改？"
	Lang[protocol.ZH_CN]["tree"]["confirmDeleteLine"] = "该产品线及其子产品线都会被删除，您确定删除吗？"
	Lang[protocol.ZH_CN]["tree"]["confirmDelete"] = "该模块及其子模块都会被删除，您确定删除吗？"
	Lang[protocol.ZH_CN]["tree"]["addChild"] = "增加子模块"
	Lang[protocol.ZH_CN]["tree"]["sort"] = "排序"
	Lang[protocol.ZH_CN]["tree"]["dragAndSort"] = "拖放排序"
	Lang[protocol.ZH_CN]["tree"]["syncFromProduct"] = "复制模块"
	Lang[protocol.ZH_CN]["tree"]["manageTaskChild"] = "维护" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "子模块"
	Lang[protocol.ZH_CN]["tree"]["manageCaselibChild"] = "维护用例库子模块"
	Lang[protocol.ZH_CN]["tree"]["manageCaseChild"] = "维护用例子模块"
	Lang[protocol.ZH_CN]["tree"]["manageBugChild"] = "维护Bug子模块"
	Lang[protocol.ZH_CN]["tree"]["manageLineChild"] = "维护产品线"
	Lang[protocol.ZH_CN]["tree"]["manageStoryChild"] = "维护子模块"
	Lang[protocol.ZH_CN]["tree"]["manageChild"] = "维护子模块"
	Lang[protocol.ZH_CN]["tree"]["updateOrder"] = "更新排序"
	Lang[protocol.ZH_CN]["tree"]["manageCustomDoc"] = "维护文档库分类"
	Lang[protocol.ZH_CN]["tree"]["manageCaseLib"] = "维护用例库模块"
	Lang[protocol.ZH_CN]["tree"]["manageCase"] = "维护用例视图模块"
	Lang[protocol.ZH_CN]["tree"]["manageBug"] = "维护测试视图模块"
	Lang[protocol.ZH_CN]["tree"]["manageLine"] = "维护产品线"
	Lang[protocol.ZH_CN]["tree"]["manageProject"] = "维护" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "视图模块"
	Lang[protocol.ZH_CN]["tree"]["manageProduct"] = "维护" + Lang[protocol.ZH_CN]["common"]["productCommon"].(string) + "视图模块"
	Lang[protocol.ZH_CN]["tree"]["fix"] = "修正数据"
	Lang[protocol.ZH_CN]["tree"]["manage"] = "维护模块"
	Lang[protocol.ZH_CN]["tree"]["browseTask"] = "任务模块维护"
	Lang[protocol.ZH_CN]["tree"]["browse"] = "通用模块维护"
	Lang[protocol.ZH_CN]["tree"]["delete"] = "删除模块"
	Lang[protocol.ZH_CN]["tree"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["tree"]["common"] = "模块维护"

	Lang[protocol.ZH_CN]["trip"]["typeList"] = map[string]string{
		"trip":   "出差",
		"egress": "外出",
	}

	Lang[protocol.ZH_CN]["user"]["roleList"] = []protocol.HtmlKeyValueStr{
		{"", ""},
		{"dev", "研发"},
		{"qa", "测试"},
		{"pm", "项目经理"},
		{"po", "产品经理"},
		{"td", "研发主管"},
		{"pd", "产品主管"},
		{"qd", "测试主管"},
		{"top", "高层管理"},
		{"others", "其他"},
		{"market", "市场"},
		{"service", "客服"},
		{"operation", "运营"},
		{"support", "技术支持"},
		{"leading", "组长"},
		{"member", "组员"},
	}

	Lang[protocol.ZH_CN]["user"]["genderList"] = []protocol.HtmlKeyValueStr{
		{"0", "男"},
		{"1", "女"},
	}

	Lang[protocol.ZH_CN]["user"]["thirdPerson"] = map[string]string{
		"m": "他",
		"f": "她",
	}

	Lang[protocol.ZH_CN]["user"]["passwordStrengthList"] = []string{"<span style='color:red'>弱</span>", "<span style='color:#000'>中</span>", "<span style='color:green'>强</span>"}

	Lang[protocol.ZH_CN]["user"]["statusList"] = map[string]string{
		"active": "正常",
		"delete": "删除",
	}

	Lang[protocol.ZH_CN]["user"]["personalData"] = map[string]string{
		"createdTodo":  "创建的待办数",
		"createdStory": "创建的" + Lang[protocol.ZH_CN]["common"]["storyCommon"].(string) + "数",
		"finishedTask": "完成的任务数",
		"resolvedBug":  "解决的Bug数",
		"createdCase":  "创建的用例数",
	}

	Lang[protocol.ZH_CN]["user"]["keepLogin"] = []protocol.HtmlKeyValueStr{
		{"on", "保持登录"},
	}

	Lang[protocol.ZH_CN]["user"]["tpl"] = map[string]string{
		"type":    "类型",
		"title":   "模板名",
		"content": "内容",
		"public":  "是否公开",
	}

	Lang[protocol.ZH_CN]["usertpl"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["usertpl"]["title"] = "模板名称"

	Lang[protocol.ZH_CN]["user"]["placeholder"] = map[string]interface{}{
		"account":          "英文、数字和下划线的组合，三位以上",
		"password1":        "六位以上",
		"role":             "职位影响内容和用户列表的顺序。",
		"group":            "分组决定用户的权限列表。",
		"commiter":         "版本控制系统(subversion)中的账号",
		"verify":           "请输入您的系统登录密码",
		"passwordStrength": []string{"6位以上，包含大小写字母，数字。", "10位以上，包含大小写字母，数字，特殊字符。"},
	}

	Lang[protocol.ZH_CN]["user"]["error"] = map[string]string{
		"account":                "【ID %s】的用户名应该为：三位以上的英文、数字或下划线的组合",
		"accountDupl":            "【ID %s】的用户名已经存在",
		"realname":               "【ID %s】的真实姓名必须填写",
		"accountrealnameDupl":    "用户名或者真实姓名已经存在",
		"password":               "【ID %s】的密码必须为六位以上",
		"mail":                   "【ID %s】的邮箱地址不正确",
		"reserved":               "【ID %s】的用户名已被系统预留",
		"dangerPassword":         "【ID %s】的密码不能使用【%s】这些常用若口令。",
		"ErrPassword":            "验证失败，请检查您的系统登录密码是否正确",
		"originalPassword":       "原密码不正确",
		"repeat":                 "%s，因为用户名重复，不能添加！，请修改用户名后再添加",
		"illaccount":             "%s，因为用户名不合法，添加失败！，请修改用户名后再添加",
		"userLimit":              "人数已经达到授权的上限，不能从LDAP导入新用户！",
		"duplicated":             "重复关联账号",
		"role":                   "%s，职位不能为空。",
		"Week":                   "密码不能使用【%s】这些常用弱口令。",
		"weakPasswordlen":        "您的密码长度小于系统设定。最少要求%d位",
		"weakPasswordLowerUpper": "密码必须包含大小写",
		"weakPasswordSpecial":    "密码必须包含特殊字符",
		"weakPasswordNum":        "密码必须包含数字",
		"loginLocked":            "密码尝试次数太多，请联系管理员解锁，或%s分钟后重试。",
		"lockWarning":            "您还有%s次尝试机会。",
		"errorView":              "抱歉，您无权访问『<b>%s</b>』视图。请联系管理员获取权限。点击后退返回上页。",
		"errorDeny":              "抱歉，您无权访问『<b>%s</b>』模块的『<b>%s</b>』功能。请联系管理员获取权限。点击后退返回上页。",
		"NotFoundUserInfo":       "没有找到该用户信息",
		"passwordrule":           "密码应该符合规则，长度至少为六位。",
		"passwordsame":           "两次密码应当相等。",
		"ErrUpdate":              "更新失败，错误%v",
		"ErrCheckaccount":        "检查账号是否存在失败，错误%v",
		"UserAccountIsexist":     "该账号已存在，请更换其他账号",
		"loginFailed":            "登录失败 用户名为空",
	}

	Lang[protocol.ZH_CN]["user"]["contacts"] = map[string]string{
		"common":        "联系人",
		"listName":      "列表名称",
		"userList":      "用户列表",
		"manage":        "维护列表",
		"contactsList":  "已有列表",
		"selectedUsers": "选择用户",
		"selectList":    "选择列表",
		"createList":    "创建新列表",
		"noListYet":     "还没有创建任何列表，请先创建联系人列表。",
		"confirmDelete": "您确定要删除这个列表吗？",
		"or":            " 或者 ",
	}

	Lang[protocol.ZH_CN]["user"]["isFeedback"] = []string{"研发用户", `非研发用户`}

	Lang[protocol.ZH_CN]["user"]["notice"] = map[string]string{
		"checkbox": "没有勾选，则不导入！",
		"ldapoff":  "LDAP处于关闭状态。",
	}
	Lang[protocol.ZH_CN]["user"]["contactFieldList"] = map[string]string{
		"phone":    Lang[protocol.ZH_CN]["user"]["phone"].(string),
		"mobile":   Lang[protocol.ZH_CN]["user"]["Mobile"].(string),
		"qq":       Lang[protocol.ZH_CN]["user"]["QQ"].(string),
		"dingding": Lang[protocol.ZH_CN]["user"]["dingding"].(string),
		"weixin":   Lang[protocol.ZH_CN]["user"]["Weixin"].(string),
		"skype":    Lang[protocol.ZH_CN]["user"]["skype"].(string),
		"slack":    Lang[protocol.ZH_CN]["user"]["slack"].(string),
		"whatsapp": Lang[protocol.ZH_CN]["user"]["whatsapp"].(string),
	}

	Lang[protocol.ZH_CN]["dept"]["error"] = map[string]string{
		"hasSons":           "该部门有子部门，不能删除！",
		"hasUsers":          "该部门有职员，不能删除！",
		"ErrDeptIDType":     "参数错误，ID %s 不是数字",
		"ErrDeptInfo":       "获取部门信息失败，错误%v",
		"ErrOrderType":      "参数错误，ID %s 排序%s 不是数字",
		"ErrDeptInfoDeptID": "无法获取ID %v 的部门信息",
		"ErrUpdate":         "更新部门信息错误失败，错误%v",
		"ErrGetDeptUser":    "获取部门职员错误失败，错误%v",
		"ErrManager":        "无法获取该部门经理的信息，修改失败",
	}
	Lang[protocol.ZH_CN]["common"]["error"] = map[string]string{
		"ErrGetMsg": "初始化失败，请联系管理员，错误%v",
	}
	Lang[protocol.ZH_CN]["file"]["error"] = map[string]string{
		"ErrImgType": "上传图片失败，格式识别错误",
		"updateTmp":  "文件上传失败，请刷新再试",
	}
	Lang[protocol.ZH_CN]["error"]["checkTypeRequire"] = "不能为空"
	Lang[protocol.ZH_CN]["error"]["checkTypeInt"] = "必须为整数"
	Lang[protocol.ZH_CN]["error"]["checkTypeUserId"] = "必须为有效的用户"
	Lang[protocol.ZH_CN]["error"]["checkHtmlKeyValueStr"] = "该选项无效，请刷新重新选择"
	Lang[protocol.ZH_CN]["error"]["checkTypeDate"] = "日期格式不对，必须为2020-01-01"
	Lang[protocol.ZH_CN]["error"]["checkPositiveAndZero"] = "数值必须大于或等于0"
	Lang[protocol.ZH_CN]["error"]["checkPositive"] = "数值必须大于0"
	Lang[protocol.ZH_CN]["error"]["checkNegativeAndZero"] = "数值必须小于或等于0"
	Lang[protocol.ZH_CN]["error"]["checkNegative"] = "数值必须小于0"
	Lang[protocol.ZH_CN]["error"]["resultType"] = "远程服务器返回的结果不符合预期"
	Lang[protocol.ZH_CN]["error"]["assignedToNotFoundUser"] = "未找到指派人"

	Lang[protocol.ZH_CN]["product"]["error"] = map[string]string{
		"NotFound": "没有找到产品",
	}
	Lang[protocol.ZH_CN]["tree"]["error"] = map[string]string{
		"ModuleNameRepeat": "模块名已经存在！",
		"ModuleNotFound":   "没有找到该模块请返回首页重试",
	}
	Lang[protocol.ZH_CN]["my"]["error"] = map[string]string{
		"managecontactsEmptyListName": "列表名称不能为空",
		"managecontactsEmptyUsers":    "用户列表不能为空",
		"managecontactsErrorID":       "列表id错误，请刷新重试",
	}
	Lang[protocol.ZH_CN]["branch"]["error"] = map[string]string{
		"ErrDelete": "删除失败，错误%v",
	}
	Lang[protocol.ZH_CN]["productplan"]["error"] = map[string]string{
		"beginGeEnd":                           "开始时间不能大于结束时间",
		"errorNoEnd":                           "结束时间不能为空",
		"errorNoBegin":                         "开始时间不能为空",
		"errorNoTitle":                         "标题不能为空",
		"NotFoundProductPlanInfo":              "没有找到产品计划信息",
		"Err_ProjectProductPlanParentNotFound": "没有找到上一级产品计划信息",
	}
	Lang[protocol.ZH_CN]["project"]["error"] = map[string]string{
		"NotFound":                 "没有找到对应的" + Lang[protocol.ZH_CN]["common"]["projectCommon"].(string) + "信息",
		"beginTime":                "开始日期格式不正确",
		"endTime":                  "结束日期格式不正",
		"beginGeEnd":               "开始时间不能大于结束时间",
		"CreateNotFoundProduct":    "找不到关联产品，请刷新重试",
		"CreateNotFoundPlan":       "找不到关联计划，请刷新重试",
		"ProjectNameIsExist":       "项目名称或者代号已存在相同的，请修改重试",
		"daysErr":                  "项目可用工作日输入错误，输入范围0-32767",
		"manageMembersUserIsExist": "团队成员%s信息重复录入",
	}
	//Lang[protocol.ZH_CN]["common"]["moduleOrder"] = []string{"index", "my", "todo", "product", "story", "productplan", "release", "project", "task", "build", "qa", "bug", "testcase", "testtask", "testsuite", "testreport", "caselib", "doc", "report", "company", "dept", "group", "user", "admin", "extension", "custom", "editor", "convert", "action", "mail", "svn", "git", "search", "tree", "api", "file", "misc", "backup", "cron", "dev", "message"}
	Lang[protocol.ZH_CN]["common"]["moduleOrder"] = []string{"index", "my", "todo", "product", "story", "productplan", "release", "project", "task", "build", "qa", "bug", "testcase", "testtask", "testsuite", "testreport", "doc", "report", "company", "dept", "group", "user", "admin", "custom", "action", "mail", "search", "tree", "file", "backup", "cron", "attend", "holiday", "leave", "makeup", "overtime", "lieu", "trip"}
	Lang[protocol.ZH_CN]["resource"] = make(map[string]interface{})

	Lang[protocol.ZH_CN]["resource"]["index"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"notice", "notice"},
	}
	/* Index module. */

	Lang[protocol.ZH_CN]["index"]["methodOrder"] = []string{"index", "notice"}
	/* My module. */
	Lang[protocol.ZH_CN]["resource"]["my"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"todo", "todo"},
		{"calendar", "calendar"},
		{"task", "task"},
		{"bug", "bug"},
		{"testTask", "testTask"},
		{"testCase", "testCase"},
		{"story", "story"},
		{"project", "myProject"},
		{"profile", "profile"},
		{"dynamic", "dynamic"},
		{"editProfile", "editProfile"},
		{"changePassword", "changePassword"},
		{"unbind", "unbind"},
		{"manageContacts", "manageContacts"},
		{"deleteContacts", "deleteContacts"},
		{"score", "score"},
	}
	Lang[protocol.ZH_CN]["my"]["methodOrder"] = []string{"index", "todo", "task", "bug", "testTask", "testCase", "story", "project", "profile", "dynamic", "editProfile", "changePassword", "unbind", "manageContacts", "deleteContacts", "score"}
	/* Todo. */
	Lang[protocol.ZH_CN]["resource"]["todo"] = []protocol.HtmlKeyValueStr{
		{"create", "create"},
		{"createcycle", "createCycle"},
		{"batchCreate", "batchCreate"},
		{"edit", "edit"},
		{"batchEdit", "batchEdit"},
		{"view", "view"},
		{"delete", "delete"},
		{"export", "export"},
		{"finish", "finish"},
		{"batchFinish", "batchFinish"},
		{"import2Today", "import2Today"},
		{"assignTo", "assignTo"},
		{"activate", "activate"},
		{"close", "close"},
		{"batchClose", "batchClose"},
	}

	Lang[protocol.ZH_CN]["todo"]["methodOrder"] = []string{"create", "createCycle", "batchCreate", "edit", "batchEdit", "view", "delete", "export", "finish", "batchFinish", "import2Today", "assignTo", "activate", "close", "batchClose"}

	/* Product. */
	Lang[protocol.ZH_CN]["resource"]["product"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"browse", "browse"},
		{"create", "create"},
		{"view", "view"},
		{"edit", "edit"},
		{"batchEdit", "batchEdit"},
		{"delete", "delete"},
		{"roadmap", "roadmap"},
		{"doc", "doc"},
		{"dynamic", "dynamic"},
		{"project", "project"},
		{"close", "close"},
		{"updateOrder", "updateOrder"},
		{"all", "all"},
		{"build", "build"},
		{"export", "export"},
	}

	Lang[protocol.ZH_CN]["product"]["methodOrder"] = []string{"index", "browse", "create", "view", "edit", "batchEdit", "delete", "roadmap", "dynamic", "project", "close", "updateOrder", "all", "build", "export"}
	/* Branch. */
	Lang[protocol.ZH_CN]["resource"]["branch"] = []protocol.HtmlKeyValueStr{
		{"manage", "manage"},
		{"sort", "sort"},
		{"delete", "delete"},
	}

	Lang[protocol.ZH_CN]["branch"]["methodOrder"] = []string{
		"manage", "sort", "delete"}

	/* Story. */
	Lang[protocol.ZH_CN]["resource"]["story"] = []protocol.HtmlKeyValueStr{
		{"create", "create"},
		{"batchCreate", "batchCreate"},
		{"edit", "edit"},
		{"linkStory", "linkStory"},
		{"batchEdit", "batchEdit"},
		{"export", "export"},
		{"delete", "delete"},
		{"view", "view"},
		{"change", "lblChange"},
		{"review", "lblReview"},
		{"batchReview", "batchReview"},
		{"close", "lblClose"},
		{"batchClose", "batchClose"},
		{"activate", "lblActivate"},
		{"tasks", "tasks"},
		{"bugs", "bugs"},
		{"cases", "cases"},
		{"zeroCase", "zeroCase"},
		{"report", "reportChart"},
		{"batchChangePlan", "batchChangePlan"},
		{"batchChangeBranch", "batchChangeBranch"},
		{"batchChangeStage", "batchChangeStage"},
		{"batchAssignTo", "batchAssignTo"},
		{"batchChangeModule", "batchChangeModule"},
	}
	Lang[protocol.ZH_CN]["story"]["methodOrder"] = []string{
		"create", "batchCreate", "edit", "export", "delete", "view", "change", "review", "batchReview", "close", "batchClose", "batchChangePlan", "batchChangeStage", "batchAssignTo", "activate", "tasks", "bugs", "cases", "zeroCase", "report", "linkStory", "batchChangeBranch", "batchChangeModule"}
	/* Product plan. */
	Lang[protocol.ZH_CN]["resource"]["productplan"] = []protocol.HtmlKeyValueStr{
		{"browse", "browse"},
		{"create", "create"},
		{"edit", "edit"},
		{"delete", "delete"},
		{"view", "view"},
		{"linkStory", "linkStory"},
		{"unlinkStory", "unlinkStory"},
		{"batchUnlinkStory", "batchUnlinkStory"},
		{"linkBug", "linkBug"},
		{"unlinkBug", "unlinkBug"},
		{"batchUnlinkBug", "batchUnlinkBug"},
		{"batchEdit", "batchEdit"},
	}

	Lang[protocol.ZH_CN]["productplan"]["methodOrder"] = []string{
		"browse", "create", "edit", "delete", "view", "linkStory", "unlinkStory", "batchUnlinkStory", "linkBug", "unlinkBug", "batchUnlinkBug", "batchEdit"}
	/* Release. */
	Lang[protocol.ZH_CN]["resource"]["release"] = []protocol.HtmlKeyValueStr{
		{"browse", "browse"},
		{"create", "create"},
		{"edit", "edit"},
		{"delete", "delete"},
		{"view", "view"},
		{"export", "export"},
		{"linkStory", "linkStory"},
		{"unlinkStory", "unlinkStory"},
		{"batchUnlinkStory", "batchUnlinkStory"},
		{"linkBug", "linkBug"},
		{"unlinkBug", "unlinkBug"},
		{"batchUnlinkBug", "batchUnlinkBug"},
		{"changeStatus", "changeStatus"},
	}

	Lang[protocol.ZH_CN]["release"]["methodOrder"] = []string{
		"browse", "create", "edit", "delete", "view", "export", "linkStory", "unlinkStory", "batchUnlinkStory", "linkBug", "unlinkBug", "batchUnlinkBug", "changeStatus"}
	/* Project. */
	Lang[protocol.ZH_CN]["resource"]["project"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"view", "view"},
		{"browse", "browse"},
		{"create", "create"},
		{"edit", "edit"},
		{"batchedit", "batchEdit"},
		{"start", "start"},
		{"activate", "activate"},
		{"putoff", "putoff"},
		{"suspend", "suspend"},
		{"close", "close"},
		{"delete", "delete"},
		{"task", "task"},
		{"grouptask", "groupTask"},
		{"importtask", "importTask"},
		{"importplanstories", "importPlanStories"},
		{"importBug", "importBug"},
		{"story", "story"},
		{"build", "build"},
		{"testtask", "testtask"},
		{"bug", "bug"},
		{"burn", "burn"},
		{"computeBurn", "computeBurn"},
		{"fixFirst", "fixFirst"},
		{"burnData", "burnData"},
		{"team", "team"},
		{"doc", "doc"},
		{"dynamic", "dynamic"},
		{"manageProducts", "manageProducts"},
		//{"manageChilds","manageChilds"},
		{"manageMembers", "manageMembers"},
		{"unlinkMember", "unlinkMember"},
		{"linkStory", "linkStory"},
		{"showFile", "showFile"},
		{"unlinkStory", "unlinkStory"},
		{"batchUnlinkStory", "batchUnlinkStory"},
		{"updateOrder", "updateOrder"},
		{"kanban", "kanban"},
		{"printKanban", "printKanban"},
		{"tree", "tree"},
		{"treeTask", "treeOnlyTask"},
		{"treeStory", "treeOnlyStory"},
		{"all", "allProjects"},
		{"kanbanHideCols", "kanbanHideCols"},
		{"kanbanColsColor", "kanbanColsColor"},
		{"export", "export"},
		{"exportfile", "exportfile"},
		{"storyKanban", "storyKanban"},
		{"storySort", "storySort"},
		{"gantt", "ganttchart"},
	}

	Lang[protocol.ZH_CN]["project"]["methodOrder"] = []string{
		"index", "view", "browse", "create", "edit", "batchedit", "start", "activate", "putoff", "suspend", "close", "delete", "task", "grouptask", "importtask", "importplanstories", "importBug", "story", "build", "testtask", "bug", "burn", "computeBurn", "fixFirst", "burnData", "team", "dynamic", "manageProducts", "manageMembers", "unlinkMember", "linkStory", "unlinkStory", "batchUnlinkStory", "updateOrder", "kanban", "printKanban", "tree", "tree", "tree", "all", "kanbanHideCols", "kanbanColsColor", "export", "exportfile", "storyKanban", "storySort"}
	/* Task. */
	Lang[protocol.ZH_CN]["resource"]["task"] = []protocol.HtmlKeyValueStr{
		{"create", "create"},
		{"edit", "edit"},
		{"assignTo", "assign"},
		{"start", "start"},
		{"pause", "pause"},
		{"restart", "restart"},
		{"finish", "finish"},
		{"cancel", "cancel"},
		{"close", "close"},
		{"batchCreate", "batchCreate"},
		{"batchEdit", "batchEdit"},
		{"batchClose", "batchClose"},
		{"batchCancel", "batchCancel"},
		{"batchAssignTo", "batchAssignTo"},
		{"batchChangeModule", "batchChangeModule"},
		{"activate", "activate"},
		{"delete", "delete"},
		{"view", "view"},
		{"export", "export"},
		{"exportfile", "exportfile"},
		{"examine", "examine"},
		{"internalaudit", "internalaudit"},
		{"proofreading", "proofreading"},
		{"finishall", "finishall"},
		{"exportfinish", "exportfinish"},
		{"placeorder", "placeOrder"},

		{"batchexamine", "batchexamine"},
		{"batchproofreading", "batchproofreading"},
		{"batchexaminec", "batchexaminec"},
		{"batchproofreadingc", "batchproofreadingc"},

		{"confirmStoryChange", "confirmStoryChange"},
		{"recordEstimate", "recordEstimate"},
		{"editEstimate", "editEstimate"},
		{"deleteEstimate", "deleteEstimate"},
		{"report", "reportChart"},
	}

	Lang[protocol.ZH_CN]["task"]["methodOrder"] = []string{
		"create", "batchCreate", "batchEdit", "edit", "assignTo", "batchAssignTo", "start", "pause", "restart", "finish", "cancel", "close", "batchClose", "activate", "delete", "view", "export", "exportfile", "confirmStoryChange", "recordEstimate", "editEstimate", "deleteEstimate", "report", "batchChangeModule"}
	/* Build. */
	Lang[protocol.ZH_CN]["resource"]["build"] = []protocol.HtmlKeyValueStr{
		{"create", "create"},
		{"edit", "edit"},
		{"delete", "delete"},
		{"view", "view"},
		{"linkStory", "linkStory"},
		{"unlinkStory", "unlinkStory"},
		{"batchUnlinkStory", "batchUnlinkStory"},
		{"linkBug", "linkBug"},
		{"unlinkBug", "unlinkBug"},
		{"batchUnlinkBug", "batchUnlinkBug"},
	}

	Lang[protocol.ZH_CN]["build"]["methodOrder"] = []string{
		"create", "edit", "delete", "view", "linkStory", "unlinkStory", "batchUnlinkStory", "linkBug", "unlinkBug", "batchUnlinkBug"}
	/* QA. */
	Lang[protocol.ZH_CN]["resource"]["qa"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
	}

	Lang[protocol.ZH_CN]["qa"]["methodOrder"] = []string{
		"index"}
	/* Bug. */
	Lang[protocol.ZH_CN]["resource"]["bug"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"browse", "browse"},
		{"create", "create"},
		{"batchCreate", "batchCreate"},
		{"confirmBug", "confirmBug"},
		{"batchConfirm", "batchConfirm"},
		{"view", "view"},
		{"edit", "edit"},
		{"linkBugs", "linkBugs"},
		{"batchEdit", "batchEdit"},
		{"batchClose", "batchClose"},
		{"assignTo", "assignTo"},
		{"batchAssignTo", "batchAssignTo"},
		{"resolve", "resolve"},
		{"batchResolve", "batchResolve"},
		{"activate", "activate"},
		{"batchActivate", "batchActivate"},
		{"close", "close"},
		{"report", "reportChart"},
		{"export", "export"},
		{"confirmStoryChange", "confirmStoryChange"},
		{"delete", "delete"},
		{"saveTemplate", "saveTemplate"},
		{"deleteTemplate", "deleteTemplate"},
		{"setPublic", "setPublic"},
		{"batchChangeModule", "batchChangeModule"},
		{"batchChangeBranch", "batchChangeBranch"},
	}

	Lang[protocol.ZH_CN]["bug"]["methodOrder"] = []string{
		"index", "browse", "create", "batchCreate", "batchEdit", "confirmBug", "batchConfirm", "view", "edit", "assignTo", "batchAssignTo", "resolve", "batchResolve", "batchClose", "batchActivate", "activate", "close", "report", "export", "confirmStoryChange", "delete", "saveTemplate", "deleteTemplate", "setPublic", "linkBugs", "batchChangeModule", "batchChangeBranch"}
	/* Test case. */
	Lang[protocol.ZH_CN]["resource"]["testcase"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"browse", "browse"},
		{"groupCase", "groupCase"},
		{"create", "create"},
		{"batchCreate", "batchCreate"},
		{"createBug", "createBug"},
		{"view", "view"},
		{"edit", "edit"},
		{"linkCases", "linkCases"},
		{"batchEdit", "batchEdit"},
		{"delete", "delete"},
		{"batchDelete", "batchDelete"},
		{"export", "export"},
		{"exportTemplet", "exportTemplet"},
		{"import", "import"},
		{"showImport", "showImport"},
		{"confirmChange", "confirmChange"},
		{"confirmStoryChange", "confirmStoryChange"},
		{"batchChangeModule", "batchChangeModule"},
		{"batchChangeBranch", "batchChangeBranch"},
		{"bugs", "bugs"},
		{"review", "review"},
		{"batchReview", "batchReview"},
		{"importFromLib", "importFromLib"},
		{"batchCaseTypeChange", "batchCaseTypeChange"},
		{"batchConfirmStoryChange", "batchConfirmStoryChange"},
	}

	Lang[protocol.ZH_CN]["testcase"]["methodOrder"] = []string{
		"index", "browse", "groupCase", "create", "batchCreate", "createBug", "view", "edit", "delete", "export", "confirmChange", "confirmStoryChange", "batchEdit", "batchDelete", "batchChangeModule", "batchChangeBranch", "linkCases", "bugs", "review", "batchReview", "batchConfirmStoryChange", "importFromLib", "batchCaseTypeChange"}
	/* Test task. */
	Lang[protocol.ZH_CN]["resource"]["testtask"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"create", "create"},
		{"browse", "browse"},
		{"view", "view"},
		{"cases", "lblCases"},
		{"groupCase", "groupCase"},
		{"edit", "edit"},
		{"start", "start"},
		{"close", "close"},
		{"delete", "delete"},
		{"batchAssign", "batchAssign"},
		{"linkcase", "linkCase"},
		{"unlinkcase", "lblUnlinkCase"},
		{"batchUnlinkCases", "batchUnlinkCases"},
		{"runcase", "lblRunCase"},
		{"results", "lblResults"},
		{"results", "lblResults"},
		{"batchRun", "batchRun"},
		{"activate", "activate"},
		{"block", "block"},
		{"report", "reportChart"},
	}

	Lang[protocol.ZH_CN]["testtask"]["methodOrder"] = []string{
		"index", "create", "browse", "view", "cases", "groupCase", "edit", "start", "activate", "block", "close", "delete", "batchAssign", "linkcase", "unlinkcase", "runcase", "results", "batchUnlinkCases", "report"}

	Lang[protocol.ZH_CN]["resource"]["testreport"] = []protocol.HtmlKeyValueStr{
		{"browse", "browse"},
		{"create", "create"},
		{"view", "view"},
		{"delete", "delete"},
		{"edit", "edit"},
	}

	Lang[protocol.ZH_CN]["testreport"]["methodOrder"] = []string{
		"browse", "create", "view", "delete", "edit"}
	Lang[protocol.ZH_CN]["resource"]["testsuite"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"browse", "browse"},
		{"create", "create"},
		{"view", "view"},
		{"edit", "edit"},
		{"delete", "delete"},
		{"linkCase", "linkCase"},
		{"unlinkCase", "unlinkCase"},
		{"batchUnlinkCases", "batchUnlinkCases"},
	}

	Lang[protocol.ZH_CN]["testsuite"]["methodOrder"] = []string{
		"index", "browse", "create", "view", "edit", "delete", "linkCase", "unlinkCase", "batchUnlinkCases"}
	Lang[protocol.ZH_CN]["resource"]["caselib"] = []protocol.HtmlKeyValueStr{
		{"library", "library"},
		{"createLib", "createLib"},
		{"edit", "editLib"},
		{"createCase", "createCase"},
		{"libView", "libView"},
		{"batchCreateCase", "batchCreateCase"},
		{"exportTemplet", "exportTemplet"},
		{"import", "import"},
		{"showImport", "showImport"},
	}

	Lang[protocol.ZH_CN]["caselib"]["methodOrder"] = []string{
		"library", "createLib", "edit", "createCase", "batchCreateCase", "libView", "exportTemplet", "import", "showImport"}
	/* Doc. */
	Lang[protocol.ZH_CN]["resource"]["doc"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"browse", "browse"},
		{"createLib", "createLib"},
		{"editLib", "editLib"},
		{"deleteLib", "deleteLib"},
		{"create", "create"},
		{"view", "view"},
		{"edit", "edit"},
		{"delete", "delete"},
		{"allLibs", "allLibs"},
		{"objectLibs", "objectLibs"},
		{"showFiles", "showFiles"},
		{"sort", "sort"},
		{"collect", "collect"},
		//{"diff","diff"},
	}

	Lang[protocol.ZH_CN]["caselib"]["methodOrder"] = []string{
		"index", "browse", "createLib", "editLib", "deleteLib", "create", "view", "edit", "delete", "allLibs", "showFiles", "objectLibs", "sort", "collect"}
	//$lang->doc->methodOrder[55] = "diff";

	/* mail. */
	Lang[protocol.ZH_CN]["resource"]["mail"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"detect", "detect"},
		{"edit", "edit"},
		{"save", "save"},
		{"test", "test"},
		{"reset", "reset"},
		{"browse", "browse"},
		{"delete", "delete"},
		{"resend", "resend"},
		{"batchDelete", "batchDelete"},
		{"sendCloud", "sendCloud"},
		{"sendcloudUser", "sendcloudUser"},
	}

	Lang[protocol.ZH_CN]["mail"]["methodOrder"] = []string{
		"index", "detect", "edit", "save", "test", "reset", "browse", "delete", "batchDelete", "resend", "sendCloud", "sendcloudUser"}
	/* custom. */
	Lang[protocol.ZH_CN]["resource"]["custom"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"set", "set"},
		{"restore", "restore"},
		{"flow", "flow"},
		{"working", "working"},
		{"setPublic", "setPublic"},
	}

	Lang[protocol.ZH_CN]["custom"]["methodOrder"] = []string{
		"index", "set", "restore", "flow", "working", "setPublic"}
	Lang[protocol.ZH_CN]["resource"]["datatable"] = []protocol.HtmlKeyValueStr{
		{"setGlobal", "setGlobal"},
	}

	Lang[protocol.ZH_CN]["datatable"]["methodOrder"] = []string{
		"setGlobal"}
	/* Subversion. */
	Lang[protocol.ZH_CN]["resource"]["svn"] = []protocol.HtmlKeyValueStr{
		{"diff", "diff"},
		{"cat", "cat"},
		{"apiSync", "apiSync"},
	}

	Lang[protocol.ZH_CN]["svn"]["methodOrder"] = []string{
		"diff", "cat", "apiSync"}
	/* Git. */
	Lang[protocol.ZH_CN]["resource"]["git"] = []protocol.HtmlKeyValueStr{
		{"diff", "diff"},
		{"cat", "cat"},
		{"apiSync", "apiSync"},
	}

	Lang[protocol.ZH_CN]["git"]["methodOrder"] = []string{
		"diff", "cat", "apiSync"}
	/* Company. */
	Lang[protocol.ZH_CN]["resource"]["company"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"browse", "browse"},
		{"edit", "edit"},
		{"view", "view"},
		{"dynamic", "dynamic"},
	}

	Lang[protocol.ZH_CN]["company"]["methodOrder"] = []string{
		"index", "browse", "edit", "dynamic"}
	/* Department. */
	Lang[protocol.ZH_CN]["resource"]["dept"] = []protocol.HtmlKeyValueStr{
		{"browse", "browse"},
		{"updateOrder", "updateOrder"},
		{"manageChild", "manageChild"},
		{"edit", "edit"},
		{"delete", "delete"},
	}

	Lang[protocol.ZH_CN]["dept"]["methodOrder"] = []string{
		"browse", "updateOrder", "manageChild", "edit", "delete"}
	/* Group. */
	Lang[protocol.ZH_CN]["resource"]["group"] = []protocol.HtmlKeyValueStr{
		{"browse", "browse"},
		{"create", "create"},
		{"edit", "edit"},
		{"copy", "copy"},
		{"delete", "delete"},
		{"manageView", "manageView"},
		{"managePriv", "managePriv"},
		{"manageMember", "manageMember"},
	}

	Lang[protocol.ZH_CN]["group"]["methodOrder"] = []string{
		"browse", "create", "edit", "copy", "delete", "managePriv", "manageMember"}
	/* User. */
	Lang[protocol.ZH_CN]["resource"]["user"] = []protocol.HtmlKeyValueStr{
		{"create", "create"},
		{"batchCreate", "batchCreate"},
		{"view", "view"},
		{"edit", "edit"},
		{"unlock", "unlock"},
		{"delete", "delete"},
		{"todo", "todo"},
		{"story", "story"},
		{"task", "task"},
		{"bug", "bug"},
		{"testTask", "testTask"},
		{"testCase", "testCase"},
		{"project", "project"},
		{"dynamic", "dynamic"},
		{"profile", "profile"},
		{"batchEdit", "batchEdit"},
		{"unbind", "unbind"},
	}

	Lang[protocol.ZH_CN]["user"]["methodOrder"] = []string{
		"create", "batchCreate", "view", "edit", "unlock", "delete", "todo", "task", "bug", "project", "dynamic", "profile", "batchEdit", "unbind"}
	/* Tree. */
	Lang[protocol.ZH_CN]["resource"]["tree"] = []protocol.HtmlKeyValueStr{
		{"browse", "browse"},
		{"browseTask", "browseTask"},
		{"updateOrder", "updateOrder"},
		{"manageChild", "manageChild"},
		{"edit", "edit"},
		{"fix", "fix"},
		{"delete", "delete"},
	}

	Lang[protocol.ZH_CN]["tree"]["methodOrder"] = []string{
		"browse", "browseTask", "updateOrder", "manageChild", "edit", "delete"}
	/* Report. */
	Lang[protocol.ZH_CN]["resource"]["report"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"projectDeviation", "projectDeviation"},
		{"productSummary", "productSummary"},
		{"bugCreate", "bugCreate"},
		{"bugAssign", "bugAssign"},
		{"workload", "workload"},
	}

	Lang[protocol.ZH_CN]["report"]["methodOrder"] = []string{
		"index", "projectDeviation", "productSummary", "bugCreate", "workload"}
	/* Search. */
	Lang[protocol.ZH_CN]["resource"]["search"] = []protocol.HtmlKeyValueStr{
		{"buildForm", "buildForm"},
		{"buildQuery", "buildQuery"},
		{"saveQuery", "saveQuery"},
		{"deleteQuery", "deleteQuery"},
		{"select", "select"},
	}

	Lang[protocol.ZH_CN]["search"]["methodOrder"] = []string{
		"buildForm", "buildQuery", "saveQuery", "deleteQuery", "select"}
	/* Admin. */
	Lang[protocol.ZH_CN]["resource"]["admin"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"checkDB", "checkDB"},
		{"safe", "safeIndex"},
		{"checkWeak", "checkWeak"},
		{"sso", "sso"},
	}

	Lang[protocol.ZH_CN]["admin"]["methodOrder"] = []string{
		"index", "checkDB", "safeIndex", "checkWeak", "sso"}
	/* Extension. */
	Lang[protocol.ZH_CN]["resource"]["extension"] = []protocol.HtmlKeyValueStr{
		{"browse", "browse"},
		{"obtain", "obtain"},
		{"structure", "structure"},
		{"install", "install"},
		{"uninstall", "uninstall"},
		{"activate", "activate"},
		{"deactivate", "deactivate"},
		{"upload", "upload"},
		{"erase", "erase"},
		{"upgrade", "upgrade"},
	}
	Lang[protocol.ZH_CN]["extension"] = map[string]interface{}{
		"methodOrder": []string{"browse", "obtain", "structure", "install", "uninstall", "activate", "deactivate", "upload", "erase", "upgrade"},
	}

	/* Editor. */
	Lang[protocol.ZH_CN]["resource"]["editor"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"extend", "extend"},
		{"edit", "edit"},
		{"newPage", "newPage"},
		{"save", "save"},
		{"delete", "delete"},
	}

	Lang[protocol.ZH_CN]["extension"]["methodOrder"] = []string{
		"index", "extend", "edit", "newPage", "save", "delete"}
	/* Convert. */
	Lang[protocol.ZH_CN]["resource"]["convert"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"selectSource", "selectSource"},
		{"setConfig", "setConfig"},
		{"setBugfree", "setBugfree"},
		{"setRedmine", "setRedmine"},
		{"checkConfig", "checkConfig"},
		{"checkBugFree", "checkBugFree"},
		{"checkRedmine", "checkRedmine"},
		{"execute", "execute"},
		{"convertBugFree", "convertBugFree"},
		{"convertRedmine", "convertRedmine"},
	}

	Lang[protocol.ZH_CN]["convert"]["methodOrder"] = []string{
		"index", "selectSource", "setConfig", "setBugfree", "setRedmine", "checkConfig", "checkBugFree", "checkRedmine", "execute", "convertBugFree", "convertRedmine"}
	/* Others. */
	Lang[protocol.ZH_CN]["resource"]["api"] = []protocol.HtmlKeyValueStr{
		{"getModel", "getModel"},
		{"debug", "debug"},
		{"sql", "sql"},
	}

	Lang[protocol.ZH_CN]["api"]["methodOrder"] = []string{
		"getModel", "debug", "sql"}
	Lang[protocol.ZH_CN]["resource"]["file"] = []protocol.HtmlKeyValueStr{
		{"download", "download"},
		{"edit", "edit"},
		{"delete", "delete"},
		{"uploadImages", "uploadImages"},
		{"setPublic", "setPublic"},
	}

	Lang[protocol.ZH_CN]["file"]["methodOrder"] = []string{
		"download", "edit", "delete", "uploadImages", "setPublic"}
	/*Lang[protocol.ZH_CN]["resource"]["misc"] = []protocol.HtmlKeyValueStr{
		{"ping", "ping"},
	}
	Lang[protocol.ZH_CN]["misc"] = make(map[string]interface{})
	Lang[protocol.ZH_CN]["misc"]["methodOrder"] = []string{
		"ping"}*/
	Lang[protocol.ZH_CN]["resource"]["message"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"setting", "setting"},
	}
	Lang[protocol.ZH_CN]["message"]["methodOrder"] = []string{
		"index", "setting"}
	Lang[protocol.ZH_CN]["resource"]["action"] = []protocol.HtmlKeyValueStr{
		{"trash", "trash"},
		{"undelete", "undelete"},
		{"hideOne", "hideOne"},
		{"hideAll", "hideAll"},
		{"comment", "comment"},
		{"editComment", "editComment"},
	}

	Lang[protocol.ZH_CN]["action"]["methodOrder"] = []string{
		"trash", "undelete", "hideOne", "hideAll", "comment", "editComment"}
	Lang[protocol.ZH_CN]["resource"]["backup"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"backup", "backup"},
		{"restore", "restore"},
		{"change", "change"},
		{"delete", "delete"},
		{"setting", "setting"},
	}

	Lang[protocol.ZH_CN]["backup"]["methodOrder"] = []string{
		"index", "backup", "restore", "delete", "setting", "rmPHPHeader"}
	Lang[protocol.ZH_CN]["resource"]["cron"] = []protocol.HtmlKeyValueStr{
		{"index", "index"},
		{"turnon", "turnon"},
		{"create", "create"},
		{"edit", "edit"},
		{"toggle", "toggle"},
		{"delete", "delete"},
		{"openProcess", "openProcess"},
	}

	Lang[protocol.ZH_CN]["cron"]["methodOrder"] = []string{
		"index", "turnon", "create", "edit", "toggle", "delete", "openProcess"}
	Lang[protocol.ZH_CN]["resource"]["dev"] = []protocol.HtmlKeyValueStr{
		{"api", "api"},
		{"db", "db"},
	}
	Lang[protocol.ZH_CN]["dev"]["methodOrder"] = []string{
		"api", "db"}
	/* Every version of new privilege. */

	Lang[protocol.ZH_CN]["resource"]["attend"] = []protocol.HtmlKeyValueStr{
		{"department", "department"},
		{"company", "company"},
		{"browseReview", "browseReview"},
		{"review", "review"},
		{"export", "exportAction"},
		{"stat", "reportAction"},
		{"saveStat", "saveStatAction"},
		{"exportJachun", "exportStat"},
		{"detail", "detailAction"},
		{"exportDetail", "exportDetail"},
		{"settings", "settings"},
		{"personalSettings", "personalSettings"},
		{"setManager", "setManager"},

		{"personal", "personal"},
		{"edit", "editAction"},
		{"clockinout", "clockinout"},
	}

	Lang[protocol.ZH_CN]["attend"]["methodOrder"] = []string{
		"department", "company", "browseReview", "review", "export", "stat", "saveStat", "exportStat", "detail", "exportDetail", "settings", "personalSettings", "setManager",
		"personal", "edit"}
	/* Holiday */
	Lang[protocol.ZH_CN]["resource"]["holiday"] = []protocol.HtmlKeyValueStr{
		{"SetOT", "createAction"},
		{"SetRE", "editAction"},
		{"SetNONE", "deleteAction"},

		{"browse", "browse"},
	}

	Lang[protocol.ZH_CN]["holiday"]["methodOrder"] = []string{
		"browse", "create", "edit", "delete"}
	/* Leave */
	Lang[protocol.ZH_CN]["resource"]["leave"] = []protocol.HtmlKeyValueStr{
		{"browseReview", "browseReview"},
		{"company", "companyAction"},
		{"review", "reviewAction"},
		{"export", "exportAction"},
		{"setReviewer", "setReviewerAction"},

		{"personal", "personal"},
		{"create", "createAction"},
		{"edit", "editAction"},
		{"delete", "deleteAction"},
		{"view", "viewAction"},
		{"switchstatus", "switchstatus"},
		{"back", "backAction"},
		{"personalannual", "personalannual"},
		{"personalieu", "personalieu"},
		{"lieulist", "lieulist"},
		{"getDays", "getDays"},
	}
	Lang[protocol.ZH_CN]["leave"]["methodOrder"] = []string{
		"browseReview", "company", "review", "export", "setReviewer",
		"personal", "create", "getDays", "edit", "delete", "view", "switchstatus", "back"}
	/* Makeup */
	Lang[protocol.ZH_CN]["resource"]["makeup"] = []protocol.HtmlKeyValueStr{
		{"browseReview", "browseReview"},
		{"company", "companyAction"},
		{"review", "reviewAction"},
		{"export", "exportAction"},
		{"setReviewer", "setReviewerAction"},

		{"personal", "personal"},
		{"create", "createAction"},
		{"edit", "editAction"},
		{"view", "viewAction"},
		{"delete", "deleteAction"},
		{"switchstatus", "switchstatus"},
	}

	Lang[protocol.ZH_CN]["makeup"]["methodOrder"] = []string{
		"browseReview", "company", "review", "export", "setReviewer",
		"personal", "create", "edit", "view", "delete", "switchstatus"}
	/* Overtime */
	Lang[protocol.ZH_CN]["resource"]["overtime"] = []protocol.HtmlKeyValueStr{
		{"browseReview", "browseReview"},
		{"company", "companyAction"},
		{"review", "reviewAction"},
		{"export", "exportAction"},
		{"setReviewer", "setReviewerAction"},

		{"personal", "personal"},
		{"create", "createAction"},
		{"edit", "editAction"},
		{"view", "viewAction"},
		{"delete", "deleteAction"},
		{"switchstatus", "switchstatus"},
	}

	Lang[protocol.ZH_CN]["overtime"]["methodOrder"] = []string{
		"browseReview", "company", "review", "export", "setReviewer",
		"personal", "create", "edit", "view", "delete", "switchstatus"}
	/* Lieu */
	Lang[protocol.ZH_CN]["resource"]["lieu"] = []protocol.HtmlKeyValueStr{
		{"company", "companyAction"},
		{"browseReview", "browseReviewAction"},
		{"review", "reviewAction"},
		{"setReviewer", "setReviewerAction"},

		{"personal", "personal"},
		{"create", "createAction"},
		{"edit", "editAction"},
		{"delete", "deleteAction"},
		{"view", "viewAction"},
		{"switchstatus", "switchstatus"},
	}

	Lang[protocol.ZH_CN]["lieu"]["methodOrder"] = []string{
		"company", "browseReview", "review", "setReviewer",
		"personal", "create", "edit", "delete", "view", "switchstatus"}
	/* trip */
	Lang[protocol.ZH_CN]["resource"]["trip"] = []protocol.HtmlKeyValueStr{
		{"browseReview", "browseReview"},
		{"company", "company"},
		{"setReviewer", "setReviewerAction"},
		{"personal", "personal"},
		{"create", "createAction"},
		{"edit", "editAction"},
		{"delete", "deleteAction"},
		{"view", "viewAction"},
	}

	Lang[protocol.ZH_CN]["trip"]["methodOrder"] = []string{
		"browseReview", "company", "setReviewer",
		"personal", "create", "edit", "delete", "view"}
	Lang[protocol.ZH_CN]["makeup"] = make(map[string]interface{})

	Lang[protocol.ZH_CN]["makeup"]["common"] = "补班"
	Lang[protocol.ZH_CN]["makeup"]["browse"] = "补班列表"
	Lang[protocol.ZH_CN]["makeup"]["create"] = "申请补班"
	Lang[protocol.ZH_CN]["makeup"]["edit"] = "编辑"
	Lang[protocol.ZH_CN]["makeup"]["view"] = "详情"
	Lang[protocol.ZH_CN]["makeup"]["delete"] = "删除"
	Lang[protocol.ZH_CN]["makeup"]["review"] = "审核"
	Lang[protocol.ZH_CN]["makeup"]["cancel"] = "撤销"
	Lang[protocol.ZH_CN]["makeup"]["commit"] = "提交"
	Lang[protocol.ZH_CN]["makeup"]["export"] = "导出补班记录"

	Lang[protocol.ZH_CN]["makeup"]["personal"] = "我的补班"
	Lang[protocol.ZH_CN]["makeup"]["browseReview"] = "审核列表"
	Lang[protocol.ZH_CN]["makeup"]["company"] = "所有补班"
	Lang[protocol.ZH_CN]["makeup"]["setReviewer"] = "补班设置"
	Lang[protocol.ZH_CN]["makeup"]["batchReview"] = "批量审核"
	Lang[protocol.ZH_CN]["makeup"]["batchPass"] = "批量通过"

	Lang[protocol.ZH_CN]["makeup"]["id"] = "编号"
	Lang[protocol.ZH_CN]["makeup"]["year"] = "年"
	Lang[protocol.ZH_CN]["makeup"]["begin"] = "开始"
	Lang[protocol.ZH_CN]["makeup"]["end"] = "结束"
	Lang[protocol.ZH_CN]["makeup"]["start"] = "开始时间"
	Lang[protocol.ZH_CN]["makeup"]["finish"] = "结束时间"
	Lang[protocol.ZH_CN]["makeup"]["hours"] = "总时长"
	Lang[protocol.ZH_CN]["makeup"]["leave"] = "请假记录"
	Lang[protocol.ZH_CN]["makeup"]["type"] = "类型"
	Lang[protocol.ZH_CN]["makeup"]["desc"] = "事由"
	Lang[protocol.ZH_CN]["makeup"]["status"] = "状态"
	Lang[protocol.ZH_CN]["makeup"]["createdBy"] = "申请者"
	Lang[protocol.ZH_CN]["makeup"]["createdDate"] = "申请时间"
	Lang[protocol.ZH_CN]["makeup"]["reviewedBy"] = "审核者"
	Lang[protocol.ZH_CN]["makeup"]["reviewedDate"] = "审核时间"
	Lang[protocol.ZH_CN]["makeup"]["date"] = "日期"
	Lang[protocol.ZH_CN]["makeup"]["time"] = "时间"
	Lang[protocol.ZH_CN]["makeup"]["rejectReason"] = "拒绝原因"

	Lang[protocol.ZH_CN]["makeup"]["typeList"] = []protocol.HtmlKeyValueStr{
		{"compensate", "补班"},
	}
	Lang[protocol.ZH_CN]["makeup"]["statusList"] = []protocol.HtmlKeyValueStr{
		{"draft", "草稿"},
		{"wait", "等待审核"},
		{"doing", "审核中"},
		{"pass", "通过"},
		{"reject", "拒绝"},
	}

	Lang[protocol.ZH_CN]["makeup"]["notExist"] = "记录不存在"
	Lang[protocol.ZH_CN]["makeup"]["denied"] = "信息访问受限"
	Lang[protocol.ZH_CN]["makeup"]["unique"] = "%s 已经存在补班记录"
	Lang[protocol.ZH_CN]["makeup"]["sameMonth"] = "不支持跨月份补班"
	Lang[protocol.ZH_CN]["makeup"]["wrongEnd"] = "结束时间应该大于开始时间"
	Lang[protocol.ZH_CN]["makeup"]["nodata"] = "没有选择数据"
	Lang[protocol.ZH_CN]["makeup"]["reviewSuccess"] = "审核成功"
	Lang[protocol.ZH_CN]["makeup"]["confirmReview"] = map[string]string{
		"pass":   "您确定要执行通过操作吗？",
		"reject": "您确定要执行拒绝操作吗？",
	}

	Lang[protocol.ZH_CN]["makeup"]["hoursTip"] = "小时"
	Lang[protocol.ZH_CN]["makeup"]["reviewStatusList"] = []protocol.HtmlKeyValueStr{

		{"pass", "通过"},
		{"reject", "拒绝"},
	}

	Lang[protocol.ZH_CN]["makeup"]["switchstatus"] = "提交或撤销"
	Lang[protocol.ZH_CN]["makeup"]["companyAction"] = "所有补班"
	Lang[protocol.ZH_CN]["makeup"]["reviewAction"] = "审核补班"
	Lang[protocol.ZH_CN]["makeup"]["exportAction"] = "导出补班"
	Lang[protocol.ZH_CN]["makeup"]["createAction"] = "申请补班"
	Lang[protocol.ZH_CN]["makeup"]["editAction"] = "编辑补班"
	Lang[protocol.ZH_CN]["makeup"]["deleteAction"] = "删除补班"
	Lang[protocol.ZH_CN]["makeup"]["viewAction"] = "补班详情"
	Lang[protocol.ZH_CN]["makeup"]["setReviewerAction"] = "补班设置"
	//lang套lang的放在最后
	LangZH_CN_ADD()
}
func LangZH_CN_ADD() {
	Lang[protocol.ZH_CN]["action"]["descValue"] = map[string]map[string]interface{}{ //2021-9-13，增加 "旧值""新值" 的翻译，参数 objtype，field
		"task": map[string]interface{}{
			"status": Lang[protocol.ZH_CN]["task"]["statusList"],
		},
	}
}
