package models

import (
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type Chapter struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   types.Null[time.Time]
	Title       string
	Number      int    // Порядковый номер главы (не уникальный т.к. главы могут быть скрыты)
	Text        string // Превью текста главы
	BookID      string
	IsPublic    bool
	MapLink     types.Null[string] // карта
	MapParamsID types.Null[string] // параметры карты (координаты и тп.)
}

type CreateChapterRequest struct {
	Chapter *Chapter `json:"chapter"`
}

type UpdateChapterRequest struct {
	Chapter *Chapter `json:"chapter"`
}

type RequestChapter struct {
	Options PageOptions `json:"options,omitempty"`
	Chapter *Chapter    `json:"chapter,omitempty"`
}
