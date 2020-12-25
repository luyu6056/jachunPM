package handler

import (
	"fmt"
	"libraries"
	"protocol"
	"strconv"
	"strings"
	"time"

	"github.com/luyu6056/gnet"
)

func init() {
	httpHandlerMap["GET"]["/product/index"] = get_product_index
	httpHandlerMap["GET"]["/product/create"] = get_product_create
	httpHandlerMap["POST"]["/product/create"] = post_product_create
	httpHandlerMap["GET"]["/product/browse"] = get_product_browse
}
func get_product_index(data *TemplateData) (action gnet.Action) {

	if data.ws.Query("locate") == "yes" {
		data.ws.Redirect(createLink("product", "browse", nil))
		return
	}
	//if($this->app->getViewType() != 'mhtml') unset($this->lang->product->menu->index);
	//$productID = $this->product->saveState($productID, $this->products);
	//$branch    = (int)$this->cookie->preBranch;
	//$this->product->setMenu($this->products, $productID, $branch);

	templateOut("product.index.html", data)
	return
}
func get_product_create(data *TemplateData) (action gnet.Action) {
	//$rootID = key($this->products);
	//if($this->session->product) $rootID = $this->session->product;
	//$this->product->setMenu($this->products, $rootID);
	data.Data["groups"], _ = user_getGroupOptionMenu()
	msg, err := HostConn.GetMsg()
	if err != nil {
		data.ws.OutErr(err)
		return
	}
	getuser := protocol.GET_MSG_USER_getPairs()
	getuser.Params = "nodeleted|pofirst|noclosed"
	var res *protocol.MSG_USER_getPairs_result
	err = msg.SendMsgWaitResult(0, getuser, &res)
	if err != nil {
		data.ws.OutErr(err)
		return
	}
	data.Data["poUsers"] = res.List
	getuser.Params = "nodeleted|qdfirst|noclosed"
	var res1 *protocol.MSG_USER_getPairs_result
	err = msg.SendMsgWaitResult(0, getuser, &res1)
	if err != nil {
		data.ws.OutErr(err)
		return
	}
	data.Data["qdUsers"] = res1.List
	getuser.Params = "nodeleted|devfirst|noclosed"
	var res2 *protocol.MSG_USER_getPairs_result
	err = msg.SendMsgWaitResult(0, getuser, &res2)
	if err != nil {
		data.ws.OutErr(err)
		return
	}
	data.Data["rdUsers"] = res2.List
	getuser.Put()
	var productTypeList []protocol.HtmlKeyValueStr
	for _, v := range data.Lang["product"]["typeList"].([]protocol.HtmlKeyValueStr) {
		tip, _ := data.Lang["product"]["typeTips"].(map[string]string)[v.Key]
		productTypeList = append(productTypeList, protocol.HtmlKeyValueStr{v.Key, v.Value + tip})

	}
	data.Data["productTypeList"] = productTypeList
	getLinePairs := protocol.GET_MSG_PROJECT_tree_getLinePairs()
	var res3 *protocol.MSG_PROJECT_tree_getLinePairs_result
	err = msg.SendMsgWaitResult(0, getLinePairs, &res3)
	if err != nil {
		data.ws.OutErr(err)
		return
	}
	res3.List = append([]protocol.HtmlKeyValueStr{{"", ""}}, res3.List...)
	data.Data["lines"] = res3.List
	templateOut("product.create.html", data)
	res.Put()
	res1.Put()
	res2.Put()
	res3.Put()
	return
}
func post_product_create(data *TemplateData) (action gnet.Action) {
	if !data.ajaxCheckPost() {
		return
	}
	msg, err := HostConn.GetMsg()
	if err != nil {
		data.ajaxResult(false, err.Error())
		return
	}
	out := protocol.GET_MSG_PROJECT_product_insert()
	insert := protocol.GET_MSG_PROJECT_product_cache()
	insert.Status = "normal"
	insert.CreatedBy = data.User.Id
	insert.CreatedDate = time.Now().Unix()

	for key, v := range data.ws.GetAllPost() {
		switch key {
		case "acl":
			insert.Acl = v[0]
		case "whitelist":
			for _, sid := range v {
				id, _ := strconv.Atoi(sid)
				if id > 0 {
					insert.Whitelist = append(insert.Whitelist, int32(id))
				}
			}
		case "name":
			insert.Name = v[0]
		case "code":
			insert.Code = v[0]
		case "line":
			id, _ := strconv.Atoi(v[0])
			insert.Line = int32(id)
		case "PO":
			id, _ := strconv.Atoi(v[0])
			insert.PO = int32(id)
		case "QD":
			id, _ := strconv.Atoi(v[0])
			insert.QD = int32(id)
		case "RD":
			id, _ := strconv.Atoi(v[0])
			insert.RD = int32(id)
		case "type":
			insert.Type = v[0]
		case "desc":
			insert.Desc = v[0]

		}
	}
	m, _ := libraries.Preg_match_result(`<img src="/file/tmpimg\?fileID=(\d+)&amp;t=([^"]+)" alt="" \/>`, insert.Desc, -1)
	var uploaderr error
	var newimgids []int64
	for _, match := range m {
		b, ok := file_getTempFile(match[1])
		if ok {
			upload := protocol.GET_MSG_FILE_upload()
			upload.AddBy = data.User.Id
			upload.Data = b
			upload.Name = time.Now().Format("20060102") + "_" + match[1] + "." + match[2]
			var res *protocol.MSG_FILE_upload_result
			uploaderr = msg.SendMsgWaitResult(0, upload, &res)
			if uploaderr == nil {
				newimgids = append(newimgids, res.FileID)
				insert.Desc = strings.ReplaceAll(insert.Desc, match[0], `<img src="/file/read?fileID=`+strconv.FormatInt(res.FileID, 10)+` alt="" />`)
			}
			res.Put()
			if uploaderr != nil {
				deleteimg := protocol.GET_MSG_FILE_DeleteByID()
				for _, id := range newimgids {
					deleteimg.FileID = id
					msg.SendMsg(0, deleteimg)
				}
				deleteimg.Put()
				data.ajaxResult(false, map[string]string{"desc": fmt.Sprintf(data.Lang["file"]["imguploadFail"].(string), uploaderr)})
				return
			}
		} else {
			insert.Desc = strings.ReplaceAll(insert.Desc, match[0], "")
		}

	}
	insert.Desc = libraries.Html2bbcode(insert.Desc)
	defer func() {
		if err != nil { //以下使用err来判断图片删除
			deleteimg := protocol.GET_MSG_FILE_DeleteByID()
			for _, id := range newimgids {
				deleteimg.FileID = id
				msg.SendMsg(0, deleteimg)
			}
			deleteimg.Put()
		}
	}()
	out.Data = insert
	out.DocName = data.Lang["doclib"]["main"].(map[string]string)["product"]
	var res *protocol.MSG_PROJECT_product_insert_result
	err = msg.SendMsgWaitResult(0, out, &res)
	if err != nil {
		data.ajaxResult(false, err.Error())
		return
	}
	locate := createLink("product", "browse", []string{"productID=", strconv.Itoa(int(res.ID))})
	data.ajaxResult(true, data.Lang["common"]["saveSuccess"], locate)
	out.Put()
	res.Put()
	return
}
func get_product_browse(data *TemplateData) (action gnet.Action) {
	templateOut("product.browse.html", data)
	return
}
