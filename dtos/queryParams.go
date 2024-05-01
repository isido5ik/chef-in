package dtos

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParentComment(c *gin.Context) (int, error) {
	parentIdStr := c.DefaultQuery("parent_comment", "0")

	parent_id, err := strconv.Atoi(parentIdStr)
	if err != nil {
		return -1, errors.New("invalid value of query parameter 'page' (it must by integer)")
	}

	if parent_id < 0 {
		return -1, errors.New("invalid value of query parameter 'parent_comment', negative number")
	}
	return parent_id, nil
}

func GetCommentId(c *gin.Context) (int, error) {
	commentIdStr := c.DefaultQuery("comment_id", "-1")

	comment_id, err := strconv.Atoi(commentIdStr)
	if err != nil {
		return -1, errors.New("invalid value of query parameter 'comment_id' (it must by integer)")
	}

	if comment_id <= 0 {
		return -1, errors.New("invalid value of query parameter 'comment_id', negative number or zero")
	}
	return comment_id, nil
}
