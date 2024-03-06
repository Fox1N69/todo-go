package handler

import (
	"rest/models"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) getAllPosts(c echo.Context) error {
	var data map[string]string

	if err := c.Bind(&data); err != nil {
		log.Fatal(err)
	}

	posts := models.Posts{
		Title: data["title"],
		Description: data["description"],
		Postbody: data["postbody"],
	}

	return c.JSON(200, &posts)
}

func (h *Handler) updatePosts(c echo.Context) error {
	return nil
}

func (h *Handler) deletePosts(c echo.Context) error {
	return nil
}


