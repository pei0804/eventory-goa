package controller

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/eventory-goa/app"
)

// PrefsController implements the prefs resource.
type PrefsController struct {
	*goa.Controller
}

// NewPrefsController creates a prefs controller.
func NewPrefsController(service *goa.Service) *PrefsController {
	return &PrefsController{Controller: service.NewController("PrefsController")}
}

// PrefFollow runs the pref follow action.
func (c *PrefsController) PrefFollow(ctx *app.PrefFollowPrefsContext) error {
	// PrefsController_PrefFollow: start_implement

	// Put your logic here

	// PrefsController_PrefFollow: end_implement
	return nil
}
