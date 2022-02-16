package userroute

import (
	"github.com/pgeowng/tamed/service"
)

type UserRoute struct {
	services *service.Manager
}

func NewUserRoute(services *service.Manager) *UserRoute {
	return &UserRoute{services}
}
