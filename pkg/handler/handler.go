package handler

import (
	"rest/pkg/repository"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo *repository.MainRepository
}

func (h *Handler) InitRouter() *echo.Echo {
	e := echo.New()

	admin := e.Group("/admin")
	{
		admin.GET("/", h.Admin)
	}

	auth := e.Group("/auth")
	{
		auth.POST("/sing-up", h.SingUp)
		auth.POST("/sing-in", h.SingIn)
		auth.GET("/users", h.getAllUsers)
		auth.DELETE("/delete:id", h.deleteUser)
	}

	post := e.Group("/posts")
	{
		post.GET("/", h.getAllPosts)
		post.POST("/set", h.setPost)
		post.PUT("/update/:id", h.updatePost)
		post.DELETE("/delete/:id", h.deletePost)

	}

	return e
}
