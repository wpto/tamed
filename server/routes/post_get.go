package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *PostRoute) Get(c *gin.Context) {
	postID := c.Param("id")

	if len(postID) == 0 {
		c.String(http.StatusMethodNotAllowed, "empty post id")
	}

	result, err := r.services.Post.Get(postID)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
