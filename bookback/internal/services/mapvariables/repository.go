package mapvariables

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/services"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"strings"
)

const (
	tableName         = "map_variables"
	columnID          = "id"
	columnInsertedAt  = "inserted_at"
	columnBookID      = "book_id"
	columnChapterID   = "chapter_id"
	columnPageID      = "page_id"
	columnParagraphID = "paragraph_id"
	columnMapLink     = "map_link"
	columnLat         = "lat"
	columnLng         = "lng"
	columnZoom        = "zoom"
	columnDate        = "date"
	columnDescription = "description"
	columnLink        = "link"
	columnLinkText    = "link_text"
	columnLinkType    = "link_type"
	columnLinkImage   = "link_image"
	columnImage       = "image"
)

type Repository interface {
	Create(ctx context.Context, variable *models.MapVariable) (string, error)
	FindByID(ctx context.Context, id string) (*models.MapVariable, error)
	Update(ctx context.Context, id string, variable *models.MapVariable) (*models.MapVariable, error)
	Delete(ctx context.Context, id string) (*models.MapVariable, error)
	GetByBookID(ctx context.Context, bookID string) ([]models.MapVariable, error)
	GetByChapterID(ctx context.Context, chapterID string) ([]models.MapVariable, error)
	GetByPageID(ctx context.Context, pageID string) ([]models.MapVariable, error)
	GetByParagraphID(ctx context.Context, paragraphID string) ([]models.MapVariable, error)

	GetByMapLinkAndBookID(ctx context.Context, mapLink, bookID string) ([]models.MapVariable, error)
	GetByMapLinkAndChapterID(ctx context.Context, mapLink, chapterID string) ([]models.MapVariable, error)
	GetByMapLinkAndPageID(ctx context.Context, mapLink, pageID string) ([]models.MapVariable, error)
	GetByMapLinkAndParagraphID(ctx context.Context, mapLink, paragraphID string) ([]models.MapVariable, error)
}

type repository struct {
	db db.Client
}

func allItems() string {
	cols := []string{columnID, columnInsertedAt, columnBookID, columnChapterID, columnPageID,
		columnParagraphID, columnMapLink, columnLat, columnLng, columnZoom, columnDate,
		columnDescription, columnLink, columnLinkText, columnLinkType, columnLinkImage, columnImage}

	return strings.Join(cols, ", ")
}

func NewRepository(database db.Client) Repository {
	return &repository{database}
}

func (r *repository) Create(ctx context.Context, variable *models.MapVariable) (string, error) {
	query := `INSERT INTO` + " " + tableName + ` (` + allItems() +
		`) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id`

	args := []interface{}{variable.ID, variable.CreatedAt, variable.BookID, variable.ChapterID, //nolint:gofmt
		variable.PageID, variable.ParagraphID, variable.MapLink, variable.Lat, variable.Lng,
		variable.Zoom, variable.Date, variable.Description, variable.Link, variable.LinkText,
		variable.LinkType, variable.LinkImage, variable.Image}

	q := db.Query{Name: "MapVariableRepository.Insert", Raw: query}

	var id string
	if err := r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (*models.MapVariable, error) {
	query := services.SelectWhere(allItems, tableName, columnID)

	q := db.Query{Name: "MapVariableRepository.FindByID", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, id)

	return readItem(row)
}

func (r *repository) Update(ctx context.Context, id string, variable *models.MapVariable) (*models.MapVariable, error) {
	query := `UPDATE ` + tableName + ` SET ` +
		columnInsertedAt + ` = $1, ` + columnBookID + ` = $2, ` + columnChapterID + ` = $3, ` +
		columnPageID + ` = $4, ` + columnParagraphID + ` = $5, ` + columnMapLink + ` = $6, ` +
		columnLat + ` = $7, ` + columnLng + ` = $8, ` + columnZoom + ` = $9, ` +
		columnDate + ` = $10, ` + columnDescription + ` = $11, ` + columnLink + ` = $12, ` +
		columnLinkText + ` = $13, ` + columnLinkType + ` = $14, ` + columnLinkImage + ` = $15, ` +
		columnImage + ` = $16 WHERE ` + columnID + ` = $17 RETURNING ` + allItems()

	args := []interface{}{variable.CreatedAt, variable.BookID, variable.ChapterID, variable.PageID, //nolint:gofmt
		variable.ParagraphID, variable.MapLink, variable.Lat, variable.Lng, variable.Zoom,
		variable.Date, variable.Description, variable.Link, variable.LinkText, variable.LinkType,
		variable.LinkImage, variable.Image, id}

	q := db.Query{Name: "MapVariableRepository.Update", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, args...)

	return readItem(row)
}

func (r *repository) Delete(ctx context.Context, id string) (*models.MapVariable, error) {
	query := `DELETE FROM` + " " + tableName + ` WHERE ` + columnID + ` = $1 RETURNING ` + allItems()

	q := db.Query{Name: "MapVariableRepository.Delete", Raw: query}
	row := r.db.DB().QueryRowContext(ctx, q, id)
	return readItem(row)
}

func (r *repository) GetByBookID(ctx context.Context, bookID string) ([]models.MapVariable, error) {
	query := services.SelectWhere(allItems, tableName, columnBookID)

	q := db.Query{Name: "MapVariableRepository.GetByBookID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, bookID)
	if err != nil {
		return nil, err
	}

	return readList(rows)
}

func (r *repository) GetByChapterID(ctx context.Context, chapterID string) ([]models.MapVariable, error) {
	query := services.SelectWhere(allItems, tableName, columnChapterID)

	q := db.Query{Name: "MapVariableRepository.GetByChapterID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, chapterID)
	if err != nil {
		return nil, err
	}

	return readList(rows)
}

func (r *repository) GetByPageID(ctx context.Context, pageID string) ([]models.MapVariable, error) {
	query := services.SelectWhere(allItems, tableName, columnPageID)

	q := db.Query{Name: "MapVariableRepository.GetByPageID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, pageID)
	if err != nil {
		return nil, err
	}

	return readList(rows)
}

func (r *repository) GetByParagraphID(ctx context.Context, paragraphID string) ([]models.MapVariable, error) {
	query := services.SelectWhere(allItems, tableName, columnParagraphID)

	q := db.Query{Name: "MapVariableRepository.GetByParagraphID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, paragraphID)
	if err != nil {
		return nil, err
	}

	return readList(rows)
}

func (r *repository) GetByMapLinkAndBookID(ctx context.Context, mapLink, bookID string) ([]models.MapVariable, error) {
	query := services.SelectWhere(allItems, tableName, columnMapLink, columnBookID)

	q := db.Query{Name: "MapVariableRepository.GetByMapLinkAndBookID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, mapLink, bookID)
	if err != nil {
		return nil, err
	}

	return readList(rows)
}

func (r *repository) GetByMapLinkAndChapterID(ctx context.Context, mapLink, chapterID string) ([]models.MapVariable, error) {
	query := services.SelectWhere(allItems, tableName, columnMapLink, columnChapterID)

	q := db.Query{Name: "MapVariableRepository.GetByMapLinkAndChapterID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, mapLink, chapterID)
	if err != nil {
		return nil, err
	}

	return readList(rows)
}

func (r *repository) GetByMapLinkAndPageID(ctx context.Context, mapLink, pageID string) ([]models.MapVariable, error) {
	query := services.SelectWhere(allItems, tableName, columnMapLink, columnPageID)

	q := db.Query{Name: "MapVariableRepository.GetByMapLinkAndPageID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, mapLink, pageID)
	if err != nil {
		return nil, err
	}

	return readList(rows)
}

func (r *repository) GetByMapLinkAndParagraphID(ctx context.Context, mapLink, paragraphID string) ([]models.MapVariable, error) {
	query := services.SelectWhere(allItems, tableName, columnMapLink, columnParagraphID)

	q := db.Query{Name: "MapVariableRepository.GetByMapLinkAndParagraphID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, mapLink, paragraphID)
	if err != nil {
		return nil, err
	}

	return readList(rows)
}
