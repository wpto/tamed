package routes

import (
	"testing"

	"github.com/pgeowng/tamed/model"
)

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
				return res.IncludeTags.Includes(model.NewTags("hello", "world")) &&
					res.ExcludeTags.Includes(model.NewTags("easy")) &&
					res.IncludeTags.Len() == 2 &&
					res.ExcludeTags.Len() == 1
			},
		},
		{tags: "           ",
			err: false,
			check: func(res *model.PostQuery) bool {
				return res.IncludeTags.Len() == 0 && res.ExcludeTags.Len() == 0
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
