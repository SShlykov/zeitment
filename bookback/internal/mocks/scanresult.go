package mocks

import (
	"errors"
	"reflect"
)

type ScanResult struct {
	Row []interface{} `json:"row,omitempty"` //nolint:gofmt
}

func NewScanResult(row []interface{}) *ScanResult {
	return &ScanResult{Row: row}
}

// Scan assigns values from the ScanResult's row to the provided dest arguments.
// It uses reflection to safely assign the values, supporting a wide range of types.
func (s *ScanResult) Scan(dest ...interface{}) error {
	if len(dest) != len(s.Row) {
		return errors.New("the number of dest does not match the number of columns in row")
	}

	for i, d := range dest {
		if reflect.TypeOf(d).Kind() != reflect.Ptr {
			return errors.New("dest must be a pointer")
		}

		val := reflect.ValueOf(d).Elem()

		if !val.CanSet() {
			return errors.New("cannot set the value of dest")
		}

		newVal := reflect.ValueOf(s.Row[i])
		if newVal.IsValid() && val.Type() == newVal.Type() {
			val.Set(newVal)
		} else {
			if newVal.Type().ConvertibleTo(val.Type()) {
				val.Set(newVal.Convert(val.Type()))
			} else {
				return errors.New("type mismatch between dest and row values")
			}
		}
	}

	return nil
}
