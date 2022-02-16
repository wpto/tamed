package fslocal

import (
	"encoding/json"

	"github.com/pgeowng/tamed/model"
	"github.com/pkg/errors"
)

type ViewRepo struct {
	artRepo   *FileRepo
	userRepo  *FileRepo
	mediaRepo *FileRepo
}

func NewViewRepo(artRepo *FileRepo, userRepo *FileRepo, mediaRepo *FileRepo) *ViewRepo {
	return &ViewRepo{artRepo, userRepo, mediaRepo}
}

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

func (rep *ViewRepo) GetMedia(mediaID string) (*model.Media, error) {
	data, err := rep.mediaRepo.Get(mediaID)
	if err != nil {
		return nil, errors.Wrap(err, "viewrepo")
	}

	var result model.Media
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.New("viewrepo: typecast error")
	}

	return &result, nil
}

func (rep *ViewRepo) Search() {}
