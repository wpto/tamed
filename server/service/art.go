package service

import (
	"fmt"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/store"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type ArtSrv struct {
	store *store.Store
}

func NewArtSrv(store *store.Store) *ArtSrv {
	return &ArtSrv{store}
}

func (srv *ArtSrv) Create(userName string, media []model.Media) (*model.Art, error) {

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

func (srv *ArtSrv) Get(artID string) (result *model.Art, err error) {
	result, err = srv.store.Art.Get(artID)
	if err != nil {
		return nil, errors.Wrap(err, "srv.art.get")
	}
	if result == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("Art '%s' not found!", artID))
	}

	return
}
