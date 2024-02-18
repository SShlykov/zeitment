package models

import (
	"database/sql"
	"encoding/json"
)

// NullTime обертка над sql.NullTime для кастомной сериализации в JSON.
type NullTime struct {
	sql.NullTime
}

// MarshalJSON кастомная сериализация CustomNullTime в JSON.
func (cnt *NullTime) MarshalJSON() ([]byte, error) {
	if cnt.Valid {
		return json.Marshal(cnt.Time)
	}
	return json.Marshal(nil)
}

func (cnt *NullTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		cnt.Valid = false
		return nil
	}

	err := json.Unmarshal(data, &cnt.Time)
	if err != nil {
		return err
	}
	cnt.Valid = true
	return nil
}
