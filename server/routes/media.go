package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/service"
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
