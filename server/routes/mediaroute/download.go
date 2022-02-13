package mediaroute

// func (r *MediaRoute) Get(c *gin.Context) {
// 	mediaID := c.Param("id")
// 	sizeVal := c.Query("size")
// 	formatVal := c.Query("format")
// 	metaVal := c.Query("meta")
// 	fmt.Println("metaval", metaVal)

// 	if len(mediaID) < 1 {
// 		c.JSON(http.StatusBadRequest, "empty file id")
// 		return
// 	}

// 	if len(metaVal) > 0 {
// 		r.serveMeta(c, mediaID)
// 		return
// 	}

// 	contentType := ""
// 	if len(formatVal) > 0 {
// 		var err error
// 		contentType, err = types.GetMime(formatVal)
// 		if err != nil {
// 			SendError(c, err)
// 			return
// 		}
// 	}

// 	width, err := ParseSize(sizeVal)
// 	if err != nil {
// 		SendError(c, err)
// 		return
// 	}
// 	content, err := r.services.MediaContent.Download(mediaID, contentType, width)

// 	if err != nil {
// 		SendError(c, err)
// 		return
// 	}

// 	if len(contentType) == 0 {
// 		// MediaContent.Download should return proper MediaObj{io.Reader, contentType}
// 		// temp solution
// 		meta, err := r.services.MediaMeta.Get(mediaID)
// 		if err != nil {
// 			SendError(c, err)
// 			return
// 		}
// 		contentType = meta.Mime
// 	}

// 	c.Data(http.StatusOK, contentType, content)
// 	return
// }

// func (r *MediaRoute) serveMeta(c *gin.Context, mediaID string) {
// 	meta, err := r.services.MediaMeta.Get(mediaID)
// 	if err != nil {
// 		SendError(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, meta)
// }

// var DefaultSizeParam = 0
// var SizeParam = map[string]int{
// 	"xs": 144, // height for vid should be divisible by 2
// 	"s":  680,
// 	"m":  1200,
// 	"l":  2048,
// 	"xl": 4096,
// }

// func ParseSize(val string) (width int, err error) {
// 	if len(val) == 0 {
// 		return DefaultSizeParam, nil
// 	}

// 	width, ok := SizeParam[val]
// 	if !ok {
// 		return 0, errors.New("bad size")
// 	}
// 	return width, nil
// }
