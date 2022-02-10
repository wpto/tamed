package fslocal

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type MediaContentRepo struct {
	localPath string
}

func NewMediaContentRepo(localPath string) *MediaContentRepo {
	return &MediaContentRepo{localPath}
}

func (rep *MediaContentRepo) Save(contentType string, upload io.Reader) error {
	mediaID, err := NewId(rep.localPath)
	if err != nil {
		return errors.Wrap(err, "fslocal.mediacontent.save")
	}

	dirPath := filepath.Join(rep.localPath, mediaID)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return errors.Wrap(err, "fslocal.mediacontent.save")
	}

	// origFileName := fmt.Sprintf("%s.%s", mediaID, types.MimeToExt[contentType])
	origFileName := mediaID
	origPath := filepath.Join(dirPath, origFileName)
	err = WriteMedia(origPath, upload)
	if err != nil {
		return errors.Wrap(err, "fslocal.mediacontent.save")
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
		return errors.Wrap(err, "fslocal.mediacontent.save")
	}

	meta := &FileMeta{
		Mime:      contentType,
		MediaType: types.MimeToMediaType[contentType],
		Filename:  origFileName,
		ID:        mediaID,
		Width:     mediaInfo.Width,
		Height:    mediaInfo.Height,
		Duration:  mediaInfo.Duration,
	}

	metaData, err := json.Marshal(meta)
	if err != nil {
		return errors.Wrap(err, "fslocal.mediacontent.save")
	}

	metaPath := filepath.Join(dirPath, "meta.json")
	err = ioutil.WriteFile(metaPath, metaData, 0644)
	if err != nil {
		return errors.Wrap(err, "fslocal.mediacontent.save")
	}

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
