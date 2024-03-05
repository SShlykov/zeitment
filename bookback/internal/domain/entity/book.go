package entity

import (
	"time"
)

type Book struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   NullTime   `json:"deleted_at"`
	Owner       string     `json:"owner"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Description string     `json:"description"`
	IsPublic    bool       `json:"is_public"`
	Publication NullTime   `json:"publication"`
	ImageLink   NullString `json:"image_link"`    // обложка
	MapLink     NullString `json:"map_link"`      // карта
	MapParamsID NullString `json:"map_params_id"` // Параметры карты (координаты и тп.)
	// Переменные мира, список ключей ? Ограничение на длину ключа
	// и значения + ограничение на количество переменных
	Variables []string `json:"variables"`
	//Chapters    []Chapter
}
