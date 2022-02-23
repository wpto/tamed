package model

type Post struct {
	PostID     string
	CreateTime string
	Tags       *Tags
	Link       string
}

type PostChanges struct {
	AddTags    *Tags
	RemoveTags *Tags
}

type Ordering int64

const (
	Recent Ordering = 0
)

var OrderingMap = map[string]Ordering{
	"recent": Recent,
}

type PostQuery struct {
	PostID      *string
	Order       Ordering
	IncludeTags *Tags
	ExcludeTags *Tags
	Limit       int
	Offset      int
}

type PostList struct {
	Page  int    `json:"page"`
	Pages int    `json:"pages"`
	Total int    `json:"total"`
	Posts []Post `json:"posts"`
	Tags  *Tags  `json:"tags"`
}
