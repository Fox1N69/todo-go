package handler

import (
	"encoding/json"
	"rest/database"
	"rest/models"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) getAllPosts(c echo.Context) error {
	var posts []models.Posts
	database.DB.Find(&posts)

	data, err := json.Marshal(posts)
	if err != nil {
		log.Fatal(err)
	}

	jsonString := string(data)

	return c.String(200, jsonString)
}

func (h *Handler) setPost(c echo.Context) error {

	data := new(models.Posts)

	if err := c.Bind(&data); err != nil {
		log.Fatal(err)
	}

	database.DB.Create(&data)

	return c.JSON(201, data)
}

func (h *Handler) updatePosts(c echo.Context) error {
	var data []models.Posts
	id := c.Param("id")

	if err := database.DB.First(&data, id); err != nil {
		log.Fatal("Error id receving ",err)
	}

	if err := c.Bind(&data); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (h *Handler) deletePosts(c echo.Context) error {
	return nil
}
