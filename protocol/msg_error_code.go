package protocol

type ErrCode int16

const (
	TokenOk                ErrCode = 2
	Success                ErrCode = 1  //成功
	Fail                   ErrCode = 0  //操作失败
	Err_db                 ErrCode = -1 //数据库操作失败
	Err_token              ErrCode = -2 //token错误
	Err_msg                ErrCode = -3
	LoginFail              ErrCode = -4
	Err_Password           ErrCode = -5 //密码错误
	Err_DeptDeleteHasSons  ErrCode = -6
	Err_DeptDeletehasUsers ErrCode = -7
	Err_DeptNotFound       ErrCode = -8
	Err_DeptNotFoundID     ErrCode = -9
	Err_UserInfoNotFound   ErrCode = -10
	Err_UserAccountIsexist ErrCode = -19
)

func (err ErrCode) String() string {
	return map[ErrCode]string{
		Err_Password:           "ErrPassword",
		Err_DeptDeleteHasSons:  "hasSons",
		Err_DeptDeletehasUsers: "hasUsers",
		Err_DeptNotFound:       "ErrDeptInfo",
		Err_DeptNotFoundID:     "ErrDeptInfoDeptID",
		Err_UserInfoNotFound:   "NotFoundUserInfo",
		Err_UserAccountIsexist: "UserAccountIsexist",
	}[err]
}
