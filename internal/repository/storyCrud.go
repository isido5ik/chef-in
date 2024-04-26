package repository

import (
	"fmt"
	"log"
	"strings"

	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

func (r *repository) CreateStory(story dtos.Post, userId int) (int, error) {
	var postId int

	query := fmt.Sprintf("INSERT INTO %s (user_id, content) VALUES($1, $2) RETURNING post_id", postsTable)
	row := r.db.QueryRow(query, userId, story.Content)
	if err := row.Scan(&postId); err != nil {
		return 0, err
	}
	log.Printf("doint request: %s", query)
	return postId, nil
}

func (r *repository) GetStories() ([]dtos.Story, error) {
	var stories []dtos.Story
	query := fmt.Sprintf("SELECT u.username, p.content, p.created_at FROM %s p JOIN %s u ON p.user_id = u.user_id", postsTable, usersTable) //TODO: add pagination and filtration
	err := r.db.Select(&stories, query)
	return stories, err
}

func (r *repository) GetUsersStories(userId int) (string, []dtos.Story, error) {
	var username string
	var stories []dtos.Story

	getUsernameQuery := fmt.Sprintf("SELECT username FROM %s WHERE user_id = $1", usersTable)
	row := r.db.QueryRow(getUsernameQuery, userId)
	if err := row.Scan(&username); err != nil {
		return "", nil, err
	}

	getStoriesQuery := fmt.Sprintf("SELECT content, created_at FROM %s WHERE user_id = $1", postsTable)

	if err := r.db.Select(&stories, getStoriesQuery, userId); err != nil {
		return "", nil, err
	}
	return username, stories, nil
}

func (r *repository) GetStory(postId int) (dtos.Story, error) {
	var story dtos.Story

	query := fmt.Sprintf("SELECT p.user_id, u.username, p.content, p.created_at FROM %s p JOIN %s u ON p.user_id = u.user_id WHERE post_id = $1", postsTable, usersTable)
	if err := r.db.Get(&story, query, postId); err != nil {
		return story, err
	}

	return story, nil
}

func (r *repository) DeleteStory(postId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE post_id = $1", postsTable)
	_, err := r.db.Exec(query, postId)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateStory(postId int, input dtos.UpdateStoryInput) error {
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
	return err
}
