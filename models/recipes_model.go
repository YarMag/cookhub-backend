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
	AuthorId string
}

type RecipeCompilationEntity struct {
	Id int
	Title string
	Recipes []RecipeEntity
}

type RecipesModel interface {
	GetLastPublishedRecipes(limit int, offset int) ([]RecipeEntity, error)
	GetUserFavoriteRecipes(uid string) ([]RecipeEntity, error)
	GetRecipesCompilations() ([]RecipeCompilationEntity, error)
}

type recipesModelImpl struct {
	database *sql.DB
}

func InitRecipes(db *sql.DB) RecipesModel {
	return recipesModelImpl { database: db }
}

func (m recipesModelImpl) GetLastPublishedRecipes(limit int, offset int) ([]RecipeEntity, error) {
	rows, err := m.database.Query("SELECT r.id, r.title, r.title_image_url, r.cooktime, r.calories, r.rating, r.author_id FROM recipes AS r OFFSET $1 LIMIT $2", offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	var recipeItems []RecipeEntity

	for rows.Next() {
		var recipe RecipeEntity
		var titleImageUrl sql.NullString
		err := rows.Scan(&recipe.Id, &recipe.Title, &titleImageUrl, &recipe.CookTime, 
			&recipe.Calories, &recipe.Rating, &recipe.AuthorId)
		if err != nil {
			return nil, err
		}
		if titleImageUrl.Valid {
			titleImageUrlValue, _ := titleImageUrl.Value()
			recipe.TitleImageUrl = titleImageUrlValue.(string)
		} else {
			recipe.TitleImageUrl = ""
		}
		recipeItems = append(recipeItems, recipe)
	}

	return recipeItems, nil
}

func (m recipesModelImpl) GetUserFavoriteRecipes(uid string) ([]RecipeEntity, error) {
	rows, err := m.database.Query("SELECT r.id, r.title, r.title_image_url, r.cooktime, r.calories, r.rating, r.author_id " + 
								  "FROM recipes AS r JOIN favorite_recipes as fr ON r.id = fr.recipe_id " + 
								  "JOIN users as u ON fr.user_id = u.id WHERE u.id = $1", uid)
	if err != nil {
		return nil, err
	}	
	defer rows.Close()

	var recipeItems []RecipeEntity

	for rows.Next() {
		var recipe RecipeEntity
		var titleImageUrl sql.NullString
		err := rows.Scan(&recipe.Id, &recipe.Title, &titleImageUrl, &recipe.CookTime, 
			&recipe.Calories, &recipe.Rating, &recipe.AuthorId)
		if err != nil {
			return nil, err
		}
		if titleImageUrl.Valid {
			titleImageUrlValue, _ := titleImageUrl.Value()
			recipe.TitleImageUrl, _ = titleImageUrlValue.(string)
		} else {
			recipe.TitleImageUrl = ""
		}
		recipeItems = append(recipeItems, recipe)
	}

	return recipeItems, nil
}

func (m recipesModelImpl) GetRecipesCompilations() ([]RecipeCompilationEntity, error) {
	rows, err := m.database.Query("SELECT * FROM recipe_compilations");

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var compilationItems []RecipeCompilationEntity

	for rows.Next() {
		var recipeCompilation RecipeCompilationEntity
		err := rows.Scan(&recipeCompilation.Id, &recipeCompilation.Title)
		if err != nil {
			return nil, err
		}
		recipes, err := m.getRecipesForCompilation(recipeCompilation.Id)
		if err != nil {
			return nil, err
		}
		recipeCompilation.Recipes = recipes

		compilationItems = append(compilationItems, recipeCompilation)
	}

	return compilationItems, nil
}

func (m recipesModelImpl) getRecipesForCompilation(id int) ([]RecipeEntity, error) {
	rows, err := m.database.Query("SELECT r.id, r.title, r.title_image_url, r.cooktime, r.calories, r.rating, r.author_id " + 
		"FROM recipes AS r JOIN recipe_compilations_recipes AS rcr ON r.id = rcr.id_recipe WHERE rcr.id_compilation = $1", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipeItems []RecipeEntity

	for rows.Next() {
		var recipe RecipeEntity
		var titleImageUrl sql.NullString
		err := rows.Scan(&recipe.Id, &recipe.Title, &titleImageUrl, &recipe.CookTime, 
			&recipe.Calories, &recipe.Rating, &recipe.AuthorId)
		if err != nil {
			return nil, err
		}
		if titleImageUrl.Valid {
			titleImageUrlValue, _ := titleImageUrl.Value()
			recipe.TitleImageUrl, _ = titleImageUrlValue.(string)
		} else {
			recipe.TitleImageUrl = ""
		}
		recipeItems = append(recipeItems, recipe)
	}

	return recipeItems, nil
}
