package controller

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/eventory-goa/app"
)

// GenreController implements the genre resource.
type GenreController struct {
	*goa.Controller
}

// NewGenreController creates a genre controller.
func NewGenreController(service *goa.Service) *GenreController {
	return &GenreController{Controller: service.NewController("GenreController")}
}

// Create runs the create action.
func (c *GenreController) Create(ctx *app.CreateGenreContext) error {
	// GenreController_Create: start_implement

	// Put your logic here

	// GenreController_Create: end_implement
	return nil
}

// List runs the list action.
func (c *GenreController) List(ctx *app.ListGenreContext) error {
	// GenreController_List: start_implement

	// Put your logic here

	// GenreController_List: end_implement
	res := app.GenreCollection{}
	return ctx.OK(res)
}
