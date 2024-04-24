package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	log.Print("this print from func (h *Handler) createPost(c *gin.Context)")
	c.JSON(http.StatusOK, "everything is fine")
}

func (h *Handler) getPost(c *gin.Context) {

}

func (h *Handler) updatePost(c *gin.Context) {

}

func (h *Handler) deletePost(c *gin.Context) {

}
