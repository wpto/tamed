package store

// type MediaMetaRepo struct {
// 	localPath string
// }

// func NewMediaMetaRepo(localPath string) *MediaMetaRepo {
// 	return &MediaMetaRepo{localPath: localPath}
// }

// func TryFileMeta(dirPath string) (*FileMeta, error) {
// 	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
// 		return nil, errors.Wrap(types.ErrNotFound, "fslocal.tryfilemeta")
// 	}

// 	metaPath := filepath.Join(dirPath, "meta.json")
// 	f, err := ioutil.ReadFile(metaPath)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "fslocal.tryfilemeta filemeta read error")
// 	}

// 	fileMeta := &FileMeta{}
// 	err = json.Unmarshal(f, fileMeta)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "fslocal.tryfilemeta filemeta unmarshal error")
// 	}

// 	return fileMeta, nil
// }

// func (rep *MediaMetaRepo) GetMeta(mediaID string) (*model.MediaMeta, error) {

// 	dirPath := filepath.Join(rep.localPath, mediaID)

// 	fileMeta, err := TryFileMeta(dirPath)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "fslocal.getmeta")
// 	}

// 	return fileMeta.ToMediaMeta(), nil
// }
