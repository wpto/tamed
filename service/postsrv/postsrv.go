package postsrv

import (
	"github.com/pgeowng/tamed/store"
)

type PostSrv struct {
	store *store.Store
}

func NewPostSrv(store *store.Store) *PostSrv {
	return &PostSrv{store}
}
