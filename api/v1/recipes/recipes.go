package recipes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"cookhub.com/app/models"
)

type AuthorShortItem struct {
	Id int `json:"id"`
	Name string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

type RecipeShortItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	ImageUrl string `json:"image_url"`
	Rating float32 `json:"rating"`
	CookTime int `json:"cook_time"`
	Calories float32 `json:"calories"`
}

type RecipeUserFeedItem struct {
	Author AuthorShortItem `json:"author"`
	Recipe RecipeShortItem `json:"recipe"`
	IsFavorite bool `json:"is_favorite"`
}

type recipeUserFeedRequestParams struct {
	UserId string `header:"UUID"`
	Offset int `query:"offset"`
	Limit int `query:"limit"`
}

func GetUserFeedRecipes(context echo.Context, model models.RecipesModel) error {

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

	// get last published recipes

	// get authors for obtained recipes

	// get user favorite recipes

	// merge recipes with authors and update according to user's favorite ones

	// get recipes according to business logic rules
	return echo.NewHTTPError(http.StatusNotImplemented, "Not implemented yet!")
}
