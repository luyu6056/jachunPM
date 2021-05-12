package handler

import (
	"jachunPM_project/db"
	"protocol"
	"strconv"
)

func branch_getByProducts(data *protocol.MSG_PROJECT_branch_getByProducts, in *protocol.Msg) {
	out := protocol.GET_MSG_PROJECT_branch_getByProducts_result()
	var branches []*protocol.MSG_PROJECT_branch_info
	products := map[int32]*protocol.MSG_PROJECT_product_cache{}
	for _, productId := range data.Products {
		if product := HostConn.GetProductById(productId); product != nil {
			products[product.Id] = product
			for _, b := range product.Branchs {
				if !b.Deleted {
					branches = append(branches, b)
				}

			}
		}
	}
	protocol.Order_branch(branches, nil)
	for i := len(data.AppendBranch) - 1; i >= 0; i-- {
		for _, b := range branches {
			if b.Id == data.AppendBranch[i] {
				data.AppendBranch = append(data.AppendBranch[:i], data.AppendBranch[i+1:]...)
				break
			}
		}
	}
	if len(data.AppendBranch) > 0 {
		var appendBranchs []*protocol.MSG_PROJECT_branch_info
		if err := in.DB.Table(db.TABLE_BRANCH).Where(map[string]interface{}{"Id": data.AppendBranch}).Select(&appendBranchs); err != nil {
			in.WriteErr(err)
			return
		}
		branches = append(branches, appendBranchs...)
	}
	out.List = make(map[int32][]protocol.HtmlKeyValueStr)
	for _, b := range branches {
		product := products[b.Product]
		if product == nil {
			product = HostConn.GetProductById(b.Product)
		}
		if product != nil {
			products[b.Product] = product
			if product.Type == "normal" {
				if out.List[product.Id] == nil {
					out.List[product.Id] = []protocol.HtmlKeyValueStr{{"0", ""}}
				}
			} else {
				if out.List[product.Id] == nil {
					out.List[product.Id] = []protocol.HtmlKeyValueStr{{strconv.Itoa(int(b.Id)), b.Name}}
				} else {
					out.List[product.Id] = append(out.List[product.Id], protocol.HtmlKeyValueStr{strconv.Itoa(int(b.Id)), b.Name})
				}

			}
		}
	}
	in.SendResult(out)
	out.Put()
}
func branch_getPairsByIds(data *protocol.MSG_PROJECT_branch_getPairsByIds, in *protocol.Msg) {
	var branchs []*db.Branch
	if err := in.DB.Table(db.TABLE_BRANCH).Field("Id,Name").Where(map[string]interface{}{"Id": data.Ids}).Limit(0).Select(&branchs); err != nil {
		in.WriteErr(err)
		return
	}
	out := protocol.GET_MSG_PROJECT_project_getPairsByIds_result()
	for _, b := range branchs {
		out.List = append(out.List, protocol.HtmlKeyValueStr{strconv.Itoa(int(b.Id)), b.Name})
	}
	in.SendResult(out)
	out.Put()
}
