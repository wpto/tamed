package types

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrNotFound       = errors.New("resource not found")
	ErrBadRequest     = errors.New("bad request")
	ErrNotImplemented = errors.New("not implemented")
)

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
}

func GetExt(mime string) (string, error) {
	ext, ok := MimeToExt[mime]
	if !ok {
		return "", errors.Wrap(ErrBadRequest, fmt.Sprintf("bad mime(%s)", mime))
	}
	return ext, nil
}
