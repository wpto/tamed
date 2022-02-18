package model

type Tag struct {
	Label string
}

type TagKey struct {
	Tag
	Exclude bool
}

type Post struct {
	PostID     string
	CreateTime string
	Tags       []Tag
	Link       string
}

type PostChanges struct {
	Tags []TagKey
}

type PostQuery struct {
	PostID *string
	Tags   []TagKey
	Order  *string
}

type PostList struct {
	Posts []Post
	Tags  []Tag
}
