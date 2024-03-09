package models

import (
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type Paragraph struct {
	ID        string                `json:"id"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	DeletedAt types.Null[time.Time] `json:"deleted_at"`
	Title     string                `json:"title"`
	Text      string                `json:"text"`
	Type      string                `json:"type"`
	IsPublic  bool                  `json:"is_public"`
	PageID    string                `json:"page_id"`
}

type CreateParagraphRequest struct {
	Paragraph *Paragraph `json:"paragraph"`
}

type UpdateParagraphRequest struct {
	Paragraph *Paragraph `json:"paragraph"`
}

type RequestParagraph struct {
	Options   PageOptions `json:"options,omitempty"`
	Paragraph *Paragraph  `json:"paragraph,omitempty"`
}
