package viewroute

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/routes/commonroute"
	"github.com/pgeowng/tamed/service"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type ViewRoute struct {
	services *service.Manager
}

func NewViewRoute(services *service.Manager) *ViewRoute {
	return &ViewRoute{services}
}

func (r *ViewRoute) ViewArt(c *gin.Context) {
	artID := c.Param("id")

	if len(artID) < 1 {
		c.Redirect(http.StatusMovedPermanently, "/")
	}

	art, err := r.services.View.ViewArt(artID)
	if err != nil {
		commonroute.SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, art)
}

func (r *ViewRoute) ViewUser(c *gin.Context) {
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

func (r *ViewRoute) Search(c *gin.Context) {
	userName := c.Query("user")
	countStr := c.Query("count")
	offsetStr := c.Query("offset")
	orderStr := c.Query("order")

	var err error
	count := uint64(80)
	if countStr != "" {
		count, err = strconv.ParseUint(countStr, 10, 64)
		if err != nil {
			commonroute.SendError(c, errors.Wrap(types.ErrBadRequest, "search.count"))
			return
		}
	}

	offset := uint64(0)
	if offsetStr != "" {
		offset, err = strconv.ParseUint(countStr, 10, 64)
		if err != nil {
			commonroute.SendError(c, errors.Wrap(types.ErrBadRequest, "search.offset"))
			return
		}
	}

	order := types.Trending
	if len(orderStr) > 0 {
		var ok bool
		order, ok = types.OrderingMap[orderStr]
		if !ok {
			commonroute.SendError(c, errors.Wrap(types.ErrBadRequest, "order"))
			return
		}
	}

	// result, err := r.services.Search.Find()
	c.JSON(http.StatusOK, gin.H{
		"count":     count,
		"offset":    offset,
		"user_name": userName,
		"order":     order,
	})
}
