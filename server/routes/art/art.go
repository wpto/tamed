package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/service"
)

type ArtRoute struct {
	services *service.Manager
}

func NewArtRoute(services *service.Manager) *ArtRoute {
	return &ArtRoute{services}
}

func (r *ArtRoute) Create(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		SendError(c, err)
		return
	}

	files := form.File["upload[]"]
	metaList := make([]model.MediaMeta, 0)
	for _, fileHeader := range files {
		mediaID, err := r.services.MediaMeta.Create()
		if err != nil {
			SendError(c, err)
			return
		}

		err := r.services.MediaContent.Upload(mediaID, fileHeader)
		if err != nil {
			SendError(c, err)
			return
		}

		meta, err := r.services.MediaMeta.Get(mediaID)
		if err != nil {
			SendError(c, err)
			return
		}

		metaList = append(mediaList, media)
	}

	userName := "kifuku"

	obj, err := r.services.Art.Create(userName, mediaList)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusCreated, obj)
}

func (r *ArtRoute) Get(c *gin.Context) {
	artID := c.Param("id")

	if len(artID) < 1 {
		c.Redirect(http.StatusMovedPermanently, "/")
	}

	art, err := r.services.Art.Get(artID)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, art)
}
