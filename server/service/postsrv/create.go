package postsrv

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/service/commonsrv"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

func (srv *PostSrv) Create(files []*multipart.FileHeader) ([]model.Post, error) {
	fmt.Printf("postsrv.create")

	for idx, header := range files {
		if header.Size == 0 {
			return nil, errors.Errorf("postsrv.create: empty data inside %d file", idx)
		}

	}

	result := make([]model.Post, 0)
	errs := make([]error, 0)
	for _, header := range files {
		fileName := header.Filename
		file, err := header.Open()
		if err != nil {
			errs = append(errs, errors.Wrap(err, "postsrv.create("+fileName+")"))
			continue
		}

		contentType, err := GuessContentType(file)
		fmt.Println("Content Type: " + contentType)
		file.Close()

		_, ok := types.AcceptedMime[contentType]
		if !ok {
			errs = append(errs, errors.Errorf("postsrv.create(%s): bad upload type %v", fileName, contentType))
			continue
		}

		ext := strings.TrimPrefix(filepath.Ext(fileName), ".")

		if len(ext) == 0 {
			myExt := types.GetExt(contentType)
			if len(myExt) == 0 {
				errs = append(errs, errors.Errorf("postsrv.create(%s): cant detect content type", fileName))
				continue
			}
			ext = myExt
		}

		id := commonsrv.UniqID()

		file, err = header.Open()
		if err != nil {
			errs = append(errs, errors.Wrap(err, "postsrv.create("+fileName+")"))
			continue
		}
		filePath, err := srv.store.Media.Upload(id, ext, file)
		file.Close()
		if err != nil {
			errs = append(errs, errors.Wrap(err, "postsrv.create("+fileName+")"))
			continue
		}

		obj := model.Post{
			PostID:     id,
			CreateTime: commonsrv.TimeNow(),
			Tags:       []model.Tag{},
			Link:       filePath,
		}

		err = srv.store.Post.Create(id, &obj)
		if err != nil {
			errs = append(errs, errors.Wrap(err, "postsrv.create("+fileName+")"))
			continue
		}
	}

	return result, nil
}

func GuessContentType(file io.ReadCloser) (string, error) {
	defer file.Close()

	buf := make([]byte, 512)
	_, err := file.Read(buf)
	if err != nil {
		return "", errors.Wrap(err, "guessmime")
	}
	contentType := http.DetectContentType(buf)
	return contentType, nil
}
