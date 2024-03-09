package models

import (
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type Page struct {
	ID          string                `json:"id"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	DeletedAt   types.Null[time.Time] `json:"deleted_at"`
	Title       string                `json:"title"`
	Text        string                `json:"text"`
	ChapterID   string                `json:"chapter_id"`
	IsPublic    bool                  `json:"is_public"`
	MapParamsID types.Null[string]    `json:"map_params_id"`
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
