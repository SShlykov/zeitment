package mapvariables

import (
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/jackc/pgx/v5"
)

func readList(rows pgx.Rows) ([]models.MapVariable, error) {
	var variables []models.MapVariable
	for rows.Next() {
		item, err := readItem(rows)
		if err != nil {
			return nil, err
		}
		variables = append(variables, *item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return variables, nil
}

func readItem(row pgx.Row) (*models.MapVariable, error) {
	var variable models.MapVariable
	if err := row.Scan(&variable.ID, &variable.InsertedAt, &variable.BookID,
		&variable.ChapterID, &variable.PageID, &variable.ParagraphID, &variable.MapLink, &variable.Lat,
		&variable.Lng, &variable.Zoom, &variable.Date, &variable.Description, &variable.Link, &variable.LinkText,
		&variable.LinkType, &variable.LinkImage, &variable.Image); err != nil {
		return nil, err
	}

	return &variable, nil
}
