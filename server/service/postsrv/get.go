package postsrv

import (
  "fmt"

  "github.com/pgeowng/tamed/model"
)

func (p *PostSrv) Get(postID string) (*model.Post, error) {
  // result, err = srv.store.View.GetArt(artID)
  // if err != nil {
  //   return nil, errors.Wrap(err, "srv.view.art")
  // }

  // if result == nil {
  //   return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("srv.view.art: Art '%s' not found!", artID))
  // }

  // return
  fmt.Printf("postsrv.get")
  return nil, nil
}
