package bookevents

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
)

type Service interface {
	CreateBookEvent(ctx context.Context, event *models.BookEvent) (*models.BookEvent, error)
	GetBookEventByID(ctx context.Context, id string) (*models.BookEvent, error)
	UpdateBookEvent(ctx context.Context, id string, event *models.BookEvent) (*models.BookEvent, error)
	DeleteBookEvent(ctx context.Context, id string) (*models.BookEvent, error)
	GetBookEventsByBookID(ctx context.Context, bookID string) ([]models.BookEvent, error)
	GetBookEventsByChapterID(ctx context.Context, chapterID string) ([]models.BookEvent, error)
	GetBookEventsByPageID(ctx context.Context, pageID string) ([]models.BookEvent, error)
	GetBookEventsByParagraphID(ctx context.Context, paragraphID string) ([]models.BookEvent, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateBookEvent(ctx context.Context, event *models.BookEvent) (*models.BookEvent, error) {
	id, err := s.repo.Create(ctx, event)
	if err != nil {
		return nil, err
	}
	event.ID = id

	return event, err
}

func (s *service) GetBookEventByID(ctx context.Context, id string) (*models.BookEvent, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateBookEvent(ctx context.Context, id string, event *models.BookEvent) (*models.BookEvent, error) {
	return s.repo.Update(ctx, id, event)
}

func (s *service) DeleteBookEvent(ctx context.Context, id string) (*models.BookEvent, error) {
	return s.repo.Delete(ctx, id)
}

func (s *service) GetBookEventsByBookID(ctx context.Context, bookID string) ([]models.BookEvent, error) {
	return s.repo.GetByBookID(ctx, bookID)
}

func (s *service) GetBookEventsByChapterID(ctx context.Context, chapterID string) ([]models.BookEvent, error) {
	return s.repo.GetByChapterID(ctx, chapterID)
}

func (s *service) GetBookEventsByPageID(ctx context.Context, pageID string) ([]models.BookEvent, error) {
	return s.repo.GetByPageID(ctx, pageID)
}

func (s *service) GetBookEventsByParagraphID(ctx context.Context, paragraphID string) ([]models.BookEvent, error) {
	return s.repo.GetByParagraphID(ctx, paragraphID)
}
