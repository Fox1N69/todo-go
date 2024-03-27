package rest

import "github.com/gofiber/fiber/v3"

type Server struct {
}

func StartServer(port string) *fiber.App {
	app := fiber.New()

	go func() {
		if err := app.Listen(":" + port); err != nil {
			panic(err)
		}
	}()

	return app
}
