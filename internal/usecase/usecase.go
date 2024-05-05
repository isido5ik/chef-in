package usecase

import (
	"github.com/isido5ik/RecipePublishingPlatform/dtos"
	"github.com/isido5ik/RecipePublishingPlatform/internal/repository"
)

type Usecase interface {
	CreateUserAsClient(input dtos.SignUpInput) (int, error)
	GenerateToken(username, password string) (string, []dtos.Roles, error)
	ParseToken(token string) (int, []dtos.Roles, error)

	CreateRecipe(Recipe dtos.CreateRecipeInput, userId int) (int, error)
	GetStories() ([]dtos.Recipe, error)
	GetUsersStories(userId int) (string, []dtos.Recipe, error)
	GetRecipe(postId int) (dtos.Recipe, error)
	DeleteRecipe(postId, userId int, role string) error
	UpdateRecipe(postId, userId int, role string, input dtos.UpdateRecipeInput) error

	AddComment(userId, postId, parentId int, comment string) error
	UpdateComment(userId, postId, commentId int, newComment dtos.UpdateCommentInput) error
	DeleteComment(userId, postId, commentId int) error
	GetAllComments(postId int) ([]dtos.Comment, error)
}

type usecase struct {
	repos repository.Repository
}

func NewUsecase(repos repository.Repository) Usecase {
	return &usecase{
		repos: repos,
	}
}
