package userroute

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/routes/commonroute"
	"github.com/pgeowng/tamed/service"
)

type UserRoute struct {
	services *service.Manager
}

func NewUserRoute(services *service.Manager) *UserRoute {
	return &UserRoute{services}
}

func (r *UserRoute) Get(c *gin.Context) {
	userName := c.Param("id")

	if len(userName) < 1 {
		c.Redirect(http.StatusMovedPermanently, "/")
	}

	art, err := r.services.View.ViewUser(userName)
	if err != nil {
		commonroute.SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, art)
}
