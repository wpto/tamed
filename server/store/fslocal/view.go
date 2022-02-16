package fslocal

import (
	"encoding/json"

	"github.com/pgeowng/tamed/model"
	"github.com/pkg/errors"
)

type ViewRepo struct {
	artRepo  *FileRepo
	userRepo *FileRepo
}

func NewViewRepo(artpath string, userpath string) *ViewRepo {
	return &ViewRepo{NewFileRepo(artpath), NewFileRepo(userpath)}
}

// func (rep *ViewRepo) Create(entry *model.Art) error {
// 	id := entry.ID
// 	body, err := json.Marshal(entry)
// 	if err != nil {
// 		return errors.Wrap(err, "viewrepo.create.marshal")
// 	}

// 	err = rep.fileRepo.Create(id, body)
// 	if err != nil {
// 		return errors.Wrap(err, "viewrepo.create")
// 	}

// 	return nil
// }

func (rep *ViewRepo) GetArt(artID string) (*model.Art, error) {
	data, err := rep.artRepo.Get(artID)
	if err != nil {
		return nil, errors.Wrap(err, "viewrepo")
	}

	var result model.Art
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.New("viewrepo: typecast error")
	}

	return &result, nil
}

func (rep *ViewRepo) GetUser(userName string) (*model.User, error) {
	data, err := rep.userRepo.Get(userName)
	if err != nil {
		return nil, errors.Wrap(err, "viewrepo")
	}

	var result model.User
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.New("viewrepo: typecast error")
	}

	return &result, nil
}

func (rep *ViewRepo) Search() {}
