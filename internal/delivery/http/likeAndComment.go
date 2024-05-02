package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

func (h *Handler) like(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.useCases.Like(userId, postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "liked",
	})

}

func (h *Handler) removeLike(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.useCases.RemoveLike(userId, postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "like removed",
	})

}

func (h *Handler) addComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	parentId, err := dtos.GetParentComment(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var comment dtos.Comment
	if err := c.BindJSON(&comment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.useCases.AddComment(userId, postId, parentId, comment.CommentText)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "comment added",
	})
}

func (h *Handler) updateComment(c *gin.Context) {

	commentId, err := dtos.GetCommentId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var newComment dtos.UpdateCommentInput
	if err := c.BindJSON(&newComment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.useCases.UpdateComment(userId, postId, commentId, newComment)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "comment edited",
	})
}

func (h *Handler) deleteComment(c *gin.Context) {
	commentId, err := dtos.GetCommentId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.useCases.DeleteComment(userId, postId, commentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "comment deleted",
	})
}
