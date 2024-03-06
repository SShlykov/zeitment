package converter

import (
	"database/sql"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
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
		DeletedAt:   models.Null[time.Time]{Valid: book.DeletedAt.Valid, Value: book.DeletedAt.Time},
		Owner:       book.Owner,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		IsPublic:    book.IsPublic,
		Publication: models.Null[time.Time]{Valid: book.Publication.Valid, Value: book.Publication.Time},
		ImageLink:   models.Null[string]{Valid: book.ImageLink.Valid, Value: book.ImageLink.String},
		MapLink:     models.Null[string]{Valid: book.MapLink.Valid, Value: book.MapLink.String},
		Variables:   book.Variables,
	}
}
