package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getAllPosts(c echo.Context) error {
	return c.String(http.StatusOK,"hello echo")
}

func (h *Handler) updatePosts(c echo.Context) error {
	return nil
}

func (h *Handler) deletePosts(c echo.Context) error {
	return nil
}
