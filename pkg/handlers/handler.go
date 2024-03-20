package handler

import (
	repository "rest/pkg/repositorys"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo *repository.MainRepository
}

func NewHandler(repo *repository.MainRepository) *Handler {
	return &Handler{repo: repo}
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
		auth.GET("/users", h.GetAllUsers)
		auth.DELETE("/delete:id", h.DeleteUser)
	}

	post := e.Group("/posts")
	{
		post.GET("/", h.GetAllPosts)
		post.POST("/create", h.CreatePost)
		post.PUT("/update/:id", h.UpdatePost)
		post.DELETE("/delete/:id", h.DeletePost)

	}
	e.POST("/test", h.Test)

	return e
}
