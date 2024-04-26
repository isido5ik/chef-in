package http

import (
	"github.com/gin-gonic/gin"
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type getStoriesResponse struct {
	Data []dtos.Story `json:"data"`
}
type getMyStoriesResponse struct {
	Username  string       `json:"username"`
	MyStories []dtos.Story `json:"stories"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		Message: message,
	})
}
