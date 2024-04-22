package usecase

import (
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/isido5ik/StoryPublishingPlatform/internal/repository"
)

type Usecase interface {
	CreateUser(input dtos.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type usecase struct {
	repos repository.Repository
}

func NewUsecase(repos repository.Repository) Usecase {
	return &usecase{
		repos: repos,
	}
}
