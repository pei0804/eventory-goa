package controller

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/eventory-goa/app"
)

// EventsController implements the events resource.
type EventsController struct {
	*goa.Controller
}

// NewEventsController creates a events controller.
func NewEventsController(service *goa.Service) *EventsController {
	return &EventsController{Controller: service.NewController("EventsController")}
}

// KeepEvent runs the keep event action.
func (c *EventsController) KeepEvent(ctx *app.KeepEventEventsContext) error {
	// EventsController_KeepEvent: start_implement

	// Put your logic here

	// EventsController_KeepEvent: end_implement
	return nil
}

// List runs the list action.
func (c *EventsController) List(ctx *app.ListEventsContext) error {
	// EventsController_List: start_implement

	// Put your logic here

	// EventsController_List: end_implement
	res := app.EventCollection{}
	return ctx.OK(res)
}
