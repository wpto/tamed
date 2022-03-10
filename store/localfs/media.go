package localfs

import (
	"fmt"
	"io"
)

type MediaStore struct {
	repo *MediaRepo
}

func NewMediaStore(repo *MediaRepo) *MediaStore {
	return &MediaStore{repo}
}

func (st *MediaStore) Upload(mediaID string, ext string, upload io.Reader) (string, error) {
	fmt.Println("mediastore.upload")

	filePath, err := st.repo.UploadReader(mediaID, ext, upload)
	return filePath, err
}

func (st *MediaStore) Delete(mediaID string) error {
	return st.repo.Delete(mediaID)
}
