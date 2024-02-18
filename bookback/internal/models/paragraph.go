package models

import (
	"time"
)

type Paragraph struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt NullTime  `json:"deleted_at"`

	Text     string `json:"text"`
	IsPublic bool   `json:"is_public"`
	PageID   string `json:"page_id"`
}
