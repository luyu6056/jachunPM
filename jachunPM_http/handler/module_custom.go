package handler

import (
	"protocol"
	"strings"
)

func customModelFuncs() {
	global_Funcs["custom_getFeatureMenu"] = func(data *TemplateData, module, method string) []moduleMenu {
		var allMenu []protocol.HtmlKeyValueStr
		if moduleconfig, ok := data.Lang[module]; ok {
			if featureBar, ok := moduleconfig["featureBar"].(map[string][]protocol.HtmlKeyValueStr); ok {
				allMenu = featureBar[method]
			}
		}
		var err error
		if data.Data["custom_queryList"], err = custom_mergeFeatureBar(data, module, method); err != nil {
			panic(err)
		}
		return custom_setMenuByConfig(data, allMenu, nil)
	}

}
func custom_mergeFeatureBar(data *TemplateData, module, method string) ([]*protocol.MSG_USER_Userquery_info, error) {
	queryModule := module
	if module == "project" {
		queryModule = "task"
	} else if module == "product" {
		queryModule = "story"
	}
	out := protocol.GET_MSG_USER_user_getUserqueryByWhere()
	out.Where = map[string]interface{}{
		"Uid":      data.User.Id,
		"Module":   queryModule,
		"Shortcut": true,
	}
	var result *protocol.MSG_USER_user_getUserqueryByWhere_result
	err := data.SendMsgWaitResultToDefault(out, &result)
	return result.List, err
}
func custom_setMenuByConfig(data *TemplateData, allMenu, customMenu []protocol.HtmlKeyValueStr) (menu []moduleMenu) {

	/*if(customMenu){

	          for _,customMenu:= range  customMenuItem){
	              if(!isset(customMenuItem->order)) customMenuItem->order = order;
	              customMenuMap[customMenuItem->name] = customMenuItem;
	              order++;
	          }

	  }else if(module){
	      menuOrder = (module == "main" and isset(lang->menuOrder)) ? lang->menuOrder : (isset(lang->module->menuOrder) ? lang->module->menuOrder : array());
	      if(menuOrder)
	      {
	          ksort(menuOrder);
	          foreach(menuOrder as name)
	          {
	              item = new stdclass();
	              item->name   = name;
	              item->hidden = false;
	              item->order  = order++;
	              customMenuMap[name] = item;
	          }
	      }
	  }*/

	/* Merge fileMenu and customMenu.
	   foreach(customMenuMap as name => item)
	   {
	       if(is_object(allMenu) and !isset(allMenu->name)) allMenu->name = item;
	       if(is_array(allMenu)  and !isset(allMenu[name])) allMenu[name] = item;
	   }*/

	for _, kv := range allMenu {

		module := ""
		method := ""
		//class := ""
		//subModule := ""
		//subMenu := ""
		//alias := ""

		var link []string
		label := kv.Value
		_hasPriv := true
		if link = strings.Split(kv.Value, "|"); len(link) > 2 {

			label, module, method = link[0], link[1], link[2]
			_hasPriv = hasPriv(data, module, method)
		}

		if _hasPriv {
			var itemLink map[string]string
			if module != "" && method != "" {
				itemLink = map[string]string{"module": module, "method": method}
				if len(link) > 3 {
					itemLink["vars"] = link[3]
				}

			}
			//hidden = isset(customMenuMap[name]) && isset(customMenuMap[name]->hidden) && customMenuMap[name]->hidden;
			//if(strpos(name, "QUERY") === 0 and !isset(customMenuMap[name])) hidden = false;
			menuItem := moduleMenu{
				Name: kv.Key,
				Link: itemLink,
				Text: label,
			}

			/* Hidden menu by config in mobile. */
			//if(app->viewType == "mhtml" and isset(config->custom->moblieHidden[menuModuleName]) and in_array(name, config->custom->moblieHidden[menuModuleName])) menuItem->hidden = 1;

			//while(isset(menu[menuItem->order])) menuItem->order++;
			menu = append(menu, menuItem)
		}
	}
	if data.Data["custom_queryList"] != nil {
		menu = append(menu, moduleMenu{
			Name: "QUERY",
			Text: data.Lang["custom"]["common"].(string),
		})
	}
	return
}
