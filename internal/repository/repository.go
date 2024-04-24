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
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
