package store

import (
	"encoding/json"
	"io"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type UserStoreImpl struct {
	artRepo       FileRepo
	mediaMetaRepo FileRepo
	mediaRepo     MediaRepo
}

func NewUserStoreImpl(artRepo FileRepo, mediaMetaRepo FileRepo, mediaRepo MediaRepo) *UserStoreImpl {
	return &UserStoreImpl{artRepo, mediaMetaRepo, mediaRepo}
}

func (rep UserStoreImpl) UploadMedia(mediaID string, contentType string, upload io.Reader) error {

	ext, err := types.GetExt(contentType)
	if err != nil {
		ext = ""
	}
	err = rep.mediaRepo.UploadReader(mediaID, ext, upload)
	if err != nil {
		return errors.Wrap(err, "mediarepo.upload")
	}

	// var mediaInfo MediaInfo
	// mediaType := types.MimeToMediaType[contentType]
	// if mediaType == "vid" {
	// 	err = VidInfo(origPath, &mediaInfo)
	// } else if mediaType == "pic" {
	// 	err = PicInfo(origPath, &mediaInfo)
	// } else {
	// 	err = errors.Errorf("unknown mediatype for %s", contentType)
	// }

	// if err != nil {
	// 	return errors.Wrap(err, "mediarepo.upload")
	// }

	// meta := &FileMeta{
	// 	Mime:      contentType,
	// 	MediaType: types.MimeToMediaType[contentType],
	// 	Filename:  origFileName,
	// 	ID:        mediaID,
	// 	Width:     mediaInfo.Width,
	// 	Height:    mediaInfo.Height,
	// 	Duration:  mediaInfo.Duration,
	// }

	// metaData, err := json.Marshal(meta)
	// if err != nil {
	// 	return errors.Wrap(err, "mediarepo.upload")
	// }

	// metaPath := filepath.Join(dirPath, "meta.json")
	// err = ioutil.WriteFile(metaPath, metaData, 0644)
	// if err != nil {
	// 	return errors.Wrap(err, "mediarepo.upload")
	// }

	return nil
}

func (rep UserStoreImpl) CreateMedia(mediaID string, obj *model.Media) error {

	enc, err := json.Marshal(*obj)
	if err != nil {
		return errors.Wrap(err, "user.create_media")
	}

	err = rep.mediaMetaRepo.Create(mediaID, enc)
	if err != nil {
		return errors.Wrap(err, "user.create_media")
	}

	return nil
}

func (rep UserStoreImpl) CreateArt(artID string, obj *model.Art) error {

	enc, err := json.Marshal(*obj)
	if err != nil {
		return errors.Wrap(err, "user.create_art")
	}

	err = rep.artRepo.Create(artID, enc)
	if err != nil {
		return errors.Wrap(err, "user.create_art")
	}

	return nil
}
