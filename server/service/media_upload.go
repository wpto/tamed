package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

func (srv *MediaContentSrv) Upload(fileHeader *multipart.FileHeader) error {
	if fileHeader.Size == 0 {
		return errors.New("srv.mediacontent.upload: empty data")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return errors.Wrap(err, "srv.mediacontent.upload")
	}

	contentType, err := GuessContentType(file)
	fmt.Println("Content Type: " + contentType)
	file.Close()

	_, ok := types.AcceptedMime[contentType]
	if !ok {
		return errors.Errorf("srv.mediacontent.upload: bad upload type %v", contentType)
	}

	file, err = fileHeader.Open()
	if err != nil {
		return errors.Wrap(err, "srv.mediacontent.upload")
	}
	defer file.Close()
	err = srv.store.MediaContent.Save(contentType, file)
	if err != nil {
		return errors.Wrap(err, "srv.mediacontent.upload: save")
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
