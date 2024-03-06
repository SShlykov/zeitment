package entity

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
	"time"
)

// MapVariable структура для хранения информации о переменных карты
type MapVariable struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BookID      string
	ChapterID   sql.NullString
	PageID      sql.NullString
	ParagraphID sql.NullString
	MapLink     string
	Lat         float64
	Lng         float64
	Zoom        sql.NullInt64
	Date        sql.NullString
	Description sql.NullString
	Link        sql.NullString
	LinkText    sql.NullString
	LinkType    sql.NullString
	LinkImage   sql.NullString
	Image       sql.NullString
}

// TableName возвращает имя таблицы для структуры MapVariable
func (mv MapVariable) TableName() string {
	return "map_variables"
}

// AllFields возвращает список всех полей структуры MapVariable
func (mv MapVariable) AllFields() []string {
	return []string{"id", "created_at", "updated_at", "book_id", "chapter_id", "page_id", "paragraph_id", "map_link",
		"lat", "lng", "zoom", "date", "description", "link", "link_text", "link_type", "link_image", "image"}
}

// InsertFields возвращает список полей, используемых при вставке новой записи
func (mv MapVariable) InsertFields() []string {
	return []string{"book_id", "chapter_id", "page_id", "paragraph_id", "map_link", "lat", "lng", "zoom", "date",
		"description", "link", "link_text", "link_type", "link_image", "image"}
}

// EntityToInsertValues преобразует сущность в список значений для вставки
func (mv MapVariable) EntityToInsertValues(entity any) []interface{} {
	if e, ok := entity.(MapVariable); ok {
		return []interface{}{
			e.BookID, e.ChapterID, e.PageID, e.ParagraphID, e.MapLink, e.Lat, e.Lng, e.Zoom, e.Date,
			e.Description, e.Link, e.LinkText, e.LinkType, e.LinkImage, e.Image,
		}
	}
	return nil
}

// ReadItem читает одну запись из строки запроса
func (mv MapVariable) ReadItem(row pgx.Row) (any, error) {
	var mapVariable MapVariable
	err := row.Scan(&mapVariable.ID, &mapVariable.CreatedAt, &mapVariable.UpdatedAt, &mapVariable.BookID, &mapVariable.ChapterID,
		&mapVariable.PageID, &mapVariable.ParagraphID, &mapVariable.MapLink, &mapVariable.Lat, &mapVariable.Lng, &mapVariable.Zoom,
		&mapVariable.Date, &mapVariable.Description, &mapVariable.Link, &mapVariable.LinkText, &mapVariable.LinkType,
		&mapVariable.LinkImage, &mapVariable.Image)
	if err != nil {
		return nil, err
	}
	return mapVariable, nil
}

// ReadList читает список записей из результатов запроса
func (mv MapVariable) ReadList(rows pgx.Rows) ([]any, error) {
	var mapVariables []any
	for rows.Next() {
		mapVariable, err := mv.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		mapVariables = append(mapVariables, mapVariable)
	}
	return mapVariables, nil
}
