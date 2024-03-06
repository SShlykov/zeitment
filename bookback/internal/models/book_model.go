package models

import (
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type Book struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   types.Null[time.Time]
	Owner       string
	Title       string
	Author      string
	Description string
	IsPublic    bool
	Publication types.Null[time.Time]
	ImageLink   types.Null[string]
	MapLink     types.Null[string]
	Variables   []string
}

type CreateBookRequest struct {
	Book *Book `json:"book"`
}
type UpdateBookRequest struct {
	Book *Book `json:"book"`
}

type RequestBook struct {
	Options PageOptions `json:"options,omitempty"`
	Book    *Book       `json:"book,omitempty"`
}
