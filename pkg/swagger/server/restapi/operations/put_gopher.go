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

// PutGopherHandlerFunc turns a function with the right signature into a put gopher handler
type PutGopherHandlerFunc func(PutGopherParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutGopherHandlerFunc) Handle(params PutGopherParams) middleware.Responder {
	return fn(params)
}

// PutGopherHandler interface for that can handle valid put gopher params
type PutGopherHandler interface {
	Handle(PutGopherParams) middleware.Responder
}

// NewPutGopher creates a new http.Handler for the put gopher operation
func NewPutGopher(ctx *middleware.Context, handler PutGopherHandler) *PutGopher {
	return &PutGopher{Context: ctx, Handler: handler}
}

/*
	PutGopher swagger:route PUT /gopher putGopher

Update a gopher
*/
type PutGopher struct {
	Context *middleware.Context
	Handler PutGopherHandler
}

func (o *PutGopher) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutGopherParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutGopherBody put gopher body
//
// swagger:model PutGopherBody
type PutGopherBody struct {

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

// Validate validates this put gopher body
func (o *PutGopherBody) Validate(formats strfmt.Registry) error {
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

func (o *PutGopherBody) validateDisplayname(formats strfmt.Registry) error {

	if err := validate.Required("gopher"+"."+"displayname", "body", o.Displayname); err != nil {
		return err
	}

	return nil
}

func (o *PutGopherBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("gopher"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *PutGopherBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("gopher"+"."+"url", "body", o.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put gopher body based on context it is used
func (o *PutGopherBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutGopherBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutGopherBody) UnmarshalBinary(b []byte) error {
	var res PutGopherBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
