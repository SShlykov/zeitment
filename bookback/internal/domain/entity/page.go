package entity

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
	"time"
)

// Page структура для хранения информации о странице
type Page struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	Title       string
	Text        string
	ChapterID   string
	IsPublic    bool
	MapParamsID sql.NullString
}

// TableName возвращает имя таблицы для структуры Page
func (p Page) TableName() string {
	return "pages"
}

// AllFields возвращает список всех полей структуры Page
func (p Page) AllFields() []string {
	return []string{"id", "created_at", "updated_at", "deleted_at", "title", "text", "chapter_id", "is_public", "map_params_id"}
}

// InsertFields возвращает список полей, используемых при вставке новой записи
func (p Page) InsertFields() []string {
	return []string{"title", "text", "chapter_id", "is_public", "map_params_id"}
}

// EntityToInsertValues преобразует сущность Page в список значений для вставки
func (p Page) EntityToInsertValues(entity any) []interface{} {
	if e, ok := entity.(Page); ok {
		return []interface{}{
			e.Title, e.Text, e.ChapterID, e.IsPublic, e.MapParamsID,
		}
	}
	return nil
}

// ReadItem читает одну запись из строки запроса
func (p Page) ReadItem(row pgx.Row) (any, error) {
	var page Page
	err := row.Scan(&page.ID, &page.CreatedAt, &page.UpdatedAt, &page.DeletedAt, &page.Title, &page.Text,
		&page.ChapterID, &page.IsPublic, &page.MapParamsID)
	if err != nil {
		return nil, err
	}
	return page, nil
}

// ReadList читает список записей из результатов запроса
func (p Page) ReadList(rows pgx.Rows) ([]any, error) {
	var pages []any
	for rows.Next() {
		page, err := p.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}
	return pages, nil
}
