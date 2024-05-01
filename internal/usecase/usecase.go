package usecase

import (
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/isido5ik/StoryPublishingPlatform/internal/repository"
)

type Usecase interface {
	CreateUserAsClient(input dtos.User) (int, error)
	GenerateToken(username, password string) (string, []dtos.Roles, error)
	ParseToken(token string) (int, []dtos.Roles, error)

	CreateStory(story dtos.Post, userId int) (int, error)
	GetStories() ([]dtos.Story, error)
	GetUsersStories(userId int) (string, []dtos.Story, error)
	GetStory(postId int) (dtos.Story, error)
	DeleteStory(postId, userId int, role string) error
	UpdateStory(postId, userId int, role string, input dtos.UpdateStoryInput) error

	Like(userId, postId int) error
	RemoveLike(userId, postId int) error
	AddComment(userId, postId, parentId int, comment string) error
	UpdateComment(userId, postId, commentId int, newComment dtos.UpdateCommentInput) error
}

type usecase struct {
	repos repository.Repository
}

func NewUsecase(repos repository.Repository) Usecase {
	return &usecase{
		repos: repos,
	}
}
