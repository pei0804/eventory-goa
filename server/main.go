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
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory-goa/app"
	"github.com/tikasan/eventory-goa/controller"
	"github.com/tikasan/eventory-goa/database"
	"github.com/tikasan/eventory-goa/models"
	"github.com/tikasan/eventory-goa/utility"
)

func init() {
	// Create service
	service := goa.New("eventory")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	cs, err := database.NewConfigsFromFile("dbconfig.yml")
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	dbcon, err := cs.Open("setting")
	if err != nil {
		log.Fatalf("database initialization failed: %s", err)
	}

	app.UseKeyMiddleware(service, NewAPIKeyMiddleware(dbcon))

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

func NewAPIKeyMiddleware(db *gorm.DB) goa.Middleware {

	// Instantiate API Key security scheme details generated from design
	scheme := app.NewKeySecurity()

	// Middleware
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log header specified by scheme
			token := req.Header.Get(scheme.Name)
			// A real app would do something more interesting here
			if len(token) == 0 || token == "Bearer" {
				goa.LogInfo(ctx, "failed api token auth")
				return errors.Unauthenticated("missing auth")
			}
			userTerminalDB := models.NewUserTerminalDB(db)
			userID, err := userTerminalDB.GetUserIDByToken(ctx, token)
			if err != nil {
				goa.LogInfo(ctx, "failed api token auth")
				return errors.Unauthenticated("missing auth")
			}
			utility.SetToken(ctx, userID)
			// Proceed.
			goa.LogInfo(ctx, "auth", "apikey", "token", token)
			return h(ctx, rw, req)
		}

	}
}
