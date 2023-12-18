package main

import (
	"github.com/savin8305/ZSASS/handlers"
	"github.com/savin8305/ZSASS/stores/car"
	"gofr.dev/pkg/gofr"
)

func main() {
	// create the application object
	app := gofr.New()

	// Bypass header validation during API calls
	app.Server.ValidateHeaders = false

	store := car.New()
	h := handlers.New(store)

	// specifying the different routes supported by this service
	app.GET("/car/:name", h.Get)
	app.GET("/cars", h.List)      // New route for listing all cars
	app.POST("/car", h.Create)
	app.DELETE("/car/:name", h.Delete)

	app.Server.HTTP.Port = 8097

	app.Start()
}
