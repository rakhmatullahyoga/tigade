package main

import (
	"github.com/rakhmatullahyoga/tigade"
	"github.com/rakhmatullahyoga/tigade/delivery"
	"github.com/rakhmatullahyoga/tigade/config"
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
	port := config.GetInstance().AppPort
	delivery.RunHttpServer(h, port)
}
