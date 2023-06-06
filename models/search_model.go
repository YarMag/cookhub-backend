package models

import (
	"database/sql"
)

type SearchResult struct {
	Recipes []RecipeEntity
}

type SearchModel interface {
	GetSearchResults(string) (*SearchResult, error)
}

type searchModelImpl struct {
	database *sql.DB
}

func NewSearchModel(db *sql.DB) SearchModel {
	searchModel := new(searchModelImpl)
	searchModel.database = db
	return searchModel
}

func (model *searchModelImpl)GetSearchResults(text string) (*SearchResult, error) {
	recipeRows, err := model.database.Query("SELECT * from recipes WHERE recipes.title LIKE '%' || ? || '%' OR recipes.description LIKE '%' || ? || '%'", text, text)

	if err != nil {
		return nil, err
	}

	foundRecipes := make([]RecipeEntity, 0)

	for recipeRows.Next() {
		var recipe RecipeEntity
		var titleImageUrl sql.NullString

		err := recipeRows.Scan(&recipe.Id, &recipe.Title, &titleImageUrl, &recipe.CookTime, 
			&recipe.Calories, &recipe.Rating, &recipe.AuthorId, &recipe.Description, &recipe.RecipeType)

		if err != nil {
			return nil, err
		}

		if titleImageUrl.Valid {
			titleImageUrlValue, _ := titleImageUrl.Value()
			recipe.TitleImageUrl = titleImageUrlValue.(string)
		} else {
			recipe.TitleImageUrl = ""
		}

		foundRecipes = append(foundRecipes, recipe)
	}

	return &SearchResult {
		Recipes: foundRecipes,
	}, nil
}