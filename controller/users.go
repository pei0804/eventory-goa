package controller

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory-goa/app"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
	db *gorm.DB
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service, db *gorm.DB) *UsersController {
	return &UsersController{
		Controller: service.NewController("UsersController"),
		db:         db,
	}
}

// AccountCreate runs the account create action.
func (c *UsersController) AccountCreate(ctx *app.AccountCreateUsersContext) error {
	// UsersController_AccountCreate: start_implement

	// Put your logic here

	// UsersController_AccountCreate: end_implement
	res := &app.Token{}
	return ctx.OK(res)
}

// TmpAccountCreate runs the tmp account create action.
func (c *UsersController) TmpAccountCreate(ctx *app.TmpAccountCreateUsersContext) error {
	// UsersController_TmpAccountCreate: start_implement

	// Put your logic here

	// UsersController_TmpAccountCreate: end_implement
	res := &app.Token{}
	return ctx.OK(res)
}
