package protocol

import (
	"sync"
	"libraries"
	"unsafe"
	"time"
)

const (
	CMD_MSG_OA_attend_getByAccount = 1153965319
	CMD_MSG_OA_attend_getByAccount_result = 542857479
	CMD_MSG_OA_attend_getAllMonth = 110893319
	CMD_MSG_OA_attend_getAllMonth_result = -1533055481
	CMD_MSG_OA_attend_year_info = 948880647
	CMD_MSG_OA_attend_info = 1143057415
	CMD_MSG_OA_attend_computeStat = -1930839289
	CMD_MSG_OA_attend_computeStat_result = 880401415
	CMD_MSG_OA_attend_statInfo = 1199344903
	CMD_MSG_OA_attend_statAttendExt = -24900345
	CMD_MSG_OA_attend_statAbsentExt = -772908025
	CMD_MSG_OA_attend_statMarkExt = 240008455
)

type MSG_OA_attend_getByAccount struct {
	Uid int32
	StartDate time.Time
	EndDate time.Time
}

var pool_MSG_OA_attend_getByAccount = sync.Pool{New: func() interface{} { return &MSG_OA_attend_getByAccount{} }}

func GET_MSG_OA_attend_getByAccount() *MSG_OA_attend_getByAccount {
	return pool_MSG_OA_attend_getByAccount.Get().(*MSG_OA_attend_getByAccount)
}

func (data *MSG_OA_attend_getByAccount) cmd() int32 {
	return CMD_MSG_OA_attend_getByAccount
}

func (data *MSG_OA_attend_getByAccount) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_getByAccount()))
}

func (data *MSG_OA_attend_getByAccount) Put() {
	data.Uid = 0
	data.StartDate = time.UnixMicro(0)
	data.EndDate = time.UnixMicro(0)
	pool_MSG_OA_attend_getByAccount.Put(data)
}
func (data *MSG_OA_attend_getByAccount) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_getByAccount,buf)
	WRITE_MSG_OA_attend_getByAccount(data, buf)
}

func WRITE_MSG_OA_attend_getByAccount(data *MSG_OA_attend_getByAccount, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
	WRITE_int64(data.StartDate.UnixMicro(), buf)
	WRITE_int64(data.EndDate.UnixMicro(), buf)
}

func READ_MSG_OA_attend_getByAccount(buf *libraries.MsgBuffer) *MSG_OA_attend_getByAccount {
	data := pool_MSG_OA_attend_getByAccount.Get().(*MSG_OA_attend_getByAccount)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_getByAccount) read(buf *libraries.MsgBuffer) {
	data.Uid = READ_int32(buf)
	data.StartDate = time.UnixMicro(READ_int64(buf))
	data.EndDate = time.UnixMicro(READ_int64(buf))

}

type MSG_OA_attend_getByAccount_result struct {
	List []*MSG_OA_attend_info
}

var pool_MSG_OA_attend_getByAccount_result = sync.Pool{New: func() interface{} { return &MSG_OA_attend_getByAccount_result{} }}

func GET_MSG_OA_attend_getByAccount_result() *MSG_OA_attend_getByAccount_result {
	return pool_MSG_OA_attend_getByAccount_result.Get().(*MSG_OA_attend_getByAccount_result)
}

func (data *MSG_OA_attend_getByAccount_result) cmd() int32 {
	return CMD_MSG_OA_attend_getByAccount_result
}

func (data *MSG_OA_attend_getByAccount_result) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_getByAccount_result()))
}

func (data *MSG_OA_attend_getByAccount_result) Put() {
	data.List = data.List[:0]
	pool_MSG_OA_attend_getByAccount_result.Put(data)
}
func (data *MSG_OA_attend_getByAccount_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_getByAccount_result,buf)
	WRITE_MSG_OA_attend_getByAccount_result(data, buf)
}

func WRITE_MSG_OA_attend_getByAccount_result(data *MSG_OA_attend_getByAccount_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_OA_attend_info(v, buf)
	}
}

func READ_MSG_OA_attend_getByAccount_result(buf *libraries.MsgBuffer) *MSG_OA_attend_getByAccount_result {
	data := pool_MSG_OA_attend_getByAccount_result.Get().(*MSG_OA_attend_getByAccount_result)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_getByAccount_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_OA_attend_info, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_OA_attend_info(buf)
	}

}

type MSG_OA_attend_getAllMonth struct {
	Uids []int32
}

var pool_MSG_OA_attend_getAllMonth = sync.Pool{New: func() interface{} { return &MSG_OA_attend_getAllMonth{} }}

func GET_MSG_OA_attend_getAllMonth() *MSG_OA_attend_getAllMonth {
	return pool_MSG_OA_attend_getAllMonth.Get().(*MSG_OA_attend_getAllMonth)
}

func (data *MSG_OA_attend_getAllMonth) cmd() int32 {
	return CMD_MSG_OA_attend_getAllMonth
}

func (data *MSG_OA_attend_getAllMonth) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_getAllMonth()))
}

func (data *MSG_OA_attend_getAllMonth) Put() {
	data.Uids = data.Uids[:0]
	pool_MSG_OA_attend_getAllMonth.Put(data)
}
func (data *MSG_OA_attend_getAllMonth) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_getAllMonth,buf)
	WRITE_MSG_OA_attend_getAllMonth(data, buf)
}

func WRITE_MSG_OA_attend_getAllMonth(data *MSG_OA_attend_getAllMonth, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.Uids), buf)
	for _, v := range data.Uids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_OA_attend_getAllMonth(buf *libraries.MsgBuffer) *MSG_OA_attend_getAllMonth {
	data := pool_MSG_OA_attend_getAllMonth.Get().(*MSG_OA_attend_getAllMonth)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_getAllMonth) read(buf *libraries.MsgBuffer) {
	Uids_len := READ_int(buf)
	if Uids_len>cap(data.Uids){
		data.Uids= make([]int32, Uids_len)
	}else{
		data.Uids = data.Uids[:Uids_len]
	}
	for i := 0; i < Uids_len; i++ {
		data.Uids[i] = READ_int32(buf)
	}

}

type MSG_OA_attend_getAllMonth_result struct {
	List []*MSG_OA_attend_year_info
}

var pool_MSG_OA_attend_getAllMonth_result = sync.Pool{New: func() interface{} { return &MSG_OA_attend_getAllMonth_result{} }}

func GET_MSG_OA_attend_getAllMonth_result() *MSG_OA_attend_getAllMonth_result {
	return pool_MSG_OA_attend_getAllMonth_result.Get().(*MSG_OA_attend_getAllMonth_result)
}

func (data *MSG_OA_attend_getAllMonth_result) cmd() int32 {
	return CMD_MSG_OA_attend_getAllMonth_result
}

func (data *MSG_OA_attend_getAllMonth_result) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_getAllMonth_result()))
}

func (data *MSG_OA_attend_getAllMonth_result) Put() {
	data.List = data.List[:0]
	pool_MSG_OA_attend_getAllMonth_result.Put(data)
}
func (data *MSG_OA_attend_getAllMonth_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_getAllMonth_result,buf)
	WRITE_MSG_OA_attend_getAllMonth_result(data, buf)
}

func WRITE_MSG_OA_attend_getAllMonth_result(data *MSG_OA_attend_getAllMonth_result, buf *libraries.MsgBuffer) {
	WRITE_int(len(data.List), buf)
	for _, v := range data.List{
		WRITE_MSG_OA_attend_year_info(v, buf)
	}
}

func READ_MSG_OA_attend_getAllMonth_result(buf *libraries.MsgBuffer) *MSG_OA_attend_getAllMonth_result {
	data := pool_MSG_OA_attend_getAllMonth_result.Get().(*MSG_OA_attend_getAllMonth_result)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_getAllMonth_result) read(buf *libraries.MsgBuffer) {
	List_len := READ_int(buf)
	if List_len>cap(data.List){
		data.List= make([]*MSG_OA_attend_year_info, List_len)
	}else{
		data.List = data.List[:List_len]
	}
	for i := 0; i < List_len; i++ {
		data.List[i] = READ_MSG_OA_attend_year_info(buf)
	}

}

type MSG_OA_attend_year_info struct {
	Year string
	MonthList []string
}

var pool_MSG_OA_attend_year_info = sync.Pool{New: func() interface{} { return &MSG_OA_attend_year_info{} }}

func GET_MSG_OA_attend_year_info() *MSG_OA_attend_year_info {
	return pool_MSG_OA_attend_year_info.Get().(*MSG_OA_attend_year_info)
}

func (data *MSG_OA_attend_year_info) cmd() int32 {
	return CMD_MSG_OA_attend_year_info
}

func (data *MSG_OA_attend_year_info) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_year_info()))
}

func (data *MSG_OA_attend_year_info) Put() {
	data.Year = ``
	data.MonthList = data.MonthList[:0]
	pool_MSG_OA_attend_year_info.Put(data)
}
func (data *MSG_OA_attend_year_info) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_year_info,buf)
	WRITE_MSG_OA_attend_year_info(data, buf)
}

func WRITE_MSG_OA_attend_year_info(data *MSG_OA_attend_year_info, buf *libraries.MsgBuffer) {
	WRITE_string(data.Year, buf)
	WRITE_int(len(data.MonthList), buf)
	for _, v := range data.MonthList{
		WRITE_string(v, buf)
	}
}

func READ_MSG_OA_attend_year_info(buf *libraries.MsgBuffer) *MSG_OA_attend_year_info {
	data := pool_MSG_OA_attend_year_info.Get().(*MSG_OA_attend_year_info)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_year_info) read(buf *libraries.MsgBuffer) {
	data.Year = READ_string(buf)
	MonthList_len := READ_int(buf)
	if MonthList_len>cap(data.MonthList){
		data.MonthList= make([]string, MonthList_len)
	}else{
		data.MonthList = data.MonthList[:MonthList_len]
	}
	for i := 0; i < MonthList_len; i++ {
		data.MonthList[i] = READ_string(buf)
	}

}

type MSG_OA_attend_info struct {
	Id int32
	Uid int32
	Account string
	Date time.Time
	SignIn string
	SignOut string
	Status string
	Ip string
	Device string
	Client string
	ManualIn string
	ManualOut string
	Reason string
	Desc string
	ReviewStatus string
	ReviewedBy string
	ReviewedDate time.Time
	EarlyMin int32
	LateMin int32
	HoursList map[string]float32 `db:"-"`
}

var pool_MSG_OA_attend_info = sync.Pool{New: func() interface{} { return &MSG_OA_attend_info{} }}

func GET_MSG_OA_attend_info() *MSG_OA_attend_info {
	return pool_MSG_OA_attend_info.Get().(*MSG_OA_attend_info)
}

func (data *MSG_OA_attend_info) cmd() int32 {
	return CMD_MSG_OA_attend_info
}

func (data *MSG_OA_attend_info) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_info()))
}

func (data *MSG_OA_attend_info) Put() {
	data.Id = 0
	data.Uid = 0
	data.Account = ``
	data.Date = time.UnixMicro(0)
	data.SignIn = ``
	data.SignOut = ``
	data.Status = ``
	data.Ip = ``
	data.Device = ``
	data.Client = ``
	data.ManualIn = ``
	data.ManualOut = ``
	data.Reason = ``
	data.Desc = ``
	data.ReviewStatus = ``
	data.ReviewedBy = ``
	data.ReviewedDate = time.UnixMicro(0)
	data.EarlyMin = 0
	data.LateMin = 0
	data.HoursList = nil
	pool_MSG_OA_attend_info.Put(data)
}
func (data *MSG_OA_attend_info) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_info,buf)
	WRITE_MSG_OA_attend_info(data, buf)
}

func WRITE_MSG_OA_attend_info(data *MSG_OA_attend_info, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Uid, buf)
	WRITE_string(data.Account, buf)
	WRITE_int64(data.Date.UnixMicro(), buf)
	WRITE_string(data.SignIn, buf)
	WRITE_string(data.SignOut, buf)
	WRITE_string(data.Status, buf)
	WRITE_string(data.Ip, buf)
	WRITE_string(data.Device, buf)
	WRITE_string(data.Client, buf)
	WRITE_string(data.ManualIn, buf)
	WRITE_string(data.ManualOut, buf)
	WRITE_string(data.Reason, buf)
	WRITE_string(data.Desc, buf)
	WRITE_string(data.ReviewStatus, buf)
	WRITE_string(data.ReviewedBy, buf)
	WRITE_int64(data.ReviewedDate.UnixMicro(), buf)
	WRITE_int32(data.EarlyMin, buf)
	WRITE_int32(data.LateMin, buf)
	WRITE_map(data.HoursList,buf)
}

func READ_MSG_OA_attend_info(buf *libraries.MsgBuffer) *MSG_OA_attend_info {
	data := pool_MSG_OA_attend_info.Get().(*MSG_OA_attend_info)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_info) read(buf *libraries.MsgBuffer) {
	data.Id = READ_int32(buf)
	data.Uid = READ_int32(buf)
	data.Account = READ_string(buf)
	data.Date = time.UnixMicro(READ_int64(buf))
	data.SignIn = READ_string(buf)
	data.SignOut = READ_string(buf)
	data.Status = READ_string(buf)
	data.Ip = READ_string(buf)
	data.Device = READ_string(buf)
	data.Client = READ_string(buf)
	data.ManualIn = READ_string(buf)
	data.ManualOut = READ_string(buf)
	data.Reason = READ_string(buf)
	data.Desc = READ_string(buf)
	data.ReviewStatus = READ_string(buf)
	data.ReviewedBy = READ_string(buf)
	data.ReviewedDate = time.UnixMicro(READ_int64(buf))
	data.EarlyMin = READ_int32(buf)
	data.LateMin = READ_int32(buf)
	READ_map(&data.HoursList,buf)

}

type MSG_OA_attend_computeStat struct {
	Year string
	Month string
	Uids []int32
}

var pool_MSG_OA_attend_computeStat = sync.Pool{New: func() interface{} { return &MSG_OA_attend_computeStat{} }}

func GET_MSG_OA_attend_computeStat() *MSG_OA_attend_computeStat {
	return pool_MSG_OA_attend_computeStat.Get().(*MSG_OA_attend_computeStat)
}

func (data *MSG_OA_attend_computeStat) cmd() int32 {
	return CMD_MSG_OA_attend_computeStat
}

func (data *MSG_OA_attend_computeStat) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_computeStat()))
}

func (data *MSG_OA_attend_computeStat) Put() {
	data.Year = ``
	data.Month = ``
	data.Uids = data.Uids[:0]
	pool_MSG_OA_attend_computeStat.Put(data)
}
func (data *MSG_OA_attend_computeStat) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_computeStat,buf)
	WRITE_MSG_OA_attend_computeStat(data, buf)
}

func WRITE_MSG_OA_attend_computeStat(data *MSG_OA_attend_computeStat, buf *libraries.MsgBuffer) {
	WRITE_string(data.Year, buf)
	WRITE_string(data.Month, buf)
	WRITE_int(len(data.Uids), buf)
	for _, v := range data.Uids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_OA_attend_computeStat(buf *libraries.MsgBuffer) *MSG_OA_attend_computeStat {
	data := pool_MSG_OA_attend_computeStat.Get().(*MSG_OA_attend_computeStat)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_computeStat) read(buf *libraries.MsgBuffer) {
	data.Year = READ_string(buf)
	data.Month = READ_string(buf)
	Uids_len := READ_int(buf)
	if Uids_len>cap(data.Uids){
		data.Uids= make([]int32, Uids_len)
	}else{
		data.Uids = data.Uids[:Uids_len]
	}
	for i := 0; i < Uids_len; i++ {
		data.Uids[i] = READ_int32(buf)
	}

}

type MSG_OA_attend_computeStat_result struct {
	Stat map[int32]*MSG_OA_attend_statInfo
}

var pool_MSG_OA_attend_computeStat_result = sync.Pool{New: func() interface{} { return &MSG_OA_attend_computeStat_result{} }}

func GET_MSG_OA_attend_computeStat_result() *MSG_OA_attend_computeStat_result {
	return pool_MSG_OA_attend_computeStat_result.Get().(*MSG_OA_attend_computeStat_result)
}

func (data *MSG_OA_attend_computeStat_result) cmd() int32 {
	return CMD_MSG_OA_attend_computeStat_result
}

func (data *MSG_OA_attend_computeStat_result) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_computeStat_result()))
}

func (data *MSG_OA_attend_computeStat_result) Put() {
	data.Stat = nil
	pool_MSG_OA_attend_computeStat_result.Put(data)
}
func (data *MSG_OA_attend_computeStat_result) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_computeStat_result,buf)
	WRITE_MSG_OA_attend_computeStat_result(data, buf)
}

func WRITE_MSG_OA_attend_computeStat_result(data *MSG_OA_attend_computeStat_result, buf *libraries.MsgBuffer) {
	WRITE_map(data.Stat,buf)
}

func READ_MSG_OA_attend_computeStat_result(buf *libraries.MsgBuffer) *MSG_OA_attend_computeStat_result {
	data := pool_MSG_OA_attend_computeStat_result.Get().(*MSG_OA_attend_computeStat_result)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_computeStat_result) read(buf *libraries.MsgBuffer) {
	READ_map(&data.Stat,buf)

}

type MSG_OA_attend_statInfo struct {
	Deserve int
	Actual float32
	Normal float32
	Late int
	Early int
	Absent int
	Trip float32
	Egress float32
	Lieu float32
	PaidLeave float32
	UnpaidLeave float32
	TimeOvertime float32
	RestOvertime float32
	HolidayOvertime float32
	SickLeave float32
	AffairsLeave float32
	AbsentDates map[int64]*MSG_OA_attend_statAbsentExt
	AnnualLeave float32
	MarryLeave float32
	MaternityLeave float32
	Mark []*MSG_OA_attend_statMarkExt
	AttendExtDesc map[int64][]*MSG_OA_attend_statAttendExt
	Workmin int
	EarlyMin int
	LateMin int
	NotSignIn int
	NotSignOut int
	Attend map[int64]*MSG_OA_attend_info
}

var pool_MSG_OA_attend_statInfo = sync.Pool{New: func() interface{} { return &MSG_OA_attend_statInfo{} }}

func GET_MSG_OA_attend_statInfo() *MSG_OA_attend_statInfo {
	return pool_MSG_OA_attend_statInfo.Get().(*MSG_OA_attend_statInfo)
}

func (data *MSG_OA_attend_statInfo) cmd() int32 {
	return CMD_MSG_OA_attend_statInfo
}

func (data *MSG_OA_attend_statInfo) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_statInfo()))
}

func (data *MSG_OA_attend_statInfo) Put() {
	data.Deserve = 0
	data.Actual = 0
	data.Normal = 0
	data.Late = 0
	data.Early = 0
	data.Absent = 0
	data.Trip = 0
	data.Egress = 0
	data.Lieu = 0
	data.PaidLeave = 0
	data.UnpaidLeave = 0
	data.TimeOvertime = 0
	data.RestOvertime = 0
	data.HolidayOvertime = 0
	data.SickLeave = 0
	data.AffairsLeave = 0
	data.AbsentDates = nil
	data.AnnualLeave = 0
	data.MarryLeave = 0
	data.MaternityLeave = 0
	data.Mark = data.Mark[:0]
	data.AttendExtDesc = nil
	data.Workmin = 0
	data.EarlyMin = 0
	data.LateMin = 0
	data.NotSignIn = 0
	data.NotSignOut = 0
	data.Attend = nil
	pool_MSG_OA_attend_statInfo.Put(data)
}
func (data *MSG_OA_attend_statInfo) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_statInfo,buf)
	WRITE_MSG_OA_attend_statInfo(data, buf)
}

func WRITE_MSG_OA_attend_statInfo(data *MSG_OA_attend_statInfo, buf *libraries.MsgBuffer) {
	WRITE_int(data.Deserve, buf)
	WRITE_float32(data.Actual, buf)
	WRITE_float32(data.Normal, buf)
	WRITE_int(data.Late, buf)
	WRITE_int(data.Early, buf)
	WRITE_int(data.Absent, buf)
	WRITE_float32(data.Trip, buf)
	WRITE_float32(data.Egress, buf)
	WRITE_float32(data.Lieu, buf)
	WRITE_float32(data.PaidLeave, buf)
	WRITE_float32(data.UnpaidLeave, buf)
	WRITE_float32(data.TimeOvertime, buf)
	WRITE_float32(data.RestOvertime, buf)
	WRITE_float32(data.HolidayOvertime, buf)
	WRITE_float32(data.SickLeave, buf)
	WRITE_float32(data.AffairsLeave, buf)
	WRITE_map(data.AbsentDates,buf)
	WRITE_float32(data.AnnualLeave, buf)
	WRITE_float32(data.MarryLeave, buf)
	WRITE_float32(data.MaternityLeave, buf)
	WRITE_int(len(data.Mark), buf)
	for _, v := range data.Mark{
		WRITE_MSG_OA_attend_statMarkExt(v, buf)
	}
	WRITE_map(data.AttendExtDesc,buf)
	WRITE_int(data.Workmin, buf)
	WRITE_int(data.EarlyMin, buf)
	WRITE_int(data.LateMin, buf)
	WRITE_int(data.NotSignIn, buf)
	WRITE_int(data.NotSignOut, buf)
	WRITE_map(data.Attend,buf)
}

func READ_MSG_OA_attend_statInfo(buf *libraries.MsgBuffer) *MSG_OA_attend_statInfo {
	data := pool_MSG_OA_attend_statInfo.Get().(*MSG_OA_attend_statInfo)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_statInfo) read(buf *libraries.MsgBuffer) {
	data.Deserve = READ_int(buf)
	data.Actual = READ_float32(buf)
	data.Normal = READ_float32(buf)
	data.Late = READ_int(buf)
	data.Early = READ_int(buf)
	data.Absent = READ_int(buf)
	data.Trip = READ_float32(buf)
	data.Egress = READ_float32(buf)
	data.Lieu = READ_float32(buf)
	data.PaidLeave = READ_float32(buf)
	data.UnpaidLeave = READ_float32(buf)
	data.TimeOvertime = READ_float32(buf)
	data.RestOvertime = READ_float32(buf)
	data.HolidayOvertime = READ_float32(buf)
	data.SickLeave = READ_float32(buf)
	data.AffairsLeave = READ_float32(buf)
	READ_map(&data.AbsentDates,buf)
	data.AnnualLeave = READ_float32(buf)
	data.MarryLeave = READ_float32(buf)
	data.MaternityLeave = READ_float32(buf)
	Mark_len := READ_int(buf)
	if Mark_len>cap(data.Mark){
		data.Mark= make([]*MSG_OA_attend_statMarkExt, Mark_len)
	}else{
		data.Mark = data.Mark[:Mark_len]
	}
	for i := 0; i < Mark_len; i++ {
		data.Mark[i] = READ_MSG_OA_attend_statMarkExt(buf)
	}
	READ_map(&data.AttendExtDesc,buf)
	data.Workmin = READ_int(buf)
	data.EarlyMin = READ_int(buf)
	data.LateMin = READ_int(buf)
	data.NotSignIn = READ_int(buf)
	data.NotSignOut = READ_int(buf)
	READ_map(&data.Attend,buf)

}

type MSG_OA_attend_statAttendExt struct {
	Type string
	Day float32
}

var pool_MSG_OA_attend_statAttendExt = sync.Pool{New: func() interface{} { return &MSG_OA_attend_statAttendExt{} }}

func GET_MSG_OA_attend_statAttendExt() *MSG_OA_attend_statAttendExt {
	return pool_MSG_OA_attend_statAttendExt.Get().(*MSG_OA_attend_statAttendExt)
}

func (data *MSG_OA_attend_statAttendExt) cmd() int32 {
	return CMD_MSG_OA_attend_statAttendExt
}

func (data *MSG_OA_attend_statAttendExt) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_statAttendExt()))
}

func (data *MSG_OA_attend_statAttendExt) Put() {
	data.Type = ``
	data.Day = 0
	pool_MSG_OA_attend_statAttendExt.Put(data)
}
func (data *MSG_OA_attend_statAttendExt) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_statAttendExt,buf)
	WRITE_MSG_OA_attend_statAttendExt(data, buf)
}

func WRITE_MSG_OA_attend_statAttendExt(data *MSG_OA_attend_statAttendExt, buf *libraries.MsgBuffer) {
	WRITE_string(data.Type, buf)
	WRITE_float32(data.Day, buf)
}

func READ_MSG_OA_attend_statAttendExt(buf *libraries.MsgBuffer) *MSG_OA_attend_statAttendExt {
	data := pool_MSG_OA_attend_statAttendExt.Get().(*MSG_OA_attend_statAttendExt)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_statAttendExt) read(buf *libraries.MsgBuffer) {
	data.Type = READ_string(buf)
	data.Day = READ_float32(buf)

}

type MSG_OA_attend_statAbsentExt struct {
	AmAbsent bool
	PmAbsent bool
}

var pool_MSG_OA_attend_statAbsentExt = sync.Pool{New: func() interface{} { return &MSG_OA_attend_statAbsentExt{} }}

func GET_MSG_OA_attend_statAbsentExt() *MSG_OA_attend_statAbsentExt {
	return pool_MSG_OA_attend_statAbsentExt.Get().(*MSG_OA_attend_statAbsentExt)
}

func (data *MSG_OA_attend_statAbsentExt) cmd() int32 {
	return CMD_MSG_OA_attend_statAbsentExt
}

func (data *MSG_OA_attend_statAbsentExt) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_statAbsentExt()))
}

func (data *MSG_OA_attend_statAbsentExt) Put() {
	data.AmAbsent = false
	data.PmAbsent = false
	pool_MSG_OA_attend_statAbsentExt.Put(data)
}
func (data *MSG_OA_attend_statAbsentExt) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_statAbsentExt,buf)
	WRITE_MSG_OA_attend_statAbsentExt(data, buf)
}

func WRITE_MSG_OA_attend_statAbsentExt(data *MSG_OA_attend_statAbsentExt, buf *libraries.MsgBuffer) {
	WRITE_bool(data.AmAbsent, buf)
	WRITE_bool(data.PmAbsent, buf)
}

func READ_MSG_OA_attend_statAbsentExt(buf *libraries.MsgBuffer) *MSG_OA_attend_statAbsentExt {
	data := pool_MSG_OA_attend_statAbsentExt.Get().(*MSG_OA_attend_statAbsentExt)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_statAbsentExt) read(buf *libraries.MsgBuffer) {
	data.AmAbsent = READ_bool(buf)
	data.PmAbsent = READ_bool(buf)

}

type MSG_OA_attend_statMarkExt struct {
	Type string
	Start int8
	Finish int8
	Days int
}

var pool_MSG_OA_attend_statMarkExt = sync.Pool{New: func() interface{} { return &MSG_OA_attend_statMarkExt{} }}

func GET_MSG_OA_attend_statMarkExt() *MSG_OA_attend_statMarkExt {
	return pool_MSG_OA_attend_statMarkExt.Get().(*MSG_OA_attend_statMarkExt)
}

func (data *MSG_OA_attend_statMarkExt) cmd() int32 {
	return CMD_MSG_OA_attend_statMarkExt
}

func (data *MSG_OA_attend_statMarkExt) SetUintptr(in uintptr) {
	*(*uintptr)(unsafe.Pointer(in)) = uintptr(unsafe.Pointer(GET_MSG_OA_attend_statMarkExt()))
}

func (data *MSG_OA_attend_statMarkExt) Put() {
	data.Type = ``
	data.Start = 0
	data.Finish = 0
	data.Days = 0
	pool_MSG_OA_attend_statMarkExt.Put(data)
}
func (data *MSG_OA_attend_statMarkExt) write(buf *libraries.MsgBuffer) {
	WRITE_int32(CMD_MSG_OA_attend_statMarkExt,buf)
	WRITE_MSG_OA_attend_statMarkExt(data, buf)
}

func WRITE_MSG_OA_attend_statMarkExt(data *MSG_OA_attend_statMarkExt, buf *libraries.MsgBuffer) {
	WRITE_string(data.Type, buf)
	WRITE_int8(data.Start, buf)
	WRITE_int8(data.Finish, buf)
	WRITE_int(data.Days, buf)
}

func READ_MSG_OA_attend_statMarkExt(buf *libraries.MsgBuffer) *MSG_OA_attend_statMarkExt {
	data := pool_MSG_OA_attend_statMarkExt.Get().(*MSG_OA_attend_statMarkExt)
	data.read(buf)
	return data
}

func (data *MSG_OA_attend_statMarkExt) read(buf *libraries.MsgBuffer) {
	data.Type = READ_string(buf)
	data.Start = READ_int8(buf)
	data.Finish = READ_int8(buf)
	data.Days = READ_int(buf)

}

