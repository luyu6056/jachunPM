package handler

import (
	"encoding/base64"
	"errors"
	"fmt"
	common_image "image"
	"io"
	"jachunPM/image"
	"libraries"
	"strconv"
	"strings"
	"time"

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
}

func post_file_ajaxPasteImage(data *TemplateData) {
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
func get_file_tmpimg(data *TemplateData) {
	b, ok := file_getTempFile(data.ws.Query("fileID"))
	if !ok {
		data.ws.SetCode(404)
		data.ws.WriteString("img Not Found")
		return
	}
	ext := data.ws.Query("t")
	if ext == "webp" && !strings.Contains(data.ws.Header("Accept"), "image/webp") {
		ext = "jpg"
		var err error
		b, err = image.ConvertImgB(b, ext, 1, 1, 80)
		if err != nil {
			data.ws.SetCode(404)
			data.ws.WriteString(err.Error())
			return
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
func get_file_read(data *TemplateData) {
	out := protocol.GET_MSG_FILE_getByID()
	out.FileID, _ = strconv.ParseInt(data.ws.Query("fileID"), 10, 64)
	var result *protocol.MSG_FILE_getByID_result
	err := HostConn.SendMsgWaitResultToDefault(out, &result)
	if err != nil {
		data.ws.SetCode(404)
		data.ws.WriteString(err.Error())
		return
	}
	if result.Ext == "webp" && !strings.Contains(data.ws.Header("Accept"), "image/webp") {
		result.Ext = "jpg"
		var err error
		result.Data, err = image.ConvertImgB(result.Data, result.Ext, 1, 1, 80)
		if err != nil {
			data.ws.SetCode(404)
			data.ws.WriteString(err.Error())
			return
		}
	}
	buf := bufpool.Get().(*libraries.MsgBuffer)
	defer func() {
		buf.Reset()
		bufpool.Put(buf)
	}()
	buf.Write(result.Data)
	data.ws.SetContentType("image/" + result.Ext)
	data.ws.Write(buf)
	return
}
func get_file_buildform(data *TemplateData) {
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
	if data.Msg == nil {
		if _, uploaderr = data.GetMsg(); uploaderr != nil {
			return
		}
	}
	//检查移除失效的老文件
	m, _ := libraries.Preg_match_result(`<img aid="attachimg_0" src="\/file\/read\?fileID=(\d+)" border="0" alt=""  \/>`, desc, -1)
	for _, match := range m {
		check := protocol.GET_MSG_FILE_getByID()
		check.FileID, _ = strconv.ParseInt(match[1], 10, 64)
		check.NoData = true
		var result *protocol.MSG_FILE_getByID_result
		if uploaderr = HostConn.SendMsgWaitResultToDefault(check, &result); uploaderr != nil {
			if strings.Index(uploaderr.Error(), protocol.Err_FileNotFount.String()) == 0 {
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
func file_deleteFromIds(newimgids []int64) {
	deleteimg := protocol.GET_MSG_FILE_DeleteByID()
	for _, id := range newimgids {
		deleteimg.FileID = id
		HostConn.SendMsgToDefault(deleteimg)
	}
	deleteimg.Put()
}
func file_updateObject(fileIds []int64, typ string, id int32) {
	out := protocol.GET_MSG_FILE_updateByIDMap()
	for _, id := range fileIds {
		out.FileID = id
		out.Update = map[string]interface{}{
			"ObjectType": typ,
			"ObjectID":   id,
		}
		HostConn.SendMsgToDefault(out)
	}
	out.Put()
}
