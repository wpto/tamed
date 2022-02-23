package postsrv

import (
	"fmt"

	"github.com/pgeowng/tamed/model"
)

func (p *PostSrv) Modify(postID string, changes *model.PostChanges) error {
	fmt.Printf("postsrv.modify")
	return nil
}
