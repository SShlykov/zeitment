package chapterrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/jackc/pgx/v5"
)

func readList(rows pgx.Rows) ([]models.Chapter, error) {
	defer rows.Close()
	chapters := make([]models.Chapter, 0)
	for rows.Next() {
		item, err := readItem(rows)
		if err != nil {
			return nil, err
		}
		chapters = append(chapters, *item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chapters, rows.Err()
}

func readItem(row pgx.Row) (*models.Chapter, error) {
	var chapter models.Chapter
	if err := row.Scan(&chapter.ID, &chapter.CreatedAt, &chapter.UpdatedAt, &chapter.DeletedAt, &chapter.Title,
		&chapter.Number, &chapter.Text, &chapter.BookID, &chapter.IsPublic,
		&chapter.MapLink, &chapter.MapParamsID); err != nil {
		return nil, err
	}

	return &chapter, nil
}
