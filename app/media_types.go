//************************************************************************//
// winecellar: Application Media Types
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

import "github.com/raphael/goa"

// A bottle of wine
// Identifier: application/vnd.goa.example.bottle+json
type Bottle struct {
	Color string
	Href  string
	// ID of bottle
	ID int
	// The bottle  name
	Name      string
	Sweetness int
}

// LoadBottle loads raw data into an instance of Bottle running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadBottle(raw interface{}) (*Bottle, error) {
	var err error
	var res *Bottle
	if val, ok := raw.(map[string]interface{}); ok {
		res = new(Bottle)
		if v, ok := val["color"]; ok {
			var tmp1 string
			if val, ok := v.(string); ok {
				tmp1 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Color`, v, "string", err)
			}
			if err == nil {
				if tmp1 != "" {
					if !(tmp1 == "red" || tmp1 == "white" || tmp1 == "rose" || tmp1 == "yellow" || tmp1 == "sparkling") {
						err = goa.InvalidEnumValueError(`.Color`, tmp1, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
					}
				}
			}
			res.Color = tmp1
		}
		if v, ok := val["href"]; ok {
			var tmp2 string
			if val, ok := v.(string); ok {
				tmp2 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp2
		}
		if v, ok := val["id"]; ok {
			var tmp3 int
			if f, ok := v.(float64); ok {
				tmp3 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp3
		}
		if v, ok := val["name"]; ok {
			var tmp4 string
			if val, ok := v.(string); ok {
				tmp4 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp4) < 1 {
					err = goa.InvalidLengthError(`.Name`, tmp4, 1, true, err)
				}
				if len(tmp4) > 255 {
					err = goa.InvalidLengthError(`.Name`, tmp4, 255, false, err)
				}
			}
			res.Name = tmp4
		}
		if v, ok := val["sweetness"]; ok {
			var tmp5 int
			if f, ok := v.(float64); ok {
				tmp5 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.Sweetness`, v, "int", err)
			}
			if err == nil {
				if tmp5 < 1 {
					err = goa.InvalidRangeError(`.Sweetness`, tmp5, 1, true, err)
				}
				if tmp5 > 5 {
					err = goa.InvalidRangeError(`.Sweetness`, tmp5, 5, false, err)
				}
			}
			res.Sweetness = tmp5
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "dictionary", err)
	}
	return res, err
}

// Dump produces raw data from an instance of Bottle running all the
// validations. See LoadBottle for the definition of raw data.
func (mt *Bottle) Dump() (map[string]interface{}, error) {
	var err error
	var res map[string]interface{}
	if mt.Color != "" {
		if !(mt.Color == "red" || mt.Color == "white" || mt.Color == "rose" || mt.Color == "yellow" || mt.Color == "sparkling") {
			err = goa.InvalidEnumValueError(`default view.color`, mt.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
		}
	}
	if len(mt.Name) < 1 {
		err = goa.InvalidLengthError(`default view.name`, mt.Name, 1, true, err)
	}
	if len(mt.Name) > 255 {
		err = goa.InvalidLengthError(`default view.name`, mt.Name, 255, false, err)
	}
	if mt.Sweetness < 1 {
		err = goa.InvalidRangeError(`default view.sweetness`, mt.Sweetness, 1, true, err)
	}
	if mt.Sweetness > 5 {
		err = goa.InvalidRangeError(`default view.sweetness`, mt.Sweetness, 5, false, err)
	}
	tmp6 := map[string]interface{}{
		"color":     mt.Color,
		"href":      mt.Href,
		"id":        mt.ID,
		"name":      mt.Name,
		"sweetness": mt.Sweetness,
	}
	res = tmp6
	return res, err
}

// Validate validates the media type instance.
func (mt *Bottle) Validate() (err error) {
	if mt.Color != "" {
		if !(mt.Color == "red" || mt.Color == "white" || mt.Color == "rose" || mt.Color == "yellow" || mt.Color == "sparkling") {
			err = goa.InvalidEnumValueError(`response.color`, mt.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
		}
	}
	if len(mt.Name) < 1 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, 1, true, err)
	}
	if len(mt.Name) > 255 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, 255, false, err)
	}
	if mt.Sweetness < 1 {
		err = goa.InvalidRangeError(`response.sweetness`, mt.Sweetness, 1, true, err)
	}
	if mt.Sweetness > 5 {
		err = goa.InvalidRangeError(`response.sweetness`, mt.Sweetness, 5, false, err)
	}
	return
}
