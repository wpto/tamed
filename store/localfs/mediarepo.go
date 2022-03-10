package localfs

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func (repo *MediaRepo) UploadReader(mediaID string, ext string, upload io.Reader) (string, error) {
	dirPath := filepath.Join(repo.localPath, mediaID)
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return "", errors.Wrap(err, "mediarepo.upload")
	}

	fileName := mediaID
	if ext != "" {
		fileName = mediaID + "." + ext
	}

	filePath := filepath.Join(dirPath, fileName)

	err = WriteMedia(filePath, upload)
	if err != nil {
		return "", errors.Wrap(err, "mediarepo.upload")
	}

	return mediaID + "/" + fileName, nil
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

func (repo *MediaRepo) Delete(mediaID string) error {
	dirPath := filepath.Join(repo.localPath, mediaID)
	err := os.RemoveAll(dirPath)
	if err != nil {
		return errors.Wrap(err, "mediarepo.delete")
	}
	return nil
}
