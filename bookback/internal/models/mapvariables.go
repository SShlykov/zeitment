package models

import (
	"time"
)

type MapVariable struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`

	BookID      string     `json:"book_id"`
	ChapterID   NullString `json:"chapter_id"`
	PageID      NullString `json:"page_id"`
	ParagraphID NullString `json:"paragraph_id"`
	MapLink     string     `json:"map_link"`
	Lat         float64    `json:"lat"`
	Lng         float64    `json:"lng"`
	Zoom        int        `json:"zoom"`
	Date        time.Time  `json:"date"`
	Description NullString `json:"description"`
	Link        NullString `json:"link"`
	LinkText    NullString `json:"link_text"`
	LinkType    NullString `json:"link_type"`
	LinkImage   NullString `json:"link_image"`
	Image       NullString `json:"image"`
}
