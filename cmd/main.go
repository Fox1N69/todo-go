package main

import (
	"context"
	"os"
	"rest"
	"rest/database"
	"rest/pkg/handler"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	database.InitDB()

	srv := new(rest.Server)
	handlers := new(handler.Handler)

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")

	log.Infoln("Server start on port" + port)
	log.Fatal(srv.Run(port, handlers.InitRouter()))
	log.Fatal(srv.Shutdown(context.Background()))
}
