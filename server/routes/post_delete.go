package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/types"
)

func (r *PostRoute) Delete(c *gin.Context) {
	postID := c.Param("id")

	if len(postID) == 0 {
		SendError(c, types.ErrNotAllowed)
	}

	err := r.services.Post.Delete(postID)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": "deleted"})
}
