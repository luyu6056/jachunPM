package main

import "time"

type MSG_USER_GET_LoginSalt struct {
	Name string
}
type MSG_USER_GET_LoginSalt_result struct {
	Salt string
}
type MSG_USER_INFO_cache struct {
	Id          int32
	Dept        int32
	Account     string
	Role        string
	Realname    string
	Group       []int32
	Commiter    string
	Gender      int8 // 0男，1女
	Email       string
	Mobile      string
	Join        int64
	Visits      int32 //访问次数
	QQ          int64
	Ip          string //上次登录ip
	Last        int64  //上次登录时间
	Fails       int8   //密码错误次数
	Locked      int64
	ClientLang  string
	AttendNo    int32 //打卡机编号
	Deleted     bool
	Weixin      string
	Address     string
	AclProducts map[int32]bool                          //允许访问的产品
	AclProjects map[int32]bool                          //允许访问的项目
	IsAdmin     bool                                    `db:"-"` //暂定id为1是admin
	Config      map[string]map[string]map[string]string `db:"-"`
}

//检查密码是否正确，有Id优先查询Id，Id为0，Name查询account，realname，Mobile
type MSG_USER_CheckPasswd struct {
	UserId int32
	Name   string
	Rand   int64
	Passwd string
}
type MSG_USER_CheckPasswd_result struct {
	UserId int32
	Result ErrCode
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
	Id int32
}
type MSG_USER_Dept_getParents_result struct {
	List []*MSG_USER_Dept_cache
}

type MSG_USER_Dept_getDataStructure struct {
	RootDeptID int32
}
type MSG_USER_Dept_getDataStructure_result struct {
	List []*MSG_USER_Dept_cache
}

type MSG_USER_Dept_update struct {
	List []*MSG_USER_Dept_cache
}
type MSG_USER_Dept_delete struct {
	DeptId int32
}
type MSG_USER_Dept_delete_result struct {
	Result ErrCode
}
type MSG_USER_Pairs struct {
	Id       int32
	Account  string
	Realname string
}
type MSG_USER_getDeptUserPairs struct {
	DeptId int32
}
type MSG_USER_getDeptUserPairs_result struct {
	List []*MSG_USER_Pairs
}
type MSG_USER_getCompanyUsers struct {
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
	List  []*MSG_USER_INFO_cache
	Total int
}
type MSG_USER_Group_cache struct {
	Id          int32
	Name        string
	Role        string
	Desc        string
	Acl         []string
	AclProducts []int32
	AclProjects []int32
	Developer   int8
	Priv        map[string]map[string]bool
}
type MSG_USER_INFO_updateByID struct {
	UserID int32
	Update map[string]string
}
type MSG_USER_CheckAccount struct {
	Account string
}
type MSG_USER_CheckAccount_result struct {
	Result ErrCode
}
type MSG_USER_getPairs struct {
	Params          string
	UsersToAppended int32
}
type MSG_USER_getPairs_result struct {
	List []HtmlKeyValueStr
}
type MSG_USER_updateUserView struct {
	ProjectIds []int32
	ProductIds []int32
	UserIds    []int32
	GroupIds   []int32
}
type MSG_USER_getContactLists struct {
	Uid    int32
	Params string
}
type MSG_USER_getContactLists_result struct {
	List []HtmlKeyValueStr
}
type MSG_USER_getContactListByUid struct {
	Uid int32
}
type MSG_USER_getContactListByUid_result struct {
	List []HtmlKeyValueStr
}
type MSG_USER_getContactListById struct {
	Id int32
}
type MSG_USER_getContactListById_result struct {
	Result *MSG_USER_ContactList
}
type MSG_USER_ContactList struct {
	Id       int32
	Uid      int32
	ListName string
	UserList []int32
	Share    bool
}
type MSG_USER_insertUpdateContactList struct {
	Insert *MSG_USER_ContactList
}
type MSG_USER_insertUpdateContactList_result struct {
	Id int32
}
type MSG_USER_getGlobalContacts struct {
}
type MSG_USER_getGlobalContacts_result struct {
	Result []*MSG_USER_ContactList
}
type MSG_USER_team_getByTypeRoot struct {
	Type string
	Root []int32
}
type MSG_USER_team_getByTypeRoot_result struct {
	List []*MSG_USER_team_info
}
type MSG_USER_team_getByIds struct {
	Ids []int32
}
type MSG_USER_team_getByIds_result struct {
	List []*MSG_USER_team_info
}
type MSG_USER_team_info struct {
	Id       int32
	Root     int32
	Type     string
	Uid      int32
	Account  string
	Role     string
	Limited  string
	Join     time.Time
	Days     int16
	Hours    float64
	Estimate float64
	Consumed float64
	Left     float64
	Order    int8
	Deleted  bool   `db:"-"`
	Realname string `db:"-"`
}
type MSG_USER_team_addByList struct {
	List []*MSG_USER_team_info
}
type MSG_USER_Group_getPairs struct {
}
type MSG_USER_Group_getPairs_result struct {
	List []HtmlKeyValueStr
}
type MSG_USER_team_getByTypeUid struct {
	Type string
	Uid  int32
}
type MSG_USER_team_getByTypeUid_result struct {
	List []*MSG_USER_team_info
}
type MSG_USER_Userquery_info struct {
	Id       int32
	Uid      int32
	Module   string
	Title    string
	Form     string
	Sql      string
	Shortcut bool
}
type MSG_USER_user_getUserqueryByWhere struct {
	Where map[string]interface{}
}
type MSG_USER_user_getUserqueryByWhere_result struct {
	List []*MSG_USER_Userquery_info
}
type MSG_USER_team_getMemberPairsByTypeRoot struct {
	Type string
	Root int32
}
type MSG_USER_team_getMemberPairsByTypeRoot_result struct {
	List []HtmlKeyValueStr
}
type MSG_USER_team_updateByWhere struct {
	Where  map[string]interface{}
	Update map[string]interface{}
}

type MSG_USER_config_save struct {
	Uid     int32
	Module  string
	Section string
	Key     string
	Value   string
	Type    string //add deleted

}
