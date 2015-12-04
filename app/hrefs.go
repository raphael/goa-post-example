//************************************************************************//
// winecellar: Application Resource Href Factories
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

import "fmt"

// BottleHref returns the resource href.
func BottleHref(bottleID interface{}) string {
	return fmt.Sprintf("/cellar/bottles/%v", bottleID)
}
