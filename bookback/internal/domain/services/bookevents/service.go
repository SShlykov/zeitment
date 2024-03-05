package bookevents

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Service interface {
	CreateBookEvent(ctx context.Context, event *entity.BookEvent) (*entity.BookEvent, error)
	GetBookEventByID(ctx context.Context, id string) (*entity.BookEvent, error)
	UpdateBookEvent(ctx context.Context, id string, event *entity.BookEvent) (*entity.BookEvent, error)
	DeleteBookEvent(ctx context.Context, id string) (*entity.BookEvent, error)
	GetBookEventsByBookID(ctx context.Context, bookID string) ([]entity.BookEvent, error)
	GetBookEventsByChapterID(ctx context.Context, chapterID string) ([]entity.BookEvent, error)
	GetBookEventsByPageID(ctx context.Context, pageID string) ([]entity.BookEvent, error)
	GetBookEventsByParagraphID(ctx context.Context, paragraphID string) ([]entity.BookEvent, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateBookEvent(ctx context.Context, event *entity.BookEvent) (*entity.BookEvent, error) {
	id, err := s.repo.Create(ctx, event)
	if err != nil {
		return nil, err
	}
	event.ID = id

	return event, err
}

func (s *service) GetBookEventByID(ctx context.Context, id string) (*entity.BookEvent, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateBookEvent(ctx context.Context, id string, event *entity.BookEvent) (*entity.BookEvent, error) {
	return s.repo.Update(ctx, id, event)
}

func (s *service) DeleteBookEvent(ctx context.Context, id string) (*entity.BookEvent, error) {
	return s.repo.Delete(ctx, id)
}

func (s *service) GetBookEventsByBookID(ctx context.Context, bookID string) ([]entity.BookEvent, error) {
	return s.repo.GetByBookID(ctx, bookID)
}

func (s *service) GetBookEventsByChapterID(ctx context.Context, chapterID string) ([]entity.BookEvent, error) {
	return s.repo.GetByChapterID(ctx, chapterID)
}

func (s *service) GetBookEventsByPageID(ctx context.Context, pageID string) ([]entity.BookEvent, error) {
	return s.repo.GetByPageID(ctx, pageID)
}

func (s *service) GetBookEventsByParagraphID(ctx context.Context, paragraphID string) ([]entity.BookEvent, error) {
	return s.repo.GetByParagraphID(ctx, paragraphID)
}
