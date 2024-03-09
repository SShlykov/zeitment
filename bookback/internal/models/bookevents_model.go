package models

import (
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type BookEvent struct {
	ID          string             `json:"id"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	BookID      string             `json:"book_id"`
	ChapterID   types.Null[string] `json:"chapter_id"`
	PageID      types.Null[string] `json:"page_id"`
	ParagraphID types.Null[string] `json:"paragraph_id"`
	EventType   types.Null[string] `json:"event_type"` // Тип события (начало главы, начало страницы, начало параграфа...)
	IsPublic    bool               `json:"is_public"`
	Key         string             `json:"key"`
	Value       string             `json:"value"`
	Link        types.Null[string] `json:"link"`
	LinkText    types.Null[string] `json:"link_text"`
	LinkType    types.Null[string] `json:"link_type"`
	LinkImage   types.Null[string] `json:"link_image"`
	Description types.Null[string] `json:"description"`
}

type CreateBookEventRequest struct {
	BookEvent *BookEvent `json:"book_event"`
}
type UpdateBookEventRequest struct {
	BookEvent *BookEvent `json:"book_event"`
}

type RequestBookEvent struct {
	Options   PageOptions `json:"options,omitempty"`
	BookEvent *BookEvent  `json:"book_event,omitempty"`
}
