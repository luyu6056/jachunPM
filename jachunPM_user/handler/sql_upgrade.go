package handler

import (
	"jachunPM_user/db"
	"libraries"
	"protocol"
	"strconv"
	"time"
)

//老格式的表升级为新格式的表
func mysqlUpgrade() {
	var uids []int32
	HostConn.DB.Table(db.TABLE_USER).Where().Delete()
	var rows []*db.User
	err := HostConn.DB.Table("zt_user").Field("`Id`,`Dept`,`Account`,`Password`,`Role`,`Realname`,IF(`Gender`='f',0,1) as Gender,`Email`,`QQ`,`Mobile`,`Phone`,`Weixin`,`Dingding`,`Address`,`Zipcode`,`Join`,`Visits`,`Ip`,`Last`,`Fails`,`Locked`,`Deleted`,`AttendNo`").Limit(0).Select(&rows)
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
		uids = append(uids, u.Id)
	}
	rows=append(rows,&db.User{Id: protocol.SYSTEMUSER, Realname: "系统缓存",Account: "systemConfig",Password: "-", Salt: "-",Deleted: true})

	_, err = HostConn.DB.Table(db.TABLE_USER).InsertAll(rows)
	//_, err = HostConn.DB.Table(db.TABLE_TASK).InsertAll(tasks_insert)
	libraries.DebugLog("插入User %d 条，错误 %v", len(rows), err)

	HostConn.DB.Table(db.TABLE_TEAM).Where().Delete()
	var teams []*db.Team
	err = HostConn.DB.Table("zt_team").Field("`Id`,`Root`,`Type`,`Account`,`Role`,`Limited`,`Join`,`Days`,`Hours`,`Estimate`,`Consumed`,`Left`,`Order`").Limit(0).Select(&teams)
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

	HostConn.DB.Table(db.TABLE_GROUP).Where().Delete()
	m, err := HostConn.DB.Table("zt_group").SelectMap()
	var groups []*db.Group
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取zt_group表 err:%v", err)
		return
	}
	groupPrivm, err := HostConn.DB.Table("zt_grouppriv").Limit(0).SelectMap()
	if err != nil {
		libraries.ReleaseLog("mysqlUpgrade无法获取zt_grouppriv表 err:%v", err)
		return
	}
	type Acl struct {
		Views    map[string]string `json:"views"`
		Products []string          `json:"products"`
		Projects []string          `json:"projects"`
	}
	for _, row := range m {
		id, _ := strconv.Atoi(row["id"])
		group := &db.Group{
			Name:      row["name"],
			Role:      row["role"],
			Desc:      row["desc"],
			TimeStamp: time.Now(),
		}
		group.Id = int32(id)
		var acl Acl
		libraries.JsonUnmarshalStr(row["acl"], &acl)
		for name := range acl.Views {
			group.Acl = append(group.Acl, name)
		}
		for _, ids := range acl.Products {
			id, _ := strconv.Atoi(ids)
			if id > 0 {
				group.AclProducts = append(group.AclProducts, int32(id))
			}
		}
		for _, ids := range acl.Projects {
			id, _ := strconv.Atoi(ids)
			if id > 0 {
				group.AclProjects = append(group.AclProjects, int32(id))
			}
		}
		if id == 1 {
			group.AclProjects = []int32{-1}
			group.AclProducts = []int32{-1}
		}
		group.Priv = make(map[string]map[string]bool)
		for _, row := range groupPrivm {
			id, _ := strconv.Atoi(row["group"])
			if id == int(group.Id) {
				if group.Priv[row["module"]] == nil {
					group.Priv[row["module"]] = make(map[string]bool)
				}
				group.Priv[row["module"]][row["method"]] = true
			}
		}
		groups = append(groups, group)
	}
	_, err = HostConn.DB.Table(db.TABLE_GROUP).InsertAll(groups)
	//_, err = HostConn.DB.Table(db.TABLE_TASK).InsertAll(tasks_insert)
	libraries.DebugLog("插入group %d 条，错误 %v", len(groups), err)

	in, _ := HostConn.GetMsg()
	out := protocol.GET_MSG_PROJECT_getAllprojectProductID()
	var result *protocol.MSG_PROJECT_getAllprojectProductID_result
	if err := in.SendMsgWaitResult(0, out, &result); err != nil {
		libraries.DebugLog("无法同步user acl信息，无法获取project,product")
		return
	}

	updateUserView(uids, nil, result.ProductID, result.ProjectID, in)
	libraries.DebugLog("成功同步user信息")
}
func init() {
	//time.AfterFunc(time.Second*5, mysqlUpgrade)
}
