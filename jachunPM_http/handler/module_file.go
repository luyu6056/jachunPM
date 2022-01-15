package handler

import (
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	common_image "image"
	"io"
	"jachunPM/image"
	"jachunPM_http/js"
	"libraries"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"

	"protocol"

	"github.com/luyu6056/cache"
	"github.com/rubenfonseca/fastimage"
	"golang.org/x/image/bmp"
	"golang.org/x/image/webp"
)

func init() {
	httpHandlerMap["POST"]["/file/ajaxPasteImage"] = post_file_ajaxPasteImage
	httpHandlerMap["GET"]["/file/tmpimg"] = get_file_tmpimg
	httpHandlerMap["GET"]["/file/read"] = get_file_read
	httpHandlerMap["GET"]["/file/buildform"] = get_file_buildform
	httpHandlerMap["POST"]["/file/ajaxUploadTmp"] = post_file_ajaxUploadTmp
	httpHandlerMap["GET"]["/file/download"] = get_file_download
	httpHandlerMap["GET"]["/file/edit"] = get_file_edit
	httpHandlerMap["POST"]["/file/edit"] = post_file_edit
	httpHandlerMap["GET"]["/file/delete"] = get_file_delete
	httpHandlerMap["GET"]["/file/buildExportTPL"] = get_file_buildExportTPL
}
func fileTemplateFuncs() {
	global_Funcs["file_printFiles"] = func(oldData *TemplateData, files []*protocol.MSG_FILE_getByID_result, fieldset bool, Type string) template.HTML {
		path := "/file/file_printFiles"
		data := getFetchInterface(oldData.ws, path, oldData.User)
		if Type != "" {
			data.Data["type"] = Type
		} else {
			data.Data["type"] = "file"
		}
		data.Data["files"] = files
		data.Data["fieldset"] = fieldset
		templateOut("file.printfiles.html", data)
		res := string(data.ws.(*CommonFetch).OutBuffer())
		putFetchInterface(data.ws.(*CommonFetch))
		return template.HTML(res)
	}
	global_Funcs["file_getCustomExport"] = func(data *TemplateData) template.HTML {
		customExport, _ := data.Data["customExport"].(bool)
		isCustomExport := customExport && data.Data["allExportFields"] != nil
		data.Data["isCustomExport"] = isCustomExport
		if isCustomExport {
			var selectedFields []string
			exportFieldPairs := map[string]string{}
			moduleName := data.App["moduleName"].(string)
			moduleLang := data.Lang[moduleName]
			for _, field := range data.Data["allExportFields"].([]string) {
				if v, ok := moduleLang[field].(string); ok {
					exportFieldPairs[field] = v
				} else if v, ok := data.Lang["common"][field].(string); ok {
					exportFieldPairs[field] = v
				} else {
					exportFieldPairs[field] = field
				}

				selectedFields = append(selectedFields, field)
			}
			data.Data["exportFieldPairs"] = exportFieldPairs
			return template.HTML(`<script type="text/javascript">var defaultExportFields='` + strings.Join(selectedFields, ",") + `'</script>`)
		}

		return template.HTML("")
	}
}
func post_file_ajaxPasteImage(data *TemplateData) (err error) {
	editor := data.ws.Post("editor")
	result, err := libraries.Preg_match_result(`^<img src="data:image/([^;]+);base64,([^"]+)"`, editor, 1)
	if len(result) != 1 && err == nil {
		err = errors.New("img data Not Found")
	}
	if err != nil {
		data.ws.SetCode(404)
		data.ws.WriteString("error" + err.Error())
		return
	}
	b, err := base64.StdEncoding.DecodeString(result[0][2])
	if err != nil {
		data.ws.SetCode(404)
		data.ws.WriteString(err.Error())
		return
	}
	imagetype, size, err := getimageTypeAndSizefastimage(b)
	if err != nil || size == nil {
		data.ws.SetCode(404)
		data.ws.WriteString(data.Lang["file"]["error"].(map[string]string)["ErrImgType"])
		return
	}
	ext := result[0][1]
	if imagetype != fastimage.GIF {
		ext = "webp"
		b, err = image.ConvertImgFromFastimage(b, size, imagetype, "webp", 1, 1, 80)
		if err != nil {
			data.ws.SetCode(404)
			data.ws.WriteString(err.Error())
			return
		}
	}
	newID := commoncache.INCRBY("commoncache", 1)
	strID := strconv.FormatUint(uint64(newID), 10)
	cache.Hset(strID, map[string][]byte{"img": b}, "tmpfile", 86400)
	data.ws.WriteString(`<img src="/file/tmpimg?fileID=` + strID + `&t=` + ext + `" alt="" />`)
	return
}
func get_file_tmpimg(data *TemplateData) (err error) {
	b, ok := file_getTempFile(data.ws.Query("fileID"))
	if !ok {
		data.ws.SetCode(404)
		data.ws.WriteString("img Not Found")
		return
	}
	ext := data.ws.Query("t")
	if ext == "webp" && !strings.Contains(data.ws.Header("Accept"), "image/webp") {
		ext = "jpg"
		b, err = image.ConvertImgB(b, ext, 1, 1, 80)
		if err != nil {
			data.ws.SetCode(404)
			data.ws.WriteString(err.Error())
			return nil
		}
	}
	buf := bufpool.Get().(*libraries.MsgBuffer)
	defer func() {
		buf.Reset()
		bufpool.Put(buf)
	}()
	buf.Write(b)
	data.ws.SetContentType("image/" + ext)
	data.ws.Write(buf)
	//data.ws.WriteString(img.Load_str("img"))
	return
}
func getimageTypeAndSizefastimage(data []byte) (imagetype fastimage.ImageType, size *fastimage.ImageSize, err error) {
	buf := bufpool.Get().(*libraries.MsgBuffer)
	defer func() {
		buf.Reset()
		bufpool.Put(buf)
	}()
	buf.Write(data)
	imagetype, size, err = fastimage.DetectImageTypeFromReader(buf)
	if err == io.EOF || size == nil {
		var imgconfig common_image.Config
		switch imagetype {
		case fastimage.BMP: //fastimage未对bmp与webp识别
			buf.Reset()
			buf.Write(data)
			imgconfig, err = bmp.DecodeConfig(buf)
			if imgconfig.Width > 0 && imgconfig.Height > 0 {
				size = &fastimage.ImageSize{Height: uint32(imgconfig.Height), Width: uint32(imgconfig.Width)}
			}
		default:
			buf.Reset()
			buf.Write(data)
			imgconfig, err = webp.DecodeConfig(buf)
			if imgconfig.Width > 0 && imgconfig.Height > 0 {
				size = &fastimage.ImageSize{Height: uint32(imgconfig.Height), Width: uint32(imgconfig.Width)}
			}
		}
	}
	return
}
func file_getTempFile(fileID string) (b []byte, ok bool) {
	img := cache.Hget(fileID, "tmpfile")
	ok = img.Get("img", &b)
	return
}
func get_file_read(data *TemplateData) (err error) {
	out := protocol.GET_MSG_FILE_getByID()
	out.FileID, _ = strconv.ParseInt(data.ws.Query("fileID"), 10, 64)
	var result *protocol.MSG_FILE_getByID_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		data.ws.SetCode(404)
		data.ws.WriteString(err.Error())
		return nil
	}
	if err = checkFileAcl(data, result.ObjectType, result.ObjectID); err != nil {
		data.ws.SetCode(404)
		data.ws.WriteString(err.Error())
		return nil
	}

	if result.Ext == "webp" || result.Ext == "bmp" || result.Ext == "jpg" || result.Ext == "png" {
		getByte := protocol.GET_MSG_FILE_RangeDown()
		getByte.FileID = result.FileID
		getByte.Start = 0
		getByte.End = result.Size
		var r *protocol.MSG_FILE_RangeDown_result
		if err = data.SendMsgWaitResultToDefault(getByte, &r); err != nil {
			data.ws.SetCode(404)
			data.ws.WriteString(err.Error())
			return nil
		}
		if result.Ext == "webp" && !strings.Contains(data.ws.Header("Accept"), "image/webp") {
			result.Ext = "jpg"
			var err error
			r.Byte, err = image.ConvertImgB(r.Byte, result.Ext, 1, 1, 80)
			if err != nil {
				data.ws.SetCode(404)
				data.ws.WriteString(err.Error())
				return nil
			}
		}
		data.ws.SetContentType("image/" + result.Ext)
		buf := bufpool.Get().(*libraries.MsgBuffer)
		buf.Write(r.Byte)
		data.ws.Write(buf)
		buf.Reset()
		bufpool.Put(buf)
		getByte.Put()
		r.Put()
	} else {
		data.ws.RangeDownload(&fileRangeDown{data: data, fileId: result.FileID, size: result.Size}, result.Size, result.Name)
	}
	out.Put()
	result.Put()
	return
}
func get_file_buildform(data *TemplateData) (err error) {
	if HostConn.Status&protocol.RpcClientStatuShutdown == protocol.RpcClientStatuShutdown {
		data.ws.WriteString("文件服务器关闭,无法上传附件")
		return
	}
	filesName := data.ws.Query("filesName")
	if filesName == "" {
		filesName = "files"
	}
	data.Data["filesName"] = filesName
	labelsName := data.ws.Query("filesName")
	if labelsName == "" {
		labelsName = "labels"
	}
	data.Data["labelsName"] = labelsName
	data.Data["examine"] = data.ws.Query("examine")
	data.Data["isadmin"] = data.User.Role == "top"
	data.Data["action"] = data.ws.Query("action")
	templateOut("file.buildform.html", data)
	return
}
func file_descProcessImgURLAnd2Bbcode(data *TemplateData, desc string) (res string, newimgids []int64, uploaderr error) {

	//检查移除失效的老文件
	m, _ := libraries.Preg_match_result(`<img aid="attachimg_0" src="\/file\/read\?fileID=(\d+)" border="0" alt=""  \/>`, desc, -1)
	for _, match := range m {
		check := protocol.GET_MSG_FILE_getByID()
		check.FileID, _ = strconv.ParseInt(match[1], 10, 64)
		check.NoData = true
		var result *protocol.MSG_FILE_getByID_result
		if uploaderr = data.SendMsgWaitResultToDefault(check, &result); uploaderr != nil {
			if strings.Index(uploaderr.Error(), protocol.Err_FileNotFound.String()) == 0 {
				//删掉失效文件
				desc = strings.Replace(desc, match[0], "", 1)
			} else {
				return
			}

		} else {
			result.Put()
		}
		check.Put()
	}
	//转换上传的临时文件
	m, _ = libraries.Preg_match_result(`<img src="/file/tmpimg\?fileID=(\d+)&amp;t=([^"]+)" alt="" \/>`, desc, -1)
	for _, match := range m {
		b, ok := file_getTempFile(match[1])
		if ok {
			upload := protocol.GET_MSG_FILE_upload()
			upload.AddBy = data.User.Id
			upload.Data = b
			upload.Name = time.Now().Format("20060102") + "_" + match[1] + "." + match[2]
			var result *protocol.MSG_FILE_upload_result
			uploaderr = data.Msg.SendMsgWaitResult(0, upload, &result)
			if uploaderr == nil {
				newimgids = append(newimgids, result.FileID)
				desc = strings.ReplaceAll(desc, match[0], `<img src="/file/read?fileID=`+strconv.FormatInt(result.FileID, 10)+` alt="" />`)
			}
			result.Put()
			if uploaderr != nil {
				deleteimg := protocol.GET_MSG_FILE_DeleteByID()
				for _, id := range newimgids {
					deleteimg.FileID = id
					data.Msg.SendMsg(0, deleteimg)
				}
				deleteimg.Put()
				uploaderr = errors.New(fmt.Sprintf(data.Lang["file"]["imguploadFail"].(string), uploaderr))
				return
			}
			upload.Put()
		} else {
			desc = strings.ReplaceAll(desc, match[0], "")
		}

	}

	res = libraries.Html2bbcode(desc)
	return
}
func file_deleteFromIds(data *TemplateData, newimgids []int64) {
	deleteimg := protocol.GET_MSG_FILE_DeleteByID()
	for _, id := range newimgids {
		deleteimg.FileID = id
		data.SendMsgToDefault(deleteimg)
	}
	deleteimg.Put()
}
func file_updateObject(data *TemplateData, fileIds []int64, typ string, id int32) {
	out := protocol.GET_MSG_FILE_updateMapByWhere()
	for _, id := range fileIds {
		out.Where = map[string]interface{}{"Id": id}
		out.Update = map[string]interface{}{
			"ObjectType": typ,
			"ObjectID":   id,
		}
		data.SendMsgToDefault(out)
	}
	out.Put()
}
func file_getByObject(data *TemplateData, object string, ID int32) (file []*protocol.MSG_FILE_getByID_result, err error) {
	out := protocol.GET_MSG_FILE_getByObject()
	out.ObjectType = object
	out.ObjectID = ID
	var result *protocol.MSG_FILE_getByObject_result
	err = data.SendMsgWaitResultToDefault(out, &result)
	if err != nil {
		return
	}
	out.Put()
	return result.List, nil
}

func post_file_ajaxUploadTmp(data *TemplateData) (err error) {
	blockSize, _ := strconv.Atoi(data.ws.Query("blockSize"))
	index, _ := strconv.Atoi(data.ws.Query("index"))
	if len(data.ws.Body()) > blockSize {
		libraries.DebugLog("上传的文件%d大于blockSize%d", len(data.ws.Body()), blockSize)
		return errors.New("Error blockSize")
	}
	out := protocol.GET_MSG_FILE_uploadTmp()
	out.BlockSize = blockSize
	out.Data = append(out.Data, data.ws.Body()...)
	out.Index = index
	out.Name = file_getRealName(data.ws.Query("name"))
	err = data.SendMsgWaitResultToDefault(out, nil)
	out.Put()
	if err != nil {
		data.ws.WriteString(err.Error())
	} else {
		data.ws.WriteString("")
	}
	return nil
}

func file_getRealName(in string) string { //与前端js upload.js一起修改
	s := strings.Split(in, "_。。_")
	if len(s) == 2 {
		return s[1]
	}
	return in
}
func file_getTitleName(in string) string { //与前端js upload.js一起修改
	s := strings.Split(in, "_。。_")
	if len(s) == 2 {
		return s[0]
	}
	return in
}
func get_file_download(data *TemplateData) (err error) {
	out := protocol.GET_MSG_FILE_getByID()
	out.FileID, _ = strconv.ParseInt(data.ws.Query("fileID"), 10, 64)
	var result *protocol.MSG_FILE_getByID_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		data.ws.SetCode(404)
		data.ws.WriteString(err.Error())
		return nil
	}
	if err = checkFileAcl(data, result.ObjectType, result.ObjectID); err != nil {
		data.ws.SetCode(404)
		data.ws.WriteString(err.Error())
		return nil
	}
	data.ws.RangeDownload(&fileRangeDown{data: data, fileId: result.FileID, size: result.Size}, result.Size, result.Name)
	out.Put()
	result.Put()
	return nil
}

//封装一个与common通讯的远程下载
type fileRangeDown struct {
	data   *TemplateData
	fileId int64
	size   int64
	offset int64
	buf    *libraries.MsgBuffer
}

func (f *fileRangeDown) Read(b []byte) (int, error) {
	if f.buf == nil {
		out := protocol.GET_MSG_FILE_RangeDown()
		out.Start = f.offset
		out.End = f.offset + int64(len(b))
		if out.End > f.size {
			out.End = f.size
		}
		out.FileID = f.fileId
		var result *protocol.MSG_FILE_RangeDown_result
		if err := f.data.SendMsgWaitResultToDefault(out, &result); err != nil {
			return 0, err
		}
		copy(b, result.Byte)
		l := len(result.Byte)
		out.Put()
		result.Put()
		f.offset += int64(l)
		return l, nil
	}
	res := f.buf.Bytes()
	if f.offset > int64(len(res)) {
		return 0, nil
	}
	res = res[f.offset:]
	copy(b, res)
	l := len(res)
	if l > len(b) {
		l = len(b)
	}
	return l, nil

}

//只做了whence为0的情况
func (f *fileRangeDown) Seek(offset int64, whence int) (ret int64, err error) {
	f.offset = offset
	return offset, nil
}
func get_file_edit(data *TemplateData) (err error) {
	out := protocol.GET_MSG_FILE_getByID()
	out.FileID, _ = strconv.ParseInt(data.ws.Query("fileID"), 10, 64)
	out.NoData = true
	var result *protocol.MSG_FILE_getByID_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	if err = checkFileAcl(data, result.ObjectType, result.ObjectID); err != nil {
		data.ws.SetCode(404)
		data.ws.WriteString(err.Error())
		return nil
	}
	if strings.Contains(result.Name, ".") {
		result.Name = result.Name[:strings.LastIndex(result.Name, ".")]
	}
	data.Data["file"] = result
	templateOut("file.edit.html", data)
	out.Put()
	result.Put()
	return nil
}
func post_file_edit(data *TemplateData) (err error) {
	if len(data.ws.Post("fileName")) > 80 || len(data.ws.Post("fileName")) == 0 {
		data.ws.WriteString(js.Alert(fmt.Sprintf(data.Lang["error"]["length"].([]string)[1], data.Lang["file"]["title"].(string), 80, 1)))
		return
	}

	out := protocol.GET_MSG_FILE_edit()
	out.FileID, _ = strconv.ParseInt(data.ws.Query("fileID"), 10, 64)
	out.Name = data.ws.Post("fileName") + "." + data.ws.Post("extension")

	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		return
	}
	data.ws.WriteString(js.Reload("parent.parent"))
	return nil
}

func get_file_delete(data *TemplateData) (err error) {
	if data.ws.Query("confirm") != "yes" {
		data.ws.WriteString(js.Confirm(data.Lang["file"]["confirmDelete"].(string), createLink("file", "delete", "fileID="+data.ws.Query("fileID")+"&confirm=yes"), ""))
		return
	}
	getfile := protocol.GET_MSG_FILE_getByID()
	getfile.FileID, _ = strconv.ParseInt(data.ws.Query("fileID"), 10, 64)
	getfile.NoData = true
	var file *protocol.MSG_FILE_getByID_result
	if err = data.SendMsgWaitResultToDefault(getfile, &file); err != nil {
		data.ws.WriteString(js.Alert(err.Error()))
		return dataErrAlreadyOut
	}
	if err = checkFileAcl(data, file.ObjectType, file.ObjectID); err != nil {
		data.ws.SetCode(404)
		data.ws.WriteString(err.Error())
		return nil
	}
	out := protocol.GET_MSG_FILE_DeleteByID()
	out.FileID = getfile.FileID
	if err = data.SendMsgWaitResultToDefault(out, nil); err != nil {
		data.ws.WriteString(js.Alert(err.Error()))
		return dataErrAlreadyOut
	}
	data.ws.WriteString(js.Reload("parent"))
	data.Msg.ActionCreate(file.ObjectType, file.ObjectID, "deletedFile", "", file.Name, nil, 0)
	return
}
func checkFileAcl(data *TemplateData, ObjectType string, ObjectID int32) error {
	if !data.User.IsAdmin {
		switch ObjectType {
		case "task":
			task, err := task_getByID(data, ObjectID)
			if err != nil {
				return err
			}
			ObjectID = task.Project
			fallthrough
		case "project":
			if !data.User.AclProjects[ObjectID] {
				return errors.New(data.Lang["project"]["accessDenied"].(string))
			}
		case "product":
			if !data.User.AclProducts[ObjectID] {
				return errors.New(data.Lang["product"]["accessDenied"].(string))
			}
		}
	}
	return nil
}
func get_file_buildExportTPL(data *TemplateData) (err error) {
	module := data.ws.Query("module")
	out := protocol.GET_MSG_USER_getExportTemplate()
	out.Module = module
	out.Uid = data.User.Id
	var result *protocol.MSG_USER_getExportTemplate_result
	if err = data.SendMsgWaitResultToDefault(out, &result); err != nil {
		return
	}
	templatePairs := []protocol.HtmlKeyValueStr{protocol.HtmlKeyValueStr{"", data.Lang["file"]["defaultTPL"].(string)}}
	for _, template := range result.List {
		name := template.Title
		if template.Public {
			name = data.Lang["common"]["public"].(string) + name
		}
		templatePairs = append(templatePairs, protocol.HtmlKeyValueStr{strconv.Itoa(int(template.Id)), name})
	}

	data.Data["templates"] = result.List
	data.Data["templatePairs"] = templatePairs
	data.Data["templateID"] = data.ws.Query("templateID")
	templateOut("file.buildexporttpl.html", data)
	out.Put()
	result.Put()
	return
}

func file_export2xlsx(data *TemplateData, filename string, fields []protocol.HtmlKeyValueStr, values []map[string]string) (err error) {
	f := excelize.NewFile()
	row := 1
	fileIndex := 0
	for k, field := range fields {
		f.SetCellValue("Sheet1", file_getxlsAxis(k, row), field.Value)
		if field.Key == "files" {
			fileIndex = k
		}
		//给desc加宽
		if field.Key == "desc" {
			f.SetColWidth("Sheet1", file_getxlsAxis(k, -1), file_getxlsAxis(k, -1), 60)
		}
	}
	//边框
	border := []excelize.Border{excelize.Border{Type: "left", Style: 1, Color: "#000000"}, excelize.Border{Type: "right", Style: 1, Color: "#000000"}, excelize.Border{Type: "top", Style: 1, Color: "#000000"}, excelize.Border{Type: "bottom", Style: 1, Color: "#000000"}}
	//首行样式
	if style, err := f.NewStyle(&excelize.Style{
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#343399"}, Pattern: 1},
		Font:      &excelize.Font{Bold: true, Size: 9, Color: "#ffffff"},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Border:    border,
	}); err != nil {
		return err
	} else {
		f.SetCellStyle("Sheet1", file_getxlsAxis(0, row), file_getxlsAxis(len(fields), row), style)
	}
	//隔行样式
	style1, _ := f.NewStyle(&excelize.Style{
		Fill:   excelize.Fill{Type: "pattern", Color: []string{"#b2d7ea"}, Pattern: 1},
		Border: border,
	})
	style2, _ := f.NewStyle(&excelize.Style{
		Fill:   excelize.Fill{Type: "pattern", Color: []string{"#dee6fb"}, Pattern: 1},
		Border: border,
	})
	style1WrapText, _ := f.NewStyle(&excelize.Style{
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#b2d7ea"}, Pattern: 1},
		Border:    border,
		Alignment: &excelize.Alignment{WrapText: true},
	})
	style2WrapText, _ := f.NewStyle(&excelize.Style{
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#dee6fb"}, Pattern: 1},
		Border:    border,
		Alignment: &excelize.Alignment{WrapText: true},
	})
	maxFileNum := 0
	for _, v := range values {
		row++
		lastindex := 0
		for k, field := range fields {
			if field.Key != "files" {
				lastindex = k
				f.SetCellValue("Sheet1", file_getxlsAxis(k, row), v[field.Key])
			}
		}
		if row%2 == 1 {
			f.SetCellStyle("Sheet1", file_getxlsAxis(0, row), file_getxlsAxis(len(fields), row), style2)
		} else {
			f.SetCellStyle("Sheet1", file_getxlsAxis(0, row), file_getxlsAxis(len(fields), row), style1)
		}

		if v["files"] != "" {
			lastindex++
			_maxFileNum := 0
			for _, file := range strings.Split(v["files"], "||||") {
				m, err := libraries.Preg_match_result(`<a href='([^']+)' >(.+)<\/a>`, file, 1)
				if err == nil && len(m) == 1 {
					f.SetCellValue("Sheet1", file_getxlsAxis(lastindex, row), m[0][2])
					f.SetCellHyperLink("Sheet1", file_getxlsAxis(lastindex, row), m[0][1], "External")
				} else {
					f.SetCellValue("Sheet1", file_getxlsAxis(lastindex, row), file)
				}
				lastindex++
				_maxFileNum++

				if row%2 == 1 {
					f.SetCellStyle("Sheet1", file_getxlsAxis(lastindex, row), file_getxlsAxis(lastindex, row), style2WrapText)
				} else {
					f.SetCellStyle("Sheet1", file_getxlsAxis(lastindex, row), file_getxlsAxis(lastindex, row), style1WrapText)
				}

			}
			if _maxFileNum > maxFileNum {
				maxFileNum = _maxFileNum
			}
		}
	}
	//设置files的宽度
	if maxFileNum > 0 {
		f.SetColWidth("Sheet1", file_getxlsAxis(fileIndex, -1), file_getxlsAxis(fileIndex+maxFileNum, -1), 35)
	}
	buf := bufpool.Get().(*libraries.MsgBuffer)
	buf.Reset()
	if err = f.Write(buf); err != nil {
		return
	}
	data.ws.SetContentType("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	data.ws.RangeDownload(&fileRangeDown{data: data, size: int64(buf.Len()), buf: buf}, int64(buf.Len()), filename+".xlsx")
	buf.Reset()
	bufpool.Put(buf)
	return
}

//index 和 row 为-1时，表示行列
func file_getxlsAxis(index int, row int) string {
	axis := ""
	if index > -1 {
		if s1 := index / 26; s1 > 0 {
			axis = string(rune(64 + s1))
		}
		axis += string(rune(65 + index%26))
	}
	if row > 0 {
		axis += strconv.Itoa(row)
	}
	return axis
}
