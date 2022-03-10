package service

import (
	"fmt"
	"mime/multipart"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/service/postsrv"
	"github.com/pgeowng/tamed/store"
)

type PostService interface {
	Get(postID string) (*model.Post, error)
	List(query *model.PostQuery) (*model.PostList, error)
	Modify(postID string, changes *model.PostChanges) error
	Delete(postID string) error
	Create(files []*multipart.FileHeader) ([]model.PostCreate, error)
}

type Manager struct {
	Post PostService
}

func NewManager(store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, fmt.Errorf("no store provided")
	} else {
		return &Manager{
			Post: postsrv.NewPostSrv(store),
		}, nil
	}
}
