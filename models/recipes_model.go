package models

import (
	"database/sql"
)

type RecipeEntity struct {
	Id int
	Title string
	TitleImageUrl string
	CookTime int
	Calories float32
	Rating float32
	AuthorId int
}

type RecipesModel interface {
	GetLastPublishedRecipes(limit int, offset int) ([]RecipeEntity, error)
	GetUserFavoriteRecipes(uid string) ([]RecipeEntity, error)
}

type recipesModelImpl struct {
	database *sql.DB
}

func InitRecipes(db *sql.DB) RecipesModel {
	return recipesModelImpl { database: db }
}

func (m recipesModelImpl) GetLastPublishedRecipes(limit int, offset int) ([]RecipeEntity, error) {
	rows, err := m.database.Query("SELECT * FROM Recipes LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	var recipeItems []RecipeEntity

	for rows.Next() {
		var recipe RecipeEntity
		err := rows.Scan(&recipe.Id, &recipe.Title, &recipe.TitleImageUrl, &recipe.CookTime, 
			&recipe.Calories, &recipe.Rating, &recipe.AuthorId)
		if err != nil {
			return nil, err
		}
		recipeItems = append(recipeItems, recipe)
	}

	return recipeItems, nil
}

func (m recipesModelImpl) GetUserFavoriteRecipes(uid string) ([]RecipeEntity, error) {
	rows, err := m.database.Query("SELECT r.id, r.title, r.title_image_url, r.cooktime, r.calories, r.rating, r.author_id " + 
								  "FROM recipe AS r JOIN favorite_recipes as fr ON r.id = fr.recipe_id " + 
								  "JOIN user as u ON fr.author_id = u.id WHERE u.id = ?", uid)
	if err != nil {
		return nil, err
	}	
	defer rows.Close()

	var recipeItems []RecipeEntity

	for rows.Next() {
		var recipe RecipeEntity
		err := rows.Scan(&recipe.Id, &recipe.Title, &recipe.TitleImageUrl, &recipe.CookTime, 
			&recipe.Calories, &recipe.Rating, &recipe.AuthorId)
		if err != nil {
			return nil, err
		}
		recipeItems = append(recipeItems, recipe)
	}

	return recipeItems, nil
}
