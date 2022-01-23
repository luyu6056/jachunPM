package handler

import (
	"jachunPM_oa/db"
	"libraries"
	"protocol"
	"strconv"
	"time"
)

//老格式的表升级为新格式的表
func mysqlUpgrade() {
	out := protocol.GET_MSG_USER_getPairs()
	out.Params = "noletter,account"
	msg, err := HostConn.GetMsg()
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法初始化msg err:%v", err)
		return
	}
	var result *protocol.MSG_USER_getPairs_result
	if err := msg.SendMsgWaitResult(0, out, &result); err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取用户信息 err:%v", err)
		return
	}
	msg.DB.Table(db.TABLE_ATTEND).Delete()
	type old_attend struct {
		Id           int32     `db:"auto_increment;pk"`
		Account      string    `db:"type:varchar(30);index"`
		Date         time.Time `db:"index"`
		SignIn       string    `db:"type:time;not null"`
		SignOut      string    `db:"type:time;not null"`
		Status       string    `db:"type:varchar(30)"`
		Ip           string    `db:"type:varchar(15)"`
		Device       string    `db:"type:varchar(30)"`
		Client       string    `db:"type:varchar(20)"`
		ManualIn     string    `db:"type:time;not null"`
		ManualOut    string    `db:"type:time;not null"`
		Reason       string    `db:"type:varchar(30)"`
		Desc         string    `db:"type:text"`
		ReviewStatus string    `db:"type:varchar(30)"`
		ReviewedBy   string    `db:"type:varchar(30)"`
		ReviewedDate time.Time `db:"not null"`
		EarlyMin     int32     `db:"null"`
		LateMin      int32     `db:"null"`
	}
	HostConn.DB.Regsiter(&old_attend{})
	var list []*old_attend
	err = msg.DB.Table("zt_attend").Field("`Id`,`Account`,`Date`,`SignIn`,`SignOut`,`Status`,`Ip`,`Device`,`Client`,`ManualIn`,`ManualOut`,`Reason`,`Desc`,`ReviewStatus`,`ReviewedBy`,`ReviewedDate`,`EarlyMin`,`LateMin`").Select(&list)
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取attend表 err:%v", err)
	}
	var insert_attend []*db.Attend
	for _, v := range list {
		a := &db.Attend{
			Id:           v.Id,
			Account:      v.Account,
			Date:         v.Date,
			SignIn:       v.SignIn,
			SignOut:      v.SignOut,
			Status:       v.Status,
			Ip:           v.Ip,
			Device:       v.Device,
			Client:       v.Client,
			ManualIn:     v.ManualIn,
			ManualOut:    v.ManualOut,
			Reason:       v.Reason,
			Desc:         v.Desc,
			ReviewStatus: v.ReviewStatus,
			ReviewedBy:   v.ReviewedBy,
			ReviewedDate: v.ReviewedDate,
			EarlyMin:     v.EarlyMin,
			LateMin:      v.LateMin,
		}
		for _, kv := range result.List {
			if kv.Value == v.Account {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					a.Uid = int32(id)
				}
			}
		}
		insert_attend = append(insert_attend, a)
	}
	for i := 0; i < len(insert_attend); i += 1000 {
		en := i + 1000
		if en > len(insert_attend) {
			en = len(insert_attend)
		}

		_, err = HostConn.DB.Table(db.TABLE_ATTEND).ReplaceAll(insert_attend[i:en])
		libraries.DebugLog("插入ATTEND %d-%d 条，错误 %v", i, en, err)
	}
	msg.DB.Table(db.TABLE_HOLIDAY).Delete()
	var holidays []*db.Holiday
	err = msg.DB.Table("zt_holiday").Field("`Year`,`Date`,`Option`,userid as Uid,Name").Select(&holidays)
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取holiday表 err:%v", err)
	}
	for _, h := range holidays {
		h.Year = int32(h.Date.Year())
	}
	_, err = HostConn.DB.Table(db.TABLE_HOLIDAY).ReplaceAll(holidays)
	libraries.DebugLog("插入holiday %d 条，错误 %v", len(holidays), err)
	type old_Leave struct {
		Id            int32     `db:"auto_increment;pk"`
		Year          int16     `db:"index"`
		Begin         time.Time `db:"index"`
		End           time.Time `db:"index"`
		Start         int8      `db:"not null"`
		Finish        int8      `db:"not null"`
		Hours         float32   `db:"default(0)"`
		BackDate      time.Time `db:"not null"`
		Type          string    `db:"type:varchar(30)"`
		Desc          string    `db:"type:text"`
		Status        string    `db:"type:varchar(30)"`
		CreatedBy     string
		CreatedDate   time.Time `db:"not null"`
		ReviewedBy    string
		ReviewedDate  time.Time `db:"not null"`
		Level         int8      `db:"not null"`
		Reviewers     string    `db:"type:text"`
		BackReviewers string    `db:"type:text"`
		Noticeleader  string    `db:"type:varchar(255)"`
	}
	HostConn.DB.Regsiter(&old_Leave{})
	msg.DB.Table(db.TABLE_LEAVE).Delete()
	var leavelist []*old_Leave
	err = msg.DB.Table("zt_leave").Field("`Id`,`Year`,`Begin`,`End`,`Start`,`Finish`,`Hours`,`CreatedBy`,`ReviewedBy`,`BackDate`,`Type`,`Desc`,`Status`,`CreatedDate`,`ReviewedDate`,`Level`,`Reviewers`,`BackReviewers`,`Noticeleader`").Select(&leavelist)
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取zt_leave表 err:%v", err)
	}
	var insert_leave []*db.Leave
	for _, v := range leavelist {
		a := &db.Leave{
			Id:           v.Id,
			Year:         v.Year,
			Begin:        v.Begin,
			End:          v.End,
			Start:        v.Start,
			Finish:       v.Finish,
			Day:          v.Hours,
			BackDate:     v.BackDate,
			Type:         v.Type,
			Desc:         v.Desc,
			Status:       v.Status,
			CreatedDate:  v.CreatedDate,
			ReviewedDate: v.ReviewedDate,
		}
		for _, kv := range result.List {
			if kv.Value == v.CreatedBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					a.CreatedBy = int32(id)
				}
			}
			if kv.Value == v.ReviewedBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					a.ReviewedBy = int32(id)
				}
			}
		}
		insert_leave = append(insert_leave, a)
	}
	for i := 0; i < len(insert_leave); i += 1000 {
		en := i + 1000
		if en > len(insert_leave) {
			en = len(insert_leave)
		}

		_, err = HostConn.DB.Table(db.TABLE_LEAVE).ReplaceAll(insert_leave[i:en])
		libraries.DebugLog("插入LEAVE %d-%d 条，错误 %v", i, en, err)
	}

	type old_Overtime struct {
		Id     int32     `db:"auto_increment;pk"`
		Year   int16     `db:"index"`
		Begin  time.Time `db:"index"`
		End    time.Time `db:"index"`
		Start  string    `db:"type:time;not null"`
		Finish string    `db:"type:time;not null"`
		Hours  float32   `db:"default(0)"`
		//Leave        string    `db:"type:varchar(255)"`
		Type         string    `db:"type:varchar(30)"`
		Desc         string    `db:"type:text"`
		Status       string    `db:"type:varchar(30)"`
		RejectReason string    `db:"type:varchar(100)"`
		CreatedBy    string    `db:"index"`
		CreatedDate  time.Time `db:"not null"`
		ReviewedBy   string    `db:"type:varchar(30)"`
		ReviewedDate time.Time `db:"not null"`
	}
	msg.DB.Table(db.TABLE_OVERTIME).Delete()
	HostConn.DB.Regsiter(&old_Overtime{})
	var overtimelist []*old_Overtime
	err = msg.DB.Table("zt_overtime").Field("`Id`,`Year`,`Begin`,`End`,`Start`,`Finish`,`Hours`,`Type`,`Desc`,`Status`,`CreatedBy`,`CreatedDate`,`ReviewedBy`,`ReviewedDate`,`RejectReason`").Select(&overtimelist)
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取zt_overtime表 err:%v", err)
	}
	var insert_overtime []*db.Overtime
	for _, v := range overtimelist {
		a := &db.Overtime{
			Id:           v.Id,
			Year:         v.Year,
			Begin:        v.Begin,
			End:          v.End,
			Start:        v.Start,
			Finish:       v.Finish,
			Day:          v.Hours,
			RejectReason: v.RejectReason,
			Type:         v.Type,
			Desc:         v.Desc,
			Status:       v.Status,
			CreatedDate:  v.CreatedDate,
			ReviewedDate: v.ReviewedDate,
		}
		for _, kv := range result.List {
			if kv.Value == v.CreatedBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					a.CreatedBy = int32(id)
				}
			}
			if kv.Value == v.ReviewedBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					a.ReviewedBy = int32(id)
				}
			}
		}
		insert_overtime = append(insert_overtime, a)
	}
	for i := 0; i < len(insert_overtime); i += 1000 {
		en := i + 1000
		if en > len(insert_overtime) {
			en = len(insert_overtime)
		}

		_, err = HostConn.DB.Table(db.TABLE_OVERTIME).ReplaceAll(insert_overtime[i:en])
		libraries.DebugLog("插入OVERTIME %d-%d 条，错误 %v", i, en, err)
	}
	msg.DB.Table(db.TABLE_OVERTIMEBASE).Delete()
	var OvertimeBases []*db.OvertimeBase
	err = msg.DB.Table("zt_overtimebase").Field("`Account`,`OffsetDay`").Select(&OvertimeBases)
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取zt_overtimebase表 err:%v", err)
	}
	for _, o := range OvertimeBases {
		for _, kv := range result.List {
			if kv.Value == o.Account {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					o.Uid = int32(id)
				}
			}

		}
	}
	_, err = HostConn.DB.Table(db.TABLE_OVERTIMEBASE).ReplaceAll(OvertimeBases)
	libraries.DebugLog("插入OVERTIMEBASE %d 条，错误 %v", len(OvertimeBases), err)

	msg.DB.Table(db.TABLE_TRIP).Delete()
	type old_Trip struct {
		Id           int32     `db:"auto_increment;pk"`
		Type         string    `db:"default(0)"` // 0=trip,1=egress,
		Customers    string    `db:"type:varchar(20)"`
		Name         string    `db:"type:varchar(30)"`
		Desc         string    `db:"type:text"`
		Status       string    `db:"type:varchar(20)"`
		Year         int16     `db:"index"`
		Begin        time.Time `db:"index"`
		End          time.Time `db:"index"`
		Start        int8      `db:"not null"`
		Finish       int8      `db:"not null"`
		Hours        float32   `db:"default(0)"`
		From         string    `db:"type:varchar(50)"`
		To           string    `db:"type:varchar(50)"`
		CreatedBy    string    `db:"index"`
		CreatedDate  time.Time `db:"not null"`
		ReviewedBy   string    `db:"type:varchar(30)"`
		ReviewedDate time.Time `db:"null"`
	}
	HostConn.DB.Regsiter(&old_Trip{})
	var trips []*old_Trip
	err = msg.DB.Table("zt_trip").Field("`Id`,`Type`,`Customers`,`Name`,`Desc`,`Status`,`Year`,`Begin`,`End`,`Start`,`Finish`,`Hours`,`From`,`To`,`CreatedBy`,`CreatedDate`,`ReviewedBy`,`ReviewedDate`").Select(&trips)
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取zt_trip表 err:%v", err)
	}

	var insert_trip []*db.Trip
	for _, v := range trips {
		a := &db.Trip{
			Id:        v.Id,
			Customers: v.Customers,
			Name:      v.Name,
			Year:      v.Year,
			Begin:     v.Begin,
			End:       v.End,
			Start:     v.Start,
			Finish:    v.Finish,
			Day:       v.Hours,

			Type:         v.Type,
			Desc:         v.Desc,
			Status:       v.Status,
			CreatedDate:  v.CreatedDate,
			ReviewedDate: v.ReviewedDate,
			From:         v.From,
			To:           v.To,
		}
		for _, kv := range result.List {
			if kv.Value == v.CreatedBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					a.CreatedBy = int32(id)
				}
			}
			if kv.Value == v.ReviewedBy {
				if id, _ := strconv.Atoi(kv.Key); id > 0 {
					a.ReviewedBy = int32(id)
				}
			}
		}
		insert_trip = append(insert_trip, a)
	}
	_, err = HostConn.DB.Table(db.TABLE_TRIP).ReplaceAll(insert_trip)
	libraries.DebugLog("插入TRIP %d 条，错误 %v", len(insert_trip), err)
}
func init() {
	//time.AfterFunc(time.Second*5, mysqlUpgrade)
}
