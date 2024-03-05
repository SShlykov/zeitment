package paragraphrepo

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/SShlykov/zeitment/bookback/pkg/querybuilder"
	"strings"
)

const (
	tableName       = "paragraphs"
	columnID        = "id"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
	columnDeletedAt = "deleted_at"
	columnTitle     = "title"
	columnText      = "text"
	columnType      = "type"
	columnIsPublic  = "is_public"
	columnPageID    = "page_id"
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
	db postgres.Client
}

// NewRepository создает новый экземпляр репозитория для книг.
func NewRepository(database postgres.Client) Repository {
	return &repository{database}
}

func allItems() string {
	cols := []string{columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt, columnTitle,
		columnText, columnType, columnIsPublic, columnPageID}

	return strings.Join(cols, ", ")
}

func insertItems() string {
	cols := []string{columnTitle, columnText, columnType, columnIsPublic, columnPageID}

	return strings.Join(cols, ", ")
}

// Create inserts a new Paragraph into the database
func (r *repository) Create(ctx context.Context, paragraphDto *models.Paragraph) (string, error) {
	query := "INSERT INTO" + " " + tableName + " (" + insertItems() + ") VALUES ($1, $2, $3, $4, $5) " +
		Returning + columnID
	q := postgres.Query{Name: "ParagraphRepository.Create", Raw: query}

	args := []interface{}{paragraphDto.Title, paragraphDto.Text, paragraphDto.Type, paragraphDto.IsPublic, paragraphDto.PageID}

	row := r.db.DB().QueryRowContext(ctx, q, args...)
	var id string
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

// FindByID retrieves a paragraph by its ID
func (r *repository) FindByID(ctx context.Context, id string) (*models.Paragraph, error) {
	query := querybuilder.SelectWhere(allItems, tableName, columnID) + " AND deleted_at IS NULL"

	q := postgres.Query{Name: "ParagraphRepository.FindByID", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, id)

	return readItem(row)
}

// Update modifies an existing paragraph's data
func (r *repository) Update(ctx context.Context, id string, updParagraph *models.Paragraph) (*models.Paragraph, error) {
	query := "UPDATE" + " " + tableName + " SET " +
		querybuilder.ParamsToQuery(", ", columnTitle, columnText, columnType, columnIsPublic, columnPageID) +
		" WHERE id = $6" + Returning + allItems()

	args := []interface{}{updParagraph.Title, updParagraph.Text, updParagraph.Type, updParagraph.IsPublic,
		updParagraph.PageID, id}

	q := postgres.Query{Name: "ParagraphRepository.Update", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, args...)

	return readItem(row)
}

// Delete removes a paragraph from the database
func (r *repository) Delete(ctx context.Context, id string) (*models.Paragraph, error) {
	query := querybuilder.DeleteQuery(tableName, columnID) + Returning + allItems()
	q := postgres.Query{Name: "ParagraphRepository.Delete", Raw: query}
	row := r.db.DB().QueryRowContext(ctx, q, id)

	return readItem(row)
}

// List retrieves all paragraphs (adjust based on your needs, e.g., by parent Page or Chapter ID)
func (r *repository) List(ctx context.Context) ([]models.Paragraph, error) {
	query := "SELECT " + allItems() + " FROM " + tableName + " WHERE deleted_at IS NULL"

	q := postgres.Query{Name: "ParagraphRepository.List", Raw: query}

	rows, err := r.db.DB().QueryContext(ctx, q)
	if err != nil {
		return nil, errors.New("params error")
	}

	return readList(rows)
}

func (r *repository) GetParagraphsByPageID(ctx context.Context, pageID string) ([]models.Paragraph, error) {
	query := querybuilder.SelectWhere(allItems, tableName, columnPageID) + " AND deleted_at IS NULL"

	q := postgres.Query{Name: "ParagraphRepository.GetParagraphsByPageID", Raw: query}

	rows, err := r.db.DB().QueryContext(ctx, q, pageID)
	if err != nil {
		return nil, errors.New("params error")
	}

	return readList(rows)
}
