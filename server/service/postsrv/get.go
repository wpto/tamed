package postsrv

import (
  "github.com/pgeowng/tamed/model"
  "github.com/pkg/errors"
)

func (p *PostSrv) Get(postID string) (*model.Post, error) {
  result, err := p.store.Post.Get(postID)
  if err != nil {
    return nil, errors.Wrap(err, "postsrv.getpost")
  }

  return result, nil
}
