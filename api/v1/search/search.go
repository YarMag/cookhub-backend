package search

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"cookhub.com/app/models"
	"cookhub.com/app/api/entities"
	"cookhub.com/app/repositories"
)

type searchRequestParams struct {
	userId string `header:"UUID"`
	text string `json:"text"`
}

type searchResponse struct {
	recipes []entities.RecipeShortItem `json:"recipes"`
}

func GetSearchResults(context echo.Context, searchModel models.SearchModel, recipesModel models.RecipesModel) error {
	queryParams := new(searchRequestParams)
	if err := context.Bind(queryParams); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	searchResult, err := searchModel.GetSearchResults(queryParams.text)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	userRepository := repositories.NewUserRepository(recipesModel)

	recipes := make([]entities.RecipeShortItem, 0)

	for _, recipe := range searchResult.Recipes {
		isFavorite, err := userRepository.IsRecipeFavorite(queryParams.userId, &recipe)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		recipeEntity := entities.RecipeShortItem {
			Id: recipe.Id,
			Title: recipe.Title,
			ImageUrl: recipe.TitleImageUrl,
			Rating: recipe.Rating,
			Calories: recipe.Calories,
			CookTime: recipe.CookTime,
			IsFavorite: isFavorite,
		}
		recipes = append(recipes, recipeEntity)
	}

	return context.JSON(http.StatusOK, searchResponse { recipes: recipes })
}