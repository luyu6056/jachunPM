package protocol

import (
	"sync"
	"libraries"
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
)

type MSG_USER_GET_LoginSalt struct {
	QueryID uint32
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
	data.QueryID = 0
	data.Name = ``
	pool_MSG_USER_GET_LoginSalt.Put(data)
}
func (data *MSG_USER_GET_LoginSalt) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_GET_LoginSalt,buf)
	WRITE_MSG_USER_GET_LoginSalt(data, buf)
}

func WRITE_MSG_USER_GET_LoginSalt(data *MSG_USER_GET_LoginSalt, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_USER_GET_LoginSalt(buf *libraries.MsgBuffer) *MSG_USER_GET_LoginSalt {
	data := pool_MSG_USER_GET_LoginSalt.Get().(*MSG_USER_GET_LoginSalt)
	data.read(buf)
	return data
}

func (data *MSG_USER_GET_LoginSalt) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Name = READ_string(buf)

}
func (data *MSG_USER_GET_LoginSalt) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_USER_GET_LoginSalt) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_USER_GET_LoginSalt_result struct {
	QueryResultID uint32
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
	data.QueryResultID = 0
	data.Salt = ``
	pool_MSG_USER_GET_LoginSalt_result.Put(data)
}
func (data *MSG_USER_GET_LoginSalt_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_GET_LoginSalt_result,buf)
	WRITE_MSG_USER_GET_LoginSalt_result(data, buf)
}

func WRITE_MSG_USER_GET_LoginSalt_result(data *MSG_USER_GET_LoginSalt_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_string(data.Salt, buf)
}

func READ_MSG_USER_GET_LoginSalt_result(buf *libraries.MsgBuffer) *MSG_USER_GET_LoginSalt_result {
	data := pool_MSG_USER_GET_LoginSalt_result.Get().(*MSG_USER_GET_LoginSalt_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_GET_LoginSalt_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Salt = READ_string(buf)

}
func (data *MSG_USER_GET_LoginSalt_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_USER_GET_LoginSalt_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_USER_INFO_cache struct {
	Id int32
	Dept int32
	Account string
	Role string
	Realname string
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

}

type MSG_USER_CheckPasswd struct {
	QueryID uint32
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
	data.QueryID = 0
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
	WRITE_uint32(data.QueryID, buf)
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
	data.QueryID = READ_uint32(buf)
	data.UserId = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Rand = READ_int64(buf)
	data.Passwd = READ_string(buf)

}
func (data *MSG_USER_CheckPasswd) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_USER_CheckPasswd) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_USER_CheckPasswd_result struct {
	QueryResultID uint32
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
	data.QueryResultID = 0
	data.UserId = 0
	pool_MSG_USER_CheckPasswd_result.Put(data)
}
func (data *MSG_USER_CheckPasswd_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_CheckPasswd_result,buf)
	WRITE_MSG_USER_CheckPasswd_result(data, buf)
}

func WRITE_MSG_USER_CheckPasswd_result(data *MSG_USER_CheckPasswd_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(data.UserId, buf)
	WRITE_ErrCode(data.Result, buf)
}

func READ_MSG_USER_CheckPasswd_result(buf *libraries.MsgBuffer) *MSG_USER_CheckPasswd_result {
	data := pool_MSG_USER_CheckPasswd_result.Get().(*MSG_USER_CheckPasswd_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_CheckPasswd_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.UserId = READ_int32(buf)
	data.Result = READ_ErrCode(buf)

}
func (data *MSG_USER_CheckPasswd_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_USER_CheckPasswd_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
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
	WRITE_int32(int32(len(data.Admins)), buf)
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
	Admins_len := int(READ_int32(buf))
	for i := 0; i < Admins_len; i++ {
		data.Admins = append(data.Admins, READ_string(buf))
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
	Children []*MSG_USER_Dept_cache
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
	WRITE_int32(int32(len(data.Path)), buf)
	for _, v := range data.Path{
		WRITE_int32(v, buf)
	}
	WRITE_int8(data.Grade, buf)
	WRITE_int8(data.Order, buf)
	WRITE_int32(data.Manager, buf)
	WRITE_string(data.ManagerName, buf)
	WRITE_int32(int32(len(data.Children)), buf)
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
	Path_len := int(READ_int32(buf))
	for i := 0; i < Path_len; i++ {
		data.Path = append(data.Path, READ_int32(buf))
	}
	data.Grade = READ_int8(buf)
	data.Order = READ_int8(buf)
	data.Manager = READ_int32(buf)
	data.ManagerName = READ_string(buf)
	Children_len := int(READ_int32(buf))
	for i := 0; i < Children_len; i++ {
		data.Children = append(data.Children, READ_MSG_USER_Dept_cache(buf))
	}

}

type MSG_USER_Dept_getParents struct {
	QueryID uint32
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
	data.QueryID = 0
	data.Id = 0
	pool_MSG_USER_Dept_getParents.Put(data)
}
func (data *MSG_USER_Dept_getParents) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_getParents,buf)
	WRITE_MSG_USER_Dept_getParents(data, buf)
}

func WRITE_MSG_USER_Dept_getParents(data *MSG_USER_Dept_getParents, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.Id, buf)
}

func READ_MSG_USER_Dept_getParents(buf *libraries.MsgBuffer) *MSG_USER_Dept_getParents {
	data := pool_MSG_USER_Dept_getParents.Get().(*MSG_USER_Dept_getParents)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_getParents) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Id = READ_int32(buf)

}
func (data *MSG_USER_Dept_getParents) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_USER_Dept_getParents) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_USER_Dept_getParents_result struct {
	QueryResultID uint32
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
	data.QueryResultID = 0
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
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
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
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	for i := 0; i < List_len; i++ {
		data.List = append(data.List, READ_MSG_USER_Dept_cache(buf))
	}

}
func (data *MSG_USER_Dept_getParents_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_USER_Dept_getParents_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_USER_Dept_getDataStructure struct {
	QueryID uint32
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
	data.QueryID = 0
	data.RootDeptID = 0
	pool_MSG_USER_Dept_getDataStructure.Put(data)
}
func (data *MSG_USER_Dept_getDataStructure) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_getDataStructure,buf)
	WRITE_MSG_USER_Dept_getDataStructure(data, buf)
}

func WRITE_MSG_USER_Dept_getDataStructure(data *MSG_USER_Dept_getDataStructure, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.RootDeptID, buf)
}

func READ_MSG_USER_Dept_getDataStructure(buf *libraries.MsgBuffer) *MSG_USER_Dept_getDataStructure {
	data := pool_MSG_USER_Dept_getDataStructure.Get().(*MSG_USER_Dept_getDataStructure)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_getDataStructure) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.RootDeptID = READ_int32(buf)

}
func (data *MSG_USER_Dept_getDataStructure) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_USER_Dept_getDataStructure) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_USER_Dept_getDataStructure_result struct {
	QueryResultID uint32
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
	data.QueryResultID = 0
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
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
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
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	for i := 0; i < List_len; i++ {
		data.List = append(data.List, READ_MSG_USER_Dept_cache(buf))
	}

}
func (data *MSG_USER_Dept_getDataStructure_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_USER_Dept_getDataStructure_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_USER_Dept_update struct {
	QueryID uint32
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
	data.QueryID = 0
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
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(int32(len(data.List)), buf)
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
	data.QueryID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	for i := 0; i < List_len; i++ {
		data.List = append(data.List, READ_MSG_USER_Dept_cache(buf))
	}

}
func (data *MSG_USER_Dept_update) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_USER_Dept_update) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_USER_Dept_delete struct {
	QueryID uint32
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
	data.QueryID = 0
	data.DeptId = 0
	pool_MSG_USER_Dept_delete.Put(data)
}
func (data *MSG_USER_Dept_delete) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_delete,buf)
	WRITE_MSG_USER_Dept_delete(data, buf)
}

func WRITE_MSG_USER_Dept_delete(data *MSG_USER_Dept_delete, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.DeptId, buf)
}

func READ_MSG_USER_Dept_delete(buf *libraries.MsgBuffer) *MSG_USER_Dept_delete {
	data := pool_MSG_USER_Dept_delete.Get().(*MSG_USER_Dept_delete)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_delete) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.DeptId = READ_int32(buf)

}
func (data *MSG_USER_Dept_delete) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_USER_Dept_delete) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_USER_Dept_delete_result struct {
	QueryResultID uint32
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
	data.QueryResultID = 0
	pool_MSG_USER_Dept_delete_result.Put(data)
}
func (data *MSG_USER_Dept_delete_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_Dept_delete_result,buf)
	WRITE_MSG_USER_Dept_delete_result(data, buf)
}

func WRITE_MSG_USER_Dept_delete_result(data *MSG_USER_Dept_delete_result, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_ErrCode(data.Result, buf)
}

func READ_MSG_USER_Dept_delete_result(buf *libraries.MsgBuffer) *MSG_USER_Dept_delete_result {
	data := pool_MSG_USER_Dept_delete_result.Get().(*MSG_USER_Dept_delete_result)
	data.read(buf)
	return data
}

func (data *MSG_USER_Dept_delete_result) read(buf *libraries.MsgBuffer) {
	data.QueryResultID = READ_uint32(buf)
	data.Result = READ_ErrCode(buf)

}
func (data *MSG_USER_Dept_delete_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_USER_Dept_delete_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
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
	QueryID uint32
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
	data.QueryID = 0
	data.DeptId = 0
	pool_MSG_USER_getDeptUserPairs.Put(data)
}
func (data *MSG_USER_getDeptUserPairs) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getDeptUserPairs,buf)
	WRITE_MSG_USER_getDeptUserPairs(data, buf)
}

func WRITE_MSG_USER_getDeptUserPairs(data *MSG_USER_getDeptUserPairs, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_int32(data.DeptId, buf)
}

func READ_MSG_USER_getDeptUserPairs(buf *libraries.MsgBuffer) *MSG_USER_getDeptUserPairs {
	data := pool_MSG_USER_getDeptUserPairs.Get().(*MSG_USER_getDeptUserPairs)
	data.read(buf)
	return data
}

func (data *MSG_USER_getDeptUserPairs) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.DeptId = READ_int32(buf)

}
func (data *MSG_USER_getDeptUserPairs) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_USER_getDeptUserPairs) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_USER_getDeptUserPairs_result struct {
	QueryResultID uint32
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
	data.QueryResultID = 0
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
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
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
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	for i := 0; i < List_len; i++ {
		data.List = append(data.List, READ_MSG_USER_Pairs(buf))
	}

}
func (data *MSG_USER_getDeptUserPairs_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_USER_getDeptUserPairs_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

type MSG_USER_getCompanyUsers struct {
	QueryID uint32
	Type string
	Query string
	DeptID int32
	Sort string
	Page int
	PerPage int
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
	data.QueryID = 0
	data.Type = ``
	data.Query = ``
	data.DeptID = 0
	data.Sort = ``
	data.Page = 0
	data.PerPage = 0
	data.Total = 0
	pool_MSG_USER_getCompanyUsers.Put(data)
}
func (data *MSG_USER_getCompanyUsers) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_USER_getCompanyUsers,buf)
	WRITE_MSG_USER_getCompanyUsers(data, buf)
}

func WRITE_MSG_USER_getCompanyUsers(data *MSG_USER_getCompanyUsers, buf *libraries.MsgBuffer) {
	WRITE_uint32(data.QueryID, buf)
	WRITE_string(data.Type, buf)
	WRITE_string(data.Query, buf)
	WRITE_int32(data.DeptID, buf)
	WRITE_string(data.Sort, buf)
	WRITE_int(data.Page, buf)
	WRITE_int(data.PerPage, buf)
	WRITE_int(data.Total, buf)
}

func READ_MSG_USER_getCompanyUsers(buf *libraries.MsgBuffer) *MSG_USER_getCompanyUsers {
	data := pool_MSG_USER_getCompanyUsers.Get().(*MSG_USER_getCompanyUsers)
	data.read(buf)
	return data
}

func (data *MSG_USER_getCompanyUsers) read(buf *libraries.MsgBuffer) {
	data.QueryID = READ_uint32(buf)
	data.Type = READ_string(buf)
	data.Query = READ_string(buf)
	data.DeptID = READ_int32(buf)
	data.Sort = READ_string(buf)
	data.Page = READ_int(buf)
	data.PerPage = READ_int(buf)
	data.Total = READ_int(buf)

}
func (data *MSG_USER_getCompanyUsers) getQueryID() uint32 {
	return data.QueryID
}
func (data *MSG_USER_getCompanyUsers) setQueryID(id uint32) {
	data.QueryID = id
}

type MSG_USER_getCompanyUsers_result struct {
	QueryResultID uint32
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
	data.QueryResultID = 0
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
	WRITE_uint32(data.QueryResultID, buf)
	WRITE_int32(int32(len(data.List)), buf)
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
	data.QueryResultID = READ_uint32(buf)
	List_len := int(READ_int32(buf))
	for i := 0; i < List_len; i++ {
		data.List = append(data.List, READ_MSG_USER_INFO_cache(buf))
	}
	data.Total = READ_int(buf)

}
func (data *MSG_USER_getCompanyUsers_result) getQueryResultID() uint32 {
	return data.QueryResultID
}
func (data *MSG_USER_getCompanyUsers_result) setQueryResultID(id uint32) {
	data.QueryResultID = id
}

