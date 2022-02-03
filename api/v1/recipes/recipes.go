package recipes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"cookhub.com/app/models"
)

type AuthorShortItem struct {
	Id string `json:"id"`
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

type UserFeedComponent struct {
	Type int `json:"type"` // 1 - recipe, 2...
	FeedRecipe RecipeUserFeedItem `json:"feed_recipe,omitempty"`
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

	// get last published recipes
	lastPublishedRecipes, err := recipesModel.GetLastPublishedRecipes(queryParams.Limit, queryParams.Offset)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var recipeIds []int
	for _, recipe := range lastPublishedRecipes {
		recipeIds = append(recipeIds, recipe.Id)
	}
	
	// get authors for obtained recipes
	var authors []models.UserEntity
	authors, err = usersModel.GetRecipesAuthors(recipeIds)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var favoriteRecipes []models.RecipeEntity
	// get user favorite recipes
	favoriteRecipes, err = recipesModel.GetUserFavoriteRecipes(queryParams.UserId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// merge recipes with authors and update according to user's favorite ones
	var components []UserFeedComponent
	for _, recipe := range lastPublishedRecipes {
		var component UserFeedComponent
		component.Type = 1

		var recipeAuthor models.UserEntity
		for _, author := range authors {
			if author.Id == recipe.AuthorId {
				recipeAuthor = author
				break
			}
		}

		isFavorite := false
		for _, userRecipe := range favoriteRecipes {
			if userRecipe.Id == recipe.Id {
				isFavorite = true
				break
			}
		}

		component.FeedRecipe = RecipeUserFeedItem {
			Author: AuthorShortItem { Id: recipeAuthor.Id, Name: recipeAuthor.Name, AvatarUrl: recipeAuthor.ImageUrl },
			Recipe: RecipeShortItem { Id: recipe.Id, Title: recipe.Title, ImageUrl: recipe.TitleImageUrl, Rating: recipe.Rating, Calories: recipe.Calories, CookTime: recipe.CookTime },
			IsFavorite: isFavorite,
		}

		components = append(components, component)
	}

	// get recipes according to business logic rules
	return context.JSON(http.StatusOK, components)
}
