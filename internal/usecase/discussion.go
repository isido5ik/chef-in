package usecase

import (
	"errors"

	"github.com/isido5ik/RecipePublishingPlatform/dtos"
)

func (u *usecase) AddComment(userId, postId, parentId int, comment string) error {
	return u.repos.AddComment(userId, postId, parentId, comment)
}

func (u *usecase) GetAllComments(postId int) ([]dtos.Comment, error) {
	return u.repos.GetAllComments(postId)
}

func (u *usecase) UpdateComment(userId, postId, commentId int, newComment dtos.UpdateCommentInput) error {
	err := u.repos.CheckComment(userId, postId, commentId)
	if err != nil {
		return errors.New("Forbidden")
	}
	return u.repos.UpdateComment(userId, postId, commentId, newComment)
}

func (u *usecase) DeleteComment(userId, postId, commentId int) error {
	if err := u.repos.CheckComment(userId, postId, commentId); err != nil {
		return errors.New("Forbidden")
	}
	return u.repos.DeleteComment(userId, postId, commentId)
}
