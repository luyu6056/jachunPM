package protocol

import (
	"libraries"
)

var cmdMapFunc = map[int32]func(*libraries.MsgBuffer) MSG_DATA{
	845234432: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_StartTicker(buf))},
	1835280896: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_WINDOW_UPDATE(buf))},
	-1094191872: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_GET(buf))},
	1735336196: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_GET_LoginSalt(buf))},
	-1423609088: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_regServer_result(buf))},
	-1961335296: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_GETPATH(buf))},
	-26660096: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_GET_Msgno(buf))},
	-1484540160: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_getByID_result(buf))},
	-1261518588: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getPairs_result(buf))},
	-1751634176: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_GET_Msgno_result(buf))},
	-1371742204: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getDataStructure_result(buf))},
	-1103486716: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_delete(buf))},
	-2041209088: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_PING(buf))},
	-1568529408: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_DeleteByID(buf))},
	-681858812: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_cache(buf))},
	2045370628: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getParents_result(buf))},
	390890500: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getCompanyUsers(buf))},
	1131689476: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_updateUserView(buf))},
	751862020: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_GET_LoginSalt_result(buf))},
	-334914812: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getDeptUserPairs(buf))},
	295929604: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_CheckAccount(buf))},
	-2099715328: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_PONG(buf))},
	679671812: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Pairs(buf))},
	-867223292: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getDeptUserPairs_result(buf))},
	-553153019: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_cache(buf))},
	-1163212800: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_DelPath(buf))},
	-581285116: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_delete_result(buf))},
	338636804: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getPairs(buf))},
	-1108380155: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_getLinePairs(buf))},
	-1979298816: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_GETPATH_result(buf))},
	714684672: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_QueryErr(buf))},
	-1777243392: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_updateByIDMap(buf))},
	-1996657916: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getParents(buf))},
	2081476612: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getCompanyUsers_result(buf))},
	-198630396: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_INFO_updateByID(buf))},
	953841924: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_CheckAccount_result(buf))},
	-686336507: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_insert_result(buf))},
	1110878976: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_upload(buf))},
	1302197764: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Company_cache(buf))},
	470640644: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_update(buf))},
	-1871273216: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_getByID(buf))},
	-1780358652: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Group_cache(buf))},
	-1262905851: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_getLinePairs_result(buf))},
	988600064: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_ResetWindow(buf))},
	-2057389056: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_upload_result(buf))},
	1508627456: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_regServer(buf))},
	-1511457280: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_SET(buf))},
	-2006312700: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_CheckPasswd(buf))},
	503974656: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_GET_result(buf))},
	-1343368448: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_COMMON_CACHE_DEL(buf))},
	-876112636: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_INFO_cache(buf))},
	-101573884: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_CheckPasswd_result(buf))},
	-136790780: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getDataStructure(buf))},
	-504988411: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_insert(buf))},
}

var CmdToName = map[int32]string{
	1508627456: `MSG_COMMON_regServer`,
	-1511457280: `MSG_COMMON_CACHE_SET`,
	-2006312700: `MSG_USER_CheckPasswd`,
	503974656: `MSG_COMMON_CACHE_GET_result`,
	-1343368448: `MSG_COMMON_CACHE_DEL`,
	-876112636: `MSG_USER_INFO_cache`,
	-101573884: `MSG_USER_CheckPasswd_result`,
	-136790780: `MSG_USER_Dept_getDataStructure`,
	-504988411: `MSG_PROJECT_product_insert`,
	845234432: `MSG_COMMON_StartTicker`,
	1835280896: `MSG_COMMON_WINDOW_UPDATE`,
	-1094191872: `MSG_COMMON_CACHE_GET`,
	1735336196: `MSG_USER_GET_LoginSalt`,
	-1423609088: `MSG_COMMON_regServer_result`,
	-1961335296: `MSG_COMMON_CACHE_GETPATH`,
	-26660096: `MSG_COMMON_GET_Msgno`,
	-1484540160: `MSG_FILE_getByID_result`,
	-1261518588: `MSG_USER_getPairs_result`,
	-1751634176: `MSG_COMMON_GET_Msgno_result`,
	-1371742204: `MSG_USER_Dept_getDataStructure_result`,
	-1103486716: `MSG_USER_Dept_delete`,
	-2041209088: `MSG_COMMON_PING`,
	-1568529408: `MSG_FILE_DeleteByID`,
	-681858812: `MSG_USER_Dept_cache`,
	2045370628: `MSG_USER_Dept_getParents_result`,
	390890500: `MSG_USER_getCompanyUsers`,
	1131689476: `MSG_USER_updateUserView`,
	751862020: `MSG_USER_GET_LoginSalt_result`,
	-334914812: `MSG_USER_getDeptUserPairs`,
	295929604: `MSG_USER_CheckAccount`,
	-2099715328: `MSG_COMMON_PONG`,
	679671812: `MSG_USER_Pairs`,
	-867223292: `MSG_USER_getDeptUserPairs_result`,
	-553153019: `MSG_PROJECT_product_cache`,
	-1163212800: `MSG_COMMON_CACHE_DelPath`,
	-581285116: `MSG_USER_Dept_delete_result`,
	338636804: `MSG_USER_getPairs`,
	-1108380155: `MSG_PROJECT_tree_getLinePairs`,
	-1979298816: `MSG_COMMON_CACHE_GETPATH_result`,
	714684672: `MSG_COMMON_QueryErr`,
	-1777243392: `MSG_FILE_updateByIDMap`,
	-1996657916: `MSG_USER_Dept_getParents`,
	2081476612: `MSG_USER_getCompanyUsers_result`,
	-198630396: `MSG_USER_INFO_updateByID`,
	953841924: `MSG_USER_CheckAccount_result`,
	-686336507: `MSG_PROJECT_product_insert_result`,
	1110878976: `MSG_FILE_upload`,
	1302197764: `MSG_USER_Company_cache`,
	470640644: `MSG_USER_Dept_update`,
	-1871273216: `MSG_FILE_getByID`,
	-1780358652: `MSG_USER_Group_cache`,
	-1262905851: `MSG_PROJECT_tree_getLinePairs_result`,
	988600064: `MSG_COMMON_ResetWindow`,
	-2057389056: `MSG_FILE_upload_result`,
}