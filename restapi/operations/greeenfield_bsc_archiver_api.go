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

	"greeenfield-bsc-archiver/restapi/operations/block"
)

// NewGreeenfieldBscArchiverAPI creates a new GreeenfieldBscArchiver instance
func NewGreeenfieldBscArchiverAPI(spec *loads.Document) *GreeenfieldBscArchiverAPI {
	return &GreeenfieldBscArchiverAPI{
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

		BlockGetBlockByHashHandler: block.GetBlockByHashHandlerFunc(func(params block.GetBlockByHashParams) middleware.Responder {
			return middleware.NotImplemented("operation block.GetBlockByHash has not yet been implemented")
		}),
		BlockGetBlockByNumberHandler: block.GetBlockByNumberHandlerFunc(func(params block.GetBlockByNumberParams) middleware.Responder {
			return middleware.NotImplemented("operation block.GetBlockByNumber has not yet been implemented")
		}),
		BlockGetBlockByNumberSimplifiedHandler: block.GetBlockByNumberSimplifiedHandlerFunc(func(params block.GetBlockByNumberSimplifiedParams) middleware.Responder {
			return middleware.NotImplemented("operation block.GetBlockByNumberSimplified has not yet been implemented")
		}),
		BlockGetBlockNumberHandler: block.GetBlockNumberHandlerFunc(func(params block.GetBlockNumberParams) middleware.Responder {
			return middleware.NotImplemented("operation block.GetBlockNumber has not yet been implemented")
		}),
		BlockGetBundleNameByBlockNumberHandler: block.GetBundleNameByBlockNumberHandlerFunc(func(params block.GetBundleNameByBlockNumberParams) middleware.Responder {
			return middleware.NotImplemented("operation block.GetBundleNameByBlockNumber has not yet been implemented")
		}),
		BlockGetBundledBlockByNumberHandler: block.GetBundledBlockByNumberHandlerFunc(func(params block.GetBundledBlockByNumberParams) middleware.Responder {
			return middleware.NotImplemented("operation block.GetBundledBlockByNumber has not yet been implemented")
		}),
	}
}

/*GreeenfieldBscArchiverAPI API for handling block query in the Greenfield BSC Archiver. */
type GreeenfieldBscArchiverAPI struct {
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

	// BlockGetBlockByHashHandler sets the operation handler for the get block by hash operation
	BlockGetBlockByHashHandler block.GetBlockByHashHandler
	// BlockGetBlockByNumberHandler sets the operation handler for the get block by number operation
	BlockGetBlockByNumberHandler block.GetBlockByNumberHandler
	// BlockGetBlockByNumberSimplifiedHandler sets the operation handler for the get block by number simplified operation
	BlockGetBlockByNumberSimplifiedHandler block.GetBlockByNumberSimplifiedHandler
	// BlockGetBlockNumberHandler sets the operation handler for the get block number operation
	BlockGetBlockNumberHandler block.GetBlockNumberHandler
	// BlockGetBundleNameByBlockNumberHandler sets the operation handler for the get bundle name by block number operation
	BlockGetBundleNameByBlockNumberHandler block.GetBundleNameByBlockNumberHandler
	// BlockGetBundledBlockByNumberHandler sets the operation handler for the get bundled block by number operation
	BlockGetBundledBlockByNumberHandler block.GetBundledBlockByNumberHandler

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
func (o *GreeenfieldBscArchiverAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *GreeenfieldBscArchiverAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *GreeenfieldBscArchiverAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *GreeenfieldBscArchiverAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *GreeenfieldBscArchiverAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *GreeenfieldBscArchiverAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *GreeenfieldBscArchiverAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *GreeenfieldBscArchiverAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *GreeenfieldBscArchiverAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the GreeenfieldBscArchiverAPI
func (o *GreeenfieldBscArchiverAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BlockGetBlockByHashHandler == nil {
		unregistered = append(unregistered, "block.GetBlockByHashHandler")
	}
	if o.BlockGetBlockByNumberHandler == nil {
		unregistered = append(unregistered, "block.GetBlockByNumberHandler")
	}
	if o.BlockGetBlockByNumberSimplifiedHandler == nil {
		unregistered = append(unregistered, "block.GetBlockByNumberSimplifiedHandler")
	}
	if o.BlockGetBlockNumberHandler == nil {
		unregistered = append(unregistered, "block.GetBlockNumberHandler")
	}
	if o.BlockGetBundleNameByBlockNumberHandler == nil {
		unregistered = append(unregistered, "block.GetBundleNameByBlockNumberHandler")
	}
	if o.BlockGetBundledBlockByNumberHandler == nil {
		unregistered = append(unregistered, "block.GetBundledBlockByNumberHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *GreeenfieldBscArchiverAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *GreeenfieldBscArchiverAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *GreeenfieldBscArchiverAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *GreeenfieldBscArchiverAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *GreeenfieldBscArchiverAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *GreeenfieldBscArchiverAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the greeenfield bsc archiver API
func (o *GreeenfieldBscArchiverAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *GreeenfieldBscArchiverAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/eth_getBlockByHash"] = block.NewGetBlockByHash(o.context, o.BlockGetBlockByHashHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"][""] = block.NewGetBlockByNumber(o.context, o.BlockGetBlockByNumberHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/"] = block.NewGetBlockByNumberSimplified(o.context, o.BlockGetBlockByNumberSimplifiedHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/eth_blockNumber"] = block.NewGetBlockNumber(o.context, o.BlockGetBlockNumberHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/bsc/v1/blocks/{block_number}/bundle/name"] = block.NewGetBundleNameByBlockNumber(o.context, o.BlockGetBundleNameByBlockNumberHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/eth_getBundledBlockByNumber"] = block.NewGetBundledBlockByNumber(o.context, o.BlockGetBundledBlockByNumberHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *GreeenfieldBscArchiverAPI) Serve(builder middleware.Builder) http.Handler {
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
func (o *GreeenfieldBscArchiverAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *GreeenfieldBscArchiverAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *GreeenfieldBscArchiverAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *GreeenfieldBscArchiverAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
