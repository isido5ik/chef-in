package repository

import (
	"fmt"
	"log"

	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

func (r *repository) CreateUser(input dtos.User) (int, error) {

	var user_id int
	createUserQuery := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)

	log.Printf("doing query to sql: %s", createUserQuery)
	row := r.db.QueryRow(createUserQuery, input.Username, input.Email, input.Password)
	if err := row.Scan(&user_id); err != nil {
		return 0, err
	}
	return user_id, nil
}

func (r *repository) GetUser(username, password string) (dtos.User, error) {
	var user dtos.User
	getUserQuery := fmt.Sprintf("SELECT * FROM %s WHERE username = $1 AND password_hash = $2", usersTable)

	err := r.db.Get(&user, getUserQuery, username, password)

	return user, err
}
