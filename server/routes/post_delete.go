package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *PostRoute) Delete(c *gin.Context) {
	postID := c.Param("id")

	if len(postID) == 0 {
		c.String(http.StatusMethodNotAllowed, "empty post id")
	}

	err := r.services.Post.Delete(postID)
	if err != nil {
		SendError(c, err)
		return
	}

	c.String(http.StatusOK, "deleted")
}
