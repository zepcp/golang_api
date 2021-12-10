// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostUsersLoginHandlerFunc turns a function with the right signature into a post users login handler
type PostUsersLoginHandlerFunc func(PostUsersLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUsersLoginHandlerFunc) Handle(params PostUsersLoginParams) middleware.Responder {
	return fn(params)
}

// PostUsersLoginHandler interface for that can handle valid post users login params
type PostUsersLoginHandler interface {
	Handle(PostUsersLoginParams) middleware.Responder
}

// NewPostUsersLogin creates a new http.Handler for the post users login operation
func NewPostUsersLogin(ctx *middleware.Context, handler PostUsersLoginHandler) *PostUsersLogin {
	return &PostUsersLogin{Context: ctx, Handler: handler}
}

/*PostUsersLogin swagger:route POST /users/login users postUsersLogin

Login

Login

*/
type PostUsersLogin struct {
	Context *middleware.Context
	Handler PostUsersLoginHandler
}

func (o *PostUsersLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostUsersLoginParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}