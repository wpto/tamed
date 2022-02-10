package store

import (
	"github.com/pgeowng/tamed/config"
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/store/fslocal"
	"github.com/pgeowng/tamed/types"
)

type MediaMetaRepo interface {
	GetMeta(string) (*model.MediaMeta, error)
}

type MediaPicRepo interface {
	GetContent(*types.GetPicOpts) ([]byte, error)
}

type MediaVidRepo interface {
	GetContent(*types.GetVidOpts) ([]byte, error)
}

type MediaContentRepo interface {
	Save(contentType string, data []byte) error
}

type Store struct {
	MediaMeta    MediaMetaRepo
	MediaPic     MediaPicRepo
	MediaVid     MediaVidRepo
	MediaContent MediaContentRepo
}

func New() (*Store, error) {
	cfg := config.Get()

	var store Store

	if cfg.LocalPath != "" {
		store.MediaMeta = fslocal.NewMediaMetaRepo(cfg.LocalPath)
		store.MediaPic = fslocal.NewMediaPicRepo(cfg.LocalPath)
		store.MediaVid = fslocal.NewMediaVidRepo(cfg.LocalPath)
		store.MediaContent = fslocal.NewMediaContentRepo(cfg.LocalPath)
	}

	return &store, nil
}
