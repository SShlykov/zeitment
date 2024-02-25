package models

import (
	"database/sql"
	"time"
)

type BookEvent struct {
	ID         string    `json:"id"`
	InsertedAt time.Time `json:"inserted_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	BookID      string         `json:"book_id"`
	ChapterID   sql.NullString `json:"chapter_id"`
	PageID      sql.NullString `json:"page_id"`
	ParagraphID sql.NullString `json:"paragraph_id"`
	EventType   string         `json:"event_type"` // Тип события (начало главы, начало страницы, начало параграфа...)
	IsPublic    bool           `json:"is_public"`
	Key         string         `json:"key"`
	Value       string         `json:"value"`
	Link        string         `json:"link"`
	LinkText    string         `json:"link_text"`
	LinkType    string         `json:"link_type"`
	LinkImage   string         `json:"link_image"`
	Description string         `json:"description"`
}
