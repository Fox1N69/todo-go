package main

import (
	"blog/internal/handlers"
	"blog/internal/repositorys"
	"blog/internal/services"
	"blog/pkg/database"
	"blog/pkg/routers"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	db := database.InitAuthDB()

	//ini repotitory auth
	repo := repositorys.NewAuthRepository(db)

	//init service auth
	service := services.NewAuthService(repo)

	//init handler auth
	handler := handlers.NewAuthHandler(service)

	//init router auth
	router := routers.NewAuthRouter(handler)
	router.RouterSetup(app)

	log.Fatal(app.Start(":8000"))
}
