package http

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

		c.Set(userCtx, strconv.Itoa(userId))
		for _, role := range roles {
			c.Set(role.RoleName, strconv.Itoa(role.RoleId))
			log.Printf("adding the role %s with id %d to header of request", role.RoleName, role.RoleId)
		}

		c.Next()

	}
}

func (h *Handler) CheckRole(roleName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, ok := c.Get(roleName)
		if !ok {
			newErrorResponse(c, http.StatusForbidden, "user doesn't have the required role")
			return
		}
		log.Printf("role value, key: %s value: %s \n", roleName, roleValue)
		c.Next()
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("could not get user id from context")
	}

	idStr, ok := id.(string)
	if !ok {
		return 0, errors.New("invalid user id")
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("could not convert user id to integer")
	}

	return idInt, nil
}

func getRole(c *gin.Context) (string, error) {
	_, isAdmin := c.Get(adminCtx)
	if isAdmin {
		return adminCtx, nil
	}
	_, isClient := c.Get(clientCtx)
	if isClient {
		return clientCtx, nil
	}
	return "", errors.New("you do not have a permission")
}
