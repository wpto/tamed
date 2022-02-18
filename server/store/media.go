package store

import (
	"fmt"
	"io"
)

type MediaStoreImpl struct {
	repo MediaRepo
}

func NewMediaStoreImpl(repo MediaRepo) *MediaStoreImpl {
	return &MediaStoreImpl{repo}
}

func (st *MediaStoreImpl) Upload(mediaID string, ext string, upload io.Reader) (string, error) {
	fmt.Println("mediastore.upload")

	filePath, err := st.repo.UploadReader(mediaID, ext, upload)
	return filePath, err
}
