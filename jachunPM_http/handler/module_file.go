package handler

import (
	"encoding/base64"
	"errors"
	"image"
	"io"
	"libraries"
	"strconv"
	"strings"

	"github.com/luyu6056/cache"
	"github.com/luyu6056/gnet"
	"github.com/rubenfonseca/fastimage"
	"golang.org/x/image/bmp"
	"golang.org/x/image/webp"
)

func init() {
	httpHandlerMap["POST"]["/file/ajaxPasteImage"] = post_file_ajaxPasteImage
	httpHandlerMap["GET"]["/file/tmpimg"] = get_file_tmpimg
}

func post_file_ajaxPasteImage(data *TemplateData) (action gnet.Action) {
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
		b, err = libraries.ConvertImgFromFastimage(b, size, imagetype, "webp", 1, 1, 80)
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
func get_file_tmpimg(data *TemplateData) (action gnet.Action) {
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
		b, err = libraries.ConvertImgB(b, ext, 1, 1, 80)
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
		var imgconfig image.Config
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
