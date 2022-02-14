package main

import "time"

type MSG_OA_attend_getByAccount struct {
	Uids      []int32
	BeginDate time.Time
	EndDate   time.Time
}
type MSG_OA_attend_getByAccount_result struct {
	List []*MSG_OA_attend_info
}
type MSG_OA_attend_getAllMonth struct {
	Uids []int32
}
type MSG_OA_attend_getAllMonth_result struct {
	List []*MSG_OA_attend_year_info
}
type MSG_OA_attend_year_info struct {
	Year      string   //2006格式
	MonthList []string //200601格式
}

type MSG_OA_attend_info struct {
	Id           int32
	Uid          int32
	Account      string
	Date         time.Time
	SignIn       string
	SignOut      string
	Status       string
	Ip           string
	Device       string
	Client       string
	ManualIn     string
	ManualOut    string
	Reason       string
	Desc         string
	ReviewStatus string
	ReviewedBy   int32
	RejectDesc   string
	ReviewedDate time.Time
	EarlyMin     int32
	LateMin      int32
	HoursList    map[string]float32 `db:"-"`
}

type MSG_OA_attend_computeStat struct {
	Year  string
	Month string
	Uids  []int32
}
type MSG_OA_attend_computeStat_result struct {
	Stat map[int32]*MSG_OA_attend_statInfo
}

type MSG_OA_attend_statInfo struct {
	Deserve         int
	Actual          float32
	Normal          float32
	Late            int
	Early           int
	Absent          int
	Trip            float32
	Egress          float32
	Lieu            float32
	PaidLeave       float32
	UnpaidLeave     float32
	TimeOvertime    float32
	RestOvertime    float32
	HolidayOvertime float32
	SickLeave       float32
	AffairsLeave    float32
	AbsentDates     map[int64]*MSG_OA_attend_statAbsentExt
	AnnualLeave     float32
	MarryLeave      float32
	MaternityLeave  float32
	Mark            []*MSG_OA_attend_statMarkExt
	AttendExtDesc   map[int64][]*MSG_OA_attend_statAttendExt
	Workmin         int
	EarlyMin        int
	LateMin         int
	NotSignIn       int
	NotSignOut      int
	Attend          map[int64]*MSG_OA_attend_info
}
type MSG_OA_attend_statAttendExt struct {
	Type string
	Day  float32
}
type MSG_OA_attend_statAbsentExt struct {
	AmAbsent bool
	PmAbsent bool
}
type MSG_OA_attend_statMarkExt struct {
	Type   string
	Start  int8
	Finish int8
	Days   int
}
type MSG_OA_attend_detail struct {
	BeginDate time.Time
	EndDate   time.Time
	User      []int32
}
type MSG_OA_attend_detail_result struct {
	List []*MSG_OA_attend_detail_info
}
type MSG_OA_attend_detail_info struct {
	Dept      string
	Realname  string
	Date      time.Time
	DayName   int
	Status    string
	Desc      string
	SignIn    string
	SignOut   string
	EarlyMin  int32
	LateMin   int32
	IP        string
	HoursList map[string]float32
}
type MSG_OA_attend_getWaitAttends struct {
	Users []int32
}

type MSG_OA_attend_getWaitAttends_result struct {
	List []*MSG_OA_attend_info
}

type MSG_OA_attend_getByDate struct {
	Uid  int32
	Date time.Time
}
type MSG_OA_attend_update struct {
	Uid          int32
	Date         time.Time
	ManualIn     time.Time
	ManualOut    time.Time
	Desc         string
	ReviewStatus string
	Reason       string
	ReviewedBy   int32
	RejectDesc   string
}

type MSG_OA_attend_getById struct {
	Id int32
}
type MSG_OA_attend_getbyId_result struct {
	Info *MSG_OA_attend_info
}
type MSG_OA_attend_getStat struct {
	Month time.Time
}
type MSG_OA_attend_getStat_result struct {
	List []map[string]string
}
type MSG_OA_attend_checkWaitReviews struct {
	Month time.Time
}
type MSG_OA_attend_checkWaitReviews_result struct {
	WaitReviews []string
}
