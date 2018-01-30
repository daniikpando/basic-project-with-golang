package models

type Article struct {
	IdArticle   int    `json:"id_article"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
}
