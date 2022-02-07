package store

import (
	"github.com/pgeowng/tamed/config"
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/store/fslocal"
)

type MediaMetaRepo interface {
	GetMeta(string) (*model.MediaMeta, error)
}

type Store struct {
	MediaMeta MediaMetaRepo
}

func New() (*Store, error) {
	cfg := config.Get()

	var store Store

	if cfg.LocalPath != "" {
		store.MediaMeta = fslocal.NewMediaMetaRepo(cfg.LocalPath)
	}

	return &store, nil
}
