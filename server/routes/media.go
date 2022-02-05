package routes

import (
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
	
}
