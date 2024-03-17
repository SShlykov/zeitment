package models

import (
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type Book struct {
	ID          string                `json:"id"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	DeletedAt   types.Null[time.Time] `json:"deleted_at"`
	Owner       string                `json:"owner"`
	Title       string                `json:"title"`
	Author      string                `json:"author"`
	Description string                `json:"description"`
	IsPublic    bool                  `json:"is_public"`
	Publication types.Null[time.Time] `json:"publication"`
	ImageLink   types.Null[string]    `json:"image_link"`
	MapLink     types.Null[string]    `json:"map_link"`
	Variables   []string              `json:"variables"`
}

type CreateBookRequest struct {
	Book *Book `json:"book"`
}
type UpdateBookRequest struct {
	Book *Book `json:"book"`
}

type RequestBook struct {
	Options PageOptions `json:"options"`
	Book    *Book       `json:"book,omitempty"`
}

type RequestTOC struct {
	BookID string `json:"book_id"`
}
