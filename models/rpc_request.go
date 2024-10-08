// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RPCRequest RPC request
//
// swagger:model RPCRequest
type RPCRequest struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// jsonrpc
	// Example: 2.0
	Jsonrpc string `json:"jsonrpc,omitempty"`

	// method
	// Example: eth_getBlockByNumber
	Method string `json:"method,omitempty"`

	// params
	// Example: ["0x1",true]
	Params []string `json:"params"`
}

// Validate validates this RPC request
func (m *RPCRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this RPC request based on context it is used
func (m *RPCRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RPCRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RPCRequest) UnmarshalBinary(b []byte) error {
	var res RPCRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
