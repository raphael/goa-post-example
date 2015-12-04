package main

import (
	"github.com/raphael/goa"
	"github.com/raphael/goa-post-example/app"
	"github.com/raphael/goa-post-example/swagger"
)

func main() {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(goa.RequestID())
	service.Use(goa.LogRequest())
	service.Use(goa.Recover())

	// Mount "Bottle" controller
	c := NewBottleController(service)
	app.MountBottleController(service, c)

	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on port 8080
	service.ListenAndServe(":8080")
}
