package repository

import (
	"fmt"
	"log"

	"github.com/isido5ik/RecipePublishingPlatform/dtos"
)

func (r *repository) CheckComment(userId, postId, commentId int) error {
	var comment dtos.Comment
	checkCommentQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND post_id = $2 AND comment_id = $3", commentsTable)
	err := r.db.Get(&comment, checkCommentQuery, userId, postId, commentId)
	if err != nil {
		log.Print("error from CheckComment -> r.db.Get(&comment, checkCommentQuery, userId, postId, commentId)")
	}
	return err
}

func (r *repository) AddComment(userId, postId, parentId int, comment string) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Print("error from AddComment -> r.db.Begin()")
		return err
	}

	incCommentQuery := fmt.Sprintf("UPDATE %s SET comments = comments + 1 WHERE post_id = $1", postsTable)
	_, err = tx.Exec(incCommentQuery, postId)
	if err != nil {
		log.Print("error from AddComment -> tx.Exec(incCommentQuery, postId)")
		tx.Rollback()
		return err
	}
	if parentId != 0 {
		addCommentQuery := fmt.Sprintf("INSERT INTO %s (comment_text, parent_id, user_id, post_id) VALUES($1, $2, $3, $4)", commentsTable)
		_, err = tx.Exec(addCommentQuery, comment, parentId, userId, postId)
		if err != nil {
			log.Print("error from AddComment -> tx.Exec(addCommentQuery, comment, parentId, userId, postId)")
			tx.Rollback()
			return err
		}
	} else {
		addCommentQuery := fmt.Sprintf("INSERT INTO %s (comment_text, user_id, post_id) VALUES($1, $2, $3)", commentsTable)
		_, err := tx.Exec(addCommentQuery, comment, userId, postId)
		if err != nil {
			log.Print("error from AddComment -> tx.Exec(addCommentQuery, comment, userId, postId)")
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *repository) GetAllComments(postId int) ([]dtos.Comment, error) {
	var comments []dtos.Comment

	getCommentsQuery := fmt.Sprintf("SELECT * FROM %s WHERE post_id = $1", commentsTable)
	if err := r.db.Select(&comments, getCommentsQuery, postId); err != nil {
		log.Print("error from GetAllComments -> r.db.Select(&comments, getCommentsQuery, postId)")
		return nil, err
	}
	return comments, nil
}

func (r *repository) UpdateComment(userId, postId, commentId int, newComment dtos.UpdateCommentInput) error {
	updateCommentQuery := fmt.Sprintf("UPDATE %s SET comment_text = $1 WHERE user_id = $2 AND post_id = $3 AND comment_id = $4", commentsTable)
	_, err := r.db.Exec(updateCommentQuery, newComment.CommentText, userId, postId, commentId)
	if err != nil {
		log.Print("error from UpdateComment -> r.db.Exec(updateCommentQuery, newComment.CommentText, userId, postId, commentId)")
		return err
	}
	return nil
}

func (r *repository) DeleteComment(userId, postId, commentId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Print("error from DeleteComment -> r.db.Begin()")
		return err
	}

	decCommentQuery := fmt.Sprintf("UPDATE %s SET comments = comments - 1 WHERE post_id = $1 AND user_id = $2", postsTable)
	_, err = r.db.Exec(decCommentQuery, postId, userId)
	if err != nil {
		log.Print("error from DeleteComment -> r.db.Exec(decCommentQuery, postId, userId)")
		tx.Rollback()
		return err
	}

	deleteCommentQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND post_id = $2 AND comment_id = $3", commentsTable)
	_, err = r.db.Exec(deleteCommentQuery, userId, postId, commentId)
	if err != nil {
		log.Print("error from DeleteComment -> r.db.Exec(deleteCommentQuery, userId, postId, commentId)")
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
