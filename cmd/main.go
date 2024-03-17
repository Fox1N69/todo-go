package main

import (
	"context"
	"os"
	"rest"
	"rest/database"
	handler "rest/pkg/handlers"
	repository "rest/pkg/repositorys"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	port := os.Getenv("PORT")

	srv := new(rest.Server)

	db := database.GetDB()
	repositorys := repository.NewMainRepository(db)

	handlers := new(handler.Handler)
	handlers.NewHandler(repositorys)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	log.Infoln("Server start on", "http://localhost:"+port+"4000")
	log.Fatal(srv.Run("4000", handlers.InitRouter()))
	log.Fatal(srv.Shutdown(context.Background()))
}
