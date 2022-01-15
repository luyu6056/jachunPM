package handler

import (
	"jachunPM_project/db"
	"libraries"
	"protocol"
	"strconv"
	"time"
)

var lastflush int64

func HandleTick(t time.Time) {

	var products []*protocol.MSG_PROJECT_product_cache
	var modules []*protocol.MSG_PROJECT_tree_cache
	var branchs []*protocol.MSG_PROJECT_branch_info
	var projects []*protocol.MSG_PROJECT_project_cache
	now := time.Now().Unix()
	if lastflush < now-86400 {
		//检查是否缺少默认admin
		err := HostConn.DB.Table(db.TABLE_PRODUCT).Limit(0).Select(&products)
		if err != nil {
			panic("检查product失败" + err.Error())
		}

		err = HostConn.DB.Table(db.TABLE_MODULE).Limit(0).Select(&modules)
		if err != nil {
			panic("检查module失败" + err.Error())
		}
		err = HostConn.DB.Table(db.TABLE_BRANCH).Order("Order desc,id desc").Limit(0).Select(&branchs)
		if err != nil {
			panic("检查branch失败" + err.Error())
		}
		for _, p := range products {
			for i := len(branchs) - 1; i >= 0; i-- {
				b := branchs[i]
				if b.Product == p.Id {
					p.Branchs = append(p.Branchs, b)
					branchs = append(branchs[:i], branchs[i+1:]...)
				}
			}
		}
		err = HostConn.DB.Table(db.TABLE_PROJECT).Limit(0).Select(&projects)
		if err != nil {
			panic("检查projects失败" + err.Error())
		}

		//检查product缓存
		res, err := HostConn.CacheGetPath(protocol.ProjectServerNo, protocol.PATH_PROJECT_PRODUCT_CACHE)
		if err != nil {
			libraries.DebugLog("检查host product缓存失败", err)
			return
			return
		}
		buf := protocol.BufPoolGet()
		buf1 := protocol.BufPoolGet()
		for _, b := range res {
			buf.Reset()
			buf.Write(b)
			buf1.Reset()
			if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_PROJECT_product_cache); ok {
				var find bool
				for i := len(products) - 1; i >= 0; i-- {
					p := products[i]
					if p.Id == v.Id {
						protocol.WRITE_MSG_PROJECT_product_cache(p, buf1)
						if buf1.String() == string(b[4:]) {
							products = append(products[:i], products[i+1:]...)
						}
						find = true
					}
				}
				if !find {
					HostConn.CacheDel(protocol.PATH_PROJECT_PRODUCT_CACHE, strconv.Itoa(int(v.Id)))
				}
				v.Put()
			}
		}
		//检查module缓存
		res, err = HostConn.CacheGetPath(protocol.ProjectServerNo, protocol.PATH_PROJECT_TREE_CACHE)
		if err != nil {
			libraries.DebugLog("检查host module缓存失败", err)
			return
		}
		for _, b := range res {
			buf.Reset()
			buf.Write(b)
			buf1.Reset()
			if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_PROJECT_tree_cache); ok {
				var find bool
				for i := len(modules) - 1; i >= 0; i-- {
					m := modules[i]
					if m.Id == v.Id {
						protocol.WRITE_MSG_PROJECT_tree_cache(m, buf1)
						if buf1.String() == string(b[4:]) {
							modules = append(modules[:i], modules[i+1:]...)
						}
						find = true
					}
				}
				if !find {
					HostConn.CacheDel(protocol.PATH_PROJECT_TREE_CACHE, strconv.Itoa(int(v.Id)))
				}
				v.Put()
			}
		}
		//检查projects缓存
		res, err = HostConn.CacheGetPath(protocol.ProjectServerNo, protocol.PATH_PROJECT_PROJECT_CACHE)
		if err != nil {
			libraries.DebugLog("检查host module缓存失败", err)
			return
		}
		for _, b := range res {
			buf.Reset()
			buf.Write(b)
			buf1.Reset()
			if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_PROJECT_project_cache); ok {
				var find bool
				for i := len(projects) - 1; i >= 0; i-- {
					p := projects[i]
					if p.Id == v.Id {
						protocol.WRITE_MSG_PROJECT_project_cache(p, buf1)
						if buf1.String() == string(b[4:]) {
							projects = append(projects[:i], projects[i+1:]...)
						}
						find = true
					}
				}
				if !find {
					HostConn.CacheDel(protocol.PATH_PROJECT_PROJECT_CACHE, strconv.Itoa(int(v.Id)))
				}
				v.Put()
			}
		}
		buf.Reset()
		protocol.BufPoolPut(buf)
		buf1.Reset()
		protocol.BufPoolPut(buf1)
		lastflush = now
	} else {
		//检查是否需要更新缓存
		err := HostConn.DB.Table(db.TABLE_PRODUCT).Prepare().Where("UNIX_TIMESTAMP(TimeStamp)>?", t.Unix()-protocol.RpcTickDefaultTime*2).Limit(0).Select(&products)
		if err != nil {
			libraries.ReleaseLog("检查product刷新缓存失败%v", err)
		}
		err = HostConn.DB.Table(db.TABLE_MODULE).Prepare().Where("UNIX_TIMESTAMP(TimeStamp)>?", t.Unix()-protocol.RpcTickDefaultTime*2).Find(&modules)
		if err != nil {
			libraries.ReleaseLog("检查module刷新缓存失败%v", err)
		}
		err = HostConn.DB.Table(db.TABLE_BRANCH).Prepare().Where("UNIX_TIMESTAMP(TimeStamp)>?", t.Unix()-protocol.RpcTickDefaultTime*2).Limit(0).Select(&branchs)
		if err != nil {
			libraries.ReleaseLog("检查branch刷新缓存失败%v", err)
		}

		err = HostConn.DB.Table(db.TABLE_PROJECT).Prepare().Where("UNIX_TIMESTAMP(TimeStamp)>?", t.Unix()-protocol.RpcTickDefaultTime*2).Limit(0).Select(&projects)
		if err != nil {
			libraries.ReleaseLog("检查projects刷新缓存失败%v", err)
		}
		for _, p := range products {
			for i := len(branchs) - 1; i >= 0; i-- {
				b := branchs[i]
				if b.Product == p.Id {
					p.Branchs = append(p.Branchs, b)
					branchs = append(branchs[:i], branchs[i+1:]...)
					break
				}
			}
		}
		if len(branchs) > 0 {
			productids := map[int32]int{}
			for _, b := range branchs {
				productids[b.Product] = 1
			}
			var ids []int32
			for id := range productids {
				ids = append(ids, id)
			}
			var new_products []*protocol.MSG_PROJECT_product_cache
			err = HostConn.DB.Table(db.TABLE_PRODUCT).Where(map[string]interface{}{"Id": ids}).Limit(0).Select(&new_products)
			if err != nil {
				libraries.ReleaseLog("检查product刷新缓存失败%v", err)
			}
			products = append(products, new_products...)
			for _, p := range products {
				for i := len(branchs) - 1; i >= 0; i-- {
					b := branchs[i]
					if b.Product == p.Id {
						p.Branchs = append(p.Branchs, b)
						branchs = append(branchs[:i], branchs[i+1:]...)
						break
					}
				}
			}
		}

	}

	//同步缓存

	for _, v := range products {
		product_setCache(v.Id)
	}

	for _, v := range modules {
		tree_setCache(v.Id)
	}
	for _, v := range projects {
		project_setCache(v.Id)
	}
}
