package store

import (
	"io"

	"github.com/pgeowng/tamed/config"
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/store/localfs"
)

type PostList struct {
	Posts []model.Post
	Tags  []model.Tag
}

type FileRepo interface {
	Create(id string, encodedJSON []byte) error
	Write(id string, encodedJSON []byte) error
	Get(id string) ([]byte, error)
	All() ([]byte, error)
}

type MediaRepo interface {
	UploadReader(mediaID, ext string, upload io.Reader) (filePath string, err error)
}

type MediaStore interface {
	Upload(mediaID string, ext string, upload io.Reader) (filePath string, err error)
}

type PostStore interface {
	Get(postID string) (*model.Post, error)
	Query(query *model.PostQuery) (*PostList, error)
	Create(postID string, post *model.Post) error
	Modify(postID string, changes *model.PostChanges) error
	Delete(postID string) error
}

type Store struct {
	// View ViewStore
	// User UserStore
	Media MediaStore
	Post  PostStore
}

func New() (*Store, error) {
	cfg := config.Get()

	var store Store

	if cfg.LocalPath != "" {
		mediaRepo := localfs.NewMediaRepo(config.Get().MediaPath)
		postRepo := localfs.NewFileRepo(config.Get().PostDBPath)

		store.Media = NewMediaStoreImpl(mediaRepo)
		store.Post = NewPostStoreImpl(postRepo)
	}

	return &store, nil
}
