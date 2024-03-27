package rest

import (
	"github.com/gofiber/fiber/v3"
)

type Server struct {
	app *fiber.App
}

func StartServer()