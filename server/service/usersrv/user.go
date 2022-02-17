package usersrv

import (
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/service/commonsrv"
	"github.com/pgeowng/tamed/store"
)

type UserSrv struct {
	store *store.Store
}

func NewUserSrv(store *store.Store) *UserSrv {
	return &UserSrv{store}
}

func (srv *UserSrv) CreateArt(userName string, media []model.Media) (*model.Art, error) {

	obj := model.Art{
		ID:         commonsrv.UniqID(),
		CreateTime: commonsrv.TimeNow(),
		UserName:   userName,
		Media:      media,
	}

	// err := srv.store.Art.Create(&obj)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "srv.art.create.store")
	// }

	return &obj, nil
}
