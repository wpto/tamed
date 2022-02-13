package service

import (
	"github.com/pgeowng/tamed/model"
	"github.com/pkg/errors"
)

func (srv *UserSrv) Create(userName string, media []model.Media) (*model.Art, error) {

	obj := model.Art{
		ID:         UniqID(),
		CreateTime: TimeNow(),
		UserName:   userName,
		Media:      media,
	}

	err := srv.store.Art.Create(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "srv.art.create.store")
	}

	return &obj, nil
}
