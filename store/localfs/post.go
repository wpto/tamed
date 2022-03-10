package localfs

import (
	"encoding/json"

	"github.com/pgeowng/tamed/model"
	"github.com/pkg/errors"
)

type PostStore struct {
	repo *FileRepo
}

func NewPostStore(repo *FileRepo) *PostStore {
	return &PostStore{repo}
}

func (store *PostStore) Get(postID string) (*model.Post, error) {
	prev, err := store.repo.Get(postID)
	if err != nil {
		return nil, errors.Wrap(err, "poststore.get")
	}

	var entry model.Post
	err = json.Unmarshal(prev, &entry)
	if err != nil {
		return nil, errors.Wrap(err, "poststore.get")
	}

	return &entry, nil
}

func (store *PostStore) Query(query *model.PostQuery) (*model.PostList, error) {

	if query.PostID != nil {
		post, err := store.Get(*query.PostID)
		if err != nil {
			return nil, errors.Wrap(err, "poststore.query")
		}

		return &model.PostList{
			Next:  false,
			Posts: []model.Post{*post},
			Tags:  post.Tags,
		}, nil
	}

	dbText, err := store.repo.All()
	if err != nil {
		return nil, errors.Wrap(err, "poststore.query")
	}

	var db []model.Post
	err = json.Unmarshal(dbText, &db)
	if err != nil {
		return nil, errors.Wrap(err, "poststore.query")
	}

	filtered := []model.Post{}

	for _, entry := range db {
		if entry.Tags.Includes(query.IncludeTags) &&
			entry.Tags.Excludes(query.ExcludeTags) {
			filtered = append(filtered, entry)
		}
	}

	total := len(filtered)
	//pages := int(math.Ceil(float64(total) / float64(query.Limit)))

	left := query.Offset * query.Limit
	//page := int(math.Floor(float64(left) / float64(query.Limit)))
	if left >= total {
		left = total
		//page = pages
	}

	hasNext := true
	right := left + query.Limit
	if right > total {
		right = total
		hasNext = false
	}

	return &model.PostList{
		Next:  hasNext,
		Posts: filtered[left:right],
		Tags:  model.NewTags(),
	}, nil
}
func (store *PostStore) Create(postID string, post *model.Post) error {
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

func (store *PostStore) Modify(postID string, changes *model.PostChanges) error {
	postBody, err := store.repo.Get(postID)
	if err != nil {
		return errors.Wrap(err, "poststore.modify")
	}

	var entry model.Post
	err = json.Unmarshal(postBody, &entry)
	if err != nil {
		return errors.Wrap(err, "poststore.modify")
	}

	entry.Tags.Include(changes.AddTags)
	entry.Tags.Exclude(changes.RemoveTags)

	postBody, err = json.Marshal(entry)
	if err != nil {
		return errors.Wrap(err, "poststore.modify")
	}

	err = store.repo.Write(postID, postBody)
	if err != nil {
		return errors.Wrap(err, "poststore.modify")
	}

	return nil
}

func (store *PostStore) Delete(postID string) error {
	err := store.repo.Delete(postID)
	if err != nil {
		return errors.Wrap(err, "poststore.delete")
	}
	return nil
}
