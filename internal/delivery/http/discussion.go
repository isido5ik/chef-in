package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isido5ik/RecipePublishingPlatform/dtos"
)

// @Summary Add comment to post
// @Security ApiKeyAuth
// @Tags comments
// @Description add comment to post
// @ID add-comment-to-post
// @Accept  json
// @Produce  json
// @Param :recipe_id path int true "Post ID"
// @Param input body dtos.NewComment true "recipe info"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/recipes/{:recipe_id}/comment [post]
func (h *Handler) addComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	postId, err := strconv.Atoi(c.Param("recipe_id"))
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

// @Summary Get all comments
// @Security ApiKeyAuth
// @Tags comments
// @Description get all comments
// @ID get-all-comments
// @Accept  json
// @Produce  json
// @Param :recipe_id path int true "Post ID"
// @Success 200 {object} getAllCommentsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/recipes/{:recipe_id}/comment [get]
func (h *Handler) getAllComments(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("recipe_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	comments, err := h.useCases.GetAllComments(postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCommentsResponse{
		Comments: comments,
	})
}

// @Summary Update comment
// @Security ApiKeyAuth
// @Tags comments
// @Description update comment
// @ID update-comment
// @Accept  json
// @Produce  json
// @Param :recipe_id path int true "Post ID"
// @Param comment_id query int true "Comment ID"
// @Param input body dtos.UpdateCommentInput true "comment info"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/recipes/{:recipe_id}/comment [put]
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
	postId, err := strconv.Atoi(c.Param("recipe_id"))
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

// @Summary Delete comment
// @Security ApiKeyAuth
// @Tags comments
// @Description delete comment
// @ID delete-comment
// @Accept  json
// @Produce  json
// @Param :recipe_id path int true "Post ID"
// @Param comment_id query int true "Comment ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/recipes/{:recipe_id}/comment [delete]
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
	postId, err := strconv.Atoi(c.Param("recipe_id"))
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
