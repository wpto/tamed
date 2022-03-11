package pg

import (
	"context"
	"database/sql"
	"fmt"

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

		db.NewDropTable().Model((*Tags)(nil)).Exec(ctx)
		// if err != nil {
		// 	return nil, errors.Wrap(err, "post.pg.drop")
		// }

		ctx = context.Background()
		db.NewDropTable().Model((*Post)(nil)).Exec(ctx)
		// if err != nil {
		// 	return nil, errors.Wrap(err, "post.pg.drop")
		// }

		ctx = context.Background()
		_, err := db.NewCreateTable().Model((*Post)(nil)).Exec(ctx)
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

	result := p.FromDB()
	result.Tags = model.NewTags(db.getTags(postID)...)

	return result, nil
}

func (db *DB) getTags(postID string) []string {
	ctx := context.Background()

	var tags []Tags
	err := db.NewSelect().
		Model(&tags).
		Column("tag").
		Where("? = ?", bun.Ident("post_id"), postID).
		Group("tag").
		Order("tag ASC").
		Scan(ctx)

	if err != nil {
		panic(err)
	}

	tagStrings := make([]string, 0, len(tags))
	for _, tag := range tags {
		tagStrings = append(tagStrings, tag.Tag)
	}
	return tagStrings
}

func (db *DB) containTags(ctx context.Context, list []string) ([]string, error) {

	var tags []Tags
	err := db.NewSelect().
		Model(&tags).
		Column("post_id").
		Where("? IN (?)", bun.Ident("tag"), bun.In(list)).
		Group("post_id").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(tags))
	for _, tag := range tags {
		result = append(result, tag.PostID)
	}

	return result, nil
}

func (db *DB) Query(query *model.PostQuery) (*model.PostList, error) {

	ctx := context.Background()

	fmt.Println(query)

	// isTagFilter := query.IncludeTags.Len() > 0 || query.ExcludeTags.Len() > 0

	var excSQL *bun.SelectQuery
	if query.ExcludeTags.Len() > 0 {
		excSQL = db.NewSelect().
			Model((*Tags)(nil)).
			Column("post_id").
			Where("? IN (?)", bun.Ident("tag"), bun.In(query.ExcludeTags.Slice())).
			Group("post_id")
	}

	var postsModel []Post
	postSQL := db.NewSelect().
		Model(&postsModel)

	if query.IncludeTags.Len() > 0 {
		postSQL = postSQL.
			Table("tags").
			Where("p.post_id = tags.post_id").
			Where("? IN (?)", bun.Ident("tags.tag"), bun.In(query.IncludeTags.Slice()))
	}
	if query.ExcludeTags.Len() > 0 {
		postSQL = postSQL.Where("? NOT IN (?)", bun.Ident("p.post_id"), excSQL)
	}

	if query.IncludeTags.Len() > 0 {
		postSQL = postSQL.Group("p.post_id").
			Having("COUNT(?) = ?", bun.Ident("p.post_id"), query.IncludeTags.Len())
	}

	postSQL = postSQL.Offset(query.Offset*query.Limit).
		Limit(query.Limit+1).
		Order("p.create_time DESC", "p.post_id DESC")

	err := postSQL.Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "post.pg.query")
	}

	hasNext := len(postsModel) == query.Limit+1
	fmt.Println("query", len(postsModel))

	length := len(postsModel)
	if hasNext {
		length = query.Limit
		postsModel = postsModel[:length]
	}

	postsId := make([]string, 0, length)
	for _, post := range postsModel {
		postsId = append(postsId, post.PostID)
	}

	var tags []Tags
	if len(postsId) > 0 {
		err = db.NewSelect().
			Model(&tags).
			Column("tag").
			Where("? IN (?)", bun.Ident("post_id"), bun.In(postsId)).
			Group("tag").
			Order("tag ASC").
			Scan(ctx)
	}

	tagStrings := make([]string, 0, len(tags))
	for _, tag := range tags {
		tagStrings = append(tagStrings, tag.Tag)
	}

	result := make([]model.Post, 0, length)
	for idx := range postsModel {
		item := postsModel[idx].FromDB()
		item.Tags = model.NewTags(db.getTags(item.PostID)...)
		result = append(result, *item)
	}
	fmt.Println("result", len(result))

	postlist := &model.PostList{
		Next:  hasNext,
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
