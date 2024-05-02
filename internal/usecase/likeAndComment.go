package usecase

import (
	"errors"

	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

func (u *usecase) Like(userId, postId int) error {
	return u.repos.Like(userId, postId)
}

func (u *usecase) RemoveLike(userId, postId int) error {
	err := u.repos.CheckLike(userId, postId)
	if err != nil {
		return errors.New("Forbidden")
	}
	return u.repos.RemoveLike(userId, postId)
}

func (u *usecase) AddComment(userId, postId, parentId int, comment string) error {
	return u.repos.AddComment(userId, postId, parentId, comment)
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
