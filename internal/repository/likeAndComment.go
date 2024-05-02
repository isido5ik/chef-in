package repository

import (
	"fmt"
	"log"

	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

func (r *repository) CheckLike(userId, postId int) error {

	var like dtos.Like
	checkLikeQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND post_id = $2", likesTable)
	err := r.db.Get(&like, checkLikeQuery, userId, postId)
	return err
}

func (r *repository) CheckComment(userId, postId, commentId int) error {
	var comment dtos.Comment
	checkCommentQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND post_id = $2 AND comment_id = $3", commentsTable)
	err := r.db.Get(&comment, checkCommentQuery, userId, postId, commentId)
	return err
}

func (r *repository) Like(userId, postId int) error {

	err := r.CheckLike(userId, postId)
	if err == nil {
		//already liked
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	incLikeQuery := fmt.Sprintf("UPDATE %s SET likes = likes + 1 WHERE post_id = $1", postsTable)
	_, err = tx.Exec(incLikeQuery, postId)
	if err != nil {
		tx.Rollback()
		return err
	}
	addLikeQuery := fmt.Sprintf("INSERT INTO %s (user_id, post_id) VALUES($1, $2)", likesTable)
	_, err = tx.Exec(addLikeQuery, userId, postId)
	if err != nil {
		tx.Rollback()
		return err
	}

	log.Printf("adding like: %s", addLikeQuery)

	return tx.Commit()

}

func (r *repository) RemoveLike(userId, postId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	decLikeQuery := fmt.Sprintf("UPDATE %s SET likes = likes - 1 WHERE post_id = $1", postsTable)
	_, err = tx.Exec(decLikeQuery, postId)
	if err != nil {
		tx.Rollback()
		return err
	}
	removeLikeQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", likesTable)
	_, err = tx.Exec(removeLikeQuery, userId)
	if err != nil {
		return err
	}
	log.Printf("removing like: %s", removeLikeQuery)
	return tx.Commit()

}

func (r *repository) AddComment(userId, postId, parentId int, comment string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	incCommentQuery := fmt.Sprintf("UPDATE %s SET comments = comments + 1 WHERE post_id = $1", postsTable)
	_, err = tx.Exec(incCommentQuery, postId)
	if err != nil {
		tx.Rollback()
		return err
	}
	if parentId != 0 {
		addCommentQuery := fmt.Sprintf("INSERT INTO %s (comment_text, parent_id, user_id, post_id) VALUES($1, $2, $3, $4)", commentsTable)
		_, err = tx.Exec(addCommentQuery, comment, parentId, userId, postId)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		addCommentQuery := fmt.Sprintf("INSERT INTO %s (comment_text, user_id, post_id) VALUES($1, $2, $3)", commentsTable)
		_, err := tx.Exec(addCommentQuery, comment, userId, postId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *repository) UpdateComment(userId, postId, commentId int, newComment dtos.UpdateCommentInput) error {
	updateCommentQuery := fmt.Sprintf("UPDATE %s SET comment_text = $1 WHERE user_id = $2 AND post_id = $3 AND comment_id = $4", commentsTable)
	_, err := r.db.Exec(updateCommentQuery, newComment.CommentText, userId, postId, commentId)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteComment(userId, postId, commentId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	decCommentQuery := fmt.Sprintf("UPDATE %s SET comments = comments - 1 WHERE post_id = $1 AND user_id = $2", postsTable)
	_, err = r.db.Exec(decCommentQuery, postId, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteCommentQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND post_id = $2 AND comment_id = $3", commentsTable)
	_, err = r.db.Exec(deleteCommentQuery, userId, postId, commentId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
