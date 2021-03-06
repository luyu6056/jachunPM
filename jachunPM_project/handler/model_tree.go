package handler

import (
	"jachunPM_project/db"
	"mysql"
	"protocol"
	"strconv"
	"time"
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

func tree_setCache(id int32) {
	data := protocol.GET_MSG_PROJECT_tree_cache()
	HostConn.DB.Table(db.TABLE_MODULE).Prepare().Where("Id=?", id).Find(&data)
	if data.Id != 0 {
		HostConn.CacheSet(protocol.PATH_PROJECT_TREE_CACHE, strconv.Itoa(int(data.Id)), data, 0)
	}
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
	session, err := in.BeginTransaction()
	if err != nil {
		in.WriteErr(err)
		return
	}
	defer session.Rollback()
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
		for _, v := range ids {
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
	for _, v := range data.Modules {
		v.TimeStamp = time.Now()
	}
	update, err := HostConn.DB.Table(db.TABLE_MODULE).ReplaceAll(data.Modules)
	in.WriteErr(err)
	if update {
		for _, v := range data.Modules {
			tree_setCache(v.Id)
		}
	}
}
func tree_delete(data *protocol.MSG_PROJECT_tree_delete, in *protocol.Msg) {
	_, err := in.DB.Table(db.TABLE_MODULE).Where(map[string]interface{}{"Id": data.Ids}).Update(map[string]interface{}{"Deleted": true, "TimeStamp": time.Now()})
	in.WriteErr(err)
	for _, id := range data.Ids {
		tree_setCache(id)
	}

}
func tree_getParents(moduleID int32) (list []*protocol.MSG_PROJECT_tree_cache, err error) {
	if moduleID == 0 {
		return
	}
	module := HostConn.GetTreeById(moduleID)
	if len(module.Path) == 0 {
		return
	}
	err = HostConn.DB.Table(db.TABLE_MODULE).Where(map[string]interface{}{"Id": module.Path, "Deleted": false}).Order("Grade").Select(&list)
	return
}
func tree_getPairsByIds(data *protocol.MSG_PROJECT_tree_getPairsByIds, in *protocol.Msg) {
	var trees []*db.Module
	if err := in.DB.Table(db.TABLE_MODULE).Field("Id,Name").Where(map[string]interface{}{"Id": data.Ids}).Limit(0).Select(&trees); err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_tree_getPairsByIds_result()
	for _, t := range trees {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(t.Id)), t.Name})
	}
	in.SendResult(out)
	out.Put()
}
func tree_getTaskTreeModules(data *protocol.MSG_PROJECT_tree_getTaskTreeModules, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_tree_getTaskTreeModules_result()
	out.ProjectModules = map[int32]int32{}

	//$field = $parent ? 'path' : 'id';

	//默认true if($linkStory){
	project := HostConn.GetProjectById(data.ProjectID)
	var modules []*db.Module
	err := in.DB.Table(db.TABLE_MODULE).Alias("t1").Field("t1.Path as Path,t1.Id as Id").LeftJoin(db.TABLE_STORY).Alias("t2").On("t2.Module = t1.Id").Where(map[string]interface{}{"t2.Id": project.Storys, "t1.Deleted": false, "t2.Deleted": false}).Limit(0).Select(&modules)
	if err != nil {
		in.WriteErr(err)
		return
	}
	for _, m := range modules {
		if data.Parent {
			for _, id := range m.Path {
				out.ProjectModules[id] = id
			}
		} else {
			out.ProjectModules[m.Id] = m.Id
		}
	}
	if err = in.DB.Table(db.TABLE_MODULE).Field("Path,Id").Where(map[string]interface{}{"Root": data.ProjectID, "Type": "task", "Deleted": false}).Limit(0).Select(&modules); err != nil {
		in.WriteErr(err)
		return
	}

	/* Add task paths of this project.*/
	for _, m := range modules {
		if data.Parent {
			for _, id := range m.Path {
				out.ProjectModules[id] = id
			}
		} else {
			out.ProjectModules[m.Id] = m.Id
		}
	}
	if err = in.DB.Table(db.TABLE_MODULE).Alias("t1").Field("t1.Path as Path,t1.Id as Id").LeftJoin(db.TABLE_TASK).Alias("t2").On("t2.Module = t1.Id").Where(map[string]interface{}{
		"t2.Module":  []interface{}{mysql.WhereOperatorNE, 0},
		"t2.Project": data.ProjectID,
		"t1.Deleted": false, "t2.Deleted": false,
		"t1.Type": "story",
	}).Limit(0).Select(&modules); err != nil {
		in.WriteErr(err)
		return
	}
	for _, m := range modules {
		if data.Parent {
			for _, id := range m.Path {
				out.ProjectModules[id] = id
			}
		} else {
			out.ProjectModules[m.Id] = m.Id
		}
	}
	in.SendResult(out)
	out.Put()
}
