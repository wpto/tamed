package service

import (
	"fmt"

	"github.com/pgeowng/tamed/store"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type MediaContentSrv struct {
	store *store.Store
}

func NewMediaContentSrv(store *store.Store) *MediaContentSrv {
	return &MediaContentSrv{
		store: store,
	}
}

func (srv *MediaContentSrv) Download(mediaID string, contentType string, width int, height int) (data []byte, err error) {

	meta, err := srv.store.MediaMeta.GetMeta(mediaID)
	if err != nil {
		return nil, errors.Wrap(err, "srv.mediacontent.dl")
	}

	resultMediaType, err := types.GetMediaType(contentType)
	if err != nil {
		return nil, errors.Wrap(err, "srv.mediacontent.dl")
	}

	localMediaType, err := types.GetMediaType(meta.Mime)
	if err != nil {
		return nil, errors.Wrap(err, "srv.mediacontent.dl")
	}

	if resultMediaType != localMediaType {
		return nil, errors.Errorf("srv.mediacontent.dl: mediatype mismatch: %s(%s) to %s(%s)", meta.Mime, localMediaType, contentType, resultMediaType)
	}

	if resultMediaType == "vid" {
		data, err = srv.store.MediaVid.GetContent(&types.GetVidOpts{
			mediaID, contentType, width, false,
		})
		if err != nil {
			return nil, errors.Wrap(err, "srv.mediacontent.vid")
		}
		return data, nil
	}

	if resultMediaType == "pic" {
		data, err = srv.store.MediaPic.GetContent(&types.GetPicOpts{mediaID, contentType, width})
		if err != nil {
			return nil, errors.Wrap(err, "srv.mediacontent.pic")
		}
		return data, nil
	}

	return nil, errors.Wrap(types.ErrNotImplemented, fmt.Sprintf("srv.mediacontent.dl(%s to %s)", meta.Mime, contentType))
}
