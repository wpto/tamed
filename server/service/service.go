package service

import (
	"fmt"
	"io"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/service/viewsrv"
	"github.com/pgeowng/tamed/store"
)

type MediaFileRes interface {
	ContentType() string
	ContentLength() int64
	Reader() io.Reader
}

type MediaService interface {
	Serve(mediaID string, qualityTag string) (*MediaFileRes, error)
	Upload(contentType string, upload io.Reader) (*model.Media, error)
}

type ViewService interface {
	ViewArt(artID string) (*model.Art, error)
	// ViewRecent()
	ViewUser(userName string) (*model.User, error)
	// ViewTag()
	// Search()
}

type UserService interface {
	CreateArt()
	DeleteArt()
	AddTag()
	LookFavorites()
}

type Manager struct {
	Media MediaService
	User  UserService
	View  ViewService
}

func NewManager(store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, fmt.Errorf("no store provided")
	} else {
		return &Manager{
			// Media: mediasrv.NewMediaSrv(store),
			// User:  usersrv.NewUserSrv(store),
			View: viewsrv.NewViewSrv(store),
		}, nil
	}
}
