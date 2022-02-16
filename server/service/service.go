package service

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/service/usersrv"
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
}

type ViewService interface {
	ViewArt(artID string) (*model.Art, error)
	// ViewRecent()
	ViewUser(userName string) (*model.User, error)
	// ViewTag()
	// Search()
}

type UserService interface {
	Upload(fileHeader *multipart.FileHeader) (*model.Media, error)
	CreateArt(userName string, media []model.Media) (*model.Art, error)
	// DeleteArt()
	// AddTag()
	// LookFavorites()
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
			View: viewsrv.NewViewSrv(store),
			User: usersrv.NewUserSrv(store),
		}, nil
	}
}
