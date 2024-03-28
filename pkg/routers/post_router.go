package routers

import (
	"blog/internal/handlers"

	"github.com/labstack/echo/v4"
)

type PostRouter struct {
	handler *handlers.PostHandler
}

func NewPostRouter(postHander *handlers.PostHandler) *PostRouter {
	return &PostRouter{
		handler: postHander,
	}
}

func (r *PostRouter) RouterSetup(e *echo.Echo) {
	api := e.Group("/api")
	{
		api.GET("/post", r.handler.GetAllPosts)
		api.POST("/post:create", r.handler.CreatePost)
		api.PUT("/post:id", r.handler.UpdatePost)
		api.DELETE("/post:id", r.handler.DeletePost)
	}
}
