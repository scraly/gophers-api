// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostGopherHandlerFunc turns a function with the right signature into a post gopher handler
type PostGopherHandlerFunc func(PostGopherParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostGopherHandlerFunc) Handle(params PostGopherParams) middleware.Responder {
	return fn(params)
}

// PostGopherHandler interface for that can handle valid post gopher params
type PostGopherHandler interface {
	Handle(PostGopherParams) middleware.Responder
}

// NewPostGopher creates a new http.Handler for the post gopher operation
func NewPostGopher(ctx *middleware.Context, handler PostGopherHandler) *PostGopher {
	return &PostGopher{Context: ctx, Handler: handler}
}

/*
	PostGopher swagger:route POST /gopher postGopher

Add a new Gopher
*/
type PostGopher struct {
	Context *middleware.Context
	Handler PostGopherHandler
}

func (o *PostGopher) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostGopherParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostGopherBody post gopher body
//
// swagger:model PostGopherBody
type PostGopherBody struct {

	// displayname
	// Required: true
	Displayname *string `json:"displayname"`

	// name
	// Required: true
	Name *string `json:"name"`

	// url
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this post gopher body
func (o *PostGopherBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDisplayname(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGopherBody) validateDisplayname(formats strfmt.Registry) error {

	if err := validate.Required("gopher"+"."+"displayname", "body", o.Displayname); err != nil {
		return err
	}

	return nil
}

func (o *PostGopherBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("gopher"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *PostGopherBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("gopher"+"."+"url", "body", o.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post gopher body based on context it is used
func (o *PostGopherBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostGopherBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGopherBody) UnmarshalBinary(b []byte) error {
	var res PostGopherBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
