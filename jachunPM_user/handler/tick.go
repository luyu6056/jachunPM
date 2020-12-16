package handler

import (
	"jachunPM_user/db"
	"libraries"
	"math/rand"
	"protocol"
	"strconv"
	"time"
)

func HandleTick(t time.Time) {

	var users []*db.User
	var company *db.Company
	var deptlist []*db.Dept
	firstFlag := protocol.RpcTickStatusFirst | protocol.RpcTickStatusFirst
	if HostConn.Status&firstFlag == firstFlag {
		HostConn.Status -= protocol.RpcTickStatusFirst
		//检查是否缺少默认admin
		err := db.DB.Table(db.TABLE_USER).Limit(0).Select(&users)
		if err != nil {
			panic("检查用户数量失败" + err.Error())
		}
		//创建一个默认密码为123456的admin
		if len(users) == 0 {
			admin := &db.User{
				Id:       1,
				Dept:     0,
				Account:  "admin",
				Salt:     libraries.SHA256_S(strconv.Itoa(int(rand.Int63()))),
				Realname: "admin",
			}
			admin.Password = libraries.SHA256_S("123456" + admin.Salt)
			_, err = db.DB.Table(db.TABLE_USER).Insert(admin)
			if err != nil {
				panic("初始化admin用户失败" + err.Error())
			}
			users = append(users, admin)
		}
		err = db.DB.Table(db.TABLE_COMPANY).Limit(0).Find(&company)
		if err != nil {
			panic("检查公司信息失败" + err.Error())
		}
		if company == nil {
			//插入一个公司信息
			company = &db.Company{
				Name:   "杰骏数码",
				Admins: []string{"admin"},
			}
			_, err = db.DB.Table(db.TABLE_COMPANY).Insert(company)
			if err != nil {
				panic("初始化company失败" + err.Error())
			}
		}
		err = db.DB.Table(db.TABLE_DEPT).Limit(0).Select(&deptlist)
		if err != nil {
			panic("检查dept失败" + err.Error())
		}
	} else {
		//检查是否需要更新缓存
		err := db.DB.Table(db.TABLE_USER).Prepare().Where("TimeStamp >?", t.Unix()-protocol.RpcTickDefaultTime*2).Select(&users)
		if err != nil {
			libraries.ReleaseLog("检查user刷新缓存失败%v", err)
		}
		err = db.DB.Table(db.TABLE_COMPANY).Prepare().Where("TimeStamp >?", t.Unix()-protocol.RpcTickDefaultTime*2).Find(&company)
		if err != nil {
			libraries.ReleaseLog("检查company刷新缓存失败%v", err)
		}
		err = db.DB.Table(db.TABLE_DEPT).Prepare().Where("TimeStamp >?", t.Unix()-protocol.RpcTickDefaultTime*2).Select(&deptlist)
		if err != nil {
			libraries.ReleaseLog("检查dept刷新缓存失败%v", err)
		}
	}

	//同步缓存
	if len(users) > 0 {
		cache := protocol.GET_MSG_USER_INFO_cache()
		for _, user := range users {

			cache.Account = user.Account
			cache.AttendNo = user.AttendNo
			cache.ClientLang = user.ClientLang
			cache.Commiter = user.Commiter
			cache.Dept = user.Dept
			cache.Email = user.Email
			cache.Fails = user.Fails
			cache.Gender = user.Gender
			cache.Id = user.Id
			cache.Ip = user.Ip
			cache.Join = user.Join.Unix()
			cache.Last = user.Last.Unix()
			cache.Locked = user.Locked.Unix()
			cache.Mobile = user.Mobile
			cache.Realname = user.Realname
			cache.Role = user.Role
			cache.Visits = user.Visits
			cache.Deleted = user.Deleted
			HostConn.CacheSet(protocol.PATH_USER_INFO_CACHE, strconv.Itoa(int(user.Id)), cache, 0)
		}
		cache.Put()
	}

	if company != nil {
		cache := protocol.GET_MSG_USER_Company_cache()
		cache.Id = company.Id
		cache.Name = company.Name
		cache.Phone = company.Phone
		cache.Fax = company.Fax
		cache.Address = company.Address
		cache.Zipcode = company.Zipcode
		cache.Website = company.Website
		cache.Backyard = company.Backyard
		cache.Admins = company.Admins
		cache.Deleted = company.Deleted
		HostConn.CacheSet(protocol.PATH_USER_COMPANY_CACHE, strconv.Itoa(int(company.Id)), cache, 0)
		cache.Put()
	}

	for _, dept := range deptlist {
		dept_setCache(dept)
	}

}
