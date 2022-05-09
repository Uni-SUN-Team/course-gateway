package model

type Categories struct {
	Data []CategoriesData `json:"data"`
}

type CategoriesData struct {
	Id         int64                `json:"id"`
	Attributes CategoriesAttributes `json:"attributes"`
}

type CategoriesAttributes struct {
	Name        string `json:"name"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	PublishedAt string `json:"publishedAt"`
	Locale      string `json:"locale"`
	Slug        string `json:"slug"`
}
