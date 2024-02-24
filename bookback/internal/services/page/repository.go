package page

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/SShlykov/zeitment/bookback/pkg/db/sq"
	"github.com/google/uuid"
)

const (
	// model fields and table name for books table
	tableName       = "pages"
	columnID        = "id"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
	columnDeletedAt = "deleted_at"
	columnText      = "text"
	columnChapterID = "chapter_id"
	columnIsPublic  = "is_public"
	Returning       = "RETURNING "
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

// Create inserts a new page into the database and returns its ID.
func (r *repository) Create(ctx context.Context, page *models.Page) (string, error) {
	page.ID = uuid.New().String()
	query, args, err := sq.Insert(tableName).
		Columns(columnID, columnText, columnChapterID, columnIsPublic).
		Values(page.ID, page.Text, page.ChapterID, page.IsPublic).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{Name: "PageRepository.Create", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, args...)
	var id string
	if err = row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

// FindByID retrieves a page by its ID.
func (r *repository) FindByID(ctx context.Context, id string) (*models.Page, error) {
	query, args, err := sq.Select(columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt, columnText,
		columnChapterID, columnIsPublic).
		From("pages").
		Where(sq.Eq{"id": id, "deleted_at": nil}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{Name: "PageRepository.FindByID", Raw: query}

	var page models.Page
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&page.ID, &page.CreatedAt, &page.UpdatedAt,
		&page.DeletedAt, &page.Text, &page.ChapterID, &page.IsPublic); err != nil {
		return nil, err
	}

	return &page, nil
}

// Update modifies an existing page's data.
func (r *repository) Update(ctx context.Context, id string, updPage *models.Page) (*models.Page, error) {
	query, args, err := sq.Update(tableName).
		Set(columnText, updPage.Text).
		Set(columnIsPublic, updPage.IsPublic).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id, created_at, updated_at, deleted_at, text, chapter_id, is_public").
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{Name: "PageRepository.Update", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, args...)

	return readItem(row)
}

// Delete removes a page from the database.
func (r *repository) Delete(ctx context.Context, id string) (*models.Page, error) {
	query, args, err := sq.SQ.Delete(tableName).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{Name: "PageRepository.Delete", Raw: query}

	var deletedID string
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&deletedID); err != nil {
		return nil, err
	}

	return &models.Page{ID: deletedID}, nil
}

// List retrieves all pages for a given chapter ID.
func (r *repository) List(ctx context.Context) ([]models.Page, error) {
	query :=
		`SELECT id, created_at, updated_at, deleted_at, text, chapter_id, is_public 
		 FROM pages 
		 WHERE deleted_at IS NULL`

	q := db.Query{Name: "PageRepository.List", Raw: query}

	rows, err := r.db.DB().QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return readList(rows)
}

// GetPagesByChapterID retrieves all pages for a given chapter ID.
func (r *repository) GetPagesByChapterID(ctx context.Context, chapterID string) ([]models.Page, error) {
	query :=
		`SELECT id, created_at, updated_at, deleted_at, text, chapter_id, is_public 
		 FROM pages 
		 WHERE chapter_id = $1 AND deleted_at IS NULL`

	q := db.Query{Name: "PageRepository.GetPagesByChapterID", Raw: query}

	rows, err := r.db.DB().QueryContext(ctx, q, chapterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return readList(rows)
}
