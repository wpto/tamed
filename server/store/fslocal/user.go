package fslocal

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type UserRepo struct {
	localPath     string
	mediaMetaRepo *FileRepo
}

func NewUserRepo(localPath string, mediaMetaRepo *FileRepo) *UserRepo {
	return &UserRepo{localPath, mediaMetaRepo}
}

func (rep *UserRepo) UploadMedia(mediaID string, contentType string, upload io.Reader) error {
	dirPath := filepath.Join(rep.localPath, mediaID)
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return errors.Wrap(err, "mediarepo.upload")
	}

	origFileName := mediaID
	ext, err := types.GetExt(contentType)
	if err == nil {
		origFileName = origFileName + "." + ext
	}

	origPath := filepath.Join(dirPath, origFileName)
	err = WriteMedia(origPath, upload)
	if err != nil {
		return errors.Wrap(err, "mediarepo.upload")
	}

	var mediaInfo MediaInfo
	mediaType := types.MimeToMediaType[contentType]
	if mediaType == "vid" {
		err = VidInfo(origPath, &mediaInfo)
	} else if mediaType == "pic" {
		err = PicInfo(origPath, &mediaInfo)
	} else {
		err = errors.Errorf("unknown mediatype for %s", contentType)
	}

	if err != nil {
		return errors.Wrap(err, "mediarepo.upload")
	}

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

func WriteMedia(localPath string, upload io.Reader) error {
	local, err := os.Create(localPath)
	if err != nil {
		return errors.Wrap(err, "writemedia")
	}
	defer local.Close()

	buf := make([]byte, 1024)
	for {
		n, err := upload.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := local.Write(buf[:n]); err != nil {
			return err
		}
	}

	return nil
}

func (rep *UserRepo) CreateMedia(mediaID string, obj *model.Media) error {

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
