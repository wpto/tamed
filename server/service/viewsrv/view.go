package viewsrv

import (
	"fmt"

	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/store"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type ViewSrv struct {
	store *store.Store
}

func NewViewSrv(store *store.Store) *ViewSrv {
	return &ViewSrv{store}
}

func (srv *ViewSrv) ViewArt(artID string) (result *model.Art, err error) {
	result, err = srv.store.View.GetArt(artID)
	if err != nil {
		return nil, errors.Wrap(err, "srv.view.art")
	}

	if result == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("srv.view.art: Art '%s' not found!", artID))
	}

	return
}

func (srv *ViewSrv) ViewUser(userName string) (result *model.User, err error) {
	result, err = srv.store.View.GetUser(userName)
	if err != nil {
		return nil, errors.Wrap(err, "srv.view.user")
	}
	if result == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("User '%s' not found!", userName))
	}

	return
}

func (srv *ViewSrv) Search() (*model.SearchResponse, error) {

	user, err := srv.store.View.GetUser("kifuku")
	if err != nil {
		return nil, errors.Wrap(err, "srv.view.search.user")
	}

	arts, err := srv.store.View.SearchArt()
	if err != nil {
		return nil, errors.Wrap(err, "srv.view.search.media")
	}

	return &model.SearchResponse{
		Page:  1,
		Pages: 1,
		Total: len(arts),
		Arts:  arts,
		Users: []model.User{*user},
	}, nil
}
