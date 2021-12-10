// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBooksHandlerFunc turns a function with the right signature into a get books handler
type GetBooksHandlerFunc func(GetBooksParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBooksHandlerFunc) Handle(params GetBooksParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetBooksHandler interface for that can handle valid get books params
type GetBooksHandler interface {
	Handle(GetBooksParams, interface{}) middleware.Responder
}

// NewGetBooks creates a new http.Handler for the get books operation
func NewGetBooks(ctx *middleware.Context, handler GetBooksHandler) *GetBooks {
	return &GetBooks{Context: ctx, Handler: handler}
}

/*GetBooks swagger:route GET /books books getBooks

FetchBooks

FetchBooks

*/
type GetBooks struct {
	Context *middleware.Context
	Handler GetBooksHandler
}

func (o *GetBooks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBooksParams()

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
