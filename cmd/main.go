package main

import (
	"log"
	"os"
	"rest"
	"rest/pkg/handler"

	"github.com/joho/godotenv"
)

func main() {
	srv := new(rest.Server)
	handlers := new(handler.Handler)

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	log.Println("Server start on port" + port)
	log.Fatal(srv.Run(port, handlers.InitRouter()))
}
