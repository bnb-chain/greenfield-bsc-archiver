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
