package services

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/domain/adapters"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/postgres/dbutils"
)

//go:generate mockgen -destination=../../../tests/mocks/domain/services/bookevents_service_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/services BookEventsService
type BookEventsService interface {
	CreateBookEvent(ctx context.Context, request models.CreateBookEventRequest) (*models.BookEvent, error)
	GetBookEventByID(ctx context.Context, id string) (*models.BookEvent, error)
	UpdateBookEvent(ctx context.Context, id string, request models.UpdateBookEventRequest) (*models.BookEvent, error)
	DeleteBookEvent(ctx context.Context, id string) (*models.BookEvent, error)

	GetBookEventsByBookID(ctx context.Context, bookID string, request models.RequestBookEvent) ([]*models.BookEvent, error)
	GetBookEventsByChapterID(ctx context.Context, chapterID string, request models.RequestBookEvent) ([]*models.BookEvent, error)
	GetBookEventsByPageID(ctx context.Context, pageID string, request models.RequestBookEvent) ([]*models.BookEvent, error)
	GetBookEventsByParagraphID(ctx context.Context, paragraphID string, request models.RequestBookEvent) ([]*models.BookEvent, error)

	TogglePublic(ctx context.Context, request models.ToggleBookEventRequest) (*models.BookEvent, error)
}

type bookEventsService struct {
	repo SimpleRepo[*entity.BookEvent]
}

func NewBookEventsService(repo SimpleRepo[*entity.BookEvent]) BookEventsService {
	return &bookEventsService{repo}
}

func (s *bookEventsService) TogglePublic(ctx context.Context, request models.ToggleBookEventRequest) (*models.BookEvent, error) {
	bookEvent, err := s.repo.FindByID(ctx, request.BookEventID)
	if err != nil {
		return nil, err
	}
	bookEvent.IsPublic = !bookEvent.IsPublic

	var updated *entity.BookEvent
	updated, err = s.repo.Update(ctx, request.BookEventID, bookEvent)
	if err != nil {
		return nil, err
	}

	return adapters.BookEventEntityToModel(updated), nil
}

func (s *bookEventsService) CreateBookEvent(ctx context.Context, request models.CreateBookEventRequest) (*models.BookEvent, error) {
	bookEvent := adapters.BookEventModelToEntity(request.BookEvent)

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

	return adapters.BookEventEntityToModel(bookEvent), nil
}

func (s *bookEventsService) UpdateBookEvent(ctx context.Context, id string,
	request models.UpdateBookEventRequest) (*models.BookEvent, error) {
	bookEvent := adapters.BookEventModelToEntity(request.BookEvent)

	updatedEvent, err := s.repo.Update(ctx, id, bookEvent)
	if err != nil {
		return nil, err
	}

	return adapters.BookEventEntityToModel(updatedEvent), nil
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

func (s *bookEventsService) GetBookEventsByBookID(ctx context.Context, bookID string,
	request models.RequestBookEvent) ([]*models.BookEvent, error) {
	options := dbutils.NewQueryOptions(
		dbutils.NewFilter("book_id", bookID),
		dbutils.NewPagination(&request.Options),
	)
	events, err := s.repo.FindByKV(ctx, options)
	if err != nil {
		return nil, err
	}
	return adapters.BookEventsEntityToModel(events), nil
}

func (s *bookEventsService) GetBookEventsByChapterID(ctx context.Context, chapterID string,
	request models.RequestBookEvent) ([]*models.BookEvent, error) {
	options := dbutils.NewQueryOptions(
		dbutils.NewFilter("chapter_id", chapterID),
		dbutils.NewPagination(&request.Options),
	)
	events, err := s.repo.FindByKV(ctx, options)
	if err != nil {
		return nil, err
	}
	return adapters.BookEventsEntityToModel(events), nil
}

func (s *bookEventsService) GetBookEventsByPageID(ctx context.Context, pageID string,
	request models.RequestBookEvent) ([]*models.BookEvent, error) {
	options := dbutils.NewQueryOptions(
		dbutils.NewFilter("page_id", pageID),
		dbutils.NewPagination(&request.Options),
	)
	events, err := s.repo.FindByKV(ctx, options)
	if err != nil {
		return nil, err
	}
	return adapters.BookEventsEntityToModel(events), nil
}

func (s *bookEventsService) GetBookEventsByParagraphID(ctx context.Context, paragraphID string,
	request models.RequestBookEvent) ([]*models.BookEvent, error) {
	options := dbutils.NewQueryOptions(
		dbutils.NewFilter("paragraph_id", paragraphID),
		dbutils.NewPagination(&request.Options),
	)
	events, err := s.repo.FindByKV(ctx, options)
	if err != nil {
		return nil, err
	}
	return adapters.BookEventsEntityToModel(events), nil
}
