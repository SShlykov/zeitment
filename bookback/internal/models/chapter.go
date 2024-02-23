package models

import (
	"time"
)

// DeletedAt sql.NullTime `json:"deleted_at"`

type Chapter struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Number    int       `json:"number"`
	Text      string    `json:"text"`
	BookID    string    `json:"book_id"`
	IsPublic  bool      `json:"is_public"`
	Pages     []Page    `json:"pages"`
}
