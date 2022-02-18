package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *PostRoute) Modify(c *gin.Context) {
	// form, err := c.MultipartForm()
	// if err != nil {
	// 	SendError(c, err)
	// 	return
	// }

	// files := form.File["upload[]"]
	// if len(files) == 0 {
	// 	SendError(c, errors.Wrap(types.ErrBadRequest, "empty upload"))
	// 	return
	// }

	// obj, err := r.services.Post.Create(files)
	// if err != nil {
	// 	SendError(c, err)
	// 	return
	// }

	// c.JSON(http.StatusCreated, obj)
}
