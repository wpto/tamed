package mediasrv

import "github.com/pgeowng/tamed/store"

type MediaSrv struct {
	store *store.Store
}

func NewMediaSrv(store *store.Store) *MediaSrv {
	return &MediaSrv{
		store: store,
	}
}
