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
	Err_DeptNotFount       ErrCode = -8
	Err_DeptNotFountID     ErrCode = -9
)

func (err ErrCode) String() string {
	return map[ErrCode]string{
		Err_Password:           "ErrPassword",
		Err_DeptDeleteHasSons:  "hasSons",
		Err_DeptDeletehasUsers: "hasUsers",
		Err_DeptNotFount:       "ErrDeptInfo",
		Err_DeptNotFountID:     "ErrDeptInfoDeptID",
	}[err]
}
