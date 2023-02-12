package models

import (
	"strings"
	"encoding/json"
	"database/sql"
)

type UserEntity struct {
	Id string
	Name string
	ImageUrl string
}

type UsersModel interface {
	GetRecipesAuthors(recipeIds []int) ([]UserEntity, error)
	GetAuthorForRecipe(recipeId int) (*UserEntity, error)
}

type usersModelImpl struct {
	database *sql.DB
}

func InitUsers(db *sql.DB) UsersModel {
	return usersModelImpl { database: db }
}

func (m usersModelImpl) GetAuthorForRecipe(recipeId int) (*UserEntity, error) {
	authorQuery := "SELECT users.id, users.name, users.image_url from users JOIN recipes ON recipes.author_id = users.id WHERE recipes.id = $1 LIMIT 1"
	rows, err := m.database.Query(authorQuery, recipeId)

	if err != nil {
		return nil, err
	}

	user := UserEntity{}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.ImageUrl)
		if err != nil {
			return nil, err
		}
		break
	}

	return &user, nil
}

func (m usersModelImpl) GetRecipesAuthors(recipeIds []int) ([]UserEntity, error) {
	if len(recipeIds) == 0 {
		return nil, nil
	}

	idsList, _ := json.Marshal(recipeIds)
	idsListString := strings.Trim(string(idsList), "[]")

	recipeAuthorsIdsQuery := "SELECT DISTINCT recipes.author_id FROM recipes WHERE recipes.id IN (" + idsListString + ")"
	rows, err := m.database.Query("SELECT * FROM users WHERE users.id IN (" + recipeAuthorsIdsQuery + ")")
	
	if err != nil {
		return nil, err
	}

	var users []UserEntity

	for rows.Next() {
		var userEntity UserEntity
		err := rows.Scan(&userEntity.Id, &userEntity.Name, &userEntity.ImageUrl)
		if err != nil {
			return nil, err
		}
		users = append(users, userEntity)
	}

	return users, nil
}