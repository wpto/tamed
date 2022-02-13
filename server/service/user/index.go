package service

import (
	"github.com/pgeowng/tamed/store"
)

type UserSrv struct {
	store *store.Store
}

func NewUserSrv(store *store.Store) *UserSrv {
	return &UserSrv{store}
}
