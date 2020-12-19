package main

type MSG_USER_GET_LoginSalt struct {
	QueryID uint32
	Name    string
}
type MSG_USER_GET_LoginSalt_result struct {
	QueryResultID uint32
	Salt          string
}
type MSG_USER_INFO_cache struct {
	Id         int32
	Dept       int32
	Account    string
	Role       string
	Realname   string
	Commiter   string
	Gender     int8 // 0男，1女
	Email      string
	Mobile     string
	Join       int64
	Visits     int32 //访问次数
	QQ         int64
	Ip         string //上次登录ip
	Last       int64  //上次登录时间
	Fails      int8   //密码错误次数
	Locked     int64
	ClientLang string
	AttendNo   int32 //打卡机编号
	Deleted    bool
}

//检查密码是否正确，有Id优先查询Id，Id为0，Name查询account，realname，Mobile
type MSG_USER_CheckPasswd struct {
	QueryID uint32
	UserId  int32
	Name    string
	Rand    int64
	Passwd  string
}
type MSG_USER_CheckPasswd_result struct {
	QueryResultID uint32
	UserId        int32
	Result        ErrCode
}
type MSG_USER_Company_cache struct {
	Id       int32
	Name     string
	Phone    string
	Fax      string
	Address  string
	Zipcode  string
	Website  string
	Backyard string
	Admins   []string
	Deleted  bool
}
type MSG_USER_Dept_cache struct {
	Id          int32
	Name        string
	Parent      int32
	Path        []int32
	Grade       int8
	Order       int8
	Manager     int32
	ManagerName string
	Children    []*MSG_USER_Dept_cache `json:"children"` //部门管理前端的js需要小写
}
type MSG_USER_Dept_getParents struct {
	QueryID uint32
	Id      int32
}
type MSG_USER_Dept_getParents_result struct {
	QueryResultID uint32
	List          []*MSG_USER_Dept_cache
}

type MSG_USER_Dept_getDataStructure struct {
	QueryID    uint32
	RootDeptID int32
}
type MSG_USER_Dept_getDataStructure_result struct {
	QueryResultID uint32
	List          []*MSG_USER_Dept_cache
}

type MSG_USER_Dept_update struct {
	QueryID uint32
	List    []*MSG_USER_Dept_cache
}
type MSG_USER_Dept_delete struct {
	QueryID uint32
	DeptId  int32
}
type MSG_USER_Dept_delete_result struct {
	QueryResultID uint32
	Result        ErrCode
}
type MSG_USER_Pairs struct {
	Id       int32
	Account  string
	Realname string
}
type MSG_USER_getDeptUserPairs struct {
	QueryID uint32
	DeptId  int32
}
type MSG_USER_getDeptUserPairs_result struct {
	QueryResultID uint32
	List          []*MSG_USER_Pairs
}
type MSG_USER_getCompanyUsers struct {
	QueryID uint32
	Type    string
	Query   string
	DeptID  int32
	Sort    string
	Page    int
	PerPage int
	Where   string //注意转入原生的语句要防sql注入
	Total   int
}

type MSG_USER_getCompanyUsers_result struct {
	QueryResultID uint32
	List          []*MSG_USER_INFO_cache
	Total         int
}
