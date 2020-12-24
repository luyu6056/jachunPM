package handler

import (
	"protocol"

	"github.com/luyu6056/gnet"
)

func init() {
	httpHandlerMap["GET"]["/product/index"] = get_product_index
	httpHandlerMap["GET"]["/product/create"] = get_product_create
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
	res, err := msg.SendMsgWaitResult(0, getuser)
	if err != nil {
		data.ws.OutErr(err)
		return
	}
	if v, ok := res.(*protocol.MSG_USER_getPairs_result); ok {
		data.Data["poUsers"] = v.List
	}
	getuser.Params = "nodeleted|qdfirst|noclosed"
	res, err = msg.SendMsgWaitResult(0, getuser)
	if err != nil {
		data.ws.OutErr(err)
		return
	}
	if v, ok := res.(*protocol.MSG_USER_getPairs_result); ok {
		data.Data["qdUsers"] = v.List
	}
	getuser.Params = "nodeleted|devfirst|noclosed"
	res, err = msg.SendMsgWaitResult(0, getuser)
	if err != nil {
		data.ws.OutErr(err)
		return
	}
	if v, ok := res.(*protocol.MSG_USER_getPairs_result); ok {
		data.Data["rdUsers"] = v.List
	}
	getuser.Put()
	var productTypeList []protocol.HtmlKeyValueStr
	for _, v := range data.Lang["product"]["typeList"].([]protocol.HtmlKeyValueStr) {
		tip, _ := data.Lang["product"]["typeTips"].(map[string]string)[v.Key]
		productTypeList = append(productTypeList, protocol.HtmlKeyValueStr{v.Key, v.Value + tip})

	}
	data.Data["productTypeList"] = productTypeList
	getLinePairs := protocol.GET_MSG_PROJECT_tree_getLinePairs()
	res, err = msg.SendMsgWaitResult(0, getLinePairs)
	if err != nil {
		data.ws.OutErr(err)
		return
	}
	if v, ok := res.(*protocol.MSG_PROJECT_tree_getLinePairs_result); ok {
		v.List = append([]protocol.HtmlKeyValueStr{{"", ""}}, v.List...)
		data.Data["lines"] = v.List
	}

	templateOut("product.create.html", data)
	return
}
func post_product_create(data *TemplateData) (action gnet.Action) {
	return
}
