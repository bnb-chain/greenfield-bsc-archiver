// Code generated by go-swagger; DO NOT EDIT.

package block

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBlockByNumberSimplifiedHandlerFunc turns a function with the right signature into a get block by number simplified handler
type GetBlockByNumberSimplifiedHandlerFunc func(GetBlockByNumberSimplifiedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBlockByNumberSimplifiedHandlerFunc) Handle(params GetBlockByNumberSimplifiedParams) middleware.Responder {
	return fn(params)
}

// GetBlockByNumberSimplifiedHandler interface for that can handle valid get block by number simplified params
type GetBlockByNumberSimplifiedHandler interface {
	Handle(GetBlockByNumberSimplifiedParams) middleware.Responder
}

// NewGetBlockByNumberSimplified creates a new http.Handler for the get block by number simplified operation
func NewGetBlockByNumberSimplified(ctx *middleware.Context, handler GetBlockByNumberSimplifiedHandler) *GetBlockByNumberSimplified {
	return &GetBlockByNumberSimplified{Context: ctx, Handler: handler}
}

/*
	GetBlockByNumberSimplified swagger:route POST // block getBlockByNumberSimplified

Returns information of the block matching the given block number.
*/
type GetBlockByNumberSimplified struct {
	Context *middleware.Context
	Handler GetBlockByNumberSimplifiedHandler
}

func (o *GetBlockByNumberSimplified) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetBlockByNumberSimplifiedParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
