package paragraph

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"strings"
)

const (
	// model fields and table name for books table
	tableName       = "paragraphs"
	columnID        = "id"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
	columnDeletedAt = "deleted_at"
	columnText      = "text"
	columnIsPublic  = "is_public"
	Returning       = " RETURNING "
)

// Repository определяет интерфейс для взаимодействия с хранилищем книг.
type Repository interface {
	Create(ctx context.Context, paragraph *models.Paragraph) (string, error)
	FindByID(ctx context.Context, id string) (*models.Paragraph, error)
	Update(ctx context.Context, id string, updParagraph *models.Paragraph) (*models.Paragraph, error)
	Delete(ctx context.Context, id string) (*models.Paragraph, error)
	List(ctx context.Context) ([]models.Paragraph, error)

	GetParagraphsByPageID(ctx context.Context, pageID string) ([]models.Paragraph, error)
}

type repository struct {
	db db.Client
}

func allItems() string {
	cols := []string{columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt, columnText, columnIsPublic}

	return strings.Join(cols, ", ")
}

// NewRepository создает новый экземпляр репозитория для книг.
func NewRepository(database db.Client) Repository {
	return &repository{database}
}

// Create inserts a new Paragraph into the database
func (r *repository) Create(ctx context.Context, paragraph *models.Paragraph) (string, error) {
	query := "INSERT INTO paragraphs (text, is_public, page_id) VALUES ($1, $2, $3) RETURNING id"

	q := db.Query{Name: "ParagraphRepository.Create", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, paragraph.Text, paragraph.IsPublic, paragraph.PageID)
	var id string
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

// FindByID retrieves a paragraph by its ID
func (r *repository) FindByID(ctx context.Context, id string) (*models.Paragraph, error) {
	query := "SELECT " + allItems() + " FROM " + tableName + " WHERE id = $1 AND deleted_at IS NULL"

	q := db.Query{Name: "ParagraphRepository.FindByID", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, id)

	return readItem(row)
}

// Update modifies an existing paragraph's data
func (r *repository) Update(ctx context.Context, id string, updParagraph *models.Paragraph) (*models.Paragraph, error) {
	query := "UPDATE paragraphs SET text = $1, is_public = $2 WHERE id = $3" + Returning + allItems()

	q := db.Query{Name: "ParagraphRepository.Update", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, updParagraph.Text, updParagraph.IsPublic, id)

	return readItem(row)
}

// Delete removes a paragraph from the database
func (r *repository) Delete(ctx context.Context, id string) (*models.Paragraph, error) {
	paragraph, err := r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	query := "DELETE FROM" + " " + tableName + " WHERE id = $1" + Returning + "id"

	q := db.Query{Name: "ParagraphRepository.Delete", Raw: query}

	var deletedID string
	if err = r.db.DB().QueryRowContext(ctx, q, id).Scan(&deletedID); err != nil {
		return nil, err
	}

	return paragraph, nil
}

// List retrieves all paragraphs (adjust based on your needs, e.g., by parent Page or Chapter ID)
func (r *repository) List(ctx context.Context) ([]models.Paragraph, error) {
	query := "SELECT " + allItems() + " FROM " + tableName + " WHERE deleted_at IS NULL"

	q := db.Query{Name: "ParagraphRepository.List", Raw: query}

	rows, err := r.db.DB().QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return readList(rows)
}

func (r *repository) GetParagraphsByPageID(ctx context.Context, pageID string) ([]models.Paragraph, error) {
	query := "SELECT " + allItems() + " FROM " + tableName + " WHERE page_id = $1 AND deleted_at IS NULL"

	q := db.Query{Name: "ParagraphRepository.GetParagraphsByPageID", Raw: query}

	rows, err := r.db.DB().QueryContext(ctx, q, pageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return readList(rows)
}
