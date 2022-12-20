package recipes

import (
	"cookhub.com/app/models"
	"cookhub.com/app/api/entities"
	"cookhub.com/app/repositories"
)

type recipesCombiner struct {
	recipesModel models.RecipesModel
	usersModel models.UsersModel

	userRepository repositories.UserRepository
}

func newRecipesCombiner(recipesModel models.RecipesModel, usersModel models.UsersModel) *recipesCombiner {
	rc := new(recipesCombiner)
	rc.recipesModel = recipesModel
	rc.usersModel = usersModel
	rc.userRepository = repositories.NewUserRepository(recipesModel)
	return rc
}

func (rc *recipesCombiner)getComponents(limit int, offset int, userId string) ([]UserFeedComponent, error) {
	recipeComponents, err := rc.getRecipesComponents(userId, limit, offset)

	if err != nil {
		return nil, err
	}

	compilationComponents, err := rc.getCompilationsComponents(userId, limit, offset)

	if err != nil {
		return nil, err
	}

	promoComponents, err := rc.getPromoComponents(userId, limit, offset)

	if err != nil {
		return nil, err
	}

	components := []UserFeedComponent{}
	components = append(components, recipeComponents[0:1]...)
	components = append(components, compilationComponents...)
	components = append(components, promoComponents...)
	components = append(components, recipeComponents[2:]...)

	return components, nil
}

func (rc *recipesCombiner)getRecipesComponents(userId string, limit int, offset int) ([]UserFeedComponent, error) {
	// get last published recipes
	lastPublishedRecipes, err := rc.recipesModel.GetLastPublishedRecipes(limit, offset)

	if err != nil {
		return nil, err
	}

	var recipeIds []int
	for _, recipe := range lastPublishedRecipes {
		recipeIds = append(recipeIds, recipe.Id)
	}
	
	// get authors for obtained recipes
	var authors []models.UserEntity
	authors, err = rc.usersModel.GetRecipesAuthors(recipeIds)

	if err != nil {
		return nil, err
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


		isFavorite, _ := rc.userRepository.IsRecipeFavorite(userId, &recipe)

		component.FeedRecipe = entities.RecipeUserFeedItem {
			Author: rc.mapAuthorShortItem(recipeAuthor),
			Recipe: rc.mapRecipeShortItem(recipe, isFavorite),
		}

		components = append(components, component)
	}

	return components, nil
}

func (rc *recipesCombiner)getCompilationsComponents(userId string, limit int, offset int) ([]UserFeedComponent, error) {
	compilations, err := rc.recipesModel.GetRecipesCompilations()
	if err != nil {
		return nil, err
	}

	var components []UserFeedComponent
	for _, compilation := range compilations {
		var component UserFeedComponent
		component.Type = 2

		var recipeItems []entities.RecipeShortItem
		for _, compilationRecipe := range compilation.Recipes {
			isFavorite, _ := rc.userRepository.IsRecipeFavorite(userId, &compilationRecipe)
			recipeItems = append(recipeItems, rc.mapRecipeShortItem(compilationRecipe, isFavorite))
		}

		component.FeedCompilation = entities.RecipeCompilationUserFeedItem {
			Id: compilation.Id,
			Title: compilation.Title,
			Recipes: recipeItems,
		}
		components = append(components, component)
	}

	return components, nil
}

func (rc *recipesCombiner)getPromoComponents(userId string, limit int, offset int) ([]UserFeedComponent, error) {
	// TODO: code smell, need to refactor with getRecipesComponents method
	// it almost duplicates this method, but with minor differences for promo (data source, id)
	promos, err := rc.recipesModel.GetPromoRecipes(limit, offset)
	if err != nil {
		return nil, err
	}

	var recipeIds []int
	for _, recipe := range promos {
		recipeIds = append(recipeIds, recipe.Id)
	}
	
	// get authors for obtained recipes
	var authors []models.UserEntity
	authors, err = rc.usersModel.GetRecipesAuthors(recipeIds)

	if err != nil {
		return nil, err
	}

	var components []UserFeedComponent
	for _, promoRecipe := range promos {
		var component UserFeedComponent
		component.Type = 3

		var recipeAuthor models.UserEntity
		for _, author := range authors {
			if author.Id == promoRecipe.AuthorId {
				recipeAuthor = author
				break
			}
		}

		isFavorite, _ := rc.userRepository.IsRecipeFavorite(userId, &promoRecipe)

		component.FeedRecipe = entities.RecipeUserFeedItem {
			Author: rc.mapAuthorShortItem(recipeAuthor),
			Recipe: rc.mapRecipeShortItem(promoRecipe, isFavorite),
		}

		components = append(components, component)
	}

	return components, nil
}

func (rc *recipesCombiner)mapRecipeShortItem(recipe models.RecipeEntity, isFavorite bool) entities.RecipeShortItem {
	return entities.RecipeShortItem {
		Id: recipe.Id,
		Title: recipe.Title,
		ImageUrl: recipe.TitleImageUrl,
		Rating: recipe.Rating,
		Calories: recipe.Calories,
		CookTime: recipe.CookTime,
		IsFavorite: isFavorite,
	}
}

func (rc *recipesCombiner)mapAuthorShortItem(author models.UserEntity) entities.AuthorShortItem {
	return entities.AuthorShortItem {
		Id: author.Id,
		Name: author.Name,
		AvatarUrl: author.ImageUrl,
	}
}
