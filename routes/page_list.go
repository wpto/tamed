package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *PageRoute) List(c *gin.Context) {
	postIDStr := c.Query("id")
	orderStr := c.Query("order")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	tagsStr := c.Query("tags")

	query, err := ListArgs(postIDStr, orderStr, limitStr, offsetStr, tagsStr)
	if err != nil {
		SendErrorPage(c, err)
		return
	}

	res, err := r.services.Post.List(query)
	if err != nil {
		SendErrorPage(c, err)
		return
	}

	fmt.Println(res)

	c.HTML(http.StatusOK, "content_list.tmpl", gin.H{
		"title": "Posts",
		"list":  res,
	})
}
