package protocol

import "errors"

type ErrCode int16

const (
	TokenOk                                   ErrCode = 2
	Success                                   ErrCode = 1  //成功
	Fail                                      ErrCode = 0  //操作失败
	Err_db                                    ErrCode = -1 //数据库操作失败
	Err_token                                 ErrCode = -2 //token错误
	Err_msg                                   ErrCode = -3
	LoginFail                                 ErrCode = -4
	Err_Password                              ErrCode = -5 //密码错误
	Err_DeptDeleteHasSons                     ErrCode = -6
	Err_DeptDeletehasUsers                    ErrCode = -7
	Err_DeptNotFound                          ErrCode = -8
	Err_DeptNotFoundID                        ErrCode = -9
	Err_UserInfoNotFound                      ErrCode = -10
	Err_UserAccountIsexist                    ErrCode = -11
	Err_TreeRepeatName                        ErrCode = -12
	Err_ProjectBranchCanNotDelete_PROJECT     ErrCode = -13
	Err_ProjectBranchCanNotDelete_STORY       ErrCode = -14
	Err_ProjectBranchCanNotDelete_MODULE      ErrCode = -15
	Err_ProjectBranchCanNotDelete_PRODUCTPLAN ErrCode = -16
	Err_ProjectBranchCanNotDelete_BUG         ErrCode = -17
	Err_ProjectBranchCanNotDelete_CASE        ErrCode = -18
	Err_ProjectBranchCanNotDelete_RELEASE     ErrCode = -19
	Err_ProjectBranchCanNotDelete_BUILD       ErrCode = -20
	Err_ProjectProductPlanNotFound            ErrCode = -21
	Err_ProjectProductPlanParentNotFound      ErrCode = -22
	Err_FileNotFount                          ErrCode = -23
	Err_ProjectStoryTitleExists               ErrCode = -24
	Err_ProjectNotFound                       ErrCode = -25
	Err_ProjectNameIsExist                    ErrCode = -26
	Err_ProjectStoryNotFount                  ErrCode = -27
)

var errCodeMap = map[ErrCode]string{
	Err_Password:                              "ErrPassword",
	Err_DeptDeleteHasSons:                     "hasSons",
	Err_DeptDeletehasUsers:                    "hasUsers",
	Err_DeptNotFound:                          "ErrDeptInfo",
	Err_DeptNotFoundID:                        "ErrDeptInfoDeptID",
	Err_UserInfoNotFound:                      "NotFoundUserInfo",
	Err_UserAccountIsexist:                    "UserAccountIsexist",
	Err_TreeRepeatName:                        "ModuleNameRepeat",
	Err_ProjectBranchCanNotDelete_PROJECT:     "BranchCanNotDeletePROJECTHasData",
	Err_ProjectBranchCanNotDelete_STORY:       "BranchCanNotDeleteSTORYHasData",
	Err_ProjectBranchCanNotDelete_MODULE:      "BranchCanNotDeleteMODULEHasData",
	Err_ProjectBranchCanNotDelete_PRODUCTPLAN: "BranchCanNotDeletePRODUCTPLANHasData",
	Err_ProjectBranchCanNotDelete_BUG:         "BranchCanNotDeleteBUGHasData",
	Err_ProjectBranchCanNotDelete_CASE:        "BranchCanNotDeleteCASEHasData",
	Err_ProjectBranchCanNotDelete_RELEASE:     "BranchCanNotDeleteRELEASEHasData",
	Err_ProjectBranchCanNotDelete_BUILD:       "BranchCanNotDeleteBUILDHasData",
	Err_ProjectProductPlanNotFound:            "NotFoundProductPlanInfo",
	Err_ProjectProductPlanParentNotFound:      "NotFoundProductPlanParent",
	Err_FileNotFount:                          "FileNotFount",
	Err_ProjectNotFound:                       "NotFountProject",
	Err_ProjectNameIsExist:                    "ProjectNameIsExist",
	Err_ProjectStoryNotFount:                  "StoryNotFount",
}

func (err ErrCode) String() string {
	return errCodeMap[err]
}
func (err ErrCode) Err() error {
	return errors.New(err.String())
}
