package paragraph

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/SShlykov/zeitment/bookback/pkg/db/sq"
	"github.com/google/uuid"
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
	Returning       = "RETURNING "
)

// Repository определяет интерфейс для взаимодействия с хранилищем книг.
type Repository interface {
	Create(ctx context.Context, paragraph *models.Paragraph) (string, error)
	FindByID(ctx context.Context, id string) (*models.Paragraph, error)
	Update(ctx context.Context, id string, updParagraph *models.Paragraph) (*models.Paragraph, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]models.Paragraph, error)
}

type repository struct {
	db db.Client
}

// NewRepository создает новый экземпляр репозитория для книг.
func NewRepository(database db.Client) Repository {
	return &repository{database}
}

// Create inserts a new Paragraph into the database
func (r *repository) Create(ctx context.Context, paragraph *models.Paragraph) (string, error) {
	paragraph.ID = uuid.New().String() // Generate a new UUID for the paragraph
	query, args, err := sq.SQ.Insert(tableName).
		Columns(columnID, columnText, columnIsPublic).
		Values(paragraph.ID, paragraph.Text, paragraph.IsPublic).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{Name: "ParagraphRepository.Create", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, args...)
	var id string
	if err = row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

// FindByID retrieves a paragraph by its ID
func (r *repository) FindByID(ctx context.Context, id string) (*models.Paragraph, error) {
	query, args, err := sq.SQ.Select(columnID, columnCreatedAt, columnUpdatedAt, columnDeletedAt, columnText, columnIsPublic).
		From(tableName).
		Where(sq.Eq{"id": id, "deleted_at": nil}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{Name: "ParagraphRepository.FindByID", Raw: query}

	var paragraph models.Paragraph
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&paragraph.ID, &paragraph.CreatedAt, &paragraph.UpdatedAt,
		&paragraph.DeletedAt, &paragraph.Text, &paragraph.IsPublic); err != nil {
		return nil, err
	}

	return &paragraph, nil
}

// Update modifies an existing paragraph's data
func (r *repository) Update(ctx context.Context, id string, updParagraph *models.Paragraph) (*models.Paragraph, error) {
	query, args, err := sq.Update(tableName).
		Set(columnText, updParagraph.Text).
		Set(columnIsPublic, updParagraph.IsPublic).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id, created_at, updated_at, deleted_at, text, is_public").
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{Name: "ParagraphRepository.Update", Raw: query}

	var paragraph models.Paragraph
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&paragraph.ID, &paragraph.CreatedAt, &paragraph.UpdatedAt,
		&paragraph.DeletedAt, &paragraph.Text, &paragraph.IsPublic); err != nil {
		return nil, err
	}

	return &paragraph, nil
}

// Delete removes a paragraph from the database
func (r *repository) Delete(ctx context.Context, id string) error {
	query, args, err := sq.SQ.Delete(tableName).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return err
	}

	q := db.Query{Name: "ParagraphRepository.Delete", Raw: query}

	var deletedID string
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&deletedID); err != nil {
		return err
	}

	return nil
}

// List retrieves all paragraphs (adjust based on your needs, e.g., by parent Page or Chapter ID)
func (r *repository) List(ctx context.Context) ([]models.Paragraph, error) {
	query, args, err := sq.SQ.Select("id", "created_at", "updated_at", "deleted_at", "text", "is_public").
		From(tableName).
		Where(sq.Eq{"deleted_at": nil}). // Adjust if you need to filter by a parent entity
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{Name: "ParagraphRepository.List", Raw: query}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paragraphs []models.Paragraph
	for rows.Next() {
		var paragraph models.Paragraph
		if err = rows.Scan(&paragraph.ID, &paragraph.CreatedAt, &paragraph.UpdatedAt, &paragraph.DeletedAt,
			&paragraph.Text, &paragraph.IsPublic); err != nil {
			return nil, err
		}
		paragraphs = append(paragraphs, paragraph)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return paragraphs, nil
}
