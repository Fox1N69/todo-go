package routers

import (
	"blog/internal/handlers"

	"github.com/labstack/echo/v4"
)

type PostRouter struct {
	handler *handlers.PostHandler
	echo    *echo.Echo
}

func NewPostRouter(postHander *handlers.PostHandler, e *echo.Echo) *PostRouter {
	return &PostRouter{handler: postHander, echo: e}
}

func (r *PostRouter) InitRouters(e *echo.Echo) {
	api := e.Group("/api")
	{
		api.GET("/posts", r.handler.GetAllPost)
	}
}
