package store

import (
	"io"
	"path/filepath"

	"github.com/pgeowng/tamed/config"
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/store/localfs"
)

type FileRepo interface {
	Create(id string, encodedJSON []byte) error
	Get(id string) ([]byte, error)
	All() ([]byte, error)
}

type MediaRepo interface {
	Alloc(mediaID string, ext string) (localPath string, err error)
	UploadReader(mediaID, ext string, upload io.Reader) error
}

type UserStore interface {
	UploadMedia(mediaID string, contentType string, upload io.Reader) error
	CreateMedia(mediaID string, obj *model.Media) error
}

type ViewStore interface {
	// Create(entry *model.Art) error
	GetArt(artID string) (*model.Art, error)
	GetUser(userName string) (*model.User, error)
	SearchArt() ([]model.Art, error)
}

type Store struct {
	View ViewStore
	User UserStore
}

func New() (*Store, error) {
	cfg := config.Get()

	var store Store

	if cfg.LocalPath != "" {

		artRepo := localfs.NewFileRepo(filepath.Join(cfg.LocalPath, "artdb.json"))
		userRepo := localfs.NewFileRepo(filepath.Join(cfg.LocalPath, "userdb.json"))
		mediaRepo := localfs.NewFileRepo(filepath.Join(cfg.LocalPath, "mediadb.json"))
		mediaFileRepo := localfs.NewMediaRepo(filepath.Join(cfg.LocalPath, "mediacontent"))

		store.View = NewViewStoreImpl(
			artRepo,
			userRepo,
			mediaRepo,
		)

		store.User = NewUserStoreImpl(
			mediaRepo,
			mediaFileRepo,
		)
	}

	return &store, nil
}
