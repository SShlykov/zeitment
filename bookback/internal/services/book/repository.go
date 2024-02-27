package book

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/services"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"strings"
)

const (
	// model fields and table name for books table
	tableName         = "books"
	columnID          = "id"
	columnCreatedAt   = "created_at"
	columnUpdatedAt   = "updated_at"
	columnDeletedAt   = "deleted_at"
	columnOwner       = "owner"
	columnTitle       = "title"
	columnAuthor      = "author"
	columnDescription = "description"
	columnIsPublic    = "is_public"
	columnPublication = "publication"
	columnImageLink   = "image_link"
	columnMapLink     = "map_link"
	columnMapParamsID = "map_params_id"
	columnVariables   = "variables"
	Returning         = " RETURNING "
	Where             = " WHERE "
)

// Repository определяет интерфейс для взаимодействия с хранилищем книг.
type Repository interface {
	Create(ctx context.Context, book *models.Book) (string, error)
	FindByID(ctx context.Context, id string) (*models.Book, error)
	Update(ctx context.Context, id string, book *models.Book) (*models.Book, error)
	Delete(ctx context.Context, id string) (*models.Book, error)
	List(ctx context.Context) ([]models.Book, error)
}

type repository struct {
	db db.Client
}

func allItems() string {
	cols := []string{columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt, columnOwner, columnTitle,
		columnAuthor, columnDescription, columnIsPublic, columnPublication, columnImageLink, columnMapLink,
		columnMapParamsID, columnVariables}

	return strings.Join(cols, ", ")
}

func insertItems() string {
	cols := []string{columnTitle, columnAuthor, columnOwner, columnDescription, columnIsPublic,
		columnPublication, columnImageLink, columnMapLink, columnMapParamsID, columnVariables}

	return strings.Join(cols, ", ")
}

// NewRepository создает новый экземпляр репозитория для книг.
func NewRepository(database db.Client) Repository {
	return &repository{database}
}

func (r *repository) Create(ctx context.Context, book *models.Book) (string, error) {
	query := "INSERT INTO" + " " + tableName + ` (` + insertItems() +
		`) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ` + Returning + columnID

	if book.Variables == nil {
		book.Variables = []string{}
	}

	args := []interface{}{book.Title, book.Author, book.Owner, book.Description, book.IsPublic, //nolint:gofmt
		book.Publication, book.ImageLink, book.MapLink, book.MapParamsID, book.Variables}
	q := db.Query{Name: "BookRepository.Insert", Raw: query}

	var id string
	if err := r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (*models.Book, error) {
	query := `SELECT ` + allItems() + ` FROM ` + tableName + Where + columnID + ` = $1 AND ` +
		columnDeletedAt + ` IS NULL LIMIT 1`

	q := db.Query{Name: "BookRepository.FindById", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, id)

	return readItem(row)
}

func (r *repository) Update(ctx context.Context, id string, updBook *models.Book) (*models.Book, error) {
	query := "Update " + tableName + " SET " +
		services.ParamsToQuery(", ", columnTitle, columnAuthor, columnOwner, columnDescription, columnIsPublic,
			columnPublication, columnImageLink, columnMapLink, columnMapParamsID, columnVariables) +
		" WHERE " + columnID + " = $11" + Returning + allItems()

	args := []interface{}{updBook.Title, updBook.Author, updBook.Owner, updBook.Description, //nolint:gofmt
		updBook.IsPublic, updBook.Publication, updBook.ImageLink, updBook.MapLink, updBook.MapParamsID,
		updBook.Variables, id}

	q := db.Query{Name: "BookRepository.Update", Raw: query}
	row := r.db.DB().QueryRowContext(ctx, q, args...)

	return readItem(row)
}

func (r *repository) Delete(ctx context.Context, id string) (*models.Book, error) {
	query := `DELETE FROM` + " " + tableName + Where + columnID + ` = $1 RETURNING ` + allItems()

	q := db.Query{Name: "BookRepository.Delete", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, id)

	return readItem(row)
}

func (r *repository) List(ctx context.Context) ([]models.Book, error) {
	query := `SELECT ` + allItems() + ` FROM ` + tableName + Where + columnDeletedAt + ` IS NULL`

	q := db.Query{Name: "BookRepository.List", Raw: query}

	rows, err := r.db.DB().QueryRawContextMulti(ctx, q)
	if err != nil {
		return nil, errors.New("params error")
	}

	return readList(rows)
}
