package handlers

import (
	"blog/internal/services"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	postService *services.PostServeci
}

func NewPostHandler(postService *services.PostServeci) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) GetAllPost(c echo.Context) error {
	return nil
}
