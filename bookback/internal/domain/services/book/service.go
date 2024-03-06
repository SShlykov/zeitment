package book

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/converter"
)

// Service описывает сервис для работы с книгами.
type Service interface {
	CreateBook(ctx context.Context, request models.CreateBookRequest) (*models.Book, error)
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	UpdateBook(ctx context.Context, id string, request models.UpdateBookRequest) (*models.Book, error)
	DeleteBook(ctx context.Context, id string) (*models.Book, error)
	ListBooks(ctx context.Context, limit uint64, offset uint64) ([]*models.Book, error)
}

type service struct {
	repo Repo
}

// NewService создает новый экземпляр Service.
func NewService(repo Repo) Service {
	return &service{repo}
}

func (s *service) CreateBook(ctx context.Context, request models.CreateBookRequest) (*models.Book, error) {
	book := converter.BookModelToEntity(request.Book)

	if book.Variables == nil {
		book.Variables = []string{}
	}
	id, err := s.repo.Create(ctx, book)
	if err != nil {
		return nil, err
	}

	var newBook *entity.Book
	newBook, err = s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.BookEntityToModel(newBook), err
}

func (s *service) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	book, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.BookEntityToModel(book), nil
}

func (s *service) UpdateBook(ctx context.Context, id string, request models.UpdateBookRequest) (*models.Book, error) {
	book, err := s.repo.Update(ctx, id, converter.BookModelToEntity(request.Book))
	if err != nil {
		return nil, err
	}

	return converter.BookEntityToModel(book), nil
}

func (s *service) DeleteBook(ctx context.Context, id string) (*models.Book, error) {
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

func (s *service) ListBooks(ctx context.Context, limit uint64, offset uint64) ([]*models.Book, error) {
	books, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return converter.BooksEntityToModel(books), nil
}
