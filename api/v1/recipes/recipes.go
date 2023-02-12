package recipes  

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"cookhub.com/app/models"
	"cookhub.com/app/api/entities"
)

type UserFeedComponent struct {
	Type int `json:"type"` // 1 - recipe, 2 - carousel, 3 - promo
	FeedRecipe entities.RecipeUserFeedItem `json:"feed_recipe,omitempty"`
	FeedCompilation entities.RecipeCompilationUserFeedItem `json:"feed_compilation,omitempty"`
}

type UserFeedResponse struct {
	Components []UserFeedComponent `json:"components"`
}

type recipeUserFeedRequestParams struct {
	UserId string `header:"UUID"`
	Offset int `query:"offset"`
	Limit int `query:"limit"`
}

func GetUserFeedRecipes(context echo.Context, recipesModel models.RecipesModel, usersModel models.UsersModel) error {

	queryParams := new(recipeUserFeedRequestParams)
	if err := context.Bind(queryParams); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validation
	if queryParams.Limit < 0 {
		queryParams.Limit = 0
	} else if queryParams.Limit > 10 {
		queryParams.Limit = 10
	}

	if queryParams.Offset < 0 {
		queryParams.Offset = 0
	} else if queryParams.Offset > 1000 {
		queryParams.Offset = 1000
	}

	recipesCombiner := newRecipesCombiner(recipesModel, usersModel)
	components, err := recipesCombiner.getComponents(queryParams.Limit, queryParams.Offset, queryParams.UserId)
	
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// get recipes according to business logic rules
	return context.JSON(http.StatusOK, UserFeedResponse { Components: components })
}

// ============================

type recipeRequestParams struct {
	Id int `query:"id"`
}

type recipeResponse struct {
	Recipe entities.RecipeFullItem `json:"recipe"`
	Author entities.AuthorShortItem `json:"author"`
}

func mapToFullRecipe(recipe *models.RecipeFullInfoEntity) entities.RecipeFullItem {
	var urls []entities.MediaItem 
	for _, media := range recipe.Medias {
		urls = append(urls, entities.MediaItem { 
			Type: media.Type,
			Url: media.Url,
		})
	}

	var ingredientItems []entities.IngredientItem
	for _, ingr := range recipe.Ingredients {
		ingredientItems = append(ingredientItems, entities.IngredientItem {
			Name: ingr.Ingredient.Name,
			Count: ingr.Unit.Count,
			Units: ingr.Unit.Name,
		})
	}

	var stepItems []entities.StepItem
	for _, step := range recipe.Steps {
		stepItems = append(stepItems, entities.StepItem {
			Desc: step.Desc,
		})
	}

	return entities.RecipeFullItem {
		Id: recipe.Recipe.Id,
		Title: recipe.Recipe.Title,
		Rating: recipe.Recipe.Rating,
		CookTime: recipe.Recipe.CookTime,
		Calories: recipe.Recipe.Calories,
		IsFavorite: false, // TODO: method is hidden in recipesCombiner, need to extract
		MediaUrls: urls,
		Ingredients: ingredientItems,
		Steps: stepItems,
	}
}

func mapAuthor(author *models.UserEntity) entities.AuthorShortItem {
	return entities.AuthorShortItem {
		Id: author.Id,
		Name: author.Name,
		AvatarUrl: author.ImageUrl,
	}
}

func GetRecipe(context echo.Context, recipesModel models.RecipesModel, usersModel models.UsersModel) error {
	queryParams := new(recipeRequestParams)
	if err := context.Bind(queryParams); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	recipe, err := recipesModel.GetFullRecipeInfo(queryParams.Id)
	
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	author, err := usersModel.GetAuthorForRecipe(recipe.Recipe.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, recipeResponse { 
		Recipe: mapToFullRecipe(recipe), 
		Author: mapAuthor(author), 
	})
}


