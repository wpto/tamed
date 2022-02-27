package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/types"
)

func (r *PostRoute) Get(c *gin.Context) {
	postID := c.Param("id")

	if len(postID) == 0 {
		SendError(c, types.ErrNotAllowed)
	}

	result, err := r.services.Post.Get(postID)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
