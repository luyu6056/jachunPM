package handler

import (
	"jachunPM_project/db"
	"libraries"
	"protocol"
	"time"
)

func HandleTick(t time.Time) {
	var products []*protocol.MSG_PROJECT_product_cache
	var modules []*db.Module
	var branchs []*protocol.MSG_PROJECT_branch_info
	firstFlag := protocol.RpcTickStatusFirst

	if HostConn.Status&firstFlag == firstFlag {
		HostConn.Status -= protocol.RpcTickStatusFirst
		//检查是否缺少默认admin
		err := HostConn.DB.Table(db.TABLE_PRODUCT).Limit(0).Select(&products)
		if err != nil {
			panic("检查product失败" + err.Error())
		}

		err = HostConn.DB.Table(db.TABLE_MODULE).Limit(0).Select(&modules)
		if err != nil {
			panic("检查module失败" + err.Error())
		}
		err = HostConn.DB.Table(db.TABLE_BRANCH).Limit(0).Select(&branchs)
		if err != nil {
			panic("检查branch失败" + err.Error())
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

	} else {
		//检查是否需要更新缓存
		err := HostConn.DB.Table(db.TABLE_PRODUCT).Prepare().Where("TimeStamp >?", t.Unix()-protocol.RpcTickDefaultTime*2).Limit(0).Select(&products)
		if err != nil {
			libraries.ReleaseLog("检查product刷新缓存失败%v", err)
		}
		err = HostConn.DB.Table(db.TABLE_MODULE).Prepare().Where("TimeStamp >?", t.Unix()-protocol.RpcTickDefaultTime*2).Find(&modules)
		if err != nil {
			libraries.ReleaseLog("检查module刷新缓存失败%v", err)
		}
		err = HostConn.DB.Table(db.TABLE_BRANCH).Limit(0).Select(&branchs)
		if err != nil {
			libraries.ReleaseLog("检查branch刷新缓存失败%v", err)
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
		var productids []int32
		for _, b := range branchs {
			productids = append(productids, b.Product)
		}
		var new_products []*protocol.MSG_PROJECT_product_cache
		err = HostConn.DB.Table(db.TABLE_PRODUCT).Where(map[string]interface{}{"Id": productids}).Limit(0).Select(&new_products)
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

	//同步缓存

	for _, v := range products {
		product_setCache(v.Id)
	}

	for _, v := range modules {
		tree_setCache(v.Id)
	}
}
