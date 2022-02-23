package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/model"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

func (r *PostRoute) List(c *gin.Context) {
	postIDStr := c.Query("id")
	orderStr := c.Query("order")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	tagsStr := c.Query("tags")

	query, err := ListArgs(postIDStr, orderStr, limitStr, offsetStr, tagsStr)
	if err != nil {
		SendError(c, err)
		return
	}

	fmt.Println(query)

	res, err := r.services.Post.List(query)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func ListArgs(postIDStr, orderStr, limitStr, offsetStr, tagsStr string) (*model.PostQuery, error) {
	var postID *string = nil
	if len(postIDStr) > 0 {
		postID = &postIDStr
	}

	order := model.Recent
	if len(orderStr) > 0 {
		var ok bool
		order, ok = model.OrderingMap[orderStr]
		if !ok {
			return nil, errors.Wrap(types.ErrBadRequest, "bad order")
		}
	}

	limit := 20
	if len(limitStr) > 0 {
		limit64, err := strconv.ParseInt(limitStr, 10, 0)
		if err != nil || limit64 < 1 {
			return nil, errors.Wrap(types.ErrBadRequest, "bad limit")
		}
		limit = int(limit64)
	}

	offset := 0
	if len(offsetStr) > 0 {
		offset64, err := strconv.ParseInt(offsetStr, 10, 0)
		if err != nil || offset64 < 0 {
			return nil, errors.Wrap(types.ErrBadRequest, "search.offset")
		}
		offset = int(offset64)
	}

	incTags := []string{}
	excTags := []string{}
	if len(tagsStr) > 0 {
		tagsList := strings.Split(tagsStr, " ")
		for _, tag := range tagsList {
			if len(tag) > 0 {
				if strings.HasPrefix(tag, "-") {
					excTags = append(excTags, strings.TrimPrefix(tag, "-"))
				} else {
					incTags = append(incTags, tag)
				}
			}
		}
	}

	return &model.PostQuery{
		PostID:      postID,
		Order:       order,
		IncludeTags: model.NewTags(incTags...),
		ExcludeTags: model.NewTags(excTags...),
		Limit:       limit,
		Offset:      offset,
	}, nil
}
