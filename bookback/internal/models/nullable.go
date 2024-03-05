package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

// NullTime обертка вокруг sql.NullTime для корректной работы с JSON
type NullTime struct {
	sql.NullTime
}

// NewNullTime создает новый экземпляр NullTime.
func NewNullTime(t time.Time, valid bool) NullTime {
	return NullTime{sql.NullTime{Time: t, Valid: valid}}
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
	if isNull(data) {
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

// NewNullString создает новый экземпляр NullString.
func NewNullString(str string, valid bool) NullString {
	return NullString{sql.NullString{String: str, Valid: valid}}
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
	if isNull(data) {
		ns.Valid = false
		return nil
	}
	err := json.Unmarshal(data, &ns.String)
	ns.Valid = err == nil
	return err
}

// NullInt обертка вокруг sql.NullInt64 для корректной работы с JSON
type NullInt struct {
	sql.NullInt64
}

// NewNullInt создает новый экземпляр NullInt.
func NewNullInt(i int64, valid bool) NullInt {
	return NullInt{sql.NullInt64{Int64: i, Valid: valid}}
}

// MarshalJSON метод для сериализация в JSON.
func (ni *NullInt) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int64)
	}
	return json.Marshal(nil)
}

func (ni *NullInt) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		ni.Valid = false
		return nil
	}
	err := json.Unmarshal(data, &ni.Int64)
	ni.Valid = err == nil
	return err
}

func isNull(data []byte) bool {
	return string(data) == "null"
}
