package main

import (
	"context"
	"os"
	"rest"
	"rest/pkg/handler"

	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
)


func main() {
	port := os.Getenv("PORT")

	srv := new(rest.Server)
	handlers := new(handler.Handler)

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	log.Infoln("Server start on", "http://localhost:"+port)
	log.Fatal(srv.Run(port, handlers.InitRouter()))
	log.Fatal(srv.Shutdown(context.Background()))
}