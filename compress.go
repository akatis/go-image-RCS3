package go_image_RCS3

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

func ImgCompress(width, height uint, quality int, imgBase64 string) (string, error) {

	index := strings.Index(imgBase64, ";base64,")
	if index < 0 {
		return "", errors.New("Invalid image")
	}
	imgExt := imgBase64[11:index]

	unbasedImage, err := base64.StdEncoding.DecodeString(imgBase64[index+8:])
	if err != nil {
		return "", err
	}

	var newImage string

	switch imgExt {
	case "png":
		img, err := png.Decode(bytes.NewReader(unbasedImage))
		if err != nil {
			panic("bad png")
		}

		resizedImage := resize.Resize(width, height, img, resize.Lanczos3)
		buf := new(bytes.Buffer)
		if err = jpeg.Encode(buf, resizedImage, &jpeg.Options{Quality: quality}); err != nil {
			return "", err
		}

		//compressedImage, err := jpeg.Decode(buf)
		compressedImageBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

		//newImage = "data:image/jpeg;base64," + compressedImageBase64
		newImage = compressedImageBase64

	case "jpeg", "jpg":
		img, _, err := image.Decode(strings.NewReader(string(unbasedImage)))
		if err != nil {
			return "", err
		}
		resizedImage := resize.Resize(width, height, img, resize.Lanczos3)
		buf := new(bytes.Buffer)
		if err = jpeg.Encode(buf, resizedImage, &jpeg.Options{Quality: quality}); err != nil {
			return "", err
		}

		//compressedImage, err := jpeg.Decode(buf)
		compressedImageBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

		//newImage = "data:image/jpeg;base64," + compressedImageBase64
		newImage = compressedImageBase64
	}
	return newImage, nil
}
