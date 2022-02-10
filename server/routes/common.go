package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

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
