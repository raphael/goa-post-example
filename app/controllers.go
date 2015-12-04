//************************************************************************//
// winecellar: Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/raphael/goa-post-example
// --design=github.com/raphael/goa-post-example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// BottleController is the controller interface for the Bottle actions.
type BottleController interface {
	goa.Controller
	Show(*ShowBottleContext) error
}

// MountBottleController "mounts" a Bottle resource controller on the given service.
func MountBottleController(service goa.Service, ctrl BottleController) {
	router := service.HTTPHandler().(*httprouter.Router)
	var h goa.Handler
	h = func(c *goa.Context) error {
		ctx, err := NewShowBottleContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/cellar/bottles/:bottleID", ctrl.NewHTTPRouterHandle("Show", h))
	service.Info("mount", "ctrl", "Bottle", "action", "Show", "route", "GET /cellar/bottles/:bottleID")
}
