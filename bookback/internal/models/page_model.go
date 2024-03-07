package models

import (
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type Page struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   types.Null[time.Time]
	Title       string
	Text        string
	ChapterID   string
	IsPublic    bool
	MapParamsID types.Null[string]
}

type CreatePageRequest struct {
	Page *Page `json:"page"`
}

type UpdatePageRequest struct {
	Page *Page `json:"page"`
}

type RequestPage struct {
	Options PageOptions `json:"options,omitempty"`
	Page    *Page       `json:"page,omitempty"`
}
