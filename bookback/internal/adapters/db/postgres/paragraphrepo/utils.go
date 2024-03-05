package paragraphrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/jackc/pgx/v5"
)

func readList(rows pgx.Rows) ([]models.Paragraph, error) {
	defer rows.Close()
	paragraphs := make([]models.Paragraph, 0)
	for rows.Next() {
		item, err := readItem(rows)
		if err != nil {
			return nil, err
		}
		paragraphs = append(paragraphs, *item)
	}

	return paragraphs, rows.Err()
}

func readItem(row pgx.Row) (*models.Paragraph, error) {
	var paragraph models.Paragraph
	if err := row.Scan(&paragraph.ID, &paragraph.CreatedAt, &paragraph.UpdatedAt, &paragraph.DeletedAt,
		&paragraph.Title, &paragraph.Text, &paragraph.Type, &paragraph.IsPublic, &paragraph.PageID); err != nil {
		return nil, err
	}

	return &paragraph, nil
}
