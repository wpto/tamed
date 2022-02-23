package store

import (
	"encoding/json"
	"fmt"

	"github.com/pgeowng/tamed/model"
	"github.com/pkg/errors"
)

type PostStoreImpl struct {
	repo FileRepo
}

func NewPostStoreImpl(repo FileRepo) *PostStoreImpl {
	return &PostStoreImpl{repo}
}

func (store *PostStoreImpl) Get(postID string) (*model.Post, error) {
	fmt.Println("poststore.get")
	return nil, nil
}

func (store *PostStoreImpl) Query(query *model.PostQuery) (*PostList, error) {
	fmt.Println("poststore.query")
	return &PostList{
		Posts: []model.Post{},
		Tags:  []model.Tag{},
	}, nil
}
func (store *PostStoreImpl) Create(postID string, post *model.Post) error {
	body, err := json.Marshal(*post)
	if err != nil {
		return errors.Wrap(err, "poststore.create.marshal")
	}

	err = store.repo.Create(postID, body)
	if err != nil {
		return errors.Wrap(err, "poststore.create")
	}

	return nil
}

func (store *PostStoreImpl) Modify(postID string, changes *model.PostChanges) error {
	// prev, err := store.repo.Get(postID)
	// if err != nil {
	// 	return errors.Wrap(err, "poststore.modify")
	// }

	// var entry model.Post
	// err = json.Unmarshal(prev, &entry)
	// if err != nil {
	// 	return errors.Wrap(err, "poststore.modify")
	// }

	// body, err := json.Marshal(entry)
	// if err != nil {
	// 	return errors.Wrap(err, "poststore.modify")
	// }

	// err = store.repo.Write(postID, body)
	// if err != nil {
	// 	return errors.Wrap(err, "poststore.modify")
	// }

	return nil
}

func (store *PostStoreImpl) Delete(postID string) error {
	fmt.Println("poststore.delete")
	return nil
}
