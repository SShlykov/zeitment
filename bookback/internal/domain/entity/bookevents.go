package entity

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
	"time"
)

type BookEvent struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BookID      string
	ChapterID   sql.NullString
	PageID      sql.NullString
	ParagraphID sql.NullString
	EventType   sql.NullString // Тип события (начало главы, начало страницы, начало параграфа...)
	IsPublic    bool
	Key         string
	Value       string
	Link        sql.NullString
	LinkText    sql.NullString
	LinkType    sql.NullString
	LinkImage   sql.NullString
	Description sql.NullString
}

func (be BookEvent) TableName() string {
	return "book_events"
}

func (be BookEvent) AllFields() []string {
	return []string{"id", "created_at", "updated_at", "book_id", "chapter_id", "page_id", "paragraph_id", "event_type",
		"is_public", "key", "value", "link", "link_text", "link_type", "link_image", "description"}
}

func (be BookEvent) InsertOrUpdateFields() []string {
	return []string{"book_id", "chapter_id", "page_id", "paragraph_id", "event_type",
		"is_public", "key", "value", "link", "link_text", "link_type", "link_image", "description"}
}

func (be BookEvent) EntityToInsertValues(entity *BookEvent) []interface{} {
	return []interface{}{entity.BookID, entity.ChapterID, entity.PageID, entity.ParagraphID, entity.EventType,
		entity.IsPublic, entity.Key, entity.Value, entity.Link, entity.LinkText, entity.LinkType, entity.LinkImage, entity.Description}
}

func (be BookEvent) ReadItem(row pgx.Row) (BookEvent, error) {
	var e BookEvent
	err := row.Scan(&e.ID, &e.CreatedAt, &e.UpdatedAt, &e.BookID, &e.ChapterID, &e.PageID, &e.ParagraphID, &e.EventType,
		&e.IsPublic, &e.Key, &e.Value, &e.Link, &e.LinkText, &e.LinkType, &e.LinkImage, &e.Description)
	if err != nil {
		return BookEvent{}, err
	}
	return e, nil
}

func (be BookEvent) ReadList(rows pgx.Rows) ([]BookEvent, error) {
	var entities []BookEvent
	for rows.Next() {
		bookEvent, err := be.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		entities = append(entities, bookEvent)
	}
	return entities, nil
}
