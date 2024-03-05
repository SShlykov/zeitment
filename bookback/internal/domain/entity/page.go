package entity

import (
	"time"
)

type Page struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   NullTime   `json:"deleted_at"`
	Title       string     `json:"title"`
	Text        string     `json:"text"`
	ChapterID   string     `json:"chapter_id"`
	IsPublic    bool       `json:"is_public"`
	MapParamsID NullString `json:"map_params"` // параметры карты (координаты и тп.)
	//Paragraphs []Paragraph
	//Variables  []Variable `json:"variables"` // переменные мира key: key, value: value
}
