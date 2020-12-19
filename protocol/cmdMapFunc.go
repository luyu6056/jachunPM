package protocol

import (
	"libraries"
)

var cmdMapFunc = map[int32]func(*libraries.MsgBuffer) MSG_DATA{
	-876112636: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_INFO_cache(buf))},
	-2099715328: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_PONG(buf))},
	714684672: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_QueryErr(buf))},
	-2006312700: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_CheckPasswd(buf))},
	2045370628: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getParents_result(buf))},
	-136790780: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getDataStructure(buf))},
	-1343368448: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_DEL(buf))},
	751862020: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_GET_LoginSalt_result(buf))},
	-101573884: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_CheckPasswd_result(buf))},
	-1996657916: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getParents(buf))},
	390890500: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getCompanyUsers(buf))},
	-1961335296: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_GETPATH(buf))},
	-1979298816: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_GETPATH_result(buf))},
	-26660096: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_GET_Msgno(buf))},
	988600064: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_ResetWindow(buf))},
	-581285116: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_delete_result(buf))},
	-334914812: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getDeptUserPairs(buf))},
	503974656: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_GET_result(buf))},
	-1163212800: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_DelPath(buf))},
	-1751634176: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_GET_Msgno_result(buf))},
	1735336196: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_GET_LoginSalt(buf))},
	-1103486716: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_delete(buf))},
	2081476612: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getCompanyUsers_result(buf))},
	-2041209088: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_PING(buf))},
	-1511457280: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_SET(buf))},
	-681858812: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_cache(buf))},
	-867223292: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getDeptUserPairs_result(buf))},
	1508627456: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_regServer(buf))},
	1302197764: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Company_cache(buf))},
	845234432: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_StartTicker(buf))},
	1835280896: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_WINDOW_UPDATE(buf))},
	-1094191872: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_GET(buf))},
	-1371742204: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getDataStructure_result(buf))},
	470640644: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_update(buf))},
	679671812: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Pairs(buf))},
	-1423609088: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_regServer_result(buf))},
}

var CmdToName = map[int32]string{
	1508627456: `MSG_COMMON_regServer`,
	-1511457280: `MSG_COMMON_CACHE_SET`,
	-681858812: `MSG_USER_Dept_cache`,
	-867223292: `MSG_USER_getDeptUserPairs_result`,
	845234432: `MSG_COMMON_StartTicker`,
	1302197764: `MSG_USER_Company_cache`,
	-1423609088: `MSG_COMMON_regServer_result`,
	1835280896: `MSG_COMMON_WINDOW_UPDATE`,
	-1094191872: `MSG_COMMON_CACHE_GET`,
	-1371742204: `MSG_USER_Dept_getDataStructure_result`,
	470640644: `MSG_USER_Dept_update`,
	679671812: `MSG_USER_Pairs`,
	-2099715328: `MSG_COMMON_PONG`,
	-876112636: `MSG_USER_INFO_cache`,
	-1343368448: `MSG_COMMON_CACHE_DEL`,
	714684672: `MSG_COMMON_QueryErr`,
	-2006312700: `MSG_USER_CheckPasswd`,
	2045370628: `MSG_USER_Dept_getParents_result`,
	-136790780: `MSG_USER_Dept_getDataStructure`,
	-1961335296: `MSG_COMMON_CACHE_GETPATH`,
	751862020: `MSG_USER_GET_LoginSalt_result`,
	-101573884: `MSG_USER_CheckPasswd_result`,
	-1996657916: `MSG_USER_Dept_getParents`,
	390890500: `MSG_USER_getCompanyUsers`,
	503974656: `MSG_COMMON_CACHE_GET_result`,
	-1979298816: `MSG_COMMON_CACHE_GETPATH_result`,
	-26660096: `MSG_COMMON_GET_Msgno`,
	988600064: `MSG_COMMON_ResetWindow`,
	-581285116: `MSG_USER_Dept_delete_result`,
	-334914812: `MSG_USER_getDeptUserPairs`,
	-2041209088: `MSG_COMMON_PING`,
	-1163212800: `MSG_COMMON_CACHE_DelPath`,
	-1751634176: `MSG_COMMON_GET_Msgno_result`,
	1735336196: `MSG_USER_GET_LoginSalt`,
	-1103486716: `MSG_USER_Dept_delete`,
	2081476612: `MSG_USER_getCompanyUsers_result`,
}