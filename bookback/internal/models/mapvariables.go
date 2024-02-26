package models

import (
	"database/sql"
	"time"
)

type MapVariable struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`

	BookID      string         `json:"book_id"`
	ChapterID   sql.NullString `json:"chapter_id"`
	PageID      sql.NullString `json:"page_id"`
	ParagraphID sql.NullString `json:"paragraph_id"`
	MapLink     string         `json:"map_link"`
	Lat         float64        `json:"lat"`
	Lng         float64        `json:"lng"`
	Zoom        int            `json:"zoom"`
	Date        time.Time      `json:"date"`
	Description string         `json:"description"`
	Link        string         `json:"link"`
	LinkText    string         `json:"link_text"`
	LinkType    string         `json:"link_type"`
	LinkImage   string         `json:"link_image"`
	Image       string         `json:"image"`
}
