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

type ViewRepo interface {
	// Create(entry *model.Art) error
	GetArt(artID string) (*model.Art, error)
	GetUser(userName string) (*model.User, error)
}

type Store struct {
	View         ViewRepo
	MediaMeta    MediaMetaRepo
	MediaPic     MediaPicRepo
	MediaVid     MediaVidRepo
	MediaContent MediaContentRepo
}

func New() (*Store, error) {
	cfg := config.Get()

	var store Store

	if cfg.LocalPath != "" {
		store.View = fslocal.NewViewRepo(
			filepath.Join(cfg.LocalPath, "artdb.json"),
			filepath.Join(cfg.LocalPath, "userdb.json"),
		)
		store.MediaMeta = fslocal.NewMediaMetaRepo(cfg.LocalPath)
		store.MediaPic = fslocal.NewMediaPicRepo(cfg.LocalPath)
		store.MediaVid = fslocal.NewMediaVidRepo(cfg.LocalPath)
		store.MediaContent = fslocal.NewMediaContentRepo(cfg.LocalPath)
	}

	return &store, nil
}
