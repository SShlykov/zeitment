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
	Number      int
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
	return []string{"id", "created_at", "updated_at", "deleted_at", "title", "number", "text", "chapter_id", "is_public", "map_params_id"}
}

// InsertOrUpdateFields возвращает список полей, используемых при вставке новой записи
func (p Page) InsertOrUpdateFields() []string {
	return []string{"title", "number", "text", "chapter_id", "is_public", "map_params_id"}
}

// EntityToInsertValues преобразует сущность Page в список значений для вставки
func (p Page) EntityToInsertValues(entity *Page) []interface{} {
	return []interface{}{entity.Title, entity.Number, entity.Text, entity.ChapterID, entity.IsPublic, entity.MapParamsID}
}

// ReadItem читает одну запись из строки запроса
func (p Page) ReadItem(row pgx.Row) (Page, error) {
	var page Page
	err := row.Scan(&page.ID, &page.CreatedAt, &page.UpdatedAt, &page.DeletedAt, &page.Title, &page.Number, &page.Text,
		&page.ChapterID, &page.IsPublic, &page.MapParamsID)
	if err != nil {
		return Page{}, err
	}
	return page, nil
}

// ReadList читает список записей из результатов запроса
func (p Page) ReadList(rows pgx.Rows) ([]Page, error) {
	var pages []Page
	for rows.Next() {
		page, err := p.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}
	return pages, nil
}
