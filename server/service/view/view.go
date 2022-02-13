package service

import "github.com/pgeowng/tamed/store"

type ViewSrv struct {
	store *store.Store
}

func NewViewSrv(store *store.Store) *ViewSrv {
	return &ViewSrv{store}
}
