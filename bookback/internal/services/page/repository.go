package page

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/services"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"strings"
)

const (
	tableName         = "pages"
	columnID          = "id"
	columnCreatedAt   = "created_at"
	columnUpdatedAt   = "updated_at"
	columnDeletedAt   = "deleted_at"
	columnTitle       = "title"
	columnText        = "text"
	columnChapterID   = "chapter_id"
	columnIsPublic    = "is_public"
	columnMapParamsID = "map_params_id"
	Returning         = " RETURNING "
)

// Repository определяет интерфейс для взаимодействия с хранилищем книг.
type Repository interface {
	Create(ctx context.Context, book *models.Page) (string, error)
	FindByID(ctx context.Context, id string) (*models.Page, error)
	Update(ctx context.Context, id string, book *models.Page) (*models.Page, error)
	Delete(ctx context.Context, id string) (*models.Page, error)
	List(ctx context.Context) ([]models.Page, error)

	GetPagesByChapterID(ctx context.Context, chapterID string) ([]models.Page, error)
}

type repository struct {
	db db.Client
}

// NewRepository создает новый экземпляр репозитория для книг.
func NewRepository(database db.Client) Repository {
	return &repository{database}
}

func allItems() string {
	cols := []string{columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt, columnTitle, columnText, columnChapterID,
		columnIsPublic, columnMapParamsID}

	return strings.Join(cols, ", ")
}

func insertItems() string {
	cols := []string{columnTitle, columnText, columnChapterID, columnIsPublic, columnMapParamsID}

	return strings.Join(cols, ", ")
}

// Create inserts a new page into the database and returns its ID.
func (r *repository) Create(ctx context.Context, page *models.Page) (string, error) {
	query := `INSERT INTO` + " " + tableName + ` (` + insertItems() + `) VALUES ($1, $2, $3, $4, $5) ` +
		Returning + columnID
	args := []interface{}{page.Title, page.Text, page.ChapterID, page.IsPublic, page.MapParamsID}

	q := db.Query{Name: "PageRepository.Create", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, args...)
	var id string
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

// FindByID retrieves a page by its ID.
func (r *repository) FindByID(ctx context.Context, id string) (*models.Page, error) {
	query := services.SelectWhere(allItems, tableName, columnID)
	q := db.Query{Name: "PageRepository.FindByID", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, id)

	return readItem(row)
}

// Update modifies an existing page's data.
func (r *repository) Update(ctx context.Context, id string, updPage *models.Page) (*models.Page, error) {
	query := "Update " + tableName + " SET " +
		services.ParamsToQuery(", ", columnTitle, columnText, columnChapterID, columnIsPublic, columnMapParamsID) +
		" WHERE " + columnID + " = $6" + Returning + allItems()

	args := []interface{}{updPage.Title, updPage.Text, updPage.ChapterID, updPage.IsPublic, updPage.MapParamsID, id}

	q := db.Query{Name: "PageRepository.Update", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, args...)

	return readItem(row)
}

// Delete removes a page from the database.
func (r *repository) Delete(ctx context.Context, id string) (*models.Page, error) {
	query := services.DeleteQuery(tableName, columnID) + Returning + allItems()

	q := db.Query{Name: "PageRepository.Delete", Raw: query}
	row := r.db.DB().QueryRowContext(ctx, q, id)

	return readItem(row)
}

// List retrieves all pages for a given chapter ID.
func (r *repository) List(ctx context.Context) ([]models.Page, error) {
	query := `Select ` + allItems() + ` FROM ` + tableName + ` WHERE ` + columnDeletedAt + ` IS NULL`

	q := db.Query{Name: "PageRepository.List", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	return readList(rows)
}

// GetPagesByChapterID retrieves all pages for a given chapter ID.
func (r *repository) GetPagesByChapterID(ctx context.Context, chapterID string) ([]models.Page, error) {
	query := services.SelectWhere(allItems, tableName, columnChapterID) + ` AND ` + columnDeletedAt + ` IS NULL`

	q := db.Query{Name: "PageRepository.GetPagesByChapterID", Raw: query}

	rows, err := r.db.DB().QueryContext(ctx, q, chapterID)
	if err != nil {
		return nil, err
	}

	return readList(rows)
}
