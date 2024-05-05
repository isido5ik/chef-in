package http

import (
	"github.com/gin-gonic/gin"
	"github.com/isido5ik/RecipePublishingPlatform/dtos"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type getRecipesResponse struct {
	Data []dtos.Recipe `json:"data"`
}
type getMyRecipesResponse struct {
	Username  string        `json:"username"`
	MyRecipes []dtos.Recipe `json:"recipes"`
}

type getAllCommentsResponse struct {
	Comments []dtos.Comment `json:"comments"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		Message: message,
	})
}
