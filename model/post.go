package model

type Post struct {
	PostID     string `json:"id"`
	CreateTime string `json:"ctime"`
	Tags       *Tags  `json:"tags"`
	Link       string `json:"link"`
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
	Next  bool   `json:"next"`
	Posts []Post `json:"posts"`
	Tags  *Tags  `json:"tags"`
}
