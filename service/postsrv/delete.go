package postsrv

import "github.com/pkg/errors"

func (p *PostSrv) Delete(postID string) error {
	err := p.store.Media.Delete(postID)
	if err != nil {
		return errors.Wrap(err, "postsrv.delete")
	}

	err = p.store.Post.Delete(postID)
	if err != nil {
		return errors.Wrap(err, "postsrv.delete")
	}

	return nil
}
