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
	mediaList := make([]model.Media, 0)
	for _, fileHeader := range files {
		media, err := r.services.MediaContent.Upload(fileHeader)
		if err != nil {
			for _, media := range mediaList {
				// r.services.MediaContent.MarkClean(media.ID)
			}
		}
		mediaList = append(mediaList, media)
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

func (r *ArtRoute) Update(c *gin.Context) {

}

func (r *ArtRoute) Delete(c *gin.Context) {
}
