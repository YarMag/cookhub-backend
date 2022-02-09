package models

import (
	"database/sql"
)

type OnboardingEntity struct {
	Id int
	Title string
	Image string
}

type OnboardingModel interface {
	All() ([]OnboardingEntity, error)
}

type onboardingModelImpl struct {
	database *sql.DB
}

func InitOnboarding(db *sql.DB) OnboardingModel {
	return onboardingModelImpl { database: db }
} 

func (m onboardingModelImpl) All() ([]OnboardingEntity, error) {
	rows, err := m.database.Query("SELECT * FROM Onboardings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var onboardingItems []OnboardingEntity

	for rows.Next() {
		var onboarding OnboardingEntity
		err := rows.Scan(&onboarding.Id, &onboarding.Title, &onboarding.Image)
		if err != nil {
			return nil, err
		}

		onboardingItems = append(onboardingItems, onboarding)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return onboardingItems, nil
}