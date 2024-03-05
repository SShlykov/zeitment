package entity

import (
	"time"
)

type Chapter struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   NullTime   `json:"deleted_at"`
	Title       string     `json:"title"`
	Number      int        `json:"number"` // Порядковый номер главы (не уникальный т.к. главы могут быть скрыты)
	Text        string     `json:"text"`   // Превью текста главы
	BookID      string     `json:"book_id"`
	IsPublic    bool       `json:"is_public"`
	MapLink     NullString `json:"map_link"`      // карта
	MapParamsID NullString `json:"map_params_id"` // параметры карты (координаты и тп.)
	//Variables   []string   `json:"variables"`     // переменные мира key: key, value: value
	//Pages      []Page    `json:"pages"`
}
