package controller

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/eventory-goa/app"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// EventFavUpdate runs the event fav update action.
func (c *UserController) EventFavUpdate(ctx *app.EventFavUpdateUserContext) error {
	// UserController_EventFavUpdate: start_implement

	// Put your logic here

	// UserController_EventFavUpdate: end_implement
	return nil
}

// GenreFavUpdate runs the genre fav update action.
func (c *UserController) GenreFavUpdate(ctx *app.GenreFavUpdateUserContext) error {
	// UserController_GenreFavUpdate: start_implement

	// Put your logic here

	// UserController_GenreFavUpdate: end_implement
	return nil
}
