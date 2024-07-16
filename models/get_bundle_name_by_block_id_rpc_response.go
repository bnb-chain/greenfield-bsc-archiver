// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetBundleNameByBlockIDRPCResponse get bundle name by block ID RPC response
//
// swagger:model GetBundleNameByBlockIDRPCResponse
type GetBundleNameByBlockIDRPCResponse struct {

	// status code
	// Example: 200
	Code string `json:"code,omitempty"`

	// actual data for request
	// Example: 1
	Data string `json:"data,omitempty"`

	// error message if there is error
	// Example: signature invalid
	Message string `json:"message,omitempty"`
}

// Validate validates this get bundle name by block ID RPC response
func (m *GetBundleNameByBlockIDRPCResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get bundle name by block ID RPC response based on context it is used
func (m *GetBundleNameByBlockIDRPCResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetBundleNameByBlockIDRPCResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBundleNameByBlockIDRPCResponse) UnmarshalBinary(b []byte) error {
	var res GetBundleNameByBlockIDRPCResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}