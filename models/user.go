package models

type Users struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
