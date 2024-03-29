package routers

import (
	"blog/internal/handlers"

	"github.com/labstack/echo/v4"
)

type AuthRouter struct {
	handler *handlers.AuthHandler
}

func NewAuthRouter(handlerAuth *handlers.AuthHandler) *AuthRouter {
	return &AuthRouter{handler: handlerAuth}
}

func (r *AuthRouter) RouterSetup(e *echo.Echo) {
	auth := e.Group("/auth")
	{
		auth.POST("/login", r.handler.Login)
		auth.POST("/singup", r.handler.Singup)
		auth.POST("/logout", r.handler.Logout)
		auth.POST("/refresh", r.handler.Refresh)
	}
}
