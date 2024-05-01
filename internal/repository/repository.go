package repository

import (
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateUserAsClient(input dtos.User) (int, error)
	GetUser(username, password string) (dtos.User, error)
	GetRoles(userId int) ([]string, error)
	GetRoleId(role string, userId int) (int, error)

	CreateStory(story dtos.Post, userId int) (int, error)
	GetStories() ([]dtos.Story, error)
	GetUsersStories(userId int) (string, []dtos.Story, error)
	GetStory(postId int) (dtos.Story, error)
	DeleteStory(postId int) error
	UpdateStory(postId int, input dtos.UpdateStoryInput) error

	Like(userId, postId int) error
	CheckLike(userId, postId int) error
	RemoveLike(userId, postId int) error
	AddComment(userId, postId, parentId int, comment string) error
	CheckComment(userId, postId, commentId int) error
	UpdateComment(userId, postId, commentId int, newComment dtos.UpdateCommentInput) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
