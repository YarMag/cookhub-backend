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
}

type usersModelImpl struct {
	database *sql.DB
}

func InitUsers(db *sql.DB) UsersModel {
	return usersModelImpl { database: db }
}

func (m usersModelImpl) GetRecipesAuthors(recipeIds []int) ([]UserEntity, error) {
	if len(recipeIds) == 0 {
		return nil, nil
	}

	idsList, _ := json.Marshal(recipeIds)
	idsListString := strings.Trim(string(idsList), "[]")

	recipeAuthorsIdsQuery := "SELECT DISTINCT recipe.author_id FROM recipe WHERE recipe.id IN (" + idsListString + ")"
	rows, err := m.database.Query("SELECT * FROM user WHERE user.id IN (" + recipeAuthorsIdsQuery + ")")
	
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