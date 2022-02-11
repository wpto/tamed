package fslocal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type FileRepo struct {
	filePath string
}

func NewFileRepo(filePath string) *FileRepo {
	return &FileRepo{filePath}
}

func (rep *FileRepo) ReadDB() map[string]interface{} {
	if _, err := os.Stat(rep.filePath); os.IsNotExist(err) {
		return make(map[string]interface{}, 0)
	}

	file, err := ioutil.ReadFile(rep.filePath)
	if err != nil {
		fmt.Printf("filerepo.readdb.read %v", err)
		return make(map[string]interface{}, 0)
	}

	var db map[string]interface{}
	err = json.Unmarshal(file, &db)
	if err != nil {
		fmt.Printf("filerepo.readdb.unmarshal %v", err)
		return make(map[string]interface{}, 0)
	}

	return db
}

func (rep *FileRepo) WriteDB(db map[string]interface{}) error {
	text, err := json.Marshal(db)
	if err != nil {
		return errors.Wrap(err, "writedb.marshal")
	}

	err = ioutil.WriteFile(rep.filePath, text, 0644)
	if err != nil {
		return errors.Wrap(err, "writedb.write")
	}

	return nil
}

func (rep *FileRepo) Create(id string, encodedJSON []byte) error {
	var entry interface{}
	err := json.Unmarshal(encodedJSON, &entry)
	if err != nil {
		return errors.Wrap(err, "filerepo.create.unmarshal_entry")
	}

	db := rep.ReadDB()
	_, ok := db[id]
	if ok {
		return errors.Errorf("filerepo.create.db: id collision for %s", id)
	}

	db[id] = entry
	rep.WriteDB(db)

	return nil
}

func (rep *FileRepo) Get(id string) (interface{}, error) {
	db := rep.ReadDB()
	result, ok := db[id]
	if !ok {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("%s not found", id))
	}
	return result, nil
}
