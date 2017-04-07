package controller

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory-goa/app"
)

// EventsController implements the events resource.
type EventsController struct {
	*goa.Controller
	db *gorm.DB
}

// NewEventsController creates a events controller.
func NewEventsController(service *goa.Service, db *gorm.DB) *EventsController {
	return &EventsController{
		Controller: service.NewController("EventsController"),
		db:         db,
	}
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
	//eventDB := models.NewEventDB(c.db)
	//events, err := eventDB.List(ctx.Context)
	//if err != nil {
	//	log.Error("miss")
	//}

	// EventsController_List: end_implement
	//return fmt.Errorf(ctx.URL.Path)
	//return fmt.Errorf(ctx.URL.Query().Get("page"))
}
