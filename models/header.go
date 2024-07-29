// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Header header
//
// swagger:model header
type Header struct {

	// base fee per gas
	// Example: 1000000000
	BaseFeePerGas *string `json:"baseFeePerGas,omitempty"`

	// blob gas used
	// Example: 1000
	BlobGasUsed *string `json:"blobGasUsed,omitempty"`

	// difficulty
	// Example: 0x186a1
	Difficulty *string `json:"difficulty,omitempty"`

	// excess blob gas
	// Example: 500
	ExcessBlobGas *string `json:"excessBlobGas,omitempty"`

	// extra data
	// Example: 0x123456
	ExtraData string `json:"extraData,omitempty"`

	// gas limit
	// Example: 8000000
	GasLimit string `json:"gasLimit,omitempty"`

	// gas used
	// Example: 21000
	GasUsed string `json:"gasUsed,omitempty"`

	// logs bloom
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	LogsBloom string `json:"logsBloom,omitempty"`

	// miner
	// Example: 0x1234567890abcdef1234567890abcdef12345678
	Miner string `json:"miner,omitempty"`

	// mix hash
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	MixHash string `json:"mixHash,omitempty"`

	// nonce
	// Example: 0x0000000000000042
	Nonce string `json:"nonce,omitempty"`

	// number
	// Example: 100000
	Number *string `json:"number,omitempty"`

	// parent beacon block root
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	ParentBeaconBlockRoot *string `json:"parentBeaconBlockRoot,omitempty"`

	// parent hash
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	ParentHash string `json:"parentHash,omitempty"`

	// receipts root
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	ReceiptsRoot string `json:"receiptsRoot,omitempty"`

	// sha3 uncles
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	Sha3Uncles string `json:"sha3Uncles,omitempty"`

	// state root
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	StateRoot string `json:"stateRoot,omitempty"`

	// timestamp
	// Example: 0x5f4e5f87
	Timestamp string `json:"timestamp,omitempty"`

	// transactions root
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	TransactionsRoot string `json:"transactionsRoot,omitempty"`

	// withdrawals root
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	WithdrawalsRoot *string `json:"withdrawalsRoot,omitempty"`
}

// Validate validates this header
func (m *Header) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this header based on context it is used
func (m *Header) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Header) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Header) UnmarshalBinary(b []byte) error {
	var res Header
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
