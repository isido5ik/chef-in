package usecase

import (
	"errors"
	"log"

	"github.com/isido5ik/RecipePublishingPlatform/dtos"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	adminCtx            = "ADMIN"
	clientCtx           = "CLIENT"
	rolesCtx            = "roles"
)

func (u *usecase) CreateRecipe(Recipe dtos.CreateRecipeInput, userId int) (int, error) {
	return u.repos.CreateRecipe(Recipe, userId)
}

func (u *usecase) GetStories() ([]dtos.Recipe, error) {
	return u.repos.GetStories()
}

func (u *usecase) GetUsersStories(userId int) (string, []dtos.Recipe, error) {
	return u.repos.GetUsersStories(userId)
}

func (u *usecase) GetRecipe(postId int) (dtos.Recipe, error) {
	return u.repos.GetRecipe(postId)
}

func (u *usecase) DeleteRecipe(postId, userId int, role string) error {
	if role == adminCtx {
		log.Printf("the user has role %s", role)
		return u.repos.DeleteRecipe(postId)
	}
	if role == clientCtx {
		log.Printf("the user has role %s", role)
		Recipe, err := u.repos.GetRecipe(postId)
		if err != nil {
			return err
		}
		if Recipe.UserId == userId {
			log.Printf("the post belongs to user")
			return u.repos.DeleteRecipe(postId)
		}
	}
	return errors.New("you do not have a permission")

}

func (u *usecase) UpdateRecipe(postId, userId int, role string, input dtos.UpdateRecipeInput) error {
	if role == adminCtx {
		log.Printf("the user has role %s", role)
		return u.repos.UpdateRecipe(postId, input)
	}
	if role == clientCtx {
		log.Printf("the user has role %s", role)
		Recipe, err := u.repos.GetRecipe(postId)
		if err != nil {
			return err
		}
		if Recipe.UserId == userId {
			log.Printf("the post belongs to user")
			return u.repos.UpdateRecipe(postId, input)
		}
	}
	return errors.New("you do not have a permission")
}
