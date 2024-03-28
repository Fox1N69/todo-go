package handlers

import (
	"blog/internal/services"

	"github.com/labstack/echo/v4"
)

type PostHandlerI interface {
	GetAllPosts(c echo.Context) error
	CreatePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}
type PostHandler struct {
	service *services.PostServeci
}

func NewPostHandler(postService *services.PostServeci) *PostHandler {
	return &PostHandler{service: postService}
}

func (h *PostHandler) GetAllPosts(c echo.Context) error {
	post, err := h.service.GetAllPosts()
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"message": "Post create success",
		"post":    post,
	})
}

func (h *PostHandler) CreatePost(c echo.Context) error {
	return nil
}

func (h *PostHandler) UpdatePost(c echo.Context) error {
	return nil
}

func (h *PostHandler) DeletePost(c echo.Context) error {
	return nil
}
