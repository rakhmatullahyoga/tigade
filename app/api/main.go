package main

import (
	"github.com/rakhmatullahyoga/tigade"
	"github.com/rakhmatullahyoga/tigade/config"
	"github.com/rakhmatullahyoga/tigade/delivery"
)

func main() {
	// Initializing core application
	core := tigade.NewCoreService()
	defer core.Shutdown()

	// Initializing http handler
	h := delivery.SetupHttpHandler()

	// Setup domain handlers and middleware
	// TODO

	// Run http server
	port := config.GetInstance().AppPort
	delivery.RunHttpServer(h, port)
}
