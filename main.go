package main

import (
	"github.com/savin8305/ZSASS/handlers"
	"github.com/savin8305/ZSASS/stores/customer"
	"gofr.dev/pkg/gofr"
)

func main() {
	// create the application object
	app := gofr.New()

	// Bypass header validation during API calls
	app.Server.ValidateHeaders = false

	store := customer.New()
	h := handlers.New(store)

	// specifying the different routes supported by this service
	app.GET("/customer/:name", h.Get)
	app.GET("/customers", h.List)      // New route for listing all customers
	app.POST("/customer", h.Create)
	app.PUT("/customer/:name", h.Update) // New route for updating a customer
	app.DELETE("/customer/:name", h.Delete)

	app.Server.HTTP.Port = 8097

	app.Start()
}
