package entity

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
	"time"
)

type Chapter struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	Title       string
	Number      int    // Порядковый номер главы (не уникальный т.к. главы могут быть скрыты)
	Text        string // Превью текста главы
	BookID      string
	IsPublic    bool
	MapLink     sql.NullString // карта
	MapParamsID sql.NullString // параметры карты (координаты и тп.)
}

func (c Chapter) TableName() string {
	return "chapters"
}

func (c Chapter) AllFields() []string {
	return []string{"id", "created_at", "updated_at", "deleted_at", "title", "number", "text", "book_id",
		"is_public", "map_link", "map_params_id"}
}

func (c Chapter) InsertFields() []string {
	return []string{"title", "number", "text", "book_id", "is_public", "map_link", "map_params_id"}
}

func (c Chapter) EntityToInsertValues(entity any) []interface{} {
	if e, ok := entity.(Chapter); ok {
		return []interface{}{e.Title, e.Number, e.Text, e.BookID, e.IsPublic, e.MapLink, e.MapParamsID}
	}
	return nil
}

func (c Chapter) ReadItem(row pgx.Row) (any, error) {
	var e Chapter
	err := row.Scan(&e.ID, &e.CreatedAt, &e.UpdatedAt, &e.Title, &e.Number, &e.Text, &e.BookID, &e.IsPublic,
		&e.MapLink, &e.MapParamsID)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (c Chapter) ReadList(rows pgx.Rows) ([]any, error) {
	var entities []any
	for rows.Next() {
		chapter, err := c.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		entities = append(entities, chapter)
	}
	return entities, nil
}
