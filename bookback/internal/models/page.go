package models

import (
	"time"
)

type Page struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   NullTime  `json:"deleted_at"`
	Text        string    `json:"text"`
	ChapterID   string    `json:"chapter_id"`
	IsPublic    bool      `json:"is_public"`
	MapParamsID string    `json:"map_params"` // параметры карты (координаты и тп.)
	//Paragraphs []Paragraph
	//Variables  []Variable `json:"variables"` // переменные мира key: key, value: value
}
