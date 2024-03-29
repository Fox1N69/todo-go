package handlers

import (
	"blog/internal/services"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(serviceAuth *services.AuthService) *AuthHandler {
	return &AuthHandler{service: serviceAuth}
}

func (h *AuthHandler) Login(c echo.Context) error {
	return c.JSON(200, "login")
}

func (h *AuthHandler) Singup(c echo.Context) error {
	return c.JSON(200, "login")
}

func (h *AuthHandler) Logout(c echo.Context) error {
	return c.JSON(200, "login")
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	return c.JSON(200, "login")
}
