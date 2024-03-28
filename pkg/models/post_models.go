package models

type Post struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PostBody    string `json:"postBody"`
}
