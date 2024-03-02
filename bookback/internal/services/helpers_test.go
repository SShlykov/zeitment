package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockAllItems() string {
	return "*"
}

func TestDeleteQuery(t *testing.T) {
	tests := []struct {
		name      string
		tableName string
		idName    string
		want      string
	}{
		{
			name:      "Single WHERE condition",
			tableName: "books",
			idName:    "id",
			want:      "DELETE FROM books WHERE id = $1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteQuery(tt.tableName, tt.idName)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSelectWhere(t *testing.T) {
	tests := []struct {
		name      string
		allItems  func() string
		tableName string
		args      []string
		want      string
	}{
		{
			name:      "No WHERE conditions",
			allItems:  mockAllItems,
			tableName: "books",
			args:      nil,
			want:      "SELECT * FROM books",
		},
		{
			name:      "Single WHERE condition",
			allItems:  mockAllItems,
			tableName: "books",
			args:      []string{"id"},
			want:      "SELECT * FROM books WHERE id = $1",
		},
		{
			name:      "Multiple WHERE conditions",
			allItems:  mockAllItems,
			tableName: "books",
			args:      []string{"id", "title"},
			want:      "SELECT * FROM books WHERE id = $1 AND title = $2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectWhere(tt.allItems, tt.tableName, tt.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParamsToQuery(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Single parameter",
			args: []string{"id"},
			want: "id = $1",
		},
		{
			name: "Multiple parameters",
			args: []string{"id", "name"},
			want: "id = $1, name = $2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParamsToQuery(", ", tt.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}
