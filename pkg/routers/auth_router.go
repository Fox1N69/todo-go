package routers

import (
	"blog/internal/controller"

	"github.com/labstack/echo/v4"
)

type AuthRouter struct {
	controller *controller.AuthController
}

func NewAuthRouter(controller *controller.AuthController) *AuthRouter {
	return &AuthRouter{controller: controller}
}

func (r *AuthRouter) RouterSetup(e *echo.Echo) {
	auth := e.Group("/auth")
	{
		auth.POST("/login", r.controller.Login)
		auth.POST("/singup", r.controller.SingUp)
		auth.POST("/logout", r.controller.Logout)
		auth.POST("/refresh", r.controller.Refresh)
	}
}
