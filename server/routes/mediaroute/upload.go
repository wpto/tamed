package mediaroute

// func (r *MediaRoute) Upload(c *gin.Context) {
// 	fileHeader, err := c.FormFile("file")
// 	if err != nil {
// 		SendError(c, err)
// 		return
// 	}

// 	err = r.services.MediaContent.Upload(fileHeader)
// 	if err != nil {
// 		SendError(c, errors.Wrap(err, "route.mediaupload"))
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{})
// }
