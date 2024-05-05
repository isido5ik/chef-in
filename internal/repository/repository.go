package repository

import (
	"github.com/isido5ik/RecipePublishingPlatform/dtos"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateUserAsClient(input dtos.SignUpInput) (int, error)
	GetUser(username, password string) (dtos.User, error)
	GetRoles(userId int) ([]string, error)
	GetRoleId(role string, userId int) (int, error)

	CreateRecipe(Recipe dtos.CreateRecipeInput, userId int) (int, error)
	GetStories() ([]dtos.Recipe, error)
	GetUsersStories(userId int) (string, []dtos.Recipe, error)
	GetRecipe(postId int) (dtos.Recipe, error)
	DeleteRecipe(postId int) error
	UpdateRecipe(postId int, input dtos.UpdateRecipeInput) error

	AddComment(userId, postId, parentId int, comment string) error
	CheckComment(userId, postId, commentId int) error
	UpdateComment(userId, postId, commentId int, newComment dtos.UpdateCommentInput) error
	DeleteComment(userId, postId, commentId int) error
	GetAllComments(postId int) ([]dtos.Comment, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
