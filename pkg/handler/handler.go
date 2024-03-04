package handler

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (h *Handler) InitRouter() *echo.Echo {
	e := echo.New()

	e.GET("/", h.Hello)

	auth := e.Group("/auth")
	{
		auth.POST("/sing-up", h.SingUp)
		auth.POST("/sing-in", h.SingIn)
	}

	list := e.Group("/list")
	{
		list.GET("/", h.GetAllList)
	}

	return e
}
