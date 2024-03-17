package models

import (
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type Chapter struct {
	ID          string                `json:"id"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	DeletedAt   types.Null[time.Time] `json:"deleted_at"`
	Title       string                `json:"title"`
	Number      int                   `json:"number"` // Порядковый номер главы (не уникальный т.к. главы могут быть скрыты)
	Text        string                `json:"text"`   // Превью текста главы
	BookID      string                `json:"book_id"`
	IsPublic    bool                  `json:"is_public"`
	MapLink     types.Null[string]    `json:"map_link"`      // карта
	MapParamsID types.Null[string]    `json:"map_params_id"` // параметры карты (координаты и тп.)
}

type CreateChapterRequest struct {
	Chapter *Chapter `json:"chapter"`
}

type UpdateChapterRequest struct {
	Chapter *Chapter `json:"chapter"`
}

type ToggleChapterRequest struct {
	ChapterID string `json:"chapter_id"`
}

type RequestChapter struct {
	Options PageOptions `json:"options,omitempty"`
	Chapter *Chapter    `json:"chapter,omitempty"`
}
