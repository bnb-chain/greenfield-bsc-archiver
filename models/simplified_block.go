// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SimplifiedBlock simplified block
//
// swagger:model simplifiedBlock
type SimplifiedBlock struct {

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

	// hash
	// Example: 0x3757a8cc692ec0ffe09c93d147e690fbc2dd8e2a94ca94050059adb6e9b7b1fc
	Hash string `json:"hash,omitempty"`

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

	// size
	// Example: 0x176d20b
	Size string `json:"size,omitempty"`

	// state root
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	StateRoot string `json:"stateRoot,omitempty"`

	// timestamp
	// Example: 0x5f4e5f87
	Timestamp string `json:"timestamp,omitempty"`

	// total difficulty
	// Example: 0x176d20b
	TotalDifficulty *string `json:"totalDifficulty,omitempty"`

	// transactions
	Transactions []string `json:"transactions"`

	// transactions root
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	TransactionsRoot string `json:"transactionsRoot,omitempty"`

	// uncles
	Uncles []*Header `json:"uncles"`

	// withdrawals
	Withdrawals []*Withdrawal `json:"withdrawals"`

	// withdrawals root
	// Example: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	WithdrawalsRoot *string `json:"withdrawalsRoot,omitempty"`
}

// Validate validates this simplified block
func (m *SimplifiedBlock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUncles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWithdrawals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimplifiedBlock) validateUncles(formats strfmt.Registry) error {
	if swag.IsZero(m.Uncles) { // not required
		return nil
	}

	for i := 0; i < len(m.Uncles); i++ {
		if swag.IsZero(m.Uncles[i]) { // not required
			continue
		}

		if m.Uncles[i] != nil {
			if err := m.Uncles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("uncles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("uncles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SimplifiedBlock) validateWithdrawals(formats strfmt.Registry) error {
	if swag.IsZero(m.Withdrawals) { // not required
		return nil
	}

	for i := 0; i < len(m.Withdrawals); i++ {
		if swag.IsZero(m.Withdrawals[i]) { // not required
			continue
		}

		if m.Withdrawals[i] != nil {
			if err := m.Withdrawals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("withdrawals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("withdrawals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this simplified block based on the context it is used
func (m *SimplifiedBlock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUncles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWithdrawals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimplifiedBlock) contextValidateUncles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Uncles); i++ {

		if m.Uncles[i] != nil {

			if swag.IsZero(m.Uncles[i]) { // not required
				return nil
			}

			if err := m.Uncles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("uncles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("uncles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SimplifiedBlock) contextValidateWithdrawals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Withdrawals); i++ {

		if m.Withdrawals[i] != nil {

			if swag.IsZero(m.Withdrawals[i]) { // not required
				return nil
			}

			if err := m.Withdrawals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("withdrawals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("withdrawals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimplifiedBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimplifiedBlock) UnmarshalBinary(b []byte) error {
	var res SimplifiedBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
