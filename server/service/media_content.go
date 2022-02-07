package service

import (
	"github.com/pgeowng/tamed/store"
	"github.com/pkg/errors"
)

type MediaContentSrv struct {
	store *store.Store
}

func (srv *MediaContentSrv) Download(mediaID string, contentType string, width int, height int) (data []byte, err error) {
	data, err = srv.store.MediaContent.GetContent(mediaID, contentType, width, height)

	if err != nil {
		return nil, errors.Wrap(err, "srv.mediacontent.download")
	}

	return
}

func NewMediaContentSrv(store *store.Store) *MediaContentSrv {
	return &MediaContentSrv{
		store: store,
	}
}
