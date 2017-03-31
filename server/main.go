// +build appengine
//go:generate goagen bootstrap -d github.com/tikasan/eventory-goa/design

package server

import (
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

	// Mount "events" controller
	c := controller.NewEventsController(service)
	app.MountEventsController(service, c)
	// Mount "genres" controller
	c2 := controller.NewGenresController(service)
	app.MountGenresController(service, c2)
	// Mount "prefs" controller
	c3 := controller.NewPrefsController(service)
	app.MountPrefsController(service, c3)
	// Mount "users" controller
	c4 := controller.NewUsersController(service)
	app.MountUsersController(service, c4)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
