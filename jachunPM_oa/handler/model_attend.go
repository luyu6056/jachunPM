package handler

import (
	"errors"
	"fmt"
	"jachunPM_oa/db"
	"mysql"
	"protocol"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func attend_getByAccount(data *protocol.MSG_OA_attend_getByAccount, in *protocol.Msg) {
	config, err := attend_LoadConfig(in)
	if err != nil {
		in.WriteErr(err)
		return
	}
	where := map[string]interface{}{
		"Uid": data.Uid,
	}
	if v, ok := config["custom"]["beginDate"]; ok && v != "" {
		where["beginDateRaw"] = mysql.WhereOperatorRaw(fmt.Sprintf("`date` > '%s'", v))
	}
	if !data.StartDate.IsZero() {
		where["StartDateRaw"] = mysql.WhereOperatorRaw(fmt.Sprintf("`date` >= '%s'", data.StartDate.Format(protocol.TIMEFORMAT_MYSQLDATE)))
	}
	if !data.EndDate.IsZero() {
		where["EndDateRaw"] = mysql.WhereOperatorRaw(fmt.Sprintf("`date` <= '%s'", data.EndDate.Format(protocol.TIMEFORMAT_MYSQLDATE)))
	}
	out := protocol.GET_MSG_OA_attend_getByAccount_result()
	if err = in.DB.Table(db.TABLE_ATTEND).Where(where).Select(&out.List); err != nil {
		in.WriteErr(err)
		return
	}
	if out.List, err = attend_fixUserAttendList(config, out.List, data.StartDate, data.EndDate, data.Uid); err != nil {
		in.WriteErr(err)
		return
	}
	if err = attend_processAttendList(config, out.List); err != nil {
		in.WriteErr(err)
		return
	}

	if err = attend_processHours(config, out.List, data.StartDate); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}

//进来的list必须按date从小到大排列
func attend_fixUserAttendList(config map[string]map[string]string, list []*protocol.MSG_OA_attend_info, startDate, endDate time.Time, uid int32) (res []*protocol.MSG_OA_attend_info, err error) {
	if !startDate.IsZero() && !endDate.IsZero() {
		for i := startDate; i.Unix() <= endDate.Unix(); i = i.AddDate(0, 0, 1) {
			find := false
			for _, attend := range list {
				if attend.Date == i {
					find = true
					break
				}
			}
			if !find {
				attend := protocol.GET_MSG_OA_attend_info()
				attend.Uid = uid
				attend.Date = i
				attend.SignIn = "00:00:00"
				attend.SignOut = "00:00:00"
				attend.ManualIn = "00:00:00"
				attend.ManualOut = "00:00:00"
				if attend.Status, err = attend_computeStatus(attend, config); err != nil {
					return
				}
				list = append(list, attend)
			}
		}
	}
	protocol.Order_attend(list, nil)
	return list, nil
}

func attend_computeStatus(attend *protocol.MSG_OA_attend_info, config map[string]map[string]string) (string, error) {
	if v, ok := config["custom"]["beginDate"]; ok && v != "" {
		if attend.Date.Format(protocol.TIMEFORMAT_MYSQLDATE) < v {
			return "normal", nil
		}
	}
	status := "normal"
	if attend.SignIn == "00:00:00" && attend.SignOut == "00:00:00" {
		status = "absent"
	} else {
		if attend.LateMin > 0 {
			status = "late"
		}
		if attend.EarlyMin > 0 {
			if status == "late" {
				status = "both"
			} else {
				status = "early"
			}

		}
		if config["custom"]["halfAbsendMin"] > "0" && strconv.Itoa(int(attend.LateMin+attend.EarlyMin)) >= config["custom"]["halfAbsendMin"] {
			//旷工半天
			status = "halfAbsent"
		}
		if config["custom"]["absendMin"] > "0" && strconv.Itoa(int(attend.LateMin+attend.EarlyMin)) >= config["custom"]["absendMin"] {
			//旷工一天
			status = "allAbsent"
		}
		if attend.Date.Format(protocol.TIMEFORMAT_MYSQLDATE) < time.Now().Format(protocol.TIMEFORMAT_MYSQLDATE) {
			if (attend.SignIn == "00:00:00" && attend.SignOut != "00:00:00") || (attend.SignIn != "00:00:00" && attend.SignOut == "00:00:00") {
				status = "halfAbsent"
			}
		}
	}
	isRestDay, err := attend_isRestDay(attend.Date, config)
	if isRestDay {
		status = "rest"
	}
	return status, err
}
func attend_LoadConfig(in *protocol.Msg) (map[string]map[string]string, error) {
	out := protocol.GET_MSG_USER_config_get()
	out.Uid = protocol.SYSTEMUSER
	out.Module = "attend"
	var result *protocol.MSG_USER_config_get_result
	err := in.SendMsgWaitResult(0, out, &result)
	out.Put()
	return result.Config, err
}
func attend_isRestDay(date time.Time, config map[string]map[string]string) (bool, error) {
	if holiday, err := getHoliday(date); err != nil {
		return false, err
	} else if holiday != nil {
		switch holiday.Option {
		case "OT": //加班
			return false, nil
		case "RE": //休息
			return true, nil
		}
	}
	week := strconv.Itoa(int(date.Weekday()+7) % 7)

	for _, work := range strings.Split(config["custom"]["workingDays"], ",") {
		if week == work {
			return false, nil
		}
	}
	return true, nil
}
func attend_processAttendList(config map[string]map[string]string, attends []*protocol.MSG_OA_attend_info) (err error) {
	for _, attend := range attends {
		if err = attend_processAttend(config, attend); err != nil {
			return
		}
	}
	return nil
}
func attend_processAttend(config map[string]map[string]string, attend *protocol.MSG_OA_attend_info) (err error) {
	if attend.Status == "" || attend.Status == "rest" {
		if attend.Status, err = attend_computeStatus(attend, config); err != nil {
			return
		}
	}
	if attend.SignIn == "00:00:00" {
		attend.SignIn = ""
	}
	if attend.SignOut == "00:00:00" {
		attend.SignOut = ""
	}
	if attend.ManualIn == "00:00:00" {
		attend.ManualIn = ""
	}
	if attend.ManualOut == "00:00:00" {
		attend.ManualOut = ""
	}
	return nil
}
func attend_processHours(config map[string]map[string]string, attends []*protocol.MSG_OA_attend_info, startDate time.Time) (err error) {
	var uids = make([]int32, len(attends))
	for k, attend := range attends {
		uids[k] = attend.Uid
	}
	year := startDate.Format("2006")
	date := startDate.Format("2006-01-%")
	var leaveList []*db.Leave
	if err = HostConn.DB.Table(db.TABLE_LEAVE).Where(map[string]interface{}{
		"CreatedBy": uids,
		"Year":      year,
		"date":      mysql.WhereOperatorRaw(fmt.Sprintf("(Begin like '%s' or End like '%s')", date, date)),
		"Status":    "pass",
	}).Order("id desc").Select(&leaveList); err != nil {
		return err
	}
	/*var lieuList []*db.Lieu
	if err = HostConn.DB.Table(db.TABLE_LIEU).Where(map[string]interface{}{
		"CreatedBy": uids,
		"Year":      year,
		"date":      mysql.WhereOperatorRaw(fmt.Sprintf("(Begin like '%s' or End like '%s')", date, date)),
		"Status":    "pass",
	}).Order("id desc").Select(&lieuList); err != nil {
		return err
	}*/
	var overtimeList []*db.Overtime
	if err = HostConn.DB.Table(db.TABLE_OVERTIME).Where(map[string]interface{}{
		"CreatedBy": uids,
		"Year":      year,
		"date":      mysql.WhereOperatorRaw(fmt.Sprintf("(Begin like '%s' or End like '%s')", date, date)),
		"Status":    "pass",
	}).Order("id desc").Select(&overtimeList); err != nil {
		return err
	}
	var tripList []*db.Trip
	if err = HostConn.DB.Table(db.TABLE_TRIP).Where(map[string]interface{}{
		"CreatedBy": uids,
		"Year":      year,
		"date":      mysql.WhereOperatorRaw(fmt.Sprintf("(Begin like '%s' or End like '%s')", date, date)),
		"Status":    "pass",
	}).Order("id desc").Select(&tripList); err != nil {
		return err
	}
	leaveM := make(map[int32][]*db.Leave)
	for _, leave := range leaveList {
		leaveM[leave.CreatedBy] = append(leaveM[leave.CreatedBy], leave)
	}
	/*lieuM := make(map[int32][]*db.Lieu)
	for _, lieu := range lieuList {
		lieuM[lieu.CreatedBy] = append(lieuM[lieu.CreatedBy], lieu)
	}*/
	overtimeM := make(map[int32][]*db.Overtime)

	for _, overtime := range overtimeList {
		overtimeM[overtime.CreatedBy] = append(overtimeM[overtime.CreatedBy], overtime)
	}
	tripM := make(map[int32][]*db.Trip)
	for _, trip := range tripList {
		tripM[trip.CreatedBy] = append(tripM[trip.CreatedBy], trip)
	}
	for _, attend := range attends {
		attend.HoursList = make(map[string]float32)
		if v, ok := leaveM[attend.Uid]; ok {
			if hour, err := attend_computeHours(config, attend, v, "leave"); err != nil {
				return err
			} else if hour != 0 {
				attend.HoursList["leave"] = hour
			}
		}
		/*if v, ok := lieuM[attend.Uid]; ok {
			if hour, err := attend_computeHours(config, attend, v, "lieu"); err != nil {
				return err
			} else if hour != 0 {
				attend.HoursList["lieu"] = hour
			}
		}*/
		if v, ok := overtimeM[attend.Uid]; ok {
			if hour, err := attend_computeHours(config, attend, v, "overtime"); err != nil {
				return err
			} else if hour != 0 {
				attend.HoursList["overtime"] = hour
			}
		}
		if v, ok := tripM[attend.Uid]; ok {
			if hour, err := attend_computeHours(config, attend, v, "trip"); err != nil {
				return err
			} else if hour != 0 {
				attend.HoursList["trip"] = hour
			}
		}

	}
	return nil
}

func attend_computeHours(config map[string]map[string]string, attend *protocol.MSG_OA_attend_info, dates interface{}, typ string) (float32, error) {

	d := reflect.ValueOf(dates)

	if typ == "overtime" {
		for i := 0; i < d.Len(); i++ {
			data := d.Index(i).Elem()
			begin := data.FieldByName("Begin").Interface().(time.Time)
			end := data.FieldByName("End").Interface().(time.Time)
			day := data.FieldByName("Day").Interface().(float32)
			if begin.After(attend.Date) || attend.Date.After(end) {
				continue
			}
			return day, nil
		}
	} else {
		isRestDay, err := attend_isRestDay(attend.Date, config)
		if isRestDay || err != nil {
			return 0, err
		}
		for i := 0; i < d.Len(); i++ {
			data := d.Index(i).Elem()
			begin := data.FieldByName("Begin").Interface().(time.Time)
			end := data.FieldByName("End").Interface().(time.Time)
			start := data.FieldByName("Start").Int()
			finish := data.FieldByName("Finish").Int()
			if begin.After(attend.Date) || attend.Date.After(end) {
				continue
			}

			if begin == end {
				if start == 1 && finish == 2 {
					return 1, nil
				}
				if start == 1 && finish == 1 {
					return protocol.AttendAM, nil
				}
				if start == 2 && finish == 2 {
					return protocol.AttendPM, nil
				}
			} else {
				if begin == attend.Date {
					if start == 1 {
						return 1, nil
					}
					if start == 2 {
						return protocol.AttendPM, nil
					}
				} else if end == attend.Date {
					if finish == 1 {
						return protocol.AttendAM, nil
					}
					if finish == 2 {
						return 1, nil
					}
				} else {
					return 1, nil
				}
			}
		}
	}

	return 0, nil
}
func attend_getAllMonth(data *protocol.MSG_OA_attend_getAllMonth, in *protocol.Msg) {
	var res []map[string]string
	var err error
	if len(data.Uids) > 0 {
		var ids = make([]string, len(data.Uids))
		for k, id := range data.Uids {
			ids[k] = strconv.Itoa(int(id))
		}
		if res, err = in.DB.Raw(fmt.Sprintf("select * from (select left(date,7) as date,left(date,4) as year from `%s` where Uid in (%s)) as x group by x.date order by x.date desc", db.TABLE_ATTEND, strings.Join(ids, ","))).SelectMap(); err != nil {
			in.WriteErr(err)
			return
		}
	} else {
		if res, err = in.DB.Raw(fmt.Sprintf("select * from (select left(date,7) as date,left(date,4) as year from `%s`) as x group by x.date order by x.date desc", db.TABLE_ATTEND)).SelectMap(); err != nil {
			in.WriteErr(err)
			return
		}
	}

	out := protocol.GET_MSG_OA_attend_getAllMonth_result()
	for _, row := range res {
		find := false
		for _, y := range out.List {
			if y.Year == row["year"] {
				y.MonthList = append(y.MonthList, row["date"])
				find = true
			}
		}
		if !find {
			out.List = append(out.List, &protocol.MSG_OA_attend_year_info{
				Year:      row["year"],
				MonthList: []string{row["date"]},
			})
		}
	}
	in.SendResult(out)
}

func attend_computeStat(data *protocol.MSG_OA_attend_computeStat, in *protocol.Msg) {
	config, err := attend_LoadConfig(in)
	if err != nil {
		in.WriteErr(err)
		return
	}
	var startDate, endDate time.Time
	if data.Month != "" {
		if startDate, err = time.ParseInLocation("2006-01-02", fmt.Sprintf("%s-%s-01", data.Year, data.Month), time.Local); err != nil {
			in.WriteErr(err)
			return
		}
		endDate = startDate.AddDate(0, 1, -1)
	} else {
		if startDate, err = time.ParseInLocation("2006-01-02", fmt.Sprintf("%s-01-01", data.Year), time.Local); err != nil {
			in.WriteErr(err)
			return
		}
		if endDate, err = time.ParseInLocation("2006-01-02", fmt.Sprintf("%s-12-31", data.Year), time.Local); err != nil {
			in.WriteErr(err)
			return
		}
	}
	//computeWorkingDates
	beginDate, _ := time.ParseInLocation("2006-01-02", config["custom"]["beginDate"], time.Local)
	var workingDays = make(map[time.Time]bool)
	for i := startDate; i.Unix() <= endDate.Unix(); i = i.AddDate(0, 0, 1) {
		if beginDate.After(i) {
			continue
		}
		if ok, err := attend_isRestDay(i, config); err != nil {
			in.WriteErr(err)
			return
		} else if !ok {
			workingDays[i] = true
		}
	}
	var noAttendUsers []string
	if config["custom"]["noAttendUsers"] != "" {
		noAttendUsers = strings.Split(config["custom"]["noAttendUsers"], ",")
	}
	out := protocol.GET_MSG_OA_attend_computeStat_result()
	if out.Stat == nil {
		out.Stat = make(map[int32]*protocol.MSG_OA_attend_statInfo)
	}
outcontinue:
	for _, uid := range data.Uids {
		for _, stdId := range noAttendUsers {
			if strconv.Itoa(int(uid)) == stdId {
				continue outcontinue
			}
		}
		stat := protocol.GET_MSG_OA_attend_statInfo()
		if stat.AbsentDates == nil {
			stat.AbsentDates = make(map[int64]*protocol.MSG_OA_attend_statAbsentExt)
		}
		if stat.AttendExtDesc == nil {
			stat.AttendExtDesc = make(map[int64][]*protocol.MSG_OA_attend_statAttendExt)
		}
		if stat.Attend == nil {
			stat.Attend = make(map[int64]*protocol.MSG_OA_attend_info)
		}
		for k := range workingDays {
			stat.AbsentDates[k.Unix()] = protocol.GET_MSG_OA_attend_statAbsentExt()
		}
		out.Stat[uid] = stat
	}
	if out.Stat, err = attend_computeTripStat(out.Stat, data.Uids, startDate, endDate, config); err != nil {
		in.WriteErr(err)
		return
	}
	if out.Stat, err = attend_computeLeaveStat(out.Stat, data.Uids, startDate, endDate, config); err != nil {
		in.WriteErr(err)
		return
	}
	//out.Stat = attend_computeLieuStat(out.Stat, data.Uids, startDate, endDate, config)
	//out.Stat = attend_computeMakeupStat(out.Stat, data.Uids, startDate, endDate, config)
	if out.Stat, err = attend_computeOvertimeStat(out.Stat, data.Uids, startDate, endDate, config); err != nil {
		in.WriteErr(err)
		return
	}
	if out.Stat, err = attend_computeAttendStat(out.Stat, data.Uids, startDate, endDate, config); err != nil {
		in.WriteErr(err)
		return
	}
	in.SendResult(out)
	out.Put()
}
func attend_computeTripStat(stat map[int32]*protocol.MSG_OA_attend_statInfo, uids []int32, startDate, endDate time.Time, config map[string]map[string]string) (res map[int32]*protocol.MSG_OA_attend_statInfo, err error) {
	var tripList []*db.Trip
	if err = HostConn.DB.Table(db.TABLE_TRIP).Where(map[string]interface{}{
		"CreatedBy": uids,
		"date":      mysql.WhereOperatorRaw(fmt.Sprintf("(End >= '%s' and  Begin <= '%s')", startDate.Format(protocol.TIMEFORMAT_MYSQLDATE), endDate.Format(protocol.TIMEFORMAT_MYSQLDATE))),
		"Status":    "pass",
	}).Order("Begin,End").Select(&tripList); err != nil {
		return
	}
	for _, trip := range tripList {
		trip.Day = 0 //重算时间
		if startDate.After(trip.Begin) {
			if trip.Finish == 1 {
				trip.Day = 0.5
			} else {
				trip.Day = 1
			}
			for i := startDate; trip.End.After(i); i = i.AddDate(0, 0, 1) {
				if ok, err := attend_isRestDay(i, config); err != nil {
					return nil, err
				} else if !ok {
					trip.Day++
				}
			}
		}
		if trip.End.After(endDate) {
			if trip.Start == 1 {
				trip.Day += 1
			} else {
				trip.Day += 0.5
			}
			for i := trip.Begin.AddDate(0, 0, 1); endDate.Unix() >= i.Unix(); i = i.AddDate(0, 0, 1) {
				if ok, err := attend_isRestDay(i, config); err != nil {
					return nil, err
				} else if !ok {
					trip.Day++
				}
			}
		}
		switch trip.Type {
		case "trip":
			stat[trip.CreatedBy].Trip += trip.Day
		case "egress":
			stat[trip.CreatedBy].Egress += trip.Day
		default:
			return nil, errors.New("attend_computeTripStat未处理1type:" + trip.Type)
		}

		for i := trip.Begin; i.Unix() <= trip.End.Unix(); i = i.AddDate(0, 0, 1) {
			if startDate.After(i) || i.After(endDate) {
				continue
			}
			key := i.Unix()
			switch {
			case trip.Begin == trip.End:
				switch {
				case trip.Start == 1 && trip.Finish == 1:
					stat[trip.CreatedBy].AttendExtDesc[key] = append(stat[trip.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: trip.Type,
						Day:  protocol.AttendAM,
					})
					stat[trip.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true}
				case trip.Start == 2 && trip.Finish == 2:
					stat[trip.CreatedBy].AttendExtDesc[key] = append(stat[trip.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: trip.Type,
						Day:  protocol.AttendPM,
					})
					stat[trip.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{PmAbsent: true}
				default:
					stat[trip.CreatedBy].AttendExtDesc[key] = append(stat[trip.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: trip.Type,
						Day:  1,
					})
					stat[trip.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true, PmAbsent: true}
				}
			case i == trip.Begin:
				if trip.Start == 1 {
					stat[trip.CreatedBy].AttendExtDesc[key] = append(stat[trip.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: trip.Type,
						Day:  1,
					})
					stat[trip.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true, PmAbsent: true}
				} else {
					stat[trip.CreatedBy].AttendExtDesc[key] = append(stat[trip.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: trip.Type,
						Day:  protocol.AttendPM,
					})
					stat[trip.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{PmAbsent: true}
				}
			case i == trip.End:
				if trip.Finish == 1 {
					stat[trip.CreatedBy].AttendExtDesc[key] = append(stat[trip.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: trip.Type,
						Day:  protocol.AttendAM,
					})
					stat[trip.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true}
				} else {
					stat[trip.CreatedBy].AttendExtDesc[key] = append(stat[trip.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: trip.Type,
						Day:  1,
					})
					stat[trip.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true, PmAbsent: true}
				}
			default:
				stat[trip.CreatedBy].AttendExtDesc[key] = append(stat[trip.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
					Type: trip.Type,
					Day:  1,
				})
				stat[trip.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true, PmAbsent: true}
			}
		}
		//添加导出备注
		if startDate.After(trip.Begin) {
			trip.Begin = startDate
			trip.Start = 1
		}
		if trip.End.After(endDate) {
			trip.End = endDate
			trip.Finish = 2
		}

	}
	return stat, nil
}

func attend_computeLeaveStat(stat map[int32]*protocol.MSG_OA_attend_statInfo, uids []int32, startDate, endDate time.Time, config map[string]map[string]string) (res map[int32]*protocol.MSG_OA_attend_statInfo, err error) {
	var leaveList []*db.Trip
	if err = HostConn.DB.Table(db.TABLE_LEAVE).Where(map[string]interface{}{
		"CreatedBy": uids,
		"date":      mysql.WhereOperatorRaw(fmt.Sprintf("(End >= '%s' and  Begin <= '%s')", startDate.Format(protocol.TIMEFORMAT_MYSQLDATE), endDate.Format(protocol.TIMEFORMAT_MYSQLDATE))),
		"Status":    "pass",
	}).Order("Begin,End").Select(&leaveList); err != nil {
		return
	}
	for _, leave := range leaveList {
		leave.Day = 0 //重算时间
		if startDate.After(leave.Begin) {
			if leave.Finish == 1 {
				leave.Day = 0.5
			} else {
				leave.Day = 1
			}
			for i := startDate; leave.End.After(i); i = i.AddDate(0, 0, 1) {
				if ok, err := attend_isRestDay(i, config); err != nil {
					return nil, err
				} else if !ok {
					leave.Day++
				}
			}
		}
		if leave.End.After(endDate) {
			if leave.Start == 1 {
				leave.Day += 1
			} else {
				leave.Day += 0.5
			}
			for i := leave.Begin.AddDate(0, 0, 1); endDate.Unix() >= i.Unix(); i = i.AddDate(0, 0, 1) {
				if ok, err := attend_isRestDay(i, config); err != nil {
					return nil, err
				} else if !ok {
					leave.Day++
				}
			}
		}
		switch leave.Type {
		case "sick":
			stat[leave.CreatedBy].SickLeave += leave.Day
			stat[leave.CreatedBy].UnpaidLeave += leave.Day
		case "affairs":
			stat[leave.CreatedBy].AffairsLeave += leave.Day
			stat[leave.CreatedBy].UnpaidLeave += leave.Day
		case "annual":
			stat[leave.CreatedBy].AnnualLeave += leave.Day
			stat[leave.CreatedBy].PaidLeave += leave.Day
		case "marry":
			stat[leave.CreatedBy].MarryLeave += leave.Day
			stat[leave.CreatedBy].PaidLeave += leave.Day
		case "maternity":
			stat[leave.CreatedBy].MaternityLeave += leave.Day
			stat[leave.CreatedBy].PaidLeave += leave.Day
		case "lieu":
			stat[leave.CreatedBy].Lieu += leave.Day
			stat[leave.CreatedBy].PaidLeave += leave.Day
		default:
			return nil, errors.New("attend_computeLeaveStat未处理1type:" + leave.Type)
		}

		for i := leave.Begin; i.Unix() <= leave.End.Unix(); i = i.AddDate(0, 0, 1) {
			if startDate.After(i) || i.After(endDate) {
				continue
			}
			key := i.Unix()
			switch {
			case leave.Begin == leave.End:
				switch {
				case leave.Start == 1 && leave.Finish == 1:
					stat[leave.CreatedBy].AttendExtDesc[key] = append(stat[leave.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: leave.Type,
						Day:  protocol.AttendAM,
					})
					stat[leave.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true}
				case leave.Start == 2 && leave.Finish == 2:
					stat[leave.CreatedBy].AttendExtDesc[key] = append(stat[leave.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: leave.Type,
						Day:  protocol.AttendPM,
					})
					stat[leave.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{PmAbsent: true}
				default:
					stat[leave.CreatedBy].AttendExtDesc[key] = append(stat[leave.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: leave.Type,
						Day:  1,
					})
					stat[leave.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true, PmAbsent: true}
				}
			case i == leave.Begin:
				if leave.Start == 1 {
					stat[leave.CreatedBy].AttendExtDesc[key] = append(stat[leave.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: leave.Type,
						Day:  1,
					})
					stat[leave.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true, PmAbsent: true}
				} else {
					stat[leave.CreatedBy].AttendExtDesc[key] = append(stat[leave.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: leave.Type,
						Day:  protocol.AttendPM,
					})
					stat[leave.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{PmAbsent: true}
				}
			case i == leave.End:
				if leave.Finish == 1 {
					stat[leave.CreatedBy].AttendExtDesc[key] = append(stat[leave.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: leave.Type,
						Day:  protocol.AttendAM,
					})
					stat[leave.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true}
				} else {
					stat[leave.CreatedBy].AttendExtDesc[key] = append(stat[leave.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
						Type: leave.Type,
						Day:  1,
					})
					stat[leave.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true, PmAbsent: true}
				}
			default:
				stat[leave.CreatedBy].AttendExtDesc[key] = append(stat[leave.CreatedBy].AttendExtDesc[key], &protocol.MSG_OA_attend_statAttendExt{
					Type: leave.Type,
					Day:  1,
				})
				stat[leave.CreatedBy].AbsentDates[key] = &protocol.MSG_OA_attend_statAbsentExt{AmAbsent: true, PmAbsent: true}
			}
		}
		//添加导出备注
		if startDate.After(leave.Begin) {
			leave.Begin = startDate
			leave.Start = 1
		}
		if leave.End.After(endDate) {
			leave.End = endDate
			leave.Finish = 2
		}

	}
	return stat, nil
}

func attend_computeOvertimeStat(stat map[int32]*protocol.MSG_OA_attend_statInfo, uids []int32, startDate, endDate time.Time, config map[string]map[string]string) (res map[int32]*protocol.MSG_OA_attend_statInfo, err error) {
	var overtimeList []*db.Trip
	if err = HostConn.DB.Table(db.TABLE_OVERTIME).Where(map[string]interface{}{
		"CreatedBy": uids,
		"date":      mysql.WhereOperatorRaw(fmt.Sprintf("(End >= '%s' and  Begin <= '%s')", startDate.Format(protocol.TIMEFORMAT_MYSQLDATE), endDate.Format(protocol.TIMEFORMAT_MYSQLDATE))),
		"Status":    "pass",
	}).Order("Begin,End").Select(&overtimeList); err != nil {
		return
	}
	for _, overtime := range overtimeList {
		switch overtime.Type {
		case "time":
			stat[overtime.CreatedBy].TimeOvertime += overtime.Day
		case "rest":
			stat[overtime.CreatedBy].RestOvertime += overtime.Day
		case "holiday":
			stat[overtime.CreatedBy].HolidayOvertime += overtime.Day
		case "compensate":
			stat[overtime.CreatedBy].Normal += overtime.Day
		}
	}
	return stat, nil
}

func attend_computeAttendStat(stat map[int32]*protocol.MSG_OA_attend_statInfo, uids []int32, startDate, endDate time.Time, config map[string]map[string]string) (res map[int32]*protocol.MSG_OA_attend_statInfo, err error) {

	var attendList []*protocol.MSG_OA_attend_info
	if err = HostConn.DB.Table(db.TABLE_ATTEND).Where(map[string]interface{}{
		"Uid":  uids,
		"date": mysql.WhereOperatorRaw(fmt.Sprintf("(Date >= '%s' and  Date <= '%s')", startDate.Format(protocol.TIMEFORMAT_MYSQLDATE), endDate.Format(protocol.TIMEFORMAT_MYSQLDATE))),
	}).Order("Date").Select(&attendList); err != nil {
		return
	}
	for _, attend := range attendList {
		if attend.Status == "normal" {
			stat[attend.Uid].Normal++
		}
		if attend.Status == "late" || attend.Status == "both" {
			stat[attend.Uid].Late++
		}
		if attend.Status == "early" || attend.Status == "both" {
			stat[attend.Uid].Early++
		}

		if attend.SignIn != "00:00:00" && attend.SignOut != "00:00:00" {
			in := strings.Split(attend.SignIn, ":")
			out := strings.Split(attend.SignOut, ":")
			inHour, _ := strconv.Atoi(in[0])
			inMin, _ := strconv.Atoi(in[1])
			outHour, _ := strconv.Atoi(out[0])
			outMin, _ := strconv.Atoi(out[1])
			workmin := outHour*60 + outMin - inHour*60 - inMin
			if workmin > 0 {
				stat[attend.Uid].Workmin += workmin
			}
		}
		if attend.SignIn == "00:00:00" && attend.SignOut != "00:00:00" {
			stat[attend.Uid].NotSignIn++
		}
		if attend.SignIn != "00:00:00" && attend.SignOut == "00:00:00" {
			stat[attend.Uid].NotSignOut++
		}
		stat[attend.Uid].EarlyMin += int(attend.EarlyMin)
		stat[attend.Uid].LateMin += int(attend.LateMin)
		key := attend.Date.Unix()
		if v, ok := stat[attend.Uid].AbsentDates[key]; ok {
			switch {
			case v.PmAbsent && v.AmAbsent:
				if attend.Status == "absent" || attend.Status == "allAbsent" || attend.Status == "halfAbsent" {
					attend.Status = "normal"
				}
				delete(stat[attend.Uid].AbsentDates, key)
			case !v.PmAbsent && !v.AmAbsent:
				if attend.Status != "absent" && attend.Status != "allAbsent" {
					if attend.Status == "halfAbsent" {
						stat[attend.Uid].Actual += 0.5
						stat[attend.Uid].AbsentDates[key].AmAbsent = true

					} else {
						stat[attend.Uid].Actual++
						delete(stat[attend.Uid].AbsentDates, key)
					}
				}
				break
			default:
				if attend.Status == "absent" || attend.Status == "allAbsent" {
					attend.Status = "halfAbsent"
				}
				if attend.Status != "halfAbsent" {
					stat[attend.Uid].Actual += 0.5
					delete(stat[attend.Uid].AbsentDates, key)
				}
				break

			}
		}
		stat[attend.Uid].Attend[key] = attend

	}
	return stat, nil
}
