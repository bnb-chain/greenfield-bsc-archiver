// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccessTuple access tuple
//
// swagger:model accessTuple
type AccessTuple struct {

	// address
	// Example: 0x1234567890abcdef1234567890abcdef12345678
	Address string `json:"address,omitempty"`

	// storage keys
	StorageKeys []string `json:"storageKeys"`
}

// Validate validates this access tuple
func (m *AccessTuple) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access tuple based on context it is used
func (m *AccessTuple) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessTuple) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessTuple) UnmarshalBinary(b []byte) error {
	var res AccessTuple
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}