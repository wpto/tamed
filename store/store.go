package store

import (
	"io"

	"github.com/pgeowng/tamed/config"
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/store/localfs"
	"github.com/pgeowng/tamed/store/pg"
)

type FileRepo interface {
	Create(id string, encodedJSON []byte) error
	Delete(id string) error
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
	Query(query *model.PostQuery) (*model.PostList, error)
	Create(postID string, post *model.Post) error
	Modify(postID string, changes *model.PostChanges) error
	Delete(postID string) error
}

type Store struct {
	Media MediaStore
	Post  PostStore
}

func New() (*Store, error) {
	cfg := config.Get()

	var store Store

	if cfg.FsLocalPath != "" {
		mediaRepo := localfs.NewMediaRepo(config.Get().FsMediaPath)
		store.Media = localfs.NewMediaStore(mediaRepo)

		if cfg.PgUrl != "" {
			db, err := pg.Dial()
			if err != nil {
				panic(err)
			}
			store.Post = db
		} else {
			postRepo := localfs.NewFileRepo(config.Get().FsPostDBPath)

			store.Post = localfs.NewPostStore(postRepo)
		}
	}

	return &store, nil
}
