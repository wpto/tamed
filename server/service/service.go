package service

import (
	"fmt"
	"mime/multipart"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/service/postsrv"
	"github.com/pgeowng/tamed/store"
)

// type MediaFileRes interface {
// 	ContentType() string
// 	ContentLength() int64
// 	Reader() io.Reader
// }

// type MediaService interface {
// 	Serve(mediaID string, qualityTag string) (*MediaFileRes, error)
// }

// type ViewService interface {
// 	ViewArt(artID string) (*model.Art, error)
// 	// ViewRecent()
// 	ViewUser(userName string) (*model.User, error)
// 	// ViewTag()
// 	Search() (*model.SearchResponse, error)
// 	ViewMedia(mediaID string) (*model.File, error)
// }

// type UserService interface {
// 	Upload(fileHeader *multipart.FileHeader) (*model.Media, error)
// 	CreateArt(userName string, media []model.Media) (*model.Art, error)
// 	// DeleteArt()
// 	// AddTag()
// 	// LookFavorites()
// }

type PostService interface {
	Get(postID string) (*model.Post, error)
	List(query *model.PostQuery) (*model.PostList, error)
	Modify(postID string, changes *model.PostChanges) error
	Delete(postID string) error
	Create(files []*multipart.FileHeader) ([]model.Post, error)
}

type Manager struct {
	Post PostService
}

func NewManager(store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, fmt.Errorf("no store provided")
	} else {
		return &Manager{
			Post: postsrv.NewPostSrv(store),
			// Media: mediasrv.NewMediaSrv(store),
			// View: viewsrv.NewViewSrv(store),
			// User: usersrv.NewUserSrv(store),
		}, nil
	}
}
