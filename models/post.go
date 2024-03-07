package models

type Posts struct {
	Id          int    `json:"-" gorm:"primaryKey; autoIncrement:true"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PostBody    string `json: "postbody"`
}
