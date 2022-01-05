package protocol

import (
	"libraries"
)

var cmdMapFunc = map[int32]func(*libraries.MsgBuffer) MSG_DATA{
	-79972603: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_GetTaskEstimateByTaskId(buf))},
	642796289: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_updateTmp(buf))},
	1302197764: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Company_cache(buf))},
	484676613: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_branch_getByProducts_result(buf))},
	-910528507: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_activate(buf))},
	882194947: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_Action_AddHistory(buf))},
	-83610363: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_getBurn(buf))},
	1624184325: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_getTaskTreeModules(buf))},
	1133272325: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_getBurn_result(buf))},
	-1798962171: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_build_getById(buf))},
	-756565247: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_StartTicker(buf))},
	-1780358652: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Group_cache(buf))},
	243028229: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getPairsForStory_result(buf))},
	1808135429: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_getProjectTasks_result(buf))},
	1735336196: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_GET_LoginSalt(buf))},
	-1549135356: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_team_getByIds_result(buf))},
	1422058245: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan(buf))},
	1153682945: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_GET_MsgUserId_result(buf))},
	-1525303291: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getForProducts_result(buf))},
	-581285116: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_delete_result(buf))},
	134391300: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_team_updateByWhere(buf))},
	-281089787: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_StoryStage(buf))},
	-1926345727: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_regServer_result(buf))},
	-1940757759: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_BeginTransaction_result(buf))},
	-810685181: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_Action_GetByWhereMap_result(buf))},
	2045178629: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getById(buf))},
	-1642266109: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_Action_transformActions_result(buf))},
	-576769789: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_Action_set_read(buf))},
	679671812: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Pairs(buf))},
	2045370628: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getParents_result(buf))},
	-2122256635: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_getStoryTaskCounts(buf))},
	-1661534459: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getPairs(buf))},
	-357570299: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_getById(buf))},
	557953793: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_QueryErr(buf))},
	-198630396: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_INFO_updateByID(buf))},
	-1400001276: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getGlobalContacts_result(buf))},
	-803107580: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_ContactList(buf))},
	-686336507: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_insert_result(buf))},
	1253762819: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_Action_GetByID_result(buf))},
	598613509: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_delete(buf))},
	1019057158: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_bug(buf))},
	-1315762175: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_GET_Msgno_result(buf))},
	-453068031: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_getByWhere_result(buf))},
	-1058626559: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_download_byIds_result(buf))},
	814416901: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_start(buf))},
	-730127098: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_CASE_getTaskCasePairs_result(buf))},
	-1348354556: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getGlobalContacts(buf))},
	2127377925: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_editBranch(buf))},
	1356456453: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_stroy_create_result(buf))},
	341559300: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_team_getMemberPairsByTypeRoot_result(buf))},
	-504988411: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_insert(buf))},
	528481285: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_getStories_result(buf))},
	-1626340351: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_Transaction_Check(buf))},
	488167427: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_Action_GetByID(buf))},
	-1137193468: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_insertUpdateContactList_result(buf))},
	155656196: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getContactListByUid_result(buf))},
	362981380: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_team_info(buf))},
	1478139909: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_getStoriesMapBySql_result(buf))},
	-1951638010: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_testsuite_getById_result(buf))},
	-1489369599: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_WINDOW_UPDATE(buf))},
	1095125505: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_CACHE_GETPATH(buf))},
	1895097345: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_CACHE_DelPath(buf))},
	-1996657916: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getParents(buf))},
	1829698565: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_cache(buf))},
	-1715112955: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_suspend(buf))},
	-160677883: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getPairs_result(buf))},
	808894213: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_getById(buf))},
	-1790767871: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_PONG(buf))},
	-1568529407: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_DeleteByID(buf))},
	-1123948283: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getList(buf))},
	939602945: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_CACHE_GET_result(buf))},
	819345413: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_build(buf))},
	-694422779: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_getProjectStoryPairs_result(buf))},
	-1677176059: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_UpdateTaskEstimate(buf))},
	-241416699: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_TaskEstimate(buf))},
	-1067127803: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_GetTaskEstimateByTaskId_result(buf))},
	-630575871: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_ResetWindow(buf))},
	775618820: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Userquery_info(buf))},
	-2052331515: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_getPairsByIds(buf))},
	829781510: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_testsuite_info(buf))},
	1618683396: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getContactLists(buf))},
	1984920837: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_getProjectStoryPairs(buf))},
	-1124872187: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_close(buf))},
	-360695547: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_putoff(buf))},
	-111920891: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_examine(buf))},
	-1722911999: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_CACHE_GETPATH_result(buf))},
	-1137975037: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_Action_Create(buf))},
	1806051844: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Group_getPairs(buf))},
	-1620079099: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_getListByWhereMap_result(buf))},
	1017210373: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_getProductStories(buf))},
	1057561345: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_uploadTmp(buf))},
	1937124355: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_transformActions_info(buf))},
	1713425925: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story(buf))},
	-1330589692: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_user_getUserqueryByWhere(buf))},
	-2065363451: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_getParents(buf))},
	421344517: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getLast_result(buf))},
	1228419589: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_insertUpdate(buf))},
	-568022779: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_create(buf))},
	-136790780: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getDataStructure(buf))},
	1131689476: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_updateUserView(buf))},
	-1892740092: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getContactLists_result(buf))},
	-1678157051: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_UpdateTaskEstimate_result(buf))},
	1967635718: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_bug_updateMapById(buf))},
	671857158: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_bug_getById(buf))},
	2097391617: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_BeginTransaction(buf))},
	-2003846907: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_manageChild_result(buf))},
	-582706938: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_testtask_getById_result(buf))},
	-446270460: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_team_getByTypeRoot_result(buf))},
	627510533: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_delete(buf))},
	1974311942: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_product_deleteBranch_check(buf))},
	-102044159: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_regServer(buf))},
	1875485187: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_Action_transformActions(buf))},
	1170940932: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_insertUpdateContactList(buf))},
	76296454: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_bug_getCountByWhere(buf))},
	640438276: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getContactListById(buf))},
	-1262905851: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_getLinePairs_result(buf))},
	-1288650235: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_start(buf))},
	-671290623: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_updateMapByWhere(buf))},
	-1650651898: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_product_deleteBranch_result(buf))},
	-1036028410: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_bug_getPairs_result(buf))},
	905195012: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_team_getByIds(buf))},
	-1827817723: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_insertUpdate_result(buf))},
	577900805: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_batchGetStoryStage(buf))},
	1779498757: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_proofreading(buf))},
	743980545: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_download_byIds(buf))},
	769103364: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_team_getByTypeUid_result(buf))},
	-815775483: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_branch_getByProducts(buf))},
	-334914812: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getDeptUserPairs(buf))},
	-1615845883: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_release(buf))},
	232008193: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_getByObject(buf))},
	-1103486716: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_delete(buf))},
	-48489211: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_getProjectTasks(buf))},
	-1896457211: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_create_result(buf))},
	-820050940: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Group_getPairs_result(buf))},
	-1195990011: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getById_result(buf))},
	1909900293: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_getPairsByIds(buf))},
	468566788: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_team_getByTypeUid(buf))},
	-1670634235: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_delete(buf))},
	298479622: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_buf_getCount_result(buf))},
	244356101: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_getPairsByIds_result(buf))},
	1583876353: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_GET_Msgno(buf))},
	226334977: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_RangeDown_result(buf))},
	2007179013: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_getStoriesMapBySql(buf))},
	908792837: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_getPairsByIds_result(buf))},
	-999177978: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_testtask_getById(buf))},
	-2057389055: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_upload_result(buf))},
	-1484540159: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_getByID_result(buf))},
	-792225787: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_release_getById_result(buf))},
	1223823365: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_Burn_info(buf))},
	1700618245: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_linkStory(buf))},
	1061688837: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_close(buf))},
	-2134257914: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_Testtask_info(buf))},
	1920226305: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_Transaction_Commit(buf))},
	-1108380155: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_getLinePairs(buf))},
	-553153019: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_cache(buf))},
	1110878977: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_upload(buf))},
	1184964097: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_GET_MsgUserId(buf))},
	-225686266: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_CASE_getTaskCasePairs(buf))},
	762193157: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_getPairsByIds_result(buf))},
	-950315259: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_getPairsByIds_result(buf))},
	1857790721: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_getByWhere(buf))},
	751862020: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_GET_LoginSalt_result(buf))},
	-242916347: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_build_getById_result(buf))},
	-1425965820: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_team_getByTypeRoot(buf))},
	633955845: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_assignTo(buf))},
	-2126260731: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_placeOrder(buf))},
	523943941: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_statRelatedData_result(buf))},
	1861930501: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_branch_getPairsByIds(buf))},
	566302981: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_finish_result(buf))},
	-1161866493: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_Action_GetByWhereMap(buf))},
	390890500: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getCompanyUsers(buf))},
	-1861023995: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_cache(buf))},
	614527750: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_bug_getCount(buf))},
	1720057093: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_create_result(buf))},
	-1868328699: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_create(buf))},
	-2101229563: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getForProducts(buf))},
	-867223292: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getDeptUserPairs_result(buf))},
	1504611846: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_bug_getCountByWhere_result(buf))},
	-2057750522: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_bug_getById_result(buf))},
	870918148: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_team_addByList(buf))},
	1709400580: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_user_getUserqueryByWhere_result(buf))},
	-261057276: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_team_getMemberPairsByTypeRoot(buf))},
	619873029: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_getStoryTaskCounts_result(buf))},
	-1871273215: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_getByID(buf))},
	-1371742204: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_getDataStructure_result(buf))},
	-1261518588: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getPairs_result(buf))},
	-1139812859: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_update(buf))},
	-114983163: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_getProjectTasksByWhere(buf))},
	-679761915: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_branch_getPairsByIds_result(buf))},
	636129793: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_RangeDown(buf))},
	346215685: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_getById_result(buf))},
	-1948818683: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_getPairsByIds(buf))},
	295929604: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_CheckAccount(buf))},
	-1549754875: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_start_result(buf))},
	-1563464187: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_activate(buf))},
	382968325: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_getProductStories_result(buf))},
	-670540539: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getPairsForStory(buf))},
	779733509: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_getById(buf))},
	1500885253: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_getPairsByIds(buf))},
	181040389: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getList_result(buf))},
	1679097605: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_getById_result(buf))},
	842619909: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_getPlanStories(buf))},
	1809855749: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_delete(buf))},
	267184385: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_CACHE_DEL(buf))},
	-876112636: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_INFO_cache(buf))},
	-942650620: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_config_save(buf))},
	-1004505085: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_History(buf))},
	368212996: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getContactListById_result(buf))},
	-143122426: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_testsuite_getById(buf))},
	2106113029: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_getTaskTreeModules_result(buf))},
	1984652549: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_getById_result(buf))},
	665022465: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_getByObject_result(buf))},
	953841924: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_CheckAccount_result(buf))},
	727801093: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_manageChild(buf))},
	-1093696765: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_Action(buf))},
	-2069873403: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_internalaudit(buf))},
	-1848741631: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_PING(buf))},
	-1144133115: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_updateList(buf))},
	-426602747: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_stroy_create(buf))},
	-1441671419: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_getPairs_result(buf))},
	516345601: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_CACHE_GET(buf))},
	338636804: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getPairs(buf))},
	-1290618875: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_productplan_getLast(buf))},
	3743749: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_batchGetStoryStage_result(buf))},
	2020131589: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_pause(buf))},
	1231270917: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_cancel(buf))},
	2009659137: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_Transaction_RollBack(buf))},
	2081476612: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getCompanyUsers_result(buf))},
	40064261: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_tree_getParents_result(buf))},
	1712503301: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_release_getById(buf))},
	-1587037691: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_getPairs(buf))},
	-1248575483: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_getListByWhereMap(buf))},
	470640644: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_update(buf))},
	1176677381: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_deleteBranch(buf))},
	-1911489786: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_TEST_bug_getPairs(buf))},
	99206145: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_HOST_CACHE_SET(buf))},
	-2090728188: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_getContactListByUid(buf))},
	439123205: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_TASK(buf))},
	894530565: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_branch_info(buf))},
	1630616325: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_project_statRelatedData(buf))},
	1409978369: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_FILE_edit(buf))},
	-681858812: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_Dept_cache(buf))},
	-1254780411: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_story_getPlanStories_result(buf))},
	1034823429: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_deleteBranch_result(buf))},
	-1288845053: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_LOG_Action_Create_result(buf))},
	-101573884: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_CheckPasswd_result(buf))},
	1365829125: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_product_getStories(buf))},
	-2006312700: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_USER_CheckPasswd(buf))},
	1069324805: func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_MSG_PROJECT_task_finish(buf))},
}

var CmdToName = map[int32]string{
	-820050940: `MSG_USER_Group_getPairs_result`,
	-1195990011: `MSG_PROJECT_productplan_getById_result`,
	1909900293: `MSG_PROJECT_tree_getPairsByIds`,
	-1896457211: `MSG_PROJECT_task_create_result`,
	468566788: `MSG_USER_team_getByTypeUid`,
	-1670634235: `MSG_PROJECT_productplan_delete`,
	298479622: `MSG_TEST_buf_getCount_result`,
	1583876353: `MSG_HOST_GET_Msgno`,
	226334977: `MSG_FILE_RangeDown_result`,
	2007179013: `MSG_PROJECT_product_getStoriesMapBySql`,
	244356101: `MSG_PROJECT_project_getPairsByIds_result`,
	-2057389055: `MSG_FILE_upload_result`,
	-1484540159: `MSG_FILE_getByID_result`,
	-792225787: `MSG_PROJECT_release_getById_result`,
	908792837: `MSG_PROJECT_product_getPairsByIds_result`,
	-999177978: `MSG_TEST_testtask_getById`,
	-2134257914: `MSG_TEST_Testtask_info`,
	1920226305: `MSG_HOST_Transaction_Commit`,
	-1108380155: `MSG_PROJECT_tree_getLinePairs`,
	-553153019: `MSG_PROJECT_product_cache`,
	1223823365: `MSG_PROJECT_project_Burn_info`,
	1700618245: `MSG_PROJECT_project_linkStory`,
	1061688837: `MSG_PROJECT_project_close`,
	1110878977: `MSG_FILE_upload`,
	1184964097: `MSG_HOST_GET_MsgUserId`,
	-225686266: `MSG_TEST_CASE_getTaskCasePairs`,
	1857790721: `MSG_FILE_getByWhere`,
	751862020: `MSG_USER_GET_LoginSalt_result`,
	-242916347: `MSG_PROJECT_build_getById_result`,
	762193157: `MSG_PROJECT_story_getPairsByIds_result`,
	-950315259: `MSG_PROJECT_tree_getPairsByIds_result`,
	-1425965820: `MSG_USER_team_getByTypeRoot`,
	633955845: `MSG_PROJECT_task_assignTo`,
	-2126260731: `MSG_PROJECT_task_placeOrder`,
	-1161866493: `MSG_LOG_Action_GetByWhereMap`,
	390890500: `MSG_USER_getCompanyUsers`,
	-1861023995: `MSG_PROJECT_tree_cache`,
	523943941: `MSG_PROJECT_project_statRelatedData_result`,
	1861930501: `MSG_PROJECT_branch_getPairsByIds`,
	566302981: `MSG_PROJECT_task_finish_result`,
	1720057093: `MSG_PROJECT_project_create_result`,
	-1868328699: `MSG_PROJECT_task_create`,
	-2101229563: `MSG_PROJECT_productplan_getForProducts`,
	614527750: `MSG_TEST_bug_getCount`,
	-867223292: `MSG_USER_getDeptUserPairs_result`,
	1504611846: `MSG_TEST_bug_getCountByWhere_result`,
	-2057750522: `MSG_TEST_bug_getById_result`,
	870918148: `MSG_USER_team_addByList`,
	1709400580: `MSG_USER_user_getUserqueryByWhere_result`,
	-1871273215: `MSG_FILE_getByID`,
	-1371742204: `MSG_USER_Dept_getDataStructure_result`,
	-1261518588: `MSG_USER_getPairs_result`,
	-261057276: `MSG_USER_team_getMemberPairsByTypeRoot`,
	619873029: `MSG_PROJECT_task_getStoryTaskCounts_result`,
	-1139812859: `MSG_PROJECT_product_update`,
	-114983163: `MSG_PROJECT_project_getProjectTasksByWhere`,
	636129793: `MSG_FILE_RangeDown`,
	346215685: `MSG_PROJECT_story_getById_result`,
	-1948818683: `MSG_PROJECT_product_getPairsByIds`,
	-679761915: `MSG_PROJECT_branch_getPairsByIds_result`,
	295929604: `MSG_USER_CheckAccount`,
	-1549754875: `MSG_PROJECT_task_start_result`,
	-670540539: `MSG_PROJECT_productplan_getPairsForStory`,
	779733509: `MSG_PROJECT_project_getById`,
	1500885253: `MSG_PROJECT_story_getPairsByIds`,
	-1563464187: `MSG_PROJECT_project_activate`,
	382968325: `MSG_PROJECT_story_getProductStories_result`,
	1809855749: `MSG_PROJECT_project_delete`,
	267184385: `MSG_HOST_CACHE_DEL`,
	-876112636: `MSG_USER_INFO_cache`,
	-942650620: `MSG_USER_config_save`,
	181040389: `MSG_PROJECT_productplan_getList_result`,
	1679097605: `MSG_PROJECT_project_getById_result`,
	842619909: `MSG_PROJECT_story_getPlanStories`,
	-1004505085: `MSG_LOG_History`,
	368212996: `MSG_USER_getContactListById_result`,
	-143122426: `MSG_TEST_testsuite_getById`,
	665022465: `MSG_FILE_getByObject_result`,
	953841924: `MSG_USER_CheckAccount_result`,
	727801093: `MSG_PROJECT_tree_manageChild`,
	2106113029: `MSG_PROJECT_tree_getTaskTreeModules_result`,
	1984652549: `MSG_PROJECT_task_getById_result`,
	-1093696765: `MSG_LOG_Action`,
	-2069873403: `MSG_PROJECT_task_internalaudit`,
	-1848741631: `MSG_HOST_PING`,
	-1144133115: `MSG_PROJECT_tree_updateList`,
	516345601: `MSG_HOST_CACHE_GET`,
	338636804: `MSG_USER_getPairs`,
	-1290618875: `MSG_PROJECT_productplan_getLast`,
	-426602747: `MSG_PROJECT_stroy_create`,
	-1441671419: `MSG_PROJECT_task_getPairs_result`,
	2009659137: `MSG_HOST_Transaction_RollBack`,
	2081476612: `MSG_USER_getCompanyUsers_result`,
	40064261: `MSG_PROJECT_tree_getParents_result`,
	3743749: `MSG_PROJECT_story_batchGetStoryStage_result`,
	2020131589: `MSG_PROJECT_task_pause`,
	1231270917: `MSG_PROJECT_task_cancel`,
	1712503301: `MSG_PROJECT_release_getById`,
	-1587037691: `MSG_PROJECT_task_getPairs`,
	-1248575483: `MSG_PROJECT_task_getListByWhereMap`,
	470640644: `MSG_USER_Dept_update`,
	1176677381: `MSG_PROJECT_product_deleteBranch`,
	99206145: `MSG_HOST_CACHE_SET`,
	-2090728188: `MSG_USER_getContactListByUid`,
	439123205: `MSG_PROJECT_TASK`,
	-1911489786: `MSG_TEST_bug_getPairs`,
	894530565: `MSG_PROJECT_branch_info`,
	1409978369: `MSG_FILE_edit`,
	-681858812: `MSG_USER_Dept_cache`,
	-1254780411: `MSG_PROJECT_story_getPlanStories_result`,
	1630616325: `MSG_PROJECT_project_statRelatedData`,
	-1288845053: `MSG_LOG_Action_Create_result`,
	-101573884: `MSG_USER_CheckPasswd_result`,
	1365829125: `MSG_PROJECT_product_getStories`,
	1034823429: `MSG_PROJECT_product_deleteBranch_result`,
	-2006312700: `MSG_USER_CheckPasswd`,
	1069324805: `MSG_PROJECT_task_finish`,
	642796289: `MSG_FILE_updateTmp`,
	1302197764: `MSG_USER_Company_cache`,
	484676613: `MSG_PROJECT_branch_getByProducts_result`,
	-79972603: `MSG_PROJECT_task_GetTaskEstimateByTaskId`,
	882194947: `MSG_LOG_Action_AddHistory`,
	-83610363: `MSG_PROJECT_project_getBurn`,
	1624184325: `MSG_PROJECT_tree_getTaskTreeModules`,
	-910528507: `MSG_PROJECT_task_activate`,
	1133272325: `MSG_PROJECT_project_getBurn_result`,
	-756565247: `MSG_HOST_StartTicker`,
	-1780358652: `MSG_USER_Group_cache`,
	243028229: `MSG_PROJECT_productplan_getPairsForStory_result`,
	-1798962171: `MSG_PROJECT_build_getById`,
	1735336196: `MSG_USER_GET_LoginSalt`,
	-1549135356: `MSG_USER_team_getByIds_result`,
	1422058245: `MSG_PROJECT_productplan`,
	1808135429: `MSG_PROJECT_project_getProjectTasks_result`,
	1153682945: `MSG_HOST_GET_MsgUserId_result`,
	-1525303291: `MSG_PROJECT_productplan_getForProducts_result`,
	-1926345727: `MSG_HOST_regServer_result`,
	-1940757759: `MSG_HOST_BeginTransaction_result`,
	-810685181: `MSG_LOG_Action_GetByWhereMap_result`,
	-581285116: `MSG_USER_Dept_delete_result`,
	134391300: `MSG_USER_team_updateByWhere`,
	-281089787: `MSG_PROJECT_StoryStage`,
	-1642266109: `MSG_LOG_Action_transformActions_result`,
	-576769789: `MSG_LOG_Action_set_read`,
	679671812: `MSG_USER_Pairs`,
	2045178629: `MSG_PROJECT_productplan_getById`,
	2045370628: `MSG_USER_Dept_getParents_result`,
	-2122256635: `MSG_PROJECT_task_getStoryTaskCounts`,
	557953793: `MSG_HOST_QueryErr`,
	-198630396: `MSG_USER_INFO_updateByID`,
	-1400001276: `MSG_USER_getGlobalContacts_result`,
	-1661534459: `MSG_PROJECT_productplan_getPairs`,
	-357570299: `MSG_PROJECT_story_getById`,
	-803107580: `MSG_USER_ContactList`,
	-686336507: `MSG_PROJECT_product_insert_result`,
	-1315762175: `MSG_HOST_GET_Msgno_result`,
	-453068031: `MSG_FILE_getByWhere_result`,
	-1058626559: `MSG_FILE_download_byIds_result`,
	1253762819: `MSG_LOG_Action_GetByID_result`,
	598613509: `MSG_PROJECT_tree_delete`,
	1019057158: `MSG_TEST_bug`,
	-1348354556: `MSG_USER_getGlobalContacts`,
	2127377925: `MSG_PROJECT_product_editBranch`,
	1356456453: `MSG_PROJECT_stroy_create_result`,
	814416901: `MSG_PROJECT_task_start`,
	-730127098: `MSG_TEST_CASE_getTaskCasePairs_result`,
	-1626340351: `MSG_HOST_Transaction_Check`,
	488167427: `MSG_LOG_Action_GetByID`,
	-1137193468: `MSG_USER_insertUpdateContactList_result`,
	341559300: `MSG_USER_team_getMemberPairsByTypeRoot_result`,
	-504988411: `MSG_PROJECT_product_insert`,
	528481285: `MSG_PROJECT_product_getStories_result`,
	-1951638010: `MSG_TEST_testsuite_getById_result`,
	-1489369599: `MSG_HOST_WINDOW_UPDATE`,
	1095125505: `MSG_HOST_CACHE_GETPATH`,
	1895097345: `MSG_HOST_CACHE_DelPath`,
	155656196: `MSG_USER_getContactListByUid_result`,
	362981380: `MSG_USER_team_info`,
	1478139909: `MSG_PROJECT_product_getStoriesMapBySql_result`,
	-1996657916: `MSG_USER_Dept_getParents`,
	1829698565: `MSG_PROJECT_project_cache`,
	-1715112955: `MSG_PROJECT_project_suspend`,
	-1790767871: `MSG_HOST_PONG`,
	-1568529407: `MSG_FILE_DeleteByID`,
	-1123948283: `MSG_PROJECT_productplan_getList`,
	-160677883: `MSG_PROJECT_productplan_getPairs_result`,
	808894213: `MSG_PROJECT_task_getById`,
	939602945: `MSG_HOST_CACHE_GET_result`,
	819345413: `MSG_PROJECT_build`,
	-694422779: `MSG_PROJECT_story_getProjectStoryPairs_result`,
	-1677176059: `MSG_PROJECT_task_UpdateTaskEstimate`,
	-630575871: `MSG_HOST_ResetWindow`,
	775618820: `MSG_USER_Userquery_info`,
	-2052331515: `MSG_PROJECT_project_getPairsByIds`,
	-241416699: `MSG_PROJECT_TaskEstimate`,
	-1067127803: `MSG_PROJECT_task_GetTaskEstimateByTaskId_result`,
	1618683396: `MSG_USER_getContactLists`,
	1984920837: `MSG_PROJECT_story_getProjectStoryPairs`,
	-1124872187: `MSG_PROJECT_task_close`,
	829781510: `MSG_TEST_testsuite_info`,
	-1722911999: `MSG_HOST_CACHE_GETPATH_result`,
	-1137975037: `MSG_LOG_Action_Create`,
	1806051844: `MSG_USER_Group_getPairs`,
	-360695547: `MSG_PROJECT_project_putoff`,
	-111920891: `MSG_PROJECT_task_examine`,
	1057561345: `MSG_FILE_uploadTmp`,
	1937124355: `MSG_LOG_transformActions_info`,
	1713425925: `MSG_PROJECT_story`,
	-1620079099: `MSG_PROJECT_task_getListByWhereMap_result`,
	1017210373: `MSG_PROJECT_story_getProductStories`,
	1228419589: `MSG_PROJECT_productplan_insertUpdate`,
	-568022779: `MSG_PROJECT_project_create`,
	-136790780: `MSG_USER_Dept_getDataStructure`,
	1131689476: `MSG_USER_updateUserView`,
	-1892740092: `MSG_USER_getContactLists_result`,
	-1330589692: `MSG_USER_user_getUserqueryByWhere`,
	-2065363451: `MSG_PROJECT_tree_getParents`,
	421344517: `MSG_PROJECT_productplan_getLast_result`,
	-1678157051: `MSG_PROJECT_task_UpdateTaskEstimate_result`,
	1967635718: `MSG_TEST_bug_updateMapById`,
	671857158: `MSG_TEST_bug_getById`,
	2097391617: `MSG_HOST_BeginTransaction`,
	-2003846907: `MSG_PROJECT_tree_manageChild_result`,
	-582706938: `MSG_TEST_testtask_getById_result`,
	-102044159: `MSG_HOST_regServer`,
	1875485187: `MSG_LOG_Action_transformActions`,
	1170940932: `MSG_USER_insertUpdateContactList`,
	-446270460: `MSG_USER_team_getByTypeRoot_result`,
	627510533: `MSG_PROJECT_task_delete`,
	1974311942: `MSG_TEST_product_deleteBranch_check`,
	640438276: `MSG_USER_getContactListById`,
	-1262905851: `MSG_PROJECT_tree_getLinePairs_result`,
	-1288650235: `MSG_PROJECT_project_start`,
	76296454: `MSG_TEST_bug_getCountByWhere`,
	-671290623: `MSG_FILE_updateMapByWhere`,
	-1650651898: `MSG_TEST_product_deleteBranch_result`,
	905195012: `MSG_USER_team_getByIds`,
	-1827817723: `MSG_PROJECT_productplan_insertUpdate_result`,
	577900805: `MSG_PROJECT_story_batchGetStoryStage`,
	-1036028410: `MSG_TEST_bug_getPairs_result`,
	743980545: `MSG_FILE_download_byIds`,
	769103364: `MSG_USER_team_getByTypeUid_result`,
	-815775483: `MSG_PROJECT_branch_getByProducts`,
	1779498757: `MSG_PROJECT_task_proofreading`,
	-334914812: `MSG_USER_getDeptUserPairs`,
	-1615845883: `MSG_PROJECT_release`,
	232008193: `MSG_FILE_getByObject`,
	-1103486716: `MSG_USER_Dept_delete`,
	-48489211: `MSG_PROJECT_project_getProjectTasks`,
}