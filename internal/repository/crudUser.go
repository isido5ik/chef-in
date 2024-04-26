package repository

import (
	"fmt"
	"log"

	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

func (r *repository) CreateUserAsClient(input dtos.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var user_id int
	createUserQuery := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) VALUES ($1, $2, $3) RETURNING user_id", usersTable)
	row := tx.QueryRow(createUserQuery, input.Username, input.Email, input.Password)
	if err := row.Scan(&user_id); err != nil {
		tx.Rollback()
		log.Printf("error with method row.Scan(&user_id): %s", err.Error())
		return 0, err
	}

	var client_id int
	createClientQuery := fmt.Sprintf("INSERT INTO %s (user_id) VALUES ($1) RETURNING client_id", clientTable)
	row = tx.QueryRow(createClientQuery, user_id)
	if err := row.Scan(&client_id); err != nil {
		tx.Rollback()
		log.Printf("error with method row.Scan(&client_id): %s", err.Error())
		return 0, err
	}

	createConnectionQuery := fmt.Sprintf("INSERT INTO %s (user_id, role_id) VALUES ($1, 2)", usersRolesTable)
	_, err = tx.Exec(createConnectionQuery, user_id)
	if err != nil {
		tx.Rollback()
		log.Printf("error while creating connection: %s", err.Error())
		return 0, err
	}
	return client_id, tx.Commit()
}

func (r *repository) GetUser(username, password string) (dtos.User, error) {
	var user dtos.User
	log.Printf("username: %s, \npassword_hash: %s", username, password)
	getUserQuery := fmt.Sprintf("SELECT * FROM %s WHERE username = $1 AND password_hash = $2", usersTable)

	err := r.db.Get(&user, getUserQuery, username, password)

	if err != nil {
		log.Printf("error with method get from db: %s \n", err.Error())
	}
	log.Printf("doing query to sql: %s", getUserQuery)
	return user, err
}

func (r *repository) GetRoles(userId int) ([]string, error) {
	var roles []string

	getRolesQuery := fmt.Sprintf("SELECT r.role_name FROM %s r JOIN %s ur ON ur.role_id = r.role_id WHERE ur.user_id = $1", rolesTable, usersRolesTable)
	if err := r.db.Select(&roles, getRolesQuery, userId); err != nil {
		log.Printf("error in GetRoles, method r.db.Select(&roles, ...): %s", err.Error())
		return nil, err
	}
	return roles, nil
}

func (r *repository) GetRoleId(role string, userId int) (int, error) {
	var role_id int
	table := "t_" + role

	getRoleIdQuery := fmt.Sprintf("SELECT %s_id FROM %s WHERE user_id = $1", role, table)
	err := r.db.Get(&role_id, getRoleIdQuery, userId)

	return role_id, err
}
