package fslocal

import (
	"github.com/h2non/bimg"
	"github.com/pgeowng/tamed/model"
)

type FileMeta struct {
	Mime      string `json:"mime"`
	MediaType string `json:"mediaType"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Filename  string `json:"filename"`
	ID        string `json:"id"`
}

func (mf *FileMeta) ToMediaMeta() *model.MediaMeta {
	return &model.MediaMeta{
		ID:     mf.ID,
		Type:   mf.MediaType,
		Mime:   mf.Mime,
		Width:  mf.Width,
		Height: mf.Height,
	}
}

var MimeToBimg = map[string]bimg.ImageType{
	"image/jpeg": bimg.JPEG,
	"image/gif":  bimg.GIF,
	"image/png":  bimg.PNG,
	"image/webp": bimg.WEBP,
}
