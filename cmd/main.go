package main

import (
	"context"
	"os"
	"rest"
	"rest/pkg/handler"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	srv := new(rest.Server)
	handlers := new(handler.Handler)
	log := logrus.New()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")

	log.Infoln("Server start on port" + port)
	log.Fatal(srv.Run(port, handlers.InitRouter()))
	log.Fatal(srv.Shutdown(context.Background()))
}
