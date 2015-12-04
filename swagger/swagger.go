//************************************************************************//
// winecellar Swagger Spec
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/raphael/goa-post-example
// --design=github.com/raphael/goa-post-example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package swagger

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// MountController mounts the swagger spec controller under "/swagger.json".
func MountController(service goa.Service) {
	ctrl := service.NewController("Swagger")
	service.Info("mount", "ctrl", "Swagger", "action", "Show", "route", "GET /swagger.json")
	h := ctrl.NewHTTPRouterHandle("Show", getSwagger)
	service.HTTPHandler().(*httprouter.Router).Handle("GET", "/swagger.json", h)
}

// getSwagger is the httprouter handle that returns the Swagger spec.
// func getSwagger(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
func getSwagger(ctx *goa.Context) error {
	ctx.Header().Set("Content-Type", "application/swagger+json")
	ctx.Header().Set("Cache-Control", "public, max-age=3600")
	return ctx.Respond(200, []byte(spec))
}

// Generated spec
const spec = `{"swagger":"2.0","info":{"description":"The winecellar service API","version":""},"host":"cellar.goa.design","basePath":"/cellar","schemes":["http","https"],"consumes":["application/json"],"produces":["application/json"],"paths":{"/bottles/{bottleID}":{"get":{"description":"Retrieve bottle with given ID","operationId":"Bottle#show","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"bottleID","in":"path","description":"The name of the bottle to retrieve","required":true,"type":"integer"}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/Bottle"}},"404":{"description":""}},"schemes":["https"]}}},"definitions":{"Bottle":{"title":"Mediatype identifier: application/vnd.goa.example.bottle+json","type":"object","properties":{"color":{"type":"string","enum":["red","white","rose","yellow","sparkling"]},"href":{"type":"string"},"id":{"type":"integer","description":"ID of bottle"},"name":{"type":"string","description":"The bottle  name","minLength":1,"maxLength":255},"sweetness":{"type":"integer","minimum":1,"maximum":5}},"description":"A bottle of wine"}}} `
