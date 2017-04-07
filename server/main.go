// +build appengine

//go:generate goagen bootstrap -d github.com/tikasan/eventory-goa/design

package server

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tikasan/eventory-goa/app"
	"github.com/tikasan/eventory-goa/controller"
	"github.com/tikasan/eventory-goa/db"
)

func init() {
	// Create service
	service := goa.New("eventory")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	app.UseKeyMiddleware(service, NewAPIKeyMiddleware())

	cs, err := db.NewConfigsFromFile()
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	dbcon, err := cs.Open()
	if err != nil {
		log.Fatalf("db initialization failed: %s", err)
	}

	// Mount "events" controller
	c := controller.NewEventsController(service, dbcon)
	app.MountEventsController(service, c)
	// Mount "genres" controller
	c2 := controller.NewGenresController(service, dbcon)
	app.MountGenresController(service, c2)
	// Mount "prefs" controller
	c3 := controller.NewPrefsController(service, dbcon)
	app.MountPrefsController(service, c3)
	// Mount "users" controller
	c4 := controller.NewUsersController(service, dbcon)
	app.MountUsersController(service, c4)

	// Setup HTTP handler
	http.HandleFunc("/", service.Mux.ServeHTTP)
}

func NewAPIKeyMiddleware() goa.Middleware {

	// Instantiate API Key security scheme details generated from design
	scheme := app.NewKeySecurity()

	// Middleware
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log header specified by scheme
			key := req.Header.Get(scheme.Name)
			// A real app would do something more interesting here
			if len(key) == 0 || key == "Bearer" {
				goa.LogInfo(ctx, "failed api key auth")
				return errors.Unauthenticated("missing auth")
			}
			// Proceed.
			goa.LogInfo(ctx, "auth", "apikey", "key", key)
			return h(ctx, rw, req)
		}

	}
}
