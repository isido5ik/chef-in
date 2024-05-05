package repository

import (
	"fmt"
	"log"
	"strings"

	"github.com/isido5ik/RecipePublishingPlatform/dtos"
)

func (r *repository) CreateRecipe(Recipe dtos.CreateRecipeInput, userId int) (int, error) {
	var postId int

	query := fmt.Sprintf("INSERT INTO %s (user_id, content) VALUES($1, $2) RETURNING post_id", postsTable)
	row := r.db.QueryRow(query, userId, Recipe.Content)
	if err := row.Scan(&postId); err != nil {
		log.Printf("error from CreateRecipe -> row.Scan(&postId): %v", err)
		return 0, err
	}
	log.Printf("doing request: %s", query)
	return postId, nil
}

func (r *repository) GetStories() ([]dtos.Recipe, error) {
	var stories []dtos.Recipe
	query := fmt.Sprintf("SELECT u.username, p.content, p.comments, p.created_at FROM %s p JOIN %s u ON p.user_id = u.user_id", postsTable, usersTable) //TODO: add pagination and filtration
	err := r.db.Select(&stories, query)
	if err != nil {
		log.Printf("error from GetStories -> r.db.Select(&stories, query): %v", err)
	}
	return stories, err
}

func (r *repository) GetUsersStories(userId int) (string, []dtos.Recipe, error) {
	var username string
	var stories []dtos.Recipe

	getUsernameQuery := fmt.Sprintf("SELECT username FROM %s WHERE user_id = $1", usersTable)
	row := r.db.QueryRow(getUsernameQuery, userId)
	if err := row.Scan(&username); err != nil {
		log.Printf("error from GetUsersStories -> row.Scan(&username): %v", err)
		return "", nil, err
	}

	getStoriesQuery := fmt.Sprintf("SELECT content, comments, created_at FROM %s WHERE user_id = $1", postsTable)

	if err := r.db.Select(&stories, getStoriesQuery, userId); err != nil {
		log.Printf("error from GetUsersStories -> r.db.Select(&stories, getStoriesQuery, userId): %v", err)
		return "", nil, err
	}
	return username, stories, nil
}

func (r *repository) GetRecipe(postId int) (dtos.Recipe, error) {
	var Recipe dtos.Recipe

	query := fmt.Sprintf("SELECT p.user_id, u.username, p.content, p.comments, p.created_at FROM %s p JOIN %s u ON p.user_id = u.user_id WHERE post_id = $1", postsTable, usersTable)
	if err := r.db.Get(&Recipe, query, postId); err != nil {
		log.Printf("error from GetRecipe -> r.db.Get(&Recipe, query, postId): %v", err)
		return Recipe, err
	}

	return Recipe, nil
}

func (r *repository) DeleteRecipe(postId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE post_id = $1", postsTable)
	_, err := r.db.Exec(query, postId)
	if err != nil {
		log.Printf("error from DeleteRecipe -> r.db.Exec(query, postId): %v", err)
		return err
	}

	return nil
}

func (r *repository) UpdateRecipe(postId int, input dtos.UpdateRecipeInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Content != "" {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, input.Content)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE post_id = %d", postsTable, setQuery, postId)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		log.Printf("error from UpdateRecipe -> r.db.Exec(query, args...): %v", err)
	}
	return err
}
