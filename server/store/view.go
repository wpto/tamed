package store

import (
	"encoding/json"

	"github.com/pgeowng/tamed/model"
	"github.com/pkg/errors"
)

type ViewStoreImpl struct {
	artRepo   FileRepo
	userRepo  FileRepo
	mediaRepo FileRepo
}

func NewViewStoreImpl(artRepo FileRepo, userRepo FileRepo, mediaRepo FileRepo) *ViewStoreImpl {
	return &ViewStoreImpl{artRepo, userRepo, mediaRepo}
}

func (rep ViewStoreImpl) GetArt(artID string) (*model.Art, error) {
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

func (rep ViewStoreImpl) GetUser(userName string) (*model.User, error) {
	data, err := rep.userRepo.Get(userName)
	if err != nil {
		return nil, errors.Wrap(err, "viewrepo.user")
	}

	var result model.User
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.New("viewrepo.user.typecast")
	}

	return &result, nil
}

func (rep ViewStoreImpl) GetMedia(mediaID string) (*model.Media, error) {
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

func (rep ViewStoreImpl) SearchArt() ([]model.Art, error) {
	data, err := rep.artRepo.All()
	if err != nil {
		return nil, errors.Wrap(err, "viewrepo.search.art")
	}

	var result []model.Art
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.New("viewrepo.search.art.typecast")
	}

	return result, nil
}

func (rep ViewStoreImpl) SearchMedia() ([]model.Media, error) {
	data, err := rep.mediaRepo.All()
	if err != nil {
		return nil, errors.Wrap(err, "viewrepo.search")
	}

	var result []model.Media
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.New("viewrepo.search.typecast")
	}

	return result, nil
}
