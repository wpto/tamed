package fslocal

import (
	"encoding/json"

	"github.com/pgeowng/tamed/model"
	"github.com/pkg/errors"
)

type ArtRepo struct {
	fileRepo *FileRepo
}

func NewArtRepo(filePath string) *ArtRepo {
	return &ArtRepo{NewFileRepo(filePath)}
}

func (rep *ArtRepo) Create(entry *model.Art) error {
	id := entry.ID
	body, err := json.Marshal(entry)
	if err != nil {
		return errors.Wrap(err, "artrepo.create.marshal")
	}

	err = rep.fileRepo.Create(id, body)
	if err != nil {
		return errors.Wrap(err, "artrepo.create")
	}

	return nil
}

func (rep *ArtRepo) Get(artID string) (*model.Art, error) {
	any, err := rep.fileRepo.Get(artID)
	if err != nil {
		return nil, errors.Wrap(err, "artrepo")
	}

	result, ok := any.(model.Art)
	if !ok {
		return nil, errors.New("artrepo: typecast error")
	}

	return &result, nil
}
