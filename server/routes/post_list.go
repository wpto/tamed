package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/model"
)

func (r *PostRoute) List(c *gin.Context) { // userName := c.Query("user") // countStr := c.Query("count") // offsetStr := c.Query("offset") // orderStr := c.Query("order") // var err error // count := uint64(80) // if countStr != "" {// 	count, err = strconv.ParseUint(countStr, 10, 64) // 	if err != nil {// 		commonroute.SendError(c, errors.Wrap(types.ErrBadRequest, "search.count")) // 		return // 	} // }
	// offset := uint64(0)
	// if offsetStr != "" {
	// 	offset, err = strconv.ParseUint(countStr, 10, 64)
	// 	if err != nil {
	// 		commonroute.SendError(c, errors.Wrap(types.ErrBadRequest, "search.offset"))
	// 		return
	// 	}
	// }

	// order := types.Trending
	// if len(orderStr) > 0 {
	// 	var ok bool
	// 	order, ok = types.OrderingMap[orderStr]
	// 	if !ok {
	// 		commonroute.SendError(c, errors.Wrap(types.ErrBadRequest, "order"))
	// 		return
	// 	}
	// }

	var query model.PostQuery
	_, err := r.services.Post.List(&query)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
