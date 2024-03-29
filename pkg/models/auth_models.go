package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"-"`
}

type LoginData struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
}
