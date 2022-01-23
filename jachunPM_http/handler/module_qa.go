package handler

import (
	"html/template"
	"jachunPM_http/js"
	"libraries"
	"strconv"
)

func init() {
	httpHandlerMap["GET"]["/qa/index"] = get_qa_index
}
func get_qa_index(data *TemplateData) (err error) {
	if products, err := product_getPairs(data, "nocode"); err != nil {
		return err
	} else if len(products) == 0 {
		data.ws.WriteString(js.Location(createLink("product", "create", nil)))
		return nil
	}
	if data.ws.Query("locate") == "yes" {
		data.ws.Redirect(createLink("bug", "browse", nil))
		return
	}
	productID, _ := strconv.Atoi(data.ws.Query("productID"))
	if productID, branch, err := product_saveState(data, int32(productID)); err != nil {
		return err
	} else if err = qa_setMenu(data, productID, branch); err != nil {
		return err
	}
	data.Data["title"] = data.Lang["qa"]["index"]
	templateOut("qa.index.html", data)
	return nil
}
func qa_setMenu(data *TemplateData, productID, branch int32) (err error) {
	buf := bufpool.Get().(*libraries.MsgBuffer)
	buf.Reset()
	buf.WriteString(`<div class="btn-group angle-btn`)
	if data.App["methodName"] == "index" {
		buf.WriteString(" active")
	}
	buf.WriteString(`"><div class="btn-group">`)
	buf.WriteString(html_a(createLink("qa", "index", "locate=no"), data.Lang["qa"]["index"].(string), "", "class='btn'"))
	buf.WriteString("</div></div>")
	if err = product_setMenu(data, productID, branch, "", true); err != nil {
		return
	}
	data.Data["modulePageNav"] = template.HTML(buf.String()) + data.Data["modulePageNav"].(template.HTML)
	return
}
