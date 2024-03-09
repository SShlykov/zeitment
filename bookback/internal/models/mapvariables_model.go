package models

import (
	"database/sql"
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type MapVariable struct {
	ID          string             `json:"id"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	BookID      string             `json:"book_id"`
	ChapterID   types.Null[string] `json:"chapter_id"`
	PageID      types.Null[string] `json:"page_id"`
	ParagraphID types.Null[string] `json:"paragraph_id"`
	MapLink     string             `json:"map_link"`
	Lat         float64            `json:"lat"`
	Lng         float64            `json:"lng"`
	Zoom        sql.NullInt64      `json:"zoom"`
	Date        types.Null[string] `json:"date"`
	Description types.Null[string] `json:"description"`
	Link        types.Null[string] `json:"link"`
	LinkText    types.Null[string] `json:"link_text"`
	LinkType    types.Null[string] `json:"link_type"`
	LinkImage   types.Null[string] `json:"link_image"`
	Image       types.Null[string] `json:"image"`
}

type CreateMapVariableRequest struct {
	MapVariable *MapVariable `json:"map_variable"`
}

type UpdateMapVariableRequest struct {
	MapVariable *MapVariable `json:"map_variable"`
}

type RequestMapVariable struct {
	Options     PageOptions  `json:"options,omitempty"`
	MapVariable *MapVariable `json:"map_variable,omitempty"`
}
