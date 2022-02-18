package localfs

import (
	"fmt"
)

type FileRepo struct {
	filePath string
}

func NewFileRepo(filePath string) *FileRepo {
	fmt.Println(filePath)
	return &FileRepo{filePath}
}

type MediaRepo struct {
	localPath string
}

func NewMediaRepo(localPath string) *MediaRepo {
	fmt.Println(localPath)
	return &MediaRepo{localPath}
}
