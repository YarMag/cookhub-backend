package recipes  

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"cookhub.com/app/models"
	"cookhub.com/app/api/entities"
)

type UserFeedComponent struct {
	Type int `json:"type"` // 1 - recipe, 2...
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
