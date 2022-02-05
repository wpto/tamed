package service

import (
	"fmt"

	"github.com/pgeowng/tamed/store"
)

type MediaContentService interface {
	Upload() error
	Download() error
}

type Manager struct {
	MediaContent MediaContentService
}

func NewManager(store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, fmt.Errorf("no store provided")
	} else {
		return &Manager{
			MediaContent: NewMediaContentService(store),
		}, nil
	}
}
