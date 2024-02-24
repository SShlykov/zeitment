package page

import (
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/jackc/pgx/v5"
)

func readList(rows pgx.Rows) ([]models.Page, error) {
	var pages []models.Page
	for rows.Next() {
		item, err := readItem(rows)
		if err != nil {
			return nil, err
		}
		pages = append(pages, *item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pages, nil
}

func readItem(row pgx.Row) (*models.Page, error) {
	var page models.Page
	if err := row.Scan(&page.ID, &page.CreatedAt, &page.UpdatedAt, &page.DeletedAt, &page.Text, &page.ChapterID,
		&page.IsPublic); err != nil {
		return nil, err
	}

	return &page, nil
}
