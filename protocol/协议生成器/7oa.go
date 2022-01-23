package main

import "time"

type MSG_OA_attend_getByAccount struct {
	Uid       int32
	StartDate time.Time
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
	ReviewedBy   string
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
