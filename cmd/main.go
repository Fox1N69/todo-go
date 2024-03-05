package main

import (
	"context"
	"os"
	"path/filepath"
	"rest"
	"rest/database"
	"rest/pkg/handler"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	database.InitGormDB()

	srv := new(rest.Server)
	handlers := new(handler.Handler)

	if err := godotenv.Load(filepath.Join(".env")); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")

	log.Infoln("Server start on port" + port)
	log.Fatal(srv.Run(port, handlers.InitRouter()))
	log.Fatal(srv.Shutdown(context.Background()))
}
