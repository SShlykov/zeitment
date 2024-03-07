package models

import (
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type BookEvent struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BookID      string
	ChapterID   types.Null[string]
	PageID      types.Null[string]
	ParagraphID types.Null[string]
	EventType   types.Null[string] // Тип события (начало главы, начало страницы, начало параграфа...)
	IsPublic    bool
	Key         string
	Value       string
	Link        types.Null[string]
	LinkText    types.Null[string]
	LinkType    types.Null[string]
	LinkImage   types.Null[string]
	Description types.Null[string]
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
