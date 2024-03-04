package models

type Users struct {
	ID       int    `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
}
