package store

import (
	"io"
	"path/filepath"

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
	Upload(id string, contentType string, upload io.Reader) error
}

type ArtRepo interface {
	Create(entry *model.Art) error
	Get(artID string) (*model.Art, error)
}

type UserRepo interface {
	Get(artID string) (*model.User, error)
}

type Store struct {
	Art          ArtRepo
	MediaMeta    MediaMetaRepo
	MediaPic     MediaPicRepo
	MediaVid     MediaVidRepo
	MediaContent MediaContentRepo
}

func New() (*Store, error) {
	cfg := config.Get()

	var store Store

	if cfg.LocalPath != "" {
		store.Art = fslocal.NewArtRepo(filepath.Join(cfg.LocalPath, "artdb.json"))
		store.MediaMeta = fslocal.NewMediaMetaRepo(cfg.LocalPath)
		store.MediaPic = fslocal.NewMediaPicRepo(cfg.LocalPath)
		store.MediaVid = fslocal.NewMediaVidRepo(cfg.LocalPath)
		store.MediaContent = fslocal.NewMediaContentRepo(cfg.LocalPath)
	}

	return &store, nil
}
