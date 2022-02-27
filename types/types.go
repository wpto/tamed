package types

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrBadRequest     = errors.New("bad request")
	ErrNotImplemented = errors.New("not implemented")
	ErrNotAllowed     = errors.New("not allowed")
)

var AcceptedMime = map[string]bool{
	"image/jpeg": true,
	"image/gif":  true,
	"image/png":  true,
	"image/webp": true,
	"video/mp4":  true,
	"video/webm": true,
}

var ExtToMime = map[string]string{
	"jpg":  "image/jpeg",
	"gif":  "image/gif",
	"png":  "image/png",
	"webp": "image/webp",
	"mp4":  "video/mp4",
}

func GetMime(ext string) (string, error) {
	mime, ok := ExtToMime[ext]
	if !ok {
		return "", errors.Wrap(ErrBadRequest, fmt.Sprintf("bad ext(%s)", ext))
	}
	return mime, nil
}

var MimeToExt = map[string]string{
	"image/jpeg": "jpg",
	"image/gif":  "gif",
	"image/png":  "png",
	"image/webp": "webp",
	"video/mp4":  "mp4",
	"video/webm": "mkv",
}

func GetExt(mime string) string {
	ext, ok := MimeToExt[mime]
	if !ok {
		fmt.Println(errors.Wrap(ErrBadRequest, fmt.Sprintf("bad mime(%s)", mime)))
		return ""
	}
	return ext
}

var MimeToMediaType = map[string]string{
	"image/jpeg": "pic",
	"image/gif":  "pic",
	"image/png":  "pic",
	"image/webp": "pic",
	"video/mp4":  "vid",
	"video/webm": "vid",
}

func GetMediaType(mime string) (string, error) {
	ext, ok := MimeToMediaType[mime]
	if !ok {
		return "", errors.Wrap(ErrBadRequest, fmt.Sprintf("mediatype: bad mime(%s)", mime))
	}
	return ext, nil
}

// probably bad place??
type GetVidOpts struct {
	MediaID     string
	ContentType string
	Width       int
	Audio       bool
}

type GetPicOpts struct {
	MediaID     string
	ContentType string
	Width       int
}
