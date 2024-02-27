package book

import (
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/jackc/pgx/v5"
)

func readList(rows pgx.Rows) ([]models.Book, error) {
	books := make([]models.Book, 0)
	for rows.Next() {
		item, err := readItem(rows)
		if err != nil {
			return nil, err
		}
		books = append(books, *item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func readItem(row pgx.Row) (*models.Book, error) {
	var book models.Book
	if err := row.Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Owner,
		&book.Title, &book.Author, &book.Description, &book.IsPublic, &book.Publication,
		&book.ImageLink, &book.MapLink, &book.MapParamsID, &book.Variables); err != nil {
		return nil, err
	}

	return &book, nil
}
