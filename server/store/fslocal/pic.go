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

type MediaPicRepo struct {
	localPath string
}

func NewMediaPicRepo(localPath string) *MediaPicRepo {
	return &MediaPicRepo{localPath: localPath}
}

func (rep *MediaPicRepo) GetContent(opts *types.GetPicOpts) (data []byte, err error) {
	// mediaID string, contentType string, width, height int
	dirPath := filepath.Join(rep.localPath, opts.MediaID)

	meta, err := TryFileMeta(dirPath)
	if err != nil {
		return nil, errors.Wrap(err, "fslocal.getpic")
	}

	// check cached resource

	origPath := filepath.Join(dirPath, meta.Filename)
	data, err = loadPic(origPath)
	if err != nil {
		return nil, errors.Wrap(err, "fslocal.getpic")
	}

	if opts.ContentType != meta.Mime {
		data, err = formatPic(data, opts.ContentType)
		if err != nil {
			return nil, errors.Wrap(err, "fslocal.getpic")
		}
	}

	if opts.Width < meta.Width {
		data, err = resizePic(data, opts.Width)
		if err != nil {
			return nil, errors.Wrap(err, "fslocal.getpic")
		}
	}

	// caching result

	return data, nil
}

func loadPic(origPath string) ([]byte, error) {
	buf, err := ioutil.ReadFile(origPath)
	if err != nil {
		return nil, errors.Wrap(err, "loadpic")
	}
	return buf, nil
}

func resizePic(buf []byte, width int) ([]byte, error) {
	img := bimg.NewImage(buf)
	size, err := img.Size()
	if err != nil {
		return nil, errors.Wrap(err, "resizepic")
	}

	coef := float64(width) / float64(size.Width)

	newWidth := int(math.Floor(float64(size.Width) * coef))
	newHeight := int(math.Floor(float64(size.Height) * coef))

	buf, err = img.ResizeAndCrop(newWidth, newHeight)
	if err != nil {
		return nil, errors.Wrap(err, "resizepic")
	}

	return buf, nil
}

func formatPic(buf []byte, mime string) ([]byte, error) {
	img := bimg.NewImage(buf)
	buf, err := img.Convert(MimeToBimg[mime])
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("formatpic(%s to %s)", img.Type(), mime))
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
