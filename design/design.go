package design

import (
	. "github.com/raphael/goa/design" // "dot" imports make the DSL easier to read.
	. "github.com/raphael/goa/design/dsl"
)

var _ = API("winecellar", func() { // The API function defines an API given its name.
	Description("The winecellar service API")
	BasePath("/cellar") // Base path or prefix to all requests.
	// Can be overridden in action definitions using an
	// absolute path starting with //.
	Host("cellar.goa.design") // Default API host used by clients and Swagger.
	Scheme("http")            // Supported API URL scheme used by clients and Swagger.
	Scheme("https")           // Scheme("http", "https") works too
})

var _ = Resource("Bottle", func() { // Define the Bottle resource
	DefaultMedia(BottleMedia) // Default media type used to render the bottle resources
	BasePath("/bottles")      // Gets appended to the API base path

	Action("show", func() { // Define the show action on the Bottle resource
		Routing(GET("/:bottleID")) // The relative path to the show endpoint. The full path is
		// built concatenating the API and resource base paths with it.
		// Uses a wildcard to capture the requested bottle ID.
		// Wildcards can start with : to capture a single path segment
		// or with * to capture the rest of the path.
		Description("Retrieve bottle with given ID")
		Params(func() { // Define the request parameters found in the URI (wildcards)
			// and the query string.
			Param( // Define a single parameter
				"bottleID", // Here it corresponds to the path segment captured by :bottleID
				Integer,    // The JSON type of the parameter
				"The name of the bottle to retrieve", // An optional description
			)
		})
		Response(OK)       // Define a potential response
		Response(NotFound) // An action may define any number of responses.
		// Their content is defined through ResponseTemplates (not shown
		// in this simplistic example). Here we use the default response
		// templates defined in goa.
	})
})

var BottleMedia = MediaType("application/vnd.goa.example.bottle+json", func() {
	Description("A bottle of wine")
	Attributes(func() {
		Attribute("id", Integer, "ID of bottle") // Attribute defines a single field in
		// the media type data structure given its
		// name, type and description.
		Attribute("href")                                      // The default type for attributes is String.
		Attribute("name", String, "The bottle  name", func() { // Like with API, Resource and Action an attribute
			// definition may use an anonymous function as
			// last argument to define additional properties.
			MinLength(1)   // Here we define validation rules specifying a
			MaxLength(255) // minimum and maximum number of characters in a bottle
			// name.
		})
		Attribute("color", func() { // Descriptions are optional.
			Enum("red", "white", "rose", "yellow", "sparkling") // Possible field values
		})
		Attribute("sweetness", Integer, func() {
			Minimum(1) // Minimum and maximum int field values.
			Maximum(5)
		})

		View("default", func() { // Views are used to render a media type.
			Attribute("id")        // A media type can have one or more views
			Attribute("href")      // and must define the "default" view.
			Attribute("name")      // The view simply lists the fields to render.
			Attribute("color")     // It can also specify the view to use to render
			Attribute("sweetness") // fields whose type is itself a media type
			// (the "default" view by default). Not used here.
		})
	})
})
