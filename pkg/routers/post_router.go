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
		api.GET("/posts", r.handler.GetAllPost)
	}

	e.GET("/test", r.handler.Test)

}
