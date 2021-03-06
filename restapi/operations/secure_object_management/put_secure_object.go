// Code generated by go-swagger; DO NOT EDIT.

package secure_object_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/object.application/models"
)

// PutSecureObjectHandlerFunc turns a function with the right signature into a put secure object handler
type PutSecureObjectHandlerFunc func(PutSecureObjectParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PutSecureObjectHandlerFunc) Handle(params PutSecureObjectParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PutSecureObjectHandler interface for that can handle valid put secure object params
type PutSecureObjectHandler interface {
	Handle(PutSecureObjectParams, *models.Principal) middleware.Responder
}

// NewPutSecureObject creates a new http.Handler for the put secure object operation
func NewPutSecureObject(ctx *middleware.Context, handler PutSecureObjectHandler) *PutSecureObject {
	return &PutSecureObject{Context: ctx, Handler: handler}
}

/* PutSecureObject swagger:route PUT /secure/objects/{bucket} Secure Object Management putSecureObject

Add an object in a bucket.

*/
type PutSecureObject struct {
	Context *middleware.Context
	Handler PutSecureObjectHandler
}

func (o *PutSecureObject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutSecureObjectParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutSecureObjectCreatedBody put secure object created body
//
// swagger:model PutSecureObjectCreatedBody
type PutSecureObjectCreatedBody struct {

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this put secure object created body
func (o *PutSecureObjectCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put secure object created body based on context it is used
func (o *PutSecureObjectCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutSecureObjectCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutSecureObjectCreatedBody) UnmarshalBinary(b []byte) error {
	var res PutSecureObjectCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
