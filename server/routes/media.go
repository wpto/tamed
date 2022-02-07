package routes

import (
	"net/http"
	"strconv"
	"strings"

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
	fileID := c.Param("id")

	if len(fileID) < 1 {
		c.JSON(http.StatusBadRequest, "empty file id")
		return
	}

	mediaMeta, err := r.services.MediaMeta.Get(fileID)

	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, mediaMeta)
}

func (r *MediaRoute) Download(c *gin.Context) {
	fileID := c.Param("id")
	sizeVal := c.Query("size")
	formatVal := c.Query("format")

	if len(fileID) < 1 {
		c.JSON(http.StatusBadRequest, "empty file id")
		return
	}

	contentType, err := types.GetMime(formatVal)
	if err != nil {
		SendError(c, err)
		return
	}

	width, height, err := ParseSize(sizeVal)
	if err != nil {
		width, height = 1200, 1200
		// SendError(c, err)
	}

	mediaContent, err := r.services.MediaContent.Download(fileID, contentType, width, height)

	if err != nil {
		SendError(c, err)
		return
	}

	c.Data(http.StatusOK, contentType, mediaContent)
		return

}

func ParseSize(val string) (width, height int, err error) {
	tokens := strings.Split(val, "x")
	if len(tokens) != 2 {
		err = errors.Wrap(types.ErrBadRequest, "bad size")
		return
	}

	num, err := strconv.ParseInt(tokens[0], 10, 0)
	if err != nil {
		err = errors.Wrap(types.ErrBadRequest, "bad width")
		return
	}

	width = int(num)

	num, err = strconv.ParseInt(tokens[1], 10, 0)
	if err != nil {
		err = errors.Wrap(types.ErrBadRequest, "bad height")
		return
	}

	height = int(num)

	return
}

func SendError(c *gin.Context, err error) {
	if err != nil {
		var status int
		switch {
		case errors.Cause(err) == types.ErrNotFound:
			status = http.StatusNotFound
		case errors.Cause(err) == types.ErrBadRequest:
			status = http.StatusBadRequest
		case errors.Cause(err) == types.ErrNotImplemented:
			status = http.StatusNotImplemented
		default:
			status = http.StatusInternalServerError
		}

		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
}
