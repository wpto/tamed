package postsrv

import (
	"fmt"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/store"
)

type PostSrv struct {
	store *store.Store
}

func NewPostSrv(store *store.Store) *PostSrv {
	return &PostSrv{store}
}

func (p *PostSrv) Get(postID string) (*model.Post, error) {
	fmt.Printf("postsrv.get")
	return nil, nil
}

func (p *PostSrv) List(query *model.PostQuery) (*model.PostList, error) {
	fmt.Printf("postsrv.list")
	return nil, nil
}

func (p *PostSrv) Modify(postID string, changes *model.PostChanges) error {
	fmt.Printf("postsrv.modify")
	return nil
}

func (p *PostSrv) Delete(postID string) error {
	fmt.Printf("postsrv.delete")
	return nil
}
