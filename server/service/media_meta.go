package service

import (
	"fmt"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/store"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type MediaMetaSrv struct {
	store *store.Store
}

func NewMediaMetaSrv(store *store.Store) *MediaMetaSrv {
	return &MediaMetaSrv{
		store: store,
	}
}

func (srv *MediaMetaSrv) Get(mediaID string) (meta *model.MediaMeta, err error) {
	meta, err = srv.store.MediaMeta.GetMeta(mediaID)
	if err != nil {
		return nil, errors.Wrap(err, "srv.mediameta.get")
	}
	if meta == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("Media '%s' not found!", mediaID))
	}

	return
}
