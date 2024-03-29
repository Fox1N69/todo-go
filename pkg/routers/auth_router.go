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

func (ar *AuthRouter) RouterSetup(e *echo.Echo) {
	
} 