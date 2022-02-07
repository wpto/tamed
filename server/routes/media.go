package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/service"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type MediaRoute struct {
	services *service.Manager
}

func NewMediaRoute(services *service.Manager) *MediaRoute {
	return &MediaRoute{
		services: services,
	}
}

func (r *MediaRoute) Get(c *gin.Context) {
	fileId := c.Param("id")

	if len(fileId) < 1 {
		c.JSON(http.StatusBadRequest, "empty file id")
		return
	}

	mediaMeta, err := r.services.MediaMeta.Get(fileId)

	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, mediaMeta)
}

func (r *MediaRoute) Download(c *gin.Context) {
	fileId := c.Param("id")
	sizeVal := c.Query("size")
	formatVal := c.Query("format")

	if len(fileId) < 1 {
		c.JSON(http.StatusBadRequest, "empty file id")
		return
	}

	width, height, err := ParseSize(sizeVal)
	if err != nil {
		width, height = 1200, 1200
		// SendError(c, err)
	}	

	if formatVal == "jpg" {


		return
	}


	c.JSON(http.StatusBadRequest, errors.Wrap(types.ErrBadRequest("bad format")))
	return
}

func ParseSize(val string) (width, height int, err error) {
	tokens := strings.Split(val, "x")
	if len(tokens) != 2 {
		err = errors.Wrap(types.ErrBadRequest, "bad size")
		return
	}

	width, err = strconv.ParseInt(tokens[0], 10, 0)
	if err != nil {
		err = errors.Wrap(types.ErrBadRequest, "bad width")
		return
	}

	height, err = strconv.ParseInt(tokens[1], 10, 0)
	if err != nil {
		err = errors.Wrap(types.ErrBadRequest, "bad height")
		return
	}

	return
}

func SendError(c *gin.Context, err error) {
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrNotFound:
			c.JSON(http.StatusNotFound, err)
		case errors.Cause(err) == types.ErrBadRequest:
			c.JSON(http.StatusBadRequest, err)
		default:
			c.JSON(http.StatusInternalServerError, err)
		}
		return
	}
}
