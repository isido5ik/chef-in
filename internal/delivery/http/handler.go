package http

import (
	"github.com/gin-gonic/gin"
	"github.com/isido5ik/RecipePublishingPlatform/internal/usecase"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/isido5ik/RecipePublishingPlatform/docs"
	swaggerFiles "github.com/swaggo/files"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	adminCtx            = "ADMIN"
	clientCtx           = "CLIENT"
	rolesCtx            = "roles"
)

type Handler struct {
	useCases usecase.Usecase
}

func NewHandler(useCase usecase.Usecase) *Handler {
	return &Handler{useCases: useCase}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}

		recipes := api.Group("/recipes")
		{
			// Middleware
			userIdentityMiddleware := h.UserIdentity()

			// Client routes
			recipes.GET("/", h.getRecipes)

			recipes.POST("/", userIdentityMiddleware, h.createRecipe)
			recipes.GET("/my", userIdentityMiddleware, h.getUsersRecipes)
			recipes.GET("/:recipe_id", userIdentityMiddleware, h.getRecipe)
			recipes.PUT("/:recipe_id", userIdentityMiddleware, h.updateRecipe)
			recipes.DELETE("/:recipe_id", userIdentityMiddleware, h.deleteRecipe)

			comment := recipes.Group("/:recipe_id/comment", userIdentityMiddleware)
			{
				comment.POST("/", h.addComment)
				comment.GET("/", h.getAllComments)
				comment.PUT("/", h.updateComment)
				comment.DELETE("/", h.deleteComment)
			}
		}

		return router
	}
}
