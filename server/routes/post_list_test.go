package routes

import (
	"testing"

	"github.com/pgeowng/tamed/model"
)

func containTag(list []model.Tag, label string) bool {
	for _, tag := range list {
		if tag.Label == label {
			return true
		}
	}
	return false
}

func TestList(t *testing.T) {
	cases := []struct {
		postid string
		order  string
		limit  string
		offset string
		tags   string
		err    bool
		check  func(res *model.PostQuery) bool
	}{
		{postid: "hello",
			err: false,
			check: func(res *model.PostQuery) bool {
				return res.PostID != nil && *res.PostID == "hello"
			},
		},
		{tags: "hello world -easy",
			err: false,
			check: func(res *model.PostQuery) bool {
				return containTag(res.IncludeTags, "hello") && containTag(res.IncludeTags, "world") && containTag(res.ExcludeTags, "easy") && len(res.IncludeTags) == 2 && len(res.ExcludeTags) == 1
			},
		},
		{tags: "           ",
			err: false,
			check: func(res *model.PostQuery) bool {
				return len(res.IncludeTags) == 0 && len(res.ExcludeTags) == 0
			},
		},
	}

	for _, c := range cases {
		res, err := ListArgs(c.postid, c.order, c.limit, c.offset, c.tags)
		if err != nil && !c.err {
			t.Logf("expected res, got err %v\n", err)
			t.Fail()
		} else if err == nil && c.err {
			t.Logf("expected err, got res %#v\n", res)
			t.Fail()
		} else if !c.check(res) {
			t.Logf("wrong res for %#v\n", c)
			t.Logf("res: %#v\n", res)
			t.Fail()
		}
	}
}
