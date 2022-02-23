package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/model"
)

func (r *PostRoute) Modify(c *gin.Context) {
	postID := c.Param("id")

	if len(postID) == 0 {
		c.String(http.StatusMethodNotAllowed, "empty post id")
	}

	addTags, ok := c.GetPostFormArray("add_tag")
	if !ok {
		addTags = []string{}
	}

	rmTags, ok := c.GetPostFormArray("rm_tag")
	if !ok {
		rmTags = []string{}
	}

	req := &model.PostChanges{
		AddTags:    []model.Tag{},
		RemoveTags: []model.Tag{},
	}

	for _, tag := range addTags {
		req.AddTags = append(req.AddTags, model.NewTag(tag))
	}

	for _, tag := range rmTags {
		req.RemoveTags = append(req.RemoveTags, model.NewTag(tag))
	}

	err := r.services.Post.Modify(postID, req)
	if err != nil {
		SendError(c, err)
		return
	}

	c.String(http.StatusOK, "changed")
}
