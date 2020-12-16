package libraries

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/rubenfonseca/fastimage"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
)

/**
 *图片转换
 *width与height小于10时，为按比例缩放,当其中一个值为0时，进行定值缩放，比如width=640,height=0，按640宽度定比缩放
 *quality最大质量100
 **/
func ConvertImg(src string, dst string, imgtype string, width float32, height float32, quality int) error {
	if width < 0 || height < 0 || quality <= 0 || quality > 100 {
		return errors.New("参数错误")
	}
	src_f, err := os.Open(src)
	defer src_f.Close()
	if err != nil {
		return err
	}
	imagetype, size, err := fastimage.DetectImageTypeFromReader(src_f)
	if err != nil {
		return errors.New("识别图片格式失败，路径" + src)
	}
	src_f.Seek(0, os.SEEK_SET)
	var src_img image.Image
	switch imagetype {
	case fastimage.JPEG:
		src_img, err = jpeg.Decode(src_f)
	case fastimage.PNG:
		src_img, err = png.Decode(src_f)
	case fastimage.GIF:
		src_img, err = gif.Decode(src_f)
	case fastimage.BMP:
		src_img, err = bmp.Decode(src_f)
	case fastimage.TIFF:
		src_img, err = tiff.Decode(src_f)
	default:
		src_img, err = webp.Decode(src_f) //尝试用webp解码
	}

	if err != nil {
		src_img, err = imaging.Decode(src_f) //尝试解码
		if err != nil {
			return errors.New("无法解码原始图片" + src + "，错误代码:" + err.Error())
		}
	}

	if width < 10 && height < 10 && width > 0 && height > 0 {
		width = float32(size.Width) * width
		height = float32(size.Height) * height
	}
	dstImage := imaging.Resize(src_img, int(width), int(height), imaging.Lanczos)
	dst_f, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0664)
	defer dst_f.Close()
	if err != nil {
		return err
	}
	var webp_tmp string
	switch imgtype {
	case "jpg", "jpeg":
		err = jpeg.Encode(dst_f, dstImage, &jpeg.Options{Quality: quality})
	case "webp":

		webp_tmp = dst[:strings.LastIndex(dst, ".")]
		tmp_f, err1 := os.OpenFile(webp_tmp+"_tmp.png", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0)
		if err1 != nil {
			tmp_f.Close()
			return err1
		}

		err1 = png.Encode(tmp_f, dstImage)
		if err1 != nil {
			tmp_f.Close()
			return err1
		}

		tmp_f.Close()
		dst_f.Close()
		err1 = exec.Command("cwebp", "-q", fmt.Sprintf("%d", quality), webp_tmp+"_tmp.png", "-o", webp_tmp+".webp").Run()
		if err1 != nil {
			return errors.New("转换webp文件失败" + err1.Error())
		}
		err = exec.Command("rm", webp_tmp+"_tmp.png").Run()
	case "png":
		err = png.Encode(dst_f, dstImage)
	case "bmp":
		err = bmp.Encode(dst_f, dstImage)
	}
	return err
}

var webp_no int

func ConvertImgB(in []byte, imgtype string, width float32, height float32, quality int) (error, []byte) {
	if width < 0 || height < 0 || quality <= 0 || quality > 100 {
		return errors.New("参数错误"), nil
	}
	buf := bytes.NewBuffer(in)
	imagetype, size, err := fastimage.DetectImageTypeFromReader(buf)
	if err != nil {
		return errors.New("识别图片格式失败"), nil
	}
	buf.Truncate(0)
	buf.Write(in)
	var src_img image.Image
	switch imagetype {
	case fastimage.JPEG:
		src_img, err = jpeg.Decode(buf)
	case fastimage.PNG:
		src_img, err = png.Decode(buf)
	case fastimage.GIF:
		src_img, err = gif.Decode(buf)
	case fastimage.BMP:
		src_img, err = bmp.Decode(buf)
	case fastimage.TIFF:
		src_img, err = tiff.Decode(buf)
	default:
		src_img, err = webp.Decode(buf) //尝试用webp解码
	}
	buf.Truncate(0)
	buf.Write(in)
	if err != nil {
		src_img, err = imaging.Decode(buf) //尝试解码
		if err != nil {
			return errors.New("无法解码原始图片，错误代码:" + err.Error()), nil
		}
	}

	if width < 10 && height < 10 && width > 0 && height > 0 {
		width = float32(size.Width) * width
		height = float32(size.Height) * height
	}
	dstImage := imaging.Resize(src_img, int(width), int(height), imaging.Lanczos)
	dst_f := new(bytes.Buffer)

	if err != nil {
		return err, nil
	}

	switch imgtype {
	case "jpg", "jpeg":
		err = jpeg.Encode(dst_f, dstImage, &jpeg.Options{Quality: quality})
	case "webp":
		if runtime.GOOS == "windows" {
			return errors.New("windows模式不支持webp"), nil
		}
		webp_no++
		dst := "./temp/webp" + strconv.Itoa(webp_no)
		tmp_f, err1 := os.OpenFile(dst+"_tmp.png", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0)
		if err1 != nil {
			tmp_f.Close()
			return err1, nil
		}

		err1 = png.Encode(tmp_f, dstImage)
		if err1 != nil {
			tmp_f.Close()
			return err1, nil
		}

		tmp_f.Close()
		err1 = exec.Command("cwebp", "-q", fmt.Sprintf("%d", quality), dst+"_tmp.png", "-o", dst+".webp").Run()
		if err1 != nil {
			return errors.New("转换webp文件失败" + err1.Error()), nil
		}
		b, _ := ioutil.ReadFile(dst + ".webp")
		err = exec.Command("rm", dst+"_tmp.png").Run()
		err = exec.Command("rm", dst+".webp").Run()
		return nil, b
	case "png":
		err = png.Encode(dst_f, dstImage)
	case "bmp":
		err = bmp.Encode(dst_f, dstImage)
	}
	return err, dst_f.Bytes()
}
func ConvertImgFromFastimage(in []byte, size *fastimage.ImageSize, imagetype fastimage.ImageType, imgtype string, width float32, height float32, quality int) (error, []byte) {
	if width < 0 || height < 0 || quality <= 0 || quality > 100 {
		return errors.New("参数错误"), nil
	}
	buf := bytes.NewBuffer(in)
	var err error
	var src_img image.Image
	switch imagetype {
	case fastimage.JPEG:
		src_img, err = jpeg.Decode(buf)
	case fastimage.PNG:
		src_img, err = png.Decode(buf)
	case fastimage.GIF:
		src_img, err = gif.Decode(buf)
	case fastimage.BMP:
		src_img, err = bmp.Decode(buf)
	case fastimage.TIFF:
		src_img, err = tiff.Decode(buf)
	default:
		src_img, err = webp.Decode(buf) //尝试用webp解码
	}
	buf.Truncate(0)
	buf.Write(in)
	if err != nil {
		src_img, err = imaging.Decode(buf) //尝试解码
		if err != nil {
			return errors.New("无法解码原始图片，错误代码:" + err.Error()), nil
		}
	}

	if width < 10 && height < 10 && width > 0 && height > 0 {
		width = float32(size.Width) * width
		height = float32(size.Height) * height
	}
	dstImage := imaging.Resize(src_img, int(width), int(height), imaging.Lanczos)
	dst_f := new(bytes.Buffer)

	if err != nil {
		return err, nil
	}

	switch imgtype {
	case "jpg", "jpeg":
		err = jpeg.Encode(dst_f, dstImage, &jpeg.Options{Quality: quality})
	case "webp":
		if runtime.GOOS == "windows" {
			return errors.New("windows模式不支持webp"), nil
		}
		webp_no++
		dst := "./temp/webp" + strconv.Itoa(webp_no)
		tmp_f, err1 := os.OpenFile(dst+"_tmp.png", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0)
		if err1 != nil {
			tmp_f.Close()
			return err1, nil
		}

		err1 = png.Encode(tmp_f, dstImage)
		if err1 != nil {
			tmp_f.Close()
			return err1, nil
		}

		tmp_f.Close()
		err1 = exec.Command("cwebp", "-q", fmt.Sprintf("%d", quality), dst+"_tmp.png", "-o", dst+".webp").Run()
		if err1 != nil {
			return errors.New("转换webp文件失败" + err1.Error()), nil
		}
		b, _ := ioutil.ReadFile(dst + ".webp")
		err = exec.Command("rm", dst+"_tmp.png").Run()
		err = exec.Command("rm", dst+".webp").Run()
		return nil, b
	case "png":
		err = png.Encode(dst_f, dstImage)
	case "bmp":
		err = bmp.Encode(dst_f, dstImage)
	}
	return err, dst_f.Bytes()
}
func init() {
	os.Mkdir("./temp", os.ModePerm)
}
