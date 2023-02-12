package entities


type RecipeShortItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	ImageUrl string `json:"image_url"`
	Rating float32 `json:"rating"`
	CookTime int `json:"cook_time"`
	Calories float32 `json:"calories"`
	IsFavorite bool `json:"is_favorite"`
}

type IngredientItem struct {
	Name string `json:"name"`
	Count float32 `json:"count"`
	Units string `json:"units"`
}

type StepItem struct {
	Desc string `json:"desc"`
}

type MediaItem struct {
	Type int `json:"type"` // 1 - image, 2 - video
	Url string `json:"url"`
}

type RecipeFullItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Rating float32 `json:"rating"`
	CookTime int `json:"cook_time"`
	Calories float32 `json:"calories"`
	IsFavorite bool `json:"is_favorite"`
	MediaUrls []MediaItem `json:"medias"`
	Ingredients []IngredientItem `json:"ingredients"`
	Steps []StepItem `json:"steps"`
}