package handler

import (
	"protocol"
	"strconv"
	"strings"
)

func branch_getPairs(data *TemplateData, productID int32, productInfo *protocol.MSG_PROJECT_product_cache, params ...string) []protocol.HtmlKeyValueStr {
	if productInfo == nil {
		productInfo = HostConn.GetProductById(productID)
	}
	var res []protocol.HtmlKeyValueStr
	for _, b := range productInfo.Branchs {
		res = append(res, protocol.HtmlKeyValueStr{strconv.Itoa(int(b.Id)), b.Name})
	}
	if len(params) == 0 || !strings.Contains(params[0], "noempty") {
		if productInfo == nil || productInfo.Type == "normal" {
			return nil
		}
		res = append([]protocol.HtmlKeyValueStr{{"0", data.Lang["branch"]["all"].(string) + data.Lang["product"]["branchName"].(map[string]string)[productInfo.Type]}}, res...)
	}
	return res
}
