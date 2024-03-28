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
	postService *services.PostServeci
}

func NewPostHandler(postService *services.PostServeci) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) GetAllPosts(c echo.Context) error {
	return nil
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
