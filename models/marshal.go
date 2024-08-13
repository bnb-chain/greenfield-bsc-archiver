package models

import "encoding/json"

func (m *Block) MarshalJSON() ([]byte, error) {
	type Alias Block
	if m.Withdrawals == nil {
		return json.Marshal(&struct {
			*Alias
			Withdrawals *[]*Withdrawal `json:"withdrawals,omitempty"`
		}{
			Alias: (*Alias)(m),
		})
	}

	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	})
}

func (m *SimplifiedBlock) MarshalJSON() ([]byte, error) {
	type Alias SimplifiedBlock
	if m.Withdrawals == nil {
		return json.Marshal(&struct {
			*Alias
			Withdrawals *[]*Withdrawal `json:"withdrawals,omitempty"`
		}{
			Alias: (*Alias)(m),
		})
	}

	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	})
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	type Alias Transaction
	if t.Type == "0x0" {
		return json.Marshal(&struct {
			*Alias
			AccessList *[]*AccessTuple `json:"accessList,omitempty"`
		}{
			Alias: (*Alias)(t),
		})
	}

	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	})
}
