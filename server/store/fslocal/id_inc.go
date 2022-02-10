package fslocal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
)

type CountObj struct {
	Count int64 `json:"count"`
}

func NewId(localPath string) (string, error) {
	countPath := filepath.Join(localPath, "count.json")
	countData, err := ioutil.ReadFile(countPath)
	if err != nil {
		return "", errors.Wrap(err, "fslocal.newid")
	}

	countJSON := &CountObj{}
	err = json.Unmarshal(countData, countJSON)
	if err != nil {
		return "", errors.Wrap(err, "fslocal.newid")
	}

	mediaID := fmt.Sprintf("%d", countJSON.Count)

	newCountObj := &CountObj{countJSON.Count + 1}
	newCountData, err := json.Marshal(newCountObj)
	if err != nil {
		return "", errors.Wrap(err, "fslocal.newid")
	}

	err = ioutil.WriteFile(countPath, newCountData, 0644)
	if err != nil {
		return "", errors.Wrap(err, "fslocal.newid")
	}

	return mediaID, nil
}
