package pagerepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/jackc/pgx/v5"
)

func readList(rows pgx.Rows) ([]models.Page, error) {
	defer rows.Close()
	pages := make([]models.Page, 0)
	for rows.Next() {
		item, err := readItem(rows)
		if err != nil {
			return nil, err
		}
		pages = append(pages, *item)
	}

	return pages, rows.Err()
}

func readItem(row pgx.Row) (*models.Page, error) {
	var page models.Page
	if err := row.Scan(&page.ID, &page.CreatedAt, &page.UpdatedAt, &page.DeletedAt, &page.Title, &page.Text, &page.ChapterID,
		&page.IsPublic, &page.MapParamsID); err != nil {
		return nil, err
	}

	return &page, nil
}
