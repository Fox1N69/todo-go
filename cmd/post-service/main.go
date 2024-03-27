package main

import (
	"blog/internal/handlers"
	"blog/internal/repositorys"
	"blog/internal/services"
	"blog/pkg/routers"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()

	postRepository := repositorys.NewPostRepository()

	postService := services.NewPostServeci(postRepository)

	postHandler := handlers.NewPostHandler(postService)

	if err := routers.NewPostRouter(postHandler, e); err != nil {
		logrus.Fatal(err)
	}

	logrus.Fatal(e.Start(":4000"))
}
