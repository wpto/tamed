package service

import "github.com/pgeowng/tamed/store"

type MediaContentSrv struct {
	store *store.Store
}

func (srv *MediaContentSrv) Upload() error {
	return nil
}

func (srv *MediaContentSrv) Download() error {
	return nil
}

func NewMediaContentSrv(store *store.Store) *	MediaContentSrv {
	return &MediaContentSrv{
		store: store,
	}
}
