package island_algorithm

import (
	"fmt"
	"database/sql"
	. "ga_operators"
	"cookhub.com/app/cache"
)

type RecipeInfoExtractor struct {
	Database *sql.DB
}

func (extractor *RecipeInfoExtractor)GetChromosomeRequirements() ([]int, *ChromosomeStructureRequirements, error) {
	recipesRows, err := extractor.Database.Query("SELECT recipes.id, recipes.recipe_type FROM recipes ORDER BY recipes.recipe_type")
	if err != nil {
		return nil, nil, err
	}
	var breakfastRecipesIds []int
	var lunchRecipesIds []int
	var dinnerRecipesIds []int
	var snackRecipesIds []int
	for recipesRows.Next() {
		var curId int
		var curType int
		err := recipesRows.Scan(&curId, &curType)
		if (err != nil) {
			return nil, nil, err
		}
		switch curType {
		case 1:
			breakfastRecipesIds = append(breakfastRecipesIds, curId)
		case 2:
			lunchRecipesIds = append(lunchRecipesIds, curId)
		case 3:
			dinnerRecipesIds = append(dinnerRecipesIds, curId)
		case 4:
			snackRecipesIds = append(snackRecipesIds, curId)
		default:
			break
		}
	}

	partsBoundsIndices := make([]int, 4)
	partsBoundsIndices[0] = len(breakfastRecipesIds) - 1
	partsBoundsIndices[1] = partsBoundsIndices[0] + len(lunchRecipesIds)
	partsBoundsIndices[2] = partsBoundsIndices[1] + len(dinnerRecipesIds)
	partsBoundsIndices[3] = partsBoundsIndices[2] + len(snackRecipesIds)

	chromosomeRequirements := ChromosomeStructureRequirements {
		MaxRecipesInPartsCount: 3,
		PartsBoundsIndices: partsBoundsIndices,
	}
	totalRecipesCount := len(breakfastRecipesIds) + len(lunchRecipesIds) + len(dinnerRecipesIds) + len(snackRecipesIds)
	allRecipesIds := make([]int, totalRecipesCount)
	copy(allRecipesIds[0:partsBoundsIndices[0]+1], breakfastRecipesIds[:])
	copy(allRecipesIds[partsBoundsIndices[0]+1:partsBoundsIndices[1]+1], lunchRecipesIds[:])
	copy(allRecipesIds[partsBoundsIndices[1]+1:partsBoundsIndices[2]+1], dinnerRecipesIds[:])
	copy(allRecipesIds[partsBoundsIndices[2]+1:], snackRecipesIds[:])
	
	return allRecipesIds, &chromosomeRequirements, nil
}

func (extractor *RecipeInfoExtractor)GetRecipePrice(id int) (int, error) {
	cacheObject := fmt.Sprintf("recipe_%d", id)
	cachedValue, err := cache.GetInt64(cacheObject, "price")
	if err == nil {
		return int(cachedValue), nil
	}


	recipesRows, err := extractor.Database.Query("SELECT recipe_food_values.price FROM recipe_food_values WHERE recipe_food_values.recipe_id = $1", id)
	if err != nil {
		return 0, err
	}
	var value int
	for recipesRows.Next() {
		err := recipesRows.Scan(&value)
		if err != nil {
			return 0, err
		}
	}

	cache.SetInt64(cacheObject, "price", int64(value))
	return value, nil
}

func (extractor *RecipeInfoExtractor)GetRecipeCookingTime(id int) (int, error) {
	cacheObject := fmt.Sprintf("recipe_%d", id)
	cachedValue, err := cache.GetInt64(cacheObject, "cooktime")
	if err == nil {
		return int(cachedValue), nil
	}

	recipesRows, err := extractor.Database.Query("SELECT recipes.cooktime FROM recipes WHERE rcipes.id = $1", id)
	if err != nil {
		return 0, err
	}
	var value int
	for recipesRows.Next() {
		err := recipesRows.Scan(&value)
		if err != nil {
			return 0, err
		}
	}

	cache.SetInt64(cacheObject, "cooktime", int64(value))
	return value, nil
}

func (extractor *RecipeInfoExtractor)GetRecipeProteins(id int) (float32, error) {
	cacheObject := fmt.Sprintf("recipe_%d", id)
	cachedValue, err := cache.GetFloat64(cacheObject, "proteins")
	if err == nil {
		return float32(cachedValue), nil
	}

	recipesRows, err := extractor.Database.Query("SELECT recipe_food_values.proteins FROM recipe_food_values WHERE recipe_food_values.recipe_id = $1", id)
	if err != nil {
		return 0, err
	}
	var value float32
	for recipesRows.Next() {
		err := recipesRows.Scan(&value)
		if err != nil {
			return 0, err
		}
	}

	cache.SetFloat64(cacheObject, "proteins", float64(value))
	return value, nil
}
