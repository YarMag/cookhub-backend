package entities

type RecipeUserFeedItem struct {
	Author AuthorShortItem `json:"author"`
	Recipe RecipeShortItem `json:"recipe"`
}

type RecipeCompilationUserFeedItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Recipes []RecipeShortItem `json:"recipes"`
	HasMore bool `json:"has_more"`
}