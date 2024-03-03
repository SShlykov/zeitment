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
	columnCreatedAt   = "created_at"
	columnUpdatedAt   = "updated_at"
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
}

type repository struct {
	db db.Client
}

func NewRepository(database db.Client) Repository {
	return &repository{database}
}

func allItems() string {
	cols := []string{columnID, columnCreatedAt, columnUpdatedAt, columnBookID, columnChapterID, columnPageID,
		columnParagraphID, columnMapLink, columnLat, columnLng, columnZoom, columnDate,
		columnDescription, columnLink, columnLinkText, columnLinkType, columnLinkImage, columnImage}

	return strings.Join(cols, ", ")
}

func insertItems() string {
	cols := []string{columnBookID, columnChapterID, columnPageID, columnParagraphID, columnMapLink,
		columnLat, columnLng, columnZoom, columnDate, columnDescription, columnLink, columnLinkText,
		columnLinkType, columnLinkImage, columnImage}

	return strings.Join(cols, ", ")
}

func (r *repository) Create(ctx context.Context, variable *models.MapVariable) (string, error) {
	query := `INSERT INTO` + " " + tableName + ` (` + insertItems() +
		`) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id`

	args := []interface{}{variable.BookID, variable.ChapterID, variable.PageID, variable.ParagraphID, variable.MapLink, //nolint:gofmt
		variable.Lat, variable.Lng, variable.Zoom, variable.Date, variable.Description,
		variable.Link, variable.LinkText, variable.LinkType, variable.LinkImage, variable.Image}

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
	query := `UPDATE ` + tableName + ` SET ` + services.ParamsToQuery(", ",
		columnBookID, columnChapterID, columnPageID, columnParagraphID, columnMapLink,
		columnLat, columnLng, columnZoom, columnDate, columnDescription, columnLink, columnLinkText,
		columnLinkType, columnLinkImage, columnImage) +
		` WHERE ` + columnID + ` = $16 RETURNING ` + allItems()

	args := []interface{}{variable.BookID, variable.ChapterID, variable.PageID, variable.ParagraphID, variable.MapLink, //nolint:gofmt
		variable.Lat, variable.Lng, variable.Zoom, variable.Date, variable.Description,
		variable.Link, variable.LinkText, variable.LinkType, variable.LinkImage, variable.Image, id}

	q := db.Query{Name: "MapVariableRepository.Update", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, args...)

	return readItem(row)
}

func (r *repository) Delete(ctx context.Context, id string) (*models.MapVariable, error) {
	query := services.DeleteQuery(tableName, columnID) + ` RETURNING ` + allItems()

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
