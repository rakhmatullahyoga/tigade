package main

import (
	"tigade"
	"tigade/delivery"
)

func main() {
	// Initializing core application
	core := tigade.NewCoreService()
	defer core.Shutdown()

	// Initializing http handler
	h := delivery.SetupHttpHandler()

	// setup domain handlers and middleware
	// TODO

	// run http server
	delivery.RunHttpServer(h, 8080)
}
