package handlers

import (
	"blog/internal/services"
	"blog/pkg/models"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	postService *services.PostServeci
}

func NewPostHandler(postService *services.PostServeci) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) Test(c echo.Context) error {
	var data models.Test

	if err := c.Bind(&data); err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"data":    &data,
	})
}
