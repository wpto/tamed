package postsrv

import "fmt"

func (p *PostSrv) Delete(postID string) error {
	fmt.Printf("postsrv.delete")
	return nil
}
