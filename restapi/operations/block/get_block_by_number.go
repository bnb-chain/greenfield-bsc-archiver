// Code generated by go-swagger; DO NOT EDIT.

package block

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBlockByNumberHandlerFunc turns a function with the right signature into a get block by number handler
type GetBlockByNumberHandlerFunc func(GetBlockByNumberParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBlockByNumberHandlerFunc) Handle(params GetBlockByNumberParams) middleware.Responder {
	return fn(params)
}

// GetBlockByNumberHandler interface for that can handle valid get block by number params
type GetBlockByNumberHandler interface {
	Handle(GetBlockByNumberParams) middleware.Responder
}

// NewGetBlockByNumber creates a new http.Handler for the get block by number operation
func NewGetBlockByNumber(ctx *middleware.Context, handler GetBlockByNumberHandler) *GetBlockByNumber {
	return &GetBlockByNumber{Context: ctx, Handler: handler}
}

/*
	GetBlockByNumber swagger:route POST / block getBlockByNumber

Returns information of the block matching the given block number.
*/
type GetBlockByNumber struct {
	Context *middleware.Context
	Handler GetBlockByNumberHandler
}

func (o *GetBlockByNumber) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetBlockByNumberParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
