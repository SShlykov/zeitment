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

func (c Chapter) InsertOrUpdateFields() []string {
	return []string{"title", "number", "text", "book_id", "is_public", "map_link", "map_params_id"}
}

func (c Chapter) EntityToInsertValues(entity *Chapter) []interface{} {
	return []interface{}{entity.Title, entity.Number, entity.Text, entity.BookID, entity.IsPublic,
		entity.MapLink, entity.MapParamsID}
}

func (c Chapter) ReadItem(row pgx.Row) (Chapter, error) {
	var e Chapter
	err := row.Scan(&e.ID, &e.CreatedAt, &e.UpdatedAt, &e.DeletedAt, &e.Title, &e.Number, &e.Text, &e.BookID,
		&e.IsPublic, &e.MapLink, &e.MapParamsID)
	if err != nil {
		return Chapter{}, err
	}
	return e, nil
}

func (c Chapter) ReadList(rows pgx.Rows) ([]Chapter, error) {
	var entities []Chapter
	for rows.Next() {
		chapter, err := c.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		entities = append(entities, chapter)
	}
	return entities, nil
}
