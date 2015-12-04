package main

import (
	"github.com/raphael/goa"
	"github.com/raphael/goa-post-example/app"
)

// BottleController implements the Bottle resource.
type BottleController struct {
	goa.Controller
}

// NewBottleController creates a Bottle controller.
func NewBottleController(service goa.Service) app.BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	res := &app.Bottle{}
	return ctx.OK(res)
}
