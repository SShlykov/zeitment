package models

import (
	"time"
)

type Page struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text"`
	ChapterID string    `json:"chapter_id"`
	IsPublic  bool      `json:"is_public"`
	//Paragraphs []Paragraph
	//DeletedAt sql.NullTime `json:"deleted_at"`
}
