package http

import (
	"github.com/gin-gonic/gin"
	"github.com/isido5ik/StoryPublishingPlatform/internal/usecase"
)

type Handler struct {
	useCases usecase.Usecase
}

func NewHandler(useCase usecase.Usecase) *Handler {
	return &Handler{useCases: useCase}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	story := router.Group("/story")
	{

		userIdentityMiddleware := h.UserIdentity()
		story.Use(userIdentityMiddleware)

		adminMiddleware := h.CheckRole(adminCtx)
		clientMiddleware := h.CheckRole(clientCtx)

		client := story.Group("/client")
		{
			client.POST("/", clientMiddleware, h.createPost)
			client.GET("/", clientMiddleware, h.getPost)
			client.DELETE("/:id", clientMiddleware, h.deletePost)
		}

		admin := story.Group("/admin")
		{
			admin.DELETE("/:id", adminMiddleware, h.deletePost)
		}

	}

	//other handlers
	return router
}
