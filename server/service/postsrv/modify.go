package postsrv

import (
	"github.com/pgeowng/tamed/model"
	"github.com/pkg/errors"
)

func (p *PostSrv) Modify(postID string, changes *model.PostChanges) error {
	err := p.store.Post.Modify(postID, changes)
	if err != nil {
		return errors.Wrap(err, "postsrv.modify")
	}
	return nil
}
