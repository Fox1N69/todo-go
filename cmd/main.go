package main

import (
	"log"
	"rest"
	"rest/pkg/handler"

)

func main() {
	srv := new(rest.Server)
	handlers := new(handler.Handler)


	log.Println("Server start on port 4000")
	log.Fatal(srv.Run("4000", handlers.InitRouter()))
}
