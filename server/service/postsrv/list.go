package postsrv

import (
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/service/commonsrv"
	"github.com/pkg/errors"
)

func (p *PostSrv) List(query *model.PostQuery) (*commonsrv.ListResponse, error) {
	list, err := p.store.Post.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "postsrv.list")
	}

	return &commonsrv.ListResponse{
		Page:  1,
		Pages: 1,
		Total: len(list.Posts),
		Posts: list.Posts,
		Tags:  list.Tags,
	}, nil

}
