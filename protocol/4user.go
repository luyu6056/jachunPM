package protocol

import (
	"sync"
	"libraries"
	"time"
)

const (
	CMD_MSG_USER_GET_LoginSalt = 1735336196
	CMD_MSG_USER_GET_LoginSalt_result = 751862020
	CMD_MSG_USER_INFO_cache = -876112636
	CMD_MSG_USER_CheckPasswd = -2006312700
	CMD_MSG_USER_CheckPasswd_result = -101573884
	CMD_MSG_USER_Company_cache = 1302197764
	CMD_MSG_USER_Dept_cache = -681858812
	CMD_MSG_USER_Dept_getParents = -1996657916
	CMD_MSG_USER_Dept_getParents_result = 2045370628
	CMD_MSG_USER_Dept_getDataStructure = -136790780
	CMD_MSG_USER_Dept_getDataStructure_result = -1371742204
	CMD_MSG_USER_Dept_update = 470640644
	CMD_MSG_USER_Dept_delete = -1103486716
	CMD_MSG_USER_Dept_delete_result = -581285116
	CMD_MSG_USER_Pairs = 679671812
	CMD_MSG_USER_getDeptUserPairs = -334914812
	CMD_MSG_USER_getDeptUserPairs_result = -867223292
	CMD_MSG_USER_getCompanyUsers = 390890500
	CMD_MSG_USER_getCompanyUsers_result = 2081476612
	CMD_MSG_USER_Group_cache = -1780358652
	CMD_MSG_USER_INFO_updateByID = -198630396
	CMD_MSG_USER_CheckAccount = 295929604
	CMD_MSG_USER_CheckAccount_result = 953841924
	CMD_MSG_USER_getPairs = 338636804
	CMD_MSG_USER_getPairs_result = -1261518588
	CMD_MSG_USER_updateUserView = 1131689476
	CMD_MSG_USER_getContactLists = 1618683396
	CMD_MSG_USER_getContactLists_result = -1892740092
	CMD_MSG_USER_getContactListByUid = -2090728188
	CMD_MSG_USER_getContactListByUid_result = 155656196
	CMD_MSG_USER_getContactListById = 640438276
	CMD_MSG_USER_getContactListById_result = 368212996
	CMD_MSG_USER_ContactList = -803107580
	CMD_MSG_USER_insertUpdateContactList = 1170940932
	CMD_MSG_USER_insertUpdateContactList_result = -1137193468
	CMD_MSG_USER_getGlobalContacts = -1348354556
	CMD_MSG_USER_getGlobalContacts_result = -1400001276
	CMD_MSG_USER_team_getByTypeRoot = -1425965820
	CMD_MSG_USER_team_getByTypeRoot_result = -446270460
	CMD_MSG_USER_team_getByIds = 905195012
	CMD_MSG_USER_team_getByIds_result = -1549135356
	CMD_MSG_USER_team_info = 362981380
	CMD_MSG_USER_team_addByList = 870918148
	CMD_MSG_USER_Group_getPairs = 1806051844
	CMD_MSG_USER_Group_getPairs_result = -820050940
	CMD_MSG_USER_team_getByTypeUid = 468566788
	CMD_MSG_USER_team_getByTypeUid_result = 769103364
	CMD_MSG_USER_Userquery_info = 775618820
	CMD_MSG_USER_user_getUserqueryByWhere = -1330589692
	CMD_MSG_USER_user_getUserqueryByWhere_result = 1709400580
	CMD_MSG_USER_team_getMemberPairsByTypeRoot = -261057276
	CMD_MSG_USER_team_getMemberPairsByTypeRoot_result = 341559300
	CMD_MSG_USER_team_updateByWhere = 134391300
	CMD_MSG_USER_config_save = -942650620
)

type MSG_USER_GET_LoginSalt struct {
	Name string
}

var pool_MSG_USER_GET_LoginSalt = sync.Pool{New: func() interface{} { return &MSG_USER_GET_LoginSalt{} }}

func GET_MSG_USER_GET_LoginSalt() *MSG_USER_GET_LoginSalt {
	return pool_MSG_USER_GET_LoginSalt.Get().(*MSG_USER_GET_LoginSalt)
}

func (data *MSG_USER_GET_LoginSalt) cmd() int32 {
	return CMD_MSG_USER_GET_LoginSalt
}

func (data *MSG_USER_GET_LoginSalt) Put() {
	data.Name = ``
	pool_MSG_USER_GET_LoginSalt.Put(data)
}
func (data *MSG_USER_GET_LoginSalt) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_GET_LoginSalt,buf)
	WRITE_MSG_USER_GET_LoginSalt(data, buf)
}

func WRITE_MSG_USER_GET_LoginSalt(data *MSG_USER_GET_LoginSalt, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
}

func READ_MSG_USER_GET_LoginSalt(buf *libraries.MsgBuffer) *MSG_USER_GET_LoginSalt {
	data := pool_MSG_USER_GET_LoginSalt.Get().(*MSG_USER_GET_LoginSalt)
	data.read(buf)
	return data
}

func (data *MSG_USER_GET_LoginSalt) read(buf *libraries.MsgBuffer) {
	data.Name = READ_string(buf)

}

type MSG_USER_GET_LoginSalt_result struct {
	Salt string
}

var pool_MSG_USER_GET_LoginSalt_result = sync.Pool{New: func() interface{} { return &MSG_USER_GET_LoginSalt_result{} }}

func GET_MSG_USER_GET_LoginSalt_result() *MSG_USER_GET_LoginSalt_result {
	return pool_MSG_USER_GET_LoginSalt_result.Get().(*MSG_USER_GET_LoginSalt_result)
}

func (data *MSG_USER_GET_LoginSalt_result) cmd() int32 {
	return CMD_MSG_USER_GET_LoginSalt_result
}

func (data *MSG_USER_GET_LoginSalt_result) Put() {
	data.Salt = ``
	pool_MSG_USER_GET_LoginSalt_result.Put(data)
}
func (data *MSG_USER_GET_LoginSalt_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_GET_LoginSalt_result,buf)
	WRITE_MSG_USER_GET_LoginSalt_result(data, buf)
}

func WRITE_MSG_USER_GET_LoginSalt_result(data *MSG_USER_GET_LoginSalt_result, buf *libraries.MsgBuffer) {
	WRITE_string(data.Salt, buf)
}

func READ_MSG_USER_GET_LoginSalt_result(buf *libraries.MsgBuffer) *MSG_USER_GET_LoginSalt_result {
	data := pool_MSG_USER_GET_LoginSalt_result.Get().(*MSG_USER_GET_LoginSalt_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_GET_LoginSalt_result) read(buf *libraries.MsgBuffer) {
	data.Salt = READ_string(buf)

}

type MSG_USER_INFO_cache struct {
	Id int32
	Dept int32
	Account string
	Role string
	Realname string
	Group []int32
	Commiter string
	Gender int8
	Email string
	Mobile string
	Join int64
	Visits int32
	QQ int64
	Ip string
	Last int64
	Fails int8
	Locked int64
	ClientLang string
	AttendNo int32
	Deleted bool
	Weixin string
	Address string
	AclProducts map[int32]bool
	AclProjects map[int32]bool
	IsAdmin bool `db:"-"`
	Config map[string]map[string]map[string]string `db:"-"`
}

var pool_MSG_USER_INFO_cache = sync.Pool{New: func() interface{} { return &MSG_USER_INFO_cache{} }}

func GET_MSG_USER_INFO_cache() *MSG_USER_INFO_cache {
	return pool_MSG_USER_INFO_cache.Get().(*MSG_USER_INFO_cache)
}

func (data *MSG_USER_INFO_cache) cmd() int32 {
	return CMD_MSG_USER_INFO_cache
}

func (data *MSG_USER_INFO_cache) Put() {
	data.Id = 0
	data.Dept = 0
	data.Account = ``
	data.Role = ``
	data.Realname = ``
	data.Group = data.Group[:0]
	data.Commiter = ``
	data.Gender = 0
	data.Email = ``
	data.Mobile = ``
	data.Join = 0
	data.Visits = 0
	data.QQ = 0
	data.Ip = ``
	data.Last = 0
	data.Fails = 0
	data.Locked = 0
	data.ClientLang = ``
	data.AttendNo = 0
	data.Deleted = false
	data.Weixin = ``
	data.Address = ``
	data.AclProducts = nil
	data.AclProjects = nil
	data.IsAdmin = false
	data.Config = nil
	pool_MSG_USER_INFO_cache.Put(data)
}
func (data *MSG_USER_INFO_cache) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_INFO_cache,buf)
	WRITE_MSG_USER_INFO_cache(data, buf)
}

func WRITE_MSG_USER_INFO_cache(data *MSG_USER_INFO_cache, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Dept, buf)
	WRITE_string(data.Account, buf)
	WRITE_string(data.Role, buf)
	WRITE_string(data.Realname, buf)
	WRITE_int(len(data.Group), buf)
	for _, v := range data.Group{
		WRITE_int32(v, buf)
	}
	WRITE_string(data.Commiter, buf)
	WRITE_int8(data.Gender, buf)
	WRITE_string(data.Email, buf)
	WRITE_string(data.Mobile, buf)
	WRITE_int64(data.Join, buf)
	WRITE_int32(data.Visits, buf)
	WRITE_int64(data.QQ, buf)
	WRITE_string(data.Ip, buf)
	WRITE_int64(data.Last, buf)
	WRITE_int8(data.Fails, buf)
	WRITE_int64(data.Locked, buf)
	WRITE_string(data.ClientLang, buf)
	WRITE_int32(data.AttendNo, buf)
	WRITE_bool(data.Deleted, buf)
	WRITE_string(data.Weixin, buf)
	WRITE_string(data.Address, buf)
	WRITE_map(data.AclProducts,buf)
	WRITE_map(data.AclProjects,buf)
	WRITE_bool(data.IsAdmin, buf)
	WRITE_map(data.Config,buf)
}

func READ_MSG_USER_INFO_cache(buf *libraries.MsgBuffer) *MSG_USER_INFO_cache {
	data := pool_MSG_USER_INFO_cache.Get().(*MSG_USER_INFO_cache)
	data.read(buf)
	return data
}

func (data *MSG_USER_INFO_cache) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Dept = READ_int32(buf)
	data.Account = READ_string(buf)
	data.Role = READ_string(buf)
	data.Realname = READ_string(buf)
	Group_len := READ_int(buf)
	if Group_len>cap(data.Group){
		data.Group= make([]int32, Group_len)
	}else{
		data.Group = data.Group[:Group_len]
	}
	for i := 0; i < Group_len; i++ {
		data.Group[i] = READ_int32(buf)
	}
	data.Commiter = READ_string(buf)
	data.Gender = READ_int8(buf)
	data.Email = READ_string(buf)
	data.Mobile = READ_string(buf)
	data.Join = READ_int64(buf)
	data.Visits = READ_int32(buf)
	data.QQ = READ_int64(buf)
	data.Ip = READ_string(buf)
	data.Last = READ_int64(buf)
	data.Fails = READ_int8(buf)
	data.Locked = READ_int64(buf)
	data.ClientLang = READ_string(buf)
	data.AttendNo = READ_int32(buf)
	data.Deleted = READ_bool(buf)
	data.Weixin = READ_string(buf)
	data.Address = READ_string(buf)
	READ_map(&data.AclProducts,buf)
	READ_map(&data.AclProjects,buf)
	data.IsAdmin = READ_bool(buf)
	READ_map(&data.Config,buf)

}

type MSG_USER_CheckPasswd struct {
	UserId int32
	Name string
	Rand int64
	Passwd string
}

var pool_MSG_USER_CheckPasswd = sync.Pool{New: func() interface{} { return &MSG_USER_CheckPasswd{} }}

func GET_MSG_USER_CheckPasswd() *MSG_USER_CheckPasswd {
	return pool_MSG_USER_CheckPasswd.Get().(*MSG_USER_CheckPasswd)
}

func (data *MSG_USER_CheckPasswd) cmd() int32 {
	return CMD_MSG_USER_CheckPasswd
}

func (data *MSG_USER_CheckPasswd) Put() {
	data.UserId = 0
	data.Name = ``
	data.Rand = 0
	data.Passwd = ``
	pool_MSG_USER_CheckPasswd.Put(data)
}
func (data *MSG_USER_CheckPasswd) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_CheckPasswd,buf)
	WRITE_MSG_USER_CheckPasswd(data, buf)
}

func WRITE_MSG_USER_CheckPasswd(data *MSG_USER_CheckPasswd, buf *libraries.MsgBuffer) {
	WRITE_int32(data.UserId, buf)
	WRITE_string(data.Name, buf)
	WRITE_int64(data.Rand, buf)
	WRITE_string(data.Passwd, buf)
}

func READ_MSG_USER_CheckPasswd(buf *libraries.MsgBuffer) *MSG_USER_CheckPasswd {
	data := pool_MSG_USER_CheckPasswd.Get().(*MSG_USER_CheckPasswd)
	data.read(buf)
	return data
}

func (data *MSG_USER_CheckPasswd) read(buf *libraries.MsgBuffer) {
	data.UserId = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Rand = READ_int64(buf)
	data.Passwd = READ_string(buf)

}

type MSG_USER_CheckPasswd_result struct {
	UserId int32
	Result ErrCode
}

var pool_MSG_USER_CheckPasswd_result = sync.Pool{New: func() interface{} { return &MSG_USER_CheckPasswd_result{} }}

func GET_MSG_USER_CheckPasswd_result() *MSG_USER_CheckPasswd_result {
	return pool_MSG_USER_CheckPasswd_result.Get().(*MSG_USER_CheckPasswd_result)
}

func (data *MSG_USER_CheckPasswd_result) cmd() int32 {
	return CMD_MSG_USER_CheckPasswd_result
}

func (data *MSG_USER_CheckPasswd_result) Put() {
	data.UserId = 0
	data.Result = 0
	pool_MSG_USER_CheckPasswd_result.Put(data)
}
func (data *MSG_USER_CheckPasswd_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_CheckPasswd_result,buf)
	WRITE_MSG_USER_CheckPasswd_result(data, buf)
}

func WRITE_MSG_USER_CheckPasswd_result(data *MSG_USER_CheckPasswd_result, buf *libraries.MsgBuffer) {
	WRITE_int32(data.UserId, buf)
	WRITE_ErrCode(data.Result, buf)
}

func READ_MSG_USER_CheckPasswd_result(buf *libraries.MsgBuffer) *MSG_USER_CheckPasswd_result {
	data := pool_MSG_USER_CheckPasswd_result.Get().(*MSG_USER_CheckPasswd_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_CheckPasswd_result) read(buf *libraries.MsgBuffer) {
	data.UserId = READ_int32(buf)
	data.Result = READ_ErrCode(buf)

}

type MSG_USER_Company_cache struct {
	Id int32
	Name string
	Phone string
	Fax string
	Address string
	Zipcode string
	Website string
	Backyard string
	Admins []string
	Deleted bool
}

var pool_MSG_USER_Company_cache = sync.Pool{New: func() interface{} { return &MSG_USER_Company_cache{} }}

func GET_MSG_USER_Company_cache() *MSG_USER_Company_cache {
	return pool_MSG_USER_Company_cache.Get().(*MSG_USER_Company_cache)
}

func (data *MSG_USER_Company_cache) cmd() int32 {
	return CMD_MSG_USER_Company_cache
}

func (data *MSG_USER_Company_cache) Put() {
	data.Id = 0
	data.Name = ``
	data.Phone = ``
	data.Fax = ``
	data.Address = ``
	data.Zipcode = ``
	data.Website = ``
	data.Backyard = ``
	data.Admins = data.Admins[:0]
	data.Deleted = false
	pool_MSG_USER_Company_cache.Put(data)
}
func (data *MSG_USER_Company_cache) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Company_cache,buf)
	WRITE_MSG_USER_Company_cache(data, buf)
}

func WRITE_MSG_USER_Company_cache(data *MSG_USER_Company_cache, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Phone, buf)
	WRITE_string(data.Fax, buf)
	WRITE_string(data.Address, buf)
	WRITE_string(data.Zipcode, buf)
	WRITE_string(data.Website, buf)
	WRITE_string(data.Backyard, buf)
	WRITE_int(len(data.Admins), buf)
	for _, v := range data.Admins{
		WRITE_string(v, buf)
	}
	WRITE_bool(data.Deleted, buf)
}

func READ_MSG_USER_Company_cache(buf *libraries.MsgBuffer) *MSG_USER_Company_cache {
	data := pool_MSG_USER_Company_cache.Get().(*MSG_USER_Company_cache)
	data.read(buf)
	return data
}

func (data *MSG_USER_Company_cache) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Phone = READ_string(buf)
	data.Fax = READ_string(buf)
	data.Address = READ_string(buf)
	data.Zipcode = READ_string(buf)
	data.Website = READ_string(buf)
	data.Backyard = READ_string(buf)
	Admins_len := READ_int(buf)
	if Admins_len>cap(data.Admins){
		data.Admins= make([]string, Admins_len)
	}else{
		data.Admins = data.Admins[:Admins_len]
	}
	for i := 0; i < Admins_len; i++ {
		data.Admins[i] = READ_string(buf)
	}
	data.Deleted = READ_bool(buf)

}

type MSG_USER_Dept_cache struct {
	Id int32
	Name string
	Parent int32
	Path []int32
	Grade int8
	Order int8
	Manager int32
	ManagerName string
	Children []*MSG_USER_Dept_cache `json:"children"`
}

var pool_MSG_USER_Dept_cache = sync.Pool{New: func() interface{} { return &MSG_USER_Dept_cache{} }}

func GET_MSG_USER_Dept_cache() *MSG_USER_Dept_cache {
	return pool_MSG_USER_Dept_cache.Get().(*MSG_USER_Dept_cache)
}

func (data *MSG_USER_Dept_cache) cmd() int32 {
	return CMD_MSG_USER_Dept_cache
}

func (data *MSG_USER_Dept_cache) Put() {
	data.Id = 0
	data.Name = ``
	data.Parent = 0
	data.Path = data.Path[:0]
	data.Grade = 0
	data.Order = 0
	data.Manager = 0
	data.ManagerName = ``
	for _,v := range data.Children {
		v.Put()
	}
	data.Children = data.Children[:0]
	pool_MSG_USER_Dept_cache.Put(data)
}
func (data *MSG_USER_Dept_cache) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_cache,buf)
	WRITE_MSG_USER_Dept_cache(data, buf)
}

func WRITE_MSG_USER_Dept_cache(data *MSG_USER_Dept_cache, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_string(data.Name, buf)
	WRITE_int32(data.Parent, buf)
	WRITE_int(len(data.Path), buf)
	for _, v := range data.Path{
		WRITE_int32(v, buf)
	}
	WRITE_int8(data.Grade, buf)
	WRITE_int8(data.Order, buf)
	WRITE_int32(data.Manager, buf)
	WRITE_string(data.ManagerName, buf)
	WRITE_int(len(data.Children), buf)
	for _, v := range data.Children{
		WRITE_MSG_USER_Dept_cache(v, buf)
	}
}

func READ_MSG_USER_Dept_cache(buf *libraries.MsgBuffer) *MSG_USER_Dept_cache {
	data := pool_MSG_USER_Dept_cache.Get().(*MSG_USER_Dept_cache)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_cache) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Parent = READ_int32(buf)
	Path_len := READ_int(buf)
	if Path_len>cap(data.Path){
		data.Path= make([]int32, Path_len)
	}else{
		data.Path = data.Path[:Path_len]
	}
	for i := 0; i < Path_len; i++ {
		data.Path[i] = READ_int32(buf)
	}
	data.Grade = READ_int8(buf)
	data.Order = READ_int8(buf)
	data.Manager = READ_int32(buf)
	data.ManagerName = READ_string(buf)
	Children_len := READ_int(buf)
	if Children_len>cap(data.Children){
		data.Children= make([]*MSG_USER_Dept_cache, Children_len)
	}else{
		data.Children = data.Children[:Children_len]
	}
	for i := 0; i < Children_len; i++ {
		data.Children[i] = READ_MSG_USER_Dept_cache(buf)
	}

}

type MSG_USER_Dept_getParents struct {
	Id int32
}

var pool_MSG_USER_Dept_getParents = sync.Pool{New: func() interface{} { return &MSG_USER_Dept_getParents{} }}

func GET_MSG_USER_Dept_getParents() *MSG_USER_Dept_getParents {
	return pool_MSG_USER_Dept_getParents.Get().(*MSG_USER_Dept_getParents)
}

func (data *MSG_USER_Dept_getParents) cmd() int32 {
	return CMD_MSG_USER_Dept_getParents
}

func (data *MSG_USER_Dept_getParents) Put() {
	data.Id = 0
	pool_MSG_USER_Dept_getParents.Put(data)
}
func (data *MSG_USER_Dept_getParents) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_getParents,buf)
	WRITE_MSG_USER_Dept_getParents(data, buf)
}

func WRITE_MSG_USER_Dept_getParents(data *MSG_USER_Dept_getParents, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
}

func READ_MSG_USER_Dept_getParents(buf *libraries.MsgBuffer) *MSG_USER_Dept_getParents {
	data := pool_MSG_USER_Dept_getParents.Get().(*MSG_USER_Dept_getParents)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_getParents) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)

}

type MSG_USER_Dept_getParents_result struct {
	List []*MSG_USER_Dept_cache
}

var pool_MSG_USER_Dept_getParents_result = sync.Pool{New: func() interface{} { return &MSG_USER_Dept_getParents_result{} }}

func GET_MSG_USER_Dept_getParents_result() *MSG_USER_Dept_getParents_result {
	return pool_MSG_USER_Dept_getParents_result.Get().(*MSG_USER_Dept_getParents_result)
}

func (data *MSG_USER_Dept_getParents_result) cmd() int32 {
	return CMD_MSG_USER_Dept_getParents_result
}

func (data *MSG_USER_Dept_getParents_result) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_USER_Dept_getParents_result.Put(data)
}
func (data *MSG_USER_Dept_getParents_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_getParents_result,buf)
	WRITE_MSG_USER_Dept_getParents_result(data, buf)
}

func WRITE_MSG_USER_Dept_getParents_result(data *MSG_USER_Dept_getParents_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_USER_Dept_cache(v, buf)
	}
}

func READ_MSG_USER_Dept_getParents_result(buf *libraries.MsgBuffer) *MSG_USER_Dept_getParents_result {
	data := pool_MSG_USER_Dept_getParents_result.Get().(*MSG_USER_Dept_getParents_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_getParents_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_USER_Dept_cache, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_USER_Dept_cache(buf)
	}

}

type MSG_USER_Dept_getDataStructure struct {
	RootDeptID int32
}

var pool_MSG_USER_Dept_getDataStructure = sync.Pool{New: func() interface{} { return &MSG_USER_Dept_getDataStructure{} }}

func GET_MSG_USER_Dept_getDataStructure() *MSG_USER_Dept_getDataStructure {
	return pool_MSG_USER_Dept_getDataStructure.Get().(*MSG_USER_Dept_getDataStructure)
}

func (data *MSG_USER_Dept_getDataStructure) cmd() int32 {
	return CMD_MSG_USER_Dept_getDataStructure
}

func (data *MSG_USER_Dept_getDataStructure) Put() {
	data.RootDeptID = 0
	pool_MSG_USER_Dept_getDataStructure.Put(data)
}
func (data *MSG_USER_Dept_getDataStructure) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_getDataStructure,buf)
	WRITE_MSG_USER_Dept_getDataStructure(data, buf)
}

func WRITE_MSG_USER_Dept_getDataStructure(data *MSG_USER_Dept_getDataStructure, buf *libraries.MsgBuffer) {
	WRITE_int32(data.RootDeptID, buf)
}

func READ_MSG_USER_Dept_getDataStructure(buf *libraries.MsgBuffer) *MSG_USER_Dept_getDataStructure {
	data := pool_MSG_USER_Dept_getDataStructure.Get().(*MSG_USER_Dept_getDataStructure)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_getDataStructure) read(buf *libraries.MsgBuffer) {
	data.RootDeptID = READ_int32(buf)

}

type MSG_USER_Dept_getDataStructure_result struct {
	List []*MSG_USER_Dept_cache
}

var pool_MSG_USER_Dept_getDataStructure_result = sync.Pool{New: func() interface{} { return &MSG_USER_Dept_getDataStructure_result{} }}

func GET_MSG_USER_Dept_getDataStructure_result() *MSG_USER_Dept_getDataStructure_result {
	return pool_MSG_USER_Dept_getDataStructure_result.Get().(*MSG_USER_Dept_getDataStructure_result)
}

func (data *MSG_USER_Dept_getDataStructure_result) cmd() int32 {
	return CMD_MSG_USER_Dept_getDataStructure_result
}

func (data *MSG_USER_Dept_getDataStructure_result) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_USER_Dept_getDataStructure_result.Put(data)
}
func (data *MSG_USER_Dept_getDataStructure_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_getDataStructure_result,buf)
	WRITE_MSG_USER_Dept_getDataStructure_result(data, buf)
}

func WRITE_MSG_USER_Dept_getDataStructure_result(data *MSG_USER_Dept_getDataStructure_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_USER_Dept_cache(v, buf)
	}
}

func READ_MSG_USER_Dept_getDataStructure_result(buf *libraries.MsgBuffer) *MSG_USER_Dept_getDataStructure_result {
	data := pool_MSG_USER_Dept_getDataStructure_result.Get().(*MSG_USER_Dept_getDataStructure_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_getDataStructure_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_USER_Dept_cache, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_USER_Dept_cache(buf)
	}

}

type MSG_USER_Dept_update struct {
	List []*MSG_USER_Dept_cache
}

var pool_MSG_USER_Dept_update = sync.Pool{New: func() interface{} { return &MSG_USER_Dept_update{} }}

func GET_MSG_USER_Dept_update() *MSG_USER_Dept_update {
	return pool_MSG_USER_Dept_update.Get().(*MSG_USER_Dept_update)
}

func (data *MSG_USER_Dept_update) cmd() int32 {
	return CMD_MSG_USER_Dept_update
}

func (data *MSG_USER_Dept_update) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_USER_Dept_update.Put(data)
}
func (data *MSG_USER_Dept_update) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_update,buf)
	WRITE_MSG_USER_Dept_update(data, buf)
}

func WRITE_MSG_USER_Dept_update(data *MSG_USER_Dept_update, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_USER_Dept_cache(v, buf)
	}
}

func READ_MSG_USER_Dept_update(buf *libraries.MsgBuffer) *MSG_USER_Dept_update {
	data := pool_MSG_USER_Dept_update.Get().(*MSG_USER_Dept_update)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_update) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_USER_Dept_cache, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_USER_Dept_cache(buf)
	}

}

type MSG_USER_Dept_delete struct {
	DeptId int32
}

var pool_MSG_USER_Dept_delete = sync.Pool{New: func() interface{} { return &MSG_USER_Dept_delete{} }}

func GET_MSG_USER_Dept_delete() *MSG_USER_Dept_delete {
	return pool_MSG_USER_Dept_delete.Get().(*MSG_USER_Dept_delete)
}

func (data *MSG_USER_Dept_delete) cmd() int32 {
	return CMD_MSG_USER_Dept_delete
}

func (data *MSG_USER_Dept_delete) Put() {
	data.DeptId = 0
	pool_MSG_USER_Dept_delete.Put(data)
}
func (data *MSG_USER_Dept_delete) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_delete,buf)
	WRITE_MSG_USER_Dept_delete(data, buf)
}

func WRITE_MSG_USER_Dept_delete(data *MSG_USER_Dept_delete, buf *libraries.MsgBuffer) {
	WRITE_int32(data.DeptId, buf)
}

func READ_MSG_USER_Dept_delete(buf *libraries.MsgBuffer) *MSG_USER_Dept_delete {
	data := pool_MSG_USER_Dept_delete.Get().(*MSG_USER_Dept_delete)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_delete) read(buf *libraries.MsgBuffer) {
	data.DeptId = READ_int32(buf)

}

type MSG_USER_Dept_delete_result struct {
	Result ErrCode
}

var pool_MSG_USER_Dept_delete_result = sync.Pool{New: func() interface{} { return &MSG_USER_Dept_delete_result{} }}

func GET_MSG_USER_Dept_delete_result() *MSG_USER_Dept_delete_result {
	return pool_MSG_USER_Dept_delete_result.Get().(*MSG_USER_Dept_delete_result)
}

func (data *MSG_USER_Dept_delete_result) cmd() int32 {
	return CMD_MSG_USER_Dept_delete_result
}

func (data *MSG_USER_Dept_delete_result) Put() {
	data.Result = 0
	pool_MSG_USER_Dept_delete_result.Put(data)
}
func (data *MSG_USER_Dept_delete_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_delete_result,buf)
	WRITE_MSG_USER_Dept_delete_result(data, buf)
}

func WRITE_MSG_USER_Dept_delete_result(data *MSG_USER_Dept_delete_result, buf *libraries.MsgBuffer) {
	WRITE_ErrCode(data.Result, buf)
}

func READ_MSG_USER_Dept_delete_result(buf *libraries.MsgBuffer) *MSG_USER_Dept_delete_result {
	data := pool_MSG_USER_Dept_delete_result.Get().(*MSG_USER_Dept_delete_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_delete_result) read(buf *libraries.MsgBuffer) {
	data.Result = READ_ErrCode(buf)

}

type MSG_USER_Pairs struct {
	Id int32
	Account string
	Realname string
}

var pool_MSG_USER_Pairs = sync.Pool{New: func() interface{} { return &MSG_USER_Pairs{} }}

func GET_MSG_USER_Pairs() *MSG_USER_Pairs {
	return pool_MSG_USER_Pairs.Get().(*MSG_USER_Pairs)
}

func (data *MSG_USER_Pairs) cmd() int32 {
	return CMD_MSG_USER_Pairs
}

func (data *MSG_USER_Pairs) Put() {
	data.Id = 0
	data.Account = ``
	data.Realname = ``
	pool_MSG_USER_Pairs.Put(data)
}
func (data *MSG_USER_Pairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Pairs,buf)
	WRITE_MSG_USER_Pairs(data, buf)
}

func WRITE_MSG_USER_Pairs(data *MSG_USER_Pairs, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_string(data.Account, buf)
	WRITE_string(data.Realname, buf)
}

func READ_MSG_USER_Pairs(buf *libraries.MsgBuffer) *MSG_USER_Pairs {
	data := pool_MSG_USER_Pairs.Get().(*MSG_USER_Pairs)
	data.read(buf)
	return data
}

func (data *MSG_USER_Pairs) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Account = READ_string(buf)
	data.Realname = READ_string(buf)

}

type MSG_USER_getDeptUserPairs struct {
	DeptId int32
}

var pool_MSG_USER_getDeptUserPairs = sync.Pool{New: func() interface{} { return &MSG_USER_getDeptUserPairs{} }}

func GET_MSG_USER_getDeptUserPairs() *MSG_USER_getDeptUserPairs {
	return pool_MSG_USER_getDeptUserPairs.Get().(*MSG_USER_getDeptUserPairs)
}

func (data *MSG_USER_getDeptUserPairs) cmd() int32 {
	return CMD_MSG_USER_getDeptUserPairs
}

func (data *MSG_USER_getDeptUserPairs) Put() {
	data.DeptId = 0
	pool_MSG_USER_getDeptUserPairs.Put(data)
}
func (data *MSG_USER_getDeptUserPairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getDeptUserPairs,buf)
	WRITE_MSG_USER_getDeptUserPairs(data, buf)
}

func WRITE_MSG_USER_getDeptUserPairs(data *MSG_USER_getDeptUserPairs, buf *libraries.MsgBuffer) {
	WRITE_int32(data.DeptId, buf)
}

func READ_MSG_USER_getDeptUserPairs(buf *libraries.MsgBuffer) *MSG_USER_getDeptUserPairs {
	data := pool_MSG_USER_getDeptUserPairs.Get().(*MSG_USER_getDeptUserPairs)
	data.read(buf)
	return data
}

func (data *MSG_USER_getDeptUserPairs) read(buf *libraries.MsgBuffer) {
	data.DeptId = READ_int32(buf)

}

type MSG_USER_getDeptUserPairs_result struct {
	List []*MSG_USER_Pairs
}

var pool_MSG_USER_getDeptUserPairs_result = sync.Pool{New: func() interface{} { return &MSG_USER_getDeptUserPairs_result{} }}

func GET_MSG_USER_getDeptUserPairs_result() *MSG_USER_getDeptUserPairs_result {
	return pool_MSG_USER_getDeptUserPairs_result.Get().(*MSG_USER_getDeptUserPairs_result)
}

func (data *MSG_USER_getDeptUserPairs_result) cmd() int32 {
	return CMD_MSG_USER_getDeptUserPairs_result
}

func (data *MSG_USER_getDeptUserPairs_result) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_USER_getDeptUserPairs_result.Put(data)
}
func (data *MSG_USER_getDeptUserPairs_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getDeptUserPairs_result,buf)
	WRITE_MSG_USER_getDeptUserPairs_result(data, buf)
}

func WRITE_MSG_USER_getDeptUserPairs_result(data *MSG_USER_getDeptUserPairs_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_USER_Pairs(v, buf)
	}
}

func READ_MSG_USER_getDeptUserPairs_result(buf *libraries.MsgBuffer) *MSG_USER_getDeptUserPairs_result {
	data := pool_MSG_USER_getDeptUserPairs_result.Get().(*MSG_USER_getDeptUserPairs_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_getDeptUserPairs_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_USER_Pairs, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_USER_Pairs(buf)
	}

}

type MSG_USER_getCompanyUsers struct {
	Type string
	Query string
	DeptID int32
	Sort string
	Page int
	PerPage int
	Where string
	Total int
}

var pool_MSG_USER_getCompanyUsers = sync.Pool{New: func() interface{} { return &MSG_USER_getCompanyUsers{} }}

func GET_MSG_USER_getCompanyUsers() *MSG_USER_getCompanyUsers {
	return pool_MSG_USER_getCompanyUsers.Get().(*MSG_USER_getCompanyUsers)
}

func (data *MSG_USER_getCompanyUsers) cmd() int32 {
	return CMD_MSG_USER_getCompanyUsers
}

func (data *MSG_USER_getCompanyUsers) Put() {
	data.Type = ``
	data.Query = ``
	data.DeptID = 0
	data.Sort = ``
	data.Page = 0
	data.PerPage = 0
	data.Where = ``
	data.Total = 0
	pool_MSG_USER_getCompanyUsers.Put(data)
}
func (data *MSG_USER_getCompanyUsers) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getCompanyUsers,buf)
	WRITE_MSG_USER_getCompanyUsers(data, buf)
}

func WRITE_MSG_USER_getCompanyUsers(data *MSG_USER_getCompanyUsers, buf *libraries.MsgBuffer) {
	WRITE_string(data.Type, buf)
	WRITE_string(data.Query, buf)
	WRITE_int32(data.DeptID, buf)
	WRITE_string(data.Sort, buf)
	WRITE_int(data.Page, buf)
	WRITE_int(data.PerPage, buf)
	WRITE_string(data.Where, buf)
	WRITE_int(data.Total, buf)
}

func READ_MSG_USER_getCompanyUsers(buf *libraries.MsgBuffer) *MSG_USER_getCompanyUsers {
	data := pool_MSG_USER_getCompanyUsers.Get().(*MSG_USER_getCompanyUsers)
	data.read(buf)
	return data
}

func (data *MSG_USER_getCompanyUsers) read(buf *libraries.MsgBuffer) {
	data.Type = READ_string(buf)
	data.Query = READ_string(buf)
	data.DeptID = READ_int32(buf)
	data.Sort = READ_string(buf)
	data.Page = READ_int(buf)
	data.PerPage = READ_int(buf)
	data.Where = READ_string(buf)
	data.Total = READ_int(buf)

}

type MSG_USER_getCompanyUsers_result struct {
	List []*MSG_USER_INFO_cache
	Total int
}

var pool_MSG_USER_getCompanyUsers_result = sync.Pool{New: func() interface{} { return &MSG_USER_getCompanyUsers_result{} }}

func GET_MSG_USER_getCompanyUsers_result() *MSG_USER_getCompanyUsers_result {
	return pool_MSG_USER_getCompanyUsers_result.Get().(*MSG_USER_getCompanyUsers_result)
}

func (data *MSG_USER_getCompanyUsers_result) cmd() int32 {
	return CMD_MSG_USER_getCompanyUsers_result
}

func (data *MSG_USER_getCompanyUsers_result) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	data.Total = 0
	pool_MSG_USER_getCompanyUsers_result.Put(data)
}
func (data *MSG_USER_getCompanyUsers_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getCompanyUsers_result,buf)
	WRITE_MSG_USER_getCompanyUsers_result(data, buf)
}

func WRITE_MSG_USER_getCompanyUsers_result(data *MSG_USER_getCompanyUsers_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_USER_INFO_cache(v, buf)
	}
	WRITE_int(data.Total, buf)
}

func READ_MSG_USER_getCompanyUsers_result(buf *libraries.MsgBuffer) *MSG_USER_getCompanyUsers_result {
	data := pool_MSG_USER_getCompanyUsers_result.Get().(*MSG_USER_getCompanyUsers_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_getCompanyUsers_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_USER_INFO_cache, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_USER_INFO_cache(buf)
	}
	data.Total = READ_int(buf)

}

type MSG_USER_Group_cache struct {
	Id int32
	Name string
	Role string
	Desc string
	Acl []string
	AclProducts []int32
	AclProjects []int32
	Developer int8
	Priv map[string]map[string]bool
}

var pool_MSG_USER_Group_cache = sync.Pool{New: func() interface{} { return &MSG_USER_Group_cache{} }}

func GET_MSG_USER_Group_cache() *MSG_USER_Group_cache {
	return pool_MSG_USER_Group_cache.Get().(*MSG_USER_Group_cache)
}

func (data *MSG_USER_Group_cache) cmd() int32 {
	return CMD_MSG_USER_Group_cache
}

func (data *MSG_USER_Group_cache) Put() {
	data.Id = 0
	data.Name = ``
	data.Role = ``
	data.Desc = ``
	data.Acl = data.Acl[:0]
	data.AclProducts = data.AclProducts[:0]
	data.AclProjects = data.AclProjects[:0]
	data.Developer = 0
	data.Priv = nil
	pool_MSG_USER_Group_cache.Put(data)
}
func (data *MSG_USER_Group_cache) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Group_cache,buf)
	WRITE_MSG_USER_Group_cache(data, buf)
}

func WRITE_MSG_USER_Group_cache(data *MSG_USER_Group_cache, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Role, buf)
	WRITE_string(data.Desc, buf)
	WRITE_int(len(data.Acl), buf)
	for _, v := range data.Acl{
		WRITE_string(v, buf)
	}
	WRITE_int(len(data.AclProducts), buf)
	for _, v := range data.AclProducts{
		WRITE_int32(v, buf)
	}
	WRITE_int(len(data.AclProjects), buf)
	for _, v := range data.AclProjects{
		WRITE_int32(v, buf)
	}
	WRITE_int8(data.Developer, buf)
	WRITE_map(data.Priv,buf)
}

func READ_MSG_USER_Group_cache(buf *libraries.MsgBuffer) *MSG_USER_Group_cache {
	data := pool_MSG_USER_Group_cache.Get().(*MSG_USER_Group_cache)
	data.read(buf)
	return data
}

func (data *MSG_USER_Group_cache) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Role = READ_string(buf)
	data.Desc = READ_string(buf)
	Acl_len := READ_int(buf)
	if Acl_len>cap(data.Acl){
		data.Acl= make([]string, Acl_len)
	}else{
		data.Acl = data.Acl[:Acl_len]
	}
	for i := 0; i < Acl_len; i++ {
		data.Acl[i] = READ_string(buf)
	}
	AclProducts_len := READ_int(buf)
	if AclProducts_len>cap(data.AclProducts){
		data.AclProducts= make([]int32, AclProducts_len)
	}else{
		data.AclProducts = data.AclProducts[:AclProducts_len]
	}
	for i := 0; i < AclProducts_len; i++ {
		data.AclProducts[i] = READ_int32(buf)
	}
	AclProjects_len := READ_int(buf)
	if AclProjects_len>cap(data.AclProjects){
		data.AclProjects= make([]int32, AclProjects_len)
	}else{
		data.AclProjects = data.AclProjects[:AclProjects_len]
	}
	for i := 0; i < AclProjects_len; i++ {
		data.AclProjects[i] = READ_int32(buf)
	}
	data.Developer = READ_int8(buf)
	READ_map(&data.Priv,buf)

}

type MSG_USER_INFO_updateByID struct {
	UserID int32
	Update map[string]string
}

var pool_MSG_USER_INFO_updateByID = sync.Pool{New: func() interface{} { return &MSG_USER_INFO_updateByID{} }}

func GET_MSG_USER_INFO_updateByID() *MSG_USER_INFO_updateByID {
	return pool_MSG_USER_INFO_updateByID.Get().(*MSG_USER_INFO_updateByID)
}

func (data *MSG_USER_INFO_updateByID) cmd() int32 {
	return CMD_MSG_USER_INFO_updateByID
}

func (data *MSG_USER_INFO_updateByID) Put() {
	data.UserID = 0
	data.Update = nil
	pool_MSG_USER_INFO_updateByID.Put(data)
}
func (data *MSG_USER_INFO_updateByID) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_INFO_updateByID,buf)
	WRITE_MSG_USER_INFO_updateByID(data, buf)
}

func WRITE_MSG_USER_INFO_updateByID(data *MSG_USER_INFO_updateByID, buf *libraries.MsgBuffer) {
	WRITE_int32(data.UserID, buf)
	WRITE_map(data.Update,buf)
}

func READ_MSG_USER_INFO_updateByID(buf *libraries.MsgBuffer) *MSG_USER_INFO_updateByID {
	data := pool_MSG_USER_INFO_updateByID.Get().(*MSG_USER_INFO_updateByID)
	data.read(buf)
	return data
}

func (data *MSG_USER_INFO_updateByID) read(buf *libraries.MsgBuffer) {
	data.UserID = READ_int32(buf)
	READ_map(&data.Update,buf)

}

type MSG_USER_CheckAccount struct {
	Account string
}

var pool_MSG_USER_CheckAccount = sync.Pool{New: func() interface{} { return &MSG_USER_CheckAccount{} }}

func GET_MSG_USER_CheckAccount() *MSG_USER_CheckAccount {
	return pool_MSG_USER_CheckAccount.Get().(*MSG_USER_CheckAccount)
}

func (data *MSG_USER_CheckAccount) cmd() int32 {
	return CMD_MSG_USER_CheckAccount
}

func (data *MSG_USER_CheckAccount) Put() {
	data.Account = ``
	pool_MSG_USER_CheckAccount.Put(data)
}
func (data *MSG_USER_CheckAccount) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_CheckAccount,buf)
	WRITE_MSG_USER_CheckAccount(data, buf)
}

func WRITE_MSG_USER_CheckAccount(data *MSG_USER_CheckAccount, buf *libraries.MsgBuffer) {
	WRITE_string(data.Account, buf)
}

func READ_MSG_USER_CheckAccount(buf *libraries.MsgBuffer) *MSG_USER_CheckAccount {
	data := pool_MSG_USER_CheckAccount.Get().(*MSG_USER_CheckAccount)
	data.read(buf)
	return data
}

func (data *MSG_USER_CheckAccount) read(buf *libraries.MsgBuffer) {
	data.Account = READ_string(buf)

}

type MSG_USER_CheckAccount_result struct {
	Result ErrCode
}

var pool_MSG_USER_CheckAccount_result = sync.Pool{New: func() interface{} { return &MSG_USER_CheckAccount_result{} }}

func GET_MSG_USER_CheckAccount_result() *MSG_USER_CheckAccount_result {
	return pool_MSG_USER_CheckAccount_result.Get().(*MSG_USER_CheckAccount_result)
}

func (data *MSG_USER_CheckAccount_result) cmd() int32 {
	return CMD_MSG_USER_CheckAccount_result
}

func (data *MSG_USER_CheckAccount_result) Put() {
	data.Result = 0
	pool_MSG_USER_CheckAccount_result.Put(data)
}
func (data *MSG_USER_CheckAccount_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_CheckAccount_result,buf)
	WRITE_MSG_USER_CheckAccount_result(data, buf)
}

func WRITE_MSG_USER_CheckAccount_result(data *MSG_USER_CheckAccount_result, buf *libraries.MsgBuffer) {
	WRITE_ErrCode(data.Result, buf)
}

func READ_MSG_USER_CheckAccount_result(buf *libraries.MsgBuffer) *MSG_USER_CheckAccount_result {
	data := pool_MSG_USER_CheckAccount_result.Get().(*MSG_USER_CheckAccount_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_CheckAccount_result) read(buf *libraries.MsgBuffer) {
	data.Result = READ_ErrCode(buf)

}

type MSG_USER_getPairs struct {
	Params string
	UsersToAppended int32
}

var pool_MSG_USER_getPairs = sync.Pool{New: func() interface{} { return &MSG_USER_getPairs{} }}

func GET_MSG_USER_getPairs() *MSG_USER_getPairs {
	return pool_MSG_USER_getPairs.Get().(*MSG_USER_getPairs)
}

func (data *MSG_USER_getPairs) cmd() int32 {
	return CMD_MSG_USER_getPairs
}

func (data *MSG_USER_getPairs) Put() {
	data.Params = ``
	data.UsersToAppended = 0
	pool_MSG_USER_getPairs.Put(data)
}
func (data *MSG_USER_getPairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getPairs,buf)
	WRITE_MSG_USER_getPairs(data, buf)
}

func WRITE_MSG_USER_getPairs(data *MSG_USER_getPairs, buf *libraries.MsgBuffer) {
	WRITE_string(data.Params, buf)
	WRITE_int32(data.UsersToAppended, buf)
}

func READ_MSG_USER_getPairs(buf *libraries.MsgBuffer) *MSG_USER_getPairs {
	data := pool_MSG_USER_getPairs.Get().(*MSG_USER_getPairs)
	data.read(buf)
	return data
}

func (data *MSG_USER_getPairs) read(buf *libraries.MsgBuffer) {
	data.Params = READ_string(buf)
	data.UsersToAppended = READ_int32(buf)

}

type MSG_USER_getPairs_result struct {
	List []HtmlKeyValueStr
}

var pool_MSG_USER_getPairs_result = sync.Pool{New: func() interface{} { return &MSG_USER_getPairs_result{} }}

func GET_MSG_USER_getPairs_result() *MSG_USER_getPairs_result {
	return pool_MSG_USER_getPairs_result.Get().(*MSG_USER_getPairs_result)
}

func (data *MSG_USER_getPairs_result) cmd() int32 {
	return CMD_MSG_USER_getPairs_result
}

func (data *MSG_USER_getPairs_result) Put() {
	data.List = data.List[:0]
	pool_MSG_USER_getPairs_result.Put(data)
}
func (data *MSG_USER_getPairs_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getPairs_result,buf)
	WRITE_MSG_USER_getPairs_result(data, buf)
}

func WRITE_MSG_USER_getPairs_result(data *MSG_USER_getPairs_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_USER_getPairs_result(buf *libraries.MsgBuffer) *MSG_USER_getPairs_result {
	data := pool_MSG_USER_getPairs_result.Get().(*MSG_USER_getPairs_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_getPairs_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]HtmlKeyValueStr, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_HtmlKeyValueStr(buf)
	}

}

type MSG_USER_updateUserView struct {
	ProjectIds []int32
	ProductIds []int32
	UserIds []int32
	GroupIds []int32
}

var pool_MSG_USER_updateUserView = sync.Pool{New: func() interface{} { return &MSG_USER_updateUserView{} }}

func GET_MSG_USER_updateUserView() *MSG_USER_updateUserView {
	return pool_MSG_USER_updateUserView.Get().(*MSG_USER_updateUserView)
}

func (data *MSG_USER_updateUserView) cmd() int32 {
	return CMD_MSG_USER_updateUserView
}

func (data *MSG_USER_updateUserView) Put() {
	data.ProjectIds = data.ProjectIds[:0]
	data.ProductIds = data.ProductIds[:0]
	data.UserIds = data.UserIds[:0]
	data.GroupIds = data.GroupIds[:0]
	pool_MSG_USER_updateUserView.Put(data)
}
func (data *MSG_USER_updateUserView) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_updateUserView,buf)
	WRITE_MSG_USER_updateUserView(data, buf)
}

func WRITE_MSG_USER_updateUserView(data *MSG_USER_updateUserView, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.ProjectIds), buf)
	for _, v := range data.ProjectIds{
		WRITE_int32(v, buf)
	}
	WRITE_int(len(data.ProductIds), buf)
	for _, v := range data.ProductIds{
		WRITE_int32(v, buf)
	}
	WRITE_int(len(data.UserIds), buf)
	for _, v := range data.UserIds{
		WRITE_int32(v, buf)
	}
	WRITE_int(len(data.GroupIds), buf)
	for _, v := range data.GroupIds{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_USER_updateUserView(buf *libraries.MsgBuffer) *MSG_USER_updateUserView {
	data := pool_MSG_USER_updateUserView.Get().(*MSG_USER_updateUserView)
	data.read(buf)
	return data
}

func (data *MSG_USER_updateUserView) read(buf *libraries.MsgBuffer) {
	ProjectIds_len := READ_int(buf)
	if ProjectIds_len>cap(data.ProjectIds){
		data.ProjectIds= make([]int32, ProjectIds_len)
	}else{
		data.ProjectIds = data.ProjectIds[:ProjectIds_len]
	}
	for i := 0; i < ProjectIds_len; i++ {
		data.ProjectIds[i] = READ_int32(buf)
	}
	ProductIds_len := READ_int(buf)
	if ProductIds_len>cap(data.ProductIds){
		data.ProductIds= make([]int32, ProductIds_len)
	}else{
		data.ProductIds = data.ProductIds[:ProductIds_len]
	}
	for i := 0; i < ProductIds_len; i++ {
		data.ProductIds[i] = READ_int32(buf)
	}
	UserIds_len := READ_int(buf)
	if UserIds_len>cap(data.UserIds){
		data.UserIds= make([]int32, UserIds_len)
	}else{
		data.UserIds = data.UserIds[:UserIds_len]
	}
	for i := 0; i < UserIds_len; i++ {
		data.UserIds[i] = READ_int32(buf)
	}
	GroupIds_len := READ_int(buf)
	if GroupIds_len>cap(data.GroupIds){
		data.GroupIds= make([]int32, GroupIds_len)
	}else{
		data.GroupIds = data.GroupIds[:GroupIds_len]
	}
	for i := 0; i < GroupIds_len; i++ {
		data.GroupIds[i] = READ_int32(buf)
	}

}

type MSG_USER_getContactLists struct {
	Uid int32
	Params string
}

var pool_MSG_USER_getContactLists = sync.Pool{New: func() interface{} { return &MSG_USER_getContactLists{} }}

func GET_MSG_USER_getContactLists() *MSG_USER_getContactLists {
	return pool_MSG_USER_getContactLists.Get().(*MSG_USER_getContactLists)
}

func (data *MSG_USER_getContactLists) cmd() int32 {
	return CMD_MSG_USER_getContactLists
}

func (data *MSG_USER_getContactLists) Put() {
	data.Uid = 0
	data.Params = ``
	pool_MSG_USER_getContactLists.Put(data)
}
func (data *MSG_USER_getContactLists) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getContactLists,buf)
	WRITE_MSG_USER_getContactLists(data, buf)
}

func WRITE_MSG_USER_getContactLists(data *MSG_USER_getContactLists, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
	WRITE_string(data.Params, buf)
}

func READ_MSG_USER_getContactLists(buf *libraries.MsgBuffer) *MSG_USER_getContactLists {
	data := pool_MSG_USER_getContactLists.Get().(*MSG_USER_getContactLists)
	data.read(buf)
	return data
}

func (data *MSG_USER_getContactLists) read(buf *libraries.MsgBuffer) {
	data.Uid = READ_int32(buf)
	data.Params = READ_string(buf)

}

type MSG_USER_getContactLists_result struct {
	List []HtmlKeyValueStr
}

var pool_MSG_USER_getContactLists_result = sync.Pool{New: func() interface{} { return &MSG_USER_getContactLists_result{} }}

func GET_MSG_USER_getContactLists_result() *MSG_USER_getContactLists_result {
	return pool_MSG_USER_getContactLists_result.Get().(*MSG_USER_getContactLists_result)
}

func (data *MSG_USER_getContactLists_result) cmd() int32 {
	return CMD_MSG_USER_getContactLists_result
}

func (data *MSG_USER_getContactLists_result) Put() {
	data.List = data.List[:0]
	pool_MSG_USER_getContactLists_result.Put(data)
}
func (data *MSG_USER_getContactLists_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getContactLists_result,buf)
	WRITE_MSG_USER_getContactLists_result(data, buf)
}

func WRITE_MSG_USER_getContactLists_result(data *MSG_USER_getContactLists_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_USER_getContactLists_result(buf *libraries.MsgBuffer) *MSG_USER_getContactLists_result {
	data := pool_MSG_USER_getContactLists_result.Get().(*MSG_USER_getContactLists_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_getContactLists_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]HtmlKeyValueStr, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_HtmlKeyValueStr(buf)
	}

}

type MSG_USER_getContactListByUid struct {
	Uid int32
}

var pool_MSG_USER_getContactListByUid = sync.Pool{New: func() interface{} { return &MSG_USER_getContactListByUid{} }}

func GET_MSG_USER_getContactListByUid() *MSG_USER_getContactListByUid {
	return pool_MSG_USER_getContactListByUid.Get().(*MSG_USER_getContactListByUid)
}

func (data *MSG_USER_getContactListByUid) cmd() int32 {
	return CMD_MSG_USER_getContactListByUid
}

func (data *MSG_USER_getContactListByUid) Put() {
	data.Uid = 0
	pool_MSG_USER_getContactListByUid.Put(data)
}
func (data *MSG_USER_getContactListByUid) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getContactListByUid,buf)
	WRITE_MSG_USER_getContactListByUid(data, buf)
}

func WRITE_MSG_USER_getContactListByUid(data *MSG_USER_getContactListByUid, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
}

func READ_MSG_USER_getContactListByUid(buf *libraries.MsgBuffer) *MSG_USER_getContactListByUid {
	data := pool_MSG_USER_getContactListByUid.Get().(*MSG_USER_getContactListByUid)
	data.read(buf)
	return data
}

func (data *MSG_USER_getContactListByUid) read(buf *libraries.MsgBuffer) {
	data.Uid = READ_int32(buf)

}

type MSG_USER_getContactListByUid_result struct {
	List []HtmlKeyValueStr
}

var pool_MSG_USER_getContactListByUid_result = sync.Pool{New: func() interface{} { return &MSG_USER_getContactListByUid_result{} }}

func GET_MSG_USER_getContactListByUid_result() *MSG_USER_getContactListByUid_result {
	return pool_MSG_USER_getContactListByUid_result.Get().(*MSG_USER_getContactListByUid_result)
}

func (data *MSG_USER_getContactListByUid_result) cmd() int32 {
	return CMD_MSG_USER_getContactListByUid_result
}

func (data *MSG_USER_getContactListByUid_result) Put() {
	data.List = data.List[:0]
	pool_MSG_USER_getContactListByUid_result.Put(data)
}
func (data *MSG_USER_getContactListByUid_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getContactListByUid_result,buf)
	WRITE_MSG_USER_getContactListByUid_result(data, buf)
}

func WRITE_MSG_USER_getContactListByUid_result(data *MSG_USER_getContactListByUid_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_USER_getContactListByUid_result(buf *libraries.MsgBuffer) *MSG_USER_getContactListByUid_result {
	data := pool_MSG_USER_getContactListByUid_result.Get().(*MSG_USER_getContactListByUid_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_getContactListByUid_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]HtmlKeyValueStr, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_HtmlKeyValueStr(buf)
	}

}

type MSG_USER_getContactListById struct {
	Id int32
}

var pool_MSG_USER_getContactListById = sync.Pool{New: func() interface{} { return &MSG_USER_getContactListById{} }}

func GET_MSG_USER_getContactListById() *MSG_USER_getContactListById {
	return pool_MSG_USER_getContactListById.Get().(*MSG_USER_getContactListById)
}

func (data *MSG_USER_getContactListById) cmd() int32 {
	return CMD_MSG_USER_getContactListById
}

func (data *MSG_USER_getContactListById) Put() {
	data.Id = 0
	pool_MSG_USER_getContactListById.Put(data)
}
func (data *MSG_USER_getContactListById) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getContactListById,buf)
	WRITE_MSG_USER_getContactListById(data, buf)
}

func WRITE_MSG_USER_getContactListById(data *MSG_USER_getContactListById, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
}

func READ_MSG_USER_getContactListById(buf *libraries.MsgBuffer) *MSG_USER_getContactListById {
	data := pool_MSG_USER_getContactListById.Get().(*MSG_USER_getContactListById)
	data.read(buf)
	return data
}

func (data *MSG_USER_getContactListById) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)

}

type MSG_USER_getContactListById_result struct {
	Result *MSG_USER_ContactList
}

var pool_MSG_USER_getContactListById_result = sync.Pool{New: func() interface{} { return &MSG_USER_getContactListById_result{} }}

func GET_MSG_USER_getContactListById_result() *MSG_USER_getContactListById_result {
	return pool_MSG_USER_getContactListById_result.Get().(*MSG_USER_getContactListById_result)
}

func (data *MSG_USER_getContactListById_result) cmd() int32 {
	return CMD_MSG_USER_getContactListById_result
}

func (data *MSG_USER_getContactListById_result) Put() {
	if data.Result != nil {
		data.Result.Put()
		data.Result = nil
	}
	pool_MSG_USER_getContactListById_result.Put(data)
}
func (data *MSG_USER_getContactListById_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getContactListById_result,buf)
	WRITE_MSG_USER_getContactListById_result(data, buf)
}

func WRITE_MSG_USER_getContactListById_result(data *MSG_USER_getContactListById_result, buf *libraries.MsgBuffer) {
	if data.Result == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_USER_ContactList(data.Result, buf)
	}
}

func READ_MSG_USER_getContactListById_result(buf *libraries.MsgBuffer) *MSG_USER_getContactListById_result {
	data := pool_MSG_USER_getContactListById_result.Get().(*MSG_USER_getContactListById_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_getContactListById_result) read(buf *libraries.MsgBuffer) {
	Result_len := int(READ_int8(buf))
	if Result_len == 1 {
		data.Result = READ_MSG_USER_ContactList(buf)
	}else{
		data.Result = nil
	}

}

type MSG_USER_ContactList struct {
	Id int32
	Uid int32
	ListName string
	UserList []int32
	Share bool
}

var pool_MSG_USER_ContactList = sync.Pool{New: func() interface{} { return &MSG_USER_ContactList{} }}

func GET_MSG_USER_ContactList() *MSG_USER_ContactList {
	return pool_MSG_USER_ContactList.Get().(*MSG_USER_ContactList)
}

func (data *MSG_USER_ContactList) cmd() int32 {
	return CMD_MSG_USER_ContactList
}

func (data *MSG_USER_ContactList) Put() {
	data.Id = 0
	data.Uid = 0
	data.ListName = ``
	data.UserList = data.UserList[:0]
	data.Share = false
	pool_MSG_USER_ContactList.Put(data)
}
func (data *MSG_USER_ContactList) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_ContactList,buf)
	WRITE_MSG_USER_ContactList(data, buf)
}

func WRITE_MSG_USER_ContactList(data *MSG_USER_ContactList, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Uid, buf)
	WRITE_string(data.ListName, buf)
	WRITE_int(len(data.UserList), buf)
	for _, v := range data.UserList{
		WRITE_int32(v, buf)
	}
	WRITE_bool(data.Share, buf)
}

func READ_MSG_USER_ContactList(buf *libraries.MsgBuffer) *MSG_USER_ContactList {
	data := pool_MSG_USER_ContactList.Get().(*MSG_USER_ContactList)
	data.read(buf)
	return data
}

func (data *MSG_USER_ContactList) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Uid = READ_int32(buf)
	data.ListName = READ_string(buf)
	UserList_len := READ_int(buf)
	if UserList_len>cap(data.UserList){
		data.UserList= make([]int32, UserList_len)
	}else{
		data.UserList = data.UserList[:UserList_len]
	}
	for i := 0; i < UserList_len; i++ {
		data.UserList[i] = READ_int32(buf)
	}
	data.Share = READ_bool(buf)

}

type MSG_USER_insertUpdateContactList struct {
	Insert *MSG_USER_ContactList
}

var pool_MSG_USER_insertUpdateContactList = sync.Pool{New: func() interface{} { return &MSG_USER_insertUpdateContactList{} }}

func GET_MSG_USER_insertUpdateContactList() *MSG_USER_insertUpdateContactList {
	return pool_MSG_USER_insertUpdateContactList.Get().(*MSG_USER_insertUpdateContactList)
}

func (data *MSG_USER_insertUpdateContactList) cmd() int32 {
	return CMD_MSG_USER_insertUpdateContactList
}

func (data *MSG_USER_insertUpdateContactList) Put() {
	if data.Insert != nil {
		data.Insert.Put()
		data.Insert = nil
	}
	pool_MSG_USER_insertUpdateContactList.Put(data)
}
func (data *MSG_USER_insertUpdateContactList) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_insertUpdateContactList,buf)
	WRITE_MSG_USER_insertUpdateContactList(data, buf)
}

func WRITE_MSG_USER_insertUpdateContactList(data *MSG_USER_insertUpdateContactList, buf *libraries.MsgBuffer) {
	if data.Insert == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_USER_ContactList(data.Insert, buf)
	}
}

func READ_MSG_USER_insertUpdateContactList(buf *libraries.MsgBuffer) *MSG_USER_insertUpdateContactList {
	data := pool_MSG_USER_insertUpdateContactList.Get().(*MSG_USER_insertUpdateContactList)
	data.read(buf)
	return data
}

func (data *MSG_USER_insertUpdateContactList) read(buf *libraries.MsgBuffer) {
	Insert_len := int(READ_int8(buf))
	if Insert_len == 1 {
		data.Insert = READ_MSG_USER_ContactList(buf)
	}else{
		data.Insert = nil
	}

}

type MSG_USER_insertUpdateContactList_result struct {
	Id int32
}

var pool_MSG_USER_insertUpdateContactList_result = sync.Pool{New: func() interface{} { return &MSG_USER_insertUpdateContactList_result{} }}

func GET_MSG_USER_insertUpdateContactList_result() *MSG_USER_insertUpdateContactList_result {
	return pool_MSG_USER_insertUpdateContactList_result.Get().(*MSG_USER_insertUpdateContactList_result)
}

func (data *MSG_USER_insertUpdateContactList_result) cmd() int32 {
	return CMD_MSG_USER_insertUpdateContactList_result
}

func (data *MSG_USER_insertUpdateContactList_result) Put() {
	data.Id = 0
	pool_MSG_USER_insertUpdateContactList_result.Put(data)
}
func (data *MSG_USER_insertUpdateContactList_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_insertUpdateContactList_result,buf)
	WRITE_MSG_USER_insertUpdateContactList_result(data, buf)
}

func WRITE_MSG_USER_insertUpdateContactList_result(data *MSG_USER_insertUpdateContactList_result, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
}

func READ_MSG_USER_insertUpdateContactList_result(buf *libraries.MsgBuffer) *MSG_USER_insertUpdateContactList_result {
	data := pool_MSG_USER_insertUpdateContactList_result.Get().(*MSG_USER_insertUpdateContactList_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_insertUpdateContactList_result) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)

}

type MSG_USER_getGlobalContacts struct {
}

var pool_MSG_USER_getGlobalContacts = sync.Pool{New: func() interface{} { return &MSG_USER_getGlobalContacts{} }}

func GET_MSG_USER_getGlobalContacts() *MSG_USER_getGlobalContacts {
	return pool_MSG_USER_getGlobalContacts.Get().(*MSG_USER_getGlobalContacts)
}

func (data *MSG_USER_getGlobalContacts) cmd() int32 {
	return CMD_MSG_USER_getGlobalContacts
}

func (data *MSG_USER_getGlobalContacts) Put() {
	pool_MSG_USER_getGlobalContacts.Put(data)
}
func (data *MSG_USER_getGlobalContacts) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getGlobalContacts,buf)
	WRITE_MSG_USER_getGlobalContacts(data, buf)
}

func WRITE_MSG_USER_getGlobalContacts(data *MSG_USER_getGlobalContacts, buf *libraries.MsgBuffer) {
}

func READ_MSG_USER_getGlobalContacts(buf *libraries.MsgBuffer) *MSG_USER_getGlobalContacts {
	data := pool_MSG_USER_getGlobalContacts.Get().(*MSG_USER_getGlobalContacts)
	data.read(buf)
	return data
}

func (data *MSG_USER_getGlobalContacts) read(buf *libraries.MsgBuffer) {

}

type MSG_USER_getGlobalContacts_result struct {
	Result []*MSG_USER_ContactList
}

var pool_MSG_USER_getGlobalContacts_result = sync.Pool{New: func() interface{} { return &MSG_USER_getGlobalContacts_result{} }}

func GET_MSG_USER_getGlobalContacts_result() *MSG_USER_getGlobalContacts_result {
	return pool_MSG_USER_getGlobalContacts_result.Get().(*MSG_USER_getGlobalContacts_result)
}

func (data *MSG_USER_getGlobalContacts_result) cmd() int32 {
	return CMD_MSG_USER_getGlobalContacts_result
}

func (data *MSG_USER_getGlobalContacts_result) Put() {
	for _,v := range data.Result {
		v.Put()
	}
	data.Result = data.Result[:0]
	pool_MSG_USER_getGlobalContacts_result.Put(data)
}
func (data *MSG_USER_getGlobalContacts_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getGlobalContacts_result,buf)
	WRITE_MSG_USER_getGlobalContacts_result(data, buf)
}

func WRITE_MSG_USER_getGlobalContacts_result(data *MSG_USER_getGlobalContacts_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.Result), buf)
	for _, v := range data.Result{
		WRITE_MSG_USER_ContactList(v, buf)
	}
}

func READ_MSG_USER_getGlobalContacts_result(buf *libraries.MsgBuffer) *MSG_USER_getGlobalContacts_result {
	data := pool_MSG_USER_getGlobalContacts_result.Get().(*MSG_USER_getGlobalContacts_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_getGlobalContacts_result) read(buf *libraries.MsgBuffer) {
	Result_len := READ_int(buf)
	if Result_len>cap(data.Result){
		data.Result= make([]*MSG_USER_ContactList, Result_len)
	}else{
		data.Result = data.Result[:Result_len]
	}
	for i := 0; i < Result_len; i++ {
		data.Result[i] = READ_MSG_USER_ContactList(buf)
	}

}

type MSG_USER_team_getByTypeRoot struct {
	Type string
	Root []int32
}

var pool_MSG_USER_team_getByTypeRoot = sync.Pool{New: func() interface{} { return &MSG_USER_team_getByTypeRoot{} }}

func GET_MSG_USER_team_getByTypeRoot() *MSG_USER_team_getByTypeRoot {
	return pool_MSG_USER_team_getByTypeRoot.Get().(*MSG_USER_team_getByTypeRoot)
}

func (data *MSG_USER_team_getByTypeRoot) cmd() int32 {
	return CMD_MSG_USER_team_getByTypeRoot
}

func (data *MSG_USER_team_getByTypeRoot) Put() {
	data.Type = ``
	data.Root = data.Root[:0]
	pool_MSG_USER_team_getByTypeRoot.Put(data)
}
func (data *MSG_USER_team_getByTypeRoot) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_team_getByTypeRoot,buf)
	WRITE_MSG_USER_team_getByTypeRoot(data, buf)
}

func WRITE_MSG_USER_team_getByTypeRoot(data *MSG_USER_team_getByTypeRoot, buf *libraries.MsgBuffer) {
	WRITE_string(data.Type, buf)
	WRITE_int(len(data.Root), buf)
	for _, v := range data.Root{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_USER_team_getByTypeRoot(buf *libraries.MsgBuffer) *MSG_USER_team_getByTypeRoot {
	data := pool_MSG_USER_team_getByTypeRoot.Get().(*MSG_USER_team_getByTypeRoot)
	data.read(buf)
	return data
}

func (data *MSG_USER_team_getByTypeRoot) read(buf *libraries.MsgBuffer) {
	data.Type = READ_string(buf)
	Root_len := READ_int(buf)
	if Root_len>cap(data.Root){
		data.Root= make([]int32, Root_len)
	}else{
		data.Root = data.Root[:Root_len]
	}
	for i := 0; i < Root_len; i++ {
		data.Root[i] = READ_int32(buf)
	}

}

type MSG_USER_team_getByTypeRoot_result struct {
	List []*MSG_USER_team_info
}

var pool_MSG_USER_team_getByTypeRoot_result = sync.Pool{New: func() interface{} { return &MSG_USER_team_getByTypeRoot_result{} }}

func GET_MSG_USER_team_getByTypeRoot_result() *MSG_USER_team_getByTypeRoot_result {
	return pool_MSG_USER_team_getByTypeRoot_result.Get().(*MSG_USER_team_getByTypeRoot_result)
}

func (data *MSG_USER_team_getByTypeRoot_result) cmd() int32 {
	return CMD_MSG_USER_team_getByTypeRoot_result
}

func (data *MSG_USER_team_getByTypeRoot_result) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_USER_team_getByTypeRoot_result.Put(data)
}
func (data *MSG_USER_team_getByTypeRoot_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_team_getByTypeRoot_result,buf)
	WRITE_MSG_USER_team_getByTypeRoot_result(data, buf)
}

func WRITE_MSG_USER_team_getByTypeRoot_result(data *MSG_USER_team_getByTypeRoot_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_USER_team_info(v, buf)
	}
}

func READ_MSG_USER_team_getByTypeRoot_result(buf *libraries.MsgBuffer) *MSG_USER_team_getByTypeRoot_result {
	data := pool_MSG_USER_team_getByTypeRoot_result.Get().(*MSG_USER_team_getByTypeRoot_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_team_getByTypeRoot_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_USER_team_info, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_USER_team_info(buf)
	}

}

type MSG_USER_team_getByIds struct {
	Ids []int32
}

var pool_MSG_USER_team_getByIds = sync.Pool{New: func() interface{} { return &MSG_USER_team_getByIds{} }}

func GET_MSG_USER_team_getByIds() *MSG_USER_team_getByIds {
	return pool_MSG_USER_team_getByIds.Get().(*MSG_USER_team_getByIds)
}

func (data *MSG_USER_team_getByIds) cmd() int32 {
	return CMD_MSG_USER_team_getByIds
}

func (data *MSG_USER_team_getByIds) Put() {
	data.Ids = data.Ids[:0]
	pool_MSG_USER_team_getByIds.Put(data)
}
func (data *MSG_USER_team_getByIds) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_team_getByIds,buf)
	WRITE_MSG_USER_team_getByIds(data, buf)
}

func WRITE_MSG_USER_team_getByIds(data *MSG_USER_team_getByIds, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.Ids), buf)
	for _, v := range data.Ids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_USER_team_getByIds(buf *libraries.MsgBuffer) *MSG_USER_team_getByIds {
	data := pool_MSG_USER_team_getByIds.Get().(*MSG_USER_team_getByIds)
	data.read(buf)
	return data
}

func (data *MSG_USER_team_getByIds) read(buf *libraries.MsgBuffer) {
	Ids_len := READ_int(buf)
	if Ids_len>cap(data.Ids){
		data.Ids= make([]int32, Ids_len)
	}else{
		data.Ids = data.Ids[:Ids_len]
	}
	for i := 0; i < Ids_len; i++ {
		data.Ids[i] = READ_int32(buf)
	}

}

type MSG_USER_team_getByIds_result struct {
	List []*MSG_USER_team_info
}

var pool_MSG_USER_team_getByIds_result = sync.Pool{New: func() interface{} { return &MSG_USER_team_getByIds_result{} }}

func GET_MSG_USER_team_getByIds_result() *MSG_USER_team_getByIds_result {
	return pool_MSG_USER_team_getByIds_result.Get().(*MSG_USER_team_getByIds_result)
}

func (data *MSG_USER_team_getByIds_result) cmd() int32 {
	return CMD_MSG_USER_team_getByIds_result
}

func (data *MSG_USER_team_getByIds_result) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_USER_team_getByIds_result.Put(data)
}
func (data *MSG_USER_team_getByIds_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_team_getByIds_result,buf)
	WRITE_MSG_USER_team_getByIds_result(data, buf)
}

func WRITE_MSG_USER_team_getByIds_result(data *MSG_USER_team_getByIds_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_USER_team_info(v, buf)
	}
}

func READ_MSG_USER_team_getByIds_result(buf *libraries.MsgBuffer) *MSG_USER_team_getByIds_result {
	data := pool_MSG_USER_team_getByIds_result.Get().(*MSG_USER_team_getByIds_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_team_getByIds_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_USER_team_info, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_USER_team_info(buf)
	}

}

type MSG_USER_team_info struct {
	Id int32
	Root int32
	Type string
	Uid int32
	Account string
	Role string
	Limited string
	Join time.Time
	Days int16
	Hours float64
	Estimate float64
	Consumed float64
	Left float64
	Order int8
	Deleted bool `db:"-"`
	Realname string `db:"-"`
}

var pool_MSG_USER_team_info = sync.Pool{New: func() interface{} { return &MSG_USER_team_info{} }}

func GET_MSG_USER_team_info() *MSG_USER_team_info {
	return pool_MSG_USER_team_info.Get().(*MSG_USER_team_info)
}

func (data *MSG_USER_team_info) cmd() int32 {
	return CMD_MSG_USER_team_info
}

func (data *MSG_USER_team_info) Put() {
	data.Id = 0
	data.Root = 0
	data.Type = ``
	data.Uid = 0
	data.Account = ``
	data.Role = ``
	data.Limited = ``
	data.Join = time.UnixMicro(0)
	data.Days = 0
	data.Hours = 0
	data.Estimate = 0
	data.Consumed = 0
	data.Left = 0
	data.Order = 0
	data.Deleted = false
	data.Realname = ``
	pool_MSG_USER_team_info.Put(data)
}
func (data *MSG_USER_team_info) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_team_info,buf)
	WRITE_MSG_USER_team_info(data, buf)
}

func WRITE_MSG_USER_team_info(data *MSG_USER_team_info, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Root, buf)
	WRITE_string(data.Type, buf)
	WRITE_int32(data.Uid, buf)
	WRITE_string(data.Account, buf)
	WRITE_string(data.Role, buf)
	WRITE_string(data.Limited, buf)
	WRITE_int64(data.Join.UnixMicro(), buf)
	WRITE_int16(data.Days, buf)
	WRITE_float64(data.Hours, buf)
	WRITE_float64(data.Estimate, buf)
	WRITE_float64(data.Consumed, buf)
	WRITE_float64(data.Left, buf)
	WRITE_int8(data.Order, buf)
	WRITE_bool(data.Deleted, buf)
	WRITE_string(data.Realname, buf)
}

func READ_MSG_USER_team_info(buf *libraries.MsgBuffer) *MSG_USER_team_info {
	data := pool_MSG_USER_team_info.Get().(*MSG_USER_team_info)
	data.read(buf)
	return data
}

func (data *MSG_USER_team_info) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Root = READ_int32(buf)
	data.Type = READ_string(buf)
	data.Uid = READ_int32(buf)
	data.Account = READ_string(buf)
	data.Role = READ_string(buf)
	data.Limited = READ_string(buf)
	data.Join = time.UnixMicro(READ_int64(buf))
	data.Days = READ_int16(buf)
	data.Hours = READ_float64(buf)
	data.Estimate = READ_float64(buf)
	data.Consumed = READ_float64(buf)
	data.Left = READ_float64(buf)
	data.Order = READ_int8(buf)
	data.Deleted = READ_bool(buf)
	data.Realname = READ_string(buf)

}

type MSG_USER_team_addByList struct {
	List []*MSG_USER_team_info
}

var pool_MSG_USER_team_addByList = sync.Pool{New: func() interface{} { return &MSG_USER_team_addByList{} }}

func GET_MSG_USER_team_addByList() *MSG_USER_team_addByList {
	return pool_MSG_USER_team_addByList.Get().(*MSG_USER_team_addByList)
}

func (data *MSG_USER_team_addByList) cmd() int32 {
	return CMD_MSG_USER_team_addByList
}

func (data *MSG_USER_team_addByList) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_USER_team_addByList.Put(data)
}
func (data *MSG_USER_team_addByList) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_team_addByList,buf)
	WRITE_MSG_USER_team_addByList(data, buf)
}

func WRITE_MSG_USER_team_addByList(data *MSG_USER_team_addByList, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_USER_team_info(v, buf)
	}
}

func READ_MSG_USER_team_addByList(buf *libraries.MsgBuffer) *MSG_USER_team_addByList {
	data := pool_MSG_USER_team_addByList.Get().(*MSG_USER_team_addByList)
	data.read(buf)
	return data
}

func (data *MSG_USER_team_addByList) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_USER_team_info, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_USER_team_info(buf)
	}

}

type MSG_USER_Group_getPairs struct {
}

var pool_MSG_USER_Group_getPairs = sync.Pool{New: func() interface{} { return &MSG_USER_Group_getPairs{} }}

func GET_MSG_USER_Group_getPairs() *MSG_USER_Group_getPairs {
	return pool_MSG_USER_Group_getPairs.Get().(*MSG_USER_Group_getPairs)
}

func (data *MSG_USER_Group_getPairs) cmd() int32 {
	return CMD_MSG_USER_Group_getPairs
}

func (data *MSG_USER_Group_getPairs) Put() {
	pool_MSG_USER_Group_getPairs.Put(data)
}
func (data *MSG_USER_Group_getPairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Group_getPairs,buf)
	WRITE_MSG_USER_Group_getPairs(data, buf)
}

func WRITE_MSG_USER_Group_getPairs(data *MSG_USER_Group_getPairs, buf *libraries.MsgBuffer) {
}

func READ_MSG_USER_Group_getPairs(buf *libraries.MsgBuffer) *MSG_USER_Group_getPairs {
	data := pool_MSG_USER_Group_getPairs.Get().(*MSG_USER_Group_getPairs)
	data.read(buf)
	return data
}

func (data *MSG_USER_Group_getPairs) read(buf *libraries.MsgBuffer) {

}

type MSG_USER_Group_getPairs_result struct {
	List []HtmlKeyValueStr
}

var pool_MSG_USER_Group_getPairs_result = sync.Pool{New: func() interface{} { return &MSG_USER_Group_getPairs_result{} }}

func GET_MSG_USER_Group_getPairs_result() *MSG_USER_Group_getPairs_result {
	return pool_MSG_USER_Group_getPairs_result.Get().(*MSG_USER_Group_getPairs_result)
}

func (data *MSG_USER_Group_getPairs_result) cmd() int32 {
	return CMD_MSG_USER_Group_getPairs_result
}

func (data *MSG_USER_Group_getPairs_result) Put() {
	data.List = data.List[:0]
	pool_MSG_USER_Group_getPairs_result.Put(data)
}
func (data *MSG_USER_Group_getPairs_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Group_getPairs_result,buf)
	WRITE_MSG_USER_Group_getPairs_result(data, buf)
}

func WRITE_MSG_USER_Group_getPairs_result(data *MSG_USER_Group_getPairs_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_USER_Group_getPairs_result(buf *libraries.MsgBuffer) *MSG_USER_Group_getPairs_result {
	data := pool_MSG_USER_Group_getPairs_result.Get().(*MSG_USER_Group_getPairs_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_Group_getPairs_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]HtmlKeyValueStr, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_HtmlKeyValueStr(buf)
	}

}

type MSG_USER_team_getByTypeUid struct {
	Type string
	Uid int32
}

var pool_MSG_USER_team_getByTypeUid = sync.Pool{New: func() interface{} { return &MSG_USER_team_getByTypeUid{} }}

func GET_MSG_USER_team_getByTypeUid() *MSG_USER_team_getByTypeUid {
	return pool_MSG_USER_team_getByTypeUid.Get().(*MSG_USER_team_getByTypeUid)
}

func (data *MSG_USER_team_getByTypeUid) cmd() int32 {
	return CMD_MSG_USER_team_getByTypeUid
}

func (data *MSG_USER_team_getByTypeUid) Put() {
	data.Type = ``
	data.Uid = 0
	pool_MSG_USER_team_getByTypeUid.Put(data)
}
func (data *MSG_USER_team_getByTypeUid) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_team_getByTypeUid,buf)
	WRITE_MSG_USER_team_getByTypeUid(data, buf)
}

func WRITE_MSG_USER_team_getByTypeUid(data *MSG_USER_team_getByTypeUid, buf *libraries.MsgBuffer) {
	WRITE_string(data.Type, buf)
	WRITE_int32(data.Uid, buf)
}

func READ_MSG_USER_team_getByTypeUid(buf *libraries.MsgBuffer) *MSG_USER_team_getByTypeUid {
	data := pool_MSG_USER_team_getByTypeUid.Get().(*MSG_USER_team_getByTypeUid)
	data.read(buf)
	return data
}

func (data *MSG_USER_team_getByTypeUid) read(buf *libraries.MsgBuffer) {
	data.Type = READ_string(buf)
	data.Uid = READ_int32(buf)

}

type MSG_USER_team_getByTypeUid_result struct {
	List []*MSG_USER_team_info
}

var pool_MSG_USER_team_getByTypeUid_result = sync.Pool{New: func() interface{} { return &MSG_USER_team_getByTypeUid_result{} }}

func GET_MSG_USER_team_getByTypeUid_result() *MSG_USER_team_getByTypeUid_result {
	return pool_MSG_USER_team_getByTypeUid_result.Get().(*MSG_USER_team_getByTypeUid_result)
}

func (data *MSG_USER_team_getByTypeUid_result) cmd() int32 {
	return CMD_MSG_USER_team_getByTypeUid_result
}

func (data *MSG_USER_team_getByTypeUid_result) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_USER_team_getByTypeUid_result.Put(data)
}
func (data *MSG_USER_team_getByTypeUid_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_team_getByTypeUid_result,buf)
	WRITE_MSG_USER_team_getByTypeUid_result(data, buf)
}

func WRITE_MSG_USER_team_getByTypeUid_result(data *MSG_USER_team_getByTypeUid_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_USER_team_info(v, buf)
	}
}

func READ_MSG_USER_team_getByTypeUid_result(buf *libraries.MsgBuffer) *MSG_USER_team_getByTypeUid_result {
	data := pool_MSG_USER_team_getByTypeUid_result.Get().(*MSG_USER_team_getByTypeUid_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_team_getByTypeUid_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_USER_team_info, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_USER_team_info(buf)
	}

}

type MSG_USER_Userquery_info struct {
	Id int32
	Uid int32
	Module string
	Title string
	Form string
	Sql string
	Shortcut bool
}

var pool_MSG_USER_Userquery_info = sync.Pool{New: func() interface{} { return &MSG_USER_Userquery_info{} }}

func GET_MSG_USER_Userquery_info() *MSG_USER_Userquery_info {
	return pool_MSG_USER_Userquery_info.Get().(*MSG_USER_Userquery_info)
}

func (data *MSG_USER_Userquery_info) cmd() int32 {
	return CMD_MSG_USER_Userquery_info
}

func (data *MSG_USER_Userquery_info) Put() {
	data.Id = 0
	data.Uid = 0
	data.Module = ``
	data.Title = ``
	data.Form = ``
	data.Sql = ``
	data.Shortcut = false
	pool_MSG_USER_Userquery_info.Put(data)
}
func (data *MSG_USER_Userquery_info) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Userquery_info,buf)
	WRITE_MSG_USER_Userquery_info(data, buf)
}

func WRITE_MSG_USER_Userquery_info(data *MSG_USER_Userquery_info, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Uid, buf)
	WRITE_string(data.Module, buf)
	WRITE_string(data.Title, buf)
	WRITE_string(data.Form, buf)
	WRITE_string(data.Sql, buf)
	WRITE_bool(data.Shortcut, buf)
}

func READ_MSG_USER_Userquery_info(buf *libraries.MsgBuffer) *MSG_USER_Userquery_info {
	data := pool_MSG_USER_Userquery_info.Get().(*MSG_USER_Userquery_info)
	data.read(buf)
	return data
}

func (data *MSG_USER_Userquery_info) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Uid = READ_int32(buf)
	data.Module = READ_string(buf)
	data.Title = READ_string(buf)
	data.Form = READ_string(buf)
	data.Sql = READ_string(buf)
	data.Shortcut = READ_bool(buf)

}

type MSG_USER_user_getUserqueryByWhere struct {
	Where map[string]interface{}
}

var pool_MSG_USER_user_getUserqueryByWhere = sync.Pool{New: func() interface{} { return &MSG_USER_user_getUserqueryByWhere{} }}

func GET_MSG_USER_user_getUserqueryByWhere() *MSG_USER_user_getUserqueryByWhere {
	return pool_MSG_USER_user_getUserqueryByWhere.Get().(*MSG_USER_user_getUserqueryByWhere)
}

func (data *MSG_USER_user_getUserqueryByWhere) cmd() int32 {
	return CMD_MSG_USER_user_getUserqueryByWhere
}

func (data *MSG_USER_user_getUserqueryByWhere) Put() {
	data.Where = nil
	pool_MSG_USER_user_getUserqueryByWhere.Put(data)
}
func (data *MSG_USER_user_getUserqueryByWhere) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_user_getUserqueryByWhere,buf)
	WRITE_MSG_USER_user_getUserqueryByWhere(data, buf)
}

func WRITE_MSG_USER_user_getUserqueryByWhere(data *MSG_USER_user_getUserqueryByWhere, buf *libraries.MsgBuffer) {
	WRITE_map(data.Where,buf)
}

func READ_MSG_USER_user_getUserqueryByWhere(buf *libraries.MsgBuffer) *MSG_USER_user_getUserqueryByWhere {
	data := pool_MSG_USER_user_getUserqueryByWhere.Get().(*MSG_USER_user_getUserqueryByWhere)
	data.read(buf)
	return data
}

func (data *MSG_USER_user_getUserqueryByWhere) read(buf *libraries.MsgBuffer) {
	READ_map(&data.Where,buf)

}

type MSG_USER_user_getUserqueryByWhere_result struct {
	List []*MSG_USER_Userquery_info
}

var pool_MSG_USER_user_getUserqueryByWhere_result = sync.Pool{New: func() interface{} { return &MSG_USER_user_getUserqueryByWhere_result{} }}

func GET_MSG_USER_user_getUserqueryByWhere_result() *MSG_USER_user_getUserqueryByWhere_result {
	return pool_MSG_USER_user_getUserqueryByWhere_result.Get().(*MSG_USER_user_getUserqueryByWhere_result)
}

func (data *MSG_USER_user_getUserqueryByWhere_result) cmd() int32 {
	return CMD_MSG_USER_user_getUserqueryByWhere_result
}

func (data *MSG_USER_user_getUserqueryByWhere_result) Put() {
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_USER_user_getUserqueryByWhere_result.Put(data)
}
func (data *MSG_USER_user_getUserqueryByWhere_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_user_getUserqueryByWhere_result,buf)
	WRITE_MSG_USER_user_getUserqueryByWhere_result(data, buf)
}

func WRITE_MSG_USER_user_getUserqueryByWhere_result(data *MSG_USER_user_getUserqueryByWhere_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_USER_Userquery_info(v, buf)
	}
}

func READ_MSG_USER_user_getUserqueryByWhere_result(buf *libraries.MsgBuffer) *MSG_USER_user_getUserqueryByWhere_result {
	data := pool_MSG_USER_user_getUserqueryByWhere_result.Get().(*MSG_USER_user_getUserqueryByWhere_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_user_getUserqueryByWhere_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_USER_Userquery_info, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_USER_Userquery_info(buf)
	}

}

type MSG_USER_team_getMemberPairsByTypeRoot struct {
	Type string
	Root int32
}

var pool_MSG_USER_team_getMemberPairsByTypeRoot = sync.Pool{New: func() interface{} { return &MSG_USER_team_getMemberPairsByTypeRoot{} }}

func GET_MSG_USER_team_getMemberPairsByTypeRoot() *MSG_USER_team_getMemberPairsByTypeRoot {
	return pool_MSG_USER_team_getMemberPairsByTypeRoot.Get().(*MSG_USER_team_getMemberPairsByTypeRoot)
}

func (data *MSG_USER_team_getMemberPairsByTypeRoot) cmd() int32 {
	return CMD_MSG_USER_team_getMemberPairsByTypeRoot
}

func (data *MSG_USER_team_getMemberPairsByTypeRoot) Put() {
	data.Type = ``
	data.Root = 0
	pool_MSG_USER_team_getMemberPairsByTypeRoot.Put(data)
}
func (data *MSG_USER_team_getMemberPairsByTypeRoot) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_team_getMemberPairsByTypeRoot,buf)
	WRITE_MSG_USER_team_getMemberPairsByTypeRoot(data, buf)
}

func WRITE_MSG_USER_team_getMemberPairsByTypeRoot(data *MSG_USER_team_getMemberPairsByTypeRoot, buf *libraries.MsgBuffer) {
	WRITE_string(data.Type, buf)
	WRITE_int32(data.Root, buf)
}

func READ_MSG_USER_team_getMemberPairsByTypeRoot(buf *libraries.MsgBuffer) *MSG_USER_team_getMemberPairsByTypeRoot {
	data := pool_MSG_USER_team_getMemberPairsByTypeRoot.Get().(*MSG_USER_team_getMemberPairsByTypeRoot)
	data.read(buf)
	return data
}

func (data *MSG_USER_team_getMemberPairsByTypeRoot) read(buf *libraries.MsgBuffer) {
	data.Type = READ_string(buf)
	data.Root = READ_int32(buf)

}

type MSG_USER_team_getMemberPairsByTypeRoot_result struct {
	List []HtmlKeyValueStr
}

var pool_MSG_USER_team_getMemberPairsByTypeRoot_result = sync.Pool{New: func() interface{} { return &MSG_USER_team_getMemberPairsByTypeRoot_result{} }}

func GET_MSG_USER_team_getMemberPairsByTypeRoot_result() *MSG_USER_team_getMemberPairsByTypeRoot_result {
	return pool_MSG_USER_team_getMemberPairsByTypeRoot_result.Get().(*MSG_USER_team_getMemberPairsByTypeRoot_result)
}

func (data *MSG_USER_team_getMemberPairsByTypeRoot_result) cmd() int32 {
	return CMD_MSG_USER_team_getMemberPairsByTypeRoot_result
}

func (data *MSG_USER_team_getMemberPairsByTypeRoot_result) Put() {
	data.List = data.List[:0]
	pool_MSG_USER_team_getMemberPairsByTypeRoot_result.Put(data)
}
func (data *MSG_USER_team_getMemberPairsByTypeRoot_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_team_getMemberPairsByTypeRoot_result,buf)
	WRITE_MSG_USER_team_getMemberPairsByTypeRoot_result(data, buf)
}

func WRITE_MSG_USER_team_getMemberPairsByTypeRoot_result(data *MSG_USER_team_getMemberPairsByTypeRoot_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_HtmlKeyValueStr(v, buf)
	}
}

func READ_MSG_USER_team_getMemberPairsByTypeRoot_result(buf *libraries.MsgBuffer) *MSG_USER_team_getMemberPairsByTypeRoot_result {
	data := pool_MSG_USER_team_getMemberPairsByTypeRoot_result.Get().(*MSG_USER_team_getMemberPairsByTypeRoot_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_team_getMemberPairsByTypeRoot_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]HtmlKeyValueStr, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_HtmlKeyValueStr(buf)
	}

}

type MSG_USER_team_updateByWhere struct {
	Where map[string]interface{}
	Update map[string]interface{}
}

var pool_MSG_USER_team_updateByWhere = sync.Pool{New: func() interface{} { return &MSG_USER_team_updateByWhere{} }}

func GET_MSG_USER_team_updateByWhere() *MSG_USER_team_updateByWhere {
	return pool_MSG_USER_team_updateByWhere.Get().(*MSG_USER_team_updateByWhere)
}

func (data *MSG_USER_team_updateByWhere) cmd() int32 {
	return CMD_MSG_USER_team_updateByWhere
}

func (data *MSG_USER_team_updateByWhere) Put() {
	data.Where = nil
	data.Update = nil
	pool_MSG_USER_team_updateByWhere.Put(data)
}
func (data *MSG_USER_team_updateByWhere) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_team_updateByWhere,buf)
	WRITE_MSG_USER_team_updateByWhere(data, buf)
}

func WRITE_MSG_USER_team_updateByWhere(data *MSG_USER_team_updateByWhere, buf *libraries.MsgBuffer) {
	WRITE_map(data.Where,buf)
	WRITE_map(data.Update,buf)
}

func READ_MSG_USER_team_updateByWhere(buf *libraries.MsgBuffer) *MSG_USER_team_updateByWhere {
	data := pool_MSG_USER_team_updateByWhere.Get().(*MSG_USER_team_updateByWhere)
	data.read(buf)
	return data
}

func (data *MSG_USER_team_updateByWhere) read(buf *libraries.MsgBuffer) {
	READ_map(&data.Where,buf)
	READ_map(&data.Update,buf)

}

type MSG_USER_config_save struct {
	Uid int32
	Module string
	Section string
	Key string
	Value string
	Type string
}

var pool_MSG_USER_config_save = sync.Pool{New: func() interface{} { return &MSG_USER_config_save{} }}

func GET_MSG_USER_config_save() *MSG_USER_config_save {
	return pool_MSG_USER_config_save.Get().(*MSG_USER_config_save)
}

func (data *MSG_USER_config_save) cmd() int32 {
	return CMD_MSG_USER_config_save
}

func (data *MSG_USER_config_save) Put() {
	data.Uid = 0
	data.Module = ``
	data.Section = ``
	data.Key = ``
	data.Value = ``
	data.Type = ``
	pool_MSG_USER_config_save.Put(data)
}
func (data *MSG_USER_config_save) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_config_save,buf)
	WRITE_MSG_USER_config_save(data, buf)
}

func WRITE_MSG_USER_config_save(data *MSG_USER_config_save, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
	WRITE_string(data.Module, buf)
	WRITE_string(data.Section, buf)
	WRITE_string(data.Key, buf)
	WRITE_string(data.Value, buf)
	WRITE_string(data.Type, buf)
}

func READ_MSG_USER_config_save(buf *libraries.MsgBuffer) *MSG_USER_config_save {
	data := pool_MSG_USER_config_save.Get().(*MSG_USER_config_save)
	data.read(buf)
	return data
}

func (data *MSG_USER_config_save) read(buf *libraries.MsgBuffer) {
	data.Uid = READ_int32(buf)
	data.Module = READ_string(buf)
	data.Section = READ_string(buf)
	data.Key = READ_string(buf)
	data.Value = READ_string(buf)
	data.Type = READ_string(buf)

}

