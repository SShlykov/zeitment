package services

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/adapters"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/dbutils"
)

// BookService описывает сервис для работы с книгами.
//
//go:generate mockgen -destination=../../../tests/mocks/domain/services/book_service_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/services BookService
type BookService interface {
	CreateBook(ctx context.Context, request models.CreateBookRequest) (*models.Book, error)
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	UpdateBook(ctx context.Context, id string, request models.UpdateBookRequest) (*models.Book, error)
	DeleteBook(ctx context.Context, id string) (*models.Book, error)
	ListBooks(ctx context.Context, request models.RequestBook) ([]*models.Book, error)
}

type bookService struct {
	repo SimpleRepo[*entity.Book]
}

// NewBookService создает новый экземпляр Service.
func NewBookService(repo SimpleRepo[*entity.Book]) BookService {
	return &bookService{repo}
}

func (s *bookService) CreateBook(ctx context.Context, request models.CreateBookRequest) (*models.Book, error) {
	book := adapters.BookModelToEntity(request.Book)

	if book.Variables == nil {
		book.Variables = []string{}
	}
	id, err := s.repo.Create(ctx, book)
	if err != nil {
		return nil, err
	}

	return s.GetBookByID(ctx, id)
}

func (s *bookService) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	book, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return adapters.BookEntityToModel(book), nil
}

func (s *bookService) UpdateBook(ctx context.Context, id string, request models.UpdateBookRequest) (*models.Book, error) {
	book, err := s.repo.Update(ctx, id, adapters.BookModelToEntity(request.Book))
	if err != nil {
		return nil, err
	}

	return adapters.BookEntityToModel(book), nil
}

func (s *bookService) DeleteBook(ctx context.Context, id string) (*models.Book, error) {
	book, err := s.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.repo.HardDelete(ctx, id)
	if err != nil {
		return nil, err
	}

	return book, err
}

func (s *bookService) ListBooks(ctx context.Context, request models.RequestBook) ([]*models.Book, error) {
	options := dbutils.NewPagination(&request.Options)

	books, err := s.repo.List(ctx, options)
	if err != nil {
		return nil, err
	}

	return adapters.BooksEntityToModel(books), nil
}
