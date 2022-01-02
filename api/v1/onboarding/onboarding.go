package onboarding

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Onboarding struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

func GetOnboarding(context echo.Context) error {
	onboardingItems := []Onboarding{
		Onboarding{
			Title: "Удобный поиск и фильтрация",
			Image: "https://yaroslavs-imac.local:80/static/onboarding/1.jpg",
		},
		Onboarding{
			Title: "Сохраняйте рецепты в избранное, создавайте новые и делитесь",
			Image: "https://yaroslavs-imac.local:80/static/onboarding/2.jpg",
		},
		Onboarding{
			Title: "Сканируйте холодильник и проверяйте наличие и срок годности продуктов",
			Image: "https://yaroslavs-imac.local:80/static/onboarding/3.jpg",
		},
	}
	return context.JSON(http.StatusOK, onboardingItems)
}