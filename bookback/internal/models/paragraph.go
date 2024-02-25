package models

import (
	"time"
)

type Paragraph struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text"`
	Type      string    `json:"type"`
	IsPublic  bool      `json:"is_public"`
	PageID    string    `json:"page_id"`
	DeletedAt NullTime  `json:"deleted_at"`
}
