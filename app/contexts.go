//************************************************************************//
// winecellar: Application Contexts
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
	"fmt"
	"strconv"

	"github.com/raphael/goa"
)

// ShowBottleContext provides the Bottle show action context.
type ShowBottleContext struct {
	*goa.Context
	BottleID int
}

// NewShowBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the Bottle controller show action.
func NewShowBottleContext(c *goa.Context) (*ShowBottleContext, error) {
	var err error
	ctx := ShowBottleContext{Context: c}
	rawBottleID, ok := c.Get("bottleID")
	if ok {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			ctx.BottleID = int(bottleID)
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowBottleContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OK(resp *Bottle) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.goa.example.bottle+json; charset=utf-8")
	return ctx.JSON(200, r)
}
