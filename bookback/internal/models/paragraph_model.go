package models

import (
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type Paragraph struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt types.Null[time.Time]
	Title     string
	Text      string
	Type      string
	IsPublic  bool
	PageID    string
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
