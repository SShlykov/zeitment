package bookevents

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/services"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"strings"
)

const (
	tableName         = "book_events"
	columnID          = "id"
	columnCreatedAt   = "created_at"
	columnUpdatedAt   = "updated_at"
	columnBookID      = "book_id"
	columnChapterID   = "chapter_id"
	columnPageID      = "page_id"
	columnParagraphID = "paragraph_id"
	columnEventType   = "event_type"
	columnIsPublic    = "is_public"
	columnKey         = "key"
	columnValue       = "value"
	columnLink        = "link"
	columnLinkText    = "link_text"
	columnLinkType    = "link_type"
	columnLinkImage   = "link_image"
	columnDescription = "description"
)

type Repository interface {
	Create(ctx context.Context, event *models.BookEvent) (string, error)
	FindByID(ctx context.Context, id string) (*models.BookEvent, error)
	Update(ctx context.Context, id string, event *models.BookEvent) (*models.BookEvent, error)
	Delete(ctx context.Context, id string) (*models.BookEvent, error)
	GetByBookID(ctx context.Context, bookID string) ([]models.BookEvent, error)
	GetByChapterID(ctx context.Context, chapterID string) ([]models.BookEvent, error)
	GetByPageID(ctx context.Context, pageID string) ([]models.BookEvent, error)
	GetByParagraphID(ctx context.Context, paragraphID string) ([]models.BookEvent, error)
}

type repository struct {
	db db.Client
}

func NewRepository(database db.Client) Repository {
	return &repository{database}
}

func allItems() string {
	cols := []string{columnID, columnCreatedAt, columnUpdatedAt, columnBookID, columnChapterID,
		columnPageID, columnParagraphID, columnEventType, columnIsPublic, columnKey,
		columnValue, columnLink, columnLinkText, columnLinkType, columnLinkImage,
		columnDescription}

	return strings.Join(cols, ", ")
}

func insertItems() string {
	cols := []string{columnBookID, columnChapterID, columnPageID, columnParagraphID, columnEventType,
		columnIsPublic, columnKey, columnValue, columnLink, columnLinkText,
		columnLinkType, columnLinkImage, columnDescription}

	return strings.Join(cols, ", ")
}

func (r *repository) Create(ctx context.Context, event *models.BookEvent) (string, error) {
	query := `INSERT INTO` + " " + tableName + ` (` + insertItems() +
		`) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING ` +
		columnID

	args := []interface{}{event.BookID, event.ChapterID, event.PageID, event.ParagraphID, event.EventType, //nolint:gofmt
		event.IsPublic, event.Key, event.Value, event.Link, event.LinkText, event.LinkType,
		event.LinkImage, event.Description}

	q := db.Query{Name: "BookEventRepository.Insert", Raw: query}

	var id string
	if err := r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (*models.BookEvent, error) {
	query := services.SelectWhere(allItems, tableName, columnID)

	q := db.Query{Name: "BookEventRepository.FindByID", Raw: query}
	row := r.db.DB().QueryRowContext(ctx, q, id)
	return readItem(row)
}

func (r *repository) Update(ctx context.Context, id string, event *models.BookEvent) (*models.BookEvent, error) {
	query := `UPDATE ` + tableName + ` SET ` +
		services.ParamsToQuery(", ",
			columnBookID, columnChapterID, columnPageID, columnParagraphID, columnEventType,
			columnIsPublic, columnKey, columnValue, columnLink, columnLinkText,
			columnLinkType, columnLinkImage, columnDescription) + ` WHERE ` +
		columnID + ` = $14` + ` RETURNING ` + allItems()

	args := []interface{}{event.BookID, event.ChapterID, event.PageID, event.ParagraphID, event.EventType, //nolint:gofmt
		event.IsPublic, event.Key, event.Value, event.Link, event.LinkText,
		event.LinkType, event.LinkImage, event.Description, id}

	q := db.Query{Name: "BookEventRepository.Update", Raw: query}

	row := r.db.DB().QueryRowContext(ctx, q, args...)

	return readItem(row)
}

func (r *repository) Delete(ctx context.Context, id string) (*models.BookEvent, error) {
	query := services.DeleteQuery(tableName, columnID) + ` RETURNING ` + allItems()

	q := db.Query{Name: "BookEventRepository.Delete", Raw: query}
	row := r.db.DB().QueryRowContext(ctx, q, id)
	return readItem(row)
}

func (r *repository) GetByBookID(ctx context.Context, bookID string) ([]models.BookEvent, error) {
	query := services.SelectWhere(allItems, tableName, columnBookID)

	q := db.Query{Name: "BookEventRepository.GetByBookID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, bookID)
	if err != nil {
		return nil, err
	}
	return readList(rows)
}

func (r *repository) GetByChapterID(ctx context.Context, chapterID string) ([]models.BookEvent, error) {
	query := services.SelectWhere(allItems, tableName, columnChapterID)

	q := db.Query{Name: "BookEventRepository.GetByChapterID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, chapterID)
	if err != nil {
		return nil, err
	}
	return readList(rows)
}

func (r *repository) GetByPageID(ctx context.Context, pageID string) ([]models.BookEvent, error) {
	query := services.SelectWhere(allItems, tableName, columnPageID)

	q := db.Query{Name: "BookEventRepository.GetByPageID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, pageID)
	if err != nil {
		return nil, err
	}
	return readList(rows)
}

func (r *repository) GetByParagraphID(ctx context.Context, paragraphID string) ([]models.BookEvent, error) {
	query := services.SelectWhere(allItems, tableName, columnParagraphID)

	q := db.Query{Name: "BookEventRepository.GetByParagraphID", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, paragraphID)
	if err != nil {
		return nil, err
	}
	return readList(rows)
}