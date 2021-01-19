package handler

import (
	"jachunPM_project/db"
	"protocol"
	"strconv"
)

func tree_getLinePairs(data *protocol.MSG_PROJECT_tree_getLinePairs, in *protocol.Msg) {
	var list []*db.Module
	err := in.DB.Table(db.TABLE_MODULE).Field("Id,Name").Where("type = 'line' and deleted = 0").Select(&list)
	if err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_tree_getLinePairs_result()
	for _, v := range list {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(v.Id)), v.Name})
	}
	in.SendResult(out)
	out.Put()
}
func tree_getAllChildId(moduleID int32) (res []int32) {

	var list []*protocol.MSG_PROJECT_tree_cache
	HostConn.DB.Table(db.TABLE_MODULE).Field("Id,Path,Deleted").Limit(0).Select(&list)
	var module *protocol.MSG_PROJECT_tree_cache
	for _, v := range list {
		if v.Id == moduleID {
			module = v
			break
		}
	}
	if module != nil {
		for _, v := range list {
			find := true
			if !v.Deleted && len(v.Path) > len(module.Path) {
				for k, id := range module.Path {
					if v.Path[k] != id {
						find = false
						break
					}
				}
			} else {
				find = false
			}
			if find {
				res = append(res, v.Id)
			}
		}
	}
	return
}

func tree_setCache(m *db.Module) {
	data := protocol.GET_MSG_PROJECT_tree_cache()
	data.Branch = m.Branch
	data.Collector = m.Collector
	data.Deleted = m.Deleted
	data.Grade = m.Grade
	data.Id = m.Id
	data.Name = m.Name
	data.Order = m.Order
	data.Owner = m.Owner
	data.OwnerID = m.OwnerID
	data.Parent = m.Parent
	data.Path = m.Path
	data.Root = m.Root
	data.Short = m.Short
	data.Type = m.Type
	HostConn.CacheSet(protocol.PATH_PROJECT_TREE_CACHE, strconv.Itoa(int(data.Id)), data, 0)
	data.Put()
}
func tree_manageChild(data *protocol.MSG_PROJECT_tree_manageChild, in *protocol.Msg) {
	repeatName, err := tree_checkUnique(&protocol.MSG_PROJECT_tree_cache{
		Id:     data.RootID,
		Type:   data.ViewType,
		Parent: data.ParentModuleID,
	}, data.Modules, nil)
	if err != nil {
		in.WriteErr(err)
		return
	}

	if repeatName {
		out := protocol.GET_MSG_PROJECT_tree_manageChild_result()
		out.Result = protocol.Err_TreeRepeatName
		out.Name = "modules"
		in.SendResult(out)
		out.Put()
		return
	}
	grade := int8(1)
	var parentPath []int32
	if data.ParentModuleID > 0 {
		parentModule := HostConn.GetTreeById(data.ParentModuleID)

		if parentModule != nil {
			grade = parentModule.Grade + 1
			parentPath = parentModule.Path
		}
	}

	var updates []*protocol.MSG_PROJECT_tree_cache
	session, err := in.DB.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer session.EndTransaction()
	var ids []int32
	for _, module := range data.Modules {

		if module.Id == 0 {

			module.Root = data.RootID
			module.Parent = data.ParentModuleID
			//module->branch  = isset(branches[moduleID]) ? branches[moduleID] : 0
			module.Grade = grade
			module.Type = data.ViewType
			id, err := session.Table(db.TABLE_MODULE).Insert(module)
			if err != nil {
				in.WriteErr(err)
				return
			}
			module.Path = append(parentPath, int32(id))
			ids = append(ids, int32(id))
			_, err = session.Table(db.TABLE_MODULE).Prepare().Where("Id=?", id).Update("Path=?", module.Path)
			if err != nil {
				in.WriteErr(err)
				return
			}
		} else {
			ids = append(ids, module.Id)
			update := HostConn.GetTreeById(module.Id)
			if update != nil {
				childs := tree_getAllChildId(module.Id)
				update.Grade = grade
				update.Path = append(parentPath, update.Id)
				update.OwnerID = module.OwnerID
				if data.RootID >= 0 {
					update.Root = data.RootID
					update.Branch = 0
					for _, id := range childs {
						child := HostConn.GetTreeById(id)
						child.Grade = update.Grade + 1
						child.OwnerID = module.OwnerID
						child.Root = data.RootID
						child.Branch = 0
						child.Path = append(update.Path, child.Id)
						ids = append(ids, id)
						updates = append(updates, child)
					}
				}
				//$this->fixModulePath(isset($module->root) ? $module->root : update->root, update->type);
				update.Name = module.Name
				update.Short = module.Short
				updates = append(updates, update)
			}
		}
	}
	if len(updates) > 0 {
		_, err = session.Table(db.TABLE_MODULE).ReplaceAll(updates)
		if err != nil {
			in.WriteErr(err)
			return
		}
	}
	session.CommitCallback(func() {
		var list []*db.Module
		session.Table(db.TABLE_MODULE).Prepare().Where(map[string]interface{}{"id": ids}).Limit(0).Select(&list)
		for _, v := range list {
			tree_setCache(v)
		}
	})
	session.Commit()
	out := protocol.GET_MSG_PROJECT_tree_manageChild_result()
	out.Result = protocol.Success
	in.SendResult(out)
	out.Put()
}
func tree_checkUnique(module *protocol.MSG_PROJECT_tree_cache, modules []*protocol.MSG_PROJECT_tree_cache, branches []int32) (bool, error) {
	if branches == nil && module.Branch > 0 {
		branches = []int32{module.Branch}
	}
	if modules == nil {
		modules = []*protocol.MSG_PROJECT_tree_cache{module}
	}
	for i := 0; i < len(modules); i++ {
		for j := i + 1; j < len(modules); j++ {
			if modules[i].Name == modules[j].Name {
				return false, nil
			}
		}
	}
	where := map[string]interface{}{
		"Root":    module.Root,
		"Type":    module.Type,
		"Parent":  module.Parent,
		"Deleted": false,
	}
	if branches != nil {
		where["Branch"] = branches
	}
	var existsModules []*protocol.MSG_PROJECT_tree_cache
	err := HostConn.DB.Table(db.TABLE_MODULE).Where(where).Select(&existsModules)
	if err != nil {
		return false, err
	}
	for _, module := range modules {
		if module.Name == "" {
			return true, nil
		}
		for _, existsModule := range existsModules {
			if module.Name == existsModule.Name && module.Id != existsModule.Id && (module.Branch == 0 || module.Branch == existsModule.Branch) {
				return true, nil
			}
		}

	}

	return false, nil
}
func tree_updateList(data *protocol.MSG_PROJECT_tree_updateList, in *protocol.Msg) {
	update, err := HostConn.DB.Table(db.TABLE_MODULE).ReplaceAll(data.Modules)
	in.WriteErr(err)
	if update {
		var ids = make([]int32, len(data.Modules))
		for k, v := range data.Modules {
			ids[k] = v.Id
		}
		var list []*db.Module
		HostConn.DB.Table(db.TABLE_MODULE).Where(map[string]interface{}{"id": ids}).Limit(0).Select(&list)
		for _, v := range list {
			tree_setCache(v)
		}
	}
}
func tree_delete(data *protocol.MSG_PROJECT_tree_delete, in *protocol.Msg) {
	_, err := in.DB.Table(db.TABLE_MODULE).Where(map[string]interface{}{"Id": data.Ids}).Delete()
	in.WriteErr(err)
	for _, id := range data.Ids {
		HostConn.CacheDel(protocol.PATH_PROJECT_TREE_CACHE, strconv.Itoa(int(id)))
	}

}
