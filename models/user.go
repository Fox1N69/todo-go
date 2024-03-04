package models

type Users struct {
	id       int    `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
}
