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