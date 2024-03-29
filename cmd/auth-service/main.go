package main

import (
	"blog/internal/handlers"
	"blog/internal/repositorys"
	"blog/internal/services"
	"blog/pkg/routers"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	repo := repositorys.NewAuthRepository()

	service := services.NewAuthService(repo)

	handler := handlers.NewAuthHandler(service)

	router := routers.NewAuthRouter(handler)
	router.RouterSetup(app)

	log.Fatal(app.Start(":8000"))
}
