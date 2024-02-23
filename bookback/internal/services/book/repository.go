package book

import (
	"context"
	"errors"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/SShlykov/zeitment/bookback/pkg/db/sq"
	"github.com/google/uuid"
	"time"
)

const (
	// model fields and table name for books table
	tableName         = "books"
	columnID          = "id"
	columnCreatedAt   = "created_at"
	columnUpdatedAt   = "updated_at"
	columnDeletedAt   = "deleted_at"
	columnTitle       = "title"
	columnAuthor      = "author"
	columnOwner       = "owner"
	columnDescription = "description"
	columnIsPublic    = "is_public"
	columnPublication = "publication"
	Returning         = "RETURNING "
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

// NewRepository создает новый экземпляр репозитория для книг.
func NewRepository(database db.Client) Repository {
	return &repository{database}
}

func (r *repository) Create(ctx context.Context, book *models.Book) (string, error) {
	query, args, err :=
		sq.SQ.Insert(tableName).
			Columns(columnID, columnTitle, columnAuthor, columnOwner,
				columnDescription, columnIsPublic, columnPublication).
			Values(uuid.New().String(), book.Title, book.Author, book.Owner,
				book.Description, book.IsPublic, book.Publication).
			Suffix("RETURNING " + columnID).
			ToSql()

	if err != nil {
		return "", err
	}

	q := db.Query{Name: "BookRepository.Insert", Raw: query}

	var id string
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		fmt.Println(err)
		return "", err
	}

	return id, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (*models.Book, error) {
	query, args, err := sq.
		Select(columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt, columnTitle, columnAuthor, columnOwner,
			columnDescription, columnIsPublic, columnPublication).
		From(tableName).
		Where(sq.And{sq.Eq{columnID: id}, sq.Eq{columnDeletedAt: nil}}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, errors.New("unknown error")
	}

	q := db.Query{Name: "BookRepository.FindById", Raw: query}

	var book models.Book
	if err = r.db.DB().QueryRowContext(ctx, q, args...).
		Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Title,
			&book.Author, &book.Owner, &book.Description, &book.IsPublic,
			&book.Publication); err != nil {
		fmt.Println(err)
		return nil, errors.New("params error")
	}

	return &book, nil
}

func (r *repository) Update(ctx context.Context, id string, updBook *models.Book) (*models.Book, error) {
	query, args, err := sq.Update(tableName).
		Where(sq.And{sq.Eq{columnID: id}, sq.Eq{columnDeletedAt: nil}}).
		Set(columnUpdatedAt, time.Now()).
		Set(columnTitle, updBook.Title).
		Set(columnAuthor, updBook.Author).
		Set(columnOwner, updBook.Owner).
		Set(columnDescription, updBook.Description).
		Set(columnIsPublic, updBook.IsPublic).
		Set(columnPublication, updBook.Publication).
		Suffix(Returning + fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s", columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt,
			columnTitle, columnAuthor, columnOwner, columnDescription, columnIsPublic, columnPublication)).
		ToSql()

	if err != nil {
		return nil, errors.New("unknown error")
	}

	q := db.Query{Name: "BookRepository.Update", Raw: query}

	var book models.Book
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt,
		&book.DeletedAt, &book.Title, &book.Author, &book.Owner, &book.Description, &book.IsPublic,
		&book.Publication); err != nil {
		return nil, errors.New("params error")
	}

	return &book, nil
}

func (r *repository) Delete(ctx context.Context, id string) (*models.Book, error) {
	query, args, err := sq.SQ.Delete(tableName).
		Where(sq.Eq{columnID: id}).
		Suffix(Returning + fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s", columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt,
			columnTitle, columnAuthor, columnOwner, columnDescription, columnIsPublic, columnPublication)).
		ToSql()

	if err != nil {
		return nil, errors.New("unknown error")
	}

	q := db.Query{Name: "BookRepository.Delete", Raw: query}
	fmt.Println(q)

	var book models.Book
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt,
		&book.DeletedAt, &book.Title, &book.Author, &book.Owner, &book.Description, &book.IsPublic,
		&book.Publication); err != nil {
		fmt.Println(err)
		return nil, errors.New("params error")
	}

	return &book, nil
}

func (r *repository) List(ctx context.Context) ([]models.Book, error) {
	query, args, err := sq.
		Select(columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt, columnTitle, columnAuthor,
			columnOwner, columnDescription, columnIsPublic, columnPublication).
		From(tableName).
		Where(sq.Eq{columnDeletedAt: nil}).
		ToSql()

	if err != nil {
		return nil, errors.New("unknown error")
	}

	q := db.Query{Name: "BookRepository.Delete", Raw: query}

	var booksList []models.Book
	rows, err := r.db.DB().QueryRawContextMulti(ctx, q, args...)
	if err != nil {
		return nil, errors.New("params error")
	}

	for rows.Next() {
		var book models.Book
		if err = rows.Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Title,
			&book.Author, &book.Owner, &book.Description, &book.IsPublic,
			&book.Publication); err != nil {
			return nil, errors.New("params error")
		}
		booksList = append(booksList, book)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("params error")
	}

	return booksList, nil
}
