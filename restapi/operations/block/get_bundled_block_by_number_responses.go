// Code generated by go-swagger; DO NOT EDIT.

package block

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"greeenfield-bsc-archiver/models"
)

// GetBundledBlockByNumberOKCode is the HTTP code returned for type GetBundledBlockByNumberOK
const GetBundledBlockByNumberOKCode int = 200

/*
GetBundledBlockByNumberOK successful operation

swagger:response getBundledBlockByNumberOK
*/
type GetBundledBlockByNumberOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetBundledBlockByNumberRPCResponse `json:"body,omitempty"`
}

// NewGetBundledBlockByNumberOK creates GetBundledBlockByNumberOK with default headers values
func NewGetBundledBlockByNumberOK() *GetBundledBlockByNumberOK {

	return &GetBundledBlockByNumberOK{}
}

// WithPayload adds the payload to the get bundled block by number o k response
func (o *GetBundledBlockByNumberOK) WithPayload(payload *models.GetBundledBlockByNumberRPCResponse) *GetBundledBlockByNumberOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bundled block by number o k response
func (o *GetBundledBlockByNumberOK) SetPayload(payload *models.GetBundledBlockByNumberRPCResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBundledBlockByNumberOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBundledBlockByNumberInternalServerErrorCode is the HTTP code returned for type GetBundledBlockByNumberInternalServerError
const GetBundledBlockByNumberInternalServerErrorCode int = 500

/*
GetBundledBlockByNumberInternalServerError internal server error

swagger:response getBundledBlockByNumberInternalServerError
*/
type GetBundledBlockByNumberInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBundledBlockByNumberInternalServerError creates GetBundledBlockByNumberInternalServerError with default headers values
func NewGetBundledBlockByNumberInternalServerError() *GetBundledBlockByNumberInternalServerError {

	return &GetBundledBlockByNumberInternalServerError{}
}

// WithPayload adds the payload to the get bundled block by number internal server error response
func (o *GetBundledBlockByNumberInternalServerError) WithPayload(payload *models.Error) *GetBundledBlockByNumberInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bundled block by number internal server error response
func (o *GetBundledBlockByNumberInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBundledBlockByNumberInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
