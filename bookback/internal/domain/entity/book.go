package entity

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
	"time"
)

type Book struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	Owner       string
	Title       string
	Author      string
	Description string
	IsPublic    bool
	Publication sql.NullTime
	ImageLink   sql.NullString
	MapLink     sql.NullString
	MapParamsID sql.NullString
	Variables   []string
}

func (b Book) TableName() string {
	return "books"
}

func (b Book) AllFields() []string {
	return []string{"id", "created_at", "updated_at", "deleted_at", "owner", "title", "author",
		"description", "is_public", "publication", "image_link", "map_link", "map_params_id", "variables"}
}

func (b Book) InsertFields() []string {
	return []string{"owner", "title", "author", "description", "is_public", "publication", "image_link", "map_link",
		"map_params_id", "variables"}
}

func (b Book) EntityToInsertValues(impl *Book) []interface{} {
	return []interface{}{
		impl.Owner, impl.Title, impl.Author, impl.Description, impl.IsPublic, impl.Publication,
		impl.ImageLink, impl.MapLink, impl.MapParamsID, impl.Variables,
	}
}

func (b Book) ReadItem(row pgx.Row) (Book, error) {
	var book Book
	err := row.Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Owner, &book.Title, &book.Author,
		&book.Description, &book.IsPublic, &book.Publication, &book.ImageLink, &book.MapLink, &book.MapParamsID, &book.Variables)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

func (b Book) ReadList(rows pgx.Rows) ([]Book, error) {
	var books []Book
	for rows.Next() {
		book, err := b.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
