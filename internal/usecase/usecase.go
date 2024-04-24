package usecase

import (
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/isido5ik/StoryPublishingPlatform/internal/repository"
)

type Usecase interface {
	CreateUserAsClient(input dtos.User) (int, error)
	GenerateToken(username, password string) (string, []dtos.Roles, error)
	ParseToken(token string) (int, []dtos.Roles, error)
}

type usecase struct {
	repos repository.Repository
}

func NewUsecase(repos repository.Repository) Usecase {
	return &usecase{
		repos: repos,
	}
}
