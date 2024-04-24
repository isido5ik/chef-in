package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

func (h *Handler) signUp(c *gin.Context) {
	var input dtos.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	client_id, err := h.useCases.CreateUserAsClient(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"client_id": client_id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input dtos.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	token, rolesHeaders, err := h.useCases.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"rolesHeaders": rolesHeaders,
		"token":        token,
	})

}
