package mediasrv

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

func (srv *MediaContentSrv) Upload(fileHeader *multipart.FileHeader) (*model.Media, error) {
	if fileHeader.Size == 0 {
		return errors.New("srv.media.upload: empty data")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return errors.Wrap(err, "srv.media.upload")
	}

	contentType, err := GuessContentType(file)
	fmt.Println("Content Type: " + contentType)
	file.Close()

	_, ok := types.AcceptedMime[contentType]
	if !ok {
		return errors.Errorf("srv.media.upload: bad upload type %v", contentType)
	}

	id := UniqID()

	file, err = fileHeader.Open()
	if err != nil {
		return errors.Wrap(err, "srv.media.upload")
	}
	defer file.Close()
	meta, err = srv.store.MediaContent.Upload(id, contentType, file)
	if err != nil {
		return errors.Wrap(err, "srv.media.upload: save")
	}

	obj := model.Media{
		ID:         id,
		CreateTime: TimeNow(),
		UserName:   "kifuku",
		Meta:       meta,
		Social:     model.NewMediaSocial(),
	}

	return nil
}

func GuessContentType(file io.Reader) (string, error) {
	buf := make([]byte, 512)
	_, err := file.Read(buf)
	if err != nil {
		return "", errors.Wrap(err, "guessmime")
	}
	contentType := http.DetectContentType(buf)
	return contentType, nil
}
