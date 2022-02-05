package service

import "github.com/pgeowng/tamed/store"

type MediaContentServiceImpl struct {
	store *store.Store
}

func (srv *MediaContentServiceImpl) Upload() error {
	return nil
}

func (srv *MediaContentServiceImpl) Download() error {
	return nil
}

func NewMediaContentServiceImpl(store *store.Store) *MediaContentServiceImpl {
	return &MediaContentServiceImpl{
		store: store,
	}
}
