package artroute

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/routes/commonroute"
	"github.com/pgeowng/tamed/service"
)

type ArtRoute struct {
	services *service.Manager
}

func NewArtRoute(services *service.Manager) *ArtRoute {
	return &ArtRoute{services}
}

// func (r *ArtRoute) Create(c *gin.Context) {
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		SendError(c, err)
// 		return
// 	}

// 	files := form.File["upload[]"]
// 	metaList := make([]model.MediaMeta, 0)
// 	for _, fileHeader := range files {
// 		media, err := r.services.Media.Upload(fileHeader)
// 		if err != nil {
// 			SendError(c, err)
// 			return
// 		}

// 		metaList = append(mediaList, media)
// 	}

// 	userName := "kifuku"

// 	obj, err := r.services.Art.Create(userName, mediaList)
// 	if err != nil {
// 		SendError(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusCreated, obj)
// }

func (r *ArtRoute) Get(c *gin.Context) {
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
