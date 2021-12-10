// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"golang_api/models"
)

// GetUsersIDOKCode is the HTTP code returned for type GetUsersIDOK
const GetUsersIDOKCode int = 200

/*GetUsersIDOK OK

swagger:response getUsersIdOK
*/
type GetUsersIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUsersIDOK creates GetUsersIDOK with default headers values
func NewGetUsersIDOK() *GetUsersIDOK {

	return &GetUsersIDOK{}
}

// WithPayload adds the payload to the get users Id o k response
func (o *GetUsersIDOK) WithPayload(payload *models.User) *GetUsersIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users Id o k response
func (o *GetUsersIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersIDUnauthorizedCode is the HTTP code returned for type GetUsersIDUnauthorized
const GetUsersIDUnauthorizedCode int = 401

/*GetUsersIDUnauthorized Unauthorized

swagger:response getUsersIdUnauthorized
*/
type GetUsersIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUsersIDUnauthorized creates GetUsersIDUnauthorized with default headers values
func NewGetUsersIDUnauthorized() *GetUsersIDUnauthorized {

	return &GetUsersIDUnauthorized{}
}

// WithPayload adds the payload to the get users Id unauthorized response
func (o *GetUsersIDUnauthorized) WithPayload(payload *models.Error) *GetUsersIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users Id unauthorized response
func (o *GetUsersIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersIDNotFoundCode is the HTTP code returned for type GetUsersIDNotFound
const GetUsersIDNotFoundCode int = 404

/*GetUsersIDNotFound Not found

swagger:response getUsersIdNotFound
*/
type GetUsersIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUsersIDNotFound creates GetUsersIDNotFound with default headers values
func NewGetUsersIDNotFound() *GetUsersIDNotFound {

	return &GetUsersIDNotFound{}
}

// WithPayload adds the payload to the get users Id not found response
func (o *GetUsersIDNotFound) WithPayload(payload *models.Error) *GetUsersIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users Id not found response
func (o *GetUsersIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}