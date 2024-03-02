package bookevents

import (
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/jackc/pgx/v5"
)

func readList(rows pgx.Rows) ([]models.BookEvent, error) {
	defer rows.Close()
	events := make([]models.BookEvent, 0)
	for rows.Next() {
		item, err := readItem(rows)
		if err != nil {
			return nil, err
		}
		events = append(events, *item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, rows.Err()
}

func readItem(row pgx.Row) (*models.BookEvent, error) {
	var event models.BookEvent
	if err := row.Scan(&event.ID, &event.CreatedAt, &event.UpdatedAt, &event.BookID, &event.ChapterID,
		&event.PageID, &event.ParagraphID, &event.EventType, &event.IsPublic, &event.Key, &event.Value,
		&event.Link, &event.LinkText, &event.LinkType, &event.LinkImage, &event.Description); err != nil {
		return nil, err
	}

	return &event, nil
}
