package paragraph

import (
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/jackc/pgx/v5"
)

func readList(rows pgx.Rows) ([]models.Paragraph, error) {
	var paragraphs []models.Paragraph
	for rows.Next() {
		item, err := readItem(rows)
		if err != nil {
			return nil, err
		}
		paragraphs = append(paragraphs, *item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return paragraphs, nil
}

func readItem(row pgx.Row) (*models.Paragraph, error) {
	var paragraph models.Paragraph
	if err := row.Scan(&paragraph.ID, &paragraph.CreatedAt, &paragraph.UpdatedAt, &paragraph.DeletedAt, &paragraph.Text, &paragraph.PageID, &paragraph.IsPublic); err != nil {
		return nil, err
	}

	return &paragraph, nil
}
