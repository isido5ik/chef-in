package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	usersTable      = "t_users"
	adminTable      = "t_admin"
	clientTable     = "t_client"
	rolesTable      = "t_roles"
	usersRolesTable = "t_users_roles"
	postsTable      = "t_posts"
	commentsTable   = "t_comments"
	likesTable      = "t_likes"
)

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	source := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	log.Print(source)
	db, err := sqlx.Open("postgres", source)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
