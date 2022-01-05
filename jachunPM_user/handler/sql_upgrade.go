package handler

import (
	"jachunPM_user/db"
	"libraries"
	"strconv"
)

//老格式的表升级为新格式的表
func mysqlUpgrade() {
	HostConn.DB.Table(db.TABLE_USER).Where().Delete()
	var rows []*db.User
	err := HostConn.DB.Table("zt_user").Field("`Id`,`Dept`,`Account`,`Password`,`Role`,`Realname`,IF(`Gender`='f',0,1) as Gender,`Email`,`QQ`,`Mobile`,`Phone`,`Weixin`,`Dingding`,`Address`,`Zipcode`,`Join`,`Visits`,`Ip`,`Last`,`Fails`,`Locked`,`Deleted`,`AttendNo`").Select(&rows)
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取zt_user表 err:%v", err)
	}
	for _, u := range rows {
		m, _ := HostConn.DB.Table("zt_usergroup").Where("account=?", u.Account).SelectMap()
		for _, r := range m {
			g, _ := strconv.Atoi(r["group"])
			if g > 0 {
				u.Group = append(u.Group, int32(g))
			}
		}
		u.Password = "323fb4507cc28c94e26c1273987027a0ae935bda1af539cc880933f74c557702"
		u.Salt = "2909941e479e5b177372ca62223e108975cabc6719732523d4f0a82176da343c"
	}
	_, err = HostConn.DB.Table(db.TABLE_USER).InsertAll(rows)
	//_, err = HostConn.DB.Table(db.TABLE_TASK).InsertAll(tasks_insert)
	libraries.DebugLog("插入User %d 条，错误 %v", len(rows), err)
	HostConn.DB.Table(db.TABLE_TEAM).Where().Delete()
	var teams []*db.Team
	err = HostConn.DB.Table("zt_team").Field("`Id`,`Root`,`Type`,`Account`,`Role`,`Limited`,`Join`,`Days`,`Hours`,`Estimate`,`Consumed`,`Left`,`Order`").Select(&teams)
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取zt_user表 err:%v", err)
	}
	for _, t := range teams {
		for _, u := range rows {
			if u.Account == t.Account {
				t.Uid = u.Id
				break
			}
		}
	}
	_, err = HostConn.DB.Table(db.TABLE_TEAM).InsertAll(teams)
	//_, err = HostConn.DB.Table(db.TABLE_TASK).InsertAll(tasks_insert)
	libraries.DebugLog("插入team %d 条，错误 %v", len(teams), err)
}
func init() {
	//go time.AfterFunc(time.Second*5, mysqlUpgrade)
}
