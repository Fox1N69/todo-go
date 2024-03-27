package main

import "rest"

func main() {
	app := rest.StartServer("4000")
	defer app.Shutdown()
	select {}
}
