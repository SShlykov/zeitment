package models

import (
	"database/sql"
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

type MapVariable struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BookID      string
	ChapterID   types.Null[string]
	PageID      types.Null[string]
	ParagraphID types.Null[string]
	MapLink     string
	Lat         float64
	Lng         float64
	Zoom        sql.NullInt64
	Date        types.Null[string]
	Description types.Null[string]
	Link        types.Null[string]
	LinkText    types.Null[string]
	LinkType    types.Null[string]
	LinkImage   types.Null[string]
	Image       types.Null[string]
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
