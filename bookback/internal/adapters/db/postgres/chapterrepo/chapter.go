package chapterrepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/SShlykov/zeitment/bookback/pkg/querybuilder"
	"strings"
)

const (
	tableName         = "chapters"
	columnID          = "id"
	columnCreatedAt   = "created_at"
	columnUpdatedAt   = "updated_at"
	columnDeletedAt   = "deleted_at"
	columnTitle       = "title"
	columnNumber      = "number"
	columnText        = "text"
	columnBookID      = "book_id"
	columnIsPublic    = "is_public"
	columnMapLink     = "map_link"
	columnMapParamsID = "map_params_id"
	Returning         = " RETURNING "
	WHERE             = " WHERE "
	FromTable         = " FROM " + tableName + " "
	NotDeleted        = " " + columnDeletedAt + " IS NULL "
)

// Repository определяет интерфейс для взаимодействия с хранилищем глав.
type Repository interface {
	Create(ctx context.Context, book *models.Chapter) (string, error)
	FindByID(ctx context.Context, id string) (*models.Chapter, error)
	Update(ctx context.Context, id string, book *models.Chapter) (*models.Chapter, error)
	Delete(ctx context.Context, id string) (*models.Chapter, error)
	List(ctx context.Context) ([]models.Chapter, error)

	GetChapterByBookID(ctx context.Context, bookID string) ([]models.Chapter, error)
}

type repository struct {
	db postgres.Client
}

// NewRepository создает новый экземпляр репозитория для книг.
func NewRepository(database postgres.Client) Repository {
	return &repository{database}
}

func allItems() string {
	cols := []string{columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt, columnTitle, columnNumber,
		columnText, columnBookID, columnIsPublic, columnMapLink, columnMapParamsID}

	return strings.Join(cols, ", ")
}

func insertItems() string {
	cols := []string{columnTitle, columnNumber, columnText, columnBookID, columnIsPublic, columnMapLink,
		columnMapParamsID}

	return strings.Join(cols, ", ")
}

func (r *repository) Create(ctx context.Context, chapter *models.Chapter) (string, error) {
	query := "INSERT INTO" + " " + tableName + ` (` + insertItems() +
		`) VALUES ($1, $2, $3, $4, $5, $6, $7) ` + Returning + columnID

	q := postgres.Query{Name: "ChapterRepository.Create", Raw: query}
	row := r.db.DB().QueryRowContext(ctx, q, chapter.Title, chapter.Number, chapter.Text, chapter.BookID,
		chapter.IsPublic, chapter.MapLink, chapter.MapParamsID)

	var id string
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (*models.Chapter, error) {
	query := querybuilder.SelectWhere(allItems, tableName, columnID) + " AND" + NotDeleted

	fmt.Println(query)

	q := postgres.Query{Name: "ChapterRepository.FindByID", Raw: query}
	row := r.db.DB().QueryRowContext(ctx, q, id)

	return readItem(row)
}

func (r *repository) Update(ctx context.Context, id string, updChapter *models.Chapter) (*models.Chapter, error) {
	query := "Update " + tableName + " SET " +
		querybuilder.ParamsToQuery(", ", columnTitle, columnNumber, columnText, columnBookID,
			columnIsPublic, columnMapLink, columnMapParamsID) +
		" WHERE " + columnID + " = $8" + Returning + allItems()

	args := []interface{}{updChapter.Title, updChapter.Number, updChapter.Text, updChapter.BookID,
		updChapter.IsPublic, updChapter.MapLink, updChapter.MapParamsID, id}

	q := postgres.Query{Name: "BookRepository.Update", Raw: query}
	row := r.db.DB().QueryRowContext(ctx, q, args...)

	return readItem(row)
}

func (r *repository) Delete(ctx context.Context, id string) (*models.Chapter, error) {
	query := querybuilder.DeleteQuery(tableName, columnID) + ` RETURNING ` + allItems()
	q := postgres.Query{Name: "BookRepository.Delete", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, id)

	return readItem(row)
}

func (r *repository) List(ctx context.Context) ([]models.Chapter, error) {
	query := `SELECT ` + allItems() + FromTable + WHERE + NotDeleted

	q := postgres.Query{Name: "ChapterRepository.List", Raw: query}

	rows, err := r.db.DB().QueryRawContextMulti(ctx, q)
	if err != nil {
		return nil, errors.New("params error")
	}

	return readList(rows)
}

func (r *repository) GetChapterByBookID(ctx context.Context, bookID string) ([]models.Chapter, error) {
	query := querybuilder.SelectWhere(allItems, tableName, columnBookID) + " AND" + NotDeleted

	q := postgres.Query{Name: "ChapterRepository.GetChapterByBookID", Raw: query}

	rows, err := r.db.DB().QueryRawContextMulti(ctx, q, bookID)
	if err != nil {
		return nil, errors.New("params error")
	}

	return readList(rows)
}
