package models

import (
	"time"
)

type Chapter struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt NullTime  `json:"deleted_at"`
	Title     string    `json:"title"`
	Number    int       `json:"number"`
	Text      string    `json:"text"`
	BookID    string    `json:"book_id"`
	IsPublic  bool      `json:"is_public"`
	Pages     []Page    `json:"pages"`
}
