package postsrv

import (
	"github.com/pgeowng/tamed/model"
	"github.com/pkg/errors"
)

func (p *PostSrv) List(query *model.PostQuery) (*model.PostList, error) {
	list, err := p.store.Post.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "postsrv.list")
	}

	return list, nil
}
