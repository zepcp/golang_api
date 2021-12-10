// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"golang_api/restapi/operations/books"
	"golang_api/restapi/operations/users"
)

// NewGoLangAPIAPI creates a new GoLangAPI instance
func NewGoLangAPIAPI(spec *loads.Document) *GoLangAPIAPI {
	return &GoLangAPIAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		BooksDeleteBooksIDHandler: books.DeleteBooksIDHandlerFunc(func(params books.DeleteBooksIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation books.DeleteBooksID has not yet been implemented")
		}),
		UsersDeleteUsersIDHandler: users.DeleteUsersIDHandlerFunc(func(params users.DeleteUsersIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.DeleteUsersID has not yet been implemented")
		}),
		BooksGetBooksHandler: books.GetBooksHandlerFunc(func(params books.GetBooksParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation books.GetBooks has not yet been implemented")
		}),
		BooksGetBooksIDHandler: books.GetBooksIDHandlerFunc(func(params books.GetBooksIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation books.GetBooksID has not yet been implemented")
		}),
		UsersGetUsersHandler: users.GetUsersHandlerFunc(func(params users.GetUsersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
		}),
		UsersGetUsersIDHandler: users.GetUsersIDHandlerFunc(func(params users.GetUsersIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsersID has not yet been implemented")
		}),
		BooksPatchBooksIDHandler: books.PatchBooksIDHandlerFunc(func(params books.PatchBooksIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation books.PatchBooksID has not yet been implemented")
		}),
		UsersPatchUsersIDHandler: users.PatchUsersIDHandlerFunc(func(params users.PatchUsersIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.PatchUsersID has not yet been implemented")
		}),
		BooksPostBooksHandler: books.PostBooksHandlerFunc(func(params books.PostBooksParams) middleware.Responder {
			return middleware.NotImplemented("operation books.PostBooks has not yet been implemented")
		}),
		UsersPostUsersHandler: users.PostUsersHandlerFunc(func(params users.PostUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUsers has not yet been implemented")
		}),
		UsersPostUsersLoginHandler: users.PostUsersLoginHandlerFunc(func(params users.PostUsersLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUsersLogin has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*GoLangAPIAPI Go-Lang API */
type GoLangAPIAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// BooksDeleteBooksIDHandler sets the operation handler for the delete books ID operation
	BooksDeleteBooksIDHandler books.DeleteBooksIDHandler
	// UsersDeleteUsersIDHandler sets the operation handler for the delete users ID operation
	UsersDeleteUsersIDHandler users.DeleteUsersIDHandler
	// BooksGetBooksHandler sets the operation handler for the get books operation
	BooksGetBooksHandler books.GetBooksHandler
	// BooksGetBooksIDHandler sets the operation handler for the get books ID operation
	BooksGetBooksIDHandler books.GetBooksIDHandler
	// UsersGetUsersHandler sets the operation handler for the get users operation
	UsersGetUsersHandler users.GetUsersHandler
	// UsersGetUsersIDHandler sets the operation handler for the get users ID operation
	UsersGetUsersIDHandler users.GetUsersIDHandler
	// BooksPatchBooksIDHandler sets the operation handler for the patch books ID operation
	BooksPatchBooksIDHandler books.PatchBooksIDHandler
	// UsersPatchUsersIDHandler sets the operation handler for the patch users ID operation
	UsersPatchUsersIDHandler users.PatchUsersIDHandler
	// BooksPostBooksHandler sets the operation handler for the post books operation
	BooksPostBooksHandler books.PostBooksHandler
	// UsersPostUsersHandler sets the operation handler for the post users operation
	UsersPostUsersHandler users.PostUsersHandler
	// UsersPostUsersLoginHandler sets the operation handler for the post users login operation
	UsersPostUsersLoginHandler users.PostUsersLoginHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *GoLangAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *GoLangAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *GoLangAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *GoLangAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *GoLangAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *GoLangAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *GoLangAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *GoLangAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *GoLangAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the GoLangAPIAPI
func (o *GoLangAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.BooksDeleteBooksIDHandler == nil {
		unregistered = append(unregistered, "books.DeleteBooksIDHandler")
	}
	if o.UsersDeleteUsersIDHandler == nil {
		unregistered = append(unregistered, "users.DeleteUsersIDHandler")
	}
	if o.BooksGetBooksHandler == nil {
		unregistered = append(unregistered, "books.GetBooksHandler")
	}
	if o.BooksGetBooksIDHandler == nil {
		unregistered = append(unregistered, "books.GetBooksIDHandler")
	}
	if o.UsersGetUsersHandler == nil {
		unregistered = append(unregistered, "users.GetUsersHandler")
	}
	if o.UsersGetUsersIDHandler == nil {
		unregistered = append(unregistered, "users.GetUsersIDHandler")
	}
	if o.BooksPatchBooksIDHandler == nil {
		unregistered = append(unregistered, "books.PatchBooksIDHandler")
	}
	if o.UsersPatchUsersIDHandler == nil {
		unregistered = append(unregistered, "users.PatchUsersIDHandler")
	}
	if o.BooksPostBooksHandler == nil {
		unregistered = append(unregistered, "books.PostBooksHandler")
	}
	if o.UsersPostUsersHandler == nil {
		unregistered = append(unregistered, "users.PostUsersHandler")
	}
	if o.UsersPostUsersLoginHandler == nil {
		unregistered = append(unregistered, "users.PostUsersLoginHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *GoLangAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *GoLangAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "Bearer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.BearerAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *GoLangAPIAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *GoLangAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *GoLangAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *GoLangAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the go lang API API
func (o *GoLangAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *GoLangAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/books/{id}"] = books.NewDeleteBooksID(o.context, o.BooksDeleteBooksIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{id}"] = users.NewDeleteUsersID(o.context, o.UsersDeleteUsersIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/books"] = books.NewGetBooks(o.context, o.BooksGetBooksHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/books/{id}"] = books.NewGetBooksID(o.context, o.BooksGetBooksIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = users.NewGetUsers(o.context, o.UsersGetUsersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{id}"] = users.NewGetUsersID(o.context, o.UsersGetUsersIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/books/{id}"] = books.NewPatchBooksID(o.context, o.BooksPatchBooksIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/users/{id}"] = users.NewPatchUsersID(o.context, o.UsersPatchUsersIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/books"] = books.NewPostBooks(o.context, o.BooksPostBooksHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = users.NewPostUsers(o.context, o.UsersPostUsersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/login"] = users.NewPostUsersLogin(o.context, o.UsersPostUsersLoginHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *GoLangAPIAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *GoLangAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *GoLangAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *GoLangAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *GoLangAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}