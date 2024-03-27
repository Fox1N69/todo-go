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
	return c.JSON(200, map[string]interface{}{"message": "hello world"})
}

func (h *PostHandler) Test(c echo.Context) error {
	return c.JSON(200, "Hello world test")
}
