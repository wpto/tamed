package localfs

import (
	"github.com/pgeowng/tamed/model"
)

type FileMeta struct {
	Mime      string `json:"mime"`
	MediaType string `json:"mediaType"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Filename  string `json:"filename"`
	ID        string `json:"id"`
	Duration  int    `json:"duration"`
}

func (mf *FileMeta) ToMediaMeta() *model.MediaMeta {
	return &model.MediaMeta{
		Type:   mf.MediaType,
		Mime:   mf.Mime,
		Width:  mf.Width,
		Height: mf.Height,
	}
}

type FileRepo struct {
	filePath string
}

func NewFileRepo(filePath string) FileRepo {
	return FileRepo{filePath}
}

type MediaRepo struct {
	localPath string
}

func NewMediaRepo(localPath string) MediaRepo {
	return MediaRepo{localPath}
}
