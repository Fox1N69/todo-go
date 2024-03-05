package models

type Posts struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Postbody   string `json: "post_body"`
}
