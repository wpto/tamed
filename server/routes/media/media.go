package routes

import "github.com/pgeowng/tamed/service"

type MediaRoute struct {
	services *service.Manager
}

func NewMediaRoute(services *service.Manager) *MediaRoute {
	return &MediaRoute{
		services: services,
	}
}
