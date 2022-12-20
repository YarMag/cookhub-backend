package repositories

import (
	"encoding/json"

	"cookhub.com/app/cache"
	"cookhub.com/app/models"
)

const (
	favoriteRecipesCacheKey string = "favorite_recipes"
	paginationOffsetInfoCacheKey string = "pagination_offset"
)

type UserPaginationOffsetInfo struct {
	recipes int `json:"recipes"`
	compilations int `json:"compilations"`
	promos int `json:"promos"`
}

type UserRepository interface {
	IsRecipeFavorite(userId string, recipe *models.RecipeEntity) (bool, error)

	GetUserPaginationInfo(userId string) (*UserPaginationOffsetInfo, error)
	UpdateUserPaginationInfo(userId string, info UserPaginationOffsetInfo)
}

type userRepositoryImpl struct {
	recipesModel models.RecipesModel

	favoriteRecipesIdsCache map[int]bool
}

func NewUserRepository(recipesModel models.RecipesModel) UserRepository {
	repo := new(userRepositoryImpl)
	repo.recipesModel = recipesModel
	repo.favoriteRecipesIdsCache = nil
	return repo
}

func (repo *userRepositoryImpl)IsRecipeFavorite(userId string, recipe *models.RecipeEntity) (bool, error) {
	if repo.favoriteRecipesIdsCache == nil {
		cachedFavoriteIds, err := repo.getFavoriteRecipesFromCache(userId)

		if err == nil {
			repo.favoriteRecipesIdsCache = cachedFavoriteIds
		} else {
			databaseFavoriteIds, err := repo.getFavoritesRecipesFromDatabase(userId)

			if err == nil {
				repo.favoriteRecipesIdsCache = databaseFavoriteIds
				go repo.cacheNewFavoriteIds(repo.getUserCacheId(userId), databaseFavoriteIds)
			} else {
				return false, err
			}
		}
	}

	_, ok := repo.favoriteRecipesIdsCache[recipe.Id]

	return ok, nil
}

func (repo *userRepositoryImpl)GetUserPaginationInfo(userId string) (*UserPaginationOffsetInfo, error) {
	userCacheId := repo.getUserCacheId(userId)
	
	paginationInfoBytes, err := cache.GetBytes(userCacheId, paginationOffsetInfoCacheKey)

	if err != nil {
		return nil, err
	}

	paginationInfo := new(UserPaginationOffsetInfo)
	err = json.Unmarshal(paginationInfoBytes, &paginationInfo)

	if err != nil {
		return nil, err
	}

	return paginationInfo, nil
}

func (repo *userRepositoryImpl)UpdateUserPaginationInfo(userId string, info UserPaginationOffsetInfo) {
	newPaginationInfoBytes, err := json.Marshal(info)

	if err != nil {
		return
	}
	cache.SetBytes(repo.getUserCacheId(userId), paginationOffsetInfoCacheKey, newPaginationInfoBytes)
}

func (repo *userRepositoryImpl)getUserCacheId(userId string) (string) {
	return "user:"+userId
}

func (repo *userRepositoryImpl)getFavoriteRecipesFromCache(userId string) (map[int]bool, error) {
	userCacheId := repo.getUserCacheId(userId)
	cachedFavoritesBytes, err := cache.GetBytes(userCacheId, favoriteRecipesCacheKey)

	if err != nil {
		return nil, err
	}

	cachedFavoritesIds := make(map[int]bool)
	err = json.Unmarshal(cachedFavoritesBytes, &cachedFavoritesIds)

	if err != nil {
		return nil, err
	}
		
	return cachedFavoritesIds, nil
}

func (repo *userRepositoryImpl)cacheNewFavoriteIds(userCacheId string, favoriteIdsMap map[int]bool) {
	newCachedFavoritesBytes, err := json.Marshal(favoriteIdsMap)
		
	if err != nil {
		return
	}
	cache.SetBytes(userCacheId, favoriteRecipesCacheKey, newCachedFavoritesBytes)
}

func (repo *userRepositoryImpl)getFavoritesRecipesFromDatabase(userId string) (map[int]bool, error) {
	databaseRecipes, err := repo.recipesModel.GetUserFavoriteRecipes(userId)

	if err != nil {
		return nil, err
	}

	cachedFavoritesIds := make(map[int]bool)

	for _, recipe := range databaseRecipes {
		cachedFavoritesIds[recipe.Id] = true
	}
		
	return cachedFavoritesIds, nil
}