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
func tree_getAllChildId(moduleID int32) ([]int32, error) {
	if moduleID < 1 {
		return nil, nil
	}
	module := HostConn.GetTreeById(moduleID)
	if module == nil {
		return nil, nil
	}
	var list []db.Module
	err := HostConn.DB.Table(db.TABLE_MODULE).Field("Id").Where("JSON_CONTAINS(Path, '?')").Select(&list)
	var res []int32
	for _, m := range list {
		res = append(res, m.Id)
	}
	return res, err

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
	//branches = isset(data->branch) ? data->branch : array()
	//orders   = isset(data->order)  ? data->order  : array()
	var updates []*protocol.MSG_PROJECT_tree_cache
	session, err := in.DB.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer session.EndTransaction()
	for _, module := range data.Modules {

		/* The new modules. */
		if module.Id == 0 {
			/*if(isset(orders[moduleID]) and !empty(orders[moduleID]))
			  {
			      order = orders[moduleID]
			  }
			  else
			  {
			      order = this->post->maxOrder + i * 10
			      i ++
			  }*/
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
			module.Path = make([]int32, len(parentPath)+1)
			copy(module.Path, parentPath)
			module.Path[len(parentPath)] = int32(id)
			_, err = session.Table(db.TABLE_MODULE).Prepare().Where("Id=?", id).Update("Path=?", module.Path)
			if err != nil {
				in.WriteErr(err)
				return
			}
			/*this->dao->insert(TABLE_MODULE)->data(module)->exec()
			  moduleID  = this->dao->lastInsertID()
			  childPath = parentPath . "moduleID,"
			  this->dao->update(TABLE_MODULE)->set("path")->eq(childPath)->where("id")->eq(moduleID)->limit(1)->exec()*/
		} else {
			update := HostConn.GetTreeById(module.Id)
			if update != nil {
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
		for _, existsModule := range existsModules {
			if module.Name == existsModule.Name && module.Id != existsModule.Id && (module.Branch == 0 || module.Branch == existsModule.Branch) {
				return true, nil
			}
		}

	}

	return false, nil
}
