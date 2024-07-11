// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Transaction transaction
//
// swagger:model Transaction
type Transaction struct {

	// from
	// Example: 0x1234567890abcdef1234567890abcdef12345678
	From string `json:"from,omitempty"`

	// hash
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	Hash string `json:"hash,omitempty"`

	// to
	// Example: 0xabcdef1234567890abcdef1234567890abcdef12
	To string `json:"to,omitempty"`

	// value
	// Example: 1000000000
	Value string `json:"value,omitempty"`
}

// Validate validates this transaction
func (m *Transaction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this transaction based on context it is used
func (m *Transaction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Transaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transaction) UnmarshalBinary(b []byte) error {
	var res Transaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
