// Code generated by go-swagger; DO NOT EDIT.

package blob

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBlobSidecarsByBlockNumHandlerFunc turns a function with the right signature into a get blob sidecars by block num handler
type GetBlobSidecarsByBlockNumHandlerFunc func(GetBlobSidecarsByBlockNumParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBlobSidecarsByBlockNumHandlerFunc) Handle(params GetBlobSidecarsByBlockNumParams) middleware.Responder {
	return fn(params)
}

// GetBlobSidecarsByBlockNumHandler interface for that can handle valid get blob sidecars by block num params
type GetBlobSidecarsByBlockNumHandler interface {
	Handle(GetBlobSidecarsByBlockNumParams) middleware.Responder
}

// NewGetBlobSidecarsByBlockNum creates a new http.Handler for the get blob sidecars by block num operation
func NewGetBlobSidecarsByBlockNum(ctx *middleware.Context, handler GetBlobSidecarsByBlockNumHandler) *GetBlobSidecarsByBlockNum {
	return &GetBlobSidecarsByBlockNum{Context: ctx, Handler: handler}
}

/*
	GetBlobSidecarsByBlockNum swagger:route GET /eth/v1/beacon/blob_sidecars/{block_id} blob getBlobSidecarsByBlockNum

Get blob sidecars by block num
*/
type GetBlobSidecarsByBlockNum struct {
	Context *middleware.Context
	Handler GetBlobSidecarsByBlockNumHandler
}

func (o *GetBlobSidecarsByBlockNum) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetBlobSidecarsByBlockNumParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}