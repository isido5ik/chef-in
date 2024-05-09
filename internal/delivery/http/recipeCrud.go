package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isido5ik/RecipePublishingPlatform/dtos"
)

// @Summary Create Recipe
// @Security ApiKeyAuth
// @Tags recipes
// @Description create recipe
// @ID create-recipe
// @Accept  json
// @Produce  json
// @Param input body dtos.CreateRecipeInput true "recipe info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/recipes [post]
func (h *Handler) createRecipe(c *gin.Context) {
	var Recipe dtos.CreateRecipeInput
	if err := c.BindJSON(&Recipe); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	postId, err := h.useCases.CreateRecipe(Recipe, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"postId": postId,
	})
}

// @Summary Get All Recipes
// @Tags resipes
// @Description get all recipes
// @ID get-all-recipes
// @Accept  json
// @Produce  json
// @Success 200 {object} getRecipesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/recipes [get]
func (h *Handler) getRecipes(c *gin.Context) {

	stories, err := h.useCases.GetStories()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getRecipesResponse{
		Data: stories,
	})
}

// @Summary Get My Recipe
// @Security ApiKeyAuth
// @Tags recipes
// @Description get users recipes
// @ID get-my-recipe
// @Accept  json
// @Produce  json
// @Success 200 {object} getMyRecipesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/recipes/my [get]
func (h *Handler) getUsersRecipes(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	username, recipes, err := h.useCases.GetUsersStories(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getMyRecipesResponse{
		Username:  username,
		MyRecipes: recipes,
	})
}

// @Summary Get Recipe By Id
// @Security ApiKeyAuth
// @Tags recipes
// @Description get recipe by id
// @ID get-recipe-by-id
// @Param :recipe_id path int true "Идентификатор рецепта"
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.Recipe
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/recipes/{:recipe_id} [get]
func (h *Handler) getRecipe(c *gin.Context) {

	log.Print(c.Param("recipe_id"))
	postId, err := strconv.Atoi(c.Param("recipe_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	Recipe, err := h.useCases.GetRecipe(postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Recipe:": Recipe,
	})
}

// @Summary Update Recipe By Id
// @Security ApiKeyAuth
// @Tags recipes
// @Description update recipe by id
// @ID update-recipe-by-id
// @Param :recipe_id path int true "Post ID"
// @Accept  json
// @Produce  json
// @Param input body dtos.UpdateRecipeInput true "account info"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/recipes/{:recipe_id} [put]
func (h *Handler) updateRecipe(c *gin.Context) {
	var input dtos.UpdateRecipeInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	postId, err := strconv.Atoi(c.Param("recipe_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	role, err := getRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	err = h.useCases.UpdateRecipe(postId, userId, role, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}

// @Summary Delete Recipe By Id
// @Security ApiKeyAuth
// @Tags recipes
// @Description delete recipe by id
// @ID delete-recipe-by-id
// @Param :recipe_id path int true "Post ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/recipes/{:recipe_id} [delete]
func (h *Handler) deleteRecipe(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("recipe_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	role, err := getRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	err = h.useCases.DeleteRecipe(postId, userId, role)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})

}
