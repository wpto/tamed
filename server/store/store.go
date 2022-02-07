package store

import (
	"github.com/pgeowng/tamed/config"
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/store/fslocal"
)

type MediaMetaRepo interface {
	GetMeta(string) (*model.MediaMeta, error)
}

type MediaContentRepo interface {
	GetContent(mediaID string, contentType string, width, height int) ([]byte, error)
}

type Store struct {
	MediaMeta MediaMetaRepo
	MediaContent MediaContentRepo
}

func New() (*Store, error) {
	cfg := config.Get()

	var store Store

	if cfg.LocalPath != "" {
		store.MediaMeta = fslocal.NewMediaMetaRepo(cfg.LocalPath)
		store.MediaContent = fslocal.NewMediaContentRepo(cfg.LocalPath)
	}

	return &store, nil
}
