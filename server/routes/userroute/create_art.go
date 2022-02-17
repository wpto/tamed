package userroute

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/routes/commonroute"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

func (r *UserRoute) CreateArt(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		commonroute.SendError(c, err)
		return
	}

	files := form.File["upload[]"]
	if len(files) == 0 {
		commonroute.SendError(c, errors.Wrap(types.ErrBadRequest, "empty upload"))
		return
	}
	mediaList := make([]model.Media, 0)
	for _, fileHeader := range files {
		media, err := r.services.User.Upload(fileHeader)
		if err != nil {
			commonroute.SendError(c, err)
			return
		}

		mediaList = append(mediaList, *media)
	}

	userName := "kifuku"
	// based on auth assign username

	obj, err := r.services.User.CreateArt(userName, mediaList)
	if err != nil {
		commonroute.SendError(c, err)
		return
	}

	c.JSON(http.StatusCreated, obj)
}
