package main

import (
	"log"
	"rest"
	"rest/pkg/handler"

)

func main() {
	srv := new(rest.Server)
	handler := new(handler.Handler)

	if err := srv.Run("8000", handler.InitRouter()); err != nil {
		log.Fatal(err)
	}
}
