// Code generated by go-swagger; DO NOT EDIT.

package object_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteObjectOKCode is the HTTP code returned for type DeleteObjectOK
const DeleteObjectOKCode int = 200

/*DeleteObjectOK Ok.

swagger:response deleteObjectOK
*/
type DeleteObjectOK struct {
}

// NewDeleteObjectOK creates DeleteObjectOK with default headers values
func NewDeleteObjectOK() *DeleteObjectOK {

	return &DeleteObjectOK{}
}

// WriteResponse to the client
func (o *DeleteObjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteObjectBadRequestCode is the HTTP code returned for type DeleteObjectBadRequest
const DeleteObjectBadRequestCode int = 400

/*DeleteObjectBadRequest Bad request.

swagger:response deleteObjectBadRequest
*/
type DeleteObjectBadRequest struct {
}

// NewDeleteObjectBadRequest creates DeleteObjectBadRequest with default headers values
func NewDeleteObjectBadRequest() *DeleteObjectBadRequest {

	return &DeleteObjectBadRequest{}
}

// WriteResponse to the client
func (o *DeleteObjectBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteObjectUnauthorizedCode is the HTTP code returned for type DeleteObjectUnauthorized
const DeleteObjectUnauthorizedCode int = 401

/*DeleteObjectUnauthorized Not authenticated.

swagger:response deleteObjectUnauthorized
*/
type DeleteObjectUnauthorized struct {
}

// NewDeleteObjectUnauthorized creates DeleteObjectUnauthorized with default headers values
func NewDeleteObjectUnauthorized() *DeleteObjectUnauthorized {

	return &DeleteObjectUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteObjectUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// DeleteObjectForbiddenCode is the HTTP code returned for type DeleteObjectForbidden
const DeleteObjectForbiddenCode int = 403

/*DeleteObjectForbidden Forbidden.

swagger:response deleteObjectForbidden
*/
type DeleteObjectForbidden struct {
}

// NewDeleteObjectForbidden creates DeleteObjectForbidden with default headers values
func NewDeleteObjectForbidden() *DeleteObjectForbidden {

	return &DeleteObjectForbidden{}
}

// WriteResponse to the client
func (o *DeleteObjectForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DeleteObjectNotFoundCode is the HTTP code returned for type DeleteObjectNotFound
const DeleteObjectNotFoundCode int = 404

/*DeleteObjectNotFound Object not found.

swagger:response deleteObjectNotFound
*/
type DeleteObjectNotFound struct {
}

// NewDeleteObjectNotFound creates DeleteObjectNotFound with default headers values
func NewDeleteObjectNotFound() *DeleteObjectNotFound {

	return &DeleteObjectNotFound{}
}

// WriteResponse to the client
func (o *DeleteObjectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteObjectInternalServerErrorCode is the HTTP code returned for type DeleteObjectInternalServerError
const DeleteObjectInternalServerErrorCode int = 500

/*DeleteObjectInternalServerError Failed to delete object.

swagger:response deleteObjectInternalServerError
*/
type DeleteObjectInternalServerError struct {
}

// NewDeleteObjectInternalServerError creates DeleteObjectInternalServerError with default headers values
func NewDeleteObjectInternalServerError() *DeleteObjectInternalServerError {

	return &DeleteObjectInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteObjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

// DeleteObjectServiceUnavailableCode is the HTTP code returned for type DeleteObjectServiceUnavailable
const DeleteObjectServiceUnavailableCode int = 503

/*DeleteObjectServiceUnavailable Not available.

swagger:response deleteObjectServiceUnavailable
*/
type DeleteObjectServiceUnavailable struct {
}

// NewDeleteObjectServiceUnavailable creates DeleteObjectServiceUnavailable with default headers values
func NewDeleteObjectServiceUnavailable() *DeleteObjectServiceUnavailable {

	return &DeleteObjectServiceUnavailable{}
}

// WriteResponse to the client
func (o *DeleteObjectServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(503)
}