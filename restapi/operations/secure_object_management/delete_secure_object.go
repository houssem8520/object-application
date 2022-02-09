// Code generated by go-swagger; DO NOT EDIT.

package secure_object_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/object.application/models"
)

// DeleteSecureObjectHandlerFunc turns a function with the right signature into a delete secure object handler
type DeleteSecureObjectHandlerFunc func(DeleteSecureObjectParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSecureObjectHandlerFunc) Handle(params DeleteSecureObjectParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteSecureObjectHandler interface for that can handle valid delete secure object params
type DeleteSecureObjectHandler interface {
	Handle(DeleteSecureObjectParams, *models.Principal) middleware.Responder
}

// NewDeleteSecureObject creates a new http.Handler for the delete secure object operation
func NewDeleteSecureObject(ctx *middleware.Context, handler DeleteSecureObjectHandler) *DeleteSecureObject {
	return &DeleteSecureObject{Context: ctx, Handler: handler}
}

/* DeleteSecureObject swagger:route DELETE /secure/objects/{bucket}/{objectID} Secure Object Management deleteSecureObject

delete an object.

*/
type DeleteSecureObject struct {
	Context *middleware.Context
	Handler DeleteSecureObjectHandler
}

func (o *DeleteSecureObject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteSecureObjectParams()
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