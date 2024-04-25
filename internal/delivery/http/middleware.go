package http

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	adminCtx            = "ADMIN"
	clientCtx           = "CLIENT"
)

func (h *Handler) UserIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
			return
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
			return
		}

		log.Printf("HEADER PARTS: %s \n %s \n", headerParts[0], headerParts[1])
		userId, roles, err := h.useCases.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Request.Header.Add(userCtx, strconv.Itoa(userId))
		for _, role := range roles {
			c.Request.Header.Add(role.RoleName, strconv.Itoa(role.RoleId))
			log.Printf("adding the role %s with id %d to header of request", role.RoleName, role.RoleId)
		}

		c.Next()

	}
}

func (h *Handler) CheckRole(roleName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue := c.Request.Header.Get(roleName)
		log.Printf("role value, key: %s value: %s \n", roleName, roleValue)
		if roleValue == "" {
			newErrorResponse(c, http.StatusForbidden, "user doesn't have the required role")
			return
		}
		c.Next()
	}
}
