package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isido5ik/RecipePublishingPlatform/dtos"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body dtos.SignUpInput true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input dtos.SignUpInput

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

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body dtos.SignInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/auth/sign-in [post]
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
