package adapters

import (
	"database/sql"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

func BooksEntityToModel(books []*entity.Book) []*models.Book {
	var result []*models.Book
	for _, book := range books {
		result = append(result, BookEntityToModel(book))
	}
	return result
}

func BookModelToEntity(book *models.Book) *entity.Book {
	return &entity.Book{
		ID:          book.ID,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
		DeletedAt:   sql.NullTime{Valid: book.DeletedAt.Valid, Time: book.DeletedAt.Value},
		Owner:       book.Owner,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		IsPublic:    book.IsPublic,
		Publication: sql.NullTime{Valid: book.Publication.Valid, Time: book.Publication.Value},
		ImageLink:   sql.NullString{Valid: book.ImageLink.Valid, String: book.ImageLink.Value},
		MapLink:     sql.NullString{Valid: book.MapLink.Valid, String: book.MapLink.Value},
		Variables:   book.Variables,
	}
}

func BookEntityToModel(book *entity.Book) *models.Book {
	return &models.Book{
		ID:          book.ID,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
		DeletedAt:   types.Null[time.Time]{Valid: book.DeletedAt.Valid, Value: book.DeletedAt.Time},
		Owner:       book.Owner,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		IsPublic:    book.IsPublic,
		Publication: types.Null[time.Time]{Valid: book.Publication.Valid, Value: book.Publication.Time},
		ImageLink:   types.Null[string]{Valid: book.ImageLink.Valid, Value: book.ImageLink.String},
		MapLink:     types.Null[string]{Valid: book.MapLink.Valid, Value: book.MapLink.String},
		Variables:   book.Variables,
	}
}
