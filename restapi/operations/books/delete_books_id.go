// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteBooksIDHandlerFunc turns a function with the right signature into a delete books ID handler
type DeleteBooksIDHandlerFunc func(DeleteBooksIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBooksIDHandlerFunc) Handle(params DeleteBooksIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteBooksIDHandler interface for that can handle valid delete books ID params
type DeleteBooksIDHandler interface {
	Handle(DeleteBooksIDParams, interface{}) middleware.Responder
}

// NewDeleteBooksID creates a new http.Handler for the delete books ID operation
func NewDeleteBooksID(ctx *middleware.Context, handler DeleteBooksIDHandler) *DeleteBooksID {
	return &DeleteBooksID{Context: ctx, Handler: handler}
}

/*DeleteBooksID swagger:route DELETE /books/{id} books deleteBooksId

DeleteBook by ID

DeleteBook by ID

*/
type DeleteBooksID struct {
	Context *middleware.Context
	Handler DeleteBooksIDHandler
}

func (o *DeleteBooksID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteBooksIDParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}