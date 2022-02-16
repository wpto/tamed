package usersrv

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/service/commonsrv"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

func (srv *UserSrv) Upload(fileHeader *multipart.FileHeader) (*model.Media, error) {
	if fileHeader.Size == 0 {
		return nil, errors.New("srv.user.upload: empty data")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, errors.Wrap(err, "srv.user.upload")
	}

	contentType, err := GuessContentType(file)
	fmt.Println("Content Type: " + contentType)
	file.Close()

	_, ok := types.AcceptedMime[contentType]
	if !ok {
		return nil, errors.Errorf("srv.user.upload: bad upload type %v", contentType)
	}

	id := commonsrv.UniqID()

	file, err = fileHeader.Open()
	if err != nil {
		return nil, errors.Wrap(err, "srv.user.upload")
	}
	defer file.Close()
	log.Println("uploading!!")
	err = srv.store.User.UploadMedia(id, contentType, file)
	if err != nil {
		return nil, errors.Wrap(err, "srv.user.upload: save")
	}

	obj := model.Media{
		ID:         id,
		CreateTime: commonsrv.TimeNow(),
		UserName:   "kifuku",
		// Meta:       meta,
		// Social: *model.NewMediaSocial(),
	}

	err = srv.store.User.CreateMedia(id, &obj)
	if err != nil {
		return nil, errors.Wrap(err, "srv.user.upload: db add")
	}

	return &obj, nil
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
