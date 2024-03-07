package services

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/converter"
)

type BookEventsService interface {
	CreateBookEvent(ctx context.Context, request models.CreateBookEventRequest) (*models.BookEvent, error)
	GetBookEventByID(ctx context.Context, id string) (*models.BookEvent, error)
	UpdateBookEvent(ctx context.Context, id string, request models.UpdateBookEventRequest) (*models.BookEvent, error)
	DeleteBookEvent(ctx context.Context, id string) (*models.BookEvent, error)

	GetBookEventsByBookID(ctx context.Context, bookID string) ([]*models.BookEvent, error)
	GetBookEventsByChapterID(ctx context.Context, chapterID string) ([]*models.BookEvent, error)
	GetBookEventsByPageID(ctx context.Context, pageID string) ([]*models.BookEvent, error)
	GetBookEventsByParagraphID(ctx context.Context, paragraphID string) ([]*models.BookEvent, error)
}

type bookEventsService struct {
	repo SimpleRepo[*entity.BookEvent]
}

func NewBookEventsService(repo SimpleRepo[*entity.BookEvent]) BookEventsService {
	return &bookEventsService{repo}
}

func (s *bookEventsService) CreateBookEvent(ctx context.Context, request models.CreateBookEventRequest) (*models.BookEvent, error) {
	bookEvent := converter.BookEventModelToEntity(request.BookEvent)

	id, err := s.repo.Create(ctx, bookEvent)
	if err != nil {
		return nil, err
	}

	var newBook *models.BookEvent
	newBook, err = s.GetBookEventByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return newBook, err
}

func (s *bookEventsService) GetBookEventByID(ctx context.Context, id string) (*models.BookEvent, error) {
	bookEvent, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.BookEventEntityToModel(bookEvent), nil
}

func (s *bookEventsService) UpdateBookEvent(ctx context.Context, id string, request models.UpdateBookEventRequest) (*models.BookEvent, error) {
	bookEvent := converter.BookEventModelToEntity(request.BookEvent)

	updatedEvent, err := s.repo.Update(ctx, id, bookEvent)
	if err != nil {
		return nil, err
	}

	return converter.BookEventEntityToModel(updatedEvent), nil
}

func (s *bookEventsService) DeleteBookEvent(ctx context.Context, id string) (*models.BookEvent, error) {
	bookEvent, err := s.GetBookEventByID(ctx, id)
	if err != nil {
		return nil, errors.Join(errors.New("BookEvent not found"), err)
	}

	err = s.repo.HardDelete(ctx, id)
	if err != nil {
		return nil, err
	}

	return bookEvent, err
}

func (s *bookEventsService) GetBookEventsByBookID(ctx context.Context, bookID string) ([]*models.BookEvent, error) {
	events, err := s.repo.FindByKV(ctx, "book_id", bookID)
	if err != nil {
		return nil, err
	}
	return converter.BookEventsEntityToModel(events), nil
}

func (s *bookEventsService) GetBookEventsByChapterID(ctx context.Context, chapterID string) ([]*models.BookEvent, error) {
	events, err := s.repo.FindByKV(ctx, "chapter_id", chapterID)
	if err != nil {
		return nil, err
	}
	return converter.BookEventsEntityToModel(events), nil
}

func (s *bookEventsService) GetBookEventsByPageID(ctx context.Context, pageID string) ([]*models.BookEvent, error) {
	events, err := s.repo.FindByKV(ctx, "page_id", pageID)
	if err != nil {
		return nil, err
	}
	return converter.BookEventsEntityToModel(events), nil
}

func (s *bookEventsService) GetBookEventsByParagraphID(ctx context.Context, paragraphID string) ([]*models.BookEvent, error) {
	events, err := s.repo.FindByKV(ctx, "paragraph_id", paragraphID)
	if err != nil {
		return nil, err
	}
	return converter.BookEventsEntityToModel(events), nil
}
