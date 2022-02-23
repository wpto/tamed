package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/types"
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
		SendError(c, types.ErrNotAllowed)
		return
	}

	var opts ModifyOpts
	if err := c.ShouldBindJSON(&opts); err != nil {
		SendError(c, err)
		return
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

	c.JSON(http.StatusOK, gin.H{"ok": "changed"})
}
