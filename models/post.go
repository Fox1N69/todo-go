package models

type Posts struct {
	id          int    `json:"id"`
	title       string `json:"title"`
	description string `json:"description"`
	post_body   string `json: "post_body"`
}
