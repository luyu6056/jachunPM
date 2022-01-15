package handler

import (
	"jachunPM_user/db"
	"protocol"
	"strconv"
	"time"
)

func dept_getDataStructure(rootDeptId int32) ([]*protocol.MSG_USER_Dept_cache, error) {
	tree, err := dept_getSons(rootDeptId)
	if err != nil {
		return nil, err
	}
	if len(tree) > 0 {
		for _, node := range tree {
			children, err := dept_getDataStructure(node.Id)
			if err != nil {
				return nil, err
			}
			node.Children = children
		}
	}
	return tree, nil
}
func dept_getSons(deptId int32) (deptList []*protocol.MSG_USER_Dept_cache, err error) {
	res, err := HostConn.CacheGetPath(protocol.UserServerNo, protocol.PATH_USER_DEPT_CACHE)
	if err != nil {
		return nil, err
	}
	buf := protocol.BufPoolGet()
	for _, b := range res {
		buf.Reset()
		buf.Write(b)
		if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_USER_Dept_cache); ok && v.Parent == deptId {
			deptList = append(deptList, v)
		}
	}
	buf.Reset()
	protocol.BufPoolPut(buf)
	if len(deptList) > 0 {
		protocol.Order_dept(deptList, func(a, b *protocol.MSG_USER_Dept_cache) bool { return a.Order < b.Order })
	}
	return
}
func dept_updateFromCache(cacheList []*protocol.MSG_USER_Dept_cache) error {
	var deptList = make([]*db.Dept, len(cacheList))
	for k, v := range cacheList {
		deptList[k] = &db.Dept{
			Grade:     v.Grade,
			Id:        v.Id,
			Manager:   v.Manager,
			Name:      v.Name,
			Order:     v.Order,
			Parent:    v.Parent,
			Path:      v.Path,
			TimeStamp: time.Now(),
		}
	}
	return dept_updateFromList(deptList)
}
func dept_updateFromList(deptList []*db.Dept) error {
	_, err := HostConn.DB.Table(db.TABLE_DEPT).ReplaceAll(deptList)
	if err == nil {
		//刷新缓存
		for _, v := range deptList {
			dept_setCache(v)
		}
	}
	return err
}
func dept_setCache(dept *db.Dept) {
	cache := protocol.GET_MSG_USER_Dept_cache()
	cache.Grade = dept.Grade
	cache.Id = dept.Id
	cache.Name = dept.Name
	cache.Order = dept.Order
	cache.Parent = dept.Parent
	cache.Path = dept.Path
	cache.Manager = dept.Manager
	cache.ManagerName = ""
	if dept.Manager > 0 {
		if user, _ := getUserInfoByID(dept.Manager); user != nil {
			cache.ManagerName = user.Realname
			if cache.ManagerName == "" {
				cache.ManagerName = user.Account
			}
		}
	}
	HostConn.CacheSet(protocol.PATH_USER_DEPT_CACHE, strconv.Itoa(int(dept.Id)), cache, 0)
	cache.Put()
}
func dept_getDeptUserPairs(deptID int32) (list []*protocol.MSG_USER_Pairs, err error) {
	conditions := map[string]interface{}{
		"Deleted": false,
	}
	if deptID > 0 {
		r, err := dept_getAllChildID(deptID)
		if err != nil {
			return nil, err
		}
		if len(r) > 0 {
			conditions["Dept"] = r
		}
	}

	err = HostConn.DB.Table(db.TABLE_USER).Field("Id,Account,Realname").Where(conditions).Order("Account").Select(&list)
	return list, err
}
func dept_getAllChildID(deptID int32) (res []int32, err error) {
	if deptID < 0 {
		return nil, nil
	}
	result, err := HostConn.CacheGetPath(protocol.UserServerNo, protocol.PATH_USER_DEPT_CACHE)
	if err != nil {
		return nil, err
	}
	buf := protocol.BufPoolGet()
	for _, b := range result {
		buf.Reset()
		buf.Write(b)
		if v, ok := protocol.READ_MSG_DATA(buf).(*protocol.MSG_USER_Dept_cache); ok {
			for _, id := range v.Path {
				if id == deptID {
					res = append(res, v.Id)
				}
			}
			v.Put()
		}
	}
	buf.Reset()
	protocol.BufPoolPut(buf)
	return res, nil
}
func dept_delete(data *protocol.MSG_USER_Dept_delete, in *protocol.Msg) {
	out := protocol.GET_MSG_USER_Dept_delete_result()
	defer out.Put()
	if data.DeptId > 0 {
		list, err := dept_getSons(data.DeptId)
		if err != nil {
			in.WriteErr(err)
			return
		}
		if len(list) > 0 {
			out.Result = protocol.Err_DeptDeleteHasSons
			in.SendResult(out)
			return
		}
		user, err := dept_getDeptUserPairs(data.DeptId)
		if err != nil {
			in.WriteErr(err)
			return
		}
		if len(user) > 0 {
			out.Result = protocol.Err_DeptDeletehasUsers
			in.SendResult(out)
			return
		}
	}
	_, err := in.DB.Table(db.TABLE_DEPT).Where("Id=?", data.DeptId).Delete()
	if err != nil {
		in.WriteErr(err)
		return
	}
	HostConn.CacheDel(protocol.PATH_USER_DEPT_CACHE, strconv.Itoa(int(data.DeptId)))
	out.Result = protocol.Success
	in.SendResult(out)
}
