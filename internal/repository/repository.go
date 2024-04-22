package repository

import (
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateUser(input dtos.User) (int, error)
	GetUser(username, password string) (dtos.User, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
