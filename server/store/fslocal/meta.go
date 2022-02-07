package fslocal

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type MediaMetaRepo struct {
	localPath string
}

func NewMediaMetaRepo(localPath string) *MediaMetaRepo {
	return &MediaMetaRepo{localPath: localPath}
}

func TryFileMeta(dirPath string) (*FileMeta, error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return nil, errors.Wrap(types.ErrNotFound, "fslocal.tryfilemeta")
	}

	metaPath := filepath.Join(dirPath, "meta.json")
	f, err := ioutil.ReadFile(metaPath)
	if err != nil {
		return nil, errors.Wrap(err, "fslocal.tryfilemeta filemeta read error")
	}

	fileMeta := &FileMeta{}
	err = json.Unmarshal(f, fileMeta)
	if err != nil {
		return nil, errors.Wrap(err, "fslocal.tryfilemeta filemeta unmarshal error")
	}

	return fileMeta, nil
}

func (rep *MediaMetaRepo) GetMeta(mediaID string) (*model.MediaMeta, error) {

	dirPath := filepath.Join(rep.localPath, mediaID)

	fileMeta, err := TryFileMeta(dirPath)
	if err != nil {
		return nil, errors.Wrap(err, "fslocal.getmeta")
	}

	return fileMeta.ToMediaMeta(), nil
}

// func DetectType(formatName string) string {
// 	mediaType := "unknown"
// 	switch formatName {
// 	case :
// 		/* code */
// 	default:
// 		/* code */
// 		return
// 	}
// }

// func GetFileContentType(out *os.File) (string, error) {
// 	f, err := os.Open(opath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	contentType, err := GetFileContentType(f)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Content Type: " + contentType)
// 	buffer := make([]byte, 512)
// 	_, err := out.Read(buffer)
// 	if err != nil {
// 		return "", err
// 	}

// 	contentType := http.DetectContentType(buffer)
// 	return contentType, nil
// }

type FileMeta struct {
	Mime      string `json:"mime"`
	MediaType string `json:"mediaType"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Filename  string `json:"filename"`
	ID        string `json:"id"`
}

func (mf *FileMeta) ToMediaMeta() *model.MediaMeta {
	return &model.MediaMeta{
		ID:     mf.ID,
		Type:   mf.MediaType,
		Width:  mf.Width,
		Height: mf.Height,
	}
}
