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
	Description string
	AuthorId string
	RecipeType int
}

type RecipeCompilationEntity struct {
	Id int
	Title string
	Recipes []RecipeEntity
}

type IngredientEntity struct {
	Id int
	Name string
}

type UnitEntity struct {
	Id int
	Name string
	Count float32
}

type RecipeIngredientEntity struct {
	Ingredient IngredientEntity
	Unit UnitEntity
}

type StepEntity struct {
	Title string
	Desc string
}

type FoodValueEntity struct {
	Name string
	Value float32
}

type RecipeMediaEntity struct {
	Url string
	Type int
}

type RecipeFullInfoEntity struct {
	Recipe RecipeEntity
	Ingredients []RecipeIngredientEntity
	Steps []StepEntity
	FoodValues []FoodValueEntity
	Medias []RecipeMediaEntity
}

type RecipesModel interface {
	GetFullRecipeInfo(id int) (*RecipeFullInfoEntity, error)
	GetLastPublishedRecipes(limit int, offset int) ([]RecipeEntity, error)
	GetUserFavoriteRecipes(uid string) ([]RecipeEntity, error)
	GetPromoRecipes(limit int, offset int) ([]RecipeEntity, error)
	GetRecipesCompilations() ([]RecipeCompilationEntity, error)
}

type recipesModelImpl struct {
	database *sql.DB
}

func InitRecipes(db *sql.DB) RecipesModel {
	recipesModelImpl := new(recipesModelImpl)
	recipesModelImpl.database = db
	return recipesModelImpl
}

func (m *recipesModelImpl) GetLastPublishedRecipes(limit int, offset int) ([]RecipeEntity, error) {
	rows, err := m.database.Query("SELECT r.id, r.title, r.title_image_url, r.cooktime, r.calories, r.rating, r.author_id, r.description, r.recipe_type FROM recipes AS r OFFSET $1 LIMIT $2", offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	var recipeItems []RecipeEntity

	for rows.Next() {
		var recipe RecipeEntity
		var titleImageUrl sql.NullString
		err := rows.Scan(&recipe.Id, &recipe.Title, &titleImageUrl, &recipe.CookTime, 
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
		recipeItems = append(recipeItems, recipe)
	}

	return recipeItems, nil
}

func (m *recipesModelImpl)GetPromoRecipes(limit int, offset int) ([]RecipeEntity, error) {
	rows, err := m.database.Query("SELECT r.id, r.title, r.title_image_url, r.cooktime, r.calories, r.rating, r.author_id, r.description, r.recipe_type " + 
								  "FROM recipes as r JOIN promo_recipes as pr ON r.id = pr.id_recipe OFFSET $1 LIMIT $2", offset, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipeItems []RecipeEntity

	for rows.Next() {
		var recipe RecipeEntity
		var titleImageUrl sql.NullString
		err := rows.Scan(&recipe.Id, &recipe.Title, &titleImageUrl, &recipe.CookTime, 
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
		recipeItems = append(recipeItems, recipe)
	}

	return recipeItems, nil
}

func (m *recipesModelImpl) GetUserFavoriteRecipes(uid string) ([]RecipeEntity, error) {
	rows, err := m.database.Query("SELECT r.id, r.title, r.title_image_url, r.cooktime, r.calories, r.rating, r.author_id, r.description, r.recipe_type " + 
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
			&recipe.Calories, &recipe.Rating, &recipe.AuthorId, &recipe.Description, &recipe.RecipeType)
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

func (m *recipesModelImpl) GetFullRecipeInfo(id int) (*RecipeFullInfoEntity, error) {
	var recipeEntity RecipeFullInfoEntity

	recipeRows, err := m.database.Query("SELECT * FROM recipes WHERE recipes.id = $1 LIMIT 1", id)

	if err != nil {
		return nil, err
	}

	for recipeRows.Next() {
		var recipe RecipeEntity
		err := recipeRows.Scan(&recipe.Id, &recipe.Title, &recipe.CookTime, &recipe.Calories, 
							   &recipe.Rating, &recipe.TitleImageUrl, &recipe.AuthorId, &recipe.Description, &recipe.RecipeType)
		if err != nil {
			return nil, err
		}
		recipeEntity.Recipe = recipe
	}

	stepsRows, err := m.database.Query("SELECT recipes_steps.description, recipes_steps.title FROM recipes_steps WHERE recipes_steps.recipe_id = $1 ORDER BY recipes_steps.step", id)
	if err != nil {
		return nil, err
	}

	var steps []StepEntity
	for stepsRows.Next() {
		var step StepEntity
		err := stepsRows.Scan(&step.Desc, &step.Title)
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}
	recipeEntity.Steps = steps

	ingredientRows, err := m.database.Query("SELECT i.name, u.name, ri.amount FROM recipes_ingredients AS ri JOIN ingredients AS i ON ri.ingredient_id = i.id JOIN units AS u ON u.id = ri.unit_id WHERE ri.recipe_id = $1", id)
	if err != nil {
		return nil, err
	}

	var ingredients []RecipeIngredientEntity
	for ingredientRows.Next() {
		var ingredient RecipeIngredientEntity
		err := ingredientRows.Scan(&ingredient.Ingredient.Name, &ingredient.Unit.Name, &ingredient.Unit.Count)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}
	recipeEntity.Ingredients = ingredients

	mediasRows, err := m.database.Query("SELECT m.url, m.type FROM recipe_medias AS m WHERE m.recipe_id = $1", id)
	if err != nil {
		return nil, err
	}

	var medias []RecipeMediaEntity
	for mediasRows.Next() {
		var media RecipeMediaEntity
		err := mediasRows.Scan(&media.Url, &media.Type)
		if err != nil {
			return nil, err
		}
		medias = append(medias, media)
	}
	recipeEntity.Medias = medias

	foodValuesRows, err := m.database.Query("SELECT fv.proteins, fv.fats, fv.carbohydrates FROM recipe_food_values AS fv WHERE fv.recipe_id = $1 LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	var foodValues []FoodValueEntity
	for foodValuesRows.Next() {
		var proteinsValue float32
		var fatsValue float32
		var carbohydratesValue float32
		err := foodValuesRows.Scan(&proteinsValue, &fatsValue, &carbohydratesValue)
		if err != nil {
			return nil, err
		}
		foodValues = append(foodValues, FoodValueEntity { Name: "proteins", Value: proteinsValue })
		foodValues = append(foodValues, FoodValueEntity { Name: "fats", Value: fatsValue })
		foodValues = append(foodValues, FoodValueEntity { Name: "carbohydrates", Value: carbohydratesValue })
	}
	recipeEntity.FoodValues = foodValues

	return &recipeEntity, nil
	
}

func (m *recipesModelImpl) GetRecipesCompilations() ([]RecipeCompilationEntity, error) {
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

func (m *recipesModelImpl) getRecipesForCompilation(id int) ([]RecipeEntity, error) {
	rows, err := m.database.Query("SELECT r.id, r.title, r.title_image_url, r.cooktime, r.calories, r.rating, r.author_id, r.description, r.recipe_type " + 
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
			&recipe.Calories, &recipe.Rating, &recipe.AuthorId, &recipe.Description, &recipe.RecipeType)
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
