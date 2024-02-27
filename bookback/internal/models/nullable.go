package models

import (
	"database/sql"
	"encoding/json"
)

// NullTime обертка вокруг sql.NullTime для корректной работы с JSON
type NullTime struct {
	sql.NullTime
}

// MarshalJSON метод для сериализация в JSON.
// Возвращает значение в формате JSON или null.
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON метод для десериализации из JSON.
func (nt *NullTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nt.Valid = false
		return nil
	}
	err := json.Unmarshal(data, &nt.Time)
	nt.Valid = err == nil
	return err
}

// NullString обертка вокруг sql.NullString для корректной работы с JSON
type NullString struct {
	sql.NullString
}

// MarshalJSON метод для сериализация в JSON.
// Возвращает значение в формате JSON или null.
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON метод для десериализации из JSON.
func (ns *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ns.Valid = false
		return nil
	}
	err := json.Unmarshal(data, &ns.String)
	ns.Valid = err == nil
	return err
}
