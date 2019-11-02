package main

import (
	"tigade/delivery"
)

func main() {
	// Initializing core application
	// TODO

	// Initializing http handler
	h := delivery.SetupHttpHandler()

	// setup domain handlers and middleware
	// TODO

	// run http server
	delivery.RunHttpServer(h)
}
