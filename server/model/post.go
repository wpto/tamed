package model

type Tag struct {
	Label string
}

func NewTag(label string) Tag {
	return Tag{Label: label}
}

type Post struct {
	PostID     string
	CreateTime string
	Tags       []Tag
	Link       string
}

type PostChanges struct {
	AddTags    []Tag
	RemoveTags []Tag
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
	IncludeTags []Tag
	ExcludeTags []Tag
	Limit       int
	Offset      int
}

type PostList struct {
	Page  int    `json:"page"`
	Pages int    `json:"pages"`
	Total int    `json:"total"`
	Posts []Post `json:"posts"`
	Tags  []Tag  `json:"tags"`
}

func ContainTag(list []Tag, label string) bool {
	for _, tag := range list {
		if tag.Label == label {
			return true
		}
	}
	return false
}
