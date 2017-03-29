// +build appengine

//go:generate goagen bootstrap -d github.com/tikasan/eventory-goa/design

package main

import (
	"net/http"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tikasan/eventory-goa/app"
	"github.com/tikasan/eventory-goa/controller"
)

func init() {
	// Create service
	service := goa.New("eventory")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "event" controller
	c := controller.NewEventController(service)
	app.MountEventController(service, c)
	// Mount "genre" controller
	c2 := controller.NewGenreController(service)
	app.MountGenreController(service, c2)
	// Mount "swagger" controller
	c3 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c3)
	// Mount "user" controller
	c4 := controller.NewUserController(service)
	app.MountUserController(service, c4)

	// Setup HTTP handler
	http.HandleFunc("/", service.Mux.ServeHTTP)
}
