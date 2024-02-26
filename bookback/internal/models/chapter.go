package models

import (
	"time"
)

type Chapter struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   NullTime  `json:"deleted_at"`
	Title       string    `json:"title"`
	Number      int       `json:"number"` // Порядковый номер главы (не уникальный т.к. главы могут быть скрыты)
	Text        string    `json:"text"`   // Превью текста главы
	BookID      string    `json:"book_id"`
	IsPublic    bool      `json:"is_public"`
	MapLink     string    `json:"map_link"`      // карта
	MapParamsID string    `json:"map_params_id"` // параметры карты (координаты и тп.)

	//Variables  []Variable `json:"variables"` // переменные мира key: key, value: value
	//Pages      []Page    `json:"pages"`
}
