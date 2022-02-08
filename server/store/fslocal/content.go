package fslocal

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"

	"github.com/h2non/bimg"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type MediaContentRepo struct {
	localPath string
}

func NewMediaContentRepo(localPath string) *MediaContentRepo {
	return &MediaContentRepo{localPath: localPath}
}

func (rep *MediaContentRepo) GetContent(mediaID string, contentType string, width, height int) (data []byte, err error) {
	dirPath := filepath.Join(rep.localPath, mediaID)

	meta, err := TryFileMeta(dirPath)
	if err != nil {
		return nil, errors.Wrap(err, "fslocal.getcontent")
	}

	if width >= meta.Width {
		width = meta.Width
	}

	if height >= meta.Height {
		height = meta.Height
	}

	// check cached resource

	origPath := filepath.Join(dirPath, meta.Filename)

	if meta.Mime != "video/mp4" && contentType != "video/mp4" {
		targetExt, err := types.GetExt(contentType)
		if err != nil {
			return nil, errors.Wrap(err, "fslocal.getcontent")
		}
		buf, err := ConvImgToImg(origPath, "", targetExt, width, height)
		if err != nil {
			return nil, errors.Wrap(err, "fslocal.getcontent")
		}

		return buf, nil
	}

	return nil, errors.Wrap(types.ErrNotImplemented, fmt.Sprintf("%s to %s not supported", meta.Mime, contentType))
}

var ExtToBimg = map[string]bimg.ImageType{
	"jpg":  bimg.JPEG,
	"gif":  bimg.GIF,
	"png":  bimg.PNG,
	"webp": bimg.WEBP,
}

func ConvImgToImg(origPath string, cachePath string, targetExt string, width, height int) ([]byte, error) {
	buf, err := ioutil.ReadFile(origPath)
	if err != nil {
		return nil, errors.Wrap(err, "fslocal.content.convimgtoimg")
	}

	img := bimg.NewImage(buf)

	if img.Type() != targetExt {
		fmt.Printf("converting from %s to %s", img.Type(), targetExt)
		buf, err = img.Convert(ExtToBimg[targetExt])
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("fslocal.content.convimgtoimg: from %s to %s", img.Type(), targetExt))
		}

		img = bimg.NewImage(buf)
	}

	size, err := img.Size()
	if err != nil {
		return nil, errors.Wrap(err, "fslocal.content.convimgtoimg")
	}

	if width < size.Width || height < size.Height {
		coef := float64(width) / float64(size.Width)
		coefH := float64(height) / float64(size.Height)

		if coefH < coef {
			coef = coefH
		}

		newWidth := int(math.Floor(float64(size.Width) * coef))
		newHeight := int(math.Floor(float64(size.Height) * coef))

		buf, err = img.ResizeAndCrop(newWidth, newHeight)
		if err != nil {
			return nil, errors.Wrap(err, "fslocal.content.convimgtoimg")
		}

		img = bimg.NewImage(buf)
	}

	if len(cachePath) > 0 {
		err := ioutil.WriteFile(cachePath, buf, 0644)
		if err != nil {
			return nil, errors.Wrap(err, "fslocal.content.convimgtoimg")
		}
	}

	return buf, nil
}

// 120x120 png
// 680x680 jpg small
// small is 680x680
// medium is 1200x1200
// large is 2048x2048
// ... is 4096x4096
// iphone - small png
// ipad - medium png
