package usecase

import (
	"errors"
	"log"

	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	adminCtx            = "ADMIN"
	clientCtx           = "CLIENT"
	rolesCtx            = "roles"
)

func (u *usecase) CreateStory(story dtos.Post, userId int) (int, error) {
	return u.repos.CreateStory(story, userId)
}

func (u *usecase) GetStories() ([]dtos.Story, error) {
	return u.repos.GetStories()
}

func (u *usecase) GetUsersStories(userId int) (string, []dtos.Story, error) {
	return u.repos.GetUsersStories(userId)
}

func (u *usecase) GetStory(postId int) (dtos.Story, error) {
	return u.repos.GetStory(postId)
}

func (u *usecase) DeleteStory(postId, userId int, role string) error {
	if role == adminCtx {
		log.Printf("the user has role %s", role)
		return u.repos.DeleteStory(postId)
	}
	if role == clientCtx {
		log.Printf("the user has role %s", role)
		story, err := u.repos.GetStory(postId)
		if err != nil {
			return err
		}
		if story.UserId == userId {
			log.Printf("the post belongs to user")
			return u.repos.DeleteStory(postId)
		}
	}
	return errors.New("you do not have a permission")

}

func (u *usecase) UpdateStory(postId, userId int, role string, input dtos.UpdateStoryInput) error {
	if role == adminCtx {
		log.Printf("the user has role %s", role)
		return u.repos.UpdateStory(postId, input)
	}
	if role == clientCtx {
		log.Printf("the user has role %s", role)
		story, err := u.repos.GetStory(postId)
		if err != nil {
			return err
		}
		if story.UserId == userId {
			log.Printf("the post belongs to user")
			return u.repos.UpdateStory(postId, input)
		}
	}
	return errors.New("you do not have a permission")
}
