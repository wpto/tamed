package pg

import (
	"context"
	"database/sql"

	"github.com/pgeowng/tamed/config"
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type DB struct {
	*bun.DB
}

func Dial() (*DB, error) {
	dsn := config.Get().PgUrl
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	if config.Get().PgReset {
		ctx := context.Background()

		_, err := db.NewDropTable().Model((*Tags)(nil)).Exec(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "post.pg.drop")
		}

		ctx = context.Background()
		_, err = db.NewDropTable().Model((*Post)(nil)).Exec(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "post.pg.drop")
		}

		ctx = context.Background()
		_, err = db.NewCreateTable().Model((*Post)(nil)).Exec(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "post.pg.create")
		}

		ctx = context.Background()
		_, err = db.NewCreateTable().Model((*Tags)(nil)).Exec(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "post.pg.create")
		}
	}

	return &DB{db}, nil
}

type Post struct {
	bun.BaseModel `bun:"table:posts,alias:p"`

	PostID     string `bun:",pk"`
	CreateTime string
	Link       string
}

type Tags struct {
	bun.BaseModel `bun:"table:tags,alias:t"`

	PostID string
	Tag    string
}

func ToDB(post *model.Post) *Post {
	return &Post{
		PostID:     post.PostID,
		CreateTime: post.CreateTime,
		Link:       post.Link,
	}
}

func (p *Post) FromDB() *model.Post {
	return &model.Post{
		PostID:     p.PostID,
		CreateTime: p.CreateTime,
		Tags:       model.NewTags(),
		Link:       p.Link,
	}
}

func (db *DB) Create(postID string, post *model.Post) error {
	ctx := context.Background()

	p := ToDB(post)
	_, err := db.NewInsert().Model(p).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "post.pg.create")
	}

	return nil
}

func (db *DB) Get(postID string) (*model.Post, error) {

	ctx := context.Background()
	p := new(Post)
	err := db.NewSelect().Model(p).Where("post_id = ?", postID).Scan(ctx)

	if err != nil {
		return nil, errors.Wrap(types.ErrNotFound, "post.pg.get")
	}

	var tags []Tags
	err = db.NewSelect().
		Model(&tags).
		Column("tag").
		Where("? = ?", bun.Ident("post_id"), postID).
		Group("tag").
		Scan(ctx)

	tagStrings := make([]string, 0, len(tags))
	for _, tag := range tags {
		tagStrings = append(tagStrings, tag.Tag)
	}

	result := p.FromDB()
	result.Tags = model.NewTags(tagStrings...)

	return result, nil
}

func (db *DB) Query(query *model.PostQuery) (*model.PostList, error) {

	ctx := context.Background()

	var postsModel []Post
	err := db.NewSelect().Model(&postsModel).Offset(query.Offset).Limit(query.Limit).Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "post.pg.query")
	}

	postsId := make([]string, 0, len(postsModel))
	for _, post := range postsModel {
		postsId = append(postsId, post.PostID)
	}

	var tags []Tags
	err = db.NewSelect().
		Model(&tags).
		Where("? IN ?", bun.Ident("post_id"), bun.In(postsId)).
		Group("tag").
		Scan(ctx)

	tagStrings := make([]string, 0, len(tags))
	for _, tag := range tags {
		tagStrings = append(tagStrings, tag.Tag)
	}

	result := make([]model.Post, 0, len(postsModel))
	for idx := range postsModel {
		result = append(result, *postsModel[idx].FromDB())
	}

	postlist := &model.PostList{
		Page:  0,
		Pages: 1,
		Total: 20,
		Posts: result,
		Tags:  model.NewTags(tagStrings...),
	}

	return postlist, nil
}

func (db *DB) Modify(postID string, changes *model.PostChanges) error {

	if changes.RemoveTags.Len() > 0 {
		var del Tags
		ctx := context.Background()
		_, err := db.NewDelete().
			Model(&del).
			Where("? = ?", bun.Ident("post_id"), postID).
			Where("? IN (?)", bun.Ident("tag"), bun.In(changes.RemoveTags.Slice())).
			Exec(ctx)

		if err != nil {
			return errors.Wrap(err, "post.pg.removetags")
		}
	}

	if changes.AddTags.Len() > 0 {
		add := make([]*Tags, 0, changes.AddTags.Len())
		for _, tag := range changes.AddTags.Slice() {
			add = append(add, &Tags{PostID: postID, Tag: tag})
		}

		ctx := context.Background()
		_, err := db.NewInsert().Model(&add).Exec(ctx)
		if err != nil {
			return errors.Wrap(err, "post.pg.addtags")
		}
	}

	return nil
}

func (db *DB) Delete(postID string) error {
	ctx := context.Background()
	model := &Post{PostID: postID}

	_, err := db.NewDelete().Model(model).WherePK().Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "post.pg.delete")
	}

	return nil
}
