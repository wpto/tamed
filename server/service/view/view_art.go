package service

import (
	"fmt"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

func (srv *ArtSrv) Get(artID string) (result *model.Art, err error) {
	result, err = srv.store.Art.Get(artID)
	if err != nil {
		return nil, errors.Wrap(err, "srv.art.get")
	}
	if result == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("Art '%s' not found!", artID))
	}

// func (srv *MediaMetaSrv) Get(mediaID string) (meta *model.MediaMeta, err error) {
// 	meta, err = srv.store.MediaMeta.GetMeta(mediaID)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "srv.mediameta.get")
// 	}
// 	if meta == nil {
// 		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("Media '%s' not found!", mediaID))
// 	}

// 	return
}
	return
}
