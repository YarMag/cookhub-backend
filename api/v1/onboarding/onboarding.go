package onboarding

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"cookhub.com/app/models"
)

type Onboarding struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

func GetOnboarding(context echo.Context, model models.OnboardingModel) error {

	items, err := model.All()

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "No onboardings in database!")
	}

	var onboardingItems []Onboarding

	for _, item := range items {
		onboardingItems = append(onboardingItems, Onboarding { 
			Title: item.Title,
			Image: item.Image,
		})
	}

	return context.JSON(http.StatusOK, onboardingItems)
}
