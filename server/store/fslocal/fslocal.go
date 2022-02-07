package fslocal

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pgeowng/tamed/model"
	"github.com/pkg/errors"
)

type FileMetaFSLocalRepo struct {
	localPath string
}

func NewMediaMetaRepo(localPath string) *FileMetaFSLocalRepo {
	return &FileMetaFSLocalRepo{localPath: localPath}
}

func (rep *FileMetaFSLocalRepo) GetMeta(mediaID string) (meta *model.MediaMeta, err error) {

	dpath := filepath.Join(rep.localPath, mediaID)

	if _, err = os.Stat(dpath); os.IsNotExist(err) {
		return nil, nil
	}

	mpath := filepath.Join(dpath, "meta.json")

	metaFile, err := ioutil.ReadFile(mpath)
	if err != nil {
		return nil, errors.Wrap(err, "fs.meta file read error")
	}

	metaJSON := MetaFile{}
	err = json.Unmarshal(metaFile, &metaJSON)
	if err != nil {
		return nil, errors.Wrap(err, "fs.meta file unmarshal error")
	}

	return metaJSON.ToMediaMeta(), nil
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

type MetaFile struct {
	Mime      string `json:"mime"`
	MediaType string `json:"mediaType"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Filename  string `json:"filename"`
	ID        string `json:"id"`
}

func (mf *MetaFile) ToMediaMeta() *model.MediaMeta {
	return &model.MediaMeta{
		ID:     mf.ID,
		Type:   mf.MediaType,
		Width:  mf.Width,
		Height: mf.Height,
	}
}
