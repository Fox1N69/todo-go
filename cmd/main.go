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

	server := &rest.Server{}

	db := database.GetDB()
	repo := repository.NewMainRepository(db)
	
	handlers := handler.NewHandler(repo)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	log.Println("Server start on", "http://localhost:"+port+"4000")
	log.Fatal(server.Run("4000", handlers.InitRouter()))
	log.Fatal(server.Shutdown(context.Background()))
}
