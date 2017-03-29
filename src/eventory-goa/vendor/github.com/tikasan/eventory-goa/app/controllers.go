// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/eventory-goa/design
// --out=$(GOPATH)/src/github.com/tikasan/eventory-goa
// --version=v1.1.0-dirty
//
// API "eventory": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// EventController is the controller interface for the Event actions.
type EventController interface {
	goa.Muxer
	List(*ListEventContext) error
}

// MountEventController "mounts" a Event resource controller on the given service.
func MountEventController(service *goa.Service, ctrl EventController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/v2/event", ctrl.MuxHandler("preflight", handleEventOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListEventContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleEventOrigin(h)
	service.Mux.Handle("GET", "/api/v2/event", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Event", "action", "List", "route", "GET /api/v2/event")
}

// handleEventOrigin applies the CORS response headers corresponding to the origin.
func handleEventOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// GenreController is the controller interface for the Genre actions.
type GenreController interface {
	goa.Muxer
	Create(*CreateGenreContext) error
	List(*ListGenreContext) error
}

// MountGenreController "mounts" a Genre resource controller on the given service.
func MountGenreController(service *goa.Service, ctrl GenreController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/v2/genre/new", ctrl.MuxHandler("preflight", handleGenreOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/v2/genre", ctrl.MuxHandler("preflight", handleGenreOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateGenreContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Create(rctx)
	}
	h = handleGenreOrigin(h)
	service.Mux.Handle("POST", "/api/v2/genre/new", ctrl.MuxHandler("Create", h, nil))
	service.LogInfo("mount", "ctrl", "Genre", "action", "Create", "route", "POST /api/v2/genre/new")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListGenreContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleGenreOrigin(h)
	service.Mux.Handle("GET", "/api/v2/genre", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Genre", "action", "List", "route", "GET /api/v2/genre")
}

// handleGenreOrigin applies the CORS response headers corresponding to the origin.
func handleGenreOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger/*filepath", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger/*filepath", "public/swagger")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "public/swagger", "route", "GET /swagger/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger/", "public/swagger/index.html")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "public/swagger/index.html", "route", "GET /swagger/")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	EventFavUpdate(*EventFavUpdateUserContext) error
	GenreFavUpdate(*GenreFavUpdateUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/v2/user/:eventID/keep", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/v2/user/:eventID/nokeep", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/v2/user/:genreID/add", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/v2/user/:genreID/remove", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewEventFavUpdateUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.EventFavUpdate(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("PUT", "/api/v2/user/:eventID/keep", ctrl.MuxHandler("EventFavUpdate", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "EventFavUpdate", "route", "PUT /api/v2/user/:eventID/keep")
	service.Mux.Handle("PUT", "/api/v2/user/:eventID/nokeep", ctrl.MuxHandler("EventFavUpdate", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "EventFavUpdate", "route", "PUT /api/v2/user/:eventID/nokeep")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGenreFavUpdateUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GenreFavUpdate(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("PUT", "/api/v2/user/:genreID/add", ctrl.MuxHandler("GenreFavUpdate", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "GenreFavUpdate", "route", "PUT /api/v2/user/:genreID/add")
	service.Mux.Handle("PUT", "/api/v2/user/:genreID/remove", ctrl.MuxHandler("GenreFavUpdate", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "GenreFavUpdate", "route", "PUT /api/v2/user/:genreID/remove")
}

// handleUserOrigin applies the CORS response headers corresponding to the origin.
func handleUserOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}