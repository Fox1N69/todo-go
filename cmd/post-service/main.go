package main

import (
	"blog/internal/handlers"
	"blog/internal/repositorys"
	"blog/internal/services"
	"blog/pkg/database"
	"blog/pkg/routers"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()
	db := database.InitPostDB()

	//init repository
	postRepository := repositorys.NewPostRepository(db)

	//init service
	postService := services.NewPostServeci(postRepository)

	//init handler
	postHandler := handlers.NewPostHandler(postService)

	//ini routing
	if err := routers.NewPostRouter(postHandler, e); err != nil {
		logrus.Fatal(err)
	}

	//start server...
	logrus.Fatal(e.Start(":4000"))
}
