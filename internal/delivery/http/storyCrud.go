package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

func (h *Handler) createStory(c *gin.Context) {
	var story dtos.Post
	if err := c.BindJSON(&story); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	postId, err := h.useCases.CreateStory(story, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"postId": postId,
	})
}

func (h *Handler) getStories(c *gin.Context) {

	stories, err := h.useCases.GetStories()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getStoriesResponse{
		Data: stories,
	})
}

func (h *Handler) getUsersStories(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	username, stories, err := h.useCases.GetUsersStories(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getMyStoriesResponse{
		Username:  username,
		MyStories: stories,
	})
}

func (h *Handler) getStory(c *gin.Context) {

	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	story, err := h.useCases.GetStory(postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"story:": story,
	})
}

func (h *Handler) updateStory(c *gin.Context) {
	var input dtos.UpdateStoryInput
	if err := c.BindJSON(&input); err != nil {
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
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	role, err := getRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	err = h.useCases.UpdateStory(postId, userId, role, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}

func (h *Handler) deleteStory(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	role, err := getRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	err = h.useCases.DeleteStory(postId, userId, role)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})

}

func (h *Handler) getUsers(c *gin.Context) {
	// TODO: Implement handler for getting users
}

func (h *Handler) getUser(c *gin.Context) {
	// TODO: Implement handler for getting a specific user
}

func (h *Handler) updateUser(c *gin.Context) {
	// TODO: Implement handler for updating a user
}

func (h *Handler) deleteUser(c *gin.Context) {
	// TODO: Implement handler for deleting a user
}
