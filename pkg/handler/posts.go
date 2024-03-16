package handler

import (
	"encoding/json"
	"net/http"
	"rest/database"
	"rest/pkg/models"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getAllPosts(c echo.Context) error {
	var posts []models.Posts
	if err := database.DB.Find(&posts).Error; err != nil {
		return err
	}

	data, err := json.Marshal(posts)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(data))
}

func (h *Handler) setPost(c echo.Context) error {
	var post models.Posts

	if err := c.Bind(&post); err != nil {
		return err
	}

	if err := h.postRepo.Create(&post);err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &post)
}

func (h *Handler) updatePost(c echo.Context) error {
	id := c.Param("id")

	post := new(models.Posts)
	if err := database.DB.Where("id = ?", id).First(post).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Error receiving data for ID")
	}

	if err := c.Bind(post); err != nil {
		return err
	}

	if err := database.DB.Save(post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Error saving post to database")
	}

	return c.JSON(http.StatusOK, "Post successfully changed")
}

func (h *Handler) deletePost(c echo.Context) error {
	id := c.Param("id")

	post := new(models.Posts)
	if err := database.DB.Where("id = ?", id).First(post).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "post not found")
	}

	if err := database.DB.Delete(post, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "error deleting post")
	}

	return c.NoContent(http.StatusOK)
}

