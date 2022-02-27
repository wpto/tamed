package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

func (r *PostRoute) Create(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		SendError(c, err)
		return
	}

	files := form.File["upload[]"]
	if len(files) == 0 {
		SendError(c, errors.Wrap(types.ErrBadRequest, "empty upload"))
		return
	}

	obj, err := r.services.Post.Create(files)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusCreated, obj)
}
