package service

import (
	"fmt"
	"net/http"

	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

func (srv *MediaContentSrv) Upload(data []byte) error {
	if len(data) == 0 {
		return errors.New("srv.mediacontent.upload: empty data")
	}

	contentType, err := GuessContentType(data)
	fmt.Println("Content Type: " + contentType)

	_, ok := types.AcceptedMime[contentType]
	if !ok {
		return errors.Errorf("srv.mediacontent.upload: bad upload type %v", contentType)
	}

	// mediaID, err := srv.store.MediaMeta.Create(contentType)
	// if err != nil {
	// 	return errors.Wrap(err, "srv.mediacontent.upload")
	// }

	err = srv.store.MediaContent.Save(contentType, data)
	if err != nil {
		return errors.Wrap(err, "srv.mediacontent.upload")
	}

	return nil
}

func GuessContentType(data []byte) (string, error) {
	contentType := http.DetectContentType(data[:512])
	return contentType, nil
}
