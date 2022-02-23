package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/model"
)

type ModifyOpts struct {
	AddTags    *model.Tags `json:"add_tags"`
	RemoveTags *model.Tags `json:"rm_tags"`
}

func (opts *ModifyOpts) PostChanges() *model.PostChanges {
	return &model.PostChanges{
		AddTags:    opts.AddTags,
		RemoveTags: opts.RemoveTags,
	}
}

func (r *PostRoute) Modify(c *gin.Context) {
	postID := c.Param("id")

	if len(postID) == 0 {
		c.String(http.StatusMethodNotAllowed, "empty post id")
	}

	var opts ModifyOpts
	if err := c.ShouldBindJSON(&opts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if opts.AddTags == nil {
		opts.AddTags = model.NewTags()
	}
	if opts.RemoveTags == nil {
		opts.RemoveTags = model.NewTags()
	}

	err := r.services.Post.Modify(postID, opts.PostChanges())
	if err != nil {
		SendError(c, err)
		return
	}

	c.String(http.StatusOK, "changed")
}
