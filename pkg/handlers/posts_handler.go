package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"rest/database"
	"rest/pkg/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Test(c echo.Context) error {
	data := struct {
		Message string `json:"message"`
	}{}

	if err := c.Bind(&data); err != nil {
		return err
	}

	return c.JSON(200, &data)
}

func (h *Handler) GetAllPosts(c echo.Context) error {
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

func (h *Handler) CreatePost(c echo.Context) error {
	var post models.Posts

	if err := c.Bind(&post); err != nil {
		return err
	}

	if err := h.repo.Post.Create(&post); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &post)
}

func (h *Handler) UpdatePost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	post, err := h.repo.Post.GetByID(uint(postID))
	if err != nil {
		return err
	}

	if err := c.Bind(&post); err != nil {
		log.Fatal(err)
	}

	if err = h.repo.Post.UpdatePost(post); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Post: successfully updated")
}

func (h *Handler) DeletePost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	post, err := h.repo.Post.GetByID(uint(postID))
	if err != nil {
		return err
	}

	if err := h.repo.Post.DeletePost(post); err != nil {
		return err
	}

	return c.JSON(200, "Post delete")
}
