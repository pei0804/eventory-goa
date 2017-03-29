package controller

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/eventory-goa/app"
)

// EventController implements the event resource.
type EventController struct {
	*goa.Controller
}

// NewEventController creates a event controller.
func NewEventController(service *goa.Service) *EventController {
	return &EventController{Controller: service.NewController("EventController")}
}

// List runs the list action.
func (c *EventController) List(ctx *app.ListEventContext) error {
	// EventController_List: start_implement

	// Put your logic here

	// EventController_List: end_implement
	res := app.EventCollection{}
	return ctx.OK(res)
}
