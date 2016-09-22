//go:generate goagen bootstrap -d study-event-api/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"study-event-api/app"
)

func main() {
	// Create service
	service := goa.New("study-event")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "keywords" controller
	c := NewKeywordsController(service)
	app.MountKeywordsController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
