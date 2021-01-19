package image

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	chai2010webp "github.com/chai2010/webp"
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

	switch imgtype {
	case "jpg", "jpeg":
		err = jpeg.Encode(dst_f, dstImage, &jpeg.Options{Quality: quality})
	case "webp":
		err = chai2010webp.Encode(dst_f, dstImage, &chai2010webp.Options{Quality: float32(quality)})
	case "png":
		err = png.Encode(dst_f, dstImage)
	case "bmp":
		err = bmp.Encode(dst_f, dstImage)
	}
	return err
}

func ConvertImgB(in []byte, imgtype string, width float32, height float32, quality int) ([]byte, error) {
	if width < 0 || height < 0 || quality <= 0 || quality > 100 {
		return nil, errors.New("参数错误")
	}
	buf := bytes.NewBuffer(in)
	imagetype, size, err := fastimage.DetectImageTypeFromReader(buf)
	if err != nil {
		return nil, errors.New("识别图片格式失败")
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
			return nil, errors.New("无法解码原始图片，错误代码:" + err.Error())
		}
	}

	if width < 10 && height < 10 && width > 0 && height > 0 {
		width = float32(size.Width) * width
		height = float32(size.Height) * height
	}
	dstImage := imaging.Resize(src_img, int(width), int(height), imaging.Lanczos)
	dst_f := new(bytes.Buffer)

	if err != nil {
		return nil, err
	}

	switch imgtype {
	case "jpg", "jpeg":
		err = jpeg.Encode(dst_f, dstImage, &jpeg.Options{Quality: quality})
	case "webp":
		err = chai2010webp.Encode(dst_f, dstImage, &chai2010webp.Options{Quality: float32(quality)})
	case "png":
		err = png.Encode(dst_f, dstImage)
	case "bmp":
		err = bmp.Encode(dst_f, dstImage)
	}
	return dst_f.Bytes(), err
}
func ConvertImgFromFastimage(in []byte, size *fastimage.ImageSize, imagetype fastimage.ImageType, imgtype string, width float32, height float32, quality int) ([]byte, error) {
	if width < 0 || height < 0 || quality <= 0 || quality > 100 {
		return nil, errors.New("参数错误")
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
			return nil, errors.New("无法解码原始图片，错误代码:" + err.Error())
		}
	}

	if width < 10 && height < 10 && width > 0 && height > 0 {
		width = float32(size.Width) * width
		height = float32(size.Height) * height
	}
	dstImage := imaging.Resize(src_img, int(width), int(height), imaging.Lanczos)
	dst_f := new(bytes.Buffer)

	if err != nil {
		return nil, err
	}

	switch imgtype {
	case "jpg", "jpeg":
		err = jpeg.Encode(dst_f, dstImage, &jpeg.Options{Quality: quality})
	case "webp":
		err = chai2010webp.Encode(dst_f, dstImage, &chai2010webp.Options{Quality: float32(quality)})
	case "png":
		err = png.Encode(dst_f, dstImage)
	case "bmp":
		err = bmp.Encode(dst_f, dstImage)
	}
	return dst_f.Bytes(), err
}
