package chapter

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
	tableName       = "chapters"
	columnID        = "id"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
	columnDeletedAt = "deleted_at"
	columnTitle     = "title"
	columnNumber    = "number"
	columnText      = "text"
	columnBookID    = "book_id"
	columnIsPublic  = "is_public"
	Returning       = "RETURNING "
)

// Repository определяет интерфейс для взаимодействия с хранилищем книг.
type Repository interface {
	Create(ctx context.Context, book *models.Chapter) (string, error)
	FindByID(ctx context.Context, id string) (*models.Chapter, error)
	Update(ctx context.Context, id string, book *models.Chapter) (*models.Chapter, error)
	Delete(ctx context.Context, id string) (*models.Chapter, error)
	List(ctx context.Context) ([]models.Chapter, error)
}

type repository struct {
	db db.Client
}

// NewRepository создает новый экземпляр репозитория для книг.
func NewRepository(database db.Client) Repository {
	return &repository{database}
}

func (r *repository) Create(ctx context.Context, chapter *models.Chapter) (string, error) {
	query, args, err :=
		sq.SQ.Insert(tableName).
			Columns(columnID, columnTitle, columnNumber, columnText, columnBookID, columnIsPublic).
			Values(uuid.New().String(), chapter.Title, chapter.Number, chapter.Text, chapter.BookID, chapter.IsPublic).
			Suffix("RETURNING " + columnID).
			ToSql()

	if err != nil {
		return "", err
	}

	q := db.Query{Name: "ChapterRepository.Insert", Raw: query}

	var id string
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (*models.Chapter, error) {
	query, args, err := sq.
		Select(columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt, columnTitle, columnNumber,
			columnText, columnBookID, columnIsPublic).
		From(tableName).
		Where(sq.And{sq.Eq{columnID: id}, sq.Eq{columnDeletedAt: nil}}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, errors.New("unknown error")
	}

	q := db.Query{Name: "ChapterRepository.FindById", Raw: query}

	var chapter models.Chapter
	if err = r.db.DB().QueryRowContext(ctx, q, args...).
		Scan(&chapter.ID, &chapter.CreatedAt, &chapter.UpdatedAt, &chapter.DeletedAt, &chapter.Title, &chapter.Number,
			&chapter.Text, &chapter.BookID, &chapter.IsPublic); err != nil {
		fmt.Println(err)
		return nil, errors.New("params error")
	}

	return &chapter, nil
}

func (r *repository) Update(ctx context.Context, id string, updChapter *models.Chapter) (*models.Chapter, error) {
	query, args, err := sq.Update(tableName).
		Where(sq.And{sq.Eq{columnID: id}, sq.Eq{columnDeletedAt: nil}}).
		Set(columnUpdatedAt, time.Now()).
		Set(columnTitle, updChapter.Title).
		Set(columnNumber, updChapter.IsPublic).
		Set(columnText, updChapter.Title).
		Set(columnBookID, updChapter.IsPublic).
		Set(columnIsPublic, updChapter.IsPublic).
		Suffix(Returning +
			fmt.Sprintf("%s, %s, %s, %s, %s, %s", columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt,
				columnTitle, columnIsPublic)).
		ToSql()

	if err != nil {
		return nil, errors.New("unknown error")
	}

	q := db.Query{Name: "ChapterRepository.Update", Raw: query}

	var chapter models.Chapter
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chapter.ID, &chapter.CreatedAt, &chapter.UpdatedAt, &chapter.DeletedAt,
		&chapter.Title, &chapter.Number, &chapter.Text, &chapter.BookID, &chapter.IsPublic); err != nil {
		return nil, errors.New("params error")
	}

	return &chapter, nil
}

func (r *repository) Delete(ctx context.Context, id string) (*models.Chapter, error) {
	query, args, err := sq.SQ.Delete(tableName).
		Where(sq.Eq{columnID: id}).
		Suffix(Returning +
			fmt.Sprintf("%s, %s, %s, %s, %s", columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt,
				columnTitle)).
		ToSql()

	if err != nil {
		return nil, errors.New("unknown error")
	}

	q := db.Query{Name: "ChapterRepository.Delete", Raw: query}
	fmt.Println(q)

	var chapter models.Chapter
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chapter.ID, &chapter.CreatedAt, &chapter.UpdatedAt, &chapter.DeletedAt,
		&chapter.Title, &chapter.Number, &chapter.Text, &chapter.BookID, &chapter.IsPublic); err != nil {
		fmt.Println(err)
		return nil, errors.New("params error")
	}

	return &chapter, nil
}

func (r *repository) List(ctx context.Context) ([]models.Chapter, error) {
	query, args, err := sq.
		Select(columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt, columnTitle, columnIsPublic).
		From(tableName).
		Where(sq.Eq{columnDeletedAt: nil}).
		ToSql()

	if err != nil {
		return nil, errors.New("unknown error")
	}

	q := db.Query{Name: "ChapterRepository.Delete", Raw: query}

	var chapterList []models.Chapter
	rows, err := r.db.DB().QueryRawContextMulti(ctx, q, args...)
	if err != nil {
		return nil, errors.New("params error")
	}

	for rows.Next() {
		var chapter models.Chapter
		if err = rows.Scan(&chapter.ID, &chapter.CreatedAt, &chapter.UpdatedAt, &chapter.DeletedAt, &chapter.Title, &chapter.Number,
			&chapter.Text, &chapter.BookID, &chapter.IsPublic); err != nil {
			return nil, errors.New("params error")
		}
		chapterList = append(chapterList, chapter)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("params error")
	}

	return chapterList, nil
}
