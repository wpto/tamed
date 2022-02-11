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

// func (srv *UserSrv) Get(userName string) (result *model.User, err error) {
// 	result, err = srv.store.User.Get(userName)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "srv.user.get")
// 	}
// 	if meta == nil {
// 		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("User '%s' not found!", userName))
// 	}

// 	return
// }
