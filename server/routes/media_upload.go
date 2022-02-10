package routes

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (r *MediaRoute) Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		SendError(c, err)
	}

	file, err := fileHeader.Open()
	if err != nil {
		SendError(c, err)
	}

	fileBody, err := ioutil.ReadAll(file)
	if err != nil {
		SendError(c, err)
	}

	err = r.services.MediaContent.Upload(fileBody)
	if err != nil {
		SendError(c, errors.Wrap(err, "route.mediaupload"))
	}

	c.JSON(http.StatusOK, gin.H{})
}
