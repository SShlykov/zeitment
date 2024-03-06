package types

import "encoding/json"

type Null[T any] struct {
	Value T
	Valid bool
}

func (n *Null[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	n.Value = v
	n.Valid = true
	return nil
}

func (n *Null[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value)
}
