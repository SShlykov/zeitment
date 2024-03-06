package entity

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
	"time"
)

// Paragraph структура для хранения информации о параграфе
type Paragraph struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Title     string
	Text      string
	Type      string
	IsPublic  bool
	PageID    string
}

// TableName возвращает имя таблицы для структуры Paragraph
func (p Paragraph) TableName() string {
	return "paragraphs"
}

// AllFields возвращает список всех полей структуры Paragraph
func (p Paragraph) AllFields() []string {
	return []string{"id", "created_at", "updated_at", "deleted_at", "title", "text", "type", "is_public", "page_id"}
}

// InsertFields возвращает список полей, используемых при вставке новой записи
func (p Paragraph) InsertFields() []string {
	return []string{"title", "text", "type", "is_public", "page_id"}
}

// EntityToInsertValues преобразует сущность Paragraph в список значений для вставки
func (p Paragraph) EntityToInsertValues(entity any) []interface{} {
	if e, ok := entity.(Paragraph); ok {
		return []interface{}{
			e.Title, e.Text, e.Type, e.IsPublic, e.PageID,
		}
	}
	return nil
}

// ReadItem читает одну запись из строки запроса
func (p Paragraph) ReadItem(row pgx.Row) (any, error) {
	var paragraph Paragraph
	err := row.Scan(&paragraph.ID, &paragraph.CreatedAt, &paragraph.UpdatedAt, &paragraph.DeletedAt, &paragraph.Title,
		&paragraph.Text, &paragraph.Type, &paragraph.IsPublic, &paragraph.PageID)
	if err != nil {
		return nil, err
	}
	return paragraph, nil
}

// ReadList читает список записей из результатов запроса
func (p Paragraph) ReadList(rows pgx.Rows) ([]any, error) {
	var paragraphs []any
	for rows.Next() {
		paragraph, err := p.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		paragraphs = append(paragraphs, paragraph)
	}
	return paragraphs, nil
}
