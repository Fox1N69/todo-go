package handler

import (
	"encoding/json"
	"net/http"
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
	var data models.Posts
	id := c.Param("id")

	if err := database.DB.Where("id = ?", id).First(&data).Error; err != nil {
		log.Fatal("Error receiving data for ID: ", err)
		c.String(http.StatusInternalServerError, "Error receiving data for ID")
		return err
	}

	if err := c.Bind(&data); err != nil {
		return err
	}

	if err := database.DB.Save(&data).Error; err != nil {
		log.Fatal("Error save to database ", err)
	}

	return c.JSON(200, data)
}

func (h *Handler) deletePosts(c echo.Context) error {
	var data models.Posts
	id := c.Param("id")

	if err := database.DB.Where("id = ?", id).First(&data).Error; err != nil {
		log.Fatal("Error recceving data for ID ", err)
		c.String(http.StatusInternalServerError, "error receiving data for ID")
		return err
	}

	if err := database.DB.Delete(&data).Error; err != nil {
		return err
	}

	return c.String(http.StatusOK, "User successfully deleted")
}
