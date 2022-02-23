package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/service"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

func DetectStatus(err error) int {
	var status int
	switch {
	case errors.Cause(err) == types.ErrNotFound:
		status = http.StatusNotFound
	case errors.Cause(err) == types.ErrBadRequest:
		status = http.StatusBadRequest
	case errors.Cause(err) == types.ErrNotImplemented:
		status = http.StatusNotImplemented
	case errors.Cause(err) == types.ErrNotAllowed:
		status = http.StatusMethodNotAllowed
	default:
		status = http.StatusInternalServerError
	}
	return status
}

func SendError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(DetectStatus(err), gin.H{"error": err.Error()})
		return
	}
}

func SendErrorPage(c *gin.Context, err error) {
	c.HTML(DetectStatus(err), "content_error.tmpl", gin.H{})
}

type PostRoute struct {
	services *service.Manager
}

type PageRoute struct {
	services *service.Manager
}

func NewPostRoute(services *service.Manager) *PostRoute {
	return &PostRoute{services}
}

func NewPageRoute(services *service.Manager) *PageRoute {
	return &PageRoute{services}
}
