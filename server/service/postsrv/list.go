package postsrv

import (
	"fmt"

	"github.com/pgeowng/tamed/model"
)

func (p *PostSrv) List(query *model.PostQuery) (*model.PostList, error) {
	fmt.Printf("postsrv.list")
	return nil, nil

	// user, err := srv.store.View.GetUser("kifuku")
	// if err != nil {
	// 	return nil, errors.Wrap(err, "srv.view.search.user")
	// }

	// arts, err := srv.store.View.SearchArt()
	// if err != nil {
	// 	return nil, errors.Wrap(err, "srv.view.search.media")
	// }

	// return &model.SearchResponse{
	// 	Page:  1,
	// 	Pages: 1,
	// 	Total: len(arts),
	// 	Arts:  arts,
	// 	Users: []model.User{*user},
	// }, nil

}
