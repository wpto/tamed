package localfs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

func (repo FileRepo) ReadDB() map[string]interface{} {
	if _, err := os.Stat(repo.filePath); os.IsNotExist(err) {
		return make(map[string]interface{}, 0)
	}

	file, err := ioutil.ReadFile(repo.filePath)
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

func (repo FileRepo) WriteDB(db map[string]interface{}) error {
	text, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return errors.Wrap(err, "writedb.marshal")
	}

	err = ioutil.WriteFile(repo.filePath, text, 0644)
	if err != nil {
		return errors.Wrap(err, "writedb.write")
	}

	return nil
}

func (repo FileRepo) Create(id string, encodedJSON []byte) error {
	var entry interface{}
	err := json.Unmarshal(encodedJSON, &entry)
	if err != nil {
		return errors.Wrap(err, "filerepo.create.unmarshal_entry")
	}

	db := repo.ReadDB()
	_, ok := db[id]
	if ok {
		return errors.Errorf("filerepo.create.db: id collision for %s", id)
	}

	db[id] = entry
	repo.WriteDB(db)

	return nil
}

func (repo FileRepo) Get(id string) ([]byte, error) {
	db := repo.ReadDB()
	entry, ok := db[id]
	if !ok {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("%s not found", id))
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return nil, errors.Errorf("marshal error %v", entry)
	}
	return data, nil
}

func (repo FileRepo) All() ([]byte, error) {
	db := repo.ReadDB()
	result := make([]interface{}, 0)
	for _, entry := range db {
		result = append(result, entry)
	}

	data, err := json.Marshal(result)
	if err != nil {
		return nil, errors.Errorf("filerepo.all.marshal %v", err)
	}

	return data, nil
}
